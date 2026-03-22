---
title: "Search App Admin List"
url: "https://open.larksuite.com/document/ukTMukTMukTM/ucDOwYjL3gDM24yN4AjN"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/user/v4/app_admin_user/list"
section: "App Information"
scopes:
  - "admin:app.admin_id:readonly"
  - "admin:app.admin:readonly"
field_scopes:
  - "contact:user.employee_id:readonly"
updated: "1646720020000"
---

# Query the list of app administrators
This API is used to query the list of administrators reviewing the app and returns the IDs of 10 latest administrator accounts.

## Request
| HTTP URL | https://open.larksuite.com/open-apis/user/v4/app_admin_user/list |
| --- | --- |
| HTTP Method | GET |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `admin:app.admin_id:readonly` `admin:app.admin:readonly` |
| Required field scopes | > The response body of the API contains the following sensitive fields, and they will be returned only after corresponding scopes are added. If you do not need the fields, it is not recommended that you request the scopes. `contact:user.employee_id:readonly` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" Learn more about how to obtain and use access_token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Request parameters
None

## Response

### Response body
|Parameter|Type|Description|
|-|-|-|
|code |int| Return code. A value other than 0 indicates failure.|
|msg  |string| Return code description|
|data | - | -|
|  ∟user_list|list| Administrator list|
|      ∟open_id|string| open_id of an administrator|
|      ∟user_id|string| user_id of an administrator|
|      ∟union_id|string| union_id of an administrator|               
###  Response example
```json
{
    "code": 0,
    "data": {
        "user_list": [
            {
                "open_id": "ou_a3b5bea08516cc528cd5546030344a",
                "union_id": "on_8ed6aa67826108097d9ee143819da2",
                "user_id": "e33ggbsd"
            },
            {
                "open_id": "ou_bb4dd8c1cbea4ec7bb7f8c556db7d4",
                "union_id": "on_a0c79d984a2dafaa1b91e97a13f183",
                "user_id": "738dfcsd"
            },
            {
                "open_id": "ou_cb02269798078d5d347f1014039aad",
                "union_id": "on_9a791f91a133e37ddc96537339db72",
                "user_id": "fab56ddf"
            },
            {
                "open_id": "ou_e050d74a4784b1e6e2403a1bff1ad7",
                "union_id": "on_06047a4315a8803e9d9e3a46268132",
                "user_id": "1ad57ckl"
            },
            {
                "open_id": "ou_a719f64bc45c930b204d813c580a2e",
                "union_id": "on_e6b7976ae9541fdb8763acfe0a6ce851",
                "user_id": "ffccc8cl"
            },
            {
                "open_id": "ou_23c7dd3b7a0666f54ce834d89afd34",
                "union_id": "on_51cfc6aecaa9441067eddf7e51b608f2",
                "user_id": "f4287g12"
            },
            {
                "open_id": "ou_ef7ee8e7b78394161f9e2be06f6a19",
                "union_id": "on_66941d406e790d275364863bdf376b",
                "user_id": "49bade34"
            },
            {
                "open_id": "ou_fa3d224bd3b9b963e6cad17946ee33",
                "union_id": "on_f219869448af8cbd27bed166878f08",
                "user_id": "3fcg5f12"
            },
            {
                "open_id": "ou_862ede29d66293c0016ebba1036e02",
                "union_id": "on_9b6d9de255d1c4555c40924a5759fc",
                "user_id": "4gge2115"
            }
        ]
    },
    "msg": "ok"
}
```
