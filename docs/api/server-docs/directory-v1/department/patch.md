---
title: "Update department"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/directory-v1/department/patch"
method: "PATCH"
api_path: "https://open.larksuite.com/open-apis/directory/v1/departments/:department_id"
service: "directory-v1"
resource: "department"
section: "Organization"
rate_limit: "10 per second"
scopes:
  - "directory:department.update:write"
  - "directory:department:write"
field_scopes:
  - "directory:employee.base.external_id:read"
updated: "1756439290000"
---

# Update department

This interface is used to update department information. Only the parts param that are explicitly passed are updated.

> - Department information can only be updated within the scope of the address book authorization of the current application (which can be viewed via the management backend > application permissions > address book authorization page).
> - To change the parent department of a department, you need to have the department permissions both before and after the change to succeed.
> - To change the head of a department, you need to have the personnel permissions before and after the change to ensure a successful change.
> - The user_access_token supported in this interface defaults to an administrator user, and the administrator's management scope will be verified. When a user has multiple administrator identities that can all update departments, the administrator's management scope will be the maximum set. For administrator permissions, please refer to the Help Center document: [Administrator creates administrator role and assigns permissions](https://www.larksuite.com/hc/zh-CN/articles/360043595213)
> 
> - Modifying the department ID (department_id) requires knowledge of the following impacts:
> - The department ID (department_id) is the unique identifier of a department within an enterprise, which may be referenced by applications to implement various internal logics. Modifying the unique ID may lead to reference failures, causing all business operations that reference and have saved the 'modified ID department' to be affected.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/directory/v1/departments/:department_id |
| HTTP Method | PATCH |
| Rate Limit | 10 per second |
| Supported app types | custom |
| Required scopes Enable any scope from the list | `directory:department.update:write` `directory:department:write` |
| Required field scopes | > The response body of the API contains the following sensitive fields, and they will be returned only after corresponding scopes are added. If you do not need the fields, it is not recommended that you request the scopes. `directory:employee.base.external_id:read` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `department_id` | `string` | Department ID, consistent with the department_id_type type **Example value**: "H12921" **Data validation rules**: - Maximum length: `64` characters | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `employee_id_type` | `string` | No | User ID categories **Example value**: open_id **Optional values are**:  - `open_id`: Identifies a user to an app. The same user has different Open IDs in different apps. How to get Open ID - `union_id`: Identifies a user to a tenant that acts as a developer. A user has the same Union ID in apps developed by the same developer, and has different Union IDs in apps developed by different developers. A developer can use Union ID to link the same user's identities in multiple apps.How to get Union ID - `employee_id`: The unique identifier of an employee within the enterprise. Supports customization, the system automatically generates it if it is not customized. ID supports modification. How to get employee_id: - Enterprise administrators go to the Admin Console > Organizational Structure > Members and Departments page, click Member Details to query employee IDs - Through the interface of Get employee list in batches, query employee ID by mobile phone number or email.  **Default value**: `open_id` **When the value is `employee_id`, the following field scopes are required**: `directory:employee.base.external_id:read` |
| `department_id_type` | `string` | No | The type of department ID used in this call **Example value**: open_department_id **Optional values are**:  - `open_department_id`: It is used to identify a department in a specific application, and the same department has different open_department_id in different applications. - `department_id`: Used to identify a unique department within a tenant  **Default value**: `open_department_id` | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `department` | `update_department` | Yes | Update department information |
|   `custom_department_id` | `string` | No | Customize the department ID. Note: 1. Except for the need to meet the regular rules, it cannot start with od- at the same time 2. Regular check: ^ [a-zA-Z0-9] [a-zA-Z0-9 _\-@.]{ 0,63} $ **Data Validation Rules**: Length range: 1-64 characters **Example value**: "eedasqwA" |
|   `name` | `i18n_text` | No | Department name |
|     `default_value` | `string` | Yes | default value **Example value**: "Zhang San Length range: 1-100" |
|     `i18n_value` | `map<string, string>` | No | Internationalized value, key is zh_cn, ja_jp, en_us, value is the corresponding value **Example value**: {"zh_cn": "Zhang San"} |
|   `parent_department_id` | `string` | No | Parent Department ID, consistent with the department_id_type type **Example value**: "100" |
|   `leaders` | `department_leader[]` | No | department head **Data validation rules**: - Length range: `0` ～ `20` |
|     `leader_type` | `int` | Yes | Type of department head **Example value**: 1 **Optional values are**:  - `1`: main - `2`: vice  |
|     `leader_id` | `string` | Yes | Department head ID, consistent with the employee_id_type type **Example value**: "u273y71" |
|   `order_weight` | `string` | No | The sorting weight under the superior department, and the returned results will be sorted in ascending order according to the value of order_weight. **Example value**: "100" |
|   `enabled_status` | `boolean` | No | Whether to enable **Example value**: true |
|   `custom_field_values` | `custom_field_value[]` | No | Department custom field values **Data validation rules**: - Length range: `0` ～ `100` |
|     `field_key` | `string` | No | Custom field key **Example value**: "C-1000001" |
|     `field_type` | `string` | No | custom field type **Example value**: "1" **Optional values are**:  - `1`: Text - `2`: URL - `3`: Options - `4`: Members - `10`: Multiple selection options - `11`: Members list  |
|     `text_value` | `i18n_text` | No | Text field value |
|       `default_value` | `string` | Yes | default value **Example value**: "Zhang San" |
|       `i18n_value` | `map<string, string>` | No | Internationalized value, key is zh_cn, ja_jp, en_us, value is the corresponding value **Example value**: {"zh_cn": "Chinese", "ja_jp": "ja_jp_name", "en_us": "en_us_name"} |
|     `url_value` | `url_value` | No | Web link field value |
|       `link_text` | `i18n_text` | Yes | page title |
|         `default_value` | `string` | Yes | default value **Example value**: "Zhang San" |
|         `i18n_value` | `map<string, string>` | No | Internationalized value, key is zh_cn, ja_jp, en_us, value is the corresponding value **Example value**: {"zh_cn": "Chinese", "ja_jp": "ja_jp_name", "en_us": "en_us_name"} |
|       `url` | `string` | Yes | Mobile end web link **Example value**: "https://m.bytedance.com/afnasjfna" |
|       `pcurl` | `string` | Yes | Desktop web link **Example value**: "http://www.fs.cn" |
|     `enum_value` | `enum_value` | No | Enumerate field values |
|       `enum_ids` | `string[]` | Yes | Option result ID **Example value**: ["1"] **Data validation rules**: - Length range: `0` ～ `100` |
|       `enum_type` | `string` | Yes | option type **Example value**: "1" **Optional values are**:  - `1`: Text - `2`: picture  |
|     `user_values` | `user_value[]` | No | Person field value **Data validation rules**: - Length range: `0` ～ `100` |
|       `ids` | `string[]` | Yes | Person ID,consistent with the employee_id_type type **Example value**: ["1"] **Data validation rules**: - Length range: `0` ～ `100` |
|       `user_type` | `string` | Yes | Personnel Type **Example value**: "1" **Optional values are**:  - `1`: employees  | ### Request body example

{"department":{"custom_department_id":"eedasqwA",
"name":{"default_value":"Zhang San

Length range: 1-100",
"i18n_value":{"zh_cn": "Zhang San"}},
"parent_department_id":"100",
"leaders":[{
    "leader_type": 1,
    "leader_id": "u273y71"
}],
"order_weight":"100",
"enabled_status":true,
"custom_field_values":[{
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
            ],
            "user_type": "1"
        }
    ]
}]}}

## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- | ### Response body example

{
    "code": 0,
    "msg": "Success",
    "data": {}
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 2221305 | Request parameter error | Request parameter error, refer to the error message for details. |
| 400 | 2221306 | ExternalID repeat | Custom ID within the tenant is duplicated. Please change the custom ID and try again. |
| 400 | 2221307 | Exceed department limit | The number of departments exceeds the limit. Please reduce the number of departments to within the limit and try again. |
| 400 | 2221309 | Department does not exist | The updated department does not exist, and the parent department does not exist. |
| 400 | 2221311 | Internationalized name duplication | The multilingual name conflicts with an existing department. Please change the multilingual name to an unused one and try again. |
| 400 | 2221312 | The upper limit of department level is 25 | The department level is too deep. The maximum limit for the department level is 25. Please adjust the department level to within 25 levels. |
| 400 | 2221313 | ExternalID invalid | The length of the custom ID cannot exceed 64 characters. Please adjust the length of the custom ID to within 64 characters. |
| 400 | 2221317 | Department directing sub departments exceeding limits | The maximum number of directly subordinate sub-departments of a department is 1000. Please adjust the number of directly subordinate sub-departments of the department to within 1000. |
| 400 | 2221319 | Department name duplicate when creating and updating | When creating or updating, the department name is duplicated. XX (department name) already exists. Please use another name |
| 400 | 2221328 | Description invalid | The default value or multilingual value exceeds 100 characters, please check the default value and multilingual value. |
| 400 | 2221329 | Department type invalid | Invalid department type. Please use a valid department type and try again. |
| 400 | 2221330 | HRBP invalid | The HRBP is an employee who has not left the company and is not frozen from a different tenant. Please ensure that the HRBP is an employee who has not left the company and is not frozen from the same tenant, and then try again. |
| 400 | 2221331 | Custom field value invalid | The custom field value is invalid, refer to the error message for details. |
| 400 | 2221332 | Department has members, can not delete department | There are still active and pending employees under this department, so the department cannot be deleted. Please remove the active and pending employees under the department first, and then delete the department. |
| 400 | 2221333 | Department name can not contain '/' | Department names cannot contain "/" |
| 400 | 2221334 | Department ID duplicate | The department ID has been used. Please change the department ID and try again. |
| 400 | 2224001 | No permission to operate | No interface permission. Please apply for interface permission for the application and try again. For specific operations, refer to the relevant documentation. |
| 400 | 2224002 | No permission to operate records | You do not have row record permission. Please contact the administrator to grant row record permission and then try again. |
| 400 | 2224003 | No permission to operate dependent objects | There is no permission for the dependent object. Please apply for permission for the dependent object for the application and then try again. For specific operations, refer to the relevant documentation. |
| 400 | 2221349 | Department has members, can not disable department | There are still members under the department and it cannot be deactivated. Please remove the members under the department first, and then deactivate the department. |
| 400 | 2221350 | Department has children, can not disabled department | The department has sub-departments under it and cannot be deactivated. Please remove the sub-departments under the department first, and then deactivate the department. |
| 400 | 2221351 | Department parent is disabled, can not enable | The parent department of this department has been deactivated. It cannot be activated. Please activate the parent department first, and then activate this department. |
| 400 | 2221352 | Department parent is disabled | The parent department of this department is disabled, and related operations cannot be performed. Please enable the parent department and try again. |
