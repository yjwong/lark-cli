---
title: "Get Correction Records"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/attendance-v1/user_task_remedy/query"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/attendance/v1/user_task_remedys/query"
service: "attendance-v1"
resource: "user_task_remedy"
section: "Attendance"
scopes:
  - "attendance:task:readonly"
updated: "1647419329000"
---

# Obtain correction record

Obtains correction records of authorized employees.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/attendance/v1/user_task_remedys/query |
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
| `check_time_from` | `string` | Yes | The start time of the check, with the timestamp accurate to seconds **Example value**: "1566641088" |
| `check_time_to` | `string` | Yes | The end time of the check, with the timestamp accurate to seconds **Example value**: "1592561088" | ### Request body example

```json
{
    "user_ids": [
        "abd754f7"
    ],
    "check_time_from": "1566641088",
    "check_time_to": "1592561088"
}
```

## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
| ∟ `user_remedys` | `user_task_remedy[]` | Correction record list |
| ∟ `user_id` | `string` | User ID |
| ∟ `remedy_date` | `int` | Correction date |
| ∟ `punch_no` | `int` | The punch in/off No., 0: the first punch in/off, 1: the second punch in/off, 2: the third punch in/off, and fill in 0 for free shift |
| ∟ `work_type` | `int` | On duty/off duty, 1: on duty, 2: off duty, and fill in 0 for free shift |
| ∟ `approval_id` | `string` | Approval ID |
| ∟ `remedy_time` | `string` | Correction time in the format of yyyy-MM-dd HH:mm |
| ∟ `status` | `int` | Correction status **Optional values are**: - `0`: Under review - `2`: Approved - `3`: Canceled - `4`: Revoked after approval |
| ∟ `reason` | `string` | Correction reason |
| ∟ `time` | `string` | Correction time, with the timestamp accurate to seconds |
| ∟ `time_zone` | `string` | Time zone of the attendance group at the time of correction |
| ∟ `create_time` | `string` | Correction initiation time, with the timestamp accurate to seconds |
| ∟ `update_time` | `string` | Time of correction status updates, with the timestamp accurate to seconds | ### Response body example

```json
{
    "code": 0,
    "msg": "success",
    "data": {
        "user_remedys": [
            {
                "user_id": "abd754f7",
                "remedy_date": 20210701,
                "punch_no": 0,
                "work_type": 1,
                "approval_id": "6737202939523236113",
                "remedy_time": "2021-07-01 08:00",
                "status": 2,
                "reason": "Forget about attendance",
                "time": "1611476284",
                "time_zone": "Asia/Shanghai",
                "create_time": "1611476284",
                "update_time": "1611476284"
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
