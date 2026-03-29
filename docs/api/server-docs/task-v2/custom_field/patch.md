---
title: "Patch Custom Field"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/task-v2/custom_field/patch"
method: "PATCH"
api_path: "https://open.larksuite.com/open-apis/task/v2/custom_fields/:custom_field_guid"
service: "task-v2"
resource: "custom_field"
section: "Tasks v2"
rate_limit: "100 per minute"
scopes:
  - "task:custom_field:write"
field_scopes:
  - "contact:user.employee_id:readonly"
updated: "1699521757000"
---

# Update Custom Field

Update the name and settings of a custom field. When updating, fill in all the task field names to be updated in the `update_fields`, and fill in the new value in the `custom_field`. Custom field type is not allowed to change. Only their settings can be modified according to the type.

`update_fields` supports:

* `name`: custom field name
* `number_setting`: number type setting (if and only if the custom field type to update is number)
* `member_setting`: member setting (if and only if the custom field type to update is member)
* `datetime_setting`: datetime type setting (if and only if the custom field type to update is datetime)
* `single_select_setting`: single select setting (if and only if the custom field type to update is single select)
* `multi_select_setting`: multi_select type setting (if and only if the custom field type to be updated is multi_select)
* `text_setting`: text type setting (currently nothing can be set)

When updating a setting, if you do not fill in a field, it means that the original setting is not changed. For example, for a number, the original settings are:
```json
"number_setting": {
  "format": "normal",
  "decimal_count": 2,
  "separator": "none",
  "custom_symbol": "€", 
  "custom_symbol_position": "right"
}
```

Invoke the API with the following parameters:

```
PATCH /task/v2/custom_fields/:custom_field_guid
{
  "custom_field": {
    "number_setting": {
      "decimal_count": 4
    }
  },
  "update_fields": ["number_setting"]
}
```

indicates that only the number of decimal places is changed from 2 to 4, and the rest of the settings 'format', 'separator ',' custom_symbol', etc. are unchanged.

For custom fields of single/multi_select type, the setting is a list of options. When updated, the usage is just like configure the options in App UI. Instead of passing in all the options for the field, the user only needs to provide the options that eventually want the UI to be visible (is_hidden = false). Options in the original field that do not appear in the input will be set to is_hidden = true and placed after all visible options.

For an updated option, if `option_guid` is provided, it will be taken as option updating (in this case guid must exist in the current field, otherwise a 400 error will be returned); if GUID not provided, it will be regarded as creating a new option (the new option option_guid will be returned in reponse).

For example, a single_select field originally had 3 options A, B, C, D. Where C is hidden. The user can update the options like this:

```
PATCH /task/v2/custom_fields/:custom_field_guid
{
   "custom_field": {
      "single_select_setting": {
         "options": [
            {
               "name": "E",
               "color_index": 25
            },
            {
               "guid": ""
               "name": "A2"
            },
            {
               "guid": "",
            },
         ]
      }
   },
   "update_fields": ["single_select_setting"]
}
```

After calling, you finally get a new list of options E, A, C, B, D, where:

* option E is created and its color_index is set to 25.
* option A is updated and its name is changed to "A2". But its color_index remains the same because it is not set;
* overall order of options follows the user's input order, that is, E, A, C, just as the direct input, and their is_hidden are set to false. Note, C was originally is_hidden = true, it will also be set to is_hidden = false.
* options B and D have their is_hidden set to true because the user did not input them, and are placed after all options input in the request.

If you simply want to modify the order of options visible to the user, such as changing from the original options A, B, C to C, B, A, you can invoke the API like this:
```
PATCH /task/v2/custom_fields/:custom_field_guid
{
   "custom_field": {
      "single_select_setting": {
         "options": [
            {
               "guid": ""
            },
            {
               "guid": ""
            },
            {
               "guid": "",
            },
         ]
      }
   },
  "update_fields": ["single_select_setting"]
}
```

If you want to directly mark all options in the field as invisible, you can invoke the API like this:
```
PATCH /task/v2/custom_fields/:custom_field_guid
{
  "custom_field": {
      "single_select_setting": {
         "options": []
      }
   },
  "update_fields": ["single_select_setting"]
}
```

The option to update the single/multi_select field must meet the "visible option name cannot be duplicated" constraint. Otherwise, an error will be returned. Developers need to ensure that the option names entered contain on duplicated items.

If you want to update only one single option, or want to set the is_hidden of an option separately, this API cannot do that. But you can use the Update Custom Field Options API to do it.

> Updating a custom field requires edit permissions to the custom field.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/task/v2/custom_fields/:custom_field_guid |
| HTTP Method | PATCH |
| Rate Limit | 100 per minute |
| Supported app types | custom,isv |
| Required scopes | `task:custom_field:write` |
| Required field scopes | > The response body of the API contains the following sensitive fields, and they will be returned only after corresponding scopes are added. If you do not need the fields, it is not recommended that you request the scopes. `contact:user.employee_id:readonly` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `custom_field_guid` | `string` | custom field GUID, which can be created byCreate Custom Field, or queried byList Custom Field. **Example value**: "5ffbe0ca-6600-41e0-a634-2b38cbcf13b8" | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `user_id_type` | `string` | No | User ID categories **Example value**: open_id **Optional values are**:  - `open_id`: Identifies a user to an app. The same user has different Open IDs in different apps. How to get Open ID - `union_id`: Identifies a user to a tenant that acts as a developer. A user has the same Union ID in apps developed by the same developer, and has different Union IDs in apps developed by different developers. A developer can use Union ID to link the same user's identities in multiple apps.How to get Union ID - `user_id`: Identifies a user to a tenant. The same user has different User IDs in different tenants. In one single tenant, a user has the same User ID in all apps （including store apps）. User ID is usually used to communicate user data between different apps. How to get User ID  **Default value**: `open_id` **When the value is `user_id`, the following field scopes are required**: `contact:user.employee_id:readonly` | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `custom_field` | `input_custom_field` | No | custom field data to update |
|   `name` | `string` | No | field name, up to 50 characters **Example value**: "priority" |
|   `number_setting` | `number_setting` | No | field number setting |
|     `format` | `string` | No | number display format **Example value**: "normal" **Optional values are**:  - `normal`: regular number - `percentage`: percentage format - `cny`: RMB format - `usd`: USD format - `custom`: custom symbol format  **Default value**: `normal` |
|     `custom_symbol` | `string` | No | Custom symbol charaters, supporting up to 4 characters. Only takes effect when `format` is set to "custom". **Example value**: "€" |
|     `custom_symbol_position` | `string` | No | Customize where symbols are displayed. **Example value**: "left" **Optional values are**:  - `left`: Custom symbols are placed to the left of the numbers - `right`: Custom symbols are placed to the right of the number  **Default value**: `right` |
|     `separator` | `string` | No | separator in the integer part **Example value**: "thousand" **Optional values are**:  - `none`: no separator - `thousand`: thousand separator  **Default value**: `none` |
|     `decimal_count` | `int` | No | The number of decimal places. If the entered number value has more decimal places than this setting, the extra digits will be rounded and discarded. If'format 'is "percentage", it means the number of decimal places after becoming a percentage. **Example value**: 2 **Default value**: `0` **Data validation rules**: - Value range: `0` ～ `6` |
|   `member_setting` | `member_setting` | No | member field setting |
|     `multi` | `boolean` | No | Whether to support multiple selection **Example value**: true **Default value**: `false` |
|   `datetime_setting` | `datetime_setting` | No | datetime field setting |
|     `format` | `string` | No | Datetime display format, which supports:  Year, month and day separated by dashes, e.g. 2023-08-24  - `yyyy/mm/dd`: Year, month and day separated by slashes, e.g. 2023/08/04 - `mm/dd/yyyy`: Months, days and years separated by slashes, e.g. 08/24/2023 - `dd/mm/yyyy`: Day, month and year separated by slashes, e.g. 24/08/2023  **Example value**: "yyyy/mm/dd" |
|   `single_select_setting` | `select_setting` | No | single_select setting |
|     `options` | `option[]` | No | single_select options **Data validation rules**: - Length range: `0` ～ `100` |
|       `guid` | `string` | No | GUID of the option. If filled in means updating the existing option; if not filled in means creating a new option. **Example value**: "4216f79b-3fda-4dc6-a0c4-a16022e47152" |
|       `name` | `string` | No | Option name,  maximum 50 characters **Example value**: "high" |
|       `color_index` | `int` | No | The color index value of the option, which can be a number from 0 to 54. **Example value**: 1 **Data validation rules**: - Value range: `0` ～ `54` |
|   `multi_select_setting` | `select_setting` | No | mult_select setting |
|     `options` | `option[]` | No | mult_select options **Data validation rules**: - Length range: `0` ～ `100` |
|       `guid` | `string` | No | GUID of the option. If filled in means updating the existing option; if not filled in means creating a new option. **Example value**: "4216f79b-3fda-4dc6-a0c4-a16022e47152" |
|       `name` | `string` | No | Option name,  maximum 50 characters **Example value**: "high" |
|       `color_index` | `int` | No | The color index value of the option, which can be a number from 0 to 54. **Example value**: 1 **Data validation rules**: - Value range: `0` ～ `54` |
|   `text_setting` | `text_setting` | No | text field setting |
| `update_fields` | `string[]` | No | The custom field type to modify supports: * `name`: custom field name. * `number_setting`: number setting (if and only if the custom field type to update is numeric) * `member_setting`: people type setting (if and only if the custom field type to update is member) * `datetime_setting`: Date type setting (if and only if the custom field type to update is date) * `single_select_setting`: single select setting (if and only if the custom field type to update is single_select) * `multi_select_setting`: multi select type setting (if and only if the custom field type to be updated is multi_select) * `text_setting`: text type setting (currently nothing can be set) **Example value**: ["name"] **Data validation rules**: - Length range: `1` ～ `20` | ### Request body example

{
    "custom_field": {
        "name": "priority",
        "number_setting": {
            "format": "normal",
            "custom_symbol": "€",
            "custom_symbol_position": "left",
            "separator": "thousand",
            "decimal_count": 2
        },
        "member_setting": {
            "multi": true
        },
        "datetime_setting": {
            "format": "yyyy/mm/dd"
        },
        "single_select_setting": {
            "options": [
                {
                    "guid": "4216f79b-3fda-4dc6-a0c4-a16022e47152",
                    "name": "high",
                    "color_index": 1
                }
            ]
        },
        "multi_select_setting": {
            "options": [
                {
                    "guid": "4216f79b-3fda-4dc6-a0c4-a16022e47152",
                    "name": "high",
                    "color_index": 1
                }
            ]
        },
        "text_setting": {}
    },
    "update_fields": [
        "name"
    ]
}

## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `custom_field` | `custom_field` | updated custom field settings |
|     `guid` | `string` | GUID for custom field |
|     `name` | `string` | Custom field name |
|     `type` | `string` | Custom field type |
|     `number_setting` | `number_setting` | number setting |
|       `format` | `string` | number display format **Optional values are**:  - `normal`: regular number - `percentage`: percentage format - `cny`: RMB format - `usd`: USD format - `custom`: custom symbol  |
|       `custom_symbol` | `string` | Custom symbols. Only takes effect when `format` is set to custom. |
|       `custom_symbol_position` | `string` | Customize where symbols are displayed. **Optional values are**:  - `left`: Custom symbols are placed to the left of the numbers - `right`: Custom symbols are placed to the right of the number  |
|       `separator` | `string` | Separator format **Optional values are**:  - `none`: No separator - `thousand`: thousand separator  |
|       `decimal_count` | `int` | The number of decimal places. If the entered number value has more decimal places than this setting, the extra digits will be rounded and discarded. If `format` is "percentage", it means the number of decimal places after becoming a percentage. |
|     `member_setting` | `member_setting` | member field setting |
|       `multi` | `boolean` | Whether to support multiple selection |
|     `datetime_setting` | `datetime_setting` | datetime display format |
|       `format` | `string` | Datetime display format, which supports:  Year, month and day separated by dashes, e.g. 2023-08-24  - `yyyy/mm/dd`: Year, month and day separated by slashes, e.g. 2023/08/04 - `mm/dd/yyyy`: Months, days and years separated by slashes, e.g. 08/24/2023 - `dd/mm/yyyy`: Day, month and year separated by slashes, e.g. 24/08/2023  |
|     `single_select_setting` | `select_setting` | single_select setting |
|       `options` | `option[]` | single select options |
|         `guid` | `string` | option GUID |
|         `name` | `string` | Option name, maximum 50 characters |
|         `color_index` | `int` | The color index value of the option, which can be a number from 0 to 54. |
|         `is_hidden` | `boolean` | Whether the option is hidden. The hidden option is not visible in the UI, and the field value can no longer be set to this option through openapi. |
|     `multi_select_setting` | `select_setting` | multi select setting |
|       `options` | `option[]` | multi select options |
|         `guid` | `string` | option GUID |
|         `name` | `string` | Option name, maximum 50 characters |
|         `color_index` | `int` | The color index value of the option, which can be a number from 0 to 54. |
|         `is_hidden` | `boolean` | Whether the option is hidden. The hidden option is not visible in the UI, and the field value can no longer be set to this option through openapi. |
|     `creator` | `member` | field creator |
|       `id` | `string` | member id |
|       `type` | `string` | member type |
|       `role` | `string` | member role |
|     `created_at` | `string` | Custom field create timestamp (ms) |
|     `updated_at` | `string` | Custom field update timestamp (ms) |
|     `text_setting` | `text_setting` | text field setting | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
        "custom_field": {
            "guid": "34d4b29f-3d58-4bc5-b752-6be80fb687c8",
            "name": "priority",
            "type": "number",
            "number_setting": {
                "format": "normal",
                "custom_symbol": "€",
                "custom_symbol_position": "left",
                "separator": "thousand",
                "decimal_count": 2
            },
            "member_setting": {
                "multi": true
            },
            "datetime_setting": {
                "format": "yyyy/mm/dd"
            },
            "single_select_setting": {
                "options": [
                    {
                        "guid": "4216f79b-3fda-4dc6-a0c4-a16022e47152",
                        "name": "high",
                        "color_index": 1,
                        "is_hidden": false
                    }
                ]
            },
            "multi_select_setting": {
                "options": [
                    {
                        "guid": "4216f79b-3fda-4dc6-a0c4-a16022e47152",
                        "name": "high",
                        "color_index": 1,
                        "is_hidden": false
                    }
                ]
            },
            "creator": {
                "id": "ou_2cefb2f014f8d0c6c2d2eb7bafb0e54f",
                "type": "user",
                "role": "creator"
            },
            "created_at": "1688196600000",
            "updated_at": "1688196600000",
            "text_setting": {}
        }
    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 1470400 | Incorrect request parameters, such as passing invalid field name to `update_fields`. | The reason for the error is shown in the returned msg prompt. |
| 403 | 1470403 | Missing permission to edit custom fields. | Confirm that the calling identity has permission to edit the custom field. |
| 404 | 1470404 | The custom field to edit does not exist or has been deleted. | Confirm that the custom field to edit does not exist or has been deleted. |
| 500 | 1470500 | Server error. | Try retrying the call. If it continues to fail, please contact support for feedback. |
