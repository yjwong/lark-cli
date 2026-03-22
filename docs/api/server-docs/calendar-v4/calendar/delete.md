---
title: "Delete shared calendar"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar/delete"
method: "DELETE"
api_path: "https://open.larksuite.com/open-apis/calendar/v4/calendars/:calendar_id"
service: "calendar-v4"
resource: "calendar"
section: "Calendar"
rate_limit: "1000 per minute、50 per second"
scopes:
  - "calendar:calendar"
  - "calendar:calendar:delete"
updated: "1736822105000"
---

# Delete shared calendar

Call this interface to delete a specified shared calendar under the current identity (application or user).

> - The current identity is determined by the Token type of Header Authorization. Tenant_access_token refers to the application identity, user_access_token refers to the user identity.
> - If you call this interface using the application identity, you need to ensure that the application has Bot capabilities.
> - The current identity must have owner permissions for the calendar to be able to delete it. You can call the Query Calendar Information interface, and check the permissions of the current identity to the calendar through the response field role.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/calendar/v4/calendars/:calendar_id |
| HTTP Method | DELETE |
| Rate Limit | 1000 per minute、50 per second |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `calendar:calendar` `calendar:calendar:delete` | ### Request header
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
