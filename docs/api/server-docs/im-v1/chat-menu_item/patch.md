---
title: "Patch chat menu item info"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat-menu_item/patch"
method: "PATCH"
api_path: "https://open.larksuite.com/open-apis/im/v1/chats/:chat_id/menu_items/:menu_item_id"
service: "im-v1"
resource: "chat-menu_item"
section: "Group Chat"
rate_limit: "1000 per minute、50 per second"
scopes:
  - "im:chat"
  - "im:chat.menu_tree:write_only"
updated: "1717574958000"
---

# Patch chat menu item info

Patch the meta information of a top_level menu or second_level menu, including the icon, name, i18n name and redirect link of the chat menu.

> Notice:
> - The bot ability needs to be enabled.
> - The bot must be in the chat.
> - This API does not currently support adding or deleting a second level menu on the top level menu.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/im/v1/chats/:chat_id/menu_items/:menu_item_id |
| HTTP Method | PATCH |
| Rate Limit | 1000 per minute、50 per second |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `im:chat` `im:chat.menu_tree:write_only` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `chat_id` | `string` | Chat ID. For details, refer to Group ID description. **Note**: Only supports group IDs whose group mode is `group`. **Example value**: "oc_a0553eda9014c201e6969b478895c230" |
| `menu_item_id` | `string` | ID of top level menu or second level menu. Get the menu's ID through the Get Chat Menu. **Example value**: "7156553273518882844" | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `update_fields` | `string[]` | Yes | Field to modify **Example value**: ["ICON"] **Optional values are**:  - `ICON`: Icon - `NAME`: Name - `I18N_NAME`: I18n name - `REDIRECT_LINK`: redirect link  |
| `chat_menu_item` | `chat_menu_item` | Yes | Item info |
|   `action_type` | `string` | No | Menu type **Notice** - If the top_level menu has a second_level menu, the value of this top_level menu must be NONE. **Example value**: "NONE" **Optional values are**:  - `NONE`: NONE type, when only the top level menu has a second level menu, the top level menu is set to NONE type. - `REDIRECT_LINK`: REDIRECT_LINK type  |
|   `redirect_link` | `chat_menu_item_redirect_link` | No | Redirect link |
|     `common_url` | `string` | No | Common redirect link, must start with http. **Example value**: "https://open.larksuite.com/" |
|     `ios_url` | `string` | No | IOS client redirect link, when this field is not filled, IOS client will use common_url. Must start with http. **Example value**: "https://open.larksuite.com/" |
|     `android_url` | `string` | No | Android client redirect link, when this field is not filled, Android client will use common_url. Must start with http. **Example value**: "https://open.larksuite.com/" |
|     `pc_url` | `string` | No | PC client redirect link, when this field is not filled, PC client will use common_url. Must start with http. After clicking the chat menu on the PC side, if you want the page corresponding to the url to be displayed on the sidebar of lark, you can add https://applink.larksuite.com/client/web_url/open?mode=sidebar-semi&url= before the url , such as https://applink.larksuite.com/client/web_url/open?mode=sidebar-semi&url=https://open.larksuite.com/ **Example value**: "https://open.larksuite.com/" |
|     `web_url` | `string` | No | Web client redirect link, when this field is not filled, Web client will use common_url. Must start with http. **Example value**: "https://open.larksuite.com/" |
|   `image_key` | `string` | No | image_key, upload a message type image through the upload image interface to obtain image_key. **Notice** - If a top_level menu has a second_level menu, this top_level menu cannot have an icon. **Example value**: "img_v2_b0fbe905-7988-4282-b882-82edd010336j" |
|   `name` | `string` | No | Menu name. **Notice** - The number of characters in the name must be in the range of 1 to 120. **Example value**: "group chat" |
|   `i18n_names` | `i18n_names` | No | Menu I18n names. **Notice** - The number of characters in the name must be in the range of 1 to 120. |
|     `zh_cn` | `string` | No | Chinese name **Example value**: "评审报名" |
|     `en_us` | `string` | No | English name **Example value**: "Sign up" |
|     `ja_jp` | `string` | No | Japanese name **Example value**: "サインアップ" | ### Request body example

{
    "update_fields": [
        "ICON"
    ],
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
        "name": "group chat",
        "i18n_names": {
            "zh_cn": "评审报名",
            "en_us": "Sign up",
            "ja_jp": "サインアップ"
        }
    }
}

## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `chat_menu_item` | `chat_menu_item` | Modified menu item info. |
|     `action_type` | `string` | Menu type **Notice** - If the top_level menu has a second_level menu, the value of this top_level menu must be NONE. **Optional values are**:  - `NONE`: NONE type, when only the top level menu has a second level menu, the top level menu is set to NONE type. - `REDIRECT_LINK`: REDIRECT_LINK type  |
|     `redirect_link` | `chat_menu_item_redirect_link` | Redirect link |
|       `common_url` | `string` | Common redirect link, must start with http. |
|       `ios_url` | `string` | IOS client redirect link, when this field is not filled, IOS client will use common_url. Must start with http. |
|       `android_url` | `string` | Android client redirect link, when this field is not filled, Android client will use common_url. Must start with http. |
|       `pc_url` | `string` | PC client redirect link, when this field is not filled, PC client will use common_url. Must start with http. After clicking the chat menu on the PC side, if you want the page corresponding to the url to be displayed on the sidebar of lark, you can add https://applink.larksuite.com/client/web_url/open?mode=sidebar-semi&url= before the url , such as https://applink.larksuite.com/client/web_url/open?mode=sidebar-semi&url=https://open.larksuite.com/ |
|       `web_url` | `string` | Web client redirect link, when this field is not filled, Web client will use common_url. Must start with http. |
|     `image_key` | `string` | image_key, upload a message type image through the upload image interface to obtain image_key. **Notice** - If a top_level menu has a second_level menu, this top_level menu cannot have an icon. |
|     `name` | `string` | Menu name. **Notice** - The number of characters in the name must be in the range of 1 to 120. |
|     `i18n_names` | `i18n_names` | Menu I18n names. **Notice** - The number of characters in the name must be in the range of 1 to 120. |
|       `zh_cn` | `string` | Chinese name |
|       `en_us` | `string` | English name |
|       `ja_jp` | `string` | Japanese name | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
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
            "name": "报名",
            "i18n_names": {
                "zh_cn": "报名",
                "en_us": "Sign up",
                "ja_jp": "サインアップ"
            }
        }
    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 232001 | Your request contains an invalid request parameter. | Parameter error. Check the input parameters according to the appropriate document. |
| 400 | 232009 | Your request specifies a chat which has already been dissolved. | The group has been deleted. |
| 400 | 232011 | Operator can NOT be out of the chat. | The operator must be a member of the chat. |
| 400 | 232014 | The token used in the request does NOT have the permissions necessary to complete the request. | The operator does not have permission to perform this operation, please check whether it has the corresponding permission. |
| 400 | 232025 | Bot ability is not activated. | The bot ability needs to be enabled for the app. |
| 400 | 232034 | The app is unavailable or inactivated by the tenant. | The app is unavailable or inactivated by the tenant. |
| 400 | 232055 | The operator does not have chat tab, chat menu, chat widget manage permission | If you do not have management rights for chat tabs, chat menus, and widgets, please check whether "Only chat owners and administrators" is enabled in "Who can manage chat tabs, chat menus, and widgets" in the chat settings. |
| 400 | 232073 | The group menu is being modified, please try again later. | The group menu is being modified, please try again later. |
| 400 | 232078 | The chat menu is illegal. | The chat menu is illegal. |
