---
title: "Query room level list"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/vc-v1/room_level/list"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/vc/v1/room_levels"
service: "vc-v1"
resource: "room_level"
section: "Video Conferencing"
rate_limit: "1000 per minute、50 per second"
scopes:
  - "vc:room"
  - "vc:room:readonly"
updated: "1689326162000"
---

# Query room level list

This API is used to query meeting room level list at a specified meeting room level.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/vc/v1/room_levels |
| HTTP Method | GET |
| Rate Limit | 1000 per minute、50 per second |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `vc:room` `vc:room:readonly` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `room_level_id` | `string` | No | Queries the ID of specifed meeting room level. **Example value**: omb_4ad1a2c7a2fbc5fc9570f38456931293 |
| `page_size` | `int` | No | **Example value**: 10 **Data validation rules**: - Maximum value: `100` |
| `page_token` | `string` | No | Page identifier. It is not filled in the first request, indicating traversal from the beginning; when there will be more groups, the new page_token will be returned at the same time, and the next traversal can use the page_token to get more groups **Example value**: | ## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `items` | `room_level[]` | Returned meeting room level list |
|     `room_level_id` | `string` | Room level ID |
|     `name` | `string` | Room level name |
|     `parent_id` | `string` | Parent room level ID |
|     `path` | `string[]` | Room level path |
|     `has_child` | `boolean` | Whether the room level has child levels |
|     `custom_group_id` | `string` | Custom room level ID |
|   `page_token` | `string` | Page identifier, when has_more is true, a new page_token will also be returned. Otherwise, page_token will not be returned |
|   `has_more` | `boolean` | Whether the response body has more parameters | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
        "items": [
            {
                "room_level_id": "omb_4ad1a2c7a2fbc5fc9570f38456931293",
                "name": "TestRoomLevel",
                "parent_id": "omb_4ad1a2c7a2fbc5fc9570f38456931293",
                "path": [
                    "omb_4ad1a2c7a2fbc5fc9570f38456931293"
                ],
                "has_child": false,
                "custom_group_id": "10000"
            }
        ],
        "page_token": "10",
        "has_more": true
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
