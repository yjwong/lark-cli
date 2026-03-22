---
title: "Create a department"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/department/create"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/contact/v3/departments"
service: "contact-v3"
resource: "department"
section: "Contacts"
scopes:
  - "contact:contact"
field_scopes:
  - "contact:department.organize:readonly"
  - "contact:contact:readonly_as_app"
  - "contact:department.base:readonly"
  - "contact:user.employee_id:readonly"
  - "contact:contact:access_as_app"
  - "contact:contact:readonly"
updated: "1713857233000"
---

# Create a department

This API is used to create a department in contacts.

> You can create departments only under departments within the app's contact scopes. If you need to create a sub-department under the root department, you must set the app's contact scope to "All members". Store apps have no scope to call this API.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/contact/v3/departments |
| HTTP Method | POST |
| Supported app types | custom |
| Required scopes | `contact:contact` |
| Required field scopes | > The response body of the API contains the following sensitive fields, and they will be returned only after corresponding scopes are added. If you do not need the fields, it is not recommended that you request the scopes. `contact:department.organize:readonly` `contact:contact:readonly_as_app` `contact:department.base:readonly` `contact:user.employee_id:readonly` `contact:contact:access_as_app` `contact:contact:readonly` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `user_id_type` | `string` | No | User ID categories **Example value**: "open_id" **Optional values are**:  - `open_id`: Identifies a user to an app. The same user has different Open IDs in different apps. How to get Open ID - `union_id`: Identifies a user to a tenant that acts as a developer. A user has the same Union ID in apps developed by the same developer, and has different Union IDs in apps developed by different developers. A developer can use Union ID to link the same user's identities in multiple apps.How to get Union ID - `user_id`: Identifies a user to a tenant. The same user has different User IDs in different tenants. In one single tenant, a user has the same User ID in all apps （including store apps）. User ID is usually used to communicate user data between different apps. How to get User ID  **Default value**: `open_id` **When the value is `user_id`, the following field scopes are required**: `contact:user.employee_id:readonly` |
| `department_id_type` | `string` | No | Type of department ID used in this call Description of different IDs refers to Department ID Description **Example value**: "open_department_id" **Optional values are**:  - `department_id`: Identify the department with the custom department_id - `open_department_id`: Identify the department with open_department_id  |
| `client_token` | `string` | No | It is used to determine whether it is the same request idempotent to avoid repeated creation. String type, self-generated. **Example value**: "473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E" | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `name` | `string` | Yes | Department name **Example value**: "DemoName" **Data validation rules**: - Minimum length: `1` characters |
| `i18n_name` | `department_i18n_name` | No | Internationalized department name |
|   `zh_cn` | `string` | No | Department's Chinese name **Example value**: "Demo name" |
|   `ja_jp` | `string` | No | Department's Japanese name **Example value**: "Name" |
|   `en_us` | `string` | No | Department's English name **Example value**: "Demo Name" |
| `parent_department_id` | `string` | Yes | Parent department ID * This value is "0" for a root department. **Example value**: "D067" |
| `department_id` | `string` | No | Department's custom department ID Note: In addition to meeting the regular rules, it cannot start with `od-` at the same time **Example value**: "D096" **Data validation rules**: - Maximum length: `64` characters - Regular expression: `^[a-zA-Z0-9][a-zA-Z0-9_\-@.]{0,63}$` |
| `leader_user_id` | `string` | No | Department manager's user ID **Example value**: "ou_7dab8a3d3cdcc9da365777c7ad535d62" |
| `order` | `string` | No | Department order, i.e. the order in which the department is displayed among the departments at the same level. **Example value**: "100" |
| `unit_ids` | `string[]` | No | List of the department unit's custom IDs. Only one custom ID is supported currently. **Example value**: custom_unit_id |
| `create_group_chat` | `boolean` | No | Whether to create a department group. It defaults to false. When creating a department group, the default group name is the department name, and the default group owner is the person in charge of the department **Example value**: false |
| `leaders` | `departmentLeader[]` | No | Head of department |
|   `leaderType` | `int` | Yes | Person in charge type **Example value**: 1 **Optional values are**:  - `1`: Main person in charge - `2`: Deputy responsible person  |
|   `leaderID` | `string` | Yes | Person in charge ID **Example value**: "ou_7dab8a3d3cdcc9da365777c7ad535d62" |
| `group_chat_employee_types` | `int[]` | No | Department group employee type restriction. [] When the list is empty, it means that there is no employee type. The type field can contain the following values, and supports multiple type values; if there are multiple values, separate them with ',' in English: 1. Regular employees 2. Intern 3. Outsourcing 4. Labor 5. Consultant 6. For other custom type fields, you can get the name of the tenant's custom employee type through the interface below, see Get Personnel Type . **Example value**: [1,2,3] | ### Request body example

{
    "name": "DemoName",
    "i18n_name": {
        "zh_cn": "Demo名称",
        "ja_jp": "デモ名",
        "en_us": "Demo Name"
    },
    "parent_department_id": "D067",
    "department_id": "D096",
    "leader_user_id": "ou_7dab8a3d3cdcc9da365777c7ad535d62",
    "order": "100",
    "unit_ids": [
        "custom_unit_id"
    ],
    "create_group_chat": false,
     "leaders": [
                {
                    "leaderID": "ou_357368f98775f04bea02afc6b1d33478",
                    "leaderType": 1
                }
            ]
}

## Response

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
          "leaders": [
                {
                    "leaderID": "ou_357368f98775f04bea02afc6b1d33478",
                    "leaderType": 1
                }
            ]
        }
    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 43005 | duplicate order error | The department order must be unique. Check the order and try again. |
| 400 | 40018 | duplicate name error | Under the same parent department, The department name must be unique. Check the name and try again. |
| 400 | 43008 | custom dept id invalid error | The custom department ID is invalid. It cannot start with "od-", cannot be "0", and cannot exceed 64 characters. |
| 400 | 43013 | dept too many children error | Too many sub-departments. |
| 400 | 40003 | internal error | Internal error. Please provide the X-Request-Id to our agent. [Contact support](https://applink.larksuite.com/client/helpdesk/open?id=6626260912531570952&extra=%7B%22channel%22%3A14%2C%22created_at%22%3A1614493146%2C%22scenario_id%22%3A6885151765134622721%2C%22signature%22%3A%22ca94c408b966dc1de2083e5bbcd418294c146e98%22%7D) |
| 403 | 40004 | no dept authority error | The department must be within the range of contacts data that the app can access. Learn more |
| 400 | 40008 | dept Info is null error | Department information cannot be empty. |
| 400 | 40010 | chat id is invalid error | Incorrect format of the department group ID |
| 403 | 40014 | no parent dept authority error | No access to the parent department(visible scope of organization). |
| 401 | 43000 | dept has already exist error | The department already exists. |
| 401 | 40016 | dept name can not be nul error | Department name cannot be empty. |
| 400 | 42006 | user has resigned error | The user has been terminated. |
| 400 | 41012 | user id invalid error | Invalid user ID. |
| 400 | 40021 | no a same request error | The two actions do not belong to the same request. Check whether the request value was changed. |
| 400 | 43004 | illegal unit error | Invalid department unit ID |
| 400 | 43007 | duplicated department custom id error | Duplicate department's custom ID in the company. |
| 403 | 40004 | no dept authority error | The department must be within the range of contacts data that the app can access. Learn more |
| 400 | 44101 | miss parent department id error | The request does not have parent department ID. |
| 400 | 43017 | relate dept over limit | Affiliates exceed limits |
| 400 | 43018 | duplicate i18n name | Duplicate i18n names |
| 400 | 43019 | exceed dept max level | More than department level |
| 400 | 43021 | department chat not exist | Department group does not exist |
| 400 | 43022 | department name duplicate | Department name duplicate |
| 400 | 43023 | dept structure no permissione | Organizational structure has no authority |
| 400 | 43024 | dept structure tenant lock fail | Department structure change Failed to acquire tenant lock |
| 400 | 43025 | top department leader unjoined | Users cannot become the person in charge of the enterprise without joining |
| 400 | 43026 | employee type is not valid | Please use the exact employee type |
