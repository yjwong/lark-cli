---
title: "Download Files"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/Attendance//rule/download-file"
section: "Attendance"
updated: "1646322856000"
---

# Download a file
This API is used to download the specified file by a file ID.
## Request
|Facts||
|---|---|
|HTTP URL|https://open.larksuite.com/open-apis/attendance/v1/files/:file_id/download|
|HTTP Method|GET|
|HTTP Content-Type|application/json; charset=utf-8|
|Token requirement|tenant_access_token|
|Required scopes|Export attendance management rules|
### Header
key|value
--|--
Authorization|Bearer tenant_access_token
Content-Type|application/json
### Path parameters
|Parameter|Type|Required|Description|
|---|---|---|---|
|file_id|string|Yes|File ID. Example value: "xxxxxb306842b1c189bc5212eefxxxxx"|
## Response
### Response header
|Parameter|Type|Description|
|---|---|---|
|content-type|string|File MIME|
|content-disposition|string|File name|
An HTTP status code of 200 indicates success

Returns the binary data stream of a file
### Error code
|HTTP status code|Error code|Description|Troubleshooting suggestions|
|---|---|---|---|
|400|1220001|Parameter error|Please check if the parameters meet the requirements|
|400|1220002|Tenant doesn't exist|Please check if the tenant_access_token is correct|
|500|1225000|System error|See error information for details|
|500|1227000|Management service system error|See the error information for details|
