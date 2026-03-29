---
title: "Create ACL"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar-acl/create"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/calendar/v4/calendars/:calendar_id/acls"
service: "calendar-v4"
resource: "calendar-acl"
section: "Calendar"
rate_limit: "1000 per minute、50 per second"
scopes:
  - "calendar:calendar"
  - "calendar:calendar.acl:create"
field_scopes:
  - "contact:user.employee_id:readonly"
updated: "1736822128000"
---

# Create access control

Call this interface to add access control to the specified calendar in your current identity (application or user), that is, calendar member permissions.

> - The current identity is determined by the Token type of Header Authorization. 'tenant_access_token' refers to the application identity, and 'user_access_token' refers to the user identity.
> - If you use the application identity to call this interface, you need to make sure that the application has turned on the bot capability.
> - The current identity needs to have the owner permission of the calendar, and the type of the calendar can only be primary or shared. You can call the query calendar information interface to get the calendar type and the current identity's access rights to the calendar.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/calendar/v4/calendars/:calendar_id/acls |
| HTTP Method | POST |
| Rate Limit | 1000 per minute、50 per second |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `calendar:calendar` `calendar:calendar.acl:create` |
| Required field scopes | > The response body of the API contains the following sensitive fields, and they will be returned only after corresponding scopes are added. If you do not need the fields, it is not recommended that you request the scopes. `contact:user.employee_id:readonly` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `calendar_id` | `string` | The calendar ID for which access control needs to be added. The calendar ID is returned when a shared calendar is created. You can also call the following interfaces to obtain the ID of a calendar. - Query primary calendar information - Query calendar list - Search calendars **Example value**: "larksuite.com_xxxxxxxxxx@group.calendar.larksuite.com" | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `user_id_type` | `string` | No | User ID categories **Example value**: open_id **Optional values are**:  - `open_id`: Identifies a user to an app. The same user has different Open IDs in different apps. How to get Open ID - `union_id`: Identifies a user to a tenant that acts as a developer. A user has the same Union ID in apps developed by the same developer, and has different Union IDs in apps developed by different developers. A developer can use Union ID to link the same user's identities in multiple apps.How to get Union ID - `user_id`: Identifies a user to a tenant. The same user has different User IDs in different tenants. In one single tenant, a user has the same User ID in all apps （including store apps）. User ID is usually used to communicate user data between different apps. How to get User ID  **Default value**: `open_id` **When the value is `user_id`, the following field scopes are required**: `contact:user.employee_id:readonly` | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `role` | `string` | Yes | Access to the calendar. **Example value**: "writer" **Optional values are**:  - `unknown`: Unknown permission. unknown is one of the enumerated values of the role parameter, but when role is used as a request parameter, it does not support the input of unknown. - `free_busy_reader`: Guest. A guest can only view availability information. - `reader`: Subscriber. A subscriber can view all event details. - `writer`: Editor. An editor can create and modify events. - `owner`: Administrator. An administrator can manage calendars and share settings.  |
| `scope` | `acl_scope` | Yes | The scope of the permission's effectiveness. |
|   `type` | `string` | Yes | The type of the scope where the permission takes effect. **Note**: Currently, only user is supported. When `type=user`, the user_id needs to be consistent with the type of user_id_type. For example, when `user_id_type=open_id`, the user_id needs to be the user's open_id. **Example value**: "user" **Optional values are**:  - `user`: User  |
|   `user_id` | `string` | No | User ID, for more information, you can refer to the Concept of User-Related ID. **Example value**: "ou_xxxxxx" | ### Request body example

{
    "role": "writer",
    "scope": {
        "type": "user",
        "user_id": "ou_xxxxxx"
    }
}

## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `calendar.acl` | \- |
|   `acl_id` | `string` | ACL resource ID. For details, see Calendar-related IDs. |
|   `role` | `string` | Access to the calendar **Optional values are**:  - `unknown`: Unknown access - `free_busy_reader`: Guest. A guest can only view availability information. - `reader`: Subscriber. A subscriber can view all event details. - `writer`: Editor. An editor can create and modify events. - `owner`: Administrator. An administrator can manage calendars and share settings.  |
|   `scope` | `acl_scope` | Scope |
|     `type` | `string` | Scope type. When type is User, the value can be open_id, user_id, or union_id. **Optional values are**:  - `user`: User  |
|     `user_id` | `string` | User ID. For details, see User-related IDs. | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
        "acl_id": "user_xxxxxx",
        "role": "writer",
        "scope": {
            "type": "user",
            "user_id": "ou_xxxxxx"
        }
    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 190002 | invalid parameters in request | Invalid request parameters. Troubleshooting suggestions are as follows: - Confirm that the field name and parameter type of the request parameter are correct. - Confirm that the permissions for the corresponding resource have been applied for. - Confirm that the corresponding resource has not been deleted. |
| 500 | 190003 | internal service error | Internal service error, please contact [Technical Support](https://applink.larksuite.com/TLJpeNdW). |
| 429 | 190004 | method rate limited | Method frequency limit. It is recommended to try again later and reduce the request QPS appropriately. |
| 429 | 190005 | app rate limited | Frequency limiting is applied. We recommend that you try again later and reduce the request QPS appropriately. |
| 403 | 190006 | wrong unit for app tenant | Request error. You need to check whether the App ID and App Secret are correct. If the problem is still not solved, please consult [Technical Support](https://applink.larksuite.com/TLJpeNdW). |
| 404 | 190007 | app bot_id not found | The bot_id of the application is not found. You need to make sure that the application has enabled bot capability. If the problem is still not solved, please contact [technical support](https://applink.larksuite.com/TLJpeNdW). |
| 429 | 190010 | current operation rate limited | The current operation is limited, usually because concurrent preemption of public resources fails. You can appropriately reduce the frequency of the current operation and try again. |
| 404 | 191000 | calendar not found | Calendar not found. You need to check and change to the correct calendar ID. |
| 400 | 191001 | invalid calendar_id | The calendar_id is invalid. You need to check and change it to the correct calendar ID. |
| 403 | 191002 | no calendar access_role | The current identity does not have permission to access the calendar. If you want to query a calendar, you need to ensure that the current identity has permission to access the calendar. |
| 403 | 191003 | calendar is deleted | The calendar has been deleted. You need to check and change to the correct calendar ID. |
| 403 | 191004 | invalid calendar type | The calendar type is incorrect. You can call the query calendar information interface to obtain the calendar type information, and then ensure that the calendar type is applicable to the current interface. |
| 404 | 192000 | acl not found | Access control (ACL) not found. You need to make sure you passed in the correct ACL ID. |
| 404 | 195100 | user is dismiss or not exist in the tenant | The current identity or specified user has resigned or is no longer in the tenant. Please check and change to the correct identity to call the interface. | For more error code information, see General error codes.
