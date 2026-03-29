---
title: "List Attachment"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/task-v2/attachment/list"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/task/v2/attachments"
service: "task-v2"
resource: "attachment"
section: "Tasks v2"
rate_limit: "1000 per minute、50 per second"
scopes:
  - "task:attachment:read"
  - "task:attachment:write"
updated: "1699521705000"
---

# List Attachment

List all attachments of a resource. The returned attachments supports paging and are sorted by upload time.

Each attachment will return a temporary url available for download.
Url is available for up to 3 minutes and can only be used for up to 3 times. If the limit is exceeded, you need to obtain a new temporary url through this api.

> Get the attachment list of the task requires read permission for the task.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/task/v2/attachments |
| HTTP Method | GET |
| Rate Limit | 1000 per minute、50 per second |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `task:attachment:read` `task:attachment:write` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `page_size` | `int` | No | paging size **Example value**: 50 **Default value**: `50` **Data validation rules**: - Value range: `1` ～ `100` |
| `page_token` | `string` | No | Page identifier. It is not filled in the first request, indicating traversal from the beginning; when there will be more groups, the new page_token will be returned at the same time, and the next traversal can use the page_token to get more groups **Example value**: aWQ9NzEwMjMzMjMxMDE= |
| `resource_type` | `string` | No | The resource type to which the attachment belongs. Currently only "task" is supported. **Example value**: task **Default value**: `task` |
| `resource_id` | `string` | Yes | The id of the attachment home resource, used with resource_type. For example, if you want to get the attachment of the task, you need to set resource_type as task, resource_id as the task GUID. To get task GUID, you can refer to Task Features Overview. **Example value**: 9842501a-9f47-4ff5-a622-d319eeecb97f |
| `user_id_type` | `string` | No | User ID categories **Example value**: open_id **Default value**: `open_id` | ## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `items` | `attachment[]` | Attachment list |
|     `guid` | `string` | Attachment guid |
|     `file_token` | `string` | Attachment token in space system |
|     `name` | `string` | Attachment name |
|     `size` | `int` | Byte size of attachment |
|     `resource` | `resource` | Resources attributed to attachments |
|       `type` | `string` | resource type |
|       `id` | `string` | Resource ID |
|     `uploader` | `member` | Attachment uploader |
|       `id` | `string` | Indicates the id of member |
|       `type` | `string` | Type of member |
|       `role` | `string` | Member role |
|     `is_cover` | `boolean` | Whether is it a cover image. |
|     `uploaded_at` | `string` | Upload timestamp (ms) |
|     `url` | `string` | The temporary download url of the attachment is available for 3 minutes, and max 3 calls are allowed to download the attachment. |
|   `page_token` | `string` | Page identifier, when has_more is true, a new page_token will also be returned. Otherwise, page_token will not be returned |
|   `has_more` | `boolean` | Whether the response body has more parameters | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
        "items": [
            {
                "guid": "f860de3e-6881-4ddd-9321-070f36d1af0b",
                "file_token": "boxcnTDqPaRA6JbYnzQsZ2doB2b",
                "name": "foo.jpg",
                "size": 62232,
                "resource": {
                    "type": "task",
                    "id": "e6e37dcc-f75a-5936-f589-12fb4b5c80c2"
                },
                "uploader": {
                    "id": "ou_2cefb2f014f8d0c6c2d2eb7bafb0e54f",
                    "type": "user",
                    "role": "editor"
                },
                "is_cover": false,
                "uploaded_at": "1675742789470",
                "url": "https://example.com/download/authcode/?code=OWMzNDlmMjJmZThkYzZkZGJlMjYwZTI0OTUxZTE2MDJfMDZmZmMwOWVj"
            }
        ],
        "page_token": "aWQ9NzEwMjMzMjMxMDE=",
        "has_more": true
    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 1470400 | Incorrect request parameters | Check the 'msg' field in the returned request body to determine the specific reason. |
| 404 | 1470404 | The resource to query the attachment does not exist or has been deleted. | Confirm that the resource to be queried still exists. |
| 500 | 1470500 | Server error. | Re-invoke the API with the same parameters. If it continues to fail, please contact support for feedback. |
| 403 | 1470403 | Missing permission to list attachments. | Confirm that the calling identity has permission to list attachments. |
