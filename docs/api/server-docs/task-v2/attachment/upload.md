---
title: "Upload Attachment"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/task-v2/attachment/upload"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/task/v2/attachments/upload"
service: "task-v2"
resource: "attachment"
section: "Tasks v2"
rate_limit: "10 per second"
scopes:
  - "task:attachment:write"
field_scopes:
  - "contact:user.employee_id:readonly"
updated: "1699521701000"
---

# Upload Attachment

Upload attachments for a resource. This API supports uploading up to 5 attachments at a time. The size of each attachment cannot exceed 50MB. Attachment file's format is not limited.

The request body format is "form-data". To upload multiple files at once, just fill in multiple "file" parameter. The result attachments order will follow the order the files are passed in.

Currently the resource type only supports "task", with which `resource_id` need to be filled in a task GUID.

> Uploading attachments for a task requires edit permissions for the task.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/task/v2/attachments/upload |
| HTTP Method | POST |
| Rate Limit | 10 per second |
| Supported app types | custom,isv |
| Required scopes | `task:attachment:write` |
| Required field scopes | > The response body of the API contains the following sensitive fields, and they will be returned only after corresponding scopes are added. If you do not need the fields, it is not recommended that you request the scopes. `contact:user.employee_id:readonly` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token |
| Content-Type | string | Yes | **Example value**: "multipart/form-data; boundary=---7MA4YWxkTrZu0gW" | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `user_id_type` | `string` | No | User ID categories **Example value**: open_id **Optional values are**:  - `open_id`: Identifies a user to an app. The same user has different Open IDs in different apps. How to get Open ID - `union_id`: Identifies a user to a tenant that acts as a developer. A user has the same Union ID in apps developed by the same developer, and has different Union IDs in apps developed by different developers. A developer can use Union ID to link the same user's identities in multiple apps.How to get Union ID - `user_id`: Identifies a user to a tenant. The same user has different User IDs in different tenants. In one single tenant, a user has the same User ID in all apps （including store apps）. User ID is usually used to communicate user data between different apps. How to get User ID  **Default value**: `open_id` **When the value is `user_id`, the following field scopes are required**: `contact:user.employee_id:readonly` | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `resource_type` | `string` | No | Resource Type the attachment belongs to. **Example value**: "task" **Default value**: `task` |
| `resource_id` | `string` | Yes | The resource ID to which the attachment belongs. When `resource_type` is "task", `resource_id` should be filled in a task GUID. Task GUID can be feched by Task APIs。 **Example value**: "fe96108d-b004-4a47-b2f8-6886e758b3a5" **Data validation rules**: - Maximum length: `100` characters |
| `file` | `file` | Yes | files to be uploaded, up to 5 files per request. The result attachments will keep the input order. **Example value**: file binary | ### Request body example

```HTTP
---7MA4YWxkTrZu0gW
Content-Disposition: form-data; name="resource_type";

task
---7MA4YWxkTrZu0gW
Content-Disposition: form-data; name="resource_id";

fe96108d-b004-4a47-b2f8-6886e758b3a5
---7MA4YWxkTrZu0gW
Content-Disposition: form-data; name="file";
Content-Type: application/octet-stream

---7MA4YWxkTrZu0gW
```

## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `items` | `attachment[]` | List of uploaded attachments |
|     `guid` | `string` | Attachment GUID |
|     `file_token` | `string` | Attached tokens in the space system |
|     `name` | `string` | Attachment name |
|     `size` | `int` | Byte size of attachment |
|     `resource` | `resource` | Resources attributed to attachments |
|       `type` | `string` | Resource Type |
|       `id` | `string` | Resource ID |
|     `uploader` | `member` | attachment uploader |
|       `id` | `string` | The id of the member |
|       `type` | `string` | Type of member |
|       `role` | `string` | member role |
|     `is_cover` | `boolean` | Whether the attachment is the cover image. |
|     `uploaded_at` | `string` | Upload timestamp (ms) | ### Response body example

{
    "code": 0,
    "msg": "Success",
    "data": {
        "items": [
            {
                "guid": "f860de3e-6881-4ddd-9321-070f36d1af0b",
                "file_token": "boxcnTDqPaRA6JbYnzQsZ2doB2b",
                "name": "Fo.jpg",
                "size": 62232,
                "resource": {
                    "type": "task",
                    "id": "e6e37dcc-f75a-5936-f589-12fb4b5c80c2"
                },
                "uploader": {
                    "id": "ou_ 2cefb2f014f8d0c6c2d2eb7bafb0e54f",
                    "type": "user",
                    "role": "creator"
                },
                "is_cover": false,
                "uploaded_at": "1675742789470"
            }
        ]
    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 1470400 | The request parameter is incorrect. | Check the `msg` in the returned response body to determine the specific reason. |
| 403 | 1470403 | The resource does not exist or has been deleted. | Verify whether the resource  still exists. |
| 404 | 1470404 | Missing permission to upload attachments. | Verify that the calling identity has permission to upload attachments. |
| 500 | 1470500 | Server error. | Re-invoke the API with the same parameters. If it continues to fail, please contact support for feedback. |
