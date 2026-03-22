---
title: "Update feed card button"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/group/im-v2/chat_button/update"
method: "PUT"
api_path: "https://open.larksuite.com/open-apis/im/v2/chat_button"
service: "im-v2"
resource: "chat_button"
section: "Feed"
rate_limit: "10 per second"
scopes:
  - "im:app_feed_card:write"
field_scopes:
  - "contact:user.employee_id:readonly"
updated: "1717121398000"
---

# Update feed card button

You can add shortcut operation buttons, update buttons, or delete buttons to Groups & Bots feed card.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/im/v2/chat_button |
| HTTP Method | PUT |
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
| `user_ids` | `string[]` | No | User ids **Example value**: ["ou_88553eda9014c201e6969b478895c223"] **Data validation rules**: - Length range: `1` ～ `20` |
| `chat_id` | `string` | Yes | chat id **Example value**: "oc_a0553eda9014c201e6969b478895c230" |
| `buttons` | `open_app_feed_card_buttons` | No | buttons |
|   `buttons` | `open_app_feed_card_button[]` | Yes | button **Data validation rules**: - Length range: `0` ～ `2` |
|     `multi_url` | `open_app_feed_card_url` | No | url (support https only) |
|       `url` | `string` | No | url **Example value**: "https://www.larksuite.com/" **Data validation rules**: - Maximum length: `700` characters |
|       `android_url` | `string` | No | Android url **Example value**: "https://www.larksuite.com/" **Data validation rules**: - Maximum length: `700` characters |
|       `ios_url` | `string` | No | iOS url **Example value**: "https://www.larksuite.com/" **Data validation rules**: - Maximum length: `700` characters |
|       `pc_url` | `string` | No | PC url **Example value**: "https://www.larksuite.com/" **Data validation rules**: - Maximum length: `700` characters |
|     `action_type` | `string` | Yes | action type **Example value**: "url_page" **Optional values are**:  - `url_page`: url page - `webhook`: webhook  |
|     `text` | `open_app_feed_card_text` | Yes | text |
|       `text` | `string` | Yes | text **Example value**: "text" **Data validation rules**: - Length range: `1` ～ `30` characters |
|     `button_type` | `string` | No | text **Example value**: "default" **Optional values are**:  - `default`: default - `primary`: primary - `success`: success  **Default value**: `default` |
|     `action_map` | `map<string, string>` | No | action map **Example value**: {"foo": "bar"} | ### Request body example

{
    "user_ids": [
        "ou_88553eda9014c201e6969b478895c223"
    ],
    "chat_id": "oc_a0553eda9014c201e6969b478895c230",
    "buttons": {
        "buttons": [
            {
                "multi_url": {
                    "url": "https://www.larksuite.com/",
                    "android_url": "https://www.larksuite.com/",
                    "ios_url": "https://www.larksuite.com/",
                    "pc_url": "https://www.larksuite.com/"
                },
                "action_type": "url_page",
                "text": {
                    "text": "text"
                },
                "button_type": "default",
                "action_map": {
                    "foo": "bar"
                }
            }
        ]
    }
}

## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `failed_user_reasons` | `failed_reason[]` | failed reason |
|     `error_code` | `int` | error code |
|     `error_message` | `string` | error message |
|     `user_id` | `string` | User ids | ### Response body example

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
| 400 | 230001 | param is invalid | Determine invalid fields and reasons according to the error message. |
| 400 | 230027 | no permission, bot not in the chat | Check if the bot is in the chat |
