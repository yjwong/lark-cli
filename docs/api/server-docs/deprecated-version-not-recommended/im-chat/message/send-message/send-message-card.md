---
title: "Send Message Card"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uYTNwUjL2UDM14iN1ATN"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/message/v4/send/"
section: "Deprecated Version (Not Recommended)"
updated: "1647175631000"
---

# Send Message Card

Sends a message card to a specified user or chat, including private chats and group chats.

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
--|--|--|---|--|--
open_id user_id  email  chat_id  | string | No | To send a private chat message to a user, use either open_id, email, or user_id. To send a message to a group, use the chat_id for the group (obtain via the Obtains Group List API). The server reads the string in the following order: chat_id > open_id > user_id > email||om_4abcdefg1234567890ad8|
msg_type|string|Yes |Message type, must be "interactive"l||interactive|
card|object|Yes|Message content, refer Card structure|
| root_id                                              | string | No   | The open_message_id of the message that needs to be replied ||om_40eb06e7b84dc71c03e009ad3c754195
| update_multi                                     | bool | No  | Controls whether the card is a shared card (all users share the same message card), the default is false,Process reference Interaction module ||true| ### Request body example
```json
{
    "chat_id": "oc_abcdefg1234567890",
    "msg_type": "interactive",
    "root_id":"om_4*********************ad8",
    "update_multi":false,
    "card": {
        	// card content
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
