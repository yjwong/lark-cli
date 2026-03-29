---
title: "Create Filter View"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/sheets-v3/spreadsheet-sheet-filter_view/create"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/sheets/v3/spreadsheets/:spreadsheet_token/sheets/:sheet_id/filter_views"
service: "sheets-v3"
resource: "spreadsheet-sheet-filter_view"
section: "Docs"
scopes:
  - "drive:drive"
  - "sheets:spreadsheet"
updated: "1664448897000"
---

# Create filter view

::: note
For range settings, see: User guide for using filter conditions in the filter view

This API is used to create a filter view based on passed in parameters. The ID and name fields are optional and default values are generated if none are provided. The range field is required. IDs are 10 characters long and composed of numbers and upper and lowercase English letters. Names can't exceed 100 characters. A single sheet can have up to 150 filter views.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/sheets/v3/spreadsheets/:spreadsheet_token/sheets/:sheet_id/filter_views |
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
| `sheet_id` | `string` | Sheet ID **Example value**: "0b**12" | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `filter_view_id` | `string` | No | Filter view ID **Example value**: "pH9hbVcCXA" |
| `filter_view_name` | `string` | No | Filter view name **Example value**: "Filter view 1" |
| `range` | `string` | No | Filter view range **Example value**: "0b**12!C1:H14" | ### Request body example

{
    "filter_view_id": "pH9hbVcCXA",
    "filter_view_name": "Filter view 1",
    "range": "0b**12!C1:H14"
}

```
curl --location --request POST 'https://open.larksuite.com/open-apis/sheets/v3/spreadsheets/shtcnspY6YDVPxWjhG09Qxxxxxx/sheets/phwh0X/filter_views' \
--header 'Authorization: Bearer t-2602ac1d050a2a308ab8a98639d25cbaaaf26c9f' \
--header 'Content-Type: application/json' \
--data-raw '{
    
    "range": "phwh0X!C1:H14"
}'
```

## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `filter_view` | `filter_view` | ID, name, and range of the created filter view |
|     `filter_view_id` | `string` | Filter view ID |
|     `filter_view_name` | `string` | Filter view name |
|     `range` | `string` | Filter view range | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
        "filter_view": {
            "filter_view_id": "pH9hbVcCXA",
            "filter_view_name": "Filter view 1",
            "range": "0b**12!C1:H14"
        }
    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 500 | 1315201 | Server Error | Internal service error. For more information, contact support. |
| 500 | 1315203 | Server Error | Internal service error. For more information, contact support. |
| 400 | 1310217 | Too Many Request | Check if requests have been sent too frequently. |
| 400 | 1310235 | Retry Later | Please try again later. |
| 400 | 1310214 | SpreadSheet Not Found | Check the spreadsheet token. |
| 400 | 1310215 | Sheet Id Not Found | Check the sheet_id. |
| 400 | 1310223 | Col Id Not Found | Column ID not found. It may have been entered incorrectly or exceed the spreadsheet's column range. |
| 400 | 1310213 | Permission Fail | No permission |
| 400 | 1310204 | Wrong Request Body | Check the request body parameter. |
| 400 | 1310229 | Wrong URL Param | Check the url parameter. |
| 400 | 1310202 | Wrong Range | Invalid range |
| 400 | 1310226 | Excess Limit | Exceeds the limit |
| 400 | 1310211 | Wrong Sheet Id | Check the sheet_id. |
| 400 | 1310232 | Wrong Style | Invalid style, such as color or font |
| 400 | 1310218 | Locked Cell | Check if the action targets a protected range. |
| 400 | 1310237 | Wrong Filter View Id | Invalid filter view ID |
| 400 | 1310238 | Exist Filter View Id | This filter view ID already exists. |
| 400 | 1310239 | Wrong Filter View Name | Invalid filter view name |
| 400 | 1310240 | Exist Filter View Name | This filter view name already exists. |
| 400 | 1310241 | Filter View Excess | The number of filter views exceeds the limit. |
| 400 | 1310242 | In Mix state | Retey Later |
| 400 | 1310249 | Spreadsheet Deleted | Restore the form and try again |
