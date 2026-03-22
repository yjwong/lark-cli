---
title: "Contact Scope Updated"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/scope/events/updated"
service: "contact-v3"
resource: "scope"
section: "Contacts"
scopes:
  - "contact:contact:readonly_as_app"
  - "contact:contact:access_as_app"
field_scopes:
  - "contact:contact:readonly_as_app"
  - "contact:department.base:readonly"
  - "contact:department.organize:readonly"
  - "contact:user.base:readonly"
  - "contact:user.department:readonly"
  - "contact:user.employee:readonly"
  - "contact:user.gender:readonly"
  - "contact:user.employee_id:readonly"
  - "contact:user.email:readonly"
  - "contact:user.phone:readonly"
  - "contact:user.user_geo"
updated: "1664162620000"
---

# Contact scope changed

This event is triggered when the range of contacts data that an app can access is changed. Apps subscribing to this event will receive the event.{Usage Examples}(url=/api/tools/api_explore/api_explore_config?project=contact&version=v3&resource=scope&event=updated)

## Event
| Facts |  |
| --- | --- |
| Event type | contact.scope.updated_v3 |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `contact:contact:readonly_as_app` `contact:contact:access_as_app` |
| Required field scopes | > The response body of the API contains the following sensitive fields, and they will be returned only after corresponding scopes are added. If you do not need the fields, it is not recommended that you request the scopes. `contact:contact:readonly_as_app` `contact:department.base:readonly` `contact:department.organize:readonly` `contact:user.base:readonly` `contact:user.department:readonly` `contact:user.employee:readonly` `contact:user.gender:readonly` `contact:user.employee_id:readonly` `contact:user.email:readonly` `contact:user.phone:readonly` `contact:user.user_geo` |
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
|   `added` | `scope` | Object added when the contact scope is changed |
|     `departments` | `department[]` | Department object |
|       `name` | `string` | Department name **Data validation rules**: - Minimum length: `1` characters **Required field scopes (Satisfy any)**: `contact:contact:readonly_as_app` `contact:department.base:readonly` |
|       `i18n_name` | `department_i18n_name` | Internationalized department name **Required field scopes (Satisfy any)**: `contact:contact:readonly_as_app` `contact:department.base:readonly` |
|         `zh_cn` | `string` | Department's Chinese name |
|         `ja_jp` | `string` | Department's Japanese name |
|         `en_us` | `string` | Department's English name |
|       `parent_department_id` | `string` | Parent department ID * This value is "0" for a root department. **Required field scopes (Satisfy any)**: `contact:contact:readonly_as_app` `contact:department.organize:readonly` |
|       `department_id` | `string` | Department's custom department ID **Data validation rules**: - Maximum length: `64` characters - Regular expression: `^0|[^od][A-Za-z0-9]*` **Required field scopes (Satisfy any)**: `contact:contact:readonly_as_app` `contact:department.base:readonly` |
|       `open_department_id` | `string` | Department's open_id |
|       `leader_user_id` | `string` | Department manager's user ID **Required field scopes (Satisfy any)**: `contact:contact:readonly_as_app` `contact:department.organize:readonly` |
|       `chat_id` | `string` | Department group ID **Required field scopes (Satisfy any)**: `contact:contact:readonly_as_app` `contact:department.base:readonly` |
|       `order` | `string` | Department order, i.e. the order in which the department is displayed among the departments at the same level. **Required field scopes (Satisfy any)**: `contact:contact:readonly_as_app` `contact:department.organize:readonly` |
|       `unit_ids` | `string[]` | List of the department unit's custom IDs. Only one custom ID is supported currently. **Required field scopes (Satisfy any)**: `contact:contact:readonly_as_app` `contact:department.organize:readonly` |
|       `member_count` | `int` | Number of users under the department **Required field scopes (Satisfy any)**: `contact:contact:readonly_as_app` `contact:department.organize:readonly` |
|       `status` | `department_status` | Department status **Required field scopes (Satisfy any)**: `contact:contact:readonly_as_app` `contact:department.base:readonly` |
|         `is_deleted` | `boolean` | Whether it is deleted |
|       `leaders` | `departmentLeader[]` | department leader |
|         `leaderType` | `int` | department leaderType **Optional values are**:  - `1`: main leader - `2`: deputy leader  |
|         `leaderID` | `string` | leader id **Required field scopes (Satisfy any)**: `contact:contact:readonly_as_app` `contact:department.organize:readonly` |
|     `users` | `user[]` | User |
|       `union_id` | `string` | User's union_id. For information about different IDs, see User-related IDs. |
|       `user_id` | `string` | Unique identifier of the user in the tenant, namely the user's user_id. For information of different IDs, refer to User-related IDs. **Required field scopes**: `contact:user.employee_id:readonly` |
|       `open_id` | `string` | User's open_id. For information about different IDs, see User-related IDs. |
|       `name` | `string` | User name **Data validation rules**: - Minimum length: `1` characters **Required field scopes (Satisfy any)**: `contact:contact:readonly_as_app` `contact:user.base:readonly` |
|       `en_name` | `string` | English name **Required field scopes (Satisfy any)**: `contact:contact:readonly_as_app` `contact:user.base:readonly` |
|       `nickname` | `string` | Alias **Required field scopes (Satisfy any)**: `contact:contact:readonly_as_app` `contact:user.base:readonly` |
|       `email` | `string` | Email **Required field scopes**: `contact:user.email:readonly` |
|       `mobile` | `string` | The mobile number cannot be repeated within the organization. Note that the area code must contain the plus sign +. **Required field scopes**: `contact:user.phone:readonly` |
|       `gender` | `int` | Gender **Optional values are**:  - `0`: Unknown - `1`: Male - `2`: Female  **Required field scopes (Satisfy any)**: `contact:contact:readonly_as_app` `contact:user.gender:readonly` |
|       `avatar` | `avatar_info` | User's profile photo **Required field scopes (Satisfy any)**: `contact:contact:readonly_as_app` `contact:user.base:readonly` |
|         `avatar_72` | `string` | Link of profile photo in 72*72 px |
|         `avatar_240` | `string` | Link of profile photo in 240*240 px |
|         `avatar_640` | `string` | Link of profile photo in 640*640 px |
|         `avatar_origin` | `string` | Link of the original profile photo |
|       `status` | `user_status` | User status **Required field scopes (Satisfy any)**: `contact:contact:readonly_as_app` `contact:user.employee:readonly` |
|         `is_frozen` | `boolean` | Whether the user is frozen. |
|         `is_resigned` | `boolean` | Whether the user is terminated. |
|         `is_activated` | `boolean` | Whether the user is activated. |
|         `is_exited` | `boolean` | Whether the user is voluntary. The user's status will be automatically changed to terminated after the user has left for a specific time of period. |
|         `is_unjoin` | `boolean` | Whether the user has not joined. The user must confirm in person to join the team. |
|       `leader_user_id` | `string` | User ID of the user's direct manager. The ID value matches the user_id_type in the query parameter. For information about different IDs, see User-related IDs. **Required field scopes (Satisfy any)**: `contact:contact:readonly_as_app` `contact:user.department:readonly` |
|       `city` | `string` | City **Required field scopes (Satisfy any)**: `contact:contact:readonly_as_app` `contact:user.employee:readonly` |
|       `country` | `string` | Country or region code. Refer to the list of country/region codes for details. **Required field scopes (Satisfy any)**: `contact:contact:readonly_as_app` `contact:user.employee:readonly` |
|       `work_station` | `string` | Seat ID **Required field scopes (Satisfy any)**: `contact:contact:readonly_as_app` `contact:user.employee:readonly` |
|       `join_time` | `int` | Date of employment **Required field scopes (Satisfy any)**: `contact:contact:readonly_as_app` `contact:user.employee:readonly` |
|       `employee_no` | `string` | Employee ID **Required field scopes (Satisfy any)**: `contact:contact:readonly_as_app` `contact:user.employee:readonly` |
|       `employee_type` | `int` | Employee type. Optional values are: - `1`: Regular - `2`: Intern - `3`: Outsourcing - `4`: Contractor - `5`: Consultant The int value of the custom employee type can also be read, and the name of the custom employee type for the tenant can be obtained through the API: Obtain workforce type **Required field scopes (Satisfy any)**: `contact:contact:readonly_as_app` `contact:user.employee:readonly` |
|       `custom_attrs` | `user_custom_attr[]` | Custom field. Make sure your company administrator has enabled "Allow to invoke by Open Platform's Contacts API" in Admin > Organization > Member Profile > Custom Fields > Global Settings; otherwise, this field will not take effect or be returned. **Required field scopes (Satisfy any)**: `contact:contact:readonly_as_app` `contact:user.employee:readonly` |
|         `type` | `string` | Custom field type - `TEXT`: Text - `HREF`: Webpage - `ENUMERATION`: Enumeration - `PICTURE_ENUM`: Image - `GENERIC_USER`: User FAQs about custom fields |
|         `id` | `string` | Custom field ID |
|         `value` | `user_custom_attr_value` | Custom field value |
|           `text` | `string` | When the field type is `TEXT`, this parameter defines the field value and is required; when the field type is `HREF`, this parameter defines the webpage title and is required. |
|           `url` | `string` | When the field type is HREF, this parameter defines the default URL. |
|           `pc_url` | `string` | When the field type is HREF, this parameter defines URL on the PC. |
|           `option_id` | `string` | When the field type is ENUMERATION or PICTURE_ENUM, this parameter defines the option value. |
|           `option_value` | `string` | Option value |
|           `name` | `string` | Name |
|           `picture_url` | `string` | Image link |
|           `generic_user` | `custom_attr_generic_user` | When the field type is GENERIC_USER, this parameter defines the user referenced. |
|             `id` | `string` | User's user_id |
|             `type` | `int` | User type    1: User |
|       `enterprise_email` | `string` | Business email. Make sure that Lark Business Mail has been enabled in Admin. **Required field scopes (Satisfy any)**: `contact:contact:readonly_as_app` `contact:user.employee:readonly` |
|       `job_title` | `string` | Job title **Required field scopes (Satisfy any)**: `contact:contact:readonly_as_app` `contact:user.employee:readonly` |
|       `is_frozen` | `boolean` | Whether to freeze the user |
|       `geo` | `string` | geo **Required field scopes**: `contact:user.user_geo` |
|     `user_groups` | `user_group[]` | User group object |
|       `user_group_id` | `string` | Custom ID of the user group **Data validation rules**: - Length range: `1` ～ `64` characters |
|       `name` | `string` | User group name **Data validation rules**: - Length range: `1` ～ `100` characters |
|       `type` | `int` | User group type **Optional values are**:  - `1`: Common user group - `2`: Dynamic group  |
|       `member_count` | `int` | Number of members |
|       `status` | `int` | User group status **Optional values are**:  - `0`: Unknown - `1`: Calculation completed - `2`: Calculating - `3`: Calculation failed  |
|   `removed` | `scope` | Object removed when the range of contacts data that an app can access is changed |
|     `departments` | `department[]` | Department object |
|       `name` | `string` | Department name **Data validation rules**: - Minimum length: `1` characters **Required field scopes (Satisfy any)**: `contact:contact:readonly_as_app` `contact:department.base:readonly` |
|       `i18n_name` | `department_i18n_name` | Internationalized department name **Required field scopes (Satisfy any)**: `contact:contact:readonly_as_app` `contact:department.base:readonly` |
|         `zh_cn` | `string` | Department's Chinese name |
|         `ja_jp` | `string` | Department's Japanese name |
|         `en_us` | `string` | Department's English name |
|       `parent_department_id` | `string` | Parent department ID * This value is "0" for a root department. **Required field scopes (Satisfy any)**: `contact:contact:readonly_as_app` `contact:department.organize:readonly` |
|       `department_id` | `string` | Department's custom department ID **Data validation rules**: - Maximum length: `64` characters - Regular expression: `^0|[^od][A-Za-z0-9]*` **Required field scopes (Satisfy any)**: `contact:contact:readonly_as_app` `contact:department.base:readonly` |
|       `open_department_id` | `string` | Department's open_id |
|       `leader_user_id` | `string` | Department manager's user ID **Required field scopes (Satisfy any)**: `contact:contact:readonly_as_app` `contact:department.organize:readonly` |
|       `chat_id` | `string` | Department group ID **Required field scopes (Satisfy any)**: `contact:contact:readonly_as_app` `contact:department.base:readonly` |
|       `order` | `string` | Department order, i.e. the order in which the department is displayed among the departments at the same level. **Required field scopes (Satisfy any)**: `contact:contact:readonly_as_app` `contact:department.organize:readonly` |
|       `unit_ids` | `string[]` | List of the department unit's custom IDs. Only one custom ID is supported currently. **Required field scopes (Satisfy any)**: `contact:contact:readonly_as_app` `contact:department.organize:readonly` |
|       `member_count` | `int` | Number of users under the department **Required field scopes (Satisfy any)**: `contact:contact:readonly_as_app` `contact:department.organize:readonly` |
|       `status` | `department_status` | Department status **Required field scopes (Satisfy any)**: `contact:contact:readonly_as_app` `contact:department.base:readonly` |
|         `is_deleted` | `boolean` | Whether it is deleted |
|       `leaders` | `departmentLeader[]` | department leader |
|         `leaderType` | `int` | department leaderType **Optional values are**:  - `1`: main leader - `2`: deputy leader  |
|         `leaderID` | `string` | leader id **Required field scopes (Satisfy any)**: `contact:contact:readonly_as_app` `contact:department.organize:readonly` |
|     `users` | `user[]` | User |
|       `union_id` | `string` | User's union_id. For information about different IDs, see User-related IDs. |
|       `user_id` | `string` | Unique identifier of the user in the tenant, namely the user's user_id. For information of different IDs, refer to User-related IDs. **Required field scopes**: `contact:user.employee_id:readonly` |
|       `open_id` | `string` | User's open_id. For information about different IDs, see User-related IDs. |
|       `name` | `string` | User name **Data validation rules**: - Minimum length: `1` characters **Required field scopes (Satisfy any)**: `contact:contact:readonly_as_app` `contact:user.base:readonly` |
|       `en_name` | `string` | English name **Required field scopes (Satisfy any)**: `contact:contact:readonly_as_app` `contact:user.base:readonly` |
|       `nickname` | `string` | Alias **Required field scopes (Satisfy any)**: `contact:contact:readonly_as_app` `contact:user.base:readonly` |
|       `email` | `string` | Email **Required field scopes**: `contact:user.email:readonly` |
|       `mobile` | `string` | The mobile number cannot be repeated within the organization. Note that the area code must contain the plus sign +. **Required field scopes**: `contact:user.phone:readonly` |
|       `gender` | `int` | Gender **Optional values are**:  - `0`: Unknown - `1`: Male - `2`: Female  **Required field scopes (Satisfy any)**: `contact:contact:readonly_as_app` `contact:user.gender:readonly` |
|       `avatar` | `avatar_info` | User's profile photo **Required field scopes (Satisfy any)**: `contact:contact:readonly_as_app` `contact:user.base:readonly` |
|         `avatar_72` | `string` | Link of profile photo in 72*72 px |
|         `avatar_240` | `string` | Link of profile photo in 240*240 px |
|         `avatar_640` | `string` | Link of profile photo in 640*640 px |
|         `avatar_origin` | `string` | Link of the original profile photo |
|       `status` | `user_status` | User status **Required field scopes (Satisfy any)**: `contact:contact:readonly_as_app` `contact:user.employee:readonly` |
|         `is_frozen` | `boolean` | Whether the user is frozen. |
|         `is_resigned` | `boolean` | Whether the user is terminated. |
|         `is_activated` | `boolean` | Whether the user is activated. |
|         `is_exited` | `boolean` | Whether the user is voluntary. The user's status will be automatically changed to terminated after the user has left for a specific time of period. |
|         `is_unjoin` | `boolean` | Whether the user has not joined. The user must confirm in person to join the team. |
|       `leader_user_id` | `string` | User ID of the user's direct manager. The ID value matches the user_id_type in the query parameter. For information about different IDs, see User-related IDs. **Required field scopes (Satisfy any)**: `contact:contact:readonly_as_app` `contact:user.department:readonly` |
|       `city` | `string` | City **Required field scopes (Satisfy any)**: `contact:contact:readonly_as_app` `contact:user.employee:readonly` |
|       `country` | `string` | Country or region code. Refer to the list of country/region codes for details. **Required field scopes (Satisfy any)**: `contact:contact:readonly_as_app` `contact:user.employee:readonly` |
|       `work_station` | `string` | Seat ID **Required field scopes (Satisfy any)**: `contact:contact:readonly_as_app` `contact:user.employee:readonly` |
|       `join_time` | `int` | Date of employment **Required field scopes (Satisfy any)**: `contact:contact:readonly_as_app` `contact:user.employee:readonly` |
|       `employee_no` | `string` | Employee ID **Required field scopes (Satisfy any)**: `contact:contact:readonly_as_app` `contact:user.employee:readonly` |
|       `employee_type` | `int` | Employee type. Optional values are: - `1`: Regular - `2`: Intern - `3`: Outsourcing - `4`: Contractor - `5`: Consultant The int value of the custom employee type can also be read, and the name of the custom employee type for the tenant can be obtained through the API: Obtain workforce type **Required field scopes (Satisfy any)**: `contact:contact:readonly_as_app` `contact:user.employee:readonly` |
|       `custom_attrs` | `user_custom_attr[]` | Custom field. Make sure your company administrator has enabled "Allow to invoke by Open Platform's Contacts API" in Admin > Organization > Member Profile > Custom Fields > Global Settings; otherwise, this field will not take effect or be returned. **Required field scopes (Satisfy any)**: `contact:contact:readonly_as_app` `contact:user.employee:readonly` |
|         `type` | `string` | Custom field type - `TEXT`: Text - `HREF`: Webpage - `ENUMERATION`: Enumeration - `PICTURE_ENUM`: Image - `GENERIC_USER`: User FAQs about custom fields |
|         `id` | `string` | Custom field ID |
|         `value` | `user_custom_attr_value` | Custom field value |
|           `text` | `string` | When the field type is `TEXT`, this parameter defines the field value and is required; when the field type is `HREF`, this parameter defines the webpage title and is required. |
|           `url` | `string` | When the field type is HREF, this parameter defines the default URL. |
|           `pc_url` | `string` | When the field type is HREF, this parameter defines URL on the PC. |
|           `option_id` | `string` | When the field type is ENUMERATION or PICTURE_ENUM, this parameter defines the option value. |
|           `option_value` | `string` | Option value |
|           `name` | `string` | Name |
|           `picture_url` | `string` | Image link |
|           `generic_user` | `custom_attr_generic_user` | When the field type is GENERIC_USER, this parameter defines the user referenced. |
|             `id` | `string` | User's user_id |
|             `type` | `int` | User type    1: User |
|       `enterprise_email` | `string` | Business email. Make sure that Lark Business Mail has been enabled in Admin. **Required field scopes (Satisfy any)**: `contact:contact:readonly_as_app` `contact:user.employee:readonly` |
|       `job_title` | `string` | Job title **Required field scopes (Satisfy any)**: `contact:contact:readonly_as_app` `contact:user.employee:readonly` |
|       `is_frozen` | `boolean` | Whether to freeze the user |
|       `geo` | `string` | geo **Required field scopes**: `contact:user.user_geo` |
|     `user_groups` | `user_group[]` | User group object |
|       `user_group_id` | `string` | Custom ID of the user group **Data validation rules**: - Length range: `1` ～ `64` characters |
|       `name` | `string` | User group name **Data validation rules**: - Length range: `1` ～ `100` characters |
|       `type` | `int` | User group type **Optional values are**:  - `1`: Common user group - `2`: Dynamic group  |
|       `member_count` | `int` | Number of members |
|       `status` | `int` | User group status **Optional values are**:  - `0`: Unknown - `1`: Calculation completed - `2`: Calculating - `3`: Calculation failed  | ### Event body example

{
    "schema": "2.0",
    "header": {
        "event_id": "5e3702a84e847582be8db7fb73283c02",
        "event_type": "contact.scope.updated_v3",
        "create_time": "1608725989000",
        "token": "rvaYgkND1GOiu5MM0E1rncYC6PLtF7JV",
        "app_id": "cli_9f5343c580712544",
        "tenant_key": "2ca1d211f64f6438"
    },
    "event": {
        "added": {
            "departments": [
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
                    },
                    "leaders": [
                        {
                            "leaderType": 1,
                            "leaderID": "ou_7dab8a3d3cdcc9da365777c7ad535d62"
                        }
                    ]
                }
            ],
            "users": [
                {
                    "union_id": "on_94a1ee5551019f18cd73d9f111898cf2",
                    "user_id": "3e3cf96b",
                    "open_id": "ou_7dab8a3d3cdcc9da365777c7ad535d62",
                    "name": "John",
                    "en_name": "San Zhang",
                    "nickname": "Alex Zhang",
                    "email": "zhangsan@gmail.com",
                    "mobile": "+41446681800",
                    "gender": 1,
                    "avatar": {
                        "avatar_72": "https://foo.icon.com/xxxx",
                        "avatar_240": "https://foo.icon.com/xxxx",
                        "avatar_640": "https://foo.icon.com/xxxx",
                        "avatar_origin": "https://foo.icon.com/xxxx"
                    },
                    "status": {
                        "is_frozen": false,
                        "is_resigned": false,
                        "is_activated": true,
                        "is_exited": false,
                        "is_unjoin": false
                    },
                    "leader_user_id": "ou_7dab8a3d3cdcc9da365777c7ad535d62",
                    "city": "Hangzhou",
                    "country": "CN",
                    "work_station": "North building-H34",
                    "join_time": 2147483647,
                    "employee_no": "1",
                    "employee_type": 1,
                    "custom_attrs": [
                        {
                            "type": "TEXT",
                            "id": "DemoId",
                            "value": {
                                "text": "DemoText",
                                "url": "http://www.larksuite.com",
                                "pc_url": "http://www.larksuite.com",
                                "option_id": "edcvfrtg",
                                "option_value": "option",
                                "name": "name",
                                "picture_url": "https://xxxxxxxxxxxxxxxxxx",
                                "generic_user": {
                                    "id": "9b2fabg5",
                                    "type": 1
                                }
                            }
                        }
                    ],
                    "enterprise_email": "demo@mail.com",
                    "job_title": "xxxxx",
                    "is_frozen": false,
                    "geo": "cn"
                }
            ],
            "user_groups": [
                {
                    "user_group_id": "test",
                    "name": "userGroupName",
                    "type": 1,
                    "member_count": 10,
                    "status": 1
                }
            ]
        },
        "removed": {
            "departments": [
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
                    },
                    "leaders": [
                        {
                            "leaderType": 1,
                            "leaderID": "ou_7dab8a3d3cdcc9da365777c7ad535d62"
                        }
                    ]
                }
            ],
            "users": [
                {
                    "union_id": "on_94a1ee5551019f18cd73d9f111898cf2",
                    "user_id": "3e3cf96b",
                    "open_id": "ou_7dab8a3d3cdcc9da365777c7ad535d62",
                    "name": "John",
                    "en_name": "San Zhang",
                    "nickname": "Alex Zhang",
                    "email": "zhangsan@gmail.com",
                    "mobile": "+41446681800",
                    "gender": 1,
                    "avatar": {
                        "avatar_72": "https://foo.icon.com/xxxx",
                        "avatar_240": "https://foo.icon.com/xxxx",
                        "avatar_640": "https://foo.icon.com/xxxx",
                        "avatar_origin": "https://foo.icon.com/xxxx"
                    },
                    "status": {
                        "is_frozen": false,
                        "is_resigned": false,
                        "is_activated": true,
                        "is_exited": false,
                        "is_unjoin": false
                    },
                    "leader_user_id": "ou_7dab8a3d3cdcc9da365777c7ad535d62",
                    "city": "Hangzhou",
                    "country": "CN",
                    "work_station": "North building-H34",
                    "join_time": 2147483647,
                    "employee_no": "1",
                    "employee_type": 1,
                    "custom_attrs": [
                        {
                            "type": "TEXT",
                            "id": "DemoId",
                            "value": {
                                "text": "DemoText",
                                "url": "http://www.larksuite.com",
                                "pc_url": "http://www.larksuite.com",
                                "option_id": "edcvfrtg",
                                "option_value": "option",
                                "name": "name",
                                "picture_url": "https://xxxxxxxxxxxxxxxxxx",
                                "generic_user": {
                                    "id": "9b2fabg5",
                                    "type": 1
                                }
                            }
                        }
                    ],
                    "enterprise_email": "demo@mail.com",
                    "job_title": "xxxxx",
                    "is_frozen": false,
                    "geo": "cn"
                }
            ],
            "user_groups": [
                {
                    "user_group_id": "test",
                    "name": "userGroupName",
                    "type": 1,
                    "member_count": 10,
                    "status": 1
                }
            ]
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
	"github.com/larksuite/oapi-sdk-go/v3/service/contact/v3"
	larkws "github.com/larksuite/oapi-sdk-go/v3/ws"
)

// SDK 使用说明 SDK user guide：https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/server-side-sdk/golang-sdk-guide/preparations
func main() {
	// 注册事件 Register event
	eventHandler := dispatcher.NewEventDispatcher("", "").
		OnP2ScopeUpdatedV3(func(ctx context.Context, event *larkcontact.P2ScopeUpdatedV3) error {
			fmt.Printf("[ OnP2ScopeUpdatedV3 access ], data: %s\n", larkcore.Prettify(event))
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

def do_p2_contact_scope_updated_v3(data: lark.contact.v3.P2ContactScopeUpdatedV3) -> None:
    print(f'[ do_p2_contact_scope_updated_v3 access ], data: {lark.JSON.marshal(data, indent=4)}')

# 注册事件 Register event
event_handler = lark.EventDispatcherHandler.builder("", "") \
    .register_p2_contact_scope_updated_v3(do_p2_contact_scope_updated_v3) \
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
import com.lark.oapi.service.contact.ContactService;
import com.lark.oapi.service.contact.v3.model.P2ScopeUpdatedV3;
import com.lark.oapi.event.EventDispatcher;
import com.lark.oapi.ws.Client;

// SDK 使用说明 SDK user guide：https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/server-side-sdk/java-sdk-guide/preparations
public class Sample {
    // 注册事件 Register event
    private static final EventDispatcher EVENT_HANDLER = EventDispatcher.newBuilder("", "")
            .onP2ScopeUpdatedV3(new ContactService.P2ScopeUpdatedV3Handler() {
                @Override
                public void handle(P2ScopeUpdatedV3 event) throws Exception {
                    System.out.printf("[ onP2ScopeUpdatedV3 access ], data: %s\n", Jsons.DEFAULT.toJson(event.getEvent()));
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
        'contact.scope.updated_v3': async (data) => {
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
	"github.com/larksuite/oapi-sdk-go/v3/service/contact/v3"
)

// SDK 使用说明 SDK user guide：https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/server-side-sdk/golang-sdk-guide/preparations
func main() {
	// 注册事件 Register event
	eventHandler := dispatcher.NewEventDispatcher("", "").
		OnP2ScopeUpdatedV3(func(ctx context.Context, event *larkcontact.P2ScopeUpdatedV3) error {
			fmt.Printf("[ OnP2ScopeUpdatedV3 access ], data: %s\n", larkcore.Prettify(event))
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

def do_p2_contact_scope_updated_v3(data: lark.contact.v3.P2ContactScopeUpdatedV3) -> None:
    print(f'[ do_p2_contact_scope_updated_v3 access ], data: {lark.JSON.marshal(data, indent=4)}')

# 注册事件 Register event
event_handler = lark.EventDispatcherHandler.builder("", "") \
    .register_p2_contact_scope_updated_v3(do_p2_contact_scope_updated_v3) \
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
import com.lark.oapi.service.contact.ContactService;
import com.lark.oapi.service.contact.v3.model.P2ScopeUpdatedV3;
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
            .onP2ScopeUpdatedV3(new ContactService.P2ScopeUpdatedV3Handler() {
                @Override
                public void handle(P2ScopeUpdatedV3 event) throws Exception {
                    System.out.printf("[ onP2ScopeUpdatedV3 access ], data: %s\n", Jsons.DEFAULT.toJson(event.getEvent()));
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
    'contact.scope.updated_v3': async (data) => {
        console.log(data);
        return 'success';
    },
});

const server = http.createServer();
// 创建路由处理器 Create route handler
server.on('request', lark.adaptDefault('/webhook/event', eventDispatcher));
server.listen(3000);
