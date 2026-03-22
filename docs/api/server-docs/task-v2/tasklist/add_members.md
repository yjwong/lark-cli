---
title: "Add Tasklist Members"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/task-v2/tasklist/add_members"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/task/v2/tasklists/:tasklist_guid/add_members"
service: "task-v2"
resource: "tasklist"
section: "Tasks v2"
rate_limit: "1000 per minute、50 per second"
scopes:
  - "task:tasklist:write"
updated: "1699521638000"
---

# Add Task Members

Add one or more members to a tasklist. Members can be set by setting `members` fields. For the format of member, please refer to the " How to represent members of tasks and tasklists?" section in Feature Overview .

A tasklist member can be a user, app or chat. Each member can as "editor" or "viewer" role. A chat member means that all chatters in chat automatically have the tasklist role of chat member

If the member to be added is already a member and the role is the same as set in the request, it will be automatically ignored and the api returns success.

If the member to be added is already a tasklist member, and the role is not set in the request (for example, the original role is viewer, and the request is set to editor), it is equivalent to updating its role.

If the member to be added is already the tasklist owner, it is automatically ignored. In this case api returns success. The role of tasklist owner does not change.

This api cannot be used to update the tasklist owner. To do it, you can use the Update Tasklist.

> Tasklist edit permission is required. Please refer to the "How are tasklist authorized?" section in Tasklist Features Overview .

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/task/v2/tasklists/:tasklist_guid/add_members |
| HTTP Method | POST |
| Rate Limit | 1000 per minute、50 per second |
| Supported app types | custom,isv |
| Required scopes | `task:tasklist:write` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `tasklist_guid` | `string` | GUID of tasklist to which new members is added **Example value**: "D300a75f-c56a-4be9-80d1-e47653028ceb" | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `user_id_type` | `string` | No | User ID categories **Example value**: open_id **Default value**: `open_id` | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `members` | `member[]` | Yes | Members to add. For the format of member, please refer to the " How to represent members of tasks and tasklists?" section in Feature Overview . **Data validation rules**: - Length range: `1` ～ `500` |
|   `id` | `string` | No | Member ID **Example value**: "ou_2cefb2f014f8d0c6c2d2eb7bafb0e54f" **Data validation rules**: - Maximum length: `100` characters |
|   `type` | `string` | No | Member types, supported types include:  - `user`: Ordinary user, in this case member id should be an ID representing the user, such as open_id. The specific format depends on the user_id_type parameter - `chat`: chat, in which case the member id shold be an Open Chat ID - `app`: application, in this case the member ID should be an appid  **Example value**: "user" **Default value**: `user` |
|   `role` | `string` | No | Member roles. Supported roles includes:  editor viewer  Default is "viewer". The tasklist owner role cannot be set through this field. **Example value**: "editor" **Data validation rules**: - Maximum length: `20` characters | ### Request body example

{
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
|   `tasklist` | `tasklist` | The updated tasklist |
|     `guid` | `string` | GUID of tasklist |
|     `name` | `string` | Taklist name |
|     `creator` | `member` | Taklist creator |
|       `id` | `string` | Indicates the id of member |
|       `type` | `string` | Type of member |
|       `role` | `string` | member type |
|     `owner` | `member` | Taklist owner |
|       `id` | `string` | Indicates the id of member |
|       `type` | `string` | Type of member |
|       `role` | `string` | Taklist role |
|     `members` | `member[]` | Taklist members |
|       `id` | `string` | Indicates the id of member |
|       `type` | `string` | Type of member |
|       `role` | `string` | Member role |
|     `url` | `string` | tasklist applink |
|     `created_at` | `string` | Taklist creation timestamp (ms) |
|     `updated_at` | `string` | Taklist recent updated timestamp (ms) | ### Response body example

{
    "code": 0,
    "msg": "Success",
    "data": {
        "tasklist": {
            "guid": "Cc371766-6584-cf50-a222-c22cd9055004",
            "name": "Annual meeting summary work task list",
            "creator": {
                "id": "ou_2cefb2f014f8d0c6c2d2eb7bafb0e54f",
                "type": "user",
                "role": "Creator"
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
| 400 | 1470400 | Incorrect request parameters, such as an invalid Taklist GUID provided. | The reason for the error is shown in the returned msg. |
| 404 | 1470404 | The taklist does not exist or has been deleted. | Verify that the taklist to which members are added exists or has been deleted. |
| 500 | 1470500 | Server error. | Try retrying the call. If it continues to fail, please contact support for feedback. |
| 400 | 1470612 | The number of tasklist members exceeds the limit. | Call the Get Tasklsit Details(/ssl:ttdoc:/uAjLw4CM/ukTMukTMukTM/task-v2/tasklist/get) api to confirm the number of list members. |
| 403 | 1470403 | Missing edit permission for taklist. | Check if the calling identity has edit permission to the taklist. Please refer to the "How are tasklist authorized?" section in [Tasklist Features Overview . |
