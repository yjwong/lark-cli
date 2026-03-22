---
title: "Update DepartmentID"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/department/update_department_id"
method: "PATCH"
api_path: "https://open.larksuite.com/open-apis/contact/v3/departments/:department_id/update_department_id"
service: "contact-v3"
resource: "department"
section: "Contacts"
rate_limit: "1000 per minute、50 per second"
scopes:
  - "contact:contact:update_department_id"
updated: "1699521193000"
---

# Update DepartmentID

This API allows the user to update the department ID(department_id). The new department ID(department_id) needs to be unoccupied within the enterprise.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/contact/v3/departments/:department_id/update_department_id |
| HTTP Method | PATCH |
| Rate Limit | 1000 per minute、50 per second |
| Supported app types | custom |
| Required scopes | `contact:contact:update_department_id` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `department_id` | `string` | The department ID that needs to be updated, the ID type is the same as in the department_id_type **Example value**: "od-d6b83d25c129775723a36f52495c4f81" **Data validation rules**: - Maximum length: `64` characters - Regular expression: `^[a-zA-Z0-9][a-zA-Z0-9_\-@.]{0,63}$` | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `department_id_type` | `string` | No | The type of department ID used in this request **Example value**: open_department_id **Optional values are**:  - `department_id`: Identify departments with custom department_id - `open_department_id`: Identify departments by open_department_id  | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `new_department_id` | `string` | Yes | New department custom ID(department_id) **Example value**: "NewDevDepartID" **Data validation rules**: - Maximum length: `128` characters - Regular expression: `^0|[^od][A-Za-z0-9]*` | ### Request body example

{
    "new_department_id": "NewDevDepartID"
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
| 400 | 40001 | param error | Parameter error, please modify the interface documentation and try again. |
| 400 | 43009 | exceed update custom dept limit error | Maximum update limit exceeded. |
| 403 | 40007 | can not update dept custom id error | Unable to update department ID. Please confirm whether you have permission to operate the corresponding department. |
| 401 | 42008 | tenant id is invalid error | Tenant ID is invalid, please check and try again |
