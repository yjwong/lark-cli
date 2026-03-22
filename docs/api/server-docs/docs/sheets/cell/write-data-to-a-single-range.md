---
title: "Write Data to a Single Range"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uAjMzUjLwIzM14CMyMTN"
method: "PUT"
api_path: "https://open.larksuite.com/open-apis/sheets/v2/spreadsheets/:spreadsheetToken/values"
section: "Docs"
scopes:
  - "drive:drive"
  - "sheets:spreadsheet"
updated: "1647001399000"
---

# Write data to a single range

This API is used to write data to a single range based on spreadsheetToken and range. Any data in the range will be overwritten. You can write up to 5,000 rows and 100 columns of data. A single cell can't exceed 50,000 characters.
##  Request
| HTTP URL | https://open.larksuite.com/open-apis/sheets/v2/spreadsheets/:spreadsheetToken/values |
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
|spreadsheetToken|string|Yes| Spreadsheet token. For more information about how to obtain the token, see Sheets Developer's Guide| URL PATH.|
###  Request body
|Parameter|Type|Required|Description|
|--|-----|--|----|
|valueRange||Yes|Value and range|
|  ∟range|string|Yes|Range to update, includes the sheetId range and cell range. Three indexing methods are supported. For details, see Sheets Developer's Guide. The range indicated by range needs to be greater than or equal to the range occupied by values.|
|  ∟values|array>|Yes|Values to write. To write formulas, hyperlinks, emails, and @mention users, see the appendix Data types that can be written to sheet.|
### Request body example
```json
{
"valueRange":{
    "range": "string",
    "values": [
      [
        "string", 1 ,"http://www.xx.com"
      ]
    ]
    }
}
```
  
###  cURL Request example
```
curl --location --request PUT 'https://open.larksuite.com/open-apis/sheets/v2/spreadsheets/shtcngNygNfuqhxTBf588jwgWbJ/values' \
--header 'Authorization: Bearer t-e346617a4acfc3a11d4ed24dca0d0c0fc8e0067e' \
--header 'Content-Type: application/json' \
--data-raw '{
"valueRange":{
    "range": "Q7PlXT!A1:B2",
    "values": [
      [
        "Hello", 1
      ],
      [
        "World", 1
      ]
    ]
    }
}'
```
  
##  Response
### Response body
  |Parameter|Type|Description|
|--|-----|--|
|spreadsheetToken|string |spreadsheet   token |
|updatedRange|string |Range to write|
|updatedRows|int|Number of rows to write|
|updatedColumns|int|Number of columns to write|
|updatedCells|int|Total cells to write|
|revision|int|Version number of sheet |
###  Response body example  
```json
{
    "code": 0,
    "data": {
        "revision": 0,
        "spreadsheetToken": "***",
        "updatedCells": 0,
        "updatedColumns": 0,
        "updatedRange": "***",
        "updatedRows": 0
    },
    "msg": "Success"
}

```
### Error code

For details, refer to Server-side error codes.
