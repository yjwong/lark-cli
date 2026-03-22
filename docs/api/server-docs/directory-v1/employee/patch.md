---
title: "Update employee information"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/directory-v1/employee/patch"
method: "PATCH"
api_path: "https://open.larksuite.com/open-apis/directory/v1/employees/:employee_id"
service: "directory-v1"
resource: "employee"
section: "Organization"
rate_limit: "10 per second"
scopes:
  - "directory:employee.update:write"
  - "directory:employee:write"
field_scopes:
  - "directory:employee.base.external_id:read"
updated: "1756439249000"
---

# Update employee information 

This interface is used to update the information of current/former employees, freeze/restore employees. Unpassed parameters will not be updated.
Employee refers to a member of Lark's enterprise who is identified as "Employee", which is equivalent to "User" in the address book OpenAPI.

> - The modification of employee status follows the rules of life cycle flow. For details, please refer to Directory-Employee Management-Resource Introduction.
> - This interface supports tenant_access_token and user_access_token.Refer to Get Access Credentials for the interface acquisition method.
> - When using tenant_access_token, employee information can only be updated within the scope of the current app's address book authorization.can be viewed in Developer Backend > Permission Management > Address Book Permission.
> - When changing an employee's department information, the application needs to have department permissions before and after the change in order for the change to be successful.
> - When using user_access_token, the default is the administrator user, and the administrator management scope will be verified. When the user has multiple administrator identities to update employee information, the administrator management scope takes the maximum set.Administrator privileges can view the Help Center document: [Administrator creates administrator role and assigns permissions](https://www.larksuite.com/hc/zh-CN/articles/360043595213)
> - Changing the contact mobile phone number and work email of employees in the "unjoined" and "inactive" states will modify the employee's login credentials, reset the employee to the "unjoined" state, and send an invitation text message/email. Modifying the contact information of employees in other states does not affect the login credentials.
> - Changing employee ID (employee_id) requires knowledge of the following effects:
> - Employee ID (employee_id) is the employee's unique ID in the enterprise, which may be referenced by applications to implement various internal logic. After the unique ID is modified, the reference may fail, resulting in all referenced and saved'modified ID employee 'businesses being affected.
> - When updating employee information with exit status, the following fields cannot be updated:
> - Email, mobile, department_ids, leader_id, is_frozen, work_city_id

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/directory/v1/employees/:employee_id |
| HTTP Method | PATCH |
| Rate Limit | 10 per second |
| Supported app types | custom |
| Required scopes Enable any scope from the list | `directory:employee.update:write` `directory:employee:write` |
| Required field scopes | > The response body of the API contains the following sensitive fields, and they will be returned only after corresponding scopes are added. If you do not need the fields, it is not recommended that you request the scopes. `directory:employee.base.external_id:read` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `employee_id` | `string` | Employee ID, consistent with the employee_id_type type. **Example value**: "EeehsDNA" | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `employee_id_type` | `string` | No | User ID categories **Example value**: open_id **Optional values are**:  - `open_id`: Identifies the identity of a user in an application. The same user has different Open IDs in different applications. Learn more: How to get an Open ID - `union_id`: Identifies the identity of a user under an application developer. The Union ID of the same user in an application under the same developer is the same, and the Union ID in an application under different developers is different. With Union ID, application developers can associate the identity of the same user in multiple applications. Learn more: How to get Union ID? - `employee_id`: The unique identifier of an employee within the enterprise. Supports customization, the system automatically generates it if it is not customized. ID supports modification. How to get employee_id: - Enterprise administrators go to the Admin Console > Organizational Structure > Members and Departments page, click Member Details to query employee IDs - Through the interface of Get employee list in batches, query employee ID by mobile phone number or email.  **Default value**: `open_id` **When the value is `employee_id`, the following field scopes are required**: `directory:employee.base.external_id:read` |
| `department_id_type` | `string` | No | Department ID type **Example value**: open_department_id **Optional values are**:  - `department_id`: department_id - `open_department_id`: open_department_id  **Default value**: `open_department_id` | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `employee` | `update_employee` | Yes | Update employee object |
|   `name` | `upsert_name` | No | Name |
|     `name` | `i18n_text` | Yes | Employee's name |
|       `default_value` | `string` | Yes | default value Length range: 1 - 64 characters **Example value**: "Zhang San " |
|       `i18n_value` | `map<string, string>` | No | Internationalized value, key is zh_cn, ja_jp, en_us, value is the corresponding value **Example value**: {"zh_cn": "张三"} |
|     `another_name` | `string` | No | Alias, up to 64 characters **Example value**: "Jack" |
|   `mobile` | `string` | No | The mobile phone number of the employee, up to 255 characters can be entered. . Note: 1. It cannot be repeated among current employees in the enterprise 2. Uncertified enterprises only support adding Chinese mainland mobile phone numbers, and enterprises certified by Lark are allowed to add overseas mobile phone numbers 3. The international area code prefix must contain the plus sign + **Example value**: "13011111111 "or" + 8613011111111" |
|   `custom_employee_id` | `string` | No | Unique identification of active employees in the enterprise. Customization is supported, and the system automatically generates it when not customized. ID supports modification. Note: 1. The ID of current employees cannot be duplicated. 2. ID cannot contain spaces. **Example value**: "Eesadeq" |
|   `avatar_key` | `string` | No | Employee's avatar key. To get the key of the picture, please use Upload Image - Server Level API - Development Documentation - Lark open platform. when uploading, you need to select for setting avatar as the picture type **Example value**: "dadwqeqwdsa" |
|   `email` | `string` | No | The employee's email address at work. Note: 1. It cannot be repeated among current employees within the enterprise. 2. Non-Chinese mainland mobile phone number members must add email at the same time. **Example value**: "zhangsan@163.com" |
|   `enterprise_email` | `string` | No | Employee's enterprise mail address. Please make sure that the Lark mailbox service has been enabled in the management background first. The domain name of the enterprise mail address needs to be applied and opened by the enterprise in the management background. If the enterprise does not open the enterprise mailbox of the corresponding domain name, setting the user's enterprise mail address will fail. **Example value**: "zhangsan@163.com" |
|   `gender` | `int` | No | gender **Example value**: 1 **Optional values are**:  - `0`: unknown - `1`: male - `2`: female - `3`: other  |
|   `employee_order_in_departments` | `upsert_user_department_sort_info[]` | No | Sorting information of employees within their department **Data validation rules**: - Length range: `0` ～ `10` |
|     `department_id` | `string` | No | Department ID, consistent with the department_id_type type **Example value**: "eediasdjw" |
|     `order_weight_in_deparment` | `string` | No | Sorting weights of employees within the department **Data verification rules:** **Example value**: "100" |
|     `order_weight_among_deparments` | `string` | No | The ranking weight of the department among the multiple departments to which the user belongs **Example value**: "20" |
|     `is_main_department` | `boolean` | No | Whether it is the user's main department (the user can only have one main department, and the ranking weight should be the largest. If you don't fill it in, the department with the first ranking will be used as the main department by default) **Example value**: True |
|   `background_image_key` | `string` | No | The key of the background cover. To get the key of the image, please use Upload Image - Server Level API - Development Documentation - Lark open platform. when uploading, you need to select for sending messages as the picture type **Example value**: "qweasdqawqeq" |
|   `description` | `string` | No | Employee's personal signature **Example value**: "New employee onboarding" |
|   `leader_id` | `string` | No | Employee's direct manager ID. Note: 1. It is impossible to form a ring, that is, the superior of A is B, and the superior of B is A. 2. The superior needs to be a current employee. **Example value**: "eeshfosd" |
|   `dotted_line_leader_ids` | `string[]` | No | The employee's dotted-line manager ID, consistent with the employee_id_type type . Note: 1. It is impossible to form a ring, that is, the superior of A is B, and the superior of B is A. 2. The superior needs to be a current employee. **Example value**: ["eefhdgsd"] **Data validation rules**: - Length range: `0` ～ `10` |
|   `work_country_or_region` | `string` | No | Workplace country code. To obtain the country code, please use [Paging Bulk Query Country] . **Example value**: "MDM2312312" |
|   `work_place_id` | `string` | No | Workplace ID **Example value**: "1111SDDA" |
|   `work_station` | `i18n_text` | No | Workstation |
|     `default_value` | `string` | Yes | default value **Example value**: "Work Station" |
|     `i18n_value` | `map<string, string>` | No | Internationalized value, key is zh_cn, ja_jp, en_us, value is the corresponding value **Example value**: {"zh_cn": "工位"} |
|   `job_number` | `string` | No | Job number. The job number of employees in the enterprise cannot be repeated. **Example value**: "28549233" |
|   `extension_number` | `string` | No | Extension number, up to 99 characters can be entered.. The extension numbers of all employees in the enterprise cannot be repeated. **Example value**: "2854923" |
|   `join_date` | `string` | No | Onboard Date The fixed format is 'YYYY-MM-DD' and the fixed length is 10 **Example value**: "2022-10-10" |
|   `employment_type` | `int` | No | Employee type **Example value**: 1 **Optional values are**:  - `0`: unknown - `1`: full-time - `2`: internship - `3`: outsourcing - `4`: labor service - `5`: consultant  |
|   `job_title_id` | `string` | No | Job title ID **Example value**: "Ewdjdssd" |
|   `job_level_id` | `string` | No | Job level ID **Example value**: "eersdfa" |
|   `job_family_id` | `string` | No | Job family ID **Example value**: "koiuytr" |
|   `resign_date` | `string` | No | date of separation The fixed format is 'YYYY-MM-DD' and the fixed length is 10 **Example value**: "2022-10-10" **Data validation rules**: - Length range: `0` ～ `20` characters |
|   `resign_reason` | `string` | No | reason for leaving **Example value**: "1" **Optional values are**:  - `0`: empty - `1`: Salary does not meet expectations - `2`: Working too long - `3`: Dissatisfied with the job content - `4`: Do not recognize superiors or management - `5`: Career development opportunities are limited - `6`: Lack of recognition of company culture - `7`: Organizational restructuring (voluntary resignation) - `8`: Contract expires - `9`: Job hopping - `10`: career change - `11`: Family reasons - `12`: Poor health - `13`: Workplace reasons - `14`: Other (voluntary resignation) - `15`: accident - `16`: passed away - `17`: dismiss - `18`: Probation period not passed - `19`: Underperforming at work - `20`: Low work output - `21`: Organizational restructuring (passive departure) - `22`: discipline violation - `23`: illegal - `24`: Other (passive resignation) - `25`: Other  |
|   `resign_remark` | `string` | No | Resign remark **Example value**: "个人原因" **Data validation rules**: - Length range: `0` ～ `255` characters |
|   `resign_type` | `string` | No | Resign type **Example value**: "1" **Optional values are**:  - `0`: empty - `1`: active - `2`: inactive - `3`: other  |
|   `is_frozen` | `boolean` | No | Whether to freeze employee accounts. True to freeze, false to restore account. **Example value**: True |
|   `custom_field_values` | `custom_field_value[]` | No | Custom field **Data validation rules**: - Length range: `0` ～ `100` |
|     `field_type` | `string` | No | Custom field type **Example value**: "1" **Optional values are**:  - `1`: Text - `2`: URL - `3`: Options - `4`: Members - `9`: Phone - `10`: Multiple selection options - `11`: Member list  |
|     `text_value` | `i18n_text` | No | Text field value |
|       `default_value` | `string` | Yes | default value **Example value**: "Name Filed" |
|       `i18n_value` | `map<string, string>` | No | Internationalized value, key is zh_cn, ja_jp, en_us, value is the corresponding value **Example value**: {"zh_cn": "姓名字段"} |
|     `url_value` | `url_value` | No | Web link field value |
|       `link_text` | `i18n_text` | Yes | page title |
|         `default_value` | `string` | Yes | default value **Example value**: "URL Title" |
|         `i18n_value` | `map<string, string>` | No | Internationalized value, key is zh_cn, ja_jp, en_us, value is the corresponding value **Example value**: {"zh_cn": "网页标题"} |
|       `url` | `string` | Yes | Mobile end web link **Example value**: "https://m.bytedance.com/afnasjfna" |
|       `pcurl` | `string` | Yes | Desktop web link **Example value**: "http://www.fs.cn" |
|     `enum_value` | `enum_value` | No | enumeration |
|       `enum_ids` | `string[]` | Yes | Option result ID **Example value**: ["1"] **Data validation rules**: - Length range: `0` ～ `100` |
|       `enum_type` | `string` | Yes | option type **Example value**: "1" **Optional values are**:  - `1`: Text - `2`: picture  |
|     `user_values` | `user_value[]` | No | Person field value **Data validation rules**: - Length range: `0` ～ `100` |
|       `ids` | `string[]` | Yes | Person ID, consistent with the employee_id_type type **Example value**: ["1"] **Data validation rules**: - Length range: `0` ～ `100` |
|     `phone_value` | `phone_value` | No |  |
|       `phone_number` | `string` | Yes | **Example value**: "18812345678" |
|       `extension_number` | `string` | No | **Example value**: "234234234" |
|     `field_key` | `string` | No | Custom field key **Example value**: "C-1000001" | ### Request body example

{
    "employee": {
        "name": {
            "name": {
                "default_value": "张三",
                "i18n_value": {
                    "zh_cn": "张三",
                    "ja_jp": "佐藤はるか",
                    "en_us": "Alex Zhang"
                }
            },
            "another_name": "Jack"
        },
        "mobile": "+8613011111111",
        "custom_employee_id": "u273y71",
        "avatar_key": "8abc397a-9950-44ea-9302-e1d8fe00858g",
        "email": "zhangsan@test.com",
        "enterprise_email": "zhangsan@test.com",
        "gender": 1,
        "employee_order_in_departments": [
            {
                "department_id": "0",
                "order_weight_in_deparment": "100",
                "order_weight_among_deparments": "20",
                "is_main_department": false
            }
        ],
        "leader_id": "eeasdqwwe",
        "dotted_line_leader_ids": [
            "hdsuqw"
        ],
        "work_country_or_region": "MDM34234234",
        "work_place_id": "eqwedas",
        "work_station": {
            "default_value": "张三",
            "i18n_value": {
                "zh_cn": "工位",
                "ja_jp": "工位",
                "en_us": "Work Station"
            }
        },
        "job_number": "2845435",
        "extension_number": "2845435",
        "join_date": "2022-10-10",
        "employment_type": "1",
        "staff_status": "",
        "job_title_id": "wqedsaqw",
        "resign_reason": "",
        "resign_type": "",
        "custom_field_values": [
            {
                "field_type": "1",
                "text_value": {
                    "default_value": "姓名字段",
                    "i18n_value": {
                        "zh_cn": "姓名字段",
                        "ja_jp": "姓名字段",
                        "en_us": "Name Filed"
                    }
                },
                "url_value": {
                    "link_text": {
                        "default_value": "网页地址",
                        "i18n_value": {
                            "zh_cn": "网页地址",
                            "ja_jp": "This is a URL",
                            "en_us": "This is a URL"
                        }
                    },
                    "url": "https://www.larksuite.com/",
                    "pcurl": "https://www.larksuite.com/"
                },
                "enum_value": {
                    "enum_ids": [
                        "75ec5g97"
                    ],
                    "enum_type": "1"
                },
                "user_values": [
                    {
                        "ids": [
                            "27al2hef"
                        ]
                ],
                "field_key": "C-1000001"
            }
        ]
    }
}

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
| 400 | 2221164 | User name exceeds limit | The name length exceeds the limit. You can enter up to 64 characters |
| 400 | 2221165 | User en_name exceeds limit | The English name exceeds the limit. The maximum number of characters is 64. |
| 400 | 2221166 | User another_name exceeds limit | The alias length exceeds the limit. You can enter up to 64 characters |
| 400 | 2221103 | Mobile already exists | The mobile phone number already exists, please modify the mobile phone number. |
| 400 | 2221106 | Invalid mobile | Invalid mobile phone number, please modify the mobile phone number. |
| 400 | 2221113 | Mobile or email not set | Mobile phone number or email address Required |
| 400 | 2221114 | Users must have a mobile in China | A Chinese mobile phone number is required for China. Please add a Chinese mobile phone number |
| 400 | 2221156 | Unable to edit verified mobile | Editing the verified mobile phone number is not allowed. Please do not edit the verified mobile phone number or modify it in other ways |
| 400 | 2221104 | Email already exists | The mailbox already exists, request to modify the mailbox. |
| 400 | 2221107 | Invalid email | Invalid mailbox, request to modify the enterprise mailbox. |
| 400 | 2221118 | Enterprise email already exists | The enterprise mailbox already exists, and the enterprise mailbox is requested to be modified. |
| 400 | 2221278 | Invalid enterprise email | Invalid mailbox, request to modify the enterprise mailbox. |
| 400 | 2221126 | Enterprise email domain unavailable | Enterprise mailbox domain name is not available, request to modify enterprise mailbox |
| 400 | 2221146 | Enterprise email alias exceeds limit | The enterprise email alias exceeds the length limit. You can enter up to 255 characters. |
| 400 | 2221147 | Enterprise email address in recycling bin | The enterprise email address is in Recycle Bin, please modify the enterprise email address. |
| 400 | 2221176 | Add Lark allows listing tenants. Email must be included with non + 86mobile | When a Lark tenant adds a mobile phone number other than +86, it must include email information. Please include email information. |
| 400 | 2221255 | Main department must be the first | The main department must be the first in the sorting information of the employee's department. Please modify the sorting information of the employee's department. |
| 400 | 2221125 | The number of members within the department exceeds the limit. Please contact an administrator for help | The number of members in the department exceeds the limit. The number cannot exceed 10,000. Please contact the administrator for help. |
| 400 | 2221129 | User department is empty | The department to which the employee belongs is empty, please modify the sorting information in the department to which the employee belongs |
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
| 400 | 2221193 | Extension number exceeds limit | The extension number length exceeds the limit. You can enter up to 99 characters |
| 400 | 2221210 | Invalid join date | Invalid entry time, please modify the entry time |
| 400 | 2221144 | EmployeeType not found | The person type does not exist, please modify the person type |
| 400 | 2221145 | EmployeeType inactive | The person type is not activated, please modify the person type |
| 400 | 2221223 | Invalid job title ID | The position does not exist, please change the position |
| 400 | 2221263 | Tenant has not activated multi geo | The tenant has not turned on Multi-Geo, please open it first and try again.Multi-Geo refers to multi-geo data residency. |
| 400 | 2221264 | User geo does not exist | The specified Geo does not exist. Please check whether the Geo parameter is correct. |
| 400 | 2221265 | The application does not have permission to write to the geo | You do not have permission to specify the employee Geo and need to apply for it. Click the api debugging console - permission configuration, and the required permissions will be displayed. Click "Operation" - "..." - "Open" to do so.  where employee data is written  |
| 400 | 2221266 | The application does not have permission to write to the SubscriptionID | You do not have permission to specify a seat. You need to apply for it. Click the API Debugging Console - Permission Configuration, and the required permissions will be displayed. Click "Operation" - "..." - "Open".  Write employee seat information  |
| 400 | 2224001 | No permission to operate | No operation permission. Please check the permission of the current application or whether the enterprise version is the business professional version or above. |
| 400 | 2224002 | No permission to operate records | You do not have permission to operate this record. Please check the permission of the data management scope of the current application or whether the member operated by the current application can be restored. |
| 400 | 2224003 | No permission to operate dependent objects | No operation permission for the dependent object. Please check whether the department to be updated has permission. |
| 400 | 2221252 | Hybrid license tenant prohibits passing empty licenses to create users | Mixed license tenants are prohibited from passing empty licenses to create users. Please pass a valid license. |
| 400 | 2221253 | Designated licenseID is insufficient | The remaining seats are insufficient, please modify the seat information |
| 400 | 2221254 | Designated licenseID is invalid | The application does not have permission to write to the employee data residence. You need to apply for permission to write to the employee data residence (directory: employee.base.geo: write) |
| 400 | 2221111 | Exceeds certified seat limit | Mixed license tenants are prohibited from passing empty licenses to create users. Please pass a valid license. |
| 400 | 2221112 | Exceeds billing plan seat limit | The remaining seats are insufficient, please modify the seat information |
| 400 | 2221115 | ExternalID is not unique | The custom ID already exists, please modify the custom ID. |
| 400 | 2221116 | Invalid ExternalID | Invalid custom ID, please modify the custom ID. |
| 400 | 2221175 | Lark only supports + 86Mobile | Lark tenants only support + 86 mobile phone number, please modify the mobile phone number. |
| 400 | 2221182 | Unable to freeze tenant founder | Enterprise administrators cannot n't freeze. Please select a non-enterprise administrator user to perform the freezing operation. |
| 400 | 2221109 | Name contains sensitive info | The name contains sensitive information, please modify the name. |
| 400 | 2221292 | User department is disabled | The user department has been deactivated. Please transfer the user to an active department or reactivate this department. |
| 400 | 2221213 | Resign date invalid or earlier than join date or empty | Please check that the separation date format may be illegal, the separation date is earlier than the joining date, or the separation date is empty. |
| 400 | 2221214 | Resign reason invalid or not match resign type | The reason for leaving is illegal, or the reason for leaving does not match the type of leaving. Please modify the reason for leaving or the type of leaving to match |
| 400 | 2221293 | Only allow update preResigned\resigned employee's resign info field | Only the departure information fields (departure remarks, departure type, reason for leaving, departure date) of employees to be separated or separated can be updated. Please update the corresponding fields of employees to be separated or separated. |
| 400 | 2221231 | Resign type invalid or not match resign reason | The type of leaving is illegal, or the reason for leaving does not match the type of leaving. Please modify the type of leaving or the reason for leaving to match it. |
