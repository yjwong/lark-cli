---
title: "List fields"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/bitable-v1/app-table-field/list"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/bitable/v1/apps/:app_token/tables/:table_id/fields"
service: "bitable-v1"
resource: "app-table-field"
section: "Docs"
rate_limit: "20 per second"
scopes:
  - "bitable:app"
  - "bitable:app:readonly"
updated: "1753668100000"
---

# List fields

Get all fields according to app_token and table_id

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/bitable/v1/apps/:app_token/tables/:table_id/fields |
| HTTP Method | GET |
| Rate Limit | 20 per second |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `bitable:app` `bitable:app:readonly` base:field:read | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token | ::: note
The instructions for AccessToken calling Docs API are detailed here Docs API Overview

### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `app_token` | `string` | Base app token **Example value**: "appbcbWCzen6D8dezhoCH2RpMAh" |
| `table_id` | `string` | table id **Example value**: "tblsRc9GRRXKqhvW" | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `view_id` | `string` | No | view id **Example value**: vewOVMEXPF |
| `text_field_as_array` | `boolean` | No | Control the return format of field description (multi-line text format) data, true means return in array rich text form **Example value**: True |
| `page_token` | `string` | No | Page identifier. It is not filled in the first request, indicating traversal from the beginning; when there will be more groups, the new page_token will be returned at the same time, and the next traversal can use the page_token to get more groups **Example value**: fldwJ4YrtB |
| `page_size` | `int` | No | paging size **Example value**: 10 **Default value**: `20` **Data validation rules**: - Maximum value: `100` | ## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `has_more` | `boolean` | Whether the response body has more parameters |
|   `page_token` | `string` | Page identifier, when has_more is true, a new page_token will also be returned. Otherwise, page_token will not be returned |
|   `total` | `int` | Total |
|   `items` | `app.table.fieldForList[]` | Field information |
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
|     `description` | `union` | Description of the field. It can be either an array or a string type, determined by the request parameter `text_field_as_array`. |
|     `is_primary` | `boolean` | Whether it is an primary field. |
|     `field_id` | `string` | Field Id |
|     `ui_type` | `string` | The type of presentation of the field. **Optional values are**:  - `Text`: multiline text - `Barcode`: barcode - `Number`: number - `Progress`: progress - `Currency`: currency - `Rating`: score - `SingleSelect`: radio - `MultiSelect`: multiple choice - `DateTime`: date - `Checkbox`: checkbox - `User`: Personnel - `GroupChat`: group - `Phone`: Phone number - `Url`: Hyperlinke - `Attachment`: Annex - `SingleLink`: one-way association - `Formula`: formula - `DuplexLink`: bidirectional association - `Location`: Geographical location - `CreatedTime`: Creation time - `ModifiedTime`: Last update time - `CreatedUser`: founder - `ModifiedUser`: Modifier - `AutoNumber`: Automatic numbering  |
|     `is_hidden` | `boolean` | Whether it is a hidden field. | ### Response body example

{
    "code": 0,
    "data": {
        "has_more": false,
        "items": [
            {
                "field_id": "fldob4JqjK",
                "field_name": "primary",
                "is_primary": true,
                "property": null,
                "type": 1,
                "ui_type": "Text"
            },
            {
                "field_id": "fldJlvK2FH",
                "field_name": "text",
                "property": null,
                "type": 1,
                "ui_type": "Text"
            },
            {
                "field_id": "fldzwmS7AJ",
                "field_name": "number",
                "property": {
                    "formatter": "0.00"
                },
                "type": 2,
                "ui_type": "Number"
            },
            {
                "field_id": "fldm4AFURc",
                "field_name": "single_select",
                "property": {
                    "options": [
                        {
                            "color": 0,
                            "id": "optjelQbc7",
                            "name": "option_1"
                        },
                        {
                            "color": 1,
                            "id": "optvln6wtL",
                            "name": "option_2"
                        },
                        {
                            "color": 2,
                            "id": "opt88b49bs",
                            "name": "option_3"
                        }
                    ]
                },
                "type": 3,
                "ui_type": "SingleSelect"
            },
            {
                "field_id": "fldm4AFUpp",
                "field_name": "multi_select",
                "property": {
                    "options": [
                        {
                            "color": 0,
                            "id": "optoZ1U6SY",
                            "name": "选项1"
                        },
                        {
                            "color": 1,
                            "id": "optG9NcxmW",
                            "name": "选项2"
                        }
                    ]
                },
                "type": 4,
                "ui_type": "MultiSelect"
            },
            {
                "field_id": "fldGXNlvnR",
                "field_name": "date",
                "property": {
                    "auto_fill": true,
                    "date_formatter": "yyyy/MM/dd"
                },
                "type": 5,
                "ui_type": "DateTime"
            },
            {
                "field_id": "fldDn51X7W",
                "field_name": "checkbox",
                "property": null,
                "type": 7,
                "ui_type": "Checkbox"
            },
            {
                "field_id": "fldX96f4cH",
                "field_name": "user",
                "property": {
                    "multiple": true
                },
                "type": 11,
                "ui_type": "User"
            },
            {
                "field_id": "fldUgDVrl7",
                "field_name": "phone",
                "property": null,
                "type": 13,
                "ui_type": "Phone"
            },
            {
                "field_id": "fldVw52oFt",
                "field_name": "url",
                "property": null,
                "type": 15,
                "ui_type": "Url"
            },
            {
                "field_id": "fldkd6eX5b",
                "field_name": "attachment",
                "property": null,
                "type": 17,
                "ui_type": "Attachment"
            },
            {
                "field_id": "fld5AzyPGi",
                "field_name": "single_link",
                "property": {
                    "multiple": true,
                    "table_id": "tblBJyX6jZteblYv",
                    "table_name": "your table"
                },
                "type": 18,
                "ui_type": "SingleLink"
            },
            {
                "field_id": "fld1cXEUhn",
                "field_name": "duplex_link",
                "property": {
                    "back_field_id": "fldfjaXrPf",
                    "back_field_name": "duplex_link_add_field",
                    "multiple": true,
                    "table_id": "tblBJyX6jZteblYv",
                    "table_name": "your table"
                },
                "type": 21,
                "ui_type": "DuplexLink"
            },
            {
                "field_id": "fld9wy0PKz",
                "field_name": "location",
                "property": {
                    "location": {
                        "input_type": "not_limit"
                    }
                },
                "type": 22,
                "ui_type": "Location"
            },
            {
                "field_id": "fldWPpN2Wm",
                "field_name": "formula",
                "property": {
                    "formatter": "0.00%",
                    "formula_expression": "CONTAINTEXT(bitable::$table[tblBJyX6jZteblYv].$field[fldJlvK2FH],\"hello world\")"
                },
                "type": 20,
                "ui_type": "Formula"
            },
            {
                "field_id": "fldWqpN1Wm",
                "field_name": "created_time",
                "property": {
                    "date_formatter": "yyyy/MM/dd"
                },
                "type": 1001,
                "ui_type": "CreatedTime"
            },
            {
                "field_id": "fldWzzN1zm",
                "field_name": "modified_time",
                "property": {
                    "date_formatter": "yyyy/MM/dd HH:mm"
                },
                "type": 1002,
                "ui_type": "ModifiedTime"
            },
            {
                "field_id": "fldWzzN1ww",
                "field_name": "created_user",
                "property": null,
                "type": 1003,
                "ui_type": "CreatedUser"
            },
            {
                "field_id": "fldnnoWrgm",
                "field_name": "modified_user",
                "property": null,
                "type": 1004,
                "ui_type": "ModifiedUser"
            },
            {
                "field_id": "fldc4YUB2m",
                "field_name": "custom_auto_number",
                "property": {
                    "auto_serial": {
                        "options": [
                            {
                                "type": "system_number",
                                "value": "3"
                            },
                            {
                                "type": "fixed_text",
                                "value": "no"
                            },
                            {
                                "type": "created_time",
                                "value": "yyyyMMdd"
                            }
                        ],
                        "type": "custom"
                    }
                },
                "type": 1005,
                "ui_type": "AutoNumber"
            },
            {
                "field_id": "fld7sQvLyF",
                "field_name": "auto_number",
                "property": {
                    "auto_serial": {
                        "type": "auto_increment_number"
                    }
                },
                "type": 1005,
                "ui_type": "AutoNumber"
            },
            {
                "field_id": "fldvDB32aX",
                "field_name": "lookup",
                "property": null,
                "type": 19,
                "ui_type": "Lookup"
            },
            {
                "field_id": "fldUHxZlqG",
                "field_name": "barcode",
                "property": null,
                "type": 1,
                "ui_type": "Barcode"
            },
            {
                "field_id": "fldVanes5j",
                "field_name": "progress",
                "property": {
                    "formatter": "%"
                },
                "type": 2,
                "ui_type": "Progress"
            },
            {
                "field_id": "fldlwg2erf",
                "field_name": "currency",
                "property": {
                    "formatter": "0.0"
                },
                "type": 2,
                "ui_type": "Currency"
            }
        ],
        "page_token": "fldlwg2erf",
        "total": 25
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
| 200 | 1254012 | NotSupportFieldOrView | Unsupported fields or views |
| 200 | 1254013 | TableNameDuplicated | TableNameDuplicated |
| 200 | 1254014 | FieldNameDuplicated | Field name duplicate |
| 200 | 1254015 | FieldTypeValueNotMatch | FieldTypeValueNotMatch |
| 200 | 1254030 | TooLargeResponse | TooLargeResponse |
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
| 200 | 1254066 | UserFieldConvFail | The value corresponding to the personnel field type is incorrect. The possible reasons are: - The ID type specified by the user_id_type parameter does not match the type of the provided ID. - An unrecognized type or structure was provided. Currently, only `id` is supported, and it must be passed as an array. - An `open_id` was passed across applications. If you are passing an ID across applications, it is recommended to use `user_id`. The `open_id` obtained from different applications cannot be used interchangeably. |
| 200 | 1254067 | LinkFieldConvFail | LinkFieldConvFail |
| 200 | 1254070 | ActionValidateFailed | ActionValidateFailed |
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
| 403 | 1254302 | Permission denied. | No access rights, usually caused by the table opening of advanced permissions, please add a group containing applications in the advanced permissions settings and give this group read and write permissions |
| 200 | 1255001 | InternalError | Internal error, have any questions can be consulting service |
| 200 | 1255002 | RpcError | Internal error, have any questions can be consulting service |
| 200 | 1255003 | MarshalError | Serialization failed, have any questions can be consulting service |
| 200 | 1255004 | UmMarshalError | Deserialization failed, have any questions can be consulting service |
| 200 | 1255005 | ConvError | Internal error, have any questions can be consulting service |
| 500 | 1255040 | Request timed out, please try again later | Try again |
| 400 | 1254607 | Data not ready, please try again later | There are usually two situations when this error occurs: 1. The last submitted modification has not been processed; 2. The data is too large and the server calculation times out; This error code can be appropriately retried. |
