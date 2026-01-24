package api

import "time"

// TimeInfo represents Lark's time structure
type TimeInfo struct {
	Date      string `json:"date,omitempty"`      // For all-day events: "2026-01-02"
	Timestamp string `json:"timestamp,omitempty"` // Unix seconds as string
	Timezone  string `json:"timezone,omitempty"`  // IANA timezone
}

// Location represents an event location
type Location struct {
	Name      string  `json:"name,omitempty"`
	Address   string  `json:"address,omitempty"`
	Latitude  float64 `json:"latitude,omitempty"`
	Longitude float64 `json:"longitude,omitempty"`
}

// Vchat represents video conference info
type Vchat struct {
	VcType      string `json:"vc_type,omitempty"`
	MeetingURL  string `json:"meeting_url,omitempty"`
	IconType    string `json:"icon_type,omitempty"`
	Description string `json:"description,omitempty"`
}

// Reminder represents an event reminder
type Reminder struct {
	Minutes int `json:"minutes"`
}

// Attendee represents an event attendee
type Attendee struct {
	Type            string `json:"type,omitempty"` // user, chat, resource, third_party
	AttendeeID      string `json:"attendee_id,omitempty"`
	UserID          string `json:"user_id,omitempty"`
	DisplayName     string `json:"display_name,omitempty"`
	RsvpStatus      string `json:"rsvp_status,omitempty"` // needs_action, accept, tentative, decline
	IsOptional      bool   `json:"is_optional,omitempty"`
	IsOrganizer     bool   `json:"is_organizer,omitempty"`
	IsExternal      bool   `json:"is_external,omitempty"`
	ThirdPartyEmail string `json:"third_party_email,omitempty"`
}

// Event represents a calendar event from Lark API
type Event struct {
	EventID             string     `json:"event_id,omitempty"`
	OrganizerCalendarID string     `json:"organizer_calendar_id,omitempty"`
	Summary             string     `json:"summary,omitempty"`
	Description         string     `json:"description,omitempty"`
	StartTime           *TimeInfo  `json:"start_time,omitempty"`
	EndTime             *TimeInfo  `json:"end_time,omitempty"`
	Vchat               *Vchat     `json:"vchat,omitempty"`
	Visibility          string     `json:"visibility,omitempty"`
	AttendeeAbility     string     `json:"attendee_ability,omitempty"`
	FreeBusyStatus      string     `json:"free_busy_status,omitempty"`
	Location            *Location  `json:"location,omitempty"`
	Color               int        `json:"color,omitempty"`
	Reminders           []Reminder `json:"reminders,omitempty"`
	Recurrence          string     `json:"recurrence,omitempty"`
	Status              string     `json:"status,omitempty"` // tentative, confirmed, cancelled
	IsException         bool       `json:"is_exception,omitempty"`
	RecurringEventID    string     `json:"recurring_event_id,omitempty"`
	CreateTime          string     `json:"create_time,omitempty"`
	Attendees           []Attendee `json:"attendees,omitempty"`
	HasMoreAttendee     bool       `json:"has_more_attendee,omitempty"`
}

// Calendar represents a Lark calendar
type Calendar struct {
	CalendarID  string `json:"calendar_id"`
	Summary     string `json:"summary,omitempty"`
	Description string `json:"description,omitempty"`
	Type        string `json:"type,omitempty"` // primary, shared, google, exchange, resource
	Color       int    `json:"color,omitempty"`
	Role        string `json:"role,omitempty"` // owner, writer, reader, free_busy_reader
}

// --- API Response Types ---

// BaseResponse is the common response structure
type BaseResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

// UserCalendar wraps calendar with user info (for primary calendar response)
type UserCalendar struct {
	Calendar *Calendar `json:"calendar,omitempty"`
	UserID   string    `json:"user_id,omitempty"`
}

// CalendarResponse is the response from calendar endpoints
type CalendarResponse struct {
	BaseResponse
	Data struct {
		Calendar  *Calendar      `json:"calendar,omitempty"`  // For single calendar get
		Calendars []UserCalendar `json:"calendars,omitempty"` // For primary calendar
	} `json:"data,omitempty"`
}

// CalendarListResponse is the response from list calendars
type CalendarListResponse struct {
	BaseResponse
	Data struct {
		HasMore   bool       `json:"has_more"`
		PageToken string     `json:"page_token,omitempty"`
		Calendars []Calendar `json:"calendars,omitempty"`
	} `json:"data,omitempty"`
}

// EventResponse is the response from single event endpoints
type EventResponse struct {
	BaseResponse
	Data struct {
		Event *Event `json:"event,omitempty"`
	} `json:"data,omitempty"`
}

// EventListResponse is the response from list events
type EventListResponse struct {
	BaseResponse
	Data struct {
		HasMore   bool    `json:"has_more"`
		PageToken string  `json:"page_token,omitempty"`
		SyncToken string  `json:"sync_token,omitempty"`
		Items     []Event `json:"items,omitempty"`
	} `json:"data,omitempty"`
}

// EventOrganizer represents the organizer info from instance_view API
type EventOrganizer struct {
	UserID      string `json:"user_id,omitempty"`
	DisplayName string `json:"display_name,omitempty"`
}

// InstanceViewItem represents an event instance from the instance_view API
// This has a slightly different structure from EventInstance
type InstanceViewItem struct {
	EventID             string          `json:"event_id"`
	Summary             string          `json:"summary,omitempty"`
	Description         string          `json:"description,omitempty"`
	StartTime           *TimeInfo       `json:"start_time,omitempty"`
	EndTime             *TimeInfo       `json:"end_time,omitempty"`
	Status              string          `json:"status,omitempty"`
	IsException         bool            `json:"is_exception,omitempty"`
	AppLink             string          `json:"app_link,omitempty"`
	OrganizerCalendarID string          `json:"organizer_calendar_id,omitempty"`
	Vchat               *Vchat          `json:"vchat,omitempty"`
	Visibility          string          `json:"visibility,omitempty"`
	AttendeeAbility     string          `json:"attendee_ability,omitempty"`
	FreeBusyStatus      string          `json:"free_busy_status,omitempty"`
	Location            *Location       `json:"location,omitempty"`
	Color               int             `json:"color,omitempty"`
	RecurringEventID    string          `json:"recurring_event_id,omitempty"`
	EventOrganizer      *EventOrganizer `json:"event_organizer,omitempty"`
	Attendees           []Attendee      `json:"attendees,omitempty"`
}

// InstanceViewResponse is the response from the instance_view API
type InstanceViewResponse struct {
	BaseResponse
	Data struct {
		Items []InstanceViewItem `json:"items,omitempty"`
	} `json:"data,omitempty"`
}

// AttendeeListResponse is the response from list event attendees
type AttendeeListResponse struct {
	BaseResponse
	Data struct {
		HasMore   bool       `json:"has_more"`
		PageToken string     `json:"page_token,omitempty"`
		Items     []Attendee `json:"items,omitempty"`
	} `json:"data,omitempty"`
}

// CreateAttendeeResponse is the response from create event attendees
type CreateAttendeeResponse struct {
	BaseResponse
	Data struct {
		Attendees []Attendee `json:"attendees,omitempty"`
	} `json:"data,omitempty"`
}

// ChatMemberAttendee represents a member of a chat group invitee
type ChatMemberAttendee struct {
	RsvpStatus  string `json:"rsvp_status,omitempty"`
	IsOptional  bool   `json:"is_optional,omitempty"`
	DisplayName string `json:"display_name,omitempty"`
	OpenID      string `json:"open_id,omitempty"`
	IsOrganizer bool   `json:"is_organizer,omitempty"`
	IsExternal  bool   `json:"is_external,omitempty"`
}

// ChatMemberAttendeesResponse is the response from list chat member attendees
type ChatMemberAttendeesResponse struct {
	BaseResponse
	Data struct {
		Items     []ChatMemberAttendee `json:"items,omitempty"`
		HasMore   bool                 `json:"has_more"`
		PageToken string               `json:"page_token,omitempty"`
	} `json:"data,omitempty"`
}

// --- CLI Output Types (JSON for Claude Code) ---

// OutputAttendee is the simplified attendee format for CLI output
type OutputAttendee struct {
	Name        string `json:"name"`
	Type        string `json:"type,omitempty"`        // user, chat, resource, third_party
	RsvpStatus  string `json:"rsvp_status,omitempty"` // needs_action, accept, tentative, decline
	IsOrganizer bool   `json:"is_organizer,omitempty"`
	IsOptional  bool   `json:"is_optional,omitempty"`
	Email       string `json:"email,omitempty"` // for third_party type
}

// OutputEvent is the simplified event format for CLI output
type OutputEvent struct {
	ID            string           `json:"id"`
	Summary       string           `json:"summary"`
	Description   string           `json:"description,omitempty"`
	Start         string           `json:"start"` // ISO 8601 with timezone
	End           string           `json:"end"`   // ISO 8601 with timezone
	AllDay        bool             `json:"all_day,omitempty"`
	Location      string           `json:"location,omitempty"`
	Color         string           `json:"color,omitempty"`      // Hex color like "#579FFF", empty if using calendar color
	Visibility    string           `json:"visibility,omitempty"` // default, public, private
	Organizer     string           `json:"organizer,omitempty"`
	Attendees     []OutputAttendee `json:"attendees,omitempty"`
	MeetingURL    string           `json:"meeting_url,omitempty"`
	Recurrence    string           `json:"recurrence,omitempty"`
	ConflictsWith []string         `json:"conflicts_with,omitempty"`
	RsvpStatus    string           `json:"rsvp_status,omitempty"` // User's RSVP status: needs_action, accept, tentative, decline
}

// Conflict represents a detected scheduling conflict
type Conflict struct {
	Type                  string   `json:"type"`                              // "overlap" or "insufficient_buffer"
	EventIDs              []string `json:"events"`                            // IDs of conflicting events
	OverlapMinutes        int      `json:"overlap_minutes,omitempty"`         // Duration of overlap
	GapMinutes            int      `json:"gap_minutes,omitempty"`             // Gap between events (for buffer)
	RequiredBufferMinutes int      `json:"required_buffer_minutes,omitempty"` // Required buffer
}

// OutputEventList is the list events response for CLI
type OutputEventList struct {
	Events       []OutputEvent `json:"events"`
	Count        int           `json:"count"`
	Conflicts    []Conflict    `json:"conflicts,omitempty"`
	HasConflicts bool          `json:"has_conflicts,omitempty"`
}

// OutputError is the error response format for CLI
type OutputError struct {
	Error   bool   `json:"error"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

// OutputAuthStatus is the auth status response for CLI
type OutputAuthStatus struct {
	Authenticated bool            `json:"authenticated"`
	User          string          `json:"user,omitempty"`
	ExpiresAt     time.Time       `json:"expires_at,omitempty"`
	RefreshAt     time.Time       `json:"refresh_token_expires_at,omitempty"`
	GrantedGroups []string        `json:"granted_groups,omitempty"`
	ScopeGroups   map[string]bool `json:"scope_groups,omitempty"`
}

// OutputSuccess is a generic success response
type OutputSuccess struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
}

// --- Freebusy Types ---

// FreebusyRequest is the request body for freebusy API
type FreebusyRequest struct {
	TimeMin                 string `json:"time_min"`
	TimeMax                 string `json:"time_max"`
	UserID                  string `json:"user_id,omitempty"`
	RoomID                  string `json:"room_id,omitempty"`
	IncludeExternalCalendar bool   `json:"include_external_calendar,omitempty"`
	OnlyBusy                bool   `json:"only_busy,omitempty"`
}

// FreebusyPeriod represents a busy time period
type FreebusyPeriod struct {
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
}

// FreebusyResponse is the response from freebusy API
type FreebusyResponse struct {
	BaseResponse
	Data struct {
		FreebusyList []FreebusyPeriod `json:"freebusy_list"`
	} `json:"data"`
}

// OutputFreebusy is the freebusy response for CLI
type OutputFreebusy struct {
	Query       OutputFreebusyQuery    `json:"query"`
	BusyPeriods []OutputFreebusyPeriod `json:"busy_periods"`
}

// OutputFreebusyQuery describes the query parameters
type OutputFreebusyQuery struct {
	From   string `json:"from"`
	To     string `json:"to"`
	UserID string `json:"user_id,omitempty"`
	RoomID string `json:"room_id,omitempty"`
}

// OutputFreebusyPeriod is a busy period for CLI output
type OutputFreebusyPeriod struct {
	Start string `json:"start"`
	End   string `json:"end"`
}

// --- User Lookup Types ---

// UserLookupRequest is the request body for user lookup API
type UserLookupRequest struct {
	Emails  []string `json:"emails,omitempty"`
	Mobiles []string `json:"mobiles,omitempty"`
}

// UserContactInfo represents a user's contact info from lookup
type UserContactInfo struct {
	UserID string `json:"user_id"`
	Email  string `json:"email,omitempty"`
	Mobile string `json:"mobile,omitempty"`
}

// UserLookupResponse is the response from user lookup API
type UserLookupResponse struct {
	BaseResponse
	Data struct {
		UserList []UserContactInfo `json:"user_list"`
	} `json:"data"`
}

// OutputUserLookup is the user lookup response for CLI
type OutputUserLookup struct {
	Users []OutputUserInfo `json:"users"`
}

// OutputUserInfo is a user info for CLI output
type OutputUserInfo struct {
	UserID string `json:"user_id"`
	Email  string `json:"email,omitempty"`
	Mobile string `json:"mobile,omitempty"`
}

// --- Common Free Time Types ---

// CommonFreeTimeRequest is the request body for common free time API
type CommonFreeTimeRequest struct {
	UserIDs                 []string `json:"user_ids"`
	StartTime               string   `json:"start_time"` // "YYYY-MM-DD HH:MM:SS"
	EndTime                 string   `json:"end_time"`   // "YYYY-MM-DD HH:MM:SS"
	Timezone                string   `json:"timezone"`
	OnlyBusy                bool     `json:"only_busy,omitempty"`
	IncludeExternalCalendar bool     `json:"include_external_calendar,omitempty"`
	EnableWorkHour          bool     `json:"enable_work_hour,omitempty"`
	MinTimeLength           int      `json:"min_time_length,omitempty"`
	Limit                   int      `json:"limit,omitempty"`
}

// FreeTimeSlot represents a free time slot from the API
type FreeTimeSlot struct {
	StartTime string `json:"start_time"` // "YYYY-MM-DD HH:MM:SS"
	EndTime   string `json:"end_time"`   // "YYYY-MM-DD HH:MM:SS"
	Length    int    `json:"length"`     // seconds
}

// CommonFreeTimeResponse is the response from common free time API
type CommonFreeTimeResponse struct {
	BaseResponse
	Data struct {
		Items []FreeTimeSlot `json:"items"`
	} `json:"data"`
}

// OutputCommonFreeTime is the common free time response for CLI
type OutputCommonFreeTime struct {
	Query     OutputCommonFreeTimeQuery `json:"query"`
	FreeSlots []OutputFreeTimeSlot      `json:"free_slots"`
}

// OutputCommonFreeTimeQuery describes the query parameters
type OutputCommonFreeTimeQuery struct {
	Users    []string `json:"users"`
	From     string   `json:"from"`
	To       string   `json:"to"`
	Timezone string   `json:"timezone"`
}

// OutputFreeTimeSlot is a free time slot for CLI output
type OutputFreeTimeSlot struct {
	Start         string `json:"start"`
	End           string `json:"end"`
	LengthMinutes int    `json:"length_minutes"`
}

// --- Contact Types ---

// ContactUserStatus represents a user's status
type ContactUserStatus struct {
	IsFrozen    bool `json:"is_frozen,omitempty"`
	IsResigned  bool `json:"is_resigned,omitempty"`
	IsActivated bool `json:"is_activated,omitempty"`
	IsExited    bool `json:"is_exited,omitempty"`
	IsUnjoin    bool `json:"is_unjoin,omitempty"`
}

// ContactUser represents a user from the Contacts API
type ContactUser struct {
	UnionID         string             `json:"union_id,omitempty"`
	UserID          string             `json:"user_id,omitempty"`
	OpenID          string             `json:"open_id,omitempty"`
	Name            string             `json:"name,omitempty"`
	EnName          string             `json:"en_name,omitempty"`
	Nickname        string             `json:"nickname,omitempty"`
	Email           string             `json:"email,omitempty"`
	Mobile          string             `json:"mobile,omitempty"`
	Gender          int                `json:"gender,omitempty"`
	Status          *ContactUserStatus `json:"status,omitempty"`
	DepartmentIDs   []string           `json:"department_ids,omitempty"`
	LeaderUserID    string             `json:"leader_user_id,omitempty"`
	City            string             `json:"city,omitempty"`
	Country         string             `json:"country,omitempty"`
	WorkStation     string             `json:"work_station,omitempty"`
	JoinTime        int64              `json:"join_time,omitempty"`
	EmployeeNo      string             `json:"employee_no,omitempty"`
	EmployeeType    int                `json:"employee_type,omitempty"`
	JobTitle        string             `json:"job_title,omitempty"`
	EnterpriseEmail string             `json:"enterprise_email,omitempty"`
}

// DepartmentI18nName represents internationalized department names
type DepartmentI18nName struct {
	ZhCn string `json:"zh_cn,omitempty"`
	JaJp string `json:"ja_jp,omitempty"`
	EnUs string `json:"en_us,omitempty"`
}

// DepartmentStatus represents a department's status
type DepartmentStatus struct {
	IsDeleted bool `json:"is_deleted,omitempty"`
}

// Department represents a department from the Contacts API
type Department struct {
	Name               string              `json:"name,omitempty"`
	I18nName           *DepartmentI18nName `json:"i18n_name,omitempty"`
	ParentDepartmentID string              `json:"parent_department_id,omitempty"`
	DepartmentID       string              `json:"department_id,omitempty"`
	OpenDepartmentID   string              `json:"open_department_id,omitempty"`
	LeaderUserID       string              `json:"leader_user_id,omitempty"`
	ChatID             string              `json:"chat_id,omitempty"`
	Order              string              `json:"order,omitempty"`
	MemberCount        int                 `json:"member_count,omitempty"`
	Status             *DepartmentStatus   `json:"status,omitempty"`
}

// --- Contact API Response Types ---

// GetUserResponse is the response from GET /contact/v3/users/:user_id
type GetUserResponse struct {
	BaseResponse
	Data struct {
		User *ContactUser `json:"user,omitempty"`
	} `json:"data,omitempty"`
}

// FindByDepartmentResponse is the response from GET /contact/v3/users/find_by_department
type FindByDepartmentResponse struct {
	BaseResponse
	Data struct {
		HasMore   bool          `json:"has_more"`
		PageToken string        `json:"page_token,omitempty"`
		Items     []ContactUser `json:"items,omitempty"`
	} `json:"data,omitempty"`
}

// GetDepartmentResponse is the response from GET /contact/v3/departments/:department_id
type GetDepartmentResponse struct {
	BaseResponse
	Data struct {
		Department *Department `json:"department,omitempty"`
	} `json:"data,omitempty"`
}

// SearchDepartmentsResponse is the response from POST /contact/v3/departments/search
type SearchDepartmentsResponse struct {
	BaseResponse
	Data struct {
		Items     []Department `json:"items,omitempty"`
		PageToken string       `json:"page_token,omitempty"`
		HasMore   bool         `json:"has_more"`
	} `json:"data,omitempty"`
}

// SearchUserAvatar represents user avatar information from search
type SearchUserAvatar struct {
	Avatar72     string `json:"avatar_72,omitempty"`
	Avatar240    string `json:"avatar_240,omitempty"`
	Avatar640    string `json:"avatar_640,omitempty"`
	AvatarOrigin string `json:"avatar_origin,omitempty"`
}

// SearchUserResult represents a user from the search API
type SearchUserResult struct {
	Avatar        *SearchUserAvatar `json:"avatar,omitempty"`
	DepartmentIDs []string          `json:"department_ids,omitempty"`
	Name          string            `json:"name,omitempty"`
	OpenID        string            `json:"open_id,omitempty"`
	UserID        string            `json:"user_id,omitempty"`
}

// SearchUsersResponse is the response from GET /search/v1/user
type SearchUsersResponse struct {
	BaseResponse
	Data struct {
		HasMore   bool               `json:"has_more"`
		PageToken string             `json:"page_token,omitempty"`
		Users     []SearchUserResult `json:"users,omitempty"`
	} `json:"data,omitempty"`
}

// --- Contact CLI Output Types ---

// OutputContact is the simplified contact format for CLI output
type OutputContact struct {
	UserID     string `json:"user_id"`
	OpenID     string `json:"open_id,omitempty"`
	Name       string `json:"name"`
	EnName     string `json:"en_name,omitempty"`
	Email      string `json:"email,omitempty"`
	JobTitle   string `json:"job_title,omitempty"`
	Department string `json:"department,omitempty"` // Primary department name
}

// OutputContactList is the list contacts response for CLI
type OutputContactList struct {
	Contacts []OutputContact `json:"contacts"`
	Count    int             `json:"count"`
	HasMore  bool            `json:"has_more,omitempty"`
}

// OutputDepartment is the simplified department format for CLI output
type OutputDepartment struct {
	DepartmentID string `json:"department_id"`
	Name         string `json:"name"`
	MemberCount  int    `json:"member_count,omitempty"`
}

// OutputDepartmentList is the list departments response for CLI
type OutputDepartmentList struct {
	Departments []OutputDepartment `json:"departments"`
	Count       int                `json:"count"`
	HasMore     bool               `json:"has_more,omitempty"`
}

// --- Document Types ---

// Document represents a Lark document
type Document struct {
	DocumentID string `json:"document_id"`
	RevisionID int    `json:"revision_id"`
	Title      string `json:"title"`
}

// TextElementStyle represents text styling
type TextElementStyle struct {
	Bold            bool `json:"bold,omitempty"`
	Italic          bool `json:"italic,omitempty"`
	Strikethrough   bool `json:"strikethrough,omitempty"`
	Underline       bool `json:"underline,omitempty"`
	InlineCode      bool `json:"inline_code,omitempty"`
	BackgroundColor int  `json:"background_color,omitempty"`
	TextColor       int  `json:"text_color,omitempty"`
}

// TextRun represents a text run element
type TextRun struct {
	Content          string            `json:"content,omitempty"`
	TextElementStyle *TextElementStyle `json:"text_element_style,omitempty"`
}

// MentionUser represents an @user element
type MentionUser struct {
	UserID           string            `json:"user_id,omitempty"`
	TextElementStyle *TextElementStyle `json:"text_element_style,omitempty"`
}

// MentionDoc represents an @document element
type MentionDoc struct {
	Token            string            `json:"token,omitempty"`
	ObjType          int               `json:"obj_type,omitempty"`
	URL              string            `json:"url,omitempty"`
	Title            string            `json:"title,omitempty"`
	TextElementStyle *TextElementStyle `json:"text_element_style,omitempty"`
}

// TextElement represents a text element within a block
type TextElement struct {
	TextRun     *TextRun     `json:"text_run,omitempty"`
	MentionUser *MentionUser `json:"mention_user,omitempty"`
	MentionDoc  *MentionDoc  `json:"mention_doc,omitempty"`
}

// TextStyle represents text block styling
type TextStyle struct {
	Align    int  `json:"align,omitempty"`
	Done     bool `json:"done,omitempty"`
	Folded   bool `json:"folded,omitempty"`
	Language int  `json:"language,omitempty"`
	Wrap     bool `json:"wrap,omitempty"`
}

// TextBlock represents text content in a block
type TextBlock struct {
	Style    *TextStyle    `json:"style,omitempty"`
	Elements []TextElement `json:"elements,omitempty"`
}

// ImageBlock represents an image block in a document
type ImageBlock struct {
	Token  string `json:"token,omitempty"`  // Image token for downloading
	Width  int    `json:"width,omitempty"`  // Image width in px
	Height int    `json:"height,omitempty"` // Image height in px
	Align  int    `json:"align,omitempty"`  // Alignment: 1=left, 2=center, 3=right
}

// DocumentBlock represents a block in a document
type DocumentBlock struct {
	BlockID   string      `json:"block_id"`
	ParentID  string      `json:"parent_id,omitempty"`
	Children  []string    `json:"children,omitempty"`
	BlockType int         `json:"block_type"`
	Page      *TextBlock  `json:"page,omitempty"`
	Text      *TextBlock  `json:"text,omitempty"`
	Image     *ImageBlock `json:"image,omitempty"`
}

// --- Document API Response Types ---

// DocumentResponse is the response from GET /docx/v1/documents/:document_id
type DocumentResponse struct {
	BaseResponse
	Data struct {
		Document *Document `json:"document,omitempty"`
	} `json:"data,omitempty"`
}

// DocumentBlocksResponse is the response from GET /docx/v1/documents/:document_id/blocks
type DocumentBlocksResponse struct {
	BaseResponse
	Data struct {
		Items     []DocumentBlock `json:"items,omitempty"`
		HasMore   bool            `json:"has_more"`
		PageToken string          `json:"page_token,omitempty"`
	} `json:"data,omitempty"`
}

// DocumentContentResponse is the response from GET /docs/v1/content
type DocumentContentResponse struct {
	BaseResponse
	Data struct {
		Content string `json:"content,omitempty"`
	} `json:"data,omitempty"`
}

// --- Document CLI Output Types ---

// OutputDocumentContent is the document content response for CLI
type OutputDocumentContent struct {
	DocumentID string `json:"document_id"`
	Title      string `json:"title,omitempty"`
	Content    string `json:"content"`
}

// OutputDocumentBlocks is the document blocks response for CLI
type OutputDocumentBlocks struct {
	DocumentID string          `json:"document_id"`
	Title      string          `json:"title,omitempty"`
	BlockCount int             `json:"block_count"`
	Blocks     []DocumentBlock `json:"blocks"`
}

// --- Media Types ---

// TmpDownloadURL represents a temporary download URL for a media file
type TmpDownloadURL struct {
	FileToken      string `json:"file_token"`
	TmpDownloadURL string `json:"tmp_download_url"`
}

// MediaTempDownloadURLResponse is the response from GET /drive/v1/medias/batch_get_tmp_download_url
type MediaTempDownloadURLResponse struct {
	BaseResponse
	Data struct {
		TmpDownloadURLs []TmpDownloadURL `json:"tmp_download_urls"`
	} `json:"data,omitempty"`
}

// --- Wiki Types ---

// WikiNode represents a wiki node from the Wiki API
type WikiNode struct {
	SpaceID         string `json:"space_id"`
	NodeToken       string `json:"node_token"`
	ObjToken        string `json:"obj_token"`
	ObjType         string `json:"obj_type"`
	ParentNodeToken string `json:"parent_node_token"`
	NodeType        string `json:"node_type"`
	OriginNodeToken string `json:"origin_node_token"`
	OriginSpaceID   string `json:"origin_space_id"`
	HasChild        bool   `json:"has_child"`
	Title           string `json:"title"`
	ObjCreateTime   string `json:"obj_create_time"`
	ObjEditTime     string `json:"obj_edit_time"`
	NodeCreateTime  string `json:"node_create_time"`
	Creator         string `json:"creator"`
	Owner           string `json:"owner"`
}

// WikiNodeResponse is the response from GET /wiki/v2/spaces/get_node
type WikiNodeResponse struct {
	BaseResponse
	Data struct {
		Node *WikiNode `json:"node,omitempty"`
	} `json:"data,omitempty"`
}

// OutputWikiNode is the wiki node response for CLI
type OutputWikiNode struct {
	NodeToken string `json:"node_token"`
	ObjToken  string `json:"obj_token"`
	ObjType   string `json:"obj_type"`
	Title     string `json:"title"`
	SpaceID   string `json:"space_id"`
	NodeType  string `json:"node_type"`
	HasChild  bool   `json:"has_child"`
}

// ListWikiChildrenResponse is the response from GET /wiki/v2/spaces/:space_id/nodes
type ListWikiChildrenResponse struct {
	BaseResponse
	Data struct {
		Items     []WikiNode `json:"items,omitempty"`
		PageToken string     `json:"page_token,omitempty"`
		HasMore   bool       `json:"has_more"`
	} `json:"data,omitempty"`
}

// OutputWikiChildren is the wiki children response for CLI
type OutputWikiChildren struct {
	ParentNodeToken string           `json:"parent_node_token"`
	SpaceID         string           `json:"space_id"`
	Children        []OutputWikiNode `json:"children"`
	Count           int              `json:"count"`
}

// WikiSearchRequest is the request body for POST /wiki/v2/nodes/search
type WikiSearchRequest struct {
	Query     string `json:"query"`
	SpaceID   string `json:"space_id,omitempty"`
	NodeID    string `json:"node_id,omitempty"`
	PageToken string `json:"page_token,omitempty"`
	PageSize  int    `json:"page_size,omitempty"`
}

// WikiSearchItem represents a wiki node from search results
type WikiSearchItem struct {
	NodeID   string `json:"node_id"`
	SpaceID  string `json:"space_id"`
	ObjType  int    `json:"obj_type"`
	ObjToken string `json:"obj_token"`
	Title    string `json:"title"`
	URL      string `json:"url"`
}

// WikiSearchResponse is the response from POST /wiki/v2/nodes/search
type WikiSearchResponse struct {
	BaseResponse
	Data struct {
		Items     []WikiSearchItem `json:"items,omitempty"`
		PageToken string           `json:"page_token,omitempty"`
		HasMore   bool             `json:"has_more"`
	} `json:"data"`
}

// OutputWikiSearchResult is the wiki search response for CLI
type OutputWikiSearchResult struct {
	Query   string                 `json:"query"`
	SpaceID string                 `json:"space_id,omitempty"`
	Results []OutputWikiSearchItem `json:"results"`
	Count   int                    `json:"count"`
}

// OutputWikiSearchItem is a search result item for CLI output
type OutputWikiSearchItem struct {
	NodeID   string `json:"node_id"`
	ObjToken string `json:"obj_token"`
	ObjType  string `json:"obj_type"` // Human-readable type
	Title    string `json:"title"`
	URL      string `json:"url"`
	SpaceID  string `json:"space_id"`
}

// --- Folder/Drive Types ---

// ShortcutInfo contains information about a shortcut's target
type ShortcutInfo struct {
	TargetType  string `json:"target_type"`
	TargetToken string `json:"target_token"`
}

// FolderItem represents an item in a Lark Drive folder
type FolderItem struct {
	Token        string        `json:"token"`
	Name         string        `json:"name"`
	Type         string        `json:"type"`
	ParentToken  string        `json:"parent_token"`
	URL          string        `json:"url"`
	ShortcutInfo *ShortcutInfo `json:"shortcut_info,omitempty"`
}

// ListFolderItemsResponse is the API response for listing folder items
type ListFolderItemsResponse struct {
	BaseResponse
	Data struct {
		Files         []FolderItem `json:"files"`
		NextPageToken string       `json:"next_page_token,omitempty"`
		HasMore       bool         `json:"has_more"`
	} `json:"data"`
}

// OutputFolderItem is the CLI output format for a folder item
type OutputFolderItem struct {
	Token        string        `json:"token"`
	Name         string        `json:"name"`
	Type         string        `json:"type"`
	ParentToken  string        `json:"parent_token,omitempty"`
	URL          string        `json:"url"`
	ShortcutInfo *ShortcutInfo `json:"shortcut_info,omitempty"`
}

// OutputFolderItemsList is the CLI output for listing folder items
type OutputFolderItemsList struct {
	FolderToken string             `json:"folder_token,omitempty"`
	Items       []OutputFolderItem `json:"items"`
	Count       int                `json:"count"`
}

// --- Document Comment Types ---

// CommentReplyElement represents an element in a comment reply
type CommentReplyElement struct {
	Type    string `json:"type,omitempty"`
	TextRun *struct {
		Text string `json:"text,omitempty"`
	} `json:"text_run,omitempty"`
	DocsLink *struct {
		URL string `json:"url,omitempty"`
	} `json:"docs_link,omitempty"`
	Person *struct {
		UserID string `json:"user_id,omitempty"`
	} `json:"person,omitempty"`
}

// CommentReply represents a reply within a comment
type CommentReply struct {
	ReplyID    string `json:"reply_id,omitempty"`
	UserID     string `json:"user_id,omitempty"`
	CreateTime int64  `json:"create_time,omitempty"`
	UpdateTime int64  `json:"update_time,omitempty"`
	Content    struct {
		Elements []CommentReplyElement `json:"elements,omitempty"`
	} `json:"content,omitempty"`
}

// DocumentComment represents a comment on a document
type DocumentComment struct {
	CommentID    string `json:"comment_id,omitempty"`
	UserID       string `json:"user_id,omitempty"`
	CreateTime   int64  `json:"create_time,omitempty"`
	UpdateTime   int64  `json:"update_time,omitempty"`
	IsSolved     bool   `json:"is_solved,omitempty"`
	SolvedTime   int64  `json:"solved_time,omitempty"`
	SolverUserID string `json:"solver_user_id,omitempty"`
	IsWhole      bool   `json:"is_whole,omitempty"`
	Quote        string `json:"quote,omitempty"`
	ReplyList    struct {
		Replies []CommentReply `json:"replies,omitempty"`
	} `json:"reply_list,omitempty"`
}

// DocumentCommentsResponse is the response from GET /drive/v1/files/:file_token/comments
type DocumentCommentsResponse struct {
	BaseResponse
	Data struct {
		Items     []DocumentComment `json:"items,omitempty"`
		HasMore   bool              `json:"has_more"`
		PageToken string            `json:"page_token,omitempty"`
	} `json:"data,omitempty"`
}

// --- Document Comment CLI Output Types ---

// OutputCommentReply is the simplified reply format for CLI output
type OutputCommentReply struct {
	ReplyID    string `json:"reply_id"`
	UserID     string `json:"user_id"`
	CreateTime string `json:"create_time"`
	Text       string `json:"text"`
}

// OutputDocumentComment is the simplified comment format for CLI output
type OutputDocumentComment struct {
	CommentID  string               `json:"comment_id"`
	UserID     string               `json:"user_id"`
	CreateTime string               `json:"create_time"`
	IsSolved   bool                 `json:"is_solved"`
	IsWhole    bool                 `json:"is_whole"`
	Quote      string               `json:"quote,omitempty"`
	Replies    []OutputCommentReply `json:"replies,omitempty"`
}

// OutputDocumentComments is the document comments response for CLI
type OutputDocumentComments struct {
	FileToken string                  `json:"file_token"`
	Comments  []OutputDocumentComment `json:"comments"`
	Count     int                     `json:"count"`
}

// --- Message Types ---

// MessageSender represents the sender of a message
type MessageSender struct {
	ID         string `json:"id,omitempty"`
	IDType     string `json:"id_type,omitempty"`     // open_id, app_id
	SenderType string `json:"sender_type,omitempty"` // user, app, anonymous, unknown
	TenantKey  string `json:"tenant_key,omitempty"`
}

// MessageBody represents the content of a message
type MessageBody struct {
	Content string `json:"content,omitempty"` // JSON string of message content
}

// MessageMention represents a mentioned user or bot in a message
type MessageMention struct {
	Key       string `json:"key,omitempty"`     // e.g., "@_user_1"
	ID        string `json:"id,omitempty"`      // open_id of mentioned user/bot
	IDType    string `json:"id_type,omitempty"` // currently only "open_id"
	Name      string `json:"name,omitempty"`    // Name of mentioned user/bot
	TenantKey string `json:"tenant_key,omitempty"`
}

// Message represents a message from the IM API
type Message struct {
	MessageID      string           `json:"message_id,omitempty"`
	RootID         string           `json:"root_id,omitempty"`     // Root message ID for replies
	ParentID       string           `json:"parent_id,omitempty"`   // Parent message ID for replies
	ThreadID       string           `json:"thread_id,omitempty"`   // Thread ID if in a thread
	MsgType        string           `json:"msg_type,omitempty"`    // text, post, image, file, audio, media, sticker, interactive, share_chat, share_user
	CreateTime     string           `json:"create_time,omitempty"` // Unix ms timestamp
	UpdateTime     string           `json:"update_time,omitempty"` // Unix ms timestamp
	Deleted        bool             `json:"deleted,omitempty"`     // Whether message is deleted/recalled
	Updated        bool             `json:"updated,omitempty"`     // Whether message is updated
	ChatID         string           `json:"chat_id,omitempty"`     // Group ID
	Sender         *MessageSender   `json:"sender,omitempty"`
	Body           *MessageBody     `json:"body,omitempty"`
	Mentions       []MessageMention `json:"mentions,omitempty"`
	UpperMessageID string           `json:"upper_message_id,omitempty"` // For combined/forwarded messages
}

// --- Message API Response Types ---

// MessageListResponse is the response from GET /im/v1/messages
type MessageListResponse struct {
	BaseResponse
	Data struct {
		HasMore   bool      `json:"has_more"`
		PageToken string    `json:"page_token,omitempty"`
		Items     []Message `json:"items,omitempty"`
	} `json:"data,omitempty"`
}

// --- Message CLI Output Types ---

// OutputMessageSender is the simplified sender format for CLI output
type OutputMessageSender struct {
	ID   string `json:"id"`
	Type string `json:"type"` // user, app, anonymous, unknown
}

// OutputMessageMention is the simplified mention format for CLI output
type OutputMessageMention struct {
	Key  string `json:"key"`
	ID   string `json:"id"`
	Name string `json:"name"`
}

// OutputMessage is the simplified message format for CLI output
type OutputMessage struct {
	MessageID  string                 `json:"message_id"`
	MsgType    string                 `json:"msg_type"`
	Content    string                 `json:"content"`
	Sender     *OutputMessageSender   `json:"sender,omitempty"`
	CreateTime string                 `json:"create_time"`
	Mentions   []OutputMessageMention `json:"mentions,omitempty"`
	IsReply    bool                   `json:"is_reply,omitempty"`
	ThreadID   string                 `json:"thread_id,omitempty"`
	Deleted    bool                   `json:"deleted,omitempty"`
}

// OutputMessageList is the message list response for CLI
type OutputMessageList struct {
	Messages []OutputMessage `json:"messages"`
	Count    int             `json:"count"`
	ChatID   string          `json:"chat_id"`
}

// --- Send Message Types ---

// SendMessageRequest is the request body for POST /im/v1/messages
type SendMessageRequest struct {
	ReceiveID string `json:"receive_id"`
	MsgType   string `json:"msg_type"` // text, post
	Content   string `json:"content"`  // JSON string
}

// UploadImageResponse is the response from POST /im/v1/images
type UploadImageResponse struct {
	BaseResponse
	Data struct {
		ImageKey string `json:"image_key"`
	} `json:"data,omitempty"`
}

// SendMessageResponse is the response from POST /im/v1/messages
type SendMessageResponse struct {
	BaseResponse
	Data struct {
		MessageID  string           `json:"message_id"`
		RootID     string           `json:"root_id,omitempty"`
		ParentID   string           `json:"parent_id,omitempty"`
		MsgType    string           `json:"msg_type"`
		CreateTime string           `json:"create_time"`
		UpdateTime string           `json:"update_time"`
		Deleted    bool             `json:"deleted"`
		Updated    bool             `json:"updated"`
		ChatID     string           `json:"chat_id"`
		Sender     *MessageSender   `json:"sender,omitempty"`
		Body       *MessageBody     `json:"body,omitempty"`
		Mentions   []MessageMention `json:"mentions,omitempty"`
	} `json:"data,omitempty"`
}

// OutputSendMessage is the simplified send message response for CLI
type OutputSendMessage struct {
	Success    bool   `json:"success"`
	MessageID  string `json:"message_id"`
	ChatID     string `json:"chat_id,omitempty"`
	CreateTime string `json:"create_time"`
}

// --- Chat Types ---

// Chat represents a chat/group from the IM API
type Chat struct {
	ChatID      string `json:"chat_id,omitempty"`
	Avatar      string `json:"avatar,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	OwnerID     string `json:"owner_id,omitempty"`
	OwnerIDType string `json:"owner_id_type,omitempty"`
	External    bool   `json:"external,omitempty"`
	TenantKey   string `json:"tenant_key,omitempty"`
	ChatStatus  string `json:"chat_status,omitempty"`
}

// SearchChatsResponse is the response from GET /im/v1/chats/search
type SearchChatsResponse struct {
	BaseResponse
	Data struct {
		Items     []Chat `json:"items,omitempty"`
		PageToken string `json:"page_token,omitempty"`
		HasMore   bool   `json:"has_more"`
	} `json:"data,omitempty"`
}

// --- Chat CLI Output Types ---

// OutputChat is the simplified chat format for CLI output
type OutputChat struct {
	ChatID      string `json:"chat_id"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	OwnerID     string `json:"owner_id,omitempty"`
	External    bool   `json:"external,omitempty"`
	ChatStatus  string `json:"chat_status,omitempty"`
}

// OutputChatList is the chat list response for CLI
type OutputChatList struct {
	Chats []OutputChat `json:"chats"`
	Count int          `json:"count"`
	Query string       `json:"query,omitempty"`
}

// --- Minutes Types ---

// Minute represents a Lark Minutes recording
type Minute struct {
	Token      string `json:"token"`
	OwnerID    string `json:"owner_id"`
	CreateTime string `json:"create_time"` // Unix ms timestamp
	Title      string `json:"title"`
	Cover      string `json:"cover"`    // Cover image URL
	Duration   string `json:"duration"` // Duration in ms
	URL        string `json:"url"`      // Minutes link
}

// MinuteResponse is the response from GET /minutes/v1/minutes/:minute_token
type MinuteResponse struct {
	BaseResponse
	Data struct {
		Minute *Minute `json:"minute,omitempty"`
	} `json:"data,omitempty"`
}

// MinuteMediaResponse is the response from GET /minutes/v1/minutes/:minute_token/media
type MinuteMediaResponse struct {
	BaseResponse
	Data struct {
		DownloadURL string `json:"download_url"`
	} `json:"data,omitempty"`
}

// --- Minutes CLI Output Types ---

// OutputMinute is the simplified minute format for CLI output
type OutputMinute struct {
	Token           string `json:"token"`
	Title           string `json:"title"`
	OwnerID         string `json:"owner_id"`
	CreateTime      string `json:"create_time"`      // ISO 8601 format
	DurationSeconds int    `json:"duration_seconds"` // Duration in seconds
	DurationDisplay string `json:"duration_display"` // Human-readable duration
	URL             string `json:"url"`
}

// OutputMinuteMedia is the media download URL response for CLI
type OutputMinuteMedia struct {
	Token       string `json:"token"`
	DownloadURL string `json:"download_url"`
}

// OutputMinuteTranscript is the transcript export response for CLI
type OutputMinuteTranscript struct {
	Token   string `json:"token"`
	Format  string `json:"format"`
	Content string `json:"content,omitempty"`
	File    string `json:"file,omitempty"` // Output file path if written to file
}

// --- Document Search Types ---

// DocSearchRequest is the request body for POST /suite/docs-api/search/object
type DocSearchRequest struct {
	SearchKey string   `json:"search_key"`
	Count     int      `json:"count,omitempty"`
	Offset    int      `json:"offset,omitempty"`
	OwnerIDs  []string `json:"owner_ids,omitempty"`
	ChatIDs   []string `json:"chat_ids,omitempty"`
	DocsTypes []string `json:"docs_types,omitempty"`
}

// DocSearchEntity represents a document from search results
type DocSearchEntity struct {
	DocsToken string `json:"docs_token"`
	DocsType  string `json:"docs_type"`
	Title     string `json:"title"`
	OwnerID   string `json:"owner_id"`
}

// DocSearchResponse is the response from POST /suite/docs-api/search/object
type DocSearchResponse struct {
	BaseResponse
	Data struct {
		DocsEntities []DocSearchEntity `json:"docs_entities,omitempty"`
		HasMore      bool              `json:"has_more"`
		Total        int               `json:"total"`
	} `json:"data"`
}

// OutputDocSearchResult is the doc search response for CLI
type OutputDocSearchResult struct {
	Query   string                `json:"query"`
	Results []OutputDocSearchItem `json:"results"`
	Total   int                   `json:"total"`
	Count   int                   `json:"count"`
}

// OutputDocSearchItem is a search result item for CLI output
type OutputDocSearchItem struct {
	Token   string `json:"token"`
	Type    string `json:"type"`
	Title   string `json:"title"`
	OwnerID string `json:"owner_id"`
}
