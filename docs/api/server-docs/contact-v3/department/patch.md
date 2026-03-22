---
title: "Modify department information in part"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/department/patch"
method: "PATCH"
api_path: "https://open.larksuite.com/open-apis/contact/v3/departments/:department_id"
service: "contact-v3"
resource: "department"
section: "Contacts"
rate_limit: "1000 per minute、50 per second"
scopes:
  - "contact:contact"
field_scopes:
  - "contact:department.organize:readonly"
  - "contact:department.base:readonly"
  - "contact:user.employee_id:readonly"
  - "contact:department.hrbp:readonly"
  - "contact:contact:access_as_app"
  - "contact:contact:readonly"
  - "contact:contact:readonly_as_app"
updated: "1732170447000"
---

# Modify department information in part

Call this interface to update some information of the specified department, including name, parent department, sorting, person in charge, etc.

## Precautions

- When calling this interface to update department information, all departments involved need to be within the application's contacts permissions, otherwise the call will fail and a no permissions error will be reported. To understand the scope of authority, see Authority Scope Resource Introduction.
- This interface is not a full update interface. If a request parameter is not passed, the corresponding department information will not be modified by default (Note: if an empty array is passed to leaders and department_hrbps, the original value will be cleared). If you need to use the full update interface, please refer to Update all department information.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/contact/v3/departments/:department_id |
| HTTP Method | PATCH |
| Rate Limit | 1000 per minute、50 per second |
| Supported app types | custom |
| Required scopes | `contact:contact` |
| Required field scopes | > The response body of the API contains the following sensitive fields, and they will be returned only after corresponding scopes are added. If you do not need the fields, it is not recommended that you request the scopes. `contact:department.organize:readonly` `contact:department.base:readonly` `contact:user.employee_id:readonly` `contact:department.hrbp:readonly` `contact:contact:access_as_app` `contact:contact:readonly` `contact:contact:readonly_as_app` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `department_id` | `string` | Department ID, ID type needs to be consistent with the value of the query parameter department_id_type. Instructions on how to obtain the ID: - After calling the Create Department interface, the department ID information can be obtained from the returned result. - The department API provides multiple ways to obtain the IDs of other departments, such as Get list of sub-departments, [Get parent department information]( /ssl:ttdoc/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/department/parent), search department, you can choose the appropriate one API for querying. **Example value**: "D096" **Data validation rules**: - Maximum length: `64` characters - Regular expression: `^[a-zA-Z0-9][a-zA-Z0-9_\-@.]{0,63}$` | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `user_id_type` | `string` | No | User ID categories **Example value**: open_id **Optional values are**:  - `open_id`: Identifies a user to an app. The same user has different Open IDs in different apps. How to get Open ID - `union_id`: Identifies a user to a tenant that acts as a developer. A user has the same Union ID in apps developed by the same developer, and has different Union IDs in apps developed by different developers. A developer can use Union ID to link the same user's identities in multiple apps.How to get Union ID - `user_id`: Identifies a user to a tenant. The same user has different User IDs in different tenants. In one single tenant, a user has the same User ID in all apps （including store apps）. User ID is usually used to communicate user data between different apps. How to get User ID  **Default value**: `open_id` **When the value is `user_id`, the following field scopes are required**: `contact:user.employee_id:readonly` |
| `department_id_type` | `string` | No | The type of department ID in this call. For a detailed introduction to department ID, please refer to Department ID Description. **Example value**: open_department_id **Optional values are**:  - `department_id`: Supports user-defined configuration of department IDs. The deleted department_id can be reused when customizing the configuration, so department_id is unique within the scope of the department that has not been deleted. - `open_department_id`: Department ID automatically generated by the system. The ID prefix is fixed to `od-` and is globally unique within the tenant.  **Default value**: `open_department_id` | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `name` | `string` | No | Department name. **Notice**: - Cannot contain slashes (`/`). - Cannot be the same as the existing department name. **Default value**: empty, indicating no modification. **Example value**: "DemoName" **Data validation rules**: - Minimum length: `1` characters |
| `i18n_name` | `department_i18n_name` | No | International configuration of department names. **Notice**: - Cannot contain slashes (`/`). - It cannot be repeated with the internationalization configuration of the existing department name. **Default value**: empty, indicating no modification. |
|   `zh_cn` | `string` | No | Department's Chinese name. **Example value**: "Demo name" |
|   `ja_jp` | `string` | No | Department's Japanese name. **Example value**: "Name" |
|   `en_us` | `string` | No | Department's English name. **Example value**: "Demo Name" |
| `parent_department_id` | `string` | No | The ID of the parent department. How to obtain department ID: - If the parent department of the department needs to be set as the root department, this parameter takes the value `0`. - You can call the Search Department interface to obtain the required department ID. **Default value**: empty, indicating no modification. **Example value**: "D067" |
| `leader_user_id` | `string` | No | The user ID of the department head. The ID type is consistent with the value of the query parameter user_id_type. For how to obtain user ID, please refer to How to obtain different user IDs. **Note**: The values of department head (leader_user_id) and department head (leaderID corresponding to leaderType value 1) are always consistent. therefore: - If the department head is set at the same time (leaderID corresponding to leaderType value 1), the department head set here must be the same person as the department head. - Only the department head is modified, and the department leader will be modified simultaneously (leaderID corresponding to leaderType value 1). **Default value**: empty, indicating no modification. **Example value**: "ou_7dab8a3d3cdcc9da365777c7ad535d62" |
| `order` | `string` | No | The sorting of departments is the order in which departments are displayed among departments at the same level. The value format is a non-negative integer of String type. The smaller the value, the higher the sorting is. **Note**: The order value is unique, that is, the value passed in cannot be the same as the order value of the stock department. **Default value**: empty, indicating no modification. **Example value**: "100" |
| `unit_ids` | `string[]` | No | A list of custom IDs for units bound to the department. Currently, only one unit is supported. **Default value**: empty, indicating no modification. **Example value**: ["custom_unit_id"] |
| `create_group_chat` | `boolean` | No | Whether to create department groups. **Optional values are:** - true: create. - false: Do not create. If a department group has been created previously, the group will continue to exist even if set to false. **Note**: When creating a department group, the group name defaults to the department name, and the group owner defaults to the department head. **Default value**: empty, indicating no modification. **Example value**: false |
| `leaders` | `departmentLeader[]` | No | Department head information. **Notice**: - If an empty array is passed to leaders, the original value will be cleared. - When configuring this parameter, you must specify a person in charge. - When setting up multiple persons in charge, only one person in charge can be set as the main person in charge. - The values of department head (leader_user_id) and department head (leaderID corresponding to leaderType value 1) are always consistent. therefore: - If a department head (leader_user_id) is set at the same time, the department head set here must be the same person as the department head. - Only the department leader is modified, and the department head (leader_user_id) will be modified simultaneously. |
|   `leaderType` | `int` | Yes | Person in charge type. **Example value**: 1 **Optional values are**:  - `1`: Main person in charge - `2`: Deputy responsible person  |
|   `leaderID` | `string` | Yes | The user ID and ID type of the person in charge are consistent with the value of the query parameter user_id_type. For how to obtain user ID, please refer to How to obtain different user IDs. **Example value**: "ou_7dab8a3d3cdcc9da365777c7ad535d62" |
| `group_chat_employee_types` | `int[]` | No | Personnel type restrictions for department groups. The value range of personnel type is as follows. This parameter supports setting multiple type values. If there are multiple types, separate them with English `,`: - 1: Regular employee - 2: Intern - 3: Outsourcing - 4: Labor services - 5: Consultant This parameter supports passing in the number corresponding to the custom person type. You can call the Query employee type interface to obtain the corresponding number (enum_value). **Default**: empty **Example value**: [1] |
| `department_hrbps` | `string[]` | No | List of user IDs for department HRBP. The ID type is consistent with the value of the query parameter user_id_type. For how to obtain user ID, please refer to How to obtain different user IDs. **Notice**: If an empty array is passed to department_hrbps, the original value will be cleared. **Example value**: ["ou_7dab8a3d3cdcc9da365777c7ad535d62"] **Data validation rules**: - Maximum length: `500` | ### Request body example

{
    "name": "DemoName",
    "i18n_name": {
        "zh_cn": "Demo name",
        "ja_jp": "Name",
        "en_us": "Demo Name"
    },
    "parent_department_id": "D067",
    "leader_user_id": "ou_7dab8a3d3cdcc9da365777c7ad535d62",
    "order": "100",
    "unit_ids": [
        "custom_unit_id"
    ],
    "create_group_chat": false,
    "leaders": [
        {
            "leaderType": 1,
            "leaderID": "ou_7dab8a3d3cdcc9da365777c7ad535d62"
        }
    ],
    "group_chat_employee_types": [
        1
    ],
    "department_hrbps": [
        "ou_7dab8a3d3cdcc9da365777c7ad535d62"
    ]
}

**Go Request Example**
```go
import (
	"context"

	"github.com/larksuite/oapi-sdk-go/v3"
	"github.com/larksuite/oapi-sdk-go/v3/service/contact/v3"
)

func main() {
	client := lark.NewClient("appID", "appSecret")

	req := larkcontact.NewPatchDepartmentReqBuilder().
		DepartmentId(`D096`).
		Department(larkcontact.NewDepartmentBuilder().
			Name(`DemoName`).
			I18nName(larkcontact.NewDepartmentI18nNameBuilder().
				ZhCn(`Demo名称`).
				EnUs(`Demo Name`).
				Build()).
			ParentDepartmentId(`D067`).
			Build()).
		Build()

	resp, err := client.Contact.Department.Patch(context.Background(), req)
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

        PatchDepartmentReq req = PatchDepartmentReq.newBuilder()
                .departmentId("D096")
                .department(Department.newBuilder()
                        .name("DemoName")
                        .i18nName(DepartmentI18nName.newBuilder()
                                .zhCn("Demo名称")
                                .enUs("Demo Name")
                                .build())
                        .parentDepartmentId("D067")
                        .build())
                .build();

        PatchDepartmentResp resp = client.contact().department().patch(req, RequestOptions.newBuilder().build());
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
|   `department` | `department` | Department information. |
|     `name` | `string` | Department name. **Required field scopes (Satisfy any)**: `contact:department.base:readonly` `contact:contact:access_as_app` `contact:contact:readonly` `contact:contact:readonly_as_app` |
|     `i18n_name` | `department_i18n_name` | Internationalized department name. **Required field scopes (Satisfy any)**: `contact:department.base:readonly` `contact:contact:access_as_app` `contact:contact:readonly` `contact:contact:readonly_as_app` |
|       `zh_cn` | `string` | Department's Chinese name. |
|       `ja_jp` | `string` | Department's Japanese name. |
|       `en_us` | `string` | Department's English name. |
|     `parent_department_id` | `string` | The department ID of the parent department. - The ID type is consistent with the department_id_type value of the query parameter. - When the parent department is the root department, the parameter value is `0`. **Required field scopes (Satisfy any)**: `contact:department.organize:readonly` `contact:contact:access_as_app` `contact:contact:readonly` `contact:contact:readonly_as_app` |
|     `department_id` | `string` | Custom department ID. This ID can be used later to delete, modify, and query department information. **Required field scopes (Satisfy any)**: `contact:department.base:readonly` `contact:contact:access_as_app` `contact:contact:readonly` `contact:contact:readonly_as_app` |
|     `open_department_id` | `string` | The open_department_id of the department is automatically generated by the system. This ID can be used later to delete, modify, and query department information. |
|     `leader_user_id` | `string` | The user ID and ID type of the department head are consistent with the user_id_type value of the query parameter. **Required field scopes (Satisfy any)**: `contact:department.organize:readonly` `contact:contact:access_as_app` `contact:contact:readonly` `contact:contact:readonly_as_app` |
|     `chat_id` | `string` | The group ID of the department group. You can later use Get Group Information to obtain the detailed information of the group. **Required field scopes (Satisfy any)**: `contact:department.base:readonly` `contact:contact:access_as_app` `contact:contact:readonly` `contact:contact:readonly_as_app` |
|     `order` | `string` | The sorting of departments is the order in which departments are displayed among departments at the same level. The smaller the value, the higher the order. **Required field scopes (Satisfy any)**: `contact:department.organize:readonly` `contact:contact:access_as_app` `contact:contact:readonly` `contact:contact:readonly_as_app` |
|     `unit_ids` | `string[]` | List of unit custom IDs bound to the department, currently only one is supported. **Required field scopes (Satisfy any)**: `contact:department.organize:readonly` `contact:contact:access_as_app` `contact:contact:readonly` `contact:contact:readonly_as_app` |
|     `member_count` | `int` | The number of users (including department heads) in the current department and its subordinate departments. **Required field scopes (Satisfy any)**: `contact:department.organize:readonly` `contact:contact:access_as_app` `contact:contact:readonly` `contact:contact:readonly_as_app` |
|     `status` | `department_status` | Department status. **Required field scopes (Satisfy any)**: `contact:department.base:readonly` `contact:contact:access_as_app` `contact:contact:readonly` `contact:contact:readonly_as_app` |
|       `is_deleted` | `boolean` | Whether it is deleted. **Possible values are:** - true - false |
|     `leaders` | `departmentLeader[]` | Head of department. |
|       `leaderType` | `int` | Person in charge type. **Optional values are**:  - `1`: Main person in charge - `2`: Deputy responsible person  |
|       `leaderID` | `string` | The user ID and ID type of the person in charge are consistent with the user_id_type value of the query parameter. **Required field scopes (Satisfy any)**: `contact:department.organize:readonly` `contact:contact:access_as_app` `contact:contact:readonly` `contact:contact:readonly_as_app` |
|     `group_chat_employee_types` | `int[]` | Personnel type restrictions for department groups. Possible values for person type are as follows: - 1: Regular employee - 2: Intern - 3: Outsourcing - 4: Labor services - 5: Consultant If it is a custom person type, the corresponding number will be returned. You can call the Query Personnel Type interface to obtain the custom employee type information corresponding to the corresponding number (enum_value) . |
|     `department_hrbps` | `string[]` | List of user IDs for department HRBP. The ID type is consistent with the value of the query parameter user_id_type. **Required field scopes**: `contact:department.hrbp:readonly` |
|     `primary_member_count` | `int` | The number of primary members of the current department and its subordinate departments (that is, the member's primary department is the current department). **Required field scopes (Satisfy any)**: `contact:department.organize:readonly` `contact:contact:access_as_app` `contact:contact:readonly` `contact:contact:readonly_as_app` | ### Response body example

{
    "code":0,
    "msg":"success",
    "data":{
        "department":{
            "name":"DemoName",
            "i18n_name":{
                "zh_cn":"Demo名称",
                "ja_jp":"デモ名",
                "en_us":"Demo Name"
            },
            "parent_department_id":"D067",
            "department_id":"D096",
            "open_department_id":"od-4e6ac4d14bcd5071a37a39de902c7141",
            "leader_user_id":"ou_7dab8a3d3cdcc9da365777c7ad535d62",
            "chat_id":"oc_5ad11d72b830411d72b836c20",
            "order":"100",
            "unit_ids":[
                "custom_unit_id"
            ],
            "member_count":100,
            "status":{
                "is_deleted":false
            },
            "leaders":[
                {
                    "leaderID":"ou_357368f98775f04bea02afc6b1d33478",
                    "leaderType":1
                }
            ],
            "department_hrbps":[
                "ou_7dab8a3d3cdcc9da365777c7ad535d62",
                "ou_7dab8a3d3cdcc9da365777c7ad535d63"
            ]
        }
    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 43005 | duplicate order error | The value of order is repeated. The order parameter value of the department must be unique and cannot be the same as the order value of the existing department. You can update the order value and try again. |
| 400 | 40002 | process root dept error | Operation on the root department is not supported. Please check whether the root department ID `0` is passed in the request parameter department ID. |
| 400 | 40003 | internal error | Internal error, please obtain the X-Request-Id of the request and provide feedback to [Technical Support](https://applink.larksuite.com/TLJpeNdW). |
| 403 | 40004 | no dept authority error | The departments involved in the current operation need to be within the permission scope of the application contacts. For more information, please refer to Permission Scope Resource Introduction. |
| 400 | 40008 | dept Info is null error | Department information cannot be empty. |
| 403 | 40014 | no parent dept authority error | There is no parent department permission. The parent department passed in needs to be within the address book permission scope of the application. For how to set the address book permission scope, see Introduction to permission scope resources. |
| 401 | 42008 | tenant id is invalid error | Invalid tenant identity. When requesting, the tenant identity corresponding to the Authorization request header needs to be the tenant to which the resource in the current operation belongs. |
| 400 | 43016 | leaders is repeat | The department head is duplicated. Please check whether the department head entered is correct. |
| 400 | 43017 | relate dept over limit | The specified unit has been associated with 1000 departments and cannot be associated with any more departments. |
| 400 | 43018 | duplicate i18n name | Duplicate internationalization configuration for department name. You need to modify the i18n_name parameter configuration and try again. |
| 400 | 43019 | exceed dept max level | The department hierarchy has reached 25 levels, and no more sub-departments can be created. |
| 400 | 43021 | department chat not exist | Department group does not exist |
| 400 | 43022 | department name duplicate | Department name is duplicated. You need to modify the passed name parameter value and try again. |
| 400 | 43023 | dept structure no permissione | There is no department permission. The department needs to be within the address book permission scope of the application. For how to set the address book permission scope, see Introduction to permission scope resources. |
| 400 | 43024 | dept structure tenant lock fail | Failed to obtain tenant lock for department structure change. Conflict caused by concurrent requests, please try again later. |
| 400 | 43025 | top department leader unjoined | The user has not joined and cannot become the department head. You need to change the department head to an active employee included in the contacts. |
| 400 | 43026 | employee type is not valid | Invalid person type. You need to refer to the parameter description of group_chat_employee_types in the interface document to set the correct employee type. |
| 400 | 43028 | invalid department hrbps | The department’s HRBP is illegal. When requesting, multiple user IDs for department HRBP can be set. You need to check whether the set user ID type is consistent with the type set by the query parameter user_id_type, and check whether the user ID value is correct. For how to obtain user ID, please refer to How to obtain different user IDs. |
| 400 | 43029 | dept name not contain separator | Department names cannot contain slashes (`/`). |
| 400 | 43030 | update department lock error,wait some seconds and retry | Concurrency is limited, please try again later | For more error code information, see General error codes.
