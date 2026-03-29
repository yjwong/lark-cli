---
title: "get export task result"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/export_task/get"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/drive/v1/export_tasks/:ticket"
service: "drive-v1"
resource: "export_task"
section: "Docs"
rate_limit: "100 per minute"
scopes:
  - "drive:export:readonly"
updated: "1684896008000"
---

# Query export task results

According to the result of the ticket polling export task returned by the Create an export task, after obtaining the file token of the export product through this API, the Download export file API can be invoked to download the export product to the local.

::: note 
The user who gets the export result needs to be the same as the user who created the export task.
 :::

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/drive/v1/export_tasks/:ticket |
| HTTP Method | GET |
| Rate Limit | 100 per minute |
| Supported app types | custom |
| Required scopes | `drive:export:readonly` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `ticket` | `string` | Export task ID, ticket field in the [Create export task] (/ssl :: ttdoc//uAjLw4CM/ukTMukTMukTM/reference/drive-v1/export_task/create) response **Example value**: "6933093124755423251" | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `token` | `string` | Yes | Token to export document [How to get the documentation token] (/ssl: ttdoc/ukTMukTMukTM/uczNzUjL3czM14yN3MTN #08bb5df6) **Example value**: "doccnZVxxxxxxxxxxxxGiyBgYqe" | ## Response

###Request example 
 "curl 
 Curl --location --request GET https://open.larksuite.com/open-apis/drive/v1/export_tasks/7143131813848809492?token=docbcZVGtv1papC6jAVGiyBgYqe '\ 
 --Header'Authorization: Bearer t-g1029efgIY34MWDJL4CEYQOVN5TZF2OMPJXTDVOP' 
 "'"

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `result` | `export_task` | Export task results |
|     `file_extension` | `string` | Export file extension **Optional values are**:  - `docx`: Microsoft Word (DOCX) format - `pdf`: PDF format - `xlsx`: Microsoft Excel (XLSX) format - `csv`: CSV format  |
|     `type` | `string` | Export document type [document type description] (/ssl :: ttdoc/ukTMukTMukTM/uczNzUjL3czM14yN3MTN #560bf735) **Optional values are**:  - `doc`: Legacy Lark cloud document type, supports export to docx, pdf format - `sheet`: Lark spreadsheet type, support export to xlsx, csv format - `bitable`: Bitable type, support export to xlsx, csv format - `docx`: The new version of Lark cloud document type supports exporting to docx, pdf formats  |
|     `file_name` | `string` | Export filename |
|     `file_token` | `string` | Export file drive token |
|     `file_size` | `int` | Export file size in units of bytes |
|     `job_error_msg` | `string` | Reason for task failure |
|     `job_status` | `int` | Task status **Optional values are**:  - `0`: Success - `1`: Initialization - `2`: Processing - `3`: Internal error - `107`: Export document is too large - `108`: Processing timeout - `109`: Export content block without permission - `110`: No permission - `111`: Export document deleted - `122`: Export is prohibited in creating a copy - `123`: Export document does not exist - `6000`: Exporting document images too many  | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
        "result": {
            "file_extension": "pdf",
            "type": "doc",
            "file_name": "docName",
            "file_token": "boxcnxe5OxxxxxxxSNdsJviENsk",
            "file_size": 34356,
            "job_error_msg": "success",
            "job_status": 0
        }
    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 500 | 1069901 | internal error | Service internal error, consult [Oncall] (/ssl: ttdoc/ukTMukTMukTM/uUDN04SN0QjL1QDN/docs-overview #51f94b41) |
| 403 | 1069902 | no permission | No read or export permission scope. |
| 400 | 1069904 | invalid param | Invalid parameter, export csv whether passed in sub_id |
| 410 | 1069906 | docs deleted | The document has been deleted. | Other code, refer to Server-side error codes.
