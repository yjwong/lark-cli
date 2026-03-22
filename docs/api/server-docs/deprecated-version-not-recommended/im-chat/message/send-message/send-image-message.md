---
title: "Send Image Message"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uIDMxEjLyATMx4iMwETM"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/message/v4/send/"
section: "Deprecated Version (Not Recommended)"
updated: "1626158290000"
---

# Sends image message
Sends a image message to a specified user or chat, including private chats and group chats.

> Bot capability needs to be enabled. For private chats, bot need to be visible to the user. For group chats, bot must be in the group.

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
open_id user_id  email  chat_id  | string | Yes | To send a private chat message to a user, use either open_id, email, or user_id. To send a message to a group, use the chat_id for the group (obtain via the Obtains Group List API). The server reads the string in the following order: chat_id > open_id > user_id > email ||ou_5ad573a6411d72b8305fda3a9c15c70e|
root_id | string | No | To reply to a particular message, add the corresponding message ID||om_40eb06e7b84dc71c03e009ad3c754195|
msg_type | string | Yes | Message type, must be “image” ||image|
content | string | Yes | Message content |-|-| ∟image_key | string | Yes | image_key can be obtained via the Upload image API||1a0c4cb9-c680-4371-924c-ddb5f2750c3d| ### Request body example
```json
{
   "open_id":"ou_5ad573a6411d72b8305fda3a9c15c70e",
   "chat_id":"oc_5ad11d72b830411d72b836c20",  
   "root_id":"om_40eb06e7b84dc71c03e009ad3c754195",
    "user_id": "92e39a99",
    "email":"fanlv@gmail.com",
    "msg_type":"image",
    "content":{
        "image_key": "1a0c4cb9-c680-4371-924c-ddb5f2750c3d"
    }
}
```

### Curl request demo

```json 
curl -X POST \
  https://open.larksuite.com/open-apis/message/v4/send/ \
  -H 'Authorization: Bearer t-fee42159a366c575f2cd2b2acde2ed1e94c89d5f' \
  -H 'Content-Type: application/json' \
  -H 'cache-control: no-cache' \
  -d '{
    "chat_id": "oc_f5b1a7eb27ae2c7b6adc2a74faf339ff",
    "msg_type": "image",
    "content": {
        "image_key": "cb46f811-98ce-4348-aeb3-859168b1053a"
    }
}'
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
