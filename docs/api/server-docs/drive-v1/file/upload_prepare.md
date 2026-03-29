---
title: "Multipart Upload Files (Prepare)"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/file/upload_prepare"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/drive/v1/files/upload_prepare"
service: "drive-v1"
resource: "file"
section: "Docs"
scopes:
  - "drive:drive"
  - "drive:file"
updated: "1665221894000"
---

# Upload a file in blocks (Pre­uploading)

Sends an initialization request to obtain an upload transaction ID and a split policy. The current policy is to split a file into 4 MB fixed-length blocks.

> You can save the upload transaction ID and upload progress within 24 hours so that you can resume uploading if needed.

> This API doesn't support high concurrency, and supports up to 5 queries per second (QPS).

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/drive/v1/files/upload_prepare |
| HTTP Method | POST |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `drive:drive` `drive:file` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" Learn more about obtaining and using access_token. |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ::: note
For more information about how to call the AccessToken in the Docs API, see Get started with the Docs API.

### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `file_name` | `string` | Yes | File name **Example value**: "test.txt" **Data validation rules**: - Maximum length: `250` characters |
| `parent_type` | `string` | Yes | Type of the upload point **Example value**: "explorer" **Optional values are**:  - `explorer`: My Space.  |
| `parent_node` | `string` | Yes | Token of the folder **Example value**: "fldbcO1UuPz8VwnpPx5a92abcef" |
| `size` | `int` | Yes | File size **Example value**: 1024 **Data validation rules**: - Minimum value: `0` | ### Request body example

{
    "file_name": "test.txt",
    "parent_type": "explorer",
    "parent_node": "fldbcO1UuPz8VwnpPx5a92abcef",
    "size": 1024
}

## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `upload_id` | `string` | ID of the block upload transaction |
|   `block_size` | `int` | Size of blocks into which you need to split the file |
|   `block_num` | `int` | Number of blocks | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
        "upload_id": "7111211691345512356",
        "block_size": 4194304,
        "block_num": 1
    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 200 | 1061001 | internal error. | Internal service error, such as timeout or failure in processing error codes. |
| 400 | 1061002 | params error. | Check whether the request parameters are correct. |
| 403 | 1061004 | forbidden. | Check whether the current user or app has the permissions for the destination upload node. For example, check whether the user has the edit permission for the specified Docs for document upload. |
| 500 | 1061022 | file version conflict. | File version number conflicts. |
| 400 | 1061043 | file size beyond limit. | Check whether the length of the file is within the limit. For more information, see [File size limits in Drive](https://www.feishu.cn/hc/zh-CN/articles/360049067549). |
| 400 | 1061044 | parent node not exist. | Check whether the upload point exists. |
| 400 | 1061109 | file name cqc not passed. | Make sure that the file to upload and the file name meet compliance. |
| 400 | 1061101 | file quota exceeded. | The tenant has reached the maximum capacity. Make sure that the tenant has sufficient capacity and then upload again. |
| 400 | 1061061 | user quota exceeded. | You have reached your maximum personal capacity. Make sure that you have sufficient capacity and then upload again. |
| 403 | 1061073 | no scope auth. | You have no access to the API. |
| 200 | 1064230 | locked_for_data_migration | Data migrating, temporarily unable to upload. |
