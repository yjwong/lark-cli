---
title: "Get User Information"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/authen-v1/user_info/get"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/authen/v1/user_info"
service: "authen-v1"
resource: "user_info"
section: "Authenticate and Authorize"
rate_limit: "Special Rate Limit"
field_scopes:
  - "contact:user.employee:readonly"
  - "contact:user.employee_id:readonly"
  - "contact:user.phone:readonly"
  - "contact:user.email:readonly"
  - "contact:contact:access_as_app"
  - "contact:contact:readonly"
  - "contact:contact:readonly_as_app"
updated: "1749459204000"
---

# Get User Information

Get related user info via `user_access_token`.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/authen/v1/user_info |
| HTTP Method | GET |
| Rate Limit | Special Rate Limit |
| Supported app types | custom,isv |
| Required scopes | None |
| Required field scopes | > The response body of the API contains the following sensitive fields, and they will be returned only after corresponding scopes are added. If you do not need the fields, it is not recommended that you request the scopes. `contact:user.employee:readonly` `contact:user.employee_id:readonly` `contact:user.phone:readonly` `contact:user.email:readonly` `contact:contact:access_as_app` `contact:contact:readonly` `contact:contact:readonly_as_app` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token | ## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `user_info` | \- |
|   `name` | `string` | User name |
|   `en_name` | `string` | User's English name |
|   `avatar_url` | `string` | User's profile photo |
|   `avatar_thumb` | `string` | User's profile photo in 72x72 pixels |
|   `avatar_middle` | `string` | User's profile photo in 240x240 pixels |
|   `avatar_big` | `string` | User's profile photo in 640x640 pixels |
|   `open_id` | `string` | User's UID in the app |
|   `union_id` | `string` | User's UID to ISVs. For an ISV, all the apps that belong to the user have the same union_id. |
|   `email` | `string` | User's email **Required field scopes**: `contact:user.email:readonly` |
|   `enterprise_email` | `string` | Business email. Make sure that Lark Business Mail has been enabled in Admin. **Required field scopes**: `contact:user.employee:readonly` |
|   `user_id` | `string` | User's user_id **Required field scopes**: `contact:user.employee_id:readonly` |
|   `mobile` | `string` | User's mobile number **Required field scopes**: `contact:user.phone:readonly` |
|   `tenant_key` | `string` | The current company ID |
|   `employee_no` | `string` | User's employee_no **Required field scopes (Satisfy any)**: `contact:user.employee:readonly` `contact:contact:access_as_app` `contact:contact:readonly` `contact:contact:readonly_as_app` | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
        "name": "zhangsan",
        "en_name": "zhangsan",
        "avatar_url": "www.larksuite.com/avatar/icon",
        "avatar_thumb": "www.larksuite.com/avatar/icon_thumb",
        "avatar_middle": "www.larksuite.com/avatar/icon_middle",
        "avatar_big": "www.larksuite.com/avatar/icon_big",
        "open_id": "ou-caecc734c2e3328a62489fe0648c4b98779515d3",
        "union_id": "on-d89jhsdhjsajkda7828enjdj328ydhhw3u43yjhdj",
        "email": "zhangsan@larksuite.com",
        "enterprise_email": "demo@mail.com",
        "user_id": "5d9bdxxx",
        "mobile": "+86130002883xx",
        "tenant_key": "736588c92lxf175d",
		"employee_no": "111222333"
    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 200 | 20001 | Invalid request. Please check request param | Invalid request. Please check your request parameters. |
| 200 | 20005 | The user access token passed is invalid. Please check the value | Given `user_access_token` is invalid, please use a valid `user_access-token`. Related docs: Get user_access_token |
| 200 | 20008 | User not exist | The user doesn't exist, you can't get information about the user |
| 200 | 20021 | User resigned | The user is resigned. Please check user status. |
| 200 | 20022 | User frozen | The user's account is frozen. Please check user status. |
| 200 | 20023 | User not registered | The user is not registered. Please check user status. |
| 500 | 20050 | System error | Internal server error. Please try again later. |
