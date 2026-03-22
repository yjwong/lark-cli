---
title: "Query statistical data"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/attendance-v1/user_stats_data/query"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/attendance/v1/user_stats_datas/query"
service: "attendance-v1"
resource: "user_stats_data"
section: "Attendance"
rate_limit: "50 per second"
scopes:
  - "attendance:task"
  - "attendance:task:readonly"
updated: "1715149653000"
---

# Query statistical data

Queries statistical data for daily or monthly statistics.

> * Calling the attendance statistics open interface api does not currently return the new field type of leave and overwork
> * No data of a day if the member was not in the attendance group then.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/attendance/v1/user_stats_datas/query |
| HTTP Method | POST |
| Rate Limit | 50 per second |
| Supported app types | custom |
| Required scopes Enable any scope from the list | `attendance:task` `attendance:task:readonly` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `employee_type` | `string` | Yes | Type of employee ID for user_ids in request body and user_id in response body **Example value**: employee_id **Optional values are**:  - `employee_id`: Employee ID, that is the user ID in [Lark Admin](https://example.larksuite.com/admin/index) > Organization > Member and Department > Member Details - `employee_no`: Employee number, that is the employee ID in [Lark Admin](https://example.larksuite.com/admin/index) > Organization > Member and Department > Member Details  | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `locale` | `string` | Yes | Language **Example value**: "zh" **Optional values are**:  - `en`: English - `ja`: Japanese - `zh`: Chinese  |
| `stats_type` | `string` | Yes | Statistics type **Example value**: "month" **Optional values are**:  - `daily`: Daily statistics - `month`: Monthly statistics  |
| `start_date` | `int` | Yes | Start time **Example value**: 20210316 |
| `end_date` | `int` | Yes | End time (The time interval can't exceed 31 days.) **Example value**: 20210323 |
| `user_ids` | `string[]` | No | User ID list to query (Can't exceed 20 users) * Mandatory field ( system has upgrade, the new system requires a required field) **Example value**: ["ec8ddg56", "af5ddg73"] |
| `need_history` | `boolean` | No | Whether historical data is required **Example value**: true |
| `current_group_only` | `boolean` | No | Display only the current attendance group **Example value**: true |
| `user_id` | `string` | No | The user_id of the operator. * Each report of different operators (administrators) may have different field settings. The system will query the statistical data of the specified report based on user_id. * Required field (all have been upgraded to the new system, and the new system requires this field to be filled in). **Example value**: "ec8ddg56" | ### Request body example

{
    "locale": "zh",
    "stats_type": "month",
    "start_date": 20210316,
    "end_date": 20210323,
    "user_ids": [
        "ec8ddg56",
        "af5ddg73"
    ],
    "need_history": true,
    "current_group_only": true,
    "user_id": "ec8ddg56"
}

## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `user_datas` | `user_stats_data[]` | User statistics data（limit 1000） |
|     `name` | `string` | User name |
|     `user_id` | `string` | User ID |
|     `datas` | `user_stats_data_cell[]` | User's statistics data |
|       `code` | `string` | Field number |
|       `value` | `string` | Data value |
|       `features` | `user_stats_data_feature[]` | Data attribute |
|         `key` | `string` | The names of additional attributes of statistics data columns |
|         `value` | `string` | The values of additional attributes of statistics data columns * First display the clock-in results of commute to and from get off work, if there is an application for leave of absence, then display the time of application for leave of absence |
|       `title` | `string` | Data title |
|   `invalid_user_list` | `string[]` | List of users without permission | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
        "user_datas": [
            {
                "name": "Brown",
                "user_id": "ec8ddg56",
                "datas": [
                    {
                        "code": "50102",
                        "value": "none",
                        "features": [
                            {
                                "key": "Abnormal",
                                "value": "false"
                            }
                        ],
                        "title": "name"
                    }
                ]
            }
        ],
        "invalid_user_list": [
            "af5ddg73"
        ]
    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 1220001 | Parameter error | Please check if the parameters meet the requirements. |
| 400 | 1220002 | Tenant doesn't exist. | Please check if the tenant_access_token is correct. |
| 400 | 1220004 | User does not exist or has no scope | Check if the user ID is correct |
| 500 | 1228000 | Statistics service system error | See error information for details. |
| 400 | 1220600 | General error message | See error message for details |
