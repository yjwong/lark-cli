---
title: "Create A Mail Group Permission Member"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/mail-v1/mailgroup-permission_member/create"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/mail/v1/mailgroups/:mailgroup_id/permission_members"
service: "mail-v1"
resource: "mailgroup-permission_member"
section: "Email"
rate_limit: "50 per second"
scopes:
  - "mail:mailgroup"
field_scopes:
  - "contact:user.employee_id:readonly"
updated: "1745840341000"
---

# Create a mail group permission member

Adds a member with custom permission to the mailing list, who will then be able to send emails to this mailing list.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/mail/v1/mailgroups/:mailgroup_id/permission_members |
| HTTP Method | POST |
| Rate Limit | 50 per second |
| Supported app types | custom |
| Required scopes | `mail:mailgroup` |
| Required field scopes | > The response body of the API contains the following sensitive fields, and they will be returned only after corresponding scopes are added. If you do not need the fields, it is not recommended that you request the scopes. `contact:user.employee_id:readonly` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `mailgroup_id` | `string` | Mailing list ID or mailing list address **Example value**: "xxxxxxxxxxxxxxx or test_mail_group@xxx.xx" | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `user_id_type` | `string` | No | User ID categories **Example value**: open_id **Optional values are**:  - `open_id`: Identifies a user to an app. The same user has different Open IDs in different apps. How to get Open ID - `union_id`: Identifies a user to a tenant that acts as a developer. A user has the same Union ID in apps developed by the same developer, and has different Union IDs in apps developed by different developers. A developer can use Union ID to link the same user's identities in multiple apps.How to get Union ID - `user_id`: Identifies a user to a tenant. The same user has different User IDs in different tenants. In one single tenant, a user has the same User ID in all apps （including store apps）. User ID is usually used to communicate user data between different apps. How to get User ID  **Default value**: `open_id` **When the value is `user_id`, the following field scopes are required**: `contact:user.employee_id:readonly` |
| `department_id_type` | `string` | No | Type of department ID used in this call **Example value**: xxx **Optional values are**:  - `department_id`: Identify the department with the custom department_id - `open_department_id`: Identify the department with open_department_id  | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `user_id` | `string` | No | Unique identifier of the user in the tenant (this field has a value when member type is USER) **Example value**: "xxxxxxxxxx" |
| `department_id` | `string` | No | Unique identifier of the department in the tenant (this field has a value when member type is DEPARTMENT) **Example value**: "xxxxxxxxxx" |
| `email` | `string` | No | The member's email address ( this field has a value when member type is MAIL_GROUP/PUBLIC_MAILBOX) **Example value**: "xxx@xx.x" |
| `type` | `string` | No | Member type **Example value**: "USER" **Optional values are**:  - `USER`: Internal user - `DEPARTMENT`: Department - `MAIL_GROUP`: Mail Group - `PUBLIC_MAILBOX`: Public MailBox  | ### Request body example

{
    "user_id": "xxxxxxxxxx",
    "department_id": "xxxxxxxxxx",
    "email": "xxx@xx.x",
    "type": "USER"
}

## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `mailgroup.permission_member` | \- |
|   `permission_member_id` | `string` | The unique identity of the member in the permission group (do not need to be filled in in the request body) |
|   `user_id` | `string` | Unique identifier of the user in the tenant (this field has a value when member type is USER) |
|   `department_id` | `string` | Unique identifier of the department in the tenant (this field has a value when member type is DEPARTMENT) |
|   `email` | `string` | The member's email address ( this field has a value when member type is MAIL_GROUP/PUBLIC_MAILBOX) |
|   `type` | `string` | Member type **Optional values are**:  - `USER`: Internal user - `DEPARTMENT`: Department - `MAIL_GROUP`: Mail Group - `PUBLIC_MAILBOX`: Public MailBox  | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
        "permission_member_id": "xxxxxxxxxxxxxxx",
        "user_id": "xxxxxxxxxx",
        "department_id": "xxxxxxxxxx",
        "email": "xxx@xx.x",
        "type": "USER"
    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 404 | 1234013 | mail group not found | Check whether the mailing list exists. |
