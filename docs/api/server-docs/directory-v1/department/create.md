---
title: "Create department"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/directory-v1/department/create"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/directory/v1/departments"
service: "directory-v1"
resource: "department"
section: "Organization"
rate_limit: "5 per second"
scopes:
  - "directory:department.create:write"
  - "directory:department:write"
field_scopes:
  - "directory:employee.base.external_id:read"
updated: "1756439286000"
---

# Create Department

This interface is used to create a new department in the enterprise address book, supporting the setting of information such as department name, parent department, and department head. 

> Attention:
> - Departments can only be created under departments within the scope of the address book authorization of the current application. If you want to create a sub-department under the root department, you must have full permissions. You can view the address book authorization scope in the developer background - application details - permission management.
> - The user_access_token supported in this interface defaults to the administrator user, and the administrator management scope will be verified. When the user has multiple administrator identities to create departments, the administrator management scope takes the maximum set.For information on administrator permissions, see the Help Center article: [Administrator creates administrator role and assigns permissions](https://www.larksuite.com/hc/zh-CN/articles/360043595213)
> - After you have the permission of this interface, you can write the department information. However, after creating the department, only the field data that the application has permission is returned. If you need to specify the field, please apply for the corresponding permission according to the description in the document.
> This interface is only open to self-built applications.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/directory/v1/departments |
| HTTP Method | POST |
| Rate Limit | 5 per second |
| Supported app types | custom |
| Required scopes Enable any scope from the list | `directory:department.create:write` `directory:department:write` |
| Required field scopes | > The response body of the API contains the following sensitive fields, and they will be returned only after corresponding scopes are added. If you do not need the fields, it is not recommended that you request the scopes. `directory:employee.base.external_id:read` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `employee_id_type` | `string` | No | User ID categories **Example value**: open_id **Optional values are**:  - `open_id`: Identifies the identity of a user in an application. The same user has different Open IDs in different applications. Learn more: How to get an Open ID - `union_id`: Identifies the identity of a user under an application developer. The Union ID of the same user in an application under the same developer is the same, and the Union ID in an application under different developers is different. With Union ID, application developers can associate the identity of the same user in multiple applications. Learn more: How to get Union ID? - `employee_id`: The unique identifier of an employee within the enterprise. Supports customization, the system automatically generates it if it is not customized. ID supports modification. How to get employee_id: - Enterprise administrators go to the Admin Console > Organizational Structure > Members and Departments page, click Member Details to query employee IDs - Through the interface of Get employee list in batches, query employee ID by mobile phone number or email.  **Default value**: `open_id` **When the value is `employee_id`, the following field scopes are required**: `directory:employee.base.external_id:read` |
| `department_id_type` | `string` | No | The type of department ID used in this call **Example value**: open_department_id **Optional values are**:  - `department_id`: Used to identify a unique department within a tenant - `open_department_id`: Used to identify a department in a specific application, where the open_department_id is different for the same department in different applications.  **Default value**: `open_department_id` | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `department` | `create_department` | Yes | Create Department |
|   `custom_department_id` | `string` | No | Identifies a unique department within the tenant, supports customization and is automatically generated by the system when not customized. id supports modification.Notes: 1. in addition to the need to meet the regular rules, at the same time can not start with od-. 2. Regular checksum: ^[a-zA-Z0-9][a-zA-Z0-9_\-@.] {0,63}$ **Example value**: "eersdf" |
|   `name` | `i18n_text` | No | Department name, up to 100 characters |
|     `default_value` | `string` | Yes | Default value **Example value**: "张三 **Data Validation Rules**: Length range: 1-64 characters" |
|     `i18n_value` | `map<string, string>` | No | Internationalization value, key is zh_cn, ja_jp, en_us, and value is the corresponding value. **Example value**: { "zh_cn": "中文", "ja_jp": "ja_jp_name", "en_us": "en_us_name" } |
|   `parent_department_id` | `string` | No | Parent department ID, consistent with the department_id_type type. if the parent department is the root department, the value of this parameter is "0" **Example value**: "h121900" |
|   `leaders` | `department_leader[]` | No | Department leader **Data validation rules**: - Length range: `0` ～ `20` |
|     `leader_type` | `int` | Yes | Department leader type **Example value**: 1 **Optional values are**:  - `1`: main - `2`: Vice  |
|     `leader_id` | `string` | Yes | Department leader id, consistent with the employee_id_type type **Example value**: "u273y71" |
|   `order_weight` | `string` | No | The sorting weight under the superior department, with the returned results sorted in descending order by order_weight **Example value**: "100" |
|   `enabled_status` | `boolean` | No | Enabled status **Example value**: true |
|   `custom_field_values` | `custom_field_value[]` | No | Departmental custom field values **Data validation rules**: - Length range: `0` ～ `100` |
|     `field_type` | `string` | No | Custom field types **Example value**: "1" **Optional values are**:  - `1`: Text - `2`: URL - `3`: Options - `4`: Members - `9`: - `10`: Multiple selection options - `11`: Members list  |
|     `text_value` | `i18n_text` | No | Text Field Values |
|       `default_value` | `string` | Yes | Default value **Example value**: "张三" |
|       `i18n_value` | `map<string, string>` | No | Internationalization value, key is zh_cn, ja_jp, en_us, and value is the corresponding value. **Example value**: { "zh_cn": "中文", "ja_jp": "ja_jp_name", "en_us": "en_us_name" } |
|     `url_value` | `url_value` | No | Web Link Field Values |
|       `link_text` | `i18n_text` | Yes | Link title |
|         `default_value` | `string` | Yes | Default value **Example value**: "张三" |
|         `i18n_value` | `map<string, string>` | No | Internationalization value, key is zh_cn, ja_jp, en_us, and value is the corresponding value. **Example value**: {"zh_cn":"张三"} |
|       `url` | `string` | Yes | Mobile web link **Example value**: "https://m.bytedance.com/afnasjfna" |
|       `pcurl` | `string` | Yes | Desktop web link **Example value**: "http://www.fs.cn" |
|     `enum_value` | `enum_value` | No | Enum field values |
|       `enum_ids` | `string[]` | Yes | Enum option id **Example value**: ["1"] **Data validation rules**: - Length range: `0` ～ `100` |
|       `enum_type` | `string` | Yes | Enum type **Example value**: "Type of Enum" **Optional values are**:  - `1`: Text - `2`: picture  |
|     `user_values` | `user_value[]` | No | User Field Values **Data validation rules**: - Length range: `0` ～ `100` |
|       `ids` | `string[]` | Yes | Employee IDs, consistent with the employee_id_type type **Example value**: ["1"] **Data validation rules**: - Length range: `0` ～ `100` |
|     `phone_value` | `phone_value` | No |  |
|       `phone_number` | `string` | Yes | **Example value**: "18812345678" |
|       `extension_number` | `string` | No | **Example value**: "234234234 Length range: 0-99 characters" |
|     `field_key` | `string` | No | Custom field key **Example value**: "C-1000001" | ### Request body example

{"department":{"custom_department_id":"eersdf",
"name":{"default_value":"张三

**Data Validation Rules**:

Length range: 1-64 characters",
"i18n_value":{
                    "zh_cn": "中文",
                    "ja_jp": "ja_jp_name",
                    "en_us": "en_us_name"
                }},
"parent_department_id":"h121900",
"leaders":[{
    "leader_type": 1,
    "leader_id": "u273y71"
}],
"order_weight":"100",
"enabled_status":true,
"custom_field_values":[{"field_type":"1",
"text_value":{"default_value":"张三",
"i18n_value":{
                    "zh_cn": "中文",
                    "ja_jp": "ja_jp_name",
                    "en_us": "en_us_name"
                }},
"url_value":{"link_text":{"default_value":"张三",
"i18n_value":{"zh_cn":"张三"}},
"url":"https://m.bytedance.com/afnasjfna",
"pcurl":"http://www.fs.cn"},
"enum_value":{"enum_ids":["1"],
"enum_type":"Type of Enum"},
"user_values":[{
    "ids": [
        "1"
    ]
}],
"phone_value":{"phone_number":"18812345678",
"extension_number":"234234234

Length range: 0-99 characters"},
"field_key":"C-1000001"}]}}

## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `department_id` | `string` | Department ID | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
        "department_id": "eesdasd"
    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 2221305 | Request parameter error | Request parameter error, refer to the error message for details. |
| 400 | 2221306 | ExternalID repeat | Custom ID within the tenant is duplicated. Please change the custom ID and try again. |
| 400 | 2221307 | Exceed department limit | The number of departments exceeds the limit. Please reduce the number of departments to within the limit. |
| 400 | 2221309 | Department does not exist | The parent department does not exist. Please confirm whether the parent department ID is correct. |
| 400 | 2221311 | Internationalized name duplication | The English name of the department already exists, please use another name or the Japanese name of the department already exists, please use another name |
| 400 | 2221312 | Exceed department level depth | The department level is too deep. The upper limit of the department level is 25 |
| 400 | 2221313 | ExternalID invalid | The custom ID length cannot exceed 64 characters |
| 400 | 2221317 | Department direct sub departments exceed limits | The maximum number of sub - departments directly under a department is 1000 |
| 400 | 2221319 | Department name duplicate when creating and updating | When creating or updating, the department name is duplicated. XX (department name) already exists. Please use another name |
| 400 | 2221300 | Department common error | Service exception, please try again |
| 400 | 2221328 | Description invalid | The default value or multilingual value exceeds 100 characters, please check and modify. |
| 400 | 2221331 | Custom field value invalid | The value of the custom field is invalid, please refer to the error message for details. |
| 400 | 2221333 | Department name can not contain '/' | Department names cannot contain "/" |
| 400 | 2221334 | Department ID duplicate | The department ID has been used. Please change the department ID and try again. |
| 400 | 2224001 | No permission to operate | No interface permission. Please apply for interface permission for the application. For specific operations, refer to the relevant documentation. |
| 400 | 2224002 | No permission to operate record | You do not have row record permissions. Please contact the administrator to obtain row record permissions. |
| 400 | 2224003 | No permission to operate dependent object | You do not have permission to operate on related entities, such as parent departments and department leaders. Please confirm whether you have permission to operate on related entities. |
| 400 | 2221349 | Department has member, can not disable department | Department has member, can not disable department |
| 400 | 2221350 | Department has child, can not disable department | Department has child, can not disable department |
| 400 | 2221351 | Department parent is disabled, can not enable | Department parent is disabled, can not enable |
| 400 | 2221352 | Department parent is disabled | The parent department is in a disabled state. Please enable the parent department first. |
