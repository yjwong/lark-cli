---
title: "Subscribe to Events"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uUTNzUjL1UzM14SN1MTN/events"
section: "Docs"
updated: "1647001461000"
---

# Event subscriptions
> To understand the scenarios and configuration process of event subscriptions, see Event subscription overview.

Use this API to subscribe to Docs or Sheets events based on a file token or file type.
-  Note: Currently, you can only subscribe to events of your own files.
-  Required permissions: View, comment, edit, and manage Sheets or View, comment, edit, and manage all Docs files or View, comment, edit, and manage documents.
-  Special note: After calling this API to subscribe, you must go to the Developer Console to add specific events, such as file read or file title changed.

**Request method**: POST  
**Request address**: https://open.larksuite.com/open-apis/drive/v1/files/:**file_token**/subscribe 
**Request header**:
key|value
--|--
Authorization|Bearer user_access_token
Content-Type|application/x-www-form-urlencoded

**Request parameter description**:  
|Parameter|Type|Required|Description|Source|
|--|-----|--|----|----|
|Authorization|string|Yes|user_access_token, obtained through the Obtain user identity or code2session API;tenant_access_token, obtained through the Obtain app identity access token API.Note that the content must include "Bearer".|Request header|
|file_token|string|Yes|File token| URL PATH|
|file_type|string|Yes|File type, "doc" or "sheet"|Query| **Return body**:  
```json
{
	"code": 0,
	"msg": "Success",
	"data": {}
}
```
