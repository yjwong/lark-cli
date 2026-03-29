---
title: "Delete Protection Scopes"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uYTM5YjL2ETO24iNxkjN"
method: "DELETE"
api_path: "https://open.larksuite.com/open-apis/sheets/v2/spreadsheets/:spreadsheetToken/protected_range_batch_del"
section: "Docs"
scopes:
  - "drive:drive"
  - "sheets:spreadsheet"
updated: "1647001385000"
---

# Delete protected range

Use this API to delete a protected range based on the protected range ID. Each operation can delete up to 10 IDs.
## Request
| HTTP URL | https://open.larksuite.com/open-apis/sheets/v2/spreadsheets/:spreadsheetToken/protected_range_batch_del |
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
| spreadsheetToken | string        | Yes   |Token of the sheet, to obtain the token, see Sheets Developer's Guide.|
### Request body
| Parameter             | Type          | Required | Description                                |
| ---------------- | ------------- | ---- | ------------------------------------------------------------ |
| protectIds       | array | Yes   | ID of the protected range to delete, obtained through the Obtain sheet metadata API.|
### Request body example

```json
{
    "protectIds": ["******"]
}
```
###  cURL Request example
```
curl --location --request DELETE 'https://open.larksuite.com/open-apis/sheets/v2/spreadsheets/shtcngNygNfuqhxTBf588jwgWbJ/protected_range_batch_del' \
--header 'Authorization: Bearer t-e346617a4acfc3a11d4ed24dca0d0c0fc8e0067e' \
--header 'Content-Type: application/json' \
--data-raw '{
    "protectIds": ["6947942538267541505","6946456074476339204"]
}'
```
## Response
### Response body
  | Parameter          |Type| Description                 |
| ------------- |-----| -------------------- |
| delProtectIds |array| ID of successfully deleted protected range | ### Response body example
```json
{
    "code": 0,
    "msg": "Success",
    "data": {
        "delProtectIds": [
            "******"
        ]
    }
}
```

### Error code

For details, refer to Server-side error codes.
