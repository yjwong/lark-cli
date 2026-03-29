---
title: "Query calendar information"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar/get"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/calendar/v4/calendars/:calendar_id"
service: "calendar-v4"
resource: "calendar"
section: "Calendar"
rate_limit: "1000 per minute、50 per second"
scopes:
  - "calendar:calendar"
  - "calendar:calendar:read"
  - "calendar:calendar:readonly"
updated: "1736822105000"
---

# Query calendar information

Call this interface to query the information of a specified calendar with the current identity (application or user).

> - The current identity is determined by the Token type of Header Authorization. Tenant_access_token refers to the application identity, user_access_token refers to the user identity.
> - If you call this interface using the application identity, you need to ensure that the application has Bot capabilities.
> - The current identity must have access to the calendar.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/calendar/v4/calendars/:calendar_id |
| HTTP Method | GET |
| Rate Limit | 1000 per minute、50 per second |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `calendar:calendar` `calendar:calendar:read` `calendar:calendar:readonly` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `calendar_id` | `string` | Calendar ID. When creating a shared calendar, the calendar ID will be returned. You can also call the following interfaces to get the ID of a calendar. - Query Primary Calendar Information - Query Calendar List - Search Calendars **Example value**: "larksuite.com_xxxxxxxxxx@group.calendar.larksuite.com" | ## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `calendar` | \- |
|   `calendar_id` | `string` | Calendar ID. Subsequently, you can query, update, or delete calendar information through this ID. For details, see Calendar-related IDs. |
|   `summary` | `string` | Calendar title. |
|   `description` | `string` | Calendar description. |
|   `permissions` | `string` | Calendar visibility range. **Optional values are**:  - `private`: Private. - `show_only_free_busy`: Only availability information is displayed. - `public`: Public. Event details are visible to other people.  |
|   `color` | `int` | Calendar color, represented by an int32 of the color RGB value. When actually displayed on the client, it will map to the closest color on the color palette, and this field only takes effect for the current identity. |
|   `type` | `string` | Calendar type. **Optional values are**:  - `unknown`: Unknown type - `primary`: Primary calendar of the user or app - `shared`: Shared calendar created by the user or app - `google`: Google calendar connected to the user - `resource`: Room calendar - `exchange`: Exchange calendar connected to the user  |
|   `summary_alias` | `string` | Calendar alias, which only takes effect for the current identity. |
|   `is_deleted` | `boolean` | Whether the calendar has been marked as deleted for the current identity. |
|   `is_third_party` | `boolean` | Whether the calendar is owned by a third party. Third-party calendars and events are read-only. |
|   `role` | `string` | Current identity's access to the calendar. **Optional values are**:  - `unknown`: Unknown access. - `free_busy_reader`: Guest. A guest can only view availability information. - `reader`: Subscriber. A subscriber can view all event details. - `writer`: Editor. An editor can create and modify events. - `owner`: Administrator. An administrator can manage calendars and share settings.  | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
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
| 404 | 191000 | calendar not found | Calendar not found. You need to check and change to the correct calendar ID. |
| 400 | 191001 | invalid calendar_id | The calendar_id is invalid. You need to check and change it to the correct calendar ID. |
| 403 | 191002 | no calendar access_role | The current identity does not have permission to access the calendar. If you want to query calendar information, you need to ensure that the current identity has permission to access the calendar. |
| 403 | 191003 | calendar is deleted | The calendar has been deleted. You need to check and change to the correct calendar ID. |
| 403 | 191004 | invalid calendar type | The calendar type is incorrect. You can call the query calendar information interface to obtain the calendar type information, and then ensure that the calendar type is applicable to the current interface. |
| 404 | 195100 | user is dismiss or not exist in the tenant | The current identity or specified user has resigned or is no longer in the tenant. Please check and change to the correct identity to call the interface. | For more error code information, see General error codes.
