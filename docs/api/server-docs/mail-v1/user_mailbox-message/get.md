---
title: "Get Email Details"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/mail-v1/user_mailbox-message/get"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/mail/v1/user_mailboxes/:user_mailbox_id/messages/:message_id"
service: "mail-v1"
resource: "user_mailbox-message"
section: "Email"
rate_limit: "100 per minute"
scopes:
  - "mail:user_mailbox.message:readonly"
field_scopes:
  - "mail:user_mailbox.message.address:read"
  - "mail:user_mailbox.message.body:read"
  - "mail:user_mailbox.message.subject:read"
updated: "1754648844000"
---

# Get Email Details

Get Email Details

> When using the tenant_access_token, it is necessary to apply for data permissions for accessing email message resources.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/mail/v1/user_mailboxes/:user_mailbox_id/messages/:message_id |
| HTTP Method | GET |
| Rate Limit | 100 per minute |
| Supported app types | custom |
| Required scopes | `mail:user_mailbox.message:readonly` |
| Required field scopes | > The response body of the API contains the following sensitive fields, and they will be returned only after corresponding scopes are added. If you do not need the fields, it is not recommended that you request the scopes. `mail:user_mailbox.message.address:read` `mail:user_mailbox.message.body:read` `mail:user_mailbox.message.subject:read` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `user_mailbox_id` | `string` | User email address, or enter me to represent the current calling interface user **Example value**: "user@xxx.xx or me" |
| `message_id` | `string` | Message ID, the method for obtaining ID is shown in List Emails **Example value**: "TUlHc1NoWFhJMXgyUi9VZTNVL3h6UnlkRUdzPQ==" | ## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `message` | `message` | Message body |
|     `subject` | `string` | Subject **Required field scopes**: `mail:user_mailbox.message.subject:read` |
|     `to` | `mail_address[]` | Recipient **Required field scopes**: `mail:user_mailbox.message.address:read` |
|       `mail_address` | `string` | Mail address |
|       `name` | `string` | Name |
|     `cc` | `mail_address[]` | CC **Required field scopes**: `mail:user_mailbox.message.address:read` |
|       `mail_address` | `string` | Mail address |
|       `name` | `string` | Name |
|     `bcc` | `mail_address[]` | Cc **Required field scopes**: `mail:user_mailbox.message.address:read` |
|       `mail_address` | `string` | Mail address |
|       `name` | `string` | Name |
|     `head_from` | `mail_address` | Sender **Required field scopes**: `mail:user_mailbox.message.address:read` |
|       `mail_address` | `string` | Mail address |
|       `name` | `string` | Name |
|     `body_html` | `string` | Body (base64url) **Required field scopes**: `mail:user_mailbox.message.body:read` |
|     `internal_date` | `string` | Create/receive/send time (milliseconds) |
|     `message_state` | `int` | Mail status |
|     `smtp_message_id` | `string` | RFC protocol id |
|     `message_id` | `string` | ID |
|     `body_plain_text` | `string` | Body plain text (base64url) **Required field scopes**: `mail:user_mailbox.message.body:read` |
|     `attachments` | `attachment[]` | attachment list **Required field scopes**: `mail:user_mailbox.message.body:read` |
|       `filename` | `string` | attachment filename |
|       `id` | `string` | attachment id |
|       `attachment_type` | `int` | attachment type **Optional values are**:  - `1`: standard attachment - `2`: oversized attachment  |
|       `is_inline` | `boolean` | `true` indicates the image is inline (embedded in email body), `false` treats it as a regular attachment. |
|       `cid` | `string` | Content ID, HTML refers to the image via cid: protocol |
|     `thread_id` | `string` | thread id | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
        "message": {
            "subject": "Mail Subject",
            "to": [
                {
                    "mail_address": "mike@outlook.com",
                    "name": "Mike"
                }
            ],
            "cc": [
                {
                    "mail_address": "mike@outlook.com",
                    "name": "Mike"
                }
            ],
            "bcc": [
                {
                    "mail_address": "mike@outlook.com",
                    "name": "Mike"
                }
            ],
            "head_from": {
                "mail_address": "mike@outlook.com",
                "name": "Mike"
            },
            "body_html": "xxxx",
            "internal_date": "1682377086000",
            "message_state": 1,
            "smtp_message_id": "ay0azrJDvbs3FJAg@outlook.com",
            "message_id": "tfuh9N4WnzU6jdDw=",
            "body_plain_text": "xxxxx",
            "attachments": [
                {
                    "filename": "helloworld.txt",
                    "id": "YQqYbQHoQoDqXjxWKhJbo8Gicjf",
                    "attachment_type": 1,
                    "is_inline": false,
                    "cid": "image1@example.com"
                }
            ],
            "thread_id": "tfuh9N4WnzU6jdDw="
        }
    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 404 | 1235013 | user mailbox not found | Check mailbox and try again |
| 400 | 1234008 | Parameter Error | Retry After Modification |
| 404 | 1234034 | user message not found | Check The Mail ID And try Again |
