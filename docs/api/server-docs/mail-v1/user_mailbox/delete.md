---
title: "Release Address From Recycle Bin"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/mail-v1/user_mailbox/delete"
method: "DELETE"
api_path: "https://open.larksuite.com/open-apis/mail/v1/user_mailboxes/:user_mailbox_id"
service: "mail-v1"
resource: "user_mailbox"
section: "Email"
rate_limit: "20 per minute"
scopes:
  - "mail:user_mailbox"
updated: "1745840274000"
---

# Release address from recycle bin

Deletes a member's email address from the email recycle bin. This deletes a member's email address forever, which means the email address cannot be recovered once deleted. This API also allows you to transfer emails from the mailbox to be released to another mailbox.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/mail/v1/user_mailboxes/:user_mailbox_id |
| HTTP Method | DELETE |
| Rate Limit | 20 per minute |
| Supported app types | custom |
| Required scopes | `mail:user_mailbox` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `user_mailbox_id` | `string` | The email address to be released **Example value**: "111111@abc.com" | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `transfer_mailbox` | `string` | No | The email address that receives the transferred emails **Example value**: 888888@abc.com | ## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- | ### Response body example

{
    "code": 0,
    "data": {},
    "msg": "release mail address success"
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 1230001 | param is invalis | Check whether the mailbox address is correct. |
