---
title: "Search Event"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar-event/search"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/calendar/v4/calendars/:calendar_id/events/search"
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

# Search for events

Call this interface to search for relevant events under the specified calendar, supporting keyword search and filter condition search.

## Precautions

Applicable to primary and shared calendars, and the current identity must have reader, writer, or owner permissions on the calendar. You can call the query calendar information interface to obtain the current identity's access permissions to the calendar.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/calendar/v4/calendars/:calendar_id/events/search |
| HTTP Method | POST |
| Rate Limit | 1000 per minute、50 per second |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `calendar:calendar` `calendar:calendar.event:read` `calendar:calendar:readonly` |
| Required field scopes | > The response body of the API contains the following sensitive fields, and they will be returned only after corresponding scopes are added. If you do not need the fields, it is not recommended that you request the scopes. `contact:user.employee_id:readonly` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `calendar_id` | `string` | Calendar ID. For details, see Calendar-related IDs. **Example value**: "larksuite.com_xxxxxxxxxx@group.calendar.larksuite.com" | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `user_id_type` | `string` | No | User ID categories **Example value**: open_id **Optional values are**:  - `open_id`: Identifies a user to an app. The same user has different Open IDs in different apps. How to get Open ID - `union_id`: Identifies a user to a tenant that acts as a developer. A user has the same Union ID in apps developed by the same developer, and has different Union IDs in apps developed by different developers. A developer can use Union ID to link the same user's identities in multiple apps.How to get Union ID - `user_id`: Identifies a user to a tenant. The same user has different User IDs in different tenants. In one single tenant, a user has the same User ID in all apps （including store apps）. User ID is usually used to communicate user data between different apps. How to get User ID  **Default value**: `open_id` **When the value is `user_id`, the following field scopes are required**: `contact:user.employee_id:readonly` |
| `page_token` | `string` | No | Page identifier. It is not filled in the first request, indicating traversal from the beginning; when there will be more groups, the new page_token will be returned at the same time, and the next traversal can use the page_token to get more groups **Example value**: xxxxx |
| `page_size` | `int` | No | The maximum number of events returned in one call. The minimum value is 10, and if fewer than 10, defaults to 10. **Example value**: 10 **Default value**: `20` **Data validation rules**: - Maximum value: `100` | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `query` | `string` | Yes | Search for the keyword. Used to fuzzy query event name. **Notice**: If the event name contains an underscore (_), you must search for it accurately. This scene fuzzy search may not search for the event. **Example value**: "query words" **Data validation rules**: - Length range: `0` ～ `200` characters |
| `filter` | `event_search_filter` | No | Search filter. |
|   `start_time` | `time_info` | No | Search filter items, start time of event search interval. **Notice**: If start_time and end_time are not passed, the default search is for events within the past month. |
|     `date` | `string` | No | Specify the start time in days, in [RFC 3339](https://datatracker.ietf.org/doc/html/rfc3339) format, for example, 2018-09-01. **Note**: This parameter cannot be specified at the same time as `timestamp`. **Example value**: "2018-09-01" |
|     `timestamp` | `string` | No | Second-level timestamp refers to the specific start time. For example, 1602504000 means 2020/10/12 20:00:00 (UTC +8 time zone). **Note**: This parameter cannot be specified at the same time as `date`. **Example value**: "1602504000" |
|     `timezone` | `string` | No | Time zone. Use IANA Time Zone Database standards such as Asia/Shanghai. - The time zone is fixed to UTC +0 throughout the day - The non-all-day time zone defaults to Asia/Shanghai **Example value**: "Asia/Shanghai" |
|   `end_time` | `time_info` | No | Search filter items, end time of event search interval. **Notice**: If start_time and end_time are not passed, the default search is for events within the past month. |
|     `date` | `string` | No | Specify the end time in days, in [RFC 3339](https://datatracker.ietf.org/doc/html/rfc3339) format, for example, 2018-09-01. **Note**: This parameter cannot be specified at the same time as `timestamp`. **Example value**: "2018-09-01" |
|     `timestamp` | `string` | No | Second-level timestamp refers to the specific end time. For example, 1602504000 means 2020/10/12 20:00:00 (UTC +8 time zone). **Note**: This parameter cannot be specified at the same time as `date`. **Example value**: "1602504000" |
|     `timezone` | `string` | No | Time zone. Use IANA Time Zone Database standards such as Asia/Shanghai. - The time zone is fixed to UTC +0 throughout the day - The non-all-day time zone defaults to Asia/Shanghai **Example value**: "Asia/Shanghai" |
|   `user_ids` | `string[]` | No | Search filter item, list of user IDs of event invitees. After setting this field, the searched event contains at least one of the invitees. **Note**: The user ID type is consistent with the value of user_id_type. For user ID, please refer to User-related ID concepts. **Default value**: empty, indicating that the filter item is not set **Example value**: ["ou_e051986ab19f80d16b7b8d74f3f1235"] |
|   `room_ids` | `string[]` | No | Search filter items, meeting room ID list. After setting this field, the searched event contains at least one of the meeting rooms. **Default value**: empty, indicating that the filter item is not set **Example value**: ["omm_eada1d61a550955240c28757e7dec3af"] |
|   `chat_ids` | `string[]` | No | Search filter items, group ID list. After setting this field, the searched event contains at least one of these groups. For details about group ID, please refer to Group ID Description. **Default value**: empty, indicating that the filter item is not set **Example value**: ["oc_a0553eda9014c201e6969b478895c230"] | ### Request body example

{
    "query": "query words",
    "filter": {
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
        "user_ids": [
            "ou_e051986ab19f80d16b7b8d74f3f1235"
        ],
        "room_ids": [
            "omm_eada1d61a550955240c28757e7dec3af"
        ],
        "chat_ids": [
            "oc_a0553eda9014c201e6969b478895c230"
        ]
    }
}

## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `items` | `calendar.event[]` | List of events matched in the search. |
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
|     `recurring_event_id` | `string` | The event_id of the original recurring event corresponding to the exception event. |
|     `event_organizer` | `event_organizer` | Event organizer information. |
|       `user_id` | `string` | Event organizer user ID. |
|       `display_name` | `string` | Name of event organizer. |
|     `app_link` | `string` | app_link of the event, jump to a specific event. |
|     `attachments` | `attachment[]` | Event attachment. |
|       `file_token` | `string` | Attachment token. |
|       `file_size` | `string` | Attachment size. |
|       `is_deleted` | `boolean` | Whether to delete attachments. |
|       `name` | `string` | Attachment name. |
|   `page_token` | `string` | Page identifier, when has_more is true, a new page_token will also be returned. Otherwise, page_token will not be returned | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
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
                "event_organizer": {
                    "user_id": "ou_xxxxxx",
                    "display_name": "Jian Li"
                },
                "app_link": "https://applink.larkoffice.com/client/calendar/event/detail?calendarId=7039673579105026066&key=aeac9c56-aeb1-4179-a21b-02f278f59048&originalTime=0&startTime=1700496000",
                "attachments": [
                    {
                        "file_token": "xAAAAA",
                        "file_size": "2345",
                        "is_deleted": true,
                        "name": "demo.jpeg"
                    }
                ]
            }
        ],
        "page_token": "xxxxx"
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
