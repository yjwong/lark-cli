---
title: "Modify All Information Of Mail Group"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/mail-v1/mailgroup/update"
method: "PUT"
api_path: "https://open.larksuite.com/open-apis/mail/v1/mailgroups/:mailgroup_id"
service: "mail-v1"
resource: "mailgroup"
section: "Email"
rate_limit: "100 per second"
scopes:
  - "mail:mailgroup"
updated: "1745840340000"
---

# Modify all information of mail group

Updates all information of a mailing list.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/mail/v1/mailgroups/:mailgroup_id |
| HTTP Method | PUT |
| Rate Limit | 100 per second |
| Supported app types | custom |
| Required scopes | `mail:mailgroup` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `mailgroup_id` | `string` | Mailing list ID or mailing list address **Example value**: "xxxxxxxxxxxxxxx or test_mail_group@xxx.xx" | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `email` | `string` | No | Mailing list address **Example value**: "test_mail_group@xxx.xx" |
| `name` | `string` | No | Mailing list name **Example value**: "test mail group" |
| `description` | `string` | No | Mailing list description **Example value**: "mail group for testing" |
| `who_can_send_mail` | `string` | No | Who can send emails to this mailing list **Example value**: "ALL_INTERNAL_USERS" **Optional values are**:  - `ANYONE`: Anyone - `ALL_INTERNAL_USERS`: Organization members only - `ALL_GROUP_MEMBERS`: Mailing list members only - `CUSTOM_MEMBERS`: Specified members  | ### Request body example

{ 
   "email": "xxx@xxx.com",
    "name": "xxx",
    "description": "xxx",
    "who_can_send_mail": "ANYONE"
}

## Response

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
        "mailgroup_id": "xxx",
        "email": "xx@xx.xx",
        "name":"xxx",
        "description":"xxx",
        "direct_members_count":"x",
        "include_external_member": true,
        "include_all_company_member":false,
        "who_can_send_mail":"ANYONE"
    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 404 | 1234013 | mail group not found | Check whether the mailing list exists. |
| 400 | 1234008 | request parameter error | Check whether the request parameters are correct. |
| 409 | 1234033 | email address has been used by another member as login account | email address has been used by another member as login account |
| 409 | 1234006 | email address has been used | The email address is already in use. Please use another email address. |
| 409 | 1234023 | email address alias exceed the number limit | email address alias exceed the number limit |
| 400 | 1234030 | mailing group count is over max limit | The maximum number of mailing groups has been reached. Please try deleting some before attempting again. |
