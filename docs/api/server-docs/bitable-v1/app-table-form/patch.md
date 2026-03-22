---
title: "Patch form"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/bitable-v1/app-table-form/patch"
method: "PATCH"
api_path: "https://open.larksuite.com/open-apis/bitable/v1/apps/:app_token/tables/:table_id/forms/:form_id"
service: "bitable-v1"
resource: "app-table-form"
section: "Docs"
rate_limit: "10 per second"
scopes:
  - "bitable:app"
updated: "1753668101000"
---

# Patch form

Update a form

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/bitable/v1/apps/:app_token/tables/:table_id/forms/:form_id |
| HTTP Method | PATCH |
| Rate Limit | 10 per second |
| Supported app types | custom,isv |
| Required scopes | `bitable:app` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `app_token` | `string` | Base unique device identifier app_token description **Example value**: "bascnv1jIEppJdTCn3jOosabcef" |
| `table_id` | `string` | table id **Example value**: "tblz8nadEUdxNMt5" |
| `form_id` | `string` | form id **Example value**: "vew6oMbAa4" | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `name` | `string` | No | Form name **Example value**: "Form" |
| `description` | `string` | No | Form description **Example value**: "Form description" |
| `shared` | `boolean` | No | Whether to enable sharing **Example value**: true |
| `shared_limit` | `string` | No | Share scope restrictions **Example value**: "tenant_editable" **Optional values are**:  - `off`: Only invited people can fill in - `tenant_editable`: People who get the link within the organization can fill in - `anyone_editable`: People who get the link on the Internet can fill in  |
| `submit_limit_once` | `boolean` | No | Fill in the number of times limit once **Example value**: true | ### Request body example

{
    "name": "Form",
    "description": "Form description",
    "shared": true,
    "shared_limit": "tenant_editable",
    "submit_limit_once": true
}

## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `form` | `app.table.form` | Form description |
|     `name` | `string` | Form name |
|     `description` | `string` | Form description |
|     `shared` | `boolean` | Whether to enable sharing |
|     `shared_url` | `string` | Share URL.  (The update API does not support this parameter) |
|     `shared_limit` | `string` | Share scope restrictions **Optional values are**:  - `off`: Only invited people can fill in - `tenant_editable`: People who get the link within the organization can fill in - `anyone_editable`: People who get the link on the Internet can fill in  |
|     `submit_limit_once` | `boolean` | Fill in the number of times limit once | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
        "form": {
            "name": "Form",
            "description": "Form description",
            "shared": true,
            "shared_url": "https://example.larksuite.com/share/base/shrcnCy1KAlpahNotmhRn1abcde",
            "shared_limit": "tenant_editable",
            "submit_limit_once": true
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
| 400 | 1254004 | WrongTableId | Table id wrong |
| 400 | 1254005 | WrongViewId | View id wrong |
| 400 | 1254006 | WrongRecordId | Record id wrong |
| 400 | 1254007 | EmptyValue | Empty value |
| 400 | 1254008 | EmptyView | Empty view |
| 400 | 1254009 | WrongFieldId | Wrong fieldId |
| 400 | 1254010 | ReqConvError | Request error |
| 400 | 1254016 | InvalidSort | invalid sort |
| 400 | 1254018 | InvalidFilter | invalid filter |
| 400 | 1254019 | InvalidViewType | Invalid view type |
| 400 | 1254020 | ViewNameDuplicated | Duplicate view name |
| 400 | 1254021 | EmptyViewName | View name is empty |
| 400 | 1254022 | InvalidViewName | Invalid view name |
| 400 | 1254030 | TooLargeResponse | TooLargeResponse |
| 404 | 1254040 | BaseTokenNotFound | AppToken not found |
| 404 | 1254041 | TableIdNotFound | Table not found |
| 404 | 1254042 | ViewIdNotFound | View not found |
| 404 | 1254043 | RecordIdNotFound | RecordIdNotFound |
| 404 | 1254044 | FieldIdNotFound | FieldIdNotFound |
| 404 | 1254045 | FieldNameNotFound | Field name does not exist |
| 404 | 1254049 | FormFieldNotFound | Form field id does not exist |
| 400 | 1254060 | TextFieldConvFail | TextFieldConvFail |
| 400 | 1254061 | NumberFieldConvFail | NumberFieldConvFail |
| 400 | 1254062 | SingleSelectFieldConvFail | SingleSelectFieldConvFail |
| 400 | 1254063 | MultiSelectFieldConvFail | MultiSelectFieldConvFail |
| 400 | 1254064 | DatetimeFieldConvFail | DatetimeFieldConvFail |
| 400 | 1254065 | CheckboxFieldConvFail | CheckboxFieldConvFail |
| 400 | 1254066 | UserFieldConvFail | The value corresponding to the personnel field type is incorrect. The possible reasons are: - The ID type specified by the user_id_type parameter does not match the type of the provided ID. - An unrecognized type or structure was provided. Currently, only `id` is supported, and it must be passed as an array. - An `open_id` was passed across applications. If you are passing an ID across applications, it is recommended to use `user_id`. The `open_id` obtained from different applications cannot be used interchangeably. |
| 400 | 1254067 | LinkFieldConvFail | LinkFieldConvFail |
| 400 | 1254100 | TableExceedLimit | TableExceedLimit, limited to 300 |
| 400 | 1254101 | ViewExceedLimit | ViewExceedLimit, limited to 200 |
| 400 | 1254103 | RecordExceedLimit | RecordExceedLimit, limited to 20,000 |
| 400 | 1254104 | RecordAddOnceExceedLimit | RecordAddOnceExceedLimit, limited to 500 |
| 400 | 1254130 | TooLargeCell | TooLargeCell |
| 429 | 1254290 | TooManyRequest | Request too fast, try again later |
| 400 | 1254291 | Write conflict | The same data table does not support concurrent calls to the write interface, please check whether there is a concurrent call to the write interface. The writing interface includes: adding, modifying, and deleting records; adding, modifying, and deleting fields; modifying forms; modifying views, etc. |
| 400 | 1254301 | OperationTypeError | Base does not have advanced permissions enabled or does not support enabling advanced permissions |
| 403 | 1254302 | Permission denied. | No access rights, usually caused by the table opening of advanced permissions, please add a group containing applications in the advanced permissions settings and give this group read and write permissions |
| 400 | 1255001 | InternalError | Internal error, have any questions can be consulting service |
| 400 | 1255002 | RpcError | Internal error, have any questions can be consulting service |
| 400 | 1255003 | MarshalError | Serialization failed, have any questions can be consulting service |
| 400 | 1255004 | UmMarshalError | Deserialization failed, have any questions can be consulting service |
| 400 | 1255005 | ConvError | Internal error, have any questions can be consulting service |
| 504 | 1255040 | Request timed out, please try again later | Try again |
| 400 | 1254036 | Base is copying, please try again later. | Base copy replicating, try again later |
| 500 | 1254200 | InternalError | Something went wrong. Please contact technical support at https://applink.larksuite.com/client/helpdesk/open?id=6626260912531570952 |
| 403 | 1254608 | Same API requests are submitted repeatedly. | Same API requests are submitted repeatedly. |
| 403 | 1254305 | PermControl | Requests may be restricted by permissions such as tenant switches, wiki space, etc. |
| 400 | 1254607 | Data not ready, please try again later | There are usually two situations when this error occurs: 1. The last submitted modification has not been processed; 2. The data is too large and the server calculation times out; This error code can be appropriately retried. |
| 403 | 1254304 | You are not authorized to perform this operation. | Advanced permissions for specific rows or columns are only available for Business and Enterprise editions |
