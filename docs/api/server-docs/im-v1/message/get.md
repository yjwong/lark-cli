---
title: "Obtain the content of the specified message"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/message/get"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/im/v1/messages/:message_id"
service: "im-v1"
resource: "message"
section: "Messaging"
rate_limit: "1000 per minute、50 per second"
scopes:
  - "im:message"
  - "im:message:readonly"
field_scopes:
  - "contact:user.employee_id:readonly"
updated: "1717574901000"
---

# Obtain the content of the specified message

Query message content with message_id.

> Notes:
> - Bot ability must be enabled.
> - The bot must be in the group chat.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/im/v1/messages/:message_id |
| HTTP Method | GET |
| Rate Limit | 1000 per minute、50 per second |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `im:message` `im:message:readonly` |
| Required field scopes | > The response body of the API contains the following sensitive fields, and they will be returned only after corresponding scopes are added. If you do not need the fields, it is not recommended that you request the scopes. `contact:user.employee_id:readonly` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `message_id` | `string` | ID of the message whose content is to be obtained. For the description, refer to Message ID description **Example value**: "om_dc13264520392913993dd051dba21dcf" | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `user_id_type` | `string` | No | User ID categories **Example value**: open_id **Optional values are**:  - `open_id`: Identifies a user to an app. The same user has different Open IDs in different apps. How to get Open ID - `union_id`: Identifies a user to a tenant that acts as a developer. A user has the same Union ID in apps developed by the same developer, and has different Union IDs in apps developed by different developers. A developer can use Union ID to link the same user's identities in multiple apps.How to get Union ID - `user_id`: Identifies a user to a tenant. The same user has different User IDs in different tenants. In one single tenant, a user has the same User ID in all apps （including store apps）. User ID is usually used to communicate user data between different apps. How to get User ID  **Default value**: `open_id` **When the value is `user_id`, the following field scopes are required**: `contact:user.employee_id:readonly` | ## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `items` | `message[]` | If the specified message type is merge_forward, the returned data includes 1 merge forward message + N sub-messages. |
|     `message_id` | `string` | Message ID; for the description, refer to Message ID description |
|     `root_id` | `string` | Root message ID; used for replying to message scenarios; for the description, refer to Message ID description |
|     `parent_id` | `string` | Parent message ID; used for replying to message scenarios; for the description, refer to Message ID description |
|     `thread_id` | `string` | The thread ID to which the message belongs (not returned indicates that the message is not a thread message).  For the description, refer to Thread Introduction |
|     `msg_type` | `string` | Message types include: text, post, image, file, audio, media, sticker, interactive, share_chat, and share_user; for the definitions of types, refer to Received message content |
|     `create_time` | `string` | Message generation timestamp (in ms) |
|     `update_time` | `string` | Message update timestamp (in ms) |
|     `deleted` | `boolean` | Whether the message is recalled or deleted |
|     `updated` | `boolean` | Whether the message is updated |
|     `chat_id` | `string` | Group of the message |
|     `sender` | `sender` | Sender, which can be the user or app |
|       `id` | `string` | This field identifies the sender's ID |
|       `id_type` | `string` | This field identifies the sender's id type. If the sender is an app, the type is app_id; if the sender is a user, the type is consistent with the `user_id_type` parameter. |
|       `sender_type` | `string` | This field identifies the sender type **Optional values are:** - `user` - `app` - `anonymous` - `unknown` |
|       `tenant_key` | `string` | A tenant's unique ID in Lark. It can be used to obtain tenant_access_token, and can also be used as a tenant's unique ID in an app. |
|     `body` | `message_body` | Message content |
|       `content` | `string` | Message content, which is a string after the serialization of the JSON structure; different msg_type values correspond to different contents; message types include text, post, image, file, audio, media, sticker, interactive, share_chat, and share_user; for the definitions of types, refer to  Received message content  **Note**: - The content of the card message is inconsistent with the card JSON obtained in the card building tool, and returning the original card JSON is not currently supported. |
|     `mentions` | `mention[]` | List of IDs of the mentioned users or bots |
|       `key` | `string` | Serial number of the mentioned user or bot; for example, for the third mentioned member, the value is "@_user_3" |
|       `id` | `string` | The ID of the mentioned user or bot **Notice**: - When `user_id_type` is not filled in, this field will return the open_id of the mentioned user or robot. - When `user_id_type` is filled in, this field will return mentioned user's corresponding type ID (open_id, union_id or user_id) or the robot's app_id. |
|       `id_type` | `string` | ID type of the mentioned user or bot |
|       `name` | `string` | Name of the mentioned user or bot |
|       `tenant_key` | `string` | A tenant's unique ID in Lark. It can be used to obtain tenant_access_token, and can also be used as a tenant's unique ID in an app. |
|     `upper_message_id` | `string` | message_id in the upper level of the combined and forwarded message; for the description, refer to Message ID description | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
        "items": [
            {
                "message_id": "om_dc13264520392913993dd051dba21dcf",
                "root_id": "om_40eb06e7b84dc71c03e009ad3c754195",
                "parent_id": "om_d4be107c616aed9c1da8ed8068570a9f",
                "thread_id": "omt_d4be107c616a",
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
                    "content": "{\"text\":\"test content\"}"
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
                "upper_message_id": "om_40eb06e7b84dc71c03e009ad3c754195"
            }
        ]
    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 230001 | Your request contains an invalid request parameter. | Parameter error. Please check the input parameter according to the error message returned by the interface and refer to the document. |
| 400 | 230002 | The bot can not be outside the group. | The bot isn't in the group chat. |
| 400 | 230006 | Bot ability is not activated. | Bot ability is not enabled. The bot ability can be added on the [Developer Console](https://open.larksuite.com/app) -> Features -> Add Features, and it will take effect after a new version is released. |
| 400 | 230013 | Bot has NO availability to this user. | The bot is not available for the user. You can edit the available scope of the application on the [Developer Console](https://open.larksuite.com/app)  -> App Versions -> Version Management & Release -> Create a version, and it will take effect after the new version is released. |
| 400 | 230027 | Lack of necessary permissions. | Please supplement the required permissions according to the permission requirements section in this API document. |
| 400 | 230050 | The message is invisible to the operator. | This operation cannot be performed on invisible messages. |
| 400 | 230110 | Action unavailable as the message has been deleted. | Action unavailable as the message has been deleted. |
