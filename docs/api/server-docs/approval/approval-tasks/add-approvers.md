---
title: "Add approvers"
url: "https://open.larksuite.com/document/ukTMukTMukTM/ukTM5UjL5ETO14SOxkTN/approval-task-addsign"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/approval/v4/instances/add_sign"
section: "Approval"
scopes:
  - "approval:approval:readonly"
updated: "1675167408000"
---

# Add approvers

Adds an approver to a single approval task.

##  Request
| HTTP URL | https://open.larksuite.com/open-apis/approval/v4/instances/add_sign |
| --- | --- |
| HTTP Method | POST |
| Supported app types | custom,isv |
| Required scopes | `approval:approval:readonly` | ###  Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes  `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" Learn more about how to obtain and use access_token |  |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Request body
|Parameter|Type|Required|Description|
|-|-|-|-|
|approval_code|String|Yes|Approval definition code|
|instance_code|String|Yes|Approval instance code|
|user_id|String|Yes|User who performed the action|
|task_id|String|Yes|Task ID|
|comment|String|No|Comments|
|add_sign_user_ids|List|Yes|ID of added approver|
|add_sign_type|String|Yes|1: Pre-approver, 2: Post-approver, 3: Co-approver|
|approval_method|String|No|Required for pre-approvers and post-approvers, 1: Anyone assigned, 2: Everyone assigned| ### Request body example

```json
{
    "approval_code": "3B68E280-CF10-4198-B4CD-2E3BB97981D8",
    "instance_code": "289330DE-FBF1-4A47-91F9-9EFCCF11BCAE",
    "user_id": "b16g66e3",
    "task_id": "6955096766400167956",
    "comment": "addSignComment",
    "add_sign_user_ids": ["d19b913b","3313g62b"],
    "add_sign_type": 1,
    "approval_method": 1
}
```

##  Response

###  Response body
|Parameter|Type|Required|Description|
|-|-|-|-|
|code|int|Yes|Error code. A value other than 0 indicates failure.|
|msg|String|Yes|Return code description.|
###  Response body example

```json
{
    "code": 0,
    "msg": "success"
}
```
