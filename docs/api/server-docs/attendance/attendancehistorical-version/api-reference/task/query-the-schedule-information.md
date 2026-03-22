---
title: "Query the Schedule Information"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/Attendance//task/GetScheduledShifts"
section: "Attendance"
updated: "1646322869000"
---

# Query schedule information
This API is used to query schedule information for multiple users. The query time span can't exceed 30 days.

## Request
|Facts||
|---|---|
|HTTP URL|https://open.larksuite.com/open-apis/attendance/v1/user_daily_shifts/query|
|HTTP Method|POST|
|HTTP Content-Type|application/json; charset=utf-8|
|Token requirement|tenant_access_token|
|Required scopes|Attendance data export|
### Header
key|value
--|--
Authorization|Bearer tenant_access_token
Content-Type|application/json
### Query parameters
|Parameter|Type|Required|Description|
|---|---|---|---|
|employee_type|string|Yes|The type of the user_ids in the request body, values: [employee_id (the employee's employee ID), employee_no (the employee's employee number)]. Example value: "employee_id"| ### Request body
|Parameter|Type|Required|Description|
|---|---|---|---|
|user_ids|string[]|Yes|employee_no or employee_id list|
|check_date_from|int|Yes|Start day of check|
|check_date_to|int|true|End working day of query|
### Request body example
```json
{
    "user_ids": [
        "abd754f7"
    ],
    "check_date_from": 20210121,
    "check_date_to": 20210122
}
```
## Response
### Response body
|Parameter|Type|Description|
|---|---|---|
|code|int|Error code. A value other than 0 indicates failure.|
|msg|string|Error description|
|data|-|-|
|  ∟user_daily_shifts|user_daily_shift[]|List of schedule information|
|    ∟group_id|string|Attendance group ID|
|    ∟shift_id|string|Shift ID, 0 indicates a day off|
|    ∟month|int|Month|
|    ∟employee_no|string|User|
|    ∟day_no|int|Date|
### Response body example
```json
{
    "code": 0,
    "msg": "success",
    "data": {
        "user_daily_shifts": [
            {
                "group_id": "6737202939523236110",
                "shift_id": "6753520403404030215",
                "month": 202101,
                "user_id": "abd754f7",
                "day_no": 21
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
|400|1220004|User doesn't exist or isn't in permission scope|Please check if the user ID is correct|
|400|1220005|No permission scope|Please go to the Attendance Admin to check the data permission scope|
|500|1225000|System error|See error information for details|
|500|1226000|Shift service system error|See error message for details|
