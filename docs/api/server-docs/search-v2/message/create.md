---
title: "search message"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/search-v2/message/create"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/search/v2/message"
service: "search-v2"
resource: "message"
section: "Search"
rate_limit: "1000 per minute、50 per second"
scopes:
  - "search:message"
field_scopes:
  - "contact:user.employee_id:readonly"
updated: "1773323430000"
---

# Search Message

Users can search for visible messages by keyword, and the visibility is consistent with the search in the suite.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/search/v2/message |
| HTTP Method | POST |
| Rate Limit | 1000 per minute、50 per second |
| Supported app types | custom |
| Required scopes | `search:message` |
| Required field scopes | > The response body of the API contains the following sensitive fields, and they will be returned only after corresponding scopes are added. If you do not need the fields, it is not recommended that you request the scopes. `contact:user.employee_id:readonly` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `user_id_type` | `string` | No | User ID categories **Example value**: open_id **Optional values are**:  - `open_id`: Identifies a user to an app. The same user has different Open IDs in different apps. How to get Open ID - `union_id`: Identifies a user to a tenant that acts as a developer. A user has the same Union ID in apps developed by the same developer, and has different Union IDs in apps developed by different developers. A developer can use Union ID to link the same user's identities in multiple apps.How to get Union ID - `user_id`: Identifies a user to a tenant. The same user has different User IDs in different tenants. In one single tenant, a user has the same User ID in all apps （including store apps）. User ID is usually used to communicate user data between different apps. How to get User ID  **Default value**: `open_id` **When the value is `user_id`, the following field scopes are required**: `contact:user.employee_id:readonly` |
| `page_size` | `int` | No | Page size **Example value**: 20 **Default value**: `20` **Data validation rules**: - Value range: `1` ～ `100` |
| `page_token` | `string` | No | Page identifier. It is not filled in the first request, indicating traversal from the beginning; when there will be more groups, the new page_token will be returned at the same time, and the next traversal can use the page_token to get more groups **Example value**: 9e91187f9107ef4d43cd71c3722cd97665e6cec51bf30a06328839bc9867 | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `query` | `string` | Yes | Search Keyword **Example value**: "测试消息" |
| `from_ids` | `string[]` | No | Messages From User IDs **Example value**: ["ou_1970b39a6730a4a8e97b530d8cb14ccb"] |
| `chat_ids` | `string[]` | No | Messages In Chat IDs **Example value**: ["oc_c063434856a818a615fd36697a9ffe09"] |
| `message_type` | `string` | No | Message type(file/image/media) **Example value**: "image" **Optional values are**:  - `file`: file - `image`: image - `media`: media  |
| `at_chatter_ids` | `string[]` | No | At User IDs **Example value**: ["ou_1970b39a6730a4a8e97b530d8cb14ccb"] |
| `from_type` | `string` | No | Messages From Type(bot/user) **Example value**: "user" **Optional values are**:  - `bot`: bot - `user`: user  |
| `chat_type` | `string` | No | Message In Chat type(group_chat/p2p_chat) **Example value**: "group_chat" **Optional values are**:  - `group_chat`: group_chat - `p2p_chat`: p2p_chat  |
| `start_time` | `string` | No | Message Start Time **Example value**: "1609296809" |
| `end_time` | `string` | No | Message End Time **Example value**: "1609296809" | ### Request body example

{
    "query": "测试消息",
    "from_ids": [
        "ou_1970b39a6730a4a8e97b530d8cb14ccb"
    ],
    "chat_ids": [
        "oc_c063434856a818a615fd36697a9ffe09"
    ],
    "message_type": "image",
    "at_chatter_ids": [
        "ou_1970b39a6730a4a8e97b530d8cb14ccb"
    ],
    "from_type": "user",
    "chat_type": "group_chat",
    "start_time": "1609296809",
    "end_time": "1609296809"
}

## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `items` | `string[]` | Message IDs |
|   `page_token` | `string` | Page identifier, when has_more is true, a new page_token will also be returned. Otherwise, page_token will not be returned |
|   `has_more` | `boolean` | Whether the response body has more parameters | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
        "items": [
            "om_a7ccbed2290bde912dd0418d19731865"
        ],
        "page_token": "GxmvlNRvP0NdQZpa7yIqf_Lv_QuBwTQ8tXkX7w-irAghVD_TvuYd1aoJ1LQph86O-XImC4X9j9FhUPhXQDvtrQ==",
        "has_more": true
    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 1278001 | invalid param | Parameter error. Check the parameter. |
| 500 | 1278002 | Internal Server Error | System error |
