---
title: "Export meeting details"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/vc-v1/export/meeting_list"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/vc/v1/exports/meeting_list"
service: "vc-v1"
resource: "export"
section: "Video Conferencing"
rate_limit: "100 per minute"
scopes:
  - "vc:export"
field_scopes:
  - "contact:user.employee_id:readonly"
updated: "1689326178000"
---

# Export meeting details

Export meeting details. For specific permission requirements, please refer to "Resource introduction".

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/vc/v1/exports/meeting_list |
| HTTP Method | POST |
| Rate Limit | 100 per minute |
| Supported app types | custom,isv |
| Required scopes | `vc:export` |
| Required field scopes | > The response body of the API contains the following sensitive fields, and they will be returned only after corresponding scopes are added. If you do not need the fields, it is not recommended that you request the scopes. `contact:user.employee_id:readonly` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `user_id_type` | `string` | No | User ID categories **Example value**: open_id **Optional values are**:  - `open_id`: Identifies a user to an app. The same user has different Open IDs in different apps. How to get Open ID - `union_id`: Identifies a user to a tenant that acts as a developer. A user has the same Union ID in apps developed by the same developer, and has different Union IDs in apps developed by different developers. A developer can use Union ID to link the same user's identities in multiple apps.How to get Union ID - `user_id`: Identifies a user to a tenant. The same user has different User IDs in different tenants. In one single tenant, a user has the same User ID in all apps （including store apps）. User ID is usually used to communicate user data between different apps. How to get User ID  **Default value**: `open_id` **When the value is `user_id`, the following field scopes are required**: `contact:user.employee_id:readonly` | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `start_time` | `string` | Yes | Query start time (unix time, unit sec) **Example value**: "1655276858" |
| `end_time` | `string` | Yes | Query end time (unix time, unit sec) **Example value**: "1655276858" |
| `meeting_no` | `string` | No | Filter by 9-digit meeting number (up to one filter) **Example value**: "123456789" |
| `user_id` | `string` | No | Filter by participating Lark users (up to one filter) **Example value**: "ou_3ec3f6a28a0d08c45d895276e8e5e19b" |
| `room_id` | `string` | No | Filter by Rooms (maximum one filter) **Example value**: "omm_eada1d61a550955240c28757e7dec3af" | ### Request body example

{
    "start_time": "1655276858",
    "end_time": "1655276858",
    "meeting_no": "123456789",
    "user_id": "ou_3ec3f6a28a0d08c45d895276e8e5e19b",
    "room_id": "omm_eada1d61a550955240c28757e7dec3af"
}

## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `task_id` | `string` | Task id | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
        "task_id": "7111325589855797267"
    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 500 | 121001 | internal error | Internal server error. If retry still doesn't work, contact the administrator. |
| 400 | 121002 | not support | This function is not supported. |
| 400 | 121003 | param error | Parameter error. Check the value range of parameters (please check yourself according to the field description above). |
| 404 | 121004 | data not exist | The requested data doesn't exist. |
| 403 | 121005 | no permission | No permission. Check the token type, operator identity and resource. |
