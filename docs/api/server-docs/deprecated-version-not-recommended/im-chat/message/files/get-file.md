---
title: "Get File"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uMDN4UjLzQDO14yM0gTN"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/open-file/v1/get?file_key=file_36r377cb-c6h2-4b6d-ag67-0ac3e796008g"
section: "Deprecated Version (Not Recommended)"
updated: "1626158296000"
---

# Get file
Get the file content according to the file_key. Currently could only fetch the file which was send by user to bot.

> Bot capability needs to be enabled. Bot can only get files uploaded by itself

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/open-file/v1/get?file_key=file_36r377cb-c6h2-4b6d-ag67-0ac3e796008g |
| HTTP Method | GET | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" Learn more about obtaining and using access_token. |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Query parameters
|Parameter|Type|Mandatory|Description|Default|Demo|
-- | -- | -- | -- | -- | --
file_key|string|Yes|After uploading successfully, you can get a file_key.||file_36r377cb-c6h2-4b6d-ag67-0ac3e796008g| ## Response
### Response body
Get file binary data when success
|Parameters|Type|Description|
|-|-|-|
code |int| Error code, anything but 0 indicates failure
msg |string| Error code description 
### Error code

For details, please refer to: Service-side error codes
