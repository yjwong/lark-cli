---
title: "Batch Delete Public Mailbox Members"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/mail-v1/public_mailbox-member/batch_delete"
method: "DELETE"
api_path: "https://open.larksuite.com/open-apis/mail/v1/public_mailboxes/:public_mailbox_id/members/batch_delete"
service: "mail-v1"
resource: "public_mailbox-member"
section: "Email"
rate_limit: "20 per second"
scopes:
  - "mail:public_mailbox"
updated: "1745840347000"
---

# Batch delete public mailbox members

A single request can delete multiple members of a public mailbox.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/mail/v1/public_mailboxes/:public_mailbox_id/members/batch_delete |
| HTTP Method | DELETE |
| Rate Limit | 20 per second |
| Supported app types | custom |
| Required scopes | `mail:public_mailbox` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `public_mailbox_id` | `string` | The unique ID or email address of a public mailbox **Example value**: "xxxxxxxxxxxxxxx or test_public_mailbox@xxx.xx" | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `member_id_list` | `string[]` | Yes | List of public mailbox member IDs deleted by this call **Example value**: ["xxxxxxx"] **Data validation rules**: - Length range: `1` ～ `200` | ### Request body example

{
    "member_id_list": [
        "xxxxxxx"
    ]
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
| 404 | 1234016 | public mailbox not found | Check whether the public mailbox exists. |
| 404 | 1234040 | public mailbox member not found | Check whether the public mailbox member exists. |
