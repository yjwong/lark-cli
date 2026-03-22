---
title: "Approval task return"
url: "https://open.larksuite.com/document/ukTMukTMukTM/ukTM5UjL5ETO14SOxkTN/approval-task-return"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/approval/v4/instances/specified_rollback"
section: "Approval"
scopes:
  - "approval:approval:readonly"
updated: "1658309938000"
---

# Roll back an approval task

Rolls back the current approval to one or more approved task nodes. After the rollback, the approved nodes will generate approval tasks again.

##  Request
| HTTP URL | https://open.larksuite.com/open-apis/approval/v4/instances/specified_rollback |
| --- | --- |
| HTTP Method | POST |
| Supported app types | custom,isv |
| Required scopes | `approval:approval:readonly` | ###  Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes  `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" Learn more about how to obtain and use access_token |  |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `user_id_type` | `string` | No | User  ID  type **Example value**: "open_id" **Optional values are**: - `open_id`: User's  open id - `union_id`: User's  union id - `user_id`: User's  user id **Default value**: `open_id` **When the value is `user_id`, the required field scopes are**: `contact:user.employee_id:readonly` | ### Request body
|Parameter|Type|Required|Description|
|-|-|-|-|
|task_id|string|Yes|ID of current approval task, obtained from the task_list in the instance details. It must be the ID of a task with the status of PENDING.|
|user_id|string|Yes|The open_id of the approver in the current approval task, obtained from the task_list in the instance details. The status of the corresponding task must be PENDING.|
|reason|string|Yes|Reason for rollback|
|task_def_key_list|list|Yes|node_key of the specified rollback task, obtained from timeline in the instance details. The status of the corresponding task must be PASS.| #### Description:

-   To roll back to the initiator, input START in task_def_key_list.

### Request body example

```json
{
    "task_id":"7023757604987891234",
    "user_id":"ou_123",
    "reason":"The initiator resubmits the task so that approvers must approve the request again",
    "task_def_key_list":[
        "START",
        "APPROVAL_141532_3632523"
    ]
}
```

##  Response

###  Response body

|Parameter|Type|Required|Description|
|-|-|-|-|
|code|int|Yes|Error code. A value other than 0 indicates failure.|
|msg|string|Yes|Return code description.| ###  Response body example

```json
{
    "code": 0,
    "data": {},
    "msg": "success"
}  
```

###
