---
title: "Get Comment"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/task-v2/comment/get"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/task/v2/comments/:comment_id"
service: "task-v2"
resource: "comment"
section: "Tasks v2"
rate_limit: "1000 per minute、50 per second"
scopes:
  - "task:comment:read"
  - "task:comment:write"
updated: "1699521682000"
---

# Get Comment Details

Given the ID of a comment, return the details of the comment, including information such as content, creator, creation time and update time.

> Getting comment of a task requires read permission of the task. Please refer to the "How are Tasks Authorized?" section in Task Features Overview.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/task/v2/comments/:comment_id |
| HTTP Method | GET |
| Rate Limit | 1000 per minute、50 per second |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `task:comment:read` `task:comment:write` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `comment_id` | `string` | Comment ID to get comment details **Example value**: "7198104824246747156" **Data validation rules**: - Maximum length: `50` characters | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `user_id_type` | `string` | No | User ID categories **Example value**: open_id **Default value**: `open_id` | ## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `comment` | `comment` | Review details |
|     `id` | `string` | Comment id |
|     `content` | `string` | Comment content |
|     `creator` | `member` | Comment creator |
|       `id` | `string` | Indicates the id of member |
|       `type` | `string` | Type of member |
|       `role` | `string` | Member role |
|     `reply_to_comment_id` | `string` | The id to reply to the comment. Empty if not replying to the comment. |
|     `created_at` | `string` | Comment creation timestamp (ms) |
|     `updated_at` | `string` | Comments Updatetimestamp (ms) |
|     `resource_type` | `string` | The resource type associated with the task |
|     `resource_id` | `string` | Resource ID associated with the task | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
        "comment": {
            "id": "7197020628442939411",
            "content": "This is a comment",
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
    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 1470400 | Incorrect request parameters, such as an invalid comment ID entered. | The reason for the error is shown in the returned msg prompt. |
| 404 | 1470404 | The comment to be retrieved does not exist or has been deleted. | Confirm that the comment you want to access exists or has been deleted. |
| 500 | 1470500 | Server error. | Try retrying the call. If it continues to fail, please contact support for feedback. |
| 403 | 1470403 | Missing permission to view comments. | Getting comment of task requires task read permission. Please refer to the "How are Tasks Authorized?" section in Task Features Overview. |
