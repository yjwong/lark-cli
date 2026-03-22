---
title: "Query tenant authorization status"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/application-v6/scope/list"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/application/v6/scopes"
service: "application-v6"
resource: "scope"
section: "App Information"
rate_limit: "Special Rate Limit"
updated: "1724640968000"
---

# Query tenant authorization status

Call this interface to query the status of the current app applying for authorization from the tenant.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/application/v6/scopes |
| HTTP Method | GET |
| Rate Limit | Special Rate Limit |
| Supported app types | custom,isv |
| Required scopes | None | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token | ## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `scopes` | `scope[]` | Status of the tenant granting permissions required by the current app. |
|     `scope_name` | `string` | Scope name. Example value: `user.phone:readonly`. |
|     `grant_status` | `int` | The scope granted status. **Optional values are**:  - `1`: Authorized - `2`: Unauthorized  | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
        "scopes": [
            {
                "scope_name": "user.phone:readonly",
                "grant_status": 1
            }
        ]
    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 212100 | app not install | The application is not installed. Check whether the current application has been installed in the enterprise. | For more error code information, see General error codes.
