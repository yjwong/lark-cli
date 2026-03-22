---
title: "Return approval task"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/approval-v4/instance/specified_rollback"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/approval/v4/instances/specified_rollback"
service: "approval-v4"
resource: "instance"
section: "Approval"
scopes:
  - "approval:approval:readonly"
field_scopes:
  - "contact:user.employee_id:readonly"
updated: "1675167405000"
---

# Return approval task

Rolls back the current approval to one or more approved task nodes. After the rollback, the approved nodes will generate approval tasks again.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/approval/v4/instances/specified_rollback |
| HTTP Method | POST |
| Supported app types | custom,isv |
| Required scopes | `approval:approval:readonly` |
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
| `user_id` | `string` | Yes | The open_id of the approver in the current approval task, obtained from the task_list in the instance details. The status of the corresponding task must be PENDING. **Example value**: "893g4c45" |
| `task_id` | `string` | Yes | ID of current approval task, obtained from the task_list in the instance details. It must be the ID of a task with the status of PENDING. **Example value**: "7026591166355210260" |
| `reason` | `string` | No | Reason for rollback **Example value**: "reason" |
| `extra` | `string` | No | extra info **Example value**: "{}" |
| `task_def_key_list` | `string[]` | Yes | node_key of the specified rollback task, obtained from timeline in the instance details. The status of the corresponding task must be PASS. **Example value**: ["START","APPROVAL_27997_285502","APPROVAL_462205_2734554"] **Data validation rules**: - Length range: `1` ～ `100` | ### Request body example

{
    "user_id": "893g4c45",
    "task_id": "7026591166355210260",
    "reason": "reason",
    "extra": "{}",
    "task_def_key_list": [
        "START","APPROVAL_27997_285502","APPROVAL_462205_2734554"
    ]
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
| 400 | 1390001 | param is invalid | Parameter error |
