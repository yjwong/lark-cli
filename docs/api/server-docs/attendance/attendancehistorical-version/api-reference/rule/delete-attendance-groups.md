---
title: "Delete Attendance Groups"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/Attendance//group_delete"
section: "Attendance"
updated: "1646322860000"
---

# Delete an attendance group
This API is used to delete an attendance group based on its attendance group ID.

## Request
|Facts||
|---|---|
|HTTP URL|https://open.larksuite.com/open-apis/attendance/v1/groups/:group_id|
|HTTP Method|DELETE|
|HTTP Content-Type|application/json; charset=utf-8|
|Token requirement|tenant_access_token|
|Required scopes|Add attendance management rules|
### Path parameters
|Parameter|Type|Required|Description|
|---|---|---|---|
|group_id|string|Yes|Attendance group ID. It needs to be obtained from the API for obtaining attendance results. Example value: "6919358128597097404".|
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
|500|1227000|Management service system error|See the error information for details|
