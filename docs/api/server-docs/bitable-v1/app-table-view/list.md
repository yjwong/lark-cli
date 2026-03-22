---
title: "List Views"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/bitable-v1/app-table-view/list"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/bitable/v1/apps/:app_token/tables/:table_id/views"
service: "bitable-v1"
resource: "app-table-view"
section: "Docs"
rate_limit: "20 per second"
scopes:
  - "bitable:app"
  - "bitable:app:readonly"
field_scopes:
  - "contact:user.employee_id:readonly"
updated: "1753668072000"
---

# List view

Get all views of the data table based on app_token and table_id

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/bitable/v1/apps/:app_token/tables/:table_id/views |
| HTTP Method | GET |
| Rate Limit | 20 per second |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `bitable:app` `bitable:app:readonly` |
| Required field scopes | > The response body of the API contains the following sensitive fields, and they will be returned only after corresponding scopes are added. If you do not need the fields, it is not recommended that you request the scopes. `contact:user.employee_id:readonly` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token | ::: note
The instructions for AccessToken calling Docs API are detailed here Docs API Overview

### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `app_token` | `string` | Base app token **Example value**: "AppbcbWCzen6D8dezhoCH2RpMAh" **Data validation rules**: - Minimum length: `1` characters |
| `table_id` | `string` | Table id **Example value**: "TblsRc9GRRXKqhvW" | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `page_size` | `int` | No | **Example value**: 10 **Data validation rules**: - Maximum value: `100` |
| `page_token` | `string` | No | Page identifier. It is not filled in the first request, indicating traversal from the beginning; when there will be more groups, the new page_token will be returned at the same time, and the next traversal can use the page_token to get more groups **Example value**: VIVWTpR1urY |
| `user_id_type` | `string` | No | User ID categories **Example value**: open_id **Optional values are**:  - `open_id`: Identifies a user to an app. The same user has different Open IDs in different apps. How to get Open ID - `union_id`: Identifies a user to a tenant that acts as a developer. A user has the same Union ID in apps developed by the same developer, and has different Union IDs in apps developed by different developers. A developer can use Union ID to link the same user's identities in multiple apps.How to get Union ID - `user_id`: Identifies a user to a tenant. The same user has different User IDs in different tenants. In one single tenant, a user has the same User ID in all apps （including store apps）. User ID is usually used to communicate user data between different apps. How to get User ID  **Default value**: `open_id` **When the value is `user_id`, the following field scopes are required**: `contact:user.employee_id:readonly` | ## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `items` | `app.table.view[]` | View information |
|     `view_id` | `string` | View Id |
|     `view_name` | `string` | View name |
|     `view_type` | `string` | View type |
|     `view_public_level` | `string` | View public class Public, Locked, Private **Optional values are**:  - `Public`: Public view - `Locked`: Lock view - `Private`: Personal view  |
|     `view_private_owner_id` | `string` | The owner_id, id type and user_id_type parameters of the personal view are consistent |
|   `page_token` | `string` | Page identifier, when has_more is true, a new page_token will also be returned. Otherwise, page_token will not be returned |
|   `has_more` | `boolean` | Whether the response body has more parameters |
|   `total` | `int` | Total | ### Response body example

{
	"code": 0,
	"msg": "success",
	"data": {
		"has_more": false,
		"items": [{
				"view_id": "vewqtI3f2u",
				"view_name": "Public grid",
				"view_public_level": "Public",
				"view_type": "grid"
			},
			{
				"view_id": "vew5Ys1Y1B",
				"view_name": "Private grid",
				"view_private_owner_id": "ou_fe4e2a0c10f41fb85620eb4b71d11182",
				"view_public_level": "Private",
				"view_type": "grid"
			}
		],
		"page_token": "vew5Ys1Y1B",
		"total": 2
	}
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
| 200 | 1254016 | InvalidSort | invalid sort |
| 200 | 1254018 | InvalidFilter | The filter parameter is incorrect. Please refer to Record filter development guide for information on how to fill in the filter parameter. |
| 200 | 1254019 | InvalidViewType | Invalid view type |
| 200 | 1254020 | ViewNameDuplicated | Duplicate view name |
| 200 | 1254021 | EmptyViewName | View name is empty |
| 200 | 1254022 | InvalidViewName | Invalid view name |
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
| 200 | 1254100 | TableExceedLimit | TableExceedLimit, limited to 300 |
| 200 | 1254101 | ViewExceedLimit | ViewExceedLimit, limited to 200 |
| 200 | 1254102 | FileExceedLimit | FileExceedLimit |
| 200 | 1254103 | RecordExceedLimit | RecordExceedLimit, limited to 20,000 |
| 200 | 1254104 | RecordAddOnceExceedLimit | RecordAddOnceExceedLimit, limited to 500 |
| 200 | 1254130 | TooLargeCell | TooLargeCell |
| 200 | 1254290 | TooManyRequest | Request too fast, try again later |
| 200 | 1254291 | Write conflict | The same data table does not support concurrent calls to the write interface, please check whether there is a concurrent call to the write interface. The writing interface includes: adding, modifying, and deleting records; adding, modifying, and deleting fields; modifying forms; modifying views, etc. |
| 200 | 1254301 | OperationTypeError | Base does not have advanced permissions enabled or does not support enabling advanced permissions |
| 403 | 1254302 | Permission denied. | No access rights, usually caused by the table opening of advanced permissions, please add a group containing applications in the advanced permissions settings and give this group read and write permissions |
| 200 | 1255001 | InternalError | Internal error, have any questions can be consulting service |
| 200 | 1255002 | RpcError | Internal error, have any questions can be consulting service |
| 200 | 1255003 | MarshalError | Serialization failed, have any questions can be consulting service |
| 200 | 1255004 | UmMarshalError | Deserialization failed, have any questions can be consulting service |
| 200 | 1255005 | ConvError | Internal error, have any questions can be consulting service |
| 504 | 1255040 | Request timed out, please try again later | Try again |
| 400 | 1254607 | Data not ready, please try again later | There are usually two situations when this error occurs: 1. The last submitted modification has not been processed; 2. The data is too large and the server calculation times out; This error code can be appropriately retried. |
