---
title: "Unbind meeting chat"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar-event-meeting_chat/delete"
method: "DELETE"
api_path: "https://open.larksuite.com/open-apis/calendar/v4/calendars/:calendar_id/events/:event_id/meeting_chat"
service: "calendar-v4"
resource: "calendar-event-meeting_chat"
section: "Calendar"
rate_limit: "1000 per minute、50 per second"
scopes:
  - "calendar:calendar"
  - "calendar:calendar.event:update"
updated: "1736835292000"
---

# Unbind meeting chat

Call this interface in the current identity (app or user) to unbind the created meeting group for the event.

> - The current identity is determined by the Token type of Header Authorization. tenant_access_token refers to the application identity, and user_access_token refers to the user identity.
> - If you use the application identity to call this interface, you need to ensure that the application turns on bot capability.
> - The calendar where the event is located needs to be the primary calendar of the current identity and have the calendar's writer permissions. You can call the Query primary calendar information interface to obtain the primary calendar information of the current identity.
> - The current operator needs to be the group owner of the conference group.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/calendar/v4/calendars/:calendar_id/events/:event_id/meeting_chat |
| HTTP Method | DELETE |
| Rate Limit | 1000 per minute、50 per second |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `calendar:calendar` `calendar:calendar.event:update` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `calendar_id` | `string` | The calendar ID where the event is located. **Example value**: "larksuite.com_xxx@group.calendar.larksuite.com" |
| `event_id` | `string` | Event ID. **Example value**: "75d28f9b-e35c-4230-8a83-xxxx_0" | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `meeting_chat_id` | `string` | Yes | Meeting group ID. The group ID is returned when the group is created. **Example value**: oc_xx | ## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {}
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
| 404 | 195100 | user is dismiss or not exist in the tenant | The current identity or specified user has resigned or is no longer in the tenant. Please check and change to the correct identity to call the interface. |
| 403 | 195107 | no permission unbind meeting chat | You do not have permission to unbind the conference group. You need to check whether you are the group owner. |
| 400 | 195108 | chat not match with event | The group ID and schedule entered do not match, and the conference group cannot be unbound. Please check whether the input parameters are filled in correctly. | For more error code information, see General error codes.
