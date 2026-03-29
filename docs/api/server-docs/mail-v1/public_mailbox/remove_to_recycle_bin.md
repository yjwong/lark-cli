---
title: "Move Public Mailbox To The Recycle Bin"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/mail-v1/public_mailbox/remove_to_recycle_bin"
method: "DELETE"
api_path: "https://open.larksuite.com/open-apis/mail/v1/public_mailboxes/:public_mailbox_id/remove_to_recycle_bin"
service: "mail-v1"
resource: "public_mailbox"
section: "Email"
rate_limit: "20 per minute"
scopes:
  - "mail:public_mailbox"
updated: "1745840346000"
---

# Move Public Mailbox To The Recycle Bin

Move Public Mailbox To The Recycle Bin

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/mail/v1/public_mailboxes/:public_mailbox_id/remove_to_recycle_bin |
| HTTP Method | DELETE |
| Rate Limit | 20 per minute |
| Supported app types | custom |
| Required scopes | `mail:public_mailbox` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `public_mailbox_id` | `string` | Public email address **Example value**: "test_public_mailbox@xxx.xx" | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `to_mail_address` | `string` | No | Please provide the email address to receive the deleted emails. If left blank, the emails of the public mailbox will be deleted without being forwarded **Example value**: "user@xxx.xx" | ### Request body example

{
    "to_mail_address": "user@xxx.xx"
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
    "msg": "success",
    "data": {}
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 1230001 | Parameter Error | Please retry after modifying the parameters. |
| 500 | 1230003 | Internal Error | Please try again later. |
| 403 | 1230002 | No Permission | Please retry after checking the relevant permissions. |
