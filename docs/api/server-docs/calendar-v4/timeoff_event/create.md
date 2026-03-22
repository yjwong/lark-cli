---
title: "Create Timeoff Event"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/timeoff_event/create"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/calendar/v4/timeoff_events"
service: "calendar-v4"
resource: "timeoff_event"
section: "Calendar"
rate_limit: "1000 per minute、50 per second"
scopes:
  - "calendar:time_off:create"
  - "calendar:timeoff"
field_scopes:
  - "contact:user.employee_id:readonly"
updated: "1736837316000"
---

# Create a leave event

Call this interface to create a leave event for a specified user. Leave events are divided into regular events and all-day events. After creating a leave event, the user's personal signature page will display leave information during the leave time.

> To use the application identity to call this interface, you need to ensure that the application turns on bot capability.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/calendar/v4/timeoff_events |
| HTTP Method | POST |
| Rate Limit | 1000 per minute、50 per second |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `calendar:time_off:create` `calendar:timeoff` |
| Required field scopes | > The response body of the API contains the following sensitive fields, and they will be returned only after corresponding scopes are added. If you do not need the fields, it is not recommended that you request the scopes. `contact:user.employee_id:readonly` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `user_id_type` | `string` | No | User ID categories **Example value**: open_id **Optional values are**:  - `open_id`: Identifies a user to an app. The same user has different Open IDs in different apps. How to get Open ID - `union_id`: Identifies a user to a tenant that acts as a developer. A user has the same Union ID in apps developed by the same developer, and has different Union IDs in apps developed by different developers. A developer can use Union ID to link the same user's identities in multiple apps.How to get Union ID - `user_id`: Identifies a user to a tenant. The same user has different User IDs in different tenants. In one single tenant, a user has the same User ID in all apps （including store apps）. User ID is usually used to communicate user data between different apps. How to get User ID  **Default value**: `open_id` **When the value is `user_id`, the following field scopes are required**: `contact:user.employee_id:readonly` | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `user_id` | `string` | Yes | User ID. The ID type needs to be consistent with the value of user_id_type. For information about user IDs, see User-related ID concepts. **Example value**: "ou_XXXXXXXXXX" |
| `timezone` | `string` | Yes | Time zone. **Example value**: "Asia/Shanghai" |
| `start_time` | `string` | Yes | Leave start time. Any of the following formats are supported: - Second-level timestamp: The leave event set through the timestamp is a normal event, that is, leave is requested by the hour. Value example `1609430400` - Date format: The leave event set by date is a full-day event. Value example `2021-01-01` **Note**: The time format selected for start_time and end_time must be consistent, otherwise it will be invalid. **Example value**: "2021-01-01" |
| `end_time` | `string` | Yes | Leave end time. Any of the following formats are supported: - Second-level timestamp: The leave event set through the timestamp is a normal event, that is, leave is requested by the hour. Value example `1609430400` - Date format: The leave event set by date is a full-day event. Value example `2021-01-01` **Note**: The time format selected for start_time and end_time must be consistent, otherwise it will be invalid. **Example value**: "2021-01-01" |
| `title` | `string` | No | Customize the leave event title. **Default**: empty, use default title **Example value**: "1-Day Time Off" |
| `description` | `string` | No | Customize leave event description. **Default**: empty, use default description **Example value**: "If you delete this event, the "leave" label automatically disappears from Lark. However, the leave application in the leave system is not canceled." | ### Request body example

{"user_id":"ou_XXXXXXXXXX",
"timezone":"Asia/Shanghai",
"start_time":"2021-01-01",
"end_time":"2021-01-01",
"title":"1-Day Time Off",
"description":"If you delete this event, the "leave" label automatically disappears from Lark. However, the leave application in the leave system is not canceled."}

## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `timeoff_event` | \- |
|   `timeoff_event_id` | `string` | Leave event ID. This ID can be used to delete the leave event later. |
|   `user_id` | `string` | User ID. For information about user IDs, see User-related ID concepts. |
|   `timezone` | `string` | Time zone. |
|   `start_time` | `string` | Leave start time. Possible time formats returned: - Second level timestamp, for example `1609430400` - Date, for example `2021-01-01` |
|   `end_time` | `string` | Leave end time. Possible time formats returned: - Second level timestamp, for example `1609430400` - Date, for example `2021-01-01` |
|   `title` | `string` | Leave event title. |
|   `description` | `string` | Leave event description. | ### Response body example

{"code":0,
"msg":"success",
"data":{"timeoff_event_id":"timeoff:XXXXXX-XXXX-0917-1623-aa493d591a39-XXXXXX",
"user_id":"ou_XXXXXXXXXX",
"timezone":"Asia/Shanghai",
"start_time":"2021-01-01",
"end_time":"2021-01-01",
"title":"1-Day Time Off",
"description":"If you delete this event, the "leave" label automatically disappears from Lark. However, the leave application in the leave system is not canceled."}}

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
| 404 | 195100 | user is dismiss or not exist in the tenant | The current identity or specified user has resigned or is no longer in the tenant. Please check and change to the correct identity to call the interface. | For more error code information, see General error codes.
