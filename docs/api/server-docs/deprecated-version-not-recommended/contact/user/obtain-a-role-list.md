---
title: "Obtain a Role List"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uYzMwUjL2MDM14iNzATN"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/contact/v2/role/list"
section: "Deprecated Version (Not Recommended)"
updated: "1646999069000"
---

# Obtain a role list

This API is used to obtain a list of user roles in your organization. Apps invoking this API must have permission to read or update the current organization’s Contacts.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/contact/v2/role/list |
| HTTP Method | GET |
| Required scopes Enable any scope from the list | Access role information Read contacts | ### Request header
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
|  ∟role_list|Role list|
|    ∟id|Role ID|
|    ∟name|Role name|
### Response body example 
```json
{
    "code": 0,
    "msg": "success",
    "data": {
        "role_list": [
            {
                "id": "or_846ea69995a259a27cc690182f27de87",
                "name": "IT"
            },
            {
                "id": "or_b337b102aff4167bd29189b1d94b77f6",
                "name": "HR"
            }
        ]
    }
}
```
### Error code

For details, please refer to: Service-side error codes
