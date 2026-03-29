---
title: "Create A Member's Email Alias"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/mail-v1/user_mailbox-alias/create"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/mail/v1/user_mailboxes/:user_mailbox_id/aliases"
service: "mail-v1"
resource: "user_mailbox-alias"
section: "Email"
rate_limit: "1000 per minute、50 per second"
scopes:
  - "mail:user_mailbox"
updated: "1745840275000"
---

# Create a member's email alias

Creates a member's email alias.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/mail/v1/user_mailboxes/:user_mailbox_id/aliases |
| HTTP Method | POST |
| Rate Limit | 1000 per minute、50 per second |
| Supported app types | custom |
| Required scopes | `mail:user_mailbox` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `user_mailbox_id` | `string` | Member's email address **Example value**: "user@xxx.xx" | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `email_alias` | `string` | No | Alias address **Example value**: "xxx@xx.xxx" | ### Request body example

{
    "email_alias": "xxx@xx.xxx"
}

## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `user_mailbox_alias` | `email_alias` | Member's email alias |
|     `primary_email` | `string` | Primary email address |
|     `email_alias` | `string` | Alias address | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
        "user_mailbox_alias": {
            "primary_email": "xxx@xx.xxx",
            "email_alias": "xxx@xx.xxx"
        }
    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 409 | 1234006 | email address has been used | The email address is already in use. Please use another email address. |
| 400 | 1234008 | request parameter error | Check whether the request parameters are correct. |
| 404 | 1234019 | mail address's domain not found | Check whether this domain exists. |
| 409 | 1235002 | email alias address has been used | The alias email address is already in use. Please use another email address. |
| 400 | 1235003 | Service unavailable | Please try again later. |
| 404 | 1235013 | user not found | Check whether the member's email exists. |
| 409 | 1234033 | email address has been used by another member as login account | email address has been used by another member as login account |
