---
title: "Send Message"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/mail-v1/user_mailbox-message/send"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/mail/v1/user_mailboxes/:user_mailbox_id/messages/send"
service: "mail-v1"
resource: "user_mailbox-message"
section: "Email"
rate_limit: "100 per minute"
scopes:
  - "mail:user_mailbox.message:send"
updated: "1755077710000"
---

# Send Message

Send Message

> This interface is based on single-user locking and can only be called sequentially.

> Send emails using base64url encoding. The difference from regular base64 encoding is that the characters '+/' are replaced with '-_', specifically when using Golang, you can use base64.URLEncoding.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/mail/v1/user_mailboxes/:user_mailbox_id/messages/send |
| HTTP Method | POST |
| Rate Limit | 100 per minute |
| Supported app types | custom |
| Required scopes | `mail:user_mailbox.message:send` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `user_mailbox_id` | `string` | User email address, or enter me to represent the current calling interface user **Example value**: "user@xxx.xx or me" | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `subject` | `string` | No | Subject **Example value**: "Mail subject" |
| `to` | `mail_address[]` | No | Recipient |
|   `mail_address` | `string` | Yes | Mail address **Example value**: "user@xxx.xx" |
|   `name` | `string` | No | Name **Example value**: "Mike" |
| `raw` | `string` | No | eml data **Example value**: "U3ViamVjdDogSGVsbG8hCkZyb206ICJtaWtlIiA8bWlrZUBtaWtlLmNvbT4KTWltZS1WZXJzaW9uOiAxLjAKQ29udGVudC1UeXBlOiBtdWx0aXBhcnQvYWx0ZXJuYXRpdmU7CiBib3VuZGFyeT1iMjhmYTIyNGExZWU2ZDY3ZjE3OTViNGUxZDEwM2Q3MTBlNzM5ZWVmYjFmZjlmOWQ4NWI4M2NlOTRmMTEKRGF0ZTogV2VkLCAyMyBKdWwgMjAyNSAxNTo0NDoxOCArMDgwMApNZXNzYWdlLUlkOiA8bW9ja3V1aWRtZXNzYWdlX2lkQGxhcmsuY29tPgpUbzogImphY2siIDxqYWNrQGphY2suY29tPgoKLS1iMjhmYTIyNGExZWU2ZDY3ZjE3OTViNGUxZDEwM2Q3MTBlNzM5ZWVmYjFmZjlmOWQ4NWI4M2NlOTRmMTEKQ29udGVudC1UcmFuc2Zlci1FbmNvZGluZzogN2JpdApDb250ZW50LVR5cGU6IHRleHQvcGxhaW47IGNoYXJzZXQ9VVRGLTgKCldlbGNvbWUgdG8gTGFyayBtYWlsIQotLWIyOGZhMjI0YTFlZTZkNjdmMTc5NWI0ZTFkMTAzZDcxMGU3MzllZWZiMWZmOWY5ZDg1YjgzY2U5NGYxMQo=" |
| `cc` | `mail_address[]` | No | CC |
|   `mail_address` | `string` | Yes | Mail address **Example value**: "user@xxx.xx" |
|   `name` | `string` | No | Name **Example value**: "Mike" |
| `bcc` | `mail_address[]` | No | CC |
|   `mail_address` | `string` | Yes | Mail address **Example value**: "user@xxx.xx" |
|   `name` | `string` | No | Name **Example value**: "Mike" |
| `body_html` | `string` | No | Body **Example value**: "xxxx" |
| `body_plain_text` | `string` | No | Body plain text **Example value**: "xxxx" |
| `attachments` | `attachment[]` | No | Mail Attachment List |
|   `body` | `string` | Yes | The body of the attachment, encoded in base64url (maximum 37MB of supported files) **Example value**: "aGVsbG8gd29ybGQK" **Data validation rules**: - Maximum length: `51729750` characters |
|   `filename` | `string` | Yes | Attachment file name **Example value**: "helloworld.txt" |
|   `is_inline` | `boolean` | No | Is it an inline image? true means it is an inline image **Example value**: false **Default value**: `false` |
|   `cid` | `string` | No | Inline image content ID (required when is_inline=true). Referenced in HTML via cid:. **Example value**: "image1@example.com" |
| `dedupe_key` | `string` | No | deduplication key **Example value**: "abc-ddd-eee-fff-ggg" |
| `head_from` | `mail_address` | No | Sender |
|   `name` | `string` | No | name **Example value**: "Mike" | ### Request body example

{
    "subject": "Mail subject",
    "to": [
        {
            "mail_address": "user@xxx.xx",
            "name": "Mike"
        }
    ],
    "raw": "U3ViamVjdDogSGVsbG8hCkZyb206ICJtaWtlIiA8bWlrZUBtaWtlLmNvbT4KTWltZS1WZXJzaW9uOiAxLjAKQ29udGVudC1UeXBlOiBtdWx0aXBhcnQvYWx0ZXJuYXRpdmU7CiBib3VuZGFyeT1iMjhmYTIyNGExZWU2ZDY3ZjE3OTViNGUxZDEwM2Q3MTBlNzM5ZWVmYjFmZjlmOWQ4NWI4M2NlOTRmMTEKRGF0ZTogV2VkLCAyMyBKdWwgMjAyNSAxNTo0NDoxOCArMDgwMApNZXNzYWdlLUlkOiA8bW9ja3V1aWRtZXNzYWdlX2lkQGxhcmsuY29tPgpUbzogImphY2siIDxqYWNrQGphY2suY29tPgoKLS1iMjhmYTIyNGExZWU2ZDY3ZjE3OTViNGUxZDEwM2Q3MTBlNzM5ZWVmYjFmZjlmOWQ4NWI4M2NlOTRmMTEKQ29udGVudC1UcmFuc2Zlci1FbmNvZGluZzogN2JpdApDb250ZW50LVR5cGU6IHRleHQvcGxhaW47IGNoYXJzZXQ9VVRGLTgKCldlbGNvbWUgdG8gTGFyayBtYWlsIQotLWIyOGZhMjI0YTFlZTZkNjdmMTc5NWI0ZTFkMTAzZDcxMGU3MzllZWZiMWZmOWY5ZDg1YjgzY2U5NGYxMQo=",
    "cc": [
        {
            "mail_address": "user@xxx.xx",
            "name": "Mike"
        }
    ],
    "bcc": [
        {
            "mail_address": "user@xxx.xx",
            "name": "Mike"
        }
    ],
    "body_html": "xxxx",
    "body_plain_text": "xxxx",
    "attachments": [
        {
            "body": "aGVsbG8gd29ybGQK",
            "filename": "helloworld.txt",
            "is_inline": false,
            "cid": "image1@example.com"
        }
    ],
    "dedupe_key": "abc-ddd-eee-fff-ggg",
    "head_from": {
        "name": "Mike"
    }
}

## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `message_id` | `string` | message id |
|   `thread_id` | `string` | thread id | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
        "message_id": "48451e97-4743-4a55-a9a3-b5c656b69c05",
        "thread_id": "14151e97-4743-4a55-a9a3-b5c656b69c05"
    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 404 | 1234013 | user mailbox not found or user mailbox not active | The email address status is wrong, please change the email address. |
| 403 | 1234017 | permission deny | Apply for email permission |
| 400 | 1234008 | request parameter error | Please check if the request parameters are correct. |
| 400 | 1236001 | raw size over limit | The email size exceeds the limit. Please reduce the size of the email. |
| 429 | 1236006 | too many request | Please do not send concurrent requests from the same user. |
| 400 | 1236002 | invalid mime format | The MIME format of the email is invalid. Please check the email format and ensure that it is encoded using base64url. |
| 400 | 1236003 | the number of recipients exceeds the limit | The number of recipients in the email exceeds the limit of 500. Please reduce the number of recipients in the email. |
| 400 | 1236004 | the number of attachments exceeds the limit | The number of attachments in the email exceeds the limit of 500. Please reduce the number of attachments in the email. |
| 409 | 1236005 | send mail repeatedly | Please avoid sending duplicate emails using the same Message-ID. |
| 429 | 1236007 | the daily number of emails sent by the user exceeds the limit | The user has exceeded the daily limit for sending emails. |
| 429 | 1236008 | the number of external recipients the user sends messages to each day exceeds the limit | The user has exceeded the daily limit for sending emails to external recipients. |
| 429 | 1236009 | the number of external recipients the tenant sends messages to each day exceeds the limit | The organization has exceeded the daily limit for sending emails to external recipients. |
| 429 | 1236010 | mail quota limit | The user's outgoing request has limited viewership of the system, please try again |
| 429 | 1236012 | reach send mail restriction | The email sent by the user reaches the sender threshold limit |
| 429 | 1236013 | tenant storage limit | The tenant storage space is full and no more emails can be sent |
| 400 | 1236014 | content risk | The email failed to be sent because the content of the email was identified as risky |
| 400 | 1236017 | sender check fail | Mail sender detection failed, please check sender information and status |
| 400 | 1236018 | receiver check fail | Mail recipient detection failed, please check recipient information |
| 500 | 1236019 | internal server error | System internal error |
