---
title: "Apply for scopes from the admin"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/application-v6/scope/apply"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/application/v6/scopes/apply"
service: "application-v6"
resource: "scope"
section: "App Information"
rate_limit: "Special Rate Limit"
updated: "1724640964000"
---

# Apply for scopes from the admin

Call this interface to apply to the tenant administrator for the API permissions that need to be reviewed within the application as an application.

> **Notice**: In the same tenant, other employees cannot apply for authorization from the administrator for more than 10 times for the same version of an application.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/application/v6/scopes/apply |
| HTTP Method | POST |
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
| `data` | `\-` | \- | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {}
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 212001 | only super scope unauthorized | The remaining unauthorized permissions of the application are highly sensitive permissions and cannot be applied for. |
| 400 | 212002 | unauthorized scopes were empty | The authorization scope is empty. Check whether the tenant has granted all the permissions required by the application. |
| 400 | 212003 | approval over limit | Approval limit exceeded. In the current tenant, other employees have applied for authorization from the administrator for more than 10 times for the same version of an application. No more applications can be processed. |
| 400 | 212004 | duplicate apply | Repeated application, operation failed. | For more error code information, see General error codes.
