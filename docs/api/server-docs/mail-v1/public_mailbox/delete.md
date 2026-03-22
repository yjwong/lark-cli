---
title: "Permanently Delete Public Mailbox Address"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/mail-v1/public_mailbox/delete"
method: "DELETE"
api_path: "https://open.larksuite.com/open-apis/mail/v1/public_mailboxes/:public_mailbox_id"
service: "mail-v1"
resource: "public_mailbox"
section: "Email"
rate_limit: "20 per minute"
scopes:
  - "mail:public_mailbox"
updated: "1745840346000"
---

# Permanently delete public mailbox address

Releases a public mailbox address from the email recycle bin. This deletes a public mailbox address forever, which means the mailbox address cannot be recovered once deleted.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/mail/v1/public_mailboxes/:public_mailbox_id |
| HTTP Method | DELETE |
| Rate Limit | 20 per minute |
| Supported app types | custom |
| Required scopes | `mail:public_mailbox` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `public_mailbox_id` | `string` | Public mailbox address to be released **Example value**: "xxxxxx@abc.com" | ## Response

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
