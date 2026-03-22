---
title: "Obtain an App Admin’s Management Permissions"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uMzN3QjLzczN04yM3cDN"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/contact/v1/user/admin_scope/get"
section: "App Information"
updated: "1646720007000"
---

# Obtain app administrator management scope
This API is used to obtain the management scope of an app administrator, that is, which departments are managed by an app administrator.  

## Request
| HTTP URL | https://open.larksuite.com/open-apis/contact/v1/user/admin_scope/get |
| --- | --- |
| HTTP Method | GET |
| Required scopes | Obtain the user ID, management scope, and other information of the app administrators | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" Learn more about how to access and use access_token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Request parameters
|Parameter|Type|Required|Description|
|-|-|-|-|-|
|employee_id, open_id|string|Yes|Either open_id or employee_id can be used for query. The employee_id is equivalent to the user_id in Contacts v3.| ## Response
### Response body
|Parameter|Description|
|-|-|
|code|Return code. A value other than 0 indicates failure.|
|msg|Return code description|
|data|Returned business data|
|  ∟is_all|Whether the administrator manages all departments|
|  ∟department_list|List of managed departments. This field is not returned if is_all is true.| ### Response example
```json
{
    "code": 0,
    "msg": "success",
    "data": {
        "is_all":false,
        "departments_list":[
            "od-a140b4eeb892b90a0ab3e616fc2054d6",
            "od-a140b4eeb892b90a0ab3e616fc205477"
        ]
    }
}
```
