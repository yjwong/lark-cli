---
title: "Add Locked Cells"
url: "https://open.larksuite.com/document/ukTMukTMukTM/ugDNzUjL4QzM14CO0MTN"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/sheets/v2/spreadsheets/:spreadsheetToken/protected_dimension"
section: "Docs"
scopes:
  - "drive:drive"
  - "sheets:spreadsheet"
updated: "1647001375000"
---

# Add protected range

Use this API to add multiple protected ranges based on  spreadsheetToken  and dimension information. You can operate on up to 5000 rows or columns.

## Request
| HTTP URL | https://open.larksuite.com/open-apis/sheets/v2/spreadsheets/:spreadsheetToken/protected_dimension |
| --- | --- |
| HTTP Method | POST |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `drive:drive` `sheets:spreadsheet` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `user_access_token`  or  `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" Learn more about how to obtain and use access_token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Query parameters
> User_id_type currently defaults to the lark_id. After Jan. 26, 2022, it will default to the open_id and lark_ids will no longer be used. Please make the necessary changes as soon as possible.

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| user_id_type | string | No | User ID type in the request, currently supports open_id and union_id | ### Path parameters
|Parameter|Type|Required|Description|
|--|-----|--|----|
|spreadsheetToken|string|Yes|Spreadsheet token. For more information about how to obtain the token, see  Sheets Developer's Guide.| ### Request body
> The editors field is no longer used. After Jan. 26, 2022, this field will no longer be supported. Please use the users field instead. Users must contain the IDs of users with edit permission for the protected range to add. The ID type is determined by the user_id_type field.

|Parameter|Type|Required|Description|
|--|-----|--|----|
|addProtectedDimension|Yes|Dimension information for the protected range(s) to be added| 
|  ∟dimension|Yes|Dimension information for the protected rows and columns| 
|    ∟sheetId|string|Yes|sheetId| 
|    ∟majorDimension|string|No|Values: ROWS (default) or COLUMNS| 
|    ∟startIndex|int|Yes|Start position|
|    ∟endIndex|int|Yes|End position| 
|  ∟editors|array|NO|userIDs of users who can edit the protected range| 
|  ∟users|array|No|IDs of users who can edit the protected range, ID type determined by user_id_type| 
|  ∟lockInfo|string|No|Protected range information| ### Request body example

```json
{
    "addProtectedDimension":[
        {
            "dimension":{
                "sheetId":"string",
                "majorDimension":"COLUMNS",
                "startIndex":10,
                "endIndex":13
            },
            "users":[
                "ou_326f4b0552770f2de069deb256de5b30"
            ],
            "lockInfo":"You can edit"
        }
    ]
}
```
  
###  cURL Request example
```
curl --location --request POST 'https://open.larksuite.com/open-apis/sheets/v2/spreadsheets/shtcngNygNfuqhxTBf588jwgWbJ/protected_dimension?user_id_type=open_id' \
--header 'Authorization: Bearer t-e346617a4acfc3a11d4ed24dca0d0c0fc8e0067e' \
--header 'Content-Type: application/json' \
--data-raw '{
    "addProtectedDimension":[
        {
            "dimension":{
                "sheetId":"Q7PlXT",
                "majorDimension":"COLUMNS",
                "startIndex":10,
                "endIndex":13
            },
            "users":[
              "ou_326f4b0552770f2de069deb256de5b30"
            ],
            "lockInfo":"You can edit"
        }
    ]
}'
```
  
 ## Response
### Response body
> The editors field is discarded. After Jan. 26, 2022, this field will no longer be supported. Please use the users field instead. Users must contain the IDs of users with edit permission for the protected range to add. The ID type is determined by the user_id_type field.

|Parameter|Type|Description|
|--|-----|--|
|addProtectedDimension|array|Dimension information for the protected range(s) to be added, multiple ranges are allowed| 
|  ∟dimension|Dimension information for the protected rows and columns| 
|    ∟sheetId|string|sheetId| 
|    ∟majorDimension|string|Values: ROWS (default) or COLUMNS| 
|    ∟startIndex|int||Start position|
|    ∟endIndex|int|End position| 
|  ∟editors|array|UserIDs of users who can edit the protected range| 
|  ∟users|array|IDs of users who can edit the protected range, ID type determined by user_id_type.|
|  ∟lockInfo|string|Protected range information|
|  ∟protectId|string|Unique uid of the protected range, can be used to remove protection later.| ### Response body example  

```json
{
    "code": 0,
    "data": {
        "addProtectedDimension": [
            {
                "dimension": {
                    "endIndex": 0,
                    "majorDimension": "COLUMNS",
                    "sheetId": "***",
                    "startIndex": 0
                },
                "users": [
                    "ou_326f4b0552770f2de069deb256de5b30"
                ],
                "lockInfo": "***",
                "protectId": "***"
            }
        ]
    },
    "msg": "Success"
}

```  
  
### Error code

For details, refer to Server-side error codes.
