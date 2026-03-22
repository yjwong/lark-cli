---
title: "Field edit development guide"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/bitable-v1/app-table-field/guide"
service: "bitable-v1"
resource: "app-table-field"
section: "Docs"
updated: "1753668094000"
---

# Field editing guide

Fields are the "columns" in a Base data table and are of the `object` structure type. When adding or updating fields in a Base data table, you can refer to this document to learn how to select field types and set field properties.

The basic structure of a field is as follows:

```json
{
    "field_id": "fldYWaldeW", // Field ID
    "field_name": "test",   // Field name
    "type": 1,  // Field type, see below for details
    "description": "description", // More description of the field
    "is_primary": true,   // Whether the field is the primary index field
    "property": null,   // Field properties, see below for details
    "ui_type": "Text",  // Field display type on the interface, e.g., progress field is a type of number display
    "is_hidden": false  // Whether the field is hidden
}
```
Parameter descriptions are as follows:
| Name | Type | Description |
| --- | --- | --- |
| field_id | string | Field ID |
| field_name | string | Field name |
| type | int | Field type: different types are distinguished by ui_type: - 1: Text (default), Barcode (requires "ui_type": "Barcode"), Email (requires "ui_type": "Email") - 2: Number (default), Progress (requires "ui_type": "Progress"), Currency (requires "ui_type": "Currency"), Rating (requires "ui_type": "Rating") - 3: Single select - 4: Multi select - 5: Date - 7: Checkbox - 11: User - 13: Phone number - 15: URL - 17: Attachment - 18: Single link - 19: Lookup - 20: Formula - 21: Duplex link - 22: Location - 23: Group - 24: Stage - 1001: Created time - 1002: Modified time - 1003: Created user - 1004: Modified user - 1005: Auto number - 3001: Button |
| description | string | More description of the field. |
| is_primary | true/false | Whether the field is the primary index field. |
| property | object | Field properties, vary by field type. See Field editing guide for details. |
| ui_type | string | Field UI type: - "Text": Text - "Email": Email - "Barcode": Barcode - "Number": Number - "Progress": Progress - "Currency": Currency - "Rating": Rating - "SingleSelect": Single select - "MultiSelect": Multi select - "DateTime": Date - "Checkbox": Checkbox - "User": User - "GroupChat": Group chat - "Stage": Stage - "Phone": Phone number - "Url": URL - "Attachment": Attachment - "SingleLink": Single link - "Formula": Formula - "Lookup": Lookup - "DuplexLink": Duplex link - "Location": Location - "CreatedTime": Created time - "ModifiedTime": Modified time - "CreatedUser": Created user - "ModifiedUser": Modified user - "AutoNumber": Auto number - "Button": Button |
| is_hidden | true/false | Whether the field is hidden. | ## Index field description

In a data table, the first column is the index column, i.e., the index field. `"is_primary": true` indicates that the field is an index field. Indexes cannot be deleted, moved, or hidden, and only support the following field types (type):
- 1: Multiline text
- 2: Number
- 5: Date
- 13: Phone number
- 15: URL
- 20: Formula
- 22: Location

## Field type `type`

Base provides a variety of field types, corresponding to the parameter `type`. A field type may have multiple display types on the interface, corresponding to the parameter `ui_type`. For example, when the field type `type` is enumerated as 1, it indicates a text type, and the corresponding UI display type `ui_type` can be:
- `Text`: Text display type
- `Barcode`: Barcode display type
- `Email`: Email display type

The following are the enumeration values for field types and field UI display types:
> **Note:** Different interfaces support different field types. Please refer to the corresponding interface documentation for the optional values of field types.

| Enumeration values of field type `type` | Enumeration values of field UI display type ui_type |
| --- | --- |
| - 1: Text (default), Barcode (requires "ui_type": "Barcode"), Email (requires "ui_type": "Email") - 2: Number (default), Progress (requires "ui_type": "Progress"), Currency (requires "ui_type": "Currency"), Rating (requires "ui_type": "Rating") - 3: Single select - 4: Multi select - 5: Date - 7: Checkbox - 11: User - 13: Phone number - 15: URL - 17: Attachment - 18: Single link - 19: Lookup - 20: Formula - 21: Duplex link - 22: Location - 23: Group - 24: Stage (not supported through write interface for adding or editing, only supported through read interface) - 1001: Created time - 1002: Modified time - 1003: Created user - 1004: Modified user - 1005: Auto number - 3001: Button (not supported through write interface for adding or editing, only supported through read interface) | - "Text": Text - "Email": Email - "Barcode": Barcode - "Number": Number - "Progress": Progress - "Currency": Currency - "Rating": Rating - "SingleSelect": Single select - "MultiSelect": Multi select - "DateTime": Date - "Checkbox": Checkbox - "User": User - "GroupChat": Group chat - "Stage": Stage - "Phone": Phone number - "Url": URL - "Attachment": Attachment - "SingleLink": Single link - "Formula": Formula - "Lookup": Lookup - "DuplexLink": Duplex link - "Location": Location - "CreatedTime": Created time - "ModifiedTime": Modified time - "CreatedUser": Created user - "ModifiedUser": Modified user - "AutoNumber": Auto number - "Button": Button | # Field properties

Field properties refer to additional functional attributes specified for a field type. For example, for a user type field, whether to enable the feature to add multiple members. The `property` structure corresponding to different types of fields is different. The following types of fields have no additional functional attributes, and their `property` is `null`:
- 1: Text (default), Email (requires `"ui_type": "Email"`)
- 7: Checkbox
- 13: Phone number
- 15: URL
- 17: Attachment
- 19: Lookup
- 1003: Created user
- 1004: Modified user

For example, when the field type is text and the field UI display type is email, its `property` should be set to `null`:

```json
{
    "field_name": "Email",
    "type": 1,
    "ui_type": "Email",
    "property": null
}
```

Other types of fields have additional functional attributes. You can refer to the following to understand and use the `property` structure of different types of fields.

### Barcode field

The field type of the barcode field is `1`, and the UI display type `ui_type` is `"Barcode"`. Its functional attributes are as follows:

| 
Name | Type | Required | Description |
| --- | --- | --- | --- |
| `allowed_edit_modes` | `allowed_edit_modes` | No | Configuration of the barcode field |
|   `manual` | bool | No | Whether manual input is allowed. Default is true. |
|   scan | bool | No | Whether mobile input is allowed. Default is true. | #### Request body example

The request body example for adding or updating a field is as follows:
```json
{
    "field_name": "Barcode",
    "type": 1,
    "ui_type": "Barcode",
    "property": {
        "allowed_edit_modes": {
            "scan": true,
            "manual": true
        }
    }
}
```

#### Response body example

```json
{
  "code": 0,
  "data": {
    "field": {
      "field_id": "fldQoEj2p3",
      "field_name": "Barcode",
      "is_primary": false,
      "property": {
        "allowed_edit_modes": {
          "manual": true,
          "scan": true
        }
      },
      "type": 1,
      "ui_type": "Barcode"
    }
  },
  "msg": "success"
}
```
### Number field

The field type of the number field is `2`, and there is no need to declare the UI display type `ui_type`. Its functional attributes are as follows:
| Name | Type | Required | Description |
| --- | --- | --- | --- |
| formatter | string | No | Number format, default is "0.0", enumeration values are as follows: - "0": Integer - "0.0": 1 decimal place - "0.00": 2 decimal places - "0.000": 3 decimal places - "0.0000": 4 decimal places - "1,000": Thousands separator - "1,000.00": Thousands separator (decimal point) - "%": Percentage - "0.00%": Percentage (decimal point) - "¥": RMB - "¥0.00": RMB (decimal point) - "$": USD - "$0.00": USD (decimal point) | #### Request body example

The request body example for adding or updating a field is as follows:
```json
{
    "field_name": "Number",
    "type": 2,
    "property": {
        "formatter": "0.00"
    }
}
```

#### Response body example

```json
{
  "code": 0,
  "data": {
    "field": {
      "field_id": "fldP9gzwe7",
      "field_name": "Number",
      "is_primary": false,
      "property": {
        "formatter": "0.00"
      },
      "type": 2,
      "ui_type": "Number"
    }
  },
  "msg": "success"
}
```

### Currency field

The field type of the currency field is `2`, and the UI display type `ui_type` is `"Currency"`. Its functional attributes are as follows:

| 
Name | Type | Required | Description |
| --- | --- | --- | --- |
| `formatter` | string | Yes | Number of decimal places for the currency, enumeration values are as follows: - "0": Integer - "0.0": 1 decimal place - "0.00": 2 decimal places - "0.000": 3 decimal places - "0.0000": 4 decimal places |
| currency_code | string | Yes |  | #### Request body example

The request body example for adding or updating a field is as follows:
```json
{
    "field_name": "Currency",
    "type": 2,
    "ui_type": "Currency",
    "property": {
        "formatter": "0", 
        "currency_code": "MOP"
    }
}
```

#### Response body example

```json
{
  "code": 0,
  "data": {
    "field": {
      "field_id": "fld6vPlPZ1",
      "field_name": "Currency",
      "is_primary": false,
      "property": {
        "currency_code": "MOP",
        "formatter": "0"
      },
      "type": 2,
      "ui_type": "Currency"
    }
  },
  "msg": "success"
}
```
  
  
### Progress field

The field type of the progress field is `2`, and the UI display type `ui_type` is `"Progress"`. Its functional attributes are as follows:

| 
Name | Type | Required | Description |
| --- | --- | --- | --- |
| `formatter` | `string` | Yes | Progress format, enumeration values are as follows: - "0": Integer, no decimal places - "0.0": Value, 1 decimal place - "0.00": Value, 2 decimal places - "0%": Percentage, integer - "0.0%": Percentage, 1 decimal place - "0.00%": Percentage, 2 decimal places |
| range_customize | bool | No | Whether to allow custom progress bar values, default is false. |
| min | number | Required when range_customize is true | Minimum value of the progress |
| max | number | Required when range_customize is true | Maximum value of the progress | #### Request body example

The request body example for adding or updating a field is as follows:
```json
{
    "field_name": "Progress",
    "type": 2,
    "ui_type": "Progress",
    "property": {
        "formatter": "0.00%", 
        "min": 0.1,
        "max": 4,
        "range_customize": true  
    }
}
```

#### Response body example

```json
{
  "code": 0,
  "data": {
    "field": {
      "field_id": "fldANhA0do",
      "field_name": "Progress",
      "is_primary": false,
      "property": {
        "formatter": "0.00%",
        "max": 100,
        "min": 0.1,
        "range_customize": true
      },
      "type": 2,
      "ui_type": "Progress"
    }
  },
  "msg": "success"
}
```
  
  
### Rating field

The type of the rating field `type` is `2`, and the UI display type `ui_type` is `"Rating"`. Its functional properties are as follows:
| 
Name | Type | Required | Description |
| --- | --- | --- | --- |
| `formatter` | `string` | Yes | The format of the rating, fixed as "0" |
| `rating` | `rating` | No | Rating settings |
|   `symbol` | `string` | No | The icon for the rating, default is "star". The enumeration values are as follows: - "star": star - "heart": heart - "thumbsup": thumbs up - "fire": fire - "smile": smiley face - "lightning": lightning - "flower": flower - "number": number |
| `min` | `number` | Yes | The minimum value of the rating, can be 0 or 1. |
| `max` | `number` | Yes | The maximum value of the rating, an integer between 1 and 10 inclusive. | #### Request body example

The request body example for adding or updating a field is as follows:
```json
{
    "field_name": "评分",
    "type": 2,
    "ui_type": "Rating",
    "property": {
        "formatter": "0",
        "min": 0,
        "max": 10,
        "rating": {
            "symbol": "lightning"
        }
    }
}
```

#### Response body example

```json
{
  "code": 0,
  "data": {
    "field": {
      "field_id": "fldTOLqnni",
      "field_name": "评分",
      "is_primary": false,
      "property": {
        "formatter": "0",
        "max": 10,
        "min": 0,
        "rating": {
          "symbol": "lightning"
        }
      },
      "type": 2,
      "ui_type": "Rating"
    }
  },
  "msg": "success"
}
```
  
### Single select field

The type of the single select field `type` is `3`. Its functional properties are as follows:
| 
Name | Type | Required | Description |
| --- | --- | --- | --- |
| `options` | `app.table.field.property.option<string>[]` | No | Option list |
|   `id` | `string` | No | Option ID. For the add field interface, there is no need to specify the option ID |
|   `name` | `string` | No | Option name |
|   `color` | `int` | No | Option color, default increments from the previous option's color. Range: 0~54 | #### Request body example for adding a field

```json
{
    "field_name": "单选",
    "type": 3,
    "property": {
        "options": [
            {
                "name": "选项 A"
            },
            {
                "name": "选项 B"
            },
            {
                "name": "选项 C",
                "color": 5
            },
            {
                "name": "选项 D"
            }
        ]
    }
}
```

#### Response body example for adding a field

```json
{
  "code": 0,
  "data": {
    "field": {
      "field_id": "fldQrcCPFs",
      "field_name": "单选",
      "is_primary": false,
      "property": {
        "options": [
          {
            "color": 0,
            "id": "optQQM0hOZ",
            "name": "选项 A"
          },
          {
            "color": 1,
            "id": "opt6wdUh3n",
            "name": "选项 B"
          },
          {
            "color": 5,
            "id": "optazha8oD",
            "name": "选项 C"
          },
          {
            "color": 6,
            "id": "optLdIy4nl",
            "name": "选项 D"
          }
        ]
      },
      "type": 3,
      "ui_type": "SingleSelect"
    }
  },
  "msg": "success"
}
```

#### Example request body for updating fields

To update the name of option A to `a++`, keep option B unchanged, delete options C and D, and add option Z, the request body is as follows:
```json
{
    "field_name": "single choice",
    "type": 3,
    "property": {
        "options": [
            {
                "color": 0,
                "id": "optQQM0hOZ",
                "name": "a++"
            },
            {
                "color": 1,
                "id": "opt6wdUh3n",
                "name": "option B"
            },
            {
                "color": 6,
                "name": "option Z"
            }
        ]
    }
}
```

#### Example response body for updating fields

Updating fields is a full update, options C and D have been deleted, the response body is as follows:
```json
{
    "code": 0,
    "data": {
        "field": {
            "field_id": "fldQrcCPFs",
            "field_name": "single choice",
            "property": {
                "options": [
                    {
                        "color": 0,
                        "id": "optQQM0hOZ",
                        "name": "a++"
                    },
                    {
                        "color": 1,
                        "id": "opt5g3xLFT",
                        "name": "option B"
                    },
                    {
                        "color": 6,
                        "id": "opt558YmTi",
                        "name": "option Z"
                    }
                ]
            },
            "type": 3
        }
    },
    "msg": "Success"
}
```
  
### Multi-select field

The type `type` of the multi-select field is `4`. Its functional properties are as follows:
| 
Name | Type | Required | Description |
| --- | --- | --- | --- |
| `options` | `app.table.field.property.option<string>[]` | No | Option list |
|   `id` | `string` | No | Option ID. For the add field interface, there is no need to specify the option ID |
|   `name` | string | No | Option name |
|   `color` | int | No | Option color, default incremented from the previous option's color. Range: 0~54 | #### Example request body for adding fields

```json
{
    "field_name": "multi-select",
    "type": 4,
    "property": {
        "options": [
            {
                "name": "a"
            },
            {
                "name": "b"
            },
            {
                "name": "c",
                "color": 5
            },
            {
                "name": "d"
            }
        ]
    }
}
```

#### Example response body for adding fields

```json
{
  "code": 0,
  "data": {
    "field": {
      "field_id": "fldLyDzkdA",
      "field_name": "multi-select",
      "is_primary": false,
      "property": {
        "options": [
          {
            "color": 0,
            "id": "opt2H40Z1F",
            "name": "a"
          },
          {
            "color": 1,
            "id": "optYQrPhiq",
            "name": "b"
          },
          {
            "color": 5,
            "id": "opt0xiSsy9",
            "name": "c"
          },
          {
            "color": 6,
            "id": "optNudypoa",
            "name": "d"
          }
        ]
      },
      "type": 4,
      "ui_type": "MultiSelect"
    }
  },
  "msg": "success"
}
```

#### Example request body for updating fields

To update the option name of `a` to `a++`, keep option `b` unchanged, delete options `c` and `d`, and add option `z`, the request body is as follows:
```json
{
    "field_name": "单选",
    "type": 3,
    "property": {
        "options": [
            {
                "color": 0,
                "id": "optpeuQVqp",
                "name": "a++"
            },
            {
                "color": 1,
                "id": "opt5g3xLFT",
                "name": "b"
            },
            {
                "color": 6,
                "name": "z"
            }
        ]
    }
}
```
  
#### Example response body for updating fields

Updating fields is a full update, options `c` and `d` have been deleted, the response body is as follows:
```json
{
    "code": 0,
    "data": {
        "field": {
            "field_id": "fld2RxOyB8",
            "field_name": "single choice",
            "property": {
                "options": [
                    {
                        "color": 0,
                        "id": "optpeuQVqp",
                        "name": "a++"
                    },
                    {
                        "color": 1,
                        "id": "opt5g3xLFT",
                        "name": "b"
                    },
                    {
                        "color": 6,
                        "id": "opt558YmTi",
                        "name": "z"
                    }
                ]
            },
            "type": 3
        }
    },
    "msg": "Success"
}
```
  
### Date field

The type `type` of the date field is `5`. Its functional properties are as follows:
| Name | Type | Required | Description |
| --- | --- | --- | --- |
| date_formatter | string | No | Date format, default is "yyyy/MM/dd", enumeration values are as follows: - "yyyy/MM/dd": format like 2021/01/30 - "yyyy-MM-dd HH:mm"：2021-01-30 14:00 - "MM-dd"：01-30 - "MM/dd/yyyy"：01/30/2021 - "dd/MM/yyyy"：30/01/2021 |
| auto_fill | boolean | No | For new records, whether to automatically fill in the creation time. Default is false. | #### Example request body

The example request body for adding and updating fields is as follows:
```json
{
    "field_name": "date",
    "type": 5,
    "property": {
        "date_formatter": "yyyy/MM/dd HH:mm",
        "auto_fill": false
    }
}
```

#### Example response body

```json
{
    "code": 0,
    "data": {
        "field": {
            "field_id": "fldHBDkAfH",
            "field_name": "date",
            "property": {
                "auto_fill": false,
                "date_formatter": "yyyy/MM/dd HH:mm"
            },
            "type": 5
        }
    },
    "msg": "Success"
}
```

### Personnel field

The type `type` of the personnel field is `11`. Its functional properties are as follows:
| Name     | Type    | Required | Description                        |
| -------- | ------- | -------- | ---------------------------------- |
| multiple | boolean | No       | Whether to allow adding multiple members, default is true. | #### Example request body

The example request body for adding and updating fields is as follows:
```json
{
    "field_name": "personnel",
    "type": 11,
    "property": {
        "multiple": true
    }
}
```

#### Example response body

```json
{
    "code": 0,
    "data": {
        "field": {
            "field_id": "fldlQDzjyK",
            "field_name": "personnel",
            "property": {
                "multiple": true
            },
            "type": 11
        }
    },
    "msg": "Success"
}
```

### Unidirectional association field

The type `type` of the unidirectional association is `18`. Its functional properties are as follows:
| Name     | Type    | Required | Description                        |
| -------- | ------- | -------- | ---------------------------------- |
| multiple | boolean | No       | Whether to allow adding multiple records, default is true. |
| table_id | string  | Yes      | ID of the associated data table    | #### Example request body

The example request body for adding and updating fields is as follows:
```json
{
    "field_name": "unidirectional association",
    "type": 18,
    "property": {
        "multiple": true,
        "table_id": "tblw92ErelmCmgHc"
    }
}
```

#### Example response body

```json
{
    "code": 0,
    "data": {
        "field": {
            "field_id": "fldNdr8VNW",
            "field_name": "unidirectional association",
            "property": {
                "multiple": true,
                "table_id": "tblw92ErelmCmgHc",
                "table_name": "Table 2"
            },
            "type": 18
        }
    },
    "msg": "Success"
}
```

### Bidirectional association field

The type `type` of the bidirectional association is `21`. Its functional properties are as follows:
| Name              | Type      | Required | Description                                 |
| ---------------   | -------   | -------- | ------------------------------------------- |
| back_field_name   | string    | No       | ID of the bidirectional association field in the associated data table, default is "associated table name-field name" |
| multiple          | boolean   | No       | Whether to allow adding multiple records, default is true.               |
| table_id          | string    | Yes      | ID of the associated data table                         | #### Example request body

The example request body for adding and updating fields is as follows:
```json
{
    "field_name": "bidirectional association",
    "type": 21,
    "property": {
        "multiple": true,
        "table_id": "tblw92ErelmCmgHc",
        "back_field_name": "bidirectional association table-auto generated"
    }
}
```

#### Example response body

```json
{
    "code": 0,
    "data": {
        "field": {
            "field_id": "fldpfIDIi0",
            "field_name": "bidirectional association",
            "property": {
                "back_field_id": "fldmQGUnWh",  // ID of the bidirectional association field in the associated data table
                "back_field_name": "bidirectional association table-auto generated",
                "multiple": true,
                "table_id": "tblw92ErelmCmgHc",
                "table_name": "Table 2"   // Name of the associated data table
            },
            "type": 21
        }
    },
    "msg": "Success"
}
```
  
### Formula field

The type `type` of the formula field is `20`. Its functional properties are as follows:
| 
Name | Type | Required | Description |
| --- | --- | --- | --- |
| `formatter` | string | No | Formula format, default is empty. It can be a number, currency, or date. Optional values are as follows: - "0": integer - "0.0": one decimal place - "0.00": two decimal places - "1,000": thousand separator - "1,000.00": thousand separator (decimal) - "%": percentage - "0.00%": percentage (decimal) - "¥": RMB - "¥0.00": RMB (decimal) - "$": USD - "$0.00": USD (decimal) - "yyyy/MM/dd HH:mm": 2021/01/30 14:00 - "yyyy/MM/dd": 2021/01/30 - "yyyy-MM-dd": 2021-01-30 - "MM-dd": 01-30 |
| `formula_expression` | string | No | Formula expression. | #### Example request body

The example request body for adding or updating fields is as follows:
```json
{
    "field_name": "formula",
    "type": 20,
    "property": {
        "formatter": "0.00%",
        "formula_expression": "IF(bitable::$table[tblxxxxxxxxxxxxx].$field[fldxxxxxxx].CONTAIN(\"Lark\"),\"aaa\",\"bbb\")"
    }
}
```

#### Example response body

```json
{
    "code": 0,
    "data": {
        "field": {
            "field_id": "fldFuAdYEI",
            "field_name": "formula",
            "property": {
                "formatter": "0.00%",
                "formula_expression": "IF(bitable::$table[tblxxxxxxxxxxxxx].$field[fldxxxxxxx].CONTAIN(\"Lark\"),\"aaa\",\"bbb\")"
            },
            "type": 20
        }
    },
    "msg": "Success"
}
```

### Location field

The type `type` of the location field is `22`. Its functional properties are as follows:
| Name | Type | Required | Description |
| --- | --- | --- | --- |
| location | app.field.property.location | No | Geolocation input method |
| input_type | string | Yes | Geolocation input restriction, optional values are: - only_mobile: only allow real-time location on mobile devices - not_limit: no restriction, can input any geolocation | #### Example request body

The example request body for adding or updating fields is as follows:
```json
{
    "field_name": "location",
    "type": 22,
    "property": {
        "location": {
            "input_type": "not_limit"
        }
    }
}
```

#### Example response body

```json
{
  "code": 0,
  "data": {
    "field": {
      "field_id": "fldc7JNkVa",
      "field_name": "location",
      "is_primary": false,
      "property": {
        "location": {
          "input_type": "not_limit"
        }
      },
      "type": 22,
      "ui_type": "Location"
    }
  },
  "msg": "success"
}
```

### Group field

The type `type` of the group field is `23`. Its functional properties are as follows:
| Name     | Type    | Required | Description                        |
| -------- | ------- | -------- | ---------------------------------- |
| multiple | boolean | No       | Whether to allow adding multiple groups, default is true. | #### Example request body

The example request body for adding or updating fields is as follows:
```json
{
    "field_name": "group field - allow adding multiple groups",
    "type": 23,
    "property": {
        "multiple": true
    },
    "description": {
        "disable_sync": false,
        "text": "group field - allow adding multiple groups"
    }
}
```

#### Example response body

```json
{
    "code": 0,
    "data": {
        "field": {
            "description": {
                "disable_sync": false,
                "text": "group field - allow adding multiple groups"
            },
            "field_id": "fldw6fSubT",
            "field_name": "group field - allow adding multiple groups",
            "is_primary": false,
            "property": {
                "multiple": true
            },
            "type": 23,
            "ui_type": "GroupChat"
        }
    },
    "msg": "success"
}
```
  
### Creation time and last update time fields

The type `type` of the creation time field is `1001`. The type `type` of the last update time field is `1002`. Their functional properties are the same, as follows:
| Name | Type | Required | Description |
| --- | --- | --- | --- |
| date_formatter | string | No | Date format, default is "yyyy/MM/dd", enum values are as follows: - "yyyy/MM/dd": format like 2021/1/30 - "yyyy-MM-dd HH:mm": 2021/01/30 14:00 - "MM-dd": 01-30 - "MM/dd/yyyy": 01/30/2021 - "dd/MM/yyyy": 30/01/2021 | #### Example request body

The example request body for adding or updating the "creation time" field is as follows:
```json
{
    "field_name": "creation time",
    "type": 1001,
    "property": {
        "date_formatter": "yyyy/MM/dd"
    }
}
```

#### Example response body

```json
{
    "code": 0,
    "data": {
        "field": {
            "field_id": "fldoblwmUC",
            "field_name": "creation time",
            "property": {
                "date_formatter": "yyyy/MM/dd"
            },
            "type": 1001
        }
    },
    "msg": "Success"
}
```

### Auto number field

The type `type` of the auto number field is `1005`. Its functional properties are as follows:

| 
Name | Type | Required | Description |
| --- | --- | --- | --- |
| `auto_serial` | app.field.property.auto_serial | No | Auto number |
|   type | string | Yes | Auto number type, optional values are as follows: - "custom": custom number - "auto_increment_number": auto-increment number |
|   reformat_existing_records | bool | No | Whether to apply changes to existing numbers, default is false. |
|   options | app.field.property.auto_serial.options[] | No | Custom number rule list |
|     type | string | Yes | Rule type, optional values are as follows: - "system_number": number of digits for auto-increment number, range is 1-9 - "fixed_text": fixed characters, maximum character length is 20 - "created_time": creation date |
|     value | string | Yes | Value corresponding to the rule type. - If the rule type is `"type": "system_number"`, value is an integer in the range of 1-9, indicating the number of digits for auto-increment number - If the rule type is `"type": "fixed_text"`, value is fixed characters within 20 characters - If the rule type is "type": "created_time", value is used to specify the date format. Optional values are as follows: - "yyyyMMdd": date format is 20220130 - "yyyyMM": date format is 202201 - "yyyy": date format is 2022 - "MMdd": date format is 130, indicating January 30 - "MM": date format is 1, indicating the month - "dd": date format is 30, indicating the day | #### Example of adding "custom" auto number

To add or update a field with the "custom" type auto number, the example request body is as follows:
```json
{
    "field_name": "custom number",
    "property": {
        "auto_serial": {
            "type": "custom",
            "reformat_existing_records": true, 
            "options": [
                {
                    "type": "system_number",
                    "value": "3" 
                },
                {
                    "type": "fixed_text", 
                    "value": "keyword"
                },
                {
                    "type": "created_time",
                    "value": "yyyyMMdd"
                }
            ]
        }
    },
    "type": 1005
}
```
The corresponding example response body is as follows:
```json
{
    "code": 0,
    "msg": "success",
    "data": {
        "field": {
            "property": {
                "auto_serial": {
                    "type": "custom",
                    "options": [
                        {
                            "type": "system_number",
                            "value": "3"
                        },
                        {
                            "value": "keyword",
                            "type": "fixed_text"
                        },
                        {
                            "type": "created_time",
                            "value": "yyyyMMdd"
                        }
                    ]
                }
            },
            "field_id": "fldmVunQuc",
            "field_name": "custom number",
            "type": 1005
        }
    }
}
```

#### Example of adding "auto_increment_number" auto number

To add or update a field with the "auto_increment_number" type auto number, the example request body is as follows:
```json
{
    "field_name": "auto-increment number auto number",
    "property": {
        "auto_serial": {
            "type": "auto_increment_number",
            "reformat_existing_records": true
        }
    },
    "type": 1005
}
```
The corresponding example response body is as follows:
```json
{
    "code": 0,
    "msg": "success",
    "data": {
        "field": {
            "field_name": "auto-increment number auto number",
            "type": 1005,
            "property": {
                "auto_serial": {
                    "type": "auto_increment_number"
                }
            },
            "field_id": "fldwq16vz2"
        }
    }
}
```
