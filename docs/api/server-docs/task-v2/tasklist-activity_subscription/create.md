---
title: "Create Activity Subscription"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/task-v2/tasklist-activity_subscription/create"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/task/v2/tasklists/:tasklist_guid/activity_subscriptions"
service: "task-v2"
resource: "tasklist-activity_subscription"
section: "Tasks v2"
rate_limit: "100 per minute"
scopes:
  - "task:tasklist:write"
field_scopes:
  - "contact:user.employee_id:readonly"
updated: "1699521654000"
---

# Create Activity Subscription

Create an activity subscription for a tasklist. Each subscription can contain multiple subscribers (currently only group chats are supported). 

After the subscription's creation, if the corresponding event occurs in the tasklist, a notification message will be sent to the subscribers defined in the subscription. Each tasklist can create up to 50 subscriptions. Each subscription supports up to 50 subscribers.

A subscription can set which kinds of events (event_key) should be notified by setting `include_keys`. If the `include_keys` are empty, no events are notified. 

When necessary, you can create supscription by which no notification messages will be sent by setting `disabled` to "true".

> Creating a subscription for the task requires the calling identity to have the edit permission of the tasklist and can access the chats in the subscribers. If the calling identity loses the tasklist edit permissions or accessibility of the chats after adding the subscription, the subscription will remain in place until deletion.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/task/v2/tasklists/:tasklist_guid/activity_subscriptions |
| HTTP Method | POST |
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
| `tasklist_guid` | `string` | tasklist GUID **Example value**: "d19e3a2a-edc0-4e4e-b7cc-950e162b53ae" | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `user_id_type` | `string` | No | User ID categories **Example value**: open_id **Optional values are**:  - `open_id`: Identifies a user to an app. The same user has different Open IDs in different apps. How to get Open ID - `union_id`: Identifies a user to a tenant that acts as a developer. A user has the same Union ID in apps developed by the same developer, and has different Union IDs in apps developed by different developers. A developer can use Union ID to link the same user's identities in multiple apps.How to get Union ID - `user_id`: Identifies a user to a tenant. The same user has different User IDs in different tenants. In one single tenant, a user has the same User ID in all apps （including store apps）. User ID is usually used to communicate user data between different apps. How to get User ID  **Default value**: `open_id` **When the value is `user_id`, the following field scopes are required**: `contact:user.employee_id:readonly` | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `name` | `string` | Yes | Subscription name, cannot be empty, up to 50 characters. **Example value**: "My subscription" |
| `subscribers` | `member[]` | Yes | Subscribers list **Data validation rules**: - Maximum length: `50` |
|   `id` | `string` | No | Indicates the id of member **Example value**: "oc_ 2cefb2f014f8d0c6c2d2eb7bafb0e54f" **Data validation rules**: - Maximum length: `100` characters |
|   `type` | `string` | No | Type of member (currently only chat is supported) **Example value**: "chat" **Default value**: `user` |
| `include_keys` | `int[]` | Yes | A list of subscribed event keys. Each event key is represented by an integer. Currently  following event keys are supported: - 100: Adde a task to a tasklist - 101: Remove a task from a tasklist - 103: Complete a task in the tasklist - 104: Uncomplte task in the tasklist - 109: Assignee of a task in the tasklist is added - 110: Assignee of a task in the tasklist  is updated - 111: Assignee of a task in the tasklist is removed - 119: Attachments are added to a task in the tasklist - 121: Create a comment of a task in the tasklist - 122: Reply a comment of a task in the tasklist - 129: Set start time of task of a task in the tasklist - 130: Set due time of task of a task in the tasklist - 131: Set both start and due time of a task in the tasklist - 132: Unset both start and due time of a task in the tasklist Empty event keys list is allowed which means no event's notification will be sent. Duplication is not allowed. **Example value**: [100] |
| `disabled` | `boolean` | No | whether the subscription is disabled **Example value**: false | ### Request body example

{
    "name": "My subscription",
    "subscribers": [
        {
            "id": "oc_ 2cefb2f014f8d0c6c2d2eb7bafb0e54f",
            "type": "chat"
        }
    ],
    "include_keys": [
        100
    ],
    "disabled": false
}

## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `activity_subscription` | `tasklist_activity_subscription` | tasklist subscription |
|     `guid` | `string` | subscription GUID |
|     `name` | `string` | subscription name |
|     `subscribers` | `member[]` | subscriber |
|       `id` | `string` | member id |
|       `type` | `string` | member type |
|       `role` | `string` | member role |
|     `include_keys` | `int[]` | event keys to send notifications |
|     `disabled` | `boolean` | whether the subscription is disabled. | ### Response body example

{
    "code": 0,
    "msg": "Success",
    "data": {
        "activity_subscription": {
            "guid": "d19e3a2a-edc0-4e4e-b7cc-950e162b53ae",
            "name": "Roadmap subscription",
            "subscribers": [
                {
                    "id": "oc_2cefb2f014f8d0c6c2d2eb7bafb0e54f",
                    "type": "user",
                    "role": "editor"
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
| 400 | 1470400 | Incorrect request parameters, such as setting an unsupported event key. | The reason for the error is shown in the returned msg. |
| 403 | 1470403 | Missing edit permission to the tasklist or accessibility to the chats set in the subscribers. | Confirm that the calling identity has the correct permissions. |
| 404 | 1470404 | The tasklist to add the subscription does not exist or has been deleted. | Determine if the tasklist to add the subscription exists. |
| 500 | 1470500 | Server error. | Try retrying the call. If it continues to fail, please contact support for feedback. |
