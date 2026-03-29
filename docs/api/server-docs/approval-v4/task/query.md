---
title: "Query the user's task list"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/approval-v4/task/query"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/approval/v4/tasks/query"
service: "approval-v4"
resource: "task"
section: "Approval"
scopes:
  - "approval:approval:readonly"
field_scopes:
  - "contact:user.employee_id:readonly"
updated: "1675167431000"
---

# Query the user's task list

Query task list by user and task group.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/approval/v4/tasks/query |
| HTTP Method | GET |
| Supported app types | custom,isv |
| Required scopes | `approval:approval:readonly` |
| Required field scopes | > The response body of the API contains the following sensitive fields, and they will be returned only after corresponding scopes are added. If you do not need the fields, it is not recommended that you request the scopes. `contact:user.employee_id:readonly` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `page_size` | `int` | No | Page size **Example value**: 100 **Default value**: `100` **Data validation rules**: - Maximum value: `200` |
| `page_token` | `string` | No | Page identifier. It is not filled in the first request, indicating traversal from the beginning; when there will be more groups, the new page_token will be returned at the same time, and the next traversal can use the page_token to get more groups **Example value**: "1" |
| `user_id` | `string` | Yes | User ID to query **Example value**: "example_user_id" |
| `topic` | `string` | Yes | Topic of task group to query, such as "Pending", "Done", etc. **Example value**: "1" **Optional values are**:  - `1`: Pending approval - `2`: Completed approvals - `3`: Approval initiated - `17`: Unread notifications - `18`: Read notifications  |
| `user_id_type` | `string` | No | User ID categories **Example value**: "open_id" **Optional values are**:  - `open_id`: Identifies a user to an app. The same user has different Open IDs in different apps. How to get Open ID - `union_id`: Identifies a user to a tenant that acts as a developer. A user has the same Union ID in apps developed by the same developer, and has different Union IDs in apps developed by different developers. A developer can use Union ID to link the same user's identities in multiple apps.How to get Union ID - `user_id`: Identifies a user to a tenant. The same user has different User IDs in different tenants. In one single tenant, a user has the same User ID in all apps （including store apps）. User ID is usually used to communicate user data between different apps. How to get User ID  **Default value**: `open_id` **When the value is `user_id`, the following field scopes are required**: `contact:user.employee_id:readonly` | ## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `tasks` | `task[]` | Task list |
|     `topic` | `string` | Task group of the task, such as "Pending", "Done", etc. **Optional values are**:  - `1`: Pending approval - `2`: Completed approvals - `3`: Approval initiated - `17`: Unread notifications - `18`: Read notifications  |
|     `user_id` | `string` | User ID for task |
|     `title` | `string` | Task title |
|     `urls` | `task_urls` | Task-related URLs |
|       `helpdesk` | `string` | Help Desk URL |
|       `mobile` | `string` | URL on mobile |
|       `pc` | `string` | URL on PC |
|     `process_external_id` | `string` | Process third-party ID, only for third-party processes, must be unique for current tenant and app |
|     `task_external_id` | `string` | Task third-party ID, only for third-party processes, must be unique within the process instance |
|     `status` | `string` | Task status **Optional values are**:  - `1`: Pending - `2`: Done - `17`: Unread - `18`: Read - `33`: Processing, mark for completion - `34`: Recall  |
|     `process_status` | `string` | Process instance status **Optional values are**:  - `0`: No process status, no label displayed - `1`: Process instance in progress - `2`: Approved - `3`: Rejected - `4`: Revoked - `5`: Terminated  |
|     `definition_code` | `string` | Process definition code |
|     `initiators` | `string[]` | List of initiator IDs |
|     `initiator_names` | `string[]` | List of initiator names |
|     `task_id` | `string` | Task ID, must be globally unique |
|     `process_id` | `string` | Process ID, globally unique |
|     `process_code` | `string` | Process code |
|     `definition_group_id` | `string` | Process definition group ID |
|     `definition_group_name` | `string` | Process definition group name |
|     `definition_id` | `string` | Process definition ID |
|     `definition_name` | `string` | Process definition name |
|   `page_token` | `string` | Page identifier, when has_more is true, a new page_token will also be returned. Otherwise, page_token will not be returned |
|   `has_more` | `boolean` | Whether the response body has more parameters |
|   `count` | `count` | List count, only returned on first page |
|     `total` | `int` | Total, only returns up to 999 items |
|     `has_more` | `boolean` | Still more, returns "true" when the quantity is at least 1,000 | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
        "tasks": [
            {
                "topic": "1",
                "user_id": "example_user_id",
                "title": "Task title example",
                "urls": {
                    "helpdesk": "https://blabla",
                    "mobile": "https://blabla",
                    "pc": "https://blabla"
                },
                "process_external_id": "example_instance_id",
                "task_external_id": "example_task_id",
                "status": "Todo",
                "process_status": "Running",
                "definition_code": "000000-00000000000000-0example",
                "initiators": [
                    "starter"
                ],
                "initiator_names": [
                    "Initiator name"
                ],
                "task_id": "1212564555454",
                "process_id": "1214564545474",
                "process_code": "123e4567-e89b-12d3-a456-426655440000",
                "definition_group_id": "1212564555454",
                "definition_group_name": "Process definition name",
                "definition_id": "1212564555454",
                "definition_name": "Process definition group name"
            }
        ],
        "page_token": "example_page_token",
        "has_more": true,
        "count": {
            "total": 123,
            "has_more": false
        }
    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 1390001 | param is invalid | Parameter error |
