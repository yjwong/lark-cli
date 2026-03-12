package api


// UserInfo represents the current user's information
type UserInfo struct {
	Name      string `json:"name"`
	EnName    string `json:"en_name"`
	AvatarURL string `json:"avatar_url"`
	OpenID    string `json:"open_id"`
	UnionID   string `json:"union_id"`
	Email     string `json:"email,omitempty"`
	TenantKey string `json:"tenant_key"`
}

// UserInfoResponse is the response from the user info API
type UserInfoResponse struct {
	BaseResponse
	Data UserInfo `json:"data"`
}

// GetCurrentUser retrieves the current user's information
func (c *Client) GetCurrentUser() (*UserInfo, error) {
	var resp UserInfoResponse

	if err := c.Get("/authen/v1/user_info", &resp); err != nil {
		return nil, err
	}

	if err := resp.Err(); err != nil {
		return nil, err
	}

	return &resp.Data, nil
}

// UserLookupOptions configures a user lookup query
type UserLookupOptions struct {
	Emails  []string
	Mobiles []string
}

// LookupUsers looks up user IDs by email or mobile number
// Note: This API requires tenant_access_token, not user_access_token
func (c *Client) LookupUsers(opts UserLookupOptions) ([]UserContactInfo, error) {
	req := UserLookupRequest{
		Emails:  opts.Emails,
		Mobiles: opts.Mobiles,
	}

	var resp UserLookupResponse
	// Use tenant token for contacts API
	if err := c.PostWithTenantToken("/contact/v3/users/batch_get_id?user_id_type=open_id", req, &resp); err != nil {
		return nil, err
	}

	if err := resp.Err(); err != nil {
		return nil, err
	}

	return resp.Data.UserList, nil
}
