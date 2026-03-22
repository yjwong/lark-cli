---
title: "Delete a Shift"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/Attendance//shift_delete"
section: "Attendance"
updated: "1646322860000"
---

# Delete a shift
This API is used to delete a shift based on its shift ID.

## Request
|Facts||
|---|---|
|HTTP URL|https://open.larksuite.com/open-apis/attendance/v1/shifts/:shift_id|
|HTTP Method|DELETE|
|HTTP Content-Type|application/json; charset=utf-8|
|Token requirement|tenant_access_token|
|Required scopes|Add attendance management rules|
### Header
key|value
--|--
Authorization|Bearer tenant_access_token
Content-Type|application/json
### Path parameters
|Parameter|Type|Required|Description|
|---|---|---|---|
|shift_id|string|Yes|Shift ID. Example value: "6919358778597097404"|
## Response
### Response body
|Parameter|Type|Description|
|---|---|---|
|code|int|Error code. A value other than 0 indicates failure.|
|msg|string|Error description|
### Response body example
```json
{
    "code": 0,
    "msg": "success"
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
|400|1226003|Shift doesn't exist|Check if the shift_id is correct|
