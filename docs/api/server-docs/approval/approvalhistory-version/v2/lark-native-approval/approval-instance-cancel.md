---
title: "Approval Instance Cancel"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uYDNyUjL2QjM14iN0ITN"
method: "POST"
api_path: "https://www.larksuite.com/approval/openapi/v2/instance/cancel"
section: "Approval"
scopes:
  - "approval:approval:readonly"
updated: "1658309952000"
---

# Cancel an approval instance
> This API has been upgraded to enhance its security. We recommend that you migrate to the new version>>

Revokes a single approval instance in the "Under review" status. The approval process ends after recall.
##  Request
| HTTP URL | https://www.larksuite.com/approval/openapi/v2/instance/cancel |
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
|open_id|String|Yes|Unique identifier for a user in an app. A Lark user obtained based on a userID, openID, or TenantId.| ### Request body example

```json
{
    "approval_code": "7C468A54-8745-2245-9675-08B7C63E7A85",
    "instance_code": "81D31358-93AF-92D6-7425-01A5D67C4E71",
    "open_id": "ou_123456",
    "user_id": "f7cb567e"
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
