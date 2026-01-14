package cmd

import (
	"fmt"
	"strconv"
	"time"

	"github.com/spf13/cobra"
	"github.com/yjwong/lark-cli/internal/api"
	"github.com/yjwong/lark-cli/internal/config"
	"github.com/yjwong/lark-cli/internal/output"
	timex "github.com/yjwong/lark-cli/internal/time"
)

var (
	updateSummary         string
	updateDescription     string
	updateStart           string
	updateEnd             string
	updateLocation        string
	updateColor           string
	updateVisibility      string
	updateAttendeeAbility string
	updateNoNotify        bool
)

var updateCmd = &cobra.Command{
	Use:   "update <event-id>",
	Short: "Update an existing event",
	Long: `Update an existing calendar event.

Only specified fields will be updated.

Examples:
  lark cal update abc123 --summary "New title"
  lark cal update abc123 --start "2026-01-03T10:00:00+08:00"
  lark cal update abc123 --location "New location"
  lark cal update abc123 --color "#9CA2A9"
  lark cal update abc123 --visibility public`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		eventID := args[0]
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

		// Build update request with only provided fields
		req := &api.UpdateEventRequest{}

		if updateSummary != "" {
			req.Summary = updateSummary
		}
		if updateDescription != "" {
			req.Description = updateDescription
		}

		// Handle start/end time updates
		// Per Lark API docs: start_time and end_time must both be provided for time changes to take effect
		if updateStart != "" || updateEnd != "" {
			existingEvent, err := client.GetEvent(cal.CalendarID, eventID)
			if err != nil {
				output.Fatalf("API_ERROR", "Failed to fetch existing event: %v", err)
			}

			var currentStart, currentEnd time.Time
			if existingEvent.StartTime != nil && existingEvent.StartTime.Timestamp != "" {
				ts, _ := strconv.ParseInt(existingEvent.StartTime.Timestamp, 10, 64)
				currentStart = time.Unix(ts, 0).In(loc)
			}
			if existingEvent.EndTime != nil && existingEvent.EndTime.Timestamp != "" {
				ts, _ := strconv.ParseInt(existingEvent.EndTime.Timestamp, 10, 64)
				currentEnd = time.Unix(ts, 0).In(loc)
			}

			duration := currentEnd.Sub(currentStart)

			var newStart, newEnd time.Time

			if updateStart != "" {
				newStart, err = timex.Parse(updateStart, loc)
				if err != nil {
					output.Fatalf("PARSE_ERROR", "Failed to parse --start: %v", err)
				}
				if updateEnd != "" {
					newEnd, err = timex.Parse(updateEnd, loc)
					if err != nil {
						output.Fatalf("PARSE_ERROR", "Failed to parse --end: %v", err)
					}
				} else {
					// Calculate end time based on original duration
					newEnd = newStart.Add(duration)
				}
			} else if updateEnd != "" {
				newEnd, err = timex.Parse(updateEnd, loc)
				if err != nil {
					output.Fatalf("PARSE_ERROR", "Failed to parse --end: %v", err)
				}
				// Keep the original start time
				newStart = currentStart
			}

			req.StartTime = &api.TimeInfo{
				Timestamp: strconv.FormatInt(newStart.Unix(), 10),
				Timezone:  tz,
			}
			req.EndTime = &api.TimeInfo{
				Timestamp: strconv.FormatInt(newEnd.Unix(), 10),
				Timezone:  tz,
			}
		}

		if updateLocation != "" {
			req.Location = &api.Location{Name: updateLocation}
		}
		if updateColor != "" {
			color, err := parseHexColor(updateColor)
			if err != nil {
				output.Fatalf("VALIDATION_ERROR", "Invalid color format: %v (use hex like #9CA2A9)", err)
			}
			req.Color = &color
		}
		if updateNoNotify {
			noNotify := false
			req.NeedNotify = &noNotify
		}
		if updateVisibility != "" {
			switch updateVisibility {
			case "default", "public", "private":
				req.Visibility = updateVisibility
			default:
				output.Fatalf("VALIDATION_ERROR", "Invalid visibility: %s (use default, public, or private)", updateVisibility)
			}
		}
		if updateAttendeeAbility != "" {
			switch updateAttendeeAbility {
			case "none", "can_see_others", "can_invite_others", "can_modify_event":
				req.AttendeeAbility = updateAttendeeAbility
			default:
				output.Fatalf("VALIDATION_ERROR", "Invalid attendee-ability: %s (must be none, can_see_others, can_invite_others, or can_modify_event)", updateAttendeeAbility)
			}
		}

		// Update event
		event, err := client.UpdateEvent(cal.CalendarID, eventID, req)
		if err != nil {
			output.Fatal("API_ERROR", err)
		}

		// Output updated event
		outputEvent := api.ConvertToOutputEvent(*event)
		output.JSON(map[string]interface{}{
			"success": true,
			"message": fmt.Sprintf("Event updated: %s", event.EventID),
			"event":   outputEvent,
		})
	},
}

func init() {
	updateCmd.Flags().StringVar(&updateSummary, "summary", "", "Event title")
	updateCmd.Flags().StringVar(&updateDescription, "description", "", "Event description")
	updateCmd.Flags().StringVar(&updateStart, "start", "", "Start time")
	updateCmd.Flags().StringVar(&updateEnd, "end", "", "End time")
	updateCmd.Flags().StringVar(&updateLocation, "location", "", "Event location")
	updateCmd.Flags().StringVar(&updateColor, "color", "", "Event color (hex format, e.g., #9CA2A9)")
	updateCmd.Flags().StringVar(&updateVisibility, "visibility", "", "Event visibility (default, public, or private)")
	updateCmd.Flags().StringVar(&updateAttendeeAbility, "attendee-ability", "", "Guest permissions (none, can_see_others, can_invite_others, can_modify_event)")
	updateCmd.Flags().BoolVar(&updateNoNotify, "no-notify", false, "Don't send notifications")
}
