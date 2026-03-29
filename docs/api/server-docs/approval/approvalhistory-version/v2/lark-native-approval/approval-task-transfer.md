---
title: "Approval Task Transfer"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uUDNyUjL1QjM14SN0ITN"
method: "POST"
api_path: "https://www.larksuite.com/approval/openapi/v2/instance/transfer"
section: "Approval"
scopes:
  - "approval:approval:readonly"
updated: "1658309960000"
---

# Transfer an approval task
> This API has been upgraded to enhance its security. We recommend that you migrate to the new version>>

Transfers a single approval task. After that, the approval process will be transferred to the transfer recipient.

##  Request
| HTTP URL | https://www.larksuite.com/approval/openapi/v2/instance/transfer |
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
|user_id|String|Yes|User who performed the action|
|task_id|String|Yes|Task ID|
|comment|String|No|Comments|
|transfer_user_id|String|Yes|Unique ID of the transfer recipient|
|open_id|String|No|User open_idAn open_id is required when there is no user_id.|
|transfer_open_id|String|No|Transfer recipient open_id A transfer_open_id is required when there is no transfer_user_id.|
### Request body example

```json
{
    "approval_code": "7C468A54-8745-2245-9675-08B7C63E7A85",
    "instance_code": "81D31358-93AF-92D6-7425-01A5D67C4E71",
    "user_id": "f7cb567e",
    "task_id": "12345",
    "comment": "check",
    "transfer_user_id": "f4ip317q",
    "open_id": "ou_123456",
    "transfer_open_id": "ou_654321"
}
```

##  Response

###  Response body

|Parameter|Type|Required|Description|
|-|-|-|-|
|code|int|Yes|Error code. A value other than 0 indicates failure.|
|msg|String|Yes|Return code description| ###  Response body example

```json
{
    "code": 0,
    "msg": "success"
}
```
