package api

import (
	"fmt"
	"net/url"
)

// GetTask retrieves task details by GUID
func (c *Client) GetTask(taskGUID string) (*TaskDetail, error) {
	path := fmt.Sprintf("/task/v2/tasks/%s", url.PathEscape(taskGUID))

	var resp TaskDetailResponse
	if err := c.Get(path, &resp); err != nil {
		return nil, err
	}

	if resp.Code != 0 {
		return nil, fmt.Errorf("API error %d: %s", resp.Code, resp.Msg)
	}

	return resp.Data.Task, nil
}
