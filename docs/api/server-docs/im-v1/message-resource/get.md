---
title: "Obtain resource files in messages"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/message-resource/get"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/im/v1/messages/:message_id/resources/:file_key"
service: "im-v1"
resource: "message-resource"
section: "Messaging"
rate_limit: "1000 per minute、50 per second"
scopes:
  - "im:message"
  - "im:message:readonly"
  - "im:message.history:readonly"
updated: "1717574901000"
---

# Obtain resource files in messages

Obtains resource files in messages, including audios, videos, images, and files. **Emoji resources can’t be downloaded**, and the resource files for download can’t exceed 100 MB.

> Notes:
> - Bot ability must be enabled.
> - The bot and message must be in the same chat.
> - Obtaining resource files of sub-messages in combined and forwarded messages is currently unsupported.
> - The type of the file can be obtained through the `Content-Type` field in the response header.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/im/v1/messages/:message_id/resources/:file_key |
| HTTP Method | GET |
| Rate Limit | 1000 per minute、50 per second |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `im:message` `im:message:readonly` `im:message.history:readonly` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `message_id` | `string` | Message ID for the resource to be queried **Example value**: "om_dc13264520392913993dd051dba21dcf" |
| `file_key` | `string` | Key of the resource to be queried. You can call the Obtain the content of the specified message API to query the resource Key in the message content through the message ID. **Note**: the requested file_key and message_id must match **Example value**: "file_456a92d6-c6ea-4de4-ac3f-7afcf44ac78g" | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `type` | `string` | Yes | Resource type **Optional values are:**: - `image`: corresponds to the image in the message or the image in the rich-text message. - `file`: corresponds to the file, audio, or video  in the message (except the emoji). **Example value**: image | ### Request example
```bash
curl --location --request GET 'https://open.larksuite.com/open-apis/im/v1/messages/om_aa0944122e2ec23545d9a93df0ceec02/resources/img_v2_98f3923d-3351-4017-8156-ee34d92196bj?type=image' \
--header 'Authorization: Bearer t-9fef198b83bfdb975f09297cf14c4a63a117c52e' 
```
```bash
curl --location --request GET 'https://open.larksuite.com/open-apis/im/v1/messages/om_4b6abce3eb423be3c2dfd11d4ca9e97b/resources/file_v2_83e69b47-5d67-4e21-bfb4-66f8bf80b7ej?type=file' \
--header 'Authorization: Bearer t-9fef198b83bfdb975f09297cf14c4a63a117c52e' 
```

## Response

When the HTTP status code is 200, it means success

Return file binary stream

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 230110 | Action unavailable as the message has been deleted. | Action unavailable as the message has been deleted. |
| 400 | 234001 | Invalid request param. | Check that the request parameters are correct. |
| 401 | 234002 | Unauthorized. | Authentication failed, contact Oncall to solve. |
| 400 | 234003 | File not in message. | The resource does not belong to the current message, please check the message ID and resource key. |
| 400 | 234004 | App not in chat. | The app is not in the same group as the message, please check if the message ID is correct. |
| 400 | 234009 | Lack of necessary permissions. | This operation is not currently supported in external groups. |
| 400 | 234019 | scope CheckAppTenant fail. | Please retry. |
| 400 | 234037 | Downloaded file size exceeds limit. | Downloaded file size exceeds 100MB limit forbid. |
| 400 | 234038 | Do not allow downloading of message resources in restricted mode. | Resource files in restricted messages cannot be downloaded. Please check whether the message has been set as restricted, or the group has enabled the restricted mode. |
| 400 | 234040 | The message is invisible to the operator. | The message is invisible to the operator, please contact the group owner or group administrator to check whether "New members can view historical messages" is turned off in the group settings. |
| 400 | 234041 | Tenant master key has been deleted, please contact the tenant administrator. | Tenant master key has been deleted, please contact the tenant administrator. |
| 400 | 234042 | Hybrid deployment tenant storage error, such as full storage space, please contact tenant administrator. | Hybrid deployment tenant storage error, such as full storage space, please contact tenant administrator. |
| 400 | 234043 | Unsupported message type. | Unsupported message type, such as card messages, merge and forward messages. |
