---
title: "Query Statistics Settings"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/attendance-v1/user_stats_view/query"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/attendance/v1/user_stats_views/query"
service: "attendance-v1"
resource: "user_stats_view"
section: "Attendance"
scopes:
  - "attendance:task"
  - "attendance:task:readonly"
updated: "1646322853000"
---

# Query statistics settings

Queries the header settings of statistical report for daily statistics or monthly statistics customized by developers.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/attendance/v1/user_stats_views/query |
| HTTP Method | POST |
| Supported app types | custom |
| Required scopes Enable any scope from the list | `attendance:task` `attendance:task:readonly` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" Learn more about obtaining and using access_token. |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `employee_type` | `string` | Yes | Type of employee ID for user_id in response body **Example value**: "employee_id" **Optional values are**: - `employee_id`: Employee ID, that is the user ID in Lark Admin > Organization > Member and Department > Member Details - `employee_no`: Employee number, that is the employee ID in Lark Admin > Organization > Member and Department > Member Details | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `locale` | `string` | Yes | Language **Example value**: "zh" **Optional values are**: - `en`: English - `ja`: Japanese - `zh`: Chinese |
| `stats_type` | `string` | Yes | Statistics type **Example value**: "daily" **Optional values are**: - `daily`: Daily statistics - `month`: Monthly statistics | ### Request body example

```json
{
    "locale": "zh",
    "stats_type": "daily"
}
```

## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
| ∟ `view` | `user_stats_view` | Statistics view |
| ∟ `view_id` | `string` | View ID |
| ∟ `stats_type` | `string` | View type **Optional values are**: - `daily`: Daily statistics - `month`: Monthly statistics |
| ∟ `user_id` | `string` | User ID |
| ∟ `items` | `item[]` | User settings field |
| ∟ `code` | `string` | Title number |
| ∟ `title` | `string` | Title name |
| ∟ `child_items` | `child_item[]` | Subtitle |
| ∟ `code` | `string` | Subtitle number |
| ∟ `value` | `string` | Switch field, 0: off, 1: on (non-switch field scenario: code attendance.v1.type.item.prop.title.desc=$$$Title name |
| ∟ `title` | `string` | Subtitle name |
| ∟ `column_type` | `int` | Column type |
| ∟ `read_only` | `boolean` | Read-only |
| ∟ `min_value` | `string` | Min. |
| ∟ `max_value` | `string` | Max. | ### Response body example

```json
{
	"code": 0,
	"msg": "",
	"data": {
		"view": {
			"items": [{
				"child_items": [{
					"code": "50101",
					"column_type": 0,
					"max_value": "",
					"min_value": "",
					"read_only": true,
					"title": "Name",
					"value": "1"
				}, {
					"code": "50102",
					"column_type": 0,
					"max_value": "",
					"min_value": "",
					"read_only": false,
					"title": "Department",
					"value": "0"
				}, {
					"code": "50111",
					"column_type": 0,
					"max_value": "",
					"min_value": "",
					"read_only": false,
					"title": "Organization",
					"value": "0"
				}, {
					"code": "50103",
					"column_type": 0,
					"max_value": "",
					"min_value": "",
					"read_only": false,
					"title": "Employee ID",
					"value": "1"
				}, {
					"code": "50104",
					"column_type": 0,
					"max_value": "",
					"min_value": "",
					"read_only": false,
					"title": "Email",
					"value": "0"
				}, {
					"code": "50105",
					"column_type": 0,
					"max_value": "",
					"min_value": "",
					"read_only": false,
					"title": "Employee type",
					"value": "0"
				}, {
					"code": "50106",
					"column_type": 0,
					"max_value": "",
					"min_value": "",
					"read_only": false,
					"title": "Sequence",
					"value": "0"
				}, {
					"code": "50107",
					"column_type": 0,
					"max_value": "",
					"min_value": "",
					"read_only": false,
					"title": "Date of employment",
					"value": "0"
				}, {
					"code": "50108",
					"column_type": 0,
					"max_value": "",
					"min_value": "",
					"read_only": false,
					"title": "Date of termination",
					"value": "0"
				}, {
					"code": "50109",
					"column_type": 0,
					"max_value": "",
					"min_value": "",
					"read_only": false,
					"title": "Status",
					"value": "0"
				}, {
					"code": "50110",
					"column_type": 0,
					"max_value": "",
					"min_value": "",
					"read_only": false,
					"title": "Direct leader",
					"value": "0"
				}],
				"code": "501",
				"title": "Basic information"
			}, {
				"child_items": [{
					"code": "52108",
					"column_type": 0,
					"max_value": "",
					"min_value": "",
					"read_only": false,
					"title": "Attendance group name",
					"value": "1"
				}, {
					"code": "52101",
					"column_type": 0,
					"max_value": "",
					"min_value": "",
					"read_only": false,
					"title": "Expected attendance days",
					"value": "1"
				}, {
					"code": "52102",
					"column_type": 0,
					"max_value": "",
					"min_value": "",
					"read_only": false,
					"title": "On-duty days on working days",
					"value": "1"
				}, {
					"code": "52103",
					"column_type": 0,
					"max_value": "",
					"min_value": "",
					"read_only": false,
					"title": "On-duty days on days off",
					"value": "0"
				}, {
					"code": "52104",
					"column_type": 0,
					"max_value": "",
					"min_value": "",
					"read_only": false,
					"title": "Expected working time",
					"value": "1"
				}, {
					"code": "52105",
					"column_type": 0,
					"max_value": "",
					"min_value": "",
					"read_only": false,
					"title": "Actual working time",
					"value": "1"
				}, {
					"code": "52106",
					"column_type": 0,
					"max_value": "",
					"min_value": "",
					"read_only": false,
					"title": "Paid working hours",
					"value": "0"
				}, {
					"code": "52107",
					"column_type": 0,
					"max_value": "",
					"min_value": "",
					"read_only": false,
					"title": "Overtime duration",
					"value": "1"
				}, {
					"code": "52109",
					"column_type": 0,
					"max_value": "",
					"min_value": "",
					"read_only": false,
					"title": "Overtime duration (calculate overtime pay)",
					"value": "0"
				}, {
					"code": "52110",
					"column_type": 0,
					"max_value": "",
					"min_value": "",
					"read_only": false,
					"title": "Overtime duration (count as compensatory leave)",
					"value": "0"
				}],
				"code": "521",
				"title": "Attendance statistics"
			}, {
				"child_items": [{
					"code": "52201",
					"column_type": 0,
					"max_value": "",
					"min_value": "",
					"read_only": false,
					"title": "Late-in count",
					"value": "1"
				}, {
					"code": "52202",
					"column_type": 0,
					"max_value": "",
					"min_value": "",
					"read_only": false,
					"title": "Late-in duration",
					"value": "0"
				}, {
					"code": "52203",
					"column_type": 0,
					"max_value": "",
					"min_value": "",
					"read_only": false,
					"title": "Early-out count",
					"value": "1"
				}, {
					"code": "52204",
					"column_type": 0,
					"max_value": "",
					"min_value": "",
					"read_only": false,
					"title": "Early-out duration",
					"value": "0"
				}, {
					"code": "52205",
					"column_type": 0,
					"max_value": "",
					"min_value": "",
					"read_only": false,
					"title": "Missed punch-in count",
					"value": "0"
				}, {
					"code": "52206",
					"column_type": 0,
					"max_value": "",
					"min_value": "",
					"read_only": false,
					"title": "Missed punch-out count",
					"value": "0"
				}, {
					"code": "52207",
					"column_type": 0,
					"max_value": "",
					"min_value": "",
					"read_only": false,
					"title": "Absence",
					"value": "1"
				}, {
					"code": "52208",
					"column_type": 0,
					"max_value": "",
					"min_value": "",
					"read_only": false,
					"title": "Leave duration",
					"value": "0"
				}, {
					"code": "52209",
					"column_type": 0,
					"max_value": "",
					"min_value": "",
					"read_only": false,
					"title": "Business trip duration",
					"value": "0"
				}, {
					"code": "52211",
					"column_type": 0,
					"max_value": "",
					"min_value": "",
					"read_only": false,
					"title": "Shift swap days",
					"value": "0"
				}, {
					"code": "52212",
					"column_type": 0,
					"max_value": "",
					"min_value": "",
					"read_only": false,
					"title": "Correction count",
					"value": "0"
				}, {
					"code": "52213",
					"column_type": 0,
					"max_value": "",
					"min_value": "",
					"read_only": false,
					"title": "Offsite count",
					"value": "0"
				}, {
					"code": "52214",
					"column_type": 0,
					"max_value": "",
					"min_value": "",
					"read_only": false,
					"title": "Offsite working time\t",
					"value": "0"
				}],
				"code": "522",
				"title": "Abnormal statistics"
			}, {
				"child_items": [{
					"code": "52401",
					"column_type": 0,
					"max_value": "",
					"min_value": "",
					"read_only": false,
					"title": "Daily attendance results",
					"value": "1"
				}],
				"code": "524",
				"title": "Daily statistics"
			}],
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
| 400 | 1220001 | Parameter error | Please check if the parameters meet the requirements. |
| 400 | 1220002 | Tenant doesn't exist. | Please check if the tenant_access_token is correct. |
| 500 | 1228000 | Statistics service system error | See error information for details. |
