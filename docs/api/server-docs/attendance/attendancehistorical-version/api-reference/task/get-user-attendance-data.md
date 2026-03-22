---
title: "Get User Attendance Data"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/Attendance//RetrieveUserApprovals"
section: "Attendance"
updated: "1646322869000"
---

# Obtain user approval data
Obtain the approval data of employees' leave, overtime, offsite and business trip within a certain period of time.

## Request
|Facts||
|---|---|
|HTTP URL|https://open.larksuite.com/open-apis/attendance/v1/user_approvals/query|
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
|employee_type|string|Yes|The type of the user_ids in the request body, required field, values: [employee_id (the employee's employee ID), employee_no (the employee's employee number)]. Example value: "employee_id"|
### Request body
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
    "check_date_from": 20210103,
    "check_date_to": 20210104
}
```
## Response
### Response body
|Parameter|Type|Description|
|---|---|---|
|code|int|Error code. A value other than 0 indicates failure.|
|msg|string|Error description|
|data|-|-|
|  ∟user_approvals|user_approval[]|Approval result list|
|    ∟user_id|string|Approval user ID|
|    ∟date|string|Approval action time|
|    ∟outs|user_out[]|Offsite information|
|      ∟uniq_id|string|Unique offsite type ID, indicates the offsite type, up to 14 characters|
|      ∟unit|int|Offsite working time unit, values: [1 (days), 2 (hours), 3 (half-days), 4 (half-hours)]|
|      ∟interval|int|Offsite working time (in sec)|
|      ∟start_time|string|Start time, format: yyyy-MM-dd HH:mm:ss|
|      ∟end_time|string|End time, format: yyyy-MM-dd HH:mm:ss|
|      ∟i18n_names|i18n_names|Offsite multilingual display, format: map, key: ["ch", "en", "ja"]; ch: Chinese, en: English, ja: Japanese|
|        ∟ch|string|Chinese description|
|        ∟en|string|English description|
|        ∟ja|string|Japanese description|
|      ∟default_locale|string|Default language type. The Lark supports Chinese, English, and Japanese. If the user switches languages and the offsite name is not available in the corresponding language, this is the language displayed by default.|
|      ∟reason|string|Offsite reason|
|      ∟approve_pass_time|string|Approved time|
|      ∟approve_apply_time|string|Approval request time|
|    ∟leaves|user_leave[]|Leave information|
|      ∟uniq_id|string|Unique leave type ID, indicates the leave type, up to 14 characters|
|      ∟unit|int|unit of holiday time, available value: [1 (days), 0 (hours), 4 (minutes)]|
|      ∟start_time|string|Start time, format: yyyy-MM-dd HH:mm:ss|
|      ∟end_time|string|End time, format: yyyy-MM-dd HH:mm:ss|
|      ∟i18n_names|i18n_names|Leave multilingual display, format: map, key: ["ch", "en", "ja"]; ch: Chinese, en: English, ja: Japanese|
|        ∟ch|string|Chinese description|
|        ∟en|string|English description|
|        ∟ja|string|Japanese description|
|      ∟default_locale|string|Default language type, values: [ch (Chinese), en (English), ja (Japanese)]. The Lark supports Chinese, English, and Japanese. If the user switches languages and the leave name is not available in the corresponding language, this is the language displayed by default.|
|      ∟reason|string|Leave reason|
|      ∟approve_pass_time|string|Approved time, format: yyyy-MM-dd HH:mm:ss|
|      ∟approve_apply_time|string|Approval request time, format: yyyy-MM-dd HH:mm:ss|
|    ∟overtime_works|user_overtime_work[]|Overtime information|
|      ∟duration|float|Overtime duration|
|      ∟unit|int|Overtime duration unit, values: [1 (days), 2 (hours)]|
|      ∟category|int|Overtime date type, values [1 (working day), 2 (day off), 3 (holiday)]|
|      ∟type|int|types of overtime rules, available value: [0 (not associated with overtime rules), 1 (compensatory leave), 2 (overtime pay)]|
|      ∟start_time|string|Start time, format: yyyy-MM-dd HH:mm:ss|
|      ∟end_time|string|End time, format: yyyy-MM-dd HH:mm:ss|
|    ∟trips|user_trip[]|Business trip information|
|      ∟start_time|string|Start time, format: yyyy-MM-dd HH:mm:ss|
|      ∟end_time|string|End time, format: yyyy-MM-dd HH:mm:ss|
|      ∟reason|string|Business trip reason|
|      ∟approve_pass_time|string|Approved time, format: yyyy-MM-dd HH:mm:ss|
|      ∟approve_apply_time|string|Approval request time, format: yyyy-MM-dd HH:mm:ss|
### Response body example
```json
{
    "code": 0,
    "msg": "success",
    "data": {
        "user_approvals": [
            {
                "user_id": "abd754f7",
                "date": "20210104",
                "outs": [
                    {
                        "uniq_id": "9496E43696967658A512969523E89870",
                        "unit": 1,
                        "interval": 28800,
                        "start_time": "2021-01-04 09:00:00",
                        "end_time": "2021-01-04 19:00:00",
                        "i18n_names": {
                            "ch": "Chinese description",
                            "en": "english description",
                            "ja": "Japanese description"
                        },
                        "default_locale": "ch",
                        "reason": "Offsite errands",
                        "approve_pass_time": "2021-01-04 12:00:00",
                        "approve_apply_time": "2021-01-04 11:00:00"
                    }
                ],
                "leaves": [
                    {
                        "uniq_id": "6852582717813440527",
                        "unit": 1,
                        "interval": 28800,
                        "start_time": "2021-01-04 09:00:00",
                        "end_time": "2021-01-04 19:00:00",
                        "i18n_names": {
                            "ch": "Chinese description",
                            "en": "english description",
                            "ja": "Japanese description"
                        },
                        "default_locale": "ch",
                        "reason": "Personal reasons",
                        "approve_pass_time": "2021-01-04 12:00:00",
                        "approve_apply_time": "2021-01-04 11:00:00"
                    }
                ],
                "overtime_works": [
                    {
                        "duration": 1.5,
                        "unit": 1,
                        "category": 2,
                        "type": 1,
                        "start_time": "2021-01-04 09:00:00",
                        "end_time": "2021-01-04 13:00:00"
                    }
                ],
                "trips": [
                    {
                        "start_time": "2021-01-04 09:00:00",
                        "end_time": "2021-01-04 19:00:00",
                        "reason": "Training",
                        "approve_pass_time": "2021-01-04 12:00:00",
                        "approve_apply_time": "2021-01-04 11:00:00"
                    }
                ]
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
|500|1226000|Shift service system error|See error message for details|
|500|1226500|Attendance service system error|See error message for details|
