---
title: "Employee modified"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/directory-v1/employee/events/updated"
service: "directory-v1"
resource: "employee"
section: "Organization"
scopes:
  - "directory:employee:read"
field_scopes:
  - "directory:department.base:read"
  - "directory:department.external_id:read"
  - "directory:employee.base.active_status:read"
  - "directory:employee.base.avatar:read"
  - "directory:employee.base.background_image:read"
  - "directory:employee.base.base:read"
  - "directory:employee.base.custom_field:read"
  - "directory:employee.base.department:read"
  - "directory:employee.base.dept_order:read"
  - "directory:employee.base.description:read"
  - "directory:employee.base.dotted_line_leaders:read"
  - "directory:employee.base.email:read"
  - "directory:employee.base.gender:read"
  - "directory:employee.base.is_resigned:read"
  - "directory:employee.base.leader:read"
  - "directory:employee.base.leader_id:read"
  - "directory:employee.base.mobile:read"
  - "directory:employee.base.name.another_name:read"
  - "directory:employee.base.name.name:read"
  - "directory:employee.base.resign_time:read"
  - "directory:employee.base.status:read"
  - "directory:employee.work.base_work:read"
  - "directory:employee.work.work_station:read"
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
  - "directory:employee.work.employment:read"
updated: "1756439278000"
---

# Employee modified

To receive the event notification when the employee's profile is updated. 

## Prerequisites
You need to configure event subscriptions in the application so that you can receive event data when the event is triggered. For more information about event subscriptions, see Event Overview. 

## Precautions
Some fields of this event require permissions. You can refer to the corresponding parameter descriptions to obtain the permissions required for the parameters. Only when the application has the corresponding field permissions enabled can it successfully receive the complete event body data. For the specific operations to apply for permission, see Apply for API Permissions. 

{Usage Examples}(url=/api/tools/api_explore/api_explore_config?project=directory&version=v1&resource=employee&event=updated)

## Event
| Facts |  |
| --- | --- |
| Event type | directory.employee.updated_v1 |
| Supported app types | custom,isv |
| Required scopes | `directory:employee:read` |
| Required field scopes | > The response body of the API contains the following sensitive fields, and they will be returned only after corresponding scopes are added. If you do not need the fields, it is not recommended that you request the scopes. `directory:department.base:read` `directory:department.external_id:read` `directory:employee.base.active_status:read` `directory:employee.base.avatar:read` `directory:employee.base.background_image:read` `directory:employee.base.base:read` `directory:employee.base.custom_field:read` `directory:employee.base.department:read` `directory:employee.base.dept_order:read` `directory:employee.base.description:read` `directory:employee.base.dotted_line_leaders:read` `directory:employee.base.email:read` `directory:employee.base.gender:read` `directory:employee.base.is_resigned:read` `directory:employee.base.leader:read` `directory:employee.base.leader_id:read` `directory:employee.base.mobile:read` `directory:employee.base.name.another_name:read` `directory:employee.base.name.name:read` `directory:employee.base.resign_time:read` `directory:employee.base.status:read` `directory:employee.work.base_work:read` `directory:employee.work.work_station:read` `directory:employee.work.employment_type:read` `directory:employee.work.extension_number:read` `directory:employee.work.job_number:read` `directory:employee.work.job_title:read` `directory:employee.work.join_date:read` `directory:employee.work.resign_date:read` `directory:employee.work.resign_reason:read` `directory:employee.work.resign_remark:read` `directory:employee.work.resign_type:read` `directory:employee.work.staff_status:read` `directory:employee.work.work_country_or_region:read` `directory:employee.work.work_place:read` `directory:employee.work.employment:read` |
| Push Mode | `Webhook` | ### Event body
| Parameter | Type | Description |
| --- | --- | --- |
| `schema` | `string` | Event schema |
| `header` | `event_header` | Event header |
|   `event_id` | `string` | Event ID |
|   `event_type` | `string` | Event type |
|   `create_time` | `string` | Event creation timestamp(in ms) |
|   `token` | `string` | Event token |
|   `app_id` | `string` | App ID |
|   `tenant_key` | `string` | Tenant key |
| `event` | `\-` | \- |
|   `changed_properties` | `string[]` | List of changed attributes. It enumerates the modified field names. If a field name appears but is missing from the modified entity information, it may be due to a lack of read permissions and requires further verification. **Data validation rules**: - Length range: `0` ～ `100` |
|   `employee_prev` | `employee_entity` | Employee entity |
|     `base_info` | `employee_base_entity` | Basic employee information |
|       `employee_id` | `string` | The user's open_id. For the ID type, see User Identity Overview |
|       `name` | `name` | Name |
|         `name` | `i18n_text` | I18n text **Required field scopes (Satisfy any)**: `directory:employee.base.base:read` `directory:employee.base.name.name:read` |
|           `default_value` | `string` | Default value |
|           `i18n_value` | `map<string, string>` | Internationalized value, key is zh_cn, ja_jp, en_us, value is the corresponding value |
|             `key` | `string` | International locale |
|             `value` | `string` | value |
|         `another_name` | `string` | Alias **Required field scopes (Satisfy any)**: `directory:employee.base.base:read` `directory:employee.base.name.another_name:read` |
|       `mobile` | `string` | Mobile phone number **Data validation rules**: - Maximum length: `255` characters **Required field scopes**: `directory:employee.base.mobile:read` |
|       `email` | `string` | Contact email **Data validation rules**: - Maximum length: `255` characters **Required field scopes**: `directory:employee.base.email:read` |
|       `gender` | `int` | Gender **Optional values are**:  - `0`: Unknown - `1`: male - `2`: female - `3`: other  **Required field scopes**: `directory:employee.base.gender:read` |
|       `departments` | `department[]` | Department information **Data validation rules**: - Length range: `0` ～ `20` **Required field scopes**: `directory:employee.base.department:read` |
|         `department_id` | `string` | Department ID. The department ID type is open_department_id. For details about the department ID, see Department Resource Introduction. **Data validation rules**: - Maximum length: `64` characters **Required field scopes (Satisfy any)**: `directory:department.base:read` `directory:department.external_id:read` |
|       `employee_order_in_departments` | `user_department_sort_info[]` | The ranking information of users within the department, the first department is the main department **Data validation rules**: - Length range: `0` ～ `20` **Required field scopes (Satisfy any)**: `directory:employee.base.department:read` `directory:employee.base.dept_order:read` |
|         `department_id` | `string` | Department ID. The department ID type is open_department_id. For details about the department ID, see Department Resource Introduction. |
|         `order_weight_in_deparment` | `string` | The ranking weight of users within the department |
|         `order_weight_among_deparments` | `string` | User ranking weights across multiple departments |
|       `description` | `string` | Personal signature **Required field scopes (Satisfy any)**: `directory:employee.base.base:read` `directory:employee.base.description:read` |
|       `active_status` | `int` | User active status **Optional values are**:  - `1`: not activated - `2`: activate - `3`: freeze - `4`: active exit - `5`: not joined  **Required field scopes (Satisfy any)**: `directory:employee.base.active_status:read` `directory:employee.base.status:read` |
|       `is_resigned` | `boolean` | Whether to quit **Required field scopes (Satisfy any)**: `directory:employee.base.is_resigned:read` `directory:employee.base.status:read` |
|       `leader_id` | `string` | The open_id of the user's line manager. For more information on user IDs, see Overview of User Identity. **Required field scopes (Satisfy any)**: `directory:employee.base.leader:read` `directory:employee.base.leader_id:read` |
|       `dotted_line_leader_ids` | `string[]` | The open_id of the user's dotted-line manager. For more information about user IDs, see User Identity Overview. **Data validation rules**: - Length range: `0` ～ `10` **Required field scopes (Satisfy any)**: `directory:employee.base.dotted_line_leaders:read` `directory:employee.base.leader:read` |
|       `custom_field_values` | `custom_field_value[]` | Custom field value **Data validation rules**: - Length range: `0` ～ `100` **Required field scopes**: `directory:employee.base.custom_field:read` |
|         `field_key` | `string` | Custom field key |
|         `field_type` | `string` | Custom field type **Optional values are**:  - `1`: multiline text - `2`: web link - `3`: Enumeration options - `4`: personnel - `10`: Multiple-choice enumeration types (currently only text types are supported) - `11`: personnel list  |
|         `text_value` | `i18n_text` | I18n text |
|           `default_value` | `string` | Default value |
|           `i18n_value` | `map<string, string>` | Internationalized value, key is zh_cn, ja_jp, en_us, value is the corresponding value |
|             `key` | `string` | International locale |
|             `value` | `string` | value |
|         `url_value` | `url_value` | Web link field value |
|           `link_text` | `i18n_text` | I18n text |
|             `default_value` | `string` | Default value |
|             `i18n_value` | `map<string, string>` | Internationalized value, key is zh_cn, ja_jp, en_us, value is the corresponding value |
|               `key` | `string` | International locale |
|               `value` | `string` | value |
|           `url` | `string` | Mobile end web link |
|           `pcurl` | `string` | Desktop web link |
|         `enum_value` | `enum_value` | enumeration |
|           `enum_ids` | `string[]` | Option result ID **Data validation rules**: - Length range: `0` ～ `100` |
|           `enum_type` | `string` | option type **Optional values are**:  - `1`: Text - `2`: picture  |
|         `user_values` | `user_value[]` | Person field value **Data validation rules**: - Length range: `0` ～ `100` |
|           `ids` | `string[]` | List of user IDs. For more information about user IDs, see User Identity Overview. **Data validation rules**: - Length range: `0` ～ `100` |
|           `user_type` | `string` | Personnel type **Optional values are**:  - `1`: Employee  |
|       `resign_time` | `string` | Turnover time **Required field scopes**: `directory:employee.base.resign_time:read` |
|       `avatar` | `image_link` | Avatar url **Required field scopes (Satisfy any)**: `directory:employee.base.avatar:read` `directory:employee.base.base:read` |
|         `avatar_72` | `string` | 72 * 72 pixel avatar link |
|         `avatar_240` | `string` | 240 * 240 pixel avatar link |
|         `avatar_640` | `string` | 640 * 640 pixel avatar link |
|         `avatar_origin` | `string` | Original avatar link |
|       `background_image` | `string` | Custom background cover url **Required field scopes (Satisfy any)**: `directory:employee.base.background_image:read` `directory:employee.base.base:read` |
|     `work_info` | `employee_work_entity` | Employee job information |
|       `work_country_or_region` | `string` | Work country/region **Required field scopes (Satisfy any)**: `directory:employee.work.base_work:read` `directory:employee.work.work_country_or_region:read` |
|       `work_place` | `place` | Work location, which is the City property in user's Job details **Required field scopes (Satisfy any)**: `directory:employee.work.base_work:read` `directory:employee.work.work_place:read` |
|         `place_id` | `string` | ID |
|       `work_station` | `i18n_text` | Physical seat ID **Required field scopes (Satisfy any)**: `directory:employee.work.base_work:read` `directory:employee.work.work_station:read` |
|         `default_value` | `string` | Default value |
|         `i18n_value` | `map<string, string>` | Internationalized value, key is zh_cn, ja_jp, en_us, value is the corresponding value |
|           `key` | `string` | International locale |
|           `value` | `string` | value |
|       `job_number` | `string` | job number **Data validation rules**: - Maximum length: `255` characters **Required field scopes (Satisfy any)**: `directory:employee.work.base_work:read` `directory:employee.work.job_number:read` |
|       `extension_number` | `string` | Extension number **Data validation rules**: - Maximum length: `99` characters **Required field scopes (Satisfy any)**: `directory:employee.work.base_work:read` `directory:employee.work.extension_number:read` |
|       `join_date` | `string` | Join date, e.g., 2007-03-20 **Required field scopes (Satisfy any)**: `directory:employee.work.join_date:read` `directory:employee.work.employment:read` |
|       `employment_type` | `int` | Employee type **Required field scopes (Satisfy any)**: `directory:employee.work.employment_type:read` `directory:employee.work.employment:read` |
|       `staff_status` | `int` | Employee personnel status **Optional values are**:  - `1`: on the job - `2`: leave - `3`: to be hired - `4`: Cancel onboarding - `5`: to leave  **Required field scopes (Satisfy any)**: `directory:employee.work.staff_status:read` `directory:employee.work.employment:read` |
|       `job_title` | `job_title` | Position **Required field scopes**: `directory:employee.work.job_title:read` |
|         `job_title_id` | `string` | ID **Data validation rules**: - Maximum length: `64` characters |
|       `resign_date` | `string` | Date of resignation, e.g., 2007-03-20 **Required field scopes (Satisfy any)**: `directory:employee.work.resign_date:read` `directory:employee.work.employment:read` |
|       `resign_reason` | `string` | Reason for leaving **Required field scopes (Satisfy any)**: `directory:employee.work.resign_reason:read` `directory:employee.work.employment:read` |
|       `resign_remark` | `string` | Resignation remarks **Data validation rules**: - Maximum length: `5000` characters **Required field scopes (Satisfy any)**: `directory:employee.work.resign_remark:read` `directory:employee.work.employment:read` |
|       `resign_type` | `string` | Type of turnover **Required field scopes (Satisfy any)**: `directory:employee.work.resign_type:read` `directory:employee.work.employment:read` |
|   `employee_curr` | `employee_entity` | Employee entity |
|     `base_info` | `employee_base_entity` | Basic employee information |
|       `employee_id` | `string` | The user's open_id. For the ID type, see User Identity Overview |
|       `name` | `name` | Name |
|         `name` | `i18n_text` | I18n text **Required field scopes (Satisfy any)**: `directory:employee.base.base:read` `directory:employee.base.name.name:read` |
|           `default_value` | `string` | Default value |
|           `i18n_value` | `map<string, string>` | Internationalized value, key is zh_cn, ja_jp, en_us, value is the corresponding value |
|             `key` | `string` | International locale |
|             `value` | `string` | value |
|         `another_name` | `string` | Alias **Required field scopes (Satisfy any)**: `directory:employee.base.base:read` `directory:employee.base.name.another_name:read` |
|       `mobile` | `string` | Mobile phone number **Data validation rules**: - Maximum length: `255` characters **Required field scopes**: `directory:employee.base.mobile:read` |
|       `email` | `string` | Contact email **Data validation rules**: - Maximum length: `255` characters **Required field scopes**: `directory:employee.base.email:read` |
|       `gender` | `int` | Gender **Optional values are**:  - `0`: Unknown - `1`: male - `2`: female - `3`: other  **Required field scopes**: `directory:employee.base.gender:read` |
|       `departments` | `department[]` | Department information **Data validation rules**: - Length range: `0` ～ `20` **Required field scopes**: `directory:employee.base.department:read` |
|         `department_id` | `string` | Department ID. The department ID type is open_department_id. For details about the department ID, see Department Resource Introduction. **Data validation rules**: - Maximum length: `64` characters **Required field scopes (Satisfy any)**: `directory:department.base:read` `directory:department.external_id:read` |
|       `employee_order_in_departments` | `user_department_sort_info[]` | The ranking information of users within the department, the first department is the main department **Data validation rules**: - Length range: `0` ～ `20` **Required field scopes (Satisfy any)**: `directory:employee.base.department:read` `directory:employee.base.dept_order:read` |
|         `department_id` | `string` | Department ID. The department ID type is open_department_id. For details about the department ID, see Department Resource Introduction. |
|         `order_weight_in_deparment` | `string` | The ranking weight of users within the department |
|         `order_weight_among_deparments` | `string` | User ranking weights across multiple departments |
|       `description` | `string` | Personal signature **Required field scopes (Satisfy any)**: `directory:employee.base.base:read` `directory:employee.base.description:read` |
|       `active_status` | `int` | User active status **Optional values are**:  - `1`: not activated - `2`: activate - `3`: freeze - `4`: active exit - `5`: not joined  **Required field scopes (Satisfy any)**: `directory:employee.base.active_status:read` `directory:employee.base.status:read` |
|       `is_resigned` | `boolean` | Whether to quit **Required field scopes (Satisfy any)**: `directory:employee.base.is_resigned:read` `directory:employee.base.status:read` |
|       `leader_id` | `string` | The open_id of the user's line manager. For more information on user IDs, see Overview of User Identity. **Required field scopes (Satisfy any)**: `directory:employee.base.leader:read` `directory:employee.base.leader_id:read` |
|       `dotted_line_leader_ids` | `string[]` | The open_id of the user's dotted-line manager. For more information about user IDs, see User Identity Overview. **Data validation rules**: - Length range: `0` ～ `10` **Required field scopes (Satisfy any)**: `directory:employee.base.dotted_line_leaders:read` `directory:employee.base.leader:read` |
|       `custom_field_values` | `custom_field_value[]` | Custom field value **Data validation rules**: - Length range: `0` ～ `100` **Required field scopes**: `directory:employee.base.custom_field:read` |
|         `field_key` | `string` | Custom field key |
|         `field_type` | `string` | Custom field type **Optional values are**:  - `1`: multiline text - `2`: web link - `3`: Enumeration options - `4`: personnel - `10`: Multiple-choice enumeration types (currently only text types are supported) - `11`: personnel list  |
|         `text_value` | `i18n_text` | I18n text |
|           `default_value` | `string` | Default value |
|           `i18n_value` | `map<string, string>` | Internationalized value, key is zh_cn, ja_jp, en_us, value is the corresponding value |
|             `key` | `string` | International locale |
|             `value` | `string` | value |
|         `url_value` | `url_value` | Web link field value |
|           `link_text` | `i18n_text` | I18n text |
|             `default_value` | `string` | Default value |
|             `i18n_value` | `map<string, string>` | Internationalized value, key is zh_cn, ja_jp, en_us, value is the corresponding value |
|               `key` | `string` | International locale |
|               `value` | `string` | value |
|           `url` | `string` | Mobile end web link |
|           `pcurl` | `string` | Desktop web link |
|         `enum_value` | `enum_value` | enumeration |
|           `enum_ids` | `string[]` | Option result ID **Data validation rules**: - Length range: `0` ～ `100` |
|           `enum_type` | `string` | option type **Optional values are**:  - `1`: Text - `2`: picture  |
|         `user_values` | `user_value[]` | Person field value **Data validation rules**: - Length range: `0` ～ `100` |
|           `ids` | `string[]` | List of user IDs. For more information about user IDs, see User Identity Overview. **Data validation rules**: - Length range: `0` ～ `100` |
|           `user_type` | `string` | Personnel type **Optional values are**:  - `1`: Employee  |
|       `resign_time` | `string` | Turnover time **Required field scopes**: `directory:employee.base.resign_time:read` |
|       `avatar` | `image_link` | Avatar url **Required field scopes (Satisfy any)**: `directory:employee.base.avatar:read` `directory:employee.base.base:read` |
|         `avatar_72` | `string` | 72 * 72 pixel avatar link |
|         `avatar_240` | `string` | 240 * 240 pixel avatar link |
|         `avatar_640` | `string` | 640 * 640 pixel avatar link |
|         `avatar_origin` | `string` | Original avatar link |
|       `background_image` | `string` | Custom background cover url **Required field scopes (Satisfy any)**: `directory:employee.base.background_image:read` `directory:employee.base.base:read` |
|     `work_info` | `employee_work_entity` | Employee job information |
|       `work_country_or_region` | `string` | Work country/region **Required field scopes (Satisfy any)**: `directory:employee.work.base_work:read` `directory:employee.work.work_country_or_region:read` |
|       `work_place` | `place` | Work location, which is the City property in user's Job details **Required field scopes (Satisfy any)**: `directory:employee.work.base_work:read` `directory:employee.work.work_place:read` |
|         `place_id` | `string` | ID |
|       `work_station` | `i18n_text` | Physical seat ID **Required field scopes (Satisfy any)**: `directory:employee.work.base_work:read` `directory:employee.work.work_station:read` |
|         `default_value` | `string` | Default value |
|         `i18n_value` | `map<string, string>` | Internationalized value, key is zh_cn, ja_jp, en_us, value is the corresponding value |
|           `key` | `string` | International locale |
|           `value` | `string` | value |
|       `job_number` | `string` | job number **Data validation rules**: - Maximum length: `255` characters **Required field scopes (Satisfy any)**: `directory:employee.work.base_work:read` `directory:employee.work.job_number:read` |
|       `extension_number` | `string` | Extension number **Data validation rules**: - Maximum length: `99` characters **Required field scopes (Satisfy any)**: `directory:employee.work.base_work:read` `directory:employee.work.extension_number:read` |
|       `join_date` | `string` | Join date, e.g., 2007-03-20 **Required field scopes (Satisfy any)**: `directory:employee.work.join_date:read` `directory:employee.work.employment:read` |
|       `employment_type` | `int` | Employee type **Required field scopes (Satisfy any)**: `directory:employee.work.employment_type:read` `directory:employee.work.employment:read` |
|       `staff_status` | `int` | Employee personnel status **Optional values are**:  - `1`: on the job - `2`: leave - `3`: to be hired - `4`: Cancel onboarding - `5`: to leave  **Required field scopes (Satisfy any)**: `directory:employee.work.staff_status:read` `directory:employee.work.employment:read` |
|       `job_title` | `job_title` | Position **Required field scopes**: `directory:employee.work.job_title:read` |
|         `job_title_id` | `string` | ID **Data validation rules**: - Maximum length: `64` characters |
|       `resign_date` | `string` | Date of resignation, e.g., 2007-03-20 **Required field scopes (Satisfy any)**: `directory:employee.work.resign_date:read` `directory:employee.work.employment:read` |
|       `resign_reason` | `string` | Reason for leaving **Required field scopes (Satisfy any)**: `directory:employee.work.resign_reason:read` `directory:employee.work.employment:read` |
|       `resign_remark` | `string` | Resignation remarks **Data validation rules**: - Maximum length: `5000` characters **Required field scopes (Satisfy any)**: `directory:employee.work.resign_remark:read` `directory:employee.work.employment:read` |
|       `resign_type` | `string` | Type of turnover **Required field scopes (Satisfy any)**: `directory:employee.work.resign_type:read` `directory:employee.work.employment:read` |
|   `abnormal` | `abnormal_record` | Field exception information |
|     `id` | `string` | abnormal entity id |
|     `row_error` | `int` | row-level exception **Optional values are**:  - `0`: success - `1000`: No permission  **Data validation rules**: - Value range: `0` ～ `100` |
|     `field_errors` | `map<string, abnormal_code>` | Column-level exception, key is the field name, value is the following enumeration |
|       `key` | `string` | field name |
|       `value` | `int` | error code **Optional values are**:  - `0`: success - `1000`: No permission - `2000`: field query exception - `2003`: Field does not exist  **Data validation rules**: - Value range: `0` ～ `3000` | ### Event body example

{
    "schema": "2.0",
    "header": {
        "event_id": "148ea70e2d95f547a170b31c194ea9d5",
        "token": "",
        "create_time": "1726297928000",
        "event_type": "directory.employee.updated_v1",
        "tenant_key": "133c1eae3c0f1748",
        "app_id": "cli_a23f3400fe78901b"
    },
    "event": {
        "abnormal": {
            "field_errors": {
                "employee_curr.base_info.departments.department_id": 1001,
                "employee_curr.base_info.employee_order_in_departments": 1001
            },
            "row_error": 0
        },
        "changed_properties": [
            "base_info.email"
        ],
        "employee_curr": {
            "base_info": {
                "active_status": 5,
                  "avatar": {
                    "avatar_72": "https://i******le.larksuite.com/static-resource/v1/as******************************io~?image_size=72x72&cut_type=&quality=&format=png&sticker_format=.webp",
                    "avatar_240": "https://i******le.larksuite.com/static-resource/v1/as******************************io~?image_size=240x240&cut_type=&quality=&format=png&sticker_format=.webp",
                    "avatar_640": "https://i******le.larksuite.com/static-resource/v1/as******************************io~?image_size=640x640&cut_type=&quality=&format=png&sticker_format=.webp",
                    "avatar_origin": "https://i******le.larksuite.com/static-resource/v1/vas******************************io~?image_size=noop&cut_type=&quality=&format=png&sticker_format=.webp"
                },
                "description": "",
                "email": "122222222@qq.com",
                "employee_id": "ou_xxxx",
                "is_resigned": false,
                "leader_id": "ou_xxxx",
                "mobile": "+86155xxxxxxxx",
                "name": {
                    "another_name": "xxxx",
                    "name": {
                        "default_value": "xxxx",
                        "i18n_value": {
                            "en_us": "",
                            "ja_jp": "",
                            "zh_cn": ""
                        }
                    }
                }
            },
            "work_info": {
                "employment_type": 1,
                "extension_number": "",
                "job_number": "",
                "job_title": {
                    "job_title_id": "0"
                },
                "join_date": "2024-07-07",
                "resign_date": "",
                "resign_reason": 0,
                "resign_remark": "",
                "resign_type": 0,
                "staff_status": 5,
                "work_country_or_region": "",
                "work_place": {
                    "place_id": "0"
                },
                "work_station": {
                    "default_value": "",
                    "i18n_value": {
                        "en_us": "",
                        "ja_jp": "",
                        "zh_cn": ""
                    }
                }
            }
        },
        "employee_prev": {
            "base_info": {
                "email": "22222222@qq.com",
                "employee_id": "ou_xxxx"
            },
            "work_info": {}
        }
    }
}

### Sample code for event subscriptions

For event subscription process, refer to:Event Subscription overview，Guide for beginners:Tutorial

  Event subscription mode
  

  
	
    
package main

import (
	"context"
	"fmt"

	larkcore "github.com/larksuite/oapi-sdk-go/v3/core"
	larkevent "github.com/larksuite/oapi-sdk-go/v3/event"
	"github.com/larksuite/oapi-sdk-go/v3/event/dispatcher"
	"github.com/larksuite/oapi-sdk-go/v3/service/directory/v1"
	larkws "github.com/larksuite/oapi-sdk-go/v3/ws"
)

// SDK 使用说明 SDK user guide：https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/server-side-sdk/golang-sdk-guide/preparations
func main() {
	// 注册事件 Register event
	eventHandler := dispatcher.NewEventDispatcher("", "").
		OnP2EmployeeUpdatedV1(func(ctx context.Context, event *larkdirectory.P2EmployeeUpdatedV1) error {
			fmt.Printf("[ OnP2EmployeeUpdatedV1 access ], data: %s\n", larkcore.Prettify(event))
			return nil
		})

	// 构建 client Build client
	cli := larkws.NewClient("YOUR_APP_ID", "YOUR_APP_SECRET",
		larkws.WithEventHandler(eventHandler),
		larkws.WithLogLevel(larkcore.LogLevelDebug),
	)

	// 建立长连接 Establish persistent connection
	err := cli.Start(context.Background())

	if err != nil {
		panic(err)
	}
}

    

    
# SDK 使用说明 SDK user guide：https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/server-side-sdk/python--sdk/preparations-before-development
import lark_oapi as lark

def do_p2_directory_employee_updated_v1(data: lark.directory.v1.P2DirectoryEmployeeUpdatedV1) -> None:
    print(f'[ do_p2_directory_employee_updated_v1 access ], data: {lark.JSON.marshal(data, indent=4)}')

# 注册事件 Register event
event_handler = lark.EventDispatcherHandler.builder("", "") \
    .register_p2_directory_employee_updated_v1(do_p2_directory_employee_updated_v1) \
    .build()

def main():
    # 构建 client Build client
    cli = lark.ws.Client("APP_ID", "APP_SECRET",
                        event_handler=event_handler, log_level=lark.LogLevel.DEBUG)
    # 建立长连接 Establish persistent connection
    cli.start()

if __name__ == "__main__":
    main()

    

    

package com.example.sample;

import com.lark.oapi.core.utils.Jsons;
import com.lark.oapi.service.directory.DirectoryService;
import com.lark.oapi.service.directory.v1.model.P2EmployeeUpdatedV1;
import com.lark.oapi.event.EventDispatcher;
import com.lark.oapi.ws.Client;

// SDK 使用说明 SDK user guide：https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/server-side-sdk/java-sdk-guide/preparations
public class Sample {
    // 注册事件 Register event
    private static final EventDispatcher EVENT_HANDLER = EventDispatcher.newBuilder("", "")
            .onP2EmployeeUpdatedV1(new DirectoryService.P2EmployeeUpdatedV1Handler() {
                @Override
                public void handle(P2EmployeeUpdatedV1 event) throws Exception {
                    System.out.printf("[ onP2EmployeeUpdatedV1 access ], data: %s\n", Jsons.DEFAULT.toJson(event.getEvent()));
                }
            })
            .build();

    public static void main(String[] args) {
        // 构建 client Build client
        Client client = new Client.Builder("APP_ID", "APP_SECRET")
                .eventHandler(EVENT_HANDLER)
                .build();
        // 建立长连接 Establish persistent connection
        client.start();
    }
}
    

    
// SDK 使用说明 SDK user guide：https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/server-side-sdk/nodejs-sdk/preparation-before-development
import * as Lark from '@larksuiteoapi/node-sdk';
const baseConfig = {
    appId: 'APP_ID',
    appSecret: 'APP_SECRET'
}
// 构建 client Build client
const wsClient = new Lark.WSClient(baseConfig);
// 建立长连接 Establish persistent connection
wsClient.start({
    // 注册事件 Register event
    eventDispatcher: new Lark.EventDispatcher({}).register({
        'directory.employee.updated_v1': async (data) => {
            console.log(data);
        }
    })
});
    

  
  
	
    
package main

import (
	"context"
	"fmt"
	"net/http"

	larkcore "github.com/larksuite/oapi-sdk-go/v3/core"
	"github.com/larksuite/oapi-sdk-go/v3/core/httpserverext"
	larkevent "github.com/larksuite/oapi-sdk-go/v3/event"
	"github.com/larksuite/oapi-sdk-go/v3/event/dispatcher"
	"github.com/larksuite/oapi-sdk-go/v3/service/directory/v1"
)

// SDK 使用说明 SDK user guide：https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/server-side-sdk/golang-sdk-guide/preparations
func main() {
	// 注册事件 Register event
	eventHandler := dispatcher.NewEventDispatcher("", "").
		OnP2EmployeeUpdatedV1(func(ctx context.Context, event *larkdirectory.P2EmployeeUpdatedV1) error {
			fmt.Printf("[ OnP2EmployeeUpdatedV1 access ], data: %s\n", larkcore.Prettify(event))
			return nil
		})

	// 创建路由处理器 Create route handler
	http.HandleFunc("/webhook/event", httpserverext.NewEventHandlerFunc(handler, larkevent.WithLogLevel(larkcore.LogLevelDebug)))

	err := http.ListenAndServe(":7777", nil)

	if err != nil {
		panic(err)
	}
}

    

    
# SDK 使用说明 SDK user guide：https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/server-side-sdk/python--sdk/preparations-before-development
from flask import Flask
from lark_oapi.adapter.flask import *
import lark_oapi as lark

app = Flask(__name__)

def do_p2_directory_employee_updated_v1(data: lark.directory.v1.P2DirectoryEmployeeUpdatedV1) -> None:
    print(f'[ do_p2_directory_employee_updated_v1 access ], data: {lark.JSON.marshal(data, indent=4)}')

# 注册事件 Register event
event_handler = lark.EventDispatcherHandler.builder("", "") \
    .register_p2_directory_employee_updated_v1(do_p2_directory_employee_updated_v1) \
    .build()

# 创建路由处理器 Create route handler
@app.route("/webhook/event", methods=["POST"])
def event():
    resp = event_handler.do(parse_req())
    return parse_resp(resp)

if __name__ == "__main__":
    app.run(port=7777)

    

    

package com.lark.oapi.sample.event;

import com.lark.oapi.core.utils.Jsons;
import com.lark.oapi.service.directory.DirectoryService;
import com.lark.oapi.service.directory.v1.model.P2EmployeeUpdatedV1;
import com.lark.oapi.sdk.servlet.ext.ServletAdapter;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;

// SDK 使用说明 SDK user guide：https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/server-side-sdk/java-sdk-guide/preparations
@RestController
public class EventController {

    // 注册事件 Register event
    private static final EventDispatcher EVENT_HANDLER = EventDispatcher.newBuilder("verificationToken", "encryptKey")
            .onP2EmployeeUpdatedV1(new DirectoryService.P2EmployeeUpdatedV1Handler() {
                @Override
                public void handle(P2EmployeeUpdatedV1 event) throws Exception {
                    System.out.printf("[ onP2EmployeeUpdatedV1 access ], data: %s\n", Jsons.DEFAULT.toJson(event.getEvent()));
                }
            })
            .build();

    // 注入 ServletAdapter 实例 Inject ServletAdapter instance
    @Autowired
    private ServletAdapter servletAdapter;

    // 创建路由处理器 Create route handler
    @RequestMapping("/webhook/event")
    public void event(HttpServletRequest request, HttpServletResponse response)
            throws Throwable {
        // 回调扩展包提供的事件回调处理器 Callback handler provided by the extension package
        servletAdapter.handleEvent(request, response, EVENT_DISPATCHER);
    }
}
    

    
// SDK 使用说明 SDK user guide：https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/server-side-sdk/nodejs-sdk/preparation-before-development
import http from 'http';
import * as lark from '@larksuiteoapi/node-sdk';

// 注册事件 Register event
const eventDispatcher = new lark.EventDispatcher({
    encryptKey: '',
    verificationToken: '',
}).register({
    'directory.employee.updated_v1': async (data) => {
        console.log(data);
        return 'success';
    },
});

const server = http.createServer();
// 创建路由处理器 Create route handler
server.on('request', lark.adaptDefault('/webhook/event', eventDispatcher));
server.listen(3000);
