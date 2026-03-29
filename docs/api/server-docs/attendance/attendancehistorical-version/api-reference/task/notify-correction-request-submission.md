---
title: "Notify Correction Request Submission"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/Attendance//task/notify-remedy-approval-initiation"
section: "Attendance"
updated: "1646322869000"
---

# Notify correction approval initiation
For the companies that only use the Lark attendance system without the Lark approval system, the correction approval data initiated in the tripartite system can be added in the Lark attendance system through this API.
## Request
|Facts||
|---|---|
|HTTP URL|https://open.larksuite.com/open-apis/attendance/v1/user_task_remedys|
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
|employee_type|string|Yes|The type of the user_id in the request body, required field, values: [employee_id (the employee's employee ID), employee_no (the employee's employee number)]. Example value: "employee_id"|
### Request body
|Parameter|Type|Required|Description|
|---|---|---|---|
|user_id|string|Yes|User ID|
|remedy_date|int|Yes|Correction date|
|punch_no|int|Yes|The specific clock-in/out record, values: [0 (1st clock-in/out), 1 (2nd clock-in/out), 2 (3rd clock-in/out)], enter 0 for a flexible shift|
|work_type|int|Yes|Clock-in or clock-out, 1: clock-in, 2: clock-out, enter 0 for a flexible shift|
|remedy_time|string|Yes|Correction time, format: yyyy-MM-dd HH:mm|
|reason|string|Yes|Correction reason|
### Request body example

```json
{ 
    "user_id" : "abd754f7",
    "remedy_date" : 20210701,
    "punch_no" : 0,
    "work_type" : 1,
    "remedy_time" : "2021-07-01 08:00",
    "reason" : ""
}
```
## Response
### Response body

|Parameter|Type|Description|
|---|---|---|
|code|int|Error code. A value other than 0 indicates failure.|
|msg|string|Error description|
|data|-|-|
|  ∟user_remedy|user_task_remedy||
|    ∟user_id| string| User ID|
|    ∟approval_id|string|Approval instance ID, can be used to notify of approval status updates|
|    ∟remedy_date|int| Correction date|
|    ∟punch_no|int|The specific clock-in/out record, values: [0 (1st clock-in/out), 1 (2nd clock-in/out), 2 (3rd clock-in/out)], enter 0 for a flexible shift|
|    ∟work_type|int|Clock-in or clock-out, 1: clock-in, 2: clock-out|
|    ∟remedy_time|string|Correction time, format: yyyy-MM-dd HH:mm|
|    ∟reason|string|Correction reason| ### Response body example

```json
{ 
    "code": 0, 
    "msg": "success",
    "data": {
        "user_remedy" : {
            "user_id" : "abd754f7",
            "approval_id" : "6737202939523236113",
            "remedy_date" : 20210701,
            "punch_no" : 0,
            "work_type" : 1,
            "remedy_time" : "2021-07-01 08:00",
            "reason" : ""
        }
    }
}
```

### Error code

|HTTP status code|Error code|Description|Troubleshooting suggestions|
| -------- | ------- | ---------- | ---------------------------- |
|400|1220001|Parameter error|Please check if the parameters meet the requirements|
| 400      | 1220002 | Tenant doesn't exist      | Please check if the tenant_access_token is correct |
|400|1220004|User doesn't exist or isn't in permission scope|Please check if the user ID is correct|
|400|1220005|No permission scope|Please go to the Attendance Admin to check the data permission scope|
| 400      | 1226501 | No attendance abnormality     | No attendance abnormality on the selected day, no correction is required                |
| 400      | 1226502 | Correction not allowed      | Attendance group settings don't allow correction                   |
| 400      | 1226503 | Correction date restriction     | Attendance group settings don't allow correction after a certain number of days and this date exceeds the limit     |
| 400      | 1226504 | Too many corrections     | No more corrections can be made in the current cycle                  |
|500|1225000|System error|See error information for details|
