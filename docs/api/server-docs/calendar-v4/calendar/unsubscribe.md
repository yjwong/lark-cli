---
title: "Unsubscribe Calendar"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar/unsubscribe"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/calendar/v4/calendars/:calendar_id/unsubscribe"
service: "calendar-v4"
resource: "calendar"
section: "Calendar"
rate_limit: "1000 per minute、50 per second"
scopes:
  - "calendar:calendar"
  - "calendar:calendar:subscribe"
updated: "1736822106000"
---

# Unsubscribe from a calendar

Call this interface with the current identity (application or user) to unsubscribe from the specified calendar.

> - The current identity is determined by the Token type of Header Authorization. The tenant_access_token refers to the application identity, and the user_access_token refers to the user identity.
> - If this interface is called with the application identity, make sure that the application has enabled the Bot ability.
> - Only calendars that have been subscribed to by the current identity can be operated.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/calendar/v4/calendars/:calendar_id/unsubscribe |
| HTTP Method | POST |
| Rate Limit | 1000 per minute、50 per second |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `calendar:calendar` `calendar:calendar:subscribe` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `calendar_id` | `string` | Calendar ID. You can call the following interfaces to get the ID of a calendar. - Query primary calendar information - Query calendar list - Search calendar **Example value**: "larksuite.com_xxxxxxxxxx@group.calendar.larksuite.com" | ## Response

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
| 400 | 190014 | invalid request parameters. check details for more info. | Invalid request parameters. Please check the details for more information. |
| 404 | 191000 | calendar not found | Calendar not found. You need to check and change to the correct calendar ID. |
| 400 | 191001 | invalid calendar_id | The calendar_id is invalid. You need to check and change it to the correct calendar ID. |
| 403 | 191002 | no calendar access_role | The current identity does not have permission to access the calendar. If you want to query a calendar, you need to ensure that the current identity has permission to access the calendar. |
| 403 | 191003 | calendar is deleted | The calendar has been deleted. You need to check and change to the correct calendar ID. |
| 403 | 191004 | invalid calendar type | The calendar type is incorrect. You can call the query calendar information interface to obtain the calendar type information, and then ensure that the calendar type is applicable to the current interface. |
| 400 | 191007 | target calendar is already unsubscribed | The target calendar has been unsubscribed. You need to check and change to the correct calendar ID. |
| 400 | 191008 | primary calendar can not be unsubscribed | The primary calendar of the current identity cannot be unsubscribed. |
| 404 | 195100 | user is dismiss or not exist in the tenant | The current identity or specified user has resigned or is no longer in the tenant. Please check and change to the correct identity to call the interface. | For more error code information, see General error codes.
