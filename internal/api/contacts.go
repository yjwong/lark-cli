package api

import (
	"fmt"
	"net/url"
)

// GetUser retrieves a single user by ID
// userID: the user identifier (open_id, union_id, or user_id based on idType)
// idType: "open_id", "union_id", or "user_id" (defaults to "open_id")
func (c *Client) GetUser(userID string, idType string) (*ContactUser, error) {
	if idType == "" {
		idType = "open_id"
	}

	path := fmt.Sprintf("/contact/v3/users/%s?user_id_type=%s", url.PathEscape(userID), idType)

	var resp GetUserResponse
	if err := c.GetWithTenantToken(path, &resp); err != nil {
		return nil, err
	}

	if resp.Code != 0 {
		return nil, fmt.Errorf("API error %d: %s", resp.Code, resp.Msg)
	}

	return resp.Data.User, nil
}

// BatchGetUsers retrieves multiple users by IDs in a single API call
func (c *Client) BatchGetUsers(userIDs []string, idType string) ([]ContactUser, error) {
	if len(userIDs) == 0 {
		return nil, nil
	}
	if idType == "" {
		idType = "open_id"
	}

	params := url.Values{}
	params.Set("user_id_type", idType)
	for _, id := range userIDs {
		params.Add("user_ids", id)
	}

	path := "/contact/v3/users/batch?" + params.Encode()

	var resp BatchGetUserResponse
	if err := c.Get(path, &resp); err != nil {
		return nil, err
	}

	if resp.Code != 0 {
		return nil, fmt.Errorf("API error %d: %s", resp.Code, resp.Msg)
	}

	return resp.Data.Items, nil
}

// ListUsersByDepartment retrieves users directly under a department
// deptID: the department ID (use "0" for root department)
// pageSize: number of results per page (max 50)
// pageToken: pagination token (empty for first page)
func (c *Client) ListUsersByDepartment(deptID string, pageSize int, pageToken string) ([]ContactUser, bool, string, error) {
	if pageSize <= 0 {
		pageSize = 50
	}
	if pageSize > 50 {
		pageSize = 50
	}

	path := fmt.Sprintf("/contact/v3/users/find_by_department?department_id=%s&department_id_type=open_department_id&user_id_type=open_id&page_size=%d",
		url.QueryEscape(deptID), pageSize)

	if pageToken != "" {
		path += "&page_token=" + url.QueryEscape(pageToken)
	}

	var resp FindByDepartmentResponse
	if err := c.GetWithTenantToken(path, &resp); err != nil {
		return nil, false, "", err
	}

	if resp.Code != 0 {
		return nil, false, "", fmt.Errorf("API error %d: %s", resp.Code, resp.Msg)
	}

	return resp.Data.Items, resp.Data.HasMore, resp.Data.PageToken, nil
}

// GetDepartment retrieves a single department by ID
// deptID: the department ID (use "0" for root department)
func (c *Client) GetDepartment(deptID string) (*Department, error) {
	path := fmt.Sprintf("/contact/v3/departments/%s?department_id_type=open_department_id", url.PathEscape(deptID))

	var resp GetDepartmentResponse
	if err := c.GetWithTenantToken(path, &resp); err != nil {
		return nil, err
	}

	if resp.Code != 0 {
		return nil, fmt.Errorf("API error %d: %s", resp.Code, resp.Msg)
	}

	return resp.Data.Department, nil
}

// SearchDepartmentsRequest is the request body for department search
type SearchDepartmentsRequest struct {
	Query string `json:"query"`
}

// SearchUsers searches for users by name keyword
// Note: This requires user_access_token (contact:user:search scope)
func (c *Client) SearchUsers(query string, pageSize int, pageToken string) ([]SearchUserResult, bool, string, error) {
	if pageSize <= 0 {
		pageSize = 20
	}
	if pageSize > 200 {
		pageSize = 200
	}

	path := fmt.Sprintf("/search/v1/user?query=%s&page_size=%d",
		url.QueryEscape(query), pageSize)

	if pageToken != "" {
		path += "&page_token=" + url.QueryEscape(pageToken)
	}

	var resp SearchUsersResponse
	// User search requires user_access_token
	if err := c.Get(path, &resp); err != nil {
		return nil, false, "", err
	}

	if resp.Code != 0 {
		return nil, false, "", fmt.Errorf("API error %d: %s", resp.Code, resp.Msg)
	}

	return resp.Data.Users, resp.Data.HasMore, resp.Data.PageToken, nil
}

// SearchDepartments searches for departments by keyword
// Note: This requires user_access_token, not tenant_access_token
// For now, we'll use tenant token and handle the limitation
func (c *Client) SearchDepartments(query string, pageSize int, pageToken string) ([]Department, bool, string, error) {
	if pageSize <= 0 {
		pageSize = 20
	}
	if pageSize > 50 {
		pageSize = 50
	}

	path := fmt.Sprintf("/contact/v3/departments/search?page_size=%d", pageSize)
	if pageToken != "" {
		path += "&page_token=" + url.QueryEscape(pageToken)
	}

	reqBody := SearchDepartmentsRequest{Query: query}

	var resp SearchDepartmentsResponse
	// Note: Department search requires user_access_token per API docs
	// Using Post (user token) instead of PostWithTenantToken
	if err := c.Post(path, reqBody, &resp); err != nil {
		return nil, false, "", err
	}

	if resp.Code != 0 {
		return nil, false, "", fmt.Errorf("API error %d: %s", resp.Code, resp.Msg)
	}

	return resp.Data.Items, resp.Data.HasMore, resp.Data.PageToken, nil
}
