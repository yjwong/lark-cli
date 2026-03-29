---
title: "Add Rows or Columns"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uUjMzUjL1IzM14SNyMTN"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/sheets/v2/spreadsheets/:spreadsheetToken/dimension_range"
section: "Docs"
scopes:
  - "drive:drive"
  - "sheets:spreadsheet"
updated: "1647001369000"
---

# Add rows and columns

This API is used to add blank rows and columns at the end of a spreadsheet based on spreadsheetToken and length. You can add up to 5,000 rows or columns at once.

## Request
| HTTP URL | https://open.larksuite.com/open-apis/sheets/v2/spreadsheets/:spreadsheetToken/dimension_range |
| --- | --- |
| HTTP Method | POST |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list |  | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `user_access_token` or `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" Learn more about how to access and use access_token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Path parameters
|Parameter|Type|Required|Description|
|--|-----|--|----|
|spreadsheetToken|string|Yes| Spreadsheet token. For more information about how to obtain the token, see Sheets Developer's Guide.|
### Request body
|Parameter|Type|Required|Description|
|--|-----|--|----|
|dimension||Yes|Dimension information of rows and columns to add|
|  ∟sheetId|string|Yes|sheetId|
|  ∟majorDimension|string|No|Values: ROWS (default) or COLUMNS|
|  ∟length|int|Yes|Number of rows and columns to add, 0<length<5,000|
### Request body example
```json
{
  "dimension":{
       "sheetId": "string",
        "majorDimension": "ROWS",
        "length": 1
     }
}
```
###  cURL request example
```
curl --location --request POST 'https://open.larksuite.com/open-apis/sheets/v2/spreadsheets/shtcngNygNfuqhxTBf588jwgWbJ/dimension_range' \
--header 'Authorization: Bearer t-e346617a4acfc3a11d4ed24dca0d0c0fc8e0067e' \
--header 'Content-Type: application/json' \
--data-raw '{
  "dimension":{
       "sheetId": "Q7PlXT",
        "majorDimension": "ROWS",
        "length": 8
     }
}'
```
## Response
### Response body
|Parameter|Type|Description|
|--|-----|--|
|addCount|int |Number of rows and columns to add|
|majorDimension|string |Insert dimension|
### Response body example 
```json
{
    "code": 0,
    "data": {
        "addCount": 1,
        "majorDimension": "ROWS"
    },
    "msg": "Success"
}

```
### Error code

For details, refer to Server-side error codes.
