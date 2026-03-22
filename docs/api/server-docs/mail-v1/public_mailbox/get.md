---
title: "Query The Specified Public Mailbox"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/mail-v1/public_mailbox/get"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/mail/v1/public_mailboxes/:public_mailbox_id"
service: "mail-v1"
resource: "public_mailbox"
section: "Email"
rate_limit: "Special Rate Limit"
scopes:
  - "mail:public_mailbox"
  - "mail:public_mailbox:readonly"
field_scopes:
  - "mail:public_mailbox.public_mailbox_geo"
updated: "1745840342000"
---

# Query the specified public mailbox

Obtains public mailbox information.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/mail/v1/public_mailboxes/:public_mailbox_id |
| HTTP Method | GET |
| Rate Limit | Special Rate Limit |
| Supported app types | custom |
| Required scopes Enable any scope from the list | `mail:public_mailbox` `mail:public_mailbox:readonly` |
| Required field scopes | > The response body of the API contains the following sensitive fields, and they will be returned only after corresponding scopes are added. If you do not need the fields, it is not recommended that you request the scopes. `mail:public_mailbox.public_mailbox_geo` | ### Request header
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
| 404 | 1234016 | public mailbox not found | Check whether the public mailbox exists. |
