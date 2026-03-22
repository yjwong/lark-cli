---
title: "Delete A Public Mailbox Alias"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/mail-v1/public_mailbox-alias/delete"
method: "DELETE"
api_path: "https://open.larksuite.com/open-apis/mail/v1/public_mailboxes/:public_mailbox_id/aliases/:alias_id"
service: "mail-v1"
resource: "public_mailbox-alias"
section: "Email"
rate_limit: "1000 per minute、50 per second"
scopes:
  - "mail:public_mailbox"
updated: "1745840347000"
---

# Delete a public mailbox alias

Deletes a public mailbox alias.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/mail/v1/public_mailboxes/:public_mailbox_id/aliases/:alias_id |
| HTTP Method | DELETE |
| Rate Limit | 1000 per minute、50 per second |
| Supported app types | custom |
| Required scopes | `mail:public_mailbox` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `public_mailbox_id` | `string` | Public mailbox ID or public mailbox address **Example value**: "xxxxxx or xxx@xx.xxx" |
| `alias_id` | `string` | Public mailbox alias **Example value**: "xxx@xx.xxx" | ## Response

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
| 404 | 1234016 | public mailbox not found | Check whether the public mailbox exists. |
| 400 | 1234008 | request parameter error | Check whether the request parameters are correct. |
| 404 | 1234012 | alias not found | Alias address does not exist. |
