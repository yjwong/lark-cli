---
title: "Transfer Ownership"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uQzNzUjL0czM14CN3MTN"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/drive/permission/member/transfer"
section: "Docs"
updated: "1647001276000"
---

# Transfer owner

This API is used to transfer a document owner based on document information and user information. 
## Request
| HTTP URL | https://open.larksuite.com/open-apis/drive/permission/member/transfer |
| --- | --- |
| HTTP Method | POST |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | View, comment, edit, and manage all files in My Space Upload and download files to My Space View, comment, edit, and manage Docs View, comment, edit, and manage Sheets | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `user_access_token` or `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" Learn more about how to access and use access_token |
| Content-Type | string | Yes | **Fixed value**: ："application/json; charset=utf-8" | ::: note
For more information about how to call the AccessToken in the Docs API, see Get started with the Docs API.

### Request body
|Parameter|Type|Required|Description|
|--|-----|--|----|
|token|string| Yes | Token of the file. For more information about how to obtain the token, see Overview |.
|type|string| Yes | Type of the document. Optional values:   "doc"  or  "sheet" or "bitable"  or "file"|
|owner|| Yes | New document owner to whom you need to transfer the document. |
|&ensp;∟member_type|string|Yes | User type. Optional values: **email, openid, and userid**|
|&ensp;∟member_id|string|Yes | Value that corresponds to the user type. For more information, see How do I obtain a User ID, Open ID, and Union ID? |
|remove_old_owner|bool|No | If the value is true, the permissions of the previous owner are deleted after the owner is transferred. The default value is false|.
|cancel_notify|bool|No | If the value is true, the new owner is not notified. The default value is false|.
### Request body example
```json
{
    "type": "string", // "doc" or "sheet" or "file",
    "token": "string",
    "owner": {  
         "member_type": "openid",
         "member_id": "string"
    },
   "remove_old_owner": false,
   "cancel_notify": false
}
```
## Response
### Response body
|Parameter|Description|
|--|--|
|is_success|Indicates whether the request is successful. | 
|type|Type of the document. Optional values:  "doc" or "sheet" or "file"| 
|token| token| of the document 
|owner|Current owner of the document | 
|&ensp;∟member_type|User type. Optional values: email, openid, and userid| 
|&ensp;∟member_id|Value that corresponds to the user type. | 
### Response body example 
```json
{
    "code": 0,
    "msg": "Success",
    "data": {
        "is_success": true,
        "token": "string",
        "type": "doc",
        "owner": {
            "member_type": "openid",
            "member_id": "string"
        }
    }
}
```
### Error code

For details, refer to Server-side error codes.
