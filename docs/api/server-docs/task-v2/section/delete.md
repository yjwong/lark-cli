---
title: "Delete Section"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/task-v2/section/delete"
method: "DELETE"
api_path: "https://open.larksuite.com/open-apis/task/v2/sections/:section_guid"
service: "task-v2"
resource: "section"
section: "Tasks v2"
rate_limit: "1000 per minute、50 per second"
scopes:
  - "task:section:write"
updated: "1699521733000"
---

# Delete Section

Delete a section. The tasks in the deleted section will be moved to the default section of the resource to which the deleted section belongs.

Default section cannot be deleted.

> Need to edit permissions of the resource the section belongs to.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/task/v2/sections/:section_guid |
| HTTP Method | DELETE |
| Rate Limit | 1000 per minute、50 per second |
| Supported app types | custom,isv |
| Required scopes | `task:section:write` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `section_guid` | `string` | section GUID to delete **Example value**: "9842501a-9f47-4ff5-a622-d319eeecb97f" | ## Response

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
| 400 | 1470400 | Incorrect request parameters, such as the section_guid passed in is not valid. | The reason for the error is shown in the returned msg prompt. |
| 404 | 1470404 | The section to delete does not exist or has been deleted. | Confirm that the section exists or not, or has been deleted. |
| 500 | 1470500 | Server error. | Retry to invoke the api with same parameters. If it continues to fail, please contact support for feedback. |
| 403 | 1470403 | Missing permission to delete section. | Confirm that the calling identity has permission to delete the section. |
