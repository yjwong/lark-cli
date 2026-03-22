---
title: "Cancel Subscribe"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/mail-v1/user_mailbox-event/unsubscribe"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/mail/v1/user_mailboxes/:user_mailbox_id/event/unsubscribe"
service: "mail-v1"
resource: "user_mailbox-event"
section: "Email"
rate_limit: "1 per second"
scopes:
  - "mail:event"
updated: "1745840320000"
---

# Cancel Subscribe

Cancel Subscribe

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/mail/v1/user_mailboxes/:user_mailbox_id/event/unsubscribe |
| HTTP Method | POST |
| Rate Limit | 1 per second |
| Supported app types | custom |
| Required scopes | `mail:event` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `user_mailbox_id` | `string` | User email address, use me when using user_access_token **Example value**: "user@xxx.xx or me" | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `event_type` | `int` | Yes | Event type **Example value**: 1 **Optional values are**:  - `1`: Mail message related events  **Data validation rules**: - Value range: `1` ～ `1` | ### Request body example

{
    "event_type": 1
}

## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {}
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 1230001 | Parameter Error | Please retry after modifying the parameters. |
| 500 | 1230003 | Internal Error | Please try again later. |
| 403 | 1230002 | No Permission | After becoming a member of the public mailbox, invoke this interface. |
