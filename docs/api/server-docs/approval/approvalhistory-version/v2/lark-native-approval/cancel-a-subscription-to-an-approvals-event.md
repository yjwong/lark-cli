---
title: "Cancel a Subscription to an Approvals Event"
url: "https://open.larksuite.com/document/ukTMukTMukTM/ugDOyUjL4gjM14CO4ITN"
method: "POST"
api_path: "https://www.larksuite.com/approval/openapi/v2/subscription/unsubscribe"
section: "Approval"
updated: "1658309936000"
---

# Unsubscribe to approval events
> This API has been upgraded to enhance its security. We recommend that you migrate to the new version>>

After an app unsubscribes to approval_code, the app can no longer receive event notifications for instances of that approval definition.

## Request
| HTTP URL | https://www.larksuite.com/approval/openapi/v2/subscription/unsubscribe |
| --- | --- |
| HTTP Method | POST | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" Learn more about how to access and use access_token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Request body

| Parameter         | Type           | Required        | Description        |
| --------- | --------------- | -------   | --------- |
| approval_code | string | Yes | Unique ID of the approval definition | ### Request body example

```json
{
	"approval_code":"7C468A54-8745-2245-9675-08B7C63E7A85"
}
````

## Response

### Response body
| Parameter         | Type           | Required        | Description        |
| --------- | ----------|----- | --------- |
|code |int |Yes |Error code. A value other than 0 indicates failure. |
|msg | string |Yes| Return code description|
### Response body example

```json
{
    "code":0,
    "msg":"success"
}
```
