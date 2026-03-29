---
title: "Get Event List"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar-event/list"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/calendar/v4/calendars/:calendar_id/events"
service: "calendar-v4"
resource: "calendar-event"
section: "Calendar"
rate_limit: "1000 per minute、50 per second"
scopes:
  - "calendar:calendar"
  - "calendar:calendar.event:read"
  - "calendar:calendar:readonly"
field_scopes:
  - "contact:user.employee_id:readonly"
updated: "1736822150000"
---

# Obtain event list

Call this interface with the current identity (app or user) to obtain the event list under a specified calendar.

> - The current identity is determined by the Token type of Header Authorization. tenant_access_token refers to the application identity, and user_access_token refers to the user identity.
> - The current identity must have reader, writer or owner permissions on the calendar, and only supports obtaining the event list of primary, shared and resource type calendars. Google and Exchange type calendars are not supported yet. You can call the Query Calendar Information interface to obtain the current identity’s access rights to the calendar and calendar type information.
> - The page_token request parameter paging is used to pull existing data, and the sync_token request parameter is used to incrementally synchronize changed data. Currently, page_token will be returned only when anchor_time is passed in.
> - In order to ensure the consistency of event synchronization data when you call the interface, when you use the sync_token request parameter, you cannot pass in start_time and end_time at the same time, otherwise the event data may be missing.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/calendar/v4/calendars/:calendar_id/events |
| HTTP Method | GET |
| Rate Limit | 1000 per minute、50 per second |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `calendar:calendar` `calendar:calendar.event:read` `calendar:calendar:readonly` |
| Required field scopes | > The response body of the API contains the following sensitive fields, and they will be returned only after corresponding scopes are added. If you do not need the fields, it is not recommended that you request the scopes. `contact:user.employee_id:readonly` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `calendar_id` | `string` | Calendar ID. For details, see Calendar-related IDs. **Example value**: "larksuite.com_xxxxxxxxxx@group.calendar.larksuite.com" | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `page_size` | `int` | No | The maximum number of entries to be returned for one request. The actual number of returned entries may be less than this value, or it may be empty, and you can determine if there are more entries based on the has_more field in the response body. **Example value**: 50 **Default value**: `500` **Data validation rules**: - Value range: `50` ～ `1000` |
| `anchor_time` | `string` | No | A time anchor represented by a Unix timestamp (in seconds). It is used to set a specific point in time for pulling events, thereby avoiding the need to pull all events. You can use page_token or sync_token for pagination or incremental retrieval of all events after the specified anchor_time. **Notice**: It can not be used together with start_time and end_time. **Default**: empty **Example value**: 1609430400 |
| `page_token` | `string` | No | Page identifier. It is not filled in the first request, indicating traversal from the beginning; when there will be more groups, the new page_token will be returned at the same time, and the next traversal can use the page_token to get more groups **Example value**: ListCalendarsPageToken_1632452910_1632539310 |
| `sync_token` | `string` | No | Incremental synchronization mark, not filled in for the first request. When the paging query ends (page_token return value is empty), the interface will return the sync_token field, and the sync_token can be used to incrementally obtain calendar change data in the next call. **Default**: empty **Example value**: ListCalendarsSyncToken_1632452910 |
| `start_time` | `string` | No | start_time is the start point of a time range, represented by a Unix timestamp (in seconds). It is used in conjunction with end_time to pull events within a specified time range. **Notice**: - This method can only return data once and cannot be paginated. The size of the data returned in a single request is limited by page_size, and any data exceeding this limit will be truncated. - When using start_time and end_time, they cannot be used together with page_token or sync_token. - When using start_time and end_time, it cannot be used with anchor_time. **Default**: empty **Example value**: 1631777271 |
| `end_time` | `string` | No | end_time is the end point of a time range, represented by a Unix timestamp (in seconds). It is used in conjunction with start_time to pull events within a specified time range. **Notice**: - This method can only return data once and cannot be paginated. The size of the data returned in a single request is limited by page_size, and any data exceeding this limit will be truncated. - When using start_time and end_time, they cannot be used together with page_token or sync_token. - When using start_time and end_time, it cannot be used with anchor_time. **Default**: empty **Example value**: 1631777271 |
| `user_id_type` | `string` | No | User ID categories **Example value**: open_id **Optional values are**:  - `open_id`: Identifies a user to an app. The same user has different Open IDs in different apps. How to get Open ID - `union_id`: Identifies a user to a tenant that acts as a developer. A user has the same Union ID in apps developed by the same developer, and has different Union IDs in apps developed by different developers. A developer can use Union ID to link the same user's identities in multiple apps.How to get Union ID - `user_id`: Identifies a user to a tenant. The same user has different User IDs in different tenants. In one single tenant, a user has the same User ID in all apps （including store apps）. User ID is usually used to communicate user data between different apps. How to get User ID  **Default value**: `open_id` **When the value is `user_id`, the following field scopes are required**: `contact:user.employee_id:readonly` | ## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `has_more` | `boolean` | Whether the response body has more parameters |
|   `page_token` | `string` | Page identifier, when has_more is true, a new page_token will also be returned. Otherwise, page_token will not be returned |
|   `sync_token` | `string` | Incremental synchronization mark. When has_more is false, a new sync_token will be returned synchronously. The next request needs to bring sync_token to incrementally obtain calendar change data. |
|   `items` | `calendar.event[]` | Event list. When returned empty, determine if there is more data based on the value of has_more. |
|     `event_id` | `string` | Event ID. Event information can be queried, updated, or deleted later through this ID. More information can be found in Event ID Description. |
|     `organizer_calendar_id` | `string` | Event organizer calendar ID. See Calendar ID description. |
|     `summary` | `string` | Event title. |
|     `description` | `string` | Event description. |
|     `start_time` | `time_info` | Event start time. |
|       `date` | `string` | Start time, only use this field for all-day events, [RFC 3339](https://datatracker.ietf.org/doc/html/rfc3339) format, for example, 2018-09-01. |
|       `timestamp` | `string` | Second-level timestamp refers to the specific start time of the event. For example, 1602504000 means 2020/10/12 20:00:00 (UTC +8 time zone). |
|       `timezone` | `string` | Time zone. Use the IANA Time Zone Database standard. |
|     `end_time` | `time_info` | Event end time. |
|       `date` | `string` | End time, only use this field for all-day events, [RFC 3339](https://datatracker.ietf.org/doc/html/rfc3339) format, for example, 2018-09-01. |
|       `timestamp` | `string` | Second-level timestamp refers to the specific end time of the event. For example, 1602504000 means 2020/10/12 20:00:00 (UTC +8 time zone). |
|       `timezone` | `string` | Time zone. Use the IANA Time Zone Database standard. |
|     `vchat` | `vchat` | Video conference information. |
|       `vc_type` | `string` | Video conference type, which can be empty, means that the Lark video conference URL will be automatically generated when a event invitee is added for the first time. **Optional values are**:  - `vc`: Lark video conferencing. When this type is selected, other fields in vchat are invalid. - `third_party`: Third-party linked video conferencing. When this type is selected, only the icon_type, description, and meeting_url fields will take effect in vchat. - `no_meeting`: No video conferencing. When this type is selected, other fields in vchat are invalid. - `lark_live`: Lark Live, read-only parameters. - `unknown`: Unknown type for compatible read-only parameters.  |
|       `icon_type` | `string` | Third-party video conferencing icon type, which can be empty to display the default icon. **Optional values are**:  - `vc`: Lark Video Conference icon - `live`: Livestream video conference icon - `default`: Default icon  |
|       `description` | `string` | Text for the third-party video conference. |
|       `meeting_url` | `string` | Video conference URL. |
|     `visibility` | `string` | Event visibility. It will only take effect for all invitees when creating a new event, and subsequent modifications to this attribute will only take effect for the current identity. **Optional values are**:  - `default`: Default range, which depends on the calendar visibility. Only the availability status is visible to other people by default. - `public`: Public. Event details are displayed. - `private`: Private. Details are visible only to the current identity.  |
|     `attendee_ability` | `string` | Event invitees' scopes. **Optional values are**:  - `none`: Cannot edit events, cannot invite others, and cannot view event invitee list - `can_see_others`: Cannot edit events, cannot invite others, and can view event invitee list - `can_invite_others`: Cannot edit events, can invite others, and can view event invitee list - `can_modify_event`: Can edit events, can invite others, and can view event invitee list  |
|     `free_busy_status` | `string` | The busy and busy status occupied by the event. It will only take effect for all invitees when creating a new event, and subsequent modifications to this attribute will only take effect for the current identity. **Optional values are**:  - `busy`: Busy - `free`: Free  |
|     `location` | `event_location` | Event location. |
|       `name` | `string` | Name of the location. |
|       `address` | `string` | Address of the location. |
|       `latitude` | `number(float)` | Latitude coordinate of the location. Locations within China Mainland should comply with the GCJ-02 standard, and locations outside China Mainland should comply with the WGS84 standard. |
|       `longitude` | `number(float)` | Longitude coordinate of the location. Locations within China Mainland should comply with the GCJ-02 standard, and locations outside China Mainland should comply with the WGS84 standard. |
|     `color` | `int` | Event color, represented by an int32 of color RGB values. **illustrate**: - Only takes effect for the current identity. - When the value is 0 or -1, it means to follow the calendar color by default. - When displayed on the client, it will be mapped to the closest color on the color palette. |
|     `reminders` | `reminder[]` | Event reminder list. |
|       `minutes` | `int` | Offset of the event reminder time. A positive value means the reminder is triggered X minutes before the event starts, and a negative value means the reminder is triggered X minutes after the event starts. This field is specified when an event is created or updated, and only valid for the current identity, not for other invitees in the event. |
|     `recurrence` | `string` | Repeat rules for repeating events, the rule format can be found in [rfc5545](https://datatracker.ietf.org/doc/html/rfc5545#section-3.3.10). |
|     `status` | `string` | Event status. **Optional values are**:  - `tentative`: Not replied - `confirmed`: Confirmed - `cancelled`: Canceled  |
|     `is_exception` | `boolean` | Whether the event is an exception to a repeating event. To learn about the exception event, see Exception event. |
|     `recurring_event_id` | `string` | The event_id of the original recurring event of an exceptional event. |
|     `create_time` | `string` | Event create time (timestamp in seconds). |
|     `schemas` | `schema[]` | Custom information of the event. This field controls the UI display on the event details page. |
|       `ui_name` | `string` | UI name. Valid values: - ForwardIcon: Event forwarding button - MeetingChatIcon: Meeting group chat button - MeetingMinutesIcon: Meeting minutes button - MeetingVideo: Video conference area - RSVP: Accept/decline/pending area - Attendee: Event invitees area - OrganizerOrCreator: Organizer/creator area |
|       `ui_status` | `string` | UI item custom state. **Optional values are**:  - `hide`: Hidden - `readonly`: Read-only - `editable`: Editable - `unknown`: Unknown. This is only used for compatibility in reading.  |
|       `app_link` | `string` | Link that is redirected to after the button is tapped/clicked. |
|     `event_organizer` | `event_organizer` | Event organizer information. |
|       `user_id` | `string` | Event organizer user ID. |
|       `display_name` | `string` | Name of event organizer. |
|     `app_link` | `string` | app_link of the event, jump to a specific event |
|     `attachments` | `attachment[]` | Event attachment. |
|       `file_token` | `string` | Attachment token. |
|       `file_size` | `string` | Attachment size. |
|       `name` | `string` | Attachment name. | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
        "has_more": false,
        "page_token": "ListCalendarsPageToken_1632452910_1632539310",
        "sync_token": "ListCalendarsSyncToken_1632452910",
        "items": [
            {
                "event_id": "00592a0e-7edf-4678-bc9d-1b77383ef08e_0",
                "organizer_calendar_id": "larksuite.com_xxxxxxxxxx@group.calendar.larksuite.com",
                "summary": "Event title",
                "description": "Event description",
                "start_time": {
                    "date": "2018-09-01",
                    "timestamp": "1602504000",
                    "timezone": "Asia/Shanghai"
                },
                "end_time": {
                    "date": "2018-09-01",
                    "timestamp": "1602504000",
                    "timezone": "Asia/Shanghai"
                },
                "vchat": {
                    "vc_type": "third_party",
                    "icon_type": "vc",
                    "description": "Initiate a video conference",
                    "meeting_url": "https://example.com"
                },
                "visibility": "default",
                "attendee_ability": "can_see_others",
                "free_busy_status": "busy",
                "location": {
                    "name": "Name of the location",
                    "address": "Address of the location",
                    "latitude": 1.100000023841858,
                    "longitude": 2.200000047683716
                },
                "color": -1,
                "reminders": [
                    {
                        "minutes": 5
                    }
                ],
                "recurrence": "FREQ=DAILY;INTERVAL=1",
                "status": "confirmed",
                "is_exception": false,
                "recurring_event_id": "1cd45aaa-fa70-4195-80b7-c93b2e208f45",
                "create_time": "1602504000",
                "schemas": [
                    {
                        "ui_name": "ForwardIcon",
                        "ui_status": "hide",
                        "app_link": "https://applink.larksuite.com/client/calendar/event/detail?calendarId=xxxxxx&key=xxxxxx&originalTime=xxxxxx&startTime=xxxxxx"
                    }
                ],
                "event_organizer": {
                    "user_id": "ou_xxxxxx",
                    "display_name": "Jian Li"
                },
                "app_link": "https://applink.larkoffice.com/client/calendar/event/detail?calendarId=7039673579105026066&key=aeac9c56-aeb1-4179-a21b-02f278f59048&originalTime=0&startTime=1700496000",
                "attachments": [
                    {
                        "file_token": "xAAAAA",
                        "file_size": "2345",
                        "name": "demo.jpeg"
                    }
                ]
            }
        ]
    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 190002 | invalid parameters in request | Invalid request parameters. Troubleshooting suggestions are as follows: - Confirm that the field name and parameter type of the request parameter are correct. - Confirm that the permissions for the corresponding resource have been applied for. - Confirm that the corresponding resource has not been deleted. |
| 500 | 190003 | internal service error | Internal service error, please contact [Technical Support](https://applink.larksuite.com/TLJpeNdW). |
| 429 | 190004 | method rate limited | Method frequency limit. It is recommended to try again later and reduce the request QPS appropriately. |
| 429 | 190005 | app rate limited | Frequency limiting is applied. We recommend that you try again later and reduce the request QPS appropriately. |
| 403 | 190006 | wrong unit for app tenant | Request error, check whether the App ID and App Secret are correct. If the problem still cannot be solved, please consult [Technical Support](https://applink.larksuite.com/TLJpeNdW). |
| 404 | 190007 | app bot_id not found | The bot_id of the application is not found. You need to make sure that the application has enabled bot capability. If the problem is still not solved, please contact [technical support](https://applink.larksuite.com/TLJpeNdW). |
| 400 | 190008 | page_token or sync_token expired | The page_token or sync_token has expired. You need to empty the token parameter value and try again. |
| 400 | 190009 | sync token cannot be used with other request restrictions | sync_token cannot be used with other conflicting parameters. You need to refer to the sync_token parameter description to modify it to the correct parameter configuration. |
| 429 | 190010 | current operation rate limited | The current operation is limited, usually because concurrent preemption of public resources fails. You can appropriately reduce the frequency of the current operation and try again. |
| 404 | 191000 | calendar not found | Calendar not found. You need to check and change to the correct calendar ID. |
| 400 | 191001 | invalid calendar_id | The calendar_id is invalid. You need to check and change it to the correct calendar ID. |
| 403 | 191002 | no calendar access_role | The current identity does not have permission to access the calendar. If you want to query a calendar, you need to ensure that the current identity has permission to access the calendar. |
| 403 | 191003 | calendar is deleted | The calendar has been deleted. You need to check and change to the correct calendar ID. |
| 403 | 191004 | invalid calendar type | The calendar type is incorrect. You can call the query calendar information interface to obtain the calendar type information, and then ensure that the calendar type is applicable to the current interface. |
| 400 | 193000 | invalid event_id | The event_id is invalid. You need to check and change it to the correct event ID. |
| 404 | 193001 | event not found | Event not found. You need to make sure you passed in the correct event ID. |
| 403 | 193002 | no permission to operate event | No permission to operate. You need to make sure you have permission to edit the calendar and event. |
| 403 | 193003 | event is deleted | The event has been deleted. You need to check and change to the correct event ID. |
| 404 | 195100 | user is dismiss or not exist in the tenant | The current identity or specified user has resigned or is no longer in the tenant. Please check and change to the correct identity to call the interface. | For more error code information, see General error codes.
