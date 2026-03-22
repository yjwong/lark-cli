---
title: "Obtain Document Content"
url: "https://open.larksuite.com/document/ukTMukTMukTM/ukzNzUjL5czM14SO3MTN"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/doc/v2/:docToken/raw_content"
section: "Deprecated Version (Not Recommended)"
updated: "1695369917000"
---

# Obtain plain text content of a document
> **Note:** This API does not support the Upgraded Docs. If you need to query the raw content of the Upgraded Docs, please use Query Document Raw Content
> .

This API is used to obtain the plain text content (not include rich text content) of a document.

## Request
| HTTP URL | https://open.larksuite.com/open-apis/doc/v2/:docToken/raw_content |
| --- | --- |
| HTTP Method | GET |
| Required scopes Enable any scope from the list | View, comment, edit, and manage all files in My Space View, comment, and download all files in My Space View, comment, edit, and manage Docs View, comment, and export Docs | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `user_access_token` or `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" Learn more about how to access and use access_token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ::: note
For more information about how to call the AccessToken in the Docs API, see Get started with the Docs API.

### Path parameters
**Request parameter description**:  
|Parameter|Type|Required|Description|
|--|-----|--|----|
|docToken|string|Yes|For more information about how to obtain the token, see How to get the token of docs resources.| ### Curl request demo
```
curl -H 'Authorization: Bearer u-s12okJw4R1VCZLWhk9Zyzg' 'https://open.larksuite.com/open-apis/doc/v2/doccnilYPZU5b34ow4ca7aNoU6a/raw_content' 
```

## Response
### Response body example
```json
{
    "code": 0,
    "msg": "Success",
    "data": {
        "content": "string"
    }
}
```

### Error code
| Error code | Description | Troubleshooting suggestions |
| --- | --- | --- |
| 91401 | PARAMERR | Check parameter validity |
| 91402 | NOTEXIST | Check if the token is valid |
| 91403 | FORBIDDEN | Check for document read permissions |
| 91404 | LOGIN_REQUIRED | Need to log in |
| 95001 | internal error | Internal error, please try again later. |
| 95003 | internal error | Internal error, please try again later. |
| 95005 | internal error | Internal error, please try again later. |
| 95006 | Failed | Check if the token is valid |
| 95007 | Failed | Deleted file cannot get document meta information |
| 95008 | FORBIDDEN | Check user's permission for doc and folder |
| 95009 | Failed | Check for document read permissions. Add permissions |
| 95010 | internal error | Internal error, please try again later. |
| 95011 | internal error | Internal error, please try again later. |
| 95017 | Specific error message | Check if the revison is correct |
| 95018 | Specific error message | See specific error messages for details |
| 95023 | revision too old | The Revision is too old, Please use the latest version number |
| 95024 | Failed | Check parameter validity |
| 95053 | this API does not support the Upgraded Docs(docx) | This API does not support the Upgraded Docs(docx) | For details, refer to Server-side error codes.
