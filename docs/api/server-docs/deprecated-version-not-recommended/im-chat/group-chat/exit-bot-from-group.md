---
title: "Exit Bot From Group"
url: "https://open.larksuite.com/document/ukTMukTMukTM/ucDO04yN4QjL3gDN"
section: "Deprecated Version (Not Recommended)"
updated: "1626158307000"
---

# Exit bot from group

Exit bot from the group.

**Permission description** ：Bot capability needs to be enabled.

**Request method**: POST  
**Request address**: https://open.larksuite.com/open-apis/bot/v4/remove/  

**Request Parameter Description** 
**Request Header** ：
|Parameter|Type|Mandatory|Description|Default|Demo|
-- | -- | -- | -- | -- | --
Authorization|string|Yes|tenant_access_token is obtained through the Obtain tenant_access_token (marketplace apps)  API or the Obtain tenant_access_token (custom apps) API ;The content must include “Bearer”.||Bearer t-394890fdlkfjdajfljajdkf 
Content-Type|string|Yes|Content-Type||application/json

**Request parameters** :   
|Parameter|Type|Mandatory|Description|Default|Demo|
-- | -- | -- | -- | -- | --
|chat_id|string|Yes| Chat ID. ||oc_4c24bbde8572c9daedd5e67f6a8ff5e4| **Request Body**:  
```json
{
    "chat_id":"oc_4c24bbde8572c9daedd5e67f6a8ff5e4"
}
```

**Description of response parameters** : 
|Parameters|Type|Description|
|-|-|-|
|code|int|Error code, anything but 0 indicates failure|
|msg|string|Error code description| **Response Body**:  
```json
{
    "code": 0,
    "msg": "ok"
}
```
