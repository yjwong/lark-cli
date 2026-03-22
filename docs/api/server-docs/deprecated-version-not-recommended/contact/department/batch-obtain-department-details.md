---
title: "Batch Obtain Department Details"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uczN3QjL3czN04yN3cDN"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/contact/v1/department/detail/batch_get?department_ids=od-2efe30807a10608754862a63b108828f&department_ids=od-da6427b2adbceb91204d7fa6aeb7e8ffL"
section: "Deprecated Version (Not Recommended)"
updated: "1646999093000"
---

# Batch obtain department details
> In order to improve the security level, we have upgraded it. Please migrate to the new version as soon as possible. New version >>

This API is used to batch obtain department details, only department authorithed would be returned。

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/contact/v1/department/detail/batch_get?department_ids=od-2efe30807a10608754862a63b108828f&department_ids=od-da6427b2adbceb91204d7fa6aeb7e8ffL |
| HTTP Method | GET |
| Required scopes Enable any scope from the list | access contacts by application identity Obtain department's basic information Obtain department's organizational structure |
| Required field scopes Enable the scopes for fields you want to return | Obtain user ID | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value Format**: "Bearer `access_token`" **Example Value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" Learn more about obtaining and using access_token. |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Query parameters
|Parameter|Type|Mandatory|Description|
|-|-|-|-|
|department_ids|string|yes|Department ID| ## Response
### Response body
|Parameter|Description|
|-|-|
|code|Error code, anything but 0 indicates failure|
|msg|Response code description|
|data|Response business information|
|  ∟department_infos|Department informations|
|    ∟chat_id|Department group chat ID|
|    ∟id|Department ID|
|    ∟name|Department Name|
|    ∟member_count|Department member count|
|    ∟parent_id|Parent department ID|
|    ∟status|Department status. 0 indicates inactive, 1 indicates active|
|    ∟leader_employee_id|Employee ID of department head. The field will be returned after apps have applied "Obtain user's userid" permission| 
|    ∟leader_open_id|Open ID of department head| 
### Response body example
```json
{
    "code": 0,
    "msg": "success",
    "data": {
        "department_infos": [
            {
                "chat_id": "oc_2a287650fc2928376b0f2a92c2a5431f",
                "id": "od-2091e298d8dffd6026b1670b90ca7f54",
                "leader_employee_id": "ed4g28gg",
                "leader_open_id": "ou_782168c229ccd7e4509ed09db13ba5d1",
                "member_count": 2,
                "name": "Finance",
                "parent_id": "od-2091e298d8dffd6026b1670b90ca7f66",
                "status": 1
            },
            {
                "chat_id": "oc_517b250144476ceeb8d325d5a86c8056",
                "id": "od-2efe30807a10608754862a63b108828f",
                "member_count": 11,
                "name": "Human Resources",
                "parent_id": "0",
                "status": 1
            }
        ]
    }
}
```
### Error code

For details, please refer to: Service-side error codes
