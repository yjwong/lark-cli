---
title: "Log out"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/passport-v1/session/logout"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/passport/v1/sessions/logout"
service: "passport-v1"
resource: "session"
section: "Authenticate and Authorize"
rate_limit: "Special Rate Limit"
scopes:
  - "passport:session:logout"
field_scopes:
  - "contact:user.employee_id:readonly"
updated: "1744168064000"
---

# Log Out

This interface is used to log out of the user login state.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/passport/v1/sessions/logout |
| HTTP Method | POST |
| Rate Limit | Special Rate Limit |
| Supported app types | custom |
| Required scopes | `passport:session:logout` |
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
| `idp_credential_id` | `string` | No | The unique identifier of the idp, required when logout_type = 2 **Example value**: "user@xxx.xx" |
| `logout_type` | `int` | Yes | Used to identify the type of logout **Example value**: 1 **Optional values are**:  - `1`: UserID - `2`: IdpCredentialID - `3`: Session uuid  |
| `terminal_type` | `int[]` | No | Logout terminal_type, default all logout. - 1：pc - 2：web - 3：android - 4：iOS - 5：server - 6：old version mini program - 8：other mobile **Example value**: [1] |
| `user_id` | `string` | No | User ID categories, is consistent with the query parameter user_id_type. required when logout_type = 1 **Example value**: "ou_7dab8a3d3cdcc9da365777c7ad535d62" |
| `logout_reason` | `int` | No | Logout prompt, optional. - 34：You have changed your login password, please log in again. - 35：Your login status has expired, please log in again. - 36：Your password has expired. Please use the Forgot Password function on the login page to change your password and log in again. **Example value**: 34 |
| `sid` | `string` | No | The session that needs to be logged out accurately, required when logout_type = 3 **Example value**: "AAAAAAAAAANll6nQoIAAFA" | ### Request body example

{
    "idp_credential_id": "user@xxx.xx",
    "logout_type": 1,
    "terminal_type": [
        1
    ],
    "user_id": "ou_7dab8a3d3cdcc9da365777c7ad535d62",
    "logout_reason": 34,
    "sid": "AAAAAAAAAANll6nQoIAAFA"
}

## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {}
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 1080001 | param is invalis | Parameter error, please modify the input parameters against the interface documentation and try again. |
| 400 | 1084001 | The parameter sid is invalid | The sid is invalid, please check and try again. |
| 400 | 1084002 | The parameter logout_reason is invalid | The logout_reason is invalid. Please refer to the "Request Body" description above and try again. |
