---
title: "List Section"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/task-v2/section/list"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/task/v2/sections"
service: "task-v2"
resource: "section"
section: "Tasks v2"
rate_limit: "1000 per minute、50 per second"
scopes:
  - "task:section:read"
  - "task:section:write"
updated: "1699521737000"
---

# List Sections

Get a list of section. Paging is supported. The returned results are sorted in the order in which the sections are placed on the UI.

> Listing sections requires read permission on the resource.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/task/v2/sections |
| HTTP Method | GET |
| Rate Limit | 1000 per minute、50 per second |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `task:section:read` `task:section:write` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `page_size` | `int` | No | paging size **Example value**: 50 **Default value**: `50` **Data validation rules**: - Value range: `1` ～ `100` |
| `page_token` | `string` | No | Page identifier. It is not filled in the first request, indicating traversal from the beginning; when there will be more groups, the new page_token will be returned at the same time, and the next traversal can use the page_token to get more groups **Example value**: aWQ9NzEwMjMzMjMxMDE= **Data validation rules**: - Maximum length: `100` characters |
| `resource_type` | `string` | Yes | The resource type to which the section belongs. Supports "my_tasks" (Owned Tasks) and "tasklist". When using tasklist, you need to provide the tasklist GUID by setting `resource_id` **Example value**: tasklist **Data validation rules**: - Maximum length: `100` characters |
| `resource_id` | `string` | No | If "resource_type" is "tasklist", you need to fill in the taslist GUID of which sections is going to list. **Example value**: caef228f-2342-23c1-c36d-91186414dc64 **Data validation rules**: - Maximum length: `100` characters |
| `user_id_type` | `string` | No | User ID categories **Example value**: open_id **Default value**: `open_id` | ## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `items` | `section_summary[]` | List of section summaries. |
|     `guid` | `string` | section GUID. |
|     `name` | `string` | section name |
|     `is_default` | `boolean` | whether is the default section. |
|   `page_token` | `string` | Page identifier, when has_more is true, a new page_token will also be returned. Otherwise, page_token will not be returned |
|   `has_more` | `boolean` | Whether the response body has more parameters | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
        "items": [
            {
                "guid": "e6e37dcc-f75a-5936-f589-12fb4b5c80c2",
                "name": "审核过的任务",
                "is_default": true
            }
        ],
        "page_token": "aWQ9NzEwMjMzMjMxMDE=",
        "has_more": true
    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 1470400 | Incorrect request parameters, such as passing a negative number to `page_size`. | The reason for the error is shown in the returned msg prompt. |
| 404 | 1470404 | The resource to list sections does not exist or has been deleted. | Confirm that the resource exists or has been deleted. |
| 500 | 1470500 | Server error. | Try retrying the call. If it continues to fail, please contact support for feedback. |
| 403 | 1470403 | Missing permission for listing sections. | Confirm that the calling identity has permission to list sections. |
