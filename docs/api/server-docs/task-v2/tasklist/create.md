---
title: "Create Tasklist"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/task-v2/tasklist/create"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/task/v2/tasklists"
service: "task-v2"
resource: "tasklist"
section: "Tasks v2"
rate_limit: "1000 per minute、50 per second"
scopes:
  - "task:tasklist:write"
updated: "1699521622000"
---

# Create Tasklist

Create a tasklist. Tasklist can be used to organize and manage multiple tasks belonging to the same project.

When creating, the name of the tasklist must be filled in. At the same time,  tasklist member can be set by specifying `members` field. For the format of member, please refer to the " How to represent members of tasks and tasklists?" section in Feature Overview .

After the tasklist is created, the creator automatically becomes the tasklist owner. If sets the creator to editor/viewer, the user eventually becomes the tasklist owner and automatically disappears from the member list. This is because the same user can only have one role on the same tasklist.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/task/v2/tasklists |
| HTTP Method | POST |
| Rate Limit | 1000 per minute、50 per second |
| Supported app types | custom,isv |
| Required scopes | `task:tasklist:write` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `user_id_type` | `string` | No | User ID categories **Example value**: open_id **Default value**: `open_id` | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `name` | `string` | Yes | Tasklist name, required. Empty is not allowed. Up to 100 characters. **Example value**: "Annual meeting work tasklist" |
| `members` | `member[]` | No | Tasklist member list. For the format of member, please refer to the " How to represent members of tasks and tasklists?" section in Feature Overview . **Data validation rules**: - Maximum length: `500` |
|   `id` | `string` | No | Member ID **Example value**: "ou_2cefb2f014f8d0c6c2d2eb7bafb0e54f" **Data validation rules**: - Maximum length: `100` characters |
|   `type` | `string` | No | Types of members, supported types include:  Ordinary user, in this case the member id should be an ID representing the user, such as open_id. The specific format depends on the user_id_type parameter - `chat`: chat, in which case the member id should be an Open Chat ID - `app`: application, in this case the member id should be the ID of an application  **Example value**: "user" **Default value**: `user` |
|   `role` | `string` | No | Member role, which can be "editor" or "viewer". Defaults to "viewer". **Example value**: "editor" **Data validation rules**: - Maximum length: `20` characters | ### Request body example

{
    "name": "Annual meeting work tasklist",
    "members": [
        {
            "id": "ou_2cefb2f014f8d0c6c2d2eb7bafb0e54f",
            "type": "user",
            "role": "editor"
        }
    ]
}

## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `tasklist` | `tasklist` | Created tasklist |
|     `guid` | `string` | The globally unique ID of the manifest |
|     `name` | `string` | Tasklist name |
|     `creator` | `member` | Tasklist creator |
|       `id` | `string` | Member ID |
|       `type` | `string` | Member type |
|       `role` | `string` | Member role |
|     `owner` | `member` | Tasklist owner |
|       `id` | `string` | Member ID |
|       `type` | `string` | Member type |
|       `role` | `string` | Member role |
|     `members` | `member[]` | Tasklist members |
|       `id` | `string` | Member ID |
|       `type` | `string` | Member type |
|       `role` | `string` | Member role |
|     `url` | `string` | Tasklist applink |
|     `created_at` | `string` | Tasklist creation timestamp (ms) |
|     `updated_at` | `string` | Tasklist recent updated timestamp (ms) | ### Response body example

{
    "code": 0,
    "msg": "Success",
    "data": {
        "tasklist": {
            "guid": "cc371766-6584-cf50-a222-c22cd9055004",
            "name": "Annual meeting summary work task list",
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
| 400 | 1470400 | Incorrect request parameters, such as not filling in the tasklist name. | The reason for the error is shown in the returned msg. |
| 500 | 1470500 | Server error. | Try retrying the call. If it continues to fail, please contact support for feedback. |
