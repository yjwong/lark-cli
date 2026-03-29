---
title: "Create A Public Mailbox"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/mail-v1/public_mailbox/create"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/mail/v1/public_mailboxes"
service: "mail-v1"
resource: "public_mailbox"
section: "Email"
rate_limit: "50 per second"
scopes:
  - "mail:public_mailbox"
field_scopes:
  - "mail:public_mailbox.public_mailbox_geo"
updated: "1745840342000"
---

# Create a public mailbox

Creates a public mailbox.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/mail/v1/public_mailboxes |
| HTTP Method | POST |
| Rate Limit | 50 per second |
| Supported app types | custom |
| Required scopes | `mail:public_mailbox` |
| Required field scopes | > The response body of the API contains the following sensitive fields, and they will be returned only after corresponding scopes are added. If you do not need the fields, it is not recommended that you request the scopes. `mail:public_mailbox.public_mailbox_geo` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `email` | `string` | No | Public mailbox address **Example value**: "test_public_mailbox@xxx.xx" |
| `name` | `string` | No | Public mailbox name **Example value**: "test public mailbox" |
| `geo` | `string` | No | geo location of public mailbox **Example value**: "cn" | ### Request body example

{
    "email": "test_public_mailbox@xxx.xx",
    "name": "test public mailbox",
    "geo": "cn"
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
|   `name` | `string` | Public mailbox name |
|   `geo` | `string` | geo location of public mailbox **Required field scopes**: `mail:public_mailbox.public_mailbox_geo` | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
        "public_mailbox_id": "xxxxxxxxxxxxxxx",
        "email": "test_public_mailbox@xxx.xx",
        "name": "test public mailbox",
        "geo": "cn"
    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 409 | 1234006 | email address has been used | The email address is already in use. Please use another email address. |
| 400 | 1234008 | email or name is invalid | The email address or name is invalid. Please modify it and try again. |
| 409 | 1234033 | email address has been used by another member as login account | email address has been used by another member as login account |
| 200 | 1234026 | The number of tenant public mailboxes has reached the limit and cannot be created | The maximum number of public mailboxes has been reached. Please delete some before trying again. |
| 400 | 1234040 | not set geo name auth | not set geo name auth |
| 400 | 1234041 | tenant not open mg not set geo name | The tenant cannot set the geo field if the MG has not been activated |
| 400 | 1234042 | req set geo not find in geo list | req set user geo not find in geo list |
