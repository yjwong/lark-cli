---
title: "Update version information"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/application-v6/application-app_version/patch"
method: "PATCH"
api_path: "https://open.larksuite.com/open-apis/application/v6/applications/:app_id/app_versions/:version_id"
service: "application-v6"
resource: "application-app_version"
section: "App Information"
scopes:
  - "application:application.app_version"
field_scopes:
  - "contact:user.employee_id:readonly"
updated: "1646274614000"
---

# Update version information

Updates the app version review status.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/application/v6/applications/:app_id/app_versions/:version_id |
| HTTP Method | PATCH |
| Supported app types | custom |
| Required scopes | `application:application.app_version` |
| Required field scopes | > The response body of the API contains the following sensitive fields, and they will be returned only after corresponding scopes are added. If you do not need the fields, it is not recommended that you request the scopes. `contact:user.employee_id:readonly` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" Learn more about obtaining and using access_token. |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `app_id` | `string` | App ID **Example value**: "cli_9f3ca975326b501b" |
| `version_id` | `string` | ID that uniquely identifies the app version **Example value**: "oav_d317f090b7258ad0372aa53963cda70d" | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `user_id_type` | `string` | Yes | User ID categories **Example value**: "open_id" **Optional values are**: - `open_id`: user's open id - `union_id`: user's union id - `user_id`: user's user id **Default value**: `open_id` **When the value is `user_id`, the following field scopes are required**: `contact:user.employee_id:readonly` |
| `operator_id` | `string` | Yes | Operator's open_id **Example value**: "ou_4065981088f8ef67a504ba8bd6b24d85" |
| `reject_reason` | `string` | No | This field is required when you want to change the version status to "Rejected". **Example value**: "Reason for rejection" | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `status` | `int` | No | Version status **Example value**: 1 **Optional values are**: - `0`: Unknown status - `1`: Approved - `2`: Rejected - `3`: In review - `4`: Not submitted for review | ### Request body example

```json
{
    "status": 1
}
```

## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- | ### Response body example

```json
{
    "code": 0,
    "msg": "success",
    "data": {

    }
}
```

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 210503 | invalid app_id | Check whether the app_id in the request path is valid. |
| 400 | 210504 | no such app in tenant | Check whether the queried app and the app used to call the API are in the same company. |
| 400 | 210505 | target app not a custom app | Check whether the queried app is a custom app. |
| 400 | 210506 | no such app | Check whether the app_id in the request path exists. |
| 400 | 210507 | no such user_id | Check whether the specified user ID exists. |
| 400 | 211002 | no such version_id | Check whether the version_id in the path is valid. |
| 400 | 211003 | no such version of desired app | Check whether the version_id matches the app specified by the app_id. |
| 400 | 211004 | no authority for quota limit | Check whether it is a Business/Enterprise tenant. |
| 400 | 211005 | invalid app id | Check the app ID. |
| 400 | 211006 | invalid department id | Check the department ID. |
