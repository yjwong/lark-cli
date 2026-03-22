---
title: "Download Files"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/file/download"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/drive/v1/files/:file_token/download"
service: "drive-v1"
resource: "file"
section: "Docs"
scopes:
  - "drive:file"
  - "drive:file:readonly"
  - "drive:drive:readonly"
  - "drive:drive"
updated: "1665221894000"
---

# Download a file

This API is used to download files from a directory in My Space, excluding online documents in Lark Docs, Sheets, and MindNotes. You can download a file based on a specified range.

> This API supports up to 5 queries per second (QPS).

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/drive/v1/files/:file_token/download |
| HTTP Method | GET |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `drive:file` `drive:file:readonly` `drive:drive:readonly` `drive:drive` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" Learn more about obtaining and using access_token. | ::: note
For more information about how to call the AccessToken in the Docs API, see Get started with the Docs API.

### Download based on a range
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Range | string | No | Specifies the part of the file that you need to download. **Example value**: "bytes=0-1024" | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `file_token` | `string` | Token of the file. For more information about how to obtain the token, see Overview. **Example value**: "boxcnabCdefg12345" | ## Response

### Response header
| `content-type` | `string` | MIME of the file |
| --- | --- | --- |
| `content-disposition` | `string` | File name | When the HTTP status code is 200, it means success

Return file binary stream
