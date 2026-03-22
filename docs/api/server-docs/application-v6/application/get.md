---
title: "Get Application Information"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/application-v6/application/get"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/application/v6/applications/:app_id"
service: "application-v6"
resource: "application"
section: "App Information"
scopes:
  - "application:application:self_manage"
  - "admin:app.info:readonly"
field_scopes:
  - "contact:user.employee_id:readonly"
updated: "1646274611000"
---

# Obtain apps information

Obtains basic information of an app based on the app_id.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/application/v6/applications/:app_id |
| HTTP Method | GET |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `application:application:self_manage` `admin:app.info:readonly` |
| Required field scopes | > The response body of the API contains the following sensitive fields, and they will be returned only after corresponding scopes are added. If you do not need the fields, it is not recommended that you request the scopes. `contact:user.employee_id:readonly` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" Learn more about obtaining and using access_token. | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `app_id` | `string` | App's app_id. To query the information of other apps, you must request the Obtain apps information scope. To only query the information of this app, enter "me" or the app_id. **Example value**: "cli_9b445f5258795107" | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `lang` | `string` | Yes | Specifies the language in which the app information is obtained **Example value**: "zh_cn" **Optional values are**: - `zh_cn`: Chinese - `en_us`: English - `ja_jp`: Japanese **Data validation rules**: - Minimum length: `1` characters |
| `user_id_type` | `string` | No | User ID categories **Example value**: "open_id" **Optional values are**: - `open_id`: user's open id - `union_id`: user's union id - `user_id`: user's user id **Default value**: `open_id` **When the value is `user_id`, the following field scopes are required**: `contact:user.employee_id:readonly` | ## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
| ∟ `app` | `application` | App data |
| ∟ `app_id` | `string` | App's app_id |
| ∟ `creator_id` | `string` | App creator (owner) |
| ∟ `status` | `int` | App status **Optional values are**: - `0`: Disabled status - `1`: Enabled status - `2`: Not-enabled status - `3`: Unknown status |
| ∟ `scene_type` | `int` | App type **Optional values are**: - `0`: Custom app - `1`: Store app - `2`: Personal store app - `3`: Unknown app type |
| ∟ `redirect_urls` | `string[]` | Redirect URLs in Security Settings |
| ∟ `online_version_id` | `string` | ID of the app version released online. This field can be left empty if there is no such ID. |
| ∟ `unaudit_version_id` | `string` | Version ID in review. This field can be left empty if there is no such version ID. |
| ∟ `app_name` | `string` | App name |
| ∟ `avatar_url` | `string` | App icon URL |
| ∟ `description` | `string` | Default app description |
| ∟ `scopes` | `app_scope[]` | App scope list |
| ∟ `scope` | `string` | App scopes |
| ∟ `description` | `string` | Internationalized description of app scopes |
| ∟ `level` | `string` | Scope level description **Optional values are**: - `1`: General scope - `2`: Advanced scope - `3`: Highly sensitive scope - `0`: Unknown level |
| ∟ `back_home_url` | `string` | Admin homepage address |
| ∟ `i18n` | `app_i18n_info[]` | Internationalized information list of the app |
| ∟ `i18n_key` | `string` | Internationalization language key **Optional values are**: - `zh_cn`: Chinese - `en_us`: English - `ja_jp`: Japanese |
| ∟ `name` | `string` | Internationalized app name |
| ∟ `description` | `string` | Internationalized app description (subtitle) |
| ∟ `help_use` | `string` | Link to the internationalized Help documentation |
| ∟ `primary_language` | `string` | Primary app language **Optional values are**: - `zh_cn`: Chinese - `en_us`: English - `ja_jp`: Japanese |
| ∟ `common_categories` | `string[]` | Internationalized description of categories | ### Response body example

```json
{
    "code": 0,
    "msg": "success",
    "data": {
        "app": {
            "app_id": "cli_9b445f5258795107",
            "creator_id": "ou_d317f090b7258ad0372aa53963cda70d",
            "status": 1,
            "scene_type": 0,
            "redirect_urls": [
                "https://www.example.com"
            ],
            "online_version_id": "oav_d317f090b7258ad0372aa53963cda70d",
            "unaudit_version_id": "oav_d317f090b7258ad0372aa53963cda70d",
            "app_name": "App name",
            "avatar_url": "https://sf1-ttcdn-tos.pstatp.com/img/avatar/d279000ca4d3f7f6aaff~72x72.jpg",
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
            "primary_language": "zh_cn",
            "common_categories": [
                "Analysis tools"
            ]
        }
    }
}
```

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 210503 | invalid app_id | Check whether the app_id in the request path is valid. |
| 400 | 210504 | no such app in tenant | Check whether the queried app and the app used to call the API are in the same company. |
| 400 | 210505 | target app not a custom app | Check whether the queried app is a custom app. |
| 400 | 210506 | no such app | Check whether the app_id in the request path exists. |
| 400 | 210508 | insufficient permission level | Check the scope you have requested and the queried app_id. This error code is returned when the queried app_id does not belong to this app and you have not requested the Obtain app version information scope. |
