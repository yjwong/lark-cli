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
	reqSize := 0
	if opts != nil {
		reqSize = opts.PageSize
	}
	pageSize := ClampPageSize(reqSize, 50, 100)

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

	if err := resp.Err(); err != nil {
		return nil, false, "", err
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

	if err := resp.Err(); err != nil {
		return nil, err
	}

	return resp.Data.Task, nil
}
