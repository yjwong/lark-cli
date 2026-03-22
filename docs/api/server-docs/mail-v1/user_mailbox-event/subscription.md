---
title: "Get Subscription Status"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/mail-v1/user_mailbox-event/subscription"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/mail/v1/user_mailboxes/:user_mailbox_id/event/subscription"
service: "mail-v1"
resource: "user_mailbox-event"
section: "Email"
rate_limit: "1 per second"
scopes:
  - "mail:event"
updated: "1745840316000"
---

# Get Subscription Status

Get Subscription Status

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/mail/v1/user_mailboxes/:user_mailbox_id/event/subscription |
| HTTP Method | GET |
| Rate Limit | 1 per second |
| Supported app types | custom |
| Required scopes | `mail:event` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `user_mailbox_id` | `string` | User email address, use me when using user_access_token **Example value**: "user@xxx.xx or me" | ## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `event_types` | `int[]` | List of subscribed events | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
        "event_types": [
            1
        ]
    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 1230001 | Parameter Error | Please retry after modifying the parameters. |
| 500 | 1230003 | Internal Error | Please try again later. |
| 403 | 1230002 | No Permission | After becoming a member of the public mailbox, invoke this interface. |
