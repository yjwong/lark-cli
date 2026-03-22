---
title: "Re-pushing app_ticket"
url: "https://open.larksuite.com/document/ukTMukTMukTM/ukDNz4SO0MjL5QzM/auth-v3/auth/app_ticket_resend"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/auth/v3/app_ticket/resend"
service: "auth-v3"
resource: "auth"
section: "Authenticate and Authorize"
updated: "1661829813000"
---

# Re-push app_ticket

Lark pushes the latest app_ticket to the app per hour. The app can also trigger Lark to re-push timely by calling this API. (This API can only trigger event push instead of directly obtaining app_ticket.)

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/auth/v3/app_ticket/resend |
| HTTP Method | POST |
| Supported app types | isv |
| Required scopes | None | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `app_id` | `string` | Yes | The app UID, which is obtained when an app is created **Example value**: "cli_slkdjalasdkjasd" |
| `app_secret` | `string` | Yes | The app's secret key, which is obtained when an app is created **Example value**: "dskLLdkasdjlasdKK" | ### Request body example

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
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions | ### Response body example

```json
{
    "code": 0,
    "msg": "success"
}
```

### Error code

For details, refer to Server-side error codes.
