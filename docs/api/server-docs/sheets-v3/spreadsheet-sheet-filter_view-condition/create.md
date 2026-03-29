---
title: "Create Filter Condition"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/sheets-v3/spreadsheet-sheet-filter_view-condition/create"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/sheets/v3/spreadsheets/:spreadsheet_token/sheets/:sheet_id/filter_views/:filter_view_id/conditions"
service: "sheets-v3"
resource: "spreadsheet-sheet-filter_view-condition"
section: "Docs"
scopes:
  - "drive:drive"
  - "sheets:spreadsheet"
updated: "1664448898000"
---

# Create a filter condition

::: note
For filter conditions, see User guide for using filter conditions in the filter view

This API is used to create filter conditions for a column in the range of a filter view.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/sheets/v3/spreadsheets/:spreadsheet_token/sheets/:sheet_id/filter_views/:filter_view_id/conditions |
| HTTP Method | POST |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `drive:drive` `sheets:spreadsheet` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" Learn more about obtaining and using access_token. |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `spreadsheet_token` | `string` | Spreadsheet token **Example value**: "shtcnmBA*****yGehy8" |
| `sheet_id` | `string` | Sheet ID **Example value**: "0b**12" |
| `filter_view_id` | `string` | Filter view ID **Example value**: "pH9hbVcCXA" | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `condition_id` | `string` | No | Set filter condition column, using a letter designation **Example value**: "E" |
| `filter_type` | `string` | No | Filter type **Example value**: "number" |
| `compare_type` | `string` | No | Comparison type **Example value**: "less" |
| `expected` | `string[]` | No | Filter parameter **Example value**: 6 | ### Request body example

{
    "condition_id": "E",
    "filter_type": "number",
    "compare_type": "less",
    "expected": [
        "6"
    ]
}

```
curl --location --request POST 'https://open.larksuite.com/open-apis/sheets/v3/spreadsheets/shtcnspY6YDVPxWjhG09Qxxxxxx/sheets/phwh0X/filter_views/1234567890/conditions' \
--header 'Authorization: Bearer t-40cdeb051222f889f4229de856517992260aa850' \
--header 'Content-Type: application/json' \
--data-raw '{
    "condition_id": "G",
    "filter_type": "text",
    "compare_type": "beginsWith",
    "expected": [
        "a"
    ]
}'
```

## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `condition` | `filter_view_condition` | Created filter condition |
|     `condition_id` | `string` | Set filter condition column, using a letter designation |
|     `filter_type` | `string` | Filter type |
|     `compare_type` | `string` | Comparison type |
|     `expected` | `string[]` | Filter parameter | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
        "condition": {
            "condition_id": "E",
            "filter_type": "number",
            "compare_type": "less",
            "expected": [
                "6"
            ]
        }
    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 1310211 | Wrong Sheet Id | Check the sheet_id. |
| 400 | 1310237 | Wrong Filter View Id | Invalid filter view ID |
| 400 | 1310236 | Wrong Filter Value | Check the filter conditions. |
| 400 | 1310202 | Wrong Range | Invalid range |
| 400 | 1310229 | Wrong URL Param | Check the url parameter. |
| 400 | 1310204 | Wrong Request Body | Check the request body parameter. |
| 400 | 1310213 | Permission Fail | No permission |
| 400 | 1310218 | Locked Cell | Check if the action targets a protected range. |
| 400 | 1310214 | SpreadSheet Not Found | Check the spreadsheet token. |
| 400 | 1310215 | Sheet Id Not Found | Check the sheet_id. |
| 400 | 1310223 | Col Id Not Found | Column ID not found. It may have been entered incorrectly or exceed the spreadsheet's column range. |
| 400 | 1310217 | Too Many Request | Check if requests have been sent too frequently. |
| 400 | 1310235 | Retry Later | Please try again later. |
| 500 | 1315201 | Server Error | Internal service error. For more information, contact support. |
| 500 | 1315203 | Server Error | Internal service error. For more information, contact support. |
| 400 | 1310242 | In Mix state | Retey Later |
| 400 | 1310232 | Wrong Style | Invalid style, such as color or font |
| 400 | 1310249 | Spreadsheet Deleted | Restore the form and try again |
