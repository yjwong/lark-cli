---
title: "Delete A Member's Email Alias"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/mail-v1/user_mailbox-alias/delete"
method: "DELETE"
api_path: "https://open.larksuite.com/open-apis/mail/v1/user_mailboxes/:user_mailbox_id/aliases/:alias_id"
service: "mail-v1"
resource: "user_mailbox-alias"
section: "Email"
rate_limit: "1000 per minute、50 per second"
scopes:
  - "mail:user_mailbox"
updated: "1745840275000"
---

# Delete a member's email alias

Deletes a member's email alias.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/mail/v1/user_mailboxes/:user_mailbox_id/aliases/:alias_id |
| HTTP Method | DELETE |
| Rate Limit | 1000 per minute、50 per second |
| Supported app types | custom |
| Required scopes | `mail:user_mailbox` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `user_mailbox_id` | `string` | Member's email address **Example value**: "user@xxx.xx" |
| `alias_id` | `string` | Alias email address **Example value**: "user_alias@xxx.xx" | ## Response

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
| 400 | 1234008 | request parameter error | Check whether the request parameters are correct. |
| 404 | 1235013 | user not found | Check whether the member's email exists. |
| 404 | 1234012 | alias not found | Alias address does not exist. |
