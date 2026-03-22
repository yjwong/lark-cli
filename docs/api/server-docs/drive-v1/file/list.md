---
title: "List items in folder"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/file/list"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/drive/v1/files"
service: "drive-v1"
resource: "file"
section: "Docs"
rate_limit: "Special Rate Limit"
scopes:
  - "drive:drive"
  - "drive:drive:readonly"
updated: "1691028139000"
---

# List items in folder

Get the list of files under the specified folder in the user's cloud space. List item types include files, various online documents (doc, sheet, bitable, mindnote), folders, and shortcuts. This interface does support paging, but does not recursively get the list of subfolders.

> To allow an app (tenant_access_token) to access a folder in a personal cloud space, please visit Space FAQs 3rd FAQ.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/drive/v1/files |
| HTTP Method | GET |
| Rate Limit | Special Rate Limit |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `drive:drive` `drive:drive:readonly` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `page_size` | `int` | No | **Example value**: 10 **Data validation rules**: - Maximum value: `200` |
| `page_token` | `string` | No | Page identifier. It is not filled in the first request, indicating traversal from the beginning; when there will be more groups, the new page_token will be returned at the same time, and the next traversal can use the page_token to get more groups **Example value**: MTY1NTA3MTA1OXw3MTA4NDc2MDc1NzkyOTI0Nabcef |
| `folder_token` | `string` | No | Folder's token. (If this parameter is not provided or an empty string is provided, then return the list of the user's cloud space and pagination is not support) **Example value**: fldbcRho46N6MQ3mJkOAuPabcef | ## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `files` | `file[]` | List of folders. |
|     `token` | `string` | File token. |
|     `name` | `string` | File name. |
|     `type` | `string` | File type. **Optional values are**: - `doc`：Docs - `sheet`：Sheets - `mindnote`：Mindnotes - `bitable`：Bitable - `file`：File - `docx`：Upgraded Docs - `folder`：Folder |
|     `parent_token` | `string` | Parent folder token. |
|     `url` | `string` | Link viewed in browser. |
|     `shortcut_info` | `shortcut_info` | Shortcut file information. |
|       `target_type` | `string` | The original file type that the shortcut points to. **Optional values are**: - `doc`：Docs - `sheet`：Sheets - `mindnote`：Mindnotes - `bitable`：Bitable - `file`：File - `docx`：Upgraded Docs |
|       `target_token` | `string` | The original file token that the shortcut points to. |
|   `next_page_token` | `string` | Paging tag, when has_more is true, the next_page_token of the next traversal will be returned at the same time, otherwise it will not be returned. |
|   `has_more` | `boolean` | Whether the response body has more parameters | ### Response body example

{
    "code": 0,
    "data": {
        "files": [
            {
                "name": "test pdf.pdf",
                "parent_token": "fldbcO1UuPz8VwnpPx5a9abcef",
                "token": "boxbc0dGSMu23m7QkC1bvabcef",
                "type": "file",
                "url": "https://larksuite.com/file/boxbc0dGSMu23m7QkC1bvabcef"
            },
            {
                "name": "test file.docx",
                "parent_token": "fldcnCEG903UUB4fUqfysdabcef",
                "shortcut_info": {
                    "target_token": "boxcnRPaXpD4I6Je9t8k8babcef",
                    "target_type": "file"
                },
                "token": "nodcnVkiLQ6LD4CsUEaANrabcef",
                "type": "shortcut",
                "url": "https://larksuite.com/file/boxcnRPaXpD4I6Je9t8k8babcef"
            },
            {
                "name": "test docx",
                "parent_token": "fldcnCEG903UUB4fUqfysdabcef",
                "token": "doxcntan34DX4QoKJu7jJyabcef",
                "type": "docx",
                "url": "https://larksuite.com/docx/doxcntan34DX4QoKJu7jJyabcef"
            },
            {
                "name": "test sheet",
                "parent_token": "fldcnCEG903UUB4fUqfysdabcef",
                "token": "shtcnOko1Ad0HU48HH8KHuabcef",
                "type": "sheet",
                "url": "https://larksuite.com/sheets/shtcnOko1Ad0HU48HH8KHuabcef"
            }
        ],
        "has_more": false
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
| 400 | 1064001 | page size out of limit. | Check whether the page size exceeds the limit. |
