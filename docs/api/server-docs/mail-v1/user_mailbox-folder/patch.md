---
title: "Update Email Folder"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/mail-v1/user_mailbox-folder/patch"
method: "PATCH"
api_path: "https://open.larksuite.com/open-apis/mail/v1/user_mailboxes/:user_mailbox_id/folders/:folder_id"
service: "mail-v1"
resource: "user_mailbox-folder"
section: "Email"
rate_limit: "1 per second"
scopes:
  - "mail:user_mailbox.folder:write"
updated: "1745840287000"
---

# Update Email Folder

Update Email Folder

> When using the tenant_access_token, it is necessary to apply for data permissions for accessing mail folder resources.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/mail/v1/user_mailboxes/:user_mailbox_id/folders/:folder_id |
| HTTP Method | PATCH |
| Rate Limit | 1 per second |
| Supported app types | custom |
| Required scopes | `mail:user_mailbox.folder:write` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `user_mailbox_id` | `string` | User email address, or enter me to represent the current calling interface user **Example value**: "user@xxx.xx or me" |
| `folder_id` | `string` | Folder ID. The method for obtaining ID is shown in List Email Folders **Example value**: "111111" | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `name` | `string` | No | Name **Example value**: "newsletter" **Data validation rules**: - Length range: `1` ～ `250` characters |
| `parent_folder_id` | `string` | No | Parent folder ID, where a value of 0 indicates the root folder. The method for obtaining ID is shown in List Email Folders **Example value**: "725627422334644" | ### Request body example

{
    "name": "newsletter",
    "parent_folder_id": "725627422334644"
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
| 403 | 1230002 | No Permission | Call this API after becoming a member of the public email or applying for relevant data permissions. |
