---
title: "Obtain bulk department information"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/department/batch"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/contact/v3/departments/batch"
service: "contact-v3"
resource: "department"
section: "Contacts"
rate_limit: "1000 per minute、50 per second"
scopes:
  - "contact:contact.base:readonly"
field_scopes:
  - "contact:department.organize:readonly"
  - "contact:contact:readonly_as_app"
  - "contact:department.base:readonly"
  - "contact:user.employee_id:readonly"
  - "contact:department.hrbp:readonly"
  - "contact:contact:access_as_app"
  - "contact:contact:readonly"
updated: "1689325063000"
---

# Obtain bulk department information

This API is to obtain the information of bulk departments in Contacts.

> - When using `tenant_access_token`, the application must obtain the access permission of the contacts of the requested department. To access root department information, the application must have an all-staff contact scope.
> - When using `user_access_token`, users must obtain the visibility scope of the requested department. To access root department information, the visibility scope must include all employees.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/contact/v3/departments/batch |
| HTTP Method | GET |
| Rate Limit | 1000 per minute、50 per second |
| Supported app types | custom,isv |
| Required scopes | `contact:contact.base:readonly` |
| Required field scopes | > The response body of the API contains the following sensitive fields, and they will be returned only after corresponding scopes are added. If you do not need the fields, it is not recommended that you request the scopes. `contact:department.organize:readonly` `contact:contact:readonly_as_app` `contact:department.base:readonly` `contact:user.employee_id:readonly` `contact:department.hrbp:readonly` `contact:contact:access_as_app` `contact:contact:readonly` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `department_ids` | `string[]` | Yes | IDs of the departments to be obtained For the description of different IDs and how to obtain them, see Department ID Description To query multiple department IDs at once, you can pass the same parameter name multiple times, and pass different parameter values each time. For example, `https://{url}?department_ids={department_id1}&department_ids={department_id2}` Of which: - `department_ids` is a parameter name that can be passed multiple times - `department_id1` and `department_id2` are parameter values **Example value**: D096 |
| `department_id_type` | `string` | No | Type of department ID used in this call **Example value**: open_department_id **Optional values are**:  - `open_department_id`: used to identify a department in a specific application. The open_department_id of the same department across different applications is the same. - `department_id`: a unique ID used to identify a department within a tenant.  **Default value**: `open_department_id` |
| `user_id_type` | `string` | No | User ID categories **Example value**: open_id **Optional values are**:  - `open_id`: Identifies a user to an app. The same user has different Open IDs in different apps. How to get Open ID - `union_id`: Identifies a user to a tenant that acts as a developer. A user has the same Union ID in apps developed by the same developer, and has different Union IDs in apps developed by different developers. A developer can use Union ID to link the same user's identities in multiple apps.How to get Union ID - `user_id`: Identifies a user to a tenant. The same user has different User IDs in different tenants. In one single tenant, a user has the same User ID in all apps （including store apps）. User ID is usually used to communicate user data between different apps. How to get User ID  **Default value**: `open_id` **When the value is `user_id`, the following field scopes are required**: `contact:user.employee_id:readonly` | ## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `items` | `department[]` | Department information |
|     `name` | `string` | Department name Note: cannot include slashes **Required field scopes (Satisfy any)**: `contact:contact:readonly_as_app` `contact:department.base:readonly` `contact:contact:access_as_app` `contact:contact:readonly` |
|     `i18n_name` | `department_i18n_name` | Internationalized department name Note: cannot include slashes **Required field scopes (Satisfy any)**: `contact:contact:readonly_as_app` `contact:department.base:readonly` `contact:contact:access_as_app` `contact:contact:readonly` |
|       `zh_cn` | `string` | Department's Chinese name |
|       `ja_jp` | `string` | Department's Japanese name |
|       `en_us` | `string` | Department's English name |
|     `parent_department_id` | `string` | Parent department ID * This value is "0" for a root department. **Required field scopes (Satisfy any)**: `contact:contact:readonly_as_app` `contact:department.organize:readonly` `contact:contact:access_as_app` `contact:contact:readonly` |
|     `department_id` | `string` | Department's custom department ID Note: In addition to meeting the regular rules, it cannot start with `od-` at the same time **Required field scopes (Satisfy any)**: `contact:contact:readonly_as_app` `contact:department.base:readonly` `contact:contact:access_as_app` `contact:contact:readonly` |
|     `open_department_id` | `string` | Department's open_id The open_id of the department, the type is the same as the department_id_type passed in through the query parameter of the request |
|     `leader_user_id` | `string` | Department manager's user ID **Required field scopes (Satisfy any)**: `contact:contact:readonly_as_app` `contact:department.organize:readonly` `contact:contact:access_as_app` `contact:contact:readonly` |
|     `chat_id` | `string` | Department group ID **Required field scopes (Satisfy any)**: `contact:contact:readonly_as_app` `contact:department.base:readonly` `contact:contact:access_as_app` `contact:contact:readonly` |
|     `order` | `string` | Department order, i.e. the order in which the department is displayed among the departments at the same level. **Required field scopes (Satisfy any)**: `contact:contact:readonly_as_app` `contact:department.organize:readonly` `contact:contact:access_as_app` `contact:contact:readonly` |
|     `unit_ids` | `string[]` | List of the department unit's custom IDs. Only one custom ID is supported currently. **Required field scopes (Satisfy any)**: `contact:contact:readonly_as_app` `contact:department.organize:readonly` `contact:contact:access_as_app` `contact:contact:readonly` |
|     `member_count` | `int` | Number of users (including department heads) under the current department and its subordinate departments **Required field scopes (Satisfy any)**: `contact:contact:readonly_as_app` `contact:department.organize:readonly` `contact:contact:access_as_app` `contact:contact:readonly` |
|     `status` | `department_status` | Department status **Required field scopes (Satisfy any)**: `contact:contact:readonly_as_app` `contact:department.base:readonly` `contact:contact:access_as_app` `contact:contact:readonly` |
|       `is_deleted` | `boolean` | Whether it is deleted |
|     `create_group_chat` | `boolean` | Whether the group chat is created |
|     `leaders` | `departmentLeader[]` | Head of department |
|       `leaderType` | `int` | Person in charge type **Optional values are**:  - `1`: Main person in charge - `2`: Deputy responsible person  |
|       `leaderID` | `string` | Person in charge ID **Required field scopes (Satisfy any)**: `contact:contact:readonly_as_app` `contact:department.organize:readonly` `contact:contact:access_as_app` `contact:contact:readonly` |
|     `group_chat_employee_types` | `int[]` | Department group employee type restriction. [] When the list is empty, it means that there is no employee type. The type field can contain the following values, and supports multiple type values; if there are multiple values, separate them with ',' in English: 1. Regular employees 2. Intern 3. Outsourcing 4. Labor 5. Consultant 6. For other custom type fields, you can get the name of the tenant's custom employee type through the interface below, see Get Personnel Type . |
|     `department_hrbps` | `string[]` | Department HRBP **Required field scopes**: `contact:department.hrbp:readonly` |
|     `primary_member_count` | `int` | The number of members whose main department is the current department and its subordinate departments **Required field scopes (Satisfy any)**: `contact:contact:readonly_as_app` `contact:department.organize:readonly` `contact:contact:access_as_app` `contact:contact:readonly` | ### Response body example

{"code":0,
"msg":"success",
"data":{"items":[{"name":"DemoName",
"i18n_name":{"zh_cn":"示例部门",
"ja_jp":"デモ名",
"en_us":"Demo name"},
"parent_department_id":"D067",
"department_id":"D096",
"open_department_id":"od-4e6ac4d14bcd5071a37a39de902c7141",
"leader_user_id":"ou_7dab8a3d3cdcc9da365777c7ad535d62",
"chat_id":"oc_5ad11d72b830411d72b836c20",
"order":"100",
"unit_ids":["custom_unit_id"],
"member_count":100,
"status":{"is_deleted":False},
"create_group_chat":false,
"leaders":[{
    "leaderType": 1,
    "leaderID": "ou_7dab8a3d3cdcc9da365777c7ad535d62"
}],
"group_chat_employee_types":[1,2,3],
"department_hrbps":["ou_7dab8a3d3cdcc9da365777c7ad535d62"],
"primary_member_count":100}]}}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 40001 | param is invalid | Parameter error. |
| 400 | 43010 | big dept forbid recursion error | Unable to query a superdepartment. |
| 400 | 40003 | internal error | Internal error. Please provide the X-Request-Id to our agent. [Contact support](https://applink.larksuite.com/client/helpdesk/open?id=6626260912531570952&extra=%7B%22channel%22:14,%22created_at%22:1614493146,%22scenario_id%22:6885151765134622721,%22signature%22:%22ca94c408b966dc1de2083e5bbcd418294c146e98%22%7D) |
| 400 | 40008 | dept Info is null error | Department information cannot be empty. |
| 403 | 40014 | no parent dept authority error | The parent department must be within the range of contacts data that the app can access. Learn more |
| 400 | 40011 | page size is more than 50 error | The parameter page_size is more than 50, please descrease the page_size and retry. |
| 400 | 40012 | page token is invalid error | The parameter page_token is invalid, and please use the token of the previous response. |
| 401 | 42008 | tenant id is invalid error | Check whether the tenant is valid. |
| 400 | 43007 | duplicated department unit custom id error | Check the department unit custom id whether is duplicated. |
| 403 | 40004 | no dept authority error | The department must be within the range of contacts data that the app can access. Learn more |
