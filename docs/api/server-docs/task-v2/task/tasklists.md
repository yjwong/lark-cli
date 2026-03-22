---
title: "List tasklists of task"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/task-v2/task/tasklists"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/task/v2/tasks/:task_guid/tasklists"
service: "task-v2"
resource: "task"
section: "Tasks v2"
rate_limit: "1000 per minute、50 per second"
scopes:
  - "task:task:read"
  - "task:task:write"
updated: "1699521582000"
---

# List Tasklists of Task

List the all lists where a task belongs to, including the tasklist GUID and section GUID.

Only the tasklists that calling identity has read permission will be returned.

> Requires the calling identity to have read access to the task. Please refer to the "How are Tasks Authorized?" section in Task Features Overview.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/task/v2/tasks/:task_guid/tasklists |
| HTTP Method | GET |
| Rate Limit | 1000 per minute、50 per second |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `task:task:read` `task:task:write` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `task_guid` | `string` | GUID of task to list tasklists. **Example value**: "d300a75f-c56a-4be9-80d1-e47653028ceb" | ## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `tasklists` | `task_in_tasklist_info[]` | Summary information of tasklists |
|     `tasklist_guid` | `string` | The guid of the task list |
|     `section_guid` | `string` | Section GUID | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
        "tasklists": [
            {
                "tasklist_guid": "cc371766-6584-cf50-a222-c22cd9055004",
                "section_guid": "e6e37dcc-f75a-5936-f589-12fb4b5c80c2"
            }
        ]
    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 1470400 | Incorrect request parameters, such as an invalid task GUID filled in. | The reason for the error is shown in the returned msg prompt. |
| 404 | 1470404 | Task data does not exist or has been deleted. | Confirm that the task data you want to access exists or has been deleted. |
| 500 | 1470500 | Server error. | Try retrying the call. If it continues to fail, please contact support for feedback. |
| 403 | 1470403 | Missing task read permission. | Check if the calling identity has read permission for the task. Please refer to the "How are Tasks Authorized?" section in Task Features Overview. |
