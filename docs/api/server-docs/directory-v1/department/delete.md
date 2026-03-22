---
title: "Delete department"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/directory-v1/department/delete"
method: "DELETE"
api_path: "https://open.larksuite.com/open-apis/directory/v1/departments/:department_id"
service: "directory-v1"
resource: "department"
section: "Organization"
rate_limit: "10 per second"
scopes:
  - "directory:department.delete:write"
  - "directory:department:write"
field_scopes:
  - "directory:employee.base.external_id:read"
updated: "1756439294000"
---

# Delete department

This interface is used to delete departments.

> Note:
> - Deleting a department requires application data permissions for the department to be deleted and its parent department Configure application data permissions

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/directory/v1/departments/:department_id |
| HTTP Method | DELETE |
| Rate Limit | 10 per second |
| Supported app types | custom |
| Required scopes Enable any scope from the list | `directory:department.delete:write` `directory:department:write` |
| Required field scopes | > The response body of the API contains the following sensitive fields, and they will be returned only after corresponding scopes are added. If you do not need the fields, it is not recommended that you request the scopes. `directory:employee.base.external_id:read` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `department_id` | `string` | Department ID, consistent with the department_id_type type **Example value**: "weasdqwe" **Data validation rules**: - Maximum length: `64` characters | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `department_id_type` | `string` | No | The type of department ID used in this call **Example value**: open_department_id **Optional values are**:  - `department_id`: Used to identify a unique department within a tenant - `open_department_id`: It is used to identify a department in a specific application, and the same department has different open_department_id in different applications.  **Default value**: `open_department_id` | ## Response

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
| 400 | 2221332 | Department has members, can not delete department | There are still current and pending employees under the department, and the department cannot be deleted. |
| 400 | 2224001 | No permission to operate | No operation permission. Please check the permission of the current application or whether the enterprise version is the business professional version or above. |
| 400 | 2224002 | No permission to operate records | You do not have permission to operate this record. Please check the permission of the data management scope of the current application or whether the department operated by the current application can be deleted. |
| 400 | 2224003 | No permission to operate dependent objects | No permission to operate the dependent object. Please check whether you have permission to delete the department. |
| 400 | 2221348 | Department has children, can not delete department | There are sub-departments under the department, please move or delete the sub-department first |
