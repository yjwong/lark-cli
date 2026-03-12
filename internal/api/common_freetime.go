package api

import (
	"time"
)

// CommonFreeTimeOptions configures a common free time query
type CommonFreeTimeOptions struct {
	UserIDs                 []string
	StartTime               time.Time
	EndTime                 time.Time
	Timezone                string
	OnlyBusy                bool
	IncludeExternalCalendar bool
	EnableWorkHour          bool
	MinTimeLength           int // seconds
	Limit                   int
}

// GetCommonFreeTime queries common free time slots for multiple users
func (c *Client) GetCommonFreeTime(opts CommonFreeTimeOptions) ([]FreeTimeSlot, error) {
	// Format time as "YYYY-MM-DD HH:MM:SS" (API requirement)
	const apiTimeFormat = "2006-01-02 15:04:05"

	req := CommonFreeTimeRequest{
		UserIDs:                 opts.UserIDs,
		StartTime:               opts.StartTime.Format(apiTimeFormat),
		EndTime:                 opts.EndTime.Format(apiTimeFormat),
		Timezone:                opts.Timezone,
		OnlyBusy:                opts.OnlyBusy,
		IncludeExternalCalendar: opts.IncludeExternalCalendar,
		EnableWorkHour:          opts.EnableWorkHour,
		MinTimeLength:           opts.MinTimeLength,
		Limit:                   opts.Limit,
	}

	var resp CommonFreeTimeResponse
	if err := c.Post("/calendar/v4/common_freetime/mget", req, &resp); err != nil {
		return nil, err
	}

	if err := resp.Err(); err != nil {
		return nil, err
	}

	return resp.Data.Items, nil
}
