---
title: "Add chat tab"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat-tab/create"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/im/v1/chats/:chat_id/chat_tabs"
service: "im-v1"
resource: "chat-tab"
section: "Group Chat"
rate_limit: "1000 per minute、50 per second"
scopes:
  - "im:chat"
  - "im:chat.tabs:write_only"
updated: "1717574957000"
---

# Add chat tab

Add a custom chat tab.

> Notice:
> - The bot ability needs to be enabled.
> - The bot or authorized user must be in the group.
> - Only allow to add group tabs of type `doc` and `url`.
> - The tab_config field is currently only valid for session tabs of type `url`
> - Only the group owner and group admins can add chat tabs when the setting Only the group owner and admins can manage tabs is turned on.
> - When adding doc type, the operator (the identity corresponding to the access token) needs to have the permission of the corresponding document.
> - When operating internal group, the operator must be under the same tenant as the group.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/im/v1/chats/:chat_id/chat_tabs |
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
| `chat_tabs` | `chat.tab[]` | Yes | Chat tags **Note**: A maximum of 20 custom group tabs are allowed to be added in a group. |
|   `tab_name` | `string` | No | Tab name **Note**: The group tab name does not exceed 30 characters (maximum 10 Chinese characters). **Example value**: "doc" |
|   `tab_type` | `string` | Yes | Tab type **Example value**: "doc" **Optional values are**:  - `message`: Message type - `doc_list`: List of Cloud Documents - `doc`: Document - `pin`: Pin - `meeting_minute`: Meeting notes - `chat_announcement`: Group announcement - `url`: URL - `file`: File  |
|   `tab_content` | `chat_tab_content` | No | Tab content |
|     `url` | `string` | No | URL type **Example value**: "https://www.larksuite.com" |
|     `doc` | `string` | No | Doc link **Example value**: "https://example.larksuite.com/wiki/wikcnPIcqWjJQwkwDzrB9t40123xz" |
|     `meeting_minute` | `string` | No | Meeting notes **Example value**: "https://example.larksuite.com/docs/doccnvIXbV22i6hSD3utar4123dx" |
|   `tab_config` | `chat_tab_config` | No | Chat tab configuration |
|     `icon_key` | `string` | No | Chat tab icon key **Example value**: "img_v2_b99741-7628-4abd-aad0-b881e4db83ig" |
|     `is_built_in` | `boolean` | No | Whether the chat tab is opened inline in the app **Example value**: false | ### Request body example

{
    "chat_tabs": [
        {
            "tab_name": "doc",
            "tab_type": "doc",
            "tab_content": {
                "url": "https://www.larksuite.com",
                "doc": "https://example.larksuite.com/wiki/wikcnPIcqWjJQwkwDzrB9t40123xz",
                "meeting_minute": "https://example.larksuite.com/docs/doccnvIXbV22i6hSD3utar4123dx"
            },
            "tab_config": {
                "icon_key": "img_v2_b99741-7628-4abd-aad0-b881e4db83ig",
                "is_built_in": false
            }
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
|   `chat_tabs` | `chat.tab[]` | Group tags |
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
               "tab_id": "7101214603622940633",
               "tab_type": "message"
            },
            {
                "tab_id": "7101214603622940671",
                "tab_name": "Doc",
                "tab_type": "doc",
                "tab_content": {
                    "doc": "https://example.larksuite.com/wiki/wikcnPIcqWjJQwkwDzrB9t40123xz"
                }
            },
            {
                "tab_id": "7158333373373759422",
                "tab_name": "Test",
                "tab_type": "url",
                "tab_content": {
                    "url": "https://www.test.cn"
                },
                "tab_config": {
                    "icon_key": "img_v2_bcaee32c-15f9-431e-87a4-48898b123zyz",
                    "is_built_in": true
                }
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
| 400 | 232046 | The number of chat tabs reaches the limit. | The number of custom chat tabs exceeds the upper limit of 20. |
| 400 | 232047 | The length of the tab name reaches the limit. | Chat tab name is too long. |
| 400 | 232048 | The chat tab content is illegal. | Chat tab content is illegal. |
| 400 | 232050 | Operate chat tab unsupported chat type. | The requested group type does not support chat tabs. |
| 400 | 232051 | The operator does not have doc permission | The operator must have document permissions |
| 400 | 232055 | The operator does not have chat tab, chat menu, chat widget manage permission | If you do not have management rights for chat tabs, chat menus, and widgets, please check whether "Only chat owners and administrators" is enabled in "Who can manage chat tabs, chat menus, and widgets" in the chat settings. |
| 400 | 232056 | The operator is not the image's owner, no permission to complete the request. | Robots need to use their own uploaded images. |
