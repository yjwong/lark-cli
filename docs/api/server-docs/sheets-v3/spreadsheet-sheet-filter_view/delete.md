---
title: "Delete Filter View"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/sheets-v3/spreadsheet-sheet-filter_view/delete"
method: "DELETE"
api_path: "https://open.larksuite.com/open-apis/sheets/v3/spreadsheets/:spreadsheet_token/sheets/:sheet_id/filter_views/:filter_view_id"
service: "sheets-v3"
resource: "spreadsheet-sheet-filter_view"
section: "Docs"
scopes:
  - "drive:drive"
  - "sheets:spreadsheet"
updated: "1664448897000"
---

# Delete filter view

This API is used to delete the filter view corresponding to the specified ID.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/sheets/v3/spreadsheets/:spreadsheet_token/sheets/:sheet_id/filter_views/:filter_view_id |
| HTTP Method | DELETE |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `drive:drive` `sheets:spreadsheet` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" Learn more about obtaining and using access_token. | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `spreadsheet_token` | `string` | Spreadsheet token **Example value**: "shtcnmBA*****yGehy8" |
| `sheet_id` | `string` | Sheet ID **Example value**: "0b**12" |
| `filter_view_id` | `string` | Filter view ID **Example value**: "pH9hbVcCXA" | ```
curl --location --request DELETE 'https://open.larksuite.com/open-apis/sheets/v3/spreadsheets/shtcnspY6YDVPxWjhG09Qxxxxxx/sheets/phwh0X/filter_views/123456789a' \
--header 'Authorization: Bearer t-2602ac1d050a2a308ab8a98639d25cbaaaf26c9f' \
--header 'Content-Type: application/json'
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
| 400 | 1310237 | Wrong Filter View Id | Invalid filter view ID |
| 400 | 1310229 | Wrong URL Param | Check the url parameter. |
| 400 | 1310213 | Permission Fail | No permission |
| 400 | 1310218 | Locked Cell | Check if the action targets a protected range. |
| 400 | 1310214 | SpreadSheet Not Found | Check the spreadsheet token. |
| 400 | 1310215 | Sheet Id Not Found | Check the sheet_id. |
| 400 | 1310217 | Too Many Request | Check if requests have been sent too frequently. |
| 400 | 1310235 | Retry Later | Please try again later. |
| 500 | 1315201 | Server Error | Internal service error. For more information, contact support. |
| 500 | 1315203 | Server Error | Internal service error. For more information, contact support. |
| 400 | 1310242 | In Mix state | Retey Later |
| 400 | 1310249 | Spreadsheet Deleted | Restore the form and try again |
