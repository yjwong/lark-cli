---
title: "Create Folder"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/file/create_folder"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/drive/v1/files/create_folder"
service: "drive-v1"
resource: "file"
section: "Docs"
rate_limit: "Special Rate Limit"
scopes:
  - "drive:drive"
updated: "1691028139000"
---

# Create Folder

Creates a new empty folder within the specified parent folder in the user's Space.

> This API doesn't allow you to concurrently create multiple folders, and supports up to 5 queries per second (QPS) and 10,000 queries per day.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/drive/v1/files/create_folder |
| HTTP Method | POST |
| Rate Limit | Special Rate Limit |
| Supported app types | custom,isv |
| Required scopes | `drive:drive` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `name` | `string` | Yes | The name of the folder you wish to create. **Example value**: "New Folder" |
| `folder_token` | `string` | Yes | Parent folder token. If you need to create folder at My Space as top level folder, pass the My Space token.How to get My Space token **Example value**: "fldbcO1UuPz8VwnpPx5a92abcef" | ### Request body example

{
    "name": "New Folder",
    "folder_token": "fldbcO1UuPz8VwnpPx5a92abcef"
}

## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `token` | `string` | The token of the created folder. |
|   `url` | `string` | The url of the created folder. | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
        "token": "fldbcddUuPz8VwnpPx5oc2abcef",
        "url": "https://larksuite.com/drive/folder/fldbcddUuPz8VwnpPx5oc2abcef"
    }
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
| 400 | 1062507 | parent node out of sibling num. | The number of nodes mounted to the directory in My Space has reached the limit of **1,500** nodes per level. |
