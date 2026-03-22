package cmd

import (
	"github.com/spf13/cobra"
	"github.com/yjwong/lark-cli/internal/api"
	"github.com/yjwong/lark-cli/internal/output"
)

var contactCmd = &cobra.Command{
	Use:   "contact",
	Short: "Contact commands",
	Long:  "Look up users and departments in the company directory",
	// No PersistentPreRun: contact get/list-dept use tenant token (can't pre-validate),
	// contact search/search-dept use user token (validated per-subcommand below).
}

// --- contact get ---

var contactGetIDType string

var contactGetCmd = &cobra.Command{
	Use:   "get <user_id>",
	Short: "Get user information by ID",
	Long: `Look up a single user's information by their user ID.

Examples:
  lark contact get ou_xxxx
  lark contact get ou_xxxx --id-type open_id
  lark contact get 12345 --id-type user_id`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		userID := args[0]

		client := api.NewClient()

		user, err := client.GetUser(userID, contactGetIDType)
		if err != nil {
			output.Fatal("API_ERROR", err)
		}

		if user == nil {
			output.Fatalf("NOT_FOUND", "User not found: %s", userID)
		}

		// Fetch department name for the first department
		var deptName string
		if len(user.DepartmentIDs) > 0 {
			dept, err := client.GetDepartment(user.DepartmentIDs[0])
			if err == nil && dept != nil {
				deptName = dept.Name
			}
		}

		result := api.OutputContact{
			UserID:     user.OpenID,
			OpenID:     user.OpenID,
			Name:       user.Name,
			EnName:     user.EnName,
			Email:      user.Email,
			JobTitle:   user.JobTitle,
			Department: deptName,
		}

		output.JSON(result)
	},
}

// --- contact list-dept ---

var contactListDeptPageSize int

var contactListDeptCmd = &cobra.Command{
	Use:   "list-dept [department_id]",
	Short: "List users in a department",
	Long: `List all users directly under a department.

If no department ID is provided, lists users in the root department.

Examples:
  lark contact list-dept
  lark contact list-dept od_xxxx
  lark contact list-dept 0`,
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		deptID := "0" // root department
		if len(args) > 0 {
			deptID = args[0]
		}

		client := api.NewClient()

		// Fetch all users (paginated)
		var allUsers []api.ContactUser
		var pageToken string
		hasMore := true

		for hasMore {
			users, more, nextToken, err := client.ListUsersByDepartment(deptID, contactListDeptPageSize, pageToken)
			if err != nil {
				output.Fatal("API_ERROR", err)
			}
			allUsers = append(allUsers, users...)
			hasMore = more
			pageToken = nextToken
		}

		// Get department name for context
		var deptName string
		if deptID != "0" {
			dept, err := client.GetDepartment(deptID)
			if err == nil && dept != nil {
				deptName = dept.Name
			}
		}

		// Convert to output format
		contacts := make([]api.OutputContact, len(allUsers))
		for i, u := range allUsers {
			contacts[i] = api.OutputContact{
				UserID:     u.OpenID,
				OpenID:     u.OpenID,
				Name:       u.Name,
				EnName:     u.EnName,
				Email:      u.Email,
				JobTitle:   u.JobTitle,
				Department: deptName,
			}
		}

		result := api.OutputContactList{
			Contacts: contacts,
			Count:    len(contacts),
		}

		output.JSON(result)
	},
}

// --- contact search ---

var contactSearchCmd = &cobra.Command{
	Use:   "search <query>",
	Short: "Search for users by name",
	Long: `Search for users by name keyword.

Examples:
  lark contact search "Zheng Peng"
  lark contact search "Bryan"`,
	Args: cobra.ExactArgs(1),
	PreRun: func(cmd *cobra.Command, args []string) {
		validateScopeGroup("contacts")
	},
	Run: func(cmd *cobra.Command, args []string) {
		query := args[0]

		client := api.NewClient()

		// Fetch all matching users (paginated)
		var allUsers []api.SearchUserResult
		var pageToken string
		hasMore := true

		for hasMore {
			users, more, nextToken, err := client.SearchUsers(query, 50, pageToken)
			if err != nil {
				output.Fatal("API_ERROR", err)
			}
			allUsers = append(allUsers, users...)
			hasMore = more
			pageToken = nextToken
		}

		// Convert to output format
		contacts := make([]api.OutputContact, len(allUsers))
		for i, u := range allUsers {
			// Get department name for first department
			var deptName string
			if len(u.DepartmentIDs) > 0 {
				dept, err := client.GetDepartment(u.DepartmentIDs[0])
				if err == nil && dept != nil {
					deptName = dept.Name
				}
			}

			contacts[i] = api.OutputContact{
				UserID:     u.OpenID,
				OpenID:     u.OpenID,
				Name:       u.Name,
				Department: deptName,
			}
		}

		result := api.OutputContactList{
			Contacts: contacts,
			Count:    len(contacts),
		}

		output.JSON(result)
	},
}

// --- contact search-dept ---

var contactSearchDeptCmd = &cobra.Command{
	Use:   "search-dept <query>",
	Short: "Search for departments by name",
	Long: `Search for departments by name keyword.

Examples:
  lark contact search-dept "Engineering"
  lark contact search-dept "Business"`,
	Args: cobra.ExactArgs(1),
	PreRun: func(cmd *cobra.Command, args []string) {
		validateScopeGroup("contacts")
	},
	Run: func(cmd *cobra.Command, args []string) {
		query := args[0]

		client := api.NewClient()

		// Fetch all matching departments (paginated)
		var allDepts []api.Department
		var pageToken string
		hasMore := true

		for hasMore {
			depts, more, nextToken, err := client.SearchDepartments(query, 50, pageToken)
			if err != nil {
				output.Fatal("API_ERROR", err)
			}
			allDepts = append(allDepts, depts...)
			hasMore = more
			pageToken = nextToken
		}

		// Convert to output format
		outputDepts := make([]api.OutputDepartment, len(allDepts))
		for i, d := range allDepts {
			outputDepts[i] = api.OutputDepartment{
				DepartmentID: d.OpenDepartmentID,
				Name:         d.Name,
				MemberCount:  d.MemberCount,
			}
		}

		result := api.OutputDepartmentList{
			Departments: outputDepts,
			Count:       len(outputDepts),
		}

		output.JSON(result)
	},
}

func init() {
	// contact get flags
	contactGetCmd.Flags().StringVar(&contactGetIDType, "id-type", "open_id", "Type of user ID (open_id, union_id, user_id)")

	// contact list-dept flags
	contactListDeptCmd.Flags().IntVar(&contactListDeptPageSize, "page-size", 50, "Number of results per page (max 50)")

	// Register subcommands
	contactCmd.AddCommand(contactGetCmd)
	contactCmd.AddCommand(contactListDeptCmd)
	contactCmd.AddCommand(contactSearchCmd)
	contactCmd.AddCommand(contactSearchDeptCmd)
}
