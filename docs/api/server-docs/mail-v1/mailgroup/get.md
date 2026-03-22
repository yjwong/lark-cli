---
title: "Query The Specified Mail Group"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/mail-v1/mailgroup/get"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/mail/v1/mailgroups/:mailgroup_id"
service: "mail-v1"
resource: "mailgroup"
section: "Email"
rate_limit: "Special Rate Limit"
scopes:
  - "mail:mailgroup"
  - "mail:mailgroup:readonly"
updated: "1745840340000"
---

# Query the specified mail group

Obtains the information of a mailing list.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/mail/v1/mailgroups/:mailgroup_id |
| HTTP Method | GET |
| Rate Limit | Special Rate Limit |
| Supported app types | custom |
| Required scopes Enable any scope from the list | `mail:mailgroup` `mail:mailgroup:readonly` | ### Request header
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
| `data` | `mailgroup` | \- |
|   `mailgroup_id` | `string` | Mailing list ID |
|   `email` | `string` | Mailing list address |
|   `name` | `string` | Mailing list name |
|   `description` | `string` | Mailing list description |
|   `direct_members_count` | `string` | Number of mailing list members |
|   `include_external_member` | `boolean` | Whether it includes external members |
|   `include_all_company_member` | `boolean` | Whether it is an all-staff mailing list |
|   `who_can_send_mail` | `string` | Who can send emails to this mailing list **Optional values are**:  - `ANYONE`: Anyone - `ALL_INTERNAL_USERS`: Organization members only - `ALL_GROUP_MEMBERS`: Mailing list members only - `CUSTOM_MEMBERS`: Specified members  | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
        "mailgroup_id": "xxxxxxxxxxxxxxx",
        "email": "test_mail_group@xxx.xx",
        "name": "test mail group",
        "description": "mail group for testing",
        "direct_members_count": "10",
        "include_external_member": true,
        "include_all_company_member": false,
        "who_can_send_mail": "ALL_INTERNAL_USERS"
    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 404 | 1234013 | mail group not found | Check whether the mailing list exists. |
