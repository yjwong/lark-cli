---
title: "Set meeting host"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/vc-v1/meeting/set_host"
method: "PATCH"
api_path: "https://open.larksuite.com/open-apis/vc/v1/meetings/:meeting_id/set_host"
service: "vc-v1"
resource: "meeting"
section: "Video Conferencing"
rate_limit: "1000 per minute、50 per second"
scopes:
  - "vc:meeting"
field_scopes:
  - "contact:user.employee_id:readonly"
updated: "1692787315000"
---

# Set meeting host

Sets a host for a meeting.

> Operators must have relevant permission scope to set hosts. If the operator is a user, they must be the current host of the meeting. This operation uses CAS concurrency security mechanism and the current host in the meeting is required. If the operation failed, try again by using the latest returned data.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/vc/v1/meetings/:meeting_id/set_host |
| HTTP Method | PATCH |
| Rate Limit | 1000 per minute、50 per second |
| Supported app types | custom,isv |
| Required scopes | `vc:meeting` |
| Required field scopes | > The response body of the API contains the following sensitive fields, and they will be returned only after corresponding scopes are added. If you do not need the fields, it is not recommended that you request the scopes. `contact:user.employee_id:readonly` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `meeting_id` | `string` | Meeting ID (Unique identifier of a video conference, generated after the conference is started) **Example value**: "6911188411932033028" | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `user_id_type` | `string` | No | User ID categories **Example value**: open_id **Optional values are**:  - `open_id`: Identifies a user to an app. The same user has different Open IDs in different apps. How to get Open ID - `union_id`: Identifies a user to a tenant that acts as a developer. A user has the same Union ID in apps developed by the same developer, and has different Union IDs in apps developed by different developers. A developer can use Union ID to link the same user's identities in multiple apps.How to get Union ID - `user_id`: Identifies a user to a tenant. The same user has different User IDs in different tenants. In one single tenant, a user has the same User ID in all apps （including store apps）. User ID is usually used to communicate user data between different apps. How to get User ID  **Default value**: `open_id` **When the value is `user_id`, the following field scopes are required**: `contact:user.employee_id:readonly` | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `host_user` | `meeting_user` | Yes | The user to be set as the host |
|   `id` | `string` | No | User ID **Example value**: "ou_3ec3f6a28a0d08c45d895276e8e5e19b" |
|   `user_type` | `int` | No | User type **Example value**: 1 **Optional values are**:  - `1`: Lark user - `2`: Rooms user - `3`: Docs user - `4`: neo Lark Meetings user - `5`: neo Lark Meetings guest - `6`: PSTN user - `7`: SIP user  |
| `old_host_user` | `meeting_user` | No | Current host (CAS concurrency security: If this field is inconsistent with the current host in the meeting, the setting will fail. Try again by using the latest returned data.) |
|   `id` | `string` | No | User ID **Example value**: "ou_3ec3f6a28a0d08c45d895276e8e5e19b" |
|   `user_type` | `int` | No | User type **Example value**: 1 **Optional values are**:  - `1`: Lark user - `2`: Rooms user - `3`: Docs user - `4`: neo Lark Meetings user - `5`: neo Lark Meetings guest - `6`: PSTN user - `7`: SIP user  | ### Request body example

{
    "host_user": {
        "id": "ou_3ec3f6a28a0d08c45d895276e8e5e19b",
        "user_type": 1
    },
    "old_host_user": {
        "id": "ou_3ec3f6a28a0d08c45d895276e8e5e19b",
        "user_type": 1
    }
}

## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `host_user` | `meeting_user` | The current host in the meeting |
|     `id` | `string` | User ID |
|     `user_type` | `int` | User type **Optional values are**:  - `1`: Lark user - `2`: Rooms user - `3`: Docs user - `4`: neo Lark Meetings user - `5`: neo Lark Meetings guest - `6`: PSTN user - `7`: SIP user  | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
        "host_user": {
            "id": "ou_3ec3f6a28a0d08c45d895276e8e5e19b",
            "user_type": 1
        }
    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 500 | 121001 | internal error | Internal server error. If retry still doesn't work, contact the administrator. |
| 400 | 121002 | not support | This function is not supported. |
| 400 | 121003 | param error | Parameter error. Check the value range of parameters (please check yourself according to the field description above). |
| 403 | 121005 | no permission | No permission. Check the token type, operator identity and resource. |
| 409 | 121006 | data conflict | This is often caused by CAS. Update the data and try again. |
| 400 | 122001 | meeting status unexpected | Meeting status is incorrect. For example, the meeting is not in progress. |
| 404 | 122002 | meeting not exist | The meeting does not exist. |
| 400 | 122003 | operator user not in meeting | The operator must be in the meeting. |
| 400 | 122004 | target user not in meeting | The target user must be in the meeting. |
| 400 | 122005 | cannot operate self | Operators cannot operate on themselves. |
| 400 | 122009 | participant status not expected | The target user status is incorrect. For example, the user is not in the meeting. |
