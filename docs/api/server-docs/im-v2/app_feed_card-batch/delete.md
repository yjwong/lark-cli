---
title: "Delete app feed card"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/group/im-v2/app_feed_card-batch/delete"
method: "DELETE"
api_path: "https://open.larksuite.com/open-apis/im/v2/app_feed_card/batch"
service: "im-v2"
resource: "app_feed_card-batch"
section: "Feed"
rate_limit: "10 per second"
scopes:
  - "im:app_feed_card:write"
field_scopes:
  - "contact:user.employee_id:readonly"
updated: "1711958935000"
---

# Delete app feed card

This interface is used to delete application information flow cards.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/im/v2/app_feed_card/batch |
| HTTP Method | DELETE |
| Rate Limit | 10 per second |
| Supported app types | custom,isv |
| Required scopes | `im:app_feed_card:write` |
| Required field scopes | > The response body of the API contains the following sensitive fields, and they will be returned only after corresponding scopes are added. If you do not need the fields, it is not recommended that you request the scopes. `contact:user.employee_id:readonly` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `user_id_type` | `string` | No | User ID categories **Example value**: open_id **Optional values are**:  - `open_id`: Identifies a user to an app. The same user has different Open IDs in different apps. How to get Open ID - `union_id`: Identifies a user to a tenant that acts as a developer. A user has the same Union ID in apps developed by the same developer, and has different Union IDs in apps developed by different developers. A developer can use Union ID to link the same user's identities in multiple apps.How to get Union ID - `user_id`: Identifies a user to a tenant. The same user has different User IDs in different tenants. In one single tenant, a user has the same User ID in all apps （including store apps）. User ID is usually used to communicate user data between different apps. How to get User ID  **Default value**: `open_id` **When the value is `user_id`, the following field scopes are required**: `contact:user.employee_id:readonly` | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `feed_cards` | `user_open_app_feed_card_deleter[]` | No | feed cards **Data validation rules**: - Length range: `1` ～ `20` |
|   `biz_id` | `string` | Yes | business id **Example value**: "ed381d34-49ac-4876-8d9e-23447acb587e" |
|   `user_id` | `string` | Yes | user id **Example value**: "ou_88553eda9014c201e6969b478895c223" | ### Request body example

{
    "feed_cards": [
        {
            "biz_id": "ed381d34-49ac-4876-8d9e-23447acb587e",
            "user_id": "ou_88553eda9014c201e6969b478895c223"
        }
    ]
}

## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `failed_cards` | `open_failed_user_app_feed_card_item[]` | failed cards |
|     `biz_id` | `string` | business id |
|     `user_id` | `string` | user id |
|     `reason` | `string` | failed reason **Optional values are**:  - `0`: unknown - `1`: no permission - `2`: not created - `3`: rate limited - `4`: duplicated  | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
        "failed_cards": [
            {
                "biz_id": "bdf22389-87ec-4890-9eb6-78a7efaeecbb",
                "user_id": "ou_88553eda9014c201e6969b478895c223",
                "reason": "NO_PERMISSION"
            }
        ]
    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 230001 | param is invalid | Determine invalid fields and reasons according to the error message |
