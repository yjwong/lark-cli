package api

import (
	"time"
)

// FreebusyOptions configures a freebusy query
type FreebusyOptions struct {
	StartTime time.Time
	EndTime   time.Time
	UserID    string // open_id of user to check
	RoomID    string // room_id of meeting room to check
}

// GetFreebusy queries availability for a user or meeting room
func (c *Client) GetFreebusy(opts FreebusyOptions) ([]FreebusyPeriod, error) {
	req := FreebusyRequest{
		TimeMin:  opts.StartTime.Format(time.RFC3339),
		TimeMax:  opts.EndTime.Format(time.RFC3339),
		UserID:   opts.UserID,
		RoomID:   opts.RoomID,
		OnlyBusy: true,
	}

	var resp FreebusyResponse
	if err := c.Post("/calendar/v4/freebusy/list", req, &resp); err != nil {
		return nil, err
	}

	if err := resp.Err(); err != nil {
		return nil, err
	}

	return resp.Data.FreebusyList, nil
}
