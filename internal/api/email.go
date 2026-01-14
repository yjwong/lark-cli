package api

import (
	"encoding/base64"
	"fmt"
	"net/url"
	"strings"
)

// ListEmailsOptions contains optional parameters for ListEmails
type ListEmailsOptions struct {
	FolderID   string // Folder ID (default: "INBOX")
	OnlyUnread bool   // Only query unread emails
	PageSize   int    // 1-20, default 20
	PageToken  string // Pagination token
}

// ListEmails retrieves email message IDs from a mailbox folder
// mailboxID: user email address or "me" for current user
// Returns list of message IDs, hasMore flag, next page token, and any error
func (c *Client) ListEmails(mailboxID string, opts *ListEmailsOptions) ([]string, bool, string, error) {
	if mailboxID == "" {
		mailboxID = "me"
	}

	pageSize := 20
	if opts != nil && opts.PageSize > 0 {
		pageSize = opts.PageSize
		if pageSize > 20 {
			pageSize = 20
		}
	}

	folderID := "INBOX"
	if opts != nil && opts.FolderID != "" {
		folderID = opts.FolderID
	}

	// Build query parameters
	params := url.Values{}
	params.Set("page_size", fmt.Sprintf("%d", pageSize))
	params.Set("folder_id", folderID)

	if opts != nil {
		if opts.OnlyUnread {
			params.Set("only_unread", "true")
		}
		if opts.PageToken != "" {
			params.Set("page_token", opts.PageToken)
		}
	}

	path := fmt.Sprintf("/mail/v1/user_mailboxes/%s/messages?%s", mailboxID, params.Encode())

	var resp ListEmailsResponse
	if err := c.Get(path, &resp); err != nil {
		return nil, false, "", err
	}

	if resp.Code != 0 {
		return nil, false, "", fmt.Errorf("API error %d: %s", resp.Code, resp.Msg)
	}

	return resp.Data.Items, resp.Data.HasMore, resp.Data.PageToken, nil
}

// GetEmail retrieves the details of a specific email message
// mailboxID: user email address or "me" for current user
// messageID: the email message ID
func (c *Client) GetEmail(mailboxID, messageID string) (*EmailMessage, error) {
	if mailboxID == "" {
		mailboxID = "me"
	}

	path := fmt.Sprintf("/mail/v1/user_mailboxes/%s/messages/%s", mailboxID, messageID)

	var resp GetEmailResponse
	if err := c.Get(path, &resp); err != nil {
		return nil, err
	}

	if resp.Code != 0 {
		return nil, fmt.Errorf("API error %d: %s", resp.Code, resp.Msg)
	}

	return resp.Data.Message, nil
}

// GetAttachmentDownloadURLs retrieves download URLs for email attachments
// mailboxID: user email address or "me" for current user
// messageID: the email message ID
// attachmentIDs: list of attachment IDs to get download URLs for
func (c *Client) GetAttachmentDownloadURLs(mailboxID, messageID string, attachmentIDs []string) ([]AttachmentDownloadURL, []string, error) {
	if mailboxID == "" {
		mailboxID = "me"
	}

	if len(attachmentIDs) == 0 {
		return nil, nil, fmt.Errorf("at least one attachment ID is required")
	}

	if len(attachmentIDs) > 20 {
		return nil, nil, fmt.Errorf("maximum 20 attachment IDs allowed")
	}

	// Build query parameters - attachment_ids is a repeated parameter
	params := url.Values{}
	for _, id := range attachmentIDs {
		params.Add("attachment_ids", id)
	}

	path := fmt.Sprintf("/mail/v1/user_mailboxes/%s/messages/%s/attachments/download_url?%s",
		mailboxID, messageID, params.Encode())

	var resp GetAttachmentDownloadURLsResponse
	if err := c.Get(path, &resp); err != nil {
		return nil, nil, err
	}

	if resp.Code != 0 {
		return nil, nil, fmt.Errorf("API error %d: %s", resp.Code, resp.Msg)
	}

	return resp.Data.DownloadURLs, resp.Data.FailedIDs, nil
}

// GetAllAttachmentDownloadURLs retrieves download URLs for all attachments in an email
// This is a convenience method that first fetches the email to get attachment IDs
func (c *Client) GetAllAttachmentDownloadURLs(mailboxID, messageID string) ([]AttachmentDownloadURL, []string, []EmailAttachment, error) {
	// First get the email to find attachment IDs
	email, err := c.GetEmail(mailboxID, messageID)
	if err != nil {
		return nil, nil, nil, err
	}

	if email == nil || len(email.Attachments) == 0 {
		return nil, nil, nil, nil
	}

	// Collect attachment IDs
	attachmentIDs := make([]string, 0, len(email.Attachments))
	for _, att := range email.Attachments {
		if att.ID != "" {
			attachmentIDs = append(attachmentIDs, att.ID)
		}
	}

	if len(attachmentIDs) == 0 {
		return nil, nil, email.Attachments, nil
	}

	// Get download URLs
	downloadURLs, failedIDs, err := c.GetAttachmentDownloadURLs(mailboxID, messageID, attachmentIDs)
	if err != nil {
		return nil, nil, email.Attachments, err
	}

	return downloadURLs, failedIDs, email.Attachments, nil
}

// DecodeEmailBody decodes a base64url encoded email body
func DecodeEmailBody(encoded string) (string, error) {
	if encoded == "" {
		return "", nil
	}

	// Base64url uses - and _ instead of + and /
	// Standard base64 uses + and /
	s := strings.ReplaceAll(encoded, "-", "+")
	s = strings.ReplaceAll(s, "_", "/")

	// Add padding if needed
	switch len(s) % 4 {
	case 2:
		s += "=="
	case 3:
		s += "="
	}

	decoded, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return "", fmt.Errorf("failed to decode email body: %w", err)
	}

	return string(decoded), nil
}
