package mail

import (
	"crypto/tls"
	"fmt"

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
