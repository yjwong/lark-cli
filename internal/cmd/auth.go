package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/yjwong/lark-cli/internal/api"
	"github.com/yjwong/lark-cli/internal/auth"
	"github.com/yjwong/lark-cli/internal/output"
	"github.com/yjwong/lark-cli/internal/scopes"
)

var loginScopes string

var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "Authentication commands",
	Long:  "Manage Lark OAuth authentication",
}

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login to Lark",
	Long: `Authenticate with Lark using OAuth browser flow.

By default, all permissions are requested. Use --scopes to request only specific
scope groups for a minimal permission setup.

Scope groups: calendar, contacts, documents, bitable, messages, mail, minutes

Examples:
  lark auth login                           # All permissions (default)
  lark auth login --scopes calendar         # Only calendar permissions
  lark auth login --scopes calendar,contacts # Calendar and contacts`,
	Run: func(cmd *cobra.Command, args []string) {
		opts := auth.LoginOptions{}

		if loginScopes != "" {
			// Parse and validate scope groups
			groups, invalid := scopes.ParseGroups(loginScopes)
			if len(invalid) > 0 {
				output.Fatal("VALIDATION_ERROR", fmt.Errorf("invalid scope groups: %s\nValid groups: %s",
					strings.Join(invalid, ", "),
					strings.Join(scopes.AllGroupNames(), ", ")))
			}
			if len(groups) == 0 {
				output.Fatal("VALIDATION_ERROR", fmt.Errorf("no valid scope groups specified\nValid groups: %s",
					strings.Join(scopes.AllGroupNames(), ", ")))
			}
			opts.ScopeGroups = groups
		}
		// If loginScopes is empty, opts.ScopeGroups remains nil, triggering default (all scopes)

		if err := auth.LoginWithOptions(opts); err != nil {
			output.Fatal("AUTH_ERROR", err)
		}
		output.Success("Successfully authenticated with Lark")
	},
}

var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "Logout from Lark",
	Long:  "Clear stored authentication credentials",
	Run: func(cmd *cobra.Command, args []string) {
		if err := auth.Logout(); err != nil {
			output.Fatal("AUTH_ERROR", err)
		}
		output.Success("Successfully logged out")
	},
}

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show authentication status",
	Long:  "Display current authentication status, token expiry, and granted permissions",
	Run: func(cmd *cobra.Command, args []string) {
		store := auth.GetTokenStore()

		status := api.OutputAuthStatus{
			Authenticated: store.IsValid(),
			ExpiresAt:     store.GetExpiresAt(),
		}

		if !status.Authenticated && store.CanRefresh() {
			// Token expired but we can refresh
			if err := auth.RefreshAccessToken(); err == nil {
				status.Authenticated = true
				status.ExpiresAt = store.GetExpiresAt()
			}
		}

		// Add scope information
		if status.Authenticated {
			status.GrantedGroups = store.GetGrantedGroupsList()
			status.ScopeGroups = store.GetGrantedGroups()
		}

		output.JSON(status)
	},
}

var (
	scopesJsonImport bool
	scopesGroups     string
)

var scopesCmd = &cobra.Command{
	Use:   "scopes",
	Short: "List scope groups or generate batch-import JSON",
	Long: `Display available scope groups and their permissions.

Use --json-import to generate JSON for Lark's scope batch import feature,
which configures your app's tenant and user scopes.

Examples:
  lark auth scopes                                    # List all scope groups
  lark auth scopes --json-import                      # Batch-import JSON (all groups)
  lark auth scopes --json-import --groups calendar     # Batch-import JSON (specific groups)
  lark auth scopes --json-import --groups calendar,contacts`,
	Run: func(cmd *cobra.Command, args []string) {
		if scopesJsonImport {
			// Generate batch-import JSON
			var groupNames []string
			if scopesGroups != "" {
				groups, invalid := scopes.ParseGroups(scopesGroups)
				if len(invalid) > 0 {
					output.Fatal("VALIDATION_ERROR", fmt.Errorf("invalid scope groups: %s\nValid groups: %s",
						strings.Join(invalid, ", "),
						strings.Join(scopes.AllGroupNames(), ", ")))
				}
				groupNames = groups
			} else {
				groupNames = scopes.AllGroupNames()
			}

			jsonBytes, err := scopes.GenerateBatchImportJSON(groupNames)
			if err != nil {
				output.Fatal("SCOPE_ERROR", err)
			}
			fmt.Println(string(jsonBytes))
			return
		}

		// Default: list scope groups
		type scopeGroupOutput struct {
			Name         string   `json:"name"`
			Description  string   `json:"description"`
			Commands     []string `json:"commands"`
			OAuthScopes  int      `json:"oauth_scopes"`
			TenantScopes int      `json:"tenant_scopes"`
			UserScopes   int      `json:"user_scopes"`
		}

		groups := make([]scopeGroupOutput, 0, len(scopes.AllGroupNames()))
		for _, name := range scopes.AllGroupNames() {
			group := scopes.Groups[name]
			cat := scopes.Catalog[name]
			groups = append(groups, scopeGroupOutput{
				Name:         group.Name,
				Description:  group.Description,
				Commands:     group.Commands,
				OAuthScopes:  len(group.Scopes),
				TenantScopes: len(cat.TenantScopes),
				UserScopes:   len(cat.UserScopes),
			})
		}

		output.JSON(map[string]interface{}{
			"groups": groups,
			"usage":  "lark auth scopes --json-import [--groups <group1,group2,...>]",
		})
	},
}

func init() {
	loginCmd.Flags().StringVar(&loginScopes, "scopes", "", "Comma-separated scope groups (calendar,contacts,documents,bitable,messages,mail,minutes)")

	scopesCmd.Flags().BoolVar(&scopesJsonImport, "json-import", false, "Output in Lark batch-import JSON format")
	scopesCmd.Flags().StringVar(&scopesGroups, "groups", "", "Comma-separated scope groups to include")

	authCmd.AddCommand(loginCmd)
	authCmd.AddCommand(logoutCmd)
	authCmd.AddCommand(statusCmd)
	authCmd.AddCommand(scopesCmd)
}
