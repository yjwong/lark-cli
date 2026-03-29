---
title: "Query the Statistics Table Header"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/Attendance//task/query-statistics-header"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/attendance/v1/user_stats_fields/query"
section: "Attendance"
updated: "1646322864000"
---

# Query statistics headers
This API is used to query statistics headers for daily or monthly statistics.

## Request
| HTTP URL | https://open.larksuite.com/open-apis/attendance/v1/user_stats_fields/query |
| --- | --- |
| HTTP Method | POST |
| Required scopes | Add attendance data Export attendance data | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" Learn more about how to access and use access_token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Query parameters
| `employee_type` | `string` | Yes | User ID type **Optional values are**: - `employee_id` - `employee_no` |
| --- | --- | --- | --- | ### Request body
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `locale` | `string` | Yes | Language **Optional values are**: - `en`: English - `ja`: Japanese - `zh`: Chinese |
| `stats_type` | `string` | Yes | Statistics type **Optional values are**: - `daily`: Daily statistics - `month`: Monthly statistics |
| `start_date` | `int` | Yes | Start time **Example value**: 20210316 (The time interval can't exceed 40 days) |
| `end_date` | `int` | Yes | End time **Example value**: 20210323 | ### Request body example

```json
{
    "locale": "zh",
    "stats_type": "month",
    "start_date": 20210316,
    "end_date": 20210323
}
```

## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error code. A value other than 0 indicates failure. |
| `msg` | `string` | Error description |
| `data` | `\-` | \- |
| ∟ `user_stats_field` | `user_stats_field` | Statistics data header |
| ∟ `stats_type` | `string` | Statistics type **Optional values are**: - `daily`: Daily statistics - `month`: Monthly statistics |
| ∟ `user_id` | `string` | User ID |
| ∟ `fields` | `field[]` | Field list |
| ∟ `code` | `string` | Field number |
| ∟ `title` | `string` | Field heading |
| ∟ `child_fields` | `child_field[]` | Sub-field list |
| ∟ `code` | `string` | Field number |
| ∟ `title` | `string` | Field name |
| ∟ `time_unit` | `string` | Time type | ### Response body example

```json
{
    "code": 0,
    "msg": "",
    "data": {
        "user_stats_field": {
            "fields": [
                {
                    "child_fields": [
                        {
                            "code": "50103",
                            "title": "Employee number"
                        }
                    ],
                    "code": "501",
                    "title": "Basic info"
                },
                {
                    "child_fields": [
                        {
                            "code": "52108",
                            "title": "Attendance group name"
                        },
                        {
                            "code": "52101",
                            "title": "Expected on-duty days"
                        },
                        {
                            "code": "52102",
                            "title": "On-duty days on working days"
                        },
                        {
                            "code": "52104",
                            "time_unit": "min",
                            "title": "Expected working time"
                        },
                        {
                            "code": "52105",
                            "time_unit": "min",
                            "title": "Actual working time"
                        },
                        {
                            "code": "52107",
                            "title": "Overtime working hours"
                        }
                    ],
                    "code": "521",
                    "title": "Attendance statistics"
                },
                {
                    "child_fields": [
                        {
                            "code": "52201",
                            "title": "Late in count"
                        },
                        {
                            "code": "52203",
                            "title": "Early out count"
                        },
                        {
                            "code": "52207",
                            "title": "Absence"
                        }
                    ],
                    "code": "522",
                    "title": "Abnormal statistics"
                },
                {
                    "child_fields": [
                        {
                            "code": "2021-03-16",
                            "title": "2021-03-16 Tuesday"
                        },
                        {
                            "code": "2021-03-17",
                            "title": "2021-03-17 Wednesday"
                        },
                        {
                            "code": "2021-03-18",
                            "title": "2021-03-18 Thursday"
                        },
                        {
                            "code": "2021-03-19",
                            "title": "2021-03-19 Friday"
                        },
                        {
                            "code": "2021-03-20",
                            "title": "2021-03-20 Saturday"
                        },
                        {
                            "code": "2021-03-21",
                            "title": "2021-03-21 Sunday"
                        },
                        {
                            "code": "2021-03-22",
                            "title": "2021-03-22 Monday"
                        },
                        {
                            "code": "2021-03-23",
                            "title": "2021-03-23 Tuesday"
                        }
                    ],
                    "code": "524",
                    "title": "Daily statistics"
                }
            ],
            "stats_type": "month",
            "user_id": "ec8ddg56"
        }
    }
}
```

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 1220001 | Parameter error | Please check if the parameters meet the requirements |
| 400 | 1220002 | Tenant doesn't exist | Please check if the tenant_access_token is correct |
| 500 | 1228000 | Statistics service system error | See error information for details |
