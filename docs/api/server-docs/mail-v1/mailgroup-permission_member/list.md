---
title: "List Permission Members Of A Mail Group"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/mail-v1/mailgroup-permission_member/list"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/mail/v1/mailgroups/:mailgroup_id/permission_members"
service: "mail-v1"
resource: "mailgroup-permission_member"
section: "Email"
rate_limit: "Special Rate Limit"
scopes:
  - "mail:mailgroup"
  - "mail:mailgroup:readonly"
field_scopes:
  - "contact:user.employee_id:readonly"
updated: "1745840341000"
---

# List permission members of a mail group

Obtains the list of members who can send emails to mailing list addresses by pages.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/mail/v1/mailgroups/:mailgroup_id/permission_members |
| HTTP Method | GET |
| Rate Limit | Special Rate Limit |
| Supported app types | custom |
| Required scopes Enable any scope from the list | `mail:mailgroup` `mail:mailgroup:readonly` |
| Required field scopes | > The response body of the API contains the following sensitive fields, and they will be returned only after corresponding scopes are added. If you do not need the fields, it is not recommended that you request the scopes. `contact:user.employee_id:readonly` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `mailgroup_id` | `string` | Mailing list ID or mailing list address **Example value**: "xxxxxxxxxxxxxxx or test_mail_group@xxx.xx" | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `user_id_type` | `string` | No | User ID categories **Example value**: open_id **Optional values are**:  - `open_id`: Identifies a user to an app. The same user has different Open IDs in different apps. How to get Open ID - `union_id`: Identifies a user to a tenant that acts as a developer. A user has the same Union ID in apps developed by the same developer, and has different Union IDs in apps developed by different developers. A developer can use Union ID to link the same user's identities in multiple apps.How to get Union ID - `user_id`: Identifies a user to a tenant. The same user has different User IDs in different tenants. In one single tenant, a user has the same User ID in all apps （including store apps）. User ID is usually used to communicate user data between different apps. How to get User ID  **Default value**: `open_id` **When the value is `user_id`, the following field scopes are required**: `contact:user.employee_id:readonly` |
| `department_id_type` | `string` | No | Type of department ID used in this call **Example value**: open_department_id **Optional values are**:  - `department_id`: Identify the department with the custom department_id - `open_department_id`: Identify the department with open_department_id  |
| `page_token` | `string` | No | Page identifier. It is not filled in the first request, indicating traversal from the beginning; when there will be more groups, the new page_token will be returned at the same time, and the next traversal can use the page_token to get more groups **Example value**: xxx |
| `page_size` | `int` | No | **Example value**: 10 **Default value**: `20` **Data validation rules**: - Maximum value: `200` | ## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `has_more` | `boolean` | Whether the response body has more parameters |
|   `page_token` | `string` | Page identifier, when has_more is true, a new page_token will also be returned. Otherwise, page_token will not be returned |
|   `items` | `mailgroup.permission_member[]` | List of members who can send emails to mailing list addresses |
|     `permission_member_id` | `string` | The unique identity of the member in the permission group (do not need to be filled in in the request body) |
|     `user_id` | `string` | Unique identifier of the user in the tenant (this field has a value when member type is USER) |
|     `department_id` | `string` | Unique identifier of the department in the tenant (this field has a value when member type is DEPARTMENT) |
|     `email` | `string` | The member's email address ( this field has a value when member type is MAIL_GROUP/PUBLIC_MAILBOX) |
|     `type` | `string` | Member type **Optional values are**:  - `USER`: Internal user - `DEPARTMENT`: Department - `MAIL_GROUP`: Mail Group - `PUBLIC_MAILBOX`: Public MailBox  | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
        "has_more": true,
        "page_token": "xxx",
        "items": [
            {
                "permission_member_id": "xxxxxxxxxxxxxxx",
                "user_id": "xxxxxxxxxx",
                "department_id": "xxxxxxxxxx",
                "email": "xxx@xx.x",
                "type": "USER"
            }
        ]
    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 404 | 1234013 | mail group not found | Check whether the mailing list exists. |
