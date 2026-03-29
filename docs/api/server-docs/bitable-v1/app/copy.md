---
title: "Copy App"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/bitable-v1/app/copy"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/bitable/v1/apps/:app_token/copy"
service: "bitable-v1"
resource: "app"
section: "Docs"
rate_limit: "20 per minute"
scopes:
  - "base:app:copy"
  - "bitable:app"
updated: "1773323331000"
---

# Copy App

Copy a base app, you can specify to copy to a folder with permissions

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/bitable/v1/apps/:app_token/copy |
| HTTP Method | POST |
| Rate Limit | 20 per minute |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `base:app:copy` `bitable:app` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `app_token` | `string` | Base App Token **Example value**: "S404b*****e9PQsYDWYcNryFn0g" | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `name` | `string` | No | Base App Name **Example value**: "test name" |
| `folder_token` | `string` | No | Base App Folder Token **Example value**: "fldbco*****CIMltVc" |
| `without_content` | `boolean` | No | Whether to copy the Base content, take the value: * true:   not copy * false:    copy **Example value**: false |
| `time_zone` | `string` | No | Timezone of Base **Example value**: "Asia/Shanghai" | ### Request body example

{
    "name": "test name",
    "folder_token": "fldbco*****CIMltVc",
    "without_content": false,
    "time_zone": "Asia/Shanghai"
}

## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `app` | `app` | Response data |
|     `app_token` | `string` | Base App Token |
|     `name` | `string` | Base App Name |
|     `folder_token` | `string` | Base App Folder Token |
|     `url` | `string` | Base App URL |
|     `default_table_id` | `string` | Default Table ID Of Base |
|     `time_zone` | `string` | Timezone of Base | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
        "app": {
            "app_token": "S404b*****e9PQsYDWYcNryFn0g",
            "name": "test name",
            "folder_token": "fldbco*****CIMltVc",
            "url": "https://by*****ce.larksuite.com/base/S404b*****e9PQsYDWYcNryFn0g",
            "default_table_id": "tbltGMo8AdY1BuAX",
            "time_zone": "Asia/Shanghai"
        }
    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 1254000 | WrongRequestJson | Request error |
| 400 | 1254001 | WrongRequestBody | Request body error |
| 400 | 1254002 | Fail | Internal error, have any questions can be consulting service |
| 400 | 1254003 | WrongBaseToken | AppToken error |
| 400 | 1254031 | InvalidAppName | App name is invalid, the length should not exceed 100 characters, and cannot contain ? / \ * : [ ] |
| 400 | 1254036 | Base is copying, please try again later. | Base copy replicating, try again later |
| 404 | 1254040 | BaseTokenNotFound | AppToken not found |
| 400 | 1254290 | TooManyRequest | Request too fast, try again later |
| 400 | 1254291 | Write conflict | The same data table does not support concurrent calls to the write interface, please check whether there is a concurrent call to the write interface. The writing interface includes: adding, modifying, and deleting records; adding, modifying, and deleting fields; modifying forms; modifying views, etc. |
| 403 | 1254304 | PermNotAllow | Advanced permissions for specific rows or columns are only available for Business and Enterprise editions |
| 403 | 1254701 | DriveNodePermNotAllow | Target folder has no permissions |
| 404 | 1254702 | DriveNodeNotExist | Target folder does not exist |
| 400 | 1254800 | ParamInvalid | param is invalid |
| 500 | 1255001 | InternalError | Internal error, have any questions can be consulting service |
| 500 | 1255002 | RpcError | Internal error, have any questions can be consulting service |
| 500 | 1255003 | MarshalError | Serialization failed, have any questions can be consulting service |
| 500 | 1255004 | UmMarshalError | Deserialization failed, have any questions can be consulting service |
| 500 | 1255005 | ConvError | Internal error, have any questions can be consulting service |
| 504 | 1255040 | Request timed out, please try again later | Try again |
| 400 | 1254292 | MountNodeOutOfSiblingNum | sibling num exceed the limit |
| 403 | 1254302 | RolePermNotAllow | no permission for the user |
