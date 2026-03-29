---
title: "Approval Task Approve"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uMDNyUjLzQjM14yM0ITN"
method: "POST"
api_path: "https://www.larksuite.com/approval/openapi/v2/instance/approve"
section: "Approval"
scopes:
  - "approval:approval:readonly"
updated: "1658309955000"
---

# Approve an approval task
> This API has been upgraded to enhance its security. We recommend that you migrate to the new version>>

Approves a single approval task. After that, the approval process will be transferred to the next approver.

##  Request
| HTTP URL | https://www.larksuite.com/approval/openapi/v2/instance/approve |
| --- | --- |
| HTTP Method | POST |
| Supported app types | custom,isv |
| Required scopes | `approval:approval:readonly` | ###  Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" Learn more about how to obtain and use access_token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Request body

|Parameter|Type|Required|Description|
|-|-|-|-|
|approval_code|String|Yes|Approval definition code|
|instance_code|String|Yes|Approval instance code|
|open_id|String|No|User open_id. An open_id is required when there is no user_id.|
|user_id|String|Yes|User who performed the action|
|task_id|String|Yes|Task ID|
|comment|String|No|Comments| ### Request body example

```
{
    "approval_code": "7C468A54-8745-2245-9675-08B7C63E7A85",
    "instance_code": "81D31358-93AF-92D6-7425-01A5D67C4E71",
    "user_id": "f7cb567e",
    "task_id": "12345",
    "comment": "OK",
    "open_id": "ou_123456"
}
```

##  Response

###  Response body

|Parameter|Type|Required|Description|
|-|-|-|-|
|code|int|Yes|Error code. A value other than 0 indicates failure.|
|msg|String|Yes|Return code description.| ###  Response body example

```json
{
    "code": 0,
    "msg": "success"
}
```
