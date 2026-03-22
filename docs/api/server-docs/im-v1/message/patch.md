---
title: "Update the message card sent by the app"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/message/patch"
method: "PATCH"
api_path: "https://open.larksuite.com/open-apis/im/v1/messages/:message_id"
service: "im-v1"
resource: "message"
section: "Messaging"
rate_limit: "1000 per minute、50 per second"
scopes:
  - "im:message"
  - "im:message:send_as_bot"
  - "im:message:update"
updated: "1717574918000"
---

# Update the message card sent by the app

Updates the content of message cards sent by an app.

> Notes:
> - Bot ability must be enabled.
> - If the message is updated with user_access_token, the operation user must be the sender of the card message.
> - Only supports updating **unrecalled** "[Shared cards](ukTMukTMukTM/uAjNwUjLwYDM14CM2ATN)" messages for everyone, you need to explicitly declare "update_multi": true in the card's config attribute.
> - **Updating batch messages is not supported.**
> - The maximum text message request body cannot exceed 150KB; the maximum card and rich text message request body cannot exceed 30KB.
> - Only supports modifying messages sent within 14 days.
> - Single message update frequency control is **5QPS**.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/im/v1/messages/:message_id |
| HTTP Method | PATCH |
| Rate Limit | 1000 per minute、50 per second |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `im:message` `im:message:send_as_bot` `im:message:update` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `message_id` | `string` | ID of the message to be updated. Only support update message card (`interactive` type). For the description, refer to Message ID description **Example value**: "om_dc13264520392913993dd051dba21dcf" | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `content` | `string` | Yes | Message content. The value is in JSON format. For more information about the card format, see Sent message content. **Example value**: "Reference URL" | ### Request body example

{
    "content": "Reference URL"
}

## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- | ### Response body example

{
    "code": 0,
    "data": {},
    "msg": "ok"
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 230001 | Your request contains an invalid request parameter. | Parameter error. Please check the input parameter according to the error message returned by the interface and refer to the document. |
| 400 | 230006 | Bot ability is not activated. | Bot ability is not enabled. The bot ability can be added on the [Developer Console](https://open.larksuite.com/app) -> Features -> Add Features, and it will take effect after a new version is released. |
| 400 | 230011 | The message is recalled. | The operation is not supported because the message has been recalled. |
| 400 | 230013 | Bot has NO availability to this user. | The bot is not available for the user. You can edit the available scope of the application on the [Developer Console](https://open.larksuite.com/app)  -> App Versions -> Version Management & Release -> Create a version, and it will take effect after the new version is released. |
| 400 | 230020 | This operation triggers the frequency limit. | The current action triggers the frequency limit. Please reduce the request frequency. |
| 400 | 230025 | The length of the message content reaches its limit. | The length of the message body exceeds the limit. The maximum size of the text message body is 150 KB and the maximum size of the card and rich text message body is 30 KB. In addition, if the message contains a large number of styles, the actual length of the message body will be greater than the length of the request body you entered. |
| 400 | 230027 | Lack of necessary permissions. | This operation is not currently supported in external groups. |
| 400 | 230028 | The messages do NOT pass the audit. | The message DLP review failed. This error may be triggered when the message content contains plaintext phone numbers, plaintext email, etc. Please check the message content according to the error information returned by the interface. |
| 400 | 230031 | Message can only be modified within 14 days after sending. | The content of a message can be modified only within 14 days after the message is sent. |
| 400 | 230099 | Failed to create card content. | Unable to create the card. For the reason of the failure, please refer to the interface error message. |
| 400 | 230110 | Action unavailable as the message has been deleted. | Action unavailable as the message has been deleted. |
| 400 | 232009 | Your request specifies a chat which has already been dissolved. | The group has been deleted. |
