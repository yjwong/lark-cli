---
title: "Create App"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/bitable-v1/app/create"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/bitable/v1/apps"
service: "bitable-v1"
resource: "app"
section: "Docs"
rate_limit: "20 per minute"
scopes:
  - "base:app:create"
  - "bitable:app"
updated: "1773323330000"
---

# Create App

Create a base app in user-defined folder

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/bitable/v1/apps |
| HTTP Method | POST |
| Rate Limit | 20 per minute |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `base:app:create` `bitable:app` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `name` | `string` | No | Base App Name **Example value**: "test name" |
| `folder_token` | `string` | No | Base App Folder Token **Example value**: "fldbco*****CIMltVc" |
| `time_zone` | `string` | No | Base Timezone **Example value**: "Asia/Macau" | ### Request body example

{
    "name":"test name",
    "folder_token": "fldbco*****CIMltVc"
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
|     `default_table_id` | `string` | Default table id |
|     `time_zone` | `string` | Base Timezone | ### Response body example

{
    "code": 0,
    "data": {
        "app": {
            "app_token": "S404b*****e9PQsYDWYcNryFn0g",
            "default_table_id": "tblY2mIl0p2oumSQ",
            "folder_token": "fldbco*****CIMltVc",
            "name": "test name",
            "url": "https://by*****ce.larksuite.com/base/S404b*****e9PQsYDWYcNryFn0g"
        }
    },
    "msg": "success"
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 200 | 1254000 | WrongRequestJson | Request error |
| 200 | 1254001 | WrongRequestBody | Request body error |
| 200 | 1254002 | Fail | Internal error, have any questions can be consulting service |
| 200 | 1254003 | WrongBaseToken | AppToken error |
| 200 | 1254004 | WrongTableId | Table id wrong |
| 200 | 1254005 | WrongViewId | View id wrong |
| 200 | 1254006 | WrongRecordId | Record id wrong |
| 200 | 1254007 | EmptyValue | Empty value |
| 200 | 1254008 | EmptyView | Empty view |
| 200 | 1254009 | WrongFieldId | Wrong fieldId |
| 200 | 1254010 | ReqConvError | Request error |
| 200 | 1254025 | InvalidCopyTypes | InvalidCopyTypes |
| 200 | 1254030 | TooLargeResponse | TooLargeResponse |
| 400 | 1254031 | InvalidAppName | The name length should not exceed 255. |
| 400 | 1254036 | Base is copying, please try again later. | Base copy replicating, try again later |
| 200 | 1254040 | BaseTokenNotFound | AppToken not found |
| 200 | 1254041 | TableIdNotFound | Table not found |
| 200 | 1254042 | ViewIdNotFound | View not found |
| 200 | 1254043 | RecordIdNotFound | RecordIdNotFound |
| 200 | 1254044 | FieldIdNotFound | FieldIdNotFound |
| 200 | 1254060 | TextFieldConvFail | TextFieldConvFail |
| 200 | 1254061 | NumberFieldConvFail | NumberFieldConvFail |
| 200 | 1254062 | SingleSelectFieldConvFail | SingleSelectFieldConvFail |
| 200 | 1254063 | MultiSelectFieldConvFail | MultiSelectFieldConvFail |
| 200 | 1254064 | DatetimeFieldConvFail | DatetimeFieldConvFail |
| 200 | 1254065 | CheckboxFieldConvFail | CheckboxFieldConvFail |
| 200 | 1254066 | UserFieldConvFail | UserFieldConvFail |
| 200 | 1254067 | LinkFieldConvFail | LinkFieldConvFail |
| 200 | 1254100 | TableExceedLimit | TableExceedLimit, limited to 300 |
| 200 | 1254101 | ViewExceedLimit | ViewExceedLimit, limited to 200 |
| 200 | 1254102 | FileExceedLimit | FileExceedLimit |
| 200 | 1254103 | RecordExceedLimit | RecordExceedLimit, limited to 20,000 |
| 200 | 1254104 | RecordAddOnceExceedLimit | RecordAddOnceExceedLimit, limited to 500 |
| 200 | 1254130 | TooLargeCell | TooLargeCell |
| 200 | 1254290 | TooManyRequest | Request too fast, try again later |
| 200 | 1254291 | Write conflict | The same data table does not support concurrent calls to the write interface, please check whether there is a concurrent call to the write interface. The writing interface includes: adding, modifying, and deleting records; adding, modifying, and deleting fields; modifying forms; modifying views, etc. |
| 200 | 1254301 | OperationTypeError | Base does not have advanced permissions enabled or does not support enabling advanced permissions |
| 200 | 1254304 | PermNotAllow | Advanced permissions for specific rows or columns are only available for Business and Enterprise editions |
| 403 | 1254701 | DriveNodePermNotAllow | DriveNodePermNotAllow |
| 404 | 1254702 | DriveNodeNotExist | DriveNodeNotExist |
| 400 | 1254800 | ParamInvalid | param is invalid, please check |
| 200 | 1255001 | InternalError | Internal error, have any questions can be consulting service |
| 200 | 1255002 | RpcError | Internal error, have any questions can be consulting service |
| 200 | 1255003 | MarshalError | Serialization failed, have any questions can be consulting service |
| 200 | 1255004 | UmMarshalError | Deserialization failed, have any questions can be consulting service |
| 200 | 1255005 | ConvError | Internal error, have any questions can be consulting service |
| 504 | 1255040 | Request timed out, please try again later | Try again |
| 400 | 1254292 | MountNodeOutOfSiblingNum | sibling num exceed the limit |
