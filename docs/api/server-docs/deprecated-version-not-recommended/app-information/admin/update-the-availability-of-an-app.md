---
title: "Update the Availability of an App"
url: "https://open.larksuite.com/document/ukTMukTMukTM/ucDN3UjL3QzN14yN0cTN"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/application/v3/app/update_visibility"
section: "Deprecated Version (Not Recommended)"
scopes:
  - "admin:app.visibility"
updated: "1682583669000"
---

# Update app availability

Update the visibility of the current Custom App or installed Store App in the enterprise, including available and disabled people. The update takes effect immediately online.

> **Note:** When adding a user or department through the api, it is necessary to judge in advance whether the corresponding user or department is already in the black list. If it is already in the black list, even if the user or department is added to the white list, the user or department cannot see the application, that is to say the disabled list has priority over the available list.

## Request
| HTTP URL | https://open.larksuite.com/open-apis/application/v3/app/update_visibility |
| --- | --- |
| HTTP Method | POST |
| Supported app types | custom |
| Required scopes | `admin:app.visibility` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" Learn more about how to obtain and use access_token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Request body
|Parameter|Type|Required|Description|
|--|-----|--|----|
|app_id|string|Yes|App ID|
|del_users|-|No|List of users to be deleted. The number of deleted users cannot exceed 500 and only users within the app availability can be deleted. **The deletion of users must be performed after the addition of users**.|
|  ∟open_id|string|No|Either open_id or user_id is required. The user_id takes priority over  open_id  if both are specified.|
|  ∟user_id|string|No||
|add_users|-|No|List of users to be added. The number of added users cannot exceed 500. **The deletion of users must be performed after the addition of users**.|
|  ∟open_id|string|No|Either open_id or user_id is required. The user_id takes priority over open_id if both are specified.|
|  ∟user_id|string|No||
|is_visiable_to_all|int|No|Whether the app is visible to all. 0: No; 1: Yes; not specified: Maintain current status|
|add_departments|string[]|No|List of departments to be added. The number of added departments cannot exceed 500. **The deletion of departments must be performed after the addition of departments**.|
|del_departments|string[]|No|List of departments to be deleted. The number of deleted departments cannot exceed 500. **The deletion of departments must be performed after the addition of departments**.| ### Request body example
```json
{
    "app_id" : "cli_9db45f86b7799104",
    "add_users": [
        {
            "open_id": "ou_d317f090b7258ad0372aa53963cda70d",
            "user_id":"79affdge"
         }
    ],
    "del_users": [
    ],
    "is_visiable_to_all" : 0,
    "add_departments": [
        "od-4b4a6907ad726ea07b27b0d2882b7c65",
        "od-2a0c3396b2cbd9a0befb104eccd1209f"
    ],
    "del_departments": [
    ]
}
```

## Response
### Response body
|Parameter|Description|
|--|--|
|code|Return code. A value other than  0  indicates failure.|
|msg|Return code description|
|data|Returned business information| ###  Response example
```json
{
    "code": 0,
    "msg": "success",
    "data": {}
}
```

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 200 | 50003 | invalid app_id | Check whether the app_id is correct and whether the app is installed for the company. |
