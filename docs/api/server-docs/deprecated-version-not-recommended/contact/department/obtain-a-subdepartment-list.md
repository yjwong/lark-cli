---
title: "Obtain a Sub-department List"
url: "https://open.larksuite.com/document/ukTMukTMukTM/ugzN3QjL4czN04CO3cDN"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/contact/v1/department/simple/list?department_id=0&page_size=10&fetch_child=false"
section: "Deprecated Version (Not Recommended)"
updated: "1646999090000"
---

# Obtain a sub-department list

This API is used to obtain a list of sub-departments under the current department.   
> To invoke this API, permission to access the Contacts of the current department must have been granted.  The company's root department ID is 0. When the specified department is the root department, the permission to access all of the organization's employees is required in order for the sub-department list to be returned.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/contact/v1/department/simple/list?department_id=0&page_size=10&fetch_child=false |
| HTTP Method | GET |
| Required scopes Enable any scope from the list | Access department's organize information Read contacts | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value Format**: "Bearer `access_token`" **Example Value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" Learn more about obtaining and using access_token. |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Query parameters
|Parameter|Type|Mandatory|Description|
|-|-|-|-|
|department_id|string|Yes|department ID|
|page_token|string|No|Page identifier. If not included in the first request, the list will iterate from the top. If a particular page is queried and there are more members in the department, a new page_token will also be returned. The next request can use that page_token to obtain more sub-departments|
|page_size|int|Yes|Page size, up to 100 items per page|
|fetch_child|bool|No|If sub-department users are recursively returned| ## Response
### Response body 

|Parameter|Description|
|-|-|
|code|Error code, anything but 0 indicates failure|
|msg|Response code description|
|data|Response business data|
|  ∟has_more|If department has more sub-department|
|  ∟page_token|Page identifier. When has_more is true, a new page_token will also be returned. Otherwise, page_token will not be returned|
|  ∟department_infos|sub-department list|
|    ∟id|department ID|
|    ∟name|department name|
|    ∟parent_id|parent department ID|
### Response body example
```json
{
    "code": 0,
    "msg": "success",
    "data": {
        "has_more":true,
        "page_token":"763bd1e74d05e95e",
        "department_infos": [
            {
                "id": "od-2efe30807a10608754862a63b108828f",
                "name": "Finance",
                "parent_id": "od-2efe30807a10608754862a63b1088266"
            },
            {
                "id": "od-846fa088e3dcac5d0ce0b4f6b4cde899",
                "name": "Human Resources",
                "parent_id": "0"
            }
        ]
    }
}
```
### Error code

For details, please refer to: Service-side error codes
