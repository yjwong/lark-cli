---
title: "Delete Task"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/task-v2/task/delete"
method: "DELETE"
api_path: "https://open.larksuite.com/open-apis/task/v2/tasks/:task_guid"
service: "task-v2"
resource: "task"
section: "Tasks v2"
rate_limit: "10 per second"
scopes:
  - "task:task:write"
updated: "1699521567000"
---

# Delete Task

Delete a task.

Tasks can no longer be restored after deletion.

> Deleting a task requires edit permissions on the task. Please refer to the "How are Tasks Authorized?" section in Task Features Overview.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/task/v2/tasks/:task_guid |
| HTTP Method | DELETE |
| Rate Limit | 10 per second |
| Supported app types | custom,isv |
| Required scopes | `task:task:write` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `task_guid` | `string` | GUID of task to delete **Example value**: "e297ddff-06ca-4166-b917-4ce57cd3a7a0" **Data validation rules**: - Maximum length: `100` characters | ## Response

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
| 404 | 1470404 | The task to delete does not exist or has been deleted. | Verify that the task data you want to delete exists or has been deleted. |
| 500 | 1470500 | Server error. | Try retrying the call. If it continues to fail, please contact support for feedback. |
| 400 | 1470400 | Incorrect request parameters. For example, the incoming task GUID is invalid. | The reason for the error is shown in the returned msg prompt. |
| 403 | 1470403 | Missing permission to delete tasks. | Check that the calling identity has edit permissions on the task. Please refer to the "How are Tasks Authorized?" section in Task Features Overview. |
