---
title: "Get Allowed Correction Time"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/attendance-v1/user_task_remedy/query_user_allowed_remedys"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/attendance/v1/user_task_remedys/query_user_allowed_remedys"
service: "attendance-v1"
resource: "user_task_remedy"
section: "Attendance"
scopes:
  - "attendance:task:readonly"
updated: "1647419329000"
---

# Obtain the time allowing for user's correction

Obtains the number of punch in/off the user can correct for a given day.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/attendance/v1/user_task_remedys/query_user_allowed_remedys |
| HTTP Method | POST |
| Supported app types | custom |
| Required scopes | `attendance:task:readonly` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" Learn more about obtaining and using access_token. |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `employee_type` | `string` | Yes | Type of employee ID for user_id in request body and in response body **Example value**: "employee_id" **Optional values are**: - `employee_id`: Employee's employee ID, that is the user ID in Lark Admin > Organization > Member and Department > Member Details - `employee_no`: Employee ID, that is the employee ID in Lark Admin > Organization > Member and Department > Member Details | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `user_id` | `string` | Yes | User ID **Example value**: "abd754f7" |
| `remedy_date` | `int` | Yes | Correction date **Example value**: 20210104 | ### Request body example

```json
{
    "user_id": "abd754f7",
    "remedy_date": 20210104
}
```

## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
| ∟ `user_allowed_remedys` | `user_allowed_remedy[]` | Time allowing for user's correction |
| ∟ `user_id` | `string` | User ID |
| ∟ `remedy_date` | `int` | Correction date |
| ∟ `is_free_punch` | `boolean` | Whether it is a free shift or not, if it is a free shift, you can directly choose the correction time, without considering the number of punch in/off |
| ∟ `punch_no` | `int` | The punch in/off No., 0: the first punch in/off, 1: the second punch in/off, 2: the third punch in/off |
| ∟ `work_type` | `int` | On duty/off duty, 1: on duty, 2: off duty |
| ∟ `punch_status` | `string` | Attendance status: Early: early out, Late: late in, Lack: no record |
| ∟ `normal_punch_time` | `string` | Normal attendance time in the format of yyyy-MM-dd HH:mm |
| ∟ `remedy_start_time` | `string` | The minimum value of the optional correction time in the format of yyyy-MM-dd HH:mm |
| ∟ `remedy_end_time` | `string` | The maximum value of the optional correction time in the format of yyyy-MM-dd HH:mm | ### Response body example

```json
{
    "code": 0,
    "msg": "success",
    "data": {
        "user_allowed_remedys": [
            {
                "user_id": "abd754f7",
                "remedy_date": 20210104,
                "is_free_punch": false,
                "punch_no": 0,
                "work_type": 1,
                "punch_status": "Lack",
                "normal_punch_time": "2021-07-01 09:00",
                "remedy_start_time": "2021-07-01 08:00",
                "remedy_end_time": "2021-07-01 10:00"
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
| 500 | 1226500 | Attendance service system error | See error information for details. |
| 400 | 1226501 | No abnormal attendance | There is no abnormal attendance on that day, and no correction is required |
| 400 | 1226502 | Correction is not allowed | Correction not allowed by the attendance group settings |
| 400 | 1226503 | Correction date restriction | It is beyond available correction date as the attendance group settings only allow to make correction for the past ** days |
| 400 | 1226504 | Exceeding correction limits | Correction times for the current period have been used up |
| 500 | 1227500 | Organization service system error | See error information for details. |
