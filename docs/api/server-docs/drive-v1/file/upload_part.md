---
title: "Multipart Upload Files (Upload)"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/file/upload_part"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/drive/v1/files/upload_part"
service: "drive-v1"
resource: "file"
section: "Docs"
scopes:
  - "drive:drive"
  - "drive:file"
updated: "1665221894000"
---

# Upload a file in blocks (Upload blocks)

Uploads blocks of a specified file.

> This API doesn't support high concurrency, and supports up to 5 queries per second (QPS).

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/drive/v1/files/upload_part |
| HTTP Method | POST |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `drive:drive` `drive:file` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" Learn more about obtaining and using access_token. |
| Content-Type | string | Yes | **Example value**: "multipart/form-data; boundary=---7MA4YWxkTrZu0gW" | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `upload_id` | `string` | Yes | ID of the block upload transaction. **Example value**: "7111211691345512356" |
| `seq` | `int` | Yes | Sequence number of the block, which starts from 0. **Example value**: 0 |
| `size` | `int` | Yes | Block size in bytes. **Example value**: 4194304 |
| `checksum` | `string` | No | Adler-32 checksum of the file block. This field is optional. **Example value**: "12342388237783212356" |
| `file` | `file` | Yes | Binary content of the file block. **Example value**: file binary | ### Request body example

```HTTP
---7MA4YWxkTrZu0gW
Content-Disposition: form-data; name="upload_id";

7111211691345512356
---7MA4YWxkTrZu0gW
Content-Disposition: form-data; name="seq";

0
---7MA4YWxkTrZu0gW
Content-Disposition: form-data; name="size";

4194304
---7MA4YWxkTrZu0gW
Content-Disposition: form-data; name="checksum";

12342388237783212356
---7MA4YWxkTrZu0gW
Content-Disposition: form-data; name="file";
Content-Type: application/octet-stream

file binary
---7MA4YWxkTrZu0gW
```

## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {

    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 200 | 1061001 | internal error. | Internal service error, such as timeout or failure in processing error codes. |
| 400 | 1061002 | params error. | Check whether the request parameters are correct. |
| 400 | 1061021 | upload id expire. | The upload transaction has expired. Please upload again. |
| 403 | 1061073 | no scope auth. | You have no access to the API. |
| 400 | 1062007 | upload user not match. | Make sure that the current request is sent by the same user or app as the upload task. |
| 400 | 1062008 | checksum param Invalid. | Make sure that the checksum of the file or file block is correct. |
| 400 | 1062009 | the actual size is inconsistent with the parameter declaration size. | The size of the file to transfer is inconsistent with that specified in the parameter. |
| 400 | 1062010 | block missing, please upload all blocks. | Some file blocks are missing. Make sure that all file blocks are uploaded. |
| 400 | 1062011 | block num out of bounds. | The number of file blocks to upload has reached the limit. Make sure that the file blocks belong to the specified file. |
