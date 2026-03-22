---
title: "Get Attachment Download Links"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/mail-v1/user_mailbox-message-attachment/download_url"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/mail/v1/user_mailboxes/:user_mailbox_id/messages/:message_id/attachments/download_url"
service: "mail-v1"
resource: "user_mailbox-message-attachment"
section: "Email"
rate_limit: "1 per second"
scopes:
  - "mail:user_mailbox.message.body:read"
updated: "1745840308000"
---

# Get Attachment Download Links

Get Attachment Download Links

> The download links are limited to two uses and expire after two hours

> When using the tenant_access_token, it is necessary to apply for data permissions for accessing email message resources.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/mail/v1/user_mailboxes/:user_mailbox_id/messages/:message_id/attachments/download_url |
| HTTP Method | GET |
| Rate Limit | 1 per second |
| Supported app types | custom,isv |
| Required scopes | `mail:user_mailbox.message.body:read` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `user_mailbox_id` | `string` | User email address, or enter me to represent the current calling interface user **Example value**: "user@xxx.xx 或 me" |
| `message_id` | `string` | Message ID, the method for obtaining ID is shown in List Emails **Example value**: "TUlHc1NoWFhJMXgyUi9VZTNVL3h6UnlkRUdzPQ==" | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `attachment_ids` | `string[]` | Yes | Attachment ID list **Example value**: YQqYbQHoQoDqXjxWKhJbo8Gicjf **Data validation rules**: - Length range: `1` ～ `20` | ## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `download_urls` | `attachment_download_url_item[]` | Download link list |
|     `attachment_id` | `string` | Attachment ID |
|     `download_url` | `string` | Download link |
|   `failed_ids` | `string[]` | List of failed attachment IDs retrieval | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
        "download_urls": [
            {
                "attachment_id": "YQqYbQHoQoDqXjxWKhJbo8Gicjf",
                "download_url": "https://api-drive-stream.blmpb.com/space/api/box/stream/download/authcode/?code=YTZiZGViMDg3NzRjMzEwOWRkMGI1MTJlYmQxYTFmYTBfZTA5ZjZiOWU4NDYzMzkxMDUyOTIxMzBmNTVjMjAyZTFfSUQ6NzI4MTE4Nzg1OTE5NTc3Mjk0N18xNjk1ODg4NjQyOjE2OTU4ODg3MDJfVjM"
            }
        ],
        "failed_ids": [
            "YQqYbQHoQoDqXjxWKhJbo8Gicjf"
        ]
    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 1230001 | Parameter Error | Please retry after modifying the parameters |
| 500 | 1230003 | No Permission | Call this API after becoming a member of the public email or applying for relevant data permissions |
| 403 | 1230002 | Internal Error | Please try again later |
