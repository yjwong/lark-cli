---
title: "Update Filter"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/sheets-v3/spreadsheet-sheet-filter/update"
method: "PUT"
api_path: "https://open.larksuite.com/open-apis/sheets/v3/spreadsheets/:spreadsheet_token/sheets/:sheet_id/filter"
service: "sheets-v3"
resource: "spreadsheet-sheet-filter"
section: "Docs"
scopes:
  - "drive:drive"
  - "sheets:spreadsheet"
updated: "1664448897000"
---

# Update a filter

::: note
For parameter values, see Filter guide

This API is used to update column filter conditions in the sheet range.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/sheets/v3/spreadsheets/:spreadsheet_token/sheets/:sheet_id/filter |
| HTTP Method | PUT |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `drive:drive` `sheets:spreadsheet` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" Learn more about obtaining and using access_token. |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `spreadsheet_token` | `string` | Spreadsheet token **Example value**: "shtcnmBA\*****yGehy8" |
| `sheet_id` | `string` | Sheet ID **Example value**: "0b\**12" | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `col` | `string` | Yes | Column with updated filter conditions **Example value**: "E" |
| `condition` | `condition` | Yes | Filter conditions |
|   `filter_type` | `string` | Yes | Filter type **Example value**: "number" |
|   `compare_type` | `string` | No | Comparison type **Example value**: "less" |
|   `expected` | `string[]` | Yes | Filter parameter **Example value**: 6 | ### Request body example

{
    "col": "E",
    "condition": {
        "filter_type": "number",
        "compare_type": "less",
        "expected": [
            "6"
        ]
    }
}

```
curl --location --request PUT 'https://open.larksuite.com/open-apis/sheets/v3/spreadsheets/shtcnSUVpFeJ7Q2yVN9/sheets/6e2914/filter' \
--header 'Authorization: Bearer t-719d578dc63f6bd37e30cdb0394798a709070fec' \
--header 'Content-Type: application/json' \
--data-raw '{
    "col": "G",
    "condition": {
        "filter_type": "text",
        "compare_type": "beginsWith",
        "expected": ["a"]
    }
}'
```

## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {

    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 200 | 1310234 | Marshal Error | Internal service error. For more information, contact support. |
| 400 | 1310235 | Retry Later | Please try again later. |
| 500 | 1315203 | Server Error | Internal service error. For more information, contact support. |
| 500 | 1315201 | Server Error | Internal service error. For more information, contact support. |
| 400 | 1310226 | Excess Limit | Exceeds the limit |
| 400 | 1310215 | Sheet Id Not Found | Check the sheet_id. |
| 400 | 1310214 | SpreadSheet Not Found | Check the spreadsheet token. |
| 400 | 1310218 | Locked Cell | Check if the action targets a protected range. |
| 400 | 1310213 | Permission Fail | No permission |
| 400 | 1310206 | Empty Sheet Id | Check the sheet_id |
| 400 | 1310229 | Wrong URL Param | Check the url parameter. |
| 400 | 1310202 | Wrong Range | Invalid range |
| 400 | 1310236 | Wrong Filter Value | Check the filter conditions. |
| 400 | 1310211 | Wrong Sheet Id | Check the sheet_id. |
| 400 | 1310242 | In Mix state | Retey Later |
| 400 | 1310249 | Spreadsheet Deleted | Restore the form and try again |
