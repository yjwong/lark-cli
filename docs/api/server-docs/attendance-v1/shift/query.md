---
title: "Query Attendance Shift by Name"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/attendance-v1/shift/query"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/attendance/v1/shifts/query"
service: "attendance-v1"
resource: "shift"
section: "Attendance"
scopes:
  - "attendance:task:readonly"
updated: "1646322852000"
---

# Search shift by name

Searches the shift information by the name of the shift.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/attendance/v1/shifts/query |
| HTTP Method | POST |
| Supported app types | custom |
| Required scopes | `attendance:task:readonly` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" Learn more about obtaining and using access_token. | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `shift_name` | `string` | Yes | Shift name **Example value**: "Morning shift" | ## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `shift` | \- |
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
```

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 1220001 | Parameter error | Please check if the parameters meet the requirements. |
| 400 | 1220002 | Tenant doesn't exist. | Please check if the tenant_access_token is correct. |
| 500 | 1225000 | System error | See error information for details. |
| 500 | 1226000 | Shift service system error | See error message for details. |
| 400 | 1226003 | The shift does not exist | Please check if the shift_id is correct |
