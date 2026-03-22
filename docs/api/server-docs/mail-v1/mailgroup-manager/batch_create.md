---
title: "Create Mailing List Managers In Batch"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/mail-v1/mailgroup-manager/batch_create"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/mail/v1/mailgroups/:mailgroup_id/managers/batch_create"
service: "mail-v1"
resource: "mailgroup-manager"
section: "Email"
rate_limit: "50 per second"
scopes:
  - "mail:mailgroup"
field_scopes:
  - "contact:user.employee_id:readonly"
updated: "1745840340000"
---

# Create mailing list managers in batch

Adds managers to a mailing list.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/mail/v1/mailgroups/:mailgroup_id/managers/batch_create |
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
| `mailgroup_id` | `string` | Mailing list ID or mailing list address **Example value**: "xxxxxx or test_mail_group@xx.xx" **Data validation rules**: - Length range: `1` ～ `255` characters | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `user_id_type` | `string` | No | User ID categories **Example value**: open_id **Optional values are**:  - `open_id`: Identifies a user to an app. The same user has different Open IDs in different apps. How to get Open ID - `union_id`: Identifies a user to a tenant that acts as a developer. A user has the same Union ID in apps developed by the same developer, and has different Union IDs in apps developed by different developers. A developer can use Union ID to link the same user's identities in multiple apps.How to get Union ID - `user_id`: Identifies a user to a tenant. The same user has different User IDs in different tenants. In one single tenant, a user has the same User ID in all apps （including store apps）. User ID is usually used to communicate user data between different apps. How to get User ID  **Default value**: `open_id` **Data validation rules**: - Length range: `0` ～ `256` characters **When the value is `user_id`, the following field scopes are required**: `contact:user.employee_id:readonly` | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `mailgroup_manager_list` | `mailgroup.manager[]` | No | Mail Group Manager List **Data validation rules**: - Length range: `0` ～ `200` |
|   `user_id` | `string` | No | Manager User ID **Example value**: "xxxxxx" **Data validation rules**: - Length range: `1` ～ `1024` characters | ### Request body example

{
    "mailgroup_manager_list": [
        {
            "user_id": "xxxxxx"
        }
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
| 200 | 1234013 | mail group not found | Check whether the mailing list exists. |
