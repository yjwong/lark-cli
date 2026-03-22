---
title: "Refresh user_access_token"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/authentication-management/access-token/refresh-user-access-token"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/authen/v2/oauth/token"
section: "Authenticate and Authorize"
updated: "1749459216000"
---

# Refresh user_access_token

OAuth token API, which can be used to refresh the user_access_token and obtain a new refresh_token.

- The `user_access_token` is a user's access credential that allows calling OpenAPI on behalf of the user. This credential has an expiration period and can be refreshed using the `refresh_token`.
- During user authorization, the user must have access to the application. Otherwise, calling this API will return error code 20010.
- The `refresh_token` is used to obtain a new `user_access_token` and can only be used once. When obtaining a new `user_access_token`, a new `refresh_token` will be returned, and the original `refresh_token` will become invalid immediately.

- For the method of obtaining the `refresh_token` for the first time, see get user_access_token.

> This API implementation adheres to [RFC 6749 - The OAuth 2.0 Authorization Framework](https://datatracker.ietf.org/doc/html/rfc6749), and you can use a [standard OAuth client library](https://oauth.net/code/) for integration (**recommended**)

## Prerequisites

### Enable offline_access Permission
To obtain a `refresh_token`, navigate to the **Permissions & Scopes** module in **Developer Console** to enable the `offline_access` permission. When initiating authorization, declare this permission in the `scope` parameter.

![Enable offline_access Permission.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/25e512721aaa6a6360211111896eed67_dB51yJmTWT.png?height=679&lazyload=true&width=1862)

After enabling the `offline_access` permission, the request parameters for obtaining a `refresh_token` should be configured as follows:

1. When initiating authorization (Get Authorization Code, the `scope` parameter in the authorization link must include `offline_access`. For example:  

```
https://accounts.larksuite.com/open-apis/authen/v1/authorize?client_id=cli_a5d611352af9d00b&redirect_uri=https%3A%2F%2Fexample.com%2Fapi%2Foauth%2Fcallback&scope=bitable:app:readonly%20offline_access
```
2. When obtaining the `user_access_token` (Get User Access Token):  
   - If scope reduction is not required (i.e., the `scope` parameter for this interface is empty), no additional configuration is needed, and the `refresh_token` will be provided automatically.  
   - If scope reduction is required (i.e., the `scope` parameter for this interface is not empty):  
     - To obtain a `refresh_token`, the `scope` parameter must include `offline_access`.  
     - If a `refresh_token` is not required, no special configuration is needed.

3. When refreshing the `user_access_token` (Refresh User Access Token), follow the same logic as in Step 2.

### Enabling the Security Setting for Refreshing user_access_token
> - If you do not see this switch, there is no need for concern, as it is enabled by default.
> - After completing the configuration, you need to publish the application to make the configuration effective. For specific operations, refer to Publishing Enterprise Self-Built Applications and Publishing Store Applications.

Navigate to the **Security Settings** module in **Developer Console**, and enable the switch for refreshing the `user_access_token`. 

![Enable the security setting for refreshing user_access_token.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/a8a5013bcf75c358ddc245124baefb2e_AWKTe9rVza.png?height=702&lazyload=true&width=1835)

## Request

> To prevent misuse of the `user_access_token` refresh mechanism, 365 days after a user authorizes the application, the application must prompt the user for re-authorization to obtain a new `user_access_token` and `refresh_token`. Continuing to refresh the `user_access_token` after the `refresh_token` has expired will result in an error (error code is 20037). You can refer to the following [error code descriptions](#错误码) for further information.

> After refreshing, please update the local `user_access_token` and `refresh_token`, as the original tokens will no longer be usable. The `user_access_token` has a one-minute grace period to allow the application to complete token rotation.

| HTTP URL | https://open.larksuite.com/open-apis/authen/v2/oauth/token |
| --- | --- |
| HTTP Method | POST |
| API Rate Limit | 1000 times/minute, 50 times/second |
| Supported Application Types | custom,isv |
| Permission Requirements | None |
| Field Permission Requirements | The `refresh_token` and `refresh_token_expires_in` fields are only returned when the following permissions are granted: `offline_access` | ### Request Headers

| Name | Type | Required | Description |
| --- | --- | --- | --- |
| Content-Type | string | Yes | Type of request body. **Fixed value:** `application/json; charset=utf-8` | ### Request Body

| Name | Type | Required | Description |
| --- | --- | --- | --- |
| `grant_type` | `string` | Yes | Authorization type. **Fixed value:** `refresh_token` |
| `client_id` | `string` | Yes | App ID of the application, which can be found on the **Credentials & Basic Info** page in the **Developer Console**. **Example Value:** `cli_a5ca35a685b0x26e` |
| `client_secret` | `string` | Yes | App Secret of the application, which can be found on the **Credentials & Basic Info** page in the **Developer Console**. For more details, see: [How to Obtain App ID](https://ssl:ttdoc/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-app-id). **Example Value:** `baBqE5um9LbFGDy3X7LcfxQX1sqpXlwy` |
| `refresh_token` | `string` | Yes | Refresh token used to refresh `user_access_token` and `refresh_token`. > Please note that this API only supports `refresh_token` returned by the Get user_access_token and Refresh user_access_token. **Example value:** `eyJhbGciOiJFUzI1NiIs**********XXOYOZz1mfgIYHwM8ZJA` |
| `scope` | `string` | No | This parameter is used to reduce the permission scope of `user_access_token`. For example: 1. When obtaining an authorization code, you can authorize three permissions using the `scope` parameter: `contact:user.base:readonly contact:contact.base:readonly contact:user.employee:readonly`. 2. In the current API, you can pass `contact:user.base:readonly` through the `scope` parameter to reduce the permissions of `user_access_token` to only `contact:user.base:readonly`. **Note:** - If this parameter is not specified, the generated `user_access_token` will include all the permissions authorized by the user. - This parameter must not contain duplicate permissions; otherwise, the API call will result in an error with error code 20067. - This parameter must not contain unauthorized permissions (i.e., permissions that were not within the scope authorized by the user when obtaining the authorization code); otherwise, the API call will result in an error with error code 20068. - Multiple calls to this API to reduce the scope of permissions will not stack. For example, if the user has granted permissions A and B, the first call to the API reduces the scope to permission A, then `user_access_token` will only contain permission A. If the second call to the API reduces the scope to permission B, then `user_access_token` will only contain permission B. - The effective permission list can be viewed via the `scope` return value of this API. **Format requirement:** Space-separated list of `scope` **Example value:** `auth:user.id:read task:task:read` | ### Example Request Body
```json
{
    "grant_type": "refresh_token",
    "client_id": "cli_a5ca35a685b0x26e",
    "client_secret": "baBqE5um9LbFGDy3X7LcfxQX1sqpXlwy",
    "refresh_token": "eyJhbGciOiJFUzI1NiIs**********XXOYOZz1mfgIYHwM8ZJA"
}
```

## Response
The response body type is `application/json; charset=utf-8`.

### Response Body
> **The length of `access_token` and `refresh_token` in the response body is considerably long**, typically between 1~2KB. However, the length may further increase due to a larger number of `scope` or subsequent changes. It is recommended to reserve 4KB of storage space.

| Name | Type | Description |
| --- | --- | --- |
| `code` | `string` | Error code, where 0 indicates a successful request, and any non-zero value indicates failure. Please refer to the error code section below for proper handling. |
| `access_token` | `string` | Equivalent to `user_access_token`, returned only when the request is successful. |
| `expires_in` | `int` | The validity period of the `user_access_token` in seconds, returned only when the request is successful. > It is recommended to use this field to determine the expiration time of the `user_access_token`, rather than hard-coding the validity period. |
| `refresh_token` | `string` | Used to refresh the `user_access_token`, this field is only returned when the request is successful and the user has granted `offline_access` permission: `offline_access` Note that if you set the `scope` request parameter during the refresh of `user_access_token` and need to return `refresh_token`, you must include `offline_access` in the `scope` parameter. |
| `refresh_token_expires_in` | `int` | The validity period of the `refresh_token` in seconds, returned only when `refresh_token` is returned. > It is recommended to call the current interface again before expiration to obtain a new `refresh_token`. |
| `token_type` | `string` | The value is fixed as `Bearer`, returned only when the request is successful. |
| `scope` | `string` | A list of permissions granted by the `access_token` for this request, separated by spaces, returned only when the request is successful. |
| `error` | `string` | The error type, returned only when the request fails. |
| `error_description` | `string` | Specific error information, returned only when the request fails. | ### Response Body Example

Successful Response Example:

```json
{
    "code": "0",
    "access_token": "eyJhbGciOiJFUzI1NiIs**********X6wrZHYKDxJkWwhdkrYg",
    "expires_in": 7200, // Variable value; please ensure to determine the validity period of the access_token based on the actual value returned in the response body
    "refresh_token": "eyJhbGciOiJFUzI1NiIs**********VXOYOZYZmfgIYHWM0ZJA",
    "refresh_token_expires_in": 604800, // Variable value; please ensure to determine the validity period of the refresh_token based on the actual value returned in the response body
    "scope": "auth:user.id:read offline_access task:task:read user_profile",
    "token_type": "Bearer"
}
```

Failed Response Example:
```json
{
    "code": "20050",
    "error": "server_error",
    "error_description": "An unexpected server error occurred. Please retry your request."
}
```

### Error Codes
| HTTP Status Code | Error Code | Description | Troubleshooting Advice |
| --- | --- | --- | --- |
| 400 | 20001 | The request is missing a required parameter. | Required parameter missing. Please check if the parameters in the request are correct. |
| 400 | 20002 | The client secret is invalid. | Application authentication failed. Please verify if the provided `client_id` and `client_secret` are correct. For the method of obtaining, refer to How to obtain the App ID of an application. |
| 400 | 20008 | The user does not exist. | User does not exist. Please check the current status of the user initiating the authorization. |
| 400 | 20009 | The specified app is not installed. | The tenant has not installed the app. Please check the app status. |
| 400 | 20010 | The user does not have permission to use this app. | User lacks permission to use the app. Please verify if the user initiating the authorization still has app usage rights. |
| 400 | 20024 | The provided authorization code or refresh token does not match the provided client ID. | The provided `refresh_token` does not match the `client_id`. Please do not mix credentials from different applications. |
| 400 | 20026 | The refresh token passed is invalid. Please check the value. | Please check the value of the `refresh_token` field in the request body. Please note that this API only supports `refresh_token` issued in v2 version API. |
| 400 | 20036 | The specified `grant_type` is not supported. | Invalid `grant_type`. Please check the value of the `grant_type` field in the request body. |
| 400 | 20037 | The refresh token passed has expired. Please generate a new one. | `refresh_token` has expired. Please re-initiate the authorization process to obtain a new `refresh_token`. |
| 400 | 20048 | The specified app does not exist. | The app does not exist. Please check the app status. |
| 500 | 20050 | An unexpected server error occurred. Please retry your request. | Internal server error. Please try again later and contact [technical support](https://applink.larksuite.com/TLJpeNdW) if the issue persists. |
| 400 | 20063 | The request is malformed. Please check your request. | Missing required fields in the request body. Please fill in the fields based on the specific error information. |
| 400 | 20064 | The refresh token has been revoked. Please note that a refresh token can only be used once. | `refresh_token` has been revoked. Please re-initiate the authorization process to obtain a new `refresh_token`. |
| 400 | 20066 | The user status is invalid. | Invalid user status. Please check the current status of the user initiating the authorization. |
| 400 | 20067 | The provided scope list contains duplicate scopes. Please ensure all scopes are unique. | Invalid `scope` list, duplicates exist. Please ensure the provided `scope` list contains no duplicates. |
| 400 | 20068 | The provided scope list contains scopes that are not permitted. Please ensure all scopes are allowed. | Invalid `scope` list; it contains permissions that the user has not authorized. The permissions passed in the `scope` parameter of the current API must be a subset of the `scope` parameter value when obtaining the authorization code. For example, if the user authorized permissions A and B when obtaining the authorization code, the `scope` value passed in the current API can only be permissions A or B. If permission C is passed, the current error code will be returned. |
| 400 | 20069 | The specified app is not enabled. | The app is not enabled. Please check the app status. |
| 400 | 20070 | Multiple authentication methods were provided. Please only use one to proceed. | Both `Basic Authentication` and `client_secret` authentication methods were used in the request. Please use only the `client_id` and `client_secret` authentication methods to call this API. |
| 503 | 20072 | The server is temporarily unavailable. Please retry your request. | The service is temporarily unavailable. Please try again later. |
| 400 | 20073 | The refresh token has been used. Please note that a refresh token can only be used once. | Please re-initiate the authorization process to obtain a new `refresh_token`. |
| 400 | 20074 | The specified app is not allowed to refresh token. | Please check in the app management console to see if the 'refresh `user_access_token`' option is enabled. Note that this takes effect after publishing. |
