---
title: "Obtain the Contacts Permission Scope"
url: "https://open.larksuite.com/document/ukTMukTMukTM/ugjNz4CO2MjL4YzM"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/contact/v1/scope/get"
section: "Deprecated Version (Not Recommended)"
scopes:
  - "contact:contact:access_as_app"
  - "contact:contact:readonly"
updated: "1646999117000"
---

# Obtain the Contacts permission scope
This API is used to obtain the scope of the permission the app has been granted for Contacts data, including accessing department lists and user lists. 

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/contact/v1/scope/get |
| HTTP Method | GET |
| Required scopes | `contact:contact:access_as_app` `contact:contact:readonly` |
| Required field scopes Enable the scopes for fields you want to return | Obtain user ID | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value Format**: "Bearer `access_token`" **Example Value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" Learn more about obtaining and using access_token. |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ## Response
### Response body 
|Parameter|Description|
|-|-|
|code|Error code, anything but 0 indicates failure|
|msg|Response code description|
|data|Response business data|
|  ∟authed_departments|Authorized Contacts department lists. If permission is granted for the entire staff, it will be the top-level department ID list|
|  ∟authed_employee_ids|Authorized Contacts user employee_id lists. The field will be returned after apps have applied "Obtain user's userid" permission. If permission is granted for the entire staff, it will be the top-level user employee ID list|
|  ∟authed_open_ids|Authorized Contacts user open ID list.If permission is granted for the entire staff, it will be the top-level user open ID list|
### Response body example
```json
{
    "code": 0,
    "msg": "success",
    "data": {
        "authed_departments": [
            "od-8756c536552a91988b1b64559356c5a4",
            "od-a140b4eeb892b90a0ab3e616fc2054d6"
        ],
        "authed_employee_ids": [
            "c248363e",
            "98a42b48"
        ],
        "authed_open_ids": [
            "ou_2ed6636cb8bdbbee6ee20f7992c10258",
            "ou_89ad6bf1fbf5a39796b556ae06a10d4d"
        ]
    }
}
```
### Error code

For details, please refer to: Service-side error codes
