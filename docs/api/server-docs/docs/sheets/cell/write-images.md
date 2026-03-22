---
title: "Write Images"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uUDNxYjL1QTM24SN0EjN"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/sheets/v2/spreadsheets/:spreadsheetToken/values_image"
section: "Docs"
scopes:
  - "drive:drive"
  - "sheets:spreadsheet"
updated: "1647001418000"
---

# Write a image

This API is used to write an image to the specified cell based on the spreadsheetToken and range.

##  Request
| HTTP URL | https://open.larksuite.com/open-apis/sheets/v2/spreadsheets/:spreadsheetToken/values_image |
| --- | --- |
| HTTP Method | POST |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `drive:drive` `sheets:spreadsheet` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `user_access_token`  or  `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" Learn more about how to obtain and use access_token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ###  Path parameters
|Parameter|Type|Required|Description|
|--|-----|--|----|
|spreadsheetToken|string|Yes|spreadsheet token. For more information about how to obtain the token, see Sheets Developer's Guide.| URL Path|
###  Request body
|Parameter|Type|Required|Description|
|--|-----|--|----|
|range|string|Yes|Query range  range=!: For example: xxxx!A1:D5.For details, see Sheets Developer's Guide. This parameter is limited to one cell here, such as: : xxxx!A1:A1.|
|image|array|Yes|Binary data stream of the image to write.   "PNG",  "JPEG",  "JPG",  "GIF",  "BMP",  "JFIF",  "EXIF",  "TIFF",  "BPG",  "WEBP", and  "HEIC"  formats are supported.|
|name|string|Yes|Name of the image to write|
### Request body example
```json
{ 
    "range": "string", 
    "image": [123,32,42,23],
    "name": "xxx.png"
}
```
### cURL Request example
```
curl --location --request POST 'https://open.larksuite.com/open-apis/sheets/v2/spreadsheets/shtcngNygNfuqhxTBf588jwgWbJ/values_image' \
--header 'Authorization: Bearer t-e346617a4acfc3a11d4ed24dca0d0c0fc8e0067e' \
--header 'Content-Type: application/json' \
--data-raw '{ 
    "range": "Q7PlXT!H7:H7",
    "image": [137,80,78,71,13,10,26,10,0,0,0,13,73,72,68,82,0,0,0,2,0,0,0,1,1,0,0,0,0,220,89,66,39,0,0,0,2,116,82,78,83,0,0,118,147,205,56,0,0,0,10,73,68,65,84,8,215,99,104,0,0,0,130,0,129,221,67,106,244,0,0,0,0,73,69,78,68,174,66,96,130],
    "name": "test.png"
}'
```
##  Response
### Response body
  |Parameter|Type|Description|
|--|-----|--|
|spreadsheetToken|string |spreadsheet   token |
### Response body example
 
```json
{ 
    "code": 0,
    "data": { 
        "spreadsheetToken": "***" 
    }, 
    "msg": "Success" 
}
```
### Error code

For details, refer to Server-side error codes.
