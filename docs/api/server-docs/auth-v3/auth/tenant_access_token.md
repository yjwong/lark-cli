---
title: "Store applications get tenant_access_token"
url: "https://open.larksuite.com/document/ukTMukTMukTM/ukDNz4SO0MjL5QzM/auth-v3/auth/tenant_access_token"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/auth/v3/tenant_access_token"
service: "auth-v3"
resource: "auth"
section: "Authenticate and Authorize"
updated: "1699856742000"
---

# Obtain tenant_access_token (for store apps)

A store app calls this API to obtain tenant_access_token, which is used as the auth token to obtain company resources.

> The token is valid for 2 hours, and remains unchanged when the API is called during this period. When the validity period of the token is less than 30 minutes, a new request leads to the creation of a new token, and the old token is still valid.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/auth/v3/tenant_access_token |
| HTTP Method | POST |
| Supported app types | isv |
| Required scopes | None | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `app_access_token` | `string` | Yes | The auth token of an app **Example value**: "a-32bd8551db2f081cbfd26293f27516390b9feb04" |
| `tenant_key` | `string` | Yes | The company ID, which can be obtained in two ways: 1. Pushed by the Platform to the app launched by a company. For details, see Initial app enabling. 2. Returned when a user logs in. For details, see Obtain user ID access token. **Example value**: "73658811060f175d" | ### Request body example

```json
{
    "app_access_token": "a-32bd8551db2f081cbfd26293f27516390b9feb04",
    "tenant_key": "73658811060f175d"
}
```

## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `tenant_access_token` | `string` | The access token |
| `expire` | `int` | The expiration time of tenant_access_token, in seconds | ### Response body example

```json
{
    "code": 0,
    "msg": "success",
    "tenant_access_token": "t-caecc734c2e3328a62489fe0648c4b98779515d3",
    "expire": 7140
}
```

### Error code

For details, refer to Server-side error codes.
