---
title: "List Custom Field"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/task-v2/custom_field/list"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/task/v2/custom_fields"
service: "task-v2"
resource: "custom_field"
section: "Tasks v2"
rate_limit: "100 per minute"
scopes:
  - "task:custom_field:read"
  - "task:custom_field:write"
field_scopes:
  - "contact:user.employee_id:readonly"
updated: "1699521761000"
---

# List Custom Fields

Get a list of custom fields accessible to the calling identity. If the resource_type and resource_id parameters are not provided, all custom fields accessible to the calling identity are returned.

If resource_type and resource_id are provided, the custom fields under that resource are returned. Currently resource_type only supports "tasklist", in which case resource_id should be a tasklist GUID.

This API supports paging.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/task/v2/custom_fields |
| HTTP Method | GET |
| Rate Limit | 100 per minute |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `task:custom_field:read` `task:custom_field:write` |
| Required field scopes | > The response body of the API contains the following sensitive fields, and they will be returned only after corresponding scopes are added. If you do not need the fields, it is not recommended that you request the scopes. `contact:user.employee_id:readonly` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `page_size` | `int` | No | paging size **Example value**: 50 **Default value**: `50` **Data validation rules**: - Value range: `1` ～ `100` |
| `page_token` | `string` | No | Page identifier. It is not filled in the first request, indicating traversal from the beginning; when there will be more groups, the new page_token will be returned at the same time, and the next traversal can use the page_token to get more groups **Example value**: aWQ9NzEwMjMzMjMxMDE= |
| `user_id_type` | `string` | No | User ID categories **Example value**: open_id **Optional values are**:  - `open_id`: Identifies a user to an app. The same user has different Open IDs in different apps. How to get Open ID - `union_id`: Identifies a user to a tenant that acts as a developer. A user has the same Union ID in apps developed by the same developer, and has different Union IDs in apps developed by different developers. A developer can use Union ID to link the same user's identities in multiple apps.How to get Union ID - `user_id`: Identifies a user to a tenant. The same user has different User IDs in different tenants. In one single tenant, a user has the same User ID in all apps （including store apps）. User ID is usually used to communicate user data between different apps. How to get User ID  **Default value**: `open_id` **When the value is `user_id`, the following field scopes are required**: `contact:user.employee_id:readonly` |
| `resource_type` | `string` | No | Resource types, such as provide custom fields that indicate that only specific resources are queried. Currently only "tasklist" is supported. **Example value**: tasklist |
| `resource_id` | `string` | No | To query the attribution resource_id of a custom field **Example value**: 5ffbe0ca-6600-41e0-a634-2b38cbcf13b8 | ## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `items` | `custom_field[]` | List of custom fields returned |
|     `guid` | `string` | custom field GUID |
|     `name` | `string` | Custom field name |
|     `type` | `string` | Custom field type |
|     `number_setting` | `number_setting` | number setting |
|       `format` | `string` | number display format **Optional values are**:  - `normal`: regular number - `percentage`: percentage format - `cny`: RMB format - `usd`: USD format - `custom`: custom symbol  |
|       `custom_symbol` | `string` | Custom symbol. Only takes effect when'format 'is set to custom. |
|       `custom_symbol_position` | `string` | Customize where symbols are displayed. **Optional values are**:  - `left`: Custom symbols are placed to the left of the numbers - `right`: Custom symbols are placed to the right of the number  |
|       `separator` | `string` | separator format in the integer part **Optional values are**:  - `none`: No separator - `thousand`: Thousand separator  |
|       `decimal_count` | `int` | The number of decimal places. If the entered number value has more decimal places than this setting, the extra digits will be rounded and discarded. If'format 'is "percentage", it means the number of decimal places after becoming a percentage. |
|     `member_setting` | `member_setting` | member field setting |
|       `multi` | `boolean` | Whether to support multiple selection |
|     `datetime_setting` | `datetime_setting` | datetime field setting |
|       `format` | `string` | datetime display format |
|     `single_select_setting` | `select_setting` | single select setting |
|       `options` | `option[]` | single select options |
|         `guid` | `string` | option GUID |
|         `name` | `string` | Option name |
|         `color_index` | `int` | The color index value of the option, which can be a number from 0 to 54. If not filled in, one will be randomly selected. |
|         `is_hidden` | `boolean` | Whether the option is hidden. The hidden option is not visible in the interface, and the field value can no longer be set to this option through openapi. |
|     `multi_select_setting` | `select_setting` | multi select setting |
|       `options` | `option[]` | multi select options |
|         `guid` | `string` | option GUID |
|         `name` | `string` | Option name, maximum 50 characters |
|         `color_index` | `int` | The color index value of the option, which can be a number from 0 to 54. If not filled in, one will be randomly selected. |
|         `is_hidden` | `boolean` | Whether the option is hidden. The hidden option is not visible in the interface, and the field value can no longer be set to this option through openapi. |
|     `creator` | `member` | field creator |
|       `id` | `string` | member ID |
|       `type` | `string` | member type |
|       `role` | `string` | member role |
|     `created_at` | `string` | Custom field create timestamp (ms) |
|     `updated_at` | `string` | Custom field update timestamp (ms) |
|     `text_setting` | `text_setting` | text field setting |
|   `page_token` | `string` | Page identifier, when has_more is true, a new page_token will also be returned. Otherwise, page_token will not be returned |
|   `has_more` | `boolean` | Whether the response body has more parameters | ### Response body example

{
    "code": 0,
    "msg": "Success",
    "data": {
        "items": [
            {
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
        ],
        "page_token": "aWQ9NzEwMjMzMjMxMDE=",
        "has_more": true
    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 1470400 | Parameter errors, such as page_size fill in a negative number. | The reason for the error is shown in the returned msg prompt. |
| 404 | 1470404 | The resource to access does not exist or has been deleted. | Confirm that the resource you want to access does not exist or has been deleted. |
| 500 | 1470500 | Server error. | Try retrying the call. If it continues to fail, please contact support for feedback. |
