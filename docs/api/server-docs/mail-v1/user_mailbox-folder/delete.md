---
title: "Delete Email Folder"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/mail-v1/user_mailbox-folder/delete"
method: "DELETE"
api_path: "https://open.larksuite.com/open-apis/mail/v1/user_mailboxes/:user_mailbox_id/folders/:folder_id"
service: "mail-v1"
resource: "user_mailbox-folder"
section: "Email"
rate_limit: "1 per second"
scopes:
  - "mail:user_mailbox.folder:write"
updated: "1745840283000"
---

# Delete Email Folder

Delete Email Folder

> When using the tenant_access_token, it is necessary to apply for data permissions for accessing mail folder resources.

> Deleting a folder will move the emails within that folder to the deleted items folder.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/mail/v1/user_mailboxes/:user_mailbox_id/folders/:folder_id |
| HTTP Method | DELETE |
| Rate Limit | 1 per second |
| Supported app types | custom |
| Required scopes | `mail:user_mailbox.folder:write` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `user_mailbox_id` | `string` | User email address, or enter me to represent the current calling interface user **Example value**: "user@xxx.xx or me" |
| `folder_id` | `string` | Folder ID. The method for obtaining ID is shown in List Email Folders **Example value**: "111111" | ## Response

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
| 403 | 1230002 | No Permission | Call this API after becoming a member of the public email or applying for relevant data permissions. |
| 500 | 1230003 | Internal Error | Please try again later. |
