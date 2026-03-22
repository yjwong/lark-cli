---
title: " Delete Rows or Columns"
url: "https://open.larksuite.com/document/ukTMukTMukTM/ucjMzUjL3IzM14yNyMTN"
method: "DELETE"
api_path: "https://open.larksuite.com/open-apis/sheets/v2/spreadsheets/:spreadsheetToken/dimension_range"
section: "Docs"
scopes:
  - "drive:drive"
  - "sheets:spreadsheet"
updated: "1647001363000"
---

# Delete rows or columns

This API is used to delete rows and columns based on spreadsheetToken and dimension information. You can delete up to 5,000 rows and columns at once.

##  Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/sheets/v2/spreadsheets/:spreadsheetToken/dimension_range |
| HTTP Method | DELETE |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `drive:drive` `sheets:spreadsheet` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `user_access_token`  or  `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" Learn more about how to obtain and use access_token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ###  Path parameters
|Parameter|Type|Required|Description                                          |
|--|-----|--|----|
|spreadsheetToken|string|Yes| Spreadsheet token. For details, see Sheets Developer's Guide.|
###  Request body
|Parameter|Type|Required|Description                                          |
|--|-----|--|----|
|dimension||Yes|Dimension information of the rows and columns to delete|
|  ∟sheetId|string|Yes|sheetId|
|  ∟majorDimension|string|No|Values: ROWS (default) or COLUMNS|
|  ∟startIndex|int|Yes|Start location|
|  ∟endIndex|int|Yes|End location|
### Request body example
```json
{
    "dimension":{
        "sheetId":"string",
        "majorDimension":"ROWS",
        "startIndex":1,
        "endIndex":3
    }
}
```
###  cURL Request example
```
curl --location --request DELETE 'https://open.larksuite.com/open-apis/sheets/v2/spreadsheets/shtcngNygNfuqhxTBf588jwgWbJ/dimension_range' \
--header 'Authorization: Bearer t-e346617a4acfc3a11d4ed24dca0d0c0fc8e0067e' \
--header 'Content-Type: application/json' \
--data-raw '{
    "dimension":{
        "sheetId":"Q7PlXT",
        "majorDimension":"ROWS",
        "startIndex":1,
        "endIndex":3
    }
}'
```
##  Response
### Response body
|Parameter|Type|Description|
|--|-----|--|
|delCount|int |Number of rows/columns to delete|
|majorDimension|string |Insert dimension|
### Response body example
```json
{
    "code": 0,
    "data": {
        "delCount": 0,
        "majorDimension": "ROWS"
    },
    "msg": "Success"
}

```
### Error code

For details, refer to Server-side error codes.
