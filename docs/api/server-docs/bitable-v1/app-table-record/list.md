---
title: "List records"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/bitable-v1/app-table-record/list"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/bitable/v1/apps/:app_token/tables/:table_id/records"
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
updated: "1752650889000"
---

# List records

list records,Up to 500 lines at a time, paging is supported（Currently, the return of search reference fields are not supported. The search reference field can be converted into a formula field. The search reference is essentially a lookup formula）

::: note
This interface is a historical interface and is no longer recommended for use. You can use Query Record instead.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/bitable/v1/apps/:app_token/tables/:table_id/records |
| HTTP Method | GET |
| Rate Limit | 20 per second |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `bitable:app` `bitable:app:readonly` |
| Required field scopes | > The response body of the API contains the following sensitive fields, and they will be returned only after corresponding scopes are added. If you do not need the fields, it is not recommended that you request the scopes. `contact:user.employee_id:readonly` `contact:user.base:readonly` `contact:contact:access_as_app` `contact:contact:readonly` `contact:contact:readonly_as_app` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `app_token` | `string` | Base unique device identifier app_token description **Example value**: "bascnd0HM3KAyiZJELxfMHRrGZc" |
| `table_id` | `string` | Base data table unique device identifier table_id description **Example value**: "tblEGB3HKvDrpj71" | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `view_id` | `string` | No | Base view unique device identifier view_id description **Example value**: vewp7nmiS4 |
| `filter` | `string` | No | filter **Example value**: AND(CurrentValue.[Height]>180, CurrentValue.[Weight]>150) |
| `sort` | `string` | No | sort **Example value**: ["fieldName1 DESC","fieldName2 ASC"] |
| `field_names` | `string` | No | field_names **Example value**: ["fieldName1"] |
| `text_field_as_array` | `boolean` | No | indicate the structure of value of text field in response **Example value**: true |
| `user_id_type` | `string` | No | User ID categories **Example value**: open_id **Optional values are**:  - `open_id`: Identifies a user to an app. The same user has different Open IDs in different apps. How to get Open ID - `union_id`: Identifies a user to a tenant that acts as a developer. A user has the same Union ID in apps developed by the same developer, and has different Union IDs in apps developed by different developers. A developer can use Union ID to link the same user's identities in multiple apps.How to get Union ID - `user_id`: Identifies a user to a tenant. The same user has different User IDs in different tenants. In one single tenant, a user has the same User ID in all apps （including store apps）. User ID is usually used to communicate user data between different apps. How to get User ID  **Default value**: `open_id` **When the value is `user_id`, the following field scopes are required**: `contact:user.employee_id:readonly` |
| `display_formula_ref` | `boolean` | No | Control whether formulas, lookup reference fields return their original values **Example value**: true |
| `automatic_fields` | `boolean` | No | Controls whether to return automatically calculated information, such as `created_by`/`created_time`/`last_modified_by`/`last_modified_time` **Example value**: true |
| `page_token` | `string` | No | Page identifier. It is not filled in the first request, indicating traversal from the beginning; when there will be more groups, the new page_token will be returned at the same time, and the next traversal can use the page_token to get more groups **Example value**: recP750ZNJ |
| `page_size` | `int` | No | paging size **Example value**: 10 **Default value**: `20` **Data validation rules**: - Maximum value: `500` | ## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `has_more` | `boolean` | Whether the response body has more parameters |
|   `page_token` | `string` | Page identifier, when has_more is true, a new page_token will also be returned. Otherwise, page_token will not be returned |
|   `total` | `int` | total |
|   `items` | `app.table.record[]` | items |
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
|     `fields` | `map<string, union>` | fields | ### Response body example

{
	"code": 0,
	"msg": "success",
	"data": {
		"has_more": true,
		"page_token": "recG70uhxh",
		"total": 8,
		"items": [{
			"fields": {
				"primary": "primary",
				"text": "text",
				"number": "100",
				"single_select": "option_1",
				"multi_select": ["option_1", "option_2"],
				"date": 1674206443000,
				"checkbox": true,
				"user": [{
					"avatar_url": "https://internal-api-lark-file.larksuite.com/static-resource/v1/b2-7619-4b8a-b27b-c72d90b06a2j~?image_size=72x72&cut_type=default-face&quality=&format=jpeg&sticker_format=.webp",
					"email": "zhangsan.leben@bytedance.com",
					"en_name": "ZhangSan",
					"id": "ou_2910013f1e6456f16a0ce75ede950a0a",
					"name": "ZhangSan"
				}, {
					"avatar_url": "https://internal-api-lark-file.larksuite.com/static-resource/v1/v2_q86-fcb6-4f18-85c7-87ca8881e50j~?image_size=72x72&cut_type=default-face&quality=&format=jpeg&sticker_format=.webp",
					"email": "lisi.00@bytedance.com",
					"en_name": "LiSi",
					"id": "ou_e04138c9633dd0d2ea166d79f548ab5d",
					"name": "LiSi"
				}],
				"phone": "13126166666",
				"url": {
					"link": "https://bitable.larksuite.com",
					"text": "lark"
				},
				"attachment": [{
					"file_token": "Vl3FbVkvnowlgpxpqsAbBrtFcrd",
					"name": "lark.jpeg",
					"size": 32975,
					"tmp_url": "https://open.larksuite.com/open-apis/drive/v1/medias/batch_get_tmp_download_url?file_tokens=Vl3FbVk11owlgpxpqsAbBrtFcrd&extra=%7B%22bitablePerm%22%3A%7B%22tableId%22%3A%22tblBJyX6jZteblYv%22%2C%22rev%22%3A90%7D%7D",
					"type": "image/jpeg",
					"url": "https://open.larksuite.com/open-apis/drive/v1/medias/Vl3FbVk11owlgpxpqsAbBrtFcrd/download?extra=%7B%22bitablePerm%22%3A%7B%22tableId%22%3A%22tblBJyX6jZteblYv%22%2C%22rev%22%3A90%7D%7D"
				}],
				"single_link": [{
					"record_ids": ["recnVYsuqV"],
					"table_id": "tblBJyX6jZteblYv",
					"text": "primary",
					"text_arr": ["primary"],
					"type": "text"
				}],
				"duplex_link": [{
					"record_ids": ["recG70uhxh"],
					"table_id": "tblBJyX6jZteblYv",
					"text": "primary",
					"text_arr": ["primary"],
					"type": "text"
				}],
				"location": {
					"address": "东长安街",
					"adname": "东城区",
					"cityname": "北京市",
					"full_address": "天安门广场，北京市东城区东长安街",
					"location": "116.397755,39.903179",
					"name": "天安门广场",
					"pname": "北京市"
				},
				"formula": [{
					"text": "false",
					"type": "text"
				}],
				"created_time": 1675244156000,
				"modified_time": 1677556020000,
				"modified_user": {
					"avatar_url": "https://internal-api-lark-file.larksuite.com/static-resource/v1/06d568cb-f464-4c2e-bd03-76512c545c5j~?image_size=72x72&cut_type=default-face&quality=&format=jpeg&sticker_format=.webp",
					"email": "",
					"en_name": "bot",
					"id": "ou_92945f86a98bba075174776959c90eda",
					"name": "bot"
				},
				"created_user": {
					"avatar_url": "https://internal-api-lark-file.larksuite.com/static-resource/v1/06d568cb-f464-4c2e-bd03-76512c545c5j~?image_size=72x72&cut_type=default-face&quality=&format=jpeg&sticker_format=.webp",
					"email": "",
					"en_name": "bot",
					"id": "ou_92945f86a98bba075174776959c90eda",
					"name": "bot"
				},
				"barcode": "123",
				"lookup": [{
					"text": "text",
					"type": "text"
				}],
				"custom_auto_number": "017no20230201",
				"auto_number": "17",
				"currency": "1",
				"progress": "0.66"
			},
			"id": "recG70uhxh",
			"record_id": "recG70uhxh"
		}]
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
| 400 | 1254011 | Page size must greater than 0. | invalid page_size |
| 200 | 1254016 | InvalidSort | invalid sort |
| 200 | 1254018 | InvalidFilter | The filter parameter is incorrect. Please refer to Record filter development guide for information on how to fill in the filter parameter. |
| 200 | 1254024 | InvalidFieldNames | InvalidFieldNames |
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
| 200 | 1254068 | URLFieldConvFail | URLFieldConvFail |
| 200 | 1254069 | AttachFieldConvFail | AttachFieldConvFail |
| 200 | 1254072 | Failed to convert phone field, please make sure it is correct. | invalid phone format |
| 200 | 1254100 | TableExceedLimit | TableExceedLimit, limited to 300 |
| 200 | 1254101 | ViewExceedLimit | ViewExceedLimit, limited to 200 |
| 200 | 1254102 | FileExceedLimit | FileExceedLimit |
| 200 | 1254103 | RecordExceedLimit | RecordExceedLimit, limited to 20,000 |
| 200 | 1254104 | RecordAddOnceExceedLimit | RecordAddOnceExceedLimit, limited to 500 |
| 200 | 1254107 | FilterLengthExceedLimit | FilterLengthExceedLimit, limited to 2,000 characters |
| 200 | 1254108 | SortLengthExceedLimit | SortLengthExceedLimit, limited to 1,000 characters |
| 200 | 1254109 | FormulaTableSizeExceedLimit | FormulaTableSizeExceedLimit |
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
