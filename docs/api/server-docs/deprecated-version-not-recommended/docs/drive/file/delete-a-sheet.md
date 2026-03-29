---
title: "Delete a Sheet"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uUTNzUjL1UzM14SN1MTN/delete-sheet"
method: "DELETE"
api_path: "https://open.larksuite.com/open-apis/drive/explorer/v2/file/spreadsheets/:spreadsheetToken"
section: "Deprecated Version (Not Recommended)"
scopes:
  - "drive:drive"
  - "sheets:spreadsheet"
updated: "1703488391000"
---

# Delete a sheet

This API is used to delete a specified sheet based on a spreadsheetToken. 

> A document can be deleted only by the owner of the document. Once deleted, the document is moved to Trash.

> **Note:** This API doesn't support concurrent calls, and supports up to 5 queries per second (QPS).

## Request
| HTTP URL | https://open.larksuite.com/open-apis/drive/explorer/v2/file/spreadsheets/:spreadsheetToken |
| --- | --- |
| HTTP Method | DELETE |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `drive:drive` `sheets:spreadsheet` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" Learn more about how to obtain and use access_token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ::: note
::: html
Before using the `tenant_access_token`, make sure that the app is the owner of the document. Otherwise, an error indicating no permissions will occur.

### Path parameters
| `spreadsheetToken` | `string` | The spreadsheet   token. For more information about how to obtain the token, see  Overview. |
| --- | --- | --- | ## Response
### Response body
|Parameter|Description|
|--|--|
|id|Sheet ID, which is a "string"|
|result|Deletion result| ### Response body example
```json
{
    "code":0,
    "msg":"Success",
    "data":
    {
        "id":"id string",
        "result":true
    }
}
```
### Error code

For details, refer to Server-side error codes.
