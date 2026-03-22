---
title: "Query room level details"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/vc-v1/room_level/get"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/vc/v1/room_levels/:room_level_id"
service: "vc-v1"
resource: "room_level"
section: "Video Conferencing"
rate_limit: "1000 per minute、50 per second"
scopes:
  - "vc:room"
  - "vc:room:readonly"
updated: "1689326161000"
---

# Query room level details

This API is used to query a specified meeting room level details by room level ID.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/vc/v1/room_levels/:room_level_id |
| HTTP Method | GET |
| Rate Limit | 1000 per minute、50 per second |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `vc:room` `vc:room:readonly` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `room_level_id` | `string` | Queries the ID of specifed meeting room level. **Example value**: "omb_4ad1a2c7a2fbc5fc9570f38456931293" | ## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `room_level` | `room_level` | Room level details |
|     `room_level_id` | `string` | Room level ID |
|     `name` | `string` | Room level name |
|     `parent_id` | `string` | Parent room level ID |
|     `path` | `string[]` | Room level path |
|     `has_child` | `boolean` | Whether the room level has child levels |
|     `custom_group_id` | `string` | Custom room level ID | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
        "room_level": {
            "room_level_id": "omb_4ad1a2c7a2fbc5fc9570f38456931293",
            "name": "TestRoomLevel",
            "parent_id": "omb_4ad1a2c7a2fbc5fc9570f38456931293",
            "path": [
                "omb_4ad1a2c7a2fbc5fc9570f38456931293"
            ],
            "has_child": false,
            "custom_group_id": "10000"
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
