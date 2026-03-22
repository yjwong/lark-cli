---
title: "Modify Email Contact's Info"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/mail-v1/user_mailbox-mail_contact/patch"
method: "PATCH"
api_path: "https://open.larksuite.com/open-apis/mail/v1/user_mailboxes/:user_mailbox_id/mail_contacts/:mail_contact_id"
service: "mail-v1"
resource: "user_mailbox-mail_contact"
section: "Email"
rate_limit: "20 per minute"
scopes:
  - "mail:user_mailbox.mail_contact:write"
updated: "1745840270000"
---

# Modify Email Contact's Info

When using the tenant_access_token, it is necessary to apply for data permissions for accessing email contacts resources.

> When using the tenant_access_token, it is necessary to apply for the relevant data permissions.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/mail/v1/user_mailboxes/:user_mailbox_id/mail_contacts/:mail_contact_id |
| HTTP Method | PATCH |
| Rate Limit | 20 per minute |
| Supported app types | custom |
| Required scopes | `mail:user_mailbox.mail_contact:write` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `user_mailbox_id` | `string` | User email address or enter "me" to represent the current user making the API call. **Example value**: "user@xxx.xx or me" |
| `mail_contact_id` | `string` | Email Contact's ID. For the acquisition method, see List Email Contacts. **Example value**: "123" | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `name` | `string` | Yes | Email Contact's Name **Example value**: "Sam" **Data validation rules**: - Length range: `1` ～ `64` characters |
| `company` | `string` | No | Email Contact's Company **Example value**: "Lark LLC" **Data validation rules**: - Maximum length: `64` characters |
| `phone` | `string` | No | Email Contact's Phone Number **Example value**: "19912341234" **Data validation rules**: - Maximum length: `40` characters |
| `mail_address` | `string` | No | Email Contact's Mail Address **Example value**: "sam@example.com" **Data validation rules**: - Maximum length: `319` characters |
| `tag` | `string` | No | Email Contact's Tag **Example value**: "Friend" **Data validation rules**: - Maximum length: `64` characters |
| `remark` | `string` | No | Email Contact's Remark **Example value**: "Met at a music party" **Data validation rules**: - Maximum length: `1000` characters |
| `position` | `string` | No | Email Contact's Position **Example value**: "CEO" **Data validation rules**: - Maximum length: `64` characters | ### Request body example

{
    "name": "Sam",
    "company": "Lark LLC",
    "phone": "19912341234",
    "mail_address": "sam@example.com",
    "tag": "Friend",
    "remark": "Met at a music party",
    "position": "CEO"
}

## Response

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
| 400 | 1230001 | Parameter Error | Please retry after modifying the parameters. |
| 500 | 1230003 | Internal Error | Please try again later. |
| 403 | 1230002 | No Permission | Call this API after becoming a member of the public email or applying for relevant data permissions. |
