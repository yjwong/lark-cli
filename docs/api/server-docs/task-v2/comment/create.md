---
title: "Create Comment"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/task-v2/comment/create"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/task/v2/comments"
service: "task-v2"
resource: "comment"
section: "Tasks v2"
rate_limit: "1000 per minute、50 per second"
scopes:
  - "task:comment:write"
updated: "1699521678000"
---

# Create Comment

Create a comment for a task, or reply to a comment for that task.

To create a reply comment, you need to set the reply_to_comment_id field at creation time. The comment being replied to and the new comment must belong to the same task.

> Creating comments for task requires read permission for this task. Please refer to the "How are Tasks Authorized?" section in Task Features Overview.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/task/v2/comments |
| HTTP Method | POST |
| Rate Limit | 1000 per minute、50 per second |
| Supported app types | custom,isv |
| Required scopes | `task:comment:write` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `user_id_type` | `string` | No | User ID categories **Example value**: open_id **Default value**: `open_id` | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `content` | `string` | Yes | Comment content. Empty is not allowed. Supports up to 3000 utf8 characters. **Example value**: "这是一条评论。" |
| `reply_to_comment_id` | `string` | No | Reply to the comment ID of the comment. If not filled in it means creating a non-reply comment. **Example value**: "6937231762296684564" **Data validation rules**: - Maximum length: `20` characters |
| `resource_type` | `string` | No | The resource type to which the comment belongs, currently only supports task "task", the default is "task". **Example value**: "task" **Default value**: `task` |
| `resource_id` | `string` | No | The resource ID of the comment attribution. When the attributed resource type is "task", the GUID of the task should be filled in here. **Example value**: "ccb55625-95d2-2e80-655f-0e40bf67953f" **Data validation rules**: - Maximum length: `100` characters | ### Request body example

{
    "content": "这是一条评论。",
    "reply_to_comment_id": "6937231762296684564",
    "resource_type": "task",
    "resource_id": "ccb55625-95d2-2e80-655f-0e40bf67953f"
}

## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `comment` | `comment` | Created comment details |
|     `id` | `string` | Comment id |
|     `content` | `string` | Comment content |
|     `creator` | `member` | Comment creator |
|       `id` | `string` | Indicates the id of member |
|       `type` | `string` | Type of member |
|       `role` | `string` | Member role |
|     `reply_to_comment_id` | `string` | Comment The id of the reply comment. Empty if not replying to the comment. |
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
    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 1470400 | Parameter errors, such as comments that are too long. | The reason for the error is shown in the returned msg prompt. |
| 404 | 1470404 | The resource to comment on does not exist or has been deleted. | Check if the resource you want to comment on exists or has been deleted. |
| 500 | 1470500 | Server error. | Try retrying the call. If it continues to fail, please contact support for feedback. |
| 403 | 1470403 | Lack of comment permission. | Read permission to comment is required when creating comment for a task. Please refer to the "How are Tasks Authorized?" section in Task Features Overview. |
