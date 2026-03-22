---
title: "Batch Query the User Settings"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/Attendance//rule/batch-query-user-settings"
section: "Attendance"
updated: "1646322856000"
---

# Query user settings in batches
Queries the user settings of authorized employees in batches, including face photo file ID and face photo last update.
## Request
|Facts||
|---|---|
|HTTP URL|https://open.larksuite.com/open-apis/attendance/v1/user_settings/query|
|HTTP Method|GET|
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
|user_ids|string[]|true|employee_no or employee_id list, which can't exceed 100|
### Request body example
```json
{
    "user_ids": [
        "abd754f7"
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
|  ∟user_settings|user_setting[]|List of user settings|
|    ∟user_id|string|Employee number|
|    ∟face_key|string|Face photo file ID|
|    ∟face_key_update_time|string|Face photo last update, with the timestamp accurate to sec|
### Response body example
```json
{
    "code": 0,
    "msg": "success",
    "data": {
        "user_settings": [
          {
            "face_key": "xxxxxb306842b1c189bc5212eefxxxxx",
            "face_key_update_time": "1625681917",
            "user_id": "abd754f7"
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
|500|1227000|Management service system error|See the error information for details|
|500|1227500|Organizational structure service system error|See error message for details|
