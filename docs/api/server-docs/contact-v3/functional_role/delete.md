---
title: "Delete role"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/functional_role/delete"
method: "DELETE"
api_path: "https://open.larksuite.com/open-apis/contact/v3/functional_roles/:role_id"
service: "contact-v3"
resource: "functional_role"
section: "Contacts"
rate_limit: "100 per minute"
scopes:
  - "contact:functional_role"
updated: "1684145710000"
---

# Delete role

A role can be deleted through this interface

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/contact/v3/functional_roles/:role_id |
| HTTP Method | DELETE |
| Rate Limit | 100 per minute |
| Supported app types | custom |
| Required scopes | `contact:functional_role` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `role_id` | `string` | Unique identity of the role, unique under a single tenant **Example value**: "7vrj3vk70xk7v5r" | ## Response

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
| 404 | 41202 | role id is not exist | Role ID does not exist |
| 400 | 41206 | functional role has members, can not delete | Roles contain members and cannot be deleted |
