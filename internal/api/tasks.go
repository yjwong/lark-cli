package api

import (
	"fmt"
	"net/url"
	"strconv"
)

// TaskListOptions configures the list tasks request
type TaskListOptions struct {
	PageSize  int    // Tasks per page (max 100)
	PageToken string // Pagination token
	Completed bool   // Include completed tasks
}

// ListTasks lists tasks for the current user
func (c *Client) ListTasks(opts *TaskListOptions) ([]Task, bool, string, error) {
	pageSize := 50
	if opts != nil && opts.PageSize > 0 {
		pageSize = opts.PageSize
		if pageSize > 100 {
			pageSize = 100
		}
	}

	params := url.Values{}
	params.Set("page_size", strconv.Itoa(pageSize))
	params.Set("user_id_type", "open_id")

	if opts != nil {
		if opts.PageToken != "" {
			params.Set("page_token", opts.PageToken)
		}
		if opts.Completed {
			params.Set("completed", "true")
		}
	}

	path := "/task/v2/tasks?" + params.Encode()

	var resp TaskListResponse
	if err := c.Get(path, &resp); err != nil {
		return nil, false, "", err
	}

	if resp.Code != 0 {
		return nil, false, "", fmt.Errorf("API error %d: %s", resp.Code, resp.Msg)
	}

	return resp.Data.Items, resp.Data.HasMore, resp.Data.PageToken, nil
}

// GetTask retrieves a single task by GUID
func (c *Client) GetTask(taskGUID string) (*Task, error) {
	params := url.Values{}
	params.Set("user_id_type", "open_id")

	path := fmt.Sprintf("/task/v2/tasks/%s?%s", url.PathEscape(taskGUID), params.Encode())

	var resp TaskResponse
	if err := c.Get(path, &resp); err != nil {
		return nil, err
	}

	if resp.Code != 0 {
		return nil, fmt.Errorf("API error %d: %s", resp.Code, resp.Msg)
	}

	return resp.Data.Task, nil
}
