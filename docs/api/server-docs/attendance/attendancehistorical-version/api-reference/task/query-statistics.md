---
title: "Query Statistics"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/Attendance//task/query-statistics-data"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/attendance/v1/user_stats_datas/query"
section: "Attendance"
updated: "1646322862000"
---

# Query statistical data
This API is used to query statistical data for daily or monthly statistics.

## Request
| HTTP URL | https://open.larksuite.com/open-apis/attendance/v1/user_stats_datas/query |
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
| `start_date` | `int` | Yes | Start time **Example value**: 20210316 |
| `end_date` | `int` | Yes | End time **Example value**: 20210323 (The time interval can't exceed 40 days) |
| `user_ids` | `string[]` | No | User ID list to query (Can't exceed 20 users) |
| `need_history` | `boolean` | No | Contains historical data **Example value**: true |
| `current_group_only` | `boolean` | No | Only contains current attendance group **Example value**: true | ### Request body example

```json
{
    "current_group_only": true,
    "end_date": 20210323,
    "locale": "zh",
    "need_history": true,
    "start_date": 20210316,
    "stats_type": "month",
    "user_ids": [
        "ec8ddg56",
        "4dbb52f2",
        "4167842e"
    ]
}
```

## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error code. A value other than 0 indicates failure. |
| `msg` | `string` | Error description |
| `data` | `\-` | \- |
| ∟ `user_datas` | `user_stats_data[]` | User statistics data |
| ∟ `name` | `string` | Name |
| ∟ `user_id` | `string` | User ID |
| ∟ `datas` | `user_stats_data_cell[]` | Data |
| ∟ `code` | `string` | Field number |
| ∟ `value` | `string` | Data value |
| ∟ `features` | `user_stats_data_feature[]` | Data attribute |
| ∟ `key` | `string` | Property |
| ∟ `value` | `string` | Property value | ### Response body example

```json
{
    "code": 0,
    "msg": "",
    "data": {
        "user_datas": [
            {
                "datas": [
                    {
                        "code": "2021-03-19",
                        "features": [
                            {
                                "key": "Abnormal",
                                "value": "true"
                            }
                        ],
                        "value": "No record (-), No record (-)"
                    },
                    {
                        "code": "2021-03-22",
                        "features": [
                            {
                                "key": "Abnormal",
                                "value": "true"
                            }
                        ],
                        "value": "No record (-), No record (-)"
                    },
                    {
                        "code": "2021-03-16",
                        "features": [
                            {
                                "key": "Abnormal",
                                "value": "true"
                            }
                        ],
                        "value": "No record (-), No record (-)"
                    },
                    {
                        "code": "2021-03-21",
                        "features": [
                            {
                                "key": "Abnormal",
                                "value": "false"
                            }
                        ],
                        "value": "No attendance required (-), No attendance required (-)"
                    },
                    {
                        "code": "50103",
                        "features": [],
                        "value": "3766663"
                    },
                    {
                        "code": "52107",
                        "features": [],
                        "value": "0.0 hours"
                    },
                    {
                        "code": "2021-03-20",
                        "features": [
                            {
                                "key": "Abnormal",
                                "value": "false"
                            }
                        ],
                        "value": "No attendance required (-), No attendance required (-)"
                    },
                    {
                        "code": "52203",
                        "features": [],
                        "value": "0"
                    },
                    {
                        "code": "52101",
                        "features": [],
                        "value": "5 days"
                    },
                    {
                        "code": "52105",
                        "features": [],
                        "value": "0"
                    },
                    {
                        "code": "52201",
                        "features": [],
                        "value": "0"
                    },
                    {
                        "code": "52207",
                        "features": [],
                        "value": "5 days"
                    },
                    {
                        "code": "2021-03-18",
                        "features": [
                            {
                                "key": "Abnormal",
                                "value": "true"
                            }
                        ],
                        "value": "No record (-), No record (-)"
                    },
                    {
                        "code": "52108",
                        "features": [],
                        "value": "Schedule test"
                    },
                    {
                        "code": "52102",
                        "features": [],
                        "value": "0 days"
                    },
                    {
                        "code": "52104",
                        "features": [],
                        "value": "2700"
                    },
                    {
                        "code": "2021-03-23",
                        "features": [
                            {
                                "key": "Abnormal",
                                "value": "false"
                            }
                        ],
                        "value": "No attendance required (-), No attendance required (-)"
                    },
                    {
                        "code": "2021-03-17",
                        "features": [
                            {
                                "key": "Abnormal",
                                "value": "true"
                            }
                        ],
                        "value": "No record (-), No record (-)"
                    }
                ],
                "name": "John",
                "user_id": "ec8ddg56"
            }
        ]
    }
}
```

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 1220001 | Parameter error | Please check if the parameters meet the requirements |
| 400 | 1220002 | Tenant doesn't exist | Please check if the tenant_access_token is correct |
| 500 | 1228000 | Statistics service system error | See error information for details |
