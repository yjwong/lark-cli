---
title: "Query Floating Image"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/sheets-v3/spreadsheet-sheet-float_image/query"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/sheets/v3/spreadsheets/:spreadsheet_token/sheets/:sheet_id/float_images/query"
service: "sheets-v3"
resource: "spreadsheet-sheet-float_image"
section: "Docs"
scopes:
  - "drive:drive"
  - "drive:drive:readonly"
  - "sheets:spreadsheet"
  - "sheets:spreadsheet:readonly"
updated: "1664448898000"
---

# Query floating images

::: note
For information about floating images, see Floating image guide

This API returns information on all floating images in a sheet.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/sheets/v3/spreadsheets/:spreadsheet_token/sheets/:sheet_id/float_images/query |
| HTTP Method | GET |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `drive:drive` `drive:drive:readonly` `sheets:spreadsheet` `sheets:spreadsheet:readonly` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" Learn more about obtaining and using access_token. | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `spreadsheet_token` | `string` | Spreadsheet token **Example value**: "shtcnmBA*****yGehy8" |
| `sheet_id` | `string` | Sheet ID **Example value**: "0b**12" | ```
curl --location --request GET 'https://open.larksuite.com/open-apis/sheets/v3/spreadsheets/shtbchuIXPxjaYxsZzQxBqxxxxx/sheets/ea131a/float_images/query' \
--header 'Authorization: Bearer t-384c15ba0644b82caecec91553386563c814c4b9' \
--header 'Content-Type: application/json'
```

## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `items` | `float_image[]` | Information on all floating images in the sheet |
|     `float_image_id` | `string` | Floating image ID |
|     `float_image_token` | `string` | Floating image token. This field is required for image creation, but not for updates. The token is obtained when you upload an image to a spreadsheet. |
|     `range` | `string` | The position of the cell where the top-left corner of the floating image is located. You can only specify a single cell. |
|     `width` | `float` | Floating image width, which is greater than or equal to 20 px |
|     `height` | `float` | Floating image height, which is greater than or equal to 20 px |
|     `offset_x` | `float` | The horizontal offset of the top-left corner of a floating image from the top-left corner of the cell where it is located, which must be greater than or equal to 0 and less than the height of the cell. |
|     `offset_y` | `float` | The vertical offset of the top-left corner of a floating image from the top-left corner of the cell where it is located, which must be greater than or equal to 0 and less than the height of the cell. | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
        "items": [
            {
                "float_image_id": "ye06SS14ph",
                "float_image_token": "boxbcbQsaSqIXsxxxxx1HCPJFbh",
                "range": "0b**12!A1:A1",
                "width": 100,
                "height": 100,
                "offset_x": 0,
                "offset_y": 0
            }
        ]
    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
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
