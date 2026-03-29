---
title: "Retrieve Protection Scopes"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uQTM5YjL0ETO24CNxkjN"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/sheets/v2/spreadsheets/:spreadsheetToken/protected_range_batch_get"
section: "Docs"
scopes:
  - "drive:drive"
  - "sheets:spreadsheet"
updated: "1647001378000"
---

# Get protected range

Use this API to get detailed row and column information for a protected range based on the protected range ID. Each operation can query up to 5 IDs.
## Request
| HTTP URL | https://open.larksuite.com/open-apis/sheets/v2/spreadsheets/:spreadsheetToken/protected_range_batch_get |
| --- | --- |
| HTTP Method | GET |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `drive:drive` `sheets:spreadsheet` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `user_access_token`  or  `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" Learn more about how to obtain and use access_token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| spreadsheetToken | string | Spreadsheet token. For more information about how to obtain the token, see Sheets Developer's Guide | ### Query parameters
| Parameter             | Type   | Required | Description                                | 
| ---------------- | ------ | ---- | ------------------------------------------------------------ | 
| protectIds       | string | Yes   | Protected range ID, can be obtained through the Obtain sheet metadata API, separate multiple IDs with comas, for example: xxxID1,xxxID2.| 
| memberType       | string | No   | Returned user type, values: userId (default), openId, or unionId| ### cURL request example
```
curl --location --request GET 'https://open.larksuite.com/open-apis/sheets/v2/spreadsheets/shtcngNygNfuqhxTBf588jwgWbJ/protected_range_batch_get?protectIds=6946456074476339204,6947648349520592923,6947942538267541505&memberType=userId' \
--header 'Authorization: Bearer t-e346617a4acfc3a11d4ed24dca0d0c0fc8e0067e' \
```

## Response  
### Response body
| Parameter                        |Type| Description                                          |
| --------------------------- |-----| --------------------------------------------- |
| protectedRange              |array| Protected rage                                      |
|   ∟protectId            |string| Protected range ID                                    |
|   ∟dimension            ||  Protected range, protects the whole sheet if empty              |
|     ∟sheetId        |string| id                                    of  sheet |
|     ∟startIndex     |int| Protected row and column start index, starts from 1                 .|
|     ∟endIndex       |int| Protected row and column end index, starts from 1                 .|
|     ∟majorDimension |string| Protected range dimensions, COLUMNS indicates protected columns and ROWS indicates protected rows .|
|   ∟sheetId              |string| ID of the sheet                                     |
|   ∟lockInfo             |string| Protection description                                      |
|   ∟editors              || User information                                      |
|     ∟users	|array|User information list|
|       ∟memberType     |string| User type                                      |
|       ∟memberId       |string| User ID                                        | ### Response body example

```json
{
    "code": 0,
    "msg": "Success",
    "data": {
        "protectedRanges": [
            {
                "protectId": "*****",
                "dimension": {
                    "sheetId": "***",
                    "startIndex": 0,
                    "endIndex": 0,
                    "majorDimension": "COLUMNS"
                },
                "editors": {
                    "users": [
                        {
                            "memberType": "userId",
                            "memberId": "***"
                        }
                    ]
                },
                "sheetId": "***"
            }
        ]
    }
}
```

### Error code

For details, refer to Server-side error codes.
