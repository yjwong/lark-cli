---
title: "Start recording"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/vc-v1/meeting-recording/start"
method: "PATCH"
api_path: "https://open.larksuite.com/open-apis/vc/v1/meetings/:meeting_id/recording/start"
service: "vc-v1"
resource: "meeting-recording"
section: "Video Conferencing"
rate_limit: "Special Rate Limit"
scopes:
  - "vc:record"
updated: "1701080966000"
---

# Start recording

Start recording in a meeting. 

> The meeting must be in progress and the operator should have relevant permission scope. (If the operator is a user, they must be the current host of the meeting.)

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/vc/v1/meetings/:meeting_id/recording/start |
| HTTP Method | PATCH |
| Rate Limit | Special Rate Limit |
| Supported app types | custom,isv |
| Required scopes | `vc:record` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `meeting_id` | `string` | Meeting ID (Unique identifier of video conference, generated after the conference is started) **Example value**: "6911188411932033028" | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `timezone` | `int` | No | Time zone used for displaying the file recording time [-12,12] **Example value**: 8 | ### Request body example

{
    "timezone": 8
}

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
    "data": {}
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 500 | 121001 | internal error | Internal server error. If retry still doesn't work, contact the administrator. |
| 400 | 121002 | not support | This function is not supported. |
| 400 | 121003 | param error | Parameter error. Check the value range of parameters (please check yourself according to the field description above). |
| 403 | 121005 | no permission | No permission. Check the token type, operator identity and resource. |
| 400 | 122001 | meeting status unexpected | Meeting status is incorrect. For example, the meeting is not in progress. |
| 404 | 122002 | meeting not exist | The meeting does not exist. |
| 400 | 122003 | operator user not in meeting | The operator must be in the meeting. |
