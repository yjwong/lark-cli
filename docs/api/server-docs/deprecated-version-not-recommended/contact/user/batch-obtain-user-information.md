---
title: "Batch Obtain User Information"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uIzNz4iM3MjLyczM"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/contact/v1/user/get?open_id=ou_337f846fd508d9b4b620478eced55d91&open_ids=2545643646&open_ids=4673565565 \nhttps://open.larksuite.com/open-apis/contact/v1/user/batch_get?employee_ids=1afbg844&employee_ids=e62e73eg"
section: "Deprecated Version (Not Recommended)"
updated: "1646999051000"
---

# Batch obtain user information

> In order to  improve the security level, we have upgraded it. Please migrate to the new version as soon as possible. New version >>

Query user information in batches. The Contacts permission for the users and the departments to which the users belong is required. Up to 100 users can be queried each time. To invoke this API, the request header must specify the tenant_access_token.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/contact/v1/user/get?open_id=ou_337f846fd508d9b4b620478eced55d91&open_ids=2545643646&open_ids=4673565565  https://open.larksuite.com/open-apis/contact/v1/user/batch_get?employee_ids=1afbg844&employee_ids=e62e73eg |
| HTTP Method | GET |
| Required scopes |  Access contacts by application identity |
| Required field scopes Enable the scopes for fields you want to return |  Obtain user ID  Obtain user's basic information  Access contacts by application identity  Obtain user's email information  Obtain user's mobile number  Obtain user's gender  Obtain user's employment information  Obtain user's organization information | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value Format**: "Bearer `access_token`" **Example Value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" Learn more about obtaining and using access_token. |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Query parameters
|Parameter|Type|Mandatory|Description|
|--|-----|--|----|
|employee_ids、open_ids|string|Yes|Both employee_ids and open_ids are supported. At least one of the two must be passed in the request. If both are passed, employee_ids will be use| ## Response
### Response body
|Parameter|Description|
|-|-|
|code|Error code, anything but 0 indicates failure|
|msg|Response code description|
|data|Response business information|
|  ∟user_info|User information|
|    ∟name|User name|
|    ∟en_name|English name|
|    ∟employee_id|User employee ID. The field will be returned after apps have applied "Obtain user's userid" permission.|
|    ∟employee_no|User Work No.|
|    ∟open_id|User open ID|
|    ∟union_Id|User union ID|
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
|    ∟leader_employee_id|Employee ID of supervisor. Returned for custom apps, not returned for marketplace apps|
|    ∟leader_open_id|Open ID of supervisor|
|    ∟leader_union_id|Union ID of supervisor|
|    ∟departments|Departments of user|
|    ∟custom_attrs|Custom user properties set by the organization. Each property includes a custom property ID and custom value set for user. When organization open custom field and set custom value for user, this field will be returned |
###  Response body example
```json
{
    "code": 0,
    "msg": "success",
    "data": {
        "user_info":[
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
