package api

import (
	"fmt"
	"net/url"
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

	if err := resp.Err(); err != nil {
		return nil, err
	}

	return resp.Data.Node, nil
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

	if err := resp.Err(); err != nil {
		return nil, err
	}

	return resp.Data.Items, nil
}

// GetWikiNodeChildren retrieves the immediate children of a wiki node
// spaceID: the wiki space ID
// parentNodeToken: the parent node token
func (c *Client) GetWikiNodeChildren(spaceID, parentNodeToken string) ([]WikiNode, error) {
	return PaginateWith(func(pageToken string) ([]WikiNode, bool, string, error) {
		params := url.Values{}
		params.Set("parent_node_token", parentNodeToken)
		params.Set("page_size", "50")
		if pageToken != "" {
			params.Set("page_token", pageToken)
		}

		path := fmt.Sprintf("/wiki/v2/spaces/%s/nodes?%s",
			url.PathEscape(spaceID), params.Encode())

		var resp ListWikiChildrenResponse
		if err := c.Get(path, &resp); err != nil {
			return nil, false, "", err
		}
		if err := resp.Err(); err != nil {
			return nil, false, "", err
		}
		return resp.Data.Items, resp.Data.HasMore, resp.Data.PageToken, nil
	})
}
