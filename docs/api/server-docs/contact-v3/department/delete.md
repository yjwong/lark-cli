---
title: "Delete Department"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/department/delete"
method: "DELETE"
api_path: "https://open.larksuite.com/open-apis/contact/v3/departments/:department_id"
service: "contact-v3"
resource: "department"
section: "Contacts"
scopes:
  - "contact:contact"
updated: "1683879166000"
---

# Delete a department

This API is used to delete a department from contacts. FAQs.

> The app must have the contact scopes of both the department to be deleted and its parent department. Store apps have no scope to call this API.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/contact/v3/departments/:department_id |
| HTTP Method | DELETE |
| Supported app types | custom |
| Required scopes | `contact:contact` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" Learn more about obtaining and using access_token. | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `department_id` | `string` | Department ID, which must be consistent with department_id_type in the query parameter. **Example value**: "od-4e6ac4d14bcd5071a37a39de902c7141" **Data validation rules**: - Maximum length: `64` characters - Regular expression: `^0|[^od][A-Za-z0-9]*` | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `department_id_type` | `string` | No | Type of department ID used in this call **Example value**: "open_department_id" **Optional values are**:  - `department_id`: Identify the department with the custom department_id - `open_department_id`: Open departmentID  **Default value**: `open_department_id` | ## Response

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
| 400 | 43011 | delete has member dept error | The department to be deleted cannot contain any user. Check whether there is any user in the department. |
| 400 | 43012 | delete has sub dept department error | The department to be deleted cannot contain any sub-department. Check whether there is any sub-department in the department. |
| 400 | 40002 | process root dept error | Root departments cannot deleted. |
| 400 | 40003 | internal error | Internal error. Please provide the X-Request-Id to our agent. [Contact support](https://applink.larksuite.com/client/helpdesk/open?id |
| 403 | 40004 | no dept authority error | The department must be within the range of contacts data that the app can access. Learn more |
| 403 | 40014 | no parent dept authority error | No access to the parent department(visible scope of organization). |
| 401 | 42008 | tenant id is invalid error | Check whether the tenant is valid. |
