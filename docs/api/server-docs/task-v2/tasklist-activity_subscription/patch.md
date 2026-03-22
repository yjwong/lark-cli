---
title: "Patch Activity Subscription"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/task-v2/tasklist-activity_subscription/patch"
method: "PATCH"
api_path: "https://open.larksuite.com/open-apis/task/v2/tasklists/:tasklist_guid/activity_subscriptions/:activity_subscription_guid"
service: "task-v2"
resource: "tasklist-activity_subscription"
section: "Tasks v2"
rate_limit: "100 per minute"
scopes:
  - "task:tasklist:write"
field_scopes:
  - "contact:user.employee_id:readonly"
updated: "1699521666000"
---

# Patch Activity Subscription

Update an activity subscription by providing a tasklist GUID and a subscription GUID. When updating, fill in all field names to be updated in the `update_fields` field, and fill in the new value of the field in the `activity_subscription` field.

`update_fields` support following fields:
* `name`: the name of the subscription
* `subscribers`: Subscriber list. when updated, the old subscribers list will be completely replaced with the new one. Each subscription supports up to 50 subscribers. And only chat subscribers are allowed.
* `include_keys`: subscribed event keys. when updated, the old list will be completely replaced with the new one.  Only supported event keys are allowed (see the field description below). Duplication is not allowed.
* `disabled`: whether the subscription is disabled or not.

> To update the subscription, the calling identity needs to have edit permission to the tasklist.
> 
> If the `subscribers` is updated, the calling identity must can access the chats set in the new subscribers.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/task/v2/tasklists/:tasklist_guid/activity_subscriptions/:activity_subscription_guid |
| HTTP Method | PATCH |
| Rate Limit | 100 per minute |
| Supported app types | custom,isv |
| Required scopes | `task:tasklist:write` |
| Required field scopes | > The response body of the API contains the following sensitive fields, and they will be returned only after corresponding scopes are added. If you do not need the fields, it is not recommended that you request the scopes. `contact:user.employee_id:readonly` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `tasklist_guid` | `string` | tasklist GUID, which can be created byCreate Tasklist, or queried by List Tasklist. **Example value**: "33991879-704f-444f-81d7-55a6aa7be80c" |
| `activity_subscription_guid` | `string` | subscription GUID, which can be created by Create Activity Subscription, or queried by List Activity Subscription. **Example value**: "f5ca6747-5ac3-422e-a97e-972c1b2c24f3" | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `user_id_type` | `string` | No | User ID categories **Example value**: open_id **Optional values are**:  - `open_id`: Identifies a user to an app. The same user has different Open IDs in different apps. How to get Open ID - `union_id`: Identifies a user to a tenant that acts as a developer. A user has the same Union ID in apps developed by the same developer, and has different Union IDs in apps developed by different developers. A developer can use Union ID to link the same user's identities in multiple apps.How to get Union ID - `user_id`: Identifies a user to a tenant. The same user has different User IDs in different tenants. In one single tenant, a user has the same User ID in all apps （including store apps）. User ID is usually used to communicate user data between different apps. How to get User ID  **Default value**: `open_id` **When the value is `user_id`, the following field scopes are required**: `contact:user.employee_id:readonly` | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `activity_subscription` | `tasklist_activity_subscription` | Yes | subscription data to update |
|   `name` | `string` | No | subscription name, when updated, it should be 1~50 characters. **Example value**: "roadmap subscription" |
|   `subscribers` | `member[]` | No | subscribers list, when updated, up to 50 subscribers are supported. **Data validation rules**: - Maximum length: `50` |
|     `id` | `string` | No | member_id **Example value**: "oc_ 2cefb2f014f8d0c6c2d2eb7bafb0e54f" **Data validation rules**: - Maximum length: `100` characters |
|     `type` | `string` | No | member type. only support 'chat'. **Example value**: "chat" **Default value**: `user` |
|   `include_keys` | `int[]` | No | subscribed event keys. Each event key is represented by an integer. Currently  following event keys are supported: - 100: Adde a task to a tasklist - 101: Remove a task from a tasklist - 103: Complete a task in the tasklist - 104: Uncomplte task in the tasklist - 109: Assignee of a task in the tasklist is added - 110: Assignee of a task in the tasklist  is updated - 111: Assignee of a task in the tasklist is removed - 119: Attachments are added to a task in the tasklist - 121: Create a comment of a task in the tasklist - 122: Reply a comment of a task in the tasklist - 129: Set start time of task of a task in the tasklist - 130: Set due time of task of a task in the tasklist - 131: Set both start and due time of a task in the tasklist - 132: Unset both start and due time of a task in the tasklist Empty event keys list is allowed which means no event's notification will be sent. Duplication is not allowed. **Example value**: [101] |
|   `disabled` | `boolean` | No | whether the subscription is disabled. **Example value**: false |
| `update_fields` | `string[]` | Yes | list of fields to update **Example value**: ["name"] **Optional values are**:  - `name`: subscription name - `include_keys`: subscribed event keys - `subscribers`: susbscribers list - `disabled`: whether the subscription is disabled  **Data validation rules**: - Length range: `1` ～ `20` | ### Request body example

{
    "activity_subscription": {
        "name": "roadmap subscription",
        "subscribers": [
            {
                "id": "oc_ 2cefb2f014f8d0c6c2d2eb7bafb0e54f",
                "type": "chat"
            }
        ],
        "include_keys": [
            101
        ],
        "disabled": false
    },
    "update_fields": [
        "name"
    ]
}

## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `activity_subscription` | `tasklist_activity_subscription` | Updated subscription |
|     `guid` | `string` | subscription guid |
|     `name` | `string` | subscription name |
|     `subscribers` | `member[]` | subscribers list |
|       `id` | `string` | member id |
|       `type` | `string` | member_type |
|       `role` | `string` | member role |
|     `include_keys` | `int[]` | subscribed event keys |
|     `disabled` | `boolean` | whether the subscription is updated. | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
        "activity_subscription": {
            "guid": "d19e3a2a-edc0-4e4e-b7cc-950e162b53ae",
            "name": "roadmap subscription",
            "subscribers": [
                {
                    "id": "oc_ 2cefb2f014f8d0c6c2d2eb7bafb0e54f",
                    "type": "chat",
                    "role": "Editor"
                }
            ],
            "include_keys": [
                101
            ],
            "disabled": false
        }
    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 1470400 | Incorrect request parameters. | Determine the specific error cause based on the returned msg. |
| 403 | 1470403 | Lack of edit permission on the tasklist. | Confirm that the calling identity has edit permission on the tasklist. |
| 404 | 1470404 | The listing/subscription does not exist or has been deleted. | Confirm that the listing/subscription does not exist or has been deleted. |
| 500 | 1470500 | Server error. | Try re-invoking the api. If it continues to fail, please contact support for feedback. |
