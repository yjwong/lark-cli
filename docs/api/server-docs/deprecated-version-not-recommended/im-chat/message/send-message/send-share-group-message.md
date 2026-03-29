---
title: "Send Share Group Message"
url: "https://open.larksuite.com/document/ukTMukTMukTM/ucjMxEjL3ITMx4yNyETM"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/message/v4/send/"
section: "Deprecated Version (Not Recommended)"
updated: "1626158291000"
---

# Send a share group message
Sends a share group message to a specified user or chat, including private chats and group chats.

> Bot capability needs to be enabled. This group is allowed to be shared. For private chats, bot need to be visible to the user. For group chats, bot must be in the group.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/message/v4/send/ |
| HTTP Method | POST |
| Required scopes | Send messages as an app(Historic Version) | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value Format**: "Bearer `access_token`" **Example Value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" Learn more about obtaining and using access_token. |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Request body
|Parameter|Type|Mandatory|Description|Default|Demo|
-- | -- | -- | -- | -- | --
open_id user_id  email  chat_id  | string | No | To send a private chat message to a user, use either open_id, email, or user_id. To send a message to a group, use the chat_id for the group (obtain via the Obtains Group List API). The server reads the string in the following order: chat_id > open_id > user_id > email ||ou_5ad573a6411d72b8305fda3a9c15c70e|
root_id | string | No | To reply to a particular message, add the corresponding message ID ||om_40eb06e7b84dc71c03e009ad3c754195|
msg_type | string | Yes | Message type, must be “share_chat” ||share_chat|
content | string | Yes | Message content ||| ∟share_chat_id | string | Yes | group's chat_id ||oc_f5b1a7eb27ae2c7b6adc2a74faf339ff| ### Request body example

```json
{
   "open_id":"ou_5ad573a6411d72b8305fda3a9c15c70e", 
   "root_id":"om_40eb06e7b84dc71c03e009ad3c754195",
   "chat_id":"oc_5ad11d72b830411d72b836c20", 
   "user_id": "92e39a99",
    "msg_type": "share_chat",
    "content":{
        "share_chat_id": "oc_f5b1a7eb27ae2c7b6adc2a74faf339ff"
    }
}
```

## Response
### Response body
|Parameters|Type|Description|
|-|-|-|
|code|int|Error code, anything but 0 indicates failure|
|msg|string|Error code description|
data | - | -
  ∟message_id |string| Message ID

### Response body example
```json
{
    "code": 0,
    "msg": "ok",
    "data":{
       "message_id": "om_92eb70a7120ga8c3ca56a12ee9ba7ca2"
    }
}
```
### Error code

For details, please refer to: Error Codes
