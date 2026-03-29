---
title: "Resubmit the task for approval"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/approval-v4/task/resubmit"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/approval/v4/tasks/resubmit"
service: "approval-v4"
resource: "task"
section: "Approval"
scopes:
  - "approval:approval"
  - "approval:approval:readonly"
field_scopes:
  - "contact:user.employee_id:readonly"
updated: "1675167408000"
---

# Resubmit the task for approval

Resume a single approval task that falls back to the sponsor. After initiation, the approval process flows to the next approver.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/approval/v4/tasks/resubmit |
| HTTP Method | POST |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `approval:approval` `approval:approval:readonly` |
| Required field scopes | > The response body of the API contains the following sensitive fields, and they will be returned only after corresponding scopes are added. If you do not need the fields, it is not recommended that you request the scopes. `contact:user.employee_id:readonly` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `user_id_type` | `string` | No | User ID categories **Example value**: "open_id" **Optional values are**:  - `open_id`: Identifies a user to an app. The same user has different Open IDs in different apps. How to get Open ID - `union_id`: Identifies a user to a tenant that acts as a developer. A user has the same Union ID in apps developed by the same developer, and has different Union IDs in apps developed by different developers. A developer can use Union ID to link the same user's identities in multiple apps.How to get Union ID - `user_id`: Identifies a user to a tenant. The same user has different User IDs in different tenants. In one single tenant, a user has the same User ID in all apps （including store apps）. User ID is usually used to communicate user data between different apps. How to get User ID  **Default value**: `open_id` **When the value is `user_id`, the following field scopes are required**: `contact:user.employee_id:readonly` | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `approval_code` | `string` | Yes | Approval Definition Code **Example value**: "7C468A54-8745-2245-9675-08B7C63E7A85" |
| `instance_code` | `string` | Yes | Approval Example Code **Example value**: "81D31358-93AF-92D6-7425-01A5D67C4E71" |
| `user_id` | `string` | Yes | Fill in the operation user id according to the user_id_type **Example value**: "f7cb567e" |
| `comment` | `string` | No | Opinion **Example value**: "OK" |
| `task_id` | `string` | Yes | Task ID, approval instance details task_list id **Example value**: "12345" |
| `form` | `string` | Yes | JSON array, control value, same as the form field in the interface of creating an approval instance **Example value**: "[{\"id\":\"user_name\", \"type\": \"input\", \"value\":\"test\"}]" | ### Request body example

{
    "approval_code": "7C468A54-8745-2245-9675-08B7C63E7A85",
    "instance_code": "81D31358-93AF-92D6-7425-01A5D67C4E71",
    "user_id": "f7cb567e",
    "comment": "OK",
    "task_id": "12345",
    "form": "[{\"id\":\"user_name\", \"type\": \"input\", \"value\":\"test\"}]"
}

## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {

    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 1395001 | There have been some errors. Please try again later | There was an error with the service, please try again later |
| 400 | 1390018 | not support handwritten signature | Handwritten signatures are not supported |
| 400 | 1390010 | task not found | Mission not found |
| 400 | 1390001 | param is invalid | Parameter error |
| 400 | 1390002 | approval code not found | Check that the review definition code is correct |
| 400 | 1390003 | instance code not found | Check that the review instance code is correct |
