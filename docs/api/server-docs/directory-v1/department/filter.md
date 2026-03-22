---
title: "List department"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/directory-v1/department/filter"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/directory/v1/departments/filter"
service: "directory-v1"
resource: "department"
section: "Organization"
rate_limit: "1000 per minute、50 per second"
scopes:
  - "directory:department:list"
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
updated: "1756439303000"
---

# Get department list

This interface is used to obtain the list of eligible department details in batches according to the specified conditions.

> Attention:
> - This API supports tenant_access_token and user_access_token.
> - When using tenant_access_token, data permissions follow the contacts permissions of the application, and the returned field data consists of fields for which the application has permissions. You can view the contacts authorization scope in the Developer Console - Application Details - Permissions Management.
> - When using user_access_token, it defaults to an administrator user, and the administrator's management scope will be verified. When a user has multiple administrator identities and can view department information with all of them, the administrator's management scope will be the maximum set. Administrators can view the Help Center document: [Administrator creates administrator role and assigns permissions](https://www.larksuite.com/hc/zh-CN/articles/360043595213)

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/directory/v1/departments/filter |
| HTTP Method | POST |
| Rate Limit | 1000 per minute、50 per second |
| Supported app types | custom,isv |
| Required scopes | `directory:department:list` |
| Required field scopes | > The response body of the API contains the following sensitive fields, and they will be returned only after corresponding scopes are added. If you do not need the fields, it is not recommended that you request the scopes. `directory:department.base:read` `directory:department.count:read` `directory:department.custom_field:read` `directory:department.data_source:read` `directory:department.department_path:read` `directory:department.external_id:read` `directory:department.has_child:read` `directory:department.leader:read` `directory:department.name:read` `directory:department.order_weight:read` `directory:department.organization:read` `directory:department.parent_id:read` `directory:department.status:read` `directory:employee.base.external_id:read` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `employee_id_type` | `string` | No | User ID categories **Example value**: open_id **Optional values are**:  - `open_id`: Identifies a user to an app. The same user has different Open IDs in different apps. How to get Open ID - `union_id`: Identifies a user to a tenant that acts as a developer. A user has the same Union ID in apps developed by the same developer, and has different Union IDs in apps developed by different developers. A developer can use Union ID to link the same user's identities in multiple apps.How to get Union ID - `employee_id`: The unique identifier of an employee within the enterprise. Supports customization, the system automatically generates it if it is not customized. ID supports modification. How to get employee_id: - Enterprise administrators go to the Admin Console > Organizational Structure > Members and Departments page, click Member Details to query employee IDs - Through the interface of Get employee list in batches, query employee ID by mobile phone number or email.  **Default value**: `open_id` **When the value is `employee_id`, the following field scopes are required**: `directory:employee.base.external_id:read` |
| `department_id_type` | `string` | No | The type of department ID used in this call **Example value**: open_department_id **Optional values are**:  - `open_department_id`: Used to identify a department in a specific application, the same department, the same open_department_id in different applications. - `department_id`: Used to identify a unique department within a tenant  **Default value**: `open_department_id` | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `filter` | `multi_filter_condition` | Yes | Query Conditions [Learn more: Query Conditions Usage] (/ssl: ttdoc/uAjLw4CM/ukTMukTMukTM/directory-v1/filter-usage) |
|   `conditions` | `filter_condition[]` | Yes | List of comparison expressions. Relationships between multiple expressions default to "and" **Data validation rules**: - Length range: `0` ～ `10` |
|     `field` | `string` | Yes | The lvalue of the filter criterion, which is the parameter name of the field. ** Optional filters are: ** - parent_department_id **Example value**: "parent_department_id" |
|     `operator` | `string` | Yes | Comparison operator ** Optional values are: ** - eq: equal to, supports any type of lvalue - in: belongs to any one, does not support parent_department_id, rvalue is an array of multiple target filter values (no more than 100) **Example value**: "eq" |
|     `value` | `string` | Yes | The right value of the filter condition, whose content is the corresponding value type under the combination of the left value field type and the operator. Its value type must be consistent with the value of the query parameter department_id_type, with a maximum length of 64 characters, and supports numbers and letters. When using the parent_department_id condition, the ID of the root department can be "0" **Example value**: "\"0\"" |
| `required_fields` | `string[]` | Yes | List of fields to query. Permissioned row and column data will be returned according to the passed field list. If not passed, no fields will be returned Learn more: Field enumeration instructions **Example value**: ["name"] **Data validation rules**: - Length range: `0` ～ `100` |
| `page_request` | `page_condition` | Yes | paging information |
|   `page_size` | `int` | No | Maximum number of entries returned per page, maximum value is 100 **Default value**：20 **Minimum Value**: 0 **Example value**: 100 |
|   `page_token` | `string` | No | Sequential paging query, can not skip page query, support deep paging, in the need to traverse all the data scene can only use this method. The first pass empty string or not, later pass the last return value in the page_token **Example value**: "iuu14140aknladna91ndas" | ### Request body example

{
    "filter": {
        "conditions": [
            {
                "field": "parent_department_id",
                "operator": "eq",
                "value": "\"0\""
            }
        ]
    },
    "required_fields": [
        "name"
    ],
    "page_request": {
        "page_size": 100,
        "page_token": "iuu14140aknladna91ndas"
    }
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
|     `department_count` | `department_count` | Department member count and sub-department count. Calculation results may be delayed **Required field scopes (Satisfy any)**: `directory:department.count:read` `directory:department.organization:read` |
|       `recursive_members_count` | `string` | number of recursive members |
|       `direct_members_count` | `string` | Number of direct members |
|       `recursive_members_count_exclude_leaders` | `string` | Number of recursive members (excluding leader) |
|       `recursive_departments_count` | `string` | Number of recursive sub-departments |
|       `direct_departments_count` | `string` | Number of direct sub-departments |
|     `has_child` | `boolean` | Whether there sub-departments **Required field scopes (Satisfy any)**: `directory:department.has_child:read` `directory:department.organization:read` |
|     `leaders` | `department_leader[]` | Department head **Required field scopes**: `directory:department.leader:read` |
|       `leader_type` | `int` | Department head type **Optional values are**:  - `1`: main - `2`: vice  |
|       `leader_id` | `string` | Department head ID, consistent with the employee_id_type type |
|     `parent_department_id` | `string` | Parent department ID, consistent with the department_id_type type **Required field scopes (Satisfy any)**: `directory:department.organization:read` `directory:department.parent_id:read` |
|     `name` | `i18n_text` | Department name **Required field scopes (Satisfy any)**: `directory:department.base:read` `directory:department.name:read` |
|       `default_value` | `string` | default value |
|       `i18n_value` | `map<string, string>` | Internationalized value, key is zh_cn, ja_jp, en_us, value is the corresponding value |
|     `enabled_status` | `boolean` | Whether to enable **Required field scopes**: `directory:department.status:read` |
|     `order_weight` | `string` | department ranking weight **Required field scopes (Satisfy any)**: `directory:department.order_weight:read` `directory:department.organization:read` |
|     `custom_field_values` | `custom_field_value[]` | Department custom field values **Required field scopes**: `directory:department.custom_field:read` |
|       `field_type` | `string` | custom field type **Optional values are**:  - `1`: multiline text - `2`: web link - `3`: Enumeration options - `4`: personnel - `9`: - `10`: - `11`:  |
|       `text_value` | `i18n_text` | Text field value |
|         `default_value` | `string` | default value |
|         `i18n_value` | `map<string, string>` | Internationalized value, key is zh_cn, ja_jp, en_us, value is the corresponding value |
|       `url_value` | `url_value` | Web link field value |
|         `link_text` | `i18n_text` | Web page title |
|           `default_value` | `string` | default value |
|           `i18n_value` | `map<string, string>` | Internationalized value, key is zh_cn, ja_jp, en_us, value is the corresponding value |
|         `url` | `string` | Mobile web link |
|         `pcurl` | `string` | Desktop web link |
|       `enum_value` | `enum_value` | enum field value |
|         `enum_ids` | `string[]` | Option Result ID |
|         `enum_type` | `string` | option type **Optional values are**:  - `1`: text - `2`: picture  |
|       `user_values` | `user_value[]` | user field value |
|         `ids` | `string[]` | user ID, consistent with the employee_id_type type |
|       `phone_value` | `phone_value` |  |
|         `phone_number` | `string` |  |
|         `extension_number` | `string` |  |
|       `field_key` | `string` | Custom field key |
|     `department_path_infos` | `department_base_info[]` | Department path information. The sorting order is from root level to last level, excluding the root department. **Required field scopes**: `directory:department.department_path:read` |
|       `department_id` | `string` | Department ID, consistent with the department_id_type type |
|       `department_name` | `i18n_text` | i18n text |
|         `default_value` | `string` | default value |
|         `i18n_value` | `map<string, string>` | Internationalized value, key is zh_cn, ja_jp, en_us, value is the corresponding value |
|     `data_source` | `int` | Data source **Optional values are**:  - `1`: LarkAdmin - `2`: CoreHR - `3`: SCIM  **Required field scopes (Satisfy any)**: `directory:department.base:read` `directory:department.data_source:read` |
|   `page_response` | `page_response` | Paginated results |
|     `has_more` | `boolean` | whether there any follow-up results |
|     `page_token` | `string` | Paging token, when has_more is true, a new page_token will be returned at the same time, otherwise page_token will not be returned. |
|   `abnormals` | `abnormal_record[]` | Abnormal info |
|     `id` | `string` | abnormal id |
|     `row_error` | `int` | row-level abnormal **Optional values are**:  - `0`: success - `1000`: no permission  |
|     `field_errors` | `map<string, int>` | Column-level exceptions, key is the field name, and value is the following enumeration **Optional values are**:  - `1000`: no permissions - `2000`: service exception - `2003`: field does not exist | ### Response body example

{
    "code": 0,
    "msg": "Success",
    "data": {
        "departments": [
            {
                "department_id": "h12921",
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
                        "leader_id": "u273y71"
                    }
                ],
                "parent_department_id": "h12921",
                "name": {
                    "default_value": "Zhang San",
                    "i18n_value": {
                        "zh_cn": "Chinese",
                        "ja_jp": "ja_jp_name",
                        "en_us": "en_us_name"
                    }
                },
                "enabled_status": true,
                "order_weight": "10",
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
                                "zh_cn": "中文",
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
            "page_token": "hiofbsabui214iokncaub"
        },
        "abnormals": [
            {
                "id": "eedasfwe",
                "row_error": 0,
                "field_errors": {
                    "name": 1000
                }
            }
        ]
    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 2221004 | Invalid page token | The pagination token is invalid. Please check whether the page_token parameter is correct or has expired. |
| 400 | 2221005 | No page request | No pagination parameters provided, please pass in a valid page_request parameter. |
| 400 | 2220009 | Filter field is invalid | Invalid filter, please check if the format of the filter parameter meets the requirements. |
| 400 | 2220010 | Exceeded the limit size | The page size exceeds the limit. Please adjust the page_size parameter to within a maximum of 100. |
| 400 | 2220012 | The field is not support filter | The relevant fields do not support filtering. Please use fields that support filtering (such as parent_department_id). |
| 400 | 2220013 | The field does not support the operator | This field does not support this operator. Please use an operator supported by the field (e.g., eq). |
| 400 | 2220014 | Invalid field value | Invalid field right value. Please check if the value of the value parameter meets the requirements of the field type and operator. |
