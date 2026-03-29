---
title: "Obtain Document Meta"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uczN3UjL3czN14yN3cTN"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/doc/v2/meta/:docToken"
section: "Deprecated Version (Not Recommended)"
updated: "1695369909000"
---

# Obtain metadata of a document
> **Note:** This API does not support the Upgraded Docs. If you want to get the metadata of the Upgraded Docs, please use Obtain metadata.
This API is used to obtain metadata based on a docToken.

## Request
| HTTP URL | https://open.larksuite.com/open-apis/doc/v2/meta/:docToken |
| --- | --- |
| HTTP Method | GET |
| Required scopes Enable any scope from the list | View, comment, edit, and manage all files in My Space View, comment, and download all files in My Space View the metadata of files in My Space View, comment, edit, and manage Docs View, comment, and export Docs | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `user_access_token` or `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" Learn more about how to access and use access_token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ::: note
For more information about how to call the AccessToken in the Docs API, see Get started with the Docs API.

### Path parameters
|Parameter|Type|Required|Description|
|--|-----|--|----|----|
|docToken|string|Yes|Document token. For more information about how to obtain the token, see How to get the token of docs resources.| ### Curl request demo
```
curl -H 'Authorization: Bearer u-s12okJw4R1VCZLWhk9Zyzg' 'https://open.larksuite.com/open-apis/doc/v2/meta/doccnilYPZU5b34ow4ca7aNoU6a' 
```

## Response
### Response body
|Parameter|Type|Description|
|--|--|--|
|create_date|string|Creation date|
|create_time|integer|Creation timestamp|
|creator|string|Creator open_id|
|create_user_name|string|Creator username|
|delete_flag|integer|Deletion flag, 0 indicates normal access and not deleted, 1 indicates the object is in the trash, 2 indicates the object is permanently deleted|
|edit_time|integer|Last edit timestamp|
|edit_user_name|string|Username of last editor|
|is_external|bool|Whether it is an external document|
|is_pined|bool|Whether it is a quick access in the API caller directory |
|is_stared|bool|Whether it is a favorite in the API caller directory |
|obj_type|string|Document type, always "doc"|
|owner|string|open_id of the current owner|
|owner_user_name|string|Username of the current owner|
|server_time|integer|Server timestamp during request processing|
|tenant_id|string|ID of the tenant of the document |
|title|string|Document name|
|type|integer|Document type, always "2"|
|url|string|Document URL| ### Response body example
```json
{
    "code": 0,
    "msg": "Success",
    "data": {
        "create_date": "integer string",
        "create_time": integer,
        "creator": "string",
        "create_user_name": "string",
        "delete_flag": integer,
        "edit_time": integer,
        "edit_user_name": "string",
        "is_external": bool,
        "is_pined": bool,
        "is_stared": bool,
        "obj_type": "doc",
        "owner": "string",
        "owner_user_name": "string",
        "server_time": integer,
        "tenant_id": "interger string",
        "title": "string",
        "type": 2,
        "url": "string"
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
