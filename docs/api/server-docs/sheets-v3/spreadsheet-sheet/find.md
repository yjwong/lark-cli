---
title: "Find cells"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/sheets-v3/spreadsheet-sheet/find"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/sheets/v3/spreadsheets/:spreadsheet_token/sheets/:sheet_id/find"
service: "sheets-v3"
resource: "spreadsheet-sheet"
section: "Docs"
scopes:
  - "drive:drive"
  - "drive:drive:readonly"
  - "sheets:spreadsheet"
  - "sheets:spreadsheet:readonly"
updated: "1664448898000"
---

# Search

Search cell locations that match specified conditions in a certain range of the sheet. The range and find fields are required in the request body.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/sheets/v3/spreadsheets/:spreadsheet_token/sheets/:sheet_id/find |
| HTTP Method | POST |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `drive:drive` `drive:drive:readonly` `sheets:spreadsheet` `sheets:spreadsheet:readonly` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" Learn more about obtaining and using access_token. |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `spreadsheet_token` | `string` | Sheet token **Example value**: "shtcnmBA*****yGehy8" |
| `sheet_id` | `string` | Sheet ID **Example value**: "0b**12" | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `find_condition` | `find_condition` | Yes | Search conditions |
|   `range` | `string` | Yes | Search range **Example value**: "0b**12!A1:H10" |
|   `match_case` | `boolean` | No | Whether it is case insensitive **Example value**: true |
|   `match_entire_cell` | `boolean` | No | Whether to match the entire cell **Example value**: false |
|   `search_by_regex` | `boolean` | No | Whether it is regular expression matching **Example value**: false |
|   `include_formulas` | `boolean` | No | Whether to search formula content **Example value**: false |
| `find` | `string` | Yes | Found string **Example value**: "hello" | ### Request body example

{
    "find_condition": {
        "range": "0b**12!A1:H10",
        "match_case": true,
        "match_entire_cell": false,
        "search_by_regex": false,
        "include_formulas": false
    },
    "find": "hello"
}

## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `find_result` | `find_replace_result` | Search return information that matches conditions |
|     `matched_cells` | `string[]` | Array of cells that meet the search conditions, not including formulas, such as: ["A1", "A2"...] |
|     `matched_formula_cells` | `string[]` | Array of cells that meet the search conditions including formulas, such as: ["B3", "H7"...] |
|     `rows_count` | `int` | Total rows that meet search conditions | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
        "find_result": {
            "matched_cells": [
                "A1"
            ],
            "matched_formula_cells": [
                "B3"
            ],
            "rows_count": 2
        }
    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 1310242 | In Mix state | Retey Later |
| 400 | 1310211 | Wrong Sheet Id | Check the sheet_id. |
| 400 | 1310248 | Wrong Regular Expression | Check Regular Expression |
| 400 | 1310202 | Wrong Range | Invalid range |
| 400 | 1310229 | Wrong URL Param | Check the url parameter. |
| 400 | 1310204 | Wrong Request Body | Check the request body parameter. |
| 400 | 1310206 | Empty Sheet Id | Check the sheet_id |
| 400 | 1310213 | Permission Fail | No permission |
| 400 | 1310218 | Locked Cell | Check if the action targets a protected range. |
| 400 | 1310214 | SpreadSheet Not Found | Check the spreadsheet token. |
| 400 | 1310215 | Sheet Id Not Found | Check the sheet_id. |
| 400 | 1310226 | Excess Limit | Exceeds the limit |
| 400 | 1310217 | Too Many Request | Check if requests have been sent too frequently. |
| 400 | 1310235 | Retry Later | Please try again later. |
| 500 | 1315201 | Server Error | Internal service error. For more information, contact support. |
| 500 | 1315203 | Server Error | Internal service error. For more information, contact support. |
| 400 | 1310249 | Spreadsheet Deleted | Restore the form and try again |
