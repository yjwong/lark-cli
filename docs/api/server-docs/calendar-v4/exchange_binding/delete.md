---
title: "Unbind the Exchange account"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/exchange_binding/delete"
method: "DELETE"
api_path: "https://open.larksuite.com/open-apis/calendar/v4/exchange_bindings/:exchange_binding_id"
service: "calendar-v4"
resource: "exchange_binding"
section: "Calendar"
rate_limit: "1000 per minute、50 per second"
scopes:
  - "calendar:calendar"
  - "calendar:exchange.bindings:delete"
updated: "1736837327000"
---

# Unbind the Exchange account

Call this interface to unbind the relationship between the Exchange account and the Lark account. After the Exchange account is unbound, it can continue to bind with other Lark accounts.

> Only a super administrator of the company can perform this operation.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/calendar/v4/exchange_bindings/:exchange_binding_id |
| HTTP Method | DELETE |
| Rate Limit | 1000 per minute、50 per second |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `calendar:calendar` `calendar:exchange.bindings:delete` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `exchange_binding_id` | `string` | The unique identification ID bound to Exchange. When calling Bind Exchange Account to Lark Account, the exchange_binding_id can be obtained from the returned result. **Example value**: "ZW1haWxfYWRtaW5fZXhhbXBsZUBvdXRsb29rLmNvbSBlbWFpbF9hY2NvdW50X2V4YW1wbGVAb3V0bG9vay5jb20=" | ## Response

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
| 404 | 195100 | user is dismiss or not exist in the tenant | The current identity or specified user has resigned or is no longer in the tenant. Please check and change to the correct identity to call the interface. |
| 403 | 195101 | user is not supper administrator | The current identity is not the super administrator of this tenant. Please check and modify your identity information. |
| 400 | 195102 | exchange_binding_id invalid | The exchange_binding_id is invalid. You need to check and modify it to the correct ID. |
| 404 | 195103 | exchange account binding is not found | The binding relationship of the exchange account was not found. You need to check whether the input parameters are filled in correctly. |
| 403 | 195104 | current tenant not match | The current tenant does not match the tenant bound to the admin account. You need to check and modify the parameter value to the correct one. |
| 403 | 195105 | admin account binding failed | Failed to bind the admin account in the management backend. Please rebind the admin account in the management backend and try again. |
| 404 | 195106 | admin account is not found | The admin account does not exist. You need to check and modify the parameter values ​​to the correct ones. | For more error code information, see General error codes.
