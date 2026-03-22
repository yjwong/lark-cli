---
title: "Add Dependency"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/task-v2/task/add_dependencies"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/task/v2/tasks/:task_guid/add_dependencies"
service: "task-v2"
resource: "task"
section: "Tasks v2"
rate_limit: "100 per minute"
scopes:
  - "task:task:write"
updated: "1699521602000"
---

# Add Dependency

Add one or more dependencies to a task. You can add prev-dependencies or next-dependencies of tasks. If a task with dependencies is in the same tasklist, you can show its dependencies through the Gantt View of the tasklist.

This API can also be used to modify the type of an existing dependency (prev to next or next to prev).

Note: The `task_guid` of added dependencies cannot contain duplicated task_guid, nor can the current task be added as its own dependency. Attempts to add an existing dependency are automatically ignored and will get a successful response.

> Adding a task dependence requires the invocation identity to have edit permissions for the current task and all tasks added as dependencies.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/task/v2/tasks/:task_guid/add_dependencies |
| HTTP Method | POST |
| Rate Limit | 100 per minute |
| Supported app types | custom,isv |
| Required scopes | `task:task:write` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `task_guid` | `string` | Task GUID **Example value**: "93b7bd05-35e6-4371-b3c9-6b7cbd7100c0" | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `dependencies` | `task_dependency[]` | No | Dependencies to add **Data validation rules**: - Length range: `1` ～ `50` |
|   `type` | `string` | Yes | dependency type **Example value**: "next" **Optional values are**:  - `prev`: prev-dependency - `next`: next-dependency  |
|   `task_guid` | `string` | Yes | GUIDs for dependent tasks **Example value**: "93b7bd05-35e6-4371-b3c9-6b7cbd7100c0" | ### Request body example

{
    "dependencies": [
        {
            "type": "next",
            "task_guid": "93b7bd05-35e6-4371-b3c9-6b7cbd7100c0"
        }
    ]
}

## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `dependencies` | `task_dependency[]` | All dependencies of the task after being added |
|     `type` | `string` | dependency type **Optional values are**:  - `prev`: pre-dependency - `next`: postdependency  |
|     `task_guid` | `string` | GUIDs for dependent tasks | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
        "dependencies": [
            {
                "type": "next",
                "task_guid": "93b7bd05-35e6-4371-b3c9-6b7cbd7100c0"
            }
        ]
    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 1470400 | Bad reqeust, such as the passing an invalid depdendency type, etc. | Refer to the returned error message for the specific error reason. |
| 403 | 1470403 | Missing edit permissions for tasks | The calling identity needs to have edit perssmions for the task that adds the dependency and the tasks that are added as dependencies. |
| 404 | 1470404 | The tasks do not exist or have been deleted. | Verify that both the task to add the dependency and the tasks to be added as dependencies exist or have not been deleted. |
| 500 | 1470500 | Server error. | Internal server error. Retry to invoke the api, and contact support if it continues to fail. |
