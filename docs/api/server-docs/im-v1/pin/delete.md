---
title: "Unpin a message"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/pin/delete"
method: "DELETE"
api_path: "https://open.larksuite.com/open-apis/im/v1/pins/:message_id"
service: "im-v1"
resource: "pin"
section: "Messaging"
rate_limit: "1000 per minute、50 per second"
scopes:
  - "im:message"
  - "im:message.pins:write_only"
  - "im:message:send_as_bot"
updated: "1717574935000"
---

# Unpin a message

Remove the pin of a specified message.

> Notes:
> - Bot ability must be enabled.
> - When you unpin a message, the bot must be in the group chat.
> - If the message has not been pinned or has been ed, return a message that the operation was successful.
> - Cannot remove a pin message that is invisible to the operator.
> - The operation of removing a pin for the same message cannot exceed 5 QPS.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/im/v1/pins/:message_id |
| HTTP Method | DELETE |
| Rate Limit | 1000 per minute、50 per second |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `im:message` `im:message.pins:write_only` `im:message:send_as_bot` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `message_id` | `string` | Message ID to be unpin. For the description, refer to Message ID description **Example value**: "om_dc13264520392913993dd051dba21dcf" | ## Response

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
| 400 | 230013 | Bot has NO availability to this user. | The bot is not available for the user. You can edit the available scope of the application on the [Developer Console](https://open.larksuite.com/app)  -> App Versions -> Version Management & Release -> Create a version, and it will take effect after the new version is released. |
| 400 | 230027 | Lack of necessary permissions. | This operation is not currently supported in external groups. |
| 400 | 230045 | The chat not exist. | The group chat to which the message belongs does not exist. Please check whether the group chat has been dissolved. |
| 400 | 230046 | No Permission to Pin/Unpin messages in the chat. | This chat setting only allows chat owners and chat administrators to pin / unpin messages. Please check the identity of the operator in the chat. |
| 400 | 230047 | Pin/Unpin message trigger message_id limit. | Pin / Unpin trigger rate-limiting strategy for the same message. Please reduce the request speed and try again later. |
| 400 | 230050 | The message is invisible to the operator. | This operation cannot be performed on invisible messages. |
| 400 | 232009 | Your request specifies a chat which has already been dissolved. | The group has been deleted. | > **Note:** For other server-side error codes, refer to Server-side error codes.
