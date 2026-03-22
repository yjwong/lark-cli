---
title: "Move Dimension"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/sheets-v3/spreadsheet-sheet/move_dimension"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/sheets/v3/spreadsheets/:spreadsheet_token/sheets/:sheet_id/move_dimension"
service: "sheets-v3"
resource: "spreadsheet-sheet"
section: "Docs"
scopes:
  - "drive:drive"
  - "sheets:spreadsheet"
updated: "1664448898000"
---

# Rows or columns to move

This API is used to move rows and columns. After rows and columns are moved to a target location, the rows and columns already at the target location are shifted right or down.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/sheets/v3/spreadsheets/:spreadsheet_token/sheets/:sheet_id/move_dimension |
| HTTP Method | POST |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `drive:drive` `sheets:spreadsheet` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" Learn more about obtaining and using access_token. |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `spreadsheet_token` | `string` | Sheet token **Example value**: "shtcnmBA\*****yGehy8" |
| `sheet_id` | `string` | Sheet ID **Example value**: "0b\**12" | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `source` | `dimension` | No | Source location parameter |
|   `major_dimension` | `string` | No | Row or column, values: ROWS or COLUMNS **Example value**: "ROWS" |
|   `start_index` | `int` | No | Start row or column number **Example value**: 0 |
|   `end_index` | `int` | No | End row or column number **Example value**: 1 |
| `destination_index` | `int` | No | Row or column number of the target location **Example value**: 4 | ### Request body example

{
    "source": {
        "major_dimension": "ROWS",
        "start_index": 0,
        "end_index": 1
    },
    "destination_index": 4
}

```
curl --location --request POST 'https://open.larksuite.com/open-apis/sheets/v3/spreadsheets/shtcnSUVpFeJ7Q2ycG1cHeSc/sheets/6e2914/move_dimension' \
--header 'Authorization: Bearer t-719d578dc63f6bd37e30cdb0394798a709070fec' \
--header 'Content-Type: application/json' \
--data-raw '{
    "source": {
        "major_dimension": "COLUMNS",
        "start_index": 1,
        "end_index": 3
    },
    "destination_index": 0
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
| 400 | 1310211 | Wrong Sheet Id | Check the sheet_id. |
| 400 | 1310202 | Wrong Range | Invalid range |
| 400 | 1310229 | Wrong URL Param | Check the url parameter. |
| 400 | 1310204 | Wrong Request Body | Check the request body parameter. |
| 400 | 1310206 | Empty Sheet Id | Check the sheet_id |
| 400 | 1310218 | Locked Cell | Check if the action targets a protected range. |
| 400 | 1310213 | Permission Fail | No permission |
| 400 | 1310214 | SpreadSheet Not Found | Check the spreadsheet token. |
| 400 | 1310215 | Sheet Id Not Found | Check the sheet_id. |
| 400 | 1310216 | Empty Value | Check the request body |
| 400 | 1310217 | Too Many Request | Check if requests have been sent too frequently. |
| 400 | 1310219 | Cell Excess | Check if it exceeds limit |
| 400 | 1310235 | Retry Later | Please try again later. |
| 500 | 1315201 | Server Error | Internal service error. For more information, contact support. |
| 500 | 1315203 | Server Error | Internal service error. For more information, contact support. |
| 200 | 1310234 | Marshal Error | Internal service error. For more information, contact support. |
| 400 | 1310242 | In Mix state | Retey Later |
| 400 | 1310249 | Spreadsheet Deleted | Restore the form and try again |
