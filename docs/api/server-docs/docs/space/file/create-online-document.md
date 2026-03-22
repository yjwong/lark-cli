---
title: "Create online document"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uQTNzUjL0UzM14CN1MTN"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/drive/explorer/v2/file/:folderToken"
section: "Docs"
scopes:
  - "drive:drive"
  - "docs:doc"
  - "sheets:spreadsheet"
updated: "1698998305000"
---

# Create online document

Create a document, sheet, or bitable in user-defined folders in the user's Space.

If the target folder is My Space, the newly created document will be listed in the "Owned by Me" list of "My Space".

> The maximum number of single-layer nodes in a folder in the cloud space is 1500. If the limit is exceeded, the create a document interface will return a failure. If there is such a requirement, you can consider creating a new document in a different folder.
> 
> This API has been upgraded to enhance its security. If you need to create a document you can try new version, if you need to create a sheet you can try new version.

> **Note:** This API doesn't allow you to concurrently create multiple documents, and supports up to 5 queries per second (QPS) and 10,000 queries per day.

## Request
| HTTP URL | https://open.larksuite.com/open-apis/drive/explorer/v2/file/:folderToken |
| --- | --- |
| HTTP Method | POST |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `drive:drive` `docs:doc` `sheets:spreadsheet` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" Learn more about how to obtain and use access_token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ::: note
For more information about how to call the AccessToken in the Docs API, see  Get started with the Docs API.

### Path parameters
| `folderToken` | `string` | Token of the folder, which is used to create a document in this folder. For more information about how to obtain the token, see How to get the token of docs resources. |
| --- | --- | --- | ### Request body
|Parameter|Type|Required|Description|
|--|--|--|------|
|title|string|Yes|Title of the document to create. Note: When the type field is "doc", this field is unavailable and optional and the request is filtered. If you need to create a titled document, you can call the  Create a document  API.|||
|type|string|Yes|Type of the document to create. Optional values:   "doc" ,  "sheet"  or  "bitable".||| ### Request body example
```json
{
   "title":"test sheet",
   "type":"sheet"
}
```

## Response
### Response body
|Parameter|Description|
|--|--|
|url|URL of the new document|
|token|Token of the new document|
|revision|Version number of the new document | ### Response body example
```json
{
   "code":0,
   "msg":"Success",
   "data":{
      "url":"https://example.larksuite.com/sheets/shtcnOko1Ad0HU48HH8KHabcef",
      "token":"shtcnOko1Ad0HU48HH8KHabcef",
      "revision":0
   }
}
```
### Error code
| Error code | Description | Troubleshooting suggestions |
| --- | --- | --- |
| 91201 | FAILED | Failed to process, try again later. |
| 91202 | PARAMERR | parameter error, check whether the parameters are correct, such as: `type`, `fileToken`, `dstFolderToken`. |
| 91203 | NOTEXIST | Please check whether the request parameters are correct, such as: `type` and `fileToken` match. |
| 91204 | FORBIDDEN | Check the permissions of the current account to documents and folders. Refer to Access Process Authorization |
| 91205 | DELETED | The source file has been deleted, check if it still exists. |
| 91206 | OUT_OF_LIMIT | Limit exceeded. |
| 91207 | DUPLICATE | Duplicate record. |
| 91208 | REVIEW | Content review failed. | For details, refer to Server-side error codes.
