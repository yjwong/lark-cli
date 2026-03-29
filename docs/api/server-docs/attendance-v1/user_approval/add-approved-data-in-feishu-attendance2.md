---
title: "Add Approved Data in Lark Attendance"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/attendance-v1/user_approval/add-approved-data-in-feishu-attendance2"
service: "attendance-v1"
resource: "user_approval"
section: "Attendance"
updated: "1707016017000"
---

# Add approval results
Because some companies use their own approval system rather than the Lark approval system, employee leave, overtime, and other data can't be connected to the Lark attendance system. As a result, employees will still receive attendance reminders when they are on vacation and will be marked as no record.

Companies that use the Lark attendance system but not the Lark approval system can use the Attendance Open API to write third-party approval result data back to the Lark attendance system.

> **Note:** Currently, only the overtime, leave, business trip, and offsite approval results can be added in this way.
## Request
|Facts||
|---|---|
|HTTP URL|https://open.larksuite.com/open-apis/attendance/v1/user_approvals|
|HTTP Method|POST|
|HTTP Content-Type|application/json; charset=utf-8|
|Token requirement|tenant_access_token|
|Required scopes|Add attendance data|
### Header
key|value
--|--
Authorization|Bearer tenant_access_token
Content-Type|application/json
### Query parameters
|Parameter|Type|Required|Description|
|---|---|---|---|
|employee_type|string|Yes|The type of employee ID for the user_id in the request body. Values: [employee_id (employee's employee ID, that is, the user ID in Lark Admin > Organization > Member and Department > Member Details), employee_no (employee ID, that is, the employee ID in Lark Admin > Organization > Member and Department > Member Details)]. Example value: "employee_id".|
### Request body
|Parameter|Type|Required|Description|
|---|---|---|---|
|user_approval|user_approval|Yes|Approval information|
|  ∟user_id|string|Yes|Approval user|
|  ∟date|string|Yes|Approval action time|
|  ∟outs|user_out[]|No|Offsite information|
|    ∟uniq_id|string|Yes|Unique offsite type ID, indicates the offsite type, up to 14 characters|
|    ∟unit|int|yes|Offsite working time unit, values: [1: forenoon, 2: afternoon, 3: whole day, 4: hour.]|
|    ∟interval|int|Yes|Leave duration (in sec)|
|    ∟start_time|string|Yes|Start time, format: yyyy-MM-dd HH:mm:ss|
|    ∟end_time|string|Yes|End time, format: yyyy-MM-dd HH:mm:ss|
|    ∟i18n_names|i18n_names|Yes|Offsite multilingual display, format: map, key: ["ch", "en", "ja"]; ch: Chinese, en: English, ja: Japanese|
|      ∟ch|string|Yes|Chinese description|
|      ∟en|string|Yes|English description|
|      ∟ja|string|Yes|Japanese description|
|    ∟default_locale|string|Yes|Default language type. The Lark supports Chinese, English, and Japanese. If the user switches languages and the offsite name is not available in the corresponding language, this is the language displayed by default.|
|    ∟reason|string|Yes|Offsite reason|
|  ∟leaves|user_leave[]|No|Leave information|
|    ∟uniq_id|string|Yes|Unique leave type ID, indicates the leave type, up to 14 characters|
|    ∟unit|int|Yes|Leave duration unit, values: [1 (days), 2 (hours), 3 (half-days), 4 (half-hours)]|
|    ∟interval|int|Yes|Leave duration (in sec)|
|    ∟start_time|string|Yes|Start time, format: yyyy-MM-dd HH:mm:ss|
|    ∟end_time|string|Yes|End time, format: yyyy-MM-dd HH:mm:ss|
|    ∟i18n_names|i18n_names|Yes|Leave multilingual display, format: map, key: ["ch", "en", "ja"]; ch: Chinese, en: English, ja: Japanese|
|      ∟ch|string|Yes|Chinese description|
|      ∟en|string|Yes|English description|
|      ∟ja|string|Yes|Japanese description|
|    ∟default_locale|string|Yes|Default language type, values: [ch (Chinese), en (English), ja (Japanese)]. The Lark supports Chinese, English, and Japanese. If the user switches languages and the leave name is not available in the corresponding language, this is the language displayed by default.|
|    ∟reason|string|Yes|Leave reason, required field|
|  ∟overtime_works|user_overtime_work[]|No|Overtime information|
|    ∟duration|float|Yes|Overtime duration|
|    ∟unit|int|Yes|Overtime duration unit, values: [1 (days), 2 (hours)]|
|    ∟category|int|Yes|Overtime date type, values [1 (working day), 2 (day off), 3 (holiday)]|
|    ∟type|int|Yes|Overtime rule type, values: [0 (no associated rules), 1 (compensatory leave), 2 (overtime pay)]|
|    ∟start_time|string|Yes|Start time, format: yyyy-MM-dd HH:mm:ss|
|    ∟end_time|string|Yes|End time, format: yyyy-MM-dd HH:mm:ss|
|  ∟trips|user_trip[]|No|Business trip information|
|    ∟start_time|string|Yes|Start time, format: yyyy-MM-dd HH:mm:ss|
|    ∟end_time|string|Yes|End time, format: yyyy-MM-dd HH:mm:ss|
|    ∟reason|string|Yes|Business trip reason|
|    ∟approve_pass_time|string|Yes|Approved time, format: yyyy-MM-dd HH:mm:ss|
|    ∟approve_apply_time|string|Yes|Approval request time, format: yyyy-MM-dd HH:mm:ss|
### Request body example
```json
{
    "user_approval": {
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
                    "ch": "中文描述",
                    "en": "english description",
                    "ja": "日本語の説明"
                },
                "default_locale": "ch",
                "reason": "Offsite errands"
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
                    "ch": "中文描述",
                    "en": "english description",
                    "ja": "日本語の説明"
                },
                "default_locale": "ch",
                "reason": "Personal reasons"
            }
        ],
        "overtime_works": [
            {
                "duration": 1.5,
                "unit": 1,
                "category": 2,
                "type": 1,
                "start_time": "2021-01-09 09:00:00",
                "end_time": "2021-01-10 13:00:00"
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
|      ∟approval_id|string|Approval instance ID
|      ∟uniq_id|string|Unique offsite type ID, indicates the offsite type, up to 14 characters|
|      ∟unit|int|Offsite working time unit, values: [1 (days), 2 (hours), 3 (half-days), 4 (half-hours)]|
|      ∟interval|int|Leave duration (in sec)|
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
|      ∟approval_id|string|Approval instance ID
|      ∟uniq_id|string|Unique leave type ID, indicates the leave type, up to 14 characters|
|      ∟unit|int|Leave duration unit, values: [1 (days), 2 (hours), 3 (half-days), 4 (half-hours)]|
|      ∟interval|int|Leave duration (in sec)|
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
|      ∟approval_id|string|Approval instance ID
|      ∟duration|float|Overtime duration|
|      ∟unit|int|Overtime duration unit, values: [1 (days), 2 (hours)]|
|      ∟category|int|Overtime date type, values [1 (working day), 2 (day off), 3 (holiday)]|
|      ∟type|int|Overtime rule type, values: [0 (no associated rules), 1 (compensatory leave), 2 (overtime pay), 3 (other associated overtime rule)]|
|      ∟start_time|string|Start time, format: yyyy-MM-dd HH:mm:ss|
|      ∟end_time|string|End time, format: yyyy-MM-dd HH:mm:ss|
|    ∟trips|user_trip[]|Business trip information|
|      ∟approval_id|string|Approval instance ID
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
                        "approval_id": "6737202939523236113",
                        "uniq_id": "9496E43696967658A512969523E89870",
                        "unit": 1,
                        "interval": 28800,
                        "start_time": "2021-01-04 09:00:00",
                        "end_time": "2021-01-04 19:00:00",
                        "i18n_names": {
                            "ch": "中文描述",
                            "en": "english description",
                            "ja": "日本語の説明"
                        },
                        "default_locale": "ch",
                        "reason": "Offsite errands",
                        "approve_pass_time": "2021-01-04 12:00:00",
                        "approve_apply_time": "2021-01-04 11:00:00"
                    }
                ],
                "leaves": [
                    {
                        "approval_id": "6737202939523236114",
                        "uniq_id": "6852582717813440527",
                        "unit": 1,
                        "interval": 28800,
                        "start_time": "2021-01-04 09:00:00",
                        "end_time": "2021-01-04 19:00:00",
                        "i18n_names": {
                            "ch": "中文描述",
                            "en": "english description",
                            "ja": "日本語の説明"
                        },
                        "default_locale": "ch",
                        "reason": "Personal reasons",
                        "approve_pass_time": "2021-01-04 12:00:00",
                        "approve_apply_time": "2021-01-04 11:00:00"
                    }
                ],
                "overtime_works": [
                    {
                        "approval_id": "6737202939523236115",
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
                        "approval_id": "6737202939523236116",
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
|400|1220004|User does not exist or has no scope|Please check if the user ID is correct|
|400|1220005|No scope|Please go to [Attendance Admin](https://oa.larksuite.com/attendance/manage/member/list) to check data scope|
|500|1225000|System error|See error message for details|
