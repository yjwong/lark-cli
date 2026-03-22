---
title: "Notify Approval Status Update"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/Attendance//task/notify-approval-status-update"
section: "Attendance"
updated: "1646322869000"
---

# Notify approval status updates
For Companies that only use the Lark attendance system without the Lark approval system, the approval status of the tripartite system written in the Lark attendance system can be updated through this API, such as approval for leave, overtime, offsite, business trip, and correction. The status includes approved, rejected and revoked.
> **Note:** Only approvals that have been initiated can be updated as approved or rejected, and only approvals that have been approved can be updated as revoked.

## Request

| Facts                |                                                                         |
| ----------------- | ----------------------------------------------------------------------- |
| HTTP URL          |  |
| HTTP Method       | POST                                                                    |
| HTTP Content-Type | application/json; charset=utf-8                                         |
| Required token              | tenant_access_token                                                     |
| Required scopes              | attendance data import                                                                  | ### Header

| key           | value                      |
| ------------- | -------------------------- |
| Authorization | Bearer tenant_access_token |
| Content-Type  | application/json           | ### Request body

| Parameter          | Type     | Required | Description      |
| ------------- | ------ | -- | -------------------------------------------------- |
| approval_id   | string | Yes  | Approval instance ID                                             |
| approval_type | string | Yes  | Approval type, leave: leave, out: offsite, overtime: overtime, trip: business trip, remedy: correction |
| status        | int    | Yes  | Approval status, 1: rejected, 2: approved, 4: revoked                               | ### Request body example

```json
{
        "approval_id" : "6737202939523236113",
        "approval_type" : "remedy",
        "status" : 4
}
```
## Response
### Response body

| Parameter             | Type            | Description                     |
| -------------- | ------------- | ---------------------- |
| code           | int           | Error code. A value other than 0 indicates failure.           |
| msg            | string        | Error description                   |
| data           | -             | -                      |
|  ∟approval_info | approval_info |                        |
|    ∟approval_id   | string        | Approval example ID                 |
|    ∟approval_type | string        | Approval type, leave: leave, out: offsite, overtime: overtime, trip: business trip, remedy: correction         |
|    ∟status        | int           | approval status, 1: rejected, 2: approved, 4: revoked | ### Response body example

```json
{ 

    "code": 0, 
    "msg": "success",
    "data": {
        "approval_info" : {
            "approval_id" : "6737202939523236113",
            "approval_type" : "remedy",
            "status" : 4
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
|500|1225000|System error|See error information for details|
