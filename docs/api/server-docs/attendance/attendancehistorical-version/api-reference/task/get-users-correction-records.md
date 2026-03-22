---
title: "Get Users’ Correction Records"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/Attendance//GetUsersRemedyRecords"
section: "Attendance"
updated: "1646322869000"
---

# Obtain correction records
Obtains correction records of authorized employees.

## Request
|Facts||
|---|---|
|HTTP URL|https://open.larksuite.com/open-apis/attendance/v1/user_task_remedys/query|
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
|employee_type|string|Yes|The type of the user_id in the request body, values: [employee_id (the employee's employee ID), employee_no (the employee's employee number)]. Example value: "employee_id"| ### Request body
|Parameter|Type|Required|Description|
|---|---|---|---|
|user_ids|string[]|Yes|employee_no or employee_id list|
|check_time_from|string|Yes|Query start time, with the timestamp accurate to sec|
|check_time_to|string|Yes|Query end time, with the timestamp accurate to sec|
### Request body example
```json
{
    "user_ids": [
        "abd754f7"
    ],
    "check_time_from": "1566641088",
    "check_time_to": "1592561088"
}
```
## Response
### Response body
|Parameter|Type|Description|
|---|---|---|
|code|int|Error code. A value other than 0 indicates failure.|
|msg|string|Error description|
|data|-|-|
|  ∟user_remedys|user_task_remedy[]|List of correction records|
|    ∟user_id|string|Employee number|
|    ∟status|int|Correction status, values: [0 (pending), 2 (approved), 3 (canceled), 4 (revoked after approval)]|
|    ∟reason|string|Correction reason|
|    ∟time|string|Correction time, with the timestamp accurate to sec|
|    ∟time_zone|string|Attendance group time zone at the time of correction|
|    ∟create_time|string|Correction initiation time, with the timestamp accurate to sec|
|    ∟update_time|string|Correction status last update, with the timestamp accurate to sec|
### Response body example
```json
{
    "code": 0,
    "msg": "success",
    "data": {
        "user_remedys": [
            {
                "user_id": "abd754f7",
                "status": 2,
                "reason": "Forgot to clock in/out",
                "time": "1611476284",
                "time_zone": "Asia/Shanghai",
                "create_time": "1611476284",
                "update_time": "1611476284"
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
|500|1226500|Attendance service system error|See error message for details|
