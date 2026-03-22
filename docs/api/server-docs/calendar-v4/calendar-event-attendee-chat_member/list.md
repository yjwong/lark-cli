---
title: "Obtain the list of members of group invitees of an event"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar-event-attendee-chat_member/list"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/calendar/v4/calendars/:calendar_id/events/:event_id/attendees/:attendee_id/chat_members"
service: "calendar-v4"
resource: "calendar-event-attendee-chat_member"
section: "Calendar"
rate_limit: "1000 per minute、50 per second"
scopes:
  - "calendar:calendar"
  - "calendar:calendar.event:read"
  - "calendar:calendar:readonly"
field_scopes:
  - "contact:user.employee_id:readonly"
updated: "1736835322000"
---

# Obtain the list of members of group invitees of an event

Call this interface with the current identity (app or user) to get the group member list of the group type invitees in the event.

> - The current identity is determined by the Token type of Header Authorization. tenant_access_token refers to the application identity, and user_access_token refers to the user identity.
> - If you use the application identity to call this interface, you need to ensure that the application turns on bot capability.
> - The current identity must have permission to view the invitee list of the event, that is, the current identity needs to be the organizer of the event, or a invitee of the event and the event has the **Invitees can view the list of invitees** permission set on the event. You can call the Get event interface to obtain the invitee permissions (attendee_ability) of the event.
> - The current identity must be in the group member list. If it is not in the group member list, the call will fail and an error message will be returned.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/calendar/v4/calendars/:calendar_id/events/:event_id/attendees/:attendee_id/chat_members |
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
| `calendar_id` | `string` | Calendar ID. For details, see Calendar-related IDs. **Example value**: "larksuite.com_xxxxxxxxxx@group.calendar.larksuite.com" |
| `event_id` | `string` | Event ID. The event ID is returned when creating a event. You can also call the following interface to obtain the ID of a certain calendar. - Get event list - Search event **Example value**: "xxxxxxxxx_0" |
| `attendee_id` | `string` | Group type invitee ID. When adding a event invitee, the invitee ID (attendee_id) will be returned. You can also call Get event invitee list Interface to query the invitee ID of the specified event. **Example value**: "chat_xxxxxx" | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `page_token` | `string` | No | Page identifier. It is not filled in the first request, indicating traversal from the beginning; when there will be more groups, the new page_token will be returned at the same time, and the next traversal can use the page_token to get more groups **Example value**: 23jhysaxxxxsysy |
| `page_size` | `int` | No | The maximum number of group members returned in one request. **Example value**: 10 **Default value**: `20` **Data validation rules**: - Maximum value: `100` |
| `user_id_type` | `string` | No | User ID categories **Example value**: open_id **Optional values are**:  - `open_id`: Identifies a user to an app. The same user has different Open IDs in different apps. How to get Open ID - `union_id`: Identifies a user to a tenant that acts as a developer. A user has the same Union ID in apps developed by the same developer, and has different Union IDs in apps developed by different developers. A developer can use Union ID to link the same user's identities in multiple apps.How to get Union ID - `user_id`: Identifies a user to a tenant. The same user has different User IDs in different tenants. In one single tenant, a user has the same User ID in all apps （including store apps）. User ID is usually used to communicate user data between different apps. How to get User ID  **Default value**: `open_id` **When the value is `user_id`, the following field scopes are required**: `contact:user.employee_id:readonly` | ## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `items` | `calendar.event.attendee.chat_member[]` | Group members. This field is valid when the type field is chat. Group members have no edit scope. |
|     `rsvp_status` | `string` | RSVP status of the event invitee. **Optional values are**:  - `needs_action`: The event invitee hasn't replied or a room is being booked. - `accept`: The event invitee accepted or a room was successfully booked. - `tentative`: The event invitee replied pending. - `decline`: The event invitee declined or a room failed to be booked. - `removed`: The event invitee or room has been removed from the event.  |
|     `is_optional` | `boolean` | Whether the event invitee is optional. |
|     `display_name` | `string` | Event invitee name. |
|     `open_id` | `string` | Invitees open_id. See User-related ID concept **Sample value**: "ou_xxxxxxxx" |
|     `is_organizer` | `boolean` | Whether the event invitee is an event organizer. |
|     `is_external` | `boolean` | Whether the event invitee is external. |
|   `has_more` | `boolean` | Whether the response body has more parameters |
|   `page_token` | `string` | Page identifier, when has_more is true, a new page_token will also be returned. Otherwise, page_token will not be returned | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
        "items": [
            {
                "rsvp_status": "needs_action",
                "is_optional": true,
                "display_name": "Zhang San",
                "open_id": "ou_143669c5a53647f00f6c80a0253aa68b",
                "is_organizer": true,
                "is_external": false
            }
        ],
        "has_more": false,
        "page_token": "73TyueXXXXX"
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
| 400 | 193000 | invalid event_id | The event_id is invalid. You need to check and change it to the correct event ID. |
| 404 | 193001 | event not found | Event not found. You need to make sure you passed in the correct event ID. |
| 403 | 193002 | no permission to operate event | No permission to operate. You need to make sure you have permission to edit the calendar and schedule. |
| 403 | 193003 | event is deleted | The event has been deleted. You need to check and change to the correct event ID. |
| 404 | 194000 | attendee not found | No participant was found. You need to make sure the participant parameters are filled in correctly. |
| 403 | 194001 | no permission to list event attendees | No operation permission. Check whether calendar_id is the organizer's calendar of the current schedule, or whether the organizer has permission to view participants. |
| 403 | 194002 | no permission to create event attendees | No operation permission. Check whether calendar_id is the organizer's calendar of the current schedule, or whether the organizer has permission to invite participants. |
| 403 | 194003 | no permission to delete event attendees | No operation permission. Check whether calendar_id is the organizer calendar of the current schedule. |
| 400 | 194004 | invalid attendee type | The participant type is invalid. Please check whether the participant type is entered correctly. |
| 404 | 195100 | user is dismiss or not exist in the tenant | The current identity or specified user has resigned or is no longer in the tenant. Please check and change to the correct identity to call the interface. | For more error code information, see General error codes.
