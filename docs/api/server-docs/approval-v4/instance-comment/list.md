---
title: "Receive comments"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/approval-v4/instance-comment/list"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/approval/v4/instances/:instance_id/comments"
service: "approval-v4"
resource: "instance-comment"
section: "Approval"
scopes:
  - "approval:approval"
  - "approval:approval:readonly"
field_scopes:
  - "contact:user.employee_id:readonly"
updated: "1675167415000"
---

# Receive comments

Obtain all comments and replies to comments under an approval instance according to Instance Code (excluding approved, rejected, transferred, and other additional reasons or opinions).

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/approval/v4/instances/:instance_id/comments |
| HTTP Method | GET |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `approval:approval` `approval:approval:readonly` |
| Required field scopes | > The response body of the API contains the following sensitive fields, and they will be returned only after corresponding scopes are added. If you do not need the fields, it is not recommended that you request the scopes. `contact:user.employee_id:readonly` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `instance_id` | `string` | Approval instance code (or tenant custom approval instance ID) **Example value**: "6A123516-FB88-470D-A428-9AF58B71B3C0" | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `user_id_type` | `string` | No | User ID categories **Example value**: "open_id" **Optional values are**:  - `open_id`: Identifies a user to an app. The same user has different Open IDs in different apps. How to get Open ID - `union_id`: Identifies a user to a tenant that acts as a developer. A user has the same Union ID in apps developed by the same developer, and has different Union IDs in apps developed by different developers. A developer can use Union ID to link the same user's identities in multiple apps.How to get Union ID - `user_id`: Identifies a user to a tenant. The same user has different User IDs in different tenants. In one single tenant, a user has the same User ID in all apps （including store apps）. User ID is usually used to communicate user data between different apps. How to get User ID  **Default value**: `open_id` **When the value is `user_id`, the following field scopes are required**: `contact:user.employee_id:readonly` |
| `user_id` | `string` | Yes | User ID **Example value**: "e5286g26" |
| `page_token` | `string` | No | Page identifier. It is not filled in the first request, indicating traversal from the beginning; when there will be more groups, the new page_token will be returned at the same time, and the next traversal can use the page_token to get more groups **Example value**: "nF1ZXJ5VGhlbkZldGNoCgAAAAAA6PZwFmUzSldvTC1yU" |
| `page_size` | `int` | No | Page size **Example value**: 10 **Data validation rules**: - Maximum value: `1000` | ## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `comments` | `comment[]` | Comment list (list may be empty) |
|     `id` | `string` | Comment ID |
|     `content` | `string` | Comment content in json format |
|     `create_time` | `string` | Comment creation time |
|     `update_time` | `string` | Comment update time |
|     `is_delete` | `int` | Deletion marks. 0: Not deleted; 1: Deleted |
|     `replies` | `comment_reply[]` | Reply to comment (list may be empty) |
|       `id` | `string` | Comment ID |
|       `content` | `string` | Comment content in json format |
|       `create_time` | `string` | Comment creation time |
|       `update_time` | `string` | Comment update time |
|       `is_delete` | `int` | Deletion marks. 0: Not deleted; 1: Deleted |
|       `at_info_list` | `comment_at_info[]` | List of members who have been mentioned in the comments (list may be empty) |
|         `user_id` | `string` | ID of the member mentioned |
|         `name` | `string` | Name of the member mentioned |
|         `offset` | `string` | Offset in the comment, starting at 0 |
|       `commentator` | `string` | Comment creator |
|       `extra` | `string` | Additional information |
|     `at_info_list` | `comment_at_info[]` | List of members who have been mentioned in the comments (list may be empty) |
|       `user_id` | `string` | ID of the member mentioned |
|       `name` | `string` | Name of the member mentioned |
|       `offset` | `string` | Offset in the comment, starting at 0 |
|     `commentator` | `string` | Comment creator |
|     `extra` | `string` | Additional information | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
        "comments": [
            {
                "id": "7081516627711524883",
                "content": "{\"text\":\"agree\",\"files\":[{\"url\":\"https://xx-xxx-xxx.bytedance.net/lark-approval-attachment/image/20220401/1/d43216ca-93b5-43a8-8a34-23c66820463a.png~tplv-k7bg0smxju-image.image?x-orig-authkey=boeorigin\&x-orig-expires=1650963890\&x-orig-sign=668QhQbRSt6638x2Ws8wFI%2FxqVg%3D#.png\",\"fileSize\":155149,\"title\":\"9a9fedc5cfb01a4a20c715098.png\",\"type\":\"image\"}]}",
                "create_time": "1648801211000",
                "update_time": "1648801211000",
                "is_delete": 1,
                "replies": [
                    {
                        "id": "7081516611634741268",
                        "content": "{\"text\":\"agree\",\"files\":[{\"url\":\"https://xx-xxx-xxx.bytedance.net/lark-approval-attachment/image/20220401/1/d43216ca-93b5-43a8-8a34-23c66820463a.png~tplv-k7bg0smxju-image.image?x-orig-authkey=boeorigin\&x-orig-expires=1650963890\&x-orig-sign=668QhQbRSt6638x2Ws8wFI%2FxqVg%3D#.png\",\"fileSize\":155149,\"title\":\"9a9fedc5cfb01a4a20c715098.png\",\"type\":\"image\"}]}",
                        "create_time": "1648803677000",
                        "update_time": "1648803677000",
                        "is_delete": 0,
                        "at_info_list": [
                            {
                                "user_id": "579fd9c4",
                                "name": "zhangsan",
                                "offset": "1"
                            }
                        ],
                        "commentator": "893g4c45",
                        "extra": "{\"a\":\"a\"}"
                    }
                ],
                "at_info_list": [
                    {
                        "user_id": "579fd9c4",
                        "name": "zhangsan",
                        "offset": "1"
                    }
                ],
                "commentator": "893g4c45",
                "extra": "{\"a\":\"a\"}"
            }
        ]
    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 1390001 | param is invalid | Parameter error |
