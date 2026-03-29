---
title: "Check All Public Mailboxes"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/mail-v1/public_mailbox/list"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/mail/v1/public_mailboxes"
service: "mail-v1"
resource: "public_mailbox"
section: "Email"
rate_limit: "Special Rate Limit"
scopes:
  - "mail:public_mailbox"
  - "mail:public_mailbox:readonly"
field_scopes:
  - "contact:user.employee_id:readonly"
  - "mail:public_mailbox.public_mailbox_geo"
updated: "1745840342000"
---

# Check all public mailboxes

Obtains the list of public mailboxes by pages.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/mail/v1/public_mailboxes |
| HTTP Method | GET |
| Rate Limit | Special Rate Limit |
| Supported app types | custom |
| Required scopes Enable any scope from the list | `mail:public_mailbox` `mail:public_mailbox:readonly` |
| Required field scopes | > The response body of the API contains the following sensitive fields, and they will be returned only after corresponding scopes are added. If you do not need the fields, it is not recommended that you request the scopes. `contact:user.employee_id:readonly` `mail:public_mailbox.public_mailbox_geo` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `page_token` | `string` | No | Page identifier. It is not filled in the first request, indicating traversal from the beginning; when there will be more groups, the new page_token will be returned at the same time, and the next traversal can use the page_token to get more groups **Example value**: xxx |
| `page_size` | `int` | No | **Example value**: 10 **Default value**: `20` **Data validation rules**: - Maximum value: `200` | ## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `has_more` | `boolean` | Whether the response body has more parameters |
|   `page_token` | `string` | Page identifier, when has_more is true, a new page_token will also be returned. Otherwise, page_token will not be returned |
|   `items` | `public_mailbox[]` | Public mailbox list |
|     `public_mailbox_id` | `string` | Unique identifier of a public mailbox |
|     `email` | `string` | Public mailbox address |
|     `name` | `string` | Public mailbox name |
|     `geo` | `string` | geo location of public mailbox **Required field scopes**: `mail:public_mailbox.public_mailbox_geo` | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
        "has_more": true,
        "page_token": "xxx",
        "items": [
            {
                "public_mailbox_id": "xxxxxxxxxxxxxxx",
                "email": "test_public_mailbox@xxx.xx",
                "name": "test public mailbox",
                "geo": "cn"
            }
        ]
    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 503 | 1235003 | Service unavailable | Please try again later. |
