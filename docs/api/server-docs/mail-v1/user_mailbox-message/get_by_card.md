---
title: "list mail of card"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/mail-v1/user_mailbox-message/get_by_card"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/mail/v1/user_mailboxes/:user_mailbox_id/messages/get_by_card"
service: "mail-v1"
resource: "user_mailbox-message"
section: "Email"
rate_limit: "10 per second"
scopes:
  - "mail:user_mailbox.message:readonly"
field_scopes:
  - "contact:user.employee_id:readonly"
updated: "1745840295000"
---

# list mail of mail card

list mail of mail card

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/mail/v1/user_mailboxes/:user_mailbox_id/messages/get_by_card |
| HTTP Method | GET |
| Rate Limit | 10 per second |
| Supported app types | custom |
| Required scopes | `mail:user_mailbox.message:readonly` |
| Required field scopes | > The response body of the API contains the following sensitive fields, and they will be returned only after corresponding scopes are added. If you do not need the fields, it is not recommended that you request the scopes. `contact:user.employee_id:readonly` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `user_mailbox_id` | `string` | User email address, or enter me to represent the current calling interface user **Example value**: "user@xxx.xx or me" | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `card_id` | `string` | Yes | Mail card ID, available via [Receive message]( /ssl:ttdoc/server-docs/im-v1/message/events/receive) events **Example value**: 512ca581-6059-4449-8150-5522e6641d32 |
| `owner_id` | `string` | Yes | Mail card OwnerID, available via [Receive message]( /ssl:ttdoc/server-docs/im-v1/message/events/receive) events (independent of `user_id_type`) **Example value**: 1234567890 |
| `user_id_type` | `string` | No | User ID categories **Example value**: open_id **Optional values are**:  - `open_id`: Identifies a user to an app. The same user has different Open IDs in different apps. How to get Open ID - `union_id`: Identifies a user to a tenant that acts as a developer. A user has the same Union ID in apps developed by the same developer, and has different Union IDs in apps developed by different developers. A developer can use Union ID to link the same user's identities in multiple apps.How to get Union ID - `user_id`: Identifies a user to a tenant. The same user has different User IDs in different tenants. In one single tenant, a user has the same User ID in all apps （including store apps）. User ID is usually used to communicate user data between different apps. How to get User ID  **Default value**: `open_id` **When the value is `user_id`, the following field scopes are required**: `contact:user.employee_id:readonly` | ## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `owner_info` | `user_info` | Mail Owner Information |
|     `type` | `string` | owner type **Optional values are**: - `user`: personal mailbox - `public_mailbox `: public mailbox |
|     `owner_user_id` | `string` | mail card owner id, non-null if owner type is `user` |
|     `public_mailbox_id` | `string` | public mailbox id, not-null if owner type is `public_mailbox` |
|   `message_ids` | `string[]` | list of message id |
|   `card_id` | `string` | mail card id | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
        "owner_info": {
            "type": "user",
            "owner_user_id": "ou_7dab8a3d3cdcc9da365777c7ad115d62",
            "public_mailbox_id": "xxxxxxxxxx"
        },
        "message_ids": [
            "TnAwLzVZbjB3TTcyeU5Tek1UMHRFa0l3aStNPQ=="
        ],
        "card_id": "512ca581-6059-4449-8150-5522e6646d32"
    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 1230001 | Parameter Error | Please retry after modifying the parameters. |
| 500 | 1230003 | Internal Error | Please try again later. |
| 403 | 1230002 | No Permission | Call this API after becoming a member of the public email or applying for relevant data permissions. |
