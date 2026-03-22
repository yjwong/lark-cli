---
title: "Download minutes audio or video file"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/minutes-v1/minute-media/get"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/minutes/v1/minutes/:minute_token/media"
service: "minutes-v1"
resource: "minute-media"
section: "Minutes"
rate_limit: "5 per second"
scopes:
  - "minutes:minute:download"
  - "minutes:minutes.media:export"
updated: "1744024590000"
---

# Download audio or video file of minutes

Get the audio or video file of minutes

> Get the download link of minutes' audio or video file(Valid for 1 day
> ) through the interface for batch download

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/minutes/v1/minutes/:minute_token/media |
| HTTP Method | GET |
| Rate Limit | 5 per second |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `minutes:minute:download` `minutes:minutes.media:export` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `minute_token` | `string` | Minute unique identifier **Example value**: "obcnq3b9jl72l83w4f149w9c" **Data validation rules**: - Length range: `24` ～ `24` characters | ## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `download_url` | `string` | Download link for audio or video file | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
        "download_url": "https://internal-api-drive-stream.larksuite-boe.cn/space/api/box/stream/download/authcode/?code=xxx"
    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 2091001 | param is invalid | Check if the parameters are correct |
| 404 | 2091002 | resource not found | Unable to find the minute, check whether the minute_token is correct. |
| 400 | 2091003 | minute not ready , try later | Check if the transcription is complete |
| 400 | 2091004 | resource deleted | Check if the minute has been deleted |
| 403 | 2091005 | permission deny | Check if you have read permission for this minute |
| 500 | 2091006 | service internal error | Server error, please try again later |
