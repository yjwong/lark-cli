package cmd

import (
	"time"

	"github.com/spf13/cobra"
	"github.com/yjwong/lark-cli/internal/api"
	"github.com/yjwong/lark-cli/internal/config"
	"github.com/yjwong/lark-cli/internal/conflicts"
	"github.com/yjwong/lark-cli/internal/output"
	timex "github.com/yjwong/lark-cli/internal/time"
)

var (
	listFrom            string
	listTo              string
	listToday           bool
	listWeek            bool
	listAttendees       bool
	listDetectConflicts bool
	listBufferMinutes   int
	listPending         bool
	listRsvp            bool
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List calendar events",
	Long: `List events from your Lark calendar.

By default, lists today's events. Use --from and --to for custom date ranges,
or --week for the current week.

Examples:
  lark cal list                                  # Today's events
  lark cal list --week                           # This week
  lark cal list --from 2026-01-02 --to 2026-01-05
  lark cal list --from 2026-01-05 --to 2026-01-10
  lark cal list --rsvp                           # Include your RSVP status
  lark cal list --attendees                      # Include all attendees info
  lark cal list --pending                        # Events awaiting your RSVP`,
	Run: func(cmd *cobra.Command, args []string) {
		client := api.NewClient()

		// Get primary calendar
		cal, err := client.GetPrimaryCalendar()
		if err != nil {
			output.Fatal("CALENDAR_ERROR", err)
		}

		// Determine time range
		loc := config.LoadTimezone()
		now := time.Now().In(loc)

		var startTime, endTime time.Time

		if listWeek {
			startTime = timex.StartOfWeek(now)
			endTime = timex.EndOfWeek(now)
		} else if listFrom != "" || listTo != "" {
			if listFrom != "" {
				startTime, err = timex.Parse(listFrom, loc)
				if err != nil {
					output.Fatalf("PARSE_ERROR", "Failed to parse --from: %v", err)
				}
				startTime = timex.StartOfDay(startTime)
			} else {
				startTime = timex.StartOfDay(now)
			}

			if listTo != "" {
				endTime, err = timex.Parse(listTo, loc)
				if err != nil {
					output.Fatalf("PARSE_ERROR", "Failed to parse --to: %v", err)
				}
				endTime = timex.EndOfDay(endTime)
			} else {
				endTime = timex.EndOfDay(now.AddDate(0, 1, 0)) // 1 month default
			}
		} else {
			// Default: today
			startTime = timex.StartOfDay(now)
			endTime = timex.EndOfDay(now)
		}

		// Fetch events
		events, err := client.ListEvents(api.ListEventsOptions{
			CalendarID: cal.CalendarID,
			StartTime:  startTime,
			EndTime:    endTime,
		})
		if err != nil {
			output.Fatal("API_ERROR", err)
		}

		// Get current user's open_id if we need to filter by pending RSVP or show RSVP status
		var currentUserOpenID string
		if listPending || listRsvp {
			user, err := client.GetCurrentUser()
			if err != nil {
				output.Fatal("USER_ERROR", err)
			}
			currentUserOpenID = user.OpenID
		}

		// Fetch attendees for each event if requested or if filtering by pending or showing RSVP
		if listAttendees || listPending || listRsvp {
			for i := range events {
				// Always use the instance EventID to get instance-specific RSVP status
				// (RecurringEventID returns series-level status which may be stale)
				attendees, err := client.ListEventAttendees(cal.CalendarID, events[i].EventID)
				if err == nil {
					events[i].Attendees = attendees
				}
				// Silently ignore errors - some events may not have permission to view attendees
			}
		}

		// Filter to events where current user's RSVP is needs_action
		if listPending {
			var pendingEvents []api.Event
			for _, event := range events {
				isPending := false

				for _, att := range event.Attendees {
					// Check direct user attendee
					if att.UserID == currentUserOpenID && att.RsvpStatus == "needs_action" {
						isPending = true
						break
					}

					// Check chat group attendees - expand to find user's individual RSVP
					if att.Type == "chat" && att.AttendeeID != "" {
						members, err := client.ListChatMemberAttendees(cal.CalendarID, event.EventID, att.AttendeeID)
						if err == nil {
							for _, member := range members {
								if member.OpenID == currentUserOpenID && member.RsvpStatus == "needs_action" {
									isPending = true
									break
								}
							}
						}
						if isPending {
							break
						}
					}
				}

				if isPending {
					pendingEvents = append(pendingEvents, event)
				}
			}
			events = pendingEvents
		}

		// Convert to output format
		outputEvents := api.ConvertToOutputEvents(events)

		// Populate user's RSVP status if requested (and we have attendees data)
		if listRsvp && currentUserOpenID != "" {
			for i := range outputEvents {
				outputEvents[i].RsvpStatus = api.ExtractUserRsvpStatus(events[i], currentUserOpenID, cal.CalendarID, client)
				// If --rsvp only (not --attendees), clear the attendees list to keep output clean
				if !listAttendees {
					outputEvents[i].Attendees = nil
					outputEvents[i].Organizer = ""
				}
			}
		}

		// Detect conflicts if requested
		var conflictResult conflicts.Result
		if listDetectConflicts {
			slots, err := conflicts.ParseEventTimes(outputEvents, loc)
			if err != nil {
				output.Fatal("CONFLICT_DETECTION_ERROR", err)
			}
			conflictResult = conflicts.Detect(slots, conflicts.Options{
				BufferMinutes: listBufferMinutes,
			})
			conflicts.ApplyToEvents(outputEvents, conflictResult)
		}

		output.JSON(api.OutputEventList{
			Events:       outputEvents,
			Count:        len(outputEvents),
			Conflicts:    conflictResult.Conflicts,
			HasConflicts: conflictResult.HasConflicts,
		})
	},
}

func init() {
	listCmd.Flags().StringVar(&listFrom, "from", "", "Start date (ISO 8601)")
	listCmd.Flags().StringVar(&listTo, "to", "", "End date (ISO 8601)")
	listCmd.Flags().BoolVar(&listToday, "today", false, "List today's events (default)")
	listCmd.Flags().BoolVar(&listWeek, "week", false, "List this week's events")
	listCmd.Flags().BoolVar(&listRsvp, "rsvp", false, "Include your RSVP status for each event")
	listCmd.Flags().BoolVar(&listAttendees, "attendees", false, "Include attendee information (slower)")
	listCmd.Flags().BoolVar(&listDetectConflicts, "detect-conflicts", false, "Detect overlapping events")
	listCmd.Flags().IntVar(&listBufferMinutes, "buffer-minutes", 0, "Minimum buffer between meetings (requires --detect-conflicts)")
	listCmd.Flags().BoolVar(&listPending, "pending", false, "Only show events awaiting your RSVP")
}
