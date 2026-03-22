---
title: "Get Department Information List"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/department/list"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/contact/v3/departments"
service: "contact-v3"
resource: "department"
section: "Deprecated Version (Not Recommended)"
scopes:
  - "contact:contact:readonly_as_app"
  - "contact:department.organize:readonly"
  - "contact:contact:access_as_app"
  - "contact:contact:readonly"
field_scopes:
  - "contact:contact:readonly_as_app"
  - "contact:department.base:readonly"
  - "contact:department.organize:readonly"
  - "contact:user.employee_id:readonly"
  - "contact:contact:readonly"
  - "contact:contact:access_as_app"
updated: "1646999084000"
---

# Obtain department information list

This API is used to obtain the list of sub-departments of a department. FAQs.

> - When user_access_token is used, all departments within the user's visible scope of organization ([Log in to the company's Admin to manage scopes](https://www.larksuite.com/admin/security/permission/visibility)) will be returned. If a recursive query is executed, only 1,000 departments at most will be filtered for visibility.
> 
> - When
> tenant_access_token is used, the departments will be filtered by verification against the range of contacts data that the app can access.
> Since parent_department_id is optional, the results returned for range verification vary depending on whether this parameter is specified:
>  1. When
> parent_department_id is set to A (root department ID is 0), A will be verified for whether it is within the access range. If yes (when parent_department_id is 0, the scope to access all employees will be verified), the list of sub-departments under the department will be returned (where the recursiveness depends on fetch_child). If no, an error code indicating no scope to access the department will be returned.
>  2. When
> parent_department_id is not specified, if the contacts access range is all employees, only the root department ID (which is 0) will be returned; otherwise, the department ID and sub-departments configured according to the contacts access range will be returned (where the recursiveness depends on
> fetch_child).

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/contact/v3/departments |
| HTTP Method | GET |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `contact:contact:readonly_as_app` `contact:department.organize:readonly` `contact:contact:access_as_app` `contact:contact:readonly` |
| Required field scopes | > The response body of the API contains the following sensitive fields, and they will be returned only after corresponding scopes are added. If you do not need the fields, it is not recommended that you request the scopes. `contact:contact:readonly_as_app` `contact:department.base:readonly` `contact:department.organize:readonly` `contact:user.employee_id:readonly` `contact:contact:readonly` `contact:contact:access_as_app` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" Learn more about obtaining and using access_token. | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `user_id_type` | `string` | No | User ID categories **Example value**: "open_id" **Optional values are**: - `open_id`: user's open id - `union_id`: user's union id - `user_id`: user's user id **Default value**: `open_id` **When the value is `user_id`, the following field scopes are required**: `contact:user.employee_id:readonly` |
| `department_id_type` | `string` | No | Type of department ID used in this call **Example value**: "open_department_id" **Optional values are**: - `department_id`: Identify the department with the custom department_id - `open_department_id`: Identify the department with open_department_id **Default value**: `open_department_id` |
| `parent_department_id` | `string` | No | Parent department ID. If this parameter is specified, all sub-departments under the department will be obtained. The ID entered must be the one specified by department_id_type. **Example value**: "od-4e6ac4d14bcd5071a37a39de902c7141" |
| `fetch_child` | `boolean` | No | Whether to obtain sub-departments recursively **Example value**: Whether to obtain sub-departments recursively. It defaults to false. |
| `page_token` | `string` | No | Page identifier. It is not filled in the first request, indicating traversal from the beginning; when there will be more groups, the new page_token will be returned at the same time, and the next traversal can use the page_token to get more groups **Example value**: "AQD9/Rn9eij9Pm39ED40/RD/cIFmu77WxpxPB/2oHfQLZ%2BG8JG6tK7%2BZnHiT7COhD2hMSICh/eBl7cpzU6JEC3J7COKNe4jrQ8ExwBCR" |
| `page_size` | `int` | No | Page size **Example value**: 10 **Data validation rules**: - Maximum value: `50` | ## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
| ∟ `has_more` | `boolean` | Whether the response body has more parameters |
| ∟ `page_token` | `string` | Page identifier, when has_more is true, a new page_token will also be returned. Otherwise, page_token will not be returned |
| ∟ `items` | `department[]` |  |
| ∟ `name` | `string` | Department name **Required field scopes (Satisfy any)**: `contact:contact:readonly_as_app` `contact:department.base:readonly` `contact:contact:readonly` `contact:contact:access_as_app` |
| ∟ `i18n_name` | `department_i18n_name` | Internationalized department name **Required field scopes (Satisfy any)**: `contact:contact:readonly_as_app` `contact:department.base:readonly` `contact:contact:readonly` `contact:contact:access_as_app` |
| ∟ `zh_cn` | `string` | Department's Chinese name |
| ∟ `ja_jp` | `string` | Department's Japanese name |
| ∟ `en_us` | `string` | Department's English name |
| ∟ `parent_department_id` | `string` | Parent department ID * This value is "0" for a root department. **Required field scopes (Satisfy any)**: `contact:contact:readonly_as_app` `contact:department.organize:readonly` `contact:contact:readonly` `contact:contact:access_as_app` |
| ∟ `department_id` | `string` | Department's custom department ID **Required field scopes (Satisfy any)**: `contact:contact:readonly_as_app` `contact:department.base:readonly` `contact:contact:readonly` `contact:contact:access_as_app` |
| ∟ `open_department_id` | `string` | Department's open_id |
| ∟ `leader_user_id` | `string` | Department manager's user ID **Required field scopes (Satisfy any)**: `contact:contact:readonly_as_app` `contact:department.organize:readonly` `contact:contact:readonly` `contact:contact:access_as_app` |
| ∟ `chat_id` | `string` | Department group ID **Required field scopes (Satisfy any)**: `contact:contact:readonly_as_app` `contact:department.base:readonly` `contact:contact:readonly` `contact:contact:access_as_app` |
| ∟ `order` | `string` | Department order, i.e. the order in which the department is displayed among the departments at the same level. **Required field scopes (Satisfy any)**: `contact:contact:readonly_as_app` `contact:department.organize:readonly` `contact:contact:readonly` `contact:contact:access_as_app` |
| ∟ `unit_ids` | `string[]` | List of the department unit's custom IDs. Only one custom ID is supported currently. **Required field scopes (Satisfy any)**: `contact:contact:readonly_as_app` `contact:department.organize:readonly` `contact:contact:readonly` `contact:contact:access_as_app` |
| ∟ `member_count` | `int` | Number of users under the department **Required field scopes (Satisfy any)**: `contact:contact:readonly_as_app` `contact:department.organize:readonly` `contact:contact:readonly` `contact:contact:access_as_app` |
| ∟ `status` | `department_status` | Department status **Required field scopes (Satisfy any)**: `contact:contact:readonly_as_app` `contact:department.base:readonly` `contact:contact:readonly` `contact:contact:access_as_app` |
| ∟ `is_deleted` | `boolean` | Whether it is deleted | ### Response body example

```json
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
                }
            }
        ]
    }
}
```

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 43010 | big dept forbid recursion error | Unable to query a superdepartment. |
| 400 | 40003 | internal error | Internal error. Please provide the X-Request-Id to our agent. [Contact support](https://applink.larksuite.com/client/helpdesk/open?id |
| 400 | 40008 | dept Info is null error | Department information cannot be empty. |
| 403 | 40014 | no parent dept authority error | No access to the parent department. |
| 400 | 40011 | page size is invalid | The page size is between 1 and 50. |
| 400 | 40012 | page token is invalid error | Page token is invalid. |
| 401 | 42008 | tenant id is invalid error | Check whether the tenant is valid. |
| 400 | 43007 | duplicated department custom id error | Duplicate department's custom ID in the company. |
| 403 | 40004 | no dept authority error | The department must be within the range of contacts data that the app can access. Learn more |
