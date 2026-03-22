---
title: "Get Folder Children"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uEjNzUjLxYzM14SM2MTN"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/drive/explorer/v2/folder/:folderToken/children"
section: "Deprecated Version (Not Recommended)"
scopes:
  - "drive:drive"
updated: "1703488380000"
---

# Obtain documents in a folder

This API is used to obtain documents in a folder based on a folderToken . The following types of documents are supported: doc, sheet, file, bitable, docx and folder.

## Request
| HTTP URL | https://open.larksuite.com/open-apis/drive/explorer/v2/folder/:folderToken/children |
| --- | --- |
| HTTP Method | GET |
| Supported app types | custom,isv |
| Required scopes | `drive:drive` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" Learn more about how to access and use access_token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ::: note
For more information about how to call the AccessToken in the Docs API, see  Get started with the Docs API.

### Path parameters
|Parameter|Type|Required|Description|
|--|-----|--|----|
|folderToken|string|Yes|Token of the folder. For more information about how to obtain the token, see  Overview.| ### Query parameters
|Parameter|Type|Required|Description|
|--|-----|--|----|
|types|array|No|Type of the file to query. By default, all children are returned. You can select multiple values for the types parameter. Optional types: doc, sheet, file, bitable, and folder. Example value:  url?types=folder&types=sheet| ### Response body
|Parameter|Description|
|--|--|
|parentToken|Token of the folder|
|children|Files in the folder|
|  ∟token|Token of the file|
|  ∟name|Title of the file|
|  ∟type|Type of the file| ### Response body example

```json
  {
	"code": 0,
	"msg": "Success",
	"data": {
		"parentToken": "token",
		"children": {
			"nodxxxxxxxx": {
				"token": "lq3xxxxxwLfExB23", // Token of the document
				"name": "test_folder_name",
				"type": "folder",
          },
                  "nodxxxxxxxx": {
                          "token": "lq3xxxxxddddB24", // Token of the document
                          "name": "test_sheet_name",
                          "type": "sheet",
                      }
		}
    }
}
```
>  When the request object is My Space with upgraded scopes, the returned data includes folders created by the requester in Shared Space in addition to the files or folders in My Space.
  
### Error code

For details, refer to Server-side error codes.
