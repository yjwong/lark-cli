---
title: "Create or Modify Attendance Groups"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/Attendance//group_create_update"
section: "Attendance"
updated: "1646322860000"
---

# Create or modify attendance groups
Attendance groups are used to set attendance policy for a department or employee in a specific place and within a specific time period (including clock-in/out, late in, early out, sick leave, marital leave, funeral leave, public holidays, business hours, and overtime).

You can use attendance groups to set attendance policy such as attendance methods, attendance times, and attendance locations for departments or specific employees.
> **Note:** For security reasons, you can only use this API to modify attendance groups created by yourself.
### Attendance group owner
The attendance group owner can modify the schedule of the attendance group and view the attendance data of the attendance group.

If the attendance group owner is assigned the role of attendance administrator by the company administrator, the owner also has the permission scope of the attendance administrator to edit and delete attendance rules.

### Attendance group member
You can set members who need to attend attendance or who don't need to attend attendance by two dimensions: department and employee.
- If a member of the attendance group is added according to the department dimension, when a new employee joins the department, they will automatically join the attendance group.
- If a member of the attendance group is added according to the employee dimension, the employee will not change the attendance group when their superior department is added to another attendance group.

### Attendance group type
There are 3 attendance types: fixed shift, scheduled shift, and flexible shift.
- Fixed shift: The working time for each member in the attendance group is same. It is suitable for attendance groups with fixed working time or without multiple shifts.
- Scheduled shift: The working time for members in the attendance group are not exactly the same, and can be customized. It is suitable for attendance groups with multiple shifts, such as morning shift and evening shifts.
- Flexible shift: There is no specific shift, and the attendance group members can clock in/out freely during the attendance period, and the working hours are counted according to the attendance period.

### Attendance shift
- Under the fixed shift, you need to arrange shifts from Monday to Sunday and for special days.
- Under the scheduled shift, you need to specify the shift for each member in the attendance group every day.
- Under the flexible shift, you need to set the earliest and the latest time of the day for attendance, and the days of the week when attendance is required.

### Attendance method
There are 3 available attendance methods: GPS attendance, Wi-Fi attendance and attendance on the attendance machine.
- GPS attendance: It requires latitude and longitude information and attendance location name.
- Wi-Fi attendance: It requires Wi-Fi name and Wi-Fi MAC address.
- Attendance on the attendance machine: It requires the name and serial number of the attendance machine.

### Other attendance settings
- Rule settings: Set whether to allow offsite attendance, whether to allow correction and the correction count allowed in one month, and whether to allow attendance on PC.
- Security settings: Set whether to enable face recognition attendance, and under what circumstances to enable face recognition.
- Statistic settings: Set whether the attendance group member can view the statistical data of certain dimensions.
- Overtime settings: Configure rules for calculating overtime hours.

## Request
|Facts||
|---|---|
|HTTP URL|https://open.larksuite.com/open-apis/attendance/v1/groups|
|HTTP Method|POST|
|HTTP Content-Type|application/json; charset=utf-8|
|Token requirement|tenant_access_token|
|Required scopes|Add attendance management rules|
### Query parameters
|Parameter|Type|Description|
|---|---|---|
|employee_type|string|User ID type, required field, values: [employee_id (the employee's employee ID), employee_no (the employee's employee number)]
|dept_type|string|Department ID type, required field, values: [open_id (only Open ID is currently supported)]. Example value: "od-fcb45c28a45311afd441b8869541ece8"|
### Request body
|Parameter|Type|Required|Description|
|---|---|---|---|
|group|group|Yes|Attendance group|
|  ∟group_id|string|No|Attendance group ID. You need to obtain it from the API that obtains the user's attendance result.|
|  ∟group_name|string|Yes|Attendance group name|
|  ∟time_zone|string|Yes|Time zone, refer to the time zone list at https://www.zeitverschiebung.net/cn/all-time-zones.html|
|  ∟bind_dept_ids|string[]|No| IDs of bound departments|
|  ∟except_dept_ids|string[]|No|IDs of excepted departments|
|  ∟bind_user_ids|string[]|No|IDs of bound users|
|  ∟except_user_ids|string[]|No|IDs of excepted users|
|  ∟group_leader_ids|string[]|Yes|Attendance owner ID list. Specifies at least 1 attendance owner|
|  ∟punch_type|int|No|Attendance method, 0: Attendance group members can clock in/out at any location and with any network environment, 1: GPS attendance, 2: Wi-Fi attendance, 4: Attendance machine attendance, 8: IP based attendance. Bit operations can support additional attendance methods, such as 3: GPS and Wi-Fi attendance, 7: GPS, Wi-Fi, and attendance machine attendance.|
|  ∟allow_out_punch|boolean|No|Whether offsite attendance is allowed|
|  ∟allow_pc_punch|boolean|No|Whether attendance on PC is allowed|
|  ∟allow_remedy|boolean|No|Whether correction is allowed|
|  ∟remedy_limit|boolean|No|Whether the number of corrections is limited|
|  ∟remedy_limit_count|int|No|Number of corrections|
|  ∟remedy_period_type|int|No|Cycle used to count corrections, 0: calendar month, 1: custom cycle|
|  ∟remedy_period_custom_date|int|No|Start day of custom correction cycle each month|
|  ∟remedy_date_limit|boolean|No|Whether the correction time is limited|
|  ∟remedy_date_num|int|No|Correction time period|
|  ∟show_cumulative_time|boolean|No|Whether to display cumulative work duration|
|  ∟show_over_time|boolean|No|Whether to display cumulative overtime duration|
|  ∟hide_staff_punch_time|boolean|No|Whether to hide specific employee attendance times|
|  ∟face_punch|boolean|No|Whether to enable face recognition attendance|
|  ∟face_punch_cfg|int|No|Face recognition attendance policy, 1: Face recognition required for each clock-in/out, 2: Face recognition only required when cheating is suspected|
|  ∟face_downgrade|boolean|No|Whether to allow normal photo attendance when face recognition fails|
|  ∟replace_basic_pic|boolean|No|Whether to allow baseline image replacement when face recognition fails|
|  ∟machines|machine[]|No|Attendance machine list|
|    ∟machine_sn|string|Yes|Attendance machine serial number|
|    ∟machine_name|string|Yes|Attendance machine name|
|  ∟locations|location[]|No|Address list|
|    ∟location_id|string|No|Address ID|
|    ∟location_name|string|Yes|Address name, required field|
|    ∟location_type|int|No|Address type, 1: GPS, 2: Wi-Fi, 8: IP|
|    ∟latitude|float|No|Address latitude|
|    ∟longitude|float|No|Address longitude|
|    ∟ssid|string|No|Wi-Fi name|
|    ∟bssid|string|No|Wi-Fi MAC address|
|    ∟map_type|int|No|Map type, 1: Amap, 2: Google|
|    ∟address|string|No|Address name|
|    ∟ip|string|No|IP address|
|    ∟feature|string|No|Additional information, such as operator information|
|    ∟gps_range|int|No|Valid range for GPS attendance, default value: 300 m|
|  ∟group_type|int|Yes|Attendance type, 0: fixed shift, 2: scheduled shift, 3: flexible shift|
|  ∟punch_day_shift_ids|string[]|Yes|Required for fixed shift, must be 7 digits|
|  ∟free_punch_cfg|free_punch_cfg|No|Flexible shift configuration|
|    ∟free_start_time|string|Yes|Flexible shift attendance start time|
|    ∟free_end_time|string|Yes|Flexible shift attendance end time|
|    ∟punch_day|int|Yes|Attendance time, 7 digits representing each day of the week from Mon. to Sun. 0 indicates no shift for the day and 1 indicates a shift for the day. The code for a normal Mon. to Fri. workweek is 1111100|
|    ∟work_day_no_punch_as_lack|boolean|No|Whether to mark employees as no record if they fail to clock in/out on a working day under the flexible shift|
|  ∟calendar_id|int|Yes|ID of the national public holiday calendar, required field, values: 0: days off not based on national public holiday calendar, 1: China (default), 2: US, 3: Japan, 4: India, 5: Singapore|
|  ∟need_punch_special_days|punch_special_date_shift[]|No|Special dates requiring attendance|
|    ∟punch_day|int|No|Attendance date, format: 20190101|
|    ∟shift_id|string|Yes |Shift ID|
|  ∟no_need_punch_special_days|punch_special_date_shift[]|No|Special dates not requiring attendance|
|    ∟punch_day|int|Yes|Attendance date, format: 20190101|
|    ∟shift_id|string|Yes |Shift ID|
|  ∟effect_now|boolean|No|Whether change takes effect immediately, default value: false|
### Request body example
```json
{
    "group": {
        "group_id": "6919358128597097404",
        "group_name": "Kaixin Attendance",
        "time_zone": "Asia/Shanghai",
        "bind_dept_ids": [
            "od-fcb45c28a45311afd440b7869541fce8"
        ],
        "except_dept_ids": [
            "od-fcb45c28a45311afd440b7869541fce8"
        ],
        "bind_user_ids": [
            "456123"
        ],
        "except_user_ids": [
            "456123"
        ],
        "group_leader_ids": [
             "456123"
        ],
        "punch_type": 7,
        "allow_out_punch": true,
        "allow_pc_punch": true,
        "allow_remedy": true,
        "remedy_limit": true,
        "remedy_limit_count": 3,
        "remedy_period_type": 0,
        "remedy_period_custom_date": 1,
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
        "locations": [
            {
                "location_id": "6919358128597017403",
                "location_name": "AVIC Plaza, Haidian District, Beijing",
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
            "6919358128597097424",
            "6919358128597097424",
            "6919358128597097424",
            "6919358128597097424",
            "6919358128597097424",
            "0",
            "0"
        ],
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
                "shift_id": "0"
            }
        ],
        "effect_now": true
    }
}
```
## Response
### Response body
|Parameter|Type|Description|
|---|---|---|
|code|int|Error code. A value other than 0 indicates failure.|
|msg|string|Error description|
|data|-|-|
|  ∟group|group|Attendance group|
|    ∟group_id|string|Attendance group ID. You need to obtain it from the API that obtains the user's attendance result. |
|    ∟group_name|string|Attendance group name|
|    ∟time_zone|string|Time zone|
|    ∟bind_dept_ids|string[]|Bound department IDs|
|    ∟except_dept_ids|string[]|Excepted department IDs|
|    ∟bind_user_ids|string[]|Bound user IDs|
|    ∟except_user_ids|string[]|Excepted user IDs|
|    ∟group_leader_ids|string[]|Attendance owner ID list, required field|
|    ∟punch_type|int|Attendance method, 0: Attendance group members can clock in/out at any location and with any network environment, 1: GPS attendance, 2: Wi-Fi attendance, 4: Attendance machine attendance, 8: IP based attendance. Bit operations can support additional attendance methods, such as 3: GPS and Wi-Fi attendance, 7: GPS, Wi-Fi, and attendance machine attendance.|
|    ∟allow_out_punch|boolean|Whether offsite attendance is allowed|
|    ∟allow_pc_punch|boolean|Whether attendance on PC is allowed|
|    ∟allow_remedy|boolean|Whether correction is allowed|
|    ∟remedy_limit|boolean|Whether the number of corrections is limited|
|    ∟remedy_limit_count|int|Number of corrections|
|    ∟remedy_period_type|int|Cycle used to count corrections, 0: calendar month, 1: custom cycle|
|    ∟remedy_period_custom_date|int|Start day of custom correction cycle each month|
|    ∟remedy_date_limit|boolean|Whether the correction time is limited|
|    ∟remedy_date_num|int|Correction time period|
|    ∟show_cumulative_time|boolean|Whether to display cumulative work duration|
|    ∟show_over_time|boolean|Whether to display cumulative overtime duration|
|    ∟hide_staff_punch_time|boolean|Whether to hide specific employee attendance times|
|    ∟face_punch|boolean|Whether to enable face recognition attendance|
|    ∟face_punch_cfg|int|Face recognition attendance policy, 1: Face recognition required for each clock-in/out, 2: Face recognition only required when cheating is suspected|
|    ∟face_downgrade|boolean|Whether to allow normal photo attendance when face recognition fails|
|    ∟replace_basic_pic|boolean|Whether to allow baseline image replacement when face recognition fails|
|    ∟machines|machine[]|Attendance machine list|
|      ∟machine_sn|string|Attendance machine serial number|
|      ∟machine_name|string|Attendance machine name|
|    ∟gps_range|int|Valid range for GPS attendance (not recommended)|
|    ∟locations|location[]|Address list|
|      ∟location_id|string|Address ID|
|      ∟location_name|string|Address name|
|      ∟location_type|int|Address type, 1: GPS, 2: Wi-Fi, 8: IP|
|      ∟latitude|float|Address latitude|
|      ∟longitude|float|Address longitude|
|      ∟ssid|string|Wi-Fi name|
|      ∟bssid|string|Wi-Fi MAC address|
|      ∟map_type|int|Map type, 1: Amap, 2: Google|
|      ∟address|string|Address name|
|      ∟ip|string|IP address|
|      ∟feature|string|Additional information, such as operator information|
|      ∟gps_range|int|Valid range for GPS attendance, default value: 300 m|
|    ∟group_type|int|Attendance type, 0: fixed shift, 2: scheduled shift, 3: flexible shift|
|    ∟punch_day_shift_ids|string[]|Required for fixed shift|
|    ∟free_punch_cfg|free_punch_cfg|Flexible shift configuration|
|      ∟free_start_time|string|Flexible shift attendance start time|
|      ∟free_end_time|string|Flexible shift attendance end time|
|      ∟punch_day|int|Attendance time, 7 digits representing each day of the week from Mon. to Sun. 0 indicates no shift for the day and 1 indicates a shift for the day. The code for a normal Mon. to Fri. workweek is 1111100|
|      ∟work_day_no_punch_as_lack|boolean|Whether to mark employees as no record if they fail to clock in/out on a working day under the flexible shift|
|    ∟calendar_id|int|ID of the national public holiday calendar, required field, values: 0: days off not based on national public holiday calendar, 1: China (default), 2: US, 3: Japan, 4: India, 5: Singapore|
|    ∟need_punch_special_days|punch_special_date_shift[]|Special dates requiring attendance|
|      ∟punch_day|int|Attendance date, format: 20190101|
|      ∟shift_id|string|Shift ID|
|    ∟no_need_punch_special_days|punch_special_date_shift[]|Special dates not requiring attendance|
|      ∟punch_day|int|Attendance date, format: 20190101|
|      ∟shift_id|string|Shift ID|
|    ∟work_day_no_punch_as_lack|boolean|Whether to mark employees as no record if they fail to clock in/out on a working day under the flexible shift|
### Response body example
```json
{
    "code": 0,
    "msg": "success",
    "data": {
        "group": {
            "group_id": "6919358128597097404",
            "group_name": "Kaixin Attendance",
            "time_zone": "Asia/Shanghai",
            "bind_dept_ids": [
                "od-fcb45c28a45311afd440b7869541fce8"
            ],
            "except_dept_ids": [
                "od-fcb45c28a45311afd440b7869541fce8"
            ],
            "bind_user_ids": [
                 "456123"
            ],
            "except_user_ids": [
                 "456123"
            ],
            "group_leader_ids": [
                "456123"
            ],
            "punch_type": 7,
            "allow_out_punch": true,
            "allow_pc_punch": true,
            "allow_remedy": true,
            "remedy_limit": true,
            "remedy_limit_count": 3,
            "remedy_period_type": 0,
            "remedy_period_custom_date": 1,
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
                    "location_id": "6919358128597017403",
                    "location_name": "AVIC Plaza, Haidian District, Beijing",
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
                "6919358128597097424",
                "6919358128597097424",
                "6919358128597097424",
                "6919358128597097424",
                "6919358128597097424",
                "0",
                "0"
            ],
            "calendar_id": 1,
            "need_punch_special_days": [
                {
                    "punch_day": 20190101,
                    "shift_id": "6919668827865513935"
                }
            ],
            "no_need_punch_special_days": [
                {
                    "punch_day": 0,
                    "shift_id": "6919668827865513935"
                }
            ],
        }
    }
}
```
### Error code
|HTTP status code|Error code|Description|Troubleshooting suggestions|
|---|---|---|---|
|400|1220001|Parameter error|Please check if the parameters meet the requirements|
|400|1220002|Tenant doesn't exist|Please check if the tenant_access_token is correct|
|400|1220003|employee_type doesn't exist|employee_type values: [employee_id (the employee's employee ID), employee_no (the employee's employee number)]|
|400|1220004|User doesn't exist|Please check if the user ID is correct|
|400|1220005|No permission scope|Please go to the Attendance Admin to check the data permission scope|
|500|1225000|System error|See error information for details|
|500|1227000|Management service system error|See the error information for details|
|500|1227500|Organizational structure service system error|See error message for details|
