---
title: "Update records"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/bitable-v1/app-table-record/batch_update"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/bitable/v1/apps/:app_token/tables/:table_id/records/batch_update"
service: "bitable-v1"
resource: "app-table-record"
section: "Docs"
rate_limit: "50 per second"
scopes:
  - "bitable:app"
field_scopes:
  - "contact:user.employee_id:readonly"
  - "contact:user.base:readonly"
  - "contact:contact:access_as_app"
  - "contact:contact:readonly"
  - "contact:contact:readonly_as_app"
updated: "1753668081000"
---

# Update records

Update records,up to 500 lines at a time.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/bitable/v1/apps/:app_token/tables/:table_id/records/batch_update |
| HTTP Method | POST |
| Rate Limit | 50 per second |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `bitable:app` base:record:update |
| Required field scopes | > The response body of the API contains the following sensitive fields, and they will be returned only after corresponding scopes are added. If you do not need the fields, it is not recommended that you request the scopes. `contact:user.employee_id:readonly` `contact:user.base:readonly` `contact:contact:access_as_app` `contact:contact:readonly` `contact:contact:readonly_as_app` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ::: note
The instructions for AccessToken calling Docs API are detailed here Docs API Overview

### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `app_token` | `string` | base app_token **Example value**: "appbcbWCzen6D8dezhoCH2RpMAh" |
| `table_id` | `string` | table_id **Example value**: "tblsRc9GRRXKqhvW" | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `user_id_type` | `string` | No | User ID categories **Example value**: open_id **Optional values are**:  - `open_id`: Identifies a user to an app. The same user has different Open IDs in different apps. How to get Open ID - `union_id`: Identifies a user to a tenant that acts as a developer. A user has the same Union ID in apps developed by the same developer, and has different Union IDs in apps developed by different developers. A developer can use Union ID to link the same user's identities in multiple apps.How to get Union ID - `user_id`: Identifies a user to a tenant. The same user has different User IDs in different tenants. In one single tenant, a user has the same User ID in all apps （including store apps）. User ID is usually used to communicate user data between different apps. How to get User ID  **Default value**: `open_id` **When the value is `user_id`, the following field scopes are required**: `contact:user.employee_id:readonly` |
| `ignore_consistency_check` | `boolean` | No | Whether to ignore consistency checks for read and write operations. The default value is `false`, meaning the system will ensure that the data read and written is consistent. Optional values: - **true**: Ignore read/write consistency checks to improve performance, but this may cause data on some nodes to be out of sync, resulting in temporary inconsistency. - **false**: Enable read/write consistency checks to ensure data consistency during read and write operations. **Example value**: true | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `records` | `app.table.record[]` | Yes | records |
|   `fields` | `map<string, union>` | Yes | fields **Example value**: {"Multiline text": "HelloWorld"} |
|   `record_id` | `string` | No | record id. This parameter is required. Refer to record_id description to get record_id **Example value**: "recqwIwhc6" |
|   `shared_url` | `string` | No | shared link **Example value**: "https://www.example.com/record/WVoXrzIaqeorcJcHgzAcg8AQnNd" |
|   `record_url` | `string` | No | record link **Example value**: "https://www.example.com/record/WVoXrzIaqeorcJcHgzAcg8AQnNd" | ### Request body example

{
    "records": [
        {
            "record_id": "recgVdQ7o7",
            "fields": {
                "manpower": 2,
                "performer": [
                    {
                        "id": "ou_debc524b2d8cb187704df652b43d29de"
                    }
                ],
                "description": "collect user feedbacks",
                "corresponding OKR": [
                    "recqwIwhc6",
                    "recOuEJMvN"
                ],
                "deadline": 1609516800000,
                "completed": true,
                "status": "complete",
                "department": [
                    "Sale",
                    "Customer service"
                ]
            }
        },
        {
            "record_id":"recAu2ReK0",
            "fields": {
                "manpower": 2,
                "performer": [
                    {
                        "id": "ou_debc524b2d8cb187704df652b43d29de"
                    }
                ],
                "description": "collect user feedbacks",
                "corresponding OKR": [
                    "recqwIwhc6",
                    "recOuEJMvN"
                ],
                "deadline": 1609516800000,
                "completed": true,
                "status": "complete",
                "department": [
                    "Sale",
                    "Customer service"
                ]
            }
        }
    ]
}

## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `records` | `app.table.record[]` | records |
|     `fields` | `map<string, union>` | fields |
|     `record_id` | `string` | record id，Update records are required |
|     `created_by` | `person` | record creator |
|       `id` | `string` | user id |
|       `name` | `string` | user name |
|       `en_name` | `string` | user english name |
|       `email` | `string` | user email |
|       `avatar_url` | `string` | user avatar url **Required field scopes (Satisfy any)**: `contact:user.base:readonly` `contact:contact:access_as_app` `contact:contact:readonly` `contact:contact:readonly_as_app` |
|     `created_time` | `int` | record create timestamp |
|     `last_modified_by` | `person` | the person who last modified the record |
|       `id` | `string` | user id |
|       `name` | `string` | user name |
|       `en_name` | `string` | user english name |
|       `email` | `string` | user email |
|       `avatar_url` | `string` | user avatar url **Required field scopes (Satisfy any)**: `contact:user.base:readonly` `contact:contact:access_as_app` `contact:contact:readonly` `contact:contact:readonly_as_app` |
|     `last_modified_time` | `int` | record last modified timestamp |
|     `shared_url` | `string` | shared link |
|     `record_url` | `string` | record link | ### Response body example

{
  "Code": 0,
  "Data": {
    "Records": [
      {
        "Fields": {
          "Personnel": [
            {
              "Id": "ou_2910013f1e6456f16a0ce75ede950a0a"
            },
            {
              "Id": "ou_e04138c9633dd0d2ea166d79f548ab5d"
            }
          ],
          "Group": [
            {
              "Id": "oc_cd07f55f14d6f4a4f1b51504e7e97f48"
            }
          ],
          "One-way association": [
            "recHTLvO7x",
            "recbS8zb2m"
          ],
          "Radio": "Option 3",
          "Bidirectional association": [
            "recHTLvO7x",
            "recbS8zb2m"
          ],
          "Location": "116.397755, 39.903179",
          "Checkbox": true,
          "Multi-line text": "Multi-line text content",
          "Multiple Choice": [
            "Option 1",
            "Option 2"
          ],
          "Number": 100,
          "Date": 1674206443000,
          "Barcode": "qawqe",
          "Phone number": "13026162666",
          "Index": "Indexed column multi-row text type",
          "Hyperlinke": {
            "Link": "https://www.larksuite.com/product/base",
            "Text": "Bitable official website"
          },
          "Attachment": [
            {
              "file_token": "Vl3FbVkvnowlgpxpqsAbBrtFcrd"
            }
          ],
          "Rating": 3,
          "Currency": 3,
          "Progress": 0.25
        },
        "Id": "reclAqylTN",
        "record_id": "reclAqylTN"
      }
    ]
  },
  "Msg": "success"
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
| 400 | 1254015 | FieldTypeValueNotMatch | Field types do not match. |
| 403 | 1254027 | UploadAttachNotAllowed | Attachments don't belong to the app, not allowed to upload |
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
| 200 | 1254066 | UserFieldConvFail | The value corresponding to the personnel field type is incorrect. The possible reasons are: - The ID type specified by the user_id_type parameter does not match the type of the provided ID. - An unrecognized type or structure was provided. Currently, only `id` is supported, and it must be passed as an array. - An `open_id` was passed across applications. If you are passing an ID across applications, it is recommended to use `user_id`. The `open_id` obtained from different applications cannot be used interchangeably. |
| 200 | 1254067 | LinkFieldConvFail | LinkFieldConvFail |
| 200 | 1254068 | URLFieldConvFail | URLFieldConvFail |
| 200 | 1254069 | AttachFieldConvFail | AttachFieldConvFail |
| 200 | 1254072 | InvalidPhoneNumber | Failed to convert phone field, please make sure it is correct. |
| 400 | 1254074 | DuplexLinkFieldConvFail | The parameters of Duplex Link field are invalid and need to be filled with an array of string. |
| 200 | 1254100 | TableExceedLimit | TableExceedLimit, limited to 300 |
| 200 | 1254101 | ViewExceedLimit | ViewExceedLimit, limited to 200 |
| 200 | 1254102 | FileExceedLimit | FileExceedLimit |
| 200 | 1254103 | RecordExceedLimit | RecordExceedLimit, limited to 20,000 |
| 200 | 1254104 | RecordAddOnceExceedLimit | RecordAddOnceExceedLimit, limited to 500 |
| 200 | 1254105 | ColumnExceedLimit | ColumnExceedLimit |
| 200 | 1254106 | AttachExceedLimit | AttachExceedLimit |
| 200 | 1254130 | TooLargeCell | TooLargeCell |
| 200 | 1254290 | TooManyRequest | Request too fast, try again later |
| 200 | 1254291 | Write conflict | The same data table does not support concurrent calls to the write interface, please check whether there is a concurrent call to the write interface. The writing interface includes: adding, modifying, and deleting records; adding, modifying, and deleting fields; modifying forms; modifying views, etc. |
| 200 | 1254301 | OperationTypeError | Base does not have advanced permissions enabled or does not support enabling advanced permissions |
| 403 | 1254303 | AttachPermNotAllow | The attachment does not belong to this base. |
| 200 | 1255001 | InternalError | Internal error, have any questions can be consulting service |
| 200 | 1255002 | RpcError | Internal error, have any questions can be consulting service |
| 200 | 1255003 | MarshalError | Serialization failed, have any questions can be consulting service |
| 200 | 1255004 | UmMarshalError | Deserialization failed, have any questions can be consulting service |
| 200 | 1255005 | ConvError | Internal error, have any questions can be consulting service |
| 504 | 1255040 | Request timed out, please try again later | Try again |
| 400 | 1254607 | Data not ready, please try again later | There are usually two situations when this error occurs: 1. The last submitted modification has not been processed; 2. The data is too large and the server calculation times out; This error code can be appropriately retried. |
| 403 | 1254302 | Permission denied. | No access rights, usually caused by the table opening of advanced permissions, please add a group containing applications in the advanced permissions settings and give this group read and write permissions |
| 403 | 1254304 | Permission denied. | Advanced permissions for specific rows or columns are only available for Business and Enterprise editions |
| 403 | 1254608 | ReqRecommited | Same API requests are submitted repeatedly. |
