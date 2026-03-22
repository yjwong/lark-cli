---
title: "Modify Protection Scopes"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uUTM5YjL1ETO24SNxkjN"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/sheets/v2/spreadsheets/:spreadsheetToken/protected_range_batch_update"
section: "Docs"
scopes:
  - "drive:drive"
  - "sheets:spreadsheet"
updated: "1647001381000"
---

# Modify protected range

Use this API to modify a protected range based on the protected range ID. Each operation can modify up to 10 IDs.
## Request
| HTTP URL | https://open.larksuite.com/open-apis/sheets/v2/spreadsheets/:spreadsheetToken/protected_range_batch_update |
| --- | --- |
| HTTP Method | POST |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `drive:drive` `sheets:spreadsheet` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `user_access_token`  or  `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" Learn more about how to obtain and use access_token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Path parameters
| Parameter              | Type   | Required | Description                                | 
| ----------------- | ------ | ---- | ------------------------------------------------------------ | ----------- |
| spreadsheetToken  | string | Yes   | token of the  sheet , to obtain the token, see Sheets Developer's Guide| 
### Request body
| Parameter              | Type   | Required | Description                                | 
| ----------------- | ------ | ---- | ------------------------------------------------------------ | ----------- |
|requests||Yes|Request
|   protectId         | string | Yes   | Protected range ID, obtained through the Obtain sheet metadata API                                          |
|   ∟dimension         |        | No   | Row and column protection information                                                 |
|     ∟sheetId          | string | Yes   | sheetId                                                      |
|     ∟startIndex       | int    | Yes   | Protected row and column start index, starts from 1.|
|     ∟endIndex         | int    | Yes   | Protected row and column end index, starts from 1.|
|     ∟majorDimension   | string | Yes   | Protected range dimensions for the protected range ID; COLUMNS indicates protected columns and ROWS indicates protected row.| Request body   |
|   ∟editors           |        | No   | Users who can edit the protected range                                         |
|     ∟addEditors       |        | No   | The list of users to add, users must have doc editing permissions.|
|       ∟memberType | string | Yes   | User type, supports userId, openId, and unionId.|
|       ∟memberId   | string | Yes   | User ID of the selected user type                                         |
|     ∟delEditors       |        | No   | List of users to delete                                         |
|       ∟memberType | string | Yes   | User type, supports userId, openId, and unionId.|
|       ∟memberId   | string | Yes   | User ID of the selected user type                                         |
|   ∟lockInfo          | string | No   | Protection description                                                     | ### Request body example
```json
{
    "requests": [
        {
            "protectId": "***",
            "dimension": {
                "majorDimension": "***",
                "sheetId": "***",
                "startIndex": 0,
                "endIndex": 0
            },
            "editors": {
                "addEditors": [
                    {
                        "memberType": "userId",
                        "memberId": "****"
                    }
                ],
                "delEditors": [
                    {
                        "memberType": "userId",
                        "memberId": "****"
                    }
                ]
            },
            "lockInfo": "****"
        }
    ]
}
```
### cURL Request example
```
curl --location --request POST 'https://open.larksuite.com/open-apis/sheets/v2/spreadsheets/shtcngNygNfuqhxTBf588jwgWbJ/protected_range_batch_update' \
--header 'Authorization: Bearer t-e346617a4acfc3a11d4ed24dca0d0c0fc8e0067e' \
--header 'Content-Type: application/json' \
--data-raw '{
    "requests": [
        {
            "protectId": "6947942538267541505",
            "dimension": {
                "majorDimension": "ROWS",
                "sheetId": "Q7PlXT",
                "startIndex": 2,
                "endIndex": 4
            },
            "editors": {
                "addEditors": [
                    {
                        "memberType": "userId",
                        "memberId": "667338922291111404"
                    }
                },
                "delEditors": [
                    {
                        "memberType": "userId",
                        "memberId": "667338922291122404"
                    }
                ]
            },
            "lockInfo": "1234"
        }
    ]
}'
```
## Response
### Response body
| Parameter              |Type| Description      |
| ----------------- |-----| ---------------------------------- |
|replies|array|Response|
|   ∟sheetId           |string| sheet ID                          |
|   ∟dimension         || Successfully modified protected row and column information             |
|     ∟sheetId          |string| sheetId                            |
|     ∟startIndex       |int| Protected row and column start index, starts from 1      .|
|     ∟endIndex         |int| Protected row and column end index, starts from 1      .|
|     ∟majorDimension   |string| Protected range dimensions                     |
|   ∟editors           | Users who can edit the protected range               |
|     ∟addEditors       |array| List of successfully added users               |
|       ∟memberType |string| User type                           |
|       ∟memberId   |string| User ID of the selected user type               |
|     ∟delEditors       |array| List of successfully deleted users               |
|       ∟memberType |string| User type                           |
|       ∟memberId   |string| User ID                of the selected user type|
|   ∟lockInfo          |string| Successfully modified protection description                 |
###  Response body example
```json
{
    "code": 0,
    "msg": "Success",
    "data": {
        "replies": [
            {
                "protectId": "***",
                "dimension": {
                    "sheetId": "***",
                    "startIndex": 0,
                    "endIndex": 0,
                    "majorDimension": "ROWS"
                },
                "editors": {
                    "addEditors": [
                        {
                            "memberType": "userId",
                            "memberId": "*"
                        }
                    },
                    "delEditors": []
                },
                "lockInfo": "Info11",
                "sheetId": "abb54d"
            }
        ]
    }
}
```

### Error code

For details, refer to Server-side error codes.
