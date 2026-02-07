package api

import (
	"fmt"
	"net/url"
	"strconv"
)

// ListBitableTables lists all tables in a Bitable app
func (c *Client) ListBitableTables(appToken string) ([]BitableTable, error) {
	var allTables []BitableTable
	pageToken := ""

	for {
		path := fmt.Sprintf("/bitable/v1/apps/%s/tables?page_size=100", url.PathEscape(appToken))
		if pageToken != "" {
			path += "&page_token=" + url.QueryEscape(pageToken)
		}

		var resp BitableTablesResponse
		if err := c.Get(path, &resp); err != nil {
			return nil, err
		}

		if resp.Code != 0 {
			return nil, fmt.Errorf("API error %d: %s", resp.Code, resp.Msg)
		}

		allTables = append(allTables, resp.Data.Items...)

		if !resp.Data.HasMore || resp.Data.PageToken == "" {
			break
		}
		pageToken = resp.Data.PageToken
	}

	return allTables, nil
}

// ListBitableFields lists all fields in a Bitable table
func (c *Client) ListBitableFields(appToken, tableID string) ([]BitableField, error) {
	var allFields []BitableField
	pageToken := ""

	for {
		path := fmt.Sprintf("/bitable/v1/apps/%s/tables/%s/fields?page_size=100",
			url.PathEscape(appToken), url.PathEscape(tableID))
		if pageToken != "" {
			path += "&page_token=" + url.QueryEscape(pageToken)
		}

		var resp BitableFieldsResponse
		if err := c.Get(path, &resp); err != nil {
			return nil, err
		}

		if resp.Code != 0 {
			return nil, fmt.Errorf("API error %d: %s", resp.Code, resp.Msg)
		}

		allFields = append(allFields, resp.Data.Items...)

		if !resp.Data.HasMore || resp.Data.PageToken == "" {
			break
		}
		pageToken = resp.Data.PageToken
	}

	return allFields, nil
}

// BitableRecordOptions configures the list records request
type BitableRecordOptions struct {
	ViewID    string   // Optional view ID to filter records
	FieldIDs  []string // Optional field IDs to return
	Sort      string   // Optional sort expression
	Filter    string   // Optional filter expression
	PageSize  int      // Records per page (max 500)
	PageToken string   // Pagination token
}

// ListBitableRecords lists records in a Bitable table
func (c *Client) ListBitableRecords(appToken, tableID string, opts *BitableRecordOptions) ([]BitableRecord, bool, string, error) {
	pageSize := 100
	if opts != nil && opts.PageSize > 0 {
		pageSize = opts.PageSize
		if pageSize > 500 {
			pageSize = 500
		}
	}

	params := url.Values{}
	params.Set("page_size", strconv.Itoa(pageSize))

	if opts != nil {
		if opts.ViewID != "" {
			params.Set("view_id", opts.ViewID)
		}
		if len(opts.FieldIDs) > 0 {
			for _, fid := range opts.FieldIDs {
				params.Add("field_names", fid)
			}
		}
		if opts.Sort != "" {
			params.Set("sort", opts.Sort)
		}
		if opts.Filter != "" {
			params.Set("filter", opts.Filter)
		}
		if opts.PageToken != "" {
			params.Set("page_token", opts.PageToken)
		}
	}

	path := fmt.Sprintf("/bitable/v1/apps/%s/tables/%s/records?%s",
		url.PathEscape(appToken), url.PathEscape(tableID), params.Encode())

	var resp BitableRecordsResponse
	if err := c.Get(path, &resp); err != nil {
		return nil, false, "", err
	}

	if resp.Code != 0 {
		return nil, false, "", fmt.Errorf("API error %d: %s", resp.Code, resp.Msg)
	}

	return resp.Data.Items, resp.Data.HasMore, resp.Data.PageToken, nil
}
