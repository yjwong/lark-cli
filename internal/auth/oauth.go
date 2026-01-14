package auth

import (
	"bytes"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os/exec"
	"runtime"
	"time"

	"github.com/yjwong/lark-cli/internal/config"
)

const (
	authorizationURL     = "https://accounts.larksuite.com/open-apis/authen/v1/authorize"
	tokenURL             = "https://open.larksuite.com/open-apis/authen/v2/oauth/token"
	tenantTokenURL       = "https://open.larksuite.com/open-apis/auth/v3/tenant_access_token/internal"
	defaultTimeout       = 5 * time.Minute
)

// TokenResponse represents the OAuth token response from Lark
type TokenResponse struct {
	Code                  int    `json:"code"`
	AccessToken           string `json:"access_token"`
	ExpiresIn             int    `json:"expires_in"`
	RefreshToken          string `json:"refresh_token"`
	RefreshTokenExpiresIn int    `json:"refresh_token_expires_in"`
	TokenType             string `json:"token_type"`
	Scope                 string `json:"scope"`
	Error                 string `json:"error"`
	ErrorDescription      string `json:"error_description"`
}

// TenantTokenResponse represents the tenant access token response from Lark
type TenantTokenResponse struct {
	Code              int    `json:"code"`
	Msg               string `json:"msg"`
	TenantAccessToken string `json:"tenant_access_token"`
	Expire            int    `json:"expire"`
}

// Login performs the OAuth login flow
func Login() error {
	appID := config.GetAppID()
	appSecret := config.GetAppSecret()

	if appID == "" {
		return fmt.Errorf("app_id not configured. Set it in .lark-cal/config.yaml or LARK_APP_ID env var")
	}
	if appSecret == "" {
		return fmt.Errorf("app_secret not configured. Set LARK_APP_SECRET env var")
	}

	// Generate state for CSRF protection
	state, err := generateState()
	if err != nil {
		return fmt.Errorf("failed to generate state: %w", err)
	}

	// Start callback server
	port := config.GetRedirectPort()
	server := NewCallbackServer(port)
	if err := server.Start(state); err != nil {
		return fmt.Errorf("failed to start callback server: %w", err)
	}
	defer server.Stop()

	// Build authorization URL
	redirectURI := server.GetRedirectURI()
	authURL := buildAuthorizationURL(appID, redirectURI, state)

	// Open browser
	fmt.Printf("Opening browser for authentication...\n")
	fmt.Printf("If the browser doesn't open, visit this URL:\n%s\n\n", authURL)

	if err := openBrowser(authURL); err != nil {
		fmt.Printf("Warning: Could not open browser automatically: %v\n", err)
	}

	fmt.Println("Waiting for authorization...")

	// Wait for callback
	code, err := server.WaitForCode(defaultTimeout)
	if err != nil {
		return fmt.Errorf("authorization failed: %w", err)
	}

	fmt.Println("Authorization code received, exchanging for tokens...")

	// Exchange code for tokens
	tokenResp, err := exchangeCodeForTokens(appID, appSecret, code, redirectURI)
	if err != nil {
		return fmt.Errorf("failed to exchange code: %w", err)
	}

	// Store tokens
	store := GetTokenStore()
	if err := store.Update(
		tokenResp.AccessToken,
		tokenResp.RefreshToken,
		tokenResp.ExpiresIn,
		tokenResp.RefreshTokenExpiresIn,
		tokenResp.Scope,
	); err != nil {
		return fmt.Errorf("failed to save tokens: %w", err)
	}

	fmt.Println("Authentication successful!")
	return nil
}

// RefreshAccessToken refreshes the access token using the refresh token
func RefreshAccessToken() error {
	store := GetTokenStore()
	if !store.CanRefresh() {
		return fmt.Errorf("no valid refresh token available, please login again")
	}

	appID := config.GetAppID()
	appSecret := config.GetAppSecret()

	if appID == "" || appSecret == "" {
		return fmt.Errorf("app credentials not configured")
	}

	tokenResp, err := refreshTokens(appID, appSecret, store.GetRefreshToken())
	if err != nil {
		return fmt.Errorf("failed to refresh token: %w", err)
	}

	if err := store.Update(
		tokenResp.AccessToken,
		tokenResp.RefreshToken,
		tokenResp.ExpiresIn,
		tokenResp.RefreshTokenExpiresIn,
		tokenResp.Scope,
	); err != nil {
		return fmt.Errorf("failed to save refreshed tokens: %w", err)
	}

	return nil
}

// EnsureValidToken checks and refreshes the token if needed
func EnsureValidToken() error {
	store := GetTokenStore()

	if store.IsValid() {
		// Token is still valid
		if store.NeedsRefresh() && store.CanRefresh() {
			// Proactively refresh before expiry
			if err := RefreshAccessToken(); err != nil {
				// Log but don't fail - current token is still valid
				fmt.Printf("Warning: Failed to proactively refresh token: %v\n", err)
			}
		}
		return nil
	}

	// Token is invalid or expired
	if store.CanRefresh() {
		return RefreshAccessToken()
	}

	return fmt.Errorf("no valid authentication, please run 'lark auth login'")
}

// Logout clears stored credentials
func Logout() error {
	store := GetTokenStore()
	return store.Clear()
}

// EnsureValidTenantToken ensures we have a valid tenant access token
func EnsureValidTenantToken() error {
	store := GetTenantTokenStore()

	if store.IsValid() {
		return nil
	}

	// Need to fetch a new tenant token
	return RefreshTenantToken()
}

// RefreshTenantToken fetches a new tenant access token
func RefreshTenantToken() error {
	appID := config.GetAppID()
	appSecret := config.GetAppSecret()

	if appID == "" {
		return fmt.Errorf("app_id not configured")
	}
	if appSecret == "" {
		return fmt.Errorf("app_secret not configured")
	}

	reqBody := map[string]string{
		"app_id":     appID,
		"app_secret": appSecret,
	}

	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return fmt.Errorf("failed to marshal request: %w", err)
	}

	req, err := http.NewRequest("POST", tenantTokenURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response: %w", err)
	}

	var tokenResp TenantTokenResponse
	if err := json.Unmarshal(body, &tokenResp); err != nil {
		return fmt.Errorf("failed to parse response: %w", err)
	}

	if tokenResp.Code != 0 {
		return fmt.Errorf("tenant token request failed (code %d): %s", tokenResp.Code, tokenResp.Msg)
	}

	if tokenResp.TenantAccessToken == "" {
		return fmt.Errorf("no tenant access token in response")
	}

	// Store the token
	store := GetTenantTokenStore()
	if err := store.Update(tokenResp.TenantAccessToken, tokenResp.Expire); err != nil {
		return fmt.Errorf("failed to save tenant token: %w", err)
	}

	return nil
}

// generateState creates a random state string for CSRF protection
func generateState() (string, error) {
	bytes := make([]byte, 16)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

// buildAuthorizationURL constructs the OAuth authorization URL
func buildAuthorizationURL(appID, redirectURI, state string) string {
	params := url.Values{}
	params.Set("client_id", appID)
	params.Set("redirect_uri", redirectURI)
	params.Set("scope", "calendar:calendar calendar:calendar:readonly contact:contact.base:readonly contact:department.base:readonly contact:department.organize:readonly contact:user:search docx:document:readonly docs:doc:readonly docs:document.content:read docs:document.comment:read drive:drive:readonly wiki:wiki:readonly space:document:retrieve offline_access")
	params.Set("state", state)

	return authorizationURL + "?" + params.Encode()
}

// exchangeCodeForTokens exchanges the authorization code for access tokens
func exchangeCodeForTokens(appID, appSecret, code, redirectURI string) (*TokenResponse, error) {
	reqBody := map[string]string{
		"grant_type":    "authorization_code",
		"client_id":     appID,
		"client_secret": appSecret,
		"code":          code,
		"redirect_uri":  redirectURI,
	}

	return doTokenRequest(reqBody)
}

// refreshTokens exchanges a refresh token for new tokens
func refreshTokens(appID, appSecret, refreshToken string) (*TokenResponse, error) {
	reqBody := map[string]string{
		"grant_type":    "refresh_token",
		"client_id":     appID,
		"client_secret": appSecret,
		"refresh_token": refreshToken,
	}

	return doTokenRequest(reqBody)
}

// doTokenRequest performs a token request to Lark's OAuth endpoint
func doTokenRequest(reqBody map[string]string) (*TokenResponse, error) {
	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	req, err := http.NewRequest("POST", tokenURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	var tokenResp TokenResponse
	if err := json.Unmarshal(body, &tokenResp); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	if tokenResp.Code != 0 {
		errMsg := tokenResp.ErrorDescription
		if errMsg == "" {
			errMsg = tokenResp.Error
		}
		return nil, fmt.Errorf("token request failed (code %d): %s", tokenResp.Code, errMsg)
	}

	if tokenResp.AccessToken == "" {
		return nil, fmt.Errorf("no access token in response")
	}

	return &tokenResp, nil
}

// openBrowser opens the specified URL in the default browser
func openBrowser(url string) error {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "darwin":
		cmd = exec.Command("open", url)
	case "linux":
		cmd = exec.Command("xdg-open", url)
	case "windows":
		cmd = exec.Command("cmd", "/c", "start", url)
	default:
		return fmt.Errorf("unsupported platform: %s", runtime.GOOS)
	}

	return cmd.Start()
}
