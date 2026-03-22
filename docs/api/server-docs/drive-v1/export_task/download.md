---
title: "download export file"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/export_task/download"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/drive/v1/export_tasks/file/:file_token/download"
service: "drive-v1"
resource: "export_task"
section: "Docs"
rate_limit: "100 per minute"
scopes:
  - "drive:export:readonly"
updated: "1684896008000"
---

# Download export file

According to the export product token returned by the Query export task results, download the export product file to the local.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/drive/v1/export_tasks/file/:file_token/download |
| HTTP Method | GET |
| Rate Limit | 100 per minute |
| Supported app types | custom |
| Required scopes | `drive:export:readonly` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `file_token` | `string` | Export document token **Example value**: "boxcnNAlfwHxxxxxxxxxxSaLSec" | ## Response

When the HTTP status code is 200, it means success

Return file binary stream

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 1060001 | param is invalis | Parameter error. |
