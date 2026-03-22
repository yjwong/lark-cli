---
title: "Update Spreadsheet Properties"
url: "https://open.larksuite.com/document/ukTMukTMukTM/ucTMzUjL3EzM14yNxMTN"
method: "PUT"
api_path: "https://open.larksuite.com/open-apis/sheets/v2/spreadsheets/:spreadsheetToken/properties"
section: "Docs"
scopes:
  - "drive:drive"
  - "sheets:spreadsheet"
updated: "1647001332000"
---

# Update spreadsheet properties

Use this API to update spreadsheet properties, including spreadsheet titles, based on a spreadsheetToken.

## Request
| HTTP URL | https://open.larksuite.com/open-apis/sheets/v2/spreadsheets/:spreadsheetToken/properties |
| --- | --- |
| HTTP Method | PUT |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `drive:drive` `sheets:spreadsheet` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `user_access_token`  or  `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" Learn more about how to obtain and use access_token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| spreadsheetToken | string | Spreadsheet token. For more information about how to obtain the token, see Sheets Developer's Guide | ### Request body  
|Parameter|Type|Required|Description          |
|--|-----|--|----|
|properties| |Yes|spreadsheet  properties| 
|  ∟title|string |Yes|Spreadsheet title, cannot exceed 100 characters| ### Request body example  
```json
{
    "properties": {
        "title": "string"
    }
}
```
### cURL Request example
```
curl --location --request PUT 'https://open.larksuite.com/open-apis/sheets/v2/spreadsheets/shtcngNygNfuqhxTBf588jwgWbJ/properties' \
--header 'Authorization: Bearer t-e346617a4acfc3a11d4ed24dca0d0c0fc8e0067e' \
--header 'Content-Type: application/json' \
--data-raw '{
    "properties":{
        "title": " Title 1"
    }
}'
```
## Response  

### Response body
|Parameter|Type|Description|
|--|-----|--|
|spreadsheetToken|string|The token of the spreadsheet|
|title|string| spreadsheet  title| ### Response body example    
```json
{
    "code": 0,
    "msg": "Success",
    "data": {
        "spreadsheetToken": "***",
        "title": "***"
    }
}
```
### Error code

For details, refer to Server-side error codes.
