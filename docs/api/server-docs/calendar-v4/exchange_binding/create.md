---
title: "Bind Exchange account to Lark account"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/exchange_binding/create"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/calendar/v4/exchange_bindings"
service: "calendar-v4"
resource: "exchange_binding"
section: "Calendar"
rate_limit: "1000 per minute、50 per second"
scopes:
  - "calendar:calendar"
  - "calendar:calendar:readonly"
  - "calendar:exchange.bindings:create"
field_scopes:
  - "contact:user.email:readonly"
  - "contact:user.employee_id:readonly"
updated: "1741687012000"
---

# Bind Exchange account to Lark account

Call this interface to bind the Exchange account to the Lark account, thereby supporting the import of the Exchange calendar.

> Only a super administrator of the company can perform this action.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/calendar/v4/exchange_bindings |
| HTTP Method | POST |
| Rate Limit | 1000 per minute、50 per second |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `calendar:calendar` `calendar:calendar:readonly` `calendar:exchange.bindings:create` |
| Required field scopes | > The response body of the API contains the following sensitive fields, and they will be returned only after corresponding scopes are added. If you do not need the fields, it is not recommended that you request the scopes. `contact:user.email:readonly` `contact:user.employee_id:readonly` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `user_id_type` | `string` | No | User ID categories **Example value**: open_id **Optional values are**:  - `open_id`: Identifies a user to an app. The same user has different Open IDs in different apps. How to get Open ID - `union_id`: Identifies a user to a tenant that acts as a developer. A user has the same Union ID in apps developed by the same developer, and has different Union IDs in apps developed by different developers. A developer can use Union ID to link the same user's identities in multiple apps.How to get Union ID - `user_id`: Identifies a user to a tenant. The same user has different User IDs in different tenants. In one single tenant, a user has the same User ID in all apps （including store apps）. User ID is usually used to communicate user data between different apps. How to get User ID  **Default value**: `open_id` **When the value is `user_id`, the following field scopes are required**: `contact:user.employee_id:readonly` | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `admin_account` | `string` | No | The admin account of Exchange. **Example value**: "email_admin_example@outlook.com" **Data validation rules**: - Length range: `1` ～ `500` characters |
| `exchange_account` | `string` | No | The Exchange account that needs to be bound. **Example value**: "email_account_example@outlook.com" **Data validation rules**: - Length range: `1` ～ `500` characters |
| `user_id` | `string` | No | User ID, that is, the Lark account ID bound to the Exchange account. For more information, see User-related IDs. **Example value**: "ou_xxxxxxxxxxxxxxxxxx" | ### Request body example

{
    "admin_account": "email_admin_example@outlook.com",
    "exchange_account": "email_account_example@outlook.com",
    "user_id": "ou_xxxxxxxxxxxxxxxxxx"
}

## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `exchange_binding` | \- |
|   `admin_account` | `string` | The admin account of Exchange. **Required field scopes**: `contact:user.email:readonly` |
|   `exchange_account` | `string` | The Exchange account that needs to be bound. **Required field scopes**: `contact:user.email:readonly` |
|   `user_id` | `string` | User ID, that is, the Lark account ID bound to the Exchange account. |
|   `status` | `string` | Sync status of Exchange account. **Optional values are**:  - `doing`: The calendar is being synced. - `cal_done`: Calendar sync is complete. - `timespan_done`: The data for the recent period of time has been synced. - `done`: Event sync is complete. - `err`: Synchronization error.  |
|   `exchange_binding_id` | `string` | The unique identification ID bound to Exchange is the unique identification ID of the admin account, Exchange account, and user triplet. You can use this ID to query the binding relationship, calendar synchronization status, or unbind the relationship. | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
        "admin_account": "email_admin_example@outlook.com",
        "exchange_account": "email_account_example@outlook.com",
        "user_id": "ou_xxxxxxxxxxxxxxxxxx",
        "status": "doing",
        "exchange_binding_id": "ZW1haWxfYWRtaW5fZXhhbXBsZUBvdXRsb29rLmNvbSBlbWFpbF9hY2NvdW50X2V4YW1wbGVAb3V0bG9vay5jb20="
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
| 404 | 195100 | user is dismiss or not exist in the tenant | The current identity or specified user has resigned or is no longer in the tenant. Please check and change to the correct identity to call the interface. |
| 403 | 195101 | user is not supper administrator | The current identity is not the super administrator of this tenant. Please check and modify your identity information. |
| 400 | 195102 | exchange_binding_id invalid | The exchange_binding_id is invalid. You need to check and modify it to the correct ID. |
| 404 | 195103 | exchange account binding is not found | The binding relationship of the exchange account was not found. You need to check whether the input parameters are filled in correctly. |
| 403 | 195104 | current tenant not match | The current tenant does not match the tenant bound to the admin account. You need to check and modify the parameter value to the correct one. |
| 403 | 195105 | admin account binding failed | Failed to bind the admin account in the management backend. Please rebind the admin account in the management backend and try again. |
| 404 | 195106 | admin account is not found | The admin account does not exist. You need to check and modify the parameter values ​​to the correct ones. | For more error code information, see General error codes.
