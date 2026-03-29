---
title: "Query Statistics Configuration"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/Attendance//task/query-user-statistics-settings"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/attendance/v1/user_stats_views/query"
section: "Attendance"
updated: "1646322869000"
---

# Query statistics settings
This API is used to query statistics setting information for daily or monthly statistics.

## Request
| HTTP URL | https://open.larksuite.com/open-apis/attendance/v1/user_stats_views/query |
| --- | --- |
| HTTP Method | POST |
| Required scopes | Add attendance data Export attendance data | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" Learn more about how to access and use access_token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Query parameters
| `employee_type` | `string` | Yes | User ID type **Optional values are**: - `employee_id` - `employee_no` |
| --- | --- | --- | --- | ### Request body
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `locale` | `string` | Yes | Language **Optional values are**: - `en`: English - `ja`: Japanese - `zh`: Chinese |
| `stats_type` | `string` | Yes | Statistics type **Optional values are**: - `daily`: Daily statistics - `month`: Monthly statistics | ### Request body example

```json
{
    "locale": "zh",
    "stats_type": "month"
}
```

## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error code. A value other than 0 indicates failure. |
| `msg` | `string` | Error description |
| `data` | `\-` | \- |
| ∟ `view` | `user_stats_view` | Statistics view |
| ∟ `view_id` | `string` | View ID |
| ∟ `stats_type` | `string` | Statistics type **Optional values are**: - `daily`: Daily statistics - `month`: Monthly statistics |
| ∟ `user_id` | `string` | User ID |
| ∟ `items` | `item[]` | Heading 1 |
| ∟ `code` | `string` | Heading number |
| ∟ `title` | `string` | Heading name |
| ∟ `child_items` | `child_item[]` | Subheading |
| ∟ `code` | `string` | Heading number |
| ∟ `value` | `string` | On **Optional values are**: - `0`: Off - `1`: On |
| ∟ `title` | `string` | Heading name |
| ∟ `column_type` | `int` | Heading type |
| ∟ `read_only` | `boolean` | Read-only |
| ∟ `min_value` | `string` | Minpoint |
| ∟ `max_value` | `string` | Maxpoint | ### Response body example

```json
{
    "code": 0,
    "msg": "",
    "data": {
        "view": {
            "items": [
                {
                    "child_items": [
                        {
                            "code": "50101",
                            "column_type": 0,
                            "max_value": "",
                            "min_value": "",
                            "read_only": true,
                            "title": "Name",
                            "value": "1"
                        },
                        {
                            "code": "50102",
                            "column_type": 0,
                            "max_value": "",
                            "min_value": "",
                            "read_only": false,
                            "title": "Department",
                            "value": "0"
                        },
                        {
                            "code": "50111",
                            "column_type": 0,
                            "max_value": "",
                            "min_value": "",
                            "read_only": false,
                            "title": "Organizational structure",
                            "value": "0"
                        },
                        {
                            "code": "50103",
                            "column_type": 0,
                            "max_value": "",
                            "min_value": "",
                            "read_only": false,
                            "title": "Employee number",
                            "value": "1"
                        },
                        {
                            "code": "50104",
                            "column_type": 0,
                            "max_value": "",
                            "min_value": "",
                            "read_only": false,
                            "title": "Email",
                            "value": "0"
                        },
                        {
                            "code": "50105",
                            "column_type": 0,
                            "max_value": "",
                            "min_value": "",
                            "read_only": false,
                            "title": "Employee type",
                            "value": "0"
                        },
                        {
                            "code": "50106",
                            "column_type": 0,
                            "max_value": "",
                            "min_value": "",
                            "read_only": false,
                            "title": "Job family",
                            "value": "0"
                        },
                        {
                            "code": "50107",
                            "column_type": 0,
                            "max_value": "",
                            "min_value": "",
                            "read_only": false,
                            "title": "Date of employment",
                            "value": "0"
                        },
                        {
                            "code": "50108",
                            "column_type": 0,
                            "max_value": "",
                            "min_value": "",
                            "read_only": false,
                            "title": "Date of termination",
                            "value": "0"
                        },
                        {
                            "code": "50109",
                            "column_type": 0,
                            "max_value": "",
                            "min_value": "",
                            "read_only": false,
                            "title": "Status",
                            "value": "0"
                        },
                        {
                            "code": "50110",
                            "column_type": 0,
                            "max_value": "",
                            "min_value": "",
                            "read_only": false,
                            "title": "Direct leader",
                            "value": "0"
                        }
                    ],
                    "code": "501",
                    "title": "Basic info"
                },
                {
                    "child_items": [
                        {
                            "code": "52108",
                            "column_type": 0,
                            "max_value": "",
                            "min_value": "",
                            "read_only": false,
                            "title": "Attendance group name",
                            "value": "1"
                        },
                        {
                            "code": "52101",
                            "column_type": 0,
                            "max_value": "",
                            "min_value": "",
                            "read_only": false,
                            "title": "Expected on-duty days",
                            "value": "1"
                        },
                        {
                            "code": "52102",
                            "column_type": 0,
                            "max_value": "",
                            "min_value": "",
                            "read_only": false,
                            "title": "On-duty days on working days",
                            "value": "1"
                        },
                        {
                            "code": "52103",
                            "column_type": 0,
                            "max_value": "",
                            "min_value": "",
                            "read_only": false,
                            "title": "On-duty days on days off",
                            "value": "0"
                        },
                        {
                            "code": "52104",
                            "column_type": 0,
                            "max_value": "",
                            "min_value": "",
                            "read_only": false,
                            "title": "Expected working time",
                            "value": "1"
                        },
                        {
                            "code": "52105",
                            "column_type": 0,
                            "max_value": "",
                            "min_value": "",
                            "read_only": false,
                            "title": "Actual working time",
                            "value": "1"
                        },
                        {
                            "code": "52106",
                            "column_type": 0,
                            "max_value": "",
                            "min_value": "",
                            "read_only": false,
                            "title": "Working hours for pay calculation",
                            "value": "0"
                        },
                        {
                            "code": "52107",
                            "column_type": 0,
                            "max_value": "",
                            "min_value": "",
                            "read_only": false,
                            "title": "Overtime working hours",
                            "value": "1"
                        },
                        {
                            "code": "52109",
                            "column_type": 0,
                            "max_value": "",
                            "min_value": "",
                            "read_only": false,
                            "title": "Overtime duration (count as overtime pay)\n",
                            "value": "0"
                        },
                        {
                            "code": "52110",
                            "column_type": 0,
                            "max_value": "",
                            "min_value": "",
                            "read_only": false,
                            "title": "Overtime duration (count as compensatory leave)\n",
                            "value": "0"
                        }
                    ],
                    "code": "521",
                    "title": "Attendance statistics"
                },
                {
                    "child_items": [
                        {
                            "code": "52201",
                            "column_type": 0,
                            "max_value": "",
                            "min_value": "",
                            "read_only": false,
                            "title": "Late in count",
                            "value": "1"
                        },
                        {
                            "code": "52202",
                            "column_type": 0,
                            "max_value": "",
                            "min_value": "",
                            "read_only": false,
                            "title": "Late in duration",
                            "value": "0"
                        },
                        {
                            "code": "52203",
                            "column_type": 0,
                            "max_value": "",
                            "min_value": "",
                            "read_only": false,
                            "title": "Early out count",
                            "value": "1"
                        },
                        {
                            "code": "52204",
                            "column_type": 0,
                            "max_value": "",
                            "min_value": "",
                            "read_only": false,
                            "title": "Early out duration",
                            "value": "0"
                        },
                        {
                            "code": "52205",
                            "column_type": 0,
                            "max_value": "",
                            "min_value": "",
                            "read_only": false,
                            "title": "Missed clock-in count",
                            "value": "0"
                        },
                        {
                            "code": "52206",
                            "column_type": 0,
                            "max_value": "",
                            "min_value": "",
                            "read_only": false,
                            "title": "Missed clock-out count",
                            "value": "0"
                        },
                        {
                            "code": "52207",
                            "column_type": 0,
                            "max_value": "",
                            "min_value": "",
                            "read_only": false,
                            "title": "Absence",
                            "value": "1"
                        },
                        {
                            "code": "52208",
                            "column_type": 0,
                            "max_value": "",
                            "min_value": "",
                            "read_only": false,
                            "title": "Leave duration",
                            "value": "0"
                        },
                        {
                            "code": "52209",
                            "column_type": 0,
                            "max_value": "",
                            "min_value": "",
                            "read_only": false,
                            "title": "Business trip duration",
                            "value": "0"
                        },
                        {
                            "code": "52211",
                            "column_type": 0,
                            "max_value": "",
                            "min_value": "",
                            "read_only": false,
                            "title": "Shift swap days",
                            "value": "0"
                        },
                        {
                            "code": "52212",
                            "column_type": 0,
                            "max_value": "",
                            "min_value": "",
                            "read_only": false,
                            "title": "Correction count",
                            "value": "0"
                        },
                        {
                            "code": "52213",
                            "column_type": 0,
                            "max_value": "",
                            "min_value": "",
                            "read_only": false,
                            "title": "Offsite count",
                            "value": "0"
                        },
                        {
                            "code": "52214",
                            "column_type": 0,
                            "max_value": "",
                            "min_value": "",
                            "read_only": false,
                            "title": "Offsite working time\t",
                            "value": "0"
                        }
                    ],
                    "code": "522",
                    "title": "Abnormal statistics"
                },
                {
                    "child_items": [
                        {
                            "code": "52401",
                            "column_type": 0,
                            "max_value": "",
                            "min_value": "",
                            "read_only": false,
                            "title": "Daily attendance results",
                            "value": "1"
                        }
                    ],
                    "code": "524",
                    "title": "Daily statistics"
                }
            ],
            "stats_type": "month",
            "user_id": "ec8ddg56",
            "view_id": "TmpZNU5qTTJORFF6T1RnNU5UTTNOakV6TWl0dGIyNTBhQT09"
        }
    }
}
```

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 1220001 | Parameter error | Please check if the parameters meet the requirements |
| 400 | 1220002 | Tenant doesn't exist | Please check if the tenant_access_token is correct |
| 500 | 1228000 | Statistics service system error | See error information for details |
