package cmd

import (
	"time"

	"github.com/spf13/cobra"
	"github.com/yjwong/lark-cli/internal/api"
	"github.com/yjwong/lark-cli/internal/config"
	"github.com/yjwong/lark-cli/internal/output"
	timex "github.com/yjwong/lark-cli/internal/time"
)

var (
	searchFrom string
	searchTo   string
)

var searchCmd = &cobra.Command{
	Use:   "search <query>",
	Short: "Search for events",
	Long: `Search for events by keyword.

Examples:
  lark cal search "standup"
  lark cal search "1:1" --from 2026-01-01 --to 2026-01-31`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		query := args[0]
		client := api.NewClient()

		// Get primary calendar
		cal, err := client.GetPrimaryCalendar()
		if err != nil {
			output.Fatal("CALENDAR_ERROR", err)
		}

		// Parse timezone
		loc := config.LoadTimezone()

		// Parse time range
		var startTime, endTime time.Time
		now := time.Now().In(loc)

		if searchFrom != "" {
			startTime, err = timex.Parse(searchFrom, loc)
			if err != nil {
				output.Fatalf("PARSE_ERROR", "Failed to parse --from: %v", err)
			}
		} else {
			// Default: 30 days ago
			startTime = now.AddDate(0, 0, -30)
		}

		if searchTo != "" {
			endTime, err = timex.Parse(searchTo, loc)
			if err != nil {
				output.Fatalf("PARSE_ERROR", "Failed to parse --to: %v", err)
			}
		} else {
			// Default: 30 days ahead
			endTime = now.AddDate(0, 0, 30)
		}

		// Search events
		events, err := client.SearchEvents(cal.CalendarID, query, startTime, endTime)
		if err != nil {
			output.Fatal("API_ERROR", err)
		}

		// Convert to output format
		outputEvents := api.ConvertToOutputEvents(events)
		output.JSON(api.OutputEventList{
			Events: outputEvents,
			Count:  len(outputEvents),
		})
	},
}

func init() {
	searchCmd.Flags().StringVar(&searchFrom, "from", "", "Start date for search range")
	searchCmd.Flags().StringVar(&searchTo, "to", "", "End date for search range")
}
