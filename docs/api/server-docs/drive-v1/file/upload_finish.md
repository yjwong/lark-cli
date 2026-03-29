---
title: "Multipart Upload Files (Finish)"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/file/upload_finish"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/drive/v1/files/upload_finish"
service: "drive-v1"
resource: "file"
section: "Docs"
scopes:
  - "drive:drive"
  - "drive:file"
updated: "1665221894000"
---

# Complete uploading a file in blocks

Completes an upload task.

> This API doesn't support high concurrency, and supports up to 5 queries per second (QPS).

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/drive/v1/files/upload_finish |
| HTTP Method | POST |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `drive:drive` `drive:file` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" Learn more about obtaining and using access_token. |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `upload_id` | `string` | Yes | ID of the block upload transaction **Example value**: "7111211691345512356" |
| `block_num` | `int` | Yes | Number of blocks **Example value**: 1 | ### Request body example

{
    "upload_id": "7111211691345512356",
    "block_num": 1
}

## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `file_token` | `string` | Token of the new file | ### Response body example

{
    "code": 0,
    "msg": "Success",
    "data": {
        "file_token": "boxcnrHpsg1QDqXAAAyachabcef"
    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 200 | 1061001 | internal error. | Internal service error, such as timeout or failure in processing error codes. |
| 400 | 1061002 | params error. | Check whether the request parameters are correct. |
| 404 | 1061003 | not found. | Check whether the resource exists. |
| 403 | 1061004 | forbidden. | Check whether the current user or app has the permissions for the destination upload node. For example, check whether the user has the edit permission for the specified Docs for document upload. |
| 400 | 1061101 | file quota exceeded. | The tenant has reached the maximum capacity. Make sure that the tenant has sufficient capacity and then upload again. |
| 400 | 1061021 | upload id expire. | The upload transaction has expired. Please upload again. |
| 500 | 1061022 | file version conflict. | File version number conflicts. |
| 400 | 1061041 | parent node has been deleted. | Check whether the upload point has been deleted. |
| 400 | 1061042 | parent node out of limit. | The number of materials to upload to the current upload node has reached the limit. Please change the upload point. |
| 400 | 1061043 | file size beyond limit. | Check whether the length of the file is within the limit. For more information, see [File size limits in Drive](https://www.feishu.cn/hc/zh-CN/articles/360049067549). |
| 400 | 1061044 | parent node not exist. | Check whether the upload point exists. |
| 200 | 1061045 | can retry. | Internal error. Please try again later. |
| 400 | 1061061 | user quota exceeded. | You have reached your maximum personal capacity. Make sure that you have sufficient capacity and then upload again. |
| 403 | 1061073 | no scope auth. | You have no access to the API. |
| 403 | 1061500 | mount node point kill | The mount point does not exist. |
| 400 | 1062007 | upload user not match. | Make sure that the current request is sent by the same user or app as the upload task. |
| 400 | 1062010 | block missing, please upload all blocks. | Some file blocks are missing. Make sure that all file blocks are uploaded. |
| 400 | 1062505 | parent node out of size. | The single tree in My Space has reached the maximum size of 400,000. |
| 400 | 1062506 | parent node out of depth. | My Space supports up to 15 levels of directories. |
| 400 | 1062507 | parent node out of sibling num. | The number of nodes mounted to the directory in My Space has reached the limit of **1,500** nodes per level. |
