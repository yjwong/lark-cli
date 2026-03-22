---
title: "Delete File/Folder"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/file/delete"
method: "DELETE"
api_path: "https://open.larksuite.com/open-apis/drive/v1/files/:file_token"
service: "drive-v1"
resource: "file"
section: "Docs"
rate_limit: "Special Rate Limit"
scopes:
  - "drive:drive"
updated: "1691028143000"
---

# Delete file

Delete a file or folder from user's Space. After the file or folder is deleted, it will go to the user's recycle bin.

> Deleting file requires ensuring that the app has one of the following two permissions:
> 1. The app is the owner of the file and has edit rights to the parent folder where the file is located.
> 2. The app is not the owner of the file, but is the owner of the parent folder where the file is located or has full access to the parent folder.

> This API doesn't support concurrent calls, and supports up to 5 queries per second (QPS). Deleting a folder will be executed asynchronously and a task_id will be returned. You can use the task_check interface to query the task execution status.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/drive/v1/files/:file_token |
| HTTP Method | DELETE |
| Rate Limit | Special Rate Limit |
| Supported app types | custom,isv |
| Required scopes | `drive:drive` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `file_token` | `string` | File token that need to be deleted. **Example value**: "boxcnrHpsg1QDqXAAAyachabcef" | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `type` | `string` | Yes | File type that need to be deleted. **Example value**: file **Optional values are**:  - `file`: Common file type - `docx`: Cloud document type of upgraded docs - `bitable`: Cloud document type of bitable - `folder`: Folder type - `doc`: Cloud document type of docs - `sheet`: Cloud document type of sheet - `mindnote`: Cloud document type of mindnote - `shortcut`: Shortcut type  | ## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `task_id` | `string` | Async task id, returned when the folder is deleted | ### Response body example

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
| 400 | 1061001 | internal error. | Internal service error, such as timeout or failure in processing error codes. |
| 400 | 1061002 | params error. | Check whether the request parameters are correct. |
| 404 | 1061003 | not found. | Check whether the resource exists. |
| 403 | 1061004 | forbidden. | Check whether the current user or app has the permissions for the destination upload node. For example, check whether the user has the edit permission for the specified Docs for document upload. |
| 401 | 1061005 | auth failed. | Assume the correct user or app mode to access the API. |
| 404 | 1061007 | file has been delete. | Check whether the node has been deleted. |
| 403 | 1062501 | operate node no permission. | Please confirm that you have permission to operate the node. |
| 403 | 1062502 | operate sub node no permission. | Please confirm that you have the permission to operate the child node of the node. |
| 404 | 1065200 | source node not found. | The source node of the operation cannot be found, which means that the database record cannot be queried, or the master-slave delay may be caused. | For more information about error codes, see Server-side error codes.
