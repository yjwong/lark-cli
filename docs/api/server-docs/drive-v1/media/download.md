---
title: "Download Media"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/media/download"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/drive/v1/medias/:file_token/download"
service: "drive-v1"
resource: "media"
section: "Docs"
scopes:
  - "bitable:app"
  - "bitable:app:readonly"
  - "docs:doc"
  - "docs:doc:readonly"
  - "drive:drive"
  - "drive:drive:readonly"
  - "moments:moments"
  - "moments:moments:readonly"
  - "sheets:spreadsheet"
  - "sheets:spreadsheet:readonly"
  - "vc:material"
  - "vc:material:readonly"
updated: "1715857429000"
---

# Download a material

This API is used to download materials. Materials refer to files in various production containers, for example, images in a document and files. You can download a material based on a specified range.

> This API doesn't support high concurrency, and supports up to 5 queries per second (QPS).

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/drive/v1/medias/:file_token/download |
| HTTP Method | GET |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `bitable:app` `bitable:app:readonly` `docs:doc` `docs:doc:readonly` `drive:drive` `drive:drive:readonly` `moments:moments` `moments:moments:readonly` `sheets:spreadsheet` `sheets:spreadsheet:readonly` `vc:material` `vc:material:readonly` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" Learn more about obtaining and using access_token. | ::: note
For more information about how to call the AccessToken in the Docs API, see Get started with the Docs API.

### Download based on a range
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Range | string | No | Specifies the part of the file that you need to download. For example, set bytes to 0-1024. | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `file_token` | `string` | Token of the file. For more information about how to obtain the token, see Overview. **Example value**: "boxcnrHpsg1QDqXAAAyachabcef" | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `extra` | `string` | No | Base with advanced permissions require additional extended information as URL query parameters for authentication during material downloads. For details, refer to the Types and the Extra parameter of upload points. Interfaces that do not correctly fill in this parameter will return an HTTP status code of 403. **Example value**: "None" | ## Response

### Response header
| `content-type` | `string` | MIME of the file |
| --- | --- | --- |
| `content-disposition` | `string` | File name | When the HTTP status code is 200, it means success

Return file binary stream
