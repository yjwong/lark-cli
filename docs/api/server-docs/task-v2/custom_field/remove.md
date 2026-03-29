---
title: "Remove Custom Field From Resource"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/task-v2/custom_field/remove"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/task/v2/custom_fields/:custom_field_guid/remove"
service: "task-v2"
resource: "custom_field"
section: "Tasks v2"
rate_limit: "100 per minute"
scopes:
  - "task:custom_field:write"
updated: "1699521768000"
---

# Remove Custom Field From Resource

Remove a custom field from the resource. After removal, the resource will no longer be able to use the field. Currently, only resource type "tasklist" is supported.

If you try to remove a custom field that does not exist in the resource, the api will return a success.

Note that custom fields are authorized through tasklist. If a custom field is removed from all associated tasklists, it means that any calling identity can no longer access the that custom field.

> Requires edit permission for the resource.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/task/v2/custom_fields/:custom_field_guid/remove |
| HTTP Method | POST |
| Rate Limit | 100 per minute |
| Supported app types | custom,isv |
| Required scopes | `task:custom_field:write` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `custom_field_guid` | `string` | custom field GUID, which can be created byCreate Custom Field, or queried byList Custom Field. **Example value**: "0110a4bd-f24b-4a93-8c1a-1732b94f9593" | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `resource_type` | `string` | Yes | The resource type to remove a custom field from a resource, currently only 'tasklist' is supported. **Example value**: "tasklist" |
| `resource_id` | `string` | Yes | To remove the resource id of a custom field from a resource, when "tasklist" is resource_type, fill in the tasklist GUID. **Example value**: "0110a4bd-f24b-4a93-8c1a-1732b94f9593" | ### Request body example

{
    "resource_type": "tasklist",
    "resource_id": "0110a4bd-f24b-4a93-8c1a-1732b94f9593"
}

## Response

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
| 400 | 1470400 | Incorrect request parameters, such as an unsupported type passed to the resource_type. | The reason for the error is shown in the returned msg prompt. |
| 403 | 1470403 | Missing permission to move custom fields from resources. | Confirm that the calling identity has edit rights to the resource. |
| 404 | 1470404 | The custom word to be removed or the resource to be removed does not exist or has been deleted. | Verify that the custom word to be removed or the resource to be removed exists or has been deleted. |
| 500 | 1470500 | Server error. | Try retrying the call. If it continues to fail, please contact support for feedback. |
