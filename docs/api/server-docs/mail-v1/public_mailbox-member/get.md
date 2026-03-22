---
title: "Get A Public Mailbox Member"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/mail-v1/public_mailbox-member/get"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/mail/v1/public_mailboxes/:public_mailbox_id/members/:member_id"
service: "mail-v1"
resource: "public_mailbox-member"
section: "Email"
rate_limit: "Special Rate Limit"
scopes:
  - "mail:public_mailbox"
  - "mail:public_mailbox:readonly"
field_scopes:
  - "contact:user.employee_id:readonly"
updated: "1745840347000"
---

# Get a public mailbox member

Obtains the information of a public mailbox member.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/mail/v1/public_mailboxes/:public_mailbox_id/members/:member_id |
| HTTP Method | GET |
| Rate Limit | Special Rate Limit |
| Supported app types | custom |
| Required scopes Enable any scope from the list | `mail:public_mailbox` `mail:public_mailbox:readonly` |
| Required field scopes | > The response body of the API contains the following sensitive fields, and they will be returned only after corresponding scopes are added. If you do not need the fields, it is not recommended that you request the scopes. `contact:user.employee_id:readonly` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `public_mailbox_id` | `string` | Unique identifier of a public mailbox or the public mailbox address **Example value**: "xxxxxxxxxxxxxxx or test_public_mailbox@xxx.xx" |
| `member_id` | `string` | Unique identifier of the member in the public mailbox **Example value**: "xxxxxxxxxxxxxxx" | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `user_id_type` | `string` | No | User ID categories **Example value**: open_id **Optional values are**:  - `open_id`: Identifies a user to an app. The same user has different Open IDs in different apps. How to get Open ID - `union_id`: Identifies a user to a tenant that acts as a developer. A user has the same Union ID in apps developed by the same developer, and has different Union IDs in apps developed by different developers. A developer can use Union ID to link the same user's identities in multiple apps.How to get Union ID - `user_id`: Identifies a user to a tenant. The same user has different User IDs in different tenants. In one single tenant, a user has the same User ID in all apps （including store apps）. User ID is usually used to communicate user data between different apps. How to get User ID  **Default value**: `open_id` **When the value is `user_id`, the following field scopes are required**: `contact:user.employee_id:readonly` | ## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `public_mailbox.member` | \- |
|   `member_id` | `string` | The unique identity of the member in the public mailbox (no need to fill in the request body) |
|   `user_id` | `string` | Unique identifier of the user in the tenant (this field has a value when member type is USER) |
|   `type` | `string` | Member type **Optional values are**:  - `USER`: Internal user  | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
        "member_id": "xxxxxxxxxxxxxxx",
        "user_id": "xxxxxxxxxx",
        "type": "USER"
    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 404 | 1234016 | public mailbox not found | Check whether the public mailbox exists. |
| 404 | 1234040 | public mailbox member not found | Check whether the public mailbox member exists. |
