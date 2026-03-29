---
title: "Copy a Doc or Sheet"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uYTNzUjL2UzM14iN1MTN"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/drive/explorer/v2/file/copy/files/:fileToken"
section: "Deprecated Version (Not Recommended)"
scopes:
  - "drive:drive"
  - "docs:doc"
  - "sheets:spreadsheet"
updated: "1703488383000"
---

# Copy a document

This API is used to copy a document or sheet to a destination folder based on a file token.
If no folder is specified to store the created document, before calling this API, you can obtain the token of My Space by calling the "Obtain the meta information of a root folder (in My Space)" API in the "Obtain meta information of a folder" document. The copied document will be displayed in "Owned by me" in "My Space".

> **Note:** This API doesn't allow you to concurrently create multiple documents, and supports up to 5 queries per second (QPS) and 10,000 queries per day.

## Request
| HTTP URL | https://open.larksuite.com/open-apis/drive/explorer/v2/file/copy/files/:fileToken |
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
| `fileToken` | `string` | Token of the source file or document that you need to copy. For more information about how to obtain the token, see Overview. |
| --- | --- | --- | ### Request body
|Parameter|Type|Required|Description|
|--|-----|--|----|
|type|string|Yes|Type of the document to create. Optional values: "doc", "sheet", "docx" or "bitable"|||
|dstFolderToken|string|Yes| Token of the destination folder. For more information about how to obtain the token, see  Overview .|||
|dstName|string|Yes|New name of the file to copy |||
|commentNeeded|bool|No|Specifies whether to copy comments. ||| ### Request body example
```json
{
  "type":"objType",
  "dstFolderToken":"string",
  "dstName":"string",
  "commentNeeded":true
}
```

## Response
### Response body
|Parameter|Description|
|--|--|
|folderToken|The  token of the destination folder|
|revision|Version number of the new document |
|token|The  token of the new document|
|type|Type of the new document. Optional values: "doc", "sheet", "docx" or "bitable"|
|url|The  url of the new document| ### Response body example
```json
{
    "code":0,
    "msg":"Success",
    "data":{
        "folderToken":"fldcne0HujIvzDmRF4Pbg0xxxxx",
        "revision":0,
        "token":"shtcnvJ358XqcZq87CCZHdxxxxx",
        "type":"sheet",
        "url":"https://bytedance.larksuite.com/space/sheet/shtcnvJ358XqcZq87CCZHdxxxxx"
    }
}
```
### Error code

For details, refer to Server-side error codes.
