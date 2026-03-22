---
title: "Delete Attachment"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/task-v2/attachment/delete"
method: "DELETE"
api_path: "https://open.larksuite.com/open-apis/task/v2/attachments/:attachment_guid"
service: "task-v2"
resource: "attachment"
section: "Tasks v2"
rate_limit: "1000 per minute、50 per second"
scopes:
  - "task:attachment:write"
updated: "1699521713000"
---

# Delete Attachment

Providing an attachment GUID, delete the attachment. After deletion, the attachment cannot be recovered.

> Deleting an attachment requires the calling identity to have the edit permission of the resource to which the deleted attachment belongs, or the calling identity is exactly the the attachment uploader.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/task/v2/attachments/:attachment_guid |
| HTTP Method | DELETE |
| Rate Limit | 1000 per minute、50 per second |
| Supported app types | custom,isv |
| Required scopes | `task:attachment:write` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `attachment_guid` | `string` | GUID of attachment to delete, which can be created by Upload AttachmentAPI, or fetched byList AttachmentAPI. **Example value**: "b59aa7a3-e98c-4830-8273-cbb29f89b837" **Data validation rules**: - Maximum length: `100` characters | ## Response

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
| 400 | 1470400 | The request parameter is incorrect. | Check the `msg` in the return to determine the specific reason. |
| 404 | 1470404 | The attachment does not exist or has been deleted. | Verify whether the attachment still exists. |
| 500 | 1470500 | Server error. | Re-invoke the API with the same parameters. If it continues to fail, please contact support for feedback. |
| 403 | 1470403 | Missing permission to delete attachments. | Confirm that the calling identity has permission to delete attachments. |
