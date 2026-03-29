---
title: "Create Import Task"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/import_task/create"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/drive/v1/import_tasks"
service: "drive-v1"
resource: "import_task"
section: "Docs"
rate_limit: "100 per minute"
scopes:
  - "drive:drive"
updated: "1684896004000"
---

# Create an import task

Create import tasks. Supports importing Upgraded Docs, Sheets, Bitable and Docs. This API is asynchronous, you need to get the import result by Query the import task result
. Detailed operation steps can be found in the Import Usage Guide.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/drive/v1/import_tasks |
| HTTP Method | POST |
| Rate Limit | 100 per minute |
| Supported app types | custom,isv |
| Required scopes | `drive:drive` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `file_extension` | `string` | Yes | Format extension of the file to import **Example value**: "xlsx" |
| `file_token` | `string` | Yes | FileToken of the file to import in the Drive **Example value**: "boxcnxe5OxxxxxxxSNdsJviENsk" |
| `type` | `string` | Yes | Format of the destination document to import in Docs **Example value**: "sheet" |
| `file_name` | `string` | No | Name of the destination document to import in Docs. If this field is empty, the file name in the Drive is used. **Example value**: "test" |
| `point` | `import_task_mount_point` | Yes | Mount point |
|   `mount_type` | `int` | Yes | Mount type **Example value**: 1 **Optional values are**:  - `1`: Mount to My Space  |
|   `mount_key` | `string` | Yes | Mount location. If the mount_type field is 1, this field is set to the token of the directory in My Space. If this field is empty, the mount location is the root directory. **Example value**: "fldxxxxxxxx" | ### Request body example

{
    "file_extension": "xlsx",
    "file_token": "boxcnxe5OxxxxxxxSNdsJviENsk",
    "type": "sheet",
    "file_name": "test",
    "point": {
        "mount_type": 1,
        "mount_key": "fldxxxxxxxx"
    }
}

## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `ticket` | `string` | ID of the import task | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
        "ticket": "6990281xxxxxxxxxxx843"
    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 500 | 1069901 | internal error | Service internal error, consult [Oncall] (/ssl: ttdoc/ukTMukTMukTM/uUDN04SN0QjL1QDN/docs-overview #51f94b41) |
| 403 | 1069902 | no permission | No read or export permission scope. |
| 500 | 1069903 | internal error | Internal service error. For more information, contact support. |
| 400 | 1069904 | invalid param | Please check whether the file extension and the imported document type are correct |
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
| 429 | 1069923 | too many requests | Request frequency limit, please reduce the request frequency and try again later. |
