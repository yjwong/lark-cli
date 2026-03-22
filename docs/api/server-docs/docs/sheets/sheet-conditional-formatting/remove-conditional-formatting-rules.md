---
title: "Remove Conditional Formatting Rules"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uATMzUjLwEzM14CMxMTN/conditionformat/condition-format-delete"
method: "DELETE"
api_path: "https://open.larksuite.com/open-apis/sheets/v2/spreadsheets/:spreadsheetToken/condition_formats/batch_delete"
section: "Docs"
scopes:
  - "drive:drive"
  - "sheets:spreadsheet"
updated: "1647001353000"
---

# Delete conditional formatting

Use this API to delete an existing conditional formatting. You can delete up to 10 conditional formats at once. Each conditional formatting deletion operation returns success or failed. Parameter verification information is returned for failed operations.

## Request
| HTTP URL | https://open.larksuite.com/open-apis/sheets/v2/spreadsheets/:spreadsheetToken/condition_formats/batch_delete |
| --- | --- |
| HTTP Method | DELETE |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `drive:drive` `sheets:spreadsheet` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `user_access_token`  or  `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" Learn more about how to obtain and use access_token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Path parameters
| Parameter             | Type          | Required | Description                                |
| ---------------- | ------------- | ---- | ------------------------------------------------------------ | 
| spreadsheetToken | string        | Yes   | token of the  sheet , to obtain the token see the Sheets Developer's Guide|
### Request body
| Parameter       |Type| Description       |
| -------- |-----| -------- |
|sheet_cf_ids| |Spreadsheet conditional formatting ID|
|   ∟sheet_id |string| sheet ID |
|   ∟cf_id    |string| Conditional formatting ID   |
### Request body example
```json
{
    "sheet_cf_ids": [
        {
            "sheet_id": "40a7b0",
            "cf_id": "6hP6Dj6gsd"
        }
    ]
}
```
###  cURL Request example
```
curl --location --request DELETE 'https://open.larksuite.com/open-apis/sheets/v2/spreadsheets/shtcngNygNfuqhxTBf588jwgWbJ/condition_formats/batch_delete' \
--header 'Authorization: Bearer t-e346617a4acfc3a11d4ed24dca0d0c0fc8e0067e' \
--header 'Content-Type: application/json' \
--data-raw '{
    "sheet_cf_ids": [
        {
            "sheet_id": "Q7PlXT",
            "cf_id": "KjRm0JyS1P"
        }
    ]
}'
```
## Response
### Response body
| Parameter       |Type| Description                           |
| -------- |-----| ---------------------------- |
|responses|array|Response|
|   ∟sheet_id |string | sheet ID                     |
|   ∟cf_id    |string| Conditional formatting ID                       |
|   ∟res_code |int| Conditional formatting deletion status code, 0 indicates success and a non-0 value indicates failure       |
|   ∟res_msg  |string| Status information returned for conditional formatting deletion, empty indicates success and any content indicates the cause of failure |
### Response body example
```json
{
    "code": 0,
    "data": {
        "responses": [
            {
                "cf_id": "6hP6Dj6gsd",
                "res_code": 555554047,
                "res_msg": "cfId not exist",
                "sheet_id": "40a7b0"
            }
        ]
    },
    "msg": "Success"
}
```
### Error code

For details, refer to Server-side error codes.
