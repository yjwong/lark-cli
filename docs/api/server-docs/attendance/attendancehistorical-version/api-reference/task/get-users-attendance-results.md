---
title: "Get Users’ Attendance Results"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/Attendance//GetCheckinResults"
section: "Attendance"
updated: "1646322869000"
---

# Obtain attendance results
Obtains the actual attendance results of employees in the Company, including the clock-in/out result.

> **Note:** * If the shift set by the company for an employee is clock-in at 9:00 AM and clock-out at 6:00 PM, even if the employee submits attendance for several times during this period, the API will only return 1 record.
> * If you want to obtain detailed data of attendance, such as information of attendance location, you can use the API of "Obtain attendance records" or "Query attendance records in batches".
## Request
|Facts||
|---|---|
|HTTP URL|https://open.larksuite.com/open-apis/attendance/v1/user_tasks/query|
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
|employee_type|string|Yes|The type of the user_ids in the request body, values: [employee_id (the employee's employee ID), employee_no (the employee's employee number)]. Example value: "employee_id"|
|ignore_invalid_users|boolean|No|Whether to ignore invalid users and users for which the current user doesn't have permission scope. If "true", data is only returned for valid users and a prompt tells the current user that some users were invalid or outside the user's permission scope. If "false", the operation returns an error if any user_ids are invalid.| ### Request body
|Parameter|Type|Required|Description|
|---|---|---|---|
|user_ids|string[]|Yes|employee_no or employee_id list|
|check_date_from|int|Yes|Start day of check|
|check_date_to|int|Yes|End day of check, which should be no more than 30 days later than the check_date_from|
### Request body example
```json
{
    "user_ids": [
        "abd754f7"
    ],
    "check_date_from": 20210104,
    "check_date_to": 20210105
}
```
## Response
### Response body
|Parameter|Type|Description|
|---|---|---|
|code|int|Error code. A value other than 0 indicates failure.|
|msg|string|Error description|
|data|-|-|
|  ∟user_task_results|user_task[]|List of attendance tasks|
|    ∟result_id|string|Attendance record ID|
|    ∟user_id|string|Employee ID|
|    ∟employee_name|string|Employee name|
|    ∟day|int|Date|
|    ∟group_id|string|Attendance group ID|
|    ∟shift_id|string|Shift ID|
|    ∟records|task_result[]|User attendance records|
|      ∟check_in_record_id|string|Clock-in attendance record ID|
|      ∟check_in_record|user_flow[]|Clock-in attendance record|
|        ∟user_id|string|Employee number|
|        ∟creator_id|string|employee_no of attendance record creator|
|        ∟location_name|string|Attendance location name|
|        ∟check_time|string|Attendance time, with the timestamp accurate to sec|
|        ∟comment|string|Attendance notes|
|        ∟record_id|string|Attendance record ID|
|        ∟longitude|float|Attendance longitude|
|        ∟latitude|float|Attendance latitude|
|        ∟ssid|string|Attendance Wi-Fi SSID|
|        ∟bssid|string|Attendance Wi-Fi MAC address|
|        ∟is_field|boolean|Whether the record is for offsite attendance|
|        ∟is_wifi|boolean|Whether the record is for Wi-Fi attendance|
|        ∟type|int|Record generation method, values: [0 (user clock-in/out), 1 (modified by administrator), 2 (user correction), 3 (auto generated), 4 (no clock-out), 5 (attendance machine clock-in/out), 6 (express attendance), 7 (imported from Attendance Open Platform), 8 (Lark custom attendance machine), 9 (Lark access control attendance machine)]|
|        ∟photo_urls|string[]|Attendance photo list|
|        ∟device_id|string|Mobile phone attendance device ID|
|      ∟check_out_record_id|string|Clock-out attendance record ID|
|      ∟check_out_record|user_flow[]|Clock-out attendance record|
|        ∟user_id|string|Employee number|
|        ∟creator_id|string|employee_no of attendance record creator|
|        ∟location_name|string|Attendance location name|
|        ∟check_time|string|Attendance time, with the timestamp accurate to sec|
|        ∟comment|string|Attendance notes|
|        ∟record_id|string|Attendance record ID|
|        ∟longitude|float|Attendance longitude|
|        ∟latitude|float|Attendance latitude|
|        ∟ssid|string|Attendance Wi-Fi SSID|
|        ∟bssid|string|Attendance Wi-Fi MAC address|
|        ∟is_field|boolean|Whether the record is for offsite attendance|
|        ∟is_wifi|boolean|Whether the record is for Wi-Fi attendance|
|        ∟type|int|Record generation method, values: [0 (user clock-in/out), 1 (modified by administrator), 2 (user correction), 3 (auto generated), 4 (no clock-out), 5 (attendance machine clock-in/out), 6 (express attendance), 7 (imported from Attendance Open Platform), 8 (Lark custom attendance machine), 9 (Lark access control attendance machine)]|
|        ∟photo_urls|string[]|Attendance photo list|
|        ∟device_id|string|Mobile phone attendance device ID|
|      ∟check_in_result|string|Clock-in attendance result, values: [NoNeedCheck (no clock-in/out required), SystemCheck (system clock-in/out), Normal (regular attendance), Early (early out), Late (late in), Lack (no record)]|
|      ∟check_out_result|string|Clock-out attendance result, values: [NoNeedCheck (no clock-in/out required), SystemCheck (system clock-in/out), Normal (regular attendance), Early (early out), Late (late in), Lack (no record)]|
|      ∟check_in_result_supplement|string|Clock-in attendance result supplement, values: [None (none), ManagerModification (modified by administrator), CardReplacement (correction approved), ShiftChange (shift swap), Travel (business trip), Leave (on leave), GoOut (offsite), CardReplacementApplication (correction application under review), FieldPunch (offsite clock-in/out)]|
|      ∟check_out_result_supplement|string|Clock-out attendance result supplement, values: [None (none), ManagerModification (modified by administrator), CardReplacement (correction approved), ShiftChange (shift swap), Travel (business trip), Leave (on leave), GoOut (offsite), CardReplacementApplication (correction application under review), FieldPunch (offsite clock-in/out)]|
|      ∟check_in_shift_time|string|Normal default clock-in time, with the timestamp accurate to sec|
|      ∟check_out_shift_time|string|Normal default clock-out time, with the timestamp accurate to sec|
|  ∟invalid_user_ids|string[]|List of invalid employee IDs of invalid users|
|  ∟unauthorized_user_ids|string[]|List of employee IDs of users outside permission scope|
### Response body example
```json
{
	"code": 0,
	"msg": "success",
	"data": {
		"user_task_results": [{
			"result_id": "6709359313699356941",
			"user_id": "abd754f7",
			"employee_name": "John",
			"day": 20210104,
			"group_id": "6737202939523236110",
			"shift_id": "6753520403404030215",
			"records": [{
				"check_in_record_id": "6709359313699356941",
				"check_in_record": {
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
				"check_out_record_id": "6709359313699356942",
				"check_out_record": {
					"user_id": "abd754f7",
					"creator_id": "abd754f7",
					"location_name": "Xixi Bafang City",
					"check_time": "1611476284",
					"comment": "Clock-out",
					"record_id": "6709359313699356942",
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
				"check_in_result": "SystemCheck",
				"check_out_result": "SystemCheck",
				"check_in_result_supplement": "None",
				"check_out_result_supplement": "None",
				"check_in_shift_time": "1611476284",
				"check_out_shift_time": "1611476284"
			}]
		}],
		"invalid_user_ids": ["abd754xx"],
		"unauthorized_user_ids": ["abd75xxx"]
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
