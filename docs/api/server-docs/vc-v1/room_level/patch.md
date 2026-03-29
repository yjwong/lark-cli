---
title: "Update room level"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/vc-v1/room_level/patch"
method: "PATCH"
api_path: "https://open.larksuite.com/open-apis/vc/v1/room_levels/:room_level_id"
service: "vc-v1"
resource: "room_level"
section: "Video Conferencing"
rate_limit: "20 per minute"
scopes:
  - "vc:room"
updated: "1689326161000"
---

# Update room level

This API is used to update a specified meeting room level details.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/vc/v1/room_levels/:room_level_id |
| HTTP Method | PATCH |
| Rate Limit | 20 per minute |
| Supported app types | custom,isv |
| Required scopes | `vc:room` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `room_level_id` | `string` | Room level ID **Example value**: "omb_4ad1a2c7a2fbc5fc9570f38456931293" **Data validation rules**: - Length range: `1` ～ `100` characters | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `name` | `string` | Yes | Room level name **Example value**: "TestRoomLevel" |
| `parent_id` | `string` | Yes | Parent room level ID **Example value**: "omb_4ad1a2c7a2fbc5fc9570f38456931293" |
| `custom_group_id` | `string` | No | Custom room level ID **Example value**: "10000" | ### Request body example

{
    "name": "TestRoomLevel",
    "parent_id": "omb_4ad1a2c7a2fbc5fc9570f38456931293",
    "custom_group_id": "10000"
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
| 404 | 121005 | no permission | No permission. Check the token type, operator identity and resource. |
| 404 | 121006 | data conflict | This is often caused by CAS. Update the data and try again. |
