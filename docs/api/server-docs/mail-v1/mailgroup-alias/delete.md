---
title: "Delete A Mailing List Alias"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/mail-v1/mailgroup-alias/delete"
method: "DELETE"
api_path: "https://open.larksuite.com/open-apis/mail/v1/mailgroups/:mailgroup_id/aliases/:alias_id"
service: "mail-v1"
resource: "mailgroup-alias"
section: "Email"
rate_limit: "1000 per minute、50 per second"
scopes:
  - "mail:mailgroup"
updated: "1745840341000"
---

# Delete a mailing list alias

Deletes a mailing list alias.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/mail/v1/mailgroups/:mailgroup_id/aliases/:alias_id |
| HTTP Method | DELETE |
| Rate Limit | 1000 per minute、50 per second |
| Supported app types | custom |
| Required scopes | `mail:mailgroup` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `mailgroup_id` | `string` | Mailing list ID or mailing list address **Example value**: "xxxxxx or test_group@xx.xxx" |
| `alias_id` | `string` | Mailing list alias email address **Example value**: "xxx@xx.xxx" | ## Response

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
| 404 | 1234012 | alias not found | Alias address does not exist. |
| 400 | 1234008 | request parameter error | Check whether the request parameters are correct. |
