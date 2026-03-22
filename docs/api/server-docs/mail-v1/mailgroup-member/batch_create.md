---
title: "Batch Create Mail Group Members"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/mail-v1/mailgroup-member/batch_create"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/mail/v1/mailgroups/:mailgroup_id/members/batch_create"
service: "mail-v1"
resource: "mailgroup-member"
section: "Email"
rate_limit: "20 per second"
scopes:
  - "mail:mailgroup"
field_scopes:
  - "contact:user.employee_id:readonly"
updated: "1745840341000"
---

# Batch create mail group members

You can add multiple members to a mail group at a time.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/mail/v1/mailgroups/:mailgroup_id/members/batch_create |
| HTTP Method | POST |
| Rate Limit | 20 per second |
| Supported app types | custom |
| Required scopes | `mail:mailgroup` |
| Required field scopes | > The response body of the API contains the following sensitive fields, and they will be returned only after corresponding scopes are added. If you do not need the fields, it is not recommended that you request the scopes. `contact:user.employee_id:readonly` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `mailgroup_id` | `string` | The unique ID or email address of a mail group **Example value**: "xxxxxxxxxxxxxxx or test_mail_group@xxx.xx" | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `user_id_type` | `string` | No | User ID categories **Example value**: open_id **Optional values are**:  - `open_id`: Identifies a user to an app. The same user has different Open IDs in different apps. How to get Open ID - `union_id`: Identifies a user to a tenant that acts as a developer. A user has the same Union ID in apps developed by the same developer, and has different Union IDs in apps developed by different developers. A developer can use Union ID to link the same user's identities in multiple apps.How to get Union ID - `user_id`: Identifies a user to a tenant. The same user has different User IDs in different tenants. In one single tenant, a user has the same User ID in all apps （including store apps）. User ID is usually used to communicate user data between different apps. How to get User ID  **Default value**: `open_id` **When the value is `user_id`, the following field scopes are required**: `contact:user.employee_id:readonly` |
| `department_id_type` | `string` | No | The type of department ID used in this call **Example value**: open_department_id **Optional values are**:  - `department_id`: Identify departments with custom department_id - `open_department_id`: Identify departments by open_department_id  | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `items` | `mailgroup.member[]` | No | List of mailing group members added this time **Data validation rules**: - Length range: `1` ～ `200` |
|   `member_id` | `string` | No | Unique identifier of the member in the mailing list（This is not required in the request body） **Example value**: "xxxxxxxxxxxxxxx" |
|   `email` | `string` | No | Member's email address (this field has a value when member type is EXTERNAL_USER/MAIL_GROUP/OTHER_MEMBER) **Example value**: "test_memeber@xxx.xx" |
|   `user_id` | `string` | No | Unique identifier of the user in the tenant (this field has a value when member type is USER) **Example value**: "xxxxxxxxxx" |
|   `department_id` | `string` | No | Unique identifier of the department in the tenant (this field has a value when member type is DEPARTMENT) **Example value**: "xxxxxxxxxx" |
|   `type` | `string` | No | Member type **Example value**: "USER" **Optional values are**:  - `USER`: Internal user - `DEPARTMENT`: Department - `COMPANY`: All staff - `EXTERNAL_USER`: External user - `MAIL_GROUP`: Mailing list - `PUBLIC_MAILBOX`: Public Mailbox - `OTHER_MEMBER`: Internal member  | ### Request body example

{
  "items": [
    {
      "user_id": "xxxxxxxxxx",
      "type": "USER"
    },
    {
      "department_id": "xxxxxxxxxx",
      "type": "DEPARTMENT"
    },
    {
      "type": "COMPANY"
    },
    {
      "email": "test_memeber@xxx.xx",
      "type": "MAIL_GROUP"
    },
    {
      "email": "test_memeber@yyy.yy",
      "type": "EXTERNAL_USER"
    }
  ]
}

## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `items` | `mailgroup.member[]` | Add the mailing group member information list after success |
|     `member_id` | `string` | Unique identifier of the member in the mailing list（This is not required in the request body） |
|     `email` | `string` | Member's email address (this field has a value when member type is EXTERNAL_USER/MAIL_GROUP/OTHER_MEMBER) |
|     `user_id` | `string` | Unique identifier of the user in the tenant (this field has a value when member type is USER) |
|     `department_id` | `string` | Unique identifier of the department in the tenant (this field has a value when member type is DEPARTMENT) |
|     `type` | `string` | Member type **Optional values are**:  - `USER`: Internal user - `DEPARTMENT`: Department - `COMPANY`: All staff - `EXTERNAL_USER`: External user - `MAIL_GROUP`: Mailing list - `PUBLIC_MAILBOX`: Public Mailbox - `OTHER_MEMBER`: Internal member  | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
        "items": [
            {
                "member_id": "xxxxxxxxxxxxxxx",
                "email": "test_memeber@xxx.xx",
                "user_id": "xxxxxxxxxx",
                "department_id": "xxxxxxxxxx",
                "type": "USER"
            }
        ]
    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 404 | 1234013 | mail group not found | Check whether the mailing list exists. |
| 400 | 1234008 | request parameter error | Check whether the request parameters are correct. |
| 400 | 1234030 | mailing group member count is over max limit | The number of mail group members has reached the maximum number limit |
