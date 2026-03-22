---
title: "Query a User's App Access"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uATNwUjLwUDM14CM1ATN"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/pay/v1/paid_scope/check_user"
section: "App Information"
updated: "1646720023000"
---

#  Query whether a user can use paid functions

When a paid plan is charged per user count or has a maximum user limit, the open platform will guide the company administrator to set the "paid function applicable range". However, the applicable range will cause some users' inability to use the paid function. This API can be used to judge whether a user is allowed to use a paid function at the entry.

## Request
| HTTP URL | https://open.larksuite.com/open-apis/pay/v1/paid_scope/check_user |
| --- | --- |
| HTTP Method | GET | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" Learn more about how to access and use access_token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Request parameters

| Parameter         | Type           | Required        | Description         |
| --------- | --------------- | -------   | ----------- | 
|open_id | string | No |User's open_id. Either open_id or user_id is required. If both are specified, the open_id prevails.|
|user_id | string | No |User's user_id. Either user_id or open_id is required. If both are specified, the open_id prevails. | ## Response
### Response body
| Parameter         | Description           | 
|-|-| 
|code | Return code. A value other than 0 indicates failure. |
|msg | Return code description |
|data | Returned business information |
|  ∟status | Whether the user is in the applicable range. "valid" - The user is in the applicable range; "not_in_scope" - The user is not in the applicable range; "no_active_license" - The company has not purchased a price plan or the price plan has expired; "exceeds_maximum_limit"- The company's paid function applicable range exceeds the limit. Notify the administrator to adjust the configuration. |
|  ∟price_plan_id | ID of the tenant's current price plan, which matches the "Plan ID" in "Configure plan" in the Developer Console. |
|  ∟is_trial | Whether it is a trial version. true - Trial version; false - Not a trial version. |
|  ∟service_stop_time | Indicates the expiration date of the tenant's valid price plan. It is a Unix timestamp. |
### Response example
```json
{
    "code": 0, 
    "msg": "success",
    "data": {
        "status": "valid", 
        "price_plan_id": "price_9daf66c96968c003",
        "is_trial": false,
        "service_stop_time":"1572280518"
    }
} 
```
