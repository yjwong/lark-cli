---
title: "Create Custom Field Option"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/task-v2/custom_field-option/create"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/task/v2/custom_fields/:custom_field_guid/options"
service: "task-v2"
resource: "custom_field-option"
section: "Tasks v2"
rate_limit: "100 per minute"
scopes:
  - "task:tasklist:read"
  - "task:tasklist:write"
updated: "1699521772000"
---

# Create Custom Field Option

Add a custom option to single/multi select fields. A single/multi select field supports up to 100 options.

If the newly added option is visible(is_hidden=false) , its name cannot be duplicated with the name of a visible option.

> Edit permissions on custom fields are required.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/task/v2/custom_fields/:custom_field_guid/options |
| HTTP Method | POST |
| Rate Limit | 100 per minute |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `task:tasklist:read` `task:tasklist:write` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `custom_field_guid` | `string` | custom field GUID to create option **Example value**: "b13adf3c-cad6-4e02-8929-550c112b5633" | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `name` | `string` | Yes | option name, up to 50 characters. **Example value**: "high" |
| `color_index` | `int` | No | Color index value, supports a number from 0 to 54. If not filled in, one will be randomly selected. **Example value**: 10 **Data validation rules**: - Value range: `0` ～ `54` |
| `insert_before` | `string` | No | option_guid to put before an option **Example value**: "2bd905f8-ef38-408b-aa1f-2b2ad33b2913" |
| `insert_after` | `string` | No | option_guid to put after an option **Example value**: "b13adf3c-cad6-4e02-8929-550c112b5633" |
| `is_hidden` | `boolean` | No | Whether to hide **Example value**: false **Default value**: `false` | ### Request body example

{
    "name": "high",
    "color_index": 10,
    "insert_before": "2bd905f8-ef38-408b-aa1f-2b2ad33b2913",
    "insert_after": "b13adf3c-cad6-4e02-8929-550c112b5633",
    "is_hidden": false
}

## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `option` | `option` | created options |
|     `guid` | `string` | option GUID |
|     `name` | `string` | Option name, cannot be empty, maximum 50 characters |
|     `color_index` | `int` | The color index value of the option, which can be a number from 0 to 54. If not filled in, one will be randomly selected. |
|     `is_hidden` | `boolean` | Whether the option is hidden. The hidden option is not visible in the interface, and the field value can no longer be set to this option through openapi. | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
        "option": {
            "guid": "4216f79b-3fda-4dc6-a0c4-a16022e47152",
            "name": "high",
            "color_index": 1,
            "is_hidden": false
        }
    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 1470400 | Parameter error, such as'name 'is too long or the field is not a single/multi select type. | The reason for the error is shown in the returned msg prompt. |
| 403 | 1470403 | Missing permission to edit custom fields | Confirm that the calling identity has permission to edit the custom field. |
| 404 | 1470404 | The custom field to add the option does not exist. | Make sure the custom field to add the option does not exist. |
| 500 | 1470500 | Server error. | Try retrying the call. If it continues to fail, please contact support for feedback. |
