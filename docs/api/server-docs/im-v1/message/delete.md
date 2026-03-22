---
title: "Recall message"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/message/delete"
method: "DELETE"
api_path: "https://open.larksuite.com/open-apis/im/v1/messages/:message_id"
service: "im-v1"
resource: "message"
section: "Messaging"
rate_limit: "1000 per minute、50 per second"
scopes:
  - "im:message"
  - "im:message:recall"
  - "im:message:send_as_bot"
updated: "1717574897000"
---

# Recall message

The bot recalls messages it sent or the group owner recalls messages within a group.

> Notes:
> -  Bot ability must be enabled, and the bot must be in the chat when recalling messages.
> - The bot can recall messages sent by itself in a single chat or group, and the sending time does not exceed the retractable time limit configured by the tenant administrator (the default is 24 hours).
> - If the bot wants to recall a message sent by one another in the group, the bot must be the owner, admin or creator of the group, and the message was sent within 1 year.
> - Messages sent via "Send messages in batches" API **can't be recalled**.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/im/v1/messages/:message_id |
| HTTP Method | DELETE |
| Rate Limit | 1000 per minute、50 per second |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `im:message` `im:message:recall` `im:message:send_as_bot` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token | > **Note:** The access token corresponds to the entity that initiates a recall.

### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `message_id` | `string` | ID of the message to be recalled **Example value**: "om_dc13264520392913993dd051dba21dcf" | ## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {}
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 230001 | Your request contains an invalid request parameter. | Parameter error. Please check the input parameter according to the error message returned by the interface and refer to the document. |
| 400 | 230002 | The bot can not be outside the group. | The bot isn't in the group chat. |
| 400 | 230006 | Bot ability is not activated. | Bot ability is not enabled. The bot ability can be added on the [Developer Console](https://open.larksuite.com/app) -> Features -> Add Features, and it will take effect after a new version is released. |
| 400 | 230009 | Message can only be recalled within one days. | The message was sent too long ago (more than 1 day) to perform the current action. |
| 400 | 230013 | Bot has NO availability to this user. | The bot is not available for the user. You can edit the available scope of the application on the [Developer Console](https://open.larksuite.com/app)  -> App Versions -> Version Management & Release -> Create a version, and it will take effect after the new version is released. |
| 400 | 230026 | No permission to recall this message. | The bot doesn't have the permission to recall the message, but the message sent by the itself. |
| 400 | 230027 | Lack of necessary permissions. | Currently, this operation is not supported in external groups. |
| 400 | 230050 | The message is invisible to the operator. | This operation cannot be performed on invisible messages. |
| 400 | 230054 | This operation is not supported for this message type. | The message type does not support this operation. For details, please refer to Some APIs under message and group domain add unsupported message type verification. |
| 400 | 230110 | Action unavailable as the message has been deleted. | Action unavailable as the message has been deleted. |
| 400 | 232009 | Your request specifies a chat which has already been dissolved. | The group has been deleted. |
