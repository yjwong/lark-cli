package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/yjwong/lark-cli/internal/api"
	"github.com/yjwong/lark-cli/internal/output"
)

var attendeeCmd = &cobra.Command{
	Use:   "attendee",
	Short: "Manage event attendees",
	Long:  "Add, remove, or list attendees for calendar events",
}

// --- Add Attendee ---

var (
	addAttendeeEmails   []string
	addAttendeeUsers    []string
	addAttendeeSelf     bool
	addAttendeeOptional bool
	addAttendeeNoNotify bool
)

var attendeeAddCmd = &cobra.Command{
	Use:   "add <event-id>",
	Short: "Add attendees to an event",
	Long: `Add one or more attendees to an existing calendar event.

Examples:
  lark cal attendee add <event-id> --self
  lark cal attendee add <event-id> --email user@example.com
  lark cal attendee add <event-id> --user ou_xxxxxxxx
  lark cal attendee add <event-id> --self --email user@example.com --optional`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		eventID := args[0]

		if !addAttendeeSelf && len(addAttendeeEmails) == 0 && len(addAttendeeUsers) == 0 {
			output.Fatalf("VALIDATION_ERROR", "at least one of --self, --email, or --user is required")
		}

		client := api.NewClient()

		// Get primary calendar
		cal, err := client.GetPrimaryCalendar()
		if err != nil {
			output.Fatal("CALENDAR_ERROR", err)
		}

		var attendees []api.Attendee

		// Add self if requested
		if addAttendeeSelf {
			currentUser, err := client.GetCurrentUser()
			if err != nil {
				output.Fatalf("USER_ERROR", "Failed to get current user: %v", err)
			}
			attendees = append(attendees, api.Attendee{
				Type:       "user",
				UserID:     currentUser.OpenID,
				IsOptional: addAttendeeOptional,
			})
		}

		// Add email attendees (auto-resolve to internal Lark users)
		var trimmedEmails []string
		for _, email := range addAttendeeEmails {
			email = strings.TrimSpace(email)
			if email != "" {
				trimmedEmails = append(trimmedEmails, email)
			}
		}
		if len(trimmedEmails) > 0 {
			resolved := resolveEmails(client, trimmedEmails)
			for _, email := range trimmedEmails {
				if userID, ok := resolved[email]; ok {
					attendees = append(attendees, api.Attendee{
						Type:       "user",
						UserID:     userID,
						IsOptional: addAttendeeOptional,
					})
				} else {
					attendees = append(attendees, api.Attendee{
						Type:            "third_party",
						ThirdPartyEmail: email,
						IsOptional:      addAttendeeOptional,
					})
				}
			}
		}

		// Add user attendees by open_id
		for _, userID := range addAttendeeUsers {
			userID = strings.TrimSpace(userID)
			if userID == "" {
				continue
			}
			attendees = append(attendees, api.Attendee{
				Type:       "user",
				UserID:     userID,
				IsOptional: addAttendeeOptional,
			})
		}

		if len(attendees) == 0 {
			output.Fatalf("VALIDATION_ERROR", "no valid attendees to add")
		}

		// Add attendees to event
		notify := !addAttendeeNoNotify
		addedAttendees, err := client.CreateEventAttendees(cal.CalendarID, eventID, attendees, notify)
		if err != nil {
			output.Fatal("API_ERROR", err)
		}

		// Output result
		var names []string
		for _, att := range addedAttendees {
			if att.DisplayName != "" {
				names = append(names, att.DisplayName)
			} else if att.ThirdPartyEmail != "" {
				names = append(names, att.ThirdPartyEmail)
			}
		}

		output.JSON(map[string]interface{}{
			"success":   true,
			"message":   fmt.Sprintf("Added %d attendee(s) to event", len(addedAttendees)),
			"attendees": names,
		})
	},
}

// --- Remove Attendee ---

var (
	removeAttendeeIDs     []string
	removeAttendeeSelf    bool
	removeAttendeeNoNotify bool
)

var attendeeRemoveCmd = &cobra.Command{
	Use:   "remove <event-id>",
	Short: "Remove attendees from an event",
	Long: `Remove one or more attendees from an existing calendar event.

Examples:
  lark cal attendee remove <event-id> --self
  lark cal attendee remove <event-id> --id user_xxxxx
  lark cal attendee remove <event-id> --id user_xxxxx --id third_party_xxxxx`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		eventID := args[0]

		client := api.NewClient()

		// Get primary calendar
		cal, err := client.GetPrimaryCalendar()
		if err != nil {
			output.Fatal("CALENDAR_ERROR", err)
		}

		var attendeeIDs []string

		// If --self, find our own attendee ID first
		if removeAttendeeSelf {
			currentUser, err := client.GetCurrentUser()
			if err != nil {
				output.Fatalf("USER_ERROR", "Failed to get current user: %v", err)
			}

			// List attendees to find our attendee_id
			attendees, err := client.ListEventAttendees(cal.CalendarID, eventID)
			if err != nil {
				output.Fatal("API_ERROR", err)
			}

			found := false
			for _, att := range attendees {
				if att.UserID == currentUser.OpenID {
					attendeeIDs = append(attendeeIDs, att.AttendeeID)
					found = true
					break
				}
			}

			if !found {
				output.Fatalf("NOT_FOUND", "You are not an attendee of this event")
			}
		}

		// Add explicit IDs
		attendeeIDs = append(attendeeIDs, removeAttendeeIDs...)

		if len(attendeeIDs) == 0 {
			output.Fatalf("VALIDATION_ERROR", "at least one of --self or --id is required")
		}

		// Remove attendees
		notify := !removeAttendeeNoNotify
		err = client.DeleteEventAttendees(cal.CalendarID, eventID, attendeeIDs, notify)
		if err != nil {
			output.Fatal("API_ERROR", err)
		}

		output.JSON(map[string]interface{}{
			"success": true,
			"message": fmt.Sprintf("Removed %d attendee(s) from event", len(attendeeIDs)),
		})
	},
}

// --- List Attendees ---

var attendeeListCmd = &cobra.Command{
	Use:   "list <event-id>",
	Short: "List attendees of an event",
	Long: `List all attendees of a calendar event.

Examples:
  lark cal attendee list <event-id>`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		eventID := args[0]

		client := api.NewClient()

		// Get primary calendar
		cal, err := client.GetPrimaryCalendar()
		if err != nil {
			output.Fatal("CALENDAR_ERROR", err)
		}

		// List attendees
		attendees, err := client.ListEventAttendees(cal.CalendarID, eventID)
		if err != nil {
			output.Fatal("API_ERROR", err)
		}

		// Convert to output format
		var outAttendees []map[string]interface{}
		for _, att := range attendees {
			outAtt := map[string]interface{}{
				"id":          att.AttendeeID,
				"name":        att.DisplayName,
				"type":        att.Type,
				"rsvp_status": att.RsvpStatus,
			}
			if att.IsOrganizer {
				outAtt["is_organizer"] = true
			}
			if att.IsOptional {
				outAtt["is_optional"] = true
			}
			if att.IsExternal {
				outAtt["is_external"] = true
			}
			if att.ThirdPartyEmail != "" {
				outAtt["email"] = att.ThirdPartyEmail
			}
			outAttendees = append(outAttendees, outAtt)
		}

		output.JSON(map[string]interface{}{
			"attendees": outAttendees,
			"count":     len(outAttendees),
		})
	},
}

func init() {
	// Add command flags
	attendeeAddCmd.Flags().StringSliceVar(&addAttendeeEmails, "email", []string{}, "Add attendee by email (repeatable)")
	attendeeAddCmd.Flags().StringSliceVar(&addAttendeeUsers, "user", []string{}, "Add attendee by Lark user ID/open_id (repeatable)")
	attendeeAddCmd.Flags().BoolVar(&addAttendeeSelf, "self", false, "Add yourself as an attendee")
	attendeeAddCmd.Flags().BoolVar(&addAttendeeOptional, "optional", false, "Mark attendee(s) as optional")
	attendeeAddCmd.Flags().BoolVar(&addAttendeeNoNotify, "no-notify", false, "Don't send notifications")

	// Remove command flags
	attendeeRemoveCmd.Flags().StringSliceVar(&removeAttendeeIDs, "id", []string{}, "Attendee ID to remove (repeatable)")
	attendeeRemoveCmd.Flags().BoolVar(&removeAttendeeSelf, "self", false, "Remove yourself from the event")
	attendeeRemoveCmd.Flags().BoolVar(&removeAttendeeNoNotify, "no-notify", false, "Don't send notifications")

	// Register subcommands
	attendeeCmd.AddCommand(attendeeAddCmd)
	attendeeCmd.AddCommand(attendeeRemoveCmd)
	attendeeCmd.AddCommand(attendeeListCmd)
}
