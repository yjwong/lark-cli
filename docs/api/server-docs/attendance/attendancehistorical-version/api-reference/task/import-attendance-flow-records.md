---
title: "Import Attendance Flow Records"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/Attendance//ImportAttendanceFlowRecords"
section: "Attendance"
updated: "1646322869000"
---

# Import attendance records
Imports the attendance records of authorized employees. After the import, the final attendance status and result can be calculated according to the shift rules of the employee's attendance group.

> **Note:** Applicable to the attendance machine data import and other scenes.
## Request
|Facts||
|---|---|
|HTTP URL|https://open.larksuite.com/open-apis/attendance/v1/user_flows/batch_create|
|HTTP Method|POST|
|HTTP Content-Type|application/json; charset=utf-8|
|Token requirement|tenant_access_token|
|Required scopes|Add attendance data|
### Header
key|value
--|--
Authorization|Bearer tenant_access_token
Content-Type|application/json
### Query parameters
|Parameter|Type|Required|Description|
|---|---|---|---|
|employee_type|string|Yes|The type of the user_id and creator_id in the request body, values: [employee_id (the employee's employee ID), employee_no (the employee's employee number)]. Example value: "employee_id"|
### Request body
|Parameter|Type|Required|Description|
|---|---|---|---|
|flow_records|user_flow[]|Yes|List of attendance records|
|  ∟user_id|string|Yes|Employee number|
|  ∟creator_id|string|Yes|Employee number of attendance record creator|
|  ∟location_name|string|Yes|Attendance location name|
|  ∟check_time|string|Yes|Attendance time, with the timestamp accurate to sec|
|  ∟comment|string|Yes|Attendance notes|
### Request body example
```json
{
    "flow_records": [
        {
            "user_id": "abd754f7",
            "creator_id": "abd754f7",
            "location_name": "Xixi Bafang City",
            "check_time": "1611476284",
            "comment": "System auto attendance"
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
|  ∟flow_records|user_flow[]|List of attendance records|
|    ∟user_id|string|Employee number|
|    ∟creator_id|string|employee_no of attendance record creator|
|    ∟location_name|string|Attendance location name|
|    ∟check_time|string|Attendance time, with the timestamp accurate to sec|
|    ∟comment|string|Attendance notes|
### Response body example
```json
{
    "code": 0,
    "msg": "success",
    "data": {
        "flow_records": [
            {
            "user_id": "abd754f7",
            "creator_id": "abd754f7",
            "location_name": "Xixi Bafang City",
            "check_time": "1611476284",
            "comment": "System auto attendance"
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
