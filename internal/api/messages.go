package api

import (
	"fmt"
	"io"
	"net/url"
)

// ListMessagesOptions contains optional parameters for ListMessages
type ListMessagesOptions struct {
	StartTime string // Unix timestamp in seconds
	EndTime   string // Unix timestamp in seconds
	SortType  string // ByCreateTimeAsc or ByCreateTimeDesc
	PageSize  int    // 1-50, default 20
	PageToken string // Pagination token
}

// ListMessages retrieves chat history from a chat or thread
// containerIDType: "chat" for groups/private chats, "thread" for thread messages
// containerID: chat_id or thread_id
func (c *Client) ListMessages(containerIDType, containerID string, opts *ListMessagesOptions) ([]Message, bool, string, error) {
	if containerIDType == "" {
		containerIDType = "chat"
	}

	pageSize := 20
	if opts != nil && opts.PageSize > 0 {
		pageSize = opts.PageSize
		if pageSize > 50 {
			pageSize = 50
		}
	}

	// Build query parameters
	params := url.Values{}
	params.Set("container_id_type", containerIDType)
	params.Set("container_id", containerID)
	params.Set("page_size", fmt.Sprintf("%d", pageSize))

	if opts != nil {
		if opts.StartTime != "" {
			params.Set("start_time", opts.StartTime)
		}
		if opts.EndTime != "" {
			params.Set("end_time", opts.EndTime)
		}
		if opts.SortType != "" {
			params.Set("sort_type", opts.SortType)
		}
		if opts.PageToken != "" {
			params.Set("page_token", opts.PageToken)
		}
	}

	path := "/im/v1/messages?" + params.Encode()

	var resp MessageListResponse
	if err := c.GetWithTenantToken(path, &resp); err != nil {
		return nil, false, "", err
	}

	if resp.Code != 0 {
		return nil, false, "", fmt.Errorf("API error %d: %s", resp.Code, resp.Msg)
	}

	return resp.Data.Items, resp.Data.HasMore, resp.Data.PageToken, nil
}

// GetMessageResource downloads a resource file (image, video, audio, file) from a message
// resourceType must be "image" or "file" (file covers files, audio, and video)
// Returns the response body (caller must close), content-type, and any error
func (c *Client) GetMessageResource(messageID, fileKey, resourceType string) (io.ReadCloser, string, error) {
	if resourceType != "image" && resourceType != "file" {
		return nil, "", fmt.Errorf("invalid resource type: %s (must be 'image' or 'file')", resourceType)
	}

	path := fmt.Sprintf("/im/v1/messages/%s/resources/%s?type=%s", messageID, fileKey, resourceType)
	return c.DownloadWithTenantToken(path)
}

// SendMessage sends a message to a user or chat
// receiveIDType: "open_id", "user_id", "email", "chat_id"
// receiveID: the recipient identifier
// msgType: "text" or "post"
// content: JSON string of message content (format depends on msgType)
func (c *Client) SendMessage(receiveIDType, receiveID, msgType, content string) (*SendMessageResponse, error) {
	path := fmt.Sprintf("/im/v1/messages?receive_id_type=%s", receiveIDType)

	req := SendMessageRequest{
		ReceiveID: receiveID,
		MsgType:   msgType,
		Content:   content,
	}

	var resp SendMessageResponse
	if err := c.PostWithTenantToken(path, req, &resp); err != nil {
		return nil, err
	}

	if resp.Code != 0 {
		return nil, fmt.Errorf("API error %d: %s", resp.Code, resp.Msg)
	}

	return &resp, nil
}

// DeleteMessage recalls a message sent by the bot
// messageID: the ID of the message to recall
// Bot can only recall its own messages sent within the configurable time limit (default 24h).
// For group chats, group admins can recall any message within 1 year.
func (c *Client) DeleteMessage(messageID string) error {
	path := fmt.Sprintf("/open-apis/im/v1/messages/%s", messageID)

	var resp BaseResponse
	if err := c.Delete(path, &resp); err != nil {
		return err
	}
	if resp.Code != 0 {
		return fmt.Errorf("API error %d: %s", resp.Code, resp.Msg)
	}
	return nil
}
