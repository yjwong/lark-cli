---
title: "Obtain App Version List"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/application-v6/application-app_version/list"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/application/v6/applications/:app_id/app_versions"
service: "application-v6"
resource: "application-app_version"
section: "App Information"
rate_limit: "100 per minute"
scopes:
  - "application:application:self_manage"
  - "application:application.app_version:readonly"
field_scopes:
  - "contact:user.employee_id:readonly"
updated: "1681902087000"
---

# Obtain app version list

Obtains the version list of an app based on app_id.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/application/v6/applications/:app_id/app_versions |
| HTTP Method | GET |
| Rate Limit | 100 per minute |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `application:application:self_manage` `application:application.app_version:readonly` |
| Required field scopes | > The response body of the API contains the following sensitive fields, and they will be returned only after corresponding scopes are added. If you do not need the fields, it is not recommended that you request the scopes. `contact:user.employee_id:readonly` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `app_id` | `string` | App's app_id. To query the version information of other apps, you must request the Obtain app version information scope. To only query the version information of this app, enter "me" or the app_id. **Example value**: "cli_9b445f5258795107" | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `lang` | `string` | Yes | Language of app information **Example value**: "zh_cn" **Optional values are**:  - `zh_cn`: Chinese - `en_us`: English - `ja_jp`: Japanese  **Data validation rules**: - Minimum length: `1` characters |
| `page_size` | `int` | No | Page size **Example value**: 10 **Default value**: `20` **Data validation rules**: - Value range: `1` ～ `50` |
| `page_token` | `string` | No | Page identifier. It is not filled in the first request, indicating traversal from the beginning; when there will be more groups, the new page_token will be returned at the same time, and the next traversal can use the page_token to get more groups **Example value**: "new-e3c5a0627cdf0c2e057da7257b90376a" |
| `order` | `int` | No | 0：Sort by creation time from large to small.   1：Sort by creation time from small to large **Example value**: 0 **Default value**: `0` |
| `user_id_type` | `string` | No | User ID categories **Example value**: "open_id" **Optional values are**:  - `open_id`: Identifies a user to an app. The same user has different Open IDs in different apps. How to get Open ID - `union_id`: Identifies a user to a tenant that acts as a developer. A user has the same Union ID in apps developed by the same developer, and has different Union IDs in apps developed by different developers. A developer can use Union ID to link the same user's identities in multiple apps.How to get Union ID - `user_id`: Identifies a user to a tenant. The same user has different User IDs in different tenants. In one single tenant, a user has the same User ID in all apps （including store apps）. User ID is usually used to communicate user data between different apps. How to get User ID  **Default value**: `open_id` **When the value is `user_id`, the following field scopes are required**: `contact:user.employee_id:readonly` | ## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `items` | `application.app_version[]` | App version list |
|     `app_id` | `string` | App ID |
|     `version` | `string` | App version number entered in the Developer Console |
|     `version_id` | `string` | ID that uniquely identifies the app version |
|     `app_name` | `string` | Default app name |
|     `avatar_url` | `string` | App profile photo URL |
|     `description` | `string` | Default app description |
|     `scopes` | `app_scope[]` | App scope list |
|       `scope` | `string` | App scopes |
|       `description` | `string` | Internationalized description of app scopes |
|       `level` | `int` | Scope level description **Optional values are**:  - `1`: General scope - `2`: Advanced scope - `3`: Highly sensitive scope - `0`: Unknown level  |
|     `back_home_url` | `string` | Admin homepage address |
|     `i18n` | `app_i18n_info[]` | Internationalized information list of the app |
|       `i18n_key` | `string` | Internationalization language key **Optional values are**:  - `zh_cn`: Chinese - `en_us`: English - `ja_jp`: Japanese  |
|       `name` | `string` | Internationalized app name |
|       `description` | `string` | Internationalized app description (subtitle) |
|       `help_use` | `string` | Link to the internationalized Help documentation |
|     `common_categories` | `string[]` | Internationalized description of categories |
|     `events` | `string[]` | List of open platform events the app has subscribed to |
|     `status` | `int` | Version status **Optional values are**:  - `0`: Unknown status - `1`: Approved - `2`: Rejected - `3`: In review - `4`: Not submitted for review  |
|     `create_time` | `string` | Version creation time, in seconds |
|     `publish_time` | `string` | Version release time, in seconds |
|     `ability` | `app_ability` | App capabilities enabled in current version |
|       `gadget` | `gadget` | Gadget |
|         `enable_pc_mode` | `int` | Gadget mode supported by PC **Optional values are**:  - `1`: Sidebar mode - `2`: PC mode - `4`: Global navigation mode  |
|         `schema_urls` | `string[]` | Schema URL list |
|         `pc_use_mobile_pkg` | `boolean` | Whether a mobile package is used on PC |
|         `pc_version` | `string` | PC gadget version |
|         `mobile_version` | `string` | Mobile gadget version |
|         `mobile_min_lark_version` | `string` | Minimum compatible Lark version on mobile |
|         `pc_min_lark_version` | `string` | Minimum compatible Lark version on PC |
|       `web_app` | `web_app` | Web |
|         `pc_url` | `string` | URL on PC |
|         `mobile_url` | `string` | URL on mobile |
|       `bot` | `bot` | Bot |
|         `card_request_url` | `string` | Message card callback address |
|       `workplace_widgets` | `workplace_widget[]` | Block |
|         `min_lark_version` | `string` | Minimum compatible Lark version |
|       `navigate` | `navigate` | Global navigation gadget |
|         `pc` | `navigate_meta` | Global navigation information on PC |
|           `version` | `string` | Version number of the global navigation gadget |
|           `image_url` | `string` | URL of the default image |
|           `hover_image_url` | `string` | URL of the selected status image |
|         `mobile` | `navigate_meta` | Global navigation information on mobile |
|           `version` | `string` | Version number of the global navigation gadget |
|           `image_url` | `string` | URL of the default image |
|           `hover_image_url` | `string` | URL of the selected status image |
|       `cloud_doc` | `cloud_doc` | Docs app |
|         `space_url` | `string` | My Space redirect URL |
|         `i18n` | `cloud_doc_i18n_info[]` | Internationalized information |
|           `i18n_key` | `string` | Internationalization language key **Optional values are**:  - `zh_cn`: Chinese - `en_us`: English - `ja_jp`: Japanese  |
|           `name` | `string` | Internationalized Docs name |
|           `read_description` | `string` | Internationalized description of read scope for Docs |
|           `write_description` | `string` | Internationalized description of write scope for Docs |
|         `icon_url` | `string` | Icon link |
|         `mode` | `int` | Mode supported by Docs **Optional values are**:  - `0`: Unknown - `1`: Mobile  |
|       `docs_blocks` | `docs_block[]` | Docs block |
|         `block_type_id` | `string` | BlockTypeID |
|         `i18n` | `block_i18n_info[]` | Internationalized block information |
|           `i18n_key` | `string` | Internationalization language key **Optional values are**:  - `zh_cn`: Chinese - `en_us`: English - `ja_jp`: Japanese  |
|           `name` | `string` | Name |
|         `mobile_icon_url` | `string` | Mobile icon link |
|         `pc_icon_url` | `string` | PC port icon link |
|       `message_action` | `message_action` | Message shortcut |
|         `pc_app_link` | `string` | Link on PC |
|         `mobile_app_link` | `string` | Link on mobile |
|         `i18n` | `message_action_i18n_info[]` | Internationalized information |
|           `i18n_key` | `string` | Internationalization language key **Optional values are**:  - `zh_cn`: Chinese - `en_us`: English - `ja_jp`: Japanese  |
|           `name` | `string` | Internationalized name |
|       `plus_menu` | `plus_menu` | "+" menu |
|         `pc_app_link` | `string` | Link on PC |
|         `mobile_app_link` | `string` | Link on mobile |
|     `remark` | `app_version_remark` | Other configuration information of the app version |
|       `remark` | `string` | Notes |
|       `update_remark` | `string` | Update description |
|       `visibility` | `app_visibility` | Visibility suggestion edited by developers of the current app version. This field is empty if no suggestion was given. |
|         `is_all` | `boolean` | Whether the app is visible to all |
|         `visible_list` | `app_visible_list` | Visibility list |
|           `open_ids` | `string[]` | List of open_ids of members to whom the app is visible |
|           `department_ids` | `string[]` | List of IDs of departments to which the app is visible |
|           `group_ids` | `string[]` | List of IDs of groups to which the app is visible |
|         `invisible_list` | `app_visible_list` | Invisibility list |
|           `open_ids` | `string[]` | Invisibility list of open_ids of members to whom the app is visible |
|           `department_ids` | `string[]` | Invisibility list of IDs of departments to which the app is visible |
|           `group_ids` | `string[]` | Invisibility list of IDs of groups to which the app is visible |
|   `page_token` | `string` | Page identifier, when has_more is true, a new page_token will also be returned. Otherwise, page_token will not be returned |
|   `has_more` | `boolean` | Whether the response body has more parameters | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
        "items": [
            {
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
                        "level": 1
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
                                "ou_1ce003a0f23484806bcab85d2a55a000"
                            ],
                            "department_ids": [
                                "od-ddee42c0f8a948a5e650341e2153243b"
                            ],
                            "group_ids": [
                                "b6d1g5dd6fd26186"
                            ]
                        },
                        "invisible_list": {
                            "open_ids": [
                                "ou_1ce003a0f23484806bcab85d2a55a000"
                            ],
                            "department_ids": [
                                "od-ddee42c0f8a948a5e650341e2153243b"
                            ],
                            "group_ids": [
                                "b6d1g5dd6fd26186"
                            ]
                        }
                    }
                }
            }
        ],
        "page_token": "new-e3c5a0627cdf0c2e057da7257b90376a",
        "has_more": true
    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 210500 | page_token does not exist or has expired | Check whether the page_token is valid. A page_token expires in 2h. Obtain a new one if the current token expires. |
| 400 | 210501 | invalid page_token | page_token is not universal between apps. Check whether this page_token is obtained by the app calling the API. |
| 400 | 210502 | page_size out of range, should be between 1 and 50 | Check whether the page_size is in the range [1, 50]. |
| 400 | 210503 | invalid app_id | Check whether the app_id in the request path is valid. |
| 400 | 210504 | no such app in tenant | Check whether the queried app and the app used to call the API are in the same company. |
| 400 | 210505 | target app not a custom app | Check whether the queried app is a custom app. |
| 400 | 210506 | no such app | Check whether the app_id in the request path exists. |
| 400 | 210508 | insufficient permission level | Check the scope you have requested and the queried app_id. This error code is returned when the queried app_id does not belong to this app and you have not requested the Obtain app version information scope. |
| 400 | 210514 | invalid order | Check whether the order is in the range [0, 1]. |
