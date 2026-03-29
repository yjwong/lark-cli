---
title: "Resource Introduction"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/directory-v1/department/resource-introduction"
service: "directory-v1"
resource: "department"
section: "Organization"
field_scopes:
  - "directory:department.base:read"
  - "directory:department.external_id:read"
  - "directory:department.count:read"
  - "directory:department.organization:read"
  - "directory:department.has_child:read"
  - "directory:department.leader:read"
  - "directory:department.parent_id:read"
  - "directory:department.name:read"
  - "directory:department.status:read"
  - "directory:department.order_weight:read"
  - "directory:department.custom_field:read"
  - "directory:department.department_path:read"
  - "directory:department.data_source:read"
updated: "1756439282000"
---

# Resource definition
A department is a base entity in the Lark organizational structure, with each employee belonging to one or more departments.
Department identifiers in Lark include `department_id`, `open_department_id`. For more information about the various types of IDs for departments, refer to the 
Department ID。

# Field description

| 名称 | 类型 | 描述 |
| --- | --- | --- |
| `department_id` | `string` | Department ID **Required field scopes (Satisfy any)**: `directory:department.base:read` `directory:department.external_id:read` |
| `department_count` | `department_count` | Department member count and subdepartment count. There may be a delay in the calculation result. **Required field scopes (Satisfy any)**: `directory:department.count:read` `directory:department.organization:read` |
|   `recursive_members_count` | `string` | The number of department members, including the number of members of all levels of sub-departments within the department |
|   `direct_members_count` | `string` | The number of department members, including only the number of direct members, excluding the number of sub - department members |
|   `recursive_members_count_exclude_leaders` | `string` | The number of department members, including the number of members of all levels of sub-departments within the department |
|   `recursive_departments_count` | `string` | The number of department members, including the number of members of all levels of sub-departments within the department |
|   `direct_departments_count` | `string` | The number of department members, including only the number of direct members, excluding the number of sub - department members |
| `has_child` | `boolean` | Is there a sub-department? **Required field scopes (Satisfy any)**: `directory:department.has_child:read` `directory:department.organization:read` |
| `leaders` | `department_leader[]` | department head **Required field scopes**: `directory:department.leader:read` |
|   `leader_type` | `int` | Type of department head **Optional values are**:  - `1`: main - `2`: vice  |
|   `leader_id` | `string` | Department head ID |
| `parent_department_id` | `string` | Parent Department ID **Required field scopes (Satisfy any)**: `directory:department.organization:read` `directory:department.parent_id:read` |
| `name` | `i18n_text` | Department name **Required field scopes (Satisfy any)**: `directory:department.base:read` `directory:department.name:read` |
|   `default_value` | `string` | default value |
|   `i18n_value` | `map<string, string>` | Internationalized value, key is zh_cn, ja_jp, en_us, value is the corresponding value |
| `enabled_status` | `boolean` | Whether to enable **Required field scopes**: `directory:department.status:read` |
| `order_weight` | `string` | department ranking weight **Required field scopes (Satisfy any)**: `directory:department.order_weight:read` `directory:department.organization:read` |
| `custom_field_values` | `custom_field_value[]` | Department custom field values **Required field scopes**: `directory:department.custom_field:read` |
|   `field_key` | `string` | Custom field key |
|   `field_type` | `string` | custom field type **Optional values are**:  - `1`: mutli-line text - `2`: web link - `3`: enum option - `4`: generic user  |
|   `text_value` | `i18n_text` | text field value |
|     `default_value` | `string` | default value |
|     `i18n_value` | `map<string, string>` | Internationalized value, key is zh_cn, ja_jp, en_us, value is the corresponding value |
|   `url_value` | `url_value` | web link field value |
|     `link_text` | `i18n_text` | web page title |
|       `default_value` | `string` | default value |
|       `i18n_value` | `map<string, string>` | Internationalized value, key is zh_cn, ja_jp, en_us, value is the corresponding value |
|     `url` | `string` | mobile web link |
|     `pcurl` | `string` | desktop web link |
|   `enum_value` | `enum_value` | enum field value |
|     `enum_ids` | `string[]` | Option Result ID |
|     `enum_type` | `string` | option type **Optional values are**:  - `1`: text - `2`: picture  |
|   `user_values` | `user_value[]` | user field value |
|     `ids` | `string[]` | user id |
| `department_path_infos` | `department_base_info[]` | Department path information. The sorting order is from root level to last level, excluding the root department. **Required field scopes**: `directory:department.department_path:read` |
|   `department_id` | `string` | Department ID |
|   `department_name` | `i18n_text` | Department name |
|     `default_value` | `string` | default value |
|     `i18n_value` | `map<string, string>` | Internationalized value, key is zh_cn, ja_jp, en_us, value is the corresponding value |
| `data_source` | `int` | Data Sources **Optional values are**:  - `1`: LarkAdmin - `2`: CoreHR - `3`: SCIM  **Required field scopes (Satisfy any)**: `directory:department.base:read` `directory:department.data_source:read` | ## Data sample

```json 
 {
    "department_id": "D100",
    "open_department_id": "od-4e6ac4d14bcd5071a37a39de902c7141",
    "name": {
        "value": "销售部",
        "i18n_value":{
            "language": "en_us"
                "value": "Sale"
            }
        },
    "parent_department_id": "D90",
    "leaders": [
        {
            "leader_ID": "u273y69",
            "DepartmentLeaderType": 1
        }
    ],
    "has_child":true,
    "department_count": {
        "recursive_members_count": 100,
        "direct_members_count": 90,
        "recursive_members_count_exclue_leaders":80,
        "recursive_departments_count": 20,
        "direct_departments_count": 10
    },
    "enabled_status": true,
    "order_weight": 1,
    "department_path_info": [
        {
            "department_id": "D102",
            "department_name":{
                "value": "北京科技有限公司",
                "i18n_value":{
                    "language": "en_us"
                        "value": "Beijing Technology Co., Ltd"
                }
            }
        },
        {
            "department_id": "D101",
            "department_name":{
                "value": "华北大区",
                "i18n_value":{
                    "language": "en_us"
                        "value": "North China Region"
                }
            }
        },
        {
            "department_id": "D100",
            "department_name":{
                "value": "销售部",
                "i18n_value":{
                    "language": "en_us"
                        "value": "Sale"
                }
            }
        }
    ],
    "custom_field_values": [
        {
            "field_id": "DemoId",
            "field_value": {
                "field_type": 1,
                "text_value": "DemoText"
            }
        },
        {
            "field_id": "DemoId_2",
            "field_value": {
                "field_type": 4,
                "user_value": [
                    "9b2fabg5"
                ]
            }
        }
    ],
    "data_source":1   
}

```
