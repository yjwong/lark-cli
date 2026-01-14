package cmd

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"github.com/yjwong/lark-cli/internal/api"
	"github.com/yjwong/lark-cli/internal/config"
	"github.com/yjwong/lark-cli/internal/output"
	timex "github.com/yjwong/lark-cli/internal/time"
)

var (
	createSummary         string
	createDescription     string
	createStart           string
	createEnd             string
	createDuration        string
	createLocation        string
	createColor           string
	createReminder        int
	createNoNotify        bool
	createAttendees       []string
	createVisibility      string
	createAttendeeAbility string
	createExcludeSelf     bool
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new event",
	Long: `Create a new calendar event.

Examples:
  lark cal create --summary "Team standup" --start 2026-01-03T09:00:00+08:00 --duration 30m
  lark cal create --summary "1:1 with John" --start 2026-01-03T14:00:00+08:00 --duration 30m --attendee john@example.com
  lark cal create --summary "Focus Time" --start 2026-01-03T14:00:00+08:00 --duration 2h --color "#9CA2A9"`,
	Run: func(cmd *cobra.Command, args []string) {
		if createSummary == "" {
			output.Fatalf("VALIDATION_ERROR", "--summary is required")
		}
		if createStart == "" {
			output.Fatalf("VALIDATION_ERROR", "--start is required")
		}
		if createEnd == "" && createDuration == "" {
			output.Fatalf("VALIDATION_ERROR", "either --end or --duration is required")
		}

		client := api.NewClient()

		// Get primary calendar
		cal, err := client.GetPrimaryCalendar()
		if err != nil {
			output.Fatal("CALENDAR_ERROR", err)
		}

		// Parse timezone
		tz := config.GetTimezone()
		loc, err := time.LoadLocation(tz)
		if err != nil {
			loc = time.Local
		}

		// Parse start time
		startTime, err := timex.Parse(createStart, loc)
		if err != nil {
			output.Fatalf("PARSE_ERROR", "Failed to parse --start: %v", err)
		}

		// Parse end time or calculate from duration
		var endTime time.Time
		if createEnd != "" {
			endTime, err = timex.Parse(createEnd, loc)
			if err != nil {
				output.Fatalf("PARSE_ERROR", "Failed to parse --end: %v", err)
			}
		} else {
			duration, err := timex.ParseDuration(createDuration)
			if err != nil {
				output.Fatalf("PARSE_ERROR", "Failed to parse --duration: %v", err)
			}
			endTime = startTime.Add(duration)
		}

		// Build request
		req := &api.CreateEventRequest{
			Summary:     createSummary,
			Description: createDescription,
			StartTime: &api.TimeInfo{
				Timestamp: strconv.FormatInt(startTime.Unix(), 10),
				Timezone:  tz,
			},
			EndTime: &api.TimeInfo{
				Timestamp: strconv.FormatInt(endTime.Unix(), 10),
				Timezone:  tz,
			},
		}

		if createLocation != "" {
			req.Location = &api.Location{Name: createLocation}
		}

		if createColor != "" {
			color, err := parseHexColor(createColor)
			if err != nil {
				output.Fatalf("VALIDATION_ERROR", "Invalid color format: %v (use hex like #9CA2A9)", err)
			}
			req.Color = color
		}

		if createReminder > 0 {
			req.Reminders = []api.Reminder{{Minutes: createReminder}}
		} else {
			// Use default reminder from config
			defaultReminder := config.Get().Defaults.ReminderMinutes
			if defaultReminder > 0 {
				req.Reminders = []api.Reminder{{Minutes: defaultReminder}}
			}
		}

		if createNoNotify {
			noNotify := false
			req.NeedNotify = &noNotify
		}

		if createVisibility != "" {
			switch createVisibility {
			case "default", "public", "private":
				req.Visibility = createVisibility
			default:
				output.Fatalf("VALIDATION_ERROR", "Invalid visibility: %s (must be default, public, or private)", createVisibility)
			}
		}

		// Default to can_invite_others to match Lark UI behavior
		attendeeAbility := createAttendeeAbility
		if attendeeAbility == "" {
			attendeeAbility = "can_invite_others"
		}
		switch attendeeAbility {
		case "none", "can_see_others", "can_invite_others", "can_modify_event":
			req.AttendeeAbility = attendeeAbility
		default:
			output.Fatalf("VALIDATION_ERROR", "Invalid attendee-ability: %s (must be none, can_see_others, can_invite_others, or can_modify_event)", attendeeAbility)
		}

		// Create event
		event, err := client.CreateEvent(cal.CalendarID, req)
		if err != nil {
			output.Fatal("API_ERROR", err)
		}

		// Build attendee list
		var attendees []api.Attendee

		// Add current user as attendee by default (unless --exclude-self)
		if !createExcludeSelf {
			currentUser, err := client.GetCurrentUser()
			if err != nil {
				output.Fatalf("USER_ERROR", "Failed to get current user: %v", err)
			}
			attendees = append(attendees, api.Attendee{
				Type:   "user",
				UserID: currentUser.OpenID,
			})
		}

		// Add explicitly specified attendees
		if len(createAttendees) > 0 {
			parsedAttendees, err := parseAttendees(createAttendees)
			if err != nil {
				output.Fatalf("ATTENDEE_ERROR", "Failed to parse attendees: %v", err)
			}
			attendees = append(attendees, parsedAttendees...)
		}

		// Add all attendees to the event
		if len(attendees) > 0 {
			notify := !createNoNotify
			addedAttendees, err := client.CreateEventAttendees(cal.CalendarID, event.EventID, attendees, notify)
			if err != nil {
				output.Fatalf("ATTENDEE_ERROR", "Failed to add attendees: %v", err)
			}
			event.Attendees = addedAttendees
		}

		// Output created event
		outputEvent := api.ConvertToOutputEvent(*event)
		output.JSON(map[string]interface{}{
			"success": true,
			"message": fmt.Sprintf("Event created: %s", event.EventID),
			"event":   outputEvent,
		})
	},
}

func init() {
	createCmd.Flags().StringVar(&createSummary, "summary", "", "Event title (required)")
	createCmd.Flags().StringVar(&createDescription, "description", "", "Event description")
	createCmd.Flags().StringVar(&createStart, "start", "", "Start time (required)")
	createCmd.Flags().StringVar(&createEnd, "end", "", "End time")
	createCmd.Flags().StringVar(&createDuration, "duration", "", "Duration (e.g., 30m, 1h)")
	createCmd.Flags().StringVar(&createLocation, "location", "", "Event location")
	createCmd.Flags().StringVar(&createColor, "color", "", "Event color (hex format, e.g., #9CA2A9)")
	createCmd.Flags().IntVar(&createReminder, "reminder", 0, "Reminder minutes before event")
	createCmd.Flags().BoolVar(&createNoNotify, "no-notify", false, "Don't send notifications")
	createCmd.Flags().StringSliceVar(&createAttendees, "attendee", []string{}, "Add attendee by email (repeatable)")
	createCmd.Flags().StringVar(&createVisibility, "visibility", "", "Event visibility (default, public, private)")
	createCmd.Flags().StringVar(&createAttendeeAbility, "attendee-ability", "", "Guest permissions (none, can_see_others, can_invite_others, can_modify_event)")
	createCmd.Flags().BoolVar(&createExcludeSelf, "exclude-self", false, "Don't add yourself as an attendee")

	createCmd.MarkFlagRequired("summary")
	createCmd.MarkFlagRequired("start")
}

// parseAttendees converts attendee strings to Attendee structs
func parseAttendees(attendeeStrs []string) ([]api.Attendee, error) {
	var attendees []api.Attendee
	for _, s := range attendeeStrs {
		// Format: email:user@example.com or just user@example.com
		if strings.HasPrefix(s, "email:") {
			email := strings.TrimPrefix(s, "email:")
			attendees = append(attendees, api.Attendee{
				Type:            "third_party",
				ThirdPartyEmail: email,
			})
		} else if strings.Contains(s, "@") {
			// Bare email address
			attendees = append(attendees, api.Attendee{
				Type:            "third_party",
				ThirdPartyEmail: s,
			})
		} else {
			return nil, fmt.Errorf("unknown attendee format: %s (use email address like user@example.com)", s)
		}
	}
	return attendees, nil
}

// parseHexColor converts a hex color string (e.g., "#9CA2A9") to an int32 RGB value for Lark API
func parseHexColor(hex string) (int, error) {
	// Remove # prefix if present
	hex = strings.TrimPrefix(hex, "#")

	// Parse as hex number
	val, err := strconv.ParseInt(hex, 16, 32)
	if err != nil {
		return 0, fmt.Errorf("invalid hex color: %s", hex)
	}

	return int(val), nil
}
