package api

import (
	"fmt"
	"net/url"
	"strconv"
)

// SearchChatsOptions configures the search chats request
type SearchChatsOptions struct {
	Query      string
	UserIDType string // open_id, union_id, user_id
	PageSize   int
	PageToken  string
}

// SearchChats searches for chats/groups visible to the user or bot
func (c *Client) SearchChats(opts *SearchChatsOptions) ([]Chat, bool, string, error) {
	// Build query parameters
	params := url.Values{}

	if opts != nil {
		if opts.Query != "" {
			params.Set("query", opts.Query)
		}
		if opts.UserIDType != "" {
			params.Set("user_id_type", opts.UserIDType)
		}
		if opts.PageSize > 0 {
			params.Set("page_size", strconv.Itoa(opts.PageSize))
		}
		if opts.PageToken != "" {
			params.Set("page_token", opts.PageToken)
		}
	}

	path := "/im/v1/chats/search"
	if len(params) > 0 {
		path += "?" + params.Encode()
	}

	var resp SearchChatsResponse
	if err := c.GetWithTenantToken(path, &resp); err != nil {
		return nil, false, "", err
	}

	if resp.Code != 0 {
		return nil, false, "", fmt.Errorf("API error %d: %s", resp.Code, resp.Msg)
	}

	return resp.Data.Items, resp.Data.HasMore, resp.Data.PageToken, nil
}

// ListChatsOptions configures the list chats request
type ListChatsOptions struct {
	UserIDType string // open_id, union_id, user_id
	PageSize   int
	PageToken  string
}

// ListChatMembersOptions configures the list chat members request.
// ChatID is required.
type ListChatMembersOptions struct {
	ChatID       string
	MemberIDType string // open_id, union_id, user_id
	PageSize     int
	PageToken    string
}

// ListChatMembers lists all members of a chat via GET /im/v1/chats/:chat_id/members
func (c *Client) ListChatMembers(opts *ListChatMembersOptions) ([]ChatMember, bool, string, error) {
	if opts == nil || opts.ChatID == "" {
		return nil, false, "", fmt.Errorf("chat_id is required")
	}

	params := url.Values{}
	if opts.MemberIDType != "" {
		params.Set("member_id_type", opts.MemberIDType)
	}
	if opts.PageSize > 0 {
		params.Set("page_size", strconv.Itoa(opts.PageSize))
	}
	if opts.PageToken != "" {
		params.Set("page_token", opts.PageToken)
	}

	path := fmt.Sprintf("/im/v1/chats/%s/members", url.PathEscape(opts.ChatID))
	if len(params) > 0 {
		path += "?" + params.Encode()
	}

	var resp ListChatMembersResponse
	if err := c.GetWithTenantToken(path, &resp); err != nil {
		return nil, false, "", err
	}

	if resp.Code != 0 {
		return nil, false, "", fmt.Errorf("API error %d: %s", resp.Code, resp.Msg)
	}

	return resp.Data.Items, resp.Data.HasMore, resp.Data.PageToken, nil
}

// ListChats lists all chats the bot has joined via GET /im/v1/chats
func (c *Client) ListChats(opts *ListChatsOptions) ([]Chat, bool, string, error) {
	params := url.Values{}

	if opts != nil {
		if opts.UserIDType != "" {
			params.Set("user_id_type", opts.UserIDType)
		}
		if opts.PageSize > 0 {
			params.Set("page_size", strconv.Itoa(opts.PageSize))
		}
		if opts.PageToken != "" {
			params.Set("page_token", opts.PageToken)
		}
	}

	path := "/im/v1/chats"
	if len(params) > 0 {
		path += "?" + params.Encode()
	}

	var resp ListChatsResponse
	if err := c.GetWithTenantToken(path, &resp); err != nil {
		return nil, false, "", err
	}

	if resp.Code != 0 {
		return nil, false, "", fmt.Errorf("API error %d: %s", resp.Code, resp.Msg)
	}

	return resp.Data.Items, resp.Data.HasMore, resp.Data.PageToken, nil
}
