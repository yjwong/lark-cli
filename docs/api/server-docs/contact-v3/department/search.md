---
title: "Search for departments"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/department/search"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/contact/v3/departments/search"
service: "contact-v3"
resource: "department"
section: "Contacts"
scopes:
  - "contact:contact:readonly_as_app"
  - "contact:contact:access_as_app"
field_scopes:
  - "contact:department.organize:readonly"
  - "contact:contact:readonly_as_app"
  - "contact:department.base:readonly"
  - "contact:user.employee_id:readonly"
  - "contact:contact:access_as_app"
  - "contact:contact:readonly"
updated: "1683879166000"
---

# Search for departments

This API is used to search for department data with keywords. Department visibility is set by administrators in Admin.

> If a department exists but the user cannot find it, this may also be because the administrator has set the department visibility.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/contact/v3/departments/search |
| HTTP Method | POST |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `contact:contact:readonly_as_app` `contact:contact:access_as_app` |
| Required field scopes | > The response body of the API contains the following sensitive fields, and they will be returned only after corresponding scopes are added. If you do not need the fields, it is not recommended that you request the scopes. `contact:department.organize:readonly` `contact:contact:readonly_as_app` `contact:department.base:readonly` `contact:user.employee_id:readonly` `contact:contact:access_as_app` `contact:contact:readonly` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `user_id_type` | `string` | No | User ID categories **Example value**: "open_id" **Optional values are**:  - `open_id`: Identifies a user to an app. The same user has different Open IDs in different apps. How to get Open ID - `union_id`: Identifies a user to a tenant that acts as a developer. A user has the same Union ID in apps developed by the same developer, and has different Union IDs in apps developed by different developers. A developer can use Union ID to link the same user's identities in multiple apps.How to get Union ID - `user_id`: Identifies a user to a tenant. The same user has different User IDs in different tenants. In one single tenant, a user has the same User ID in all apps （including store apps）. User ID is usually used to communicate user data between different apps. How to get User ID  **Default value**: `open_id` **When the value is `user_id`, the following field scopes are required**: `contact:user.employee_id:readonly` |
| `department_id_type` | `string` | No | Type of department ID used in this call **Example value**: "open_department_id" **Optional values are**:  - `department_id`: Identify the department with the custom department_id - `open_department_id`: Identify the department with open_department_id  |
| `page_token` | `string` | No | Page identifier. It is not filled in the first request, indicating traversal from the beginning; when there will be more groups, the new page_token will be returned at the same time, and the next traversal can use the page_token to get more groups **Example value**: "AQD9/Rn9eij9Pm39ED40/RD/cIFmu77WxpxPB/2oHfQLZ+G8JG6tK7+ZnHiT7COhD2hMSICh/eBl7cpzU6JEC3J7COKNe4jrQ8ExwBCR" |
| `page_size` | `int` | No | Page size **Example value**: 10 **Default value**: `20` **Data validation rules**: - Maximum value: `50` | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `query` | `string` | Yes | Search keywords. The matching field is the department name (The department's internationalized name cannot be matched). **Example value**: "DemoName" | ### Request body example

{
    "query": "DemoName"
}

## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `items` | `department[]` |  |
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
|     `group_chat_employee_types` | `int[]` | Department group employee type restriction. [] When the list is empty, it means that there is no employee type. The type field can contain the following values, and supports multiple type values; if there are multiple values, separate them with ',' in English: 1. Regular employees 2. Intern 3. Outsourcing 4. Labor 5. Consultant 6. For other custom type fields, you can get the name of the tenant's custom employee type through the interface below, see Get Personnel Type . |
|   `page_token` | `string` | Page identifier, when has_more is true, a new page_token will also be returned. Otherwise, page_token will not be returned |
|   `has_more` | `boolean` | Whether the response body has more parameters | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
        "items": [
            {
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
                "create_group_chat": false,
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
        ],
        "page_token": "AQD9/Rn9eij9Pm39ED40/RD/cIFmu77WxpxPB/2oHfQLZ%2BG8JG6tK7%2BZnHiT7COhD2hMSICh/eBl7cpzU6JEC3J7COKNe4jrQ8ExwBCR",
        "has_more": true
    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 48001 | invalid param | Parameter error. Check the parameter. |
| 500 | 48002 | Internal Server Error | System error |
