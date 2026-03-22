---
title: "Subscribe to approval events"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/approval-v4/approval/subscribe"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/approval/v4/approvals/:approval_code/subscribe"
service: "approval-v4"
resource: "approval"
section: "Approval"
scopes:
  - "approval:approval"
updated: "1676430228000"
---

# Subscribe to review events

After the app subscribes approval_code, the app can receive event notifications for the corresponding instance of the review definition. The same app only needs to subscribe once, no need to subscribe repeatedly. 
  
 When the application does not want to receive the review event again, it can use the unsubscribe interface to cancel, and no push message will be given to the application after cancellation. 
  
 Subscribing and unsubscribing are application-dimensional, multiple applications can subscribe to the same approval_code at the same time, and each application can receive review events.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/approval/v4/approvals/:approval_code/subscribe |
| HTTP Method | POST |
| Supported app types | custom,isv |
| Required scopes | `approval:approval` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `approval_code` | `string` | Review Definition Unique Identifier **Example value**: "7C468A54-8745-2245-9675-08B7C63E7A85" | ## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {

    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 1390001 | param is invalid | Parameter error |
| 400 | 1390002 | approval code not found | Check that the review definition code is correct |
| 400 | 1395001 | There have been some errors. Please try again later | There was an error with the service, please try again later |
| 400 | 1390007 | subscription existed | Subscription already exists |
