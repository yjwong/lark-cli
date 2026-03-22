---
title: "Obtain Department Details"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uAzNz4CM3MjLwczM"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/contact/v1/department/info/get?department_id=od-8756c536552a91988b1b64559356c5a4"
section: "Deprecated Version (Not Recommended)"
updated: "1646999084000"
---

# Obtain department details
> In order to improve the security level, we have upgraded it. Please migrate to the new version as soon as possible. New version >>

This API is used to obtain department information. 

- The application should apply for Department Data Permissions. To invoke this API, the request header must specify the tenant_access_token. Marketplace apps do not have permission to invoke this API.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/contact/v1/department/info/get?department_id=od-8756c536552a91988b1b64559356c5a4 |
| HTTP Method | GET |
| Required scopes | Access contacts by application identity |
| Required field scopes Enable the scopes for fields you want to return | Obtain user ID Obtain department's basic information Obtain department's organizational structure | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value Format**: "Bearer `access_token`" **Example Value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" Learn more about obtaining and using access_token. |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Query parameters
|Parameter|Type|Mandatory|Description|
|--|-----|--|----|
|department_id|string|Yes|Department ID| ## Response
### Response body

|Parameter|Type|Description|
|-|-|-
|code|int|Error code, anything but 0 indicates failure|
|msg|string|Response code description|
|data|-|Response business information|
|  ∟department_info|-|Department information|
|    ∟id|string|Department ID|
|    ∟i18n_name|-|Internationalized department names|
|      ∟en_us|-|Department name in English|
|      ∟ja_jp|-|Department name in Japanese|
|      ∟zh_cn|-|Department name in Chinese|
|    ∟open_department_id|string|Department openID|
|    ∟name|string|Department name|
|    ∟chat_id|string|Department Chat Group ID|
|    ∟member_count|int|Department member count|
|    ∟parent_id|string|Parent department ID|
|    ∟status|int|Department status. 0 indicates inactive, 1 indicates active|
|    ∟leader_employee_id|string|Employee ID of department head. The field will be returned after apps have applied "Obtain user's userid" permission| 
|    ∟leader_open_id|string|Open ID of department head| 
### Response body example
```json
{
    "code": 0,
    "msg": "success",
    "data": {
        "department_info": {
            "id":"od-8756c536552a91988b1b64559356c5a4",
            "leader_employee_id":"612a67ef",
            "leader_open_id":"ou_05065996251935ada9c2b0ecc50be91e",
            "chat_id": "oc_405333f8fc89c3262865b014ccbbb274",
            "member_count": 79,
            "name": "business",
            "parent_id": "0",
            "status": 1
        }
    }
}
```
### Error code

For details, please refer to: Service-side error codes
