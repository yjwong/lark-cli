---
title: "Get Event"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar-event/get"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/calendar/v4/calendars/:calendar_id/events/:event_id"
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

# Obtain an event

Call this interface with the current identity (app or user) to obtain the event information within a specified calendar, including the title of the event, time period, video conference information, public scope, and invitee rights, etc.

> - The current identity is determined by the Token type of Header Authorization. tenant_access_token refers to the application identity, and user_access_token refers to the user identity.
> - The current identity must have reader, writer, or owner permissions on the calendar. You can call the Query calendar information interface to obtain the current identity's access rights to the calendar.
> - If the timestamp suffix of event_id is non-0, it means that the event is an exception event of a repeating event, and you can obtain the information of the exception event from this. For example, if the timestamp suffix of event_id `xxxxxxxxx_1602504000` is non-0, it means that this is an exception event that is a repeating event.
> - You can obtain the time information of the exception event through the timestamp suffix of event_id. For example, an exception event with event_id `xxxxxxxxx_1602504000` has a timestamp of 160250400. For instructions on exception events, please see Schedule Resource Introduction.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/calendar/v4/calendars/:calendar_id/events/:event_id |
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
| `calendar_id` | `string` | The calendar ID where the event is located. For information about calendar ID, please refer to Calendar ID Description. **Example value**: "larksuite.com_xxxxxxxxxx@group.calendar.larksuite.com" |
| `event_id` | `string` | Event ID. The event ID is returned when creating a event. You can also call the following interface to obtain the ID of a certain calendar. - Get event list - Search event **Example value**: "xxxxxxxxx_0" | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `need_meeting_settings` | `boolean` | No | Do you need to return to the pre-meeting settings of Lark Video Conferencing (VC)? The following conditions must be met to obtain the returned results: - The meeting type (vc_type) of the event needs to be vc. - Requires editing permissions for event. **Optional values are**: - true - false (default) **Example value**: false |
| `need_attendee` | `boolean` | No | Do you need to return invitee information? **Optional values are**: - true - false (default) **Example value**: false |
| `max_attendee_num` | `int` | No | The maximum number of invitees returned. Call Get event invitee list to get complete invitee information of the event. **Example value**: false **Default value**: `10` **Data validation rules**: - Maximum value: `100` |
| `user_id_type` | `string` | No | User ID categories **Example value**: open_id **Optional values are**:  - `open_id`: Identifies a user to an app. The same user has different Open IDs in different apps. How to get Open ID - `union_id`: Identifies a user to a tenant that acts as a developer. A user has the same Union ID in apps developed by the same developer, and has different Union IDs in apps developed by different developers. A developer can use Union ID to link the same user's identities in multiple apps.How to get Union ID - `user_id`: Identifies a user to a tenant. The same user has different User IDs in different tenants. In one single tenant, a user has the same User ID in all apps （including store apps）. User ID is usually used to communicate user data between different apps. How to get User ID  **Default value**: `open_id` **When the value is `user_id`, the following field scopes are required**: `contact:user.employee_id:readonly` | ## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `event` | `calendar.event` | Event details. |
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
|       `icon_type` | `string` | Icon type of the third-party video conference. If this parameter is left empty, the default icon is displayed. **Optional values are**:  - `vc`: Lark Video Conference icon - `live`: Livestream video conference icon - `default`: Default icon  |
|       `description` | `string` | Text for the third-party video conference. |
|       `meeting_url` | `string` | Video conference URL. |
|       `meeting_settings` | `meeting_settings` | Pre-meeting settings for Lark Video Conferencing (VC). |
|         `owner_id` | `string` | The user ID information of the conference owner. |
|         `join_meeting_permission` | `string` | Determine who can join the meeting. **Optional values are**:  - `anyone_can_join`: Any user can join the meeting - `only_organization_employees`: Only users that from the same organization can join the meeting - `only_event_attendees`: Only event attendees can join the meeting  |
|         `password` | `string` | The moderator's user ID information. |
|         `assign_hosts` | `string[]` | The moderator's user ID information. |
|         `auto_record` | `boolean` | Whether to enable automatic recording. |
|         `open_lobby` | `boolean` | Whether to allow event invitees to initiate meetings. |
|         `allow_attendees_start` | `boolean` | Whether to allow event invitees to initiate meetings. |
|     `visibility` | `string` | Event visibility. It only takes effect for all invitees when creating a new event, and subsequent modifications to this attribute only take effect for the current identity. **Optional values are**:  - `default`: Default range, which depends on the calendar visibility. Only the availability status is visible to other people by default. - `public`: Public. Event details are displayed. - `private`: Private. Details are visible only to the current identity.  |
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
|     `recurrence` | `string` | Recurrence rule for recurring events. For details, see [RFC 5545](https://datatracker.ietf.org/doc/html/rfc5545#section-3.3.10). |
|     `status` | `string` | Event status. **Optional values are**:  - `tentative`: Not replied - `confirmed`: Confirmed - `cancelled`: Canceled  |
|     `is_exception` | `boolean` | Whether the event is an exception to a repeating event. To learn about the exception event, see Exception event. |
|     `recurring_event_id` | `string` | The event_id of the original recurring event of an exceptional event. |
|     `create_time` | `string` | Event create time (timestamp in seconds). |
|     `schemas` | `schema[]` | Custom information of the event. This field controls the UI display on the event details page. |
|       `ui_name` | `string` | UI name. Valid values: - ForwardIcon: Event forwarding button - MeetingChatIcon: Meeting group chat button - MeetingMinutesIcon: Meeting minutes button - MeetingVideo: Video conference area - RSVP: Accept/decline/pending area - Attendee: Event invitees area - OrganizerOrCreator: Organizer/creator area |
|       `ui_status` | `string` | Custom status of UI items. **Optional values are**:  - `hide`: Hidden - `readonly`: Read-only - `editable`: Editable - `unknown`: Unknown. This is only used for compatibility in reading.  |
|       `app_link` | `string` | Link that is redirected to after the button is tapped/clicked. |
|     `event_organizer` | `event_organizer` | Event organizer information. |
|       `user_id` | `string` | Event organizer user ID. |
|       `display_name` | `string` | Name of event organizer. |
|     `app_link` | `string` | app_link of the event, jump to a specific event. |
|     `attendees` | `calendar.event.attendee[]` | Event invitee Information |
|       `type` | `string` | Invitee type. **Optional values are**:  - `user`: user - `chat`: group - `resource`: meeting room - `third_party`: email  |
|       `attendee_id` | `string` | Invitee ID, the unique identifier of the event invitee in the current event. |
|       `rsvp_status` | `string` | Invitee RSVP status, that is, event reply status. **Optional values are**:  - `needs_action`: The invitee has not replied to the status, or indicates that the meeting room reservation is in progress. - `accept`: The invitee replies to accept, or indicates that the meeting room reservation is successful - `tentative`: Invitee reply to be determined - `decline`: Invitees respond with rejection, or indicate that the meeting room reservation failed - `removed`: Invitees or meeting rooms have been removed from the event  |
|       `is_optional` | `boolean` | Whether the invitee is optional, this parameter value does not take effect on the group members. |
|       `is_organizer` | `boolean` | Is the invitee an agenda organizer? |
|       `is_external` | `boolean` | Invitees are external invitees; external invitees do not support editing. |
|       `display_name` | `string` | Invitee name |
|       `chat_members` | `attendee_chat_member[]` | Group members in a group, valid when type is chat; group members do not support editing. |
|         `rsvp_status` | `string` | Invitee RSVP status, that is, event reply status. **Optional values are**:  - `needs_action`: The invitee has not replied to the status, or indicates that the meeting room reservation is in progress. - `accept`: The invitee replies to accept, or indicates that the meeting room reservation is successful - `tentative`: Invitee reply to be determined - `decline`: Invitees respond with rejection, or indicate that the meeting room reservation failed - `removed`: Invitees or meeting rooms have been removed from the event  |
|         `is_optional` | `boolean` | Whether the invitee is optional. |
|         `display_name` | `string` | Invitee name. |
|         `is_organizer` | `boolean` | Is the invitee an agenda organizer? |
|         `is_external` | `boolean` | Whether the invitee is an external invitee. |
|       `user_id` | `string` | User type The user ID of the invitee. The ID type is consistent with the value of user_id_type. For information about user IDs, see User-related ID concepts. **Note**: When is_external returns true, this field will only return open_id or union_id. |
|       `chat_id` | `string` | Group type The group ID of the invitee. For details about group ID, please refer to Group ID Description. |
|       `room_id` | `string` | Room Type The room ID of the invitee. |
|       `third_party_email` | `string` | External email type The email address of the invitee. |
|       `operate_id` | `string` | If the event was created using App Identity, specify the room contact ID when adding the room. The ID type is consistent with the value of user_id_type. |
|       `resource_customization` | `calendar.attendee.resource_customization[]` | Personalized configuration of meeting rooms. |
|         `index_key` | `string` | A unique ID for each configuration. |
|         `input_content` | `string` | Fill-in-the-blank type value. |
|         `options` | `customization.option[]` | Options for each configuration. |
|           `option_key` | `string` | A unique ID for each option. |
|           `others_content` | `string` | Values for other option types. |
|     `has_more_attendee` | `boolean` | Are there more invitees? |
|     `attachments` | `attachment[]` | Event attachment. |
|       `file_token` | `string` | Attachment token. |
|       `file_size` | `string` | Attachment size. |
|       `name` | `string` | Attachment name. | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
        "event": {
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
                "meeting_url": "https://example.com",
                "meeting_settings": {
                    "owner_id": "ou_7d8a6e6df7621556ce0d21922b676706ccs",
                    "join_meeting_permission": "only_organization_employees",
                    "password": "971024",
                    "assign_hosts": [
                        "ou_7d8a6e6df7621556ce0d21922b676706ccs"
                    ],
                    "auto_record": false,
                    "open_lobby": true,
                    "allow_attendees_start": true
                }
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
            "attendees": [
                {
                    "type": "user",
                    "attendee_id": "user_xxxxxx",
                    "rsvp_status": "Accept",
                    "is_optional": false,
                    "is_organizer": true,
                    "is_external": false,
                    "display_name": "Zhang San",
                    "chat_members": [
                        {
                            "rsvp_status": "Accept",
                            "is_optional": true,
                            "display_name": "zhangsan",
                            "is_organizer": false,
                            "is_external": false
                        }
                    ],
                    "user_id": "ou_ 8bfcecaf02c9d8d026de984db04cf5b9",
                    "chat_id": "oc_ a0553eda9014c201e6969b478895c230",
                    "room_id": "omm_ 83d09ad4f6896e02029a6a075f71c9d1",
                    "third_party_email": "test@example.com",
                    "operate_id": "4D7A3C6G",
                    "resource_customization": [
                        {
                            "index_key": "16281481596100",
                            "input_content": "16281481596100",
                            "options": [
                                {
                                    "option_key": "16281481596185xxx",
                                    "others_content": "XXXX"
                                }
                            ]
                        }
                    ]
                }
            ],
            "has_more_attendee": false,
            "attachments": [
                {
                    "file_token": "xAAAAA",
                    "file_size": "2345",
                    "name": "demo.jpeg"
                }
            ]
        }
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
| 429 | 190010 | current operation rate limited | The current operation is limited, usually because concurrent preemption of public resources fails. You can appropriately reduce the frequency of the current operation and try again. |
| 403 | 190011 | tenant encrypt key has been deleted | The autonomous key in the encryption and decryption state is deleted, and the data encrypted by the key is unavailable. |
| 403 | 190012 | tenant decrypt key has been deleted | Only the autonomous key in the decryption state is deleted, and the data encrypted by the key is unavailable. |
| 404 | 191000 | calendar not found | Calendar not found. You need to check and change to the correct calendar ID. |
| 400 | 191001 | invalid calendar_id | The calendar_id is invalid. You need to check and change it to the correct calendar ID. |
| 403 | 191002 | no calendar access_role | The current identity does not have permission to access the calendar. If you want to query a calendar, you need to ensure that the current identity has permission to access the calendar. |
| 403 | 191003 | calendar is deleted | The calendar has been deleted. You need to check and change to the correct calendar ID. |
| 403 | 191004 | invalid calendar type | The calendar type is incorrect. You can call the query calendar information interface to obtain the calendar type information, and then ensure that the calendar type is applicable to the current interface. |
| 400 | 193000 | invalid event_id | The event_id is invalid. You need to check and change it to the correct event ID. |
| 404 | 193001 | event not found | Event not found. You need to make sure you passed in the correct event ID. |
| 403 | 193002 | no permission to operate event | No permission to operate. You need to make sure you have permission to edit the calendar and event. |
| 403 | 193003 | event is deleted | The event has been deleted. You need to check and change to the correct event ID. |
| 403 | 194001 | no permission to list event attendees | No operation permission. Check whether calendar_id is the organizer's calendar of the current event, or whether the organizer has permission to view participants. |
| 404 | 195100 | user is dismiss or not exist in the tenant | The current identity or specified user has resigned or is no longer in the tenant. Please check and change to the correct identity to call the interface. | For more error code information, see General error codes.
