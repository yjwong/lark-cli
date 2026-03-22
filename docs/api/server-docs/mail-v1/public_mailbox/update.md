---
title: "Modify All Information Of Public Mailbox"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/mail-v1/public_mailbox/update"
method: "PUT"
api_path: "https://open.larksuite.com/open-apis/mail/v1/public_mailboxes/:public_mailbox_id"
service: "mail-v1"
resource: "public_mailbox"
section: "Email"
rate_limit: "Special Rate Limit"
scopes:
  - "mail:public_mailbox"
updated: "1745840342000"
---

# Modify all information of public mailbox

Updates all information of a public mailbox.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/mail/v1/public_mailboxes/:public_mailbox_id |
| HTTP Method | PUT |
| Rate Limit | Special Rate Limit |
| Supported app types | custom |
| Required scopes | `mail:public_mailbox` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `public_mailbox_id` | `string` | Unique identifier of a public mailbox or the public mailbox address **Example value**: "xxxxxxxxxxxxxxx or test_public_mailbox@xxx.xx" | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `email` | `string` | No | Public mailbox address **Example value**: "test_public_mailbox@xxx.xx" |
| `name` | `string` | No | Public mailbox name **Example value**: "test public mailbox" | ### Request body example

{
   "name": "xxx",
   "email": "xxx@xxx.xxx"
}

## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `public_mailbox` | \- |
|   `public_mailbox_id` | `string` | Unique identifier of a public mailbox |
|   `email` | `string` | Public mailbox address |
|   `name` | `string` | Public mailbox name | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
        "public_mailbox_id": "xxx",
        "email": "xx@xx.xx",
        "name":"xxx"
    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 404 | 1234016 | public mailbox not found | Check whether the public mailbox exists. |
| 409 | 1234033 | email address has been used by another member as login account | email address has been used by another member as login account |
| 400 | 1234008 | request parameter error | Check whether the request parameters are correct. |
| 409 | 1234006 | email address has been used | The email address is already in use. Please use another email address. |
| 409 | 1234023 | email address alias exceed the number limit | email address alias exceed the number limit |
