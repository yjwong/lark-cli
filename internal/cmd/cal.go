package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/yjwong/lark-cli/internal/auth"
	"github.com/yjwong/lark-cli/internal/output"
	"github.com/yjwong/lark-cli/internal/scopes"
)

var calCmd = &cobra.Command{
	Use:   "cal",
	Short: "Calendar commands",
	Long:  "Manage Lark calendar events",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		validateScopeGroup("calendar")
	},
}

// validateScopeGroup checks if the required scope group is granted
// and exits with a helpful error message if not
func validateScopeGroup(groupName string) {
	store := auth.GetTokenStore()

	// First check if authenticated at all
	if !store.IsValid() && !store.CanRefresh() {
		output.Fatalf("AUTH_ERROR", "Not authenticated. Run: lark auth login")
	}

	// Check scope group
	if !store.HasScopeGroup(groupName) {
		group := scopes.Groups[groupName]
		output.Fatal("SCOPE_ERROR", fmt.Errorf(
			"missing required permissions for %s commands\n\n"+
				"Required scopes: %v\n\n"+
				"To add these permissions, run:\n"+
				"  lark auth login --scopes %s",
			group.Description, group.Scopes, groupName))
	}
}

func init() {
	calCmd.AddCommand(listCmd)
	calCmd.AddCommand(showCmd)
	calCmd.AddCommand(createCmd)
	calCmd.AddCommand(updateCmd)
	calCmd.AddCommand(deleteCmd)
	calCmd.AddCommand(searchCmd)
	calCmd.AddCommand(freebusyCmd)
	calCmd.AddCommand(rsvpCmd)
	calCmd.AddCommand(lookupUserCmd)
	calCmd.AddCommand(commonFreetimeCmd)
	calCmd.AddCommand(attendeeCmd)
}
