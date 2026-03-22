---
title: "Resource introduction"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/vc-v1/report/meeting-report-overview"
service: "vc-v1"
resource: "report"
section: "Video Conferencing"
updated: "1690944921000"
---

# Resource introduction
##  Resource definition
This resource is used to record a tenant's meeting usage information for a specified period of time, including: obtain meeting reports and obtain the top user list.

##  Field description
| Parameter | Type | Description |
| --- | --- | --- |
| `start_time` | `string` | Start time (unix time in sec) **Example Value**: "1608888867" |
| `end_time` | `string` | End time (unix time in sec) **Example Value**: "1608888966” |
| `limit` | `int` | Number of users to be obtained **Example value**: 10 |
| `order_by` | `int` | Sorting criteria (descending) **Example value**: 1 **Optional values are**: - `1`: Number of meetings - `2`: Meeting duration |
| `data` | `-` | `-` |
| ∟ `meeting_report` | `report` | Meeting report |
|  ∟ `total_meeting_count` | `string` | Total number of meetings |
|  ∟ `total_meeting_duration` | `string` | Total meeting duration (in sec) |
|  ∟ `total_participant_count` | `string` | Total number of participants |
|  ∟ `daily_report` | `report_meeting_daily[]` | List of daily meeting reports |
|  ∟ `date` | `string` | Date (unix time in sec) |
|  ∟ `meeting_count` | `string` | Number of meetings |
|  ∟ `meeting_duration` | `string` | Meeting duration (in sec) |
|  ∟ `participant_count` | `string` | Number of participants |
| `data` | `-` | `-` |
| ∟ `top_user_report` | `report_top_user[]` | Top user list |
|  ∟ `id` | `string` | User ID |
|  ∟ `name` | `string` | User name |
|  ∟ `user_type` | `int` | User type **Optional values are**: - `1`: lark user - `2`: rooms user - `3`: Docs user - `4`: neo Lark Meetings user - `5`: neo Lark Meetings guest - `6`: pstn user - `7`: sip user |
|  ∟ `meeting_count` | `string` | Number of meetings |
|  ∟ `meeting_duration` | `string` | Meeting duration (in sec) | ### Data example
```json
{
    "start_time":"1608888867",
    "end_time":"1608888966",
    "limit":10,
    "order_by":1,
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
    "data": {
        "top_user_report": [
            {
                "id": "ou_3ec3f6a28a0d08c45d895276e8e5e19b",
                "name": "name",
                "user_type": 1,
                "meeting_count": "100",
                "meeting_duration": "3000"
            }
        ]
    }
}
```
