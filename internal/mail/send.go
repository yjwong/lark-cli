package mail

import (
	"crypto/rand"
	"crypto/tls"
	"encoding/base64"
	"fmt"
	"mime"
	"net/smtp"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// SendOptions configures the email to send
type SendOptions struct {
	From            string
	To              []string
	CC              []string
	BCC             []string
	Subject         string
	Body            string
	InReplyTo       string   // Message-ID to reply to
	References      []string // Chain of Message-IDs for threading
	Attachments     []string // File paths to attach
	LarkThreadPrefix string  // Lark thread ID prefix for generating threaded Message-IDs
}

// SendResult contains the result of sending an email
type SendResult struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	From    string `json:"from"`
	To      string `json:"to"`
}

// Send sends an email via SMTP using stored credentials
func Send(opts *SendOptions) (*SendResult, error) {
	creds, err := LoadCredentials()
	if err != nil {
		return nil, err
	}

	smtpHost := creds.SMTPHost
	smtpPort := creds.SMTPPort
	if smtpHost == "" {
		// Default: derive SMTP host from IMAP host
		smtpHost = strings.Replace(creds.Host, "imap.", "smtp.", 1)
	}
	if smtpPort == 0 {
		smtpPort = 465
	}

	from := opts.From
	if from == "" {
		from = creds.Username
	}

	// Build all recipients
	var allRecipients []string
	allRecipients = append(allRecipients, opts.To...)
	allRecipients = append(allRecipients, opts.CC...)
	allRecipients = append(allRecipients, opts.BCC...)

	if len(allRecipients) == 0 {
		return nil, fmt.Errorf("no recipients specified")
	}

	// Build the MIME message
	msg, err := buildMessage(from, opts)
	if err != nil {
		return nil, fmt.Errorf("building message: %w", err)
	}

	// Connect via TLS (port 465)
	addr := fmt.Sprintf("%s:%d", smtpHost, smtpPort)
	tlsConfig := &tls.Config{ServerName: smtpHost}

	conn, err := tls.Dial("tcp", addr, tlsConfig)
	if err != nil {
		return nil, fmt.Errorf("connecting to SMTP server %s: %w", addr, err)
	}

	client, err := smtp.NewClient(conn, smtpHost)
	if err != nil {
		return nil, fmt.Errorf("creating SMTP client: %w", err)
	}
	defer client.Close()

	// Authenticate
	auth := smtp.PlainAuth("", creds.Username, creds.Password, smtpHost)
	if err := client.Auth(auth); err != nil {
		return nil, fmt.Errorf("SMTP authentication failed: %w", err)
	}

	// Set sender
	if err := client.Mail(from); err != nil {
		return nil, fmt.Errorf("setting sender: %w", err)
	}

	// Set recipients
	for _, rcpt := range allRecipients {
		if err := client.Rcpt(rcpt); err != nil {
			return nil, fmt.Errorf("setting recipient %s: %w", rcpt, err)
		}
	}

	// Write message body
	w, err := client.Data()
	if err != nil {
		return nil, fmt.Errorf("opening data writer: %w", err)
	}

	if _, err := w.Write(msg); err != nil {
		return nil, fmt.Errorf("writing message: %w", err)
	}

	if err := w.Close(); err != nil {
		return nil, fmt.Errorf("closing data writer: %w", err)
	}

	client.Quit()

	return &SendResult{
		Success: true,
		Message: "email sent successfully",
		From:    from,
		To:      strings.Join(opts.To, ", "),
	}, nil
}

func buildMessage(from string, opts *SendOptions) ([]byte, error) {
	boundary := "==LARK_CLI_BOUNDARY_" + fmt.Sprintf("%d", os.Getpid()) + "=="

	var b strings.Builder

	// Generate Message-ID
	randBytes := make([]byte, 16)
	rand.Read(randBytes)
	var messageID string
	if opts.LarkThreadPrefix != "" {
		// Use Lark's thread prefix so the message is associated with the thread in Lark UI
		uuidBytes := make([]byte, 16)
		rand.Read(uuidBytes)
		uuid := fmt.Sprintf("%x.%x.%x.%x.%x",
			uuidBytes[0:4], uuidBytes[4:6], uuidBytes[6:8], uuidBytes[8:10], uuidBytes[10:16])
		messageID = fmt.Sprintf("<%s.%s@larksuite.com>", opts.LarkThreadPrefix, uuid)
	} else {
		domain := "larksuite.com"
		if parts := strings.SplitN(from, "@", 2); len(parts) == 2 {
			domain = parts[1]
		}
		messageID = fmt.Sprintf("<%x.%d@%s>", randBytes, time.Now().UnixNano(), domain)
	}

	// Headers
	b.WriteString(fmt.Sprintf("From: %s\r\n", from))
	b.WriteString(fmt.Sprintf("To: %s\r\n", strings.Join(opts.To, ", ")))
	if len(opts.CC) > 0 {
		b.WriteString(fmt.Sprintf("Cc: %s\r\n", strings.Join(opts.CC, ", ")))
	}
	b.WriteString(fmt.Sprintf("Subject: %s\r\n", opts.Subject))
	b.WriteString(fmt.Sprintf("Date: %s\r\n", time.Now().Format(time.RFC1123Z)))
	b.WriteString(fmt.Sprintf("Message-ID: %s\r\n", messageID))

	if opts.InReplyTo != "" {
		b.WriteString(fmt.Sprintf("In-Reply-To: %s\r\n", opts.InReplyTo))
	}
	if len(opts.References) > 0 {
		b.WriteString(fmt.Sprintf("References: %s\r\n", strings.Join(opts.References, " ")))
	}

	b.WriteString("MIME-Version: 1.0\r\n")

	hasAttachments := len(opts.Attachments) > 0

	if hasAttachments {
		b.WriteString(fmt.Sprintf("Content-Type: multipart/mixed; boundary=\"%s\"\r\n", boundary))
		b.WriteString("\r\n")

		// Body part
		b.WriteString(fmt.Sprintf("--%s\r\n", boundary))
		b.WriteString("Content-Type: text/plain; charset=\"utf-8\"\r\n")
		b.WriteString("Content-Transfer-Encoding: 7bit\r\n")
		b.WriteString("\r\n")
		b.WriteString(opts.Body)
		b.WriteString("\r\n")

		// Attachment parts
		for _, path := range opts.Attachments {
			data, err := os.ReadFile(path)
			if err != nil {
				return nil, fmt.Errorf("reading attachment %s: %w", path, err)
			}

			filename := filepath.Base(path)
			mimeType := mime.TypeByExtension(filepath.Ext(path))
			if mimeType == "" {
				mimeType = "application/octet-stream"
			}

			b.WriteString(fmt.Sprintf("--%s\r\n", boundary))
			b.WriteString(fmt.Sprintf("Content-Type: %s; name=\"%s\"\r\n", mimeType, filename))
			b.WriteString("Content-Transfer-Encoding: base64\r\n")
			b.WriteString(fmt.Sprintf("Content-Disposition: attachment; filename=\"%s\"\r\n", filename))
			b.WriteString("\r\n")

			encoded := base64.StdEncoding.EncodeToString(data)
			// Split into 76-char lines per RFC 2045
			for i := 0; i < len(encoded); i += 76 {
				end := i + 76
				if end > len(encoded) {
					end = len(encoded)
				}
				b.WriteString(encoded[i:end])
				b.WriteString("\r\n")
			}
		}

		b.WriteString(fmt.Sprintf("--%s--\r\n", boundary))
	} else {
		b.WriteString("Content-Type: text/plain; charset=\"utf-8\"\r\n")
		b.WriteString("\r\n")
		b.WriteString(opts.Body)
		b.WriteString("\r\n")
	}

	return []byte(b.String()), nil
}

// DraftResult contains the result of saving a draft
type DraftResult struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Mailbox string `json:"mailbox"`
	UID     uint32 `json:"uid,omitempty"`
}

// SaveDraft builds a MIME message and saves it as a draft via IMAP APPEND
func SaveDraft(opts *SendOptions) (*DraftResult, error) {
	creds, err := LoadCredentials()
	if err != nil {
		return nil, err
	}

	from := opts.From
	if from == "" {
		from = creds.Username
	}

	msg, err := buildMessage(from, opts)
	if err != nil {
		return nil, fmt.Errorf("building message: %w", err)
	}

	client, err := Connect()
	if err != nil {
		return nil, fmt.Errorf("connecting to IMAP: %w", err)
	}
	defer client.Close()

	// Try common draft folder names
	mailbox := "Drafts"
	uid, err := client.AppendDraft(mailbox, msg)
	if err != nil {
		// Some servers use "Draft" instead of "Drafts"
		uid, err = client.AppendDraft("Draft", msg)
		if err != nil {
			return nil, fmt.Errorf("saving draft: %w", err)
		}
		mailbox = "Draft"
	}

	return &DraftResult{
		Success: true,
		Message: "draft saved successfully",
		Mailbox: mailbox,
		UID:     uint32(uid),
	}, nil
}

// ExtractAndQuoteBody extracts the text/plain body from a raw RFC822 message
// and formats it as a quoted reply block.
func ExtractAndQuoteBody(rawMsg []byte, envelope *Envelope, fromAddr string) string {
	msgStr := string(rawMsg)

	// Extract text/plain body from MIME message
	var textBody string

	// Simple MIME parser: find text/plain section
	if strings.Contains(msgStr, "multipart/") {
		// Find boundary
		boundaryIdx := strings.Index(msgStr, "boundary=")
		if boundaryIdx >= 0 {
			rest := msgStr[boundaryIdx+len("boundary="):]
			// Extract boundary value (may be quoted)
			var boundary string
			if strings.HasPrefix(rest, "\"") {
				endQuote := strings.Index(rest[1:], "\"")
				if endQuote >= 0 {
					boundary = rest[1 : endQuote+1]
				}
			} else {
				end := strings.IndexAny(rest, "\r\n; ")
				if end >= 0 {
					boundary = rest[:end]
				} else {
					boundary = rest
				}
			}

			if boundary != "" {
				parts := strings.Split(msgStr, "--"+boundary)
				for _, part := range parts {
					if strings.Contains(part, "text/plain") {
						// Find the body after the empty line
						headerEnd := strings.Index(part, "\r\n\r\n")
						if headerEnd < 0 {
							headerEnd = strings.Index(part, "\n\n")
						}
						if headerEnd >= 0 {
							textBody = strings.TrimSpace(part[headerEnd:])
						}
						break
					}
				}
			}
		}
	} else {
		// Single-part message
		headerEnd := strings.Index(msgStr, "\r\n\r\n")
		if headerEnd < 0 {
			headerEnd = strings.Index(msgStr, "\n\n")
		}
		if headerEnd >= 0 {
			textBody = strings.TrimSpace(msgStr[headerEnd:])
		}
	}

	if textBody == "" {
		return ""
	}

	// Build the quote header
	var header string
	if envelope != nil {
		dateStr := ""
		if envelope.Date != 0 {
			dateStr = time.Unix(envelope.Date, 0).Format("Mon, Jan 2, 2006, 3:04 PM")
		}
		fromDisplay := fromAddr
		if envelope.FromName != "" {
			fromDisplay = fmt.Sprintf("\"%s\" <%s>", envelope.FromName, fromAddr)
		}
		header = fmt.Sprintf("On %s, %s wrote:", dateStr, fromDisplay)
	} else {
		header = fmt.Sprintf("On an earlier date, %s wrote:", fromAddr)
	}

	// Quote each line
	var quoted strings.Builder
	quoted.WriteString(header)
	quoted.WriteString("\n")
	for _, line := range strings.Split(textBody, "\n") {
		line = strings.TrimRight(line, "\r")
		quoted.WriteString("> ")
		quoted.WriteString(line)
		quoted.WriteString("\n")
	}

	return quoted.String()
}

// ExtractLarkThreadPrefix finds the Lark thread prefix from a list of Message-IDs.
// Lark Message-IDs follow the pattern: <threadprefix.uuid@larksuite.com>
// The thread prefix is shared across all messages in the same thread.
func ExtractLarkThreadPrefix(messageIDs []string) string {
	for _, id := range messageIDs {
		// Strip angle brackets
		id = strings.TrimPrefix(id, "<")
		id = strings.TrimSuffix(id, ">")

		if !strings.HasSuffix(id, "@larksuite.com") {
			continue
		}

		// Remove @larksuite.com
		local := strings.TrimSuffix(id, "@larksuite.com")

		// The thread prefix is the part before the first UUID-like segment
		// Format: threadprefix.xxxxxxxx.xxxx.xxxx.xxxx.xxxxxxxxxxxx
		// Thread prefix itself is a 40-char hex string (SHA1)
		dotIdx := strings.Index(local, ".")
		if dotIdx > 0 {
			prefix := local[:dotIdx]
			// Validate it looks like a hex string (at least 32 chars)
			if len(prefix) >= 32 {
				return prefix
			}
		}
	}
	return ""
}

// TestSMTPConnection tests SMTP connectivity with given credentials
func TestSMTPConnection(creds *Credentials) error {
	smtpHost := creds.SMTPHost
	smtpPort := creds.SMTPPort
	if smtpHost == "" {
		smtpHost = strings.Replace(creds.Host, "imap.", "smtp.", 1)
	}
	if smtpPort == 0 {
		smtpPort = 465
	}

	addr := fmt.Sprintf("%s:%d", smtpHost, smtpPort)
	tlsConfig := &tls.Config{ServerName: smtpHost}

	conn, err := tls.Dial("tcp", addr, tlsConfig)
	if err != nil {
		return fmt.Errorf("connecting to %s: %w", addr, err)
	}

	client, err := smtp.NewClient(conn, smtpHost)
	if err != nil {
		conn.Close()
		return fmt.Errorf("SMTP handshake: %w", err)
	}
	defer client.Close()

	auth := smtp.PlainAuth("", creds.Username, creds.Password, smtpHost)
	if err := client.Auth(auth); err != nil {
		return fmt.Errorf("SMTP auth: %w", err)
	}

	client.Quit()
	return nil
}
