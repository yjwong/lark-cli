package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/yjwong/lark-cli/internal/auth"
	"github.com/yjwong/lark-cli/internal/config"
)

const (
	defaultTimeout = 30 * time.Second
)

func getBaseURL() string {
	if config.GetRegion() == "feishu" {
		return "https://open.feishu.cn/open-apis"
	}
	return "https://open.larksuite.com/open-apis"
}

// Client is the Lark API client
type Client struct {
	httpClient *http.Client
}

// NewClient creates a new API client
func NewClient() *Client {
	return &Client{
		httpClient: &http.Client{
			Timeout: defaultTimeout,
		},
	}
}

// doRequest performs an authenticated HTTP request
func (c *Client) doRequest(method, path string, body interface{}, result interface{}) error {
	// Ensure we have a valid token
	if err := auth.EnsureValidToken(); err != nil {
		return err
	}

	var reqBody io.Reader
	if body != nil {
		jsonBody, err := json.Marshal(body)
		if err != nil {
			return fmt.Errorf("failed to marshal request body: %w", err)
		}
		reqBody = bytes.NewBuffer(jsonBody)
	}

	url := getBaseURL() + path
	req, err := http.NewRequest(method, url, reqBody)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	// Set headers
	token := auth.GetTokenStore().GetAccessToken()
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	// Execute request
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	// Read response
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response: %w", err)
	}

	// Parse response
	if result != nil {
		if err := json.Unmarshal(respBody, result); err != nil {
			return fmt.Errorf("failed to parse response: %w", err)
		}
	}

	return nil
}

// Get performs a GET request
func (c *Client) Get(path string, result interface{}) error {
	return c.doRequest("GET", path, nil, result)
}

// Post performs a POST request
func (c *Client) Post(path string, body interface{}, result interface{}) error {
	return c.doRequest("POST", path, body, result)
}

// Patch performs a PATCH request
func (c *Client) Patch(path string, body interface{}, result interface{}) error {
	return c.doRequest("PATCH", path, body, result)
}

// Delete performs a DELETE request
func (c *Client) Delete(path string, result interface{}) error {
	return c.doRequest("DELETE", path, nil, result)
}

// doRequestWithTenantToken performs an HTTP request using tenant access token
func (c *Client) doRequestWithTenantToken(method, path string, body interface{}, result interface{}) error {
	// Ensure we have a valid tenant token
	if err := auth.EnsureValidTenantToken(); err != nil {
		return err
	}

	var reqBody io.Reader
	if body != nil {
		jsonBody, err := json.Marshal(body)
		if err != nil {
			return fmt.Errorf("failed to marshal request body: %w", err)
		}
		reqBody = bytes.NewBuffer(jsonBody)
	}

	url := getBaseURL() + path
	req, err := http.NewRequest(method, url, reqBody)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	// Set headers with tenant token
	token := auth.GetTenantTokenStore().GetAccessToken()
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	// Execute request
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	// Read response
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response: %w", err)
	}

	// Parse response
	if result != nil {
		if err := json.Unmarshal(respBody, result); err != nil {
			return fmt.Errorf("failed to parse response: %w", err)
		}
	}

	return nil
}

// PostWithTenantToken performs a POST request using tenant access token
func (c *Client) PostWithTenantToken(path string, body interface{}, result interface{}) error {
	return c.doRequestWithTenantToken("POST", path, body, result)
}

// GetWithTenantToken performs a GET request using tenant access token
func (c *Client) GetWithTenantToken(path string, result interface{}) error {
	return c.doRequestWithTenantToken("GET", path, nil, result)
}

// DeleteWithTenantToken performs a DELETE request using tenant access token
func (c *Client) DeleteWithTenantToken(path string, result interface{}) error {
	return c.doRequestWithTenantToken("DELETE", path, nil, result)
}

// DownloadWithTenantToken performs a GET request that returns binary data
// The caller is responsible for closing the returned ReadCloser
func (c *Client) DownloadWithTenantToken(path string) (io.ReadCloser, string, error) {
	// Ensure we have a valid tenant token
	if err := auth.EnsureValidTenantToken(); err != nil {
		return nil, "", err
	}

	url := getBaseURL() + path
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, "", fmt.Errorf("failed to create request: %w", err)
	}

	// Set headers with tenant token
	token := auth.GetTenantTokenStore().GetAccessToken()
	req.Header.Set("Authorization", "Bearer "+token)

	// Execute request
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, "", fmt.Errorf("request failed: %w", err)
	}

	// Check for error response (non-2xx status)
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		defer resp.Body.Close()
		body, _ := io.ReadAll(resp.Body)
		return nil, "", fmt.Errorf("API error (HTTP %d): %s", resp.StatusCode, string(body))
	}

	contentType := resp.Header.Get("Content-Type")
	return resp.Body, contentType, nil
}

// Download performs a GET request that returns binary data using user access token
// The caller is responsible for closing the returned ReadCloser
func (c *Client) Download(path string) (io.ReadCloser, string, error) {
	// Ensure we have a valid token
	if err := auth.EnsureValidToken(); err != nil {
		return nil, "", err
	}

	url := getBaseURL() + path
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, "", fmt.Errorf("failed to create request: %w", err)
	}

	// Set headers with user token
	token := auth.GetTokenStore().GetAccessToken()
	req.Header.Set("Authorization", "Bearer "+token)

	// Execute request
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, "", fmt.Errorf("request failed: %w", err)
	}

	// Check for error response (non-2xx status)
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		defer resp.Body.Close()
		body, _ := io.ReadAll(resp.Body)
		return nil, "", fmt.Errorf("API error (HTTP %d): %s", resp.StatusCode, string(body))
	}

	contentType := resp.Header.Get("Content-Type")
	return resp.Body, contentType, nil
}
