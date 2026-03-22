---
title: "Modify User Settings"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/Attendance//rule/user-setting-modify"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/attendance/v1/user_settings/modify"
section: "Attendance"
updated: "1646322858000"
---

# Modify user settings
This API is used to modify the user settings of employees in your authorization scope, including face photo file IDs.

## Request
| HTTP URL | https://open.larksuite.com/open-apis/attendance/v1/user_settings/modify |
| --- | --- |
| HTTP Method | POST |
| Required scopes | Write attendance management rules | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" Learn more about how to access and use access_token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Query parameters
| `employee_type` | `string` | Yes | User type **Optional values are**: - `employee_id`: Employee ID - `employee_no`: Employee number |
| --- | --- | --- | --- | ### Request body
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `user_setting` | `user_setting` | No | User information |
| ∟ `user_id` | `string` | Yes | User ID |
| ∟ `face_key` | `string` | Yes | Face photo key (use the file ID obtained upon upload) | ### Request body example

```json
{
    "user_setting": {
        "user_id": "61gc44e1",
        "face_key": "1013d6cb59555a26ff3e5f721342a2a7"
    }
}
```

## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error code. A value other than 0 indicates failure. |
| `msg` | `string` | Error description |
| `data` | `\-` | \- |
| ∟ `user_setting` | `user_setting` | User settings |
| ∟ `user_id` | `string` | User ID |
| ∟ `face_key` | `string` | Face photo key | ### Response body example

```json
{
    "code": 0,
    "msg": "success",
    "data": {
        "user_setting": {
            "user_id": "61gc44e1",
            "face_key": "1013d6cb59555a26ff3e5f721342a2a7",
        }
    }
}
```

### Error code
|HTTP status code|Error code|Description|Troubleshooting suggestions|
|---|---|---|---|
|400|1220001|Parameter error|Please check if the parameters meet the requirements|
|400|1220002|Tenant doesn't exist|Please check if the tenant_access_token is correct|
|400|1220004|User doesn't exist or isn't in permission scope|Please check if the user ID is correct|
|400|1220005|No permission scope|Please go to the Attendance Admin to check the data permission scope|
|500|1225000|System error|See error information for details|
|500|1227000|Management service system error|See the error information for details|
