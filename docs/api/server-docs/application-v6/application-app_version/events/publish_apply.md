---
title: "Apply To\u00a0Publish an Application"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/application-v6/application-app_version/events/publish_apply"
service: "application-v6"
resource: "application-app_version"
section: "App Information"
scopes:
  - "application:application.app_version:readonly"
field_scopes:
  - "contact:user.employee_id:readonly"
updated: "1646720045000"
---

# App release application

This event is pushed when an app release application is submitted.{Usage Examples}(url=/api/tools/api_explore/api_explore_config?project=application&version=v6&resource=application.app_version&event=publish_apply)

## Event
| Facts |  |
| --- | --- |
| Event type | application.application.app_version.publish_apply_v6 |
| Supported app types | custom |
| Required scopes | `application:application.app_version:readonly` |
| Required field scopes | > The response body of the API contains the following sensitive fields, and they will be returned only after corresponding scopes are added. If you do not need the fields, it is not recommended that you request the scopes. `contact:user.employee_id:readonly` |
| Push Mode | `Webhook` | ### Event body
| Parameter | Type | Description |
| --- | --- | --- |
| `schema` | `string` | Event schema |
| `header` | `event_header` | Event header |
| ∟ `event_id` | `string` | Event ID |
| ∟ `event_type` | `string` | Event type |
| ∟ `create_time` | `string` | Event creation timestamp(in ms) |
| ∟ `token` | `string` | Event token |
| ∟ `app_id` | `string` | App ID |
| ∟ `tenant_key` | `string` | Tenant key |
| `event` | `\-` | \- |
| ∟ `operator_id` | `user_id` | User ID |
| ∟ `union_id` | `string` | User's union_id |
| ∟ `user_id` | `string` | user's user id **Required field scopes**: `contact:user.employee_id:readonly` |
| ∟ `open_id` | `string` | User's open_id |
| ∟ `online_version` | `application.app_version_event` | Information of current online version |
| ∟ `app_id` | `string` | App ID |
| ∟ `version` | `string` | App version ID entered by the developer **Data validation rules**: - Minimum length: `1` characters |
| ∟ `version_id` | `string` | ID that uniquely identifies the app version |
| ∟ `app_name` | `string` | Default app name **Data validation rules**: - Minimum length: `1` characters |
| ∟ `avatar_url` | `string` | App profile photo URL |
| ∟ `description` | `string` | Default app description **Data validation rules**: - Minimum length: `1` characters |
| ∟ `scopes` | `app_scope[]` | App scope list |
| ∟ `scope` | `string` | App scopes |
| ∟ `description` | `string` | Internationalized description of app scopes |
| ∟ `level` | `string` | Scope level description **Optional values are**: - `1`: General scope - `2`: Advanced scope - `3`: Highly sensitive scope - `0`: Unknown level |
| ∟ `back_home_url` | `string` | Admin homepage address |
| ∟ `i18n` | `app_i18n_info[]` | Internationalized information list of the app **Data validation rules**: - Minimum length: `1` |
| ∟ `i18n_key` | `string` | Internationalization language key **Optional values are**: - `zh_cn`: Chinese - `en_us`: English - `ja_jp`: Japanese |
| ∟ `name` | `string` | Internationalized app name **Data validation rules**: - Minimum length: `1` characters |
| ∟ `description` | `string` | Internationalized app description (subtitle) **Data validation rules**: - Minimum length: `1` characters |
| ∟ `help_use` | `string` | Link to the internationalized Help documentation |
| ∟ `common_categories` | `string[]` | Internationalized description of categories **Data validation rules**: - Maximum length: `3` |
| ∟ `events` | `string[]` | List of open platform events the app has subscribed to |
| ∟ `status` | `int` | Version status **Optional values are**: - `0`: Unknown status - `1`: Approved - `2`: Rejected - `3`: In review - `4`: Not submitted for review |
| ∟ `create_time` | `string` | Version creation time, in seconds |
| ∟ `publish_time` | `string` | Version release time, in seconds |
| ∟ `ability` | `app_ability` | App capabilities enabled in current version |
| ∟ `gadget` | `gadget` | Gadget |
| ∟ `enable_pc_mode` | `int` | Gadget mode supported by PC **Optional values are**: - `1`: Sidebar mode - `2`: PC mode - `4`: Global navigation mode |
| ∟ `schema_urls` | `string[]` | Schema URL list |
| ∟ `pc_use_mobile_pkg` | `boolean` | Whether a mobile package is used on PC |
| ∟ `pc_version` | `string` | PC gadget version **Data validation rules**: - Minimum length: `1` characters |
| ∟ `mobile_version` | `string` | Mobile gadget version **Data validation rules**: - Minimum length: `1` characters |
| ∟ `mobile_min_lark_version` | `string` | Minimum compatible Lark version on mobile **Data validation rules**: - Minimum length: `1` characters |
| ∟ `pc_min_lark_version` | `string` | Minimum compatible Lark version on PC **Data validation rules**: - Minimum length: `1` characters |
| ∟ `web_app` | `web_app` | Web |
| ∟ `pc_url` | `string` | URL on PC |
| ∟ `mobile_url` | `string` | URL on mobile |
| ∟ `bot` | `bot` | Bot |
| ∟ `card_request_url` | `string` | Message card callback address |
| ∟ `workplace_widgets` | `workplace_widget[]` | Block |
| ∟ `min_lark_version` | `string` | Minimum compatible Lark version |
| ∟ `navigate` | `navigate` | Global navigation gadget |
| ∟ `pc` | `navigate_meta` | Global navigation information on PC |
| ∟ `version` | `string` | Version number of the global navigation gadget **Data validation rules**: - Minimum length: `1` characters |
| ∟ `image_url` | `string` | URL of the default image |
| ∟ `hover_image_url` | `string` | URL of the selected status image |
| ∟ `mobile` | `navigate_meta` | Global navigation information on mobile |
| ∟ `version` | `string` | Version number of the global navigation gadget **Data validation rules**: - Minimum length: `1` characters |
| ∟ `image_url` | `string` | URL of the default image |
| ∟ `hover_image_url` | `string` | URL of the selected status image |
| ∟ `cloud_doc` | `cloud_doc` | Docs app |
| ∟ `space_url` | `string` | My Space redirect URL |
| ∟ `i18n` | `cloud_doc_i18n_info[]` | Internationalized information **Data validation rules**: - Minimum length: `1` |
| ∟ `i18n_key` | `string` | Internationalization language key **Optional values are**: - `zh_cn`: Chinese - `en_us`: English - `ja_jp`: Japanese |
| ∟ `name` | `string` | Internationalized Docs name |
| ∟ `read_description` | `string` | Internationalized description of read scope for Docs |
| ∟ `write_description` | `string` | Internationalized description of write scope for Docs |
| ∟ `icon_url` | `string` | Icon link |
| ∟ `mode` | `int` | Mode supported by Docs **Optional values are**: - `0`: Unknown - `1`: Mobile |
| ∟ `docs_blocks` | `docs_block[]` | Docs block |
| ∟ `block_type_id` | `string` | BlockTypeID |
| ∟ `i18n` | `block_i18n_info[]` | Internationalized block information **Data validation rules**: - Minimum length: `1` |
| ∟ `i18n_key` | `string` | Internationalization language key **Optional values are**: - `zh_cn`: Chinese - `en_us`: English - `ja_jp`: Japanese |
| ∟ `name` | `string` | Name **Data validation rules**: - Minimum length: `1` characters |
| ∟ `mobile_icon_url` | `string` | Mobile icon link |
| ∟ `pc_icon_url` | `string` | PC port icon link |
| ∟ `message_action` | `message_action` | Message shortcut |
| ∟ `pc_app_link` | `string` | Link on PC |
| ∟ `mobile_app_link` | `string` | Link on mobile |
| ∟ `i18n` | `message_action_i18n_info[]` | Internationalized information **Data validation rules**: - Minimum length: `1` |
| ∟ `i18n_key` | `string` | Internationalization language key **Optional values are**: - `zh_cn`: Chinese - `en_us`: English - `ja_jp`: Japanese |
| ∟ `name` | `string` | Internationalized name **Data validation rules**: - Minimum length: `1` characters |
| ∟ `plus_menu` | `plus_menu` | "+" menu |
| ∟ `pc_app_link` | `string` | Link on PC |
| ∟ `mobile_app_link` | `string` | Link on mobile |
| ∟ `remark` | `app_version_remark_event` | Other configuration information of the app version |
| ∟ `remark` | `string` | Notes |
| ∟ `update_remark` | `string` | Update description |
| ∟ `visibility` | `app_visibility_event` | Visibility list |
| ∟ `is_all` | `boolean` | Whether the app is visible to all |
| ∟ `visible_list` | `app_visible_list_event` | Visibility list |
| ∟ `open_ids` | `user_id[]` | List of IDs of members to whom the app is visible |
| ∟ `union_id` | `string` | User's union ID |
| ∟ `user_id` | `string` | User's user ID |
| ∟ `open_id` | `string` | User's open ID |
| ∟ `department_ids` | `string[]` | List of IDs of departments to which the app is visible |
| ∟ `invisible_list` | `app_visible_list_event` | Invisibility list |
| ∟ `open_ids` | `user_id[]` | List of IDs of members to whom the app is visible |
| ∟ `union_id` | `string` | User's union ID |
| ∟ `user_id` | `string` | User's user ID |
| ∟ `open_id` | `string` | User's open ID |
| ∟ `department_ids` | `string[]` | List of IDs of departments to which the app is visible |
| ∟ `under_audit_version` | `application.app_version_event` | Information of current online version |
| ∟ `app_id` | `string` | App ID |
| ∟ `version` | `string` | App version ID entered by the developer **Data validation rules**: - Minimum length: `1` characters |
| ∟ `version_id` | `string` | ID that uniquely identifies the app version |
| ∟ `app_name` | `string` | Default app name **Data validation rules**: - Minimum length: `1` characters |
| ∟ `avatar_url` | `string` | App profile photo URL |
| ∟ `description` | `string` | Default app description **Data validation rules**: - Minimum length: `1` characters |
| ∟ `scopes` | `app_scope[]` | App scope list |
| ∟ `scope` | `string` | App scopes |
| ∟ `description` | `string` | Internationalized description of app scopes |
| ∟ `level` | `string` | Scope level description **Optional values are**: - `1`: General scope - `2`: Advanced scope - `3`: Highly sensitive scope - `0`: Unknown level |
| ∟ `back_home_url` | `string` | Admin homepage address |
| ∟ `i18n` | `app_i18n_info[]` | Internationalized information list of the app **Data validation rules**: - Minimum length: `1` |
| ∟ `i18n_key` | `string` | Internationalization language key **Optional values are**: - `zh_cn`: Chinese - `en_us`: English - `ja_jp`: Japanese |
| ∟ `name` | `string` | Internationalized app name **Data validation rules**: - Minimum length: `1` characters |
| ∟ `description` | `string` | Internationalized app description (subtitle) **Data validation rules**: - Minimum length: `1` characters |
| ∟ `help_use` | `string` | Link to the internationalized Help documentation |
| ∟ `common_categories` | `string[]` | Internationalized description of categories **Data validation rules**: - Maximum length: `3` |
| ∟ `events` | `string[]` | List of open platform events the app has subscribed to |
| ∟ `status` | `int` | Version status **Optional values are**: - `0`: Unknown status - `1`: Approved - `2`: Rejected - `3`: In review - `4`: Not submitted for review |
| ∟ `create_time` | `string` | Version creation time, in seconds |
| ∟ `publish_time` | `string` | Version release time, in seconds |
| ∟ `ability` | `app_ability` | App capabilities enabled in current version |
| ∟ `gadget` | `gadget` | Gadget |
| ∟ `enable_pc_mode` | `int` | Gadget mode supported by PC **Optional values are**: - `1`: Sidebar mode - `2`: PC mode - `4`: Global navigation mode |
| ∟ `schema_urls` | `string[]` | Schema URL list |
| ∟ `pc_use_mobile_pkg` | `boolean` | Whether a mobile package is used on PC |
| ∟ `pc_version` | `string` | PC gadget version **Data validation rules**: - Minimum length: `1` characters |
| ∟ `mobile_version` | `string` | Mobile gadget version **Data validation rules**: - Minimum length: `1` characters |
| ∟ `mobile_min_lark_version` | `string` | Minimum compatible Lark version on mobile **Data validation rules**: - Minimum length: `1` characters |
| ∟ `pc_min_lark_version` | `string` | Minimum compatible Lark version on PC **Data validation rules**: - Minimum length: `1` characters |
| ∟ `web_app` | `web_app` | Web |
| ∟ `pc_url` | `string` | URL on PC |
| ∟ `mobile_url` | `string` | URL on mobile |
| ∟ `bot` | `bot` | Bot |
| ∟ `card_request_url` | `string` | Message card callback address |
| ∟ `workplace_widgets` | `workplace_widget[]` | Block |
| ∟ `min_lark_version` | `string` | Minimum compatible Lark version |
| ∟ `navigate` | `navigate` | Global navigation gadget |
| ∟ `pc` | `navigate_meta` | Global navigation information on PC |
| ∟ `version` | `string` | Version number of the global navigation gadget **Data validation rules**: - Minimum length: `1` characters |
| ∟ `image_url` | `string` | URL of the default image |
| ∟ `hover_image_url` | `string` | URL of the selected status image |
| ∟ `mobile` | `navigate_meta` | Global navigation information on mobile |
| ∟ `version` | `string` | Version number of the global navigation gadget **Data validation rules**: - Minimum length: `1` characters |
| ∟ `image_url` | `string` | URL of the default image |
| ∟ `hover_image_url` | `string` | URL of the selected status image |
| ∟ `cloud_doc` | `cloud_doc` | Docs app |
| ∟ `space_url` | `string` | My Space redirect URL |
| ∟ `i18n` | `cloud_doc_i18n_info[]` | Internationalized information **Data validation rules**: - Minimum length: `1` |
| ∟ `i18n_key` | `string` | Internationalization language key **Optional values are**: - `zh_cn`: Chinese - `en_us`: English - `ja_jp`: Japanese |
| ∟ `name` | `string` | Internationalized Docs name |
| ∟ `read_description` | `string` | Internationalized description of read scope for Docs |
| ∟ `write_description` | `string` | Internationalized description of write scope for Docs |
| ∟ `icon_url` | `string` | Icon link |
| ∟ `mode` | `int` | Mode supported by Docs **Optional values are**: - `0`: Unknown - `1`: Mobile |
| ∟ `docs_blocks` | `docs_block[]` | Docs block |
| ∟ `block_type_id` | `string` | BlockTypeID |
| ∟ `i18n` | `block_i18n_info[]` | Internationalized block information **Data validation rules**: - Minimum length: `1` |
| ∟ `i18n_key` | `string` | Internationalization language key **Optional values are**: - `zh_cn`: Chinese - `en_us`: English - `ja_jp`: Japanese |
| ∟ `name` | `string` | Name **Data validation rules**: - Minimum length: `1` characters |
| ∟ `mobile_icon_url` | `string` | Mobile icon link |
| ∟ `pc_icon_url` | `string` | PC port icon link |
| ∟ `message_action` | `message_action` | Message shortcut |
| ∟ `pc_app_link` | `string` | Link on PC |
| ∟ `mobile_app_link` | `string` | Link on mobile |
| ∟ `i18n` | `message_action_i18n_info[]` | Internationalized information **Data validation rules**: - Minimum length: `1` |
| ∟ `i18n_key` | `string` | Internationalization language key **Optional values are**: - `zh_cn`: Chinese - `en_us`: English - `ja_jp`: Japanese |
| ∟ `name` | `string` | Internationalized name **Data validation rules**: - Minimum length: `1` characters |
| ∟ `plus_menu` | `plus_menu` | "+" menu |
| ∟ `pc_app_link` | `string` | Link on PC |
| ∟ `mobile_app_link` | `string` | Link on mobile |
| ∟ `remark` | `app_version_remark_event` | Other configuration information of the app version |
| ∟ `remark` | `string` | Notes |
| ∟ `update_remark` | `string` | Update description |
| ∟ `visibility` | `app_visibility_event` | Visibility list |
| ∟ `is_all` | `boolean` | Whether the app is visible to all |
| ∟ `visible_list` | `app_visible_list_event` | Visibility list |
| ∟ `open_ids` | `user_id[]` | List of IDs of members to whom the app is visible |
| ∟ `union_id` | `string` | User's union ID |
| ∟ `user_id` | `string` | User's user ID |
| ∟ `open_id` | `string` | User's open ID |
| ∟ `department_ids` | `string[]` | List of IDs of departments to which the app is visible |
| ∟ `invisible_list` | `app_visible_list_event` | Invisibility list |
| ∟ `open_ids` | `user_id[]` | List of IDs of members to whom the app is visible |
| ∟ `union_id` | `string` | User's union ID |
| ∟ `user_id` | `string` | User's user ID |
| ∟ `open_id` | `string` | User's open ID |
| ∟ `department_ids` | `string[]` | List of IDs of departments to which the app is visible |
| ∟ `app_status` | `int` | App status **Optional values are**: - `0`: Disabled status - `1`: Enabled status | ### Event body example

```json
{
    "schema": "2.0",
    "header": {
        "event_id": "5e3702a84e847582be8db7fb73283c02",
        "event_type": "application.application.app_version.publish_apply_v6",
        "create_time": "1608725989000",
        "token": "rvaYgkND1GOiu5MM0E1rncYC6PLtF7JV",
        "app_id": "cli_9f5343c580712544",
        "tenant_key": "2ca1d211f64f6438"
    },
    "event": {
        "operator_id": {
            "union_id": "on_8ed6aa67826108097d9ee143816345",
            "user_id": "e33ggbyz",
            "open_id": "ou_84aad35d084aa403a838cf73ee18467"
        },
        "online_version": {
            "app_id": "cli_9f3ca975326b501b",
            "version": "1.0.0",
            "version_id": "oav_d317f090b7258ad0372aa53963cda70d",
            "app_name": "App name",
            "avatar_url": "https://www.example.com",
            "description": "App description",
            "scopes": [
                {
                    "scope": "contact:user.base",
                    "description": "Obtain apps information",
                    "level": "low_level"
                }
            ],
            "back_home_url": "https://www.example.com",
            "i18n": [
                {
                    "i18n_key": "zh_cn",
                    "name": "App name",
                    "description": "App description",
                    "help_use": "https://www.example.com"
                }
            ],
            "common_categories": [
                "Analysis tools"
            ],
            "events": [
                "App review event"
            ],
            "status": 1,
            "create_time": "1610462759",
            "publish_time": "1610462759",
            "ability": {
                "gadget": {
                    "enable_pc_mode": 1,
                    "schema_urls": [
                        "*:*"
                    ],
                    "pc_use_mobile_pkg": false,
                    "pc_version": "1.0.0",
                    "mobile_version": "1.0.0",
                    "mobile_min_lark_version": "2.0",
                    "pc_min_lark_version": "2.0"
                },
                "web_app": {
                    "pc_url": "https://www.example.com",
                    "mobile_url": "https://www.example.com"
                },
                "bot": {
                    "card_request_url": "https://www.example.com"
                },
                "workplace_widgets": [
                    {
                        "min_lark_version": "1.0.0"
                    }
                ],
                "navigate": {
                    "pc": {
                        "version": "1.0.0",
                        "image_url": "https://www.example.com",
                        "hover_image_url": "https://www.example.com"
                    },
                    "mobile": {
                        "version": "1.0.0",
                        "image_url": "https://www.example.com",
                        "hover_image_url": "https://www.example.com"
                    }
                },
                "cloud_doc": {
                    "space_url": "https://www.example.com",
                    "i18n": [
                        {
                            "i18n_key": "zh_cn",
                            "name": "Name",
                            "read_description": "Read scope description",
                            "write_description": "Write scope description"
                        }
                    ],
                    "icon_url": "https://www.example.com",
                    "mode": 1
                },
                "docs_blocks": [
                    {
                        "block_type_id": "blk_4fb61568435880110854c1d0",
                        "i18n": [
                            {
                                "i18n_key": "zh_cn",
                                "name": "Name"
                            }
                        ],
                        "mobile_icon_url": "https://www.example.com",
                        "pc_icon_url": "https://www.example.com"
                    }
                ],
                "message_action": {
                    "pc_app_link": "https://www.example.com",
                    "mobile_app_link": "https://www.example.com",
                    "i18n": [
                        {
                            "i18n_key": "zh_cn",
                            "name": "Name"
                        }
                    ]
                },
                "plus_menu": {
                    "pc_app_link": "https://www.example.com",
                    "mobile_app_link": "https://www.example.com"
                }
            },
            "remark": {
                "remark": "Notes",
                "update_remark": "Update description",
                "visibility": {
                    "is_all": false,
                    "visible_list": {
                        "open_ids": [
                            {
                                "union_id": "on_8ed6aa67826108097d9ee143816345",
                                "user_id": "e33ggbyz",
                                "open_id": "ou_84aad35d084aa403a838cf73ee18467"
                            }
                        ],
                        "department_ids": [
                            "od-ddee42c0f8a948a5e650341e2153243b"
                        ]
                    },
                    "invisible_list": {
                        "open_ids": [
                            {
                                "union_id": "on_8ed6aa67826108097d9ee143816345",
                                "user_id": "e33ggbyz",
                                "open_id": "ou_84aad35d084aa403a838cf73ee18467"
                            }
                        ],
                        "department_ids": [
                            "od-ddee42c0f8a948a5e650341e2153243b"
                        ]
                    }
                }
            }
        },
        "under_audit_version": {
            "app_id": "cli_9f3ca975326b501b",
            "version": "1.0.0",
            "version_id": "oav_d317f090b7258ad0372aa53963cda70d",
            "app_name": "App name",
            "avatar_url": "https://www.example.com",
            "description": "App description",
            "scopes": [
                {
                    "scope": "contact:user.base",
                    "description": "Obtain apps information",
                    "level": "low_level"
                }
            ],
            "back_home_url": "https://www.example.com",
            "i18n": [
                {
                    "i18n_key": "zh_cn",
                    "name": "App name",
                    "description": "App description",
                    "help_use": "https://www.example.com"
                }
            ],
            "common_categories": [
                "Analysis tools"
            ],
            "events": [
                "App review event"
            ],
            "status": 1,
            "create_time": "1610462759",
            "publish_time": "1610462759",
            "ability": {
                "gadget": {
                    "enable_pc_mode": 1,
                    "schema_urls": [
                        "*:*"
                    ],
                    "pc_use_mobile_pkg": false,
                    "pc_version": "1.0.0",
                    "mobile_version": "1.0.0",
                    "mobile_min_lark_version": "2.0",
                    "pc_min_lark_version": "2.0"
                },
                "web_app": {
                    "pc_url": "https://www.example.com",
                    "mobile_url": "https://www.example.com"
                },
                "bot": {
                    "card_request_url": "https://www.example.com"
                },
                "workplace_widgets": [
                    {
                        "min_lark_version": "1.0.0"
                    }
                ],
                "navigate": {
                    "pc": {
                        "version": "1.0.0",
                        "image_url": "https://www.example.com",
                        "hover_image_url": "https://www.example.com"
                    },
                    "mobile": {
                        "version": "1.0.0",
                        "image_url": "https://www.example.com",
                        "hover_image_url": "https://www.example.com"
                    }
                },
                "cloud_doc": {
                    "space_url": "https://www.example.com",
                    "i18n": [
                        {
                            "i18n_key": "zh_cn",
                            "name": "Name",
                            "read_description": "Read scope description",
                            "write_description": "Write scope description"
                        }
                    ],
                    "icon_url": "https://www.example.com",
                    "mode": 1
                },
                "docs_blocks": [
                    {
                        "block_type_id": "blk_4fb61568435880110854c1d0",
                        "i18n": [
                            {
                                "i18n_key": "zh_cn",
                                "name": "Name"
                            }
                        ],
                        "mobile_icon_url": "https://www.example.com",
                        "pc_icon_url": "https://www.example.com"
                    }
                ],
                "message_action": {
                    "pc_app_link": "https://www.example.com",
                    "mobile_app_link": "https://www.example.com",
                    "i18n": [
                        {
                            "i18n_key": "zh_cn",
                            "name": "Name"
                        }
                    ]
                },
                "plus_menu": {
                    "pc_app_link": "https://www.example.com",
                    "mobile_app_link": "https://www.example.com"
                }
            },
            "remark": {
                "remark": "Notes",
                "update_remark": "Update description",
                "visibility": {
                    "is_all": false,
                    "visible_list": {
                        "open_ids": [
                            {
                                "union_id": "on_8ed6aa67826108097d9ee143816345",
                                "user_id": "e33ggbyz",
                                "open_id": "ou_84aad35d084aa403a838cf73ee18467"
                            }
                        ],
                        "department_ids": [
                            "od-ddee42c0f8a948a5e650341e2153243b"
                        ]
                    },
                    "invisible_list": {
                        "open_ids": [
                            {
                                "union_id": "on_8ed6aa67826108097d9ee143816345",
                                "user_id": "e33ggbyz",
                                "open_id": "ou_84aad35d084aa403a838cf73ee18467"
                            }
                        ],
                        "department_ids": [
                            "od-ddee42c0f8a948a5e650341e2153243b"
                        ]
                    }
                }
            }
        },
        "app_status": 1
    }
}
```

### Sample code for event subscriptions

For event subscription process, refer to:Event Subscription overview，Guide for beginners:Tutorial

  Event subscription mode
  

  
	
    
package main

import (
	"context"
	"fmt"

	larkcore "github.com/larksuite/oapi-sdk-go/v3/core"
	larkevent "github.com/larksuite/oapi-sdk-go/v3/event"
	"github.com/larksuite/oapi-sdk-go/v3/event/dispatcher"
	"github.com/larksuite/oapi-sdk-go/v3/service/application/v6"
	larkws "github.com/larksuite/oapi-sdk-go/v3/ws"
)

// SDK 使用说明 SDK user guide：https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/server-side-sdk/golang-sdk-guide/preparations
func main() {
	// 注册事件 Register event
	eventHandler := dispatcher.NewEventDispatcher("", "").
		OnP2ApplicationAppVersionPublishApplyV6(func(ctx context.Context, event *larkapplication.P2ApplicationAppVersionPublishApplyV6) error {
			fmt.Printf("[ OnP2ApplicationAppVersionPublishApplyV6 access ], data: %s\n", larkcore.Prettify(event))
			return nil
		})

	// 构建 client Build client
	cli := larkws.NewClient("YOUR_APP_ID", "YOUR_APP_SECRET",
		larkws.WithEventHandler(eventHandler),
		larkws.WithLogLevel(larkcore.LogLevelDebug),
	)

	// 建立长连接 Establish persistent connection
	err := cli.Start(context.Background())

	if err != nil {
		panic(err)
	}
}

    

    
# SDK 使用说明 SDK user guide：https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/server-side-sdk/python--sdk/preparations-before-development
import lark_oapi as lark

def do_p2_application_application_app_version_publish_apply_v6(data: lark.application.v6.P2ApplicationApplicationAppVersionPublishApplyV6) -> None:
    print(f'[ do_p2_application_application_app_version_publish_apply_v6 access ], data: {lark.JSON.marshal(data, indent=4)}')

# 注册事件 Register event
event_handler = lark.EventDispatcherHandler.builder("", "") \
    .register_p2_application_application_app_version_publish_apply_v6(do_p2_application_application_app_version_publish_apply_v6) \
    .build()

def main():
    # 构建 client Build client
    cli = lark.ws.Client("APP_ID", "APP_SECRET",
                        event_handler=event_handler, log_level=lark.LogLevel.DEBUG)
    # 建立长连接 Establish persistent connection
    cli.start()

if __name__ == "__main__":
    main()

    

    

package com.example.sample;

import com.lark.oapi.core.utils.Jsons;
import com.lark.oapi.service.application.ApplicationService;
import com.lark.oapi.service.application.v6.model.P2ApplicationAppVersionPublishApplyV6;
import com.lark.oapi.event.EventDispatcher;
import com.lark.oapi.ws.Client;

// SDK 使用说明 SDK user guide：https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/server-side-sdk/java-sdk-guide/preparations
public class Sample {
    // 注册事件 Register event
    private static final EventDispatcher EVENT_HANDLER = EventDispatcher.newBuilder("", "")
            .onP2ApplicationAppVersionPublishApplyV6(new ApplicationService.P2ApplicationAppVersionPublishApplyV6Handler() {
                @Override
                public void handle(P2ApplicationAppVersionPublishApplyV6 event) throws Exception {
                    System.out.printf("[ onP2ApplicationAppVersionPublishApplyV6 access ], data: %s\n", Jsons.DEFAULT.toJson(event.getEvent()));
                }
            })
            .build();

    public static void main(String[] args) {
        // 构建 client Build client
        Client client = new Client.Builder("APP_ID", "APP_SECRET")
                .eventHandler(EVENT_HANDLER)
                .build();
        // 建立长连接 Establish persistent connection
        client.start();
    }
}
    

    
// SDK 使用说明 SDK user guide：https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/server-side-sdk/nodejs-sdk/preparation-before-development
import * as Lark from '@larksuiteoapi/node-sdk';
const baseConfig = {
    appId: 'APP_ID',
    appSecret: 'APP_SECRET'
}
// 构建 client Build client
const wsClient = new Lark.WSClient(baseConfig);
// 建立长连接 Establish persistent connection
wsClient.start({
    // 注册事件 Register event
    eventDispatcher: new Lark.EventDispatcher({}).register({
        'application.application.app_version.publish_apply_v6': async (data) => {
            console.log(data);
        }
    })
});
    

  
  
	
    
package main

import (
	"context"
	"fmt"
	"net/http"

	larkcore "github.com/larksuite/oapi-sdk-go/v3/core"
	"github.com/larksuite/oapi-sdk-go/v3/core/httpserverext"
	larkevent "github.com/larksuite/oapi-sdk-go/v3/event"
	"github.com/larksuite/oapi-sdk-go/v3/event/dispatcher"
	"github.com/larksuite/oapi-sdk-go/v3/service/application/v6"
)

// SDK 使用说明 SDK user guide：https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/server-side-sdk/golang-sdk-guide/preparations
func main() {
	// 注册事件 Register event
	eventHandler := dispatcher.NewEventDispatcher("", "").
		OnP2ApplicationAppVersionPublishApplyV6(func(ctx context.Context, event *larkapplication.P2ApplicationAppVersionPublishApplyV6) error {
			fmt.Printf("[ OnP2ApplicationAppVersionPublishApplyV6 access ], data: %s\n", larkcore.Prettify(event))
			return nil
		})

	// 创建路由处理器 Create route handler
	http.HandleFunc("/webhook/event", httpserverext.NewEventHandlerFunc(handler, larkevent.WithLogLevel(larkcore.LogLevelDebug)))

	err := http.ListenAndServe(":7777", nil)

	if err != nil {
		panic(err)
	}
}

    

    
# SDK 使用说明 SDK user guide：https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/server-side-sdk/python--sdk/preparations-before-development
from flask import Flask
from lark_oapi.adapter.flask import *
import lark_oapi as lark

app = Flask(__name__)

def do_p2_application_application_app_version_publish_apply_v6(data: lark.application.v6.P2ApplicationApplicationAppVersionPublishApplyV6) -> None:
    print(f'[ do_p2_application_application_app_version_publish_apply_v6 access ], data: {lark.JSON.marshal(data, indent=4)}')

# 注册事件 Register event
event_handler = lark.EventDispatcherHandler.builder("", "") \
    .register_p2_application_application_app_version_publish_apply_v6(do_p2_application_application_app_version_publish_apply_v6) \
    .build()

# 创建路由处理器 Create route handler
@app.route("/webhook/event", methods=["POST"])
def event():
    resp = event_handler.do(parse_req())
    return parse_resp(resp)

if __name__ == "__main__":
    app.run(port=7777)

    

    

package com.lark.oapi.sample.event;

import com.lark.oapi.core.utils.Jsons;
import com.lark.oapi.service.application.ApplicationService;
import com.lark.oapi.service.application.v6.model.P2ApplicationAppVersionPublishApplyV6;
import com.lark.oapi.sdk.servlet.ext.ServletAdapter;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;

// SDK 使用说明 SDK user guide：https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/server-side-sdk/java-sdk-guide/preparations
@RestController
public class EventController {

    // 注册事件 Register event
    private static final EventDispatcher EVENT_HANDLER = EventDispatcher.newBuilder("verificationToken", "encryptKey")
            .onP2ApplicationAppVersionPublishApplyV6(new ApplicationService.P2ApplicationAppVersionPublishApplyV6Handler() {
                @Override
                public void handle(P2ApplicationAppVersionPublishApplyV6 event) throws Exception {
                    System.out.printf("[ onP2ApplicationAppVersionPublishApplyV6 access ], data: %s\n", Jsons.DEFAULT.toJson(event.getEvent()));
                }
            })
            .build();

    // 注入 ServletAdapter 实例 Inject ServletAdapter instance
    @Autowired
    private ServletAdapter servletAdapter;

    // 创建路由处理器 Create route handler
    @RequestMapping("/webhook/event")
    public void event(HttpServletRequest request, HttpServletResponse response)
            throws Throwable {
        // 回调扩展包提供的事件回调处理器 Callback handler provided by the extension package
        servletAdapter.handleEvent(request, response, EVENT_DISPATCHER);
    }
}
    

    
// SDK 使用说明 SDK user guide：https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/server-side-sdk/nodejs-sdk/preparation-before-development
import http from 'http';
import * as lark from '@larksuiteoapi/node-sdk';

// 注册事件 Register event
const eventDispatcher = new lark.EventDispatcher({
    encryptKey: '',
    verificationToken: '',
}).register({
    'application.application.app_version.publish_apply_v6': async (data) => {
        console.log(data);
        return 'success';
    },
});

const server = http.createServer();
// 创建路由处理器 Create route handler
server.on('request', lark.adaptDefault('/webhook/event', eventDispatcher));
server.listen(3000);
