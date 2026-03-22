---
title: "Get Activity Subscription"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/task-v2/tasklist-activity_subscription/get"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/task/v2/tasklists/:tasklist_guid/activity_subscriptions/:activity_subscription_guid"
service: "task-v2"
resource: "tasklist-activity_subscription"
section: "Tasks v2"
rate_limit: "100 per minute"
scopes:
  - "task:tasklist:read"
  - "task:tasklist:write"
field_scopes:
  - "contact:user.employee_id:readonly"
updated: "1699521658000"
---

# Get Activity Subscription

Providing a tasklist GUID and tasklist's subscription GUID, get the details of the subscription data, including name, subscriber, list of event keys that can be notified, etc.

> Obtaining a dynamic subscription requires read permission on the tasklist.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/task/v2/tasklists/:tasklist_guid/activity_subscriptions/:activity_subscription_guid |
| HTTP Method | GET |
| Rate Limit | 100 per minute |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `task:tasklist:read` `task:tasklist:write` |
| Required field scopes | > The response body of the API contains the following sensitive fields, and they will be returned only after corresponding scopes are added. If you do not need the fields, it is not recommended that you request the scopes. `contact:user.employee_id:readonly` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `tasklist_guid` | `string` | tasklist GUID, which can be created byCreate Tasklist, or queried by List Tasklist. **Example value**: "33991879-704f-444f-81d7-55a6aa7be80c" |
| `activity_subscription_guid` | `string` | subscription GUID, which can be created by Create Activity Subscription, or queried by List Activity Subscription. **Example value**: "33991879-704f-444f-81d7-55a6aa7be80c" | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `user_id_type` | `string` | No | User ID categories **Example value**: open_id **Optional values are**:  - `open_id`: Identifies a user to an app. The same user has different Open IDs in different apps. How to get Open ID - `union_id`: Identifies a user to a tenant that acts as a developer. A user has the same Union ID in apps developed by the same developer, and has different Union IDs in apps developed by different developers. A developer can use Union ID to link the same user's identities in multiple apps.How to get Union ID - `user_id`: Identifies a user to a tenant. The same user has different User IDs in different tenants. In one single tenant, a user has the same User ID in all apps （including store apps）. User ID is usually used to communicate user data between different apps. How to get User ID  **Default value**: `open_id` **When the value is `user_id`, the following field scopes are required**: `contact:user.employee_id:readonly` | ## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `activity_subscription` | `tasklist_activity_subscription` | Subscription Details |
|     `guid` | `string` | Subscription GUID |
|     `name` | `string` | subscription name |
|     `subscribers` | `member[]` | subscriber list |
|       `id` | `string` | member id |
|       `type` | `string` | member type |
|       `role` | `string` | member role |
|     `include_keys` | `int[]` | subscribed event keys |
|     `disabled` | `boolean` | whether the subscription is disabled. | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
        "activity_subscription": {
            "guid": "d19e3a2a-edc0-4e4e-b7cc-950e162b53ae",
            "name": "Roadmap subscription",
            "subscribers": [
                {
                    "id": "oc_ 2cefb2f014f8d0c6c2d2eb7bafb0e54f",
                    "type": "chat",
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
| 400 | 1470400 | Invalid request parameters, such as an invalid subscription GUID entered. | Confirm the cause of the error based on the returned msg. |
| 403 | 1470403 | Missing read permission for tasklist. | Confirm that the calling identity has read access to the tasklist. |
| 404 | 1470404 | The tasklist/subscription does not exist or has been deleted. | Confirm that the list exists. |
| 500 | 1470500 | Server error. | Retry invoking the api. If it continues to fail, please contact support for feedback. |
