---
title: "Delete Comment"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/task-v2/comment/delete"
method: "DELETE"
api_path: "https://open.larksuite.com/open-apis/task/v2/comments/:comment_id"
service: "task-v2"
resource: "comment"
section: "Tasks v2"
rate_limit: "1000 per minute、50 per second"
scopes:
  - "task:comment:write"
updated: "1699521690000"
---

# Delete Comment

Delete a comment.

After a comment is deleted, nothing can be done and it cannot be restored.

> When deleting a comment of a task, read permission of that task is required. Besides, comment can only deleted by its creator. Please refer to the "How are Tasks Authorized?" section in Task Features Overview.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/task/v2/comments/:comment_id |
| HTTP Method | DELETE |
| Rate Limit | 1000 per minute、50 per second |
| Supported app types | custom,isv |
| Required scopes | `task:comment:write` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `comment_id` | `string` | Comment id to delete **Example value**: "7198104824246747156" | ## Response

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
| 400 | 1470400 | Incorrect request parameters, such as an invalid comment ID entered. | The reason for the error is shown in the returned msg prompt. |
| 404 | 1470404 | The comment to delete does not exist or has been deleted. | Confirm that the task data you want to access exists or has been deleted. |
| 500 | 1470500 | Server error. | Try retrying the call. If it continues to fail, please contact support for feedback. |
| 403 | 1470403 | Missing permission to delete comments. | Deleting task comments requires the read permission of the task to which the comment belongs, and can only delete comments sent by yourself. Please refer to the "How are Tasks Authorized?" section in Task Features Overview. |
