---
title: "Remove Tasklist Members"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/task-v2/tasklist/remove_members"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/task/v2/tasklists/:tasklist_guid/remove_members"
service: "task-v2"
resource: "tasklist"
section: "Tasks v2"
rate_limit: "1000 per minute、50 per second"
scopes:
  - "task:tasklist:write"
updated: "1699521642000"
---

# Remove Tasklist Members

Removes one or more members of the tasklist. Members to remove can be specified by setting `members` field。For the format of member, please refer to the " How to represent members of tasks and tasklists?" section in Feature Overview .

The "role" field does not need to be filled in for the member to be removed in the request, since the id and type of a member can determine the unique member in the tasklist.

If the member to be removed is not in the tasklist, it is automatically ignored and the api returns success.

This api cannot be used to remove the tasklist owner. If the member to be removed is the list owner, it will be automatically ignored. To update the tasklist owner, you need to call Update Tasklist.

> Takslist edit permission is required. Please refer to the "How are tasklist authorized?" section in Tasklist Features Overview .

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/task/v2/tasklists/:tasklist_guid/remove_members |
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
| `tasklist_guid` | `string` | GUID of tasklist from which members are removed. **Example value**: "d300a75f-c56a-4be9-80d1-e47653028ceb" **Data validation rules**: - Maximum length: `100` characters | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `user_id_type` | `string` | No | User ID categories **Example value**: open_id **Default value**: `open_id` | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `members` | `member[]` | Yes | Members to remove.  For the format of member, please refer to the " How to represent members of tasks and tasklists?" section in Feature Overview . **Data validation rules**: - Length range: `1` ～ `500` |
|   `id` | `string` | No | Member ID **Example value**: "ou_2cefb2f014f8d0c6c2d2eb7bafb0e54f" **Data validation rules**: - Maximum length: `100` characters |
|   `type` | `string` | No | Member types, supported types include:  - `user`: ordinary user, in this case member id should be an ID representing the user, such as open_id. The specific format depends on the user_id_type parameter - `chat`: chat, in which case the member id shold be an Open Chat ID - `app`: application, in this case the member ID should be an appid  **Example value**: "user" **Default value**: `user` |
|   `role` | `string` | No | Tasklist role. This field is not needed in this api. **Example value**: "editor" **Data validation rules**: - Maximum length: `20` characters | ### Request body example

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
|   `tasklist` | `tasklist` | Updated tasklist |
|     `guid` | `string` | The globally unique ID of the manifest |
|     `name` | `string` | Tasklist name |
|     `creator` | `member` | Tasklist creator |
|       `id` | `string` | Indicates the id of member |
|       `type` | `string` | member type |
|       `role` | `string` | Member role |
|     `owner` | `member` | Tasklist owner |
|       `id` | `string` | Indicates the id of member |
|       `type` | `string` | member type |
|       `role` | `string` | Member role |
|     `members` | `member[]` | Tasklist members |
|       `id` | `string` | Indicates the id of member |
|       `type` | `string` | member type |
|       `role` | `string` | member role |
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
                "role": "Tasklist creator"
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
    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 1470400 | Incorrect request parameters, such as an invalid Tasklist GUID provided. | The reason for the error is shown in the returned msg. |
| 404 | 1470404 | The Tasklist does not exist or has been deleted. | Verify the tasklist exists. |
| 500 | 1470500 | Server error. | Try retrying the call. If it continues to fail, please contact support for feedback. |
| 403 | 1470403 | Missing Tasklist edit permission. | Check that the calling identity has the edit permission of the tasklist. Please refer to the "How are tasklist authorized?" section in Tasklist Features Overview . |
