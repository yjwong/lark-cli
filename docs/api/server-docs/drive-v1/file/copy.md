---
title: "Copy File"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/file/copy"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/drive/v1/files/:file_token/copy"
service: "drive-v1"
resource: "file"
section: "Docs"
rate_limit: "Special Rate Limit"
scopes:
  - "drive:drive"
updated: "1691028143000"
---

# Copy File

Copy a file to a different folder in the user's Space. Folders cannot be copied.

If the destination folder is My Space, the copied files will be listed in the "**Owned by Me**" list in "**My Space**".

> This API doesn't allow you to concurrently copy multiple files, and supports up to 5 queries per second (QPS) and 10,000 queries per day.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/drive/v1/files/:file_token/copy |
| HTTP Method | POST |
| Rate Limit | Special Rate Limit |
| Supported app types | custom,isv |
| Required scopes | `drive:drive` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `file_token` | `string` | File token that needs to be copied. **Example value**: "doccngpahSdXrFPIBD4XdIabcef" | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `name` | `string` | Yes | The new name of the copied file. **Example value**: "test.txt" |
| `type` | `string` | No | File type that need to be copied. If the value is empty or does not match the actual file type, the interface will return failure. **Example value**: "doc" **Optional values are**:  - `file`: Common file type - `doc`: Cloud document type of docs - `sheet`: Cloud document type of sheet - `bitable`: Cloud document type of bitable - `docx`: Cloud document type of upgraded docs - `mindnote`: Cloud document type of mindnote  |
| `folder_token` | `string` | Yes | The target folder token to which the file is copied. **Example value**: "fldbcO1UuPz8VwnpPx5a92abcef" |
| `extra` | `property[]` | No | Additional parameters for special replication semantics |
|   `key` | `string` | Yes | Custom property key **Example value**: "target_type" |
|   `value` | `string` | Yes | Custom property value **Example value**: "docx" | ### Request body example

{
    "name": "name",
    "type": "file",
    "folder_token": "fldbcO1UuPz8VwnpPx5a92abcef"
}

## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `file` | `file` | The copied file resource. |
|     `token` | `string` | File token. |
|     `name` | `string` | File name. |
|     `type` | `string` | File type. |
|     `parent_token` | `string` | Parent folder token. |
|     `url` | `string` | Link viewed in browser. |
|     `shortcut_info` | `shortcut_info` | Shortcut file information. |
|       `target_type` | `string` | The original file type that the shortcut points to. |
|       `target_token` | `string` | The original file token that the shortcut points to. | ### Response body example

{
    "code": 0,
    "data": {
        "file": {
            "name": "name",
            "parent_token": "fldcnBh8LrnX42dr1pBYclabcef",
            "token": "doxcnUkUOWtOelpFcha2Z9abcef",
            "type": "docx",
            "url": "https://larksuite.com/docx/doxcnUkUOWtOelpFcha2Zabcef"
        }
    },
    "msg": "success"
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 500 | 1061001 | internal error. | Internal service error, such as timeout or failure in processing error codes. |
| 400 | 1061002 | params error. | Check whether the request parameters are correct. |
| 404 | 1061003 | not found. | Check whether the resource exists. |
| 403 | 1061004 | forbidden. | Check whether the current user or app has the permissions for the destination upload node. For example, check whether the user has the edit permission for the specified Docs for document upload. |
| 401 | 1061005 | auth failed. | Assume the correct user or app mode to access the API. |
| 404 | 1061007 | file has been delete. | Check whether the node has been deleted. |
| 400 | 1061045 | resource contention occurred, please retry. | Please retry. |
| 400 | 1062507 | parent node out of sibling num. | The number of nodes mounted to the directory in My Space has reached the limit of **1,500** nodes per level. |
| 403 | 1064510 | cross tenant and unit not support. | Cross-tenant and cross-region requests are not supported. |
| 403 | 1064511 | cross brand not support. | Cross-brand requests are not supported. |
