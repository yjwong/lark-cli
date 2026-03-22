---
title: "Prepend Data"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uIjMzUjLyIzM14iMyMTN"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/sheets/v2/spreadsheets/:spreadsheetToken/values_prepend"
section: "Docs"
scopes:
  - "drive:drive"
  - "sheets:spreadsheet"
updated: "1647001388000"
---

# Insert data

This API is used to write data to add rows and corresponding data before a specified range based on spreadsheetToken and range. This is equivalent to an array insert action. You can write up to 5,000 rows and 100 columns of data. A single cell can't exceed 50,000 characters.
##  Request
| HTTP URL | https://open.larksuite.com/open-apis/sheets/v2/spreadsheets/:spreadsheetToken/values_prepend |
| --- | --- |
| HTTP Method | POST |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `drive:drive` `sheets:spreadsheet` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `user_access_token` or  `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" Learn more about how to obtain and use access_token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ###  Path parameters
|Parameter|Type|Required|Description|
|--|-----|--|----|
|spreadsheetToken|string|Yes|Sheet token. For more information about how to obtain the token, see  Sheets Developer's Guide.| URL PATH|
###  Request body
|Parameter|Type|Required|Description|
|--|-----|--|----|
|valueRange||Yes|Value and range|
|  ∟range|string|Yes|⁣The insert range, includes the sheetId ranges and cell ranges. Three indexing methods are supported. For details, see  Sheets Developer's Guide. The range indicated by range needs to be greater than or equal to the range occupied by values.|
|  ∟values|array>|Yes|Values to write. To write formulas, hyperlinks, emails, and @mention users, see the appendix Data types that can be written to sheet.|
### Request body example
```json
{
"valueRange":{
  "range": "string",
  "values": [
    [
      "string",1 ,"http://www.xx.com"
    ]
  ]
  }
}
```
  
### cURL Request example
```
curl --location --request POST 'https://open.larksuite.com/open-apis/sheets/v2/spreadsheets/shtcngNygNfuqhxTBf588jwgWbJ/values_prepend' \
--header 'Authorization: Bearer t-e346617a4acfc3a11d4ed24dca0d0c0fc8e0067e' \
--header 'Content-Type: application/json' \
--data-raw '{
"valueRange":{
  "range": "Q7PlXT!C6:F9",
  "values": [
    [
      "a",1,"http://www.xx.com",12
    ],
    [
      "b",2,8,"me@HelloWorld.com"
    ],
    [
      "c",3,2,6
    ],
    [
      "d",4,6,"@Jack"
    ]
  ]
  }
}'
```
  
##  Response
### Response body
  |Parameter|Type|Description|
|--|-----|--|
|spreadsheetToken|string|Spreadsheet token|
|tableRange|string|Range to write|
|revision|int|Version number of sheet |
|updates||Range and number of rows and columns for data insertion|
|  ∟spreadsheetToken|string|Spreadsheet token|
|  ∟updatedRange|string|Range to write|
|  ∟updatedRows|int|Number of rows to write|
|  ∟updatedColumns|int|Number of columns to write|
|  ∟updatedCells|int|Total cells to write|
|  ∟revision|int|Version number of sheet |
###  Response body example  
```json
{
    "code": 0,
    "msg": "Success",
    "data": {
        "revision": 0,
        "spreadsheetToken": "***",
        "tableRange": "***",
        "updates": {
            "spreadsheetToken": "***",
            "updatedRange": "***",
            "updatedRows": 0,
            "updatedColumns": 0,
            "updatedCells": 0,
            "revision": 0
        }
    }
}
```
### Error code

For details, refer to Server-side error codes.
