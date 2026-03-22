---
title: "Delete Users From Chat"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uADMwUjLwADM14CMwATN"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/chat/v4/chatter/delete/"
section: "Deprecated Version (Not Recommended)"
updated: "1626158301000"
---

# Delete Users From Group
A bot deletes users from the group.

> Bot capability needs to be enabled. Bot must be the group owner or Bot is in the group, it is creator of the group and has the permission to update information of groups that are created by an app.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/chat/v4/chatter/delete/ |
| HTTP Method | POST | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value Format**: "Bearer `access_token`" **Example Value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" Learn more about obtaining and using access_token. |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Request body
|Parameter|Type|Mandatory|Description|Default|Demo|
-- | -- | -- | -- | -- | --
|chat_id|string|Yes| Chat ID.||oc_e03a63d98c0f3b36c329d3546405c490|
|user_ids|list|No|List of members' user_ids(max size is 200)( Either `open_ids` or `user_ids` should be set with one request. )|nil|["33417745","cb93bdca"]|
|open_ids|list|No|List of members' open_id(max size is 200)( Either `open_ids` or `user_ids` should be set with one request. )|nil|["ou_4065981088f8ef67a504ba8bd6b24d85","ou_111111111111111111111111111111111"]| ### Request body example
```json
{
    "chat_id": "oc_e03a63d98c0f3b36c329d3546405c490",
    "user_ids": [
        "33417745",
        "cb93bdca"
    ],
    "open_ids": [
        "ou_4065981088f8ef67a504ba8bd6b24d85",
        "ou_111111111111111111111111111111111"
    ]
}
```

## Response

### Response body
|Parameters|Type|Description|
|-|-|-|
|code|int|Error code, anything but 0 indicates failure|
|msg|string|Error code description|
|data|-|-
|  ∟invalid_open_ids|list| List of invalid open_ids|
|  ∟invalid_user_ids|list|  List of invalid user_ids| ### Response body example
```json
{
    "code": 0,
    "data": {
        "invalid_open_ids": [
            "ou_111111111111111111111111111111111"
        ],
        "invalid_user_ids": [
            "33417745"
        ]
    },
    "msg": "ok"
}
```

### Error code

For details, please refer to: Service-side error codes
