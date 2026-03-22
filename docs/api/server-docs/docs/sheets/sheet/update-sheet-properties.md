---
title: "Update Sheet Properties"
url: "https://open.larksuite.com/document/ukTMukTMukTM/ugjMzUjL4IzM14COyMTN"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/sheets/v2/spreadsheets/:spreadsheetToken/sheets_batch_update"
section: "Docs"
scopes:
  - "drive:drive"
  - "sheets:spreadsheet"
updated: "1647001335000"
---

#  Update sheet properties
::: note
This API uses the same request URL as  Sheet operations , but has different parameters. Read this document carefully before calling.

Use this API to update sheet properties based on a spreadsheetToken.

## Request
| HTTP URL | https://open.larksuite.com/open-apis/sheets/v2/spreadsheets/:spreadsheetToken/sheets_batch_update |
| --- | --- |
| HTTP Method | POST |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `drive:drive` `sheets:spreadsheet` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `user_access_token` or  `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" Learn more about how to obtain and use access_token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Query parameters
> user_id_type currently defaults to the lark_id. After January 26, 2022, it will default to the open_id and lark_id will no longer be used. Please make the necessary changes as soon as possible.

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| user_id_type | string | No | Request user ID type, values: open_id or union_id | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| spreadsheetToken | string | spreadsheet , to obtain the token, see Sheets Developer's Guide | ### Request body 
> The userIds field is discarded. After Jan. 26, 2022, this field will no longer be supported. Please use the userIDs field instead. UserIDs must contain the IDs of users with edit permission for the protected range to add. The ID type is determined by the user_id_type field.

|Parameter|Type|Required|Description|
|--|-----|--|-----------|
|requests| |Yes|Request operation, which can add, delete, or copy sheet. Only one operation can be selected at a time.| 
|  ∟updateSheet | |No|Update sheet| 
|    ∟properties| |Yes| Sheet properties|
|      ∟sheetId|string|Yes|The unique identifier of the spreadsheet (read-only)| 
|      ∟title|string |No|Change sheet title| 
|      ∟index|int |No|Move sheet location| 
|      ∟hidden|bool |No|Hidden spreadsheet, default value: false| 
|      ∟frozen        RowCount|int |No|Number of frozen rows, less than or equal to the maximum number of sheet rows. 0 indicates rows are unfrozen.|
|      ∟frozen        ColCount|int |No|Number of frozen columns of the sheet , less than or equal to the maximum number of sheet columns. 0 indicates columns are unfrozen.|
|      ∟protect|  |No|Lock spreadsheet| 
|        ∟lock| string |yes|LOCK or UNLOCK | 
|        ∟lockInfo|string  |No|Lock information| 
|        ∟userIds| array |No|Add other users with edit permission besides this user and the owner (discarded) .| 
|        ∟userIDs| array |No|Add other users with edit permission besides this user and the owner (use this field when user_id_type is not blank) .| 
### Request body example    
```json
{
  "requests": [
    {
      "updateSheet": {
        "properties": {
          "sheetId": "string",
          "title": "string",
          "index": "int",
          "hidden": "bool",
          "frozenColCount": "int",
          "frozenRowCount": "int",
          "protect": {
            "lock": "LOCK",
            "lockInfo": "111",
            "userIDs": [
              "ou_xxxxxxxxxx"
            ]
          }
        }
      }
    }
  ]
}
```
  ### cURL Request example
```
curl --location --request POST 'https://open.larksuite.com/open-apis/sheets/v2/spreadsheets/shtcngNygNfuqhxTBf588jwgWbJ/sheets_batch_update?user_id_type=open_id' \
--header 'Authorization: Bearer t-e346617a4acfc3a11d4ed24dca0d0c0fc8e0067e' \
--header 'Content-Type: application/json' \
--data-raw '{
  "requests": [
    {
      "updateSheet": {
        "properties": {
          "sheetId": "zajIJ",
          "title": "",
          "index": 0,
          "frozenColCount": 10,
          "frozenRowCount": 20,
          "protect": {
            "lock": "LOCK",
            "lockInfo": "Lock Test"
            "userIDs": [
               "ou_xxxxxxxxxxxxx"
            ]
          }
        }
      }
    }
  ]
}'
```
## Response  
### Response body
> The userIds field is discarded. After Jan. 26, 2022, this field will no longer be supported. Please use the userIDs field instead. UserIDs must contain the IDs of users with edit permission for the protected range to add. The ID type is determined by the user_id_type field.

|Parameter|Type|Description|
|--|-----|--|
|replies| |Return result for this sheet operation|
|  ∟updateSheet| |Updated sheet properties|
|    ∟properties| |Sheet properties|
|      ∟sheetId|string| SheetId of a spreadsheet|
|      ∟title|string |Updated sheet title|
|      ∟index|int |Moved sheet location|
|      ∟hidden|bool |Whether it is a hidden spreadsheet|
|      ∟frozenRowCount|int |Number of frozen rows|
|      ∟frozenColCount|int |Number of frozen columns|
|      ∟protect| |Protect sheet| 
|        ∟lock|string |LOCK or UNLOCK, or Protected/Remove Protection | 
|        ∟lockInfo|string |Protection information| 
|        ∟userIds|array |Other users added with edit permission besides this user and the owner (discarded) | 
|        ∟userIDs|array |Other users added with edit permission besides this user and the owner (this field is returned when user_id_type is not blank) | ### Response body example    
```json
 {
  "code": 0,
  "msg": "Success",
  "data": {
    "replies": [
      {
        "updateSheet": {
          "properties": {
            "sheetId": "string",
            "title": "string",
            "index": 0,
            "hidden": true,
            "frozenColCount": 0,
            "frozenRowCount": 0,
            "protect": {
              "lock": "LOCK",
              "sheetName": "",
              "permId": "",
              "userIDs": [
                "ou_xxxxxxxxx"
              ]
            }
          }
        }
      }
    ]
  }
}
```
### Error code

For details, refer to Server-side error codes.
