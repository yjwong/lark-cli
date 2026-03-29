---
title: "Resource introduction"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat-menu_tree/overview"
service: "im-v1"
resource: "chat-menu_tree"
section: "Group Chat"
updated: "1717574958000"
---

# Resource introduction
## Introduction to chat menu
The menu is divided into a top_level menu (chat_menu_top_level) and a second_level menu (chat_menu_second_level). There are at most 3 top_level menus in a chat, and each top_level menu can have 0 to 5 second_level menus.

| 名称         | 描述        |
| --------- | --------------- | -------   | ----------- | --------- |
|`menu_tree` |The menu tree consists of several top_level menus. |
|`chat_menu_top_level` | top_level menus, there are at most 3 top_level menus in a chat, and each top_level menu can have 0 to 5 second_level menus. The picture ① below is a top_level menu with a second_level menu, and the picture ② below is a top_level menu without a second_level menu. |
|`chat_menu_second_level` | The second_level menu is attached to the top_level menu. The figure ③ below shows several second_level menus attached to the top_level menu.|
|`chat_menu_item` | Menu meta information, a structure shared by the top_level menu and the second_level menu to describe the meta-information. | ![20221208-120155.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/ba23f2965c4542d6be885c6336e35e72_Q8OVD9SRlV.png?lazyload=true&width=750&height=1624)

## Introduction to Chat Menu Fields
**chat_menu_item Field Introduction**
| 名称         | 描述        |
| --------- | --------------- | -------   | ----------- | --------- |
|`action_type` | Menu type, there are two types: NONE and REDIRECT_LINK. In general, the two types of REDIRECT_LINK are filled. When only the top_level menu has a second_level menu, the top_level menu is set to NONE type.|
|`redirect_link` | redirect_link |
|      `∟ common_url` | 	Common redirect link, must start with http. |
|      `∟ ios_url` | 	IOS client redirect link, when this field is not filled, IOS client will use common_url. Must start with http. |
|      `∟ android_url` |Android client redirect link, when this field is not filled, Android client will use common_url. Must start with http. |
|      `∟ pc_url` | 	PC client redirect link, when this field is not filled, PC client will use common_url. Must start with http. |
|      `∟ web_url` | Web client redirect link, when this field is not filled, Web client will use common_url. Must start with http. |
|`image_key` | image_key |
|`name` | 	Name, the number of characters in the name of the top level menu must be in the range of 1 to 8, and the number of characters in the name of the second level menu must be in the range of 1 to 24. |
|`i18n_names` | 	I18n names, the number of characters in the name of the top level menu must be in the range of 1 to 8, and the number of characters in the name of the second level menu must be in the range of 1 to 24. |
|      `∟ zh_cn` | Chinese name|
|      `∟ en_us` | English name |
|      `∟ ja_jp` | Japanese name |
