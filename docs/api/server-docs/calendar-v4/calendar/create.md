---
title: "Create a shared calendar"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar/create"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/calendar/v4/calendars"
service: "calendar-v4"
resource: "calendar"
section: "Calendar"
rate_limit: "1000 per minute、50 per second"
scopes:
  - "calendar:calendar"
  - "calendar:calendar:create"
updated: "1736822105000"
---

# Create a shared calendar

Call this interface to create a shared calendar for the current identity (app or user).

> - The current identity is determined by the Token type of the Header Authorization. Tenant_access_token refers to app identity, and user_access_token refers to user identity.
> - If this interface is called using the app identity, ensure that the app has activated the bot capabilities.
> - When using this interface to create a shared calendar, the current identity will automatically subscribe to the calendar. The maximum number of calendars that a single identity can subscribe to is 1000.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/calendar/v4/calendars |
| HTTP Method | POST |
| Rate Limit | 1000 per minute、50 per second |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `calendar:calendar` `calendar:calendar:create` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `summary` | `string` | No | Calendar title. **Default value**: Empty **Example value**: "Test calendar" **Data validation rules**: - Maximum length: `255` characters |
| `description` | `string` | No | Calendar description. **Default value**: Empty **Example value**: "Create a calendar by calling an open API" **Data validation rules**: - Maximum length: `255` characters |
| `permissions` | `string` | No | Calendar visibility range. **Default value**: show_only_free_busy **Example value**: "private" **Optional values are**:  - `private`: Private. - `show_only_free_busy`: Only availability information is displayed. - `public`: Public. Event details are visible to other people.  |
| `color` | `int` | No | Calendar color, the value is represented by the int32 of the RGB color value, where, 24 ~ 31 bits are transparency, 16 ~ 23 bits are red, 8 ~ 15 bits are green, 0 ~ 7 bits are blue. For example, -11034625 represents the RGB value (87, 159, 255). **Note**: - The value range is  -2^31 ~ 2^31-1 - The color of the calendar will be mapped to the closest color on the Lark client palette for display. - This color only applies to the current identity. **Default value**: -14513409 **Example value**: -1 |
| `summary_alias` | `string` | No | Calendar alias, setting this field (including subsequent modification of this field) only takes effect for the current identity. **Default value**: Empty **Example value**: "Calendar alias" **Data validation rules**: - Maximum length: `255` characters | ### Request body example

{
    "summary": "Test calendar",
    "description": "Create a calendar by calling an open API",
    "permissions": "private",
    "color": -1,
    "summary_alias": "Calendar alias"
}

## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `calendar` | `calendar` | Newly created calendar entity. |
|     `calendar_id` | `string` | Calendar ID. Subsequent queries, updates, or deletion of calendar information can be completed through this ID. For details, see Calendar-related IDs. |
|     `summary` | `string` | Calendar title. |
|     `description` | `string` | Calendar description. |
|     `permissions` | `string` | Calendar visibility range. **Optional values are**:  - `private`: Private. - `show_only_free_busy`: Only availability information is displayed. - `public`: Public. Event details are visible to other people.  |
|     `color` | `int` | Calendar color, represented by the int32 value of the color RGB. When actually displayed on the client, it will be mapped to the color that is closest on the color palette, and this field only takes effect for the current identity. |
|     `type` | `string` | Calendar type. **Optional values are**:  - `unknown`: Unknown type. - `primary`: Primary calendar of the user or app. - `shared`: Shared calendar created by the user or app. - `google`: Google calendar connected to the user. - `resource`: Room calendar. - `exchange`: Exchange calendar connected to the user.  |
|     `summary_alias` | `string` | Calendar alias, effective only for the current identity. |
|     `is_deleted` | `boolean` | Whether the calendar has been marked as deleted for the current identity. |
|     `is_third_party` | `boolean` | Whether the calendar is owned by a third party. Third-party calendars and events are read-only. |
|     `role` | `string` | Current identity's access to the calendar. **Optional values are**:  - `unknown`: Unknown access. - `free_busy_reader`: Guest. A guest can only view availability information. - `reader`: Subscriber. A subscriber can view all event details. - `writer`: Editor. An editor can create and modify events. - `owner`: Administrator. An administrator can manage calendars and share settings.  | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
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
        }
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
| 404 | 195100 | user is dismiss or not exist in the tenant | The current identity or the specified user has been terminated or is not in the tenant. | For more error code information, see General error codes.
