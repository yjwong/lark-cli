---
title: "Add Custom Field To Resource"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/task-v2/custom_field/add"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/task/v2/custom_fields/:custom_field_guid/add"
service: "task-v2"
resource: "custom_field"
section: "Tasks v2"
rate_limit: "100 per minute"
scopes:
  - "task:custom_field:write"
updated: "1699521765000"
---

# Add Custom Field to Resource

Add custom fields to a resource. Currently resource types support tasklist. A custom field can be added to multiple tasklists. After adding, the tasklist can display the value of the custom fields of the task, by which functions such as filtering and grouping based on the field are available.

If the settings of the custom field are updated, all fields added by the field can receive the update and display it accordingly.

> Adding a custom field to a resource requires edit permission of both the custom field and the resource.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/task/v2/custom_fields/:custom_field_guid/add |
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
| `resource_type` | `string` | Yes | The resource type to add custom fields to a resource. Currently only tasklist is supported **Example value**: "tasklist" |
| `resource_id` | `string` | Yes | The resource id to add a custom field to, currently only tasklist_guid **Example value**: "0110a4bd-f24b-4a93-8c1a-1732b94f9593" | ### Request body example

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
| 400 | 1470400 | Incorrect request parameters, such as passing an unsupported type to resource_type. | The reason for the error is shown in the returned msg prompt. |
| 403 | 1470403 | Missing permission to add custom fields to resources. | Confirm that the calling identity has both editable permissions on the field and resource. |
| 404 | 1470404 | The custom word to be added or the resource to be added does not exist or has been deleted. | Verify that both the custom field and the resource exist. |
| 500 | 1470500 | Server error. | Try retrying the call. If it continues to fail, please contact support for feedback. |
