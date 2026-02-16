package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"

	"github.com/yjwong/lark-cli/internal/auth"
)

// ListMessagesOptions contains optional parameters for ListMessages
type ListMessagesOptions struct {
	StartTime string // Unix timestamp in seconds
	EndTime   string // Unix timestamp in seconds
	SortType  string // ByCreateTimeAsc or ByCreateTimeDesc
	PageSize  int    // 1-50, default 20
	PageToken string // Pagination token
}

// ListMessageReactionsOptions contains optional parameters for ListMessageReactions
type ListMessageReactionsOptions struct {
	ReactionType string // Emoji type (e.g., SMILE)
	PageSize     int    // 1-50, default 20
	PageToken    string // Pagination token
	UserIDType   string // open_id, union_id, user_id
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

// ListMessageReactions retrieves reactions for a message
func (c *Client) ListMessageReactions(messageID string, opts *ListMessageReactionsOptions) ([]MessageReaction, bool, string, error) {
	pageSize := 20
	if opts != nil && opts.PageSize > 0 {
		pageSize = opts.PageSize
		if pageSize > 50 {
			pageSize = 50
		}
	}

	params := url.Values{}
	params.Set("page_size", fmt.Sprintf("%d", pageSize))
	if opts != nil {
		if opts.ReactionType != "" {
			params.Set("reaction_type", opts.ReactionType)
		}
		if opts.PageToken != "" {
			params.Set("page_token", opts.PageToken)
		}
		if opts.UserIDType != "" {
			params.Set("user_id_type", opts.UserIDType)
		}
	}

	path := fmt.Sprintf("/im/v1/messages/%s/reactions?%s", messageID, params.Encode())

	var resp MessageReactionListResponse
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

// UploadMessageImage uploads an image for message sending and returns the image key
func (c *Client) UploadMessageImage(filePath string) (string, error) {
	if err := auth.EnsureValidTenantToken(); err != nil {
		return "", err
	}

	file, err := os.Open(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to open image: %w", err)
	}
	defer file.Close()

	var buf bytes.Buffer
	writer := multipart.NewWriter(&buf)
	if err := writer.WriteField("image_type", "message"); err != nil {
		return "", fmt.Errorf("failed to write image_type: %w", err)
	}

	part, err := writer.CreateFormFile("image", filepath.Base(filePath))
	if err != nil {
		return "", fmt.Errorf("failed to create image form: %w", err)
	}
	if _, err := io.Copy(part, file); err != nil {
		return "", fmt.Errorf("failed to read image: %w", err)
	}
	if err := writer.Close(); err != nil {
		return "", fmt.Errorf("failed to finalize upload: %w", err)
	}

	url := getBaseURL() + "/im/v1/images"
	req, err := http.NewRequest("POST", url, &buf)
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	token := auth.GetTenantTokenStore().GetAccessToken()
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response: %w", err)
	}

	var uploadResp UploadImageResponse
	if err := json.Unmarshal(respBody, &uploadResp); err != nil {
		return "", fmt.Errorf("failed to parse response: %w", err)
	}

	if uploadResp.Code != 0 {
		return "", fmt.Errorf("API error %d: %s", uploadResp.Code, uploadResp.Msg)
	}

	if uploadResp.Data.ImageKey == "" {
		return "", fmt.Errorf("API error: missing image_key")
	}

	return uploadResp.Data.ImageKey, nil
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

// ReplyMessage replies to a message by message_id
// msgType: "text" or "post"
// content: JSON string of message content (format depends on msgType)
// rootID: optional root message ID for thread replies
// replyInThread: whether to reply in thread
func (c *Client) ReplyMessage(messageID, msgType, content, rootID string, replyInThread bool) (*SendMessageResponse, error) {
	path := fmt.Sprintf("/im/v1/messages/%s/reply", messageID)

	req := ReplyMessageRequest{
		MsgType:       msgType,
		Content:       content,
		RootID:        rootID,
		ReplyInThread: replyInThread,
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

// RecallMessage recalls/deletes a message
// messageID: the ID of the message to recall
func (c *Client) RecallMessage(messageID string) error {
	path := fmt.Sprintf("/im/v1/messages/%s", messageID)

	var resp BaseResponse
	if err := c.DeleteWithTenantToken(path, &resp); err != nil {
		return err
	}

	if resp.Code != 0 {
		return fmt.Errorf("API error %d: %s", resp.Code, resp.Msg)
	}

	return nil
}

// DeleteMessageReaction removes a reaction from a message
func (c *Client) DeleteMessageReaction(messageID, reactionID string) (*MessageReaction, error) {
	path := fmt.Sprintf("/im/v1/messages/%s/reactions/%s", messageID, reactionID)

	var resp DeleteMessageReactionResponse
	if err := c.DeleteWithTenantToken(path, &resp); err != nil {
		return nil, err
	}

	if resp.Code != 0 {
		return nil, fmt.Errorf("API error %d: %s", resp.Code, resp.Msg)
	}

	return resp.Data, nil
}

// AddMessageReaction adds an emoji reaction to a message
// messageID: the ID of the message to react to
// emojiType: emoji type key (e.g., "SMILE")
func (c *Client) AddMessageReaction(messageID, emojiType string) (*MessageReaction, error) {
	path := fmt.Sprintf("/im/v1/messages/%s/reactions", messageID)
	req := AddMessageReactionRequest{
		ReactionType: ReactionType{
			EmojiType: emojiType,
		},
	}

	var resp AddMessageReactionResponse
	if err := c.PostWithTenantToken(path, req, &resp); err != nil {
		return nil, err
	}

	if resp.Code != 0 {
		return nil, fmt.Errorf("API error %d: %s", resp.Code, resp.Msg)
	}

	return resp.Data, nil
}
