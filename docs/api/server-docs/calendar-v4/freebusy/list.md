---
title: "Query availability of the primary calendar"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/freebusy/list"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/calendar/v4/freebusy/list"
service: "calendar-v4"
resource: "freebusy"
section: "Calendar"
rate_limit: "1000 per minute、50 per second"
scopes:
  - "calendar:calendar"
  - "calendar:calendar.free_busy:read"
  - "calendar:calendar:readonly"
field_scopes:
  - "contact:user.employee_id:readonly"
updated: "1736822105000"
---

# Query availability of the primary calendar

Call this interface to query the busy/free information of the specified user's primary calendar, or to query the busy/free information of the specified meeting room.

> If you call this interface using the application identity, you need to ensure that the application has Bot capabilities.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/calendar/v4/freebusy/list |
| HTTP Method | POST |
| Rate Limit | 1000 per minute、50 per second |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `calendar:calendar` `calendar:calendar.free_busy:read` `calendar:calendar:readonly` |
| Required field scopes | > The response body of the API contains the following sensitive fields, and they will be returned only after corresponding scopes are added. If you do not need the fields, it is not recommended that you request the scopes. `contact:user.employee_id:readonly` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `user_id_type` | `string` | No | User ID categories **Example value**: open_id **Optional values are**:  - `open_id`: Identifies a user to an app. The same user has different Open IDs in different apps. How to get Open ID - `union_id`: Identifies a user to a tenant that acts as a developer. A user has the same Union ID in apps developed by the same developer, and has different Union IDs in apps developed by different developers. A developer can use Union ID to link the same user's identities in multiple apps.How to get Union ID - `user_id`: Identifies a user to a tenant. The same user has different User IDs in different tenants. In one single tenant, a user has the same User ID in all apps （including store apps）. User ID is usually used to communicate user data between different apps. How to get User ID  **Default value**: `open_id` **When the value is `user_id`, the following field scopes are required**: `contact:user.employee_id:readonly` | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `time_min` | `string` | Yes | The start time of the query period, in [RFC 3339](https://datatracker.ietf.org/doc/html/rfc3339) date_time format. **Note**: The time interval between time_min and time_max cannot be more than 90 days. **Example value**: "2020-10-28T12:00:00+08:00" |
| `time_max` | `string` | Yes | The end time of the query period, in [RFC 3339](https://datatracker.ietf.org/doc/html/rfc3339) date_time format. **Note**: The time interval between time_min and time_max cannot be more than 90 days. **Example value**: "2020-12-28T12:00:00+08:00" |
| `user_id` | `string` | No | User ID, you need to input an id that matches the query parameter user_id_type. For example, when user_id_type=open_id, you need to input the user's open_id. Refer to User-related ID concepts for understanding user IDs. **Note**: Either user_id or room_id needs to be entered. If both are entered at the same time, only user_id will take effect. **Example value**: "ou_xxxxxxxxxx" |
| `room_id` | `string` | No | Meeting room room_id. You can call the Query Meeting Room List interface or the Search for Meeting Room interface to get the corresponding meeting room's room_id. **Note**: Either user_id or room_id needs to be entered. If both are entered at the same time, only user_id will take effect. **Example value**: "omm_xxxxxxxxxx" |
| `include_external_calendar` | `boolean` | No | Whether to include the schedule in the bound third-party calendar. **Values**： - true (default value): Include - false: Do not include **Example value**: true |
| `only_busy` | `boolean` | No | Whether to only query busy schedule information. **Value**: - true (default): Yes, the query results do not include free schedules. - false: No, the query results include free schedules. **Example value**: true | ### Request body example

{
    "time_min": "2020-10-28T12:00:00+08:00",
    "time_max": "2020-12-28T12:00:00+08:00",
    "user_id": "ou_xxxxxxxxxx",
    "room_id": "omm_xxxxxxxxxx",
    "include_external_calendar": true,
    "only_busy": true
}

## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `freebusy_list` | `freebusy[]` | List of busy time periods within the requested time interval. |
|     `start_time` | `string` | Start time of the availability information. The value must be in the same format as specified in the date_time parameter in [RFC 3339](https://datatracker.ietf.org/doc/html/rfc3339). |
|     `end_time` | `string` | End time of the availability information. The value must be in the same format as specified in the date_time parameter in [RFC 3339](https://datatracker.ietf.org/doc/html/rfc3339). | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
        "freebusy_list": [
            {
                "start_time": "2020-10-28T22:30:00+08:00",
                "end_time": "2020-10-28T22:45:00+08:00"
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
| 429 | 190010 | current operation rate limited | The current operation is limited, usually because concurrent preemption of public resources fails. You can appropriately reduce the frequency of the current operation and try again. |
| 404 | 195100 | user is dismiss or not exist in the tenant | The current identity or specified user has resigned or is no longer in the tenant. Please check and change to the correct identity to call the interface. | For more error code information, see General error codes.
