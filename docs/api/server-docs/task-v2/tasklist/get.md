---
title: "Get Tasklist Details"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/task-v2/tasklist/get"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/task/v2/tasklists/:tasklist_guid"
service: "task-v2"
resource: "tasklist"
section: "Tasks v2"
rate_limit: "1000 per minute、50 per second"
scopes:
  - "task:tasklist:read"
  - "task:tasklist:write"
updated: "1699521626000"
---

# Get Tasklist Details

Get the details of a tasklist, including list name, owner, list members, etc.

> Require tasklist read permission. Please refer to the "How are tasklist authorized?" section in Tasklist Features Overview .

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/task/v2/tasklists/:tasklist_guid |
| HTTP Method | GET |
| Rate Limit | 1000 per minute、50 per second |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `task:tasklist:read` `task:tasklist:write` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `tasklist_guid` | `string` | Tasklist GUID **Example value**: "d300a75f-c56a-4be9-80d1-e47653028ceb" | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `user_id_type` | `string` | No | User ID categories **Example value**: open_id **Default value**: `open_id` | ## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `tasklist` | `tasklist` | Tasklist details |
|     `guid` | `string` | The globally unique ID of the tasklist |
|     `name` | `string` | Tasklist name |
|     `creator` | `member` | Tasklist creator |
|       `id` | `string` | Member ID |
|       `type` | `string` | Member Type |
|       `role` | `string` | Member role |
|     `owner` | `member` | Tasklist owner |
|       `id` | `string` | Member ID |
|       `type` | `string` | Member Type |
|       `role` | `string` | Member role |
|     `members` | `member[]` | Tasklist members |
|       `id` | `string` | Member ID |
|       `type` | `string` | Member Type |
|       `role` | `string` | Member role |
|     `url` | `string` | The tasklist applink |
|     `created_at` | `string` | Tasklist creation timestamp (ms) |
|     `updated_at` | `string` | Tasklist recent updated timestamp (ms) | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
        "tasklist": {
            "guid": "cc371766-6584-cf50-a222-c22cd9055004",
            "name": "年会总结工作任务清单",
            "creator": {
                "id": "ou_2cefb2f014f8d0c6c2d2eb7bafb0e54f",
                "type": "user",
                "role": "editor"
            },
            "owner": {
                "id": "ou_2cefb2f014f8d0c6c2d2eb7bafb0e54f",
                "type": "user",
                "role": "editor"
            },
            "members": [
                {
                    "id": "ou_2cefb2f014f8d0c6c2d2eb7bafb0e54f",
                    "type": "user",
                    "role": "editor"
                }
            ],
            "url": "https://applink.larksuite.com/client/todo/task_list?guid=b45b360f-1961-4058-b338-7f50c96e1b52",
            "created_at": "1675742789470",
            "updated_at": "1675742789470"
        }
    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 1470400 | Incorrect request parameters, such as an invalid tasklist GUID is provided. | The reason for the error is shown in the returned msg. |
| 404 | 1470404 | The tasklist does not exist or has been deleted. | Confirm that the tasklist you want to get exists or has been deleted. |
| 500 | 1470500 | Server error. | Try retrying the call. If it continues to fail, please contact support for feedback. |
| 403 | 1470403 | Missing tasklist read permission. | Check if the calling identity has read permission to the tasklist. Please refer to the "How are tasklist authorized?" section in Tasklist Features Overview . |
