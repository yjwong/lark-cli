---
title: "Unsubscribe from approval events"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/approval-v4/approval/unsubscribe"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/approval/v4/approvals/:approval_code/unsubscribe"
service: "approval-v4"
resource: "approval"
section: "Approval"
scopes:
  - "approval:approval"
updated: "1676430228000"
---

# Unsubscribe from review events

After unsubscribing approval_code, you can no longer receive event notifications for the corresponding instance of the review definition

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/approval/v4/approvals/:approval_code/unsubscribe |
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
| 400 | 1390007 | subscription existed | Subscription already exists |
| 400 | 1395001 | There have been some errors. Please try again later | There was an error with the service, please try again later |
