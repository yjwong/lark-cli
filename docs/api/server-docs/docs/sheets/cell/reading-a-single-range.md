---
title: "Reading a Single Range"
url: "https://open.larksuite.com/document/ukTMukTMukTM/ugTMzUjL4EzM14COxMTN"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/sheets/v2/spreadsheets/:spreadsheetToken/values/:range"
section: "Docs"
scopes:
  - "drive:drive"
  - "drive:drive:readonly"
  - "sheets:spreadsheet"
  - "sheets:spreadsheet:readonly"
updated: "1647001394000"
---

# Read a single range

This API is used to read the values in a single range based on spreadsheetToken and range. No more than 10 MB of data can be returned.

##  Request
| HTTP URL | https://open.larksuite.com/open-apis/sheets/v2/spreadsheets/:spreadsheetToken/values/:range |
| --- | --- |
| HTTP Method | GET |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `drive:drive` `drive:drive:readonly` `sheets:spreadsheet` `sheets:spreadsheet:readonly` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `user_access_token` or  `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" Learn more about how to obtain and use access_token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ###  Path parameters
|Parameter|Type|Required|Description|
|--|-----|--|----|
|spreadsheetToken|string|Yes|Spreadsheet token. For details, see Overview.|
|range|string|Yes|Query range, includes the sheetId range and cell range. For details, see Sheets Developer's Guide .| ###  Query parameters  
> **Note:** For cells containing formulas, only the values of formulas other than data references and array formulas can be obtained.
> user_id_type currently defaults to the lark_id. After January 26, 2022, it will default to the open_id and lark_id will no longer be used. Please make the necessary changes as soon as possible.

| Parameter                 | Type   | Required | Description                                                         |
| -------------------- | ------ | ---- | ------------------------------------------------------------ |
| valueRenderOption    | string | No   | valueRenderOption=ToString can return plain text values; valueRenderOption=FormattedValue calculates and formats cells; valueRenderOption=Formula returns the formula body when cells contain formulas; valueRenderOption=UnformattedValue calculates but doesn't format cells. |
| dateTimeRenderOption | string | No   | dateTimeRenderOption=FormattedString calculates and formats the date and time according to the specified format, but doesn't format the number. It returns the formatted string. For details, see Sheets FAQs .|
| user_id_type         | string | No   | Returned user ID type, values: open_id or union_id                       |
###  cURL Request example
```
curl --location --request GET 'https://open.larksuite.com/open-apis/sheets/v2/spreadsheets/shtcngNygNfuqhxTBf588jwgWbJ/values/Q7PlXT!A3:D200?valueRenderOption=ToString&dateTimeRenderOption=FormattedString' \
--header 'Authorization: Bearer t-ce3540c5f02ac074535f1f14d64fa90fa49621c0'
```
##  Response
### Response body 
|Parameter|Type|Description|
|--|-----|--|
|revision|Int |Version number of sheet |
|spreadsheetToken|string | Spreadsheet token. For details, see Overview.|
|valueRange||Value and range|
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
        "valueRange": {
            "majorDimension": "ROWS",
            "range": "***",
            "revision": 0,
            "values": [
                [
                    "***"
                ]
            ]
        }
    },
    "msg": "Success"
}
```
### Error code

For details, refer to Server-side error codes.
