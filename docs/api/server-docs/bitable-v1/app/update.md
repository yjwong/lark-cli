---
title: "Update App Name"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/bitable-v1/app/update"
method: "PUT"
api_path: "https://open.larksuite.com/open-apis/bitable/v1/apps/:app_token"
service: "bitable-v1"
resource: "app"
section: "Docs"
rate_limit: "10 per second"
scopes:
  - "bitable:app"
updated: "1773323333000"
---

# Update App Information

Update app information according to app_token

> - Advanced permissions are not supported for Base in doc, sheet and wiki
> - This API is not an atomic operation. It firstly modifies name, and then switches advanced permissions. It may be partial success

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/bitable/v1/apps/:app_token |
| HTTP Method | PUT |
| Rate Limit | 10 per second |
| Supported app types | custom,isv |
| Required scopes | `bitable:app` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `app_token` | `string` | base app token **Example value**: "appbcbWCzen6D8dezhoCH2RpMAh" | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `name` | `string` | No | Base App Name **Example value**: "new name" |
| `is_advanced` | `boolean` | No | Is advanced **Example value**: false | You can modify the name or switch advanced permissions separately, and the parameters that are not filled in the request body will not be affected

### Request body example

```json
{
    "name": "new name",
    "is_advanced": false
}
```

## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `app` | `display_app_v2` | App information |
|     `app_token` | `string` | Base App Token |
|     `name` | `string` | Base App Name |
|     `is_advanced` | `boolean` | Advanced permissions on or off |
|     `time_zone` | `string` | Base App Time Zone | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
        "app": {
            "app_token": "appbcbWCzen6D8dezhoCH2RpMAh",
            "name": "new name",
            "is_advanced": true,
            "time_zone":"Asia/Beijing"
        }
    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 200 | 1254000 | WrongRequestJson | Request error |
| 200 | 1254001 | WrongRequestBody | Request body error |
| 200 | 1254002 | Fail | Internal error, have any questions can be consulting service |
| 200 | 1254003 | WrongBaseToken | AppToken error |
| 200 | 1254010 | ReqConvError | Request error |
| 200 | 1254031 | InvalidAppName | App name is invalid, the length should not exceed 100 characters, and cannot contain ? / \ * : [ ] |
| 400 | 1254036 | Base is copying, please try again later. | Base copy replicating, try again later |
| 200 | 1254040 | BaseTokenNotFound | AppToken not found |
| 200 | 1254043 | RecordIdNotFound | RecordIdNotFound |
| 200 | 1254200 | internal error | Internal error |
| 200 | 1254290 | TooManyRequest | Request too fast, try again later |
| 200 | 1254291 | Write conflict | The same data table does not support concurrent calls to the write interface, please check whether there is a concurrent call to the write interface. The writing interface includes: adding, modifying, and deleting records; adding, modifying, and deleting fields; modifying forms; modifying views, etc. |
| 400 | 1254301 | OperationTypeError | Base does not have advanced permissions enabled or does not support enabling advanced permissions |
| 403 | 1254302 | Permission denied. | No access rights, usually caused by the table opening of advanced permissions, please add a group containing applications in the advanced permissions settings and give this group read and write permissions |
| 403 | 1254304 | The role has no permissions. | Advanced permissions for specific rows or columns are only available for Business and Enterprise editions |
| 200 | 1255001 | InternalError | Internal error, have any questions can be consulting service |
| 200 | 1255002 | RpcError | Internal error, have any questions can be consulting service |
| 200 | 1255003 | MarshalError | Serialization failed, have any questions can be consulting service |
| 200 | 1255004 | UmMarshalError | Deserialization failed, have any questions can be consulting service |
| 504 | 1255040 | Request timed out, please try again later | Try again | ## Other error code

**Error code** | **Cause**  | **Suggestion**          |
| ------- | ------- | ----------------- |
| 1254061 | The field's format is incorrect. | Make sure the parameter's format is correct.
