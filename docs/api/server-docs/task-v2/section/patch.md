---
title: "Patch Section"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/task-v2/section/patch"
method: "PATCH"
api_path: "https://open.larksuite.com/open-apis/task/v2/sections/:section_guid"
service: "task-v2"
resource: "section"
section: "Tasks v2"
rate_limit: "1000 per minute、50 per second"
scopes:
  - "task:section:write"
updated: "1699521729000"
---

# Patch Section

Update a section. you can update section's name and postion.

When updating, fill in all the field names to be modified in the 'update_fields 'field, and fill in the new value of the field to be modified in the'section' field. For details of the calling convention, see the "Updates to Resources" part in [Function Overview](/ssl: ttdoc/uAjLw4CM/ukTMukTMukTM/task-v2/overview).

Section fields that currently support updates include:
* 'name' - section name;
* 'insert_before' - section guid before which the current section will be put.
* 'insert_after' -section guid after which the current section will be put.

`insert_before` and `insert_after`, if filled in, must be a valid section_guid of the same resource. Note that `insert_before` and `insert_after` cannot be set at the same time.

> Require the edit permission of the resource the specified section belongs to.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/task/v2/sections/:section_guid |
| HTTP Method | PATCH |
| Rate Limit | 1000 per minute、50 per second |
| Supported app types | custom,isv |
| Required scopes | `task:section:write` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `section_guid` | `string` | section GUID to update **Example value**: "9842501a-9f47-4ff5-a622-d319eeecb97f" **Data validation rules**: - Maximum length: `100` characters | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `user_id_type` | `string` | No | User ID categories **Example value**: open_id **Default value**: `open_id` | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `section` | `input_section` | Yes | The section data to update. |
|   `name` | `string` | No | Section name. When updating, empty is not allowed. Supports up to 100 utf8 characters. **Example value**: "reviewd section." **Data validation rules**: - Maximum length: `300` characters |
|   `insert_before` | `string` | No | The section_guid before which current section will be put. **Example value**: "e6e37dcc-f75a-5936-f589-12fb4b5c80c2" **Data validation rules**: - Maximum length: `100` characters |
|   `insert_after` | `string` | No | The section_guid after which current section will be put. **Example value**: "e6e37dcc-f75a-5936-f589-12fb4b5c80c2" **Data validation rules**: - Maximum length: `100` characters |
| `update_fields` | `string[]` | Yes | The field name to be updated. Currently supports: * `name`: section name; * `insert_before` - The section_guid before which current section will be put. * 'insert_after' - The section_guid after which current section will be put. **Example value**: ["name"] **Data validation rules**: - Length range: `1` ～ `10` | ### Request body example

{
    "section": {
        "name": "reviewd section.",
        "insert_before": "e6e37dcc-f75a-5936-f589-12fb4b5c80c2",
        "insert_after": "e6e37dcc-f75a-5936-f589-12fb4b5c80c2"
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
|   `section` | `section` | Updated section |
|     `guid` | `string` | section GUID |
|     `name` | `string` | section name |
|     `resource_type` | `string` | resource type |
|     `is_default` | `boolean` | whether is the default section |
|     `creator` | `member` | seciton creator |
|       `id` | `string` | Indicates the id of member |
|       `type` | `string` | Type of member |
|       `role` | `string` | Member role |
|     `tasklist` | `tasklist_summary` | If the group belongs to the list, show brief information about the list |
|       `guid` | `string` | tasklist GUID |
|       `name` | `string` | tasklist name |
|     `created_at` | `string` | Section creation timestamp (ms) |
|     `updated_at` | `string` | Section last update timestamp (ms) | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
        "section": {
            "guid": "e6e37dcc-f75a-5936-f589-12fb4b5c80c2",
            "name": "reviewed tasks",
            "resource_type": "tasklist",
            "is_default": true,
            "creator": {
                "id": "ou_2cefb2f014f8d0c6c2d2eb7bafb0e54f",
                "type": "user",
                "role": "editor"
            },
            "tasklist": {
                "guid": "cc371766-6584-cf50-a222-c22cd9055004",
                "name": "活动分工任务列表"
            },
            "created_at": "1675742789470",
            "updated_at": "1675742789470"
        }
    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 1470400 | Incorrect request parameters, such as passing in unsupported field name in `update_fields`。 | The reason for the error is shown in the returned msg prompt. |
| 404 | 1470404 | The section to update does not exist or has been deleted. | Verify that the section  to update does not exist or has been deleted. |
| 500 | 1470500 | Server error. | Retry to invoke tha api with same parameters. If it continues to fail, please contact support for feedback. |
| 403 | 1470403 | Missing permission to update section. | Confirm that the calling identity has permission to edit sections. |
