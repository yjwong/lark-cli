---
title: "Disband Chat"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uUDN5QjL1QTO04SN0kDN"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/chat/v4/disband"
section: "Deprecated Version (Not Recommended)"
updated: "1626158299000"
---

# Disband Group
A bot disbands the group.

> Bot capability needs to be enabled. Bot must be the group owner or Bot is in the group, it is the the creator of the group and has the permission to update information of groups that are created by an app.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/chat/v4/disband |
| HTTP Method | POST | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value Format**: "Bearer `access_token`" **Example Value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" Learn more about obtaining and using access_token. |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Request body 
|Parameter|Type|Mandatory|Description|Default|Demo|
-- | -- | -- | -- | -- | --
|chat_id|string|Yes| Chat ID.||oc_e03a63d98c0f3b36c329d3546405c490| ### Request body example
```json
{
   "chat_id":"oc_4c24bbde8572c9daedd5e67f6a8ff5e4"
}
```

## Response

### Response body
|Parameters|Type|Description|
|-|-|-|
|code|int|Error code, anything but 0 indicates failure|
|msg|string|Error code description| ### Response body example
```json
{
    "code": 0,
    "msg": "ok"
}
```

### Error code

For details, please refer to: Service-side error codes
