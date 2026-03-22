---
title: "Verify App Admin"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uITN1EjLyUTNx4iM1UTM"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/application/v3/is_user_admin"
section: "App Information"
scopes:
  - "admin:app.admin:check"
  - "admin:app.admin:readonly"
updated: "1646720004000"
---

# Query whether a user has app management permission

This API is used to query whether a user is an app administrator.

## Request
| HTTP URL | https://open.larksuite.com/open-apis/application/v3/is_user_admin |
| --- | --- |
| HTTP Method | GET |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `admin:app.admin:check` `admin:app.admin:readonly` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" Learn more about how to obtain and use access_token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Request parameters
|Parameter|Type|Required|Description|
|--|-----|--|----|
|open_id|string|No|User's open_id. Either open_id or employee_id is required. If both are specified, the  open_id prevails.|
|employee_id|string|No|User's employee_id (equivalent to the user_id in Contacts v3). Either open_id or employee_id is required. If both are specified, the open_id prevails.| ## Response

### Response body
|Parameter|Description|
|--|--|
|code|Return code. A value other than  0  indicates failure.|
|msg|Return code description|
|data|Returned business information|
|  ∟is_app_admin|Whether the user is an administrator. true: Yes, false: No| ### Response example
```json
{ 
    "code": 0, 
    "msg": "ok", 
    "data": { 
        "is_app_admin": false
    } 
}
```
