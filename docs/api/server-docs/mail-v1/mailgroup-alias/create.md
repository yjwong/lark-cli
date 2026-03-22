---
title: "Create A Mailing List Alias"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/mail-v1/mailgroup-alias/create"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/mail/v1/mailgroups/:mailgroup_id/aliases"
service: "mail-v1"
resource: "mailgroup-alias"
section: "Email"
rate_limit: "1000 per minute、50 per second"
scopes:
  - "mail:mailgroup"
updated: "1745840341000"
---

# Create a mailing list alias

Creates a mailing list alias.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/mail/v1/mailgroups/:mailgroup_id/aliases |
| HTTP Method | POST |
| Rate Limit | 1000 per minute、50 per second |
| Supported app types | custom |
| Required scopes | `mail:mailgroup` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `mailgroup_id` | `string` | Mailing list ID or mailing list address **Example value**: "xxxxxx or xxx@xx.xxx" | ### Request body

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
|   `mailgroup_alias` | `email_alias` | Mailing list alias |
|     `primary_email` | `string` | Primary email address |
|     `email_alias` | `string` | Alias address | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
        "mailgroup_alias": {
            "primary_email": "xxx@xx.xxx",
            "email_alias": "xxx@xx.xxx"
        }
    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 409 | 1234006 | email address has been used | The email address is already in use. Please use another email address. |
| 404 | 1234013 | mail group not found | Check whether the mailing list exists. |
| 400 | 1234008 | request parameter error | Check whether the request parameters are correct. |
| 409 | 1235002 | email alias address has been used | The alias email address is already in use. Please use another email address. |
| 400 | 1235003 | Service unavailable | Please try again later. |
| 404 | 1234019 | mail address's domain not found | Check whether this domain exists. |
| 409 | 1234033 | email address has been used by another member as login account | email address has been used by another member as login account |
