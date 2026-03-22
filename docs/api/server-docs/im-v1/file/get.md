---
title: "Download file"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/file/get"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/im/v1/files/:file_key"
service: "im-v1"
resource: "file"
section: "Messaging"
rate_limit: "1000 per minute、50 per second"
updated: "1717574914000"
---

# Download file

Download files. Only files uploaded by an app itself can be downloaded.

> Notes:
> - Bot ability must be enabled.
> - A bot can only download files uploaded by the bot.
> - To download resources sent by users, please use Obtain the Resource in the Message.
> - The size of the downloaded resource cannot exceed 100M.
> - If you need the Content-Disposition header, you need to set the Content-Type in the header to application/json when initiating the request.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/im/v1/files/:file_key |
| HTTP Method | GET |
| Rate Limit | 1000 per minute、50 per second |
| Supported app types | custom,isv |
| Required scopes | None | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `file_key` | `string` | File key, obtained by uploading a picture through the Upload files API. **Example value**: "file_456a92d6-c6ea-4de4-ac3f-7afcf44ac78g" | ## Response

When the HTTP status code is 200, it means success

Return file binary stream

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 234001 | Invalid request param. | Check that the request parameters are correct. |
| 401 | 234002 | Unauthorized. | Authentication failed, contact Oncall to solve. |
| 400 | 234008 | The app is not the resource sender | The app is not the owner of the resource. |
| 400 | 234005 | Image has been deleted | The resource doesn't exist. |
| 400 | 234037 | Downloaded file size exceeds limit. | Downloaded file size exceeds 100MB limit forbid. |
| 400 | 234041 | Tenant master key has been deleted, please contact the tenant administrator. | Tenant master key has been deleted, please contact the tenant administrator. |
| 400 | 234042 | Hybrid deployment tenant storage error, such as full storage space, please contact tenant administrator. | Hybrid deployment tenant storage error, such as full storage space, please contact tenant administrator. |
