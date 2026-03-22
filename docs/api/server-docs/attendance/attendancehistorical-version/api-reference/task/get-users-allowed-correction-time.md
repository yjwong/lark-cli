---
title: "Get Users’ Allowed Correction Time"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/Attendance//task/query-user-allowed-remedys"
section: "Attendance"
updated: "1646322869000"
---

# Obtain the correction time for users

Obtains the number of clock in/out the user can correction for a given day.

## Request
|Facts||
|---|---|
|HTTP URL|https://open.larksuite.com/open-apis/attendance/v1/user_task_remedys/query_user_allowed_remedys|
|HTTP Method|POST|
|HTTP Content-Type|application/json; charset=utf-8|
|Token requirement|tenant_access_token|
|Required scopes| Attendance data export| ### Header
| key           | value                      |
| ------------- | -------------------------- |
| Authorization | Bearer tenant_access_token |
| Content-Type  | application/json           | ### Query parameters
|Parameter|Type|Required|Description|
|---|---|---|---|
|employee_type|string|Yes|The type of the user_id in the request body, required field, values: [employee_id (the employee's employee ID), employee_no (the employee's employee number)]. Example value: "employee_id"| ### Request body

| Parameter          | Type     | Required | Description      |
| ----------- | ------ | -- | ------- |
| user_id     | string | Yes  | User ID    |
| remedy_date | int    | Yes  | Correction date to query | ### Request body example

```json
{ 
    "user_id" : "abd754f7",
    "remedy_date" : 20210704
}
```

## Response
### Response body

|Parameter|Type|Description|
|---|---|---|
|code|int|Error code. A value other than 0 indicates failure.|
|msg|string|Error description|
|data|-|-|
|  ∟user_allowed_remedys|user_allowed_remedy[]||
|    ∟user_id| string| User ID|
|    ∟remedy_date|int| Correction date|
|    ∟is_free_punch|bool|Whether it is a flexible shift. If yes, you don't need to select the specific clock-in/out record. Simply select the correction time.|
|    ∟punch_no|int|The specific clock-in/out record, values: [0 (1st clock-in/out), 1 (2nd clock-in/out), 2 (3rd clock-in/out)]|
|    ∟work_type|int|Clock-in or clock-out, 1: clock-in, 2: clock-out|
|    ∟punch_status|string| Attendance status, values: [Early (early out), Late (late in), Lack (no record)]|
|    ∟normal_punch_time|string|Normal attendance time, format: yyyy-MM-dd HH:mm|
|    ∟remedy_start_time|string|Minpoint of selectable correction time, format: yyyy-MM-dd HH:mm|
|    ∟remedy_end_time|string|Maxpoint of selectable correction time, format: yyyy-MM-dd HH:mm| ### Response body example

```json
{ 
    "code": 0, 
    "msg": "success",
    "data": {
        "user_allowed_remedy_times" : [
            {
                "user_id" : "abd754f7",
                "remedy_date" : 20210104,
                "is_free_punch" : false,
                "punch_no" : 0,
                "work_type" : 1,
                "punch_status" : "Lack",
                "normal_punch_time" : "2021-07-01 09:00",
                "remedy_start_time" : "2021-07-01 08:00",
                "remedy_end_time" : "2021-07-01 10:00"
            }
        ]
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
