---
title: "Search department"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/directory-v1/department/search"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/directory/v1/departments/search"
service: "directory-v1"
resource: "department"
section: "Organization"
rate_limit: "1000 per minute、50 per second"
scopes:
  - "directory:department:search"
field_scopes:
  - "directory:department.base:read"
  - "directory:department.count:read"
  - "directory:department.custom_field:read"
  - "directory:department.data_source:read"
  - "directory:department.department_path:read"
  - "directory:department.external_id:read"
  - "directory:department.has_child:read"
  - "directory:department.leader:read"
  - "directory:department.name:read"
  - "directory:department.order_weight:read"
  - "directory:department.organization:read"
  - "directory:department.parent_id:read"
  - "directory:department.status:read"
  - "directory:employee.base.external_id:read"
updated: "1756439307000"
---

# SearchDepartment

This interface is used to search for department information, search for department information through keywords such as department names, and return a list of departments that meet the criteria. 

> Attention:
> - This interface supports tenant_access_token and user_access_token.
> - When using tenant_access_token, the data permissions follow the contacts permissions of the app, and the field data returned is the field that the app has permissions for.
> - When using user_access_token, the default is the administrator user, and the administrator management scope will be verified. When the user has multiple administrator identities to view employee information, the administrator management scope takes the maximum set.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/directory/v1/departments/search |
| HTTP Method | POST |
| Rate Limit | 1000 per minute、50 per second |
| Supported app types | custom,isv |
| Required scopes | `directory:department:search` |
| Required field scopes | > The response body of the API contains the following sensitive fields, and they will be returned only after corresponding scopes are added. If you do not need the fields, it is not recommended that you request the scopes. `directory:department.base:read` `directory:department.count:read` `directory:department.custom_field:read` `directory:department.data_source:read` `directory:department.department_path:read` `directory:department.external_id:read` `directory:department.has_child:read` `directory:department.leader:read` `directory:department.name:read` `directory:department.order_weight:read` `directory:department.organization:read` `directory:department.parent_id:read` `directory:department.status:read` `directory:employee.base.external_id:read` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `employee_id_type` | `string` | No | User ID categories **Example value**: open_id **Optional values are**:  - `open_id`: Identifies a user to an app. The same user has different Open IDs in different apps. How to get Open ID - `union_id`: Identifies a user to a tenant that acts as a developer. A user has the same Union ID in apps developed by the same developer, and has different Union IDs in apps developed by different developers. A developer can use Union ID to link the same user's identities in multiple apps.How to get Union ID - `employee_id`: The unique identifier of an employee within the enterprise. Supports customization, the system automatically generates it if it is not customized. ID supports modification. How to get employee_id: - Enterprise administrators go to the Admin Console > Organizational Structure > Members and Departments page, click Member Details to query employee IDs - Through the interface of Get employee list in batches, query employee ID by mobile phone number or email.  **Default value**: `open_id` **When the value is `employee_id`, the following field scopes are required**: `directory:employee.base.external_id:read` |
| `department_id_type` | `string` | No | The type of department ID used in this call **Example value**: open_department_id **Optional values are**:  - `open_department_id`: It is used to identify a department in a specific application, and the same department has different open_department_id in different applications. - `department_id`: Used to identify a unique department within a tenant  **Default value**: `open_department_id` | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `query` | `string` | Yes | Search for keywords. Supports searching by department name, with a maximum input of 100 characters. **Example value**: "Zhang" |
| `page_request` | `page_condition` | Yes | paging information |
|   `page_size` | `int` | No | Maximum number of entries returned per page, maximum value is 100 **Default value**：20 **Minimum Value**: 0 **Example value**: 100 |
|   `page_token` | `string` | No | Sequential paging query, can not skip page query, support deep paging, in the need to traverse all the data scene can only use this method. The first pass empty string or not, later pass the last return value in the page_token **Example value**: "Xdcftvygbuhijhgrexr" |
| `required_fields` | `string[]` | Yes | List of fields to be queried. Permissioned row and column data will be returned according to the passed field list. If not passed, no fields will be returned Learn more: Field enumeration instructions **Example value**: ["name"] **Data validation rules**: - Length range: `0` ～ `100` | ### Request body example

{
    "query": "Zhang",
    "page_request": {
        "page_size": 100,
        "page_token": "Xdcftvygbuhijhgrexr"
    },
    "required_fields": [
        "name"
    ]
}

## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `departments` | `department[]` | Department information |
|     `department_id` | `string` | Department ID, consistent with the department_id_type type **Required field scopes (Satisfy any)**: `directory:department.base:read` `directory:department.external_id:read` |
|     `department_count` | `department_count` | Department member count and subdepartment count. There may be a delay in the calculation result. **Required field scopes (Satisfy any)**: `directory:department.count:read` `directory:department.organization:read` |
|       `recursive_members_count` | `string` | number of recursive members |
|       `direct_members_count` | `string` | Number of direct members |
|       `recursive_members_count_exclude_leaders` | `string` | Number of recursive members (excluding leaders) |
|       `recursive_departments_count` | `string` | Number of recursive subdepartments |
|       `direct_departments_count` | `string` | Number of direct sub-departments |
|     `has_child` | `boolean` | Whether has sub-department **Required field scopes (Satisfy any)**: `directory:department.has_child:read` `directory:department.organization:read` |
|     `leaders` | `department_leader[]` | department head **Required field scopes**: `directory:department.leader:read` |
|       `leader_type` | `int` | Type of department head **Optional values are**:  - `1`: main - `2`: vice  |
|       `leader_id` | `string` | Department head ID, consistent with the employee_id_type type |
|     `parent_department_id` | `string` | Parent Department ID **Required field scopes (Satisfy any)**: `directory:department.organization:read` `directory:department.parent_id:read` |
|     `name` | `i18n_text` | Department name **Required field scopes (Satisfy any)**: `directory:department.base:read` `directory:department.name:read` |
|       `default_value` | `string` | default value |
|       `i18n_value` | `map<string, string>` | Internationalized value, key is zh_cn, ja_jp, en_us, value is the corresponding value |
|     `enabled_status` | `boolean` | Whether enabled **Required field scopes**: `directory:department.status:read` |
|     `order_weight` | `string` | Department ranking weight **Required field scopes (Satisfy any)**: `directory:department.order_weight:read` `directory:department.organization:read` |
|     `custom_field_values` | `custom_field_value[]` | Deparment custom field **Required field scopes**: `directory:department.custom_field:read` |
|       `field_type` | `string` | Custom field key **Optional values are**:  - `1`: Multi-line text - `2`: Web page link - `3`: Enum options - `4`: User - `9`: phone - `10`: Multiple-choice enumeration type (currently only supports text type) - `11`: Personnel List  |
|       `text_value` | `i18n_text` | Custom field type |
|         `default_value` | `string` | default value |
|         `i18n_value` | `map<string, string>` | Internationalized value, key is zh_cn, ja_jp, en_us, value is the corresponding value |
|       `url_value` | `url_value` | Text field value |
|         `link_text` | `i18n_text` |  |
|           `default_value` | `string` | default value |
|           `i18n_value` | `map<string, string>` | Internationalized value, key is zh_cn, ja_jp, en_us, value is the corresponding value |
|         `url` | `string` |  |
|         `pcurl` | `string` |  |
|       `enum_value` | `enum_value` | Web link field value |
|         `enum_ids` | `string[]` | Web page title |
|         `enum_type` | `string` | Mobile web link **Optional values are**:  - `1`: Text - `2`: Image  |
|       `user_values` | `user_value[]` | Enum Field value |
|         `ids` | `string[]` | Person ID，consistent with the employee_id_type type |
|       `phone_value` | `phone_value` | Person field value |
|         `phone_number` | `string` |  |
|         `extension_number` | `string` |  |
|       `field_key` | `string` |  |
|     `department_path_infos` | `department_base_info[]` | Department path information. The sorting order is from root level to last level, excluding the root department. **Required field scopes**: `directory:department.department_path:read` |
|       `department_id` | `string` | Department ID, consistent with the department_id_type type |
|       `department_name` | `i18n_text` | Department name |
|         `default_value` | `string` | default value |
|         `i18n_value` | `map<string, string>` | Internationalized value, key is zh_cn, ja_jp, en_us, value is the corresponding value |
|     `data_source` | `int` | Data Source **Optional values are**:  - `1`: LarkAdmin - `2`: CoreHR - `3`: SCIM  **Required field scopes (Satisfy any)**: `directory:department.base:read` `directory:department.data_source:read` |
|   `page_response` | `page_response` | Paginated results |
|     `has_more` | `boolean` | Whether there any follow-up results |
|     `page_token` | `string` | Paging mark, when has_more is true, a new page_token will be returned at the same time, otherwise page_token will not be returned. |
|   `abnormals` | `abnormal_record[]` | Field exception information |
|     `id` | `string` | Exception ID |
|     `row_error` | `int` | Row level exception **Optional values are**:  - `0`: Success - `1000`: No permission  |
|     `field_errors` | `map<string, int>` | Column level exception, key is the field name, value is the following enumeration **Optional values ​​are**:  - `1000`: No permission - `2000`: Service exception - `2003`: Field does not exist  | ### Response body example

{
    "code": 0,
    "msg": "Success",
    "data": {
        "departments": [
            {
                "department_id": "H12921",
                "department_count": {
                    "recursive_members_count": "100",
                    "direct_members_count": "100",
                    "recursive_members_count_exclude_leaders": "100",
                    "recursive_departments_count": "100",
                    "direct_departments_count": "100"
                },
                "has_child": true,
                "leaders": [
                    {
                        "leader_type": 1,
                        "leader_id": "U273y71"
                    }
                ],
                "parent_department_id": "H12921",
                "name": {
                    "default_value": "Zhang San",
                    "i18n_value": {
                        "zh_cn": "Chinese",
                        "ja_jp": "ja_jp_name",
                        "en_us": "en_us_name"
                    }
                },
                "enabled_status": true,
                "order_weight": "3000",
                "custom_field_values": [
                    {
                        "field_type": "1",
                        "text_value": {
                            "default_value": "Zhang San",
                            "i18n_value": {
                                "zh_cn": "Chinese",
                                "ja_jp": "ja_jp_name",
                                "en_us": "en_us_name"
                            }
                        },
                        "url_value": {
                            "link_text": {
                                "default_value": "Zhang San",
                                "i18n_value": {
                                    "zh_cn": "Chinese",
                                    "ja_jp": "ja_jp_name",
                                    "en_us": "en_us_name"
                                }
                            },
                            "url": "https://m.bytedance.com/afnasjfna",
                            "pcurl": "http://www.fs.cn"
                        },
                        "enum_value": {
                            "enum_ids": [
                                "1"
                            ],
                            "enum_type": "1"
                        },
                        "user_values": [
                            {
                                "ids": [
                                    "1"
                                ]
                            }
                        ],
                        "phone_value": {
                            "phone_number": "18812345678",
                            "extension_number": "234234234"
                        },
                        "field_key": "C-1000001"
                    }
                ],
                "department_path_infos": [
                    {
                        "department_id": "1",
                        "department_name": {
                            "default_value": "Zhang San",
                            "i18n_value": {
                                "zh_cn": "Chinese",
                                "ja_jp": "ja_jp_name",
                                "en_us": "en_us_name"
                            }
                        }
                    }
                ],
                "data_source": 1
            }
        ],
        "page_response": {
            "has_more": true,
            "page_token": "aihsdniandafoafa"
        },
        "abnormals": [
            {
                "id": "eedasfwe",
                "row_error": 0,
                "field_errors": {
                    "name": "1000"
                }
            }
        ]
    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 2220001 | param is invalid | param is invalid |
| 400 | 2221004 | invalid page token | Invalid pagination token. Please check if the pagination token is correct, or re-obtain a valid page_token. |
| 400 | 2221005 | no page request | No pagination parameters provided, please add valid pagination parameters (page_request). |
| 400 | 2220010 | Exceeded the limit size | Offset pagination is not supported, as it exceeds the limit size. Offset pagination is not supported, please use page_token pagination. |
