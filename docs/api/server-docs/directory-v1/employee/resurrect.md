---
title: "Reinstate departed employees"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/directory-v1/employee/resurrect"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/directory/v1/employees/:employee_id/resurrect"
service: "directory-v1"
resource: "employee"
section: "Organization"
rate_limit: "10 per minute"
scopes:
  - "directory:employee.resurrect:write"
  - "directory:employee:write"
field_scopes:
  - "directory:employee.base.external_id:read"
updated: "1756439257000"
---

# Reinstatement of resigned employees

This interface is used to restore the members who have left the company and restore them to the working state. 

> Attention:
> - To restore the resigned employee as in-service, the enterprise version needs to be in the commercial professional version and above,You can view the current version of the enterprise through the management background> settings> version information , and the employee needs to leave within 30 days. After recovery, some user data is still unrecoverable, please call with caution.
> - The user ID of the member to be restored cannot be used by other members in the enterprise.You can query whether the User Identification exists through the Batch Get Employee List interface . If there is any duplication, please leave the corresponding member first, otherwise the interface will report an error.
> - The mobile phone number and email address of the member to be restored cannot be used by other members in the enterprise.You can query whether the mobile phone number/email exists through the Batch Get Employee List interface . If there is any duplication, please modify the information of the corresponding member first, otherwise the interface will report an error.
> - This interface supports tenant_access_token and user_access_token.For the two ways to obtain tokens, please refer to Obtain Access Credentials.
> - When using tenant_access_token, only restore the departing employee to the department within the current application address book authorization.
> - When using user_access_token, the default is the administrator user, and the administrator management scope will be verified. When the user has multiple administrator identities, the administrator management scope can be restored to the maximum set.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/directory/v1/employees/:employee_id/resurrect |
| HTTP Method | POST |
| Rate Limit | 10 per minute |
| Supported app types | custom |
| Required scopes Enable any scope from the list | `directory:employee.resurrect:write` `directory:employee:write` |
| Required field scopes | > The response body of the API contains the following sensitive fields, and they will be returned only after corresponding scopes are added. If you do not need the fields, it is not recommended that you request the scopes. `directory:employee.base.external_id:read` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `employee_id` | `string` | Employee ID, consistent with the employee_id_type type **Example value**: "Eedasdas" | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `employee_id_type` | `string` | No | User ID categories **Example value**: open_id **Optional values are**:  - `open_id`: Identifies a user's identity in an app. The same user has different Open IDs in different apps. Learn more: How to get an Open ID - `union_id`: Identifies a user's identity under an app developer. The Union ID is the same for the same user in apps under the same developer, and different for apps under different developers. With Union ID, app developers can associate the identities of the same user across multiple apps. Read more: How do I get a Union ID? - `employee_id`: The unique identifier of an employee within the enterprise. Supports customization, the system automatically generates it if it is not customized. ID supports modification. How to get employee_id: - Enterprise administrators go to the Admin Console > Organizational Structure > Members and Departments page, click Member Details to query employee IDs - Through the interface of Get employee list in batches, query employee ID by mobile phone number or email.  **Default value**: `open_id` **When the value is `employee_id`, the following field scopes are required**: `directory:employee.base.external_id:read` |
| `department_id_type` | `string` | No | Department ID type **Example value**: open_department_id **Optional values are**:  - `department_id`: department_id - `open_department_id`: open_department_id  **Default value**: `open_department_id` | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `employee_order_in_departments` | `upsert_user_department_sort_info[]` | No | Department information **Data validation rules**: - Length range: `0` ～ `10` |
|   `department_id` | `string` | No | Department ID, consistent with the department_id_type type **Example value**: "EasdiedQDS" |
|   `order_weight_in_deparment` | `string` | No | The ranking weight of users within the department **Data verification rules:** Length range: 1 to 3 **Example value**: "100" |
|   `order_weight_among_deparments` | `string` | No | User ranking weights across multiple departments **Data verification rules:** Length range: 1 to 3 **Example value**: "20 " |
|   `is_main_department` | `boolean` | No | Whether it is the user's main department (the user can only have one main department, and the ranking weight should be the largest. If you don't fill it in, the department with the first ranking will be used as the main department by default) **Example value**: true |
| `options` | `resurrect_employee_options` | No | option |
|   `subscription_ids` | `string[]` | No | List of seat IDs assigned to employees. Available seat IDs for this tenant can be obtained through the interface below, see [Get seat information] (/ssl: ttdoc/uAjLw4CM/ukTMukTMukTM/tenant-v2/tenant-product_assign_info/query). This field is required when in mixed license mode. **Example value**: ["435456345245"] **Data validation rules**: - Length range: `0` ～ `20` | ### Request body example

{"employee_order_in_departments":[{"department_id":"EasdiedQDS",
"order_weight_in_deparment":"100",
"order_weight_among_deparments":"20

",
"is_main_department":true}],
"options":{"subscription_ids":["435456345245"]}}

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
| 400 | 2221268 | Tenant not in the access list | The tenant cannot restore the employee. Please check whether the tenant is in the access list. |
| 400 | 2221269 | Resurrect user info duplicated | The employee information conflicts with that of the current employees. Please check whether the employee information is the same as that of the current employees. |
| 400 | 2221270 | Can't resurrect user in delete process | The employee is already in the departure process and cannot be restored for the time being. Please try again later. |
| 400 | 2221271 | User exceeds billing seats limit | The seat limit has been exceeded. You cannot restore employees at this time. Please contact your administrator to increase the number of billed seats. |
| 400 | 2221272 | User exceeds tenant limit | The number of employees exceeds the tenant limit. Employees cannot be restored for the time being. Please check the tenant's employee limit and adjust it. |
| 400 | 2221273 | User exceeds certification seats limit | The number of employees exceeds the limit of the certified seats. Please contact the administrator to increase the certified seats. |
| 400 | 2221274 | User exceed feature contact seats limit | The number of employees exceeds the address book seat limit. Please adjust the address book seat limit. |
| 400 | 2221275 | Tenant version not allow resurrection | The tenant version does not allow employee restoration. Please upgrade the tenant version to a version that supports employee restoration. |
| 400 | 2221123 | The user refused to join the tenant and cannot be resurrected | The employee refused to join the tenant and couldn't be restored. Please contact the employee to confirm whether they agree to join the tenant. |
| 400 | 2221140 | Department number exceeds limit | The number of people in the department exceeds the limit and cannot be restored. Please check the department's number limit and adjust it. |
| 400 | 2221252 | Hybrid license tenant prohibits passing empty licenses to create users | Empty license transfer is prohibited for tenants with mixed licenses. Please add seat information. (You can obtain the available seat ID through the Get Enterprise Seat Information Interface) |
| 400 | 2221253 | Designated licenseID is insufficient | Insufficient seats. Please modify the seat information. (You can obtain the available seat IDs through the Get Enterprise Seat Information API |
| 400 | 2221254 | Designated licenseID is invalid | Invalid seat information. Please modify the seat information. (You can obtain the available seat ID through the Get Tenant Seat Information API) |
| 400 | 2224001 | No permission to operate | No operation permission. Please check the permission of the current application or whether the enterprise version is the business professional version or above. |
| 400 | 2224002 | No permission to operate records | You do not have permission to operate this record. Please check the permissions of the data management scope of the current application or whether the members operated by the current application can be restored. |
| 400 | 2224003 | No permission to operate dependent objects | No operation permission for the dependent object. Please check whether the department to be restored has the permission. |
| 400 | 2221248 | Resurrect exceeds seat limit | The number of seats exceeds the limit. Please modify the seat information. (You can obtain the available seat ID through the Get Enterprise Seat Information Interface.) |
| 400 | 2221255 | Main department must be the first | The main department must be first, please modify the sorting information of the department to which it belongs |
