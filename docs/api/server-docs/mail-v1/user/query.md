---
title: "Email Address Query"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/mail-v1/user/query"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/mail/v1/users/query"
service: "mail-v1"
resource: "user"
section: "Email"
rate_limit: "20 per second"
scopes:
  - "mail:user_mailbox"
updated: "1745840275000"
---

# Email address query

Using the query API, you can enter an email address to query the type and status corresponding to the email address.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/mail/v1/users/query |
| HTTP Method | POST |
| Rate Limit | 20 per second |
| Supported app types | custom |
| Required scopes | `mail:user_mailbox` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `email_list` | `string[]` | Yes | email list for query **Example value**: ["aaa@lark.com"] | ### Request body example

{
    "email_list": [
        "aaa@lark.com"
    ]
}

## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `user_list` | `user[]` | Email address,status,type |
|     `email` | `string` | Email Address |
|     `status` | `int` | Email Address Status **Optional values are**:  - `1`: Email address invalid - `2`: Email address domain name does not exist - `3`: Email address does not exist - `4`: Enable - `5`: Deleted (Mailbox Recycle Bin) - `6`: Disable  |
|     `type` | `int` | Email Address Type **Optional values are**:  - `1`: Member mailbox - `2`: Member mailbox alias - `3`: Public mailbox - `4`: Public mailbox alias - `5`: Mail group - `6`: Mail group alias  | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
        "user_list": [
            {
                "email": "aaa@lark.com",
                "status": 4,
                "type": 1
            }
        ]
    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 1230001 | param is invalis | Check whether the mailbox address is correct. |
