package mail

import (
	"crypto/tls"
	"fmt"
	"strings"
	"time"

	"github.com/emersion/go-imap/v2"
	"github.com/emersion/go-imap/v2/imapclient"
)

// UID is an alias for imap.UID for external use
type UID = imap.UID

// Client wraps an IMAP connection
type Client struct {
	imap  *imapclient.Client
	creds *Credentials
}

// Connect establishes an IMAP connection using stored credentials
func Connect() (*Client, error) {
	creds, err := LoadCredentials()
	if err != nil {
		return nil, err
	}

	return ConnectWithCredentials(creds)
}

// ConnectWithCredentials establishes an IMAP connection with explicit credentials
func ConnectWithCredentials(creds *Credentials) (*Client, error) {
	addr := fmt.Sprintf("%s:%d", creds.Host, creds.Port)

	var client *imapclient.Client
	var err error

	if creds.UseSSL {
		client, err = imapclient.DialTLS(addr, &imapclient.Options{
			TLSConfig: &tls.Config{},
		})
	} else {
		client, err = imapclient.DialInsecure(addr, nil)
	}

	if err != nil {
		return nil, fmt.Errorf("connecting to %s: %w", addr, err)
	}

	if err := client.Login(creds.Username, creds.Password).Wait(); err != nil {
		client.Close()
		return nil, fmt.Errorf("login failed: %w", err)
	}

	return &Client{imap: client, creds: creds}, nil
}

// Close closes the IMAP connection
func (c *Client) Close() error {
	if c.imap != nil {
		return c.imap.Close()
	}
	return nil
}

// Mailbox represents a mailbox with metadata
type Mailbox struct {
	Name        string
	NumMessages uint32
	UIDValidity uint32
}

// ListMailboxes returns all mailboxes
func (c *Client) ListMailboxes() ([]string, error) {
	mailboxes, err := c.imap.List("", "*", nil).Collect()
	if err != nil {
		return nil, fmt.Errorf("listing mailboxes: %w", err)
	}

	names := make([]string, len(mailboxes))
	for i, mbox := range mailboxes {
		names[i] = mbox.Mailbox
	}
	return names, nil
}

// SelectMailbox selects a mailbox and returns its metadata
func (c *Client) SelectMailbox(name string) (*Mailbox, error) {
	mbox, err := c.imap.Select(name, nil).Wait()
	if err != nil {
		return nil, fmt.Errorf("selecting mailbox %s: %w", name, err)
	}

	return &Mailbox{
		Name:        name,
		NumMessages: mbox.NumMessages,
		UIDValidity: mbox.UIDValidity,
	}, nil
}

// Envelope represents email metadata
type Envelope struct {
	UID       imap.UID
	MessageID string
	Date      int64 // Unix timestamp
	FromAddr  string
	FromName  string
	Subject   string
}

// FetchEnvelopes fetches envelope data for a range of sequence numbers
func (c *Client) FetchEnvelopes(start, end uint32) ([]Envelope, error) {
	var seqSet imap.SeqSet
	seqSet.AddRange(start, end)

	fetchOptions := &imap.FetchOptions{
		Envelope: true,
		UID:      true,
	}

	messages, err := c.imap.Fetch(seqSet, fetchOptions).Collect()
	if err != nil {
		return nil, fmt.Errorf("fetching envelopes: %w", err)
	}

	envelopes := make([]Envelope, 0, len(messages))
	for _, msg := range messages {
		env := msg.Envelope
		if env == nil {
			continue
		}

		e := Envelope{
			UID:       msg.UID,
			MessageID: env.MessageID,
			Subject:   env.Subject,
		}

		if !env.Date.IsZero() {
			e.Date = env.Date.Unix()
		}

		if len(env.From) > 0 {
			e.FromAddr = env.From[0].Addr()
			e.FromName = env.From[0].Name
		}

		envelopes = append(envelopes, e)
	}

	return envelopes, nil
}

// FetchEnvelopesByUID fetches envelope data for specific UIDs
func (c *Client) FetchEnvelopesByUID(uids []imap.UID) ([]Envelope, error) {
	if len(uids) == 0 {
		return nil, nil
	}

	uidSet := imap.UIDSetNum(uids...)

	fetchOptions := &imap.FetchOptions{
		Envelope: true,
		UID:      true,
	}

	messages, err := c.imap.Fetch(uidSet, fetchOptions).Collect()
	if err != nil {
		return nil, fmt.Errorf("fetching envelopes by UID: %w", err)
	}

	envelopes := make([]Envelope, 0, len(messages))
	for _, msg := range messages {
		env := msg.Envelope
		if env == nil {
			continue
		}

		e := Envelope{
			UID:       msg.UID,
			MessageID: env.MessageID,
			Subject:   env.Subject,
		}

		if !env.Date.IsZero() {
			e.Date = env.Date.Unix()
		}

		if len(env.From) > 0 {
			e.FromAddr = env.From[0].Addr()
			e.FromName = env.From[0].Name
		}

		envelopes = append(envelopes, e)
	}

	return envelopes, nil
}

// GetAllUIDs returns all message UIDs in the selected mailbox
func (c *Client) GetAllUIDs() ([]imap.UID, error) {
	// Search for all messages (UID 1:*)
	var uidSet imap.UIDSet
	uidSet.AddRange(1, 0) // 1 to * (all messages)

	criteria := &imap.SearchCriteria{
		UID: []imap.UIDSet{uidSet},
	}

	searchData, err := c.imap.UIDSearch(criteria, nil).Wait()
	if err != nil {
		return nil, fmt.Errorf("searching for all UIDs: %w", err)
	}

	return searchData.AllUIDs(), nil
}

// FetchNewEnvelopes fetches envelopes for messages with UID > lastUID
func (c *Client) FetchNewEnvelopes(lastUID imap.UID) ([]Envelope, error) {
	// Create a UID range from lastUID+1 to * (all newer messages)
	var uidSet imap.UIDSet
	uidSet.AddRange(lastUID+1, 0) // 0 means * (highest UID)

	// UID SEARCH for UIDs in range
	criteria := &imap.SearchCriteria{
		UID: []imap.UIDSet{uidSet},
	}

	// Use UID search
	searchData, err := c.imap.UIDSearch(criteria, nil).Wait()
	if err != nil {
		return nil, fmt.Errorf("searching for new UIDs: %w", err)
	}

	if len(searchData.AllUIDs()) == 0 {
		return nil, nil
	}

	return c.FetchEnvelopesByUID(searchData.AllUIDs())
}

// FetchMessage fetches the full RFC822 message for a UID
func (c *Client) FetchMessage(uid imap.UID) ([]byte, *Envelope, error) {
	uidSet := imap.UIDSetNum(uid)

	fetchOptions := &imap.FetchOptions{
		Envelope:    true,
		UID:         true,
		BodySection: []*imap.FetchItemBodySection{{}},
	}

	messages, err := c.imap.Fetch(uidSet, fetchOptions).Collect()
	if err != nil {
		return nil, nil, fmt.Errorf("fetching message: %w", err)
	}

	if len(messages) == 0 {
		return nil, nil, fmt.Errorf("message not found: UID %d", uid)
	}

	msg := messages[0]

	var body []byte
	for _, data := range msg.BodySection {
		body = data
		break
	}

	var envelope *Envelope
	if msg.Envelope != nil {
		env := msg.Envelope
		envelope = &Envelope{
			UID:       msg.UID,
			MessageID: env.MessageID,
			Subject:   env.Subject,
		}
		if !env.Date.IsZero() {
			envelope.Date = env.Date.Unix()
		}
		if len(env.From) > 0 {
			envelope.FromAddr = env.From[0].Addr()
			envelope.FromName = env.From[0].Name
		}
	}

	return body, envelope, nil
}

// AppendDraft saves a message to the specified mailbox with the \Draft flag
func (c *Client) AppendDraft(mailbox string, msg []byte) (imap.UID, error) {
	options := &imap.AppendOptions{
		Flags: []imap.Flag{imap.FlagDraft, imap.FlagSeen},
		Time:  time.Now(),
	}

	appendCmd := c.imap.Append(mailbox, int64(len(msg)), options)
	if _, err := appendCmd.Write(msg); err != nil {
		return 0, fmt.Errorf("writing message to %s: %w", mailbox, err)
	}
	if err := appendCmd.Close(); err != nil {
		return 0, fmt.Errorf("closing append to %s: %w", mailbox, err)
	}

	data, err := appendCmd.Wait()
	if err != nil {
		return 0, fmt.Errorf("APPEND to %s failed: %w", mailbox, err)
	}

	return data.UID, nil
}

// DeleteMessage flags a message as \Deleted and expunges it
func (c *Client) DeleteMessage(uid imap.UID) error {
	uidSet := imap.UIDSetNum(uid)

	storeFlags := &imap.StoreFlags{
		Op:    imap.StoreFlagsAdd,
		Flags: []imap.Flag{imap.FlagDeleted},
	}
	if err := c.imap.Store(uidSet, storeFlags, nil).Close(); err != nil {
		return fmt.Errorf("flagging message %d as deleted: %w", uid, err)
	}

	expungeCmd := c.imap.UIDExpunge(uidSet)
	for expungeCmd.Next() != 0 {
		// drain
	}
	if err := expungeCmd.Close(); err != nil {
		return fmt.Errorf("expunging message %d: %w", uid, err)
	}

	return nil
}

// MoveMessage moves a message to another mailbox
func (c *Client) MoveMessage(uid imap.UID, destMailbox string) error {
	uidSet := imap.UIDSetNum(uid)

	_, err := c.imap.Move(uidSet, destMailbox).Wait()
	if err != nil {
		return fmt.Errorf("moving message %d to %s: %w", uid, destMailbox, err)
	}

	return nil
}

// FetchThreadingInfo fetches the Message-ID, References, and In-Reply-To headers for a message
func (c *Client) FetchThreadingInfo(uid imap.UID) (messageID string, references []string, subject string, from string, err error) {
	uidSet := imap.UIDSetNum(uid)

	fetchOptions := &imap.FetchOptions{
		Envelope: true,
		UID:      true,
		BodySection: []*imap.FetchItemBodySection{
			{Specifier: imap.PartSpecifierHeader, HeaderFields: []string{"References", "In-Reply-To"}},
		},
	}

	messages, err := c.imap.Fetch(uidSet, fetchOptions).Collect()
	if err != nil {
		return "", nil, "", "", fmt.Errorf("fetching message: %w", err)
	}

	if len(messages) == 0 {
		return "", nil, "", "", fmt.Errorf("message not found: UID %d", uid)
	}

	msg := messages[0]

	if msg.Envelope != nil {
		messageID = msg.Envelope.MessageID
		subject = msg.Envelope.Subject
		if len(msg.Envelope.From) > 0 {
			from = msg.Envelope.From[0].Addr()
		}
	}

	// Parse References from header section
	for _, data := range msg.BodySection {
		headerStr := string(data)
		for _, line := range strings.Split(headerStr, "\r\n") {
			trimmed := strings.TrimSpace(line)
			lower := strings.ToLower(trimmed)
			if strings.HasPrefix(lower, "references:") {
				refStr := strings.TrimSpace(trimmed[len("references:"):])
				// Split on whitespace to get individual message IDs
				for _, ref := range strings.Fields(refStr) {
					if strings.HasPrefix(ref, "<") {
						references = append(references, ref)
					}
				}
			}
		}
		break
	}

	return messageID, references, subject, from, nil
}

// TestConnection attempts to connect and list mailboxes
func TestConnection(creds *Credentials) error {
	client, err := ConnectWithCredentials(creds)
	if err != nil {
		return err
	}
	defer client.Close()

	_, err = client.ListMailboxes()
	return err
}
