---
title: "Obtain recording files"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/vc-v1/meeting-recording/get"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/vc/v1/meetings/:meeting_id/recording"
service: "vc-v1"
resource: "meeting-recording"
section: "Video Conferencing"
rate_limit: "1000 per minute、50 per second"
scopes:
  - "vc:record:readonly"
updated: "1681980229000"
---

# Obtain recording files

Obtain recording files of a meeting. 

> Recording files can be obtained only after the meeting is ended and “Recording is completed” event is received; Using user_access_token, only the meeting owner (the person who reserves the meeting on the open platform is considered as reserving person) has permission to obtain it, and using tenant_access_token can obtain the recording files under the scope of the tenant; too short recording time (<5s) may result in failure to generate recording files

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/vc/v1/meetings/:meeting_id/recording |
| HTTP Method | GET |
| Rate Limit | 1000 per minute、50 per second |
| Supported app types | custom,isv |
| Required scopes | `vc:record:readonly` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `meeting_id` | `string` | Meeting ID (Unique identifier of a video conference, generated after the conference is started) **Example value**: "6911188411932033028" | ## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `recording` | `meeting.recording` | Recording file data |
|     `url` | `string` | Recording file URL |
|     `duration` | `string` | Total recording time (in ms) | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
        "recording": {
            "url": "https://www.larksuite.com/minutes/obcn37dxcftoc3656rgyejm7",
            "duration": "30000"
        }
    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 500 | 121001 | internal error | Internal server error. If retry still doesn't work, contact the administrator. |
| 400 | 121002 | not support | This function is not supported. |
| 400 | 121003 | param error | Parameter error. Check the value range of parameters (please check yourself according to the field description above). |
| 404 | 121004 | data not exist | The requested data doesn't exist. |
| 403 | 121005 | no permission | No permission. Check the token type, operator identity and resource. |
| 400 | 122001 | meeting status unexpected | Meeting status is incorrect. For example, the meeting is not in progress. |
| 404 | 122002 | meeting not exist | The meeting does not exist. |
| 400 | 124002 | record processing | Recording file is being generated |
