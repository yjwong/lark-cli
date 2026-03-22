---
title: "Get User IDs Using Union IDs"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uUTO5UjL1kTO14SN5kTN"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/user/v1/union_id/batch_get/list?union_ids=on_94a1ee5551019f18cd73d9f111898cf2&union_ids=on_42f2ef9d07319a4d96fffd7ef5cbfc79"
section: "Deprecated Version (Not Recommended)"
updated: "1646999078000"
---

# Get user IDs using union IDs
Query a user’s open_id and user_id using the user’s union ID. Batch queries are supported.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/user/v1/union_id/batch_get/list?union_ids=on_94a1ee5551019f18cd73d9f111898cf2&union_ids=on_42f2ef9d07319a4d96fffd7ef5cbfc79 |
| HTTP Method | GET |
| Required field scopes Enable the scopes for fields you want to return | Obtain user ID | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value Format**: "Bearer `access_token`" **Example Value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" Learn more about obtaining and using access_token. |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Query parameters
Parameter | Type | Mandatory | Description
-- | -- | -- | --
union_ids | string | Yes | The union IDs of the users to be queried, max. 100
## Response
### Response body
Parameter | Description
-- | --
code | Error code, anything but 0 indicates failure
msg | Response code description
data | -
  ∟user_infos | The users found based on the union IDs. The key is the union ID and the value is the users found with that address. Pay attention that multiple users may be found for the same union ID
    ∟open_id | User's open_id
    ∟user_id | User's user_id，returned when the employee_id permission has been granted
### Response body example
```json
{
    "code": 0,
    "msg": "",
    "data": {
        "user_infos": {
            "on_7dba11ff38a2119f89349876b12af65c": {
                "user_id": "b6e596g8",
                "open_id": "ou_ef74363752064ca799c2bb7b6bdca9f4"
            },
            "on_c132837f686587dd494aa54f5f65b552": {
                "user_id": "384cf662",
                "open_id": "ou_47e071d2656df0df08a49d87e0c7f518"
            },
            "on_d12d180970dcc5e89f26929745166fb3": {
                "user_id": "3cc7eaf6",
                "open_id": "ou_c1ba95fa08f067545a99c8dba5bf588e"
            }
        }
    }
}
```
### Error code

For details, please refer to: Service-side error codes
