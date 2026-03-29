---
title: "List Emails"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/mail-v1/user_mailbox-message/list"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/mail/v1/user_mailboxes/:user_mailbox_id/messages"
service: "mail-v1"
resource: "user_mailbox-message"
section: "Email"
rate_limit: "10 per second"
scopes:
  - "mail:user_mailbox.message:readonly"
updated: "1745840299000"
---

# List Emails

List Emails

> When using the tenant_access_token, it is necessary to apply for data permissions for accessing email message resources.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/mail/v1/user_mailboxes/:user_mailbox_id/messages |
| HTTP Method | GET |
| Rate Limit | 10 per second |
| Supported app types | custom |
| Required scopes | `mail:user_mailbox.message:readonly` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `user_mailbox_id` | `string` | User email address, or enter me to represent the current calling interface user **Example value**: "user@xxx.xx or me" | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `page_size` | `int` | Yes | page size **Example value**: 1 **Data validation rules**: - Value range: `1` ～ `20` |
| `page_token` | `string` | No | Page identifier. It is not filled in the first request, indicating traversal from the beginning; when there will be more groups, the new page_token will be returned at the same time, and the next traversal can use the page_token to get more groups **Example value**: xxx |
| `folder_id` | `string` | Yes | Folder ID. The method for obtaining ID is shown in List Email Folders **Example value**: INBOX or user folder id |
| `only_unread` | `boolean` | No | Only query unread emails **Example value**: true | ## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `items` | `string[]` | message id list |
|   `page_token` | `string` | Page identifier, when has_more is true, a new page_token will also be returned. Otherwise, page_token will not be returned |
|   `has_more` | `boolean` | Whether the response body has more parameters | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
        "items": [
            "ZWEyNGRmY2QtOTVlNy00NzlhLTg5MjItMjFjOTQ5ZjIzZjJl"
        ],
        "page_token": "xxx",
        "has_more": true
    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 1230001 | Parameter Error | Please retry after modifying the parameters. |
| 403 | 1230002 | No Permission | Call this API after becoming a member of the public email or applying for relevant data permissions. |
| 500 | 1230003 | Internal Error | Please try again later. |
