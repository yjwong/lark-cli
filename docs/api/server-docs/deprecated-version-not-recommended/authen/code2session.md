---
title: "code2session"
url: "https://open.larksuite.com/document/uYjL24iN/ukjM04SOyQjL5IDN"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/mina/v2/tokenLoginValidate"
section: "Deprecated Version (Not Recommended)"
updated: "1710751621000"
---

# code2session
After obtaining the login credential `code` through the login API, you can get the session_key and user information by sending a request to the server
> This API should be called on the server. For details, refer to How to call server APIs

##  Request
| HTTP URL | https://open.larksuite.com/open-apis/mina/v2/tokenLoginValidate |
| --- | --- |
| HTTP Method | POST |
| Supported app types | custom,isv |
| Required scopes | None | ###  Request header
| Name | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `app_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" Learn more about how to obtain and use access_token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Request body
Parameter | Type | Required | Description 
-- | -- | -- | -- 
code | string | Yes | The  code obtained at login

###  Example request body
```json
{
     "code": "2ef0bb04e272d274"
}
```

##  Response
###  Response body

| Name | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error code. A value other than 0 indicates failure |
| `msg` | `string` | Error description |
| `data` | `\-` | \- |
| ∟ `open_id` | `string` | Open ID of the user, which is used to identify the user in an app |
| ∟ `employee_id` | `string` | User ID of the user, which is the unique identifier of an in-service employee in the company **This field is returned only when the following permission is enabled**: `contact:user.employee_id:readonly` |
| ∟ `session_key` | `string` | Session key |
| ∟ `tenant_key` | `string` | Unique identifier of the user's company |
| ∟ `access_token` | `string` | user_access_token, the user's access token |
| ∟ `expires_in` | `int` | user_access_token expiration timestamp |
| ∟ `refresh_token` | `string` | The token used to refresh the user's access_token; it is valid for 30 days | ###  Example response body
```json
{
    "code": 0,
    "msg": "success",
    "data": {
    	"open_id": "ou_194fcfc5e4b78db556a040ff5e42c0",
    	"employee_id":"6c486g",
    	"union_id":"ou_be2c1742f2bb189469bdd33f0b1516",
    	"session_key": "e3aeb7df000c835365c630dac91bcf",
    	"tenant_key":"2c5914ac018f97",
    	"access_token":"u-tpwcnx2XzIcq8yHyJ6KL",
    	"expires_in":1565512680,
    	"refresh_token":"ur-W9dGvBJyVtwZmrwh0vBn"
    }
}
```

###  Error code
| HTTPStatus code | Error code | Description | Possible cause and suggestion |
| --- | --- | --- | --- |
| 200 | 10202 | access token invalid | Check whether  access_token  has expired |
| 200 | 10213 | code appid not match | 1. The app that is used to obtain the code must be the same one that calls the API. Check whether the API is called by another app 2. The code can only be used once. Check whether it was reused or expired |
| 200 | 10228 | user to app has no visibility | Invisible to the user | ## Known issues
The returned data will contain a union_id, which is obsolete and is not the same concept as the **union_id** commonly used in open platforms. Please do not use it; if necessary, it can be obtained by Get individual user information
