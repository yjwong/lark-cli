---
title: "Get Document"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uUDM2YjL1AjN24SNwYjN"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/doc/v2/:docToken/content"
section: "Deprecated Version (Not Recommended)"
updated: "1695369921000"
---

# Obtain rich text content of a document
> **Note:** Before using this API, carefully read Overview and Prepare to integrate the Docs API to understand the calling rules and limitations of the Docs API. This will ensure your Docs data is not lost or corrupted.
> For a definition of the document data structure, see: Document Data Structure Overview. 
> **Note:** This API does not support the Upgraded Docs. If you need to obtain the rich text content of the Upgraded Docs, please call the related API of the Upgraded Docs:
> - Query All Blocks
> - Query Block
> - Query Block Children
This API is used to obtain structured document content.

## Request
| HTTP URL | https://open.larksuite.com/open-apis/doc/v2/:docToken/content |
| --- | --- |
| HTTP Method | GET |
| Required scopes Enable any scope from the list | View, comment, edit, and manage all files in My Space View, comment, and download all files in My Space View, comment, edit, and manage Docs View, comment, and export Docs | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `user_access_token` or `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" Learn more about how to access and use access_token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ::: note
For more information about how to call the AccessToken in the Docs API, see Get started with the Docs API.

### Path parameters

| Parameter | Type | Required | Description | 
| ------------- | ------ | ---- | --------- | 
| docToken      | string | Yes   | For more information about how to obtain the token, see How to get the token of docs resources. | ### Curl request demo
```
curl -H 'Authorization: Bearer u-CMtIuEpUAu7NdYOI4fLNjg' 'https://open.larksuite.com/open-apis/doc/v2/doccng20xjcYZ2C4FNEoHFGQZng/content'
```

## Response
### Response body
| Parameter | Description |
| -------- | --------------------------------- |
| content  | For details, see Document Data Structure |
| revision | Current document version                    |
### Response body example

```json
{
    "code": 0,
    "msg": "Success",
    "data": {
        "content": "{}",
        "revision": 1000  
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
