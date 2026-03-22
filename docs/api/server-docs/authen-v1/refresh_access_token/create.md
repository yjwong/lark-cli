---
title: "Refresh user_access_token (v1)"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/authen-v1/refresh_access_token/create"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/authen/v1/refresh_access_token"
service: "authen-v1"
resource: "refresh_access_token"
section: "Deprecated Version (Not Recommended)"
rate_limit: "Special Rate Limit"
field_scopes:
  - "contact:user.employee:readonly"
  - "contact:user.email:readonly"
  - "contact:user.employee_id:readonly"
  - "contact:user.phone:readonly"
updated: "1750924433000"
---

# Refresh user_access_token (v1)

> This API is now deprecated. Its use is not recommended. Please use the latest version: Refresh user_access_token.

The maximum valid period of user_access_token is about 2 hours. When the user_access_token expires, you can call this interface to get a new user_access_token.

> Please update the local user_access_token and refresh_token after refresh, do not continue to use the old value to refresh repeatedly. Make sure the parameter is the latest value

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/authen/v1/refresh_access_token |
| HTTP Method | POST |
| Rate Limit | Special Rate Limit |
| Supported app types | custom,isv |
| Required scopes | None |
| Required field scopes | > The response body of the API contains the following sensitive fields, and they will be returned only after corresponding scopes are added. If you do not need the fields, it is not recommended that you request the scopes. `contact:user.employee:readonly` `contact:user.email:readonly` `contact:user.employee_id:readonly` `contact:user.phone:readonly` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `app_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer a-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `grant_type` | `string` | Yes | The auth type. Fixed value. **Example value**: "refresh_token" |
| `refresh_token` | `string` | Yes | Refresh and fetch the user_access_token interface returns refresh_token, **For each request, please pay attention to using the latest acquired refresh_token** **Example value**: "ur-t9HHgRCsMqGqIU9vw5Zhof" | ### Request body example

{
    "grant_type": "refresh_token",
    "refresh_token": "ur-t9HHgRCsMqGqIU9vw5Zhof"
}

## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `user_access_token_info` | \- |
|   `access_token` | `string` | Field access_token is user_access_token for getting user resources and accessing certain open APIs |
|   `token_type` | `string` | token type |
|   `expires_in` | `int` | user_access_token valid period, unit: seconds, the valid time is about two hours, and the return result shall prevail |
|   `name` | `string` | User name |
|   `en_name` | `string` | User's English name |
|   `avatar_url` | `string` | User's profile photo |
|   `avatar_thumb` | `string` | User's profile photo in 72x72 pixels |
|   `avatar_middle` | `string` | User's profile photo in 240x240 pixels |
|   `avatar_big` | `string` | User's profile photo in 640x640 pixels |
|   `open_id` | `string` | User's UID in the app |
|   `union_id` | `string` | The unified user ID |
|   `email` | `string` | User's email **Required field scopes**: `contact:user.email:readonly` |
|   `enterprise_email` | `string` | Business email. Make sure that Lark Business Mail has been enabled in Admin. **Required field scopes**: `contact:user.employee:readonly` |
|   `user_id` | `string` | User's user_id **Required field scopes**: `contact:user.employee_id:readonly` |
|   `mobile` | `string` | User's mobile number **Required field scopes**: `contact:user.phone:readonly` |
|   `tenant_key` | `string` | The current company ID |
|   `refresh_expires_in` | `int` | refresh_token valid period, unit: seconds, generally about 30 days, the return result shall prevail |
|   `refresh_token` | `string` | The refresh_token used when refreshing user_access_token |
|   `sid` | `string` | Unique identifier of the user's current login session, not returned if empty | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
        "access_token": "u-Q7JWnaIM_kRChuLfreHmpArjOEayt.5XUBJcZr.V0Gst4FdQCtvrd9sAViLXQnQgkpL19brGOjKZQTxb",
        "token_type": "Bearer",
        "expires_in": 7140,
        "name": "zhangsan",
        "en_name": "Three Zhang",
        "avatar_url": "www.larksuite.com/avatar/icon",
        "avatar_thumb": "www.larksuite.com/avatar/icon_thumb",
        "avatar_middle": "www.larksuite.com/avatar/icon_middle",
        "avatar_big": "www.larksuite.com/avatar/icon_big",
        "open_id": "ou_caecc734c2e3328a62489fe0648c4b98779515d3",
        "union_id": "on_d89jhsdhjsajkda7828enjdj328ydhhw3u43yjhdj",
        "email": "zhangsan@larksuite.com",
        "enterprise_email": "demo@mail.com",
        "user_id": "5d9bdxxx",
        "mobile": "+86130002883xx",
        "tenant_key": "736588c92lxf175d",
        "refresh_expires_in": 2591940,
        "refresh_token": "ur-oQ0mMq6MCcueAv0pwx2fQQhxqv__CbLu6G8ySFwafeKww2Def2BJdOkW3.9gCFM.LBQgFri901QaqeuL",
		"sid": "AAAAAAAAAANjgHsqKEAAEw=="
    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 200 | 20001 | Invalid request. Please check request param | Please check the request parameters |
| 200 | 20007 | Failed to generate a user access token. Please try again | Please check if the parameter is valid and try again |
| 200 | 20008 | User not exist | User does not exist, change to a valid account |
| 200 | 20009 | Tenant does not install app | The tenant has not installed the app. Please contact the tenant administrator to install it |
| 200 | 20021 | User resigned | User leaves, please use a valid account |
| 200 | 20022 | User frozen | User frozen, please use a valid account |
| 200 | 20023 | User not registered | User is not registered, please use a valid account |
| 200 | 20024 | App id in user_access_token or refresh_token diff with app id in app_access_token or tenant_access_token. Please keep the app id consistent | Please check whether the app that generates two tokens is the same |
| 200 | 20025 | Lack of app_id or app_secret in request | Missing app_id or app_ secret, please check the parameters |
| 200 | 20028 | Invalid app id | Invalid app_id, please check parameters |
| 500 | 20050 | System error | System error, please try again |
