---
title: "Update un-resigned employees to be resigned"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/directory-v1/employee/to_be_resigned"
method: "PATCH"
api_path: "https://open.larksuite.com/open-apis/directory/v1/employees/:employee_id/to_be_resigned"
service: "directory-v1"
resource: "employee"
section: "Organization"
rate_limit: "10 per minute"
scopes:
  - "directory:employee.to_be_resigned:write"
field_scopes:
  - "directory:employee.base.external_id:read"
updated: "1756439261000"
---

# Update unresigned employees to pre-resigned

This interface is used to handle the departure of current employees and update them to the status of "pending departure". "pending departure" employees will not automatically leave, and need to use the "delete employee" API to terminate and transfer resources.
When using user_access_token, it defaults to an administrator user, and only administrators who have the "CoreHR Management" role can operate it.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/directory/v1/employees/:employee_id/to_be_resigned |
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
| `employee_id` | `string` | Employee ID,consistent with the employee_id_type type **Example value**: "cad2cafa" **Data validation rules**: - Length range: `1` ～ `255` characters | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `employee_id_type` | `string` | No | User ID categories **Example value**: open_id **Optional values are**:  - `open_id`: Identifies a user to an app. The same user has different Open IDs in different apps. How to get Open ID - `union_id`: Identifies a user to a tenant that acts as a developer. A user has the same Union ID in apps developed by the same developer, and has different Union IDs in apps developed by different developers. A developer can use Union ID to link the same user's identities in multiple apps.How to get Union ID - `employee_id`: The unique identifier of an employee within the enterprise. Supports customization, the system automatically generates it if it is not customized. ID supports modification. How to get employee_id: - Enterprise administrators go to the Admin Console > Organizational Structure > Members and Departments page, click Member Details to query employee IDs - Through the interface of Get employee list in batches, query employee ID by mobile phone number or email.  **Default value**: `open_id` **When the value is `employee_id`, the following field scopes are required**: `directory:employee.base.external_id:read` |
| `department_id_type` | `string` | No | Department ID type **Example value**: open_department_id **Optional values are**:  - `department_id`: Used to identify a unique department within the tenant - `open_department_id`: Used to identify a department in a specific application. The same department has the same open_department_id in different applications.  **Default value**: `open_department_id` | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `employee` | `set_employee_pre_resigned` | Yes | Current employees transfer to pending departure |
|   `resign_date` | `string` | Yes | date of separation **Example value**: "2024-06-21" |
|   `resign_reason` | `string` | Yes | reason for leaving **Example value**: "1" **Optional values are**:  - `1`: Salary does not meet expectations - `2`: Working too long - `3`: Dissatisfied with the job content - `4`: Do not recognize superiors or management - `5`: Career development opportunities are limited - `6`: Lack of recognition of company culture - `7`: Organizational restructuring (voluntary resignation) - `8`: Contract expires - `9`: Job hopping - `10`: Career change - `11`: Family reasons - `12`: Poor health - `13`: Workplace reasons - `14`: Other (voluntary resignation) - `15`: accident - `16`: passed away - `17`: dismiss - `18`: Probation period not passed - `19`: Underperforming at work - `20`: Low work output - `21`: Organizational restructuring (passive departure) - `22`: discipline violation - `23`: illegal - `24`: Other (passive resignation) - `25`: Other (other)  |
|   `resign_type` | `string` | Yes | type of turnover **Example value**: "1" **Optional values are**:  - `1`: active - `2`: passive - `3`: other  |
|   `resign_remark` | `string` | No | Resignation remarks **Example value**: "study abroad" **Data validation rules**: - Length range: `0` ～ `255` characters | ### Request body example

{
    "employee": {
        "resign_date": "2024-06-21",
        "resign_reason": "1",
        "resign_type": "1",
        "resign_remark": "study abroad"
    }
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
    "msg": "Success",
    "data": {}
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 2221246 | User staff status not updated | The user status has not been updated and is pending departure. |
| 400 | 2221213 | Resign date invalid or earlier than join date or empty | Please check that the date is not empty and formatted correctly, or if the departure date is earlier than the onboard date |
| 400 | 2221214 | Resigned reason invalid or not match resigned type | Reasons for leaving and type of leaving do not match |
| 400 | 2221231 | Resign type invalid or not match resign reason | Reasons for leaving and type of leaving do not match |
