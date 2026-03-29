---
title: "List Email Folders"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/mail-v1/user_mailbox-folder/list"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/mail/v1/user_mailboxes/:user_mailbox_id/folders"
service: "mail-v1"
resource: "user_mailbox-folder"
section: "Email"
rate_limit: "5 per second"
scopes:
  - "mail:user_mailbox.folder:read"
  - "mail:user_mailbox.folder:write"
updated: "1745840291000"
---

# List Email Folders

List Email Folders

> When using the tenant_access_token, it is necessary to apply for data permissions for accessing mail folder resources.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/mail/v1/user_mailboxes/:user_mailbox_id/folders |
| HTTP Method | GET |
| Rate Limit | 5 per second |
| Supported app types | custom |
| Required scopes Enable any scope from the list | `mail:user_mailbox.folder:read` `mail:user_mailbox.folder:write` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `user_mailbox_id` | `string` | User email address, or enter me to represent the current calling interface user **Example value**: "user@xxx.xx or me" | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `folder_type` | `int` | No | Folder type **Example value**: 1 **Optional values are**:  - `1`: System Folder - `2`: User folder  **Data validation rules**: - Value range: `1` ～ `2` | ## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `items` | `folder[]` | Folder list |
|     `id` | `string` | ID |
|     `name` | `string` | Name |
|     `parent_folder_id` | `string` | Parent folder ID, where a value of 0 indicates the root folder. |
|     `folder_type` | `int` | Folder type **Optional values are**:  - `1`: System Folder - `2`: User folder  |
|     `unread_message_count` | `int` | Unread message count |
|     `unread_thread_count` | `int` | Unread thread count | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
        "items": [
            {
                "id": "12314123123123123",
                "name": "newsletter",
                "parent_folder_id": "725627422334644",
                "folder_type": 1,
                "unread_message_count": 3,
                "unread_thread_count": 4
            }
        ]
    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 1230001 | Parameter Error | Please retry after modifying the parameters. |
| 500 | 1230003 | Internal Error | Please try again later. |
| 403 | 1230002 | No Permission | Call this API after becoming a member of the public email or applying for relevant data permissions. |
