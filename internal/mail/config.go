package mail

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/yjwong/lark-cli/internal/config"
)

// Credentials holds IMAP connection settings
type Credentials struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	UseSSL   bool   `json:"use_ssl"`
}

// CredentialsFilePath returns the path to the mail credentials file
func CredentialsFilePath() string {
	return filepath.Join(config.GetConfigDir(), "mail.json")
}

// CacheFilePath returns the path to the mail cache database
func CacheFilePath() string {
	return filepath.Join(config.GetConfigDir(), "mail_cache.db")
}

// LoadCredentials reads IMAP credentials from disk
func LoadCredentials() (*Credentials, error) {
	path := CredentialsFilePath()
	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("mail not configured; run 'lark mail setup' first")
		}
		return nil, fmt.Errorf("failed to read mail credentials: %w", err)
	}

	var creds Credentials
	if err := json.Unmarshal(data, &creds); err != nil {
		return nil, fmt.Errorf("failed to parse mail credentials: %w", err)
	}

	return &creds, nil
}

// SaveCredentials writes IMAP credentials to disk
func SaveCredentials(creds *Credentials) error {
	data, err := json.MarshalIndent(creds, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal credentials: %w", err)
	}

	path := CredentialsFilePath()
	if err := os.WriteFile(path, data, 0600); err != nil {
		return fmt.Errorf("failed to write credentials: %w", err)
	}

	return nil
}

// ClearCredentials removes stored credentials
func ClearCredentials() error {
	path := CredentialsFilePath()
	if err := os.Remove(path); err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("failed to remove credentials: %w", err)
	}
	return nil
}

// HasCredentials checks if credentials are configured
func HasCredentials() bool {
	path := CredentialsFilePath()
	_, err := os.Stat(path)
	return err == nil
}
