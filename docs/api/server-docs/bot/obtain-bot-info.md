---
title: "Obtain Bot Info"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uAjMxEjLwITMx4CMyETM"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/bot/v3/info"
section: "Bot"
updated: "1646918676000"
---

# Access Bot information

Access the basic information about the Bot. 

> The Bot capability needs to be enabled

## Request
| HTTP URL | https://open.larksuite.com/open-apis/bot/v3/info |
| --- | --- |
| HTTP Method | GET |
| Permission scope requirement | None | ### Request header
| Name | Type | Required | Description |
| --- | --- | --- | --- | ## Response

### Response body
| Name | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error code. A value other than 0 indicates failure |
| `msg` | `string` | Error description |
| `bot` | `bot_info` | Bot information |
| ∟ `activate_status` | `int` | Current status of the app. 0: Initialized, to be installed by the tenant 1: Disabled 2: Enabled 3: To be enabled after installation 4: Upgraded, to be enabled 5: Disabled due to license expiration 6: Lark plan expires or is downgraded and disabled |
| ∟ `app_name` | `string` | app name |
| ∟ `avatar_url` | `string` | URL of app image |
| ∟ `ip_white_list` | `string[]` | IP allowlist address of app |
| ∟ `open_id` | `string` | open_id of Bot | ### Example response body

```json
{
    "code":0,
    "msg":"ok",
    "bot":{
        "activate_status":2,
        "app_name":"name",
        "avatar_url":"https://xxxxxxxxxxxxxx",
        "ip_white_list":[

        ],
        "open_id":"ou_e6e14f667cfe239d7b129b521dce0569"
    }
}
```
