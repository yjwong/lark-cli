---
title: "Obtain single department information"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/department/get"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/contact/v3/departments/:department_id"
service: "contact-v3"
resource: "department"
section: "Contacts"
scopes:
  - "contact:contact:readonly_as_app"
  - "contact:contact:access_as_app"
  - "contact:contact:readonly"
field_scopes:
  - "contact:department.organize:readonly"
  - "contact:contact:readonly_as_app"
  - "contact:department.base:readonly"
  - "contact:user.employee_id:readonly"
  - "contact:contact:access_as_app"
  - "contact:contact:readonly"
updated: "1671073180000"
---

# Obtain single department information

This API is used to obtain the information of a single department in contacts.

> - When using `tenant_access_token`, the app must have the contacts scope of the department to be queried. To obtain root department information, the app must have the all-staff contact scope.
> - When using `user_access_token`, you must have the visibility of the department to be queried. To obtain root department information, you must have the visibility of all employees.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/contact/v3/departments/:department_id |
| HTTP Method | GET |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `contact:contact:readonly_as_app` `contact:contact:access_as_app` `contact:contact:readonly` |
| Required field scopes | > The response body of the API contains the following sensitive fields, and they will be returned only after corresponding scopes are added. If you do not need the fields, it is not recommended that you request the scopes. `contact:department.organize:readonly` `contact:contact:readonly_as_app` `contact:department.base:readonly` `contact:user.employee_id:readonly` `contact:contact:access_as_app` `contact:contact:readonly` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `department_id` | `string` | ID of the department to be obtained For the description of different IDs and how to obtain them, see Department ID Description **Example value**: "od-4e6ac4d14bcd5071a37a39de902c7141" **Data validation rules**: - Maximum length: `64` characters - Regular expression: `^[a-zA-Z0-9][a-zA-Z0-9_\-@.]{0,63}$` | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `user_id_type` | `string` | No | User ID categories **Example value**: "open_id" **Optional values are**:  - `open_id`: Identifies a user to an app. The same user has different Open IDs in different apps. How to get Open ID - `union_id`: Identifies a user to a tenant that acts as a developer. A user has the same Union ID in apps developed by the same developer, and has different Union IDs in apps developed by different developers. A developer can use Union ID to link the same user's identities in multiple apps.How to get Union ID - `user_id`: Identifies a user to a tenant. The same user has different User IDs in different tenants. In one single tenant, a user has the same User ID in all apps （including store apps）. User ID is usually used to communicate user data between different apps. How to get User ID  **Default value**: `open_id` **When the value is `user_id`, the following field scopes are required**: `contact:user.employee_id:readonly` |
| `department_id_type` | `string` | No | Type of department ID used in this call **Example value**: "open_department_id" **Optional values are**:  - `department_id`: Identify the department with the custom department_id - `open_department_id`: Identify the department with open_department_id  **Default value**: `open_department_id` | ## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `department` | `department` | Department information |
|     `name` | `string` | Department name **Required field scopes (Satisfy any)**: `contact:contact:readonly_as_app` `contact:department.base:readonly` `contact:contact:access_as_app` `contact:contact:readonly` |
|     `i18n_name` | `department_i18n_name` | Internationalized department name **Required field scopes (Satisfy any)**: `contact:contact:readonly_as_app` `contact:department.base:readonly` `contact:contact:access_as_app` `contact:contact:readonly` |
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
|     `member_count` | `int` | Number of users under the department **Required field scopes (Satisfy any)**: `contact:contact:readonly_as_app` `contact:department.organize:readonly` `contact:contact:access_as_app` `contact:contact:readonly` |
|     `status` | `department_status` | Department status **Required field scopes (Satisfy any)**: `contact:contact:readonly_as_app` `contact:department.base:readonly` `contact:contact:access_as_app` `contact:contact:readonly` |
|       `is_deleted` | `boolean` | Whether it is deleted |
|     `leaders` | `departmentLeader[]` | Head of department |
|       `leaderType` | `int` | Person in charge type **Optional values are**:  - `1`: Main person in charge - `2`: Deputy responsible person  |
|       `leaderID` | `string` | Person in charge ID **Required field scopes (Satisfy any)**: `contact:contact:readonly_as_app` `contact:department.organize:readonly` `contact:contact:access_as_app` `contact:contact:readonly` |
|     `group_chat_employee_types` | `int[]` | Department group employee type restriction. [] When the list is empty, it means that there is no employee type. The type field can contain the following values, and supports multiple type values; if there are multiple values, separate them with ',' in English: 1. Regular employees 2. Intern 3. Outsourcing 4. Labor 5. Consultant 6. For other custom type fields, you can get the name of the tenant's custom employee type through the interface below, see Get Personnel Type . | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
        "department": {
            "name": "DemoName",
            "i18n_name": {
                "zh_cn": "Demo name",
                "ja_jp": "Name",
                "en_us": "Demo Name"
            },
            "parent_department_id": "D067",
            "department_id": "D096",
            "open_department_id": "Od-4e6ac4d14bcd5071a37a39de902c7141",
            "leader_user_id": "ou_7dab8a3d3cdcc9da365777c7ad535d62",
            "chat_id": "oc_5ad11d72b830411d72b836c20",
            "order": "100",
            "unit_ids": [
                "custom_unit_id"
            ],
            "member_count": 100,
            "status": {
                "is_deleted": False
            },
            "leaders": [
                {
                    "leaderType": 1,
                    "leaderID": "ou_7dab8a3d3cdcc9da365777c7ad535d62"
                }
            ],
            "group_chat_employee_types": [
                [1,2,3]
            ]
        }
    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 40001 | param error | Parameter error. |
| 400 | 43010 | big dept forbid recursion error | Unable to query a superdepartment. |
| 400 | 40003 | internal error | Internal error. Please provide the X-Request-Id to our agent. [Contact support](https://applink.larksuite.com/client/helpdesk/open?id |
| 403 | 40004 | no dept authority error | The department must be within the range of contacts data that the app can access. Learn more |
| 400 | 40008 | dept Info is null error | Department information cannot be empty. |
| 401 | 42008 | tenant id is invalid error | Check whether the tenant is valid. |
| 400 | 40015 | department not exist | Department does not exist. Check whether the parameters are correct. |
