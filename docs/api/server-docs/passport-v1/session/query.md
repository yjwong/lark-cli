---
title: "Obtain desensitized user login information in batches"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/passport-v1/session/query"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/passport/v1/sessions/query"
service: "passport-v1"
resource: "session"
section: "Authenticate and Authorize"
rate_limit: "10 per minute"
scopes:
  - "passport:session_mask:readonly"
field_scopes:
  - "contact:user.employee_id:readonly"
updated: "1744168060000"
---

# Obtain desensitized user login information in batches

This interface is used to query the user's login information.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/passport/v1/sessions/query |
| HTTP Method | POST |
| Rate Limit | 10 per minute |
| Supported app types | custom |
| Required scopes | `passport:session_mask:readonly` |
| Required field scopes | > The response body of the API contains the following sensitive fields, and they will be returned only after corresponding scopes are added. If you do not need the fields, it is not recommended that you request the scopes. `contact:user.employee_id:readonly` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `user_id_type` | `string` | No | User ID categories **Example value**: open_id **Optional values are**:  - `open_id`: Identifies a user to an app. The same user has different Open IDs in different apps. How to get Open ID - `union_id`: Identifies a user to a tenant that acts as a developer. A user has the same Union ID in apps developed by the same developer, and has different Union IDs in apps developed by different developers. A developer can use Union ID to link the same user's identities in multiple apps.How to get Union ID - `user_id`: Identifies a user to a tenant. The same user has different User IDs in different tenants. In one single tenant, a user has the same User ID in all apps （including store apps）. User ID is usually used to communicate user data between different apps. How to get User ID  **Default value**: `open_id` **When the value is `user_id`, the following field scopes are required**: `contact:user.employee_id:readonly` | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `user_ids` | `string[]` | No | User ID **Example value**: ["47f621ff"] **Data validation rules**: - Maximum length: `100` | ### Request body example

{"user_ids":["7aeddb6a1","7aeddb6a2","7aeddb6a3"]}

## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `mask_sessions` | `mask_session[]` | User Login Information |
|     `create_time` | `string` | Creation time |
|     `terminal_type` | `int` | Client side type **Optional values are**:  - `0`: Unknown - `1`: Personal computer - `2`: Browser - `3`: Android phone - `4`: Apple phone - `5`: Server - `6`: old version mini program - `8`: other mobile  |
|     `user_id` | `string` | User ID |
|     `sid` | `string` | The session identifier that needs to be logged out | ### Response body example

{
    "code":0,
    "data":{
        "mask_sessions":[
            {
                "create_time":"1644980493",
                "terminal_type":2,
                "user_id":"47f183f1f1",
                "sid":"ABBAAAAAAANll6nQoIAAFA"
            },
            {
                "create_time":"1644983127",
                "terminal_type":2,
                "user_id":"47f183ff1",
                "sid":"ACCAAAAAAANll6nQoIAAFA"
            },
            {
                "create_time":"1644983127",
                "terminal_type":2,
                "user_id":"47f183ff2",
                "sid":"BBBAAAAAAANll6nQoIAAFA=="
            }
        ]
    },
    "msg":""
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 1080001 | param is invalis | Invalid parameter |
