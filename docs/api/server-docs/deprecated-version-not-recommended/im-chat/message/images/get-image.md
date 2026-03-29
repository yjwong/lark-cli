---
title: "Get Image"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uYzN5QjL2cTO04iN3kDN"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/image/v4/get?image_key=24383920-9321-4ecd-8b33-bf8ce74e84c8"
section: "Deprecated Version (Not Recommended)"
updated: "1626158295000"
---

# Get image
Get the image content according to the image_key.

> Bot capability needs to be enabled. Bot can only get images uploaded by itself

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/image/v4/get?image_key=24383920-9321-4ecd-8b33-bf8ce74e84c8 |
| HTTP Method | GET | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" Learn more about obtaining and using access_token. |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Query parameters
|Parameter|Type|Mandatory|Description|Default|Demo|
-- | -- | -- | -- | -- | --
image_key|string|Yes|After uploading successfully, you can get a image_key.||img_fdeb9536-8ec4-485f-988b-6ebd338ad47g| ## Response
### Response body  
Get image binary data when success
|Parameters|Type|Description|
|-|-|-|
code |int| Error code, anything but 0 indicates failure
msg |string| Error code description.
### Error code

For details, please refer to: Service-side error codes
