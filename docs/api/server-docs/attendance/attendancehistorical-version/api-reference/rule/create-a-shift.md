---
title: "Create a Shift"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/Attendance//shift_create"
section: "Attendance"
updated: "1646322860000"
---

# Create shifts
Shift is a general term used to describe the time rules for an attendance task, such as how many attendances are required for a shift, on and off time of each attendance, how long is the attendance delay being considered as late in, and how long is the attendance delay being considered as no record.

> **Note:** - You must create one or more shifts before creating an attendance group.
> - Shifts within a company are shared, and you can refer directly to the shifts created by others, but it should be noted that any modification of the shift by others will affect your attendance group and attendance results.
> 
## Request
|Facts||
|---|---|
|HTTP URL|https://open.larksuite.com/open-apis/attendance/v1/shifts|
|HTTP Method|POST|
|HTTP Content-Type|application/json; charset=utf-8|
|Token requirement|tenant_access_token|
|Required scopes|Add attendance management rules|
### Header
key|value
--|--
Authorization|Bearer tenant_access_token
Content-Type|application/json
### Request body
|Parameter|Type|Required|Description|
|---|---|---|---|
|shift_name|string|Yes|Shift name|
|punch_times|int|Yes|Number of clock-ins/outs. Less than or equal to 3.|
|is_flexible|boolean|No|Flextime or not|
|flexible_minutes|int|No|Flextime hours|
|no_need_off|boolean|No|Whether clock-out is not required|
|punch_time_rule|punch_time_rule[]|Yes|Attendance policy|
|  ∟on_time|string|Yes|Clock-in time, for example: "09:00"|
|  ∟off_time|string|Yes|Clock-out time, for example "18:00". Note: express 2:00 AM the next day as 26:00.|
|  ∟late_minutes_as_late|int|Yes|How many minutes late is late in|
|  ∟late_minutes_as_lack|int|Yes|How many minutes late is no record|
|  ∟on_advance_minutes|int|Yes|How early can employees clock in|
|  ∟early_minutes_as_early|int|Yes|How many minutes early is early out|
|  ∟early_minutes_as_lack|int|Yes|How many minutes early is no record|
|  ∟off_delay_minutes|int|Yes|How late can employees clock out|
|late_off_late_on_rule|late_off_late_on_rule[]|No|Late clock-in/out rules|
|  ∟late_off_minutes|int|Yes|Late clock-out time period|
|  ∟late_on_minutes|int|Yes|Late clock-in time period|
|rest_time_rule|rest_rule[]|No| Time rules of rest period|
|  ∟rest_begin_time|string|Yes|Start time of rest period, for example: "13:00"|
|  ∟rest_end_time|string|Yes|End time of rest period, for example: "14:00"|
### Request body example
```json
{
    "shift_name": "Morning shift",
    "punch_times": 1,
    "is_flexible": false,
    "flexible_minutes": 60,
    "no_need_off": true,
    "punch_time_rule": [
        {
            "on_time": "09:00",
            "off_time": "18:00",
            "late_minutes_as_late": 30,
            "late_minutes_as_lack": 60,
            "on_advance_minutes": 60,
            "early_minutes_as_early": 30,
            "early_minutes_as_lack": 60,
            "off_delay_minutes": 60
        }
    ],
    "late_off_late_on_rule": [
        {
            "late_off_minutes": 60,
            "late_on_minutes": 30
        }
    ],
    "rest_time_rule": [
        {
            "rest_begin_time": "13:00",
            "rest_end_time": "14:00"
        }
    ]
}
```
## Response
### Response body
|Parameter|Type|Description|
|---|---|---|
|code|int|Error code. A value other than 0 indicates failure.|
|msg|string|Error description|
|data|-|-|
|  ∟shift|shift|Shift|
|    ∟shift_id|string|Shift ID|
|    ∟shift_name|string|Shift name|
|    ∟punch_times|int|Clock-in/out count|
|    ∟is_flexible|boolean|Flextime or not|
|    ∟flexible_minutes|int|Flextime hours|
|    ∟no_need_off|boolean|Whether clock-out is not required|
|    ∟punch_time_rule|punch_time_rule[]|Attendance policy|
|      ∟on_time|string|Clock-in time|
|      ∟off_time|string|Clock-out time|
|      ∟late_minutes_as_late|int|How many minutes late is late in|
|      ∟late_minutes_as_lack|int|How many minutes late is no record|
|      ∟on_advance_minutes|int|How early can employees clock in|
|      ∟early_minutes_as_early|int|How many minutes early is early out|
|      ∟early_minutes_as_lack|int|How many minutes early is no record|
|      ∟off_delay_minutes|int|How late can employees clock out|
|    ∟late_off_late_on_rule|late_off_late_on_rule[]|Late clock-in/out rules|
|      ∟late_off_minutes|int|Late clock-out time period|
|      ∟late_on_minutes|int|Late clock-in time period|
|    ∟rest_time_rule|rest_rule[]|Time rules of rest period|
|      ∟rest_begin_time|string|Start time of rest period|
|      ∟rest_end_time|string|End time of rest period|
### Response body example
```json
{
    "code": 0,
    "msg": "success",
    "data": {
        "shift": {
            "shift_id": "6919358778597097404",
            "shift_name": "Morning shift",
            "punch_times": 1,
            "is_flexible": false,
            "flexible_minutes": 60,
            "no_need_off": true,
            "punch_time_rule": [
                {
                    "on_time": "09:00",
                    "off_time": "18:00",
                    "late_minutes_as_late": 30,
                    "late_minutes_as_lack": 60,
                    "on_advance_minutes": 60,
                    "early_minutes_as_early": 30,
                    "early_minutes_as_lack": 60,
                    "off_delay_minutes": 60
                }
            ],
            "late_off_late_on_rule": [
                {
                    "late_off_minutes": 60,
                    "late_on_minutes": 30
                }
            ],
            "rest_time_rule": [
                {
                    "rest_begin_time": "13:00",
                    "rest_end_time": "14:00"
                }
            ]
        }
    }
}
```
### Error code
|HTTP status code|Error code|Description|Troubleshooting suggestions|
|---|---|---|---|
|400|1220001|Parameter error|Please check if the parameters meet the requirements|
|400|1220002|Tenant doesn't exist|Please check if the tenant_access_token is correct|
|400|1220005|No permission scope|Please go to the Attendance Admin to check the data permission scope|
|500|1225000|System error|See error information for details|
|500|1226000|Shift service system error|See error message for details|
|400|1226001|Shift already in use|Use another shift name|
|400|1226002|Shift name already in use|Use another shift name|
