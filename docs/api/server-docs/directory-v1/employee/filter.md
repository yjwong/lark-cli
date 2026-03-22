---
title: "Get employee list in batches"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/directory-v1/employee/filter"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/directory/v1/employees/filter"
service: "directory-v1"
resource: "employee"
section: "Organization"
rate_limit: "1000 per minute、50 per second"
scopes:
  - "directory:employee:list"
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
  - "directory:employee.base.active_status:read"
  - "directory:employee.base.avatar:read"
  - "directory:employee.base.background_image:read"
  - "directory:employee.base.base:read"
  - "directory:employee.base.custom_field:read"
  - "directory:employee.base.data_source:read"
  - "directory:employee.base.department:read"
  - "directory:employee.base.department_path:read"
  - "directory:employee.base.dept_order:read"
  - "directory:employee.base.description:read"
  - "directory:employee.base.dotted_line_leaders:read"
  - "directory:employee.base.email:read"
  - "directory:employee.base.enterprise_email:read"
  - "directory:employee.base.enterprise_email_alias:read"
  - "directory:employee.base.external_id:read"
  - "directory:employee.base.gender:read"
  - "directory:employee.base.geo:read"
  - "directory:employee.base.is_admin:read"
  - "directory:employee.base.is_primary_admin:read"
  - "directory:employee.base.is_resigned:read"
  - "directory:employee.base.leader:read"
  - "directory:employee.base.leader_id:read"
  - "directory:employee.base.mobile:read"
  - "directory:employee.base.name.another_name:read"
  - "directory:employee.base.name.name:read"
  - "directory:employee.base.resign_time:read"
  - "directory:employee.base.role:read"
  - "directory:employee.base.status:read"
  - "directory:employee.base.subscription_ids:read"
  - "directory:employee.work.base_work:read"
  - "directory:employee.work.employment_type:read"
  - "directory:employee.work.extension_number:read"
  - "directory:employee.work.job_number:read"
  - "directory:employee.work.job_title:read"
  - "directory:employee.work.join_date:read"
  - "directory:employee.work.resign_date:read"
  - "directory:employee.work.resign_reason:read"
  - "directory:employee.work.resign_remark:read"
  - "directory:employee.work.resign_type:read"
  - "directory:employee.work.staff_status:read"
  - "directory:employee.work.work_country_or_region:read"
  - "directory:employee.work.work_place:read"
  - "directory:employee.work.work_station:read"
  - "directory:job_title.base:read"
  - "directory:job_title.status:read"
  - "directory:place.base:read"
  - "directory:place.status:read"
  - "directory:job_family.path:read"
  - "directory:job_family.status:read"
  - "directory:job_level.base:read"
  - "directory:job_level.order:read"
  - "directory:job_level.status:read"
  - "directory:employee.work.job_level:read"
  - "directory:employee.work.employment:read"
  - "directory:employee.work.job_family:read"
  - "directory:job_family.base:read"
updated: "1756439273000"
---

# Get the employee list in bulk

This interface is used to obtain a batch list of eligible employee details according to the specified conditions.
Employee refers to a member of Lark's enterprise who is identified as "Employee", which is equivalent to "User" in the contact OpenAPI.

> Attention:
> - This API supports tenant_access_token and user_access_token.
> - When using tenant_access_token, the data permissions follow the contacts permissions of the application, and the returned field data is the field for which the application has permissions. You can determine the contacts permissions of the application by Getting the contacts permissions configuration.
> - When using the user_access_token, the user is considered an administrator by default, and the administrator's management scope will be verified. When a user has multiple administrator identities and can view employee information, the administrator's management scope will be the maximum set. For information on administrator permissions, see the Help Center article: [Administrator creates administrator role and assigns permissions](https://www.larksuite.com/hc/zh-CN/articles/360043595213)

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/directory/v1/employees/filter |
| HTTP Method | POST |
| Rate Limit | 1000 per minute、50 per second |
| Supported app types | custom,isv |
| Required scopes | `directory:employee:list` |
| Required field scopes | > The response body of the API contains the following sensitive fields, and they will be returned only after corresponding scopes are added. If you do not need the fields, it is not recommended that you request the scopes. `directory:department.base:read` `directory:department.count:read` `directory:department.custom_field:read` `directory:department.data_source:read` `directory:department.department_path:read` `directory:department.external_id:read` `directory:department.has_child:read` `directory:department.leader:read` `directory:department.name:read` `directory:department.order_weight:read` `directory:department.organization:read` `directory:department.parent_id:read` `directory:department.status:read` `directory:employee.base.active_status:read` `directory:employee.base.avatar:read` `directory:employee.base.background_image:read` `directory:employee.base.base:read` `directory:employee.base.custom_field:read` `directory:employee.base.data_source:read` `directory:employee.base.department:read` `directory:employee.base.department_path:read` `directory:employee.base.dept_order:read` `directory:employee.base.description:read` `directory:employee.base.dotted_line_leaders:read` `directory:employee.base.email:read` `directory:employee.base.enterprise_email:read` `directory:employee.base.enterprise_email_alias:read` `directory:employee.base.external_id:read` `directory:employee.base.gender:read` `directory:employee.base.geo:read` `directory:employee.base.is_admin:read` `directory:employee.base.is_primary_admin:read` `directory:employee.base.is_resigned:read` `directory:employee.base.leader:read` `directory:employee.base.leader_id:read` `directory:employee.base.mobile:read` `directory:employee.base.name.another_name:read` `directory:employee.base.name.name:read` `directory:employee.base.resign_time:read` `directory:employee.base.role:read` `directory:employee.base.status:read` `directory:employee.base.subscription_ids:read` `directory:employee.work.base_work:read` `directory:employee.work.employment_type:read` `directory:employee.work.extension_number:read` `directory:employee.work.job_number:read` `directory:employee.work.job_title:read` `directory:employee.work.join_date:read` `directory:employee.work.resign_date:read` `directory:employee.work.resign_reason:read` `directory:employee.work.resign_remark:read` `directory:employee.work.resign_type:read` `directory:employee.work.staff_status:read` `directory:employee.work.work_country_or_region:read` `directory:employee.work.work_place:read` `directory:employee.work.work_station:read` `directory:job_title.base:read` `directory:job_title.status:read` `directory:place.base:read` `directory:place.status:read` `directory:job_family.path:read` `directory:job_family.status:read` `directory:job_level.base:read` `directory:job_level.order:read` `directory:job_level.status:read` `directory:employee.work.job_level:read` `directory:employee.work.employment:read` `directory:employee.work.job_family:read` `directory:job_family.base:read` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `employee_id_type` | `string` | No | User ID categories **Example value**: open_id **Optional values are**:  - `open_id`: Identifies the identity of a user in an application. The same user has different Open IDs in different applications. Learn more: How to get an Open ID - `union_id`: Identifies the identity of a user under an application developer. The Union ID of the same user in an application under the same developer is the same, and the Union ID in an application under different developers is different. With Union ID, application developers can associate the identity of the same user in multiple applications. Learn more: How to get Union ID? - `employee_id`: The unique identifier of an employee within the enterprise. Supports customization, the system automatically generates it if it is not customized. ID supports modification. How to get employee_id: - Enterprise administrators go to the Admin Console > Organizational Structure > Members and Departments page, click Member Details to query employee IDs - Through the interface of Get employee list in batches, query employee ID by mobile phone number or email.  **Default value**: `open_id` **When the value is `employee_id`, the following field scopes are required**: `directory:employee.base.external_id:read` |
| `department_id_type` | `string` | No | The type of department ID used in this call **Example value**: open_department_id **Optional values are**:  - `department_id`: Used to identify a unique department within a tenant - `open_department_id`: It is used to identify a department in a specific application, and the same department has the same open_department_id in different applications.  **Default value**: `open_department_id` | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `filter` | `multi_filter_condition` | Yes | Query conditions Learn more: query condition usage |
|   `conditions` | `filter_condition[]` | Yes | A list of comparison expressions, such as the comparison conditions base_info.mobile eq "+ 8613000000001", and the relationship between multiple expressions is and. **Data validation rules**: - Length range: `0` ～ `10` |
|     `field` | `string` | Yes | The lvalue of the filter condition, the value is the parameter name of the field. Optional filter conditions are: - base_info.mobile - base_info.email - base_info departments department_id - work_info staff_status **Example value**: "base_info.mobile" |
|     `operator` | `string` | Yes | Comparison operator. Optional values are: - eq: equal to, supports any type of lvalue - in: belongs to either **Example value**: "eq" |
|     `value` | `string` | Yes | The rvalue of the filter condition is an escaped JSON string. - For `eq` matching strings, use `"\"str\""` - For `in` matching string lists, use `"[\"str\"]"` - For `eq` matching numbers, use `"123"` - For `in` matching number lists, use `"[123]"` **Example value**: "1" |
| `required_fields` | `string[]` | Yes | List of fields to be queried. Permissioned row and column data will be returned according to the passed field list. No fields will be returned if not passed. Learn more: Field enumeration instructions **Example value**: ["base_info.gender"] **Data validation rules**: - Length range: `0` ～ `100` |
| `page_request` | `page_condition` | Yes | paging parameter |
|   `page_size` | `int` | No | The number of requests this time, the minimum is 0 and the maximum is 100 **Example value**: 10 |
|   `page_token` | `string` | No | Sequential paging query, can not skip page query, support deep paging, in the need to traverse all the data scene can only use this method. The first pass empty string or not, after the return value of the page_token **Example value**: "1" | ### Request body example

{
    "filter": {
        "conditions": [
            {
                "field": "base_info.mobile",
                "operator": "eq",
                "value": "\"8619922333322\""
            }
        ]
    },
    "required_fields": [
        "base_info.gender"
    ],
    "page_request": {
        "page_size": 10,
        "page_token": ""
    }
}

## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `employees` | `employee_entity[]` | employee information |
|     `base_info` | `employee_base_entity` | Staff basic information |
|       `employee_id` | `string` | Unique identification of active employees within the enterprise.  **Field permission requirements**: `directory:employee.base.external_id:read` |
|       `name` | `name` | name |
|         `name` | `i18n_text` | Employee's name **Required field scopes (Satisfy any)**: `directory:employee.base.base:read` `directory:employee.base.name.name:read` |
|           `default_value` | `string` | default value |
|           `i18n_value` | `map<string, string>` | Internationalized value, key is zh_cn, ja_jp, en_us, value is the corresponding value |
|         `another_name` | `string` | alias **Required field scopes (Satisfy any)**: `directory:employee.base.base:read` `directory:employee.base.name.another_name:read` |
|       `mobile` | `string` | Employee's mobile phone number **Required field scopes**: `directory:employee.base.mobile:read` |
|       `email` | `string` | Employee's email address at work **Required field scopes**: `directory:employee.base.email:read` |
|       `enterprise_email` | `string` | Employee's business email address **Required field scopes**: `directory:employee.base.enterprise_email:read` |
|       `gender` | `int` | Gender **Optional values are**:  - `0`: Unknown - `1`: Male - `2`: Female - `3`: Other  **Required field scopes**: `directory:employee.base.gender:read` |
|       `departments` | `department[]` | Information about the department to which the employee belongs. An employee can belong to multiple departments, with the first in line being the main department **Required field scopes**: `directory:employee.base.department:read` |
|         `department_id` | `string` | Department ID **Required field scopes (Satisfy any)**: `directory:department.base:read` `directory:department.external_id:read` |
|         `department_count` | `department_count` | Department members and sub-department counts **Required field scopes (Satisfy any)**: `directory:department.count:read` `directory:department.organization:read` |
|           `recursive_members_count` | `string` | Number of recursive members |
|           `direct_members_count` | `string` | Number of direct members |
|           `recursive_members_count_exclude_leaders` | `string` | Number of recursive members (excluding leaders) |
|           `recursive_departments_count` | `string` | The number of recursive subdepartments |
|           `direct_departments_count` | `string` | Number of direct sub-departments |
|         `has_child` | `boolean` | Whether there are sub department **Required field scopes (Satisfy any)**: `directory:department.has_child:read` `directory:department.organization:read` |
|         `leaders` | `department_leader[]` | Head of Department **Required field scopes**: `directory:department.leader:read` |
|           `leader_type` | `int` | Type of department head **Optional values are**:  - `1`: Main - `2`: Deputy  |
|           `leader_id` | `string` | Leader ID of the Department |
|         `parent_department_id` | `string` | Parent department ID **Required field scopes (Satisfy any)**: `directory:department.organization:read` `directory:department.parent_id:read` |
|         `name` | `i18n_text` | Department name **Required field scopes (Satisfy any)**: `directory:department.base:read` `directory:department.name:read` |
|           `default_value` | `string` | default value |
|           `i18n_value` | `map<string, string>` | Internationalized value, key is zh_cn, ja_jp, en_us, value is the corresponding value |
|         `enabled_status` | `boolean` | Whether to enable **Required field scopes**: `directory:department.status:read` |
|         `order_weight` | `string` | Department sort weights **Required field scopes (Satisfy any)**: `directory:department.order_weight:read` `directory:department.organization:read` |
|         `custom_field_values` | `custom_field_value[]` | Custom fields **Required field scopes**: `directory:department.custom_field:read` |
|           `field_type` | `string` | Custom field types **Optional values are**:  - `1`: Multi-line text - `2`: Web Links - `3`: Enumeration options - `4`: Employee - `9`: Phone - `10`: Multiple-choice enumeration type (currently only supports text type) - `11`: Personnel List  |
|           `text_value` | `i18n_text` | Text field value |
|             `default_value` | `string` | default value |
|             `i18n_value` | `map<string, string>` | Internationalized value, key is zh_cn, ja_jp, en_us, value is the corresponding value |
|           `url_value` | `url_value` | Webpage link field value |
|             `link_text` | `i18n_text` | The title of the page |
|               `default_value` | `string` | default value |
|               `i18n_value` | `map<string, string>` | Internationalized value, key is zh_cn, ja_jp, en_us, value is the corresponding value |
|             `url` | `string` | Links to mobile web pages |
|             `pcurl` | `string` | Desktop web links |
|           `enum_value` | `enum_value` | Enumeration values |
|             `enum_ids` | `string[]` | Option result ID |
|             `enum_type` | `string` | Option type **Optional values are**:  - `1`: Text type - `2`: Image Type  |
|           `user_values` | `user_value[]` | Employee Type |
|             `ids` | `string[]` | Employee ids |
|           `phone_value` | `phone_value` |  |
|             `phone_number` | `string` |  |
|             `extension_number` | `string` |  |
|           `field_key` | `string` | Custom field key |
|         `department_path_infos` | `department_base_info[]` | Department path information. The order is from the root department to the last department **Required field scopes**: `directory:department.department_path:read` |
|           `department_id` | `string` | Department ID |
|           `department_name` | `i18n_text` | Department name |
|             `default_value` | `string` | default value |
|             `i18n_value` | `map<string, string>` | Internationalized value, key is zh_cn, ja_jp, en_us, value is the corresponding value |
|         `data_source` | `int` | data source **Optional values are**:  - `1`: LarkAdmin - `2`: CoreHR - `3`: SCIM  **Required field scopes (Satisfy any)**: `directory:department.base:read` `directory:department.data_source:read` |
|       `employee_order_in_departments` | `user_department_sort_info[]` | The ranking information of users within the department, the first department is the main department **Required field scopes (Satisfy any)**: `directory:employee.base.department:read` `directory:employee.base.dept_order:read` |
|         `department_id` | `string` | Department ID |
|         `order_weight_in_deparment` | `string` | The ranking weight of users within the department |
|         `order_weight_among_deparments` | `string` | User ranking weights across multiple departments |
|       `description` | `string` | personal signature **Required field scopes (Satisfy any)**: `directory:employee.base.base:read` `directory:employee.base.description:read` |
|       `active_status` | `int` | user active status **Optional values are**:  - `1`: not activated - `2`: activate - `3`: freeze - `4`: active exit - `5`: not joined  **Required field scopes (Satisfy any)**: `directory:employee.base.active_status:read` `directory:employee.base.status:read` |
|       `is_resigned` | `boolean` | whether to quit **Required field scopes (Satisfy any)**: `directory:employee.base.is_resigned:read` `directory:employee.base.status:read` |
|       `leader_id` | `string` | Direct manager ID **Required field scopes (Satisfy any)**: `directory:employee.base.leader:read` `directory:employee.base.leader_id:read` |
|       `dotted_line_leader_ids` | `string[]` | The ID of the dotted line leader **Required field scopes (Satisfy any)**: `directory:employee.base.dotted_line_leaders:read` `directory:employee.base.leader:read` |
|       `is_primary_admin` | `boolean` | Whether is a super administrator **Required field scopes (Satisfy any)**: `directory:employee.base.is_primary_admin:read` `directory:employee.base.role:read` |
|       `enterprise_email_aliases` | `string[]` | The employee's enterprise email alias. When members want to use different email addresses for different people to send emails, they can add aliases to their enterprise email addresses. Once added, members can send messages using aliases and accept messages sent to aliases **Required field scopes (Satisfy any)**: `directory:employee.base.enterprise_email:read` `directory:employee.base.enterprise_email_alias:read` |
|       `custom_field_values` | `custom_field_value[]` | Custom field values **Required field scopes**: `directory:employee.base.custom_field:read` |
|         `field_type` | `string` | Custom field types **Optional values are**:  - `1`: Multi-line text - `2`: Web Links - `3`: Enumeration options - `4`: Employee - `9`: phone - `10`: Multiple-choice enumeration type (currently only supports text type) - `11`: Personnel List  |
|         `text_value` | `i18n_text` | Text field value |
|           `default_value` | `string` | default value |
|           `i18n_value` | `map<string, string>` | Internationalized value, key is zh_cn, ja_jp, en_us, value is the corresponding value |
|         `url_value` | `url_value` | Webpage link field value |
|           `link_text` | `i18n_text` | The title of the page |
|             `default_value` | `string` | default value |
|             `i18n_value` | `map<string, string>` | Internationalized value, key is zh_cn, ja_jp, en_us, value is the corresponding value |
|           `url` | `string` | Links to mobile web pages |
|           `pcurl` | `string` | Desktop web links |
|         `enum_value` | `enum_value` | Enumeration values |
|           `enum_ids` | `string[]` | Option result ID |
|           `enum_type` | `string` | Option type **Optional values are**:  - `1`: text - `2`: Image  |
|         `user_values` | `user_value[]` | Employee field value |
|           `ids` | `string[]` | Employee ID |
|         `phone_value` | `phone_value` |  |
|           `phone_number` | `string` |  |
|           `extension_number` | `string` |  |
|         `field_key` | `string` | Custom field key |
|       `department_path_infos` | `department_base_info[][]` | Department path information. The order is from the root department to the last department **Required field scopes**: `directory:employee.base.department_path:read` |
|       `resign_time` | `string` | Departure Time. The time of the departure operation, automatically generated by the system and cannot be written. Date format: YYYY-MM-DD **Required field scopes**: `directory:employee.base.resign_time:read` |
|       `avatar` | `image_link` | employee avatar **Required field scopes (Satisfy any)**: `directory:employee.base.avatar:read` `directory:employee.base.base:read` |
|         `avatar_72` | `string` | 72*72 pixel avatar link |
|         `avatar_240` | `string` | 240*240 pixel avatar link |
|         `avatar_640` | `string` | 640*640 pixel avatar link |
|         `avatar_origin` | `string` | Original avatar link |
|       `background_image` | `string` | Custom background image URL **Required field scopes (Satisfy any)**: `directory:employee.base.background_image:read` `directory:employee.base.base:read` |
|       `is_admin` | `boolean` | Whether is a normal administrator **Required field scopes (Satisfy any)**: `directory:employee.base.is_admin:read` `directory:employee.base.role:read` |
|       `data_source` | `int` | data source **Optional values are**:  - `1`: admin background - `2`: Personnel Enterprise Edition - `3`: SCIM  **Required field scopes (Satisfy any)**: `directory:employee.base.base:read` `directory:employee.base.data_source:read` |
|       `geo_name` | `string` | Where employee data resides **Required field scopes (Satisfy any)**: `directory:employee.base.base:read` `directory:employee.base.geo:read` |
|       `subscription_ids` | `string[]` | A list of seat IDs assigned to the employee. **Required field scopes (Satisfy any)**: `directory:employee.base.base:read` `directory:employee.base.subscription_ids:read` |
|     `work_info` | `employee_work_entity` | Employee work information |
|       `work_country_or_region` | `string` | Country/region of work How to query the meaning of the code of the country **Required field scopes (Satisfy any)**: `directory:employee.work.base_work:read` `directory:employee.work.work_country_or_region:read` |
|       `work_place` | `place` | Work place **Required field scopes (Satisfy any)**: `directory:employee.work.base_work:read` `directory:employee.work.work_place:read` |
|         `place_id` | `string` | ID, which returns "0" by default if it does not exist. |
|         `place_name` | `i18n_text` | Name of the work location **Required field scopes**: `directory:place.base:read` |
|           `default_value` | `string` | default value |
|           `i18n_value` | `map<string, string>` | Internationalized value, key is zh_cn, ja_jp, en_us, value is the corresponding value |
|         `is_enabled` | `boolean` | Whether it is enabled **Required field scopes**: `directory:place.status:read` |
|         `description` | `i18n_text` | description **Required field scopes**: `directory:place.base:read` |
|           `default_value` | `string` | default value |
|           `i18n_value` | `map<string, string>` | Internationalized value, key is zh_cn, ja_jp, en_us, value is the corresponding value |
|       `work_station` | `i18n_text` | Work station **Required field scopes (Satisfy any)**: `directory:employee.work.base_work:read` `directory:employee.work.work_station:read` |
|         `default_value` | `string` | default value |
|         `i18n_value` | `map<string, string>` | Internationalized value, key is zh_cn, ja_jp, en_us, value is the corresponding value |
|       `job_number` | `string` | Job number **Required field scopes (Satisfy any)**: `directory:employee.work.base_work:read` `directory:employee.work.job_number:read` |
|       `extension_number` | `string` | Extension number **Required field scopes (Satisfy any)**: `directory:employee.work.base_work:read` `directory:employee.work.extension_number:read` |
|       `join_date` | `string` | Employee joining date **Required field scopes (Satisfy any)**: `directory:employee.work.join_date:read` `directory:employee.work.employment:read` |
|       `employment_type` | `int` | Employment type **Required field scopes (Satisfy any)**: `directory:employee.work.employment_type:read` `directory:employee.work.employment:read` |
|       `staff_status` | `int` | Employee staff status **Optional values are**:  - `1`: Unresigned - `2`: Resigned - `3`: Pre onboarding - `4`: Cancel onboarding - `5`: Prepared to quit  **Required field scopes (Satisfy any)**: `directory:employee.work.staff_status:read` `directory:employee.work.employment:read` |
|       `job_title` | `job_title` | Job title **Required field scopes**: `directory:employee.work.job_title:read` |
|         `job_title_id` | `string` | ID, which returns "0" by default if it does not exist. |
|         `job_title_name` | `i18n_text` | Job title name **Required field scopes**: `directory:job_title.base:read` |
|           `default_value` | `string` | default value |
|           `i18n_value` | `map<string, string>` | Internationalized value, key is zh_cn, ja_jp, en_us, value is the corresponding value |
|         `is_enabled` | `boolean` | Whether it is enabled **Required field scopes**: `directory:job_title.status:read` |
|         `description` | `i18n_text` | description **Required field scopes**: `directory:job_title.base:read` |
|           `default_value` | `string` | default value |
|           `i18n_value` | `map<string, string>` | Internationalized value, key is zh_cn, ja_jp, en_us, value is the corresponding value |
|       `job_level` | `job_level` | **Required field scopes**: `directory:employee.work.job_level:read` |
|         `job_level_id` | `string` |  |
|         `job_level_name` | `i18n_text` | **Required field scopes**: `directory:job_level.base:read` |
|           `default_value` | `string` |  |
|           `i18n_value` | `map<string, string>` |  |
|         `is_enabled` | `boolean` | **Required field scopes**: `directory:job_level.status:read` |
|         `is_deleted` | `boolean` | **Required field scopes**: `directory:job_level.status:read` |
|         `order` | `string` | **Required field scopes**: `directory:job_level.order:read` |
|         `description` | `i18n_text` | **Required field scopes**: `directory:job_level.base:read` |
|           `default_value` | `string` |  |
|           `i18n_value` | `map<string, string>` |  |
|       `job_family` | `job_family` | **Required field scopes**: `directory:employee.work.job_family:read` |
|         `job_family_id` | `string` |  |
|         `job_family_name` | `i18n_text` | **Required field scopes**: `directory:job_family.base:read` |
|           `default_value` | `string` |  |
|           `i18n_value` | `map<string, string>` |  |
|         `is_enabled` | `boolean` | **Required field scopes**: `directory:job_family.status:read` |
|         `parent_job_family_id` | `string` | **Required field scopes**: `directory:job_family.path:read` |
|         `description` | `i18n_text` | **Required field scopes**: `directory:job_family.base:read` |
|           `default_value` | `string` |  |
|           `i18n_value` | `map<string, string>` |  |
|       `resign_date` | `string` | Date of separation 2007-03-20 **Required field scopes (Satisfy any)**: `directory:employee.work.resign_date:read` `directory:employee.work.employment:read` |
|       `resign_reason` | `string` | Reason for leaving  **Example value**: "1" **Optional values are**:  - `1`: Salary does not meet expectations - `2`: Working hours are too long - `3`: Dissatisfied with the job content - `4`: Does not approve of superiors or management - `5`: Limited career opportunities - `6`: Lack of identification with company culture - `7`: Organizational restructuring (voluntary departure) - `8`: Expiration of contract - `9`: Job hopping - `10`: Transfer line - `11`: Family reasons - `12`: Poor health - `13`: Reason for duty station - `14`: Other (voluntary departure) - `15`: Unexpected - `16`: Death - `17`: Dismissal - `18`: probation period not passed - `19`: Poor job performance - `20`: Work output low - `21`: Organizational restructuring (passive departure) - `22`: Discipline - `23`: Illegal - `24`: Other (passive departure) - `25`: Other (other)  **Required field scopes (Satisfy any)**: `directory:employee.work.resign_reason:read` `directory:employee.work.employment:read` |
|       `resign_remark` | `string` | Resignation remarks **Required field scopes (Satisfy any)**: `directory:employee.work.resign_remark:read` `directory:employee.work.employment:read` |
|       `resign_type` | `string` | Resign Type  **Example Value**: "1" **Optional values are**:  - `1`: Active - `2`: Passive - `3`: Other  **Required field scopes (Satisfy any)**: `directory:employee.work.resign_type:read` `directory:employee.work.employment:read` |
|   `page_response` | `page_response` | paging results |
|     `has_more` | `boolean` | Whether there are follow-up results |
|     `page_token` | `string` | Paging token, when has_more true, will return the new page_token at the same time, otherwise not return page_token |
|   `abnormals` | `abnormal_record[]` | exception information |
|     `id` | `string` | exception ID |
|     `row_error` | `int` | row-level exception **Optional values are**:  - `0`: success - `1000`: No permission  |
|     `field_errors` | `map<string, int>` | Column-level exceptions, key is the field name, and value is the following enumeration **Optional values are**:  - `1000`: no permissions - `2000`: service exception - `2002`: user does not exist - `2003`: field does not exist | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data":
    {
        "employees":
        [
            {
                "base_info":
                {
                    "employee_id": "sddasdeqwe",
                    "name":
                    {
                        "name":
                        {
                            "default_value": "张三",
                            "i18n_value":
                            {
                                "zh_cn": "张三"
                            }
                        },
                        "another_name": "张小明"
                    },
                    "mobile": "+8613011111111",
                    "email": "zhangsan@company.com",
                    "enterprise_email": "zhangsan@company.com",
                    "gender": 1,
                    "departments":
                    [
                        {
                            "department_id": "h12921",
                            "department_count":
                            {
                                "recursive_members_count": "100",
                                "direct_members_count": "100",
                                "recursive_members_count_exclude_leaders": "100",
                                "recursive_departments_count": "100",
                                "direct_departments_count": "100"
                            },
                            "has_child": true,
                            "leaders":
                            [
                                {
                                    "leader_type": "1",
                                    "leader_id": "u273y71"
                                }
                            ],
                            "hrbps":
                            [
                                "eefasdqw"
                            ],
                            "parent_department_id": "h12921",
                            "name":
                            {
                                "default_value": "张三",
                                "i18n_value":
                                {
                                    "zh_cn": "张三"
                                }
                            },
                            "order_weight": "100",
                            "custom_field_values":
                            [
                                {
                                    "field_key": "C-1000001",
                                    "field_type": "1",
                                    "text_value":
                                    {
                                        "default_value": "张三",
                                        "i18n_value":
                                        {
                                            "zh_cn": "张三"
                                        }
                                    },
                                    "url_value":
                                    {
                                        "link_text":
                                        {
                                            "default_value": "张三",
                                            "i18n_value":
                                            {
                                                "zh_cn": "张三"
                                            }
                                        },
                                        "url": "https://m.bytedance.com/afnasjfna",
                                        "pcurl": "http://www.fs.cn"
                                    },
                                    "enum_value":
                                    {
                                        "enum_ids":
                                        [
                                            "1"
                                        ],
                                        "enum_name": 1,
                                        "enum_type": "1"
                                    },
                                    "user_values":
                                    [
                                        {
                                            "ids":
                                            [
                                                "1"
                                            ],
                                            "user_type": "1"
                                        }
                                    ]
                                }
                            ],
                            "department_path_infos":
                            [
                                {
                                    "department_id": "1",
                                    "department_name":
                                    {
                                        "default_value": "张三",
                                        "i18n_value":
                                        {
                                            "zh_cn": "张三"
                                        }
                                    }
                                }
                            ],
                            "data_source": 1
                        }
                    ],
                    "employee_order_in_departments":
                    [
                        {
                            "department_id": "h12921",
                            "order_weight_in_deparment": "100",
                            "order_weight_among_deparments": "100"
                        }
                    ],
                    "description": "新成员请多关照",
                    "active_status": 1,
                    "is_resigned": true,
                    "leader_id": "uyg77nx",
                    "dotted_line_leader_ids":
                    [
                        "asdasdqwe"
                    ],
                    "is_primary_admin": true,
                    "enterprise_email_aliases":
                    [
                        "saqwe@163.com"
                    ],
                    "custom_field_values":
                    [
                        {
                            "field_key": "C-1000001",
                            "field_type": "1",
                            "text_value":
                            {
                                "default_value": "张三",
                                "i18n_value":
                                {
                                    "zh_cn": "张三"
                                }
                            },
                            "url_value":
                            {
                                "link_text":
                                {
                                    "default_value": "张三",
                                    "i18n_value":
                                    {
                                        "zh_cn": "张三"
                                    }
                                },
                                "url": "https://m.bytedance.com/afnasjfna",
                                "pcurl": "http://www.fs.cn"
                            },
                            "enum_value":
                            {
                                "enum_ids":
                                [
                                    "1"
                                ],
                                "enum_name": "选项结果名称",
                                "enum_type": "1"
                            },
                            "user_values":
                            [
                                {
                                    "ids":
                                    [
                                        "1"
                                    ],
                                    "user_type": "1"
                                }
                            ]
                        }
                    ],
                    "department_path_infos":
                    [
                        [
                            {
                                "department_id": "1",
                                "department_name":
                                {
                                    "default_value": "张三",
                                    "i18n_value":
                                    {
                                        "zh_cn": "张三"
                                    }
                                }
                            }
                        ]
                    ],
                    "resign_time": "2023-10-01",
                    "avatar":
                    {
                        "avatar_72": "http://qwed.com",
                        "avatar_240": "http://wssd.com",
                        "avatar_640": "http://wssd.com",
                        "avatar_origin": "https:inernal-api/image"
                    },
                    "background_image": "http://sadui.com",
                    "is_admin": true,
                    "data_source": 1,
                    "geo_name": "china",
                    "subscription_ids":
                    [
                        "458694723562"
                    ]
                },
                "work_info":
                {
                    "work_country_or_region": "MDCT00000012",
                    "work_place":
                    {
                        "place_id": "place100",
                        "place_name":
                        {
                            "default_value": "张三",
                            "i18n_value":
                            {
                                "zh_cn": "张三"
                            }
                        },
                        "is_enabled": true,
                        "description":
                        {
                            "default_value": "张三",
                            "i18n_value":
                            {
                                "zh_cn": "张三"
                            }
                        }
                    },
                    "work_station":
                    {
                        "default_value": "张三",
                        "i18n_value":
                        {
                            "zh_cn": "张三"
                        }
                    },
                    "job_number": "2845435",
                    "extension_number": "2845435",
                    "join_date": "2007-03-20",
                    "employment_type": 1,
                    "staff_status": 1,
                    "positions":
                    [
                        {
                            "position_code": "PO18890",
                            "position_name": "总经理",
                            "leader_id": "2e1cf73b",
                            "leader_position_code": "e71b94gb",
                            "is_main_position": true,
                            "department_id": "D100"
                        }
                    ],
                    "job_title":
                    {
                        "job_title_id": "od-4e6ac4d14bcd5071a37a39de902c7141",
                        "job_title_name":
                        {
                            "default_value": "张三",
                            "i18n_value":
                            {
                                "zh_cn": "张三"
                            }
                        },
                        "is_enabled": true,
                        "description":
                        {
                            "default_value": "张三",
                            "i18n_value":
                            {
                                "zh_cn": "张三"
                            }
                        }
                    }
                }
            }
        ],
        "page_response":
        {
            "has_more": true,
            "page_token": "sdefsd"
        },
        "abnormals":
        [
            {
                "row_error": 0,
                "field_errors":
                {
                    "base_info.mobile": "1000"
                },
                "id": "eedasfwe"
            }
        ]
    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 2220001 | param is invalid | If the request parameters are invalid, modify the parameters |
| 400 | 2221004 | Invalid page token | Invalid token. Please check whether the page_token is correct or obtain a valid token again. |
| 400 | 2221005 | No page requests | No page request parameters. Please pass in a valid page_request parameter (such as page_token or page_size). |
| 400 | 2220009 | Filter field is invalid | Invalid field in filter, please modify filter |
| 400 | 2220010 | Exceeded the limit size | The page size exceeds the limit, please modify the page size |
| 400 | 2220012 | The field is not supported filter | The field in filter does not support filtering, please modify the filter |
| 400 | 2220013 | The field does not support the operator | The field in the filter does not support this operator, please modify the filter |
| 400 | 2220014 | Invalid field value | The value in the filter is invalid, please modify the filter |
