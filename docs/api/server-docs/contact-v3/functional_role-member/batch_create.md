---
title: "Batch create role members"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/functional_role-member/batch_create"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/contact/v3/functional_roles/:role_id/members/batch_create"
service: "contact-v3"
resource: "functional_role-member"
section: "Contacts"
rate_limit: "100 per minute"
scopes:
  - "contact:functional_role"
field_scopes:
  - "contact:user.employee_id:readonly"
updated: "1685539097000"
---

# Batch create role members

Members can be added in batches through the "Add role members in batches" interface, and the member information is synchronously displayed to the tenant's management background - role management module.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/contact/v3/functional_roles/:role_id/members/batch_create |
| HTTP Method | POST |
| Rate Limit | 100 per minute |
| Supported app types | custom |
| Required scopes | `contact:functional_role` |
| Required field scopes | > The response body of the API contains the following sensitive fields, and they will be returned only after corresponding scopes are added. If you do not need the fields, it is not recommended that you request the scopes. `contact:user.employee_id:readonly` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `role_id` | `string` | Unique identity of the role, unique under a single tenant **Example value**: "7vrj3vk70xk7v5r" | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `user_id_type` | `string` | No | User ID categories **Example value**: "open_id" **Optional values are**:  - `open_id`: Identifies a user to an app. The same user has different Open IDs in different apps. How to get Open ID - `union_id`: Identifies a user to a tenant that acts as a developer. A user has the same Union ID in apps developed by the same developer, and has different Union IDs in apps developed by different developers. A developer can use Union ID to link the same user's identities in multiple apps.How to get Union ID - `user_id`: Identifies a user to a tenant. The same user has different User IDs in different tenants. In one single tenant, a user has the same User ID in all apps （including store apps）. User ID is usually used to communicate user data between different apps. How to get User ID  **Default value**: `open_id` **When the value is `user_id`, the following field scopes are required**: `contact:user.employee_id:readonly` | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `members` | `string[]` | Yes | List of role members added by the role (UserID list of a batch of users) **Example value**: ["ou-12832197382"] **Data validation rules**: - Length range: `1` ～ `100` | ### Request body example

{
    "members": [
        "ou-12832197382"
    ]
}

## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `results` | `functional_role_member_result[]` | Batch new role member result set |
|     `user_id` | `string` | User ID |
|     `reason` | `int` | Member processing results **Optional values are**:  - `1`: Processing successful - `2`: Invalid user ID - `3`: User ID No permission - `4`: User already exists in this role - `5`: User does not exist in this role - `6`: No permission to the user's old administrative scope in this role  | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
        "results": [
            {
                "user_id": "od-4e6ac4d14bcd5071a37a39de902c7141",
                "reason": 1
            }
        ]
    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 404 | 41202 | role id is not exist | Role ID does not exist |
| 400 | 41209 | tenant role is not more 1000 | The number of members of a single role within a tenant cannot exceed 1000 |
