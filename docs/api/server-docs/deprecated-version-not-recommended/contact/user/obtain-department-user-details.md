---
title: "Obtain Department User Details"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uYzN3QjL2czN04iN3cDN"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/contact/v1/department/user/detail/list?department_id=od-846fa088e3dcac5d0ce0b4f6b4cde899&page_size=10&fetch_child=true"
section: "Deprecated Version (Not Recommended)"
updated: "1646999057000"
---

# Obtain department user details

This API is used to obtain a department user detail information. 

> In addtion, user Data Permissions will put a constaraint on this API.The Contacts permission for this department must have been granted

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/contact/v1/department/user/detail/list?department_id=od-846fa088e3dcac5d0ce0b4f6b4cde899&page_size=10&fetch_child=true |
| HTTP Method | GET |
| Required scopes Enable any scope from the list |  Access department's organize information  Read contacts |
| Required field scopes Enable the scopes for fields you want to return |  Obtain user ID  Obtain user's basic information  Access contacts by application identity  Obtain user's email information  Obtain user's mobile number  Obtain user's gender  Obtain user's employment information  Obtain user's organization information | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value Format**: "Bearer `access_token`" **Example Value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" Learn more about obtaining and using access_token. |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Query parameters
|Parameter|Type|Mandatory|Description|
|-|-|-|-|
|department_id|string|Yes|Department ID|
|page_token|string|No|Page identifier. If not included in the first request, the list will iterate from the top. If a particular page is queried and there are more members in the role, a new page_token will also be returned. The next request can use that page_token to obtain more members|
|page_size|int|Yes|Page size, up to 100 items per page|
|fetch_child|bool|No|Whether sub-department users are recursively returned. Within the application authorized scope, if the number of recursive departments exceeds 500, code 40162 is returned. You can call  Obtain a sub-department list  to query the first level sub-department, and then call this API to query the user list.| ## Response
### Response body

|Parameter|Description|
|-|-|
|code|Error code, anything but 0 indicates failure|
|msg|Response code description|
|data|Response business information|
|  ∟has_nore|If department has more users|
|  ∟page_token|Page identifier. When has_more is true, a new page_token will also be returned. Otherwise, page_token will not be returned|
|  ∟user_infos|User information list|
|    ∟name|User name|
|    ∟en_name|English name|
|    ∟employee_id|User employee ID. The field will be returned after apps have applied "Obtain user's userid" permission|
|    ∟employee_no|User Work No.|
|    ∟open_id|User open ID|
|    ∟union_id|User union ID|
|    ∟status|User statuses. bit0 (lowest): 1 frozen, 0 unfrozen; bit1: 1 left the company, 0 current employee; bit2: 1 unactivated, 0 activated|
|    ∟employee_type|Employee type. 1: Regular employee, 2: Intern, 3:Freelancer, 4: Agency worker, 5: Consultant|
|    ∟avatar_url|Profile picture|
|    ∟gender|Gender,1:male；2:female|
|    ∟email|User email address, only returned when the email address permission has been granted|
|    ∟description|User signature|
|    ∟country|Country|
|    ∟city|City|
|    ∟work_station|Seat|
|    ∟is_tenant_manager|If is an administrator of a company|
|    ∟join_time|Attendance time|
|    ∟update_time|Update time|
|    ∟leader_employee_id|Employee ID of supervisor. The field will be returned after apps have applied "Obtain user's userid" permission|
|    ∟leader_open_id|Open ID of supervisor|
|    ∟leader_union_id|Union ID of supervisor|
|    ∟departments|Departments of user|
|    ∟custom_attrs|Custom user properties set by the organization. Each property includes a custom property ID and custom value set for user. When organization open custom field and set custom value for user, this field will be returned |
### Response body example
```json
{
    "code": 0,
    "msg": "success",
    "data": {
        "has_more": true,
        "page_token":"763bd1e74d05e95e",
        "user_infos": [
            {
                "name":"zhang san",
                "name_py":"zhang san",
                "en_name":"John",
                "employee_id":"a0615a67",
                "employee_no":"235634",
                "open_id":"ou_e03053f0541cecc3269d7a9dc34a0b21",  
                "union_id":"on_7dba11ff38a2119f89349876b12af65c",
                "status":2,
                "employee_type": 1,
                "avatar_url":"https://open.larksuite.com/avatar/d27500044f02b6c68b6e",
                "gender":1,
                "email":"zhangsan@gmail.com",
                "description": "",
                "country": "CN",
                "city":"Beijing",
                "work_station":"Poly, F6-123",  
                "is_tenant_manager":false,
                "join_time":1562342314,
                "update_time":1569140032,
                "leader_employee_id":"a0615a67",
                "leader_open_id":"ou_e03053f0541cecc3269d7a9dc34a0b21",
                "leader_union_id":"on_c132837f686587dd494aa54f5f65b552",
                "departments":[
                    "od-8c6c97ab9a34c1a649001d7ad36b97a7"
                ],
                "custom_attrs": {
                    "C-6702376000044400907": {
                        "value": "value1"
                    },
                    "C-6702376000048595214": {
                        "value": "value2"
                    }
                }  
            },
            {
                "name":"li si",
                "name_py":"li si",
                "en_name":"Jack",
                ...
            }
        ]
    }
}

```
### Error code

For details, please refer to: Service-side error codes
