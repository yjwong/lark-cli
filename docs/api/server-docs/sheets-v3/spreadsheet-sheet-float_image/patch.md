---
title: "Update Floating Image"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/sheets-v3/spreadsheet-sheet-float_image/patch"
method: "PATCH"
api_path: "https://open.larksuite.com/open-apis/sheets/v3/spreadsheets/:spreadsheet_token/sheets/:sheet_id/float_images/:float_image_id"
service: "sheets-v3"
resource: "spreadsheet-sheet-float_image"
section: "Docs"
scopes:
  - "drive:drive"
  - "sheets:spreadsheet"
updated: "1664448898000"
---

# Update a floating image

::: note
For information about how to update a floating image, see Floating image guide.

This API is used to update the position and size parameters of an existing floating image, including range, width, height, offset_x, and offset_y. This operation cannot update float_image_id or float_image_token.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/sheets/v3/spreadsheets/:spreadsheet_token/sheets/:sheet_id/float_images/:float_image_id |
| HTTP Method | PATCH |
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
| `float_image_id` | `string` | Floating image ID **Example value**: "ye06SS14ph" | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `float_image_token` | `string` | No | Floating image token. This field is required for image creation, but not for updates. The token is obtained when you upload an image to a spreadsheet. **Example value**: "boxbcbQsaSqIXsxxxxx1HCPJFbh" |
| `range` | `string` | No | The position of the cell where the top-left corner of the floating image is located. You can only specify a single cell. **Example value**: "0b**12!A1:A1" |
| `width` | `float` | No | Floating image width, which is greater than or equal to 20 px **Example value**: 100 |
| `height` | `float` | No | Floating image height, which is greater than or equal to 20 px **Example value**: 100 |
| `offset_x` | `float` | No | The horizontal offset of the top-left corner of a floating image from the top-left corner of the cell where it is located, which must be greater than or equal to 0 and less than the height of the cell. **Example value**: 0 |
| `offset_y` | `float` | No | The vertical offset of the top-left corner of a floating image from the top-left corner of the cell where it is located, which must be greater than or equal to 0 and less than the height of the cell. **Example value**: 0 | ### Request body example

{
    "float_image_token": "boxbcbQsaSqIXsxxxxx1HCPJFbh",
    "range": "0b**12!A1:A1",
    "width": 100,
    "height": 100,
    "offset_x": 0,
    "offset_y": 0
}

```
curl --location --request PATCH 'https://open.larksuite.com/open-apis/sheets/v3/spreadsheets/shtbchuIXPxjaYxsZzQxBqPxxxxx/sheets/ea131a/float_images/ye06SS14pr' \
--header 'Authorization: Bearer t-384c15ba0644b82caecec91553386563c814c4b9' \
--header 'Content-Type: application/json' \
--data-raw '{
    "offset_y": 20,
    "width": 100
}'
```

## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `float_image` | `float_image` | Floating image information |
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
        "float_image": {
            "float_image_id": "ye06SS14ph",
            "float_image_token": "boxbcbQsaSqIXsxxxxx1HCPJFbh",
            "range": "0b**12!A1:A1",
            "width": 100,
            "height": 100,
            "offset_x": 0,
            "offset_y": 0
        }
    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 500 | 1315201 | Server Error | Internal service error. For more information, contact support. |
| 500 | 1315203 | Server Error | Internal service error. For more information, contact support. |
| 400 | 1310235 | Retry Later | Please try again later. |
| 400 | 1310217 | Too Many Request | Check if requests have been sent too frequently. |
| 400 | 1310226 | Excess Limit | Exceeds the limit |
| 400 | 1310215 | Sheet Id Not Found | Check the sheet_id. |
| 400 | 1310214 | SpreadSheet Not Found | Check the spreadsheet token. |
| 400 | 1310218 | Locked Cell | Check if the action targets a protected range. |
| 400 | 1310213 | Permission Fail | No permission |
| 400 | 1310229 | Wrong URL Param | Check the url parameter. |
| 400 | 1310202 | Wrong Range | Invalid range |
| 400 | 1310246 | Wrong Float Image Value | Incorrect image width/height or offset. |
| 400 | 1310243 | Wrong Float Image Id | Check the floating image ID. |
| 400 | 1310211 | Wrong Sheet Id | Check the sheet_id. |
| 400 | 1310242 | In Mix state | Retey Later |
| 400 | 1310249 | Spreadsheet Deleted | Restore the form and try again |
