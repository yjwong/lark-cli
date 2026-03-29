---
title: "Sort chat tab"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat-tab/sort_tabs"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/im/v1/chats/:chat_id/chat_tabs/sort_tabs"
service: "im-v1"
resource: "chat-tab"
section: "Group Chat"
rate_limit: "1000 per minute、50 per second"
scopes:
  - "im:chat"
  - "im:chat.tabs:write_only"
updated: "1717574958000"
---

# Sort chat tab

Chat tab sorting.

> Notice:
> - The bot ability needs to be enabled.
> - The bot or authorized user must be in the group.
> - The current message tab is fixed to the first rank and does not participate in sorting, but must be included in the request body.
> - When operating internal group, the operator must be under the same tenant as the group.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/im/v1/chats/:chat_id/chat_tabs/sort_tabs |
| HTTP Method | POST |
| Rate Limit | 1000 per minute、50 per second |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `im:chat` `im:chat.tabs:write_only` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `chat_id` | `string` | Chat ID. For details, refer to Group ID description. **Note**: Only supports group IDs whose group mode is `p2p` or `group`. **Example value**: "oc_a0553eda9014c201e6969b478895c230" | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `tab_ids` | `string[]` | No | List of chat tag IDs. Tab ID can be found in Add Chat Tab and Pull Chat Tab return value. **Note**: `tab_ids` must contain the full list of Tab IDs. **Example value**: ["6936075528890826780"] | ### Request body example

{
    "tab_ids": [
        "6936075528890826780"
    ]
}

## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `chat_tabs` | `chat.tab[]` | Chat tags |
|     `tab_id` | `string` | Tab ID |
|     `tab_name` | `string` | Tab name **Note**: The group tab name does not exceed 30 characters. |
|     `tab_type` | `string` | Tab type **Optional values are**:  - `message`: Message type - `doc_list`: List of Cloud Documents - `doc`: Document - `pin`: Pin - `meeting_minute`: Meeting notes - `chat_announcement`: Group announcement - `url`: URL - `file`: File  |
|     `tab_content` | `chat_tab_content` | Tab content |
|       `url` | `string` | URL type |
|       `doc` | `string` | Doc link |
|       `meeting_minute` | `string` | Meeting notes |
|     `tab_config` | `chat_tab_config` | Chat tab configuration |
|       `icon_key` | `string` | Chat tab icon key |
|       `is_built_in` | `boolean` | Whether the chat tab is opened inline in the app | ### Response body example

{
    "code": 0,
    "msg": "ok",
    "data": {
        "chat_tabs": [
            {
                "tab_content": {},
                "tab_id": "7104164142520467475",
                "tab_type": "message"
            },
            {
                "tab_content": {},
                "tab_id": "7104164246245605395",
                "tab_type": "pin"
            },
            {
                "tab_content": {
                    "url": "https://www.larksuite.com/"
                },
                "tab_id": "7104168465417633811",
                "tab_name": "url2",
                "tab_type": "url"
            },
            {
                "tab_content": {
                    "doc": "https://example.larksuite.com/docx/doxbcjoYDoEtuwC0k0hryQBkSV1"
                },
                "tab_id": "7104168465379885076",
                "tab_name": "doc2",
                "tab_type": "doc"
            },
            {
                "tab_content": {
                    "url": "https://www.larksuite.com/"
                },
                "tab_id": "7104168141097287699",
                "tab_name": "url1",
                "tab_type": "url"
            },
            {
                "tab_content": {
                    "doc": "https://example.larksuite.com/docx/doxbcjoYDoEtuwC0k0hryQBkSV1"
                },
                "tab_id": "7104168141063716884",
                "tab_name": "doc1",
                "tab_type": "doc"
            }
        ]
    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 232001 | Your request contains an invalid request parameter. | Parameter error. Check the input parameters according to the appropriate document. |
| 400 | 232006 | Your request specifies a chat_id which is invalid. | Invalid chat_id, please check if the chat_id is correct. |
| 400 | 232009 | Your request specifies a chat which has already been dissolved. | The group has been deleted. |
| 400 | 232010 | Operator and chat can NOT be in different tenants. | The operator and the group being operated should be under the same tenant. |
| 400 | 232011 | Operator can NOT be out of the chat. | The operator must be a member of the chat. |
| 400 | 232025 | Bot ability is not activated. | The bot ability needs to be enabled for the app. |
| 400 | 232033 | The operator or invited bots does NOT have the authority to manage external chats without the scope. | No permission to operate external groups. |
| 400 | 232034 | The app is unavailable or inactivated by the tenant. | The app is unavailable or inactivated by the tenant. |
| 400 | 232050 | Operate chat tab unsupported chat type. | The requested group type does not support chat tabs. |
| 400 | 232055 | The operator does not have chat tab, chat menu, chat widget manage permission | If you do not have management rights for chat tabs, chat menus, and widgets, please check whether "Only chat owners and administrators" is enabled in "Who can manage chat tabs, chat menus, and widgets" in the chat settings. |
