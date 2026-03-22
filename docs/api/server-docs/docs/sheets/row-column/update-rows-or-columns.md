---
title: "Update Rows or Columns"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uYjMzUjL2IzM14iNyMTN"
method: "PUT"
api_path: "https://open.larksuite.com/open-apis/sheets/v2/spreadsheets/:spreadsheetToken/dimension_range"
section: "Docs"
scopes:
  - "drive:drive"
  - "sheets:spreadsheet"
updated: "1647001366000"
---

# Update rows and columns

This API is used to update hidden rows and columns and cell sizes based on spreadsheetToken and dimension information. You can act on up to 5,000 rows and columns at once.

##  Request
| HTTP URL | https://open.larksuite.com/open-apis/sheets/v2/spreadsheets/:spreadsheetToken/dimension_range |
| --- | --- |
| HTTP Method | PUT |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `drive:drive` `sheets:spreadsheet` | ###  Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `user_access_token` or `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" Learn more about how to obtain and use access_token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| spreadsheetToken | string | Spreadsheet token. For more information about how to obtain the token, see Sheets Developer's Guide | ### Request body  
|Parameter|Type|Required|Description|Source|
|--|-----|--|----|----|
|dimension||Yes|Dimension information of rows and columns to update| Request body |
|  ∟sheetId|string|Yes|sheetId| Request body |
|  ∟majorDimension|string|No|Values: ROWS (default) or COLUMNS| Request body |
|  ∟startIndex|int|Yes|Start location| Request body |
|  ∟endIndex|int|Yes|End location| Request body |
|dimensionProperties||Yes|Properties of rows and columns to update| Request body |
|  ∟visible|bool|No|true: displayed, false: hidden| Request body |
|  ∟fixedSize|int|No|Row/Column size| Request body | ###  Request body example    
```json
{
    "dimension":{
        "sheetId":"string",
        "majorDimension":"ROWS",
        "startIndex":1,
        "endIndex":3
    },
    "dimensionProperties":{
        "visible":false,
        "fixedSize":50
    }
}
```
###  cURL Request example
```
curl --location --request PUT 'https://open.larksuite.com/open-apis/sheets/v2/spreadsheets/shtcngNygNfuqhxTBf588jwgWbJ/dimension_range' \
--header 'Authorization: Bearer t-e346617a4acfc3a11d4ed24dca0d0c0fc8e0067e' \
--header 'Content-Type: application/json' \
--data-raw '{
    "dimension":{
        "sheetId":"Q7PlXT",
        "majorDimension":"ROWS",
        "startIndex":1,
        "endIndex":3
    },
    "dimensionProperties":{
        "visible":true,
        "fixedSize":60
    }
}'
```

## Response  

 ###  Response body example    
```json
{
    "code": 0,
    "data": {},
    "msg": "Success"
}
```

###  Error code

For details, refer to Server-side error codes.
