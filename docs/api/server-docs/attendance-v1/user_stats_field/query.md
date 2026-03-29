---
title: "Query Statistics Header"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/attendance-v1/user_stats_field/query"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/attendance/v1/user_stats_fields/query"
service: "attendance-v1"
resource: "user_stats_field"
section: "Attendance"
scopes:
  - "attendance:task"
  - "attendance:task:readonly"
updated: "1646322853000"
---

# Query statistics headers

Queries the statistics headers for daily or monthly statistics supported by the attendance statistics.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/attendance/v1/user_stats_fields/query |
| HTTP Method | POST |
| Supported app types | custom |
| Required scopes Enable any scope from the list | `attendance:task` `attendance:task:readonly` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" Learn more about obtaining and using access_token. |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `employee_type` | `string` | Yes | Type of employee ID for user_id in response body **Example value**: "employee_id" **Optional values are**: - `employee_id`: Employee ID, that is the user ID in Lark Admin > Organization > Member and Department > Member Details - `employee_no`: Employee number, that is the employee ID in Lark Admin > Organization > Member and Department > Member Details | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `locale` | `string` | Yes | Language **Example value**: "zh" **Optional values are**: - `en`: English - `ja`: Japanese - `zh`: Chinese |
| `stats_type` | `string` | Yes | Statistics type **Example value**: "daily" **Optional values are**: - `daily`: Daily statistics - `month`: Monthly statistics |
| `start_date` | `int` | Yes | Start time **Example value**: 20210316 |
| `end_date` | `int` | Yes | End time (The time interval can't exceed 40 days.) **Example value**: 20210323 | ### Request body example

```json
{
    "locale": "zh",
    "stats_type": "daily",
    "start_date": 20210316,
    "end_date": 20210323
}
```

## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
| ∟ `user_stats_field` | `user_stats_field` | Statistics data header |
| ∟ `stats_type` | `string` | Statistics type **Optional values are**: - `daily`: Daily statistics - `month`: Monthly statistics |
| ∟ `user_id` | `string` | User ID |
| ∟ `fields` | `field[]` | Field list |
| ∟ `code` | `string` | Field number |
| ∟ `title` | `string` | Field name |
| ∟ `child_fields` | `child_field[]` | Sub-field list |
| ∟ `code` | `string` | Sub-field No. |
| ∟ `title` | `string` | Sub-field name |
| ∟ `time_unit` | `string` | Time unit | ### Response body example

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
                            "title": "Employee ID"
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
                            "title": "Required attendance days"
                        },
                        {
                            "code": "52102",
                            "title": "Days of attendance"
                        },
                        {
                            "code": "52104",
                            "time_unit": "Minute",
                            "title": "Required attendance duration"
                        },
                        {
                            "code": "52105",
                            "time_unit": "Minute",
                            "title": "Actual attendance duration"
                        },
                        {
                            "code": "52107",
                            "title": "Overtime hours"
                        }
                    ],
                    "code": "521",
                    "title": "Attendance statistics"
                },
                {
                    "child_fields": [
                        {
                            "code": "52201",
                            "title": "Late in times"
                        },
                        {
                            "code": "52203",
                            "title": "Early out times"
                        },
                        {
                            "code": "52207",
                            "title": "No records"
                        }
                    ],
                    "code": "522",
                    "title": "Abnormal statistics"
                },
                {
                    "child_fields": [
                        {
                            "code": "2021-03-16",
                            "title": "2021-03-16 Tue"
                        },
                        {
                            "code": "2021-03-17",
                            "title": "2021-03-17 Wed"
                        },
                        {
                            "code": "2021-03-18",
                            "title": "2021-03-18 Thu"
                        },
                        {
                            "code": "2021-03-19",
                            "title": "2021-03-19 Fri"
                        },
                        {
                            "code": "2021-03-20",
                            "title": "2021-03-20 Sat"
                        },
                        {
                            "code": "2021-03-21",
                            "title": "2021-03-21 Sun"
                        },
                        {
                            "code": "2021-03-22",
                            "title": "2021-03-22 Mon"
                        },
                        {
                            "code": "2021-03-23",
                            "title": "2021-03-23 Tue"
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
| 400 | 1220001 | Parameter error | Please check if the parameters meet the requirements. |
| 400 | 1220002 | Tenant doesn't exist. | Please check if the tenant_access_token is correct. |
| 500 | 1228000 | Statistics service system error | See error information for details. |
