---
title: "Query primary calendar information"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar/primary"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/calendar/v4/calendars/primary"
service: "calendar-v4"
resource: "calendar"
section: "Calendar"
rate_limit: "1000 per minute、50 per second"
scopes:
  - "calendar:calendar:read"
  - "calendar:calendar:readonly"
field_scopes:
  - "contact:user.employee_id:readonly"
updated: "1736822105000"
---

# Query primary calendar information

Call this interface to get the main calendar information of the current identity (application or user).

> - The current identity is determined by the Token type of Header Authorization. Tenant_access_token refers to the application identity, user_access_token refers to the user identity.
> - Before calling this interface with the application identity, it's necessary to ensure that this application has enabled bot capabilities.
> - When calling this interface using the application identity, the query parameter user_id_type cannot be set to user_id. You can choose open_id or union_id. In the returned result, the user_id parameter value will contain the open_id or union_id corresponding to the application robot.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/calendar/v4/calendars/primary |
| HTTP Method | POST |
| Rate Limit | 1000 per minute、50 per second |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `calendar:calendar:read` `calendar:calendar:readonly` |
| Required field scopes | > The response body of the API contains the following sensitive fields, and they will be returned only after corresponding scopes are added. If you do not need the fields, it is not recommended that you request the scopes. `contact:user.employee_id:readonly` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `user_id_type` | `string` | No | User ID categories **Example value**: open_id **Optional values are**:  - `open_id`: Identifies a user to an app. The same user has different Open IDs in different apps. How to get Open ID - `union_id`: Identifies a user to a tenant that acts as a developer. A user has the same Union ID in apps developed by the same developer, and has different Union IDs in apps developed by different developers. A developer can use Union ID to link the same user's identities in multiple apps.How to get Union ID - `user_id`: Identifies a user to a tenant. The same user has different User IDs in different tenants. In one single tenant, a user has the same User ID in all apps （including store apps）. User ID is usually used to communicate user data between different apps. How to get User ID  **Default value**: `open_id` **When the value is `user_id`, the following field scopes are required**: `contact:user.employee_id:readonly` | ## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `calendars` | `user_calendar[]` | Primary calendar list. |
|     `calendar` | `calendar` | Calendar entity information. |
|       `calendar_id` | `string` | Calendar ID. You can use this ID to query and update the primary calendar information. For details, see Calendar-related IDs. |
|       `summary` | `string` | Calendar title. |
|       `description` | `string` | Calendar description. |
|       `permissions` | `string` | Calendar visibility range. **Optional values are**:  - `private`: Private - `show_only_free_busy`: Only availability information is displayed. - `public`: Public. Event details are visible to other people  |
|       `color` | `int` | Calendar color, represented by an int32 of the color RGB value. When actually displayed on the client, it will map to the closest color on the color palette, and this field only takes effect for the current identity. |
|       `type` | `string` | Calendar type. **Optional values are**:  - `unknown`: Unknown type - `primary`: Primary calendar of the user or app - `shared`: Shared calendar created by the user or app - `google`: Google calendar connected to the user - `resource`: Room calendar - `exchange`: Exchange calendar connected to the user  |
|       `summary_alias` | `string` | Calendar alias, which only takes effect for the current identity. |
|       `is_deleted` | `boolean` | Whether the calendar has been marked as deleted for the current identity |
|       `is_third_party` | `boolean` | Whether the calendar is owned by a third party. Third-party calendars and events are read-only. |
|       `role` | `string` | Current identity's access to the calendar. **Optional values are**:  - `unknown`: Unknown access. - `free_busy_reader`: Guest. A guest can only view availability information. - `reader`: Subscriber. A subscriber can view all event details. - `writer`: Editor. An editor can create and modify events. - `owner`: Administrator. An administrator can manage calendars and share settings.  |
|     `user_id` | `string` | The creator's User ID of the calendar. Understand different types of user IDs. See User-related ID concept. | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
        "calendars": [
            {
                "calendar": {
                    "calendar_id": "larksuite.com_xxxxxxxxxx@group.calendar.larksuite.com",
                    "summary": "Test calendar",
                    "description": "Create a calendar by calling an open API",
                    "permissions": "private",
                    "color": -1,
                    "type": "shared",
                    "summary_alias": "Calendar alias",
                    "is_deleted": false,
                    "is_third_party": false,
                    "role": "owner"
                },
                "user_id": "ou_xxxxxx"
            }
        ]
    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 190002 | invalid parameters in request | Invalid request parameters. Troubleshooting suggestions are as follows: - Confirm that the field name and parameter type of the request parameter are correct. - Confirm that the permissions for the corresponding resource have been applied for. - Confirm that the corresponding resource has not been deleted. |
| 500 | 190003 | internal service error | Internal service error. Please consult [Technical Support](https://applink.larksuite.com/TLJpeNdW). |
| 429 | 190004 | method rate limited | Reduce the QPS and try again later. |
| 429 | 190005 | app rate limited | Reduce the QPS and try again later. |
| 403 | 190006 | wrong unit for app tenant | Request error, check whether the App ID and App Secret are correct. If the problem still cannot be solved, please consult [Technical Support](https://applink.larksuite.com/TLJpeNdW). |
| 404 | 190007 | app bot_id not found | The bot_id of the application is not found. Please check whether the robot capability is enabled in the application. If the problem is still not solved, please contact [Technical Support](https://applink.larksuite.com/TLJpeNdW). |
| 429 | 190010 | current operation rate limited | The action is throttled. This is usually caused by the failure to preempt public resources. Reduce the QPS and try again. |
| 404 | 191000 | calendar not found | No calendar found. Check whether the calendar ID is correct. |
| 400 | 191001 | invalid calendar_id | Invalid calendar ID. Check whether the calendar ID is correct. |
| 403 | 191002 | no calendar access_role | No access to the calendar. Check whether the calendar ID is correct. |
| 403 | 191003 | calendar is deleted | The calendar has been deleted. Check whether the calendar ID is correct. |
| 403 | 191004 | invalid calendar type | Incorrect calendar type. Check the calendar type. |
| 404 | 195100 | user is dismiss or not exist in the tenant | The current identity or the specified user has been terminated or is not in the tenant. | For more error code information, see General error codes.
