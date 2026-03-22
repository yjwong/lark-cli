---
title: "Get Document Sharing Settings V2"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uITM3YjLyEzN24iMxcjN"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/drive/permission/v2/public/"
section: "Docs"
updated: "1647001279000"
---

# Obtain the common settings of a document V2

This API is used to obtain the common settings of a document based on a filetoken. 
## Request
| HTTP URL | https://open.larksuite.com/open-apis/drive/permission/v2/public/ |
| --- | --- |
| HTTP Method | POST |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | View, comment, edit, and manage all files in My Space View, comment, and download all files in My Space View the metadata of files in My Space Upload and download files to My Space View and download files in My Space View, comment, edit, and manage Docs View, comment, and export Docs View, comment, edit, and manage Sheets View, comment, and export Sheets | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `user_access_token` or `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" Learn more about how to access and use access_token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ::: note
For more information about how to call the AccessToken in the Docs API, see Get started with the Docs API.

### Request body
|Parameter|Type|Required|Description|
|--|-----|--|----|
|token|string| Yes | Token of the file. For more information about how to obtain the token, see Overview |.
|type|string| Yes | Type of the document. Optional values:  "doc", "sheet" or "isv"|
### Request body example
```json
{
        "token": "doccnBKgoMyY5OMbUG6FioTXuBe",
        "type": "doc"
}
```
## Response
### Response body
|Parameter|Type|Description|
|--|--|----|
|security_entity|string|Settings of the permission to make a copy of, print, export, and copy the document:"anyone_can_view" - All users who can access this document"anyone_can_edit" - Users who have the edit permission|
|comment_entity|string|Settings of the permission to comment on the document:"anyone_can_view" - All users who can access this document"anyone_can_edit" - Users who have the edit permission| 
|share_entity|string|Indicates who can add and manage collaborators:"anyone" - All users who can read or edit the document"same_tenant" - All users who can read or edit the document in the organization"only_me" - Only me| Request body |
|link_share_entity|string|Settings of link sharing:"tenant_readable" - Users who obtain the link in the organization can read the document."tenant_editable" - Users who obtain the link in the organization can edit the document."anyone_readable" - All users who obtain the link can read the document."anyone_editable" - All users who obtain the link can edit the document.| 
|external_access|bool|Indicates whether sharing out of the tenant is allowed.| 
|invite_external|bool|Indicates whether non-owners are allowed to invite external participants.| 
|permission_version|string|Version number of the permission.|
### Response body example
```json
{
    "code": 0,
    "data": {
        "security_entity": "anyone_can_view",
        "comment_entity": "anyone_can_view",
        "share_entity":"only_me",
        "link_share_entity":"tenant_editable",
        "external_access":false,
        "invite_external":false,
        "permission_version":"1024",
    },
    "msg": "Success"
}
```
### Error code

For details, refer to Server-side error codes.
