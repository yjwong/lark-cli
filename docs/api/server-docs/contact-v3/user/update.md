---
title: "Fully Update User Information"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/user/update"
method: "PUT"
api_path: "https://open.larksuite.com/open-apis/contact/v3/users/:user_id"
service: "contact-v3"
resource: "user"
section: "Contacts"
scopes:
  - "contact:contact"
field_scopes:
  - "contact:contact:readonly_as_app"
  - "contact:user.base:readonly"
  - "contact:user.department:readonly"
  - "contact:user.employee:readonly"
  - "contact:user.gender:readonly"
  - "contact:user.employee_id:readonly"
  - "contact:user.email:readonly"
  - "contact:user.phone:readonly"
updated: "1676364039000"
---

# Update user information in whole

This API is used to update a user's fields in Contacts. FAQs.

> The app must have the scope to access contacts data of the user to be updated. To change the user's departments, the app must also have the scope to access all new departments. Store apps have no scope to call this API.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/contact/v3/users/:user_id |
| HTTP Method | PUT |
| Supported app types | custom |
| Required scopes | `contact:contact` |
| Required field scopes | > The response body of the API contains the following sensitive fields, and they will be returned only after corresponding scopes are added. If you do not need the fields, it is not recommended that you request the scopes. `contact:contact:readonly_as_app` `contact:user.base:readonly` `contact:user.department:readonly` `contact:user.employee:readonly` `contact:user.gender:readonly` `contact:user.employee_id:readonly` `contact:user.email:readonly` `contact:user.phone:readonly` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `user_id` | `string` | User ID, which must match the user_id_type in the query parameter. **Example value**: "ou_7dab8a3d3cdcc9da365777c7ad535d62" | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `user_id_type` | `string` | No | User ID categories **Example value**: "open_id" **Optional values are**:  - `open_id`: Identifies a user to an app. The same user has different Open IDs in different apps. How to get Open ID - `union_id`: Identifies a user to a tenant that acts as a developer. A user has the same Union ID in apps developed by the same developer, and has different Union IDs in apps developed by different developers. A developer can use Union ID to link the same user's identities in multiple apps.How to get Union ID - `user_id`: Identifies a user to a tenant. The same user has different User IDs in different tenants. In one single tenant, a user has the same User ID in all apps （including store apps）. User ID is usually used to communicate user data between different apps. How to get User ID  **Default value**: `open_id` **When the value is `user_id`, the following field scopes are required**: `contact:user.employee_id:readonly` |
| `department_id_type` | `string` | No | Type of department ID used in this call **Example value**: "open_department_id" **Optional values are**:  - `department_id`: Identify the department with the custom department_id - `open_department_id`: Identify the department with open_department_id  **Default value**: `open_department_id` | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `name` | `string` | Yes | User name **Example value**: "John" **Data validation rules**: - Minimum length: `1` characters |
| `en_name` | `string` | No | English name **Example value**: "San Zhang" |
| `nickname` | `string` | No | Alias **Example value**: "Alex Zhang" |
| `email` | `string` | No | Email **Example value**: "zhangsan@gmail.com" |
| `mobile` | `string` | Yes | The mobile number cannot be repeated within the organization. Note that the area code must contain the plus sign +. **Example value**: "+41446681800" |
| `mobile_visible` | `boolean` | No | Visibility of mobile number. true: visible (default), false: invisible. If it is set to false, the employee's mobile number will be invisible to other company members. **Example value**: false |
| `gender` | `int` | No | Gender **Example value**: 1 **Optional values are**:  - `0`: Unknown - `1`: Male - `2`: Female  |
| `avatar_key` | `string` | No | File key of the profile photo. You can upload an image by calling the "Upload Image" API in "Messenger > Message > Images" to obtain the file key. **Example value**: "2500c7a9-5fff-4d9a-a2de-3d59614ae28g" |
| `department_ids` | `string[]` | Yes | List of IDs of the user's departments. A user can belong to multiple departments. The ID value matches the department_id_type in the query parameter. For information about different IDs, see Department IDs. **Example value**: od-4e6ac4d14bcd5071a37a39de902c7141 |
| `leader_user_id` | `string` | No | User ID of the user's direct manager. The ID value matches the user_id_type in the query parameter. For information about different IDs, see User-related IDs. **Example value**: "ou_7dab8a3d3cdcc9da365777c7ad535d62" |
| `city` | `string` | No | City **Example value**: "Hangzhou" |
| `country` | `string` | No | Country or region code. Refer to the list of country/region codes for details. **Example value**: "CN" |
| `work_station` | `string` | No | Seat ID **Example value**: "North building-H34" |
| `join_time` | `int` | No | Date of employment **Example value**: 2147483647 |
| `employee_no` | `string` | No | Employee ID **Example value**: "1" |
| `employee_type` | `int` | Yes | Employee type. Optional values are: - `1`: Regular - `2`: Intern - `3`: Outsourcing - `4`: Contractor - `5`: Consultant The int value of the custom employee type can also be read, and the name of the custom employee type for the tenant can be obtained through the API: Obtain workforce type **Example value**: 1 |
| `orders` | `user_order[]` | No | User sorting information |
|   `department_id` | `string` | No | Department ID for the sorting information. The ID value matches the department_id_type in the query parameter. For information about different IDs, see Department IDs. **Example value**: "od-4e6ac4d14bcd5071a37a39de902c7141" |
|   `user_order` | `int` | No | The user's ranking in their direct department. A larger value indicates a higher ranking. **Example value**: 100 |
|   `department_order` | `int` | No | Ranking of the user's departments. A larger value indicates a higher ranking. **Example value**: 100 |
| `custom_attrs` | `user_custom_attr[]` | No | Custom field. Make sure your company administrator has enabled "Allow to invoke by Open Platform's Contacts API" in Admin > Organization > Member Profile > Custom Fields > Global Settings; otherwise, this field will not take effect or be returned. |
|   `type` | `string` | No | Custom field type - `TEXT`: Text - `HREF`: Webpage - `ENUMERATION`: Enumeration - `PICTURE_ENUM`: Image - `GENERIC_USER`: User FAQs about custom fields **Example value**: "TEXT" |
|   `id` | `string` | No | Custom field ID **Example value**: "DemoId" |
|   `value` | `user_custom_attr_value` | No | Custom field value |
|     `text` | `string` | No | When the field type is `TEXT`, this parameter defines the field value and is required; when the field type is `HREF`, this parameter defines the webpage title and is required. **Example value**: "DemoText" |
|     `url` | `string` | No | When the field type is HREF, this parameter defines the default URL. **Example value**: "http://www.larksuite.com" |
|     `pc_url` | `string` | No | When the field type is HREF, this parameter defines URL on the PC. **Example value**: "http://www.larksuite.com" |
|     `option_id` | `string` | No | When the field type is ENUMERATION or PICTURE_ENUM, this parameter defines the option value. **Example value**: "edcvfrtg" |
|     `generic_user` | `custom_attr_generic_user` | No | When the field type is GENERIC_USER, this parameter defines the user referenced. |
|       `id` | `string` | Yes | User's user_id **Example value**: "9b2fabg5" |
|       `type` | `int` | Yes | User type    1: User **Example value**: 1 |
| `enterprise_email` | `string` | No | Business email. Make sure that Lark Business Mail has been enabled in Admin. **Example value**: "demo@mail.com" |
| `job_title` | `string` | No | Job title **Example value**: "xxxxx" |
| `is_frozen` | `boolean` | No | Whether to freeze the user **Example value**: false | ### Request body example

{
    "name": "John",
    "en_name": "San Zhang",
    "nickname": "Alex Zhang",
    "email": "zhangsan@gmail.com",
    "mobile": "+41446681800",
    "mobile_visible": false,
    "gender": 1,
    "avatar_key": "2500c7a9-5fff-4d9a-a2de-3d59614ae28g",
    "department_ids": [
        "od-4e6ac4d14bcd5071a37a39de902c7141"
    ],
    "leader_user_id": "ou_7dab8a3d3cdcc9da365777c7ad535d62",
    "city": "Hangzhou",
    "country": "CN",
    "work_station": "North building-H34",
    "join_time": 2147483647,
    "employee_no": "1",
    "employee_type": 1,
    "orders": [
        {
            "department_id": "od-4e6ac4d14bcd5071a37a39de902c7141",
            "user_order": 100,
            "department_order": 100
        }
    ],
    "custom_attrs": [
        {
            "type": "TEXT",
            "id": "DemoId",
            "value": {
                "text": "DemoText",
                "url": "http://www.larksuite.com",
                "pc_url": "http://www.larksuite.com",
                "option_id": "edcvfrtg",
                "generic_user": {
                    "id": "9b2fabg5",
                    "type": 1
                }
            }
        }
    ],
    "enterprise_email": "demo@mail.com",
    "job_title": "xxxxx",
    "is_frozen": false
}

## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `user` | `user` | User information |
|     `union_id` | `string` | User's union_id. For information about different IDs, see User-related IDs. |
|     `user_id` | `string` | Unique identifier of the user in the tenant, namely the user's user_id. For information of different IDs, refer to User-related IDs. **Required field scopes**: `contact:user.employee_id:readonly` |
|     `open_id` | `string` | User's open_id. For information about different IDs, see User-related IDs. |
|     `name` | `string` | User name **Required field scopes (Satisfy any)**: `contact:contact:readonly_as_app` `contact:user.base:readonly` |
|     `en_name` | `string` | English name **Required field scopes (Satisfy any)**: `contact:contact:readonly_as_app` `contact:user.base:readonly` |
|     `nickname` | `string` | Alias **Required field scopes (Satisfy any)**: `contact:contact:readonly_as_app` `contact:user.base:readonly` |
|     `email` | `string` | Email **Required field scopes**: `contact:user.email:readonly` |
|     `mobile` | `string` | The mobile number cannot be repeated within the organization. Note that the area code must contain the plus sign +. **Required field scopes**: `contact:user.phone:readonly` |
|     `mobile_visible` | `boolean` | Visibility of mobile number. true: visible (default), false: invisible. If it is set to false, the employee's mobile number will be invisible to other company members. |
|     `gender` | `int` | Gender **Optional values are**:  - `0`: Unknown - `1`: Male - `2`: Female  **Required field scopes (Satisfy any)**: `contact:contact:readonly_as_app` `contact:user.gender:readonly` |
|     `avatar_key` | `string` | File key of the profile photo. You can upload an image by calling the "Upload Image" API in "Messenger > Message > Images" to obtain the file key. |
|     `avatar` | `avatar_info` | User's profile photo **Required field scopes (Satisfy any)**: `contact:contact:readonly_as_app` `contact:user.base:readonly` |
|       `avatar_72` | `string` | Link of profile photo in 72*72 px |
|       `avatar_240` | `string` | Link of profile photo in 240*240 px |
|       `avatar_640` | `string` | Link of profile photo in 640*640 px |
|       `avatar_origin` | `string` | Link of the original profile photo |
|     `status` | `user_status` | User status **Required field scopes (Satisfy any)**: `contact:contact:readonly_as_app` `contact:user.employee:readonly` |
|       `is_frozen` | `boolean` | Whether the user is frozen. |
|       `is_resigned` | `boolean` | Whether the user is terminated. |
|       `is_activated` | `boolean` | Whether the user is activated. |
|       `is_exited` | `boolean` | Whether the user is voluntary. The user's status will be automatically changed to terminated after the user has left for a specific time of period. |
|       `is_unjoin` | `boolean` | Whether the user has not joined. The user must confirm in person to join the team. |
|     `department_ids` | `string[]` | List of IDs of the user's departments. A user can belong to multiple departments. The ID value matches the department_id_type in the query parameter. For information about different IDs, see Department IDs. **Required field scopes (Satisfy any)**: `contact:contact:readonly_as_app` `contact:user.department:readonly` |
|     `leader_user_id` | `string` | User ID of the user's direct manager. The ID value matches the user_id_type in the query parameter. For information about different IDs, see User-related IDs. **Required field scopes (Satisfy any)**: `contact:contact:readonly_as_app` `contact:user.department:readonly` |
|     `city` | `string` | City **Required field scopes (Satisfy any)**: `contact:contact:readonly_as_app` `contact:user.employee:readonly` |
|     `country` | `string` | Country or region code. Refer to the list of country/region codes for details. **Required field scopes (Satisfy any)**: `contact:contact:readonly_as_app` `contact:user.employee:readonly` |
|     `work_station` | `string` | Seat ID **Required field scopes (Satisfy any)**: `contact:contact:readonly_as_app` `contact:user.employee:readonly` |
|     `join_time` | `int` | Date of employment **Required field scopes (Satisfy any)**: `contact:contact:readonly_as_app` `contact:user.employee:readonly` |
|     `is_tenant_manager` | `boolean` | Whether the user is a super administrator of the tenant. **Required field scopes (Satisfy any)**: `contact:contact:readonly_as_app` `contact:user.employee:readonly` |
|     `employee_no` | `string` | Employee ID **Required field scopes (Satisfy any)**: `contact:contact:readonly_as_app` `contact:user.employee:readonly` |
|     `employee_type` | `int` | Employee type. Optional values are: - `1`: Regular - `2`: Intern - `3`: Outsourcing - `4`: Contractor - `5`: Consultant The int value of the custom employee type can also be read, and the name of the custom employee type for the tenant can be obtained through the API: Obtain workforce type **Required field scopes (Satisfy any)**: `contact:contact:readonly_as_app` `contact:user.employee:readonly` |
|     `orders` | `user_order[]` | User sorting information **Required field scopes (Satisfy any)**: `contact:contact:readonly_as_app` `contact:user.department:readonly` |
|       `department_id` | `string` | Department ID for the sorting information. The ID value matches the department_id_type in the query parameter. For information about different IDs, see Department IDs. |
|       `user_order` | `int` | The user's ranking in their direct department. A larger value indicates a higher ranking. |
|       `department_order` | `int` | Ranking of the user's departments. A larger value indicates a higher ranking. |
|     `custom_attrs` | `user_custom_attr[]` | Custom field. Make sure your company administrator has enabled "Allow to invoke by Open Platform's Contacts API" in Admin > Organization > Member Profile > Custom Fields > Global Settings; otherwise, this field will not take effect or be returned. **Required field scopes (Satisfy any)**: `contact:contact:readonly_as_app` `contact:user.employee:readonly` |
|       `type` | `string` | Custom field type - `TEXT`: Text - `HREF`: Webpage - `ENUMERATION`: Enumeration - `PICTURE_ENUM`: Image - `GENERIC_USER`: User FAQs about custom fields |
|       `id` | `string` | Custom field ID |
|       `value` | `user_custom_attr_value` | Custom field value |
|         `text` | `string` | When the field type is `TEXT`, this parameter defines the field value and is required; when the field type is `HREF`, this parameter defines the webpage title and is required. |
|         `url` | `string` | When the field type is HREF, this parameter defines the default URL. |
|         `pc_url` | `string` | When the field type is HREF, this parameter defines URL on the PC. |
|         `option_id` | `string` | When the field type is ENUMERATION or PICTURE_ENUM, this parameter defines the option value. |
|         `option_value` | `string` | Option value |
|         `name` | `string` | Name |
|         `picture_url` | `string` | Image link |
|         `generic_user` | `custom_attr_generic_user` | When the field type is GENERIC_USER, this parameter defines the user referenced. |
|           `id` | `string` | User's user_id |
|           `type` | `int` | User type    1: User |
|     `enterprise_email` | `string` | Business email. Make sure that Lark Business Mail has been enabled in Admin. **Required field scopes (Satisfy any)**: `contact:contact:readonly_as_app` `contact:user.employee:readonly` |
|     `job_title` | `string` | Job title **Required field scopes (Satisfy any)**: `contact:contact:readonly_as_app` `contact:user.employee:readonly` |
|     `is_frozen` | `boolean` | Whether to freeze the user | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
        "user": {
            "union_id": "on_94a1ee5551019f18cd73d9f111898cf2",
            "user_id": "3e3cf96b",
            "open_id": "ou_7dab8a3d3cdcc9da365777c7ad535d62",
            "name": "John",
            "en_name": "San Zhang",
            "nickname": "Alex Zhang",
            "email": "zhangsan@gmail.com",
            "mobile": "+41446681800",
            "mobile_visible": false,
            "gender": 1,
            "avatar_key": "2500c7a9-5fff-4d9a-a2de-3d59614ae28g",
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
            "department_ids": [
                "od-4e6ac4d14bcd5071a37a39de902c7141"
            ],
            "leader_user_id": "ou_7dab8a3d3cdcc9da365777c7ad535d62",
            "city": "Hangzhou",
            "country": "CN",
            "work_station": "North building-H34",
            "join_time": 2147483647,
            "is_tenant_manager": false,
            "employee_no": "1",
            "employee_type": 1,
            "orders": [
                {
                    "department_id": "od-4e6ac4d14bcd5071a37a39de902c7141",
                    "user_order": 100,
                    "department_order": 100
                }
            ],
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
            "is_frozen": false
        }
    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 40001 | param error | Parameter error. |
| 400 | 40016 | dept name can not be nul error | Department name cannot be empty. |
| 400 | 41001 | mobile has already exist error | Mobile number already exists. |
| 400 | 41002 | email has already exist error | Email already exists. |
| 409 | 41003 | user account conflict error | User's contact information involves two Lark accounts. Unable to add. The user can change the mobile number or email, or delete the account linked to the mobile number or email first, and create a new user or update the user. For how to delete an account, visit [Help Center](https://www.larksuite.com/hc/en-US/articles/360048488018). |
| 400 | 41004 | mobile is invalid error | Invalid mobile number. Check whether the mobile number format is correct. |
| 400 | 41005 | email is invalid error | Invalid email address. |
| 400 | 41006 | no user name error | No user name set. |
| 400 | 41009 | no email or mobile error | Email and mobile number cannot be both empty. |
| 400 | 41011 | userID already exist error | user_id is the unique ID of a user in a company and cannot be duplicated. |
| 400 | 41010 | no mobile error | Mobile number cannot be empty. |
| 400 | 41014 | user name sensitive error | Name contains sensitive information. For any questions, contact support. |
| 400 | 41013 | exceed userID update limit error | The number of user ID updates exceeds the limit. |
| 400 | 41015 | idp type invalid error error | Invalid login type. |
| 400 | 41016 | department has too many users  error | More than 500 users in a department. |
| 400 | 41017 | department is required error | Department information cannot be empty. |
| 400 | 41018 | position info is invalid error | Invalid position information. |
| 400 | 41019 | position department is invalid error | Invalid position department. |
| 400 | 41020 | position code has already exist  error | Invalid position code. |
| 400 | 41021 | position multiple main count error | Only one key position can be set for a user. |
| 400 | 41022 | user tenant not match error | Check if another company's credential is used to access the resources of this company. |
| 400 | 41023 | update department conflict position department error | User's position department must be one of the user's departments. When the user's departments are updated, the position department also needs to be updated; otherwise the update action will be blocked. |
| 400 | 41024 | update position department conflict department error | User's position department must be one of the user's departments. The user's new position department must be also one of the user's departments; otherwise the update action will be blocked. |
| 400 | 41025 | order department invalid error | The department ID in the requested user sorting information must be one of the user's department IDs. |
| 405 | 41028 | user multi department need upgrade visibility error | Update the supplemental rules for "Visible scope of organization" in Admin. |
| 400 | 41029 | create or update user multi department error | The company does not allow a user to join multiple departments. For any questions, contact support. |
| 400 | 41030 | set leader to oneself error | Check the user's direct manager parameter. |
| 504 | 41031 | position feature not enable error | Unable to set user's position information in the company. For any questions, contact support. |
| 504 | 41032 | user multi department feature not enable error | The company does not allow a user to join multiple departments. For any questions, contact support. |
| 400 | 41033 | user in too many departments  error | A user cannot belong to more than 50 departments. Please check. |
| 400 | 41034 | email prefix already exist error | Email prefix already exists. |
| 400 | 41035 | email prefix is invalid error | Invalid email prefix. Check the spelling. |
| 400 | 41036 | avatar key is invalid error | Invalid profile photo key. |
| 400 | 41037 | avatar key is sensitive error | Profile photo key contains sensitive information. |
| 400 | 41038 | gender is invalid error | Invalid gender. Please check. |
| 400 | 41040 | user name is null error | User name cannot be empty. |
| 400 | 41041 | department id is not assigned  error | User's department cannot be empty. |
| 400 | 41042 | join time is invalid error | User's join time cannot be empty. |
| 400 | 41043 | employee id is invalid error | Invalid user ID. The length must be between 1 and 64 bytes. |
| 400 | 41044 | Custom attribute is not set error | Please set the custom field. The field ID must be specified. The field ID can be obtained through the API "Obtain company's custom fields". |
| 400 | 41045 | Custom attribute id is not exist error | Custom field ID does not exist. Check the source of the custom field ID. The custom field ID can be obtained through the API "Obtain company's custom fields". |
| 400 | 41046 | Custom attribute value is not set error | Please set the custom field. The value field must be specified. |
| 400 | 41047 | Custom attribute href text  is null error | Please set the custom field of the HREF type. The text field is required. |
| 400 | 41048 | Custom attribute href url  is null error | Please set the custom field of the HREF type. The url field is required. |
| 400 | 41050 | no user authority error | The user must be within the range of contacts data that the app can access. Learn more. |
| 400 | 41051 | user id info not provide error | User ID not specified. |
| 403 | 40004 | no dept authority error | The department must be within the range of contacts data that the app can access. Learn more |
| 400 | 41059 | invalid employee type error | User's employee type is incorrect. Fill in a number between 1 and 5. 1: Regular, 2: Intern, 3: Outsourcing, 4: Contractor, 5: Consultant. |
| 400 | 41060 | inactive employee type error | Employee type not used by the company. Consult the administrator. |
| 400 | 41063 | job_title length exceed 100 character | Job title exceeds 100 characters. Check the field length. |
| 400 | 41068 | Number of email aliases exceeds the upper limit | The number of business email accounts has reached the limit. Contact the company administrator. |
| 400 | 41069 | Business email is in the recycle bin | Business email is being recycled and cannot be used. |
| 400 | 41070 | name length exceed 64 character | Name exceeds 64 characters. |
| 400 | 41071 | en_name length exceed 64 character | English name exceeds 64 characters. |
| 400 | 41072 | nickname length exceed 64 character | Alias exceeds 64 characters. |
| 400 | 44001 | business email domain not available error | No business email domain for the company. Contact the company administrator. |
| 400 | 44002 | update order must update department together | The list of department IDs must be included to change the user's department ranking. |
| 400 | 44006 | name length exceed 64 character | Name exceeds 64 characters. |
| 400 | 44007 | en_name length exceed 64 character | English name exceeds 64 characters. |
| 400 | 44008 | nickname length exceed 64 character | Alias exceeds 64 characters. |
| 400 | 44010 | unJoined user not allow to update | A user who has not joined cannot be updated. |
| 400 | 44011 | existed user not allow to update | Activated user cannot be updated. |
| 400 | 42006 | user has resigned error | The user has been terminated. |
| 400 | 44013 | User enterprise Email password is not valid | - |
| 400 | 44014 | Can not update inactive user email when email equal enterprise | - |
| 400 | 44015 | can not update password when user already have password | - |
| 400 | 44016 | can not set enterprise email password | - |
| 400 | 44017 | Suite_Admin_Common_UnableToEditUpper | Unable to edit. Maximum number of times a single user's contact information can be modified has reached the upper limit. |
| 400 | 44018 | lark not support +86 mobile | +86 mobile numbers are not supported |
| 400 | 44019 | - | - |
| 400 | 44020 | - | - |
| 400 | 44024 | User enterprise email has already been registered as a member's account | User's email has been registered |
| 400 | 44025 | update user lock error,wait some seconds and retry | update user lock error, wait some seconds and retry |
| 400 | 44035 | departmentID is invaild | departmentID is invaild, modify and retry |
| 400 | 44036 | freeze tenant founder is forbidden | this user can not been update by security issues |
| 400 | 44038 | this user can not been update by security issues | req set user geo not find in geo list |
