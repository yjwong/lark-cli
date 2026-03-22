---
title: "Subscribe to an Approvals Event "
url: "https://open.larksuite.com/document/ukTMukTMukTM/ucDOyUjL3gjM14yN4ITN"
method: "POST"
api_path: "https://www.larksuite.com/approval/openapi/v2/subscription/subscribe"
section: "Approval"
updated: "1658309933000"
---

# Subscribe to approval events
> This API has been upgraded to enhance its security. We recommend that you migrate to the new version>>

After an app subscribes to approval_code, the app can receive event notifications for instances of that approval definition. The app only needs to subscribe once.

If the app no longer wants to receive approval events, it can use the unsubscription API to unsubscribe so that no more messages will be pushed to the app.

Both subscription and unsubscription are specific to apps. Multiple apps can subscribe to the same approval_code at the same time, and each of the apps can receive approval events.

## Request
| HTTP URL | https://www.larksuite.com/approval/openapi/v2/subscription/subscribe |
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
