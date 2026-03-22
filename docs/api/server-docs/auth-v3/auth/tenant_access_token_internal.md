---
title: "Get custom app tenant_access_token"
url: "https://open.larksuite.com/document/ukTMukTMukTM/ukDNz4SO0MjL5QzM/auth-v3/auth/tenant_access_token_internal"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/auth/v3/tenant_access_token/internal"
service: "auth-v3"
resource: "auth"
section: "Authenticate and Authorize"
updated: "1699856742000"
---

# Custom apps get tenant_access_token

Custom apps can obtain `tenant_access_token` through this interface.

> **Note:** The maximum validity period of `tenant_a ccess_token` is 2 hours. If called when the validity is less than 30 minutes, this interface will return a new `tenant_a ccess_token`, resulting in two valid tokens existing simultaneously.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/auth/v3/tenant_access_token/internal |
| HTTP Method | POST |
| Supported app types | custom |
| Required scopes | None | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `app_id` | `string` | Yes | The unique identifier of the application is obtained after the application is created. For a detailed introduction to `app_id`, please refer to the introduction of General parameters. **Example value:** "cli_slkdjalasdkjasd" |
| `app_secret` | `string` | Yes | Application key, obtained after creating the application. For a detailed introduction to `app_secret`, please refer to the introduction of General parameters. **Example value:** "dskLLdkasdjlasdKK" | ### Request body example

```json
{
    "app_id": "cli_slkdjalasdkjasd",
    "app_secret": "dskLLdkasdjlasdKK"
}
```

## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error code, non-0 indicates failure. |
| `msg` | `string` | Error description. |
| `tenant_access_token` | `string` | The access token. |
| `expire` | `int` | The expiration time of tenant_access_token, in seconds. | ### Response body example

```json
{
    "code": 0,
    "msg": "ok",
    "tenant_access_token": "t-caecc734c2e3328a62489fe0648c4b98779515d3",
    "expire": 7200
}
```

### Error code

For a detailed description of error codes, see Introduction to Common error codes.
