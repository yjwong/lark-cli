---
title: "Get Shift Details"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/Attendance//shift_by_id"
section: "Attendance"
updated: "1646322860000"
---

# Obtain shift details
This API is used to obtain shift details via a shift ID.

## Request
|Facts||
|---|---|
|HTTP URL|https://open.larksuite.com/open-apis/attendance/v1/shifts/:shift_id|
|HTTP Method|GET|
|HTTP Content-Type|application/json; charset=utf-8|
|Token requirement|tenant_access_token|
|Required scopes|Add or export attendance management rules
### Path parameters
|Parameter|Type|Required|Description|
|---|---|---|---|
|shift_id|string|Yes|Shift ID. Example value: "6919358778597097404"|
### Header
key|value
--|--
Authorization|Bearer tenant_access_token
Content-Type|application/json
## Response
### Response body
|Parameter|Type|Description|
|---|---|---|
|code|int|Error code. A value other than 0 indicates failure.|
|msg|string|Error description|
|data|-|-|
|  ∟shift_id|string|Shift ID|
|  ∟shift_name|string|Shift name|
|  ∟punch_times|int|Clock-in/out count|
|  ∟is_flexible|boolean|Flextime or not|
|  ∟flexible_minutes|int|Flextime hours|
|  ∟no_need_off|boolean|Whether clock-out is not required|
|  ∟punch_time_rule|punch_time_rule[]|Attendance policy|
|    ∟on_time|string|Clock-in time|
|    ∟off_time|string|Clock-out time|
|    ∟late_minutes_as_late|int|How many minutes late is late in|
|    ∟late_minutes_as_lack|int|How many minutes late is no record|
|    ∟on_advance_minutes|int|How early can employees clock in|
|    ∟early_minutes_as_early|int|How many minutes early is early out|
|    ∟early_minutes_as_lack|int|How many minutes early is no record|
|    ∟off_delay_minutes|int|How late can employees clock out|
|  ∟late_off_late_on_rule|late_off_late_on_rule[]|Late clock-in/out rules|
|    ∟late_off_minutes|int|Late clock-out time period|
|    ∟late_on_minutes|int|Late clock-in time period|
|  ∟rest_time_rule|rest_rule[]|Time rules of rest period|
|    ∟rest_begin_time|string|Start time of rest period|
|    ∟rest_end_time|string|End time of rest period|
### Response body example
```json
{
    "code": 0,
    "msg": "success",
    "data": {
        "shift_id": "6919358778597097404",
        "shift_name": "Morning shift",
        "punch_times": 1,
        "is_flexible": false,
        "flexible_minutes": 60,
        "no_need_off": true,
        "punch_time_rule": [
            {
                "on_time": "9:00",
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
```
### Error code
|HTTP status code|Error code|Description|Troubleshooting suggestions|
|---|---|---|---|
|400|1220001|Parameter error|Please check if the parameters meet the requirements|
|400|1220002|Tenant doesn't exist|Please check if the tenant_access_token is correct|
|500|1225000|System error|See error information for details|
|500|1226000|Shift service system error|See error message for details|
|400|1226003|Shift doesn't exist|Check if the shift_id is correct|
