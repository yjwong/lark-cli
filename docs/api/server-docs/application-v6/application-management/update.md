---
title: "Enable or disable application"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/application-v6/application-management/update"
method: "PUT"
api_path: "https://open.larksuite.com/open-apis/application/v6/applications/:app_id/management"
service: "application-v6"
resource: "application-management"
section: "App Information"
rate_limit: "10 per minute"
scopes:
  - "admin:app.enable:write"
updated: "1706237733000"
---

# Enable or disable application

You can automatically deactivate or enable self-built apps and store apps that have been installed in the enterprise.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/application/v6/applications/:app_id/management |
| HTTP Method | PUT |
| Rate Limit | 10 per minute |
| Supported app types | custom |
| Required scopes | `admin:app.enable:write` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `app_id` | `string` | App ID **Example value**: "cli_a4517c8461f8100a" | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `enable` | `boolean` | No | Enable/Deactivate Apps **Example value**: true | ### Request body example

{
    "enable": true
}

## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {}
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 210001 | param is invalid | Parameter error |
