---
title: "Query all member information under the role"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/functional_role-member/list"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/contact/v3/functional_roles/:role_id/members"
service: "contact-v3"
resource: "functional_role-member"
section: "Contacts"
rate_limit: "100 per minute"
scopes:
  - "contact:functional_role"
  - "contact:functional_role:readonly"
field_scopes:
  - "contact:user.employee_id:readonly"
updated: "1684145710000"
---

# Query all member information under the role

Through this interface, you can query the member information under the role ID (including the member ID and its management scope).

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/contact/v3/functional_roles/:role_id/members |
| HTTP Method | GET |
| Rate Limit | 100 per minute |
| Supported app types | custom |
| Required scopes Enable any scope from the list | `contact:functional_role` `contact:functional_role:readonly` |
| Required field scopes | > The response body of the API contains the following sensitive fields, and they will be returned only after corresponding scopes are added. If you do not need the fields, it is not recommended that you request the scopes. `contact:user.employee_id:readonly` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `role_id` | `string` | Unique identity of the role, unique under a single tenant **Example value**: "7vrj3vk70xk7v5r" | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `page_size` | `int` | No | Page size **Example value**: 50 **Default value**: `20` **Data validation rules**: - Value range: `0` ～ `100` |
| `page_token` | `string` | No | Page identifier. It is not filled in the first request, indicating traversal from the beginning; when there will be more groups, the new page_token will be returned at the same time, and the next traversal can use the page_token to get more groups **Example value**: "dawdewd" |
| `user_id_type` | `string` | No | User ID categories **Example value**: "open_id" **Optional values are**:  - `open_id`: Identifies a user to an app. The same user has different Open IDs in different apps. How to get Open ID - `union_id`: Identifies a user to a tenant that acts as a developer. A user has the same Union ID in apps developed by the same developer, and has different Union IDs in apps developed by different developers. A developer can use Union ID to link the same user's identities in multiple apps.How to get Union ID - `user_id`: Identifies a user to a tenant. The same user has different User IDs in different tenants. In one single tenant, a user has the same User ID in all apps （including store apps）. User ID is usually used to communicate user data between different apps. How to get User ID  **Default value**: `open_id` **When the value is `user_id`, the following field scopes are required**: `contact:user.employee_id:readonly` |
| `department_id_type` | `string` | No | The type of department ID used in this call **Example value**: "open_department_id" **Optional values are**:  - `department_id`: Identify departments with custom department_id - `open_department_id`: Identify departments by open_department_id  **Default value**: `open_department_id` | ## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `members` | `functional_role.member[]` | Role member list |
|     `user_id` | `string` | Member ID |
|     `scope_type` | `string` | Types of management scope **Optional values are**:  - `All`: Management scope is all - `Part`: Management scope is part - `None`: Management scope is empty  |
|     `department_ids` | `string[]` | Represents the administrative scope of a member of this role, which is returned scope_type "Specified Scope" |
|   `page_token` | `string` | Page identifier, when has_more is true, a new page_token will also be returned. Otherwise, page_token will not be returned |
|   `has_more` | `boolean` | Whether the response body has more parameters | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
        "members": [
            {
                "user_id": "od-4e6ac4d14bcd5071a37a39de902c7141",
                "scope_type": "All",
                "department_ids": [
                    "ADASF"
                ]
            }
        ],
        "page_token": "2132323",
        "has_more": false
    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 40001 | param error | Parameter error. |
| 404 | 41202 | role id is not exist | Role ID does not exist |
