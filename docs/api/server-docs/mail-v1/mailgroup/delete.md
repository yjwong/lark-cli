---
title: "Delete Mail Group"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/mail-v1/mailgroup/delete"
method: "DELETE"
api_path: "https://open.larksuite.com/open-apis/mail/v1/mailgroups/:mailgroup_id"
service: "mail-v1"
resource: "mailgroup"
section: "Email"
rate_limit: "Special Rate Limit"
scopes:
  - "mail:mailgroup"
updated: "1745840340000"
---

# Delete a mailing list

Deletes a mailing list.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/mail/v1/mailgroups/:mailgroup_id |
| HTTP Method | DELETE |
| Rate Limit | Special Rate Limit |
| Supported app types | custom |
| Required scopes | `mail:mailgroup` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `mailgroup_id` | `string` | Mailing list ID or mailing list address **Example value**: "xxxxxxxxxxxxxxx or test_mail_group@xxx.xx" | ## Response

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
