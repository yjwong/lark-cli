---
title: "Obtain All Member's Email Aliases"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/mail-v1/user_mailbox-alias/list"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/mail/v1/user_mailboxes/:user_mailbox_id/aliases"
service: "mail-v1"
resource: "user_mailbox-alias"
section: "Email"
rate_limit: "1000 per minute、50 per second"
scopes:
  - "mail:user_mailbox"
  - "mail:user_mailbox:readonly"
updated: "1745840275000"
---

# Obtain all member's email aliases

Obtains all member's email aliases.

> This interface returns all the data at one time, and the pagination parameter is invalid

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/mail/v1/user_mailboxes/:user_mailbox_id/aliases |
| HTTP Method | GET |
| Rate Limit | 1000 per minute、50 per second |
| Supported app types | custom |
| Required scopes Enable any scope from the list | `mail:user_mailbox` `mail:user_mailbox:readonly` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `user_mailbox_id` | `string` | Member's email address **Example value**: "user@xxx.xx" | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `page_token` | `string` | No | Page identifier. It is not filled in the first request, indicating traversal from the beginning; when there will be more groups, the new page_token will be returned at the same time, and the next traversal can use the page_token to get more groups **Example value**: xxx |
| `page_size` | `int` | No | **Example value**: 10 **Default value**: `20` **Data validation rules**: - Maximum value: `20` | ## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `items` | `email_alias[]` | Member's email alias |
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
| 404 | 1235013 | user not found | Check whether the member's email exists. |
| 400 | 1234008 | request parameter error | Check whether the request parameters are correct. |
