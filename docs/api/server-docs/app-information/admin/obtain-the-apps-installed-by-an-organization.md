---
title: "Obtain the Apps Installed by an Organization"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uYDN3UjL2QzN14iN0cTN"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/application/v3/app/list"
section: "App Information"
scopes:
  - "admin:app.info:readonly"
updated: "1646720015000"
---

# Obtain apps installed by a company

This API is used to query the list of apps installed by a company. This API is only available to custom apps.

## Request
| HTTP URL | https://open.larksuite.com/open-apis/application/v3/app/list |
| --- | --- |
| HTTP Method | GET |
| Supported app types | custom |
| Required scopes | `admin:app.info:readonly` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" Learn more about how to obtain and use access_token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Request parameters
|Parameter|Type|Required|Description|Source|
|--|-----|--|----|----|
|page_token|string|No|Indicates the paging start position. The paging starts from the beginning if this parameter is not specified.|Request URL|
|page_size|int|No|The maximum number of results requested per page (up to 100). The value of 0 indicates 100 results.|Request URL|
|lang|string|No|Preferred language used to display app information (zh_cn: Chinese, en_us: English, ja_jp: Japanese)|Request URL|
|status|int|No|Status of apps to be returned. 0: Disabled; 1: Enabled; -1: All. The default value is -1.|Request URL| ## Response
### Response body
|Parameter|Description|
|--|--|
|code|Return code. A value other than  0  indicates failure.|
|msg|Return code description|
|data|Returned business information. This is only valid when code is 0.|
|  ∟page_token|Start position for the next request page|
|  ∟page_size|Actual page size returned for this request|
|  ∟total_count|Total number of apps available|
|  ∟has_more|Whether there are more apps|
|  ∟lang|Version language currently selected|
|  ∟app_list|App list|
|    ∟app_id|App ID|
|    ∟primary_language|Preferred app language|
|    ∟app_name|App name|
|    ∟description|App description|
|    ∟avatar_url|App icon|
|    ∟app_scene_type|App type. 0: Custom app; 1: Store app|
|    ∟status|App status. 0: Disabled; 1: Enabled|
|    ∟mobile_default_ability|Default app feature on mobile. 0: Not enabled; 1: Gadget; 2: H5; 8: Bot|
|    ∟pc_default_ability|Default app feature on PC. 0: Not enabled; 1: Gadget; 2: H5; 8: Bot| ###  Response example
```json
{
    "code": 0,
    "data": {
        "app_list": [
            {
                "app_id": "cli_9cb844403dbb9108",
                "app_name": "wenjuan.com",
                "app_scene_type": 1,
                "avatar_url": "https://s1-imfile.feishucdn.com/static-resource/v1/d5ca5971-437b-4b1d-b295-679268f9a2cg",
                "description": "Professional and easy-to-use platform for questionnaire survey,   registration forms, and   exam assessment",
                "mobile_default_ability": 1,
                "pc_default_ability": 1,
                "primary_language": "zh_cn",
                "status": 1
            },
            {
                "app_id": "cli_9cb844403dbb9108",
                "app_name": "Approval",
                "app_scene_type": 1,
                "avatar_url": "https://s3-imfile.feishucdn.com/static-resource/v1/e37af67a-b012-4ee0-80ea-a4d28c94b4eg",
                "description": "Simple, efficient, and open approval tool",
                "mobile_default_ability": 1,
                "pc_default_ability": 1,
                "primary_language": "zh_cn",
                "status": 1
            },
            {
                "app_id": "cli_9cb844403dbb9108",
                "app_name": "Payroll",
                "app_scene_type": 1,
                "avatar_url": "https://s3-imfile.feishucdn.com/static-resource/v1/da710014fc4c975ce66b~?image_size=noop&cut_type=&quality=_q100&format=image&sticker_format=.webp",
                "description": "Convenient and secure payroll management, and instant payroll publication",
                "mobile_default_ability": 1,
                "pc_default_ability": 1,
                "primary_language": "zh_cn",
                "status": 1
            }
        ],
        "has_more": 1,
        "lang": "zh_cn",
        "page_size": 3,
        "page_token": "3",
        "total_count": 34
    },
    "msg": "success"
}
```
