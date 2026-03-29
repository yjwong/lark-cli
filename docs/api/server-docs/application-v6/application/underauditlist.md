---
title: "Get Application Release Request List"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/application-v6/application/underauditlist"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/application/v6/applications/underauditlist"
service: "application-v6"
resource: "application"
section: "App Information"
scopes:
  - "admin:app.info:readonly"
field_scopes:
  - "contact:user.employee_id:readonly"
updated: "1646274609000"
---

# View the list of app release applications

Displays the list of all custom apps pending review for this company.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/application/v6/applications/underauditlist |
| HTTP Method | GET |
| Supported app types | custom |
| Required scopes | `admin:app.info:readonly` |
| Required field scopes | > The response body of the API contains the following sensitive fields, and they will be returned only after corresponding scopes are added. If you do not need the fields, it is not recommended that you request the scopes. `contact:user.employee_id:readonly` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" Learn more about obtaining and using access_token. | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `lang` | `string` | Yes | Specifies the language of returned results **Example value**: "zh_cn" **Optional values are**: - `zh_cn`: Chinese - `en_us`: English - `ja_jp`: Japanese **Data validation rules**: - Minimum length: `1` characters |
| `page_token` | `string` | No | Page identifier. It is not filled in the first request, indicating traversal from the beginning; when there will be more groups, the new page_token will be returned at the same time, and the next traversal can use the page_token to get more groups **Example value**: "new-e3c5a0627cdf0c2e057da7257b90376a" |
| `page_size` | `int` | No | Page size **Example value**: 10 **Data validation rules**: - Maximum value: `50` |
| `user_id_type` | `string` | No | User ID categories **Example value**: "open_id" **Optional values are**: - `open_id`: user's open id - `union_id`: user's union id - `user_id`: user's user id **Default value**: `open_id` **When the value is `user_id`, the following field scopes are required**: `contact:user.employee_id:readonly` | ## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
| ∟ `items` | `application[]` | List of apps pending review |
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
| ∟ `common_categories` | `string[]` | Internationalized description of categories |
| ∟ `has_more` | `boolean` | Whether the response body has more parameters |
| ∟ `page_token` | `string` | Page identifier, when has_more is true, a new page_token will also be returned. Otherwise, page_token will not be returned | ### Response body example

```json
{
    "code": 0,
    "msg": "success",
    "data": {
        "items": [
            {
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
        ],
        "has_more": true,
        "page_token": "new-xxxxxxxxxxx"
    }
}
```

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 210500 | page_token does not exist or has expired | Check whether the page_token is valid. A page_token expires in 2h. Obtain a new one if the current token expires. |
| 400 | 210501 | invalid page_token | page_token is not universal between apps. Check whether this page_token is obtained by the app calling the API. |
| 400 | 210502 | page_size out of range, should be between 1 and 50 | Check whether the page_size is in the range [1, 50]. |
