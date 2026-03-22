---
title: "Obtain the list of users directly under a department"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/user/find_by_department"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/contact/v3/users/find_by_department"
service: "contact-v3"
resource: "user"
section: "Contacts"
rate_limit: "1000 per minute、50 per second"
scopes:
  - "contact:contact.base:readonly"
  - "contact:contact:readonly_as_app"
  - "contact:department.organize:readonly"
  - "contact:contact:access_as_app"
  - "contact:contact:readonly"
field_scopes:
  - "contact:contact:readonly_as_app"
  - "contact:user.base:readonly"
  - "contact:user.department:readonly"
  - "contact:user.department_path:readonly"
  - "contact:user.dotted_line_leader_info.read"
  - "contact:user.employee:readonly"
  - "contact:user.employee_number:read"
  - "contact:user.gender:readonly"
  - "contact:user.user_geo"
  - "contact:user.employee_id:readonly"
  - "contact:user.phone:readonly"
  - "contact:user.email:readonly"
  - "contact:user.job_family:readonly"
  - "contact:user.job_level:readonly"
  - "contact:contact:access_as_app"
  - "contact:contact:readonly"
updated: "1692861218000"
---

# Obtain the list of users directly under a department

This API is used to obtain the list of users directly under a department based on the department ID.

> - Department ID is required. The root department ID is 0.
> - When `user_access_token` is used, the filtering is based on the users' visible scope of organization, and the user data within the users' visible scope of organization ([Log in to the company's Admin to manage scopes]) will be returned.
> - When `tenant_access_token` is used, the filtering is based on the app's contact scope. If the requested department ID is 0, the system verifies whether the app has the all-staff contact scope. If the department ID is not 0, the system verifies whether the app has the contact scope of the department. If no, an error code indicating no scope will be returned. If yes, the list of users directly under the department will be returned.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/contact/v3/users/find_by_department |
| HTTP Method | GET |
| Rate Limit | 1000 per minute、50 per second |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `contact:contact.base:readonly` `contact:contact:readonly_as_app` `contact:department.organize:readonly` `contact:contact:access_as_app` `contact:contact:readonly` |
| Required field scopes | > The response body of the API contains the following sensitive fields, and they will be returned only after corresponding scopes are added. If you do not need the fields, it is not recommended that you request the scopes. `contact:contact:readonly_as_app` `contact:user.base:readonly` `contact:user.department:readonly` `contact:user.department_path:readonly` `contact:user.dotted_line_leader_info.read` `contact:user.employee:readonly` `contact:user.employee_number:read` `contact:user.gender:readonly` `contact:user.user_geo` `contact:user.employee_id:readonly` `contact:user.phone:readonly` `contact:user.email:readonly` `contact:user.job_family:readonly` `contact:user.job_level:readonly` `contact:contact:access_as_app` `contact:contact:readonly` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `user_id_type` | `string` | No | User ID categories **Example value**: open_id **Optional values are**:  - `open_id`: Identifies a user to an app. The same user has different Open IDs in different apps. How to get Open ID - `union_id`: Identifies a user to a tenant that acts as a developer. A user has the same Union ID in apps developed by the same developer, and has different Union IDs in apps developed by different developers. A developer can use Union ID to link the same user's identities in multiple apps.How to get Union ID - `user_id`: Identifies a user to a tenant. The same user has different User IDs in different tenants. In one single tenant, a user has the same User ID in all apps （including store apps）. User ID is usually used to communicate user data between different apps. How to get User ID  **Default value**: `open_id` **When the value is `user_id`, the following field scopes are required**: `contact:user.employee_id:readonly` |
| `department_id_type` | `string` | No | Type of department ID used in this call Difference between department type refer to The Description of department_id **Example value**: open_department_id **Optional values are**:  - `department_id`: Identify the department with the custom department_id - `open_department_id`: Identify the department with the custom department_id  **Default value**: `open_department_id` |
| `department_id` | `string` | Yes | The department under which users are to be obtained. This field is required. The department ID of root department is 0. **Example value**: od-xxxxxxxxxxxxx |
| `page_size` | `int` | No | paging size **Example value**: 10 **Default value**: `10` **Data validation rules**: - Maximum value: `50` |
| `page_token` | `string` | No | Page identifier. It is not filled in the first request, indicating traversal from the beginning; when there will be more groups, the new page_token will be returned at the same time, and the next traversal can use the page_token to get more groups **Example value**: AQD9/Rn9eij9Pm39ED40/dk53s4Ebp882DYfFaPFbz00L4CMZJrqGdzNyc8BcZtDbwVUvRmQTvyMYicnGWrde9X56TgdBuS+JKiSIkdexPw contact.v3.type.user.prop.city.desc=$$$City | **Go Request Example**
```go
import (
	"context"

	"github.com/larksuite/oapi-sdk-go/v3"
	"github.com/larksuite/oapi-sdk-go/v3/service/contact/v3"
)

func main() {
	client := lark.NewClient("appID", "appSecret")

	req := larkcontact.NewFindByDepartmentUserReqBuilder().
		UserIdType(`open_id`).
		DepartmentIdType(`open_department_id`).
		DepartmentId(`od-xxxxxxxxxxxxx`).
		PageSize(10).
		Build()

	resp, err := client.Contact.User.FindByDepartment(context.Background(), req)
}
```

**Java Request Example**
```java
import com.lark.oapi.Client;
import com.lark.oapi.service.contact.v3.model.*;
import com.lark.oapi.core.request.RequestOptions;

public class Main {
    public static void main(String arg[]) throws Exception {
        Client client = Client.newBuilder("appId", "appSecret").build();

        FindByDepartmentUserReq req = FindByDepartmentUserReq.newBuilder()
                .userIdType("open_id")
                .departmentIdType("open_department_id")
                .departmentId("od-xxxxxxxxxxxxx")
                .pageSize(10)
                .build();

        FindByDepartmentUserResp resp = client.contact().user().findByDepartment(req, RequestOptions.newBuilder().build());
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
|   `has_more` | `boolean` | Whether the response body has more parameters |
|   `page_token` | `string` | Page identifier, when has_more is true, a new page_token will also be returned. Otherwise, page_token will not be returned |
|   `items` | `user[]` | User information list |
|     `union_id` | `string` | User's union_id. For information about different IDs, see User-related IDs. |
|     `user_id` | `string` | Unique identifier of the user in the tenant, namely the user's user_id. For information of different IDs, refer to User-related IDs. **Required field scopes**: `contact:user.employee_id:readonly` |
|     `open_id` | `string` | User's open_id. For information about different IDs, see User-related IDs. |
|     `name` | `string` | User name **Required field scopes (Satisfy any)**: `contact:contact:readonly_as_app` `contact:user.base:readonly` `contact:contact:access_as_app` `contact:contact:readonly` |
|     `en_name` | `string` | English name **Required field scopes (Satisfy any)**: `contact:contact:readonly_as_app` `contact:user.base:readonly` `contact:contact:access_as_app` `contact:contact:readonly` |
|     `nickname` | `string` | Alias **Required field scopes (Satisfy any)**: `contact:contact:readonly_as_app` `contact:user.base:readonly` `contact:contact:access_as_app` `contact:contact:readonly` |
|     `email` | `string` | Email Address. Note that 1. non +86 mobile number members must also add an email address 2. Email cannot be repeated **Required field scopes**: `contact:user.email:readonly` |
|     `mobile` | `string` | 1. The mobile number cannot be repeated within the organization. 2. Uncertified enterprises only support adding mobile numbers in mainland China. Enterprises certified by Lark are allowed to add overseas mobile numbers. 3. The area code must contain the plus sign +. 4. This mobile field is not required in the overseas version of Lark **Required field scopes**: `contact:user.phone:readonly` |
|     `mobile_visible` | `boolean` | Visibility of mobile number. true: visible (default), false: invisible. If it is set to false, the employee's mobile number will be invisible to other company members. |
|     `gender` | `int` | Gender **Optional values are**:  - `0`: Unknown - `1`: Male - `2`: Female  **Required field scopes (Satisfy any)**: `contact:contact:readonly_as_app` `contact:user.gender:readonly` `contact:contact:access_as_app` `contact:contact:readonly` |
|     `avatar_key` | `string` | File key of the profile photo. You can upload an image by calling the "Upload Image" API in "Messenger > Message > Images" to obtain the file key. The document refer to upload images |
|     `avatar` | `avatar_info` | User's profile photo **Required field scopes (Satisfy any)**: `contact:contact:readonly_as_app` `contact:user.base:readonly` `contact:contact:access_as_app` `contact:contact:readonly` |
|       `avatar_72` | `string` | Link of profile photo in 72*72 px |
|       `avatar_240` | `string` | Link of profile photo in 240*240 px |
|       `avatar_640` | `string` | Link of profile photo in 640*640 px |
|       `avatar_origin` | `string` | Link of the original profile photo |
|     `status` | `user_status` | User status，enumeration types, including is_frozen, is_resigned, is_activated, is_exited User state transition see: User State Diagram **Required field scopes (Satisfy any)**: `contact:contact:readonly_as_app` `contact:user.employee:readonly` `contact:contact:access_as_app` `contact:contact:readonly` |
|       `is_frozen` | `boolean` | Whether the user is frozen. |
|       `is_resigned` | `boolean` | Whether the user is terminated. |
|       `is_activated` | `boolean` | Whether the user is activated. |
|       `is_exited` | `boolean` | Whether the user is voluntary. The user's status will be automatically changed to terminated after the user has left for a specific time of period. |
|       `is_unjoin` | `boolean` | Whether the user has not joined. The user must confirm in person to join the team. |
|     `department_ids` | `string[]` | List of IDs of the user's departments. A user can belong to multiple departments. The ID value matches the department_id_type in the query parameter. For information about different IDs, see Department IDs. **Required field scopes (Satisfy any)**: `contact:contact:readonly_as_app` `contact:user.department:readonly` `contact:contact:access_as_app` `contact:contact:readonly` |
|     `leader_user_id` | `string` | User ID of the user's direct manager. The ID value matches the user_id_type in the query parameter. For information about different IDs, see User-related IDs. **Required field scopes (Satisfy any)**: `contact:contact:readonly_as_app` `contact:user.department:readonly` `contact:contact:access_as_app` `contact:contact:readonly` |
|     `city` | `string` | work city **Required field scopes (Satisfy any)**: `contact:contact:readonly_as_app` `contact:user.employee:readonly` `contact:contact:access_as_app` `contact:contact:readonly` |
|     `country` | `string` | Country or region code. Refer to the list of country/region codes for details. **Required field scopes (Satisfy any)**: `contact:contact:readonly_as_app` `contact:user.employee:readonly` `contact:contact:access_as_app` `contact:contact:readonly` |
|     `work_station` | `string` | Seat ID **Required field scopes (Satisfy any)**: `contact:contact:readonly_as_app` `contact:user.employee:readonly` `contact:contact:access_as_app` `contact:contact:readonly` |
|     `join_time` | `int` | Entry time, in timestamp format, in seconds since January 1, 1970 **Required field scopes (Satisfy any)**: `contact:contact:readonly_as_app` `contact:user.employee:readonly` `contact:contact:access_as_app` `contact:contact:readonly` |
|     `is_tenant_manager` | `boolean` | Whether the user is a super administrator of the tenant. **Required field scopes (Satisfy any)**: `contact:contact:readonly_as_app` `contact:user.employee:readonly` `contact:contact:access_as_app` `contact:contact:readonly` |
|     `employee_no` | `string` | Employee ID **Required field scopes (Satisfy any)**: `contact:user.employee:readonly` `contact:user.employee_number:read` `contact:contact:readonly_as_app` `contact:contact:access_as_app` `contact:contact:readonly` |
|     `employee_type` | `int` | Employee type. Optional values are: - `1`: Regular - `2`: Intern - `3`: Outsourcing - `4`: Contractor - `5`: Consultant The int value of the custom employee type can also be read, and the name of the custom employee type for the tenant can be obtained through the API: Obtain workforce type **Required field scopes (Satisfy any)**: `contact:contact:readonly_as_app` `contact:user.employee:readonly` `contact:contact:access_as_app` `contact:contact:readonly` |
|     `orders` | `user_order[]` | User sorting information . Indicating the order of personnel in the organizational structure. Personnel may exist in multiple departments and have different orders. **Required field scopes (Satisfy any)**: `contact:contact:readonly_as_app` `contact:user.department:readonly` `contact:contact:access_as_app` `contact:contact:readonly` |
|       `department_id` | `string` | Department ID for the sorting information. The ID value matches the department_id_type in the query parameter. For information about different IDs, see Department IDs. |
|       `user_order` | `int` | The user's ranking in their direct department. A larger value indicates a higher ranking. |
|       `department_order` | `int` | Ranking of the user's departments. A larger value indicates a higher ranking. |
|       `is_primary_dept` | `boolean` | Identifies the user's unique main department. The main department is the department that ranks first among the user's departments (department_order is the max) |
|     `custom_attrs` | `user_custom_attr[]` | Custom field. Make sure your company administrator has enabled "Allow to invoke by Open Platform's Contacts API" in Admin > Organization > Member Profile > Custom Fields > Global Settings; otherwise, this field will not take effect or be returned. more detail see FQAs **Required field scopes (Satisfy any)**: `contact:contact:readonly_as_app` `contact:user.employee:readonly` `contact:contact:access_as_app` `contact:contact:readonly` |
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
|     `enterprise_email` | `string` | Business email. Make sure that Lark Business Mail has been enabled in Admin. When create user, the detail of enterprise email see FQAs **Required field scopes (Satisfy any)**: `contact:contact:readonly_as_app` `contact:user.employee:readonly` `contact:contact:access_as_app` `contact:contact:readonly` |
|     `job_title` | `string` | Job title **Required field scopes (Satisfy any)**: `contact:contact:readonly_as_app` `contact:user.employee:readonly` `contact:contact:access_as_app` `contact:contact:readonly` |
|     `is_frozen` | `boolean` | Whether to freeze the user |
|     `geo` | `string` | geo **Required field scopes**: `contact:user.user_geo` |
|     `job_level_id` | `string` | Job level ID **Required field scopes**: `contact:user.job_level:readonly` |
|     `job_family_id` | `string` | Job family ID **Required field scopes**: `contact:user.job_family:readonly` |
|     `department_path` | `department_detail[]` | Department path **Required field scopes**: `contact:user.department_path:readonly` |
|       `department_id` | `string` | Department ID |
|       `department_name` | `department_path_name` | Department name |
|         `name` | `string` | Department name |
|         `i18n_name` | `department_i18n_name` | Department international name |
|           `zh_cn` | `string` | Chinese name of the department |
|           `ja_jp` | `string` | Japanese name of the department |
|           `en_us` | `string` | English name of the department |
|       `department_path` | `department_path` | Department path |
|         `department_ids` | `string[]` | Department Path IDs |
|         `department_path_name` | `department_path_name` | Department path name |
|           `name` | `string` | Department name |
|           `i18n_name` | `department_i18n_name` | Department international name |
|             `zh_cn` | `string` | Chinese name of the department |
|             `ja_jp` | `string` | Japanese name of the department |
|             `en_us` | `string` | English name of the department |
|     `dotted_line_leader_user_ids` | `string[]` | dotted_line_leader_user_ids **Required field scopes**: `contact:user.dotted_line_leader_info.read` | ### Response body example

{"code":0,
"msg":"success",
"data":{"has_more":true,
"page_token":"AQD9/Rn9eij9Pm39ED40/RD/cIFmu77WxpxPB/2oHfQLZ%2BG8JG6tK7%2BZnHiT7COhD2hMSICh/eBl7cpzU6JEC3J7COKNe4jrQ8ExwBCR",
"items":[{"union_id":"on_94a1ee5551019f18cd73d9f111898cf2",
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
"department_path":[{
    "department_id": "od-4e6ac4d14bcd5071a37a39de902c7141",
    "department_name": {
        "name": "测试部门名1",
        "i18n_name": {
            "zh_cn": "Test department name 1",
            "ja_jp": "Deployment name 1",
            "en_us": "Testing department name 1"
        }
    },
    "department_path": {
        "department_ids": [
            "od-4e6ac4d14bcd5071a37a39de902c7141"
        ],
        "department_path_name": {
            "name": "测试部门名1",
            "i18n_name": {
                "zh_cn": "Test department name 1",
                "ja_jp": "Deployment name 1",
                "en_us": "Testing department name 1"
            }
        }
    }
}],
"dotted_line_leader_user_ids":["ou_7dab8a3d3cdcc9da365777c7ad535d62"]}]}}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 41050 | no user authority error | The user must be within the range of contacts data that the app can access. Learn more. |
| 400 | 40011 | page size is invalid | The page size is between 1 and 50. |
| 400 | 40012 | page token is invalid error | Page token is invalid. |
| 403 | 40004 | no dept authority error | The department must be within the range of contacts data that the app can access. Learn more |
