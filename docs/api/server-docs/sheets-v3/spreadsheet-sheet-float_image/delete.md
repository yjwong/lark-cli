---
title: "Delete Floating Image"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/sheets-v3/spreadsheet-sheet-float_image/delete"
method: "DELETE"
api_path: "https://open.larksuite.com/open-apis/sheets/v3/spreadsheets/:spreadsheet_token/sheets/:sheet_id/float_images/:float_image_id"
service: "sheets-v3"
resource: "spreadsheet-sheet-float_image"
section: "Docs"
scopes:
  - "drive:drive"
  - "sheets:spreadsheet"
updated: "1664448898000"
---

# Delete a floating image

This API is used to delete the floating image with the specified float_image_id.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/sheets/v3/spreadsheets/:spreadsheet_token/sheets/:sheet_id/float_images/:float_image_id |
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
| `float_image_id` | `string` | Floating image ID **Example value**: "ye06SS14ph" | ```
curl --location --request DELETE 'https://open.larksuite.com/open-apis/sheets/v3/spreadsheets/shtbchuIXPxjaYxsZzQxBqxxxxx/sheets/ea131a/float_images/ye06SS14pr' \
--header 'Authorization: Bearer t-384c15ba0644b82caecec91553386563c814c4b9' \
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
| 400 | 1310243 | Wrong Float Image Id | Check the floating image ID. |
| 400 | 1310229 | Wrong URL Param | Check the url parameter. |
| 400 | 1310213 | Permission Fail | No permission |
| 400 | 1310217 | Too Many Request | Check if requests have been sent too frequently. |
| 400 | 1310214 | SpreadSheet Not Found | Check the spreadsheet token. |
| 400 | 1310215 | Sheet Id Not Found | Check the sheet_id. |
| 400 | 1310235 | Retry Later | Please try again later. |
| 500 | 1315201 | Server Error | Internal service error. For more information, contact support. |
| 500 | 1315203 | Server Error | Internal service error. For more information, contact support. |
| 400 | 1310242 | In Mix state | Retey Later |
| 400 | 1310249 | Spreadsheet Deleted | Restore the form and try again |
