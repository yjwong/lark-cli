---
title: "Clear Public Mailbox Members"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/mail-v1/public_mailbox-member/clear"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/mail/v1/public_mailboxes/:public_mailbox_id/members/clear"
service: "mail-v1"
resource: "public_mailbox-member"
section: "Email"
rate_limit: "Special Rate Limit"
scopes:
  - "mail:public_mailbox"
updated: "1745840347000"
---

# Clear public mailbox members

Deletes all members of a public mailbox.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/mail/v1/public_mailboxes/:public_mailbox_id/members/clear |
| HTTP Method | POST |
| Rate Limit | Special Rate Limit |
| Supported app types | custom |
| Required scopes | `mail:public_mailbox` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `public_mailbox_id` | `string` | Unique identifier of a public mailbox or the public mailbox address **Example value**: "xxxxxxxxxxxxxxx or test_public_mailbox@xxx.xx" | ## Response

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
