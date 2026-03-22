---
title: "Get authorization code"
url: "https://open.larksuite.com/document/common-capabilities/sso/api/obtain-oauth-code"
method: "GET"
api_path: "https://accounts.larksuite.com/open-apis/authen/v1/authorize"
resource: "common-capabilities"
section: "Authenticate and Authorize"
updated: "1749459208000"
---

# Get authorization code

This API is used to initiate user authorization. After the user consents to the authorization, the application will receive an authorization code, referred to as `code`. Please note that the authorization code is valid for 5 minutes and can only be used once.

## Notes

- This interface is actually an authorization page, suitable for the authorization scenario of web applications. When user authorization is required, the application should redirect the user to this authorization page. After the user clicks authorize on the authorization page (when the web application is opened in the Lark client, it can directly jump without confirmation), the browser will redirect to the address specified by `redirect_uri` and carry the `code` query parameter (i.e., authorization code).
- Developers can use the authorization code to obtain a `user_access_token` to call the OpenAPI (e.g. Get User Information). For detailed steps on obtaining a `user_access_token`, refer to Get user_access_token.
- By using this endpoint along with Get user_access_token and Get user information, the application can implement Lark OAuth login.
- For a complete implementation of the user authorization workflow, refer to the Web App End User Athorization Guide.
- Note that when opening the authorization page, you need to declare the user authorization scopes required by the application by appending the `scope` query parameter. For example, the scope key for obtaining basic contact information is `contact:contact.base:readonly`.
- The permissions granted by the user to the application accumulate over time, and the most recently generated `user_access_token` will include all the permissions the user has granted in the past.
- When using the `user_access_token` to call a specific OpenAPI, ensure that this `user_access_token` possesses the required permissions for the target OpenAPI; otherwise, the call will fail.

## Request
| HTTP URL | https://accounts.larksuite.com/open-apis/authen/v1/authorize |
| --- | --- |
| HTTP Method | GET |
| API Rate Limit | 1000 times/minute, 50 times/second |
| Supported Application Types | custom,isv |
| Permission Requirements | None | ### Query Parameters

> To ensure the correct construction and encoding of URLs, it is recommended to use relevant URL standard libraries for URL parsing and building, rather than manual concatenation.

| Name | Type | Required | Description |
| --- | --- | --- | --- |
| `client_id` | `string` | Yes | The App ID of the application, which can be found on the **Credentials and Basic Information** page in the developer console. For a detailed introduction to the App ID, please refer to General Parameters. **Example Value:** `cli_a5d611352af9d00b` |
| `redirect_uri` | `string` | Yes | The application redirect address, to which the user will be redirected upon successful authorization. The address will include the `code` and, if provided, the `state` parameter. Please note: 1. This address must be URL encoded; 2. Before calling this API, you need to configure the HTTP GET request API address that accepts OAuth callbacks as the application's redirect URL on the **Security Settings** page in the developer console. Multiple redirect URLs can be configured, and only those listed will pass the security verification of the open platform. For details, refer to Configure redirect URL. **Example Value:** `https://example.com/api/oauth/callback` |
| `scope` | `string` | No | An incremental grant of permissions needed by the application. **Format Requirement:** The `scope` parameter is a space-separated, case-sensitive string. **Note**: - Developers need to apply for the necessary `scope` for calling the Open API in the [Developer Console](https://open.larksuite.com/app) based on the business scenario and manually concatenate the `scope` parameter. If the corresponding permissions are not applied for in the console, the user will encounter a 20027 error when using the application. - Up to 50 `scope` permissions can be requested from the user at once. For details, see API permission list. - If you need to obtain a `refresh_token` later, you must add the `offline_access` scope here (see Refresh user_access_token for details): `offline_access` **Example Value:** `contact:contact bitable:app:readonly` |
| `state` | `string` | No | An additional string to maintain the state between request and callback. It will be returned as is in the callback upon completion of authorization. The application can use this string to determine context relations and prevent CSRF attacks. Ensure that the `state` parameter is consistent before and after. **Example Value:** `RANDOMSTRING` |
| `code_challenge` | `string` | No | Used to enhance the security of the authorization code through the PKCE (Proof Key for Code Exchange) flow. **Example Value:** `E9Melhoa2OwvFrEMTJguCHaoeK1t8URWbuGJSstw-cM` For detailed information on PKCE, please refer to [RFC-7636 - Proof Key for Code Exchange by OAuth Public Clients](https://datatracker.ietf.org/doc/html/rfc7636). |
| `code_challenge_method` | `string` | No | The method used to generate the `code_challenge`. **Optional Values**: 1. **`S256`** (Recommended): The SHA-256 hash algorithm is used to calculate the hash of the `code_verifier`, and the result is Base64URL encoded to generate the `code_challenge`. 2. **`plain`** (Default Value): Directly use `code_verifier` as `code_challenge` without additional processing. The above `code_verifier` refers to the randomly generated string created locally before initiating the authorization. | ### Request Example

> Note: This is only a sample request URL. Please replace the query parameters with actual values as described earlier.
> 
``` http
https://accounts.larksuite.com/open-apis/authen/v1/authorize?client_id=cli_a5d611352af9d00b&redirect_uri=https%3A%2F%2Fexample.com%2Fapi%2Foauth%2Fcallback&scope=bitable:app:readonly%20contact:contact&state=RANDOMSTRING
```

## Response

### Successful Response
Once the user consents to authorization, the browser will redirect to the specified `redirect_uri` address provided at the initiation of the authorization, accompanied by the `code` and `state` parameters.

| Name | Description |
| --- | --- |
| `code` | Authorization code, used to obtain the user_access_token. **Character Set:** [A-Z] / [a-z] / [0-9] / "-" / "_" **Length:** Developers should reserve at least 64 characters **Example Value:** `2Wd5g337vo5BZXUz-3W5KECsWUmIzJ_FJ1eFD59fD1AJIibIZljTu3OLK-HP_UI1` |
| `state` | The original value of the `state` parameter passed in when opening the authorization page. If not provided, it will not return here. | Example:
```http
https://example.com/api/oauth/callback?code=2Wd5g337vo5BZXUz-3W5KECsWUmIzJ_FJ1eFD59fD1AJIibIZljTu3OLK-HP_UI1&state=RANDOMSTRING
```

### Failure Response
When a user denies authorization, the browser will redirect to the `redirect_uri` given during the initial authorization request, along with the `error` and `state` query parameters. The current fixed value for the `error` parameter is `access_denied`. Be sure to handle scenarios where authorization is denied appropriately.

| Name | Description |
| --- | --- |
| `error` | Error message, currently fixed as `access_denied` |
| `state` | The original value of the `state` parameter passed during the opening of the authorization page. If not provided initially, it will not be returned here. | Example:
```http
https://example.com/api/oauth/callback?error=access_denied&state=RANDOMSTRING
```

## Frequently Asked Questions

### Error 20027 During User Authorization

**Issue:** Users encounter error 20027 when authorizing an application.

**Cause:** The issue arises when the concatenated scope parameter on the authorization page includes permissions that the current application has not applied.

**Solution:**

1. Confirm scopes that the user needs to authorize.
2. Go to the [Developer Console](https://open.larksuite.com/app), and apply the corresponding permissions on the **Development Configuration** > **Permission Management** > **API Permissions** page for the relevant application. For detailed steps, refer to Request API Permissions.
3. Call the API with scopes that the application has applied.

### How to Obtain a user_access_token with Desired Permissions

When calling an API, if the `user_access_token` lacks the required permissions, the following error will be returned:

```json
{
  "code": 99991679,
  "error": {
    "log_id": "202407260711088FB107A76E0100002087",
    "permission_violations": [
      {
        "subject": "task:task:read",
        "type": "action_privilege_required"
      },
      {
        "subject": "task:task:write",
        "type": "action_privilege_required"
      }
    ]
  },
  "msg": "Unauthorized. You do not have permission to perform the requested operation on the resource. Please request user re-authorization and try again. required one of these privileges: [task:task:read, task:task:write]"
}
```

To prevent API call failures due to insufficient `user_access_token` permissions, developers can request users to grant the necessary permissions using the `scope` query parameter. There are two methods to achieve this:

1. Concatenate all the necessary `scope` values for user authorization at once, eliminating the need for repeated authorization if no new permissions are required. Note that there is a limit of 50 `scope` values per concatenation.
2. Alternatively, based on the error code returned by the API call and the `permission_violations` field, identify the additional permissions required for the current operation. Then, regenerate the authorization link to prompt the user for supplementary authorization, and use the newly generated `user_access_token` to continue calling the API.

Developers are advised to follow the principle of least privilege, requesting only the necessary permissions from users.

### Description of redirect_uri with \#

According to the standard [RFC 3986 - Uniform Resource Identifier (URI): Generic Syntax](https://datatracker.ietf.org/doc/html/rfc3986#section-3), the content following a `#` in a URI is referred to as the fragment, which is positioned at the end of the URI. If the business authorization request parameter `redirect_uri` is concatenated with a `#`, the redirection after authorization will append the `#` and its fragment content to the end of the URI. Special attention is required when parsing to retrieve the `code`.

Example of `redirect_uri`:

``` 
https://example.com/api/oauth/callback/#/login
```

Request example:
``` 
GET https://accounts.larksuite.com/open-apis/authen/v1/authorize?client_id=cli_a5d611352af9d00b&redirect_uri=https%3A%2F%2Fexample.com%2Fapi%2Foauth%2Fcallback%2F%23%2Flogin%0A&scope=bitable:app:readonly%20contact:contact&state=RANDOMSTRING
```

Example of the values in the browser address bar after the callback:
```shell 
https://example.com/api/oauth/callback?code=2Wd5g337vo5BZXUz-3W5KECsWUmIzJ_FJ1eFD59fD1AJIibIZljTu3OLK-HP_UI1&state=RANDOMSTRING#/login
```
