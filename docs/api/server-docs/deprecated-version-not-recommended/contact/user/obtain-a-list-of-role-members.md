---
title: "Obtain a List of Role Members"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uczMwUjL3MDM14yNzATN"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/contact/v2/role/members?role_id=or_846ea69995a259a27cc690182f27de87&page_size=2&page_token=763bd1e74d05e958"
section: "Deprecated Version (Not Recommended)"
updated: "1646999072000"
---

# Obtain a list of role members

This API is used to obtain a list of users in a given role. Apps invoking this API must have permission to read or write the organization’s Contacts. The result returned is a list of users in a given role within the app’s Contacts permission.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/contact/v2/role/members?role_id=or_846ea69995a259a27cc690182f27de87&page_size=2&page_token=763bd1e74d05e958 |
| HTTP Method | GET |
| Required scopes Enable any scope from the list | Access role information Read contacts | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value Format**: "Bearer `access_token`" **Example Value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" Learn more about obtaining and using access_token. |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Query parameters
|Parameter|Type|Mandatory|Description|
|-|-|-|-|
|role_id|string|yes|Role ID|
|page_token|string|no|Page identifier. If not included in the first request, the list will iterate from the top. If a particular page is queried and there are more members in the role, a new page_token will also be returned. The next request can use that page_token to obtain more members|
|page_size|int|no|Page size，up to 200 items per page; default 20 items| ## Response
### Response body
|Parameter|Description|
|-|-|
|code|Error code, anything but 0 indicates failure|
|msg|Response code description|
|data|Response business data|
|  ∟has_more|If there are more members. When has_more is true, a new page_token will also be returned|
|  ∟page_token|Page identifier. When has_more is true, a new page_token will also be returned. Otherwise, page_token will not be returned|
|  ∟user_list|User list|
|    ∟name|User name|
|    ∟open_id|User open_id|
|    ∟user_id|User ID. The field will be returned after apps have applied "Obtain user's userid" permission|
|    ∟scope|User's management scope within the role|
|      ∟is_all_department|If management scope covers all departments of the company, this parameter is only returned when true|
|      ∟department_ids|Custom department ID. If there is no is_all_department parameter, returns specific management department|
|      ∟open_department_ids|Department open_id. If there is no is_all_department parameter, returns specific management department|
### Response body example
```json
{
    "code": 0,
    "msg": "success",
    "data": {
        "has_more": true,
        "page_token": "763bd1e74d05e95e",
        "user_list": [
            {
                "name": "Jack",
                "open_id": "ou_84aad35d084aa403a838cf73ee144ec1",
                "user_id": "gbdfb31g"
                "scope": { 
                	"department_ids": ["TT-1234","TT-1235"],
                	"open_department_ids": ["od-c02cc3b685a711db3a0f14fc4cdb76dc","od-fa83d3690de01787618b85bb27a013bc"],
                }
            },
            {
                "name": "Hency",
                "open_id": "ou_6b77dbe8459863069c1a62f9cd810217",
                "user_id": "977ea122"
                "scope": {
                  "is_all_department": true
            	}
            }
        ]
    }
}
```
### Error code

For details, please refer to: Service-side error codes
