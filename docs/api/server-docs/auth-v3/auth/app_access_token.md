---
title: "Store applications get app_access_token"
url: "https://open.larksuite.com/document/ukTMukTMukTM/ukDNz4SO0MjL5QzM/auth-v3/auth/app_access_token"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/auth/v3/app_access_token"
service: "auth-v3"
resource: "auth"
section: "Authenticate and Authorize"
updated: "1699856742000"
---

# Obtain app_access_token (for store apps)

A store app calls this API to obtain app_access_token, which is used as the auth token to obtain app resources.

> The token is valid for 2 hours, and remains unchanged when the API is called during this period. When the validity period of the token is less than 30 minutes, a new request leads to the creation of a new token, and the old token is still valid.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/auth/v3/app_access_token |
| HTTP Method | POST |
| Supported app types | isv |
| Required scopes | None | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `app_id` | `string` | Yes | The app UID, which is obtained when an app is created **Example value**: "cli_slkdjalasdkjasd" |
| `app_secret` | `string` | Yes | The app's secret key, which is obtained when an app is created **Example value**: "dskLLdkasdjlasdKK" |
| `app_ticket` | `string` | Yes | A temporary token that is pushed to an app by the Platform at the scheduled time, and obtained through the event listener. For details, see Subscription events. **Example value**: "dskLLdkasd" | ### Request body example

```json
{
    "app_id": "cli_slkdjalasdkjasd",
    "app_secret": "dskLLdkasdjlasdKK",
    "app_ticket": "dskLLdkasd"
}
```

## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `app_access_token` | `string` | The access token |
| `expire` | `int` | The expiration time of app_access_token, in seconds | ### Response body example

```json
{
    "code": 0,
    "msg": "success",
    "app_access_token": "a-6U1SbDiM6XIH2DcTCPyeub",
    "expire": 7140
}
```

### Error code

For details, refer to Server-side error codes.
