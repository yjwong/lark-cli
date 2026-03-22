---
title: "Obtain user_access_token (v1)"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/authen-v1/access_token/create"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/authen/v1/access_token"
service: "authen-v1"
resource: "access_token"
section: "Deprecated Version (Not Recommended)"
rate_limit: "Special Rate Limit"
field_scopes:
  - "contact:user.employee:readonly"
  - "contact:user.email:readonly"
  - "contact:user.employee_id:readonly"
  - "contact:user.phone:readonly"
updated: "1750924433000"
---

# Obtain user_access_token (v1)

> This API is now deprecated. Its use is not recommended. Please use the latest version: Get user_access_token.

Obtain `user_access_token` according to login pre-authorization code.

> This interface is used for web application login-free application scenarios. For the method of obtaining user_access_token for Mini Program applications, please refer to the code2session interface provided by Mini Program applications.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/authen/v1/access_token |
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
| `grant_type` | `string` | Yes | Authorization type, **fixed value**: **Example value**: "authorization_code" |
| `code` | `string` | Yes | Login pre-authorization code, call pre-authorization code interface to get **Example value**: "xMSldislSkdK" | ### Request body example

{
    "grant_type": "authorization_code",
    "code": "xMSldislSkdK"
}

## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `user_access_token_info` | \- |
|   `access_token` | `string` | user_access_token for obtaining user resources |
|   `token_type` | `string` | Token type |
|   `expires_in` | `int` | access_token valid period in seconds |
|   `name` | `string` | User name |
|   `en_name` | `string` | User English name |
|   `avatar_url` | `string` | user avatar |
|   `avatar_thumb` | `string` | User avatar 72x72 |
|   `avatar_middle` | `string` | User avatar 240x240 |
|   `avatar_big` | `string` | User avatar 640x640 |
|   `open_id` | `string` | Unique identity of the user within the app |
|   `union_id` | `string` | User Unified ID |
|   `email` | `string` | User mailbox **Required field scopes**: `contact:user.email:readonly` |
|   `enterprise_email` | `string` | Enterprise mailbox, please make sure that the Lark mailbox service is enabled in the management background first **Required field scopes**: `contact:user.employee:readonly` |
|   `user_id` | `string` | User user_id **Required field scopes**: `contact:user.employee_id:readonly` |
|   `mobile` | `string` | User's mobile phone number **Required field scopes**: `contact:user.phone:readonly` |
|   `tenant_key` | `string` | Current corporate identity |
|   `refresh_expires_in` | `int` | refresh_token valid period in seconds |
|   `refresh_token` | `string` | Token to use when refreshing user access_token |
|   `sid` | `string` | The unique identifier of the user's current login session, if it is empty, it will not be returned | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
        "access_token": "u-6U1SbDiM6XIH2DcTCPyeub",
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
        "enterprise_email": "demo@mail.com"
        "user_id": "5d9bdxxx",
        "mobile": "+86130002883xx",
        "tenant_key": "736588c92lxf175d",
        "refresh_expires_in": 2591940,
        "refresh_token": "ur-t9HHgRCjMqGqIU9v05Zhos",
        "sid": "AAAAAAAAAANjgHsqKEAAEw=="
    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 200 | 20001 | Invalid request. Please check request param | Please check the request parameters |
| 200 | 20007 | Failed to generate a user access token. Please try again | Generation failed, please check if the parameter is valid and try again |
| 200 | 20008 | User not exist | User does not exist, change to a valid account |
| 200 | 20009 | Tenant does not install app | The tenant does not have the app installed, please contact the tenant administrator to install it. |
| 200 | 20021 | User resigned | User leaves, please use a valid account |
| 200 | 20022 | User frozen | User frozen, please use a valid account |
| 200 | 20023 | User not registered | User is not registered, please use a valid account |
| 200 | 20024 | App id in user_access_token or refresh_token diff with app id in app_access_token or tenant_access_token. Please keep the app id consistent | Please check whether the app that generates two tokens is the same |
| 200 | 20025 | Lack of app_id or app_secret in request | Missing app_id or app_secret, please check parameters |
| 200 | 20028 | Invalid app id | Invalid app_id, please check parameters |
| 500 | 20050 | System error | System error, please try again |
