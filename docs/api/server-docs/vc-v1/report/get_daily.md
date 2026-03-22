---
title: "Obtain meeting reports"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/vc-v1/report/get_daily"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/vc/v1/reports/get_daily"
service: "vc-v1"
resource: "report"
section: "Video Conferencing"
rate_limit: "Special Rate Limit"
scopes:
  - "vc:report:readonly"
updated: "1690944921000"
---

# Obtain meeting reports

Obtains an organizations' daily meeting usage reports for a period of time.

> Data for the last 90 days can be queried.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/vc/v1/reports/get_daily |
| HTTP Method | GET |
| Rate Limit | Special Rate Limit |
| Supported app types | custom,isv |
| Required scopes | `vc:report:readonly` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `start_time` | `string` | Yes | Start time (unix time in sec) **Example value**: 1608888867 |
| `end_time` | `string` | Yes | End time (unix time in sec) **Example value**: 1608888966 | ## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `meeting_report` | `report` | Meeting reports |
|     `total_meeting_count` | `string` | Total number of meetings |
|     `total_meeting_duration` | `string` | Total meeting duration (in sec) |
|     `total_participant_count` | `string` | Total number of participants |
|     `daily_report` | `report_meeting_daily[]` | List of daily meeting reports |
|       `date` | `string` | Date (unix time in sec) |
|       `meeting_count` | `string` | Number of meetings |
|       `meeting_duration` | `string` | Meeting duration (in sec) |
|       `participant_count` | `string` | Number of participants | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
        "meeting_report": {
            "total_meeting_count": "100",
            "total_meeting_duration": "300000",
            "total_participant_count": "20000",
            "daily_report": [
                {
                    "date": "1609113600",
                    "meeting_count": "100",
                    "meeting_duration": "147680",
                    "participant_count": "2000"
                }
            ]
        }
    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 500 | 121001 | internal error | Internal server error. If retry still doesn't work, contact the administrator. |
| 400 | 121002 | not support | This function is not supported. |
| 400 | 121003 | param error | Parameter error. Check the value range of parameters (please check yourself according to the field description above). |
