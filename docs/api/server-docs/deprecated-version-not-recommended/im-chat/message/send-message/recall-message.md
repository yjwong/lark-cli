---
title: "Recall Message"
url: "https://open.larksuite.com/document/ukTMukTMukTM/ukjN1UjL5YTN14SO2UTN"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/message/v4/recall/"
section: "Deprecated Version (Not Recommended)"
updated: "1626158292000"
---

# Recall message
Recall a message.

> - Attention should be paid to calling this interface:
> - The application needs to enable the robot ability
> - The message to be recalled cannot be sent out more than 24 hours
> - You can only recall messages sent as apps (messages sent by the robot itself)
> - **Unable to recall** messages sent through the 「Batch send messages」

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/message/v4/recall/ |
| HTTP Method | POST | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value Format**: "Bearer `access_token`" **Example Value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" Learn more about obtaining and using access_token. |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Request body
|Parameter|Type|Mandatory|Description|Default|Demo|
-- | -- | -- | -- | -- | --
message_id|string|Yes|message_id||om_f32a6454a616f4123b72fd9ecd976c41| ### Request body example
```json
{
   "message_id": "om_f32a6454a616f4123b72fd9ecd976c41"
}
```
## Response
### Response body
|Parameters|Type|Description|
|-|-|-|
code |int| Error code, anything but 0 indicates failure
msg |string| Error code description.

### Response body example
```json
{
    "code": 0,
    "msg": "ok"
}
```
### Error code

For details, please refer to: Service-side error codes
