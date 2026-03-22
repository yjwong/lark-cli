---
title: "Obtain Filter"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/sheets-v3/spreadsheet-sheet-filter/get"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/sheets/v3/spreadsheets/:spreadsheet_token/sheets/:sheet_id/filter"
service: "sheets-v3"
resource: "spreadsheet-sheet-filter"
section: "Docs"
scopes:
  - "drive:drive"
  - "drive:drive:readonly"
  - "sheets:spreadsheet"
  - "sheets:spreadsheet:readonly"
updated: "1664448897000"
---

# Obtain filter

This API is used to obtain the filter details for a sheet.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/sheets/v3/spreadsheets/:spreadsheet_token/sheets/:sheet_id/filter |
| HTTP Method | GET |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `drive:drive` `drive:drive:readonly` `sheets:spreadsheet` `sheets:spreadsheet:readonly` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" Learn more about obtaining and using access_token. | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `spreadsheet_token` | `string` | Spreadsheet token **Example value**: "shtcnmBA\*****yGehy8" |
| `sheet_id` | `string` | Sheet ID **Example value**: "0b\**12" | ```
curl --location --request GET 'https://open.feishu-pre.cn/open-apis/sheets/v3/spreadsheets/shtcnSUVpFeJ7QyVN9G1cHSc/sheets/6e2914/filter' \
--header 'Authorization: Bearer t-3bf3d4463b8f1956f14240c2517aa8ba2c93d8ec' \
--header 'Content-Type: application/json'
```

## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `sheet_filter_info` | `sheet_filter_info` | Filter information |
|     `range` | `string` | Range |
|     `filtered_out_rows` | `int[]` | Rows that are filtered out and hidden |
|     `filter_infos` | `filter_info[]` | Sheet filter conditions |
|       `col` | `string` | Columns with set filter conditions |
|       `conditions` | `condition[]` | Filter conditions |
|         `filter_type` | `string` | Filter type |
|         `compare_type` | `string` | Comparison type |
|         `expected` | `string[]` | Filter parameter | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
        "sheet_filter_info": {
            "range": "xxxxxx!A1:H14",
            "filtered_out_rows": [
                4
            ],
            "filter_infos": [
                {
                    "col": "E",
                    "conditions": [
                        {
                            "filter_type": "number",
                            "compare_type": "less",
                            "expected": [
                                "6"
                            ]
                        }
                    ]
                }
            ]
        }
    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 500 | 1315201 | Server Error | Internal service error. For more information, contact support. |
| 500 | 1315203 | Server Error | Internal service error. For more information, contact support. |
| 400 | 1310213 | Permission Fail | No permission |
| 400 | 1310215 | Sheet Id Not Found | Check the sheet_id. |
| 400 | 1310229 | Wrong URL Param | Check the url parameter. |
| 400 | 1310226 | Excess Limit | Exceeds the limit |
| 400 | 1310217 | Too Many Request | Check if requests have been sent too frequently. |
| 400 | 1310235 | Retry Later | Please try again later. |
| 400 | 1310242 | In Mix state | Retey Later |
| 400 | 1310249 | Spreadsheet Deleted | Restore the form and try again |
