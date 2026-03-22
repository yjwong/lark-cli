---
title: "Get chat history"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/message/list"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/im/v1/messages"
service: "im-v1"
resource: "message"
section: "Messaging"
rate_limit: "1000 per minute、50 per second"
scopes:
  - "im:message"
  - "im:message:readonly"
  - "im:message.history:readonly"
updated: "1717574901000"
---

# Get chat history

Obtains chat history (chat records) of chats (including private chats and group chats).

> By default, the API-level scope can only obtain private chat messages. If you need to obtain group chat messages, the app must also have the scope **Read all messages in associated group chat**.

> - Bot ability must be enabled.
> - When you obtain messages, the bot must be in the group chat.
> - For thread messages in ordinary conversation groups, only the root message of the thread can be obtained through the "chat" container type, and all messages in the thread replies can be obtained by specifying the container type as "thread".

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/im/v1/messages |
| HTTP Method | GET |
| Rate Limit | 1000 per minute、50 per second |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `im:message` `im:message:readonly` `im:message.history:readonly` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `container_id_type` | `string` | Yes | Container type **Options**： - `chat`：including p2p chat and group chat. - `thread`: thread **Example value**: chat |
| `container_id` | `string` | Yes | Container ID, which can be filled in: - chat_id, refer to Group ID Description - thread_id, refer to Thread Introduction **Example value**: oc_234jsi43d3ssi993d43545f |
| `start_time` | `string` | No | Start time of historical information **Notice:** - The thread container type currently does not support obtaining messages within a specified time range. **Example value**: 1608594809 |
| `end_time` | `string` | No | End time of historical information **Notice:** - The thread container type currently does not support obtaining messages within a specified time range. **Example value**: 1609296809 |
| `sort_type` | `string` | No | message sorting type **Example value**: ByCreateTimeAsc **Optional values are**:  - `ByCreateTimeAsc`: Sort by message creation time ascending - `ByCreateTimeDesc`: Sort by message creation time descending  **Default value**: `ByCreateTimeAsc` |
| `page_size` | `int` | No | **Example value**: 20 **Default value**: `20` **Data validation rules**: - Value range: `1` ～ `50` |
| `page_token` | `string` | No | Page identifier. It is not filled in the first request, indicating traversal from the beginning; when there will be more groups, the new page_token will be returned at the same time, and the next traversal can use the page_token to get more groups **Example value**: GxmvlNRvP0NdQZpa7yIqf_Lv_QuBwTQ8tXkX7w-irAghVD_TvuYd1aoJ1LQph86O-XImC4X9j9FhUPhXQDvtrQ | ::: note
When using the `page_token`, the sorting type is consistent with the first request, and it does not support changing the sorting method midway.

## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `has_more` | `boolean` | Whether the response body has more parameters |
|   `page_token` | `string` | Page identifier, when has_more is true, a new page_token will also be returned. Otherwise, page_token will not be returned |
|   `items` | `message[]` | message[] |
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
|       `id_type` | `string` | This field identifies the sender's ID type **Optional values are:** - `open_id` - `app_id` |
|       `sender_type` | `string` | This field identifies the sender type **Optional values are:** - `user` - `app` - `anonymous` - `unknown` |
|       `tenant_key` | `string` | A tenant's unique ID in Lark. It can be used to obtain tenant_access_token, and can also be used as a tenant's unique ID in an app. |
|     `body` | `message_body` | Message content |
|       `content` | `string` | Message content, which is a string after the serialization of the JSON structure; different msg_type values correspond to different contents; message types include text, post, image, file, audio, media, sticker, interactive, share_chat, and share_user; for the definitions of types, refer to Received message content  **Note**: - The content of the card message is inconsistent with the card JSON obtained in the card building tool, and returning the original card JSON is not currently supported. |
|     `mentions` | `mention[]` | List of IDs of the mentioned users or bots |
|       `key` | `string` | Serial number of the mentioned user or bot; for example, for the third mentioned member, the value is "@_user_3" |
|       `id` | `string` | open_id of the mentioned user or bot |
|       `id_type` | `string` | ID type of the mentioned user or bot; currently, only `open_id` is supported (What is Open ID?) |
|       `name` | `string` | Name of the mentioned user or bot |
|       `tenant_key` | `string` | A tenant's unique ID in Lark. It can be used to obtain tenant_access_token, and can also be used as a tenant's unique ID in an app. |
|     `upper_message_id` | `string` | message_id in the upper level of the combined and forwarded message; for the description, refer to Message ID description | ### Response body example

{
	"code": 0,
	"msg": "success",
	"data": {
		"has_more": false,
		"page_token": "GxmvlNRvP0NdQZpa7yIqf_Lv_QuBwTQ8tXkX7w-irAghVD_TvuYd1aoJ1LQph86O-XImC4X9j9FhUPhXQDvtrQ",
		"items": [{
			"message_id": "om_dc13264520392913993dd051dba21dcf",
			"root_id": "om_40eb06e7b84dc71c03e009ad3c754195",
			"parent_id": "om_d4be107c616aed9c1da8ed8068570a9f",
			"msg_type": "interactive",
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
			"mentions": [{
				"key": "@_user_1",
				"id": "ou_155184d1e73cbfb8973e5a9e698e74f2",
				"id_type": "open_id",
				"name": "Tom",
				"tenant_key": "736588c9260f175e"
			}],
			"upper_message_id": "om_40eb06e7b84dc71c03e009ad3c754195"
		}]
	}
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 230001 | Your request contains an invalid request parameter. | Parameter error. Please check the input parameter according to the error message returned by the interface and refer to the document. |
| 400 | 230002 | The bot can not be outside the group. | The bot isn't in the group chat. |
| 400 | 230006 | Bot ability is not activated. | Bot ability is not enabled. The bot ability can be added on the [Developer Console](https://open.larksuite.com/app) -> Features -> Add Features, and it will take effect after a new version is released. |
| 400 | 230027 | Lack of necessary permissions. | Please supplement the required permissions according to the permission requirements in this document, such as the permission to obtain all messages in the group. In addition, this operation is not currently supported in external groups. |
| 400 | 230073 | The thread is invisible to the operator. | The thread is not visible to the operator. If "New members can see chat history" is turned off in the group chat and the thread was created before the operator entered the group chat, the thread needs to be passively subscribed by the operator to be visible, such as other users @ the operator in the thread. |
| 400 | 230110 | Action unavailable as the message has been deleted. | Action unavailable as the message has been deleted. | For other unlisted error codes, please refer to Server Error Codes.
