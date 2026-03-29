---
title: "reset user enterprise email password"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/admin-v1/password/reset"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/admin/v1/password/reset"
service: "admin-v1"
resource: "password"
section: "Admin"
scopes:
  - "admin:ent_email_password"
field_scopes:
  - "contact:user.employee_id:readonly"
updated: "1652772927000"
---

# Password reset API

A user's business email password can only be reset if the user’s email address and business email address (alias) are the same.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/admin/v1/password/reset |
| HTTP Method | POST |
| Supported app types | custom |
| Required scopes | `admin:ent_email_password` |
| Required field scopes | > The response body of the API contains the following sensitive fields, and they will be returned only after corresponding scopes are added. If you do not need the fields, it is not recommended that you request the scopes. `contact:user.employee_id:readonly` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" Learn more about obtaining and using access_token. |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `user_id_type` | `string` | Yes | User ID categories **Example value**: "open_id" **Optional values are**:  - `open_id`: user's open id - `union_id`: user's union id - `user_id`: user's user id  **Default value**: `open_id` **When the value is `user_id`, the following field scopes are required**: `contact:user.employee_id:readonly` | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `password` | `password` | Yes | Password reset parameters: the password should contain at least 8 characters, including at least 2 character types (letters, numbers and symbols). |
|   `ent_email_password` | `string` | Yes | Enterprise Email Password **Example value**: "abcd*efg" |
| `user_id` | `string` | Yes | The user ID password can only be reset if the user’s login email address and business email address (including aliases) are the same.  Data validation rule: Length range: 0 to 200 characters **Example value**: "abc123" **Data validation rules**: - Length range: `0` ～ `200` characters | ### Request body example

{
    "password": {
        "ent_email_password": "abcd*efg"
    },
    "user_id": "abc123"
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
    "data": {

    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 1050023 | reach reset password num limit | Please try again the next day |
| 400 | 1050022 | password format is not valid | Please try again after changing the password format |
| 400 | 1050024 | user cant reset password.may be user enterprise email not equal email | The password can only be changed if the user's email address and business email address (alias) are the same. |
