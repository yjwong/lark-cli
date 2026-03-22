---
title: "Create employee"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/directory-v1/employee/create"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/directory/v1/employees"
service: "directory-v1"
resource: "employee"
section: "Organization"
rate_limit: "5 per second"
scopes:
  - "directory:employee.create:write"
  - "directory:employee:write"
field_scopes:
  - "directory:employee.base.external_id:read"
updated: "1756439245000"
---

# Create Employee

This API is used to create an employee under the enterprise. It supports passing in information such as name and mobile phone number to generate an employee object in the active state. 
Employees refer to members with the identity of "Employee" in the Lark enterprise, which is equivalent to the "User" in the address book OpenAPI. 

> Attention:
> - Employees can only be created under the departments within the address book authorization scope of the current application. To create employees under the root department, you must have full permissions. You can view the address book authorization scope in the developer background - application details - permission management.
> - The user_access_token supported in this interface is the administrator user by default, and the administrator management scope will be verified. When a user has multiple administrator identities and can create employees, the administrator management scope will be the maximum set. For administrator permissions, please refer to the Help Center documentation: [Administrator creates administrator role and assigns permissions](https://www.larksuite.com/hc/zh-CN/articles/360043595213)
> - After having the permission for this interface, you can write employee information. However, only the field data that the application has permission to access will be returned after creating an employee. If you need specific fields, please apply for the corresponding permissions as described in the documentation.
> - This interface is only available for Custom Apps.
> - The created employee is in the active state.
> After creating an employee, an invitation SMS/email will be sent. The invitee must click "Agree" to join the enterprise.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/directory/v1/employees |
| HTTP Method | POST |
| Rate Limit | 5 per second |
| Supported app types | custom |
| Required scopes Enable any scope from the list | `directory:employee.create:write` `directory:employee:write` |
| Required field scopes | > The response body of the API contains the following sensitive fields, and they will be returned only after corresponding scopes are added. If you do not need the fields, it is not recommended that you request the scopes. `directory:employee.base.external_id:read` | ### Request header
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
| `employee` | `create_employee` | Yes | Create employee object |
|   `name` | `upsert_name` | No | name |
|     `name` | `i18n_text` | Yes | The name of the employee, up to 64 characters can be entered |
|       `default_value` | `string` | Yes | default value Minimum length: 1 character **Example value**: "Zhang San" |
|       `i18n_value` | `map<string, string>` | No | Internationalized value, key is zh_cn, ja_jp, en_us, value is the corresponding value. **Example value**: {"zh_cn": "Zhang San"} |
|     `another_name` | `string` | No | Alias, up to 64 characters **Example value**: "Jack" |
|   `mobile` | `string` | No | Employee's mobile phone number, enter up to 255 characters. Note: 1. It cannot be repeated among current employees within the enterprise. 2. Uncertified enterprises only support adding Chinese mainland mobile phone numbers, and enterprises certified by Lark are allowed to add overseas mobile phone numbers. 3. The plus sign + must be included in the international area code prefix. **Example value**: "13011111111 "or" + 8613011111111" |
|   `custom_employee_id` | `string` | No | Unique identification of active employees in the enterprise. Customization is supported, and the system automatically generates it when not customized. ID supports modification. Note: 1. The ID of current employees cannot be repeated 2. ID cannot contain spaces **Example value**: "U273y71 **Data verification rules:** Length range: 1-64 characters" |
|   `avatar_key` | `string` | No | Employee's avatar key. To get the key of the picture, please use [Upload Image - Server Level API - Development Documentation - Lark open platform] (/ssl: ttdoc/uAjLw4CM/ukTMukTMukTM/reference/im-v1/image/create). when uploading, you need to select for setting avatar as the picture type **Example value**: "8abc397a-9950-44ea-9302-e1d8fe00858g" **Data validation rules**: - Length range: `0` ～ `255` characters |
|   `email` | `string` | No | The employee's email address at work. Note: 1. It cannot be repeated among current employees within the enterprise. 2. Non-Chinese mainland mobile phone number members must add email at the same time. **Example value**: "zhangsan@gmail.com" **Data validation rules**: - Length range: `0` ～ `255` characters |
|   `enterprise_email` | `string` | No | Employee's enterprise mailbox. Please make sure that the Lark mailbox service has been enabled in the management background first. The domain name of the enterprise mailbox needs to be applied and opened by the enterprise in the management background. If the enterprise does not open the enterprise mailbox of the corresponding domain name, setting the user's enterprise mailbox will fail. **Example value**: "zhangsan@gmail.com" **Data validation rules**: - Length range: `0` ～ `255` characters |
|   `gender` | `int` | No | gender **Example value**: 1 **Optional values are**:  - `0`: unknown - `1`: male - `2`: female - `3`: other  |
|   `employee_order_in_departments` | `upsert_user_department_sort_info[]` | No | The ranking information of employees within their department. **Data validation rules**: - Length range: `0` ～ `10` |
|     `department_id` | `string` | No | Specify the department where the employee is located, identify a unique department within the enterprise, and keep it consistent with the department_id_type type. **Example value**: "Eeddjisdwe" |
|     `order_weight_in_deparment` | `string` | No | The ranking weight of employees within the department. **Example value**: "100" |
|     `order_weight_among_deparments` | `string` | No | The ranking weight of this department among the multiple departments to which the user belongs. **Example value**: "20" |
|     `is_main_department` | `boolean` | No | Whether it is the user's main department (a user can only have one main department, and the sorting weight should be the largest. If not filled in, the first department in the sorting will be used as the main department by default), optional values: true/false. **Example value**: true |
|   `leader_id` | `string` | No | The ID of the employee's line manager, consistent with the employee_id_type. Note: 1. There should be no circular reporting, i.e. A reports to B and B reports to A. 2. The line manager must be an active employee. **Example value**: "Eeasdqwwe" |
|   `dotted_line_leader_ids` | `string[]` | No | The dotted-line manager ID of the employee, which is consistent with the employee_id_type. Note: 1. There should be no circular reporting, i.e. A reports to B and B reports to A. 2. The line manager must be an active employee. **Example value**: ["HDSUQW"] **Data validation rules**: - Length range: `0` ～ `20` |
|   `work_country_or_region` | `string` | No | Workplace country code. To obtain the country code, please use [Paging Bulk Query Country] (/ssl: ttdoc/uAjLw4CM/ukTMukTMukTM/mdm-v3/country_region/list). **Example value**: "MDM34234234" |
|   `work_place_id` | `string` | No | Workplace ID **Example value**: "Eqwedas" |
|   `work_station` | `i18n_text` | No | Workstation |
|     `default_value` | `string` | Yes | default value **Example value**: "Zhang San" |
|     `i18n_value` | `map<string, string>` | No | Internationalized value, key is zh_cn, ja_jp, en_us, value is the corresponding value **Example value**: {"zh_cn": "Zhang San"} |
|   `job_number` | `string` | No | Job number. The job number of employees in the enterprise cannot be repeated. **Example value**: "2845435 **Data verification rules:** Length range: 0-255 characters" |
|   `extension_number` | `string` | No | Extension number, up to 99 characters can be entered . The extension numbers of all employees in the enterprise cannot be repeated. **Example value**: "2845435" |
|   `join_date` | `string` | No | Onboard date **Example value**: "2022-10-10 **Data verification rules:** Length range: Fixed length: 10 characters, fixed format: “yyyy-mm-dd”" |
|   `employment_type` | `int` | No | Employee type **Example value**: 1 **Optional values are**:  - `1`: full-time - `2`: internship - `3`: outsourcing - `4`: labor service - `5`: consultant  |
|   `job_title_id` | `string` | No | Job ID **Example value**: "Wqedsaqw" |
|   `custom_field_values` | `custom_field_value[]` | No | custom field **Data validation rules**: - Length range: `0` ～ `100` |
|     `field_type` | `string` | No | Custom field type **Example value**: "1" **Optional values are**:  - `1`: Text - `2`: URL - `3`: Options - `4`: Members - `9`: - `10`: Multiple selection options - `11`: Members list  |
|     `text_value` | `i18n_text` | No | Text field value |
|       `default_value` | `string` | Yes | default Minimum length: 1 character **Example value**: "Inspire creativity" |
|       `i18n_value` | `map<string, string>` | No | Internationalized value, key is zh_cn, ja_jp, en_us, value is the corresponding value **Example value**: {"zh_cn": "激发创造","en_us":"Inspire creativity"} |
|     `url_value` | `url_value` | No |  |
|       `link_text` | `i18n_text` | Yes |  |
|         `default_value` | `string` | Yes | default value Length range: 1-40 characters **Example value**: "Inspire creativity" |
|         `i18n_value` | `map<string, string>` | No | Internationalized value, key is zh_cn, ja_jp, en_us, value is the corresponding value **Example value**: {"zh_cn": "激发创造","en_us":"Inspire creativity"} |
|       `url` | `string` | Yes | **Example value**: "https://m.bytedance.com/afnasjfna" |
|       `pcurl` | `string` | Yes | **Example value**: "http://www.fs.cn" |
|     `enum_value` | `enum_value` | No | Url field value |
|       `enum_ids` | `string[]` | Yes | link title **Example value**: ["1"] **Data validation rules**: - Length range: `0` ～ `100` |
|       `enum_type` | `string` | Yes | mobile web link **Example value**: "1" **Optional values are**:  - `1`: text - `2`: image  |
|     `user_values` | `user_value[]` | No | enum field value **Data validation rules**: - Length range: `0` ～ `100` |
|       `ids` | `string[]` | Yes | employee id，consistent with the employee_id_type type. **Example value**: ["1"] **Data validation rules**: - Length range: `0` ～ `100` |
|     `phone_value` | `phone_value` | No | field representing employee |
|       `phone_number` | `string` | Yes | **Example value**: "18812345678" |
|       `extension_number` | `string` | No | **Example value**: "234234234" |
|     `field_key` | `string` | No | Custom field key **Example value**: "C-1000001" |
| `options` | `create_employee_options` | No | api extend options |
|   `geo_name` | `string` | No | Employee's Data Residency. Optional only for organizations with Multi-Geo enabled, and can only be populated for Geo's in the list of organization data residences. The following permissions need to be requested to write:Write to Employee Data Location **Example value**: "cn" |
|   `subscription_ids` | `string[]` | No | A list of seat IDs assigned to employees. The available seat IDs for this tenant can be obtained through the interface below, see Get Seat Information. This field is required when in hybrid license mode.Permissions are required in order to write: `directory:employee.base.subscription_ids:write` **Example value**: ["123123123"] **Data validation rules**: - Length range: `0` ～ `20` | ### Request body example

{
    "employee":
    {
        "name":
        {
            "name":
            {
                "default_value": "张三",
                "i18n_value":
                {
                    "zh_cn": "中文",
                    "ja_jp": "ja_jp_name",
                    "en_us": "en_us_name"
                }
            },
            "another_name": "jack"
        },
        "mobile": "13011111111 或 +8613011111111",
        "custom_employee_id": "u273y71",
        "avatar_key": "8abc397a-9950-44ea-9302-e1d8fe00858g",
        "email": "zhangsan@gmail.com",
        "enterprise_email": "zhangsan@gmail.com",
        "gender": 1,
        "employee_order_in_departments":
        [
            {
                "department_id": "",
                "order_weight_in_deparment": "100",
                "order_weight_among_deparments": "20",
                "is_main_department": false
            }
        ],
        "leader_id": "eeasdqwwe",
        "dotted_line_leader_ids":
        [
            "hdsuqw"
        ],
        "work_country_or_region": "MDM34234234",
        "work_place_id": "eqwedas",
        "work_station":
        {
            "default_value": "张三",
            "i18n_value":
            {
                "zh_cn": "中文",
                "ja_jp": "ja_jp_name",
                "en_us": "en_us_name"
            }
        },
        "job_number": "2845435",
        "extension_number": "2845435",
        "join_date": "2022-10-10",
        "employment_type": "",
        "staff_status": "",
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
        "job_title_id": "wqedsaqw",
        "probation_period": "10",
        "convert_status": "",
        "resign_reason": "",
        "resign_type": "",
        "cancelled_entry_type": "",
        "custom_field_values":
        [
            {
                "type": "",
                "text_value":
                {
                    "default_value": "张三",
                    "i18n_value":
                    {
                        "zh_cn": "中文",
                        "ja_jp": "ja_jp_name",
                        "en_us": "en_us_name"
                    }
                },
                "url_value":
                {
                    "link_text":
                    {
                        "default_value": "张三",
                        "i18n_value":
                        {
                            "zh_cn": "中文",
                            "ja_jp": "ja_jp_name",
                            "en_us": "en_us_name"
                        }
                    },
                    "url": "https://m.bytedance.com/afnasjfna",
                    "pcurl": "http://www.fs.cn"
                },
                "enum_value":
                {
                    "ids":
                    [
                        "1"
                    ],
                    "name":
                    {
                        "zh_cn": "中文",
                        "ja_jp": "ja_jp_name",
                        "en_us": "en_us_name"
                    },
                    "enum_type": ""
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
                ],
                "field_key": "C-1000001"
            }
        ]
    },
    "options":
    {
        "geo_name": "cn",
        "subscription_ids":
        [
            "123123123"
        ],
        "need_send_notification":
        {
            "need_send_notification": false,
            "language": 1
        }
    }
}

## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `employee_id` | `string` | Employee IDWhen the employee_id_type value is employee_id, the field permission requirements: `directory:employee.base.external_id:read` | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
        "employee_id": "sderdt"
    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 2221164 | User name exceeds limit | The name is too long. The maximum length is 64 characters. |
| 400 | 2221165 | User en_name exceeds limit | The English name exceeds the limit. You can enter up to 64 characters. |
| 400 | 2221166 | User another_name exceeds limit | The alias length exceeds the limit. You can enter up to 64 characters |
| 400 | 2221103 | Up to 64 characters can be entered | The mobile phone number already exists, please modify the mobile phone number. |
| 400 | 2221106 | Invalid mobile | Invalid mobile phone number, please modify the mobile phone number. |
| 400 | 2221113 | Mobile or email not set | Mobile phone number or email address Required |
| 400 | 2221114 | Users must have a mobile in China | China must have a mobile phone number |
| 400 | 2221104 | Email already exists | The mailbox already exists, request to modify the mailbox. |
| 400 | 2221107 | Invalid email | Invalid mailbox, request to modify the enterprise mailbox. |
| 400 | 2221118 | Enterprise email already exists | The enterprise mailbox already exists, and the enterprise mailbox is requested to be modified. |
| 400 | 2221278 | Invalid enterprise email | Invalid enterprise mailbox, request to modify enterprise mailbox |
| 400 | 2221126 | Enterprise email domain unavailable | Enterprise mailbox domain name is not available, request to modify enterprise mailbox |
| 400 | 2221146 | Enterprise email alias exceeds limit | The enterprise email alias exceeds the length limit. You can enter up to 255 characters. |
| 400 | 2221147 | Enterprise email address in recycling bin | The enterprise email address is in Recycle Bin, please modify the enterprise email address. |
| 400 | 2221176 | Add Lark allows listing tenants. Email must be included with non + 86mobile | Lark tenants must include email information when adding a mobile phone number other than + 86. |
| 400 | 2221255 | Main department must be the first | The main department must be the first in the sorting information of the employee's department. Please modify the sorting information of the employee's department. |
| 400 | 2221125 | The number of members within the department exceeds the limit. Please contact an administrator for help | The number of members in the department cannot exceed 10,000. Please contact the administrator for help. Access the Admin Console (larksuite.com/admin) on the desktop, and the administrator information will be displayed on the guidance page. |
| 400 | 2221129 | User department is empty | The employee's department is empty, please modify the sorting information in the employee's department |
| 400 | 2221141 | Unable to join multiple departments. Please upgrade relevant'Organizational Structure Visible '. | Unable to join multiple departments, please modify the sorting information within the department to which the employee belongs |
| 400 | 2221181 | Department does not exist | The department does not exist, please modify the sorting information in the department to which the employee belongs |
| 400 | 2221239 | Leader loop error | Direct manager, please modify direct manager |
| 400 | 2221238 | DottedLineLeaderID loop error | Dotted-line manager is a ring, please modify dotted-line manager |
| 400 | 2221221 | DottedLineLeaderID exceeds length limit | The length of the dotted-line manager exceeds the limit, please modify the dotted-line manager. |
| 400 | 2221222 | Invalid dottedLineLeaderID | Invalid dotted-line manager ID, please modify dotted-line manager |
| 400 | 2221242 | Invalid custom field | Invalid custom field, please modify the custom field |
| 400 | 2221216 | Invalid working country or region | Invalid work country or region, please modify work country or region |
| 400 | 2221217 | WorkplaceID not found | The working city does not exist, please modify the working city |
| 400 | 2221240 | JobNumber not unique | The job number already exists, please modify the job number. |
| 400 | 2221191 | Invalid extension number | The extension number is invalid, please modify the extension number |
| 400 | 2221192 | Repeated extension number within the tenant | The extension number already exists, please modify the extension number |
| 400 | 2221193 | Up to 99 characters can be entered | The extension number length exceeds the limit. You can enter up to 99 characters |
| 400 | 2221210 | Invalid join date | Invalid entry time, please modify the entry time |
| 400 | 2221144 | EmployeeType not found | The person type does not exist, please modify the person type |
| 400 | 2221145 | EmployeeType inactive | The person type is not activated, please modify the person type |
| 400 | 2221223 | Invalid job title ID | The position does not exist, please change the position |
| 400 | 2221263 | Tenant has not activated multi geo | The tenant has not turned on Multi-Geo, please open it first and try again. **Multi-Geo** refers to multi-geo data residency. |
| 400 | 2221264 | User geo does not exist | The specified Geo does not exist. Please check if the Geo parameter is correct. |
| 400 | 2221265 | The application does not have permission to write to the geo | No permission to specify employee Geo, need to apply ,Click the API Debugging Console - Permission Configuration, and the required permissions will be displayed. Click "Operation" - "..." - "Activate".  where employee data is written  |
| 400 | 2221266 | The application does not have permission to write to the SubscriptionID | No permission to specify seats, you need to apply ,Click the API Debugging Console - Permission Configuration, and the required permissions will be displayed. Click "Operation" - "..." - "Activate".  Write employee seat information  |
| 400 | 2224001 | No permission to operate | No operation permission. Please check the permission of the current application or whether the enterprise version is the business professional version or above. |
| 400 | 2224002 | No permission to operate records | You do not have permission to operate this record. Please check the permissions of the data management scope of the current application or whether the members operated by the current application can be created. |
| 400 | 2224003 | No permission to operate dependent objects | You do not have permission to operate the dependent object. Please check whether you have permission to create the department. |
| 400 | 2221252 | Hybrid license tenant prohibits passing empty licenses to create users | Mixed license tenant prohibits passing empty licenses to create users |
| 400 | 2221253 | Designated licenseID is insufficient | The remaining seats are insufficient, please modify the seat information |
| 400 | 2221254 | Designated licenseID is invalid | The specified subscription_ids is illegal. Please check the seat ID. |
| 400 | 2221111 | Exceeds certified seat limit | If the seat limit is exceeded, please modify the seat information |
| 400 | 2221112 | Exceeds billing plan seat limit | The employee data residence does not exist, please modify the employee data residence |
| 400 | 2221115 | ExternalID is not unique | The custom ID already exists, please modify the custom ID. |
| 400 | 2221116 | Invalid ExternalID | If the seat limit is exceeded, please modify the seat information |
| 400 | 2221175 | Lark only supports + 86Mobile | Lark tenants only support + 86 mobile phone number, please modify the mobile phone number. |
| 400 | 2221163 | Users are created too frequently | Employees create too often, please try again later |
| 400 | 2221109 | Name contains sensitive info | The name contains sensitive information, please modify the name. |
| 400 | 2221292 | User department is disabled | The user department has been disabled. Please contact the administrator to enable the department or modify the department to which the employee belongs. |
