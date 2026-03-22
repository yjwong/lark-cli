---
title: "Querying if a Collaborator Has a Specific Permission"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uYzN3UjL2czN14iN3cTN"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/drive/permission/member/permitted"
section: "Deprecated Version (Not Recommended)"
updated: "1703488373000"
---

# Check whether the current user has a specific permission on a document

This API is used to check whether the current login user has a specific permission on a document based on a filetoken. 

## Request
| HTTP URL | https://open.larksuite.com/open-apis/drive/permission/member/permitted |
| --- | --- |
| HTTP Method | POST |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | View, comment, edit, and manage all files in My Space View, comment, and download all files in My Space Upload and download files to My Space View and download files in My Space View, comment, edit, and manage Docs View, comment, and export Docs View, comment, edit, and manage Sheets View, comment, and export Sheets | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `user_access_token`  or  `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" Learn more about how to access and use access_token |
| Content-Type | string | Yes | **Fixed value**: ："application/json; charset=utf-8" | ::: note
For more information about how to call the AccessToken in the Docs API, see Get started with the Docs API.

### Request body
|Parameter|Type|Required|Description|
|--|-----|--|----|
|token|string| Yes | Token of the file. For more information about how to obtain the token, see Overview |. 
|type|string| Yes | Type of the document. Optional values:   "doc"  or  "sheet" or "file"| 
|perm|string|Yes | Permission. Optional values: "view" or "edit" or "share"| ### Request body example
```json
{
	"token": "doccnBKgoMyY5OMbUG6FioTXuBe",
	"type": "doc",
 	"perm": "string"
}
```
## Response
### Response body
|Parameter|Description|
|--|--|
|is_permitted|Indicates whether the user has the specified permission. | ### Response body example
```json
{
    "code": 0,
    "msg": "Success",
    "data": {
        "is_permitted": true
    }
}
```
### Error code

For details, refer to Server-side error codes.
