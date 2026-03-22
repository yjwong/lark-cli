---
title: "Query Schedule Information"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/attendance-v1/user_daily_shift/query"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/attendance/v1/user_daily_shifts/query"
service: "attendance-v1"
resource: "user_daily_shift"
section: "Attendance"
scopes:
  - "attendance:task:readonly"
updated: "1647419328000"
---

# Query schedule information

This API is used to query schedule information for multiple users. The query time span can't exceed 30 days.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/attendance/v1/user_daily_shifts/query |
| HTTP Method | POST |
| Supported app types | custom |
| Required scopes | `attendance:task:readonly` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" Learn more about obtaining and using access_token. |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `employee_type` | `string` | Yes | Type of employee ID for user_ids in request body and user_id in response body **Example value**: "employee_id" **Optional values are**: - `employee_id`: Employee's employee ID, that is the user ID in Lark Admin > Organization > Member and Department > Member Details - `employee_no`: Employee ID, that is the employee ID in Lark Admin > Organization > Member and Department > Member Details | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `user_ids` | `string[]` | Yes | employee_no or employee_id list **Example value**: ["abd754f7"] |
| `check_date_from` | `int` | Yes | The start workday of the query **Example value**: 20190817 |
| `check_date_to` | `int` | Yes | The end workday of the query **Example value**: 20190820 | ### Request body example

```json
{
    "user_ids": [
        "abd754f7"
    ],
    "check_date_from": 20190817,
    "check_date_to": 20190820
}
```

## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
| ∟ `user_daily_shifts` | `user_daily_shift[]` | Schedule information list |
| ∟ `group_id` | `string` | Attendance group ID is obtained as follows: 1) Create or modify attendance groups 2) Query attendance group by name 3) Obtain attendance results |
| ∟ `shift_id` | `string` | Shift ID, which can be obtained as follows: 1) Search shift by name 2) Create a shift |
| ∟ `month` | `int` | Month |
| ∟ `user_id` | `string` | User ID |
| ∟ `day_no` | `int` | Date | ### Response body example

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
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 1220001 | Parameter error | Please check if the parameters meet the requirements. |
| 400 | 1220002 | Tenant doesn't exist. | Please check if the tenant_access_token is correct. |
| 400 | 1220004 | User does not exist or has no scope | Check if the user ID is correct |
| 400 | 1220005 | No scope | Please go to [Attendance Admin](https://oa.larksuite.com/attendance/manage/member/list) to check permission scope on data |
| 500 | 1225000 | System error | See error information for details. |
| 500 | 1226000 | Shift service system error | See error message for details. |
| 500 | 1226003 | The shift does not exist | Please check if the shift_id is correct |
