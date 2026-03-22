---
title: "Authorize recording files"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/vc-v1/meeting-recording/set_permission"
method: "PATCH"
api_path: "https://open.larksuite.com/open-apis/vc/v1/meetings/:meeting_id/recording/set_permission"
service: "vc-v1"
resource: "meeting-recording"
section: "Video Conferencing"
rate_limit: "1000 per minute、50 per second"
scopes:
  - "vc:record"
field_scopes:
  - "contact:user.employee_id:readonly"
updated: "1680514408000"
---

# Authorize recording files

Authorizes a meeting's recording files to an organization/user or makes them public to the public network.

> Authorization can be granted only after the meeting is ended and “Recording is completed” event is received; Only the meeting owner (the person who reserves the meeting on the open platform is considered as reserving person) has the right to perform the operation

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/vc/v1/meetings/:meeting_id/recording/set_permission |
| HTTP Method | PATCH |
| Rate Limit | 1000 per minute、50 per second |
| Supported app types | custom,isv |
| Required scopes | `vc:record` |
| Required field scopes | > The response body of the API contains the following sensitive fields, and they will be returned only after corresponding scopes are added. If you do not need the fields, it is not recommended that you request the scopes. `contact:user.employee_id:readonly` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `meeting_id` | `string` | Meeting ID (Unique identifier of a video conference, generated after the conference is started) **Example value**: "6911188411932033028" | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `user_id_type` | `string` | No | User ID categories **Example value**: "open_id" **Optional values are**:  - `open_id`: Identifies a user to an app. The same user has different Open IDs in different apps. How to get Open ID - `union_id`: Identifies a user to a tenant that acts as a developer. A user has the same Union ID in apps developed by the same developer, and has different Union IDs in apps developed by different developers. A developer can use Union ID to link the same user's identities in multiple apps.How to get Union ID - `user_id`: Identifies a user to a tenant. The same user has different User IDs in different tenants. In one single tenant, a user has the same User ID in all apps （including store apps）. User ID is usually used to communicate user data between different apps. How to get User ID  **Default value**: `open_id` **When the value is `user_id`, the following field scopes are required**: `contact:user.employee_id:readonly` | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `permission_objects` | `recording_permission_object[]` | Yes | Authorized object list |
|   `id` | `string` | No | Authorized object ID **Example value**: "ou_3ec3f6a28a0d08c45d895276e8e5e19b" |
|   `type` | `int` | Yes | Authorized object type **Example value**: 1 **Optional values are**:  - `1`: User authorization (fill in the user ID in the id field) - `2`: Group authorization (fill in the group open_chat_id in the id field) - `3`: Authorize in tenant (id field left blank) - `4`: Authorize in public network (id field left blank)  |
|   `permission` | `int` | Yes | Permissions **Example value**: 1 **Optional values are**:  - `1`: View  |
| `action_type` | `int` | No | Authorize or cancel authorization, default authorization **Example value**: 1 **Optional values are**:  - `0`: Authorize - `1`: Cancel authorization  | ### Request body example

{
    "permission_objects": [
        {
            "id": "ou_3ec3f6a28a0d08c45d895276e8e5e19b",
            "type": 1,
            "permission": 1
        }
    ],
    "action_type": 1
}

## Response

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
| 500 | 121001 | internal error | Internal server error. If retry still doesn't work, contact the administrator. |
| 400 | 121002 | not support | This function is not supported. |
| 400 | 121003 | param error | Parameter error. Check the value range of parameters (please check yourself according to the field description above). |
| 404 | 121004 | data not exist | The requested data doesn't exist. |
| 403 | 121005 | no permission | No permission. Check the token type, operator identity and resource. |
| 404 | 122002 | meeting not exist | The meeting does not exist. |
