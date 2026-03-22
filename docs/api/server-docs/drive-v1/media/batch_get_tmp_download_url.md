---
title: "Get Temporary Download URL of Media"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/media/batch_get_tmp_download_url"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/drive/v1/medias/batch_get_tmp_download_url"
service: "drive-v1"
resource: "media"
section: "Docs"
scopes:
  - "vc:material:readonly"
  - "vc:material"
  - "moments:moments:readonly"
  - "drive:drive:readonly"
  - "bitable:app:readonly"
  - "docs:doc:readonly"
  - "sheets:spreadsheet:readonly"
  - "bitable:app"
  - "docs:doc"
  - "sheets:spreadsheet"
  - "drive:drive"
  - "moments:moments"
updated: "1647178822000"
---

# Obtain a temporary material download URL

Obtain the temporary download URL of a material based on a file_token. The URL is valid for 24 hours.

> This API doesn't support high concurrency, and supports up to 5 queries per second (QPS).

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/drive/v1/medias/batch_get_tmp_download_url |
| HTTP Method | GET |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `vc:material:readonly` `vc:material` `moments:moments:readonly` `drive:drive:readonly` `bitable:app:readonly` `docs:doc:readonly` `sheets:spreadsheet:readonly` `bitable:app` `docs:doc` `sheets:spreadsheet` `drive:drive` `moments:moments` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" Learn more about obtaining and using access_token. | ::: note
For more information about how to call the AccessToken in the Docs API, see Get started with the Docs API.

### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `file_tokens` | `string[]` | Yes | List of file tokens **Example value**: Boxcnabcdefg |
| `extra` | `string` | No | Extension information. This field is optional. **Example value**: "For more information, see Types and the Extra parameter of upload points." | ## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
| ∟ `tmp_download_urls` | `tmp_download_url[]` | Temporary download list |
| ∟ `file_token` | `string` | Token of the file |
| ∟ `tmp_download_url` | `string` | Temporary download URL of the file | ### Response body example

```json
{
    "code": 0,
    "msg": "success",
    "data": {
        "tmp_download_urls": [
            {
                "file_token": "boxcnabcdefg",
                "tmp_download_url": "https://internal-api-drive-stream.larksuite.com/space/api/box/stream/download/authcode/?code

drive.v1.media.method.batch_get_tmp_download_url.query.prop.file_tokens.list.example=$$$boxcnabcdefg"
            }
        ]
    }
}
```
