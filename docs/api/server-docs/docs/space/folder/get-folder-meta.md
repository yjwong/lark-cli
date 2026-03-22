---
title: "Get Folder Meta"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uAjNzUjLwYzM14CM2MTN"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/drive/explorer/v2/folder/:folderToken/meta"
section: "Docs"
scopes:
  - "drive:drive"
  - "drive:drive.metadata:readonly"
updated: "1691028139000"
---

# Obtain meta information of a folder

This API is used to obtain the meta information of the current folder based on a folderToken.

## Request
| HTTP URL | https://open.larksuite.com/open-apis/drive/explorer/v2/folder/:folderToken/meta |
| --- | --- |
| HTTP Method | GET |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `drive:drive` `drive:drive.metadata:readonly` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793aabcef" Learn more about how to access and use access_token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ::: note
For more information about how to call the AccessToken in the Docs API, see  Get started with the Docs API.

### Path parameters
| `folderToken` | `string` | Token of the folder. For more information about how to obtain the token, see How to get the token of docs resources. |
| --- | --- | --- | ## Response
  
### Response body
|Parameter|Description|
|--|--|
|id|ID of the folder|
|name|Title of the folder|
|token|Token of the folder|
|createUid|Creator of the folder id|
|editUid|Last editor of the folder id|
|parentId|Parent directory of the folder id|
|ownUid|If it is a personal folder, this parameter indicates the ID of the folder owner. If it is a shared folder, this parameter indicates a folder tree.id| ### Response body example
```json
{
	"code": 0,
	"msg": "Success",
	"data": {
      "id": "7110173013420512356",
      "name": "name",
      "token": "nodbcbHUdOsS613xVzTzFEabcef",
      "createUid": "7103496998321312356",
      "editUid": "7103496998321312356",
      "parentId": "0",
      "ownUid": "7110173013420512356"
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
