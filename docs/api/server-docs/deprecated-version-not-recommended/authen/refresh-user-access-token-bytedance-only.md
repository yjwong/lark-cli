---
title: "Refresh user_access_token (Bytedance only)"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/authentication-management/access-token/refresh-user_access_token-bytedance"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/authen/v1/oidc/refresh_access_token"
section: "Deprecated Version (Not Recommended)"
rate_limit: "1000 per minute、50 per second"
updated: "1750924428000"
---

# Refresh user_access_token (Bytedance only)

> This API is now deprecated. Its use is not recommended. Please use the latest version: Refresh user_access_token.

The maximum valid period of user_access_token is about 2 hours. When the user_access_token expires, you can call this interface to get a new user_access_token.

> - Starting from September 23, 2024, **newly created applications within Bytedance, by default, will not be able to obtain a refresh_token through the "Get user_access_token" interface.** To obtain refresh_token and refresh the access_token normally, you need to enable the "Refresh user_access_token" option in the application's security settings in Developer Console. Applications created before September 19 will not be affected and can still obtain and refresh tokens directly through the original interface.
> 
> - Please update the local user_access_token and refresh_token after refresh, do not continue to use the old value to refresh repeatedly. Make sure the parameter is the latest value

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/authen/v1/oidc/refresh_access_token |
| HTTP Method | POST |
| Rate Limit | 1000 per minute、50 per second |
| Supported app types | custom,isv |
| Required scopes | None | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `app_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer a-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `grant_type` | `string` | Yes | Authorization type, **fixed value**: **Example value**: "refresh_token" |
| `refresh_token` | `string` | Yes | Refresh and fetch the user_access_token interface returns refresh_token, **For each request, please pay attention to using the latest acquired refresh_token** **Example value**: "ur-h4_5nUXdJ4O8rqfGe.YJCwM13Gjc557xUG20hkk00f7K" | ### Request body example

{
    "grant_type": "refresh_token",
    "refresh_token": "ur-h4_5nUXdJ4O8rqfGe.YJCwM13Gjc557xUG20hkk00f7K"
}

## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `message` | `string` | Error descriptions |
| `data` | `token_info` | \- |
|   `access_token` | `string` | Field access_token is user_access_token for getting user resources and accessing certain open APIs |
|   `refresh_token` | `string` | The refresh_token used when refreshing user_access_token |
|   `token_type` | `string` | Token type, fixed value |
|   `expires_in` | `int` | user_access_token valid period, unit: seconds, the valid time is about two hours, and the return result shall prevail |
|   `refresh_expires_in` | `int` | refresh_token valid period, unit: seconds, generally about 30 days, the return result shall prevail |
|   `scope` | `string` | The complete set of permissions granted to the app by the user | ### Response body example

{
    "code": 0,
    "message": "success",
    "data": {
        "access_token": "u-5Dak9ZAxJ9tFUn8MaTD_BFM51FNdg5xzO0y010000HWb",
        "refresh_token": "ur-6EyFQZyplb9URrOx5NtT_HM53zrJg59HXwy040400G.e",
        "token_type": "Bearer",
        "expires_in": 7199,
        "refresh_expires_in": 2591999,
        "scope": "auth:user.id:read bitable:app"
    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 200 | 20001 | Invalid request. Please check request param | Please check the request parameters |
| 200 | 20002 | The app_id or app_secret passed is incorrect. Please check the value | Check app_id and key are correct |
| 200 | 20007 | Failed to generate a user access token. Please try again | Please check if the parameter is valid and try again |
| 200 | 20008 | User not exist | User does not exist, change to a valid account |
| 200 | 20013 | The tenant access token passed is invalid. Please check the value | Check if the tenant_access_token is valid |
| 200 | 20014 | The app access token passed is invalid. Please check the value | Check if the app_access_token is valid |
| 200 | 20021 | User resigned | User leaves, please use a valid account |
| 200 | 20022 | User frozen | User frozen, please use a valid account |
| 200 | 20023 | User not registered | User is not registered, please use a valid account |
| 200 | 20024 | App id in user_access_token or refresh_token diff with app id in app_access_token or tenant_access_token. Please keep the app id consistent | Please check whether the app that generates two tokens is the same |
| 200 | 20026 | The refresh token passed is invalid. Please check the value | Invalid refresh_token, please check if it has expired or has been consumed |
| 200 | 20028 | Invalid app id | Invalid app_id, please check parameters |
| 200 | 20029 | Invalid redirect uri | redirect_uri is invalid. Troubleshooting: 1. Make sure the value of Authorization is correct. 2. Make sure that the callback address redirect_uri parameter set when obtaining the login authorization code has been configured in the Developer Console > Application Details > Security Settings > Redirect URL. For detailed solutions to this error, see How to resolve the authorization page 20029 error. |
| 200 | 20036 | The grant_type passed is not supported | Invalid grant_type, please match the interface requirements |
| 200 | 20037 | The refresh token passed has expired. Please generate a new one | Expired refresh_token, please pass valid parameters |
| 200 | 20038 | The refresh token passed is invalid. Please check the value | Can't find refresh_token, please pass valid parameters |
| 200 | 20042 | App disabled | App is unavailable, please check app status |
| 200 | 20046 | Brand inconsistency | The application brand and the domain name brand are inconsistent. Please ensure that the Lark application is used under the Lark domain name, and lark is similar. |
