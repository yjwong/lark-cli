---
title: "Update Statistics Settings"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/Attendance//task/update-user-stats-settings"
method: "PUT"
api_path: "https://open.larksuite.com/open-apis/attendance/v1/user_stats_views/:user_stats_view_id"
section: "Attendance"
updated: "1646322867000"
---

# Update statistics settings
This API is used to update statistics setting information for daily or monthly statistics.

## Request
| HTTP URL | https://open.larksuite.com/open-apis/attendance/v1/user_stats_views/:user_stats_view_id |
| --- | --- |
| HTTP Method | PUT |
| Required scopes | Add attendance data | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" Learn more about how to access and use access_token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Path parameters
| `user_stats_view_id` | `string` | User view ID **Example value**: "TmpZNU5qTTJORFF6T1RnNU5UTTNOakV6TWl0dGIyNTBhQT09" |
| --- | --- | --- | ### Query parameters
| `employee_type` | `string` | Yes | User ID type **Optional values are**: - `employee_id`: User employee ID - `employee_no`: User employee number |
| --- | --- | --- | --- | ### Request body
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `view` | `user_stats_view` | Yes | Statistics view |
| ∟ `view_id` | `string` | Yes | View ID **Example value**: "TmpnNU5EQXpPVGN3TmpVMU16Y3lPVEEwTXl0dGIyNTBhQT09" |
| ∟ `stats_type` | `string` | Yes | Statistics type **Optional values are**: - `daily`: Daily statistics - `month`: Monthly statistics |
| ∟ `user_id` | `string` | Yes | User ID |
| ∟ `items` | `item[]` | No | Heading 1 |
| ∟ `code` | `string` | Yes | Number **Example value**: "501" |
| ∟ `title` | `string` | No | Heading name **Example value**: "Basic info" |
| ∟ `child_items` | `child_item[]` | No | Subheading |
| ∟ `code` | `string` | Yes | Heading number **Example value**: "50101" |
| ∟ `value` | `string` | Yes | Switch field **Optional values are**: - `0`: Off - `1`: On Non-switch field scene -  code = 51501  **Value range: 1-6** | ### Request body example

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
| `code` | `int` | Error code. A value other than 0 indicates failure. |
| `msg` | `string` | Error description |
| `data` | `\-` | \- |
| ∟ `view` | `user_stats_view` | User view |
| ∟ `view_id` | `string` | Statistics view ID **Example value**: "TmpnNU5EQXpPVGN3TmpVMU16Y3lPVEEwTXl0dGIyNTBhQT09" |
| ∟ `stats_type` | `string` | Statistics type **Optional values are**: - `daily`: Daily statistics - `month`: Monthly statistics |
| ∟ `user_id` | `string` | User ID |
| ∟ `items` | `item[]` | Heading 1 |
| ∟ `code` | `string` | Heading number |
| ∟ `title` | `string` | Heading name |
| ∟ `child_items` | `child_item[]` | Subheading |
| ∟ `code` | `string` | Heading number |
| ∟ `value` | `string` | On **Optional values are**: - `0`: Off - `1`: On | ### Response body example

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
| 400 | 1220001 | Parameter error | Please check if the parameters meet the requirements |
| 400 | 1220002 | Tenant doesn't exist | Please check if the tenant_access_token is correct |
| 500 | 1228000 | Statistics service system error | See error information for details |
