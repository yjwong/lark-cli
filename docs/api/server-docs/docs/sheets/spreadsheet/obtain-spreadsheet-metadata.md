---
title: "Obtain Spreadsheet Metadata"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uETMzUjLxEzM14SMxMTN"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/sheets/v2/spreadsheets/:spreadsheetToken/metainfo"
section: "Docs"
scopes:
  - "drive:drive"
  - "drive:drive:readonly"
  - "sheets:spreadsheet"
  - "sheets:spreadsheet:readonly"
updated: "1647001329000"
---

#  Get spreadsheet metadata

Use this API to get spreadsheet metadata based on a spreadsheetToken.

## Request

| HTTP URL | https://open.larksuite.com/open-apis/sheets/v2/spreadsheets/:spreadsheetToken/metainfo |
| --- | --- |
| HTTP Method | GET |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `drive:drive` `drive:drive:readonly` `sheets:spreadsheet` `sheets:spreadsheet:readonly` | ### Request header

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `user_access_token`  or  `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" Learn more about how to access and use access_token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Path parameters

| Parameter | Type | Description |
| --- | --- | --- |
| spreadsheetToken | string | Spreadsheet  token. For more information about how to obtain the token, see Sheets Developer's Guide | ### Query parameters
> user_id_type currently defaults to the lark_id. After January 26, 2022, it will default to the open_id and lark_id will no longer be used. Please make the necessary changes as soon as possible.

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| extFields | string | No | Additional returned fields, protected row and column information is returned when extFields = protectedRange |
| user_id_type | string | No | Returned user ID type, values: open_id,union_id | ### cURL Request example

```
curl --location --request GET 'https://open.larksuite.com/open-apis/sheets/v2/spreadsheets/shtcngNygNfuqhxTBf588jwgWbJ/metainfo?extFields=protectedRange&user_id_type=open_id' \
--header 'Authorization: Bearer t-e346617a4acfc3a11d4ed24dca0d0c0fc8e0067e'
```

## Response

### Response body
> The ownerUser field is discarded. After Jan. 26, 2022, this field will no longer be supported. Please use the ownerUserID field instead. OwnerUserID must contain the ID of the document owner. The ID type is determined by the user_id_type field.

| Parameter                              | Type             | Description                                                         |
| --------------------------------- | ---------------- | ------------------------------------------------------------ |
| spreadsheetToken                  | string           | spreadsheet  token                                         |
| properties                        |                  | spreadsheet  properties                                           |
|   ∟title                      | string           | spreadsheet  title                                           |
|   ∟ownerUser                  | int64              | Owner ID                                                  |
|   ∟ownerUserID                  | string              | Owner ID, determined by the user_id_type field. This value is returned when user_id_type is not blank.                                                  |
|   ∟sheetCount                 | int              |Number of sheet in the  spreadsheet                                     |
|   ∟revision                   | int              | Version of this sheet                                              |
| sheets                            | array |List of sheets for the  spreadsheet                                     |
|   ∟sheetId                    | string           | Sheet ID                                                  |
|   ∟title                      | string           |Sheet title                                                 |
|   ∟index                      | int              |Sheet location                                                 |
|   ∟rowCount                   | int              |Max sheet rows                                             |
|   ∟columnCount                | int              |Max sheet columns                                             |
|   ∟frozenRowCount             | int              | Number of frozen rows for this sheet, less than or equal to the maximum number of sheet rows. 0 indicates rows are unfrozen. |
|   ∟frozenColCount             | int              | Number of frozen columns for this sheet, less than or equal to the maximum number of sheet columns. 0 indicates columns are unfrozen. |
|   ∟merges                     | array | Range of merged cells for this sheet                                  |
|     ∟startRowIndex        | int              | Merged cell range start row index , starts from 0                   |
|     ∟startColumnIndex     | int              | Merged cell range start column index , starts from 0                   |
|     ∟rowCount             | int              | Number of rows in merged cell range                                       |
|     ∟columnCount          | int              | Number of columns in merged cell range                                       |
|   ∟protectedRange             | array | Protected range for this sheet                                          |
|     ∟dimension            |                  | Information of protected rows and columns. This field is blank for a protected sheet.               |
|       ∟startIndex     | int              | Protected rows and columns start location, starts from 1                              |
|       ∟endIndex       | int              | Protected rows and columns end location, starts from 1                              |
|       ∟majorDimension | string           | ROWS indicates protected rows, COLUMNS indicates protected columns                  |
|       ∟sheetId        | string           | ID of the sheet that includes the protected range                                        |
|     ∟protectId            | string           | Protected range ID                                                   |
|     ∟lockInfo             | string           | Protection description                                                     |
|     ∟sheetId              | string           | ID of the protected sheet                                                |
|   ∟blockInfo                  |                  | If this field is included, the sheet is not a table                             |
|     ∟blockToken           | string           |Block token                                                 |
|     ∟blockType            | string           |Block type                                                  | ### Response body example

```json
{
    "code": 0,
    "msg": "Success",
    "data": {
        "properties": {
            "title": "",
            "ownerUser": 0,
            "sheetCount": 0,
            "revision": 0
        },
        "sheets": [
            {
                "sheetId": "***",
                "title": "***",
                "index": 0,
                "rowCount": 0,
                "columnCount": 0,
                "frozenColCount": 0,
                "frozenRowCount": 0,
                "merges": [
                    {
                        "columnCount": 0,
                        "rowCount": 0,
                        "startColumnIndex": 0,
                        "startRowIndex": 0
                    }
                ],
                "protectedRange": [
                    {
                        "dimension": {
                            "endIndex": 0,
                            "majorDimension": "ROWS",
                            "sheetId": "***",
                            "startIndex": 0
                        },
                        "protectId": "***",
                        "sheetId": "***",
                        "lockInfo": "***"
                    }
                ]
            },
            {
                "blockInfo": {
                    "blockToken": "***",
                    "blockType": "***"
                },
                "columnCount": 0,
                "frozenColCount": 0,
                "frozenRowCount": 0,
                "index": 0,
                "rowCount": 0,
                "sheetId": "***",
                "title": "*** "
            }
        ],
        "spreadsheetToken": "***"
    }
}
```

### Error code

For details, refer to Server-side error codes.
