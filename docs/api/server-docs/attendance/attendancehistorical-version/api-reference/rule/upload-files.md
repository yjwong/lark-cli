---
title: "Upload Files"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/Attendance//rule/file_upload"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/attendance/v1/files/upload"
section: "Attendance"
updated: "1646322860000"
---

# Upload a file
This API is used to upload a file and obtain its file ID. This ID can be used as the face_key parameter in the "Modify user settings" API.

## Request
| HTTP URL | https://open.larksuite.com/open-apis/attendance/v1/files/upload |
| --- | --- |
| HTTP Method | POST |
| Required scopes | Write attendance management rules | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" Learn more about how to access and use access_token |
| Content-Type | string | Yes | **Example value**: "multipart/form-data" | ### Query parameters
| `file_name` | `string` | Yes | File name |
| --- | --- | --- | --- | ### Request body
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `file` | `file` | No | File | ### Request body example

```HTTP
Content-Disposition: form-data; name="file"
Content-Type: application/octet-stream
```

## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error code. A value other than 0 indicates failure. |
| `msg` | `string` | Error description |
| `data` | `\-` | \- |
| ∟ `file` | `file` | File |
| ∟ `file_id` | `string` | File ID | ### Response body example

```json
{
    "code": 0,
    "msg": "success",
    "data": {
        "file": {
            "file_id": "6b30e7636a38861bbe02869c726a4612"
        }
    }
}
```

### Error code
|HTTP status code|Error code|Description|Troubleshooting suggestions|
|---|---|---|---|
|400|1220001|Parameter error|Please check if the parameters meet the requirements|
|400|1220002|Tenant doesn't exist|Please check if the tenant_access_token is correct|
|500|1225000|System error|See error information for details|
|500|1227000|Management service system error|See the error information for details|
