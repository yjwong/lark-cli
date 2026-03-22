---
title: "Create field"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/bitable-v1/app-table-field/create"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/bitable/v1/apps/:app_token/tables/:table_id/fields"
service: "bitable-v1"
resource: "app-table-field"
section: "Docs"
rate_limit: "10 per second"
scopes:
  - "bitable:app"
updated: "1753668100000"
---

# Create field

Create a field

> This API supports up to 10 queries per second (QPS).

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/bitable/v1/apps/:app_token/tables/:table_id/fields |
| HTTP Method | POST |
| Rate Limit | 10 per second |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `bitable:app` base:field:create | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `app_token` | `string` | Base unique device identifier app_token description **Example value**: "appbcbWCzen6D8dezhoCH2RpMAh" |
| `table_id` | `string` | Base data table unique device identifier table_id description **Example value**: "tblsRc9GRRXKqhvW" | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `client_token` | `string` | No | The format is a standard uuidv4, the unique identifier of the operation, used for idempotent update operations. This value is null to indicate that a new request will be initiated, and this value is non-null to indicate idempotent update operations. **Example value**: fe599b60-450f-46ff-b2ef-9f6675625b97 | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `field_name` | `string` | Yes | Field Name Please note: 1. The first and last spaces in the name will be removed. **Example value**: "Multiline text" |
| `type` | `int` | Yes | Field Type **Example value**: 1 **Optional values are**:  - `1`: Multiline - `2`: Number - `3`: Single option - `4`: Multiple options - `5`: Date - `7`: Checkbox - `11`: Person - `13`: PhoneNumber - `15`: Link - `17`: Attachment - `18`: One-way link - `20`: Formula - `21`: Two-way link - `22`: Location - `23`: group - `1001`: Date created - `1002`: Last modified date - `1003`: Created by - `1004`: Modified by - `1005`: AutoSerial  |
| `property` | `app.table.field.property` | No | Field Property, ref: Field edit development guide |
|   `options` | `app.table.field.property.option[]` | No | Option information for radio and multi-select fields |
|     `name` | `string` | No | Option name **Example value**: "Red." |
|     `id` | `string` | No | Option ID, not allowed to specify ID at creation time **Example value**: "optKl35lnG" |
|     `color` | `int` | No | Option color **Example value**: 0. **Data validation rules**: - Value range: `0` ～ `54` |
|   `formatter` | `string` | No | Display format of numbers and formula fields **Example value**: "0." |
|   `date_formatter` | `string` | No | The display format of the date, creation time, and last updated time fields **Example value**: "Date format" |
|   `auto_fill` | `boolean` | No | New records in the date field are automatically filled in Creation time **Example value**: False |
|   `multiple` | `boolean` | No | Multiple members are allowed to be added in the personnel field, and multiple records are allowed in one-way association and two-way association **Example value**: False |
|   `table_id` | `string` | No | The id of the associated data table in the one-way association, two-way association field **Example value**: "tblsRc9GRRXKqhvW" |
|   `back_field_name` | `string` | No | The name of the corresponding bidirectional association field in the associated data table in the bidirectional association field **Example value**: ""Table1 - Bidirectional Association"" |
|   `auto_serial` | `app.field.property.auto_serial` | No | Automatic numbering type |
|     `type` | `string` | Yes | Automatic numbering type **Example value**: "auto_increment_number" **Optional values are**:  - `custom`: Custom number - `auto_increment_number`: Autoincrement number  |
|     `options` | `app.field.property.auto_serial.options[]` | No | List of auto-numbering rules |
|       `type` | `string` | Yes | Optional rule item types for auto-numbering **Example value**: "created_time" **Optional values are**:  - `system_number`: Incremental digits, value range 1-9 - `fixed_text`: Fixed characters, maximum length: 20 - `created_time`: Creation time, supports formats "yyyyMMdd", "yyyyMM", "yyyy", "MMdd", "MM", "dd"  |
|       `value` | `string` | Yes | Values corresponding to auto-numbered optional rule item types **Example value**: "YyyyMMdd" |
|   `location` | `app.field.property.location` | No | Geolocation input method |
|     `input_type` | `string` | Yes | Geolocation input restrictions **Example value**: "not_limit" **Optional values are**:  - `only_mobile`: Only allow uploads on mobile ends - `not_limit`: Unlimited  |
|   `formula_expression` | `string` | No | Expression of formula field **Example value**: "Bitable:: $table [tblNj92WQBAasdEf]. $field [fldMV60rYs] * 2" |
|   `allowed_edit_modes` | `allowed_edit_modes` | No | Editing modes supported by the field |
|     `manual` | `boolean` | No | Whether to allow manual entry **Example value**: true |
|     `scan` | `boolean` | No | Whether to allow mobile end entry **Example value**: true |
|   `min` | `number(float)` | No | Minimum data range for fields such as progress, score, etc **Example value**: 0 |
|   `max` | `number(float)` | No | Maximum data range for fields such as progress, score, etc **Example value**: 10 |
|   `range_customize` | `boolean` | No | Whether fields such as progress support custom ranges **Example value**: true |
|   `currency_code` | `string` | No | Currency Currency **Example value**: "CNY" |
|   `rating` | `rating` | No | Relevant settings for scoring fields |
|     `symbol` | `string` | No | Symbol display for rating fields **Example value**: "star" |
|   `type` | `app.table.field.property.type` | No | Set the type of the formula field |
|     `data_type` | `int` | Yes | Set the data type of the formula field **Example value**: 1 **Optional values are**:  - `1`: Text - `2`: Number - `3`: SingleSelect - `4`: MultiSelect - `5`: DateTime - `7`: Checkbox - `11`: User - `13`: PhoneNumber - `15`: Url - `17`: Attachment - `18`: Link - `20`: Formula - `21`: DuplexLink - `22`: Location - `23`: GroupChat - `1001`: CreatedTime - `1002`: ModifiedTime - `1003`: CreatedUser - `1004`: ModifiedUser - `1005`: AutoSerial  |
|     `ui_property` | `app.table.field.property.type.ui_property` | No | Field Property, ref: Field edit development guide |
|       `currency_code` | `string` | No | Currency **Example value**: "CNY" **Data validation rules**: - Length range: `0` ～ `20` characters |
|       `formatter` | `string` | No | Display format of numbers and formula fields **Example value**: "0" **Data validation rules**: - Length range: `0` ～ `50` characters |
|       `range_customize` | `boolean` | No | Whether fields such as progress support custom ranges **Example value**: true |
|       `min` | `number(float)` | No | Minimum data range for fields such as progress, score, etc **Example value**: 1 **Data validation rules**: - Value range: `0` ～ `1` |
|       `max` | `number(float)` | No | Maximum data range for fields such as progress, score, etc **Example value**: 100 **Data validation rules**: - Value range: `1` ～ `100` |
|       `date_formatter` | `string` | No | The display format of the date, creation time, and last updated time fields **Example value**: "yyyy/MM/dd" **Data validation rules**: - Length range: `0` ～ `50` characters |
|       `rating` | `rating` | No | Relevant settings for scoring fields |
|         `symbol` | `string` | No | Symbol display for rating fields **Example value**: "star" |
|     `ui_type` | `string` | No | The type of display of the formula field on the interface, such as the progress field, which is a display form of numbers **Example value**: "Progress" **Optional values are**:  - `Number`: Number - `Progress`: Progress - `Currency`: Currency - `Rating`: Rating - `DateTime`: DateTime  |
| `description` | `app.table.field.description` | No | Field description |
|   `disable_sync` | `boolean` | No | Whether to disable synchronization, if true, it means that synchronization of the description content to the problem description of the form is prohibited (only valid when fields are added or modified) **Example value**: True **Default value**: `true` |
|   `text` | `string` | No | Field description content **Example value**: "This is a field description" |
| `ui_type` | `string` | No | The type of display of the field on the interface, such as the progress field, which is a display form of numbers **Example value**: "Progress" **Optional values are**:  - `Text`: multiline text - `Email`: Email - `Barcode`: barcode - `Number`: number - `Progress`: progress - `Currency`: currency - `Rating`: score - `SingleSelect`: radio - `MultiSelect`: multiple choice - `DateTime`: date - `Checkbox`: checkbox - `User`: Personnel - `GroupChat`: group - `Phone`: Phone number - `Url`: Hyperlink - `Attachment`: Attachment - `SingleLink`: one-way association - `Formula`: formula - `DuplexLink`: bidirectional association - `Location`: Geographical location - `CreatedTime`: Creation time - `ModifiedTime`: Last update time - `CreatedUser`: founder - `ModifiedUser`: Modifier - `AutoNumber`: Automatic numbering  | ### Request body example

{
    "field_name":"Text",
    "type":1
}

## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `field` | `app.table.field` | Field |
|     `field_name` | `string` | Field Name Please note: 1. The first and last spaces in the name will be removed. |
|     `type` | `int` | Field Type **Optional values are**:  - `1`: Multiline - `2`: Number - `3`: Single option - `4`: Multiple options - `5`: Date - `7`: Checkbox - `11`: Person - `13`: PhoneNumber - `15`: Link - `17`: Attachment - `18`: One-way link - `20`: Formula - `21`: Two-way link - `22`: Location - `23`: group - `1001`: Date created - `1002`: Last modified date - `1003`: Created by - `1004`: Modified by - `1005`: AutoSerial  |
|     `property` | `app.table.field.property` | Field Property, ref: Field edit development guide |
|       `options` | `app.table.field.property.option[]` | Option information for radio and multi-select fields |
|         `name` | `string` | Option name |
|         `id` | `string` | Option ID, not allowed to specify ID at creation time |
|         `color` | `int` | Option color |
|       `formatter` | `string` | Display format of numbers and formula fields |
|       `date_formatter` | `string` | The display format of the date, creation time, and last updated time fields |
|       `auto_fill` | `boolean` | New records in the date field are automatically filled in Creation time |
|       `multiple` | `boolean` | Multiple members are allowed to be added in the personnel field, and multiple records are allowed in one-way association and two-way association |
|       `table_id` | `string` | The id of the associated data table in the one-way association, two-way association field |
|       `table_name` | `string` | The name of the associated data table in the one-way association, two-way association field |
|       `back_field_name` | `string` | The name of the corresponding bidirectional association field in the associated data table in the bidirectional association field |
|       `auto_serial` | `app.field.property.auto_serial` | Automatic numbering type |
|         `type` | `string` | Automatic numbering type **Optional values are**:  - `custom`: Custom number - `auto_increment_number`: Autoincrement number  |
|         `options` | `app.field.property.auto_serial.options[]` | List of auto-numbering rules |
|           `type` | `string` | Optional rule item types for auto-numbering **Optional values are**:  - `system_number`: Incremental digits, value range 1-9 - `fixed_text`: Fixed characters, maximum length: 20 - `created_time`: Creation time, supports formats "yyyyMMdd", "yyyyMM", "yyyy", "MMdd", "MM", "dd"  |
|           `value` | `string` | Values corresponding to auto-numbered optional rule item types |
|       `location` | `app.field.property.location` | Geolocation input method |
|         `input_type` | `string` | Geolocation input restrictions **Optional values are**:  - `only_mobile`: Only allow uploads on mobile ends - `not_limit`: Unlimited  |
|       `formula_expression` | `string` | Expression of formula field |
|       `allowed_edit_modes` | `allowed_edit_modes` | Editing modes supported by the field |
|         `manual` | `boolean` | Whether to allow manual entry |
|         `scan` | `boolean` | Whether to allow mobile end entry |
|       `min` | `number(float)` | Minimum data range for fields such as progress, score, etc |
|       `max` | `number(float)` | Maximum data range for fields such as progress, score, etc |
|       `range_customize` | `boolean` | Whether fields such as progress support custom ranges |
|       `currency_code` | `string` | Currency Currency |
|       `rating` | `rating` | Relevant settings for scoring fields |
|         `symbol` | `string` | Symbol display for rating fields |
|       `type` | `app.table.field.property.type` | Set the type of the formula field |
|         `data_type` | `int` | Set the data type of the formula field **Optional values are**:  - `1`: Text - `2`: Number - `3`: SingleSelect - `4`: MultiSelect - `5`: DateTime - `7`: Checkbox - `11`: User - `13`: PhoneNumber - `15`: Url - `17`: Attachment - `18`: Link - `20`: Formula - `21`: DuplexLink - `22`: Location - `23`: GroupChat - `1001`: CreatedTime - `1002`: ModifiedTime - `1003`: CreatedUser - `1004`: ModifiedUser - `1005`: AutoSerial  |
|         `ui_property` | `app.table.field.property.type.ui_property` | Field Property, ref: Field edit development guide |
|           `currency_code` | `string` | Currency |
|           `formatter` | `string` | Display format of numbers and formula fields |
|           `range_customize` | `boolean` | Whether fields such as progress support custom ranges |
|           `min` | `number(float)` | Minimum data range for fields such as progress, score, etc |
|           `max` | `number(float)` | Maximum data range for fields such as progress, score, etc |
|           `date_formatter` | `string` | The display format of the date, creation time, and last updated time fields |
|           `rating` | `rating` | Relevant settings for scoring fields |
|             `symbol` | `string` | Symbol display for rating fields |
|         `ui_type` | `string` | The type of display of the formula field on the interface, such as the progress field, which is a display form of numbers **Optional values are**:  - `Number`: Number - `Progress`: Progress - `Currency`: Currency - `Rating`: Rating - `DateTime`: DateTime  |
|     `description` | `app.table.field.description` | Field description |
|       `disable_sync` | `boolean` | Whether to disable synchronization, if true, it means that synchronization of the description content to the problem description of the form is prohibited (only valid when fields are added or modified) |
|       `text` | `string` | Field description content |
|     `is_primary` | `boolean` | Identifies whether the current field is the primary key |
|     `field_id` | `string` | Field Id |
|     `ui_type` | `string` | Identifies whether the current field is hidden **Optional values are**:  - `Text`: multiline text - `Email`: Email - `Barcode`: barcode - `Number`: number - `Progress`: progress - `Currency`: currency - `Rating`: score - `SingleSelect`: radio - `MultiSelect`: multiple choice - `DateTime`: date - `Checkbox`: checkbox - `User`: Personnel - `GroupChat`: group - `Phone`: Phone number - `Url`: Hyperlinke - `Attachment`: Annex - `SingleLink`: one-way association - `Formula`: formula - `DuplexLink`: Two-way link - `Location`: Geographical location - `CreatedTime`: Creation time - `ModifiedTime`: Last update time - `CreatedUser`: creator - `ModifiedUser`: Modifier - `AutoNumber`: Automatic numbering  |
|     `is_hidden` | `boolean` | Is it a hidden field | ### Response body example

{
    "code": 0,
    "data": {
        "field": {
            "field_id": "fld4bocNLY",
            "field_name": "Text",
            "type": 1,
            "property": null
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
| 200 | 1254012 | NotSupportFieldOrView | Unsupported fields or views |
| 200 | 1254013 | TableNameDuplicated | TableNameDuplicated |
| 200 | 1254014 | FieldNameDuplicated | Field name duplicate |
| 200 | 1254015 | FieldTypeValueNotMatch | FieldTypeValueNotMatch |
| 200 | 1254026 | EmptyOptionName | option name should not be empty |
| 400 | 1254028 | EmptyFieldName | field name should not be empty |
| 400 | 1254029 | InvalidFieldName | Invalid field name |
| 200 | 1254030 | TooLargeResponse | TooLargeResponse |
| 400 | 1254036 | Base is copying, please try again later. | Base copy replicating, try again later |
| 400 | 1254037 | Invalid client token, make sure that it complies with the specification. | Idempotent key format is wrong, you need to pass in uuidv4 format |
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
| 200 | 1254066 | UserFieldConvFail | The value corresponding to the personnel field type is incorrect. The possible reasons are: - The ID type specified by the user_id_type parameter does not match the type of the provided ID. - An unrecognized type or structure was provided. Currently, only `id` is supported, and it must be passed as an array. - An `open_id` was passed across applications. If you are passing an ID across applications, it is recommended to use `user_id`. The `open_id` obtained from different applications cannot be used interchangeably. |
| 200 | 1254067 | LinkFieldConvFail | LinkFieldConvFail |
| 200 | 1254070 | ActionValidateFailed | ActionValidateFailed |
| 400 | 1254080 | TextFieldPropertyError | TextFieldPropertyError |
| 400 | 1254081 | NumberFieldPropertyError | NumberFieldPropertyError |
| 400 | 1254082 | SingleSelectFieldPropertyError | SingleSelectFieldPropertyError |
| 400 | 1254083 | MultiSelectFieldPropertyError | MultiSelectFieldPropertyError |
| 400 | 1254084 | DateFieldPropertyError | DateFieldPropertyError |
| 400 | 1254085 | CheckboxFieldPropertyError | CheckboxFieldPropertyError |
| 400 | 1254086 | UserFieldPropertyError | UserFieldPropertyError |
| 400 | 1254087 | URLFieldPropertyError | URLFieldPropertyError |
| 400 | 1254088 | AttachFieldPropertyError | AttachFieldPropertyError |
| 400 | 1254089 | LinkFieldPropertyError | LinkFieldPropertyError |
| 400 | 1254090 | LookUpFieldPropertyError | LookUpFieldPropertyError |
| 400 | 1254091 | FormulaFieldPropertyError | FormulaFieldPropertyError |
| 400 | 1254092 | DuplexLinkFieldPropertyError | DuplexLinkFieldPropertyError |
| 400 | 1254093 | CreatedTimeFieldPropertyError | CreatedTimeFieldPropertyError |
| 400 | 1254094 | ModifiedTimeFieldPropertyError | ModifiedTimeFieldPropertyError |
| 400 | 1254095 | CreatedUserFieldPropertyError | CreatedUserFieldPropertyError |
| 400 | 1254096 | ModifiedUserFieldPropertyError | ModifiedUserFieldPropertyError |
| 200 | 1254100 | TableExceedLimit | TableExceedLimit, limited to 300 |
| 200 | 1254101 | ViewExceedLimit | ViewExceedLimit, limited to 200 |
| 200 | 1254102 | FileExceedLimit | FileExceedLimit |
| 200 | 1254103 | RecordExceedLimit | RecordExceedLimit, limited to 20,000 |
| 200 | 1254104 | RecordAddOnceExceedLimit | RecordAddOnceExceedLimit, limited to 500 |
| 200 | 1254105 | ColumnExceedLimit | ColumnExceedLimit |
| 200 | 1254130 | TooLargeCell | TooLargeCell |
| 200 | 1254290 | TooManyRequest | Request too fast, try again later |
| 200 | 1254291 | Write conflict | The same data table does not support concurrent calls to the write interface, please check whether there is a concurrent call to the write interface. The writing interface includes: adding, modifying, and deleting records; adding, modifying, and deleting fields; modifying forms; modifying views, etc. |
| 200 | 1254301 | OperationTypeError | Base does not have advanced permissions enabled or does not support enabling advanced permissions |
| 200 | 1255001 | InternalError | Internal error, have any questions can be consulting service |
| 200 | 1255002 | RpcError | Internal error, have any questions can be consulting service |
| 200 | 1255003 | MarshalError | Serialization failed, have any questions can be consulting service |
| 200 | 1255004 | UmMarshalError | Deserialization failed, have any questions can be consulting service |
| 200 | 1255005 | ConvError | Internal error, have any questions can be consulting service |
| 400 | 1255006 | Client token conflict, please generate a new client token and try again. | Idempotent key conflict, you need to randomly generate an idempotent key |
| 504 | 1255040 | Request timed out, please try again later | Try again |
| 400 | 1254607 | Data not ready, please try again later | There are usually two situations when this error occurs: 1. The last submitted modification has not been processed; 2. The data is too large and the server calculation times out; This error code can be appropriately retried. |
| 403 | 1254302 | Permission denied. | No access rights, usually caused by the table opening of advanced permissions, please add a group containing applications in the advanced permissions settings and give this group read and write permissions |
| 403 | 1254304 | Permission denied. | Advanced permissions for specific rows or columns are only available for Business and Enterprise editions |
| 403 | 1254608 | ReqRecommited | Same API requests are submitted repeatedly. |
