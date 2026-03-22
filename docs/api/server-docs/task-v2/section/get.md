---
title: "Get Section"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/task-v2/section/get"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/task/v2/sections/:section_guid"
service: "task-v2"
resource: "section"
section: "Tasks v2"
rate_limit: "1000 per minute、50 per second"
scopes:
  - "task:section:read"
  - "task:section:write"
updated: "1699521725000"
---

# Get Section

Gets the details of a section, including name, creator, etc. If the section belongs to a tasklist, the summary of the tasklist is also returned.

> Requires read permission of the tasklist if the section belongs to a tasklist

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/task/v2/sections/:section_guid |
| HTTP Method | GET |
| Rate Limit | 1000 per minute、50 per second |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `task:section:read` `task:section:write` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `section_guid` | `string` | Section GUID to get **Example value**: "9842501a-9f47-4ff5-a622-d319eeecb97f" **Data validation rules**: - Maximum length: `100` characters | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `user_id_type` | `string` | No | User ID categories **Example value**: open_id **Default value**: `open_id` | ## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `section` | `section` | Section data |
|     `guid` | `string` | Section GUID |
|     `name` | `string` | Section name |
|     `resource_type` | `string` | resource type |
|     `is_default` | `boolean` | Whether the group is the default section |
|     `creator` | `member` | Creator of section |
|       `id` | `string` | member id |
|       `type` | `string` | Type of member |
|       `role` | `string` | Member role |
|     `tasklist` | `tasklist_summary` | If the tasklist belongs to the tasklist, providing brief information about the tasklist. |
|       `guid` | `string` | Tasklist GUID |
|       `name` | `string` | Tasklist name. |
|     `created_at` | `string` | Section creation timestamp (ms) |
|     `updated_at` | `string` | Section last update timestamp (ms) | ### Response body example

{
    "code": 0,
    "msg": "Success",
    "data": {
        "section": {
            "guid": "E6e37dcc-f75a-5936-f589-12fb4b5c80c2",
            "name": "Tasks that have been reviewed",
            "resource_type": "Tasklist",
            "is_default": true,
            "creator": {
                "id": "ou_2cefb2f014f8d0c6c2d2eb7bafb0e54f",
                "type": "user",
                "role": "Creator"
            },
            "tasklist": {
                "guid": "Cc371766-6584-cf50-a222-c22cd9055004",
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
| 400 | 1470400 | The request parameter is incorrect, such as the section_guid passed in is not valid. | For the specific error reason, please refer to the error message of the returned msg. |
| 404 | 1470404 | The section  to get does not exist or has been deleted. | Confirm that the section does not exist or has been deleted. |
| 500 | 1470500 | Server error. | Retry to invoke the api. If it continues to fail, please contact support for feedback. |
| 403 | 1470403 | Missing permission to access section | Confirm that the calling identity has access to the section. |
