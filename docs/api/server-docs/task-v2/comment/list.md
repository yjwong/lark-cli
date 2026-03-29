---
title: "List Comments"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/task-v2/comment/list"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/task/v2/comments"
service: "task-v2"
resource: "comment"
section: "Tasks v2"
rate_limit: "1000 per minute、50 per second"
scopes:
  - "task:comment:read"
  - "task:comment:write"
updated: "1699521694000"
---

# List Comments

Given a resource, returns a list of comments for that resource.

Pagination is supported. Comments can return data in positive order (asc, oldest to newest) or reverse order (desc, oldest to newest) of creation time.

> Listing comment of a task requires the read permission of the task. Please refer to the "How are Tasks Authorized?" section in Task Features Overview.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/task/v2/comments |
| HTTP Method | GET |
| Rate Limit | 1000 per minute、50 per second |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `task:comment:read` `task:comment:write` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `page_size` | `int` | No | Page size, default is 50. **Example value**: 50 **Default value**: `50` **Data validation rules**: - Value range: `1` ～ `100` |
| `page_token` | `string` | No | Page identifier. It is not filled in the first request, indicating traversal from the beginning; when there will be more groups, the new page_token will be returned at the same time, and the next traversal can use the page_token to get more groups **Example value**: aWQ9NzEwMjMzMjMxMDE= **Data validation rules**: - Maximum length: `100` characters |
| `resource_type` | `string` | No | To get the resource type of the comment list, only "task" is currently supported, and the default is "task". **Example value**: task **Default value**: `task` |
| `resource_id` | `string` | Yes | To get the resource ID of the comment. For example, to get the comment list of the task, the globally unique ID of the task should be filled in here **Example value**: d300a75f-c56a-4be9-80d1-e47653028ceb |
| `direction` | `string` | No | Returns how the data is sorted. "Asc" means return from oldest to newest order; "desc" means return from newest to oldest order. Default is "asc". **Example value**: asc **Optional values are**:  - `asc`: Comments posted in ascending order - `desc`: Comments posted in descending order  **Default value**: `asc` |
| `user_id_type` | `string` | No | User ID categories **Example value**: open_id **Default value**: `open_id` | ## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `items` | `comment[]` | Comment list data |
|     `id` | `string` | Comment id |
|     `content` | `string` | Comment content |
|     `creator` | `member` | Comment creator |
|       `id` | `string` | Indicates the id of member |
|       `type` | `string` | Type of member |
|       `role` | `string` | Member role |
|     `reply_to_comment_id` | `string` | The ID of the comment being replied to, or empty if it is not a reply comment. |
|     `created_at` | `string` | Comment creation timestamp (ms) |
|     `updated_at` | `string` | Comments Updatetimestamp (ms) |
|     `resource_type` | `string` | The resource type associated with the task |
|     `resource_id` | `string` | Resource ID associated with the task |
|   `page_token` | `string` | Page identifier, when has_more is true, a new page_token will also be returned. Otherwise, page_token will not be returned |
|   `has_more` | `boolean` | Whether the response body has more parameters | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
        "items": [
            {
                "id": "7197020628442939411",
                "content": "这是一条评论",
                "creator": {
                    "id": "ou_2cefb2f014f8d0c6c2d2eb7bafb0e54f",
                    "type": "user",
                    "role": "editor"
                },
                "reply_to_comment_id": "7166825117308174356",
                "created_at": "1675742789470",
                "updated_at": "1675742789470",
                "resource_type": "task",
                "resource_id": "ccb55625-95d2-2e80-655f-0e40bf67953f"
            }
        ],
        "page_token": "aWQ9NzEwMjMzMjMxMDE=",
        "has_more": true
    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 1470400 | Incorrect request parameters, such as an unsupported resource_type. | The reason for the error is shown in the returned msg prompt. |
| 404 | 1470404 | The resource to get the list does not exist or has been deleted. | Confirm that the resource to get the list does not exist or has been deleted. |
| 500 | 1470500 | Server error. | Try retrying the call. If it continues to fail, please contact support for feedback. |
| 403 | 1470403 | Missing permission to listing commments. | Listing comment of a task requires the read permission of the task. Please refer to the "How are Tasks Authorized?" section in Task Features Overview. |
