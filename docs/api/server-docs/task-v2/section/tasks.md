---
title: "List Tasks of Section"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/task-v2/section/tasks"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/task/v2/sections/:section_guid/tasks"
service: "task-v2"
resource: "section"
section: "Tasks v2"
rate_limit: "1000 per minute、50 per second"
scopes:
  - "task:section:read"
  - "task:section:write"
updated: "1699521741000"
---

# List Tasks of Section

List tasks of a section. Paging is supported. Tasks are returned in the order as "custom" order of UI. This API supports simple filtering.

> Require read permission of the resource to which the section belongs.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/task/v2/sections/:section_guid/tasks |
| HTTP Method | GET |
| Rate Limit | 1000 per minute、50 per second |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `task:section:read` `task:section:write` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `section_guid` | `string` | section GUID of which tasks are listed. **Example value**: "9842501a-9f47-4ff5-a622-d319eeecb97f" | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `page_size` | `int` | No | paging size **Example value**: 50 **Default value**: `50` **Data validation rules**: - Value range: `1` ～ `100` |
| `page_token` | `string` | No | Page identifier. It is not filled in the first request, indicating traversal from the beginning; when there will be more groups, the new page_token will be returned at the same time, and the next traversal can use the page_token to get more groups **Example value**: aWQ9NzEwMjMzMjMxMDE= |
| `completed` | `boolean` | No | Filter by task status. If you don't fill it in, it means you don't filter by completion status. **Example value**: true |
| `created_from` | `string` | No | Start timestamp (ms) filtered by creation time. Optional, leaving empty means the creation time of the first task. **Example value**: 1675742789470 **Data validation rules**: - Maximum length: `30` characters |
| `created_to` | `string` | No | The starting timestamp (ms) filtered by the creation time. Optional, leaving empty means the creation time of the first task. **Example value**: 1675742789470 **Data validation rules**: - Maximum length: `30` characters | ## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `items` | `task_summary[]` | Task Summaries |
|     `guid` | `string` | Task GUID |
|     `summary` | `string` | Task summary |
|     `completed_at` | `string` | The timestamp (ms) of the task completion, 0 means not completed |
|     `start` | `start` | Task start time |
|       `timestamp` | `string` | The timestamp of the start time/date in milliseconds from 1970-01-01 00:00:00. If the start time is a date, you need to convert the date to timestamp and set is_all_day = true |
|       `is_all_day` | `boolean` | Whether to start on a date. If set to true, only the date part of the timestamp will be parsed and stored. |
|     `due` | `due` | Task due |
|       `timestamp` | `string` | The timestamp of the due time/date in milliseconds from 1970-01-01 00:00:00. If the expiration time is a date, you need to convert the date to timestamp and set is_all_day = true |
|       `is_all_day` | `boolean` | Whether to due on a date. If set to true, only the date part of the timestamp will be parsed and stored. |
|     `members` | `member[]` | Task member list |
|       `id` | `string` | Member ID |
|       `type` | `string` | Type of member |
|       `role` | `string` | Member Role |
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
                        "role": "assignee"
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
| 400 | 1470400 | The request parameter is incorrect, such as the `section_guid` provided is not valid. | The reason for the error is shown in the returned msg prompt. |
| 404 | 1470404 | The section to access does not exist or has been deleted. | Make sure the section you want to access does not exist or has been deleted. |
| 500 | 1470500 | Server error. | Retry to invoke the api with the same parameters. If it continues to fail, please contact support for feedback. |
| 403 | 1470403 | Missing permission to list tasks of section. | Confirm that the calling identity has read permission of the resource to which the section belongs. |
