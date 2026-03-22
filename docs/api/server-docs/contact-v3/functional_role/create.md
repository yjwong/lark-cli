---
title: "Create role"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/functional_role/create"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/contact/v3/functional_roles"
service: "contact-v3"
resource: "functional_role"
section: "Contacts"
rate_limit: "100 per minute"
scopes:
  - "contact:functional_role"
updated: "1684145710000"
---

# Create role

Through the "Create Role" interface, role creation can be completed in batches, and new roles can be displayed synchronously to the tenant's management background - role management module.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/contact/v3/functional_roles |
| HTTP Method | POST |
| Rate Limit | 100 per minute |
| Supported app types | custom |
| Required scopes | `contact:functional_role` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `role_name` | `string` | Yes | Role name, unique under single tenant **Example value**: "考勤管理员" **Data validation rules**: - Length range: `1` ～ `50` characters | ### Request body example

{
    "role_name": "考勤管理员"
}

## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `role_id` | `string` | Role ID, unique under single tenant | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
        "role_id": "7vrj3vk70xk7v5r"
    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 41201 | role name is existed | Role names cannot be repeated within a tenant |
| 400 | 41208 | tenant role not more 500 | The number of roles in the tenant exceeds 500 |
