---
title: "Create Floating Image"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/sheets-v3/spreadsheet-sheet-float_image/create"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/sheets/v3/spreadsheets/:spreadsheet_token/sheets/:sheet_id/float_images"
service: "sheets-v3"
resource: "spreadsheet-sheet-float_image"
section: "Docs"
scopes:
  - "drive:drive"
  - "sheets:spreadsheet"
updated: "1664448898000"
---

# Create a floating image

::: note
For information about the settings of a floating image, see Floating image guide.

This API is used to create a floating image based on the parameters passed in. Float_image_token (obtained after uploading an image to a spreadsheet) and range (only one cell) are required parameters. Float_image_id is optional, and a default ID is generated if this field is not specified. IDs are 10 characters long and composed of numbers (0-9) and upper and lowercase English letters. A spreadsheet cannot have duplicate images, and the total of floating images and cell images cannot exceed 4,000. The width and height specify the display width and height of an image, which are optional. If not specified, the image is displayed according to its actual width and height. offset_x and offset_y specify the offset of the image from the top-left corner of the cell where it is located. These two parameters are optional and default to 0.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/sheets/v3/spreadsheets/:spreadsheet_token/sheets/:sheet_id/float_images |
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
| `float_image_id` | `string` | No | Floating image ID **Example value**: "ye06SS14ph" |
| `float_image_token` | `string` | No | Floating image token. This field is required for image creation, but not for updates. The token is obtained when you upload an image to a spreadsheet. **Example value**: "boxbcbQsaSqIXsxxxxx1HCPJFbh" |
| `range` | `string` | No | The position of the cell where the top-left corner of the floating image is located. You can only specify a single cell. **Example value**: "0b**12!A1:A1" |
| `width` | `float` | No | Floating image width, which is greater than or equal to 20 px **Example value**: 100 |
| `height` | `float` | No | Floating image height, which is greater than or equal to 20 px **Example value**: 100 |
| `offset_x` | `float` | No | The horizontal offset of the top-left corner of a floating image from the top-left corner of the cell where it is located, which must be greater than or equal to 0 and less than the height of the cell. **Example value**: 0 |
| `offset_y` | `float` | No | The vertical offset of the top-left corner of a floating image from the top-left corner of the cell where it is located, which must be greater than or equal to 0 and less than the height of the cell. **Example value**: 0 | ### Request body example

{
    "float_image_id": "ye06SS14ph",
    "float_image_token": "boxbcbQsaSqIXsxxxxx1HCPJFbh",
    "range": "0b**12!A1:A1",
    "width": 100,
    "height": 100,
    "offset_x": 0,
    "offset_y": 0
}

```
curl --location --request POST 'https://open.larksuite.com/open-apis/sheets/v3/spreadsheets/shtbchuIXPxjaYxsZzQxBqxxxxx/sheets/ea131a/float_images' \
--header 'Authorization: Bearer t-f148d9ee6b5c07373a2671e795e9e855a6f171f6' \
--header 'Content-Type: application/json' \
--data-raw '{
    "float_image_id": "ye06SS14p9",
   "float_image_token": "boxbcbQEhbFAe0XJvGzkD165FGb",
    "offset_x": 0,
    "offset_y": 0,
    "range": "ea131a!C3:C3"
}'
```

## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `float_image` | `float_image` | Return value of a floating image |
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
| 400 | 1310229 | Wrong URL Param | Check the url parameter. |
| 400 | 1310217 | Too Many Request | Check if requests have been sent too frequently. |
| 400 | 1310235 | Retry Later | Please try again later. |
| 500 | 1315201 | Server Error | Internal service error. For more information, contact support. |
| 400 | 1310214 | SpreadSheet Not Found | Check the spreadsheet token. |
| 400 | 1310213 | Permission Fail | No permission |
| 400 | 1310215 | Sheet Id Not Found | Check the sheet_id. |
| 400 | 1310202 | Wrong Range | Invalid range |
| 400 | 1310245 | Wrong Float Image Token | Check the floating image token. |
| 400 | 1310244 | Exist Float Image Id | Check the floating image ID. |
| 400 | 1310243 | Wrong Float Image Id | Check the floating image ID. |
| 400 | 1310211 | Wrong Sheet Id | Check the sheet_id. |
| 400 | 1310247 | Image Excess | The number of images exceeds the limit. |
| 400 | 1310246 | Wrong Float Image Value | Incorrect image width/height or offset. |
| 400 | 1310218 | Locked Cell | Check if the action targets a protected range. |
| 400 | 1310226 | Excess Limit | Exceeds the limit |
| 400 | 1310242 | In Mix state | Retey Later |
| 500 | 1315203 | Server Error | Internal service error. For more information, contact support. |
| 400 | 1310249 | Spreadsheet Deleted | Restore the form and try again |
