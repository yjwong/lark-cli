---
title: "Create an export task"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/export_task/create"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/drive/v1/export_tasks"
service: "drive-v1"
resource: "export_task"
section: "Docs"
rate_limit: "100 per minute"
scopes:
  - "drive:export:readonly"
updated: "1686566831000"
---

# Create an export task

Create an export task and export the Docs to a local file in the specified format. Currently, the Upgraded Docs, Sheets, Bitable and Docs are supported. This API is asynchronous. The task is created and returned immediately. It will not block and wait until the task is executed successfully. Therefore, it is necessary to combine the Query Export Task Results to obtain the export results.

::: note 
 The requesting user needs to have export permission on this Docs. 
 :::

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/drive/v1/export_tasks |
| HTTP Method | POST |
| Rate Limit | 100 per minute |
| Supported app types | custom |
| Required scopes | `drive:export:readonly` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `file_extension` | `string` | Yes | Export file extension **Example value**: "csv" **Optional values are**:  - `docx`: Microsoft Word (DOCX) format - `pdf`: PDF format - `xlsx`: Microsoft Excel (XLSX) format - `csv`: CSV format  |
| `token` | `string` | Yes | Export document token [get document token] (/ssl: ttdoc/ukTMukTMukTM/uczNzUjL3czM14yN3MTN #08bb5df6) **Example value**: "shtcnxe5OxxxxxxxSNdsJviENsk" |
| `type` | `string` | Yes | Export document type [document type description] (/ssl :: ttdoc/ukTMukTMukTM/uczNzUjL3czM14yN3MTN #560bf735) **Example value**: "sheet" **Optional values are**:  - `doc`: Legacy Lark cloud document type, supports export to docx, pdf format - `sheet`: Lark spreadsheet type, support export to xlsx, csv format - `bitable`: Bitable type, support export to xlsx, csv format - `docx`: The new version of Lark cloud document type supports exporting to docx, pdf formats  |
| `sub_id` | `string` | No | Export subtable ID, only used when exporting spreadsheet/Bitable to csv **Example value**: "tblKz5D60T4JlfcT" | ### Request body example

{
    "file_extension": "csv",
    "token": "shtcnxe5OxxxxxxxSNdsJviENsk",
    "type": "sheet",
    "sub_id": "tblKz5D60T4JlfcT"
}

## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `ticket` | `string` | Export task ID | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
        "ticket": "6933093124755423251"
    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 500 | 1069901 | internal error | Service internal error, consult [Oncall] (/ssl: ttdoc/ukTMukTMukTM/uUDN04SN0QjL1QDN/docs-overview #51f94b41) |
| 403 | 1069902 | no permission | No read or export permission scope. |
| 400 | 1069904 | invalid param | Invalid parameter, export csv whether passed in sub_id |
| 404 | 1069906 | docs deleted | The document has been deleted. |
| 404 | 1069914 | export file token invalid | Export document token is invalid |
| 400 | 1069918 | export file extension mismatch | File extension and type mismatch |
| 429 | 1069923 | too many requests | Request frequency limit, please reduce the request frequency and try again later. | Other code, refer to Server-side error codes.
