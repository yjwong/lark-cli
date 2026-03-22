---
title: "Query meeting room details"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/vc-v1/room/get"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/vc/v1/rooms/:room_id"
service: "vc-v1"
resource: "room"
section: "Video Conferencing"
rate_limit: "1000 per minute、50 per second"
scopes:
  - "vc:room"
  - "vc:room:readonly"
field_scopes:
  - "contact:user.employee_id:readonly"
updated: "1689326166000"
---

# Query meeting room details

This API is used to query a specified meeting room details by room ID.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/vc/v1/rooms/:room_id |
| HTTP Method | GET |
| Rate Limit | 1000 per minute、50 per second |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `vc:room` `vc:room:readonly` |
| Required field scopes | > The response body of the API contains the following sensitive fields, and they will be returned only after corresponding scopes are added. If you do not need the fields, it is not recommended that you request the scopes. `contact:user.employee_id:readonly` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `room_id` | `string` | Room ID **Example value**: "omm_4de32cf10a4358788ff4e09e37ebbf9c" | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `user_id_type` | `string` | No | User ID categories **Example value**: open_id **Optional values are**:  - `open_id`: Identifies a user to an app. The same user has different Open IDs in different apps. How to get Open ID - `union_id`: Identifies a user to a tenant that acts as a developer. A user has the same Union ID in apps developed by the same developer, and has different Union IDs in apps developed by different developers. A developer can use Union ID to link the same user's identities in multiple apps.How to get Union ID - `user_id`: Identifies a user to a tenant. The same user has different User IDs in different tenants. In one single tenant, a user has the same User ID in all apps （including store apps）. User ID is usually used to communicate user data between different apps. How to get User ID  **Default value**: `open_id` **When the value is `user_id`, the following field scopes are required**: `contact:user.employee_id:readonly` | ## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `room` | `room` | Meeting room details |
|     `room_id` | `string` | Room ID |
|     `name` | `string` | Room name |
|     `capacity` | `int` | Room capacity |
|     `description` | `string` | Room description |
|     `display_id` | `string` | Room display ID |
|     `custom_room_id` | `string` | Custom room ID |
|     `room_level_id` | `string` | Parent room level ID |
|     `path` | `string[]` | Room level path |
|     `room_status` | `room_status` | Room status |
|       `status` | `boolean` | Available status of room |
|       `schedule_status` | `boolean` | Future available status of room |
|       `disable_start_time` | `string` | Room disable start time |
|       `disable_end_time` | `string` | Room disable end time |
|       `disable_reason` | `string` | Room disable reason |
|       `contact_ids` | `string[]` | Contact user ID |
|       `disable_notice` | `boolean` | Notice content sent to users when room disabled |
|       `resume_notice` | `boolean` | Notice content sent to contact users when room resumed |
|     `device` | `device[]` | device information |
|       `name` | `string` | device name | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
        "room": {
            "room_id": "omm_4de32cf10a4358788ff4e09e37ebbf9b",
            "name": "TestRoom",
            "capacity": 10,
            "description": "Test Description",
            "display_id": "LM134742334",
            "custom_room_id": "1234",
            "room_level_id": "omb_4ad1a2c7a2fbc5fc9570f38456931293",
            "path": [
                "omb_4ad1a2c7a2fbc5fc9570f38456931293"
            ],
            "room_status": {
                "status": true,
                "schedule_status": true,
                "disable_start_time": "1652356050",
                "disable_end_time": "1652442450",
                "disable_reason": "TestReason",
                "contact_ids": [
                    "ou_3ec3f6a28a0d08c45d895276e8e5e19b"
                ],
                "disable_notice": true,
                "resume_notice": true
            },
            "device": [
                {
                    "name": "电话"
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
| 400 | 121003 | param error | Parameter error. Check the value range of parameters (please check yourself according to the field description above). |
| 404 | 121004 | data not exist | The requested data doesn't exist. |
| 404 | 121005 | no permission | No permission. Check the token type, operator identity and resource. |
| 404 | 121006 | data conflict | This is often caused by CAS. Update the data and try again. |
