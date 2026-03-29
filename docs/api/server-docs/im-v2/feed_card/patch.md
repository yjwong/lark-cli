---
title: "Instant reminder"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/group/im-v2/feed_card/patch"
method: "PATCH"
api_path: "https://open.larksuite.com/open-apis/im/v2/feed_cards/:feed_card_id"
service: "im-v2"
resource: "feed_card"
section: "Feed"
rate_limit: "1000 per minute、50 per second"
scopes:
  - "im:datasync.feed_card.time_sensitive:write"
field_scopes:
  - "contact:user.employee_id:readonly"
updated: "1711958944000"
---

# Instant reminder

The instant reminder capability is a strong reminder capability provided by Lark in the message list. When there are important notifications or tasks that need to be reached in time, users can use the instant reminder capability to display group or robot conversations at the top of the message list. Open Lark. Tackle important tasks right from the first page of the book.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/im/v2/feed_cards/:feed_card_id |
| HTTP Method | PATCH |
| Rate Limit | 1000 per minute、50 per second |
| Supported app types | custom,isv |
| Required scopes | `im:datasync.feed_card.time_sensitive:write` |
| Required field scopes | > The response body of the API contains the following sensitive fields, and they will be returned only after corresponding scopes are added. If you do not need the fields, it is not recommended that you request the scopes. `contact:user.employee_id:readonly` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `feed_card_id` | `string` | Group id, currently only group chat is supported. **Example value**: "oc_679eaeb583654bff73fefcc6e6371370" | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `user_id_type` | `string` | Yes | User ID categories **Example value**: open_id **Optional values are**:  - `open_id`: Identifies a user to an app. The same user has different Open IDs in different apps. How to get Open ID - `union_id`: Identifies a user to a tenant that acts as a developer. A user has the same Union ID in apps developed by the same developer, and has different Union IDs in apps developed by different developers. A developer can use Union ID to link the same user's identities in multiple apps.How to get Union ID - `user_id`: Identifies a user to a tenant. The same user has different User IDs in different tenants. In one single tenant, a user has the same User ID in all apps （including store apps）. User ID is usually used to communicate user data between different apps. How to get User ID  **Default value**: `open_id` **When the value is `user_id`, the following field scopes are required**: `contact:user.employee_id:readonly` | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `time_sensitive` | `boolean` | Yes | Instant reminder state, true-open, false-close **Example value**: true |
| `user_ids` | `string[]` | Yes | List of user IDs, support OpenID, UnionID, UserID. **Example value**: ["ou_9d2beb613c85a2412862a49a924558c5"] | ### Request body example

{
    "time_sensitive": true,
    "user_ids": [
        "ou_9d2beb613c85a2412862a49a924558c5"
    ]
}

## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `failed_user_reasons` | `failed_reason[]` | Reason for failure |
|     `error_code` | `int` | error code |
|     `error_message` | `string` | error message |
|     `user_id` | `string` | User ID | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
        "failed_user_reasons": [
            {
                "error_code": 0,
                "error_message": "The user is not in the chat",
                "user_id": "ou_679eaeb583654bff73fefcc6e6371301"
            }
        ]
    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 230001 | param is invalid | Check request parameters |
| 400 | 230027 | no permission | Bot has no permission |
