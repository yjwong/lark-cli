---
title: "Export minutes transcript"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/minutes-v1/minute-transcript/get"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/minutes/v1/minutes/:minute_token/transcript"
service: "minutes-v1"
resource: "minute-transcript"
section: "Minutes"
rate_limit: "5 per second"
scopes:
  - "minutes:minute:download"
  - "minutes:minutes.transcript:export"
updated: "1744024586000"
---

# Export minutes transcript

Get the transcript of the minute

> Download transcript through the interface for batch download

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/minutes/v1/minutes/:minute_token/transcript |
| HTTP Method | GET |
| Rate Limit | 5 per second |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `minutes:minute:download` `minutes:minutes.transcript:export` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `minute_token` | `string` | Minute unique identifier **Example value**: "obcnq3b9jl72l83w4f149w9c" **Data validation rules**: - Length range: `24` ～ `24` characters | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `need_speaker` | `boolean` | No | Does it include the speaker? **Example value**: true |
| `need_timestamp` | `boolean` | No | Does it include timestamp? **Example value**: true |
| `file_format` | `string` | No | Export file format, optional values include: - txt - srt **Example value**: txt | ## Response

When the HTTP status code is 200, it means success

Return file binary stream

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 2091001 | param is invalid | Check if the parameters are correct |
| 404 | 2091002 | resource not found | Unable to find the minute, check whether the minute_token is correct. |
| 400 | 2091003 | minute not ready , try later | Check if the transcription is complete |
| 400 | 2091004 | resource deleted | Check if the minute has been deleted |
| 403 | 2091005 | permission deny | Check if you have export permission for this minute |
| 500 | 2091006 | service internal error | Server error, please try again later |
