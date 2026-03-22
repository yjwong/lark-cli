---
title: "Update Statistics Settings"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/attendance-v1/user_stats_view/update"
method: "PUT"
api_path: "https://open.larksuite.com/open-apis/attendance/v1/user_stats_views/:user_stats_view_id"
service: "attendance-v1"
resource: "user_stats_view"
section: "Attendance"
scopes:
  - "attendance:task"
updated: "1646322853000"
---

# Update statistics settings

Updates the header settings of statistical report for daily statistics or monthly statistics customized by developers.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/attendance/v1/user_stats_views/:user_stats_view_id |
| HTTP Method | PUT |
| Supported app types | custom |
| Required scopes | `attendance:task` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" Learn more about obtaining and using access_token. |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `user_stats_view_id` | `string` | User view ID is obtained as follows: 1) Query statistics settings **Example value**: "TmpZNU5qTTJORFF6T1RnNU5UTTNOakV6TWl0dGIyNTBhQT09" | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `employee_type` | `string` | Yes | Employee ID type **Example value**: "employee_id" **Optional values are**: - `employee_id`: Employee ID, that is the user ID in Lark Admin > Organization > Member and Department > Member Details - `employee_no`: Employee number, that is the employee ID in Lark Admin > Organization > Member and Department > Member Details | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `view` | `user_stats_view` | Yes | Statistics settings |
| ∟ `view_id` | `string` | Yes | View ID **Example value**: "TmpZNU5qTTJORFF6T1RnNU5UTTNOakV6TWl0dGIyNTBhQT09" |
| ∟ `stats_type` | `string` | Yes | View type **Example value**: "month" **Optional values are**: - `daily`: Daily statistics - `month`: Monthly statistics |
| ∟ `user_id` | `string` | Yes | User ID **Example value**: "ec8ddg56" |
| ∟ `items` | `item[]` | No | User settings field |
| ∟ `code` | `string` | Yes | Title number **Example value**: "522" |
| ∟ `child_items` | `child_item[]` | No | Subtitle |
| ∟ `code` | `string` | Yes | Subtitle number **Example value**: "50101" |
| ∟ `value` | `string` | Yes | Switch field, 0: off, 1: on (non-switch field scenario: code attendance.v1.type.item.prop.title.desc=$$$Title name **Example value**: "0" | ### Request body example

```json
{
    "view": {
        "items": [
            {
                "child_items": [
                    {
                        "code": "50102",
                        "value": "0"
                    },
                    {
                        "code": "50111",
                        "value": "0"
                    },
                    {
                        "code": "50104",
                        "value": "0"
                    }
                ],
                "code": "501"
            }
        ],
        "stats_type": "month",
        "user_id": "ec8ddg56",
        "view_id": "TmpnNU5EQXpPVGN3TmpVMU16Y3lPVEEwTXl0dGIyNTBhQT09"
    }
}
```

## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
| ∟ `view` | `user_stats_view` | View |
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
            "items": [
                {
                    "child_items": [
                        {
                            "code": "50102",
                            "value": "0"
                        },
                        {
                            "code": "50111",
                            "value": "0"
                        },
                        {
                            "code": "50104",
                            "value": "0"
                        }
                    ],
                    "code": "501"
                }
            ],
            "stats_type": "month",
            "user_id": "ec8ddg56",
            "view_id": "TmpnNU5EQXpPVGN3TmpVMU16Y3lPVEEwTXl0dGIyNTBhQT09"
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
