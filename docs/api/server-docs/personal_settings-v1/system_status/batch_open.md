---
title: "Batch open system status"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/personal_settings-v1/system_status/batch_open"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/personal_settings/v1/system_statuses/:system_status_id/batch_open"
service: "personal_settings-v1"
resource: "system_status"
section: "Personal Settings"
scopes:
  - "personal_settings:status:system_status_operate"
field_scopes:
  - "contact:user.employee_id:readonly"
updated: "1672726533000"
---

# Batch open system status

Batch open user system status is available.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/personal_settings/v1/system_statuses/:system_status_id/batch_open |
| HTTP Method | POST |
| Supported app types | custom |
| Required scopes | `personal_settings:status:system_status_operate` |
| Required field scopes | > The response body of the API contains the following sensitive fields, and they will be returned only after corresponding scopes are added. If you do not need the fields, it is not recommended that you request the scopes. `contact:user.employee_id:readonly` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `system_status_id` | `string` | System stauts ID Get system status ID **Example value**: "7101214603622940672" | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `user_id_type` | `string` | No | User ID categories **Example value**: "open_id" **Optional values are**:  - `open_id`: Identifies a user to an app. The same user has different Open IDs in different apps. How to get Open ID - `union_id`: Identifies a user to a tenant that acts as a developer. A user has the same Union ID in apps developed by the same developer, and has different Union IDs in apps developed by different developers. A developer can use Union ID to link the same user's identities in multiple apps.How to get Union ID - `user_id`: Identifies a user to a tenant. The same user has different User IDs in different tenants. In one single tenant, a user has the same User ID in all apps （including store apps）. User ID is usually used to communicate user data between different apps. How to get User ID  **Default value**: `open_id` **When the value is `user_id`, the following field scopes are required**: `contact:user.employee_id:readonly` | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `user_list` | `system_status_user_open_param[]` | Yes | User list **Data validation rules**: - Length range: `1` ～ `50` |
|   `user_id` | `string` | Yes | UserID, the ID type is determined by user_id_type. Open ID is recommended here. For details, refer to How to get Open ID. **Example value**: "ou_53edd3282dbc2fdbe5c593cfa5ce82ab" |
|   `end_time` | `string` | Yes | End time **Example value**: "1665990378" | ### Request body example

{
    "user_list": [
        {
            "user_id": "ou_53edd3282dbc2fdbe5c593cfa5ce82ab",
            "end_time": "1665990378"
        }
    ]
}

## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `result_list` | `system_status_user_open_result_entity[]` | Open result |
|     `user_id` | `string` | User id |
|     `end_time` | `string` | End time, the incoming timestamp should be in seconds, and the current time span cannot exceed 365 days. |
|     `result` | `string` | Open result **Optional values are**:  - `success_show`: Successful and displayed on the client - `success_user_close_syn`: Succeeded but the user has set it to be disabled by default - `success_user_in_higher_priority_system_status`: Success but the user has a higher priority system state - `fail`: Failed - `invisible_user_id`: User is not visible - `invalid_user_id`: User id is not valid - `resign_user_id`: User leaves  | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
        "result_list": [
            {
                "user_id": "ou_53edd3282dbc2fdbe5c593cfa5ce82ab",
                "end_time": "1665990378",
                "result": "success_show"
            }
        ]
    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 2005001 | Your request contains an invalid request parameter. | The parameter is incorrect. Please check the input parameter according to the error message returned by the interface and refer to the documentation. |
| 400 | 2005006 | Request contains other tenant user. | Please check if the user ID is in your tenant. |
| 400 | 2005007 | Tenant does not have permission to api. | Tenant does not have permission to api, Please apply to use. |
