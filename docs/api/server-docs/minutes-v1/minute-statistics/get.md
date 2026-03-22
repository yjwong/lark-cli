---
title: "Get minutes statistics"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/minutes-v1/minute-statistics/get"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/minutes/v1/minutes/:minute_token/statistics"
service: "minutes-v1"
resource: "minute-statistics"
section: "Minutes"
rate_limit: "5 per second"
scopes:
  - "minutes:minutes"
  - "minutes:minutes:readonly"
  - "minutes:minutes.statistics:read"
field_scopes:
  - "contact:user.employee_id:readonly"
updated: "1750044550000"
---

# Get minutes statistics

Through this API, you can get access statistics of Lark Minutes, including PV, UV, visited user id, visited user timestamp.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/minutes/v1/minutes/:minute_token/statistics |
| HTTP Method | GET |
| Rate Limit | 5 per second |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `minutes:minutes` `minutes:minutes:readonly` `minutes:minutes.statistics:read` |
| Required field scopes | > The response body of the API contains the following sensitive fields, and they will be returned only after corresponding scopes are added. If you do not need the fields, it is not recommended that you request the scopes. `contact:user.employee_id:readonly` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `minute_token` | `string` | Minute uniquely identifies,it can be obtained from the minute link, usually the last string of characters in the link **Example value**: "obcnq3b9jl72l83w4f14xxxx" **Data validation rules**: - Length range: `24` ～ `24` characters | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `user_id_type` | `string` | No | User ID categories **Example value**: open_id **Optional values are**:  - `open_id`: Identifies a user to an app. The same user has different Open IDs in different apps. How to get Open ID - `union_id`: Identifies a user to a tenant that acts as a developer. A user has the same Union ID in apps developed by the same developer, and has different Union IDs in apps developed by different developers. A developer can use Union ID to link the same user's identities in multiple apps.How to get Union ID - `user_id`: Identifies a user to a tenant. The same user has different User IDs in different tenants. In one single tenant, a user has the same User ID in all apps （including store apps）. User ID is usually used to communicate user data between different apps. How to get User ID  **Default value**: `open_id` **When the value is `user_id`, the following field scopes are required**: `contact:user.employee_id:readonly` | ## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `statistics` | `statictics` | Minute browsing information statistics |
|     `user_view_count` | `string` | Number of user views |
|     `page_view_count` | `string` | Number of page views |
|     `user_view_list` | `user_view_detail[]` | User browse list |
|       `user_id` | `string` | User ID |
|       `view_time` | `string` | User's most recently viewed timestamp (ms level) | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
        "statistics": {
            "user_view_count": "3",
            "page_view_count": "20",
            "user_view_list": [
                {
                    "user_id": "ou_612b787ccd3259fb3c816b3f678dxxxx",
                    "view_time": "1669121332000"
                }
            ]
        }
    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 2091001 | Param is invalid | Check the parameters |
| 404 | 2091002 | Resource not found | Check minute token is valid |
| 400 | 2091003 | Minute not ready , try later | Check minute is transcribed |
| 400 | 2091004 | Resource deleted | Check minute is deleted |
| 403 | 2091005 | Permission deny | Check if you have read permission |
