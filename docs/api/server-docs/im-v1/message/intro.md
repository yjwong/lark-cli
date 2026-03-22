---
title: "Resource introduction"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/message/intro"
service: "im-v1"
resource: "message"
section: "Messaging"
updated: "1717574888000"
---

# Resource introduction

## Resource definition
This API indicates a message in the Lark chat.

## Field description
| Parameter | Type | Description |
| --- | --- | --- |
| `message_id` | `string` | Message ID; for the description, refer to Message ID description. |
| `root_id` | `string` | Root message ID; used for replying to message scenarios; for the description, refer to Message ID description. |
| `parent_id` | `string` | Parent message ID; used for replying to message scenarios; for the description, refer to Message ID description. |
| `msg_type` | `string` | Message types include: text, post, image, file, audio, media, sticker, interactive, share_chat, and share_user; for the definitions of types, refer to Sent message content. |
| `create_time` | `string` | Message generation timestamp (in ms). |
| `update_time` | `string` | Message update timestamp (in ms). |
| `deleted` | `boolean` | Whether the message is recalled. |
| `updated` | `boolean` | Whether the message is updated. |
| `chat_id` | `string` | Group of the message. |
| `sender` | `sender` | Sender, which can be the user or app. |
| ∟ `id` | `string` | This field identifies the sender's ID |
| ∟ `id_type` | `string` | This field identifies the sender's ID type, such as `open_id`, `app_id`. |
| ∟ `sender_type` | `string` | This field identifies the sender type, Optional values are: `user`, `app`. |
| ∟ `tenant_key` | `string` | A tenant's unique ID in Lark. It can be used to obtain tenant_access_token, and can also be used as a tenant's unique ID in an app. |
| `body` | `message_body` | Message content |
| ∟ `content` | `string` | Message content, which is a JSON string; different msg_type values correspond to different contents; message types include text, post, image, file, audio, media, sticker, interactive, share_chat, and share_user; for the description of the format, refer to Sent message content. |
| `mentions` | `mention[]` | List of IDs of the mentioned users or bots |
| ∟ `key` | `string` | Serial number of the mentioned user or bot; for example, for the third mentioned member, the value is "@_user_3". |
| ∟ `id` | `string` | The ID of the mentioned user or bot. |
| ∟ `id_type` | `string` | ID type of the mentioned user or bot; currently, such as `open_id` (What is Open ID?). |
| ∟ `name` | `string` | Name of the mentioned user or bot. |
| ∟ `tenant_key` | `string` | Tenant key. A tenant's unique ID in Lark. It can be used to obtain tenant_access_token, and can also be used as a tenant's unique ID in an app. |
| `thread_id` | `string` | The thread ID to which the message belongs. This ID can be used to forward the thread, obtain all the messages in the thread, etc. Please refer to Thread Introduction. |
| `upper_message_id` | `string` | message_id in the upper level of the combined and forwarded message; for the description, refer to Message ID description. | ### Data example
```json
{
        "message_id": "om_dc13264520392913993dd051dba21dcf",
        "root_id": "om_40eb06e7b84dc71c03e009ad3c754195",
        "parent_id": "om_d4be107c616aed9c1da8ed8068570a9f",
        "msg_type": "card",
        "create_time": "1615380573411",
        "update_time": "1615380573411",
        "deleted": false,
        "updated": false,
        "chat_id": "oc_5ad11d72b830411d72b836c20",
        "sender": {
            "id": "cli_9f427eec54ae901b",
            "id_type": "app_id",
            "sender_type": "app",
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
        "thread_id": "omt_16f3c7e1268f1749",
        "upper_message_id": "om_40eb06e7b84dc71c03e009ad3c754195"
}
```

## Message-related ID description

### 1. message_id description
`message_id` is the unique identifier of a message. When you send a message, the system automatically generates a unique ID for the message.
The `message_id` value is a string starting with om_, for example: om_934be5776f5a87239a298af9e74c0f72.

**How to obtain message_id?**

- Method 1: Call the send messages API and obtain the `message_id` field from the response.

- Method 2: Monitor the receive messages event and obtain the `message_id` field of the message entity from the event body.

- Method 3: Call the Read chat history API and obtain the `message_id` field from the response.

### 2. Description of root_id and parent_id
- `root_id`: If a message replies to another message, then `root_id` is the `message_id` of the root message in the reply tree.
- `parent_id`: If a message replies to another message, then `parent_id` is the `message_id` of the latter message.

**The figure below shows an example message reply:**
- msg2 replies to msg1, then both `root_id` and `parent_id` of msg2 are `message_id` of msg1.
- msg3 replies to msg2, then `root_id` of msg3 is `message_id` of msg1, and `parent_id` is `message_id` of msg2.

![Group 645 (1).png.lin.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/3b75512e590a54b656c3d8d1acaec8df_8eVxOoCKqw.png?height=1132&lazyload=true&width=1640)
### 3. Description of upper_message_id
`upper_message_id` is `message_id` of the combined and forwarded message where sub-messages of the combined and forwarded message are located.

**The figure below shows an example message reply:**

msg1, msg2, and msg3 are three sub-messages in the combined and forwarded message, and their `upper_message` is `message_id` of the combined and forwarded message where they are located.

![Group 646.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/51638a19a1d02857cee04a27a285e8f1_qhptW2G5O0.png?height=1132&lazyload=true&width=1640)
