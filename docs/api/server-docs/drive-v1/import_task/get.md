---
title: "Query import task result"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/import_task/get"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/drive/v1/import_tasks/:ticket"
service: "drive-v1"
resource: "import_task"
section: "Docs"
rate_limit: "1000 per minute、50 per second"
scopes:
  - "drive:drive"
  - "drive:drive:readonly"
updated: "1698998306000"
---

# Query the import task result

Polling the import results based on the `ticket` returned from the creation of the import task. Detailed operation steps can be found in the Import Usage Guide.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/drive/v1/import_tasks/:ticket |
| HTTP Method | GET |
| Rate Limit | 1000 per minute、50 per second |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `drive:drive` `drive:drive:readonly` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `ticket` | `string` | ID of the import task **Example value**: "6990281865xxxxxxxx7843" | ## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `result` | `import_task` | The result of an import task |
|     `ticket` | `string` | Task ID |
|     `type` | `string` | Format of the destination document to import in Docs |
|     `job_status` | `int` | Task status **Optional values are**:  - `0`: Success - `1`: Initializing - `2`: Processing - `3`: Internal error. - `100`: The document to import is encrypted. - `101`: Internal error. - `102`: Internal error. - `103`: Internal error. - `104`: The available capacity for the tenant is insufficient. - `105`: The number of folders has reached the limit. - `106`: Internal error. - `108`: Processing timed out. - `109`: Internal error. - `110`: No permission scope. - `112`: The format is not supported. - `113`: The office format is not supported. - `114`: Internal error. - `115`: The file to import has reached the maximum size. - `116`: You have no access to the directory. - `117`: The directory has been deleted. - `118`: The extension of the file to import is inconsistent with that specified for the task. - `119`: The directory doesn't exist. - `120`: The type of the file to import is inconsistent with that specified for the task. - `121`: The file to import has expired. - `122`: Export is not allowed when a copy is being made. - `5000`: Internal error. - `7000`: The number of blocks in the DOCX file has reached the system limit. - `7001`: The level of blocks in the DOCX file has reached the system limit. - `7002`: The size of blocks in the DOCX file has reached the system limit.  |
|     `job_error_msg` | `string` | Cause of task failure |
|     `token` | `string` | Token of the document to import in Docs |
|     `url` | `string` | URL of the document to import in Docs |
|     `extra` | `string[]` | Notification message upon a successful task | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
        "result": {
            "ticket": "6990281865xxxxxxxx7843",
            "type": "sheet",
            "job_status": 0,
            "job_error_msg": "success",
            "token": "shtcnVBTG6SuxxxxxxxkM2tUX",
            "url": "https://example.larksuite.com/sheets/shtcnVBTG6SuxxxxxxxkM2tUX",
            "extra": [
                "2000"
            ]
        }
    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 500 | 1069901 | internal error | Service internal error, consult [Oncall] (/ssl: ttdoc/ukTMukTMukTM/uUDN04SN0QjL1QDN/docs-overview #51f94b41) |
| 403 | 1069902 | no permission | No read or export permission scope. |
| 500 | 1069903 | internal error | Internal service error. For more information, contact support. |
| 400 | 1069904 | invalid param | Invalid parameter, export csv whether passed in sub_id |
| 500 | 1069905 | internal error | Internal service error. For more information, contact support. |
| 403 | 1069906 | docs deleted | The document has been deleted. |
| 404 | 1069907 | file token not found | The file token doesn't exist. |
| 403 | 1069908 | mount point not found or no permission | The mount point doesn't exist or you have no access to the mount point. |
| 400 | 1069909 | import file size over limit | The file exceeds 20 MB. |
| 400 | 1069910 | import file extension not match | The extension of the uploaded file is inconsistent with that of the file in the import task. |
| 400 | 1069911 | import file type not match | The type of the uploaded file is inconsistent with that of the file specified for the import task. |
| 400 | 1069912 | folder not exist | The directory doesn't exist. |
| 400 | 1069913 | import file token expired | The uploaded file has expired. An uploaded file is valid for 5 minutes. |
| 400 | 1069914 | export file token invalid | Export document token is invalid |
