---
title: "Get Event instances"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar-event/instances"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/calendar/v4/calendars/:calendar_id/events/:event_id/instances"
service: "calendar-v4"
resource: "calendar-event"
section: "Calendar"
rate_limit: "1000 per minute、50 per second"
scopes:
  - "calendar:calendar"
  - "calendar:calendar.event:read"
  - "calendar:calendar:readonly"
updated: "1736822155000"
---

# Get recurring schedule instances

Call this interface with the current identity (app or user) to obtain information about a specific recurring event in a specified calendar.

> The current identity is determined by the Token type of Header Authorization. tenant_access_token refers to the application identity, and user_access_token refers to the user identity.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/calendar/v4/calendars/:calendar_id/events/:event_id/instances |
| HTTP Method | GET |
| Rate Limit | 1000 per minute、50 per second |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `calendar:calendar` `calendar:calendar.event:read` `calendar:calendar:readonly` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `calendar_id` | `string` | Calendar ID. For information about calendar ID, please refer to Calendar ID Description. **Example value**: "larksuite.com_HF9U2MbibE8PPpjro6xjqa@group.calendar.larksuite.com" |
| `event_id` | `string` | Event ID. The event ID is returned when creating a event. You can also call the following interface to obtain the ID of a certain calendar. - Get event list - Search event **Example value**: "75d28f9b-e35c-4230-8a83-4a661497db54_0" | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `start_time` | `string` | Yes | Start time, Unix timestamp in seconds. This parameter and end_time are used to set the time range, that is, the query interval for repeated events is (start_time, end_time) **Note**: The time interval between start_time and end_time cannot exceed 2 years. **Example value**: 1631777271 |
| `end_time` | `string` | Yes | End time, Unix timestamp in seconds. This parameter and start_time are used to set the time range, that is, the query interval for repeated events is (start_time, end_time) **Note**: The time interval between start_time and end_time cannot exceed 2 years. **Example value**: 1631777271 |
| `page_size` | `int` | No | The maximum number of events returned in one call. **Example value**: 10 **Default value**: `50` **Data validation rules**: - Value range: `10` ～ `500` |
| `page_token` | `string` | No | Page identifier. It is not filled in the first request, indicating traversal from the beginning; when there will be more groups, the new page_token will be returned at the same time, and the next traversal can use the page_token to get more groups **Example value**: eVQrYzJBNDNONlk4VFZBZVlSdzlKdFJ4bVVHVExENDNKVHoxaVdiVnViQT0= | ## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `items` | `instance[]` | Event instance list of recurring events. |
|     `event_id` | `string` | Event instance ID. **Note**: The ID of a recurring event instance is different from other event IDs. Its ID contains the original time of the instance (Original time), and the data format is a second-level timestamp. For example: `2cf525f0-1e67-4b04-ad4d-30b7f003903c_1713168000`, where `1713168000` is the original time of the instance. |
|     `summary` | `string` | Event topic |
|     `description` | `string` | Event description. |
|     `start_time` | `time_info` | Event start time. |
|       `date` | `string` | Start time, only use this field for all-day events, [RFC 3339](https://datatracker.ietf.org/doc/html/rfc3339) format, for example, 2018-09-01. |
|       `timestamp` | `string` | Second-level timestamp refers to the specific start time of the event. For example, 1602504000 means 2020/10/12 20:00:00 (UTC +8 time zone). |
|       `timezone` | `string` | Time zone. Use the IANA Time Zone Database standard. |
|     `end_time` | `time_info` | Event end time. |
|       `date` | `string` | End time, only use this field for all-day events, [RFC 3339](https://datatracker.ietf.org/doc/html/rfc3339) format, for example, 2018-09-01. |
|       `timestamp` | `string` | Second-level timestamp refers to the specific end time of the event. For example, 1602504000 means 2020/10/12 20:00:00 (UTC +8 time zone). |
|       `timezone` | `string` | Time zone. Use the IANA Time Zone Database standard. |
|     `status` | `string` | Event status. **Optional values are**:  - `tentative`: No response - `confirmed`: Confirmed - `cancelled`: The schedule has been cancelled.  |
|     `is_exception` | `boolean` | Whether the event is an exception to a repeating event. To learn about the exception event, see Exception event. |
|     `app_link` | `string` | app_link of the event, jump to a specific event. |
|     `location` | `event_location` | Event location. |
|       `name` | `string` | Location name. |
|       `address` | `string` | Location Address. |
|       `latitude` | `number(float)` | Location coordinates and latitude information, for domestic locations, the GCJ-02 standard is used, and the WGS84 standard is used for overseas locations |
|       `longitude` | `number(float)` | Location coordinate longitude information, for domestic locations, use the GCJ-02 standard, and overseas locations use the WGS84 standard |
|   `page_token` | `string` | Page identifier, when has_more is true, a new page_token will also be returned. Otherwise, page_token will not be returned |
|   `has_more` | `boolean` | Whether the response body has more parameters | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
        "items": [
            {
                "event_id": "75d28f9b-e35c-4230-8a83-4a661497db54_1602504000",
                "summary": "Event topic",
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
                "app_link": "https://applink.larkoffice.com/client/calendar/event/detail?calendarId=7039673579105026066&key=aeac9c56-aeb1-4179-a21b-02f278f59048&originalTime=0&startTime=1700496000",
                "location": {
                    "name": "Shanghai",
                    "address": "Xuhui District",
                    "latitude": 23.4444,
                    "longitude": 23.4444
                }
            }
        ],
        "page_token": "eVQrYzJBNDNONlk4VFZBZVlSdzlKdFJ4bVVHVExENDNKVHoxaVdiVnViQT0=",
        "has_more": true
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
| 403 | 193002 | no permission to operate event | No permission to operate. You need to make sure you have permission to edit the calendar and event. |
| 403 | 193003 | event is deleted | The event has been deleted. You need to check and change to the correct event ID. |
| 404 | 195100 | user is dismiss or not exist in the tenant | The current identity or specified user has resigned or is no longer in the tenant. Please check and change to the correct identity to call the interface. | For more error code information, see General error codes.
