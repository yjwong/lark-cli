---
title: "Get Tasks of Tasklist"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/task-v2/tasklist/tasks"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/task/v2/tasklists/:tasklist_guid/tasks"
service: "task-v2"
resource: "tasklist"
section: "Tasks v2"
rate_limit: "1000 per minute、50 per second"
scopes:
  - "task:tasklist:read"
  - "task:tasklist:write"
updated: "1699521646000"
---

# Get Tasks of Tasklist

Gets the summary of tasks belonging to a tasklist. 

This API supports pagination. Tasks in the tasklist are returned in the "custom" order.

This API supports simple filtering by task completion status or task creation time range.

> Tasklist read permission is required. Please refer to the "How are tasklist authorized?" section in Tasklist Features Overview .

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/task/v2/tasklists/:tasklist_guid/tasks |
| HTTP Method | GET |
| Rate Limit | 1000 per minute、50 per second |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `task:tasklist:read` `task:tasklist:write` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `tasklist_guid` | `string` | GUID of the tasklist **Example value**: "d300a75f-c56a-4be9-80d1-e47653028ceb" | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `page_size` | `int` | No | The number of tasks in one page **Example value**: 50 **Default value**: `50` **Data validation rules**: - Value range: `1` ～ `100` |
| `page_token` | `string` | No | Page identifier. It is not filled in the first request, indicating traversal from the beginning; when there will be more groups, the new page_token will be returned at the same time, and the next traversal can use the page_token to get more groups **Example value**: aWQ9NzEwMjMzMjMxMDE= **Data validation rules**: - Maximum length: `100` characters |
| `completed` | `boolean` | No | Only return tasks with a specific completion status. Filling in "true" means returning completed tasks; "false" means returning only uncompleted tasks; not filling in means not filtering by completion status. **Example value**: true |
| `created_from` | `string` | No | Task creation start timestamp (ms), inclusive. Default is the timestamp of the first created task int the tasklist. **Example value**: 1675742789470 |
| `created_to` | `string` | No | Task end start timestamp (ms), inclusive. Default is the timestamp of the last created task int the tasklist. **Example value**: 1675742789470 |
| `user_id_type` | `string` | No | User ID categories **Example value**: open_id **Default value**: `open_id` | ## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `items` | `task_summary[]` | Task Summary Data |
|     `guid` | `string` | Task GUID |
|     `summary` | `string` | Task summary |
|     `completed_at` | `string` | The timestamp (ms) of the task completion, 0 means not completed |
|     `start` | `start` | Task start time |
|       `timestamp` | `string` | The timestamp of the start time/date in milliseconds from 1970-01-01 00:00:00. If the start time is a date, you need to convert the date to timestamp and set is_all_day= true |
|       `is_all_day` | `boolean` | Whether to start on a date. |
|     `due` | `due` | Task due |
|       `timestamp` | `string` | The task due represented by timestamp of the time/date in milliseconds from 1970-01-01 00:00:00. |
|       `is_all_day` | `boolean` | Whether due on a date. If set to true, only the date part of the timestamp will be parsed and stored. |
|     `members` | `member[]` | Task member list |
|       `id` | `string` | Indicates the id of member |
|       `type` | `string` | Type of member |
|       `role` | `string` | Member role |
|     `subtask_count` | `int` | Number of subtasks |
|   `page_token` | `string` | Page identifier, when has_more is true, a new page_token will also be returned. Otherwise, page_token will not be returned |
|   `has_more` | `boolean` | Whether the response body has more parameters | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
        "items": [
            {
                "guid": "e297ddff-06ca-4166-b917-4ce57cd3a7a0",
                "summary": "annual summary",
                "completed_at": "1675742789470",
                "start": {
                    "timestamp": "1675454764000",
                    "is_all_day": true
                },
                "due": {
                    "timestamp": "1675454764000",
                    "is_all_day": true
                },
                "members": [
                    {
                        "id": "ou_2cefb2f014f8d0c6c2d2eb7bafb0e54f",
                        "type": "user",
                        "role": "editor"
                    }
                ],
                "subtask_count": 1
            }
        ],
        "page_token": "aWQ9NzEwMjMzMjMxMDE=",
        "has_more": true
    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 1470400 | The request parameter is incorrect, such as the tasklist GUID is invalid, or the page_size is a negative number, etc. | The reason for the error is shown in the returned msg. |
| 404 | 1470404 | The tasklist does not exist or has been deleted. | Verify that the tasklist to get exists or has been deleted. |
| 500 | 1470500 | Server error. | Try retrying the call. If it continues to fail, please contact support for feedback. |
| 403 | 1470403 | Missing read permission for tasklist. | Check if the calling identity has read permission to the tasklist. See How is the tasklist authorized? |
