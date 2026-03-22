---
title: "Delete a Doc"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uATM2UjLwEjN14CMxYTN"
method: "DELETE"
api_path: "https://open.larksuite.com/open-apis/drive/explorer/v2/file/docs/:docToken"
section: "Deprecated Version (Not Recommended)"
updated: "1703488387000"
---

# Delete a document

This API is used to delete a specified document from Docs based on a docToken.

> A document can be deleted only by the owner of the document. Once deleted, the document is moved to Trash.

> **Note:** This API doesn't support concurrent calls, and supports up to 5 queries per second (QPS).

## Request
| HTTP URL | https://open.larksuite.com/open-apis/drive/explorer/v2/file/docs/:docToken |
| --- | --- |
| HTTP Method | DELETE |
| Required scopes Enable any scope from the list | View, comment, edit, and manage all files in My Space View, comment, edit, and manage Docs | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" Learn more about how to access and use access_token | ::: note
::: html
Before using the `tenant_access_token`, make sure that the app is the owner of the document. Otherwise, an error indicating no permissions will occur.

### Path parameters
| `docToken` | `string` | Token of the document. For more information about how to obtain the token, see Overview. |
| --- | --- | --- | ## Response
### Response body
|Parameter|Description|
|--|--|
|id|Document ID, which is a "string"|
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
