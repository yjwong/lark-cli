---
title: "Update Department Information"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uczNz4yN3MjL3czM"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/contact/v1/department/update"
section: "Deprecated Version (Not Recommended)"
updated: "1646999102000"
---

# Update department information
> In order to improve the security level, we have upgraded it. Please migrate to the new version as soon as possible. New version >>

This API is used to update department information in Contacts. 

- The application should apply for `update contacts` and `access contacts by application identity` permissions to invoke this API. Contacts permission for this department and other related departments is required. 
To invoke this API, the request header must specify the tenant_access_token. 
Marketplace apps do not have permission to invoke this API.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/contact/v1/department/update |
| HTTP Method | POST |
| Required scopes Enable any scope from the list | update contacts access contacts by application identity | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value Format**: "Bearer `access_token`" **Example Value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" Learn more about obtaining and using access_token. |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Request body   
|Parameter|Type|Mandatory|Description|
|--|-----|--|----|
|name|String|No|Department name|
|parent_id|String|No|Parent department ID|
|id|String|Yes|Department ID|
|leader_open_id|String|No|Line manager open ID|
|leader_employee_id|String|No|Line manager employee ID|
|create_group_chat|Bool|No|If need create group chat|
### Request body example
```json
{
    "name":"business",
    "parent_id":od-455efa262dc736b3e45a8b17fe945293,
    "id":"od-45aecb1f629c735b3e45cd017fe945124",
    "leader_employee_id":"2fab234c",
    "leader_open_id":"ou_4a2eb24a52b27c0b7fc6fd04162c0246",
    "create_group_chat":true
}
```

## Response 
### Response body
|Parameter|Description|
|-|-|
|code|Error code, anything but 0 indicates failure|
|msg|Response code description|
|data|Response business information|
|  ∟department_info    |Department information|
|    ∟id|Department group chat ID|
|    ∟id|Department ID|
|    ∟parent_id|Parent department ID|
|    ∟status|Department status|
|    ∟leader_employee_id|Employee ID of department head, The field will be returned after apps have applied "Obtain user's userid" permission|
|    ∟leader_open_id|Open ID of department head|
### Response body example
```json
{
    "code": 0,
    "msg": "success",
    "data": {
        "department_info":{
            "chat_id": "oc_a3c1a99d5736e3a269ef7302eb4763ab",
            "id":"od-3c03f7714260801568e82064f4423887",
            "parent_id":"od-c042a4980ba8e1466050e3e8da2378fe",
            "name":"business",
            "status":1,
            "leader_employee_id":"2fe234cb",
            "leader_open_id":"ou_e03053f0541cecc3269d7a9dc34a0b21"
        }
    }
}
```
### Error code

For details, please refer to: Service-side error codes
