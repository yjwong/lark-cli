---
title: "Get reserve admin"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/vc-v1/reserve_config-admin/get"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/vc/v1/reserve_configs/:reserve_config_id/admin"
service: "vc-v1"
resource: "reserve_config-admin"
section: "Video Conferencing"
rate_limit: "100 per minute"
scopes:
  - "vc:room"
  - "vc:room:readonly"
field_scopes:
  - "contact:user.employee_id:readonly"
updated: "1689326170000"
---

# Get reserve admin

Get reserve admin.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/vc/v1/reserve_configs/:reserve_config_id/admin |
| HTTP Method | GET |
| Rate Limit | 100 per minute |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `vc:room` `vc:room:readonly` |
| Required field scopes | > The response body of the API contains the following sensitive fields, and they will be returned only after corresponding scopes are added. If you do not need the fields, it is not recommended that you request the scopes. `contact:user.employee_id:readonly` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `reserve_config_id` | `string` | Room ID or room level ID **Example value**: "omm_3c5dd7e09bac0c1758fcf9511bd1a771" | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `scope_type` | `int` | Yes | 1 means room level, 2 means room **Example value**: 2 **Data validation rules**: - Value range: `1` ～ `2` |
| `user_id_type` | `string` | No | User ID categories **Example value**: open_id **Optional values are**:  - `open_id`: Identifies a user to an app. The same user has different Open IDs in different apps. How to get Open ID - `union_id`: Identifies a user to a tenant that acts as a developer. A user has the same Union ID in apps developed by the same developer, and has different Union IDs in apps developed by different developers. A developer can use Union ID to link the same user's identities in multiple apps.How to get Union ID - `user_id`: Identifies a user to a tenant. The same user has different User IDs in different tenants. In one single tenant, a user has the same User ID in all apps （including store apps）. User ID is usually used to communicate user data between different apps. How to get User ID  **Default value**: `open_id` **When the value is `user_id`, the following field scopes are required**: `contact:user.employee_id:readonly` | ## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `reserve_admin_config` | `reserve_admin_config` | Reserve admin configuration |
|     `depts` | `subscribe_department[]` | Reserve admin departments |
|       `department_id` | `string` | Reserve admin department ID |
|     `users` | `subscribe_user[]` | Reserve admin users |
|       `user_id` | `string` | Reserve admin user ID | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
        "reserve_admin_config": {
            "depts": [
                {
                    "department_id": "od-47d8b570b0a011e9679a755efcc5f61a"
                }
            ],
            "users": [
                {
                    "user_id": "ou_a27b07a9071d90577c0177bcec98f856"
                }
            ]
        }
    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 500 | 121001 | internal error | Internal server error. If retry still doesn't work, contact the administrator. |
| 400 | 121002 | not support | This function is not supported. |
| 400 | 125002 | please contact customer service to upgrade the package to use this function | contact customer service to upgrade the package to use this function |
| 400 | 126001 | invalid room or level id, check itself or whether the scope_type corresponds to it | The room or level ID does not exist, or check if the scope_type corresponds |
| 404 | 121004 | data not exist | The requested data doesn't exist. |
| 403 | 121005 | no permission | No permission. Check the token type, operator identity and resource. |
