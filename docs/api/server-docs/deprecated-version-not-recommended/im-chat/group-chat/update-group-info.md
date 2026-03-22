---
title: "Update Group Info"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uYTO5QjL2kTO04iN5kDN"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/chat/v4/update/"
section: "Deprecated Version (Not Recommended)"
updated: "1626158298000"
---

# Update Group Info 
Update group name, group settings, change group owner, etc.

> Bot capability needs to be enabled. Bot must be the owner of the group  or Bot is  in the group, it is the the creator and has the permission to update information of groups that are created by an app.

> - If the bot is the group owner or the group creator with update information of groups that are created by an app scope, the bot can edit any group info.
> - If the bot is not the above character, and the "Only group owner can edit group info" is enabled，the bot can not eit any group info.
> - If the bot is not the above character, and the "Only group owner can edit group info" is disabled，the bot can edit "group name" and "group description".

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/chat/v4/update/ |
| HTTP Method | POST | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value Format**: "Bearer `access_token`" **Example Value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" Learn more about obtaining and using access_token. |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Request body 

|Parameter|Type|Mandatory|Description|Default|Demo|
-- | -- | -- | -- | -- | --
|chat_id| string| Yes | Chat ID.||oc_020ff1d91a0295fc3961032768d41f39|
|owner_open_id| string| No | Owner's open_id （When changing group owner, fill in either owner_open_id or owner_user_id）||ou_4065981088f8ef67a504ba8bd6b24d85|
|owner_user_id| string| No | Owner's user_id （When changing group owner, fill in either owner_open_id or owner_user_id）||cb93bdca|
|name | string| No | Default group name||group name|
|i18n_names | map| No | International group name||{"zh_cn":"zh_cn name","en_us":"en_us name","ja_jp":"ja_jp name"}|
|only_owner_add|bool|No|Whether only group owners can add group members. (Bot must be group owner to modify chat info: only_owner_add)|false|false| 
|share_allowed|bool|No|Whether to allow sharing group cards.(Bot must be group owner to modify chat info: share_allowed)|true|true|  
|only_owner_at_all|bool|No|Whether only group owners can@all(Bot must be group owner to modify chat info: only_owner_at_all)|false|false|  
|only_owner_edit|bool|No|Whether only group owner can edit group info.(Including group name, profile picture, about, annoucement, etc.)(Bot must be group owner to modify chat info: only_owner_edit)|false|false| ### Request body example
```json
{
    "chat_id": "oc_020ff1d91a0295fc3961032768d41f39",
    "owner_user_id":"cb93bdca",
    "owner_open_id":"ou_4065981088f8ef67a504ba8bd6b24d85",
    "name":"group name",
    "i18n_names":{
		"zh_cn":"zh_cn name",
		"en_us":"en_us name",
		"ja_jp":"ja_jp name"
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
code |int|  Error code, anything but 0 indicates failure
msg |string| Error code description
data|-|-
  ∟chat_id |string| chat_id (same as the chat_id in the request)

### Response body example
```json
{
    "code": 0,
    "data": {
        "chat_id": "oc_020ff1d91a0295fc3961032768d41f39"
    },
    "msg": "ok"
}
```

### Error code

For details, please refer to: Service-side error codes
