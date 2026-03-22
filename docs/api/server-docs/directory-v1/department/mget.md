---
title: "Batch get department info"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/directory-v1/department/mget"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/directory/v1/departments/mget"
service: "directory-v1"
resource: "department"
section: "Organization"
rate_limit: "1000 per minute、50 per second"
scopes:
  - "directory:department:read"
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
updated: "1756439299000"
---

# Get departmental information in bulk

This interface supports passing in multiple department IDs and returns detailed information for each department (such as name, person in charge, sub-departments, etc.). 

> **Attention**:
> - This interface supports tenant_access_token and user_access_token.
> - When using tenant_access_token, data permissions follow the application's contacts permissions, and the returned field data is limited to fields for which the application has permissions. You can determine the application's contacts permissions by referring to Get Application Contacts Permissions Configuration.
> - When using user_access_token, it defaults to an administrator user, and the administrator's management scope will be verified. When a user has multiple administrator identities that can all view department information, the administrator's management scope will be the maximum set. Administrators can view the Help Center documentation: [Administrator creates administrator role and assigns permissions](https://www.larksuite.com/hc/zh-CN/articles/360043595213)
> - To enhance the flexibility of the Lark Organization OpenAPI, an update and upgrade were made to this API interface on **October 21, 2024**, with the upgrade content including: optimizing the return data structure for querying information on deleted departments.
> -Before the upgrade, when querying information about a deleted department, the information of the department head was not returned; after the upgrade, when querying information about a deleted department, the returned data will include the information of the department head.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/directory/v1/departments/mget |
| HTTP Method | POST |
| Rate Limit | 1000 per minute、50 per second |
| Supported app types | custom,isv |
| Required scopes | `directory:department:read` |
| Required field scopes | > The response body of the API contains the following sensitive fields, and they will be returned only after corresponding scopes are added. If you do not need the fields, it is not recommended that you request the scopes. `directory:department.base:read` `directory:department.count:read` `directory:department.custom_field:read` `directory:department.data_source:read` `directory:department.department_path:read` `directory:department.external_id:read` `directory:department.has_child:read` `directory:department.leader:read` `directory:department.name:read` `directory:department.order_weight:read` `directory:department.organization:read` `directory:department.parent_id:read` `directory:department.status:read` `directory:employee.base.external_id:read` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `employee_id_type` | `string` | No | User ID categories **Example value**: user_id **Optional values are**:  - `open_id`: Identifies a user to an app. The same user has different Open IDs in different apps. How to get Open ID - `union_id`: Identifies a user to a tenant that acts as a developer. A user has the same Union ID in apps developed by the same developer, and has different Union IDs in apps developed by different developers. A developer can use Union ID to link the same user's identities in multiple apps.How to get Union ID - `employee_id`: The unique identifier of an employee within the enterprise. Supports customization, the system automatically generates it if it is not customized. ID supports modification. How to get employee_id: - Enterprise administrators go to the Admin Console > Organizational Structure > Members and Departments page, click Member Details to query employee IDs - Through the interface of Get employee list in batches, query employee ID by mobile phone number or email.  **Default value**: `open_id` **When the value is `employee_id`, the following field scopes are required**: `directory:employee.base.external_id:read` |
| `department_id_type` | `string` | No | The type of department ID used in this call **Example value**: department_id **Optional values are**:  - `department_id`: Used to identify a unique department within a tenant - `open_department_id`: Used to identify a department in a specific application, the same department, the same open_department_id in different applications.  **Default value**: `open_department_id` | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `department_ids` | `string[]` | Yes | Department ID, consistent with the department_id_type. The way to get the ID: you can query it through the management background. **Example value**: ["adqwea"] **Data validation rules**: - Length range: `1` ～ `100` |
| `required_fields` | `string[]` | Yes | List of fields to query. Permissioned row and column data will be returned according to the passed field list. If not passed, no fields will be returned Learn more: Field enumeration instructions **Example value**: ["name"] **Data validation rules**: - Length range: `0` ～ `100` | ### Request body example

{
    "department_ids": [
        "adqwea"
    ],
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
|       `recursive_members_count` | `string` | number of recursive members Unit: piece |
|       `direct_members_count` | `string` | Number of direct members Unit: piece |
|       `recursive_members_count_exclude_leaders` | `string` | Number of recursive members (excluding leaders) Unit: piece |
|       `recursive_departments_count` | `string` | Number of recursive subdepartments Unit: piece |
|       `direct_departments_count` | `string` | Number of direct sub-departments Unit: piece |
|     `has_child` | `boolean` | Is there a sub-department? **Required field scopes (Satisfy any)**: `directory:department.has_child:read` `directory:department.organization:read` |
|     `leaders` | `department_leader[]` | department head **Required field scopes**: `directory:department.leader:read` |
|       `leader_type` | `int` | Type of department head **Optional values are**:  - `1`: main - `2`: vice  |
|       `leader_id` | `string` | Department head ID, consistent with the employee_id_type type |
|     `parent_department_id` | `string` | Parent Department ID, consistent with the department_id_type type **Required field scopes (Satisfy any)**: `directory:department.organization:read` `directory:department.parent_id:read` |
|     `name` | `i18n_text` | Department name **Required field scopes (Satisfy any)**: `directory:department.base:read` `directory:department.name:read` |
|       `default_value` | `string` | default value |
|       `i18n_value` | `map<string, string>` | Internationalized value, key is zh_cn, ja_jp, en_us, value is the corresponding value |
|     `enabled_status` | `boolean` | Whether to enable **Required field scopes**: `directory:department.status:read` |
|     `order_weight` | `string` | department ranking weight **Required field scopes (Satisfy any)**: `directory:department.order_weight:read` `directory:department.organization:read` |
|     `custom_field_values` | `custom_field_value[]` | Department custom field values **Required field scopes**: `directory:department.custom_field:read` |
|       `field_key` | `string` | Custom field key |
|       `field_type` | `string` | custom field type **Optional values are**:  - `1`: mutli-line text - `2`: web link - `3`: enum option - `4`: generic user  |
|       `text_value` | `i18n_text` | text field value |
|         `default_value` | `string` | default value |
|         `i18n_value` | `map<string, string>` | Internationalized value, key is zh_cn, ja_jp, en_us, value is the corresponding value |
|       `url_value` | `url_value` | web link field value |
|         `link_text` | `i18n_text` | web page title |
|           `default_value` | `string` | default value |
|           `i18n_value` | `map<string, string>` | Internationalized value, key is zh_cn, ja_jp, en_us, value is the corresponding value |
|         `url` | `string` | mobile web link |
|         `pcurl` | `string` | desktop web link |
|       `enum_value` | `enum_value` | enum field value |
|         `enum_ids` | `string[]` | Option Result ID |
|         `enum_type` | `string` | option type **Optional values are**:  - `1`: text - `2`: picture  |
|       `user_values` | `user_value[]` | user field value |
|         `ids` | `string[]` | user id, consistent with the employee_id_type type |
|     `department_path_infos` | `department_base_info[]` | Department path information. The sorting order is from root level to last level, excluding the root department. **Required field scopes**: `directory:department.department_path:read` |
|       `department_id` | `string` | Department ID, consistent with the department_id_type type |
|       `department_name` | `i18n_text` | Department name |
|         `default_value` | `string` | default value |
|         `i18n_value` | `map<string, string>` | Internationalized value, key is zh_cn, ja_jp, en_us, value is the corresponding value |
|     `data_source` | `int` | Data Sources **Optional values are**:  - `1`: LarkAdmin - `2`: CoreHR - `3`: SCIM  **Required field scopes (Satisfy any)**: `directory:department.base:read` `directory:department.data_source:read` |
|   `abnormals` | `abnormal_record[]` | Field abnormal info |
|     `id` | `string` | Abnormal ID |
|     `row_error` | `int` | Row level abnormal **Optional values are**:  - `0`: success - `1000`: No permission  |
|     `field_errors` | `map<string, int>` | Column level abnormal, key is field name, value is following enum.  **Optional values are**：  - `1000`: No permission - `2000`: Service exception - `2003`: Field not exist  | ### Response body example

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
                "order_weight": "none",
                "custom_field_values": [
                    {
                        "field_key": "C-1000001",
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
                        ]
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
| 400 | 2220001 | param is invalid | Invalid request parameters. Please check if the request parameters meet the requirements of the documentation, or refer to the parameter description document to confirm the parameter format. |
