---
title: "Insert Rows or Columns"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uQjMzUjL0IzM14CNyMTN"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/sheets/v2/spreadsheets/:spreadsheetToken/insert_dimension_range"
section: "Docs"
scopes:
  - "drive:drive"
  - "sheets:spreadsheet"
updated: "1647001372000"
---

# Insert rows and columns

This API is used to insert rows and columns based on spreadsheetToken and dimension information.  
If startIndex=3 and endIndex=7, rows and columns are inserted from the 4th row until the 7th row, for a total of 4 rows. You can insert up to 5,000 rows or columns at once.

##  Request
| HTTP URL | https://open.larksuite.com/open-apis/sheets/v2/spreadsheets/:spreadsheetToken/insert_dimension_range |
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
| spreadsheetToken | string | Spreadsheet token. For more information about how to obtain the token, see Sheets Developer's Guide | ###  Request body  
|Parameter|Type|Required|Description|
|--|-----|--|----|
|dimension|Yes|Dimension information of rows and columns to insert| 
|  ∟sheetId|string|Yes|Sheet ID| 
|  ∟majorDimension|string|No|Values: ROWS (default) or COLUMNS| 
|  ∟startIndex|int|Yes|Start location|
|  ∟endIndex|int|Yes|End location|
|inheritStyle|string|No|BEFORE or AFTER. Style won't be inherited if this field is empty.| ###  Request body example  
```json
{
    "dimension":{
        "sheetId":"string",
        "majorDimension":"ROWS",
        "startIndex":0,
        "endIndex":0
    },
    "inheritStyle":"BEFORE"
}
```
### cURL Request example
```
curl --location --request POST 'https://open.larksuite.com/open-apis/sheets/v2/spreadsheets/shtcngNygNfuqhxTBf588jwgWbJ/insert_dimension_range' \
--header 'Authorization: Bearer t-e346617a4acfc3a11d4ed24dca0d0c0fc8e0067e' \
--header 'Content-Type: application/json' \
--data-raw '{
    "dimension":{
        "sheetId":"Q7PlXT",
        "majorDimension":"ROWS",
        "startIndex":3,
        "endIndex":5
    },
    "inheritStyle":"AFTER"
}'
```
##  Response  

###  Response body example    

```json
{
    "code": 0,
    "msg": "Success",
    "data": {}
}
```

### Error code

For details, refer to Server-side error codes.
