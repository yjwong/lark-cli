---
title: "Reading Multiple Ranges"
url: "https://open.larksuite.com/document/ukTMukTMukTM/ukTMzUjL5EzM14SOxMTN"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/sheets/v2/spreadsheets/:spreadsheetToken/values_batch_get"
section: "Docs"
scopes:
  - "drive:drive"
  - "drive:drive:readonly"
  - "sheets:spreadsheet"
  - "sheets:spreadsheet:readonly"
updated: "1647001397000"
---

# Read multiple ranges

This API is used to read the values in multiple ranges based on spreadsheetToken and range. No more than 10 MB of data can be returned.

##  Request
| HTTP URL | https://open.larksuite.com/open-apis/sheets/v2/spreadsheets/:spreadsheetToken/values_batch_get |
| --- | --- |
| HTTP Method | GET |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `drive:drive` `drive:drive:readonly` `sheets:spreadsheet` `sheets:spreadsheet:readonly` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `user_access_token`  or  `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" Learn more about how to obtain and use access_token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ###  Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| spreadsheetToken | string | Spreadsheet token. For more information about how to obtain the token, see Sheets Developer's Guide | ###  Query parameters  
> **Note:** For cells containing formulas, only the values of formulas other than data references and array formulas can be obtained.
> user_id_type currently defaults to the lark_id. After January 26, 2022, it will default to the open_id and lark_id will no longer be used. Please make the necessary changes as soon as possible.

| Parameter                 | Type   | Required | Description                                                         |
| -------------------- | ------ | ---- | ------------------------------------------------------------ |
| ranges               | array | Yes   | Multiple query ranges,  such as:  url?ranges=range1,range2 . ⁣Here, range includes the sheetId range and cell range. Four indexing methods are supported. For details, see  Sheets Developer's Guide .|
| valueRenderOption    | string | No   | valueRenderOption=ToString can return plain text values; valueRenderOption=FormattedValue calculates and formats cells; valueRenderOption=Formula returns the formula body when cells contain formulas; valueRenderOption=UnformattedValue calculates but doesn't format cells.|
| dateTimeRenderOption  | string | No   | dateTimeRenderOption=FormattedString calculates and formats the date and time according to the specified format, but doesn't format the number. It returns the formatted string. For details, see Sheets FAQs .|
| user_id_type | string | No | Returned user ID type, values: open_id or union_id| ###  cURL Request example
```
curl --location --request GET 'https://open.larksuite.com/open-apis/sheets/v2/spreadsheets/shtcngNygNfuqhxTBf588jwgWbJ/values_batch_get?ranges=Q7PlXT!A2:B6,0b6377!B1:C8&valueRenderOption=ToString&dateTimeRenderOption=FormattedString' \
--header 'Authorization: Bearer t-e346617a4acfc3a11d4ed24dca0d0c0fc8e0067e' \
```
##  Response  

### Response body
|Parameter|Type|Description|
|--|-----|--|
|revision|int |Version number of sheet |
|spreadsheetToken|string | Spreadsheet token. For details, see Sheets Developer's Guide.|
|totalCells|int|Total cells read|
|valueRanges|array|Value and range|
|  ∟majorDimension|string|Insert dimension|
|  ∟range|string|Range of returned data. When empty, this indicates there is no data in the query range.|
|  ∟revision|int|Version number of sheet |
|  ∟values|array|Values obtained by the query| ###  Response body example    
```json
{
  "code": 0,
  "data": {
    "revision": 0,
    "spreadsheetToken": "***",
    "totalCells": 0,
    "valueRanges": [
      {
        "majorDimension": "ROWS",
        "range": "range1",
        "revision": 0,
        "values": [
          [
            "***"
          ]
        ]
      },
      {
        "majorDimension": "ROWS",
        "range": "range2",
        "revision": 0,
        "values": [
          [
            "***"
          ]
        ]
      }
    ]
  },
  "msg": "Success"
}
```  
  
### Error code

For details, refer to Server-side error codes.
