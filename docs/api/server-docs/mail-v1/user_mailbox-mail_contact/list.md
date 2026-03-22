---
title: "List Email Contacts"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/mail-v1/user_mailbox-mail_contact/list"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/mail/v1/user_mailboxes/:user_mailbox_id/mail_contacts"
service: "mail-v1"
resource: "user_mailbox-mail_contact"
section: "Email"
rate_limit: "20 per minute"
scopes:
  - "mail:user_mailbox.mail_contact:read"
  - "mail:user_mailbox.mail_contact:write"
field_scopes:
  - "mail:user_mailbox.mail_contact.mail_address:read"
  - "mail:user_mailbox.mail_contact.phone:read"
updated: "1745840274000"
---

# List Email Contacts

List Email Contacts

> When using the tenant_access_token, it is necessary to apply for data permissions for accessing email contacts resources.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/mail/v1/user_mailboxes/:user_mailbox_id/mail_contacts |
| HTTP Method | GET |
| Rate Limit | 20 per minute |
| Supported app types | custom |
| Required scopes Enable any scope from the list | `mail:user_mailbox.mail_contact:read` `mail:user_mailbox.mail_contact:write` |
| Required field scopes | > The response body of the API contains the following sensitive fields, and they will be returned only after corresponding scopes are added. If you do not need the fields, it is not recommended that you request the scopes. `mail:user_mailbox.mail_contact.mail_address:read` `mail:user_mailbox.mail_contact.phone:read` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `user_mailbox_id` | `string` | User email address, or enter me to represent the current calling interface user **Example value**: "user@xxx.xx or me" | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `page_size` | `int` | Yes | paging size **Example value**: 20 **Data validation rules**: - Value range: `1` ～ `20` |
| `page_token` | `string` | No | Page identifier. It is not filled in the first request, indicating traversal from the beginning; when there will be more groups, the new page_token will be returned at the same time, and the next traversal can use the page_token to get more groups **Example value**: xxx | ## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `items` | `mail_contact[]` | Mail contact list |
|     `id` | `string` | Email Contact's ID |
|     `name` | `string` | Email Contact's Name |
|     `company` | `string` | Email Contact's Company |
|     `phone` | `string` | Email Contact's Phone Number **Required field scopes**: `mail:user_mailbox.mail_contact.phone:read` |
|     `mail_address` | `string` | Email Contact's Mail Address **Required field scopes**: `mail:user_mailbox.mail_contact.mail_address:read` |
|     `tag` | `string` | Email Contact's Tag |
|     `remark` | `string` | Email Contact's Remark |
|     `avatar` | `string` | Email Contact's Avatar |
|     `position` | `string` | Email Contact's Position |
|   `page_token` | `string` | Page identifier, when has_more is true, a new page_token will also be returned. Otherwise, page_token will not be returned |
|   `has_more` | `boolean` | Whether the response body has more parameters | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
        "items": [
            {
                "id": "7256274576546463764",
                "name": "Sam",
                "company": "Lark LLC",
                "phone": "19912341234",
                "mail_address": "sam@example.com",
                "tag": "Friend",
                "remark": "Met at a music party",
                "avatar": "https://exampeimg.com/xxxx.jpg",
                "position": "CEO"
            }
        ],
        "page_token": "xxx",
        "has_more": true
    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 1230001 | Parameter Error | Please retry after modifying the parameters. |
| 403 | 1230002 | Internal Error | Please try again later. |
| 500 | 1230003 | No Permission | Call this API after becoming a member of the public email or applying for relevant data permissions. |
