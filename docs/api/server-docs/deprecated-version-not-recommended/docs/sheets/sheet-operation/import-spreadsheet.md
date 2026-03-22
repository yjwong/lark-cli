---
title: "Import Spreadsheet"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uATO2YjLwkjN24CM5YjN"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/sheets/v2/import"
section: "Deprecated Version (Not Recommended)"
scopes:
  - "drive:drive"
  - "sheets:spreadsheet"
updated: "1647001464000"
---

# Import spreadsheet
Use this API to import a local spreadsheet to My Space.
## Request
| HTTP URL | https://open.larksuite.com/open-apis/sheets/v2/import |
| --- | --- |
| HTTP Method | POST |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `drive:drive` `sheets:spreadsheet` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `user_access_token`  or  `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" Learn more about how to obtain and use access_token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Request body
| Parameter          | Type       | Required | Description                                                         |
| ------------- | ---------- | ---- | ------------------------------------------------------------ |
| file          | array | Yes   | The file data to import, converted into a byte array. The "xlsx" and "csv" formats are supported and the data cannot exceed 20 MB.  |
| name          | string     | Yes   | The file name with the file extension, such as: "hello.csv" or "hello.xlsx". After the import, sheet titles remove the extensions, so that "hello.xlsx" becomes "hello".           |
| folderToken   | string     | No   | The token of the folder to import to, imports to the root directory by default                        |
### Request body example

```json
{
    "file": [80,75,3,......],
    "name": "name.xlsx",
    "folderToken": "fldxxxxx"
}
```
### cURL Request example
```
curl --location --request POST 'https://open.larksuite.com/open-apis/sheets/v2/import' \
--header 'Authorization: Bearer t-e346617a4acfc3a11d4ed24dca0d0c0fc8e0067e' \
--header 'Content-Type: application/json' \
--data-raw '{
    "file": [49,44,51,44,52,53,44,50,51,50,44,49,10,53,44,53,50,54,44,49,50,44,53,44,51,54,50,51,44,52,10,49,50,51,50,44,53,49,52,44,53,44,50,54,44,51,50,44,50,44,49,51,49,50],
    "name": "test.csv",
    "folderToken": "fldcncFpcEKOpsQhFpZD2VV6kXf"
}'
```
## Response
### Response body
| Parameter   |Type| Description                                             |
| ------ |-----| ------------------------------------------------ |
| ticket |string| Credentials with a one-to-one correspondence with import files, which are used to query file import progress. For details, see Query import result.|
### Response body example

```json
{
    "code": 0,
    "msg": "Success",
    "data": {
        "ticket": "6891610400000000"
    }
}
```
### Error code

For details, refer to Server-side error codes.
