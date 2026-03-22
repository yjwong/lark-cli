---
title: "Obtain the list of sub-departments"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/department/children"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/contact/v3/departments/:department_id/children"
service: "contact-v3"
resource: "department"
section: "Contacts"
scopes:
  - "contact:contact:readonly_as_app"
  - "contact:department.organize:readonly"
  - "contact:contact:access_as_app"
  - "contact:contact:readonly"
field_scopes:
  - "contact:department.organize:readonly"
  - "contact:contact:readonly_as_app"
  - "contact:department.base:readonly"
  - "contact:user.employee_id:readonly"
  - "contact:contact:access_as_app"
  - "contact:contact:readonly"
updated: "1683879166000"
---

# Obtain the list of sub-departments

This API is used to obtain the list of sub-departments by using the department ID.

> - Department ID is required. The root department ID is 0.
> - When `user_access_token` is used, all departments within the user's visible scope of organization ([Log in to the company's Admin to manage scopes](https://www.feishu.cn/admin/security/permission/visibility)) will be returned. If a recursive query is executed, only 1,000 departments at most will be filtered for visibility.
> 
> - When
> `tenant_access_token` is used, the departments will be filtered by verification against the range of contacts data that the app can access.
> If the department ID is 0, the system checks whether the app has the all-staff contact scope. If the department ID is not 0, the system checks whether the app has the contact scope of the department. If no, an error code indicating no scope to access the department will be returned. If yes, the list of sub-departments under the department will be returned (where the recursiveness depends on fetch_child).

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/contact/v3/departments/:department_id/children |
| HTTP Method | GET |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `contact:contact:readonly_as_app` `contact:department.organize:readonly` `contact:contact:access_as_app` `contact:contact:readonly` |
| Required field scopes | > The response body of the API contains the following sensitive fields, and they will be returned only after corresponding scopes are added. If you do not need the fields, it is not recommended that you request the scopes. `contact:department.organize:readonly` `contact:contact:readonly_as_app` `contact:department.base:readonly` `contact:user.employee_id:readonly` `contact:contact:access_as_app` `contact:contact:readonly` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `department_id` | `string` | Department ID. The root department ID is 0. **Example value**: "od-4e6ac4d14bcd5071a37a39de902c7141" **Data validation rules**: - Maximum length: `64` characters - Regular expression: `^[a-zA-Z0-9][a-zA-Z0-9_\-@.]{0,63}$` | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `user_id_type` | `string` | No | User ID categories **Example value**: "open_id" **Optional values are**:  - `open_id`: Identifies a user to an app. The same user has different Open IDs in different apps. How to get Open ID - `union_id`: Identifies a user to a tenant that acts as a developer. A user has the same Union ID in apps developed by the same developer, and has different Union IDs in apps developed by different developers. A developer can use Union ID to link the same user's identities in multiple apps.How to get Union ID - `user_id`: Identifies a user to a tenant. The same user has different User IDs in different tenants. In one single tenant, a user has the same User ID in all apps （including store apps）. User ID is usually used to communicate user data between different apps. How to get User ID  **Default value**: `open_id` **When the value is `user_id`, the following field scopes are required**: `contact:user.employee_id:readonly` |
| `department_id_type` | `string` | No | Type of department ID used in this call **Example value**: "open_department_id" **Optional values are**:  - `department_id`: Identify the department with the custom department_id - `open_department_id`: Identify the department with open_department_id  **Default value**: `open_department_id` |
| `fetch_child` | `boolean` | No | Whether to obtain sub-departments recursively **Example value**: false |
| `page_size` | `int` | No | Page size **Example value**: 10 **Default value**: `10` **Data validation rules**: - Maximum value: `50` |
| `page_token` | `string` | No | Page identifier. It is not filled in the first request, indicating traversal from the beginning; when there will be more groups, the new page_token will be returned at the same time, and the next traversal can use the page_token to get more groups **Example value**: "AQD9/Rn9eij9Pm39ED40/RD/cIFmu77WxpxPB/2oHfQLZ+G8JG6tK7+ZnHiT7COhD2hMSICh/eBl7cpzU6JEC3J7COKNe4jrQ8ExwBCR" | ## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `has_more` | `boolean` | Whether the response body has more parameters |
|   `page_token` | `string` | Page identifier, when has_more is true, a new page_token will also be returned. Otherwise, page_token will not be returned |
|   `items` | `department[]` | Department list |
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
|     `create_group_chat` | `boolean` | Whether to create a department group. It defaults to false. When creating a department group, the default group name is the department name, and the default group owner is the person in charge of the department |
|     `leaders` | `departmentLeader[]` | Head of department |
|       `leaderType` | `int` | Person in charge type **Optional values are**:  - `1`: Main person in charge - `2`: Deputy responsible person  |
|       `leaderID` | `string` | Person in charge ID **Required field scopes (Satisfy any)**: `contact:contact:readonly_as_app` `contact:department.organize:readonly` `contact:contact:access_as_app` `contact:contact:readonly` |
|     `group_chat_employee_types` | `int[]` | Department group employee type restriction. [] When the list is empty, it means that there is no employee type. The type field can contain the following values, and supports multiple type values; if there are multiple values, separate them with ',' in English: 1. Regular employees 2. Intern 3. Outsourcing 4. Labor 5. Consultant 6. For other custom type fields, you can get the name of the tenant's custom employee type through the interface below, see Get Personnel Type . | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
        "has_more": true,
        "page_token": "AQD9/Rn9eij9Pm39ED40/RD/cIFmu77WxpxPB/2oHfQLZ%2BG8JG6tK7%2BZnHiT7COhD2hMSICh/eBl7cpzU6JEC3J7COKNe4jrQ8ExwBCR",
        "items": [
            {
                "name": "DemoName",
                "i18n_name": {
                    "zh_cn": "Demo名称",
                    "ja_jp": "デモ名",
                    "en_us": "Demo Name"
                },
                "parent_department_id": "D067",
                "department_id": "D096",
                "open_department_id": "od-4e6ac4d14bcd5071a37a39de902c7141",
                "leader_user_id": "ou_7dab8a3d3cdcc9da365777c7ad535d62",
                "chat_id": "oc_5ad11d72b830411d72b836c20",
                "order": "100",
                "unit_ids": [
                    "custom_unit_id"
                ],
                "member_count": 100,
                "status": {
                    "is_deleted": false
                },
                "create_group_chat": false,
                "leaders": [
                    {
                        "leaderType": 1,
                        "leaderID": "ou_7dab8a3d3cdcc9da365777c7ad535d62"
                    }
                ]
            }
        ]
    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 43010 | big dept forbid recursion error | Unable to query a superdepartment. |
| 400 | 40003 | internal error | Internal error. Please provide the X-Request-Id to our agent. [Contact support](https://applink.larksuite.com/client/helpdesk/open?id |
| 400 | 40008 | dept Info is null error | Department information cannot be empty. |
| 403 | 40014 | no parent dept authority error | No access to the parent department(visible scope of organization). |
| 400 | 40011 | page size is invalid | The page size is between 1 and 50. |
| 400 | 40012 | page token is invalid error | Page token is invalid. |
| 401 | 42008 | tenant id is invalid error | Check whether the tenant is valid. |
| 400 | 43007 | duplicated department custom id error | Duplicate department's custom ID in the company. |
| 403 | 40004 | no dept authority error | The department must be within the range of contacts data that the app can access. Learn more |
