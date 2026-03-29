---
title: "Obtain a Department User List"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uEzNz4SM3MjLxczM"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/contact/v1/department/user/list?department_id=od-8756c536552a91988b1b64559356c5a4&page_size=3&fetch_child=true"
section: "Deprecated Version (Not Recommended)"
updated: "1646999054000"
---

# Obtain a department user list
> In order to improve the security level, we have upgraded it. Please migrate to the new version as soon as possible. New version >>

This API is used to obtain a department user list.  Support for page-based and recursive queries is provided. To invoke this API, the request header must specify the tenant_access_token.

- The Contacts permission for this department must have been granted.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/contact/v1/department/user/list?department_id=od-8756c536552a91988b1b64559356c5a4&page_size=3&fetch_child=true |
| HTTP Method | GET |
| Required scopes Enable any scope from the list |  Access department's organize information  Read contacts |
| Required field scopes Enable the scopes for fields you want to return |  Obtain user's basic information  Obtain user's employment information | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value Format**: "Bearer `access_token`" **Example Value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" Learn more about obtaining and using access_token. |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Query parameters
|Parameter|Type|Mandatory|Description|
|--|-----|--|----|
|department_id|string|Yes|Department ID|
|page_token|string|No|Page identifier. If not included in the first request, the list will iterate from the top. If a particular page is queried and there are more members in the role, a new page_token will also be returned. The next request can use that page_token to obtain more members|
|page_size|int|Yes|Page size, up to 100 items per page|
|fetch_child|bool|Yes|Whether sub-department users are recursively returned. Within the application authorized scope, if the number of recursive departments exceeds 500, code 40162 is returned. You can call  Obtain a sub-department list  to query the first level sub-department, and then call this API to query the user list.|  
## Response 
### Response body 
|Parameter|Description|
|-|-|
|code|Error code, anything but 0 indicates failure|
|msg|Response code description|
|data|Response business information|
|  ∟has_more|If there are more members. When has_more is true, a new page_token will also be returned
|
|  ∟page_token|Page identifier. When has_more is true, a new page_token will also be returned. Otherwise, page_token will not be returned|
|  ∟user_list|User list|
|    ∟employee_id|User employee IDs. The field will be returned after apps have applied "Obtain user's userid" permission.|
|    ∟open_id|User open IDs|
|    ∟name|User name|
|    ∟employee_no|User Work.No|
|    ∟union_id|User union ID|
### Response body example
```json
{
    "code": 0,
    "msg": "success",
    "data": {
        "has_more": true,
        "page_token":"763bd1e74d05e95e",
        "user_list": [
            {
                "employee_id": "612a67ef",
                "open_id": "ou_e03053f0541cecc3269d7a9dc34a0b21",
                "name":"Jack",
                "employee_no":"264301",
                "union_id":"on_7dba11ff38a2119f89349876b12af65c"
            },
            {
                "employee_id": "a0615a67",
                "open_id": "ou_6dfd8d7e5e881bed9be19c043940bf60",
                "name":"Robert",
                "employee_no":"264302",
                "union_id":"on_c132837f686587dd494aa54f5f65b552"
            },
            {
                "employee_id": "2cb23fa3",
                "open_id": "ou_76a95e7f1b7841f74576d2fc00838f47",
                "name":"Hance",
                "employee_no":"264303",
                "union_id":"on_d12d180970dcc5e89f26929745166fb3"
            }
        ]
    }
}
```
### Error code

For details, please refer to: Service-side error codes
