---
title: "Create Section"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/task-v2/section/create"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/task/v2/sections"
service: "task-v2"
resource: "section"
section: "Tasks v2"
rate_limit: "1000 per minute、50 per second"
scopes:
  - "task:section:write"
updated: "1699521721000"
---

# Create Section

Create a section for a tasklist or "Owned Tasks". You may provide a name and optional locating parameters. If no locating parameters are specified, the new section will be placed at the end of the sections of the specified resource.

When creating a section in the tasklist, you need to set `resourse_type` to "tasklist" and `resource_id` to the tasklist guid.

When creating section for "Owned Tasks", you need to set `resource_type` to "my_tasks",  and leave `resource_id` unset. The calling identity can only create sections in "Owned Tasks" of itself.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/task/v2/sections |
| HTTP Method | POST |
| Rate Limit | 1000 per minute、50 per second |
| Supported app types | custom,isv |
| Required scopes | `task:section:write` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `user_id_type` | `string` | No | User ID categories **Example value**: open_id **Default value**: `open_id` | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `name` | `string` | Yes | Section name. Empty is not allowed. Supports up to 100 utf8 characters. **Example value**: "reviewed tasks" |
| `resource_type` | `string` | Yes | Resource type of new creation, supporting "tasklist" or "my_tasks" (Owned Tasks). **Example value**: "tasklist" **Default value**: `tasklist` |
| `resource_id` | `string` | No | Specify the resource id to which the section belongs. When `resource_type` is "tasklist", the tasklist GUID needs to be spedified here; when `resource_type` is "my_tasks", no need to fill in. **Example value**: "cc371766-6584-cf50-a222-c22cd9055004" |
| `insert_before` | `string` | No | The section_guid before which the new created section will be put. When neither `insert_before` nor `insert_after` is set, the new section is placed after all existing sections. If both `insert_before` and `insert_after` are set, the api will report an error. **Example value**: "e6e37dcc-f75a-5936-f589-12fb4b5c80c2" **Data validation rules**: - Maximum length: `100` characters |
| `insert_after` | `string` | No | The section_guid after which the new created section will be put. When neither `insert_before` nor `insert_after` is set, the new section is placed after all existing sections. If both `insert_before` and `insert_after` are set, the api will report an error. **Example value**: "e6e37dcc-f75a-5936-f589-12fb4b5c80c2" **Data validation rules**: - Maximum length: `100` characters | ### Request body example

{
    "name": "reviewed tasks",
    "resource_type": "tasklist",
    "resource_id": "cc371766-6584-cf50-a222-c22cd9055004",
    "insert_before": "e6e37dcc-f75a-5936-f589-12fb4b5c80c2",
    "insert_after": "e6e37dcc-f75a-5936-f589-12fb4b5c80c2"
}

## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `section` | `section` | section data |
|     `guid` | `string` | section GUID |
|     `name` | `string` | section name |
|     `resource_type` | `string` | resource type |
|     `is_default` | `boolean` | whether is the default section. |
|     `creator` | `member` | section creator |
|       `id` | `string` | member id |
|       `type` | `string` | member type |
|       `role` | `string` | member role |
|     `tasklist` | `tasklist_summary` | tasklist summary of the section, available only when the section belongs to a tasklist. |
|       `guid` | `string` | tasklist GUID |
|       `name` | `string` | tasklist name |
|     `created_at` | `string` | section creation timestamp (ms) |
|     `updated_at` | `string` | section last update timestamp (ms) | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
        "section": {
            "guid": "e6e37dcc-f75a-5936-f589-12fb4b5c80c2",
            "name": "reviewd tasks",
            "resource_type": "tasklist",
            "is_default": true,
            "creator": {
                "id": "ou_2cefb2f014f8d0c6c2d2eb7bafb0e54f",
                "type": "user",
                "role": "editor"
            },
            "tasklist": {
                "guid": "cc371766-6584-cf50-a222-c22cd9055004",
                "name": "worklist"
            },
            "created_at": "1675742789470",
            "updated_at": "1675742789470"
        }
    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 1470400 | Incorrect request parameters, such as setting both `insert_before` and `insert_after` simutaneously. | Refer to the returned msg for the specific error reason. |
| 404 | 1470404 | Th resource to create as section does not exist or has been deleted. | Confirm that resource exists or has been deleted. |
| 500 | 1470500 | Server error. | Retry invoke the api with same parameters. If it continues to fail, please contact support for feedback. |
| 403 | 1470403 | Missing permission to create sections | Verify the calling identity has permission to create sections. |
