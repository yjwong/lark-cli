package api

import (
	"fmt"
)

// GetPrimaryCalendar retrieves the user's primary calendar
func (c *Client) GetPrimaryCalendar() (*Calendar, error) {
	var resp CalendarResponse

	if err := c.Post("/calendar/v4/calendars/primary", nil, &resp); err != nil {
		return nil, err
	}

	if err := resp.Err(); err != nil {
		return nil, err
	}

	// Primary calendar returns an array of calendars
	if len(resp.Data.Calendars) > 0 && resp.Data.Calendars[0].Calendar != nil {
		return resp.Data.Calendars[0].Calendar, nil
	}

	// Fallback for single calendar response
	if resp.Data.Calendar != nil {
		return resp.Data.Calendar, nil
	}

	return nil, fmt.Errorf("no calendar data in response")
}

// GetCalendar retrieves a specific calendar by ID
func (c *Client) GetCalendar(calendarID string) (*Calendar, error) {
	var resp CalendarResponse

	path := fmt.Sprintf("/calendar/v4/calendars/%s", calendarID)
	if err := c.Get(path, &resp); err != nil {
		return nil, err
	}

	if err := resp.Err(); err != nil {
		return nil, err
	}

	return resp.Data.Calendar, nil
}

// ListCalendars retrieves all calendars for the user
func (c *Client) ListCalendars() ([]Calendar, error) {
	return PaginateWith(func(pageToken string) ([]Calendar, bool, string, error) {
		path := "/calendar/v4/calendars"
		if pageToken != "" {
			path += "?page_token=" + pageToken
		}

		var resp CalendarListResponse
		if err := c.Get(path, &resp); err != nil {
			return nil, false, "", err
		}
		if err := resp.Err(); err != nil {
			return nil, false, "", err
		}
		return resp.Data.Calendars, resp.Data.HasMore, resp.Data.PageToken, nil
	})
}
