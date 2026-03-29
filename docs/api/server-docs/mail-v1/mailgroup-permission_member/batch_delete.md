---
title: "Batch Delete Mail Group Permission Members"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/mail-v1/mailgroup-permission_member/batch_delete"
method: "DELETE"
api_path: "https://open.larksuite.com/open-apis/mail/v1/mailgroups/:mailgroup_id/permission_members/batch_delete"
service: "mail-v1"
resource: "mailgroup-permission_member"
section: "Email"
rate_limit: "20 per second"
scopes:
  - "mail:mailgroup"
updated: "1745840342000"
---

# Batch delete mail group permission members

A single request can delete multiple permission members in a mail group.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/mail/v1/mailgroups/:mailgroup_id/permission_members/batch_delete |
| HTTP Method | DELETE |
| Rate Limit | 20 per second |
| Supported app types | custom |
| Required scopes | `mail:mailgroup` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `mailgroup_id` | `string` | The unique ID or email address of a mail group **Example value**: "xxxxxxxxxxxxxxx or test_mail_group@xxx.xx" | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `permission_member_id_list` | `string[]` | Yes | List of permission member IDs deleted by this call **Example value**: ["xxxxxxx"] **Data validation rules**: - Length range: `1` ～ `200` | ### Request body example

{
    "permission_member_id_list": [
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
| 404 | 1234013 | mail group not found | Check whether the mailing list exists. |
| 404 | 1234015 | mail group permission member not found | Check whether the member who can send emails to mailing list addresses exists. |
