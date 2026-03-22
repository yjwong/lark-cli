---
title: "Delete Activity Subscription"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/task-v2/tasklist-activity_subscription/delete"
method: "DELETE"
api_path: "https://open.larksuite.com/open-apis/task/v2/tasklists/:tasklist_guid/activity_subscriptions/:activity_subscription_guid"
service: "task-v2"
resource: "tasklist-activity_subscription"
section: "Tasks v2"
rate_limit: "100 per minute"
scopes:
  - "task:tasklist:write"
updated: "1699521670000"
---

# Delete Activity Subscription

Delete an activity subscription by giving a tasklist GUID and its subscription GUID. The deleted subscription is not recoverable.

> Deleting a subscription requires edit permission to the tasklist.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/task/v2/tasklists/:tasklist_guid/activity_subscriptions/:activity_subscription_guid |
| HTTP Method | DELETE |
| Rate Limit | 100 per minute |
| Supported app types | custom,isv |
| Required scopes | `task:tasklist:write` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `tasklist_guid` | `string` | tasklist GUID, which can be created byCreate Tasklist, or queried by List Tasklist. **Example value**: "f5ca6747-5ac3-422e-a97e-972c1b2c24f3" |
| `activity_subscription_guid` | `string` | subscription GUID, which can be created by Create Activity Subscription, or queried by List Activity Subscription. **Example value**: "d19e3a2a-edc0-4e4e-b7cc-950e162b53ae" | ## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {}
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 1470400 | Incorrect request parameters. | Confirm the specific cause of the error based on the msg in the return body. |
| 403 | 1470403 | Missing editable permissions to the tasklist. | Confirm that the calling identity has edit permission to the tasklist. |
| 404 | 1470404 | The tasklist/subscription does not exist or has been deleted. | Confirm that the subscription to delete exists. |
| 500 | 1470500 | Server error. | Try re-invoking the api. If it continues to fail, please contact support for feedback. |
