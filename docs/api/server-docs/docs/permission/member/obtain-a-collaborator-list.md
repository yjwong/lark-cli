---
title: "Obtain a Collaborator List"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uATN3UjLwUzN14CM1cTN"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/drive/permission/member/list"
section: "Docs"
updated: "1647001273000"
---

# Obtain a collaborator list

This API is used to query collaborators based on a filetoken. The collaborator type can be "user" or "chat." 

> To obtain a collaborator list, you must have the permission to share the document.

## Request

| HTTP URL | https://open.larksuite.com/open-apis/drive/permission/member/list |
| --- | --- |
| HTTP Method | POST |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | View, comment, edit, and manage all files in My Space View, comment, and download all files in My Space Upload and download files to My Space View and download files in My Space View, comment, edit, and manage Docs View, comment, and export Docs View, comment, edit, and manage Sheets View, comment, and export Sheets | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `user_access_token` or `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" Learn more about how to access and use access_token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ::: note
For more information about how to call the AccessToken in the Docs API, see Get started with the Docs API.

### Request body
|Parameter|Type|Required|Description|
|--|-----|--|----|
|token|string| Yes | Token of the file. For more information about how to obtain the token, see Overview |.
|type|string| Yes | Type of the document. Optional values:    "doc"  or  "sheet" or "bitable"  or "file"|
### Request body example
```json
{
	"token": "doccnBKgoMyY5OMbUG6FioTXuBe",
	"type": "doc"
}
```
## Response
### Response body
|Parameter|Description|
|--|--|
|members|Collaborator list|
|&ensp;∟member_type|Collaborator type "user" or "chat"|
|&ensp;∟member_open_id|Collaboratoropenid|
|&ensp;∟member_user_id|User ID of the collaborator. This field is valid only when the member_type field is "user.")|
|&ensp;∟perm|Permissions for the collaborator. Note that **a collaborator with the "edit" permission must have the "view" permission.**)|
### Response body example
```json
{
    "code": 0,
    "msg": "Success",
    "data": {
        "members": [
            {
                "member_type": "chat",
                "member_open_id": "oc_b9be4164d821f466310bc22bb2979cc7",
                "member_user_id": "",
                "perm": "edit"
            },
            {
                "member_type": "user",
                "member_open_id": "ou_65b0affcc6c342a50e4c66f700137b64",
                "member_user_id": "96g3c421",
                "perm": "view"
            },
            {
                "member_type": "user",
                "member_open_id": "ou_b47765834b6bdc18c47a57340f98c0e5",
                "member_user_id": "bg36b129",
                "perm": "edit"
            }
        ]
    }
}
```
### Error code

For details, refer to Server-side error codes.
