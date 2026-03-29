---
title: "Create or modify an Attendance Group"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/attendance-v1/group/create"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/attendance/v1/groups"
service: "attendance-v1"
resource: "group"
section: "Attendance"
scopes:
  - "attendance:rule"
updated: "1698225706000"
---

# Create or modify attendance groups

Attendance groups are used to set attendance policy for a department or employee in a specific place and within a specific time period (including clock-in/out, late in, early out, sick leave, marital leave, funeral leave, public holidays, business hours, and overtime).

You can use attendance groups to set attendance rules such as attendance methods, attendance time, and attendance locations for departments or specific employees.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/attendance/v1/groups |
| HTTP Method | POST |
| Supported app types | custom |
| Required scopes | `attendance:rule` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `employee_type` | `string` | Yes | User ID type **Example value**: "employee_id" **Optional values are**:  - `employee_id`: Employee ID, that is the user ID in [Lark Admin](https://n17wxurp1s.larksuite.com/admin/index) > Organization > Member and Department > Member Details - `employee_no`: Employee number, that is the employee ID in [Lark Admin](https://n17wxurp1s.larksuite.com/admin/index) > Organization > Member and Department > Member Details  |
| `dept_type` | `string` | Yes | Department ID type **Example value**: "open_id" **Optional values are**:  - `open_id`: Only Open ID of the department is supported  | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `group` | `group` | Yes | 6921319402260496386 |
|   `group_id` | `string` | No | Attendance group ID (only available when modified). You need to obtain the groupId from the API of "Obtain Attendance Results" **Example value**: "6919358128597097404" |
|   `group_name` | `string` | Yes | Attendance group name **Example value**: "开心考勤" |
|   `time_zone` | `string` | Yes | Time zone **Example value**: "Asia/Shanghai" |
|   `bind_dept_ids` | `string[]` | No | IDs of connected departments **Example value**: od-fcb45c28a45311afd440b7869541fce8 |
|   `except_dept_ids` | `string[]` | No | IDs of excepted departments **Example value**: od-fcb45c28a45311afd440b7869541fce8 |
|   `bind_user_ids` | `string[]` | No | IDs of connected users **Example value**: 52aa1fa1 |
|   `except_user_ids` | `string[]` | No | IDs of excepted users **Example value**: 52aa1fa1 |
|   `group_leader_ids` | `string[]` | Yes | Attendance owner ID list (required) **Example value**: 2bg4a9be |
|   `sub_group_leader_ids` | `string[]` | No | Attendance sub owner ID list **Example value**: 4fasdtc2 |
|   `allow_out_punch` | `boolean` | No | Whether to allow offsite attendance **Example value**: true |
|   `out_punch_need_approval` | `boolean` | No | Whether offsite attendance need approval **Example value**: true |
|   `out_punch_need_remark` | `boolean` | No | Whether offsite attendance need remark **Example value**: true |
|   `out_punch_need_photo` | `boolean` | No | Whether offsite attendance need upload photo **Example value**: true |
|   `out_punch_allowed_hide_addr` | `boolean` | No | Whether offsite attendance should hide address **Example value**: true |
|   `allow_pc_punch` | `boolean` | No | Whether to allow attendance on PC **Example value**: true |
|   `allow_remedy` | `boolean` | No | Whether to limit correction **Example value**: true |
|   `remedy_limit` | `boolean` | No | Whether to limit the correction count **Example value**: true |
|   `remedy_limit_count` | `int` | No | Correction count **Example value**: 3 |
|   `remedy_date_limit` | `boolean` | No | Whether to limit the days allowed to make corrections **Example value**: true |
|   `remedy_date_num` | `int` | No | Days allowed to make corrections **Example value**: 3 |
|   `allow_remedy_type_lack` | `boolean` | No | Whether to make corrections allowed lack **Example value**: true |
|   `allow_remedy_type_late` | `boolean` | No | Whether to make corrections allowed late **Example value**: true |
|   `allow_remedy_type_early` | `boolean` | No | Whether to make corrections allowed early **Example value**: true |
|   `allow_remedy_type_normal` | `boolean` | No | Whether to make corrections allowed normal **Example value**: true |
|   `show_cumulative_time` | `boolean` | No | Whether to display cumulative duration **Example value**: true |
|   `show_over_time` | `boolean` | No | Whether to display overtime duration **Example value**: true |
|   `hide_staff_punch_time` | `boolean` | No | Whether to hide employee attendance details **Example value**: true |
|   `face_punch` | `boolean` | No | Whether to enable face recognition attendance **Example value**: true |
|   `face_punch_cfg` | `int` | No | Face recognition attendance policy **Optional values are:** * 1: Face recognition required for each clock-in/out * 2: Face recognition only required when cheating is suspected **Example value**: 1 |
|   `face_downgrade` | `boolean` | No | Whether to allow normal photo attendance when face recognition fails **Example value**: true |
|   `replace_basic_pic` | `boolean` | No | Whether to allow baseline image replacement when face recognition fails **Example value**: true |
|   `machines` | `machine[]` | No | Attendance machine list |
|     `machine_sn` | `string` | Yes | Attendance machine serial number **Example value**: "FS0701" |
|     `machine_name` | `string` | Yes | Attendance machine name **Example value**: "创实 9 楼" |
|   `gps_range` | `int` | No | Valid range for GPS attendance (not recommended) **Example value**: 300 |
|   `locations` | `location[]` | No | Address list |
|     `location_name` | `string` | Yes | Address name **Example value**: "Xixi Bafang City, Muqiaotou, Wuchang Street, Yuhang District, Hangzhou, Zhejiang Province" |
|     `location_type` | `int` | Yes | Address type **Optional values are:** * 1: GPS * 2: Wi-Fi * 8: IP **Example value**: 1 |
|     `latitude` | `float` | No | Address latitude **Example value**: 39.971827 |
|     `longitude` | `float` | No | Address longitude **Example value**: 116.32989 |
|     `ssid` | `string` | No | Wi-Fi name **Example value**: "TP-Link-af12ca" |
|     `bssid` | `string` | No | Wi-Fi MAC address **Example value**: "08:00:20:0A:8C:6D" |
|     `map_type` | `int` | No | Map type, 1: AMap, 2: Google **Example value**: 1 |
|     `address` | `string` | No | Address name **Example value**: "AVIC Plaza, Haidian District, Beijing" |
|     `ip` | `string` | No | IP address **Example value**: "122.224.123.146" |
|     `feature` | `string` | No | Additional information, such as operator information **Example value**: "China Telecom" |
|     `gps_range` | `int` | No | Valid range for GPS attendance **Example value**: 300 |
|   `group_type` | `int` | Yes | Attendance type, 0: fixed shift, 2: scheduled shift, 3: flexible shift **Example value**: 0 |
|   `punch_day_shift_ids` | `string[]` | Yes | Required for fixed shift **Example value**: 6921319402260496386 |
|   `free_punch_cfg` | `free_punch_cfg` | No | Flexible shift configuration |
|     `free_start_time` | `string` | Yes | Attendance start time under the flexible shift **Example value**: "7:00" |
|     `free_end_time` | `string` | Yes | Attendance end time under the Flexible shift **Example value**: "18:00" |
|     `punch_day` | `int` | Yes | Attendance time, 7 digits representing each day of the week from Monday to Sunday. 0 indicates no shift for the day and 1 indicates a shift for the day. **Example value**: 1111100 |
|     `work_day_no_punch_as_lack` | `boolean` | No | Whether to mark as no record if they fail to clock in/out on a working day **Example value**: true |
|   `calendar_id` | `int` | Yes | National calendar ID, 0: schedule not based on the national calendar, 1: Chinese mainland (default), 2: US, 3: Japan, 4: India, 5: Singapore **Example value**: 1 |
|   `need_punch_special_days` | `punch_special_date_shift[]` | No | Special dates requiring attendance |
|     `punch_day` | `int` | Yes | Attendance date **Example value**: 20190101 |
|     `shift_id` | `string` | Yes | Shift ID **Example value**: "6919668827865513935" |
|   `no_need_punch_special_days` | `punch_special_date_shift[]` | No | Special dates not requiring attendance |
|     `punch_day` | `int` | Yes | Attendance date **Example value**: 20190101 |
|     `shift_id` | `string` | Yes | Shift ID **Example value**: "6919668827865513935" |
|   `work_day_no_punch_as_lack` | `boolean` | No | Whether to mark as no record if they fail to clock in/out on a working day under the flexible shift **Example value**: true |
|   `effect_now` | `boolean` | No | Whether the change takes effect immediately, default value: false **Example value**: true |
|   `remedy_period_type` | `int` | No | Correction cycle type **Example value**: 0 |
|   `remedy_period_custom_date` | `int` | No | Start date of custom correction cycle **Example value**: 1 |
|   `punch_type` | `int` | No | Attendance type, bit operation. 1: GPS attendance, 2: Wi-Fi attendance, 4: attendance on the attendance machine, 8: IP-based attendance **Example value**: 1 |
|   `rest_clockIn_need_approval` | `boolean` | No | Whether to clocckin at rest need approval. When `rest_clockIn_need_approval=true` is set, the rest day start time is reset to 4:00. **Example value**: true |
|   `clockIn_need_photo` | `boolean` | No | Whether to clocckin at rest need upload photo **Example value**: true |
| `operator_id` | `string` | No | operator user id. If you didn't finish the "API Integration" at Attendance Admin. This parameter is required. **Example value**: "dd31248a" | ### Request body example

{
    "group": {
        "group_id": "6919358128597097404",
        "group_name": "开心考勤",
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
        "sub_group_leader_ids": [
            "4fasdtc2"
        ],
        "allow_out_punch": true,
        "out_punch_need_approval": true,
        "out_punch_need_remark": true,
        "out_punch_need_photo": true,
        "out_punch_allowed_hide_addr": true,
        "allow_pc_punch": true,
        "allow_remedy": true,
        "remedy_limit": true,
        "remedy_limit_count": 3,
        "remedy_date_limit": true,
        "remedy_date_num": 3,
        "allow_remedy_type_lack": true,
        "allow_remedy_type_late": true,
        "allow_remedy_type_early": true,
        "allow_remedy_type_normal": true,
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
                "machine_name": "创实 9 楼"
            }
        ],
        "gps_range": 300,
        "locations": [
            {
                "location_name": "Xixi Bafang City, Muqiaotou, Wuchang Street, Yuhang District, Hangzhou, Zhejiang Province",
                "location_type": 1,
                "latitude": 39.971827,
                "longitude": 116.32989,
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
        "punch_type": 1,
        "rest_clockIn_need_approval": true,
        "clockIn_need_photo": true
    },
    "operator_id": "dd31248a"
}

## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `group` | `group` | 6921319402260496386 |
|     `group_id` | `string` | Attendance group ID (only available when modified). You need to obtain the groupId from the API of "Obtain Attendance Results" |
|     `group_name` | `string` | Attendance group name |
|     `time_zone` | `string` | Time zone |
|     `bind_dept_ids` | `string[]` | IDs of connected departments |
|     `except_dept_ids` | `string[]` | IDs of excepted departments |
|     `bind_user_ids` | `string[]` | IDs of connected users |
|     `except_user_ids` | `string[]` | IDs of excepted users |
|     `group_leader_ids` | `string[]` | Attendance owner ID list (required) |
|     `sub_group_leader_ids` | `string[]` | Attendance sub owner ID list |
|     `allow_out_punch` | `boolean` | Whether to allow offsite attendance |
|     `out_punch_need_approval` | `boolean` | Whether offsite attendance need approval |
|     `out_punch_need_remark` | `boolean` | Whether offsite attendance need remark |
|     `out_punch_need_photo` | `boolean` | Whether offsite attendance need upload photo |
|     `out_punch_allowed_hide_addr` | `boolean` | Whether offsite attendance should hide address |
|     `allow_pc_punch` | `boolean` | Whether to allow attendance on PC |
|     `allow_remedy` | `boolean` | Whether to limit correction |
|     `remedy_limit` | `boolean` | Whether to limit the correction count |
|     `remedy_limit_count` | `int` | Correction count |
|     `remedy_date_limit` | `boolean` | Whether to limit the days allowed to make corrections |
|     `remedy_date_num` | `int` | Days allowed to make corrections |
|     `allow_remedy_type_lack` | `boolean` | Whether to make corrections allowed lack |
|     `allow_remedy_type_late` | `boolean` | Whether to make corrections allowed late |
|     `allow_remedy_type_early` | `boolean` | Whether to make corrections allowed early |
|     `allow_remedy_type_normal` | `boolean` | Whether to make corrections allowed normal |
|     `show_cumulative_time` | `boolean` | Whether to display cumulative duration |
|     `show_over_time` | `boolean` | Whether to display overtime duration |
|     `hide_staff_punch_time` | `boolean` | Whether to hide employee attendance details |
|     `face_punch` | `boolean` | Whether to enable face recognition attendance |
|     `face_punch_cfg` | `int` | Face recognition attendance policy **Optional values are:** * 1: Face recognition required for each clock-in/out * 2: Face recognition only required when cheating is suspected |
|     `face_downgrade` | `boolean` | Whether to allow normal photo attendance when face recognition fails |
|     `replace_basic_pic` | `boolean` | Whether to allow baseline image replacement when face recognition fails |
|     `machines` | `machine[]` | Attendance machine list |
|       `machine_sn` | `string` | Attendance machine serial number |
|       `machine_name` | `string` | Attendance machine name |
|     `gps_range` | `int` | Valid range for GPS attendance (not recommended) |
|     `locations` | `location[]` | Address list |
|       `location_id` | `string` | Address ID |
|       `location_name` | `string` | Address name |
|       `location_type` | `int` | Address type **Optional values are:** * 1: GPS * 2: Wi-Fi * 8: IP |
|       `latitude` | `float` | Address latitude |
|       `longitude` | `float` | Address longitude |
|       `ssid` | `string` | Wi-Fi name |
|       `bssid` | `string` | Wi-Fi MAC address |
|       `map_type` | `int` | Map type, 1: AMap, 2: Google |
|       `address` | `string` | Address name |
|       `ip` | `string` | IP address |
|       `feature` | `string` | Additional information, such as operator information |
|       `gps_range` | `int` | Valid range for GPS attendance |
|     `group_type` | `int` | Attendance type, 0: fixed shift, 2: scheduled shift, 3: flexible shift |
|     `punch_day_shift_ids` | `string[]` | Required for fixed shift |
|     `free_punch_cfg` | `free_punch_cfg` | Flexible shift configuration |
|       `free_start_time` | `string` | Attendance start time under the flexible shift |
|       `free_end_time` | `string` | Attendance end time under the Flexible shift |
|       `punch_day` | `int` | Attendance time, 7 digits representing each day of the week from Monday to Sunday. 0 indicates no shift for the day and 1 indicates a shift for the day. |
|       `work_day_no_punch_as_lack` | `boolean` | Whether to mark as no record if they fail to clock in/out on a working day |
|     `calendar_id` | `int` | National calendar ID, 0: schedule not based on the national calendar, 1: Chinese mainland (default), 2: US, 3: Japan, 4: India, 5: Singapore |
|     `need_punch_special_days` | `punch_special_date_shift[]` | Special dates requiring attendance |
|       `punch_day` | `int` | Attendance date |
|       `shift_id` | `string` | Shift ID |
|     `no_need_punch_special_days` | `punch_special_date_shift[]` | Special dates not requiring attendance |
|       `punch_day` | `int` | Attendance date |
|       `shift_id` | `string` | Shift ID |
|     `work_day_no_punch_as_lack` | `boolean` | Whether to mark as no record if they fail to clock in/out on a working day under the flexible shift |
|     `effect_now` | `boolean` | Whether the change takes effect immediately, default value: false |
|     `remedy_period_type` | `int` | Correction cycle type |
|     `remedy_period_custom_date` | `int` | Start date of custom correction cycle |
|     `punch_type` | `int` | Attendance type, bit operation. 1: GPS attendance, 2: Wi-Fi attendance, 4: attendance on the attendance machine, 8: IP-based attendance |
|     `effect_time` | `string` | Effective time, timestamp accurate to seconds |
|     `rest_clockIn_need_approval` | `boolean` | Whether to clocckin at rest need approval. When `rest_clockIn_need_approval=true` is set, the rest day start time is reset to 4:00. |
|     `clockIn_need_photo` | `boolean` | Whether to clocckin at rest need upload photo | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
        "group": {
            "group_id": "6919358128597097404",
            "group_name": "开心考勤",
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
            "sub_group_leader_ids": [
                "4fasdtc2"
            ],
            "allow_out_punch": true,
            "out_punch_need_approval": true,
            "out_punch_need_remark": true,
            "out_punch_need_photo": true,
            "out_punch_allowed_hide_addr": true,
            "allow_pc_punch": true,
            "allow_remedy": true,
            "remedy_limit": true,
            "remedy_limit_count": 3,
            "remedy_date_limit": true,
            "remedy_date_num": 3,
            "allow_remedy_type_lack": true,
            "allow_remedy_type_late": true,
            "allow_remedy_type_early": true,
            "allow_remedy_type_normal": true,
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
                    "machine_name": "创实 9 楼"
                }
            ],
            "gps_range": 300,
            "locations": [
                {
                    "location_id": "6921213751454744578",
                    "location_name": "Xixi Bafang City, Muqiaotou, Wuchang Street, Yuhang District, Hangzhou, Zhejiang Province",
                    "location_type": 1,
                    "latitude": 39.971827,
                    "longitude": 116.32989,
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
            "punch_type": 1,
            "effect_time": "1611476284",
            "rest_clockIn_need_approval": true,
            "clockIn_need_photo": true
        }
    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 1220001 | Parameter error | Please check if the parameters meet the requirements. |
| 400 | 1220002 | Tenant doesn't exist. | Please check if the tenant_access_token is correct. |
| 400 | 1220003 | employee_type does not exist. | employee_type, employee_id: Employee ID, employee_no: Employee number |
| 400 | 1220004 | User does not exist or has no scope | Check if the user ID is correct |
| 400 | 1220005 | No scope | Please go to [Attendance Admin](https://oa.larksuite.com/attendance/manage/member/list) to check permission scope on data |
| 500 | 1225000 | System error | See error information for details. |
| 400 | 1227500 | Organization service system error | See error information for details. |
| 400 | 1227000 | Management service system error | See error information for details. |
