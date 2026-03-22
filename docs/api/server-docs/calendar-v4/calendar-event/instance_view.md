---
title: "Query event view"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar-event/instance_view"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/calendar/v4/calendars/:calendar_id/events/instance_view"
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
updated: "1736822159000"
---

# Query event view

Call this interface with user identity to query the event view under a specified calendar. Unlike Getting event list, the current interface will expand into multiple event instances according to the repetitiveness rules of the recurring event, and return the corresponding event instance information according to the queried time interval.

> - The current identity is determined by the Token type of Header Authorization. tenant_access_token refers to the application identity, and user_access_token refers to the user identity.
> - The current identity must have reader, writer or owner permissions on the calendar, and only supports obtaining the event list of primary and shared calendars. Google and exchange calendars are not supported yet. You can call the Query Calendar Information interface to obtain the current identity’s access rights to the calendar and calendar type information.
> - The upper limit of the number of event instances obtained by querying the event view is less than 1,000.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/calendar/v4/calendars/:calendar_id/events/instance_view |
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
| `calendar_id` | `string` | Calendar ID. For information about calendar ID, please refer to Calendar ID Description. **Example value**: "larksuite.com_HF9U2MbibE8PPpjro6xjqa@group.calendar.larksuite.com" | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `start_time` | `string` | Yes | Start time, Unix timestamp in seconds. This parameter and end_time are used to set the time range of the query. **Note**: The time interval between start_time and end_time needs to be less than 40 days. **Example value**: 1631777271 |
| `end_time` | `string` | Yes | End time, Unix timestamp in seconds. This parameter is used with start_time to set the time range of the query. **Note**: The time interval between start_time and end_time needs to be less than 40 days. **Example value**: 1631777271 |
| `user_id_type` | `string` | No | User ID categories **Example value**: open_id **Optional values are**:  - `open_id`: Identifies a user to an app. The same user has different Open IDs in different apps. How to get Open ID - `union_id`: Identifies a user to a tenant that acts as a developer. A user has the same Union ID in apps developed by the same developer, and has different Union IDs in apps developed by different developers. A developer can use Union ID to link the same user's identities in multiple apps.How to get Union ID - `user_id`: Identifies a user to a tenant. The same user has different User IDs in different tenants. In one single tenant, a user has the same User ID in all apps （including store apps）. User ID is usually used to communicate user data between different apps. How to get User ID  **Default value**: `open_id` **When the value is `user_id`, the following field scopes are required**: `contact:user.employee_id:readonly` | ## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `items` | `instance[]` | Event instance list. |
|     `event_id` | `string` | Event instance ID. Event information can be queried, updated, or deleted later through this ID. |
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
|     `status` | `string` | Event status. **Optional values are**:  - `tentative`: No response - `confirmed`: Confirmed - `cancelled`: The event has been cancelled.  |
|     `is_exception` | `boolean` | Whether the event is an exception to a repeating event. To learn about the exception event, see Exception event. |
|     `app_link` | `string` | App_link of the event, jump to a specific event. |
|     `organizer_calendar_id` | `string` | The calendar ID of the event organizer. For information about calendar ID, please refer to Calendar ID Description. |
|     `vchat` | `vchat` | Video conferencing information. Only takes effect if the event has at least one invitee. |
|       `vc_type` | `string` | Video conferencing type. **Optional values are**:  - `vc`: Lark Video Conference - `third_party`: Third-party link video conferencing - `no_meeting`: No video conferencing - `lark_live`: Lark Live - `unknown`: Unknown type  |
|       `icon_type` | `string` | Third-party video conferencing icon types. **Optional values are**:  - `vc`: Lark video conferencing icon - `live`: Live video conference icon - `default`: Default icon  |
|       `description` | `string` | Third-party video conferencing copywriting. |
|       `meeting_url` | `string` | Video conferencing URL. |
|     `visibility` | `string` | Event visibility. It will only take effect for all invitees when creating a new event, and subsequent modifications to this attribute will only take effect for the current identity. **Optional values are**:  - `default`: Default permission to only show others if they are busy - `public`: Open, showing event details - `private`: Private, visible only to yourself  |
|     `attendee_ability` | `string` | Invitee permissions. **Optional values are**:  - `none`: Cannot edit event, cannot invite other invitees, cannot view invitee list - `can_see_others`: Can't edit event, can't invite other invitees, can view invitee list - `can_invite_others`: Can't edit event, can invite other invitees, can view invitee list - `can_modify_event`: You can edit the event, invite other invitees, and view the list of invitees  |
|     `free_busy_status` | `string` | The busy and busy status occupied by the event. It will only take effect for all invitees when creating a new event, and subsequent modifications to this attribute will only take effect for the current identity. **Optional values are**:  - `busy`: busy - `free`: idle  |
|     `location` | `event_location` | Event location. |
|       `name` | `string` | Location name. |
|       `address` | `string` | Location Address. |
|       `latitude` | `number(float)` | Location coordinates and latitude information, for domestic locations, the GCJ-02 standard is used, and the WGS84 standard is used for overseas locations |
|       `longitude` | `number(float)` | Location coordinate longitude information, for domestic locations, use the GCJ-02 standard, and overseas locations use the WGS84 standard |
|     `color` | `int` | Event color, represented by an int32 of color RGB values. **illustrate**: - Only takes effect for the current identity. - When the value is 0 or -1, it means to follow the calendar color by default. - When displayed on the client, it will be mapped to the closest color on the color palette. |
|     `recurring_event_id` | `string` | The event_id of the original duplicate event of the exception event. |
|     `event_organizer` | `event_organizer` | Event organizer information. |
|       `user_id` | `string` | Event organizer user ID. |
|       `display_name` | `string` | Name of event organizer. |
|     `attendees` | `calendar.event.attendee[]` | Event invitee information, currently only returns to the conference room. If you need other types of invitee information, please use Get event invitee list interface. |
|       `type` | `string` | Invitee type. **Optional values are**:  - `user`: user - `chat`: group - `resource`: meeting room - `third_party`: email  |
|       `attendee_id` | `string` | Invitee ID. The unique identifier of the event invitee in the current event. |
|       `rsvp_status` | `string` | Invitee RSVP status, that is, event reply status. **Optional values are**:  - `needs_action`: The invitee has not replied to the status, or indicates that the meeting room reservation is in progress. - `accept`: The invitee replies to accept, or indicates that the meeting room reservation is successful - `tentative`: Invitee reply to be determined - `decline`: Invitees respond with rejection, or indicate that the meeting room reservation failed - `removed`: Invitees or meeting rooms have been removed from the event  |
|       `is_optional` | `boolean` | Whether the invitee is optional, this parameter value does not take effect on the group members. |
|       `is_organizer` | `boolean` | Whether the invitee is the event organizer. |
|       `is_external` | `boolean` | Whether the invitee is an external invitee. Editing is not supported by external invitees. |
|       `display_name` | `string` | Invitee name. |
|       `chat_members` | `attendee_chat_member[]` | Group members in the group, valid when the invitee type is chat. Group members do not support editing. |
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
|       `approval_reason` | `string` | Reason for approval of conference room. | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
        "items": [
            {
                "event_id": "75d28f9b-e35c-4230-8a83-4a661497db54_1602504000",
                "summary": "Agenda topic",
                "description": "desc",
                "start_time": {
                    "date": "2018-09-01",
                    "timestamp": "1602504000",
                    "timezone": "Asia/Shanghai"
                },
                "end_time": {
                    "date": "2018-09-01",
                    "timestamp": "1602504400",
                    "timezone": "Asia/Shanghai"
                },
                "status": "confirmed",
                "is_exception": false,
                "app_link": "https://applink.larksuite.com/client/calendar/event/detail?calendarId=7039673579105026066&key=aeac9c56-aeb1-4179-a21b-02f278f59048&originalTime=0&startTime=1700496000",
                "organizer_calendar_id": "larksuite.com_HF9U2MbibE8PPpjro6xjqa@group.calendar.larksuite.com",
                "vchat": {
                    "vc_type": "vc",
                    "icon_type": "vc",
                    "description": "Start a video conference",
                    "meeting_url": "vc.larksuite.com/j/152568231"
                },
                "visibility": "Default",
                "attendee_ability": "None",
                "free_busy_status": "busy",
                "location": {
                    "name": "Location name",
                    "address": "Location Address",
                    "latitude": 1.100000023841858,
                    "longitude": 2.200000047683716
                },
                "color": 0,
                "recurring_event_id": "75d28f9b-e35c-4230-8a83-4a661497db54_0",
                "event_organizer": {
                    "user_id": "ou_xxxxxx",
                    "display_name": "Jian Li"
                },
                "attendees": [
                    {
                        "type": "user",
                        "attendee_id": "user_xxxxxx",
                        "rsvp_status": "Accept",
                        "is_optional": false,
                        "is_organizer": true,
                        "is_external": false,
                        "display_name": "San Zhang",
                        "chat_members": [
                            {
                                "rsvp_status": "Accept",
                                "is_optional": true,
                                "display_name": "xxxx",
                                "is_organizer": true,
                                "is_external": false
                            }
                        ],
                        "user_id": "ou_xxxxxxxx",
                        "chat_id": "oc_a0553eda9014c201e6969b478895c230",
                        "room_id": "omm_83d09ad4f6896e02029a6a075f71c9d1",
                        "third_party_email": "test@example.com",
                        "operate_id": "4d7a3c6g",
                        "resource_customization": [
                            {
                                "index_key": "16281481596100",
                                "input_content": "XXX",
                                "options": [
                                    {
                                        "option_key": "16281481596185",
                                        "others_content": "XXXX"
                                    }
                                ]
                            }
                        ],
                        "approval_reason": "Reason for application approval"
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
| 429 | 190010 | current operation rate limited | The current operation is limited, usually because concurrent preemption of public resources fails. You can appropriately reduce the frequency of the current operation and try again. |
| 403 | 190011 | tenant encrypt key has been deleted | The autonomous key in the encryption and decryption state is deleted, and the data encrypted by the key is unavailable. |
| 403 | 190012 | tenant decrypt key has been deleted | Only the autonomous key in the decryption state is deleted, and the data encrypted by the key is unavailable. |
| 404 | 191000 | calendar not found | Calendar not found. You need to check and change to the correct calendar ID. |
| 400 | 191001 | invalid calendar_id | The calendar_id is invalid. You need to check and change it to the correct calendar ID. |
| 403 | 191002 | no calendar access_role | The current identity does not have permission to access the calendar. If you want to query a calendar, you need to ensure that the current identity has permission to access the calendar. |
| 403 | 191003 | calendar is deleted | The calendar has been deleted. You need to check and change to the correct calendar ID. |
| 403 | 191004 | invalid calendar type | The calendar type is incorrect. You can call the query calendar information interface to obtain the calendar type information, and then ensure that the calendar type is applicable to the current interface. |
| 400 | 193103 | requested time exceeds the allowed range | Time range limit exceeded. You need to check the start and end timespan passed in. The timespan cannot exceed 40 days. |
| 400 | 193104 | exceeded the maximum instance number | The maximum number of instances has been exceeded. The number of schedule instances in the time interval entered exceeds 1000. It is recommended to reduce the time interval range. |
| 404 | 195100 | user is dismiss or not exist in the tenant | The current identity or specified user has resigned or is no longer in the tenant. Please check and change to the correct identity to call the interface. | For more error code information, see General error codes.
