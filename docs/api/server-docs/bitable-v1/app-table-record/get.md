---
title: "Get records"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/bitable-v1/app-table-record/get"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/bitable/v1/apps/:app_token/tables/:table_id/records/:record_id"
service: "bitable-v1"
resource: "app-table-record"
section: "Deprecated Version (Not Recommended)"
rate_limit: "20 per second"
scopes:
  - "bitable:app"
  - "bitable:app:readonly"
field_scopes:
  - "contact:user.employee_id:readonly"
  - "contact:user.base:readonly"
  - "contact:contact:access_as_app"
  - "contact:contact:readonly"
  - "contact:contact:readonly_as_app"
updated: "1752650888000"
---

# Get records

Get records

> This interface is a historical version interface and is no longer recommended for use. You can use Batch Record Retrieval.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/bitable/v1/apps/:app_token/tables/:table_id/records/:record_id |
| HTTP Method | GET |
| Rate Limit | 20 per second |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `bitable:app` `bitable:app:readonly` |
| Required field scopes | > The response body of the API contains the following sensitive fields, and they will be returned only after corresponding scopes are added. If you do not need the fields, it is not recommended that you request the scopes. `contact:user.employee_id:readonly` `contact:user.base:readonly` `contact:contact:access_as_app` `contact:contact:readonly` `contact:contact:readonly_as_app` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token | ::: note
The instructions for AccessToken calling Docs API are detailed here Docs API Overview

### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `app_token` | `string` | base app_token **Example value**: "bascnd0HM3KAyiZJELxfMHRrGZc" |
| `table_id` | `string` | table id **Example value**: "tblEGB3HKvDrpj71" |
| `record_id` | `string` | record id **Example value**: "recP750ZNJ" | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `text_field_as_array` | `boolean` | No | indicate the structure of value of text field in response **Example value**: true |
| `user_id_type` | `string` | No | User ID categories **Example value**: open_id **Optional values are**:  - `open_id`: Identifies a user to an app. The same user has different Open IDs in different apps. How to get Open ID - `union_id`: Identifies a user to a tenant that acts as a developer. A user has the same Union ID in apps developed by the same developer, and has different Union IDs in apps developed by different developers. A developer can use Union ID to link the same user's identities in multiple apps.How to get Union ID - `user_id`: Identifies a user to a tenant. The same user has different User IDs in different tenants. In one single tenant, a user has the same User ID in all apps （including store apps）. User ID is usually used to communicate user data between different apps. How to get User ID  **Default value**: `open_id` **When the value is `user_id`, the following field scopes are required**: `contact:user.employee_id:readonly` |
| `display_formula_ref` | `boolean` | No | Control whether formulas, lookup reference fields return their original values **Example value**: true |
| `with_shared_url` | `boolean` | No | Controls whether to return a link to the record **Example value**: true |
| `automatic_fields` | `boolean` | No | Controls whether to return automatically calculated information, such as `created_by`/`created_time`/`last_modified_by`/`last_modified_time` **Example value**: true | ## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `record` | `app.table.record` | record |
|     `fields` | `map<string, union>` | fields |
|     `record_id` | `string` | record id，Update records are required |
|     `created_by` | `person` | record creator |
|       `id` | `string` | user id |
|       `name` | `string` | user name |
|       `en_name` | `string` | user english name |
|       `email` | `string` | user email |
|       `avatar_url` | `string` | \- **Required field scopes (Satisfy any)**: `contact:user.base:readonly` `contact:contact:access_as_app` `contact:contact:readonly` `contact:contact:readonly_as_app` |
|     `created_time` | `int` | record create timestamp |
|     `last_modified_by` | `person` | the person who last modified the record |
|       `id` | `string` | user id |
|       `name` | `string` | user name |
|       `en_name` | `string` | user english name |
|       `email` | `string` | user email |
|       `avatar_url` | `string` | \- **Required field scopes (Satisfy any)**: `contact:user.base:readonly` `contact:contact:access_as_app` `contact:contact:readonly` `contact:contact:readonly_as_app` |
|     `last_modified_time` | `int` | record last modified timestamp | ### Response body example

{
    "code": 0,
    "data": {
        "record": {
            "fields": {
                "Departments": [
                    "Product",
                    "R&D",
                    "Design"
                ],
                "Images": null,
                "Priority": "Normal",
                "Related OKRs": [
                    {
                        "text": "Assess needs and divide tasks",
                        "type": "text"
                    }
                ],
                "Start date": 1620316800000,
                "Task": "Review new features ",
                "Task leader": [
                    {
                        "email": "james@larksuite.com",
                        "en_name": "James",
                        "id": "ou_623c1831f27674bb94b3d2367aa99237",
                        "name": "James"
                    }
                ],
                "URL": {
                    "link": "https://bytedance.larksuite.com/docs/doccnM2J1HnL82yp20r2E6xw8te",
                    "text": "WeChat"
                },
                "formula": 7,
                "one-way link":[
                    {
                        "type":"text",
                        "table_id":"tbltAvx3DYBw7PVj",
                        "record_ids":[
                            "recl1IWVnB"
                        ],
                        "text":"content"
                    }
                ],
                "two-way link":[
                    {
                        "table_id":"tbltAvx3DYBw7PVj",
                        "record_ids":[
                            "recl1IWVnB",
                            "recrJk7SXT"
                        ],
                        "text":"content",
                        "type":"text"
                    }
                ]
            },
            "record_id": "recP750ZNJ",
            "record_url": "https://bytedance.larksuite.com/record/1sfvuxxxxxxxxxxxxxKdupE5Q"

        }
    },
    "msg": "Success"
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
| 200 | 1254030 | TooLargeResponse | TooLargeResponse |
| 400 | 1254036 | Base is copying, please try again later. | Base copy replicating, try again later |
| 200 | 1254040 | BaseTokenNotFound | AppToken not found |
| 200 | 1254041 | TableIdNotFound | Table not found |
| 200 | 1254042 | ViewIdNotFound | View not found |
| 200 | 1254043 | RecordIdNotFound | RecordIdNotFound |
| 200 | 1254044 | FieldIdNotFound | FieldIdNotFound |
| 200 | 1254045 | FieldNameNotFound | Field name does not exist |
| 200 | 1254060 | TextFieldConvFail | TextFieldConvFail |
| 200 | 1254061 | NumberFieldConvFail | NumberFieldConvFail |
| 200 | 1254062 | SingleSelectFieldConvFail | SingleSelectFieldConvFail |
| 200 | 1254063 | MultiSelectFieldConvFail | MultiSelectFieldConvFail |
| 200 | 1254064 | DatetimeFieldConvFail | DatetimeFieldConvFail |
| 200 | 1254065 | CheckboxFieldConvFail | CheckboxFieldConvFail |
| 200 | 1254066 | UserFieldConvFail | UserFieldConvFail |
| 200 | 1254067 | LinkFieldConvFail | LinkFieldConvFail |
| 200 | 1254072 | Failed to convert phone field, please make sure it is correct. | invalid phone format |
| 200 | 1254100 | TableExceedLimit | TableExceedLimit, limited to 300 |
| 200 | 1254101 | ViewExceedLimit | ViewExceedLimit, limited to 200 |
| 200 | 1254102 | FileExceedLimit | FileExceedLimit |
| 200 | 1254103 | RecordExceedLimit | RecordExceedLimit, limited to 20,000 |
| 200 | 1254104 | RecordAddOnceExceedLimit | RecordAddOnceExceedLimit, limited to 500 |
| 200 | 1254130 | TooLargeCell | TooLargeCell |
| 200 | 1254290 | TooManyRequest | Request too fast, try again later |
| 200 | 1254291 | Write conflict | The same data table does not support concurrent calls to the write interface, please check whether there is a concurrent call to the write interface. The writing interface includes: adding, modifying, and deleting records; adding, modifying, and deleting fields; modifying forms; modifying views, etc. |
| 200 | 1254301 | OperationTypeError | Base does not have advanced permissions enabled or does not support enabling advanced permissions |
| 200 | 1254302 | Permission denied. | No access rights, usually caused by the table opening of advanced permissions, please add a group containing applications in the advanced permissions settings and give this group read and write permissions |
| 200 | 1254303 | AttachPermNotAllow | No attach permission |
| 200 | 1255001 | InternalError | Internal error, have any questions can be consulting service |
| 200 | 1255002 | RpcError | Internal error, have any questions can be consulting service |
| 200 | 1255003 | MarshalError | Serialization failed, have any questions can be consulting service |
| 200 | 1255004 | UmMarshalError | Deserialization failed, have any questions can be consulting service |
| 200 | 1255005 | ConvError | Internal error, have any questions can be consulting service |
| 504 | 1255040 | Request timed out, please try again later | Try again |
| 400 | 1254607 | Data not ready, please try again later | There are usually two situations when this error occurs: 1. The last submitted modification has not been processed; 2. The data is too large and the server calculation times out; This error code can be appropriately retried. |
