---
title: "whiteboard image"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/board-v1/whiteboard/download_as_image"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/board/v1/whiteboards/:whiteboard_id/download_as_image"
service: "board-v1"
resource: "whiteboard"
section: "Docs"
rate_limit: "10 per second"
scopes:
  - "board:whiteboard:node:read"
updated: "1744960864000"
---

# Download Whiteboard As Image

Download whiteboard as an image, Content-Type=image/png

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/board/v1/whiteboards/:whiteboard_id/download_as_image |
| HTTP Method | GET |
| Rate Limit | 10 per second |
| Supported app types | custom,isv |
| Required scopes | `board:whiteboard:node:read` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `whiteboard_id` | `string` | **Example value**: "Ru8nwrWFOhEmaFbEU2VbPRsHcxb" | ## Response

When the HTTP status code is 200, it means success

Return file binary stream

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 2890001 | invalid format | check the format |
| 400 | 2890002 | invalid arg | argument is invalid |
| 400 | 2890003 | record missing | can not find the record |
| 401 | 2890004 | auth failed | auth failed |
| 403 | 2890005 | forbidden | check if user has permission of whiteboard |
| 429 | 2890006 | too many request | request over qps limit |
| 500 | 2891001 | server internal error | server error happened |
