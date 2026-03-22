---
title: "Get Attachment"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/task-v2/attachment/get"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/task/v2/attachments/:attachment_guid"
service: "task-v2"
resource: "attachment"
section: "Tasks v2"
rate_limit: "1000 per minute、50 per second"
scopes:
  - "task:attachment:read"
  - "task:attachment:write"
updated: "1699521709000"
---

# Get Attachment

Providing an attachment GUID, get the detail of the attachment, including GUID, name, size, uploaded time, temporary downloadable url, etc.

> Getting attachment requires read permission to the resource to which the attachment belongs.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/task/v2/attachments/:attachment_guid |
| HTTP Method | GET |
| Rate Limit | 1000 per minute、50 per second |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `task:attachment:read` `task:attachment:write` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `attachment_guid` | `string` | Attachment GUID, which can be created by Upload AttachmentAPI, or fetched byList AttachmentAPI. **Example value**: "b59aa7a3-e98c-4830-8273-cbb29f89b837" | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `user_id_type` | `string` | No | User ID categories **Example value**: open_id **Default value**: `open_id` | ## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `attachment` | `attachment` | Attachment details |
|     `guid` | `string` | Attachment GUID |
|     `file_token` | `string` | Attached tokens in the space system |
|     `name` | `string` | Attachment name |
|     `size` | `int` | Attachment size in bytes |
|     `resource` | `resource` | Resource to which the attachment belongs to. |
|       `type` | `string` | Resource type |
|       `id` | `string` | Resource ID |
|     `uploader` | `member` | attachment uploader |
|       `id` | `string` | Member ID |
|       `type` | `string` | Member Type |
|       `role` | `string` | member role |
|     `is_cover` | `boolean` | Whether the attachment is the cover image. |
|     `uploaded_at` | `string` | Upload timestamp (ms) |
|     `url` | `string` | The temporary download URL for the attachment. The url is avaible in 3 minutes, and only allows up to 3 downloads. It is only dynamically generated when the attachement is got. | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
        "attachment": {
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
    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 1470400 | The request parameter is incorrect. | Check the `msg` in the return to determine the specific reason. |
| 404 | 1470404 | The attachment does not exist or has been deleted. | Verify whether the attachment still exists. |
| 500 | 1470500 | Server error. | Re-invoke the API with the same parameters. If it continues to fail, please contact support for feedback. |
| 403 | 1470403 | Missing read permission to the attachment. | Verify that the calling identity has read permission to the attachment. |
