---
title: "Obtain a Sub-department ID List"
url: "https://open.larksuite.com/document/ukTMukTMukTM/ukjNz4SO2MjL5YzM"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/contact/v1/department/list?department_id=od-8756c536552a91988b1b64559356c5a4"
section: "Deprecated Version (Not Recommended)"
updated: "1646999087000"
---

# Obtain a sub-department ID list

This API is used to obtain sub-department ID list. It only returns department lists for which permission has been granted. The ID of the root department is 0. When the root department is specified, all authorized department lists will be returned. When a particular department is specified, and the Contacts permission has been granted for this department, all sub-departments under that department will be returned. To invoke this API, the request header must specify the tenant_access_token.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/contact/v1/department/list?department_id=od-8756c536552a91988b1b64559356c5a4 |
| HTTP Method | GET |
| Required scopes Enable any scope from the list | Access department's organize information Read contacts | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value Format**: "Bearer `access_token`" **Example Value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" Learn more about obtaining and using access_token. |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Query parameters
**Description of request parameters**:    
|Parameter|Type|Mandatory|Description|
|--|-----|--|----|
|department_id|string|Yes|Department ID| ## Response
### Response body
|Parameter|Description|
|-|-|
|code|Error code, anything but 0 indicates failure|
|msg|Response code description|
|data|Response business data|
|  ∟department_list|Department list|
### Response body example
```json
{
    "code": 0,
    "msg": "success",
    "data": {
        "departments_list": [
            "od-a140b4eeb892b90a0ab3e616fc2054d6",
            "od-c02cc3b685a711db3a0f14fc4cdb76dc",
            "od-5a921e75b6adfb9f607bbaa40e50ba24",
            "od-fa83d3690de01787618b85bb27a013bc"
        ]
    }
}
```
### Error code

For details, please refer to: Service-side error codes
