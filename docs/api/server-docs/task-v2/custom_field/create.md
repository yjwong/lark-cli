---
title: "Create Custom Field"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/task-v2/custom_field/create"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/task/v2/custom_fields"
service: "task-v2"
resource: "custom_field"
section: "Tasks v2"
rate_limit: "100 per minute"
scopes:
  - "task:custom_field:write"
field_scopes:
  - "contact:user.employee_id:readonly"
updated: "1699521749000"
---

# Create Custom Field

Create a custom field and add it to a resource (currently resources only support "tasklist"). To create a custom field, you must provide the field name, type and corresponding field settings.

Currently task custom fields support several types: number, member, datetime, single_select and multi_select. Use "number_setting", "member_setting", "datetime_setting", "single_select_setting", "multi_select_setting", "text_setting" to set their setting correspondingly.

For example, to create a custom field of number type and add it to the tasklist with guid "ec5ed63d-a4a9-44de-a935-7ba243471c0a", you can invoke list this:

```
POST /task/v2/custom_fields
{
  "name": "price",
  "type": "number",
  "resource_type": 
  "tasklist",
  "resource_id": 
  "ec5ed63d-a4a9-44de-a935-7ba243471c0a",
  "number_setting": {
    "format": "usd",
    "decimal_count": 2,
    "separator": "thousand"
  }
}
```

which means to create a custom field named as "price" with 2 decimal palces. When displayed on the app UI, it is in USD format and displays the thousand separator.

Similarly, to create a single_select custom field, you may invoke api like this:

```
POST /task/v2/custom_fields
{
  "name": "priority",
  "type": "single_select",
  "resource_type": "tasklist",
  "resource_id":  "ec5ed63d-a4a9-44de-a935-7ba243471c0a",
  "single_select_setting": {
    "options": [
       {
          "name": "high",
          "color_index": 1
       },
       {
          "name": "mid",
          "color_index": 11
       },
       {
          "name": "low",
          "color_index": 16
       }
   ]
  }
}
```

which means creating a "priority" single select custom field with 3 options: high, mid, low，with color set to each of them.

> Creating custom fields on a resource requires edit permissions for that resource.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/task/v2/custom_fields |
| HTTP Method | POST |
| Rate Limit | 100 per minute |
| Supported app types | custom,isv |
| Required scopes | `task:custom_field:write` |
| Required field scopes | > The response body of the API contains the following sensitive fields, and they will be returned only after corresponding scopes are added. If you do not need the fields, it is not recommended that you request the scopes. `contact:user.employee_id:readonly` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `user_id_type` | `string` | No | User ID categories **Example value**: open_id **Optional values are**:  - `open_id`: Identifies a user to an app. The same user has different Open IDs in different apps. How to get Open ID - `union_id`: Identifies a user to a tenant that acts as a developer. A user has the same Union ID in apps developed by the same developer, and has different Union IDs in apps developed by different developers. A developer can use Union ID to link the same user's identities in multiple apps.How to get Union ID - `user_id`: Identifies a user to a tenant. The same user has different User IDs in different tenants. In one single tenant, a user has the same User ID in all apps （including store apps）. User ID is usually used to communicate user data between different apps. How to get User ID  **Default value**: `open_id` **When the value is `user_id`, the following field scopes are required**: `contact:user.employee_id:readonly` | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `resource_type` | `string` | Yes | Customize the resource type to which the field belongs, support "tasklist" **Example value**: "Tasklist" **Default value**: `tasklist` |
| `resource_id` | `string` | Yes | The resource ID to which the custom field belongs, the GUID of the list must be filled in when the resource_type is "tasklist". **Example value**: "ec5ed63d-a4a9-44de-a935-7ba243471c0a" **Data validation rules**: - Maximum length: `100` characters |
| `name` | `string` | Yes | Field name, up to 50 characters. **Example value**: "priority" |
| `type` | `string` | Yes | custom field type **Example value**: "number" **Optional values are**:  - `number`: number - `datetime`: datetime - `member`: member - `single_select`: single select - `multi_select`: multiple select - `text`: text  |
| `number_setting` | `number_setting` | No | number field setting |
|   `format` | `string` | No | number format in which number value are displayed in the app format. Note that this setting only affects the display format of field values, and does not affect the format of field values input/output in openapi. **Example value**: "normal" **Optional values are**:  - `normal`: regular number - `percentage`: percentage format - `cny`: RMB format - `usd`: USD format - `custom`: custom symbol  **Default value**: `normal` |
|   `custom_symbol` | `string` | No | When'format 'is set to "custom", set the specific custom symbol characters, supporting up to 4 characters. Note that this setting only affects the display format of field values, and does not affect the format of field values input/output in openapi. **Example value**: "custom symbol" |
|   `custom_symbol_position` | `string` | No | When'format 'is set to "custom", the display position of the custom symbol relative to the number is defined. Note that this setting only affects the display format of field values, and does not affect the format of field values input/output in openapi. **Example value**: "left" **Optional values are**:  - `left`: Custom symbols are displayed to the left of the numbers - `right`: Custom symbols are displayed to the right of the number  **Default value**: `right` |
|   `separator` | `string` | No | number separator in the integer part. Note that this setting only affects the display format of field values, and does not affect the format of field values input/output in openapi. **Example value**: "thousand" **Optional values are**:  - `none`: no separator - `thousand`: thousand separator  **Default value**: `none` |
|   `decimal_count` | `int` | No | The number of decimal places reserved for the value of a custom field. Extra digits will be rounded. Default is 0. **Example value**: 2 **Default value**: `0` **Data validation rules**: - Value range: `0` ～ `6` |
| `member_setting` | `member_setting` | No | member field setting |
|   `multi` | `boolean` | No | Whether to support multiple selection. Defaults to false. **Example value**: true **Default value**: `false` |
| `datetime_setting` | `datetime_setting` | No | datetime setting |
|   `format` | `string` | No | Date time format, supported   Year, month and day separated by dashes, e.g. 2023-08-24  Year, month and day separated by slashes, e.g. 2023/08/04  months, days and years separated by slashes, e.g. 08/24/2023 Day, month and year separated by slashes, e.g. 24/08/2023  Default is "yyyy-mm-dd". Note that this setting only affects the display format of field values, and does not affect the format of field values input/output in openapi. **Example value**: "yyyy/mm/dd" |
| `single_select_setting` | `select_setting` | No | single select setting |
|   `options` | `option[]` | No | single select options **Data validation rules**: - Length range: `0` ～ `100` |
|     `name` | `string` | Yes | Option name, cannot be empty, supports up to 50 characters **Example value**: "high" |
|     `color_index` | `int` | No | The color index value of the option, ranging from 0 to 54. If not filled in, it will automatically randomly select one of the unused color index values. **Example value**: 1 **Data validation rules**: - Value range: `0` ～ `54` |
|     `is_hidden` | `boolean` | No | Whether the option is hidden. The hidden option is not visible in the interface, and the field value can no longer be set to this option through openapi. **Example value**: false |
| `multi_select_setting` | `select_setting` | No | multiple select setting |
|   `options` | `option[]` | No | multi select options **Data validation rules**: - Length range: `0` ～ `100` |
|     `name` | `string` | Yes | Option name, cannot be empty, up to 50 characters. **Example value**: "high" |
|     `color_index` | `int` | No | The color index value of the option, which can be a number from 0 to 54. If not filled in, one will be randomly selected. **Example value**: 1 **Data validation rules**: - Value range: `0` ～ `54` |
|     `is_hidden` | `boolean` | No | Whether the option is hidden. The hidden option is not visible in the App interface, and the field value cannot be set to this option through oapi. **Example value**: false |
| `text_setting` | `text_setting` | No | text field setting | ### Request body example

{
    "resource_type": "Tasklist",
    "resource_id": "ec5ed63d-a4a9-44de-a935-7ba243471c0a",
    "name": "priority",
    "type": "number",
    "number_setting": {
        "format": "normal",
        "custom_symbol": "custom symbol",
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
                "name": "high",
                "color_index": 1,
                "is_hidden": false
            }
        ]
    },
    "multi_select_setting": {
        "options": [
            {
                "name": "high",
                "color_index": 1,
                "is_hidden": false
            }
        ]
    },
    "text_setting": {}
}

## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `custom_field` | `custom_field` | Created custom fields |
|     `guid` | `string` | GUID for custom fields |
|     `name` | `string` | Custom field names |
|     `type` | `string` | Custom field type |
|     `number_setting` | `number_setting` | number field setting |
|       `format` | `string` | Format of digital display **Optional values are**:  - `normal`: regular number - `percentage`: percentage format - `cny`: RMB format - `usd`: USD format - `custom`: Custom Symbol  |
|       `custom_symbol` | `string` | Custom symbols. Only takes effect when'format 'is set to custom. |
|       `custom_symbol_position` | `string` | Customize the position of the symbol. Defaults to right. **Optional values are**:  - `left`: Custom symbols are placed to the left of the numbers - `right`: Custom symbols are placed to the right of the number  |
|       `separator` | `string` | separator style in the integer part **Optional values are**:  - `none`: no separator - `thousand`: thousand separator  |
|       `decimal_count` | `int` | The number of decimal places. If the entered number value has more decimal places than this setting, the extra digits will be rounded and discarded. If'format 'is "percentage", it means the number of decimal places after becoming a percentage. |
|     `member_setting` | `member_setting` | field settings for member types |
|       `multi` | `boolean` | Whether to support multiple selection |
|     `datetime_setting` | `datetime_setting` | field settings for time and date types |
|       `format` | `string` | Datetime display format |
|     `single_select_setting` | `select_setting` | Field settings for radio types |
|       `options` | `option[]` | option |
|         `guid` | `string` | GUID of options |
|         `name` | `string` | Option name, cannot be empty, maximum 50 characters |
|         `color_index` | `int` | The color index value of the option, ranging from 0 to 54. If not filled in, it will automatically randomly select one of the unused color index values. |
|         `is_hidden` | `boolean` | Whether the option is hidden. The hidden option is not visible in the interface, and the field value can no longer be set to this option through openapi. |
|     `multi_select_setting` | `select_setting` | Field settings for multiple selection types |
|       `options` | `option[]` | option |
|         `guid` | `string` | GUID of options |
|         `name` | `string` | Option name, cannot be empty, maximum 50 characters. |
|         `color_index` | `int` | The color index value of the option. |
|         `is_hidden` | `boolean` | Whether the option is hidden. The hidden option is not visible in the interface, and the field value can no longer be set to this option through openapi. |
|     `creator` | `member` | founder |
|       `id` | `string` | Indicates the id of member |
|       `type` | `string` | Type of member |
|       `role` | `string` | Member role |
|     `created_at` | `string` | Custom field created timestamp (ms) |
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
                "custom_symbol": "Custom Symbol",
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
| 400 | 1470400 | Parameter error, such as setting the custom field type to "number", but not providing "number_setting". | The reason for the error is shown in the returned msg prompt. |
| 403 | 1470403 | Missing permission to create custom fields in the specified resource | Check whether the calling identity has edit permission for the corresponding resource. |
| 404 | 1470404 | The resource to create the custom field does not exist or has been deleted. | Verify that the field for which you want to create the custom field still exists. |
| 500 | 1470500 | Server error. | Try retrying the call. If it continues to fail, please contact support for feedback. |
