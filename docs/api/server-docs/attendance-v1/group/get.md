---
title: "Get Attendance Group Details"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/attendance-v1/group/get"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/attendance/v1/groups/:group_id"
service: "attendance-v1"
resource: "group"
section: "Attendance"
scopes:
  - "attendance:rule"
  - "attendance:rule:readonly"
updated: "1646322852000"
---

# Obtain attendance group details

Obtains attendance group details via an attendance group ID.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/attendance/v1/groups/:group_id |
| HTTP Method | GET |
| Supported app types | custom |
| Required scopes Enable any scope from the list | `attendance:rule` `attendance:rule:readonly` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" Learn more about obtaining and using access_token. | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `group_id` | `string` | Attendance group ID is obtained as follows: 1) Create or modify attendance groups 2) Query attendance group by name 3) Obtain attendance results **Example value**: "6919358128597097404" | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `employee_type` | `string` | Yes | User ID type **Example value**: "employee_id" **Optional values are**: - `employee_id`: Employee ID - `employee_no`: Employee ID |
| `dept_type` | `string` | Yes | Department ID type **Example value**: "od-fcb45c28a45311afd441b8869541ece8" **Optional values are**: - `open_id`: Only Open ID of the department is supported | ## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `group` | \- |
| ∟ `group_id` | `string` | Attendance group ID (only available when modified). You need to obtain the groupId from the API of "Obtain Attendance Results" |
| ∟ `group_name` | `string` | Attendance group name |
| ∟ `time_zone` | `string` | Time zone |
| ∟ `bind_dept_ids` | `string[]` | IDs of connected departments |
| ∟ `except_dept_ids` | `string[]` | IDs of excepted departments |
| ∟ `bind_user_ids` | `string[]` | IDs of connected users |
| ∟ `except_user_ids` | `string[]` | IDs of excepted users |
| ∟ `group_leader_ids` | `string[]` | Attendance owner ID list (required) |
| ∟ `allow_out_punch` | `boolean` | Whether to allow offsite attendance |
| ∟ `allow_pc_punch` | `boolean` | Whether to allow attendance on PC |
| ∟ `allow_remedy` | `boolean` | Whether to limit correction |
| ∟ `remedy_limit` | `boolean` | Whether to limit the correction count |
| ∟ `remedy_limit_count` | `int` | Correction count |
| ∟ `remedy_date_limit` | `boolean` | Whether to limit the days allowed to make corrections |
| ∟ `remedy_date_num` | `int` | Days allowed to make corrections |
| ∟ `show_cumulative_time` | `boolean` | Whether to display cumulative duration |
| ∟ `show_over_time` | `boolean` | Whether to display overtime duration |
| ∟ `hide_staff_punch_time` | `boolean` | Whether to hide employee attendance details |
| ∟ `face_punch` | `boolean` | Whether to enable face recognition attendance |
| ∟ `face_punch_cfg` | `int` | Face recognition attendance policy, 1: Face recognition required for each clock-in/out, 2: Face recognition only required when cheating is suspected |
| ∟ `face_downgrade` | `boolean` | Whether to allow normal photo attendance when face recognition fails |
| ∟ `replace_basic_pic` | `boolean` | Whether to allow baseline image replacement when face recognition fails |
| ∟ `machines` | `machine[]` | Attendance machine list |
| ∟ `machine_sn` | `string` | Attendance machine serial number |
| ∟ `machine_name` | `string` | Attendance machine name |
| ∟ `gps_range` | `int` | Valid range for GPS attendance (not recommended) |
| ∟ `locations` | `location[]` | Address list |
| ∟ `location_id` | `string` | Address ID |
| ∟ `location_name` | `string` | Address name |
| ∟ `location_type` | `int` | Address type, 1: GPS, 2: Wi-Fi, 8: IP |
| ∟ `latitude` | `float` | Address latitude |
| ∟ `longitude` | `float` | Address longitude |
| ∟ `ssid` | `string` | Wi-Fi name |
| ∟ `bssid` | `string` | Wi-Fi MAC address |
| ∟ `map_type` | `int` | Map type, 1: AMap, 2: Google |
| ∟ `address` | `string` | Address name |
| ∟ `ip` | `string` | IP address |
| ∟ `feature` | `string` | Additional information, such as operator information |
| ∟ `gps_range` | `int` | Valid range for GPS attendance |
| ∟ `group_type` | `int` | Attendance type, 0: fixed shift, 2: scheduled shift, 3: flexible shift |
| ∟ `punch_day_shift_ids` | `string[]` | Required for fixed shift |
| ∟ `free_punch_cfg` | `free_punch_cfg` | Flexible shift configuration |
| ∟ `free_start_time` | `string` | Attendance start time under the flexible shift |
| ∟ `free_end_time` | `string` | Attendance end time under the Flexible shift |
| ∟ `punch_day` | `int` | Attendance time, 7 digits representing each day of the week from Monday to Sunday. 0 indicates no shift for the day and 1 indicates a shift for the day. |
| ∟ `work_day_no_punch_as_lack` | `boolean` | Whether to mark as no record if they fail to clock in/out on a working day |
| ∟ `calendar_id` | `int` | National calendar ID, 0: schedule not based on the national calendar, 1: Chinese mainland (default), 2: US, 3: Japan, 4: India, 5: Singapore |
| ∟ `need_punch_special_days` | `punch_special_date_shift[]` | Special dates requiring attendance |
| ∟ `punch_day` | `int` | Attendance date |
| ∟ `shift_id` | `string` | Shift ID |
| ∟ `no_need_punch_special_days` | `punch_special_date_shift[]` | Special dates not requiring attendance |
| ∟ `punch_day` | `int` | Attendance date |
| ∟ `shift_id` | `string` | Shift ID |
| ∟ `work_day_no_punch_as_lack` | `boolean` | Whether to mark as no record if they fail to clock in/out on a working day under the flexible shift |
| ∟ `effect_now` | `boolean` | Whether the change takes effect immediately, default value: false |
| ∟ `remedy_period_type` | `int` | Correction cycle type |
| ∟ `remedy_period_custom_date` | `int` | Start date of custom correction cycle |
| ∟ `punch_type` | `int` | Attendance type, bit operation. 1: GPS attendance, 2: Wi-Fi attendance, 4: attendance on the attendance machine, 8: IP-based attendance | ### Response body example

```json
{
    "code": 0,
    "msg": "success",
    "data": {
        "group_id": "6919358128597097404",
        "group_name": "Kaixin Attendance",
        "time_zone": "Asia/Shanghai",
        "bind_dept_ids": [
            "a6a64f94gg768492"
        ],
        "except_dept_ids": [
            "f864f94dg73459845"
        ],
        "bind_user_ids": [
            "52aa1fa1"
        ],
        "except_user_ids": [
            "4fasdtc2"
        ],
        "group_leader_ids": [
            "52aa1fa1"
        ],
        "allow_out_punch": true,
        "allow_pc_punch": true,
        "allow_remedy": true,
        "remedy_limit": true,
        "remedy_limit_count": 3,
        "remedy_date_limit": true,
        "remedy_date_num": 3,
        "show_cumulative_time": true,
        "show_over_time": true,
        "hide_staff_punch_time": true,
        "face_punch": true,
        "face_punch_cfg": 1,
        "face_downgrade": true,
        "replace_basic_pic": true,
        "machines": [
            {
                "machine_sn": "FS0701",
                "machine_name": "Chuangshi 9th Floor"
            }
        ],
        "gps_range": 300,
        "locations": [
            {
                "location_id": "6921213751454744578",
                "location_name": "Xixi Bafang City, Muqiaotou, Wuchang Street, Yuhang District, Hangzhou, Zhejiang Province",
                "location_type": 1,
                "latitude": 30.28994,
                "longitude": 120.04509,
                "ssid": "TP-Link-af12ca",
                "bssid": "08:00:20:0A:8C:6D",
                "map_type": 1,
                "address": "AVIC Plaza, Haidian District, Beijing",
                "ip": "122.224.123.146",
                "feature": "China Telecom",
                "gps_range": 300
            }
        ],
        "group_type": 0,
        "punch_day_shift_ids": [
            "6919668824125513935"
        ],
        "free_punch_cfg": {
            "free_start_time": "7:00",
            "free_end_time": "18:00",
            "punch_day": 1111100,
            "work_day_no_punch_as_lack": true
        },
        "calendar_id": 1,
        "need_punch_special_days": [
            {
                "punch_day": 20190101,
                "shift_id": "6919668827865513935"
            }
        ],
        "no_need_punch_special_days": [
            {
                "punch_day": 20190101,
                "shift_id": "6919668827865513935"
            }
        ],
        "work_day_no_punch_as_lack": true,
        "effect_now": true,
        "remedy_period_type": 0,
        "remedy_period_custom_date": 1,
        "punch_type": 1
    }
}
```

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 1220001 | Parameter error | Please check if the parameters meet the requirements. |
| 400 | 1220002 | Tenant doesn't exist. | Please check if the tenant_access_token is correct. |
| 400 | 1220003 | employee_type does not exist. | employee_type, employee_id: Employee ID, employee_no: Employee number |
| 500 | 1225000 | System error | See error information for details. |
| 500 | 1227000 | Management service system error | See error information for details. |
