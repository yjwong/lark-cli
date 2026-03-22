---
title: "Delete Tasklist"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/task-v2/tasklist/delete"
method: "DELETE"
api_path: "https://open.larksuite.com/open-apis/task/v2/tasklists/:tasklist_guid"
service: "task-v2"
resource: "tasklist"
section: "Tasks v2"
rate_limit: "1000 per minute、50 per second"
scopes:
  - "task:tasklist:write"
updated: "1699521634000"
---

# Delete Tasklist

Delete a tasklist.

After a tasklist is deleted, no action can be performed on the tasklist and the tasklist can no longer be accessed. A deleted tasklist cannot be recovered.

> Deleting a tasklist requires tasklist manage permissions. Please refer to the "How are tasklist authorized?" section in Tasklist Features Overview .

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/task/v2/tasklists/:tasklist_guid |
| HTTP Method | DELETE |
| Rate Limit | 1000 per minute、50 per second |
| Supported app types | custom,isv |
| Required scopes | `task:tasklist:write` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `tasklist_guid` | `string` | GUID of tasklist to delete **Example value**: "d300a75f-c56a-4be9-80d1-e47653028ceb" | ## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {}
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 1470400 | Request parameter error | The reason for the error is shown in the returned msg. |
| 404 | 1470404 | The tasklist does not exist or has been deleted. | Confirm that the tasklist you want to delete exists or has been deleted. |
| 500 | 1470500 | Server error. | Try retrying the call. If it continues to fail, please contact support for feedback. |
| 403 | 1470403 | Missing tasklist manage permissions. | Check if the calling identity has the owner permission of the task. Please refer to the "How are tasklist authorized?" section in Tasklist Features Overview . |
