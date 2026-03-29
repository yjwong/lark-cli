---
title: "Get chat menus"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat-menu_tree/get"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/im/v1/chats/:chat_id/menu_tree"
service: "im-v1"
resource: "chat-menu_tree"
section: "Group Chat"
rate_limit: "1000 per minute、50 per second"
scopes:
  - "im:chat"
  - "im:chat.menu_tree:read"
  - "im:chat:readonly"
updated: "1717574958000"
---

# Get chat menus

Get the menu in the chat through the chat ID.

> Notice:
> - The bot ability needs to be enabled.
> - The bot must be in the chat.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/im/v1/chats/:chat_id/menu_tree |
| HTTP Method | GET |
| Rate Limit | 1000 per minute、50 per second |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `im:chat` `im:chat.menu_tree:read` `im:chat:readonly` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `chat_id` | `string` | Chat ID. For details, refer to Group ID description. **Note**: Only supports group IDs whose group mode is `group`. **Example value**: "oc_a0553eda9014c201e6969b478895c230" | ## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `menu_tree` | `chat.menu_tree` | All menus in the chat. |
|     `chat_menu_top_levels` | `chat_menu_top_level[]` | Top level menu |
|       `chat_menu_top_level_id` | `string` | Top level chat menu's id |
|       `chat_menu_item` | `chat_menu_item` | Top level menu info |
|         `action_type` | `string` | Menu type **Notice** - If the top_level menu has a second_level menu, the value of this top_level menu must be NONE. **Optional values are**:  - `NONE`: NONE type, when only the top level menu has a second level menu, the top level menu is set to NONE type. - `REDIRECT_LINK`: REDIRECT_LINK type  |
|         `redirect_link` | `chat_menu_item_redirect_link` | Redirect link |
|           `common_url` | `string` | Common redirect link, must start with http. |
|           `ios_url` | `string` | IOS client redirect link, when this field is not filled, IOS client will use common_url. Must start with http. |
|           `android_url` | `string` | Android client redirect link, when this field is not filled, Android client will use common_url. Must start with http. |
|           `pc_url` | `string` | PC client redirect link, when this field is not filled, PC client will use common_url. Must start with http. After clicking the chat menu on the PC side, if you want the page corresponding to the url to be displayed on the sidebar of lark, you can add https://applink.larksuite.com/client/web_url/open?mode=sidebar-semi&url= before the url , such as https://applink.larksuite.com/client/web_url/open?mode=sidebar-semi&url=https://open.larksuite.com/ |
|           `web_url` | `string` | Web client redirect link, when this field is not filled, Web client will use common_url. Must start with http. |
|         `image_key` | `string` | image_key, upload a message type image through the upload image interface to obtain image_key. **Notice** - If a top_level menu has a second_level menu, this top_level menu cannot have an icon. |
|         `name` | `string` | Menu name. **Notice** - The number of characters in the name must be in the range of 1 to 120. |
|         `i18n_names` | `i18n_names` | Menu I18n names. **Notice** - The number of characters in the name must be in the range of 1 to 120. |
|           `zh_cn` | `string` | Chinese name |
|           `en_us` | `string` | English name |
|           `ja_jp` | `string` | Japanese name |
|       `children` | `chat_menu_second_level[]` | Top level menu |
|         `chat_menu_second_level_id` | `string` | Second level chat menu's id |
|         `chat_menu_item` | `chat_menu_item` | Second level menu |
|           `action_type` | `string` | Menu type **Notice** - If the top_level menu has a second_level menu, the value of this top_level menu must be NONE. **Optional values are**:  - `NONE`: NONE type, when only the top level menu has a second level menu, the top level menu is set to NONE type. - `REDIRECT_LINK`: REDIRECT_LINK type  |
|           `redirect_link` | `chat_menu_item_redirect_link` | Redirect link |
|             `common_url` | `string` | Common redirect link, must start with http. |
|             `ios_url` | `string` | IOS client redirect link, when this field is not filled, IOS client will use common_url. Must start with http. |
|             `android_url` | `string` | Android client redirect link, when this field is not filled, Android client will use common_url. Must start with http. |
|             `pc_url` | `string` | PC client redirect link, when this field is not filled, PC client will use common_url. Must start with http. After clicking the chat menu on the PC side, if you want the page corresponding to the url to be displayed on the sidebar of lark, you can add https://applink.larksuite.com/client/web_url/open?mode=sidebar-semi&url= before the url , such as https://applink.larksuite.com/client/web_url/open?mode=sidebar-semi&url=https://open.larksuite.com/ |
|             `web_url` | `string` | Web client redirect link, when this field is not filled, Web client will use common_url. Must start with http. |
|           `image_key` | `string` | image_key, upload a message type image through the upload image interface to obtain image_key. **Notice** - If a top_level menu has a second_level menu, this top_level menu cannot have an icon. |
|           `name` | `string` | Menu name. **Notice** - The number of characters in the name must be in the range of 1 to 120. |
|           `i18n_names` | `i18n_names` | Menu I18n names. **Notice** - The number of characters in the name must be in the range of 1 to 120. |
|             `zh_cn` | `string` | Chinese name |
|             `en_us` | `string` | English name |
|             `ja_jp` | `string` | Japanese name | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
        "menu_tree": {
            "chat_menu_top_levels": [
                {
                    "chat_menu_top_level_id": "7117116451961487361",
                    "chat_menu_item": {
                        "action_type": "NONE",
                        "redirect_link": {
                            "common_url": "https://open.larksuite.com/",
                            "ios_url": "https://open.larksuite.com/",
                            "android_url": "https://open.larksuite.com/",
                            "pc_url": "https://open.larksuite.com/",
                            "web_url": "https://open.larksuite.com/"
                        },
                        "image_key": "img_v2_b0fbe905-7988-4282-b882-82edd010336j",
                        "name": "菜单",
                        "i18n_names": {
                            "zh_cn": "菜单",
                            "en_us": "Menu",
                            "ja_jp": "メニュー"
                        }
                    },
                    "children": [
                        {
                            "chat_menu_second_level_id": "7039638308221468675",
                            "chat_menu_item": {
                                "action_type": "REDIRECT_LINK",
                                "redirect_link": {
                                    "common_url": "https://open.larksuite.com/",
                                    "ios_url": "https://open.larksuite.com/",
                                    "android_url": "https://open.larksuite.com/",
                                    "pc_url": "https://open.larksuite.com/",
                                    "web_url": "https://open.larksuite.com/"
                                },
                                "image_key": "img_v2_b0fbe905-7988-4282-b882-82edd010336j",
                                "name": "报名",
                                "i18n_names": {
                                    "zh_cn": "报名",
                                    "en_us": "Sign up",
                                    "ja_jp": "サインアップ"
                                }
                            }
                        }
                    ]
                }
            ]
        }
    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 232001 | Your request contains an invalid request parameter. | Parameter error. Check the input parameters according to the appropriate document. |
| 400 | 232011 | Operator can NOT be out of the chat. | The operator must be a member of the chat. |
| 400 | 232014 | The token used in the request does NOT have the permissions necessary to complete the request. | The operator does not have permission to perform this operation, please check whether it has the corresponding permission. |
| 400 | 232025 | Bot ability is not activated. | The bot ability needs to be enabled for the app. |
| 400 | 232034 | The app is unavailable or inactivated by the tenant. | The app is unavailable or inactivated by the tenant. |
