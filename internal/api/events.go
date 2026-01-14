package api

import (
	"fmt"
	"net/url"
	"strconv"
	"time"

	"github.com/yjwong/lark-cli/internal/config"
)

// ListEventsOptions contains options for listing events
type ListEventsOptions struct {
	CalendarID string
	StartTime  time.Time
	EndTime    time.Time
	PageSize   int
}

// maxInstanceViewDays is the maximum time range for the instance_view API (40 days)
const maxInstanceViewDays = 40

// ListEvents retrieves events from a calendar using the instance_view API.
// This API automatically expands recurring events into individual instances.
func (c *Client) ListEvents(opts ListEventsOptions) ([]Event, error) {
	if opts.CalendarID == "" {
		return nil, fmt.Errorf("calendar ID is required")
	}

	if opts.StartTime.IsZero() || opts.EndTime.IsZero() {
		return nil, fmt.Errorf("start_time and end_time are required")
	}

	// The instance_view API has a 40-day limit. If the range is longer,
	// we need to make multiple requests.
	var allItems []InstanceViewItem
	chunkStart := opts.StartTime

	for chunkStart.Before(opts.EndTime) {
		chunkEnd := chunkStart.AddDate(0, 0, maxInstanceViewDays)
		if chunkEnd.After(opts.EndTime) {
			chunkEnd = opts.EndTime
		}

		items, err := c.getInstanceView(opts.CalendarID, chunkStart, chunkEnd)
		if err != nil {
			return nil, err
		}
		allItems = append(allItems, items...)

		chunkStart = chunkEnd
	}

	// Convert InstanceViewItems to Events and deduplicate
	var events []Event
	seen := make(map[string]bool)

	for _, item := range allItems {
		// Skip cancelled events
		if item.Status == "cancelled" {
			continue
		}

		// Skip duplicates (can happen at chunk boundaries)
		if seen[item.EventID] {
			continue
		}
		seen[item.EventID] = true

		events = append(events, Event{
			EventID:             item.EventID,
			OrganizerCalendarID: item.OrganizerCalendarID,
			Summary:             item.Summary,
			Description:         item.Description,
			StartTime:           item.StartTime,
			EndTime:             item.EndTime,
			Vchat:               item.Vchat,
			Visibility:          item.Visibility,
			AttendeeAbility:     item.AttendeeAbility,
			FreeBusyStatus:      item.FreeBusyStatus,
			Location:            item.Location,
			Color:               item.Color,
			Status:              item.Status,
			IsException:         item.IsException,
			RecurringEventID:    item.RecurringEventID,
			Attendees:           item.Attendees,
		})
	}

	return events, nil
}

// getInstanceView fetches event instances for a time range (max 40 days)
func (c *Client) getInstanceView(calendarID string, startTime, endTime time.Time) ([]InstanceViewItem, error) {
	params := url.Values{}
	params.Set("start_time", strconv.FormatInt(startTime.Unix(), 10))
	params.Set("end_time", strconv.FormatInt(endTime.Unix(), 10))

	path := fmt.Sprintf("/calendar/v4/calendars/%s/events/instance_view?%s", calendarID, params.Encode())

	var resp InstanceViewResponse
	if err := c.Get(path, &resp); err != nil {
		return nil, err
	}

	if resp.Code != 0 {
		return nil, fmt.Errorf("API error (code %d): %s", resp.Code, resp.Msg)
	}

	return resp.Data.Items, nil
}

// ListEventAttendees retrieves all attendees for an event
func (c *Client) ListEventAttendees(calendarID, eventID string) ([]Attendee, error) {
	var allAttendees []Attendee
	var pageToken string

	for {
		params := url.Values{}
		if pageToken != "" {
			params.Set("page_token", pageToken)
		}

		path := fmt.Sprintf("/calendar/v4/calendars/%s/events/%s/attendees", calendarID, eventID)
		if len(params) > 0 {
			path += "?" + params.Encode()
		}

		var resp AttendeeListResponse
		if err := c.Get(path, &resp); err != nil {
			return nil, err
		}

		if resp.Code != 0 {
			return nil, fmt.Errorf("API error (code %d): %s", resp.Code, resp.Msg)
		}

		allAttendees = append(allAttendees, resp.Data.Items...)

		if !resp.Data.HasMore {
			break
		}
		pageToken = resp.Data.PageToken
	}

	return allAttendees, nil
}

// CreateEventAttendees adds attendees to an existing event
func (c *Client) CreateEventAttendees(calendarID, eventID string, attendees []Attendee, notify bool) ([]Attendee, error) {
	reqBody := map[string]interface{}{
		"attendees":         attendees,
		"need_notification": notify,
	}

	path := fmt.Sprintf("/calendar/v4/calendars/%s/events/%s/attendees?user_id_type=open_id", calendarID, eventID)

	var resp CreateAttendeeResponse
	if err := c.Post(path, reqBody, &resp); err != nil {
		return nil, err
	}

	if resp.Code != 0 {
		return nil, fmt.Errorf("API error (code %d): %s", resp.Code, resp.Msg)
	}

	return resp.Data.Attendees, nil
}

// GetEvent retrieves a single event by ID, including attendees
func (c *Client) GetEvent(calendarID, eventID string) (*Event, error) {
	var resp EventResponse

	path := fmt.Sprintf("/calendar/v4/calendars/%s/events/%s?need_attendee=true", calendarID, eventID)
	if err := c.Get(path, &resp); err != nil {
		return nil, err
	}

	if resp.Code != 0 {
		return nil, fmt.Errorf("API error (code %d): %s", resp.Code, resp.Msg)
	}

	return resp.Data.Event, nil
}

// CreateEventRequest contains the data for creating an event
type CreateEventRequest struct {
	Summary         string     `json:"summary,omitempty"`
	Description     string     `json:"description,omitempty"`
	StartTime       *TimeInfo  `json:"start_time"`
	EndTime         *TimeInfo  `json:"end_time"`
	Location        *Location  `json:"location,omitempty"`
	Color           int        `json:"color,omitempty"`
	Reminders       []Reminder `json:"reminders,omitempty"`
	Recurrence      string     `json:"recurrence,omitempty"`
	Vchat           *Vchat     `json:"vchat,omitempty"`
	Visibility      string     `json:"visibility,omitempty"`
	AttendeeAbility string     `json:"attendee_ability,omitempty"`
	NeedNotify      *bool      `json:"need_notification,omitempty"`
}

// CreateEvent creates a new event
func (c *Client) CreateEvent(calendarID string, req *CreateEventRequest) (*Event, error) {
	var resp EventResponse

	path := fmt.Sprintf("/calendar/v4/calendars/%s/events", calendarID)
	if err := c.Post(path, req, &resp); err != nil {
		return nil, err
	}

	if resp.Code != 0 {
		return nil, fmt.Errorf("API error (code %d): %s", resp.Code, resp.Msg)
	}

	return resp.Data.Event, nil
}

// UpdateEventRequest contains the data for updating an event
type UpdateEventRequest struct {
	Summary         string     `json:"summary,omitempty"`
	Description     string     `json:"description,omitempty"`
	StartTime       *TimeInfo  `json:"start_time,omitempty"`
	EndTime         *TimeInfo  `json:"end_time,omitempty"`
	Location        *Location  `json:"location,omitempty"`
	Color           *int       `json:"color,omitempty"`
	Reminders       []Reminder `json:"reminders,omitempty"`
	Visibility      string     `json:"visibility,omitempty"`
	AttendeeAbility string     `json:"attendee_ability,omitempty"`
	NeedNotify      *bool      `json:"need_notification,omitempty"`
}

// UpdateEvent updates an existing event
func (c *Client) UpdateEvent(calendarID, eventID string, req *UpdateEventRequest) (*Event, error) {
	var resp EventResponse

	path := fmt.Sprintf("/calendar/v4/calendars/%s/events/%s", calendarID, eventID)
	if err := c.Patch(path, req, &resp); err != nil {
		return nil, err
	}

	if resp.Code != 0 {
		return nil, fmt.Errorf("API error (code %d): %s", resp.Code, resp.Msg)
	}

	return resp.Data.Event, nil
}

// DeleteEvent deletes an event
func (c *Client) DeleteEvent(calendarID, eventID string) error {
	var resp BaseResponse

	path := fmt.Sprintf("/calendar/v4/calendars/%s/events/%s", calendarID, eventID)
	if err := c.Delete(path, &resp); err != nil {
		return err
	}

	if resp.Code != 0 {
		return fmt.Errorf("API error (code %d): %s", resp.Code, resp.Msg)
	}

	return nil
}

// SearchEvents searches for events by query
func (c *Client) SearchEvents(calendarID, query string, startTime, endTime time.Time) ([]Event, error) {
	reqBody := map[string]interface{}{
		"query": query,
	}

	if !startTime.IsZero() {
		reqBody["filter"] = map[string]interface{}{
			"start_time": map[string]string{
				"timestamp": strconv.FormatInt(startTime.Unix(), 10),
			},
			"end_time": map[string]string{
				"timestamp": strconv.FormatInt(endTime.Unix(), 10),
			},
		}
	}

	var resp EventListResponse
	path := fmt.Sprintf("/calendar/v4/calendars/%s/events/search", calendarID)
	if err := c.Post(path, reqBody, &resp); err != nil {
		return nil, err
	}

	if resp.Code != 0 {
		return nil, fmt.Errorf("API error (code %d): %s", resp.Code, resp.Msg)
	}

	return resp.Data.Items, nil
}

// ConvertToOutputEvent converts a Lark API event to CLI output format
func ConvertToOutputEvent(e Event) OutputEvent {
	tz := config.GetTimezone()
	loc, _ := time.LoadLocation(tz)
	if loc == nil {
		loc = time.Local
	}

	out := OutputEvent{
		ID:          e.EventID,
		Summary:     e.Summary,
		Description: e.Description,
		Visibility:  e.Visibility,
		Recurrence:  e.Recurrence,
	}

	// Convert start time
	if e.StartTime != nil {
		if e.StartTime.Date != "" {
			out.Start = e.StartTime.Date
			out.AllDay = true
		} else if e.StartTime.Timestamp != "" {
			ts, _ := strconv.ParseInt(e.StartTime.Timestamp, 10, 64)
			t := time.Unix(ts, 0).In(loc)
			out.Start = t.Format(time.RFC3339)
		}
	}

	// Convert end time
	if e.EndTime != nil {
		if e.EndTime.Date != "" {
			out.End = e.EndTime.Date
		} else if e.EndTime.Timestamp != "" {
			ts, _ := strconv.ParseInt(e.EndTime.Timestamp, 10, 64)
			t := time.Unix(ts, 0).In(loc)
			out.End = t.Format(time.RFC3339)
		}
	}

	// Location
	if e.Location != nil && e.Location.Name != "" {
		out.Location = e.Location.Name
	}

	// Color (0 or -1 means use calendar color, so we omit it)
	if e.Color != 0 && e.Color != -1 {
		// Convert int32 RGB to hex string
		r := (e.Color >> 16) & 0xFF
		g := (e.Color >> 8) & 0xFF
		b := e.Color & 0xFF
		out.Color = fmt.Sprintf("#%02X%02X%02X", r, g, b)
	}

	// Meeting URL
	if e.Vchat != nil && e.Vchat.MeetingURL != "" {
		out.MeetingURL = e.Vchat.MeetingURL
	}

	// Convert attendees
	for _, att := range e.Attendees {
		outAtt := OutputAttendee{
			Name:        att.DisplayName,
			RsvpStatus:  att.RsvpStatus,
			IsOrganizer: att.IsOrganizer,
			IsOptional:  att.IsOptional,
		}
		// Only include type if not a regular user
		if att.Type != "user" {
			outAtt.Type = att.Type
		}
		// Include email for third-party attendees
		if att.Type == "third_party" && att.ThirdPartyEmail != "" {
			outAtt.Email = att.ThirdPartyEmail
		}
		out.Attendees = append(out.Attendees, outAtt)
		// Track organizer name
		if att.IsOrganizer {
			out.Organizer = att.DisplayName
		}
	}

	return out
}

// ConvertToOutputEvents converts a slice of events
func ConvertToOutputEvents(events []Event) []OutputEvent {
	out := make([]OutputEvent, len(events))
	for i, e := range events {
		out[i] = ConvertToOutputEvent(e)
	}
	return out
}

// ReplyToEvent sends an RSVP response (accept, decline, tentative) to an event invitation
func (c *Client) ReplyToEvent(calendarID, eventID, rsvpStatus string) error {
	reqBody := map[string]interface{}{
		"rsvp_status": rsvpStatus,
	}

	path := fmt.Sprintf("/calendar/v4/calendars/%s/events/%s/reply", calendarID, eventID)

	var resp BaseResponse
	if err := c.Post(path, reqBody, &resp); err != nil {
		return err
	}

	if resp.Code != 0 {
		return fmt.Errorf("API error (code %d): %s", resp.Code, resp.Msg)
	}

	return nil
}

// DeleteEventAttendees removes attendees from an existing event
func (c *Client) DeleteEventAttendees(calendarID, eventID string, attendeeIDs []string, notify bool) error {
	reqBody := map[string]interface{}{
		"attendee_ids":      attendeeIDs,
		"need_notification": notify,
	}

	path := fmt.Sprintf("/calendar/v4/calendars/%s/events/%s/attendees/batch_delete", calendarID, eventID)

	var resp BaseResponse
	if err := c.Post(path, reqBody, &resp); err != nil {
		return err
	}

	if resp.Code != 0 {
		return fmt.Errorf("API error (code %d): %s", resp.Code, resp.Msg)
	}

	return nil
}

// ExtractUserRsvpStatus finds the current user's RSVP status from an event's attendees.
// It checks both direct user attendees and chat group memberships.
func ExtractUserRsvpStatus(event Event, userOpenID, calendarID string, client *Client) string {
	for _, att := range event.Attendees {
		// Check direct user attendee
		if att.UserID == userOpenID {
			return att.RsvpStatus
		}

		// Check chat group attendees - expand to find user's individual RSVP
		if att.Type == "chat" && att.AttendeeID != "" && client != nil {
			members, err := client.ListChatMemberAttendees(calendarID, event.EventID, att.AttendeeID)
			if err == nil {
				for _, member := range members {
					if member.OpenID == userOpenID {
						return member.RsvpStatus
					}
				}
			}
		}
	}
	return ""
}

// ListChatMemberAttendees retrieves the individual member RSVP status for a chat group invitee
func (c *Client) ListChatMemberAttendees(calendarID, eventID, attendeeID string) ([]ChatMemberAttendee, error) {
	var allMembers []ChatMemberAttendee
	var pageToken string

	for {
		params := url.Values{}
		params.Set("user_id_type", "open_id")
		if pageToken != "" {
			params.Set("page_token", pageToken)
		}

		path := fmt.Sprintf("/calendar/v4/calendars/%s/events/%s/attendees/%s/chat_members?%s",
			calendarID, eventID, attendeeID, params.Encode())

		var resp ChatMemberAttendeesResponse
		if err := c.Get(path, &resp); err != nil {
			return nil, err
		}

		if resp.Code != 0 {
			return nil, fmt.Errorf("API error (code %d): %s", resp.Code, resp.Msg)
		}

		allMembers = append(allMembers, resp.Data.Items...)

		if !resp.Data.HasMore {
			break
		}
		pageToken = resp.Data.PageToken
	}

	return allMembers, nil
}
