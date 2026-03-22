---
title: "Resources Introduction"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/directory-v1/employee/resources-introduction"
service: "directory-v1"
resource: "employee"
section: "Organization"
field_scopes:
  - "directory:employee.base.base:read"
  - "directory:employee.base.name.name:read"
  - "directory:employee.base.name.another_name:read"
  - "directory:employee.base.mobile:read"
  - "directory:employee.base.email:read"
  - "directory:employee.base.enterprise_email:read"
  - "directory:employee.base.gender:read"
  - "directory:employee.base.department:read"
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
  - "directory:employee.base.dept_order:read"
  - "directory:employee.base.description:read"
  - "directory:employee.base.active_status:read"
  - "directory:employee.base.status:read"
  - "directory:employee.base.is_resigned:read"
  - "directory:employee.base.leader:read"
  - "directory:employee.base.leader_id:read"
  - "directory:employee.base.dotted_line_leaders:read"
  - "directory:employee.base.is_primary_admin:read"
  - "directory:employee.base.role:read"
  - "directory:employee.base.enterprise_email_alias:read"
  - "directory:employee.base.custom_field:read"
  - "directory:employee.base.department_path:read"
  - "directory:employee.base.resign_time:read"
  - "directory:employee.base.avatar:read"
  - "directory:employee.base.background_image:read"
  - "directory:employee.base.is_admin:read"
  - "directory:employee.base.data_source:read"
  - "directory:employee.base.geo:read"
  - "directory:employee.base.subscription_ids:read"
  - "directory:employee.work.base_work:read"
  - "directory:employee.work.work_country_or_region:read"
  - "directory:employee.work.work_place:read"
  - "directory:place.base:read"
  - "directory:place.status:read"
  - "directory:employee.work.work_station:read"
  - "directory:employee.work.job_number:read"
  - "directory:employee.work.extension_number:read"
  - "directory:employee.work.join_date:read"
  - "directory:employee.work.employment:read"
  - "directory:employee.work.employment_type:read"
  - "directory:employee.work.staff_status:read"
  - "directory:employee.work.job_title:read"
  - "directory:job_title.base:read"
  - "directory:job_title.status:read"
  - "directory:employee.work.resign_date:read"
  - "directory:employee.work.resign_reason:read"
  - "directory:employee.work.resign_remark:read"
  - "directory:employee.work.resign_type:read"
updated: "1756439240000"
---

# Resource definition
Employee means a member of the Lark enterprise with the identity `Employee`, which is equivalent to `User` in the Address Book OpenAPI.

Employee identifiers in Lark include `employee_id`, `open_id` and `union_id`, where the value of `employee_id` is equivalent to `user_id` in Lark Contact API, and the other two are also the same as the value of User in Contact. For more information about the various types of employee IDs, refer to theUser identity overview。

The logic of the Employee Lifecycle flow is aligned with the Contact's User, see below:

![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/3d1df8724712e1ffc99aa3aaa07d1759_0o9r5pTnJj.png?height=679&lazyload=true&width=1201)

# Field description

| 名称 | 类型 | 描述 |
| --- | --- | --- |
| `base_info` | `employee_base_entity` | Staff basic information |
|   `employee_id` | `string` | Unique identification of active employees within the enterprise.  **Field permission requirements**: `directory:employee.base.external_id:read` |
|   `name` | `name` | name |
|     `name` | `i18n_text` | Employee's name **Required field scopes (Satisfy any)**: `directory:employee.base.base:read` `directory:employee.base.name.name:read` |
|       `default_value` | `string` | default value |
|       `i18n_value` | `map<string, string>` | Internationalized value, key is zh_cn, ja_jp, en_us, value is the corresponding value |
|       `default_locale` | `string` | Default language, values include zh_cn, ja_jp, en_us, etc |
|     `another_name` | `string` | alias **Required field scopes (Satisfy any)**: `directory:employee.base.base:read` `directory:employee.base.name.another_name:read` |
|   `mobile` | `string` | Employee's mobile phone number **Required field scopes**: `directory:employee.base.mobile:read` |
|   `email` | `string` | Employee's contact email address at work **Required field scopes**: `directory:employee.base.email:read` |
|   `enterprise_email` | `string` | Employee's business email address **Required field scopes**: `directory:employee.base.enterprise_email:read` |
|   `gender` | `int` | Gender **Optional values are**:  - `0`: Unknown - `1`: Male - `2`: Female - `3`: Other  **Required field scopes**: `directory:employee.base.gender:read` |
|   `departments` | `department[]` | Information about the department to which the employee belongs. An employee can belong to multiple departments, with the first in line being the main department **Required field scopes**: `directory:employee.base.department:read` |
|     `department_id` | `string` | Department ID **Required field scopes (Satisfy any)**: `directory:department.base:read` `directory:department.external_id:read` |
|     `department_count` | `department_count` | Department members and sub-department counts **Required field scopes (Satisfy any)**: `directory:department.count:read` `directory:department.organization:read` |
|       `recursive_members_count` | `string` | Number of recursive members |
|       `direct_members_count` | `string` | Number of direct members |
|       `recursive_members_count_exclude_leaders` | `string` | Number of recursive members (excluding leaders) |
|       `recursive_departments_count` | `string` | The number of recursive subdepartments |
|       `direct_departments_count` | `string` | Number of direct sub-departments |
|     `has_child` | `boolean` | Whether there are sub department **Required field scopes (Satisfy any)**: `directory:department.has_child:read` `directory:department.organization:read` |
|     `leaders` | `department_leader[]` | Head of Department **Required field scopes**: `directory:department.leader:read` |
|       `leader_type` | `int` | Type of department head **Optional values are**:  - `1`: Main - `2`: Deputy  |
|       `leader_id` | `string` | Leader ID of the Department |
|     `parent_department_id` | `string` | Parent department ID **Required field scopes (Satisfy any)**: `directory:department.organization:read` `directory:department.parent_id:read` |
|     `name` | `i18n_text` | Department name **Required field scopes (Satisfy any)**: `directory:department.base:read` `directory:department.name:read` |
|       `default_value` | `string` | default value |
|       `i18n_value` | `map<string, string>` | Internationalized value, key is zh_cn, ja_jp, en_us, value is the corresponding value |
|       `default_locale` | `string` | Default language, values include zh_cn, ja_jp, en_us, etc |
|     `enabled_status` | `boolean` | Whether to enable **Required field scopes**: `directory:department.status:read` |
|     `order_weight` | `string` | Department sort weights **Required field scopes (Satisfy any)**: `directory:department.order_weight:read` `directory:department.organization:read` |
|     `custom_field_values` | `custom_field_value[]` | Custom fields **Required field scopes**: `directory:department.custom_field:read` |
|       `field_key` | `string` | Custom field key |
|       `field_type` | `string` | Custom field types **Optional values are**:  - `1`: Multi-line text - `2`: Web Links - `3`: Enumeration options - `4`: Employee  |
|       `text_value` | `i18n_text` | Text field value |
|           `default_value` | `string` | default value |
|           `i18n_value` | `map<string, string>` | Internationalized value, key is zh_cn, ja_jp, en_us, value is the corresponding value |
|           `default_locale` | `string` | Default language, values include zh_cn, ja_jp, en_us, etc |
|       `url_value` | `url_value` | Webpage link field value |
|           `link_text` | `i18n_text` | The title of the page |
|           `default_value` | `string` | default value |
|           `i18n_value` | `map<string, string>` | Internationalized value, key is zh_cn, ja_jp, en_us, value is the corresponding value |
|           `default_locale` | `string` | Default language, values include zh_cn, ja_jp, en_us, etc |
|           `url` | `string` | Links to mobile web pages |
|           `pcurl` | `string` | Desktop web links |
|       `enum_value` | `enum_value` | Enumeration values |
|           `enum_ids` | `string[]` | Option result ID |
|           `enum_type` | `string` | Option type **Optional values are**:  - `1`: Text type - `2`: Image Type  |
|       `user_values` | `user_value[]` | Employee Type |
|           `ids` | `string[]` | Employee ids |
|     `department_path_infos` | `department_base_info[]` | Department path information. The order is from the root department to the last department **Required field scopes**: `directory:department.department_path:read` |
|       `department_id` | `string` | Department ID |
|       `department_name` | `i18n_text` | Department name |
|           `default_value` | `string` | default value |
|           `i18n_value` | `map<string, string>` | Internationalized value, key is zh_cn, ja_jp, en_us, value is the corresponding value |
|           `default_locale` | `string` | Default language, values include zh_cn, ja_jp, en_us, etc |
|     `data_source` | `int` | data source **Optional values are**:  - `1`: LarkAdmin - `2`: CoreHR - `3`: SCIM  **Required field scopes (Satisfy any)**: `directory:department.base:read` `directory:department.data_source:read` |
|   `employee_order_in_departments` | `user_department_sort_info[]` | The ranking information of users within the department, the first department is the main department **Required field scopes (Satisfy any)**: `directory:employee.base.department:read` `directory:employee.base.dept_order:read` |
|     `department_id` | `string` | Department ID |
|     `order_weight_in_deparment` | `string` | The ranking weight of users within the department |
|     `order_weight_among_deparments` | `string` | User ranking weights across multiple departments |
|   `description` | `string` | personal signature **Required field scopes (Satisfy any)**: `directory:employee.base.base:read` `directory:employee.base.description:read` |
|   `active_status` | `int` | user active status **Optional values are**:  - `1`: not activated - `2`: activate - `3`: freeze - `4`: active exit - `5`: not joined  **Required field scopes (Satisfy any)**: `directory:employee.base.active_status:read` `directory:employee.base.status:read` |
|   `is_resigned` | `boolean` | whether to quit **Required field scopes (Satisfy any)**: `directory:employee.base.is_resigned:read` `directory:employee.base.status:read` |
|   `leader_id` | `string` | Direct manager ID **Required field scopes (Satisfy any)**: `directory:employee.base.leader:read` `directory:employee.base.leader_id:read` |
|   `dotted_line_leader_ids` | `string[]` | The ID of the dotted line leader **Required field scopes (Satisfy any)**: `directory:employee.base.dotted_line_leaders:read` `directory:employee.base.leader:read` |
|   `is_primary_admin` | `boolean` | Whether is a super administrator **Required field scopes (Satisfy any)**: `directory:employee.base.is_primary_admin:read` `directory:employee.base.role:read` |
|   `enterprise_email_aliases` | `string[]` | The employee's enterprise email alias. When members want to use different email addresses for different people to send emails, they can add aliases to their enterprise email addresses. Once added, members can send messages using aliases and accept messages sent to aliases **Required field scopes (Satisfy any)**: `directory:employee.base.enterprise_email:read` `directory:employee.base.enterprise_email_alias:read` |
|   `custom_field_values` | `custom_field_value[]` | Custom field values **Required field scopes**: `directory:employee.base.custom_field:read` |
|     `field_key` | `string` | Custom field key |
|     `field_type` | `string` | Custom field types **Optional values are**:  - `1`: Multi-line text - `2`: Web Links - `3`: Enumeration options - `4`: Employee  |
|     `text_value` | `i18n_text` | Text field value |
|       `default_value` | `string` | default value |
|       `i18n_value` | `map<string, string>` | Internationalized value, key is zh_cn, ja_jp, en_us, value is the corresponding value |
|       `default_locale` | `string` | Default language, values include zh_cn, ja_jp, en_us, etc |
|     `url_value` | `url_value` | Webpage link field value |
|       `link_text` | `i18n_text` | The title of the page |
|           `default_value` | `string` | default value |
|           `i18n_value` | `map<string, string>` | Internationalized value, key is zh_cn, ja_jp, en_us, value is the corresponding value |
|           `default_locale` | `string` | Default language, values include zh_cn, ja_jp, en_us, etc |
|       `url` | `string` | Links to mobile web pages |
|       `pcurl` | `string` | Desktop web links |
|     `enum_value` | `enum_value` | Enumeration values |
|       `enum_ids` | `string[]` | Option result ID |
|       `enum_type` | `string` | Option type **Optional values are**:  - `1`: text - `2`: Image  |
|     `user_values` | `user_value[]` | Employee field value |
|       `ids` | `string[]` | Employee ID |
|   `department_path_infos` | `department_base_info[][]` | Department path information. The order is from the root department to the last department **Required field scopes**: `directory:employee.base.department_path:read` |
|   `resign_time` | `string` | offboarding timestamp **Required field scopes**: `directory:employee.base.resign_time:read` |
|   `avatar` | `image_link` | employee avatar **Required field scopes (Satisfy any)**: `directory:employee.base.avatar:read` `directory:employee.base.base:read` |
|     `avatar_72` | `string` | 72*72 pixel avatar link |
|     `avatar_240` | `string` | 240*240 pixel avatar link |
|     `avatar_640` | `string` | 640*640 pixel avatar link |
|     `avatar_origin` | `string` | Original avatar link |
|   `background_image` | `string` | Custom background image URL **Required field scopes (Satisfy any)**: `directory:employee.base.background_image:read` `directory:employee.base.base:read` |
|   `is_admin` | `boolean` | Whether is a normal administrator **Required field scopes (Satisfy any)**: `directory:employee.base.is_admin:read` `directory:employee.base.role:read` |
|   `data_source` | `int` | data source **Optional values are**:  - `1`: admin background - `2`: Personnel Enterprise Edition - `3`: SCIM  **Required field scopes (Satisfy any)**: `directory:employee.base.base:read` `directory:employee.base.data_source:read` |
|   `geo_name` | `string` | Where employee data resides **Required field scopes (Satisfy any)**: `directory:employee.base.base:read` `directory:employee.base.geo:read` |
|   `subscription_ids` | `string[]` | A list of seat IDs assigned to the employee. **Required field scopes (Satisfy any)**: `directory:employee.base.base:read` `directory:employee.base.subscription_ids:read` |
| `work_info` | `employee_work_entity` | Employee work information |
|   `work_country_or_region` | `string` | Country/region of work How to query the meaning of the code of the country **Required field scopes (Satisfy any)**: `directory:employee.work.base_work:read` `directory:employee.work.work_country_or_region:read` |
|   `work_place` | `place` | Work place **Required field scopes (Satisfy any)**: `directory:employee.work.base_work:read` `directory:employee.work.work_place:read` |
|     `place_id` | `string` | ID, which returns "0" by default if it does not exist. |
|     `place_name` | `i18n_text` | Name of the work location **Required field scopes**: `directory:place.base:read` |
|       `default_value` | `string` | default value |
|       `i18n_value` | `map<string, string>` | Internationalized value, key is zh_cn, ja_jp, en_us, value is the corresponding value |
|       `default_locale` | `string` | Default language, values include zh_cn, ja_jp, en_us, etc |
|     `is_enabled` | `boolean` | Whether it is enabled **Required field scopes**: `directory:place.status:read` |
|     `description` | `i18n_text` | description **Required field scopes**: `directory:place.base:read` |
|       `default_value` | `string` | default value |
|       `i18n_value` | `map<string, string>` | Internationalized value, key is zh_cn, ja_jp, en_us, value is the corresponding value |
|       `default_locale` | `string` | Default language, values include zh_cn, ja_jp, en_us, etc |
|   `work_station` | `i18n_text` | Work station **Required field scopes (Satisfy any)**: `directory:employee.work.base_work:read` `directory:employee.work.work_station:read` |
|     `default_value` | `string` | default value |
|     `i18n_value` | `map<string, string>` | Internationalized value, key is zh_cn, ja_jp, en_us, value is the corresponding value |
|     `default_locale` | `string` | Default language, values include zh_cn, ja_jp, en_us, etc |
|   `job_number` | `string` | Work station **Required field scopes (Satisfy any)**: `directory:employee.work.base_work:read` `directory:employee.work.job_number:read` |
|   `extension_number` | `string` | Zhang San **Required field scopes (Satisfy any)**: `directory:employee.work.base_work:read` `directory:employee.work.extension_number:read` |
|   `join_date` | `string` | {"zh_cn":"Zhang San"} **Required field scopes (Satisfy any)**: `directory:employee.work.join_date:read` `directory:employee.work.employment:read` |
|   `employment_type` | `int` | 员工类型 **Optional values are**:  - `0`: Unknown - `1`: Regular - `2`: Intern - `3`: Outsourcing - `4`: Contractor - `5`: Consultant  **Required field scopes (Satisfy any)**: `directory:employee.work.employment_type:read` `directory:employee.work.employment:read` |
|   `staff_status` | `int` | Employee staff status **Optional values are**:  - `1`: Unresigned - `2`: Resigned - `3`: Pre onboarding - `4`: Cancel onboarding - `5`: Prepared to quit  **Required field scopes (Satisfy any)**: `directory:employee.work.staff_status:read` `directory:employee.work.employment:read` |
|   `job_title` | `job_title` | Job title **Required field scopes**: `directory:employee.work.job_title:read` |
|     `job_title_id` | `string` | ID, which returns "0" by default if it does not exist. |
|     `job_title_name` | `i18n_text` | Job title name **Required field scopes**: `directory:job_title.base:read` |
|       `default_value` | `string` | default value |
|       `i18n_value` | `map<string, string>` | Internationalized value, key is zh_cn, ja_jp, en_us, value is the corresponding value |
|       `default_locale` | `string` | Default language, values include zh_cn, ja_jp, en_us, etc |
|     `is_enabled` | `boolean` | Whether it is enabled **Required field scopes**: `directory:job_title.status:read` |
|     `description` | `i18n_text` | description **Required field scopes**: `directory:job_title.base:read` |
|       `default_value` | `string` | default value |
|       `i18n_value` | `map<string, string>` | Internationalized value, key is zh_cn, ja_jp, en_us, value is the corresponding value |
|       `default_locale` | `string` | Default language, values include zh_cn, ja_jp, en_us, etc |
|   `resign_date` | `string` | Date of separation 2007-03-20 **Required field scopes (Satisfy any)**: `directory:employee.work.resign_date:read` `directory:employee.work.employment:read` |
|   `resign_reason` | `string` | Reason for leaving  **Example value**: "1" **Optional values are**:  - `1`: Salary does not meet expectations - `2`: Working hours are too long - `3`: Dissatisfied with the job content - `4`: Does not approve of superiors or management - `5`: Limited career opportunities - `6`: Lack of identification with company culture - `7`: Organizational restructuring (voluntary departure) - `8`: Expiration of contract - `9`: Job hopping - `10`: Transfer line - `11`: Family reasons - `12`: Poor health - `13`: Reason for duty station - `14`: Other (voluntary departure) - `15`: Unexpected - `16`: Death - `17`: Dismissal - `18`: probation period not passed - `19`: Poor job performance - `20`: Work output low - `21`: Organizational restructuring (passive departure) - `22`: Discipline - `23`: Illegal - `24`: Other (passive departure) - `25`: Other (other)  **Required field scopes (Satisfy any)**: `directory:employee.work.resign_reason:read` `directory:employee.work.employment:read` |
|   `resign_remark` | `string` | Resignation remarks **Required field scopes (Satisfy any)**: `directory:employee.work.resign_remark:read` `directory:employee.work.employment:read` |
|   `resign_type` | `string` | Resign Type  **Example Value**: "1" **Optional values are**:  - `1`: Active - `2`: Passive - `3`: Other  **Required field scopes (Satisfy any)**: `directory:employee.work.resign_type:read` `directory:employee.work.employment:read` | ## Data sample

```json 
 {
    "base_info": {
        "employee_id": "u273y71",
        "name": {
            "last_name": {
                "value": "王"，
                "i18n_value": [
                    {
                        "language": "zh_cn"
                        "value": "王"
                    },
                    {
                        "language": "en_us"
                        "value": "Wang"
                    }
                ],
                "default_locale":"zh_cn"                
            },
            "first_name": {
                "value": "小明"，
                "i18n_value": [
                    {
                        "language": "zh_cn"
                        "value": "小明"
                    },
                    {
                        "language": "en_us"
                        "value": "Xiaoming"
                    }
                ],
                "default_locale":"zh_cn"                
            },
            "name": {
                "value": "王小明"，
                "i18n_value": [
                    {
                        "language": "zh_cn"
                        "value": "王小明"
                    },
                    {
                        "language": "en_us"
                        "value": "Wang Xiaoming"
                    }
                ],
                "default_locale":"zh_cn"                
            },
            "another_name": {
                "value": "王明"
            }
        },
        "avatar_key": "2500c7a9-5fff-4d9a-a2de-3d59614ae28g",
        "background_image_key": "2500c7a9-5fff-4d9a-a2de-3d59614ae28g",
        "description": "新成员请多关照",
        "mobile": "+8613011111111",
        "email": "zhangsan@gmail.com",
        "enterprise_email": "zhangsan@gmail.com",
        "enterprise_email_alias": [
            "zhangsan_1@gmail.com",
            "zhangsan_2@gmail.com"
        ],
        "gender": 1,
        "departments":[
            {
                "department_id": "D100",
                "name": {
                    "value": "销售部",
                    "i18n_value":{
                        "language": "en_us"
                        "value": "Sale"
                    },
                    "default_locale":"zh_cn"
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
                "order_weight": 1
            }
        ],
        "employee_order_in_departments": [
            {
                "department_id": "D100",
                "order_weight_in_deparment": 100,
                "order_weight_among_deparments": 1
            }
        ],
        "department_path_infos": [
            {
                "department_id": "D100",
                "path_infos": [
                    {
                        "department_id": "h121921",
                        "department_name":{
                            "value": "根部门"
                            "i18n_value":{
                                "language": "en_us"
                                "value": "root"
                            },
                            "default_locale":"zh_cn"
                        }
                    },
                    {
                        "department_id": "D100",
                        "department_name":{
                            "value": "子部门"
                            "i18n_value":{
                                "language": "en_us"
                                "value": "leaf"
                            },
                            "default_locale":"zh_cn"
                        }
                    }                    
                ]
            }
        ],
        "leader_id": "2e1cf73b",
        "dotted_line_leader_ids": [
            "2b1ae26b"
        ],
        "active_status": 1,
        "is_resigned": false,
        "is_primary_admin": false,
        "is_admin": true,
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
        "resign_time": 348970
        "data_source":1,
        "geo_name":"cn",
        "subscription_ids":[
            "23213213213123123"
        ]
    },
    "work_info": {        
        "work_country_or_region": "MDCT00000040",
        "work_place": {
            "place_id": "mga5oa8ayjlp9rb", 
            "place_name": {
                "value": "北京市"
            },
            "is_enable": true,
            "description": "主要的工作地点"
        },
        "work_station": {
            "value": "北楼-H34"，
            "i18n_value": {
                "language": "en_us",
                "value": "North-H34"
            },
            "default_locale":"zh_cn"
        },
        "job_number": "2845435",
        "extension_number": "080", 
        "join_date": "2007-03-20",
        "employment_type": 2,
        "staff_status": 1,
        "positions": [
            {
                "position_code": "fga7b946",
                "position_name": "销售",
                "is_main_position": true,
                "department_id": "D100"
                "leader_id": "2e1cf73b",
                "leader_position_code": "e71b94gb"
            }
        ],
        "job_title": {
            "job_title_id": "mga5oa8ayjlp9rb",
            "job_title_name": {
                "value": "销售"
            },
            "is_enable": true,
            "description": {
                "value": "销售岗位"
            }
        },
    }
}

```
