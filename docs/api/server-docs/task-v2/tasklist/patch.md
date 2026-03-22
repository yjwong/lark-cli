---
title: "Patch Tasklist"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/task-v2/tasklist/patch"
method: "PATCH"
api_path: "https://open.larksuite.com/open-apis/task/v2/tasklists/:tasklist_guid"
service: "task-v2"
resource: "tasklist"
section: "Tasks v2"
rate_limit: "1000 per minute、50 per second"
scopes:
  - "task:tasklist:write"
updated: "1699521630000"
---

# Patch Tasklist

Update a tasklist. You can update the name and owner of the tasklist.

When updating the tasklist, fill in all the tasklist field names to be modified in the 'update_fields 'field, and fill in the new value of the field to be modified in the'tasklist' field. For details of update api specification, please refer to the "About Resource Update" section in Feature Overview .

Updatable fields include:

* `name`: tasklist name
* `owner`: tasklist owner

When updating tasklist owner, if the new owner is already an "editor" or "viewer", it will be directly upgraded to the owner and automatically disappear from the member list. This is because the same user can only have one role in the same tasklist. Meanwhile, you can set the new role of original owner by setting `origin_owner_to_role` field.

This api cannot be used to update members of the tasklist and add or delete tasks in the tasklist.
* To add or delete members from the tasklist, you can use the Add Tasklist Member and Remove Tasklist Member api.
* To add or delete tasks in the tasklist, you can use the Add Task to Tasklist and Remove Task from Tasklist api.

> Updating the tasklist name requires edit permission. Updating owner must have manage permission. Please refer to the "How are tasklist authorized?" section in Tasklist Features Overview .

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/task/v2/tasklists/:tasklist_guid |
| HTTP Method | PATCH |
| Rate Limit | 1000 per minute、50 per second |
| Supported app types | custom,isv |
| Required scopes | `task:tasklist:write` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `tasklist_guid` | `string` | Globally unique GUID of the manifest to update **Example value**: "D300a75f-c56a-4be9-80d1-e47653028ceb" | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `user_id_type` | `string` | No | User ID categories **Example value**: open_id **Default value**: `open_id` | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `tasklist` | `input_tasklist` | Yes | Data to update the manifest |
|   `name` | `string` | No | List name. To update, cannot be set to null. Maximum 100 characters. **Example value**: "Annual meeting work task list" |
|   `owner` | `member` | No | Updated list owner. |
|     `id` | `string` | No | Indicates the id of member **Example value**: "ou_2cefb2f014f8d0c6c2d2eb7bafb0e54f" **Data validation rules**: - Maximum length: `100` characters |
|     `type` | `string` | No | The type of member, which can be "user" or "app". The type of owner cannot be "chat". **Example value**: "user" **Default value**: `user` |
|     `role` | `string` | No | Member role, which must be "owner" at this point **Example value**: "Owner" **Data validation rules**: - Maximum length: `20` characters |
| `update_fields` | `string[]` | Yes | Field name to update, support  - `name`: Tasklist name - `owner`: Tasklist owner  **Example value**: ["owner"] **Data validation rules**: - Minimum length: `1` |
| `origin_owner_to_role` | `string` | No | This field indicates that if a new owner is updated, the original owner is set to the specified collaborator role. Effective only when the manifest owner is updated. Available options are "editor", "viewer" or "none", Default option is "none". If not set of setting to "none", the original tasklist owner will not have any roles. If not auhorized through other channels (such as through chat members), the original tasklist owner will lose all permissions to the tasklist. **Example value**: "editor" **Optional values are**:  - `editor`: Original owner becomes tasklist editor - `viewer`: Original owner becomes tasklist viewer - `none`: The original owner loses role of the tasklist  **Default value**: `none` | ### Request body example

{
    "tasklist": {
        "name": "Annual meeting work task list",
        "owner": {
            "id": "ou_2cefb2f014f8d0c6c2d2eb7bafb0e54f",
            "type": "user",
            "role": "Owner"
        }
    },
    "update_fields": [
        "owner"
    ],
    "origin_owner_to_role": "editor"
}

## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `tasklist` | `tasklist` | Revised task list |
|     `guid` | `string` | The globally unique ID of the manifest |
|     `name` | `string` | tasklist name |
|     `creator` | `member` | Tasklist creator |
|       `id` | `string` | Indicates the id of member |
|       `type` | `string` | Member type |
|       `role` | `string` | Member role |
|     `owner` | `member` | Tasklist owner |
|       `id` | `string` | Owner member ID |
|       `type` | `string` | Type of member |
|       `role` | `string` | Member role |
|     `members` | `member[]` | Tasklist member |
|       `id` | `string` | Indicates the id of member |
|       `type` | `string` | Type of member |
|       `role` | `string` | member type |
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
                "type": "Owner",
                "role": "Editor"
            },
            "members": [
                {
                    "id": "ou_2cefb2f014f8d0c6c2d2eb7bafb0e54f",
                    "type": "user",
                    "role": "Editor"
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
| 400 | 1470400 | Incorrect request parameters, such as an invalid tasklist GUID provided. | The reason for the error is shown in the returned msg. |
| 404 | 1470404 | The tasklist does not exist or has been deleted. | Verify that the tasklist to be updated exists or has been deleted. |
| 500 | 1470500 | Server error. | Try retrying the call. If it continues to fail, please contact support for feedback. |
| 403 | 1470403 | Missing tasklist edit permission. If the tasklist owner is updated, the manage permission is missing. Please refer to the "How are tasklist authorized?" section in Tasklist Features Overview . | Check whether the calling identity has the permission of the tasklist to perform the update. |
