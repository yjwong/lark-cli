---
title: "Recall messages in batches"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/batch_message/delete"
method: "DELETE"
api_path: "https://open.larksuite.com/open-apis/im/v1/batch_messages/:batch_message_id"
service: "im-v1"
resource: "batch_message"
section: "Messaging"
rate_limit: "1000 per minute、50 per second"
scopes:
  - "im:message:send_multi_depts"
  - "im:message:send_multi_users"
updated: "1717574914000"
---

# Recall messages in batches

Recall messages in batches
Recall messages sent through the Send messages in batches API.

> Notes:
> - Bot ability must be enabled.
> - To recall a single sent message, please use the Recall a message API.
> - Not supported to recall messages that have been sent for more than 1 day.
> - Each call involves a large number of messages and therefore this API is an asynchronous API, which incurs a delay.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/im/v1/batch_messages/:batch_message_id |
| HTTP Method | DELETE |
| Rate Limit | 1000 per minute、50 per second |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `im:message:send_multi_depts` `im:message:send_multi_users` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `batch_message_id` | `string` | Task ID of the batch message that you need to recall, which is the `message_id` field in the return value of the Send a batch message API, which is used to identify a batch message request. **Example value**: "bm-dc13264520392913993dd051dba21dcf" | ### Request example
```
curl --location --request DELETE 'https://open.larksuite.com/open-apis/im/v1/batch_messages/bm-dc13264520392913993dd051dba21dcf' \
--header 'Authorization: Bearer t-e811da948d9a2803433be10458f650c842799970'
```

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
| 400 | 230009 | Message can only be recalled within one days. | The message was sent too long ago (more than 1 day) to perform the current action. |
| 400 | 230026 | No permission to recall this message. | The bot doesn't have the permission to recall the message, but the message sent by the itself. |
| 400 | 230027 | Lack of necessary permissions. | Please supplement the required permissions according to the permission requirements section in this document. |
| 400 | 230030 | The message id is invalid, record not found. | Check whether the batch_message_id parameter passed in is correct. |
