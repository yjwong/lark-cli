---
title: "Split Cells"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uATNzUjLwUzM14CM1MTN"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/sheets/v2/spreadsheets/:spreadsheetToken/unmerge_cells"
section: "Docs"
scopes:
  - "drive:drive"
  - "sheets:spreadsheet"
updated: "1647001415000"
---

# Split cells

This API is used to split cells based on spreadsheetToken and dimension information. You can act on up to 5,000 rows and 100 columns at once.
##  Request
| HTTP URL | https://open.larksuite.com/open-apis/sheets/v2/spreadsheets/:spreadsheetToken/unmerge_cells |
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
|spreadsheetToken|string|Yes|Spreadsheet token. For more information about how to obtain the token, see Sheets Developer's Guide.|
###  Request body
|Parameter|Type|Required|Description|
|--|-----|--|----|
|range|string|Yes|Query range, includes the sheetId range and cell range. Four indexing methods are supported. For details, see Sheets Developer's Guide.|
### Request body example
```json
  {
     "range": "string"
  }
```
### cURL Request example
```
curl --location --request POST 'https://open.larksuite.com/open-apis/sheets/v2/spreadsheets/shtcngNygNfuqhxTBf588jwgWbJ/unmerge_cells' \
--header 'Authorization: Bearer t-e346617a4acfc3a11d4ed24dca0d0c0fc8e0067e' \
--header 'Content-Type: application/json' \
--data-raw '  {
     "range": "Q7PlXT!F7:F8"
  }'
```
##  Response
### Response body
  |Parameter|Type|Description|
|--|-----|--|
|spreadsheetToken|string |Spreadsheet token|
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
