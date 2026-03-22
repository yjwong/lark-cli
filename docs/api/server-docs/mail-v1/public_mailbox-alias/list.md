---
title: "Obtain All Public Mailbox Aliases"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/mail-v1/public_mailbox-alias/list"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/mail/v1/public_mailboxes/:public_mailbox_id/aliases"
service: "mail-v1"
resource: "public_mailbox-alias"
section: "Email"
rate_limit: "1000 per minute、50 per second"
scopes:
  - "mail:public_mailbox"
  - "mail:public_mailbox:readonly"
updated: "1745840347000"
---

# Obtain all public mailbox aliases

Obtains all public mailbox aliases.

> This interface returns all the data at one time, and the pagination parameter is invalid

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/mail/v1/public_mailboxes/:public_mailbox_id/aliases |
| HTTP Method | GET |
| Rate Limit | 1000 per minute、50 per second |
| Supported app types | custom |
| Required scopes Enable any scope from the list | `mail:public_mailbox` `mail:public_mailbox:readonly` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `public_mailbox_id` | `string` | Public mailbox ID or public mailbox address **Example value**: "xxxxxx or xxx@xx.xxx" | ## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `items` | `email_alias[]` | Public mailbox alias |
|     `primary_email` | `string` | Primary email address |
|     `email_alias` | `string` | Alias address | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
        "items": [
            {
                "primary_email": "xxx@xx.xxx",
                "email_alias": "xxx@xx.xxx"
            }
        ]
    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 404 | 1234016 | public mailbox not found | Check whether the public mailbox exists. |
| 400 | 1234008 | request parameter error | Check whether the request parameters are correct. |
