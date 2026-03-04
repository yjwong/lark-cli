package api

import (
	"fmt"
	"net/url"
	"strconv"
)

// ListBitableTables lists all tables in a Bitable app
func (c *Client) ListBitableTables(appToken string) ([]BitableTable, error) {
	return PaginateWith(func(pageToken string) ([]BitableTable, bool, string, error) {
		path := fmt.Sprintf("/bitable/v1/apps/%s/tables?page_size=100", url.PathEscape(appToken))
		if pageToken != "" {
			path += "&page_token=" + url.QueryEscape(pageToken)
		}

		var resp BitableTablesResponse
		if err := c.Get(path, &resp); err != nil {
			return nil, false, "", err
		}
		if err := resp.Err(); err != nil {
			return nil, false, "", err
		}
		return resp.Data.Items, resp.Data.HasMore, resp.Data.PageToken, nil
	})
}

// ListBitableFields lists all fields in a Bitable table
func (c *Client) ListBitableFields(appToken, tableID string) ([]BitableField, error) {
	return PaginateWith(func(pageToken string) ([]BitableField, bool, string, error) {
		path := fmt.Sprintf("/bitable/v1/apps/%s/tables/%s/fields?page_size=100",
			url.PathEscape(appToken), url.PathEscape(tableID))
		if pageToken != "" {
			path += "&page_token=" + url.QueryEscape(pageToken)
		}

		var resp BitableFieldsResponse
		if err := c.Get(path, &resp); err != nil {
			return nil, false, "", err
		}
		if err := resp.Err(); err != nil {
			return nil, false, "", err
		}
		return resp.Data.Items, resp.Data.HasMore, resp.Data.PageToken, nil
	})
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
	reqSize := 0
	if opts != nil {
		reqSize = opts.PageSize
	}
	pageSize := ClampPageSize(reqSize, 100, 500)

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

	if err := resp.Err(); err != nil {
		return nil, false, "", err
	}

	return resp.Data.Items, resp.Data.HasMore, resp.Data.PageToken, nil
}
