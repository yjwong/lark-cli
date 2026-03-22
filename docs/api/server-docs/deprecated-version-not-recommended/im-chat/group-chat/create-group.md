---
title: "Create Group"
url: "https://open.larksuite.com/document/ukTMukTMukTM/ukDO5QjL5gTO04SO4kDN"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/chat/v4/create/"
section: "Deprecated Version (Not Recommended)"
updated: "1626158296000"
---

# Create Group
The bot creates a group chat and invites specific users to the group chat.

> Bot capability needs to be enabled

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/chat/v4/create/ |
| HTTP Method | POST | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value Format**: "Bearer `access_token`" **Example Value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" Learn more about obtaining and using access_token. |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Request body

|Parameter|Type|Mandatory|Description|Default|Demo|
-- | -- | -- | -- | -- | --
|name|string|Yes|Group name|null|group name|
|description|string|No|Group description|null|group description|
|open_ids|list|No|List of members' open_id (max size is 200)（ Either `open_ids` or `user_ids` should be set with one request. ）|nil|["ou_4065981088f8ef67a504ba8bd6b24d85","ou_111111111111111111111111111111111"]|
|user_ids|list|No|List of members' user_ids (max size is 200)（ Either `open_ids` or `user_ids` should be set with one request. ）|nil|["33417745","cb93bdca"]|
|i18n_names|map|No|Group name in different languages, currently supporting Chinese and English. i18n_names has a higher priority than name. |nil|{"zh_cn": "zh_cn name","en_us": "en_us name"}| 
|only_owner_add|bool|No|Whether only group owners can add group members.|false|false| 
|share_allowed|bool|No|Whether to allow sharing group cards.|true|true|  
|only_owner_at_all|bool|No|Whether only group owners can@all|false|false|  
|only_owner_edit|bool|No|Whether only group owner can edit group info.(Including group name, profile picture, about, annoucement, etc.)|false|false|  
### Request body example
```json
{
    "name": "group name",
    "description": "group description",
    "user_ids": [
        "33417745",
        "cb93bdca"
    ],
    "open_ids": [
        "ou_4065981088f8ef67a504ba8bd6b24d85",
        "ou_111111111111111111111111111111111"
    ],
    "i18n_names": {
        "zh_cn": "zh_cn name",
        "en_us": "en_us name"
    },
    "only_owner_add": false,
    "share_allowed": true,
    "only_owner_at_all": false,
    "only_owner_edit": false
}
```

## Response
### Response body
|Parameters|Type|Description|
|-|-|-|
|code|int|Error code, anything but 0 indicates failure|
|msg|string|Error code description|
|data|-|-
|  ∟chat_id|string|ID of Group chat created|
|  ∟invalid_open_ids|list| List of invalid open_ids|
|  ∟invalid_user_ids|list|  List of invalid user_ids| ### Response body example
```json
{
    "code": 0,
    "data": {
        "chat_id": "oc_4f65b883a624c59414157668c91637ab",
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
