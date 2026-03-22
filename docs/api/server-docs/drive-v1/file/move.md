---
title: "Move File/Folder"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/file/move"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/drive/v1/files/:file_token/move"
service: "drive-v1"
resource: "file"
section: "Docs"
rate_limit: "Special Rate Limit"
scopes:
  - "drive:drive"
updated: "1691028143000"
---

# Move File

Move a file or folder to a different location in the user's Space.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/drive/v1/files/:file_token/move |
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
| `file_token` | `string` | File token that needs to be moved. **Example value**: "boxcnrHpsg1QDqXAAAyachabcef" | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `type` | `string` | No | File type that need to be moved. If the value is empty or does not match the actual file type, the interface will return failure. **Example value**: "file" **Optional values are**:  - `file`: Common file type - `docx`: Cloud document type of upgraded docs - `bitable`: Cloud document type of bitable - `doc`: Cloud document type of docs - `sheet`: Cloud document type of sheet - `mindnote`: Cloud document type of mindnote - `folder`: Folder type  |
| `folder_token` | `string` | No | Target folder token. **Example value**: "fldbcO1UuPz8VwnpPx5a92abcef" | ### Request body example

{
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
|   `task_id` | `string` | Async task id, returned when moving a folder. | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
        "task_id": "12345"
    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 500 | 1061001 | internal error. | Internal service error, such as timeout or failure in processing error codes. |
| 400 | 1061002 | params error. | Check whether the request parameters are correct. |
| 404 | 1061003 | not found. | Check whether the resource exists. |
| 403 | 1061004 | forbidden. | Check whether the current user or app has the permissions for the destination upload node. For example, check whether the user has the edit permission for the specified Docs for document upload. |
| 404 | 1061007 | file has been delete. | Check whether the node has been deleted. |
| 403 | 1062524 | source parent no permission. | Please check that you have edit permissions to the folder where the moved file is located. |
| 403 | 1062535 | destination parent no permission. | Please check if the target folder has edit permissions. |
| 400 | 1062507 | parent node out of sibling num. | The number of nodes mounted to the directory in My Space has reached the limit of **1,500** nodes per level. |
