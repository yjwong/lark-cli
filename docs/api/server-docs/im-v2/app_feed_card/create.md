---
title: "Create app feed card"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/group/im-v2/app_feed_card/create"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/im/v2/app_feed_card"
service: "im-v2"
resource: "app_feed_card"
section: "Feed"
rate_limit: "10 per second"
scopes:
  - "im:app_feed_card:write"
field_scopes:
  - "contact:user.employee_id:readonly"
updated: "1717121390000"
---

# Create app feed card

The app feed card is a tool that allows applications to directly send messages within the feed list, ensuring that important messages reach users more quickly.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/im/v2/app_feed_card |
| HTTP Method | POST |
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
| `app_feed_card` | `open_app_feed_card` | No | app feed card |
|   `biz_id` | `string` | No | business id **Example value**: "096e2927-40a6-41a3-9562-314d641d09ae" **Data validation rules**: - Length range: `0` ～ `200` characters |
|   `title` | `string` | No | title **Example value**: "main title" **Data validation rules**: - Length range: `0` ～ `60` characters |
|   `avatar_key` | `string` | No | avatar key. Call the upload images API to get this value **Example value**: "v3_0041_007bca9f-67ba-4199-bf00-4031b12cf226" **Data validation rules**: - Length range: `0` ～ `100` characters |
|   `preview` | `string` | No | preview text **Example value**: "this is preview text" **Data validation rules**: - Length range: `0` ～ `120` characters |
|   `status_label` | `open_feed_status_label` | No | status label |
|     `text` | `string` | Yes | text **Example value**: "label text" **Data validation rules**: - Length range: `1` ～ `20` characters |
|     `type` | `string` | Yes | label type **Example value**: "primary" **Optional values are**:  - `primary`: primary - `secondary`: secondary - `success`: success - `danger`: Danger  **Default value**: `primary` |
|   `buttons` | `open_app_feed_card_buttons` | No | button |
|     `buttons` | `open_app_feed_card_button[]` | Yes | button **Data validation rules**: - Length range: `0` ～ `2` |
|       `multi_url` | `open_app_feed_card_url` | No | url (support https only) |
|         `url` | `string` | No | url **Example value**: "https://www.larksuite.com/" **Data validation rules**: - Maximum length: `700` characters |
|         `android_url` | `string` | No | Android url **Example value**: "https://www.larksuite.com/" **Data validation rules**: - Maximum length: `700` characters |
|         `ios_url` | `string` | No | iOS url **Example value**: "https://www.larksuite.com/" **Data validation rules**: - Maximum length: `700` characters |
|         `pc_url` | `string` | No | PC url **Example value**: "https://www.larksuite.com/" **Data validation rules**: - Maximum length: `700` characters |
|       `action_type` | `string` | Yes | action type **Example value**: "url_page" **Optional values are**:  - `url_page`: url page - `webhook`: webhook  |
|       `text` | `open_app_feed_card_text` | Yes | text |
|         `text` | `string` | Yes | text **Example value**: "text demo" **Data validation rules**: - Length range: `1` ～ `30` characters |
|       `button_type` | `string` | No | button type **Example value**: "default" **Optional values are**:  - `default`: default - `primary`: primary - `success`: success  **Default value**: `default` |
|       `action_map` | `map<string, string>` | No | action map **Example value**: {"foo": "bar"} |
|   `link` | `open_app_feed_link` | No | link |
|     `link` | `string` | No | link (support https only) **Example value**: "https://www.larksuite.com/" **Data validation rules**: - Maximum length: `700` characters |
|   `time_sensitive` | `boolean` | No | Instant reminder state, true-open, false-close **Example value**: false **Default value**: `false` |
|   `notify` | `app_feed_notify` | No | notify |
|     `close_notify` | `boolean` | No | without notify **Example value**: true |
| `user_ids` | `string[]` | No | user ids **Example value**: ["ou_88553eda9014c201e6969b478895c223"] **Data validation rules**: - Length range: `1` ～ `20` | ### Request body example

{
    "app_feed_card": {
        "biz_id": "096e2927-40a6-41a3-9562-314d641d09ae",
        "title": "main title",
        "avatar_key": "v3_0041_007bca9f-67ba-4199-bf00-4031b12cf226",
        "preview": "this is preview text",
        "status_label": {
            "text": "label text",
            "type": "primary"
        },
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
                        "text": "text demo"
                    },
                    "button_type": "default",
                    "action_map": {
                        "foo": "bar"
                    }
                }
            ]
        },
        "link": {
            "link": "https://www.larksuite.com/"
        },
        "time_sensitive": false,
        "notify": {
            "close_notify": true
        }
    },
    "user_ids": [
        "ou_88553eda9014c201e6969b478895c223"
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
|     `reason` | `string` | reason **Optional values are**:  - `0`: unknown - `1`: no permission - `2`: not created - `3`: rate limited - `4`: duplicated  |
|   `biz_id` | `string` | business id | ### Response body example

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
        ],
        "biz_id": "b90ce43a-fca8-4f42-92f4-794bff206ee5"
    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 230001 | param is invalid | Determine invalid fields and reasons according to the error message |
