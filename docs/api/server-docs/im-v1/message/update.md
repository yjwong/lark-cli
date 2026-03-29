---
title: "Edit message"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/message/update"
method: "PUT"
api_path: "https://open.larksuite.com/open-apis/im/v1/messages/:message_id"
service: "im-v1"
resource: "message"
section: "Messaging"
rate_limit: "1000 per minute、50 per second"
scopes:
  - "im:message"
  - "im:message:send_as_bot"
  - "im:message:update"
updated: "1717574897000"
---

# Edit message

Edit the content of the sent message, currently supports editing text and rich text messages. To update the message card, please refer to Update the message card sent by the app.

> **Notes**:
> - A message can be edited up to 20 times.
> - Only messages sent by the operator himself can be edited.
> - Uneditable messages that have been withdrawn, deleted, beyond editable time.
> - The operator must be in the group to which the message belongs and have permission to speak before editing the message.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/im/v1/messages/:message_id |
| HTTP Method | PUT |
| Rate Limit | 1000 per minute、50 per second |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `im:message` `im:message:send_as_bot` `im:message:update` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `message_id` | `string` | The ID of the message to be edited, only text or post type messages are supported, see Message ID Description **Example value**: "om_dc13264520392913993dd051dba21dcf" | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `msg_type` | `string` | Yes | The type of message, currently only text and post types are supported **Example value**: "text" |
| `content` | `string` | Yes | Message content, string after serialization of JSON structure. Different msg_type correspond to different content. For specific format description, please refer to: Send Message Content **Note:** - JSON string needs to be escaped, such as'\\ n' after newline escape - The maximum text message request body cannot exceed 150KB - Rich text message request body cannot exceed 30KB at most **Example value**: "{\"text\":\"test content\"}" | ### Request body example

{
    "msg_type": "text",
    "content": "{\"text\":\"test content\"}"
}

## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `message` | \- |
|   `message_id` | `string` | Message ID open_message_id |
|   `root_id` | `string` | Root message ID open_message_id |
|   `parent_id` | `string` | The ID of the parent message open_message_id |
|   `thread_id` | `string` | The thread ID to which the message belongs (not returned indicates that the message is not a thread message).  For the description, refer to Thread Introduction |
|   `msg_type` | `string` | Message type text, post, card, image, etc. |
|   `create_time` | `string` | Message generation timestamp (milliseconds) |
|   `update_time` | `string` | Message update timestamp |
|   `deleted` | `boolean` | Whether the message was withdrawn |
|   `updated` | `boolean` | Is the message updated? |
|   `chat_id` | `string` | Group to which it belongs |
|   `sender` | `sender` | Sender, which can be a user or an application |
|     `id` | `string` | This field identifies the sender's id |
|     `id_type` | `string` | This field identifies the sender's id type |
|     `sender_type` | `string` | This field identifies the type of sender |
|     `tenant_key` | `string` | Tenant key |
|   `body` | `message_body` | Message content, JSON structure |
|     `content` | `string` | Message jsonContent |
|   `mentions` | `mention[]` | The id of the person or application @ |
|     `key` | `string` | Mention key |
|     `id` | `string` | open_id of the mentioned user or bot |
|     `id_type` | `string` | ID type of the mentioned user or bot; currently, only `open_id` is supported (What is Open ID?) |
|     `name` | `string` | Name of the user at issue |
|     `tenant_key` | `string` | Tenant key |
|   `upper_message_id` | `string` | The previous level message id of the merged message open_message_id | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
        "message_id": "om_dc13264520392913993dd051dba21dcf",
        "root_id": "om_40eb06e7b84dc71c03e009ad3c754195",
        "parent_id": "om_d4be107c616aed9c1da8ed8068570a9f",
        "thread_id": "omt_d4be107c616a",
        "msg_type": "card",
        "create_time": "1609296809",
        "update_time": "1609336806",
        "deleted": false,
        "updated": true,
        "chat_id": "oc_5ad11d72b830411d72b836c20",
        "sender": {
            "id": "cli_9f427eec54ae901b",
            "id_type": "app_id",
            "sender_type": "App",
            "tenant_key": "736588c9260f175e"
        },
        "body": {
            "content": "{\"text\":\"@_user_1 test content\"}"
        },
        "mentions": [
            {
                "key": "@_user_1",
                "id": "ou_155184d1e73cbfb8973e5a9e698e74f2",
                "id_type": "open_id",
                "name": "Tom",
                "tenant_key": "736588c9260f175e"
            }
        ],
        "upper_message_id": "om_40eb06e7b84dc71c03e00ida3c754892"
    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 230001 | Your request contains an invalid request parameter. | Parameter error. Please check the input parameter according to the error message returned by the interface and refer to the document. |
| 400 | 230002 | The bot can not be outside the group. | The bot isn't in the group chat. |
| 400 | 230006 | Bot ability is not activated. | Bot ability is not enabled. The bot ability can be added on the [Developer Console](https://open.larksuite.com/app) -> Features -> Add Features, and it will take effect after a new version is released. |
| 400 | 230011 | The message is recalled. | The operation is not supported because the message has been recalled. |
| 400 | 230013 | Bot has NO availability to this user. | The bot is not available for the user. You can edit the available scope of the application on the [Developer Console](https://open.larksuite.com/app)  -> App Versions -> Version Management & Release -> Create a version, and it will take effect after the new version is released. |
| 400 | 230018 | These operations are NOT allowed at current group settings. | The current action is blocked by group settings. Please check the group settings or contact the group administrator. |
| 400 | 230020 | This operation triggers the frequency limit. | The current action triggers the frequency limit. Please reduce the request frequency. |
| 400 | 230022 | The content of the message contains sensitive information. | The message contains sensitive information. Please check the message content. |
| 400 | 230025 | The length of the message content reaches its limit. | The length of the message body exceeds the limit. The maximum size of the text message body is 150 KB and the maximum size of the card and rich text message body is 30 KB. In addition, if the message contains a large number of styles, the actual length of the message body will be greater than the length of the request body you entered. |
| 400 | 230027 | Lack of necessary permissions. | Editing messages in external groups or editing messages as a user is not currently supported. |
| 400 | 230028 | The messages do NOT pass the audit. | The message DLP review failed. This error may be triggered when the message content contains plaintext phone numbers, plaintext email, etc. Please check the message content according to the error information returned by the interface. |
| 400 | 230035 | Send message permission deny. | Do not have permission to send messages. Please check whether the group has been muted, or is controlled by the communication permission of the tenant dimension. |
| 400 | 230038 | Cross tenant p2p chat operate forbid. | Cross tenant P2P chat is not allowed to send messages through this interface. |
| 400 | 230050 | The message is invisible to the operator. | This operation cannot be performed on invisible messages. |
| 400 | 230054 | This operation is not supported for this message type. | The message type does not support this operation. |
| 400 | 230055 | The type of file upload does not match the type of message being sent. | The type of file upload does not match the type of message being sent. |
| 400 | 230071 | The operator is not the sender of the message. | The operator is not the sender of the message. |
| 400 | 230072 | The message has reached the number of times it can be edited. | The message has reached the number of times it can be edited. |
| 400 | 230073 | The message in the secret group is not allowed to be edited. | The message in the secret group is not allowed to be edited. |
| 400 | 230074 | The message in the third-party encryption group is not allowed to be edited. | The message in the third-party encryption group is not allowed to be edited. |
| 400 | 230075 | The message has exceeded the time that can be edited. | The message has exceeded the time that can be edited. |
| 400 | 230110 | Action unavailable as the message has been deleted. | The message has been deleted. |
| 400 | 232009 | Your request specifies a chat which has already been dissolved. | The group the message belongs to has been disbanded. |
