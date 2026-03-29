---
title: "Download image"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/image/get"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/im/v1/images/:image_key"
service: "im-v1"
resource: "image"
section: "Messaging"
rate_limit: "1000 per minute、50 per second"
updated: "1717574914000"
---

# Download image

Download image resources. Only images of the message type uploaded by the current application can be downloaded.

> Notes:
> - Bot ability must be enabled.
> - A bot can only download "message" images uploaded by the bot. "avatar" images are not supported.
> - To download resources sent by users, please use Obtain the Resource in the Message.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/im/v1/images/:image_key |
| HTTP Method | GET |
| Rate Limit | 1000 per minute、50 per second |
| Supported app types | custom,isv |
| Required scopes | None | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `image_key` | `string` | Image key, obtained by uploading a picture through the Upload images API. **Example value**: "img_8d5181ca-0aed-40f0-b0d1-b1452132afbg" | ### Request example
```
curl --location --request GET 'https://open.larksuite.com/open-apis/im/v1/images/img_v2_b99741-7628-4abd-aad0-b881e4db83ig' \
--header 'Authorization: Bearer t-39eec8560faa3dded7971873eb649fd40e70e0f1'
```

## Response

When the HTTP status code is 200, it means success

Return file binary stream

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 234001 | Invalid request param. | Check that the request parameters are correct. |
| 401 | 234002 | Unauthorized. | Authentication failed, contact Oncall to solve. |
| 400 | 234005 | Image has been deleted | The resource doesn't exist. |
| 400 | 234007 | App does not enable bot feature. | The bot ability is not enabled for the app. |
| 400 | 234008 | The app is not the resource sender | The app is not the owner of the resource. |
| 400 | 234041 | Tenant master key has been deleted, please contact the tenant administrator. | Tenant master key has been deleted, please contact the tenant administrator. |
| 400 | 234042 | Hybrid deployment tenant storage error, such as full storage space, please contact tenant administrator. | Hybrid deployment tenant storage error, such as full storage space, please contact tenant administrator. |
