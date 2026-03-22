---
title: "List Tasklists"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/task-v2/tasklist/list"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/task/v2/tasklists"
service: "task-v2"
resource: "tasklist"
section: "Tasks v2"
rate_limit: "1000 per minute、50 per second"
scopes:
  - "task:tasklist:read"
  - "task:tasklist:write"
updated: "1699521650000"
---

# List Tasklists

List all the tasklists the calling identity has read permission.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/task/v2/tasklists |
| HTTP Method | GET |
| Rate Limit | 1000 per minute、50 per second |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `task:tasklist:read` `task:tasklist:write` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `page_size` | `int` | No | Number of tasklist returned in one page **Example value**: 50 **Default value**: `50` **Data validation rules**: - Value range: `1` ～ `100` |
| `page_token` | `string` | No | Page identifier. It is not filled in the first request, indicating traversal from the beginning; when there will be more groups, the new page_token will be returned at the same time, and the next traversal can use the page_token to get more groups **Example value**: aWQ9NzEwMjMzMjMxMDE = |
| `user_id_type` | `string` | No | User ID categories **Example value**: open_id **Default value**: `open_id` | ## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `items` | `tasklist[]` | List of inventory data |
|     `guid` | `string` | Tasklist GUID |
|     `name` | `string` | Tasklist name |
|     `creator` | `member` | Tasklist creator |
|       `id` | `string` | Indicates the id of member |
|       `type` | `string` | member type |
|       `role` | `string` | member role |
|     `owner` | `member` | Tasklist owner |
|       `id` | `string` | Indicates the id of member |
|       `type` | `string` | member type |
|       `role` | `string` | Member role |
|     `members` | `member[]` | Tasklist members |
|       `id` | `string` | Indicates the id of member |
|       `type` | `string` | memers type |
|       `role` | `string` | Member role |
|     `url` | `string` | Tasklist  applink |
|     `created_at` | `string` | Tasklist creation timestamp (ms) |
|     `updated_at` | `string` | Tasklist recent updated timestamp (ms) |
|   `page_token` | `string` | Page identifier, when has_more is true, a new page_token will also be returned. Otherwise, page_token will not be returned |
|   `has_more` | `boolean` | Whether the response body has more parameters | ### Response body example

{
    "code": 0,
    "msg": "Success",
    "data": {
        "items": [
            {
                "guid": "cc371766-6584-cf50-a222-c22cd9055004",
                "name": "Annual meeting summary work task list",
                "creator": {
                    "id": "ou_2cefb2f014f8d0c6c2d2eb7bafb0e54f",
                    "type": "user",
                    "role": "creator"
                },
                "owner": {
                    "id": "ou_2cefb2f014f8d0c6c2d2eb7bafb0e54f",
                    "type": "user",
                    "role": "owner"
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
        ],
        "page_token": "aWQ9NzEwMjMzMjMxMDE =",
        "has_more": true
    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 1470400 | Incorrect request parameters, such as setting the page_size to a negative number, etc. | The reason for the error is shown in the returned msg. |
| 500 | 1470500 | Server error. | Try retrying the call. If it continues to fail, please contact support for feedback. |
