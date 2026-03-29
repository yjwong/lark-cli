---
title: "Get Root Folder Meta"
url: "https://open.larksuite.com/document/ukTMukTMukTM/ugTNzUjL4UzM14CO1MTN/get-root-folder-meta"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/drive/explorer/v2/root_folder/meta"
section: "Docs"
scopes:
  - "drive:drive"
  - "drive:drive.metadata:readonly"
updated: "1691028135000"
---

# Obtain the meta information of a root folder in My Space

This API is used to obtain the meta information of "My Space."

## Request
| HTTP URL | https://open.larksuite.com/open-apis/drive/explorer/v2/root_folder/meta |
| --- | --- |
| HTTP Method | GET |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `drive:drive` `drive:drive.metadata:readonly` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793aabcef" Learn more about how to obtain and use access_token | ::: note
For more information about how to call the AccessToken in the Docs API, see  Get started with the Docs API.

## Response
  
### Response body
|Parameter|Description|
|--|--|
|token|Token of the folder|
|id|Token of the folder|
|user_id|ID of the folder owner

### Response body example
```json
{
  "code": 0,
  "msg": "Success",
  "data": {
    "token": "nodbcbHUdOsS613xVzTzFEabcef",
    "id": "7110173013420512356",
    "user_id": "7103496998321312356"
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
