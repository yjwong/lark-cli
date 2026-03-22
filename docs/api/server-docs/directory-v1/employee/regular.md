---
title: "Update pre-resigned members to un-resigned employees"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/directory-v1/employee/regular"
method: "PATCH"
api_path: "https://open.larksuite.com/open-apis/directory/v1/employees/:employee_id/regular"
service: "directory-v1"
resource: "employee"
section: "Organization"
rate_limit: "10 per minute"
scopes:
  - "directory:employee.to_be_resigned:write"
field_scopes:
  - "directory:employee.base.external_id:read"
updated: "1756439265000"
---

# Update pre-resigned employee to unresigned

This interface is used to cancel the departure of the pending employee and update it to the "on-the-job" status. When canceling the departure, the departure information will be cleared.
When using user_access_token, it defaults to an administrator user, and only administrators who can has "CoreHR Management" role can operate it.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/directory/v1/employees/:employee_id/regular |
| HTTP Method | PATCH |
| Rate Limit | 10 per minute |
| Supported app types | custom |
| Required scopes | `directory:employee.to_be_resigned:write` |
| Required field scopes | > The response body of the API contains the following sensitive fields, and they will be returned only after corresponding scopes are added. If you do not need the fields, it is not recommended that you request the scopes. `directory:employee.base.external_id:read` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `employee_id` | `string` | Employee ID, consistent with the employee_id_type type **Example value**: "d2e1jas" **Data validation rules**: - Length range: `1` ～ `255` characters | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `employee_id_type` | `string` | No | User ID categories **Example value**: open_id **Optional values are**:  - `open_id`: Identifies a user to an app. The same user has different Open IDs in different apps. How to get Open ID - `union_id`: Identifies a user to a tenant that acts as a developer. A user has the same Union ID in apps developed by the same developer, and has different Union IDs in apps developed by different developers. A developer can use Union ID to link the same user's identities in multiple apps.How to get Union ID - `employee_id`: The unique identifier of an employee within the enterprise. Supports customization, the system automatically generates it if it is not customized. ID supports modification. How to get employee_id: - Enterprise administrators go to the Admin Console > Organizational Structure > Members and Departments page, click Member Details to query employee IDs - Through the interface of Get employee list in batches, query employee ID by mobile phone number or email.  **Default value**: `open_id` **When the value is `employee_id`, the following field scopes are required**: `directory:employee.base.external_id:read` |
| `department_id_type` | `string` | No | Department ID type **Example value**: open_department_id **Optional values are**:  - `department_id`: Used to identify a unique department within the tenant - `open_department_id`: Used to identify a department in a specific application. The same department has the same open_department_id in different applications.  **Default value**: `open_department_id` | ### Request body example

{}

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
| 400 | 2221246 | User staff status not updated | User staff status not updated |
