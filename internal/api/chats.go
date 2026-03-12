package api

import (
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

	if err := resp.Err(); err != nil {
		return nil, false, "", err
	}

	return resp.Data.Items, resp.Data.HasMore, resp.Data.PageToken, nil
}
