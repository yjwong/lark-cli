---
title: "Reorder Auto Filters"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/mail-v1/user_mailbox-rule/reorder"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/mail/v1/user_mailboxes/:user_mailbox_id/rules/reorder"
service: "mail-v1"
resource: "user_mailbox-rule"
section: "Email"
rate_limit: "1 per second"
scopes:
  - "mail:user_mailbox.rule:write"
updated: "1745840339000"
---

# Reorder Auto Filters

Reorder Auto Filters

> When using the tenant_access_token, it is necessary to apply for data permissions for accessing auto filter resources.

> When using this interface, you need to pass all auto filter IDs.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/mail/v1/user_mailboxes/:user_mailbox_id/rules/reorder |
| HTTP Method | POST |
| Rate Limit | 1 per second |
| Supported app types | custom |
| Required scopes | `mail:user_mailbox.rule:write` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `user_mailbox_id` | `string` | **Example value**: "user@xxx.xx or me" | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `rule_ids` | `string[]` | Yes | Auto filter ID list. The method for obtaining ID is shown in List Auto Filters **Example value**: ["11111111"] **Data validation rules**: - Minimum length: `1` | ### Request body example

{
    "rule_ids": [
        "11111111"
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
| 400 | 1230001 | Parameter Error | Please retry after modifying the parameters. |
| 500 | 1230003 | Internal Error | Please try again later. |
| 403 | 1230002 | No Permission | Call this API after becoming a member of the public email or applying for relevant data permissions. |
