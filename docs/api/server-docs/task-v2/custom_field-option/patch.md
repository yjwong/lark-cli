---
title: "Patch Custom Field Option"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/task-v2/custom_field-option/patch"
method: "PATCH"
api_path: "https://open.larksuite.com/open-apis/task/v2/custom_fields/:custom_field_guid/options/:option_guid"
service: "task-v2"
resource: "custom_field-option"
section: "Tasks v2"
rate_limit: "100 per minute"
scopes:
  - "task:tasklist:write"
updated: "1699521776000"
---

# Update Custom Field Option

Update the a custom field option by specifying the field GUID and option GUID. The field to be updated must be of a single or multi-select type, and the option to be updated must belong to that custom field.

When updating, fill in the `update_fields` field, and fill in the new values of the field to be updated in the `option` field. Fields supported by update_fields include:

* 'name': option name
* 'color_index': Color index value of option
* 'is_hidden': whether to hide from the interface
* 'insert_before': Put the current option in the option_guid before an option in the same field.
* 'insert_after': The option_guid that puts the current option after an option in the same field.

> Update options require edit permissions for custom fields

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/task/v2/custom_fields/:custom_field_guid/options/:option_guid |
| HTTP Method | PATCH |
| Rate Limit | 100 per minute |
| Supported app types | custom,isv |
| Required scopes | `task:tasklist:write` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `custom_field_guid` | `string` | custom field GUID whose option is to update **Example value**: "b13adf3c-cad6-4e02-8929-550c112b5633" |
| `option_guid` | `string` | option GUID to be udpate **Example value**: "b13adf3c-cad6-4e02-8929-550c112b5633" | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `option` | `input_option` | No | The option data to update |
|   `name` | `string` | No | Option name, up to 50 characters **Example value**: "high" |
|   `color_index` | `int` | No | Color index value, supports a number from 0 to 54. **Example value**: 10 **Data validation rules**: - Value range: `0` ～ `54` |
|   `insert_before` | `string` | No | option_guid to put before an option **Example value**: "2bd905f8-ef38-408b-aa1f-2b2ad33b2913" |
|   `insert_after` | `string` | No | option_guid to put after an option **Example value**: "b13adf3c-cad6-4e02-8929-550c112b5633" |
|   `is_hidden` | `boolean` | No | Whether to hide **Example value**: false **Default value**: `false` |
| `update_fields` | `string[]` | No | Field name to update, support * 'name': option name * 'color_index': color index value of option * 'is_hidden': whether to hide from the interface * 'insert_before': put the current option before an option in the same field. * 'insert_after': put the current option after an option in the same field. **Example value**: ["name"] **Data validation rules**: - Length range: `1` ～ `20` | ### Request body example

{
    "option": {
        "name": "high",
        "color_index": 10,
        "insert_before": "2bd905f8-ef38-408b-aa1f-2b2ad33b2913",
        "insert_after": "b13adf3c-cad6-4e02-8929-550c112b5633",
        "is_hidden": false
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
|   `option` | `option` | Updated option data |
|     `guid` | `string` | GUID option |
|     `name` | `string` | Option name, cannot be empty, maximum 50 characters |
|     `color_index` | `int` | The color index value of the option, which can be a number from 0 to 54. |
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
| 400 | 1470400 | Incorrect request parameters, such as a name that does not support updating in the provided update_fields. | Determine the actual cause of the error based on the returned msg. |
| 403 | 1470403 | The calling identity lacks permission to update the custom group. | Confirm that the calling identity has permission to edit custom groups. |
| 404 | 1470404 | The custom field or option to update does not exist. | Verify that the custom field or option to update exists. |
| 500 | 1470500 | Server error. | Try retrying the call. If it continues to fail, please contact support for feedback. |
