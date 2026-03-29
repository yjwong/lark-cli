---
title: "Append Data"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uMjMzUjLzIzM14yMyMTN"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/sheets/v2/spreadsheets/:spreadsheetToken/values_append"
section: "Docs"
scopes:
  - "drive:drive"
  - "sheets:spreadsheet"
updated: "1647001391000"
---

# Append data

This API is used to append data by overwriting blank rows or adding new rows based on spreadsheetToken and range. Blank row: By default, when the first cell of the row is empty, the row will be considered blank. You can write up to 5,000 rows and 100 columns of data at once. A single cell can't exceed 50,000 characters.

##  Request
| HTTP URL | https://open.larksuite.com/open-apis/sheets/v2/spreadsheets/:spreadsheetToken/values_append |
| --- | --- |
| HTTP Method | POST |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `drive:drive` `sheets:spreadsheet` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `user_access_token`  or  `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" Learn more about how to obtain and use access_token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ###  Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| spreadsheetToken | string | Spreadsheet token. For more information about how to obtain the token, see Sheets Developer's Guide | ### Query parameters  
|Parameter|Type|Required|Description|
|--|-----|--|----|
|insertDataOption|string|No|When a blank row is encountered, the default action is OVERWRITE. If there are not enough blank rows to fit the rows of data to append, overwrite data append is performed. Select INSERT_ROWS to insert a sufficient number of rows and then append.| ### Request body 
|Parameter|Type|Required|Description|
|--|-----|--|----|
|valueRange||Yes|Value and range|
|  ∟range|string|Yes|The query range, includes the sheetId ranges and cell ranges. Three indexing methods are supported. For details, see Sheets Developer's Guide. The range indicated by range needs to be greater than or equal to the range occupied by values.|
|  ∟values|array>|Yes|Values to write. To write formulas, hyperlinks, emails, and @mention users, see the appendix Data types that can be written to sheet.| ###  Request body example  
```json
{
  "valueRange": {
    "range": "string",
    "values": [
      [
        "string",
        1,
        "https://www.xxx.com"
      ]
    ]
  }
}
```
  
  ### cURL Request example
```
curl --location --request POST 'https://open.larksuite.com/open-apis/sheets/v2/spreadsheets/shtcngNygNfuqhxTBf588jwgWbJ/values_append' \
--header 'Authorization: Bearer t-e346617a4acfc3a11d4ed24dca0d0c0fc8e0067e' \
--header 'Content-Type: application/json' \
--data-raw '{
  "valueRange": {
    "range": "Q7PlXT!A1:B4",
    "values": [
      [
        "=",
        "https://www.xxx.com/"
      ],
      [
        "Hello",
        "https://www.xxx.com/"
      ],
      [
        "World",
        "https://www.xxx.com/"
      ],
      [
        "=",
        "https://www.xxx.com/"
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
|  ∟revision|int|Version number of sheet | ###  Response body example  
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
