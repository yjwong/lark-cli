---
title: "Query Task Status"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/file/task_check"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/drive/v1/files/task_check"
service: "drive-v1"
resource: "file"
section: "Docs"
scopes:
  - "drive:drive"
  - "drive:drive.metadata:readonly"
  - "drive:drive:readonly"
updated: "1665221859000"
---

# Query asynchronous task status

Query the status information of asynchronous tasks such as delete folder.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/drive/v1/files/task_check |
| HTTP Method | GET |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `drive:drive` `drive:drive.metadata:readonly` `drive:drive:readonly` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" Learn more about obtaining and using access_token. | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `task_id` | `string` | Yes | File-related asynchronous task id **Example value**: "12345" | ## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `status` | `string` | The execution status of the asynchronous task. If the task is successfully executed, it returns success, if the task fails, it returns fail, and if the task is still executing, it returns process. | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
        "status": "success"
    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 500 | 1061001 | internal error. | Internal service error, such as timeout or failure in processing error codes. |
| 400 | 1061002 | params error. | Check whether the request parameters are correct. |
| 404 | 1061003 | not found. | Check whether the resource exists. |
