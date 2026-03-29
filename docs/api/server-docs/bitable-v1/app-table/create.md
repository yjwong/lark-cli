---
title: "Create table"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/bitable-v1/app-table/create"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/bitable/v1/apps/:app_token/tables"
service: "bitable-v1"
resource: "app-table"
section: "Docs"
rate_limit: "10 per second"
scopes:
  - "base:table:create"
  - "bitable:app"
updated: "1773323335000"
---

# Create a table

Add a new table in Base, supporting the input of table name, view name, and fields.

## Prerequisite

Before calling this API, ensure that the current calling identity (tenant_access_token or user_access_token) has the necessary document permissions, such as edit permissions for Base. Otherwise, the API will return an HTTP 403 or 400 status code. For more information, refer to How to Grant Document Permissions to Apps or Users.

## Usage limits

In each Base, the total number of tables and dashboards is capped at 100.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/bitable/v1/apps/:app_token/tables |
| HTTP Method | POST |
| Rate Limit | 10 per second |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `base:table:create` `bitable:app` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `app_token` | `string` | Unique identifier for Base app_token parameter description **Example value**: "appbcbWCzen6D8dezhoCH2RpMAh" **Data validation rules**: - Minimum length: `1` characters | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `table` | `req_table` | No | table |
|   `name` | `string` | No | Name for the table. Please note: 1. The first and last spaces in the name will be removed. **Example value**: "table name" **Data validation rules**: - Length range: `1` ～ `100` characters |
|   `default_view_name` | `string` | No | Name of the default table view Please note: 1. The first and last spaces in the name will be removed. 2. The characters [ ] are not allowed to be included. **Example value**: "Table" |
|   `fields` | `app.table.create_header[]` | No | The initial fields of the data table. Please note: 1. If neither the default_view_name field nor the fields field is filled in, an empty data table containing only index columns will be created. 2. If the fields field is specified, a data table containing initial fields will be created and the first field will be the index column by default. **Data validation rules**: - Length range: `1` ～ `300` |
|     `field_name` | `string` | Yes | Field name **Example value**: "Text" |
|     `type` | `int` | Yes | Field type **Example value**: 1 **Optional values are**:  - `1`: Multiline - `2`: Number - `3`: Single option - `4`: Multiple options - `5`: Date - `7`: Checkbox - `11`: Person - `13`: PhoneNumber - `15`: Link - `17`: Attachment - `18`: One-way link - `20`: Formula - `21`: Two-way link - `22`: Location - `23`: group - `1001`: Date created - `1002`: Last modified date - `1003`: Created by - `1004`: Modified by - `1005`: AutoSerial  |
|     `ui_type` | `string` | No | Field Property, ref: Field edit development guide **Example value**: "Progress" **Optional values are**:  - `Text`: multiline text - `Barcode`: barcode - `Number`: number - `Progress`: progress - `Currency`: currency - `Rating`: score - `SingleSelect`: radio - `MultiSelect`: multiple choice - `DateTime`: date - `Checkbox`: checkbox - `User`: Personnel - `GroupChat`: group - `Phone`: Phone number - `Url`: Hyperlink - `Attachment`: Attachment - `SingleLink`: one-way association - `Formula`: formula - `DuplexLink`: Two-way link - `Location`: Geographical location - `CreatedTime`: Creation time - `ModifiedTime`: Last update time - `CreatedUser`: creator - `ModifiedUser`: Modifier - `AutoNumber`: Automatic numbering  |
|     `property` | `app.table.field.property` | No | Field description |
|       `options` | `app.table.field.property.option[]` | No | Whether to disable synchronization, if true, it means that synchronization of the description content to the problem description of the form is prohibited (only valid when fields are added or modified) |
|         `name` | `string` | No | Option name **Example value**: "Red." |
|         `id` | `string` | No | Option ID, not allowed to specify ID at creation time **Example value**: "optKl35lnG" |
|         `color` | `int` | No | Option color **Example value**: 0. **Data validation rules**: - Value range: `0` ～ `54` |
|       `formatter` | `string` | No | Display format of numbers and formula fields **Example value**: "0." |
|       `date_formatter` | `string` | No | The display format of the date, creation time, and last updated time fields **Example value**: "Date format" |
|       `auto_fill` | `boolean` | No | New records in the date field are automatically filled in Creation time **Example value**: False |
|       `multiple` | `boolean` | No | Multiple members are allowed to be added in the personnel field, and multiple records are allowed in one-way association and two-way association **Example value**: False |
|       `table_id` | `string` | No | The id of the associated data table in the one-way association, two-way association field **Example value**: "tblsRc9GRRXKqhvW" |
|       `table_name` | `string` | No | The name of the associated data table in the one-way association, two-way association field **Example value**: ""Table2"" |
|       `back_field_name` | `string` | No | The name of the corresponding bidirectional association field in the associated data table in the bidirectional association field **Example value**: ""Table1 - Bidirectional Association"" |
|       `auto_serial` | `app.field.property.auto_serial` | No | Automatic numbering type |
|         `type` | `string` | Yes | Automatic numbering type **Example value**: "auto_increment_number" **Optional values are**:  - `custom`: Custom number - `auto_increment_number`: Autoincrement number  |
|         `options` | `app.field.property.auto_serial.options[]` | No | List of auto-numbering rules |
|           `type` | `string` | Yes | Optional rule item types for auto-numbering **Example value**: "created_time" **Optional values are**:  - `system_number`: Incremental digits, value range 1-9 - `fixed_text`: Fixed characters, maximum length: 20 - `created_time`: Creation time, supports formats "yyyyMMdd", "yyyyMM", "yyyy", "MMdd", "MM", "dd"  |
|           `value` | `string` | Yes | Values corresponding to auto-numbered optional rule item types **Example value**: "YyyyMMdd" |
|       `location` | `app.field.property.location` | No | Geolocation input method |
|         `input_type` | `string` | Yes | Geolocation input restrictions **Example value**: "not_limit" **Optional values are**:  - `only_mobile`: Only allow uploads on mobile ends - `not_limit`: Unlimited  |
|       `formula_expression` | `string` | No | Expression of formula field **Example value**: "Bitable:: $table [tblNj92WQBAasdEf]. $field [fldMV60rYs] * 2" |
|       `allowed_edit_modes` | `allowed_edit_modes` | No | Editing modes supported by the field |
|         `manual` | `boolean` | No | Whether to allow manual entry **Example value**: true |
|         `scan` | `boolean` | No | Whether to allow mobile end entry **Example value**: true |
|       `min` | `number(float)` | No | Minimum data range for fields such as progress, score, etc **Example value**: 0 |
|       `max` | `number(float)` | No | Maximum data range for fields such as progress, score, etc **Example value**: 10 |
|       `range_customize` | `boolean` | No | Whether fields such as progress support custom ranges **Example value**: true |
|       `currency_code` | `string` | No | Currency **Example value**: "CNY" |
|       `rating` | `rating` | No | Relevant settings for scoring fields |
|         `symbol` | `string` | No | Symbol display for rating fields **Example value**: "star" |
|       `type` | `app.table.field.property.type` | No | / |
|         `data_type` | `int` | Yes | / **Example value**: 1 **Optional values are**:  - `1`: Text - `2`: Number - `3`: SingleSelect - `4`: MultiSelect - `5`: DateTime - `7`: Checkbox - `11`: User - `13`: PhoneNumber - `15`: Url - `17`: Attachment - `18`: Link - `20`: Formula - `21`: DuplexLink - `22`: Location - `23`: GroupChat - `1001`: CreatedTime - `1002`: ModifiedTime - `1003`: CreatedUser - `1004`: ModifiedUser - `1005`: AutoSerial  |
|         `ui_property` | `app.table.field.property.type.ui_property` | No | / |
|           `currency_code` | `string` | No | / **Example value**: "CNY" **Data validation rules**: - Length range: `0` ～ `20` characters |
|           `formatter` | `string` | No | / **Example value**: "0" **Data validation rules**: - Length range: `0` ～ `50` characters |
|           `range_customize` | `boolean` | No | / **Example value**: true |
|           `min` | `number(float)` | No | / **Example value**: 1 **Data validation rules**: - Value range: `0` ～ `1` |
|           `max` | `number(float)` | No | / **Example value**: 100 **Data validation rules**: - Value range: `1` ～ `100` |
|           `date_formatter` | `string` | No | / **Example value**: "yyyy/MM/dd" **Data validation rules**: - Length range: `0` ～ `50` characters |
|           `rating` | `rating` | No | / |
|             `symbol` | `string` | No | / **Example value**: "star" |
|         `ui_type` | `string` | No | / **Example value**: "Progress" **Optional values are**:  - `Number`: Number - `Progress`: Progress - `Currency`: Currency - `Rating`: Rating - `DateTime`: DateTime  |
|     `description` | `app.table.field.description` | No | Field description |
|       `disable_sync` | `boolean` | No | Whether to disable synchronization, if true, it means that synchronization of the description content to the problem description of the form is prohibited (only valid when fields are added or modified) **Example value**: true **Default value**: `true` |
|       `text` | `string` | No | Field description content **Example value**: "This is a field description" | ### Request body example

{
    "table":{
        "name":"table name",
        "default_view_name":"Grid",
        "fields":[
            {
                "field_name":"Text",
                "type":1
            }
        ]
    }
}

## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `table_id` | `string` | Bitable data table unique device identifier [table_id parameter description] (/ssl: ttdoc/uAjLw4CM/ukTMukTMukTM/bitable/notification #735fe883) |
|   `default_view_id` | `string` | The id of the default table view, which is returned only if default_view_name or fields are filled in the request parameters |
|   `field_id_list` | `string[]` | The id list of the initial field of the data table, which will only be returned if fields are filled in the request parameter | ### Response body example

{
	"Code": 0,
	"Msg": "success",
	"Data": {
		"table_id": "tblDBTWm6Es84d8c",
		"default_view_id": "vewUuKOz2R",
		"field_id_list": [
			"Fldhr2hBEA"
		]
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
| 200 | 1254007 | EmptyValue | Empty value |
| 200 | 1254008 | EmptyView | Empty view |
| 200 | 1254009 | WrongFieldId | Wrong fieldId |
| 200 | 1254010 | ReqConvError | Request error |
| 400 | 1254012 | NotSupportFieldOrView | Unsupported fields or views |
| 200 | 1254013 | TableNameDuplicated | TableNameDuplicated |
| 400 | 1254014 | FieldNameDuplicated | Field name duplicate |
| 400 | 1254021 | EmptyViewName | View name is empty |
| 400 | 1254022 | InvalidViewName | Invalid view name |
| 400 | 1254029 | InvalidFieldName | Invalid field name |
| 200 | 1254030 | TooLargeResponse | TooLargeResponse |
| 200 | 1254036 | Base is copying, please try again later. | Base copy replicating, try again later |
| 200 | 1254040 | BaseTokenNotFound | AppToken not found |
| 200 | 1254041 | TableIdNotFound | Table not found |
| 200 | 1254044 | FieldIdNotFound | FieldIdNotFound |
| 200 | 1254060 | TextFieldConvFail | TextFieldConvFail |
| 200 | 1254061 | NumberFieldConvFail | NumberFieldConvFail |
| 200 | 1254062 | SingleSelectFieldConvFail | SingleSelectFieldConvFail |
| 200 | 1254063 | MultiSelectFieldConvFail | MultiSelectFieldConvFail |
| 200 | 1254064 | DatetimeFieldConvFail | DatetimeFieldConvFail |
| 200 | 1254065 | CheckboxFieldConvFail | CheckboxFieldConvFail |
| 200 | 1254066 | UserFieldConvFail | The value corresponding to the personnel field type is incorrect. The possible reasons are: - The ID type specified by the user_id_type parameter does not match the type of the provided ID. - An unrecognized type or structure was provided. Currently, only `id` is supported, and it must be passed as an array. - An `open_id` was passed across applications. If you are passing an ID across applications, it is recommended to use `user_id`. The `open_id` obtained from different applications cannot be used interchangeably. |
| 200 | 1254067 | LinkFieldConvFail | LinkFieldConvFail |
| 200 | 1254100 | TableExceedLimit | TableExceedLimit, limited to 100 |
| 200 | 1254101 | ViewExceedLimit | ViewExceedLimit, limited to 200 |
| 200 | 1254130 | TooLargeCell | TooLargeCell |
| 200 | 1254290 | TooManyRequest | Request too fast, try again later |
| 200 | 1254291 | Write conflict | The same data table does not support concurrent calls to the write interface, please check whether there is a concurrent call to the write interface. The writing interface includes: adding, modifying, and deleting records; adding, modifying, and deleting fields; modifying forms; modifying views, etc. |
| 200 | 1254301 | OperationTypeError | Base does not have advanced permissions enabled or does not support enabling advanced permissions |
| 403 | 1254302 | Permission denied. | No access rights, usually caused by the table opening of advanced permissions, please add a group containing applications in the advanced permissions settings and give this group read and write permissions |
| 400 | 1254607 | Data not ready, please try again later | Internal error, have any questions can be consulting service |
| 200 | 1255001 | InternalError | Internal error, have any questions can be consulting service |
| 200 | 1255002 | RpcError | Internal error, have any questions can be consulting service |
| 200 | 1255003 | MarshalError | Serialization failed, have any questions can be consulting service |
| 200 | 1255004 | UmMarshalError | Deserialization failed, have any questions can be consulting service |
| 200 | 1255005 | ConvError | Internal error, have any questions can be consulting service |
