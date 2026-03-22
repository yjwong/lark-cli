---
title: "Create Attendance Shift"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/attendance-v1/shift/create"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/attendance/v1/shifts"
service: "attendance-v1"
resource: "shift"
section: "Attendance"
scopes:
  - "attendance:rule"
updated: "1647419328000"
---

# Create shifts

Shift is a general term used to describe the time rules for an attendance task, such as how many attendances are required for a shift, on and off time of each attendance, how long is the attendance delay being considered as late, and how long is the attendance delay being considered as lack.

> - You must create one or more shifts before creating an attendance group.
> - Shifts within a company are shared, and you can refer directly to the shifts created by others, but it should be noted that any modification of the shift by others will affect your attendance group and attendance results.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/attendance/v1/shifts |
| HTTP Method | POST |
| Supported app types | custom |
| Required scopes | `attendance:rule` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" Learn more about obtaining and using access_token. |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `shift_name` | `string` | Yes | Shift name **Example value**: "Morning shift" |
| `punch_times` | `int` | Yes | Attendances **Example value**: 1 |
| `is_flexible` | `boolean` | No | Whether the attendance rule is flexible **Example value**: false |
| `flexible_minutes` | `int` | No | Flexible attendance time **Example value**: 60 |
| `no_need_off` | `boolean` | No | Punch off is not required **Example value**: true |
| `punch_time_rule` | `punch_time_rule[]` | Yes | Attendance policy |
| ∟ `on_time` | `string` | Yes | On time **Example value**: "9:00" |
| ∟ `off_time` | `string` | Yes | Off time **Example value**: "18:00, 2:00 a.m. the next day, 26:00" |
| ∟ `late_minutes_as_late` | `int` | Yes | How long is the time later than on time being recorded as late **Example value**: 30 |
| ∟ `late_minutes_as_lack` | `int` | Yes | How long is the time later than on time being recorded as lack **Example value**: 60 |
| ∟ `on_advance_minutes` | `int` | Yes | What is the earliest time for punch in **Example value**: 60 |
| ∟ `early_minutes_as_early` | `int` | Yes | How long is the time earlier than off time being recorded as early leave **Example value**: 30 |
| ∟ `early_minutes_as_lack` | `int` | Yes | How long is the time earlier than off time being recorded as lack **Example value**: 60 |
| ∟ `off_delay_minutes` | `int` | Yes | What is the latest time for punch off **Example value**: 60 |
| `late_off_late_on_rule` | `late_off_late_on_rule[]` | No | Late off and late on rules |
| ∟ `late_off_minutes` | `int` | Yes | How long is the off delay **Example value**: 60 |
| ∟ `late_on_minutes` | `int` | Yes | How long is the time later than on time **Example value**: 30 |
| `rest_time_rule` | `rest_rule[]` | No | Break rule |
| ∟ `rest_begin_time` | `string` | Yes | Break starts **Example value**: "13:00" |
| ∟ `rest_end_time` | `string` | Yes | End of break **Example value**: "14:00" | ### Request body example

```json
{
    "shift_name": "Morning shift",
    "punch_times": 1,
    "is_flexible": false,
    "flexible_minutes": 60,
    "no_need_off": true,
    "punch_time_rule": [
        {
            "on_time": "9:00",
            "off_time": "18:00, 2:00 a.m. the next day, 26:00",
            "late_minutes_as_late": 30,
            "late_minutes_as_lack": 60,
            "on_advance_minutes": 60,
            "early_minutes_as_early": 30,
            "early_minutes_as_lack": 60,
            "off_delay_minutes": 60
        }
    ],
    "late_off_late_on_rule": [
        {
            "late_off_minutes": 60,
            "late_on_minutes": 30
        }
    ],
    "rest_time_rule": [
        {
            "rest_begin_time": "13:00",
            "rest_end_time": "14:00"
        }
    ]
}
```

## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
| ∟ `shift` | `shift` | Shift |
| ∟ `shift_id` | `string` | Shift ID |
| ∟ `shift_name` | `string` | Shift name |
| ∟ `punch_times` | `int` | Attendances |
| ∟ `is_flexible` | `boolean` | Whether the attendance rule is flexible |
| ∟ `flexible_minutes` | `int` | Flexible attendance time |
| ∟ `no_need_off` | `boolean` | Punch off is not required |
| ∟ `punch_time_rule` | `punch_time_rule[]` | Attendance policy |
| ∟ `on_time` | `string` | On time |
| ∟ `off_time` | `string` | Off time |
| ∟ `late_minutes_as_late` | `int` | How long is the time later than on time being recorded as late |
| ∟ `late_minutes_as_lack` | `int` | How long is the time later than on time being recorded as lack |
| ∟ `on_advance_minutes` | `int` | What is the earliest time for punch in |
| ∟ `early_minutes_as_early` | `int` | How long is the time earlier than off time being recorded as early leave |
| ∟ `early_minutes_as_lack` | `int` | How long is the time earlier than off time being recorded as lack |
| ∟ `off_delay_minutes` | `int` | What is the latest time for punch off |
| ∟ `late_off_late_on_rule` | `late_off_late_on_rule[]` | Late off and late on rules |
| ∟ `late_off_minutes` | `int` | How long is the off delay |
| ∟ `late_on_minutes` | `int` | How long is the time later than on time |
| ∟ `rest_time_rule` | `rest_rule[]` | Break rule |
| ∟ `rest_begin_time` | `string` | Break starts |
| ∟ `rest_end_time` | `string` | End of break | ### Response body example

```json
{
    "code": 0,
    "msg": "success",
    "data": {
        "shift": {
            "shift_id": "6919358778597097404",
            "shift_name": "Morning shift",
            "punch_times": 1,
            "is_flexible": false,
            "flexible_minutes": 60,
            "no_need_off": true,
            "punch_time_rule": [
                {
                    "on_time": "9:00",
                    "off_time": "18:00, 2:00 a.m. the next day, 26:00",
                    "late_minutes_as_late": 30,
                    "late_minutes_as_lack": 60,
                    "on_advance_minutes": 60,
                    "early_minutes_as_early": 30,
                    "early_minutes_as_lack": 60,
                    "off_delay_minutes": 60
                }
            ],
            "late_off_late_on_rule": [
                {
                    "late_off_minutes": 60,
                    "late_on_minutes": 30
                }
            ],
            "rest_time_rule": [
                {
                    "rest_begin_time": "13:00",
                    "rest_end_time": "14:00"
                }
            ]
        }
    }
}
```

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 1220001 | Parameter error | Please check if the parameters meet the requirements. |
| 400 | 1220002 | Tenant doesn't exist. | Please check if the tenant_access_token is correct. |
| 400 | 1220005 | No scope | Please go to [Attendance Admin](https://oa.larksuite.com/attendance/manage/member/list) to check permission scope on data |
| 500 | 1225000 | System error | See error information for details. |
| 500 | 1226000 | Shift service system error | See error message for details. |
| 400 | 1226001 | The shift has been used | Please change the shift name |
| 400 | 1226002 | Shift name has already been used | Please change the shift name |
