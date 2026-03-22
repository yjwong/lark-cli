---
title: "Set Cell Style"
url: "https://open.larksuite.com/document/ukTMukTMukTM/ukjMzUjL5IzM14SOyMTN"
method: "PUT"
api_path: "https://open.larksuite.com/open-apis/sheets/v2/spreadsheets/:spreadsheetToken/style"
section: "Docs"
scopes:
  - "drive:drive"
  - "sheets:spreadsheet"
updated: "1647001406000"
---

# Set cell style

This API is used to set the cell style based on spreadsheetToken, range, and style information. You can act on up to 5,000 rows and 100 columns.

##  Request
| HTTP URL | https://open.larksuite.com/open-apis/sheets/v2/spreadsheets/:spreadsheetToken/style |
| --- | --- |
| HTTP Method | PUT |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `drive:drive` `sheets:spreadsheet` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `user_access_token`  or  `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" Learn more about how to obtain and use access_token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ###  Path parameters
|Parameter|Type|Required|Description|
|--|-----|--|----|
|spreadsheetToken|string|Yes| Spreadsheet token. For more information about how to obtain the token, see  Sheets Developer's Guide.|
###  Request body
**Request parameter description** :  
|Parameter|Type|Required|Description|
|--|-----|--|----|
|appendStyle||Yes|Set cell style
|  ∟range|string|Yes|Query range, includes the sheetId range and cell range. Four indexing methods are supported. For details, see Sheets Developer's Guide.|
|  ∟style||Yes|The style to update| 
|    ∟font| |No|Font style|
|      ∟bold|bool|No|Bold|
|      ∟italic|bool|No|Italics| 
|      ∟fontSize|string|No|Font size: 9 to 36, line spacing fixed at 1.5, for example: 10pt/1.5| 
|      ∟clean|bool|No|Clear font formatting, default: false| Request body |
|    ∟textDecoration|int|No|Text decoration, 0: default, 1: underline, 2: strikethrough, 3: underline and strikethrough|
|    ∟formatter|string|No|Number format. For details, see the appendix  Number formats supported by sheet .| 
|    ∟hAlign|int|No|Horizontal align, 0: Align left, 1: Align center, 2: Align right | 
|    ∟vAlign|int|No|Vertical align, 0: Align top, 1: Align center, 2: Align bottom| 
|    ∟foreColor|string|No|Text color| 
|    ∟backColor|string|No|Background color|
|    ∟borderType|string|No|Border type, values: "FULL_BORDER", "OUTER_BORDER", "INNER_BORDER", "NO_BORDER", "LEFT_BORDER", "RIGHT_BORDER", "TOP_BORDER", and "BOTTOM_BORDER"| 
|    ∟borderColor|string|No|Border color| 
|    ∟clean|bool|No|Whether to clear all formatting, default: false| 
### Request body example

  ```json
{
	"appendStyle":{
       "range": "string",
       "style":{
            "font":{
                "bold":true,
                "italic":true,
                "fontSize":"10pt/1.5",
                "clean":false  
                },    
            "textDecoration":0,
            "formatter":"",
            "hAlign": 0 , 
            "vAlign":0,   
            "foreColor":"#000000",
            "backColor":"#21d11f",
            "borderType":"FULL_BORDER",
            "borderColor": "#ff0000",
            "clean": false 
            }
        }
}
```
### cURL Request example
```
curl --location --request PUT 'https://open.larksuite.com/open-apis/sheets/v2/spreadsheets/shtcngNygNfuqhxTBf588jwgWbJ/style' \
--header 'Authorization: Bearer t-e346617a4acfc3a11d4ed24dca0d0c0fc8e0067e' \
--header 'Content-Type: application/json' \
--data-raw '{
  "appendStyle":{
     "range": "BzY8T5!A3:C4",
     "style":{
          "font":{
              "bold":true,
              "italic":true,
              "fontSize":"10pt/1.5",
              "clean":false  
              },    
          "textDecoration":0,
          "formatter":"",
          "hAlign": 0 , 
          "vAlign":0,   
          "foreColor":"#000000",
          "backColor":"#21d11f",
          "borderType":"FULL_BORDER",
          "borderColor": "#ff0000",
          "clean": false 
          }
      }
}'
```
##  Response
### Response body
  |Parameter|Type|Description|
|--|-----|--|
|spreadsheetToken|string |spreadsheet   token|
|updatedRange|string |Range for which to set the style|
|updatedRows|int|Number of rows for which to set the style|
|updatedColumns|int|Number of columns for which to set the style|
|updatedCells|int|Total cells for which to set the style|
|revision|int|Version number of sheet | ###  Response body example 
```json
{
    "code": 0,
    "msg": "Success",
    "data":{
      "spreadsheetToken": "string",
      "updatedRange": "string",
      "updatedRows": 0,
      "updatedColumns": 0,
      "updatedCells": 0,
      "revision": 0
	}
}
```
  
### Error code

For details, refer to Server-side error codes.
