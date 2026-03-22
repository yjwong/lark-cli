---
title: "Create a New Folder"
url: "https://open.larksuite.com/document/ukTMukTMukTM/ukTNzUjL5UzM14SO1MTN"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/drive/explorer/v2/folder/:folderToken"
section: "Deprecated Version (Not Recommended)"
scopes:
  - "drive:drive"
updated: "1703488376000"
---

# Create a folder

This API is used to create a folder in the current folder based on a folderToken.

> This API doesn't allow you to concurrently create multiple documents, and supports up to 5 queries per second (QPS) and 10,000 queries per day.

## Request
| HTTP URL | https://open.larksuite.com/open-apis/drive/explorer/v2/folder/:folderToken |
| --- | --- |
| HTTP Method | POST |
| Supported app types | custom,isv |
| Required scopes | `drive:drive` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" Learn more about how to obtain and use access_token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ::: note
For more information about how to call the AccessToken in the Docs API, see  Get started with the Docs API.

### Path parameters
|Parameter|Type|Required|Description|
|--|-----|--|----|
|folderToken|string|Yes|Token of the folder. For more information about how to obtain the token, see  Overview.| ### Request body
|Parameter|Type|Required|Description|
|--|-----|--|----|
|title|string|Yes|Title of the folder| ### Request body example
```json
{
	"title": "string"
}
```
### Response

### Response body
|Parameter|Description|
|--|--|
|url|URL of the new folder|
|revision|Version number of the new folder|
|token|Token of the new folder| ### Response body example

```json
{
	"code": 0,
	"msg": "Success",
	"data": {
		"url": "string",
    	"revision": 0,
    	"token": "string"
    }
}
```

### Error code

For details, refer to Server-side error codes.
