---
title: "Operate Sheets"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uYTMzUjL2EzM14iNxMTN"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/sheets/v2/spreadsheets/:spreadsheetToken/sheets_batch_update"
section: "Docs"
scopes:
  - "drive:drive"
  - "sheets:spreadsheet"
updated: "1647001338000"
---

# Sheet operations

Use this API to add, copy, or delete sheets based on a spreadsheetToken.
::: note
This API uses the same request URL as  Update sheet properties, but has different parameters. Read this document carefully before calling.

## Request
| HTTP URL | https://open.larksuite.com/open-apis/sheets/v2/spreadsheets/:spreadsheetToken/sheets_batch_update |
| --- | --- |
| HTTP Method | POST |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `drive:drive` `sheets:spreadsheet` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `user_access_token`  or  `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" Learn more about how to access and use access_token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| spreadsheetToken | string | Spreadsheet token. For more information about how to obtain the token, see Sheets Developer's Guide | ###  Request body   
|Parameter|Type|Required|Description
|--|-----|--|----|
|requests| |Yes|Request operation, which can add, delete, or copy sheet.  Only one operation can be selected at a time.| 
|  ∟addSheet| |No|Add sheet | 
|    ∟properties| |Yes| Sheet properties| 
|      ∟title|string |Yes| Sheet title| 
|      ∟index|int |No|Location of added sheet, added at the front if not specified| Request body|
|  ∟copySheet| |No|Copy sheet| 
|    ∟source| |Yes|Sheet resources to copy| 
|      ∟sheetId|string|Yes|Source sheetId| 
|    ∟destination| |Yes|Sheet  properties| 
|      ∟title|string |No|Target sheet name, be "old_title (copy_0)" if not filled| 
|  ∟deleteSheet| |No|Delete sheet| 
|    ∟sheetId| |Yes| sheetId| ### Request body example  
```json
{
  "requests": [
    {
      "addSheet": {
        "properties": {
          "title": "string",
          "index": 0
        }
      }
    },
    {
      "copySheet": {
        "source": {
          "sheetId": "string"
        },
        "destination": {
          "title": "string"
        }
      }
    },
    {
      "deleteSheet": {
        "sheetId": "string"
      }
    }
  ]
}

```
### cURL Request example
```
curl --location --request POST 'https://open.larksuite.com/open-apis/sheets/v2/spreadsheets/shtcngNygNfuqhxTBf588jwgWbJ/sheets_batch_update' \
--header 'Authorization: Bearer t-e346617a4acfc3a11d4ed24dca0d0c0fc8e0067e' \
--header 'Content-Type: application/json' \
--data-raw '{
  "requests": [
    {
      "addSheet": {
        "properties": {
          "title": "Title 1",
          "index": 1
        }
      }
    },
    {
      "copySheet": {
        "source": {
          "sheetId": "0b6377"
        },
        "destination": {
          "title": "Title 2"
        }
      }
    },
    {
      "deleteSheet": {
        "sheetId": "l8Gdub"
      }
    }
  ]
}
'
```
## Response  

### Response body
|Parameter|Type|Description|
|--|-----|--|
|replies| |Return result for this sheet operation|
|  ∟addSheet/copySheet| |Properties of added/copied sheet|
|    ∟properties|string| spreadsheet properties|
|      ∟sheetId|string| sheetId|
|      ∟title|string| Sheet title|
|      ∟index|int| Sheet location|
|  ∟deleteSheet| |Delete sheet|
|    ∟result|bool|Whether deletion was successful|
|    ∟sheetId|string| sheetId| ### Response body example  
```json
{
  "code": 0,
  "msg": "Success",
  "data": {
    "replies": [
      {
        "addSheet": {
          "properties": {
            "sheetId": "string",
            "title": "string",
            "index": 0
          }
        },
        "copySheet": {
          "properties": {
            "sheetId": "string",
            "title": "string",
            "index": 0
          }
        },
        "deleteSheet": {
          "result": true,
          "sheetId": "string"
        }
      }
    ]
  }
}
```
### Error code

For details, refer to Server-side error codes.
