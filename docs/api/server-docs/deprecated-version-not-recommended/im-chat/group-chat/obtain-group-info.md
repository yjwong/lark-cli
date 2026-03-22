---
title: "Obtain Group Info"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uMTO5QjLzkTO04yM5kDN"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/chat/v4/info?chat_id=oc_e865ccc90fe736d4fa8db9b9a2ca986e"
section: "Deprecated Version (Not Recommended)"
updated: "1626158297000"
---

# Obtain Group Info
Obtain group basic information such as group name, owner's ID, members' ID, etc.

> Bot capability needs to be enabled. Bot must be in the chat.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/chat/v4/info?chat_id=oc_e865ccc90fe736d4fa8db9b9a2ca986e |
| HTTP Method | GET | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value Format**: "Bearer `access_token`" **Example Value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" Learn more about obtaining and using access_token. |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Query parameters
|Parameter|Type|Mandatory|Description|Default|Demo|
-- | -- | -- | -- | -- | --
chat_id| string| Yes | Chat ID. | |oc_eb9e82d5657777ebf1bb5b9024f549ef| ## Response
### Response body
|Parameters|Type|Description|
|-|-|-|
code |int| Error code, anything but 0 indicates failure
msg  |string| Error code description
data | - | - | ∟i18n_names |map| Group international name. This field is only available if an international name is set.
  ∟avatar|string| avatar
  ∟chat_id|string| Group ID
  ∟description|string| Group description
  ∟members|list| List of members
    ∟open_id|string| member's open_id
    ∟user_id|string| member's user_id
  ∟name |string| Group name, valid when type is "group"
  ∟type |string| Group type. "group" means group chat, "p2p" means one-to-one private chat
  ∟owner_user_id |string| owner's user_id（If a bot is the group owner, this field does not exist.）
  ∟owner_open_id |string| owner's open_id （If a bot is the group owner, this field does not exist.）

### Response body example
```json
{
    "code": 0,
    "data": {
        "chat_id": "oc_fb4fd349d4xxxxxxxxxxxxxxxxxx",
        "type": "group"
        "name": "group name",
        "i18n_names": {
            "en_us": "en_us name",
            "ja_jp": "ja_jp name",
            "zh_cn": "zh_cn name"
        },
        "description": "group description",
        "avatar": "https://p3-lark-file.byteimg.com/img/lark.avatar/default-avatar_xxxxxxxxxxxxxx.jpg",
        "members": [
            {
                "open_id": "ou_xxxxxxx",
                "user_id": "8e17d887"
            }
        ],
        "only_owner_add": false,
        "only_owner_at_all": false,
        "only_owner_edit": false,
        "share_allowed": true,
        "add_member_verify": false,
        "send_message_permission": "all",
        "join_message_visibility": "all",
        "leave_message_visibility": "owner",
        "group_email_enabled": false,
        "send_group_email_permission": "tenant_member",
        "owner_open_id": "ou_6edde2deccabb76c12b30f0345f19aa1",
        "owner_user_id": "8e17d887"
    },
    "msg": "ok"
}
```
### Error code

For details, please refer to: Service-side error codes
