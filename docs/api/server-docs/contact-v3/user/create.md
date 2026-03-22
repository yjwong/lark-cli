---
title: "Create a user"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/user/create"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/contact/v3/users"
service: "contact-v3"
resource: "user"
section: "Contacts"
rate_limit: "1000 per minute、50 per second"
scopes:
  - "contact:contact"
field_scopes:
  - "contact:user.base:readonly"
  - "contact:user.department:readonly"
  - "contact:user.dotted_line_leader_info.read"
  - "contact:user.employee:readonly"
  - "contact:user.employee_number:read"
  - "contact:user.gender:readonly"
  - "contact:user.email:readonly"
  - "contact:user.job_level:readonly"
  - "contact:user.employee_id:readonly"
  - "contact:user.phone:readonly"
  - "contact:user.user_geo"
  - "contact:user.job_family:readonly"
  - "contact:contact:access_as_app"
  - "contact:contact:readonly"
  - "contact:contact:readonly_as_app"
updated: "1703526088000"
---

# Create a user

This API is used to create a user in Contacts when a new employee is employed. After a user is created, only data within the data scope will be returned. For the relationship between data scopes and fields, refer to App scopes.

> - When in mixed license mode, the Seat field is required.

> - A user can be added only when the app has the scope to access contacts of all departments of the user. To add a user under the root department, the app must have the scope to access all users.
> - Store apps have no scope to call this API.
> - After a user is created, an invitation text message or email will be sent to the user. The user must agree to join the team before they can access the team.
> -The mobile phone number is not returned in the returned data. If necessary, please re-query the user information to obtain the mobile phone number.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/contact/v3/users |
| HTTP Method | POST |
| Rate Limit | 1000 per minute、50 per second |
| Supported app types | custom |
| Required scopes | `contact:contact` |
| Required field scopes | > The response body of the API contains the following sensitive fields, and they will be returned only after corresponding scopes are added. If you do not need the fields, it is not recommended that you request the scopes. `contact:user.base:readonly` `contact:user.department:readonly` `contact:user.dotted_line_leader_info.read` `contact:user.employee:readonly` `contact:user.employee_number:read` `contact:user.gender:readonly` `contact:user.email:readonly` `contact:user.job_level:readonly` `contact:user.employee_id:readonly` `contact:user.phone:readonly` `contact:user.user_geo` `contact:user.job_family:readonly` `contact:contact:access_as_app` `contact:contact:readonly` `contact:contact:readonly_as_app` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `user_id_type` | `string` | No | User ID categories **Example value**: open_id **Optional values are**:  - `open_id`: Identifies a user to an app. The same user has different Open IDs in different apps. How to get Open ID - `union_id`: Identifies a user to a tenant that acts as a developer. A user has the same Union ID in apps developed by the same developer, and has different Union IDs in apps developed by different developers. A developer can use Union ID to link the same user's identities in multiple apps.How to get Union ID - `user_id`: Identifies a user to a tenant. The same user has different User IDs in different tenants. In one single tenant, a user has the same User ID in all apps （including store apps）. User ID is usually used to communicate user data between different apps. How to get User ID  **Default value**: `open_id` **When the value is `user_id`, the following field scopes are required**: `contact:user.employee_id:readonly` |
| `department_id_type` | `string` | No | Type of department ID used in this call. For the information of different IDs and how to obtain them, refer to Department IDs. **Example value**: open_department_id **Optional values are**:  - `department_id`: Identify the department with the custom department_id - `open_department_id`: Identify the department with open_department_id  **Default value**: `open_department_id` |
| `client_token` | `string` | No | Used to determine whether it is the same request. **Example value**: xxxx-xxxxx-xxx | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `user_id` | `string` | No | Unique identifier of the user in the tenant, namely the user's user_id. For information of different IDs, refer to User-related IDs. **Example value**: "3e3cf96b" |
| `name` | `string` | Yes | User name **Example value**: "John" **Data validation rules**: - Minimum length: `1` characters |
| `en_name` | `string` | No | English name **Example value**: "San Zhang" |
| `nickname` | `string` | No | Alias **Example value**: "Alex Zhang" |
| `email` | `string` | No | Email Address. Note that 1. non +86 mobile number members must also add an email address 2. Email cannot be repeated **Example value**: "Alex@gmail.com" |
| `mobile` | `string` | Yes | 1. The mobile number cannot be repeated within the organization. 2. Uncertified enterprises only support adding mobile numbers in mainland China. Enterprises certified by Lark are allowed to add overseas mobile numbers. 3. The area code must contain the plus sign +. 4. This mobile field is not required in the overseas version of Lark **Example value**: "13011111111(Other examples, Mobile numbers in mainland China：13011111111 or +8613011111111； Overseas mobile numbers: +41446681800)" |
| `mobile_visible` | `boolean` | No | Visibility of mobile number. true: visible (default), false: invisible. If it is set to false, the employee's mobile number will be invisible to other company members. **Example value**: false |
| `gender` | `int` | No | Gender **Example value**: 1 **Optional values are**:  - `0`: Unknown - `1`: Male - `2`: Female  |
| `avatar_key` | `string` | No | File key of the profile photo. You can upload an image by calling the "Upload Image" API in "Messenger > Message > Images" to obtain the file key. The document refer to upload images **Example value**: "2500c7a9-5fff-4d9a-a2de-3d59614ae28g" |
| `department_ids` | `string[]` | Yes | List of IDs of the user's departments. A user can belong to multiple departments. The ID value matches the department_id_type in the query parameter. For information about different IDs, see Department IDs. **Example value**: ["od-4e6ac4d14bcd5071a37a39de902c7141"] |
| `leader_user_id` | `string` | No | User ID of the user's direct manager. The ID value matches the user_id_type in the query parameter. For information about different IDs, see User-related IDs. **Example value**: "ou_7dab8a3d3cdcc9da365777c7ad535d62" |
| `city` | `string` | No | work city **Example value**: "Hangzhou" |
| `country` | `string` | No | Country or region code. Refer to the list of country/region codes for details. **Example value**: "CN" |
| `work_station` | `string` | No | Seat ID **Example value**: "North building-H34" |
| `join_time` | `int` | No | Entry time, in timestamp format, in seconds since January 1, 1970 **Example value**: 2147483647 |
| `employee_no` | `string` | No | Employee ID **Example value**: "1" |
| `employee_type` | `int` | Yes | Employee type. Optional values are: - `1`: Regular - `2`: Intern - `3`: Outsourcing - `4`: Contractor - `5`: Consultant The int value of the custom employee type can also be read, and the name of the custom employee type for the tenant can be obtained through the API: Obtain workforce type **Example value**: 1 |
| `orders` | `user_order[]` | No | User sorting information . Indicating the order of personnel in the organizational structure. Personnel may exist in multiple departments and have different orders. |
|   `department_id` | `string` | No | Department ID for the sorting information. The ID value matches the department_id_type in the query parameter. For information about different IDs, see Department IDs. **Example value**: "od-4e6ac4d14bcd5071a37a39de902c7141" |
|   `user_order` | `int` | No | The user's ranking in their direct department. A larger value indicates a higher ranking. **Example value**: 100 |
|   `department_order` | `int` | No | Ranking of the user's departments. A larger value indicates a higher ranking. **Example value**: 100 |
|   `is_primary_dept` | `boolean` | No | Identifies the user's unique main department. The main department is the department that ranks first among the user's departments (department_order is the max) **Example value**: true |
| `custom_attrs` | `user_custom_attr[]` | No | Custom field. Make sure your company administrator has enabled "Allow to invoke by Open Platform's Contacts API" in Admin > Organization > Member Profile > Custom Fields > Global Settings; otherwise, this field will not take effect or be returned. more detail see FQAs |
|   `type` | `string` | No | Custom field type - `TEXT`: Text - `HREF`: Webpage - `ENUMERATION`: Enumeration - `PICTURE_ENUM`: Image - `GENERIC_USER`: User FAQs about custom fields **Example value**: "TEXT" |
|   `id` | `string` | No | Custom field ID **Example value**: "DemoId" |
|   `value` | `user_custom_attr_value` | No | Custom field value |
|     `text` | `string` | No | When the field type is `TEXT`, this parameter defines the field value and is required; when the field type is `HREF`, this parameter defines the webpage title and is required. **Example value**: "DemoText" |
|     `url` | `string` | No | When the field type is HREF, this parameter defines the default URL. , such as mobile phone jump to applet, PC jump to webpage . **Example value**: "http://www.fs.cn" |
|     `pc_url` | `string` | No | When the field type is HREF, this parameter defines URL on the PC. **Example value**: "http://www.fs.cn" |
|     `option_id` | `string` | No | When the field type is ENUMERATION or PICTURE_ENUM, this parameter defines the option value. **Example value**: "edcvfrtg" |
|     `generic_user` | `custom_attr_generic_user` | No | When the field type is GENERIC_USER, this parameter defines the user referenced. |
|       `id` | `string` | Yes | User's user_id **Example value**: "9b2fabg5" |
|       `type` | `int` | Yes | User type    1: User **Example value**: 1 |
| `enterprise_email` | `string` | No | Business email. Make sure that Lark Business Mail has been enabled in Admin. When create user, the detail of enterprise email see FQAs **Example value**: "demo@mail.com" |
| `job_title` | `string` | No | Job title **Example value**: "xxxxx" |
| `geo` | `string` | No | geo **Example value**: "cn" |
| `job_level_id` | `string` | No | Job level **Example value**: "mga5oa8ayjlp9rb" |
| `job_family_id` | `string` | No | Job family **Example value**: "mga5oa8ayjlp9rb" |
| `subscription_ids` | `string[]` | No | The list of seat IDs assigned to users needs to be enabled with the "Assign User Seats" permission. The available seat ID of the tenant can be obtained through the interface below, see [Get Enterprise Seat Information] (/ssl: ttdoc/uAjLw4CM/ukTMukTMukTM/tenant-v2/tenant-product_assign_info/query). This field is required in mixed license mode. **Example value**: ["23213213213123123"] |
| `dotted_line_leader_user_ids` | `string[]` | No | dotted_line_leader_user_ids **Example value**: ["ou_7dab8a3d3cdcc9da365777c7ad535d62"] | ### Request body example

{"user_id":"3e3cf96b",
"name":"John",
"en_name":"San Zhang",
"nickname":"Alex Zhang",
"email":"Alex@gmail.com",
"mobile":"13011111111(Other examples, Mobile numbers in mainland China：13011111111 or +8613011111111；
Overseas mobile numbers: +41446681800)",
"mobile_visible":false,
"gender":1,
"avatar_key":"2500c7a9-5fff-4d9a-a2de-3d59614ae28g",
"department_ids":["od-4e6ac4d14bcd5071a37a39de902c7141"],
"leader_user_id":"ou_7dab8a3d3cdcc9da365777c7ad535d62",
"city":"Hangzhou",
"country":"CN",
"work_station":"North building-H34",
"join_time":2147483647,
"employee_no":"1",
"employee_type":1,
"orders":[{
    "department_id": "od-4e6ac4d14bcd5071a37a39de902c7141",
    "user_order": 100,
    "department_order": 100,
    "is_primary_dept": true
}],
"custom_attrs":[{
    "type": "TEXT",
    "id": "DemoId",
    "value": {
        "text": "DemoText",
        "url": "http://www.fs.cn",
        "pc_url": "http://www.fs.cn",
        "option_id": "edcvfrtg",
        "generic_user": {
            "id": "9b2fabg5",
            "type": 1
        }
    }
}],
"enterprise_email":"demo@mail.com",
"job_title":"xxxxx",
"geo":"cn",
"job_level_id":"mga5oa8ayjlp9rb",
"job_family_id":"mga5oa8ayjlp9rb",
"subscription_ids":["23213213213123123"],
"dotted_line_leader_user_ids":["ou_7dab8a3d3cdcc9da365777c7ad535d62"]}

**Go Request Example**

```go
import (
	"context"

	"github.com/larksuite/oapi-sdk-go/v3"
	"github.com/larksuite/oapi-sdk-go/v3/service/contact/v3"
)

func main() {
	client := lark.NewClient("appID", "appSecret")

	req := larkcontact.NewCreateUserReqBuilder().
		UserIdType(`open_id`).
		User(larkcontact.NewUserBuilder().
			Name(`zhangsan`).
			Mobile(`13011111111`).
			DepartmentIds([]string{`od-4e6ac4d14bcd5071a37a39de902c7141`}).
			EmployeeType(1).
			Build()).
		Build()

	resp, err := client.Contact.User.Create(context.Background(), req)
}
```

**Java Request Example**
```java
import com.lark.oapi.Client;
import com.lark.oapi.core.request.RequestOptions;
import com.lark.oapi.service.contact.v3.model.CreateUserReq;
import com.lark.oapi.service.contact.v3.model.CreateUserResp;
import com.lark.oapi.service.contact.v3.model.User;

public class Main {
    public static void main(String arg[]) throws Exception {
        Client client = Client.newBuilder("appId", "appSecret").build();

        CreateUserReq req = CreateUserReq.newBuilder()
                .userIdType("open_id")
                .user(User.newBuilder()
                        .name("zhangsan")
                        .mobile("13011111111")
                        .departmentIds(new String[]{"od-4e6ac4d14bcd5071a37a39de902c7141"})
                        .employeeType(1)
                        .build())
                .build();

        CreateUserResp resp = client.contact().user().create(req, RequestOptions.newBuilder().build());
    }
}
```

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
|     `name` | `string` | User name **Required field scopes (Satisfy any)**: `contact:user.base:readonly` `contact:contact:readonly_as_app` |
|     `en_name` | `string` | English name **Required field scopes (Satisfy any)**: `contact:user.base:readonly` `contact:contact:readonly_as_app` |
|     `nickname` | `string` | Alias **Required field scopes (Satisfy any)**: `contact:user.base:readonly` `contact:contact:readonly_as_app` |
|     `email` | `string` | Email Address. Note that 1. non +86 mobile number members must also add an email address 2. Email cannot be repeated **Required field scopes**: `contact:user.email:readonly` |
|     `mobile` | `string` | 1. The mobile number cannot be repeated within the organization. 2. Uncertified enterprises only support adding mobile numbers in mainland China. Enterprises certified by Lark are allowed to add overseas mobile numbers. 3. The area code must contain the plus sign +. 4. This mobile field is not required in the overseas version of Lark **Required field scopes**: `contact:user.phone:readonly` |
|     `mobile_visible` | `boolean` | Visibility of mobile number. true: visible (default), false: invisible. If it is set to false, the employee's mobile number will be invisible to other company members. |
|     `gender` | `int` | Gender **Optional values are**:  - `0`: Unknown - `1`: Male - `2`: Female  **Required field scopes (Satisfy any)**: `contact:user.gender:readonly` `contact:contact:readonly_as_app` |
|     `avatar_key` | `string` | File key of the profile photo. You can upload an image by calling the "Upload Image" API in "Messenger > Message > Images" to obtain the file key. The document refer to upload images |
|     `avatar` | `avatar_info` | User's profile photo **Required field scopes (Satisfy any)**: `contact:user.base:readonly` `contact:contact:readonly_as_app` |
|       `avatar_72` | `string` | Link of profile photo in 72*72 px |
|       `avatar_240` | `string` | Link of profile photo in 240*240 px |
|       `avatar_640` | `string` | Link of profile photo in 640*640 px |
|       `avatar_origin` | `string` | Link of the original profile photo |
|     `status` | `user_status` | User status，enumeration types, including is_frozen, is_resigned, is_activated, is_exited User state transition see: User State Diagram **Required field scopes (Satisfy any)**: `contact:user.employee:readonly` `contact:contact:readonly_as_app` |
|       `is_frozen` | `boolean` | Whether the user is frozen. |
|       `is_resigned` | `boolean` | Whether the user is terminated. |
|       `is_activated` | `boolean` | Whether the user is activated. |
|       `is_exited` | `boolean` | Whether the user is voluntary. The user's status will be automatically changed to terminated after the user has left for a specific time of period. |
|       `is_unjoin` | `boolean` | Whether the user has not joined. The user must confirm in person to join the team. |
|     `department_ids` | `string[]` | List of IDs of the user's departments. A user can belong to multiple departments. The ID value matches the department_id_type in the query parameter. For information about different IDs, see Department IDs. **Required field scopes (Satisfy any)**: `contact:user.department:readonly` `contact:contact:readonly_as_app` |
|     `leader_user_id` | `string` | User ID of the user's direct manager. The ID value matches the user_id_type in the query parameter. For information about different IDs, see User-related IDs. **Required field scopes (Satisfy any)**: `contact:user.department:readonly` `contact:contact:readonly_as_app` |
|     `city` | `string` | work city **Required field scopes (Satisfy any)**: `contact:user.employee:readonly` `contact:contact:readonly_as_app` |
|     `country` | `string` | Country or region code. Refer to the list of country/region codes for details. **Required field scopes (Satisfy any)**: `contact:user.employee:readonly` `contact:contact:readonly_as_app` |
|     `work_station` | `string` | Seat ID **Required field scopes (Satisfy any)**: `contact:user.employee:readonly` `contact:contact:readonly_as_app` |
|     `join_time` | `int` | Entry time, in timestamp format, in seconds since January 1, 1970. If you do not specify the entry time when creating a user, the entry time is configured by default. **Required field scopes (Satisfy any)**: `contact:user.employee:readonly` `contact:contact:readonly_as_app` |
|     `is_tenant_manager` | `boolean` | Whether the user is a super administrator of the tenant. **Required field scopes (Satisfy any)**: `contact:user.employee:readonly` `contact:contact:readonly_as_app` |
|     `employee_no` | `string` | Employee ID **Required field scopes (Satisfy any)**: `contact:user.employee_number:read` `contact:user.employee:readonly` `contact:contact:access_as_app` `contact:contact:readonly` `contact:contact:readonly_as_app` |
|     `employee_type` | `int` | Employee type. Optional values are: - `1`: Regular - `2`: Intern - `3`: Outsourcing - `4`: Contractor - `5`: Consultant The int value of the custom employee type can also be read, and the name of the custom employee type for the tenant can be obtained through the API: Obtain workforce type **Required field scopes (Satisfy any)**: `contact:user.employee:readonly` `contact:contact:readonly_as_app` |
|     `orders` | `user_order[]` | User sorting information . Indicating the order of personnel in the organizational structure. Personnel may exist in multiple departments and have different orders. **Required field scopes (Satisfy any)**: `contact:user.department:readonly` `contact:contact:readonly_as_app` |
|       `department_id` | `string` | Department ID for the sorting information. The ID value matches the department_id_type in the query parameter. For information about different IDs, see Department IDs. |
|       `user_order` | `int` | The user's ranking in their direct department. A larger value indicates a higher ranking. |
|       `department_order` | `int` | Ranking of the user's departments. A larger value indicates a higher ranking. |
|       `is_primary_dept` | `boolean` | Identifies the user's unique main department. The main department is the department that ranks first among the user's departments (department_order is the max) |
|     `custom_attrs` | `user_custom_attr[]` | Custom field. Make sure your company administrator has enabled "Allow to invoke by Open Platform's Contacts API" in Admin > Organization > Member Profile > Custom Fields > Global Settings; otherwise, this field will not take effect or be returned. more detail see FQAs **Required field scopes (Satisfy any)**: `contact:user.employee:readonly` `contact:contact:readonly_as_app` |
|       `type` | `string` | Custom field type - `TEXT`: Text - `HREF`: Webpage - `ENUMERATION`: Enumeration - `PICTURE_ENUM`: Image - `GENERIC_USER`: User FAQs about custom fields |
|       `id` | `string` | Custom field ID |
|       `value` | `user_custom_attr_value` | Custom field value |
|         `text` | `string` | When the field type is `TEXT`, this parameter defines the field value and is required; when the field type is `HREF`, this parameter defines the webpage title and is required. |
|         `url` | `string` | When the field type is HREF, this parameter defines the default URL. , such as mobile phone jump to applet, PC jump to webpage . |
|         `pc_url` | `string` | When the field type is HREF, this parameter defines URL on the PC. |
|         `option_id` | `string` | When the field type is ENUMERATION or PICTURE_ENUM, this parameter defines the option value. |
|         `option_value` | `string` | Option value |
|         `name` | `string` | Name |
|         `picture_url` | `string` | Image link |
|         `generic_user` | `custom_attr_generic_user` | When the field type is GENERIC_USER, this parameter defines the user referenced. |
|           `id` | `string` | User's user_id |
|           `type` | `int` | User type    1: User |
|     `enterprise_email` | `string` | Business email. Make sure that Lark Business Mail has been enabled in Admin. When create user, the detail of enterprise email see FQAs **Required field scopes (Satisfy any)**: `contact:user.employee:readonly` `contact:contact:readonly_as_app` |
|     `job_title` | `string` | Job title **Required field scopes (Satisfy any)**: `contact:user.employee:readonly` `contact:contact:readonly_as_app` |
|     `is_frozen` | `boolean` | Whether to freeze the user |
|     `geo` | `string` | geo **Required field scopes**: `contact:user.user_geo` |
|     `job_level_id` | `string` | Job level ID **Required field scopes**: `contact:user.job_level:readonly` |
|     `job_family_id` | `string` | Job family ID **Required field scopes**: `contact:user.job_family:readonly` |
|     `dotted_line_leader_user_ids` | `string[]` | dotted_line_leader_user_ids **Required field scopes**: `contact:user.dotted_line_leader_info.read` | ### Response body example

{"code":0,
"msg":"success",
"data":{"user":{"union_id":"on_94a1ee5551019f18cd73d9f111898cf2",
"user_id":"3e3cf96b",
"open_id":"ou_7dab8a3d3cdcc9da365777c7ad535d62",
"name":"John",
"en_name":"San Zhang",
"nickname":"Alex Zhang",
"email":"Alex@gmail.com",
"mobile":"13011111111(Other examples, Mobile numbers in mainland China：13011111111 or +8613011111111；
Overseas mobile numbers: +41446681800)",
"mobile_visible":false,
"gender":1,
"avatar_key":"2500c7a9-5fff-4d9a-a2de-3d59614ae28g",
"avatar":{"avatar_72":"https://foo.icon.com/xxxx",
"avatar_240":"https://foo.icon.com/xxxx",
"avatar_640":"https://foo.icon.com/xxxx",
"avatar_origin":"https://foo.icon.com/xxxx"},
"status":{"is_frozen":false,
"is_resigned":false,
"is_activated":true,
"is_exited":false,
"is_unjoin":false},
"department_ids":["od-4e6ac4d14bcd5071a37a39de902c7141"],
"leader_user_id":"ou_7dab8a3d3cdcc9da365777c7ad535d62",
"city":"Hangzhou",
"country":"CN",
"work_station":"North building-H34",
"join_time":2147483647,
"is_tenant_manager":false,
"employee_no":"1",
"employee_type":1,
"orders":[{
    "department_id": "od-4e6ac4d14bcd5071a37a39de902c7141",
    "user_order": 100,
    "department_order": 100,
    "is_primary_dept": true
}],
"custom_attrs":[{
    "type": "TEXT",
    "id": "DemoId",
    "value": {
        "text": "DemoText",
        "url": "http://www.fs.cn",
        "pc_url": "http://www.fs.cn",
        "option_id": "edcvfrtg",
        "option_value": "option",
        "name": "name",
        "picture_url": "https://xxxxxxxxxxxxxxxxxx",
        "generic_user": {
            "id": "9b2fabg5",
            "type": 1
        }
    }
}],
"enterprise_email":"demo@mail.com",
"job_title":"xxxxx",
"is_frozen":false,
"geo":"cn",
"job_level_id":"mga5oa8ayjlp9rb",
"job_family_id":"mga5oa8ayjlp9rb",
"dotted_line_leader_user_ids":["ou_7dab8a3d3cdcc9da365777c7ad535d62"]}}}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 41001 | mobile has already exist error | Mobile number already exists. |
| 400 | 41002 | email has already exist error | Email already exists. |
| 409 | 41003 | user account conflict error | User's contact information involves two Lark accounts. Unable to add. The user can change the mobile number or email, or delete the account linked to the mobile number or email first, and create a new user or update the user. For how to delete an account, visit [Help Center] For more information about this error, please refer to the description in [General Error Codes]() |
| 400 | 41004 | mobile is invalid error | Invalid mobile number. Check whether the mobile number format is correct. |
| 400 | 41005 | email is invalid error | Invalid email address. |
| 400 | 41006 | no user name error | No user name set. |
| 400 | 41007 | exceed uncertain tenant seat limit error | The number of users of an unverified tenant cannot exceed the limit. |
| 400 | 41008 | exceed bill seat limit error | The number of users in the tenant exceeds the limit. |
| 400 | 41009 | no email or mobile error | Email and mobile number cannot be both empty. |
| 400 | 41010 | no mobile error | Mobile number cannot be empty. |
| 400 | 41011 | userID already exist error | user_id is the unique ID of a user in a company and cannot be duplicated. |
| 400 | 41012 | user id invalid error | Invalid user ID. |
| 400 | 41013 | exceed userID update limit error | The number of user ID updates exceeds the limit. |
| 400 | 41014 | user name sensitive error | Name contains sensitive information. For any questions, contact support. |
| 400 | 41015 | idp type invalid error error | Invalid login type. |
| 400 | 41016 | department has too many users  error | There are too many users in a department and the number of users exceeds the maximum limit |
| 400 | 41017 | department is required error | Department information cannot be empty. |
| 400 | 41018 | position info is invalid error | Invalid position information. |
| 400 | 41019 | position department is invalid error | Invalid position department. |
| 400 | 41020 | position code has already exist  error | Invalid position code. |
| 400 | 41021 | position multiple main count error | Only one key position can be set for a user. |
| 400 | 41022 | user tenant not match error | Check if another company's credential is used to access the resources of this company. |
| 400 | 41025 | order department invalid error | The department ID in the requested user sorting information must be one of the user's department IDs. |
| 504 | 41027 | create account failed error | Failed to create user. |
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
| 400 | 41051 | user id info not provide error | User ID not specified. |
| 400 | 41052 | user resign acceptor is invalid error | Invalid user resignation approver. |
| 409 | 41053 | user has already exist error | User already exists. |
| 400 | 40003 | internal error | Internal error. Please provide the X-Request-Id to our agent. [Contact support](https://applink.larksuite.com/client/helpdesk/open?id |
| 403 | 40004 | no dept authority error | The department must be within the range of contacts data that the app can access. Learn more |
| 400 | 40021 | no a same request error | The two actions do not belong to the same request. Check whether the request value was changed. |
| 400 | 41054 | need send email but not set mail | Check whether an email is set. |
| 400 | 41055 | need send sms but not set mobile | Check whether a mobile number is set. |
| 400 | 41059 | invalid employee type error | User's employee type is incorrect. Fill in a number between 1 and 5. 1: Regular, 2: Intern, 3: Outsourcing, 4: Contractor, 5: Consultant. |
| 400 | 41060 | inactive employee type error | Employee type not used by the company. Consult the administrator. |
| 400 | 41063 | job_title length exceed 100 character | Job title exceeds 100 characters. Check the field length. |
| 400 | 41068 | Number of email aliases exceeds the upper limit | The number of business email accounts has reached the limit. Contact the company administrator. |
| 400 | 41069 | Business email is in the recycle bin | Business email is being recycled and cannot be used. |
| 400 | 41070 | name length exceed 64 character | Name exceeds 64 characters. |
| 400 | 41071 | en_name length exceed 64 character | English name exceeds 64 characters. |
| 400 | 41072 | nickname length exceed 64 character | Alias exceeds 64 characters. |
| 400 | 44001 | business email domain not available error | No business email domain for the company. Contact the company administrator. |
| 400 | 44004 | this user has been joined too many tenants recently | The user has been added to too many teams. |
| 400 | 44006 | name length exceed 64 character | Name exceeds 64 characters. |
| 400 | 44007 | en_name length exceed 64 character | English name exceeds 64 characters. |
| 400 | 44008 | nickname length exceed 64 character | Alias exceeds 64 characters. |
| 400 | 44009 | this tenant has been create too many users recently | Too many members have been added to the team. |
| 400 | 44012 | Adding user has been intercepted. Contact Lark Customer Service | Failed to add the user. For details, contact the Lark agent. |
| 400 | 44013 | User enterprise Email password is not valid | The user's enterprise email password entered is invalid, please modify it and try again. |
| 400 | 44016 | can not set enterprise email password | Failed to set the enterprise email password, please confirm whether you have the corresponding permission. |
| 400 | 44018 | lark not support +86 mobile | Lark does not support the use of + 86 mobile phone number, please modify the contact information and try again. |
| 400 | 44019 | Lark does not support the use of + 86 mobile phone number, please modify the contact information and try again. | Uncertified enterprises only support adding +86 mobile numbers. Please complete Lark verification before adding non +86 phone numbers. |
| 400 | 44020 | mobile and email need together exist | Note that non +86 mobile number members must also add an email address |
| 400 | 44021 | Leader is resigned | The person in charge has resigned |
| 400 | 44022 | leaderID is Invalid | Person in charge is an invalid parameter |
| 400 | 44023 | exceed feature contact seat limit | Exceeds the address book limit of the plan |
| 400 | 44024 | User enterprise email has already been registered as a member's account | User's email has been registered |
| 400 | 44038 | req set user geo not find in geo list | req set user geo not find in geo list |
| 400 | 44039 | not set geo name auth | not set geo name auth |
| 400 | 44040 | tenant not open mg not set geo name | The tenant cannot set the geo field if the MG has not been activated |
| 400 | 44044 | invalid job level id | invalid job level id |
| 400 | 44045 | invalid job family id | invalid job family id |
| 400 | 44046 | User license subscription id must not empty in multi-license tenant | Specify seat id when creating user |
| 400 | 44047 | License subscription id exceeds limit | The seat id has exceeded the upper limit |
| 400 | 44048 | User license subscription id invalid | Please confirm to pass in the correct seat id |
| 403 | 44050 | Not setting subscription ids auth | Please open the "Assign User Seats" permission. |
| 400 | 44051 | employee_no already existed | Employee job number duplicate, please modify and try again. |
| 400 | 41410 | The main department must be the first department among the user's departments (department_order is the max) | The main department must be the first department among the user's departments (department_order is the max) |
