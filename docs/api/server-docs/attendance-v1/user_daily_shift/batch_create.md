---
title: "Create or modify schedule"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/attendance-v1/user_daily_shift/batch_create"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/attendance/v1/user_daily_shifts/batch_create"
service: "attendance-v1"
resource: "user_daily_shift"
section: "Attendance"
rate_limit: "50 per second"
scopes:
  - "attendance:rule"
updated: "1719396279000"
---

# Create or modify schedule

A schedule describes the shifts that the employees in an attendance group are assigned to work each day. Currently, a schedule can be created for one or more people for a period of one full month.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/attendance/v1/user_daily_shifts/batch_create |
| HTTP Method | POST |
| Rate Limit | 50 per second |
| Supported app types | custom |
| Required scopes | `attendance:rule` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `employee_type` | `string` | Yes | Type of employee ID for user_id in request body and in response body **Example value**: employee_id **Optional values are**:  - `employee_id`: Employee's employee ID, that is the user ID in [Lark Admin](https://example.larksuite.com/admin/index) > Organization > Member and Department > Member Details - `employee_no`: Employee ID, that is the employee ID in [Lark Admin](https://example.larksuite.com/admin/index) > Organization > Member and Department > Member Details  | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `user_daily_shifts` | `user_daily_shift[]` | Yes | Schedule information list（limit 50） |
|   `group_id` | `string` | Yes | Attendance group ID **Example value**: "6737202939523236110" |
|   `shift_id` | `string` | Yes | Shift ID **Example value**: "6753520403404030215" |
|   `month` | `int` | Yes | Month **Example value**: 202101 |
|   `user_id` | `string` | Yes | User ID. The ID type needs to be consistent with the value of employee_type **Example value**: "abd754f7" |
|   `day_no` | `int` | Yes | Date, the number of each month **Example value**: 21 |
| `operator_id` | `string` | No | operator user id. If you didn't finish the "API Integration" at Attendance Admin. This parameter is required. **Example value**: "dd31248a" | ### Request body example

{
    "user_daily_shifts": [
        {
            "group_id": "6737202939523236110",
            "shift_id": "6753520403404030215",
            "month": 202101,
            "user_id": "abd754f7",
            "day_no": 21
        }
    ],
    "operator_id": "dd31248a"
}

## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `user_daily_shifts` | `user_daily_shift[]` | Schedule information list |
|     `group_id` | `string` | Attendance group ID is obtained as follows: 1) Create or modify attendance groups 2) Query attendance group by name 3) Obtain attendance results |
|     `shift_id` | `string` | Shift ID, which can be obtained as follows: 1) Search shift by name 2) Create a shift |
|     `month` | `int` | Month |
|     `user_id` | `string` | User ID |
|     `day_no` | `int` | Date | ### Response body example

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

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 1220001 | Parameter error | Please check if the parameters meet the requirements. |
| 400 | 1220002 | Tenant doesn't exist. | Please check if the tenant_access_token is correct. |
| 400 | 1220004 | User does not exist or has no scope | Check if the user ID is correct |
| 400 | 1220005 | No scope | Please go to [Attendance Admin](https://oa.larksuite.com/attendance/manage/member/list) to check permission scope on data |
| 400 | 1220600 | General error message | See error message for details |
| 400 | 1225000 | System error | See error information for details. |
| 400 | 1226000 | Shift service system error | See error message for details. |
| 400 | 1226003 | The shift does not exist | Please check if the shift_id is correct |
