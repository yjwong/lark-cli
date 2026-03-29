---
title: "Batch Set Cell Style"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uAzMzUjLwMzM14CMzMTN"
method: "PUT"
api_path: "https://open.larksuite.com/open-apis/sheets/v2/spreadsheets/:spreadsheetToken/styles_batch_update"
section: "Docs"
scopes:
  - "drive:drive"
  - "sheets:spreadsheet"
updated: "1647001409000"
---

# Set cell styles in batches

This API is used to set the style of multiple cells based on spreadsheetToken, range, and style information. You can act on up to 5,000 rows and 100 columns.
##  Request
| HTTP URL | https://open.larksuite.com/open-apis/sheets/v2/spreadsheets/:spreadsheetToken/styles_batch_update |
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
|spreadsheetToken|string|Yes|Spreadsheet token. For more information about how to obtain the token, see Sheets Developer's Guide| URL PATH.|
###  Request body
|Parameter|Type|Required|Description|
|--|-----|--|----|
|data||Yes|Request data|
|  ∟ranges|array|Yes|Query range, includes the sheetId range and cell range. Four indexing methods are supported. For details, see  Sheets Developer's Guide.|
|  ∟style||Yes|The style to update|
|    ∟font| |No|Font style|
|      ∟bold|bool|No|Bold|
|      ∟italic|bool|No|Italics|
|      ∟fontSize|string|No|Font size: 9 to 36, line spacing fixed at 1.5, for example:10pt/1.5| |    ∟clean|bool|No|Clear font style, default value: false|
|    ∟textDecoration|int|No|Text decoration, 0: default, 1: underline, 2: strikethrough, 3: underline and strikethrough|  
|    ∟formatter|string| No |Number format. For details, see the appendix Number formats supported by sheets .|
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
    "data":[
        {
            "ranges":[
                "string",
                "string"
            ],
            "style":{
                "font":{
                    "bold":true,
                    "italic":true,
                    "fontSize":"10pt/1.5",
                    "clean":false
                },
                "textDecoration":0,
                "formatter":"",
                "hAlign":0,
                "vAlign":0,
                "foreColor":"#000000",
                "backColor":"#21d11f",
                "borderType":"FULL_BORDER",
                "borderColor":"#ff0000",
                "clean":false
            }
        },
        {
            "ranges":[
                "string"
            ],
            "style":{
              ...
            }
        }
    ]
}
```
  
###  cURL Request example
```
curl --location --request PUT 'https://open.larksuite.com/open-apis/sheets/v2/spreadsheets/shtcngNygNfuqhxTBf588jwgWbJ/styles_batch_update' \
--header 'Authorization: Bearer t-e346617a4acfc3a11d4ed24dca0d0c0fc8e0067e' \
--header 'Content-Type: application/json' \
--data-raw '{
    "data":[
        {
            "ranges":[
                "Q7PlXT!C7:E12",
                "Q7PlXT!I20:K27"
            ],
            "style":{
                "font":{
                    "bold":true,
                    "italic":true,
                    "fontSize":"10pt/1.5",
                    "clean":false
                },
                "textDecoration":0,
                "formatter":"",
                "hAlign":0,
                "vAlign":0,
                "foreColor":"#000000",
                "backColor":"#21d11f",
                "borderType":"FULL_BORDER",
                "borderColor":"#ff0000",
                "clean":false
            }
        },
        {
            "ranges":[
                "BzY8T5!A1:C2"
            ],
            "style":{
                "font":{
                    "bold":true,
                    "italic":true,
                    "fontSize":"10pt/1.5",
                    "clean":false
                },
                "textDecoration":0,
                "formatter":"",
                "hAlign":0,
                "vAlign":0,
                "foreColor":"#000000",
                "backColor":"#21d11f",
                "borderType":"FULL_BORDER",
                "borderColor":"#ff0000",
                "clean":false
            }
        }
    ]
}
'
```
  
##  Response
### Response body
  |Parameter|Type|Description|
|--|-----|--|
|spreadsheetToken|string | Spreadsheet token|
|totalUpdatedRows|int|Total rows for which to set the style|
|totalUpdatedColumns|int|Total columns for which to set the style|
|totalUpdatedCells|int|Total cells for which to set the style|
|revision|int|Version number of sheet |
|responses||Range and number of rows and columns for which to set the cell style for each range|
|  ∟spreadsheetToken|string |Spreadsheet token|
|  ∟updatedRange|string |Range for which to set the style|
|  ∟updatedRows|int|Number of rows for which to set the style|
|  ∟updatedColumns|int|Number of columns for which to set the style|
|  ∟updatedCells|int|Number of cells for which to set the style|
### Response body example
```json
    {
    "code": 0,
    "msg": "Success",
    "data":{
        "spreadsheetToken": "string",
        "totalUpdatedCells": 0,
        "totalUpdatedColumns": 0,
        "totalUpdatedRows": 0,
        "revision": 0,
       "responses": [
          {
            "spreadsheetToken": "string",
            "updatedRange": "string",
            "updatedRows": 0,
            "updatedColumns": 0,
            "updatedCells": 0
          }
        ]
}
```
### Error code

For details, refer to Server-side error codes.
