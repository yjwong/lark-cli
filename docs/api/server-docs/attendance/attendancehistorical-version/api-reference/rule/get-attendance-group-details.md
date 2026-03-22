---
title: "Get Attendance Group Details"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/Attendance//group"
section: "Attendance"
updated: "1646322860000"
---

# Obtain attendance group details
This API is used to obtain attendance group details via an attendance group ID.

## Request
|Facts||
|---|---|
|HTTP URL|https://open.larksuite.com/open-apis/attendance/v1/groups/:group_id|
|HTTP Method|GET|
|HTTP Content-Type|application/json; charset=utf-8|
|Token requirement|tenant_access_token|
|Required scopes|Export attendance management rules|
### Path parameters
|Parameter|Type|Required|Description|
|---|---|---|---|
|group_id|string|Yes|Attendance group ID. It needs to be obtained from the API for obtaining attendance results. Example value: "6919358128597097404".|
### Query parameters
|Parameter|Type|Required|Description|
|---|---|---|---|
|employee_type|string|Yes|User ID type, values: [employee_id (the employee's employee ID), employee_no (the employee's employee number)]
|dept_type|string|Yes|Department ID type, values: [open_id (only Open ID of the department is currently supported)]. Example value: "od-fcb45c28a45311afd441b8869541ece8"|
## Response
### Response body
|Parameter|Type|Description|
|---|---|---|
|code|int|Error code. A value other than 0 indicates failure.|
|msg|string|Error description|
|data|-|-|
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
|      ∟gps_range|int|Valid range for GPS attendance|
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
|    ∟work_day_no_punch_as_lack|boolean|Whether to mark employees as no record if they fail to clock in/out on a working day under flexible shift|
### Response body example
```json
{
    "code": 0,
    "msg": "",
    "data": {
        "allow_out_punch": true,
        "allow_pc_punch": false,
        "allow_remedy": true,
        "bind_dept_ids": [
            "a6a64f94gg768492"
        ],
        "bind_user_ids": [
            "52aa1fa1"
        ],
        "calendar_id": 1,
        "except_dept_ids": [],
        "except_user_ids": [],
        "face_downgrade": true,
        "face_punch": true,
        "face_punch_cfg": 1,
        "gps_range": 300,
        "group_id": "6921319402260496386",
        "group_leader_ids": [
            "52aa1fa1",
            "2bg4a9be"
        ],
        "group_name": "now",
        "group_type": 0,
        "hide_staff_punch_time": true,
        "locations": [
            {
                "address": "",
                "bssid": "",
                "feature": "",
                "ip": "",
                "latitude": 30.28994,
                "location_id": "6921213751454744578",
                "location_name": "Xixi Bafang City, Muqiaotou, Wuchang Street, Yuhang District, Hangzhou, Zhejiang Province",
                "location_type": 1,
                "longitude": 120.04509,
                "map_type": 0,
                "ssid": "",
                "gps_range": 300,
            }
        ],
        "machines": [],
        "need_punch_special_days": [],
        "no_need_punch_special_days": [],
        "punch_day_shift_ids": [
            "",
            "",
            "",
            "",
            "",
            "",
            ""
        ],
        "punch_type": 1,
        "remedy_date_limit": true,
        "remedy_date_num": 30,
        "remedy_limit": true,
        "remedy_limit_count": 3,
        "remedy_period_type": 0,
        "remedy_period_custom_date": 1,
        "replace_basic_pic": true,
        "show_cumulative_time": true,
        "show_over_time": false,
        "time_zone": "Asia/Shanghai"
    }
}
```
### Error code
|HTTP status code|Error code|Description|Troubleshooting suggestions|
|---|---|---|---|
|400|1220001|Parameter error|Please check if the parameters meet the requirements|
|400|1220002|Tenant doesn't exist|Please check if the tenant_access_token is correct|
|400|1220003|employee_type doesn't exist|employee_type values: [employee_id (the employee's employee ID), employee_no (the employee's employee number)]|
|500|1225000|System error|See error information for details|
|500|1227000|Management service system error|See the error information for details|
