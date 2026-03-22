---
title: "Batch Query the Attendance Flow Records"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/Attendance//BatchQueryCheckinFlowHistory"
section: "Attendance"
updated: "1646322869000"
---

# Query attendance records in batches
Queries the actual attendance records of employees within your authorization scope in batches. If the shift set by the company for an employee is clock-in at 9:00 AM and clock-out at 6:00 PM and the employee clocks in and out several times during this period, this API will return all attendance records.

> **Note:** If you only need to obtain attendance results without detailed data, you can use the "Obtain attendance results" API.
## Request
|Facts||
|---|---|
|HTTP URL|https://open.larksuite.com/open-apis/attendance/v1/user_flows/query|
|HTTP Method|POST|
|HTTP Content-Type|application/json; charset=utf-8|
|Token requirement|tenant_access_token|
|Required scopes|Attendance data export|
### Header
key|value
--|--
Authorization|Bearer tenant_access_token
Content-Type|application/json
### Query parameters
|Parameter|Type|Required|Description|
|---|---|---|---|
|employee_type|string|Yes|The type of the user_ids in the request body, values: [employee_id (the employee's employee ID), employee_no (the employee's employee number)]. Example value: "employee_id"| ### Request body
|Parameter|Type|Required|Description|
|---|---|---|---|
|user_ids|string[]|true|employee_no or employee_id list, which can't exceed 50|
|check_time_from|string|Yes|Query start time, a timestamp|
|check_time_to|string|Yes|Query end time, a timestamp|
### Request body example
```json
{
    "user_ids": [
        "abd754f7"
    ],
    "check_time_from": "1566631088",
    "check_time_to": "1566641088"
}
```
## Response
### Response body
|Parameter|Type|Description|
|---|---|---|
|code|int|Error code. A value other than 0 indicates failure.|
|msg|string|Error description|
|data|-|-|
|  ∟user_flow_results|user_flow[]|Attendance record list|
|    ∟user_id|string|Employee number|
|    ∟creator_id|string|employee_no of attendance record creator|
|    ∟location_name|string|Attendance location name|
|    ∟check_time|string|Attendance time, with the timestamp accurate to sec|
|    ∟comment|string|Attendance notes| ∟record_id|string|Attendance record ID|
|    ∟longitude|float|Attendance longitude|
|    ∟latitude|float|Attendance latitude|
|    ∟ssid|string|Attendance Wi-Fi SSID|
|    ∟bssid|string|Attendance Wi-Fi MAC address|
|    ∟is_field|boolean|Whether the record is for offsite attendance|
|    ∟is_wifi|boolean|Whether the record is for Wi-Fi attendance|
|    ∟type|int|Record generation method, values: [0 (user clock-in/out), 1 (modified by administrator), 2 (user correction), 3 (auto generated), 4 (no clock-out), 5 (attendance machine clock-in/out), 6 (express attendance), 7 (imported from Attendance Open Platform), 8 (Lark custom attendance machine), 9 (Lark access control attendance machine)]|
|    ∟photo_urls|string[]|Attendance photo list|
|    ∟device_id|string|Mobile phone attendance device ID|
### Response body example
```json
{
    "code": 0,
    "msg": "success",
    "data": {
        "user_flow_results": [
            {
                "user_id": "abd754f7",
                "creator_id": "abd754f7",
                "location_name": "Xixi Bafang City",
                "check_time": "1611476284",
                "comment": "Clock-in",
                "record_id": "6709359313699356941",
                "longitude": 30.28991,
                "latitude": 120.04513,
                "ssid": "b0:b8:67:5c:1d:72",
                "bssid": "b0:b8:67:5c:1d:72",
                "is_field": true,
                "is_wifi": true,
                "type": 0,
                "photo_urls": [
                    "https://time.clockin.biz/manage/download/6840389754748502021"
                ],
                "device_id": "99e0609ee053448596502691a81428654d7ded64c7bd85acd982d26b3636c37d"
            }
        ]
    }
}
```
### Error code
|HTTP status code|Error code|Description|Troubleshooting suggestions|
|---|---|---|---|
|400|1220001|Parameter error|Please check if the parameters meet the requirements|
|400|1220002|Tenant doesn't exist|Please check if the tenant_access_token is correct|
|400|1220004|User doesn't exist or isn't in permission scope|Please check if the user ID is correct|
|400|1220005|No permission scope|Please go to the Attendance Admin to check the data permission scope|
|500|1225000|System error|See error information for details|
|500|1226500|Attendance service system error|See error message for details|
|500|1227500|Organizational structure service system error|See error message for details|
