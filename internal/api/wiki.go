package api

import (
	"fmt"
	"net/url"
	"strconv"
)

// GetWikiNode retrieves wiki node information
// nodeToken: the wiki node token from the wiki URL
func (c *Client) GetWikiNode(nodeToken string) (*WikiNode, error) {
	path := fmt.Sprintf("/wiki/v2/spaces/get_node?token=%s",
		url.QueryEscape(nodeToken))

	var resp WikiNodeResponse
	if err := c.Get(path, &resp); err != nil {
		return nil, err
	}

	if resp.Code != 0 {
		return nil, fmt.Errorf("API error %d: %s", resp.Code, resp.Msg)
	}

	return resp.Data.Node, nil
}

// ListWikiSpaces lists wiki spaces with pagination
// pageSize: number of items per page (max 50)
// pageToken: pagination token
func (c *Client) ListWikiSpaces(pageSize int, pageToken string) ([]WikiSpace, bool, string, error) {
	params := url.Values{}
	if pageSize > 0 {
		params.Set("page_size", strconv.Itoa(pageSize))
	}
	if pageToken != "" {
		params.Set("page_token", pageToken)
	}

	path := "/wiki/v2/spaces"
	if len(params) > 0 {
		path += "?" + params.Encode()
	}

	var resp ListWikiSpacesResponse
	if err := c.Get(path, &resp); err != nil {
		return nil, false, "", err
	}

	if resp.Code != 0 {
		return nil, false, "", fmt.Errorf("API error %d: %s", resp.Code, resp.Msg)
	}

	return resp.Data.Items, resp.Data.HasMore, resp.Data.PageToken, nil
}

// ListWikiNodes lists child nodes in a wiki space with pagination
// spaceID: the wiki space ID
// parentNodeToken: optional parent node token (empty means top-level nodes)
// pageSize: number of items per page (max 50)
// pageToken: pagination token
func (c *Client) ListWikiNodes(spaceID, parentNodeToken string, pageSize int, pageToken string) ([]WikiNode, bool, string, error) {
	params := url.Values{}
	if parentNodeToken != "" {
		params.Set("parent_node_token", parentNodeToken)
	}
	if pageSize > 0 {
		params.Set("page_size", strconv.Itoa(pageSize))
	}
	if pageToken != "" {
		params.Set("page_token", pageToken)
	}

	path := fmt.Sprintf("/wiki/v2/spaces/%s/nodes", url.PathEscape(spaceID))
	if len(params) > 0 {
		path += "?" + params.Encode()
	}

	var resp ListWikiChildrenResponse
	if err := c.Get(path, &resp); err != nil {
		return nil, false, "", err
	}

	if resp.Code != 0 {
		return nil, false, "", fmt.Errorf("API error %d: %s", resp.Code, resp.Msg)
	}

	return resp.Data.Items, resp.Data.HasMore, resp.Data.PageToken, nil
}

// SearchWikiNodes searches wiki nodes by keyword
// query: search keyword (required)
// spaceID: optional filter to specific wiki space
// nodeID: optional filter to search within a node (requires spaceID)
// Returns matching wiki nodes (limited to first page of 50 results to avoid rate limits)
func (c *Client) SearchWikiNodes(query, spaceID, nodeID string) ([]WikiSearchItem, error) {
	req := WikiSearchRequest{
		Query:    query,
		PageSize: 50,
	}
	if spaceID != "" {
		req.SpaceID = spaceID
	}
	if nodeID != "" {
		req.NodeID = nodeID
	}

	var resp WikiSearchResponse
	if err := c.Post("/wiki/v2/nodes/search", req, &resp); err != nil {
		return nil, err
	}

	if resp.Code != 0 {
		return nil, fmt.Errorf("API error %d: %s", resp.Code, resp.Msg)
	}

	return resp.Data.Items, nil
}

// GetWikiNodeChildren retrieves the immediate children of a wiki node
// spaceID: the wiki space ID
// parentNodeToken: the parent node token
func (c *Client) GetWikiNodeChildren(spaceID, parentNodeToken string) ([]WikiNode, error) {
	var allItems []WikiNode
	var pageToken string

	for {
		items, hasMore, nextPageToken, err := c.ListWikiNodes(spaceID, parentNodeToken, 50, pageToken)
		if err != nil {
			return nil, err
		}

		allItems = append(allItems, items...)

		if !hasMore {
			break
		}
		pageToken = nextPageToken
	}

	return allItems, nil
}
