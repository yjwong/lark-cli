---
title: "Remove Dependency"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/task-v2/task/remove_dependencies"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/task/v2/tasks/:task_guid/remove_dependencies"
service: "task-v2"
resource: "task"
section: "Tasks v2"
rate_limit: "100 per minute"
scopes:
  - "task:task:write"
updated: "1699521606000"
---

# Remove Dependency

Removes one or more dependencies from a task. You can remove dependencies by passing their `task_guid`.

Note that if the dependency to be removed is not a dependency of the current task, it will be automatically ignored. The API will return success.

> When removing task dependencies, edit permission for the current task is required.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/task/v2/tasks/:task_guid/remove_dependencies |
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
| `task_guid` | `string` | GUID of task to remove the dependencies **Example value**: "93b7bd05-35e6-4371-b3c9-6b7cbd7100c0" | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `dependencies` | `task_dependency[]` | Yes | Dependencies to remove **Data validation rules**: - Length range: `1` ～ `50` |
|   `task_guid` | `string` | Yes | GUIDs for dependent tasks **Example value**: "93b7bd05-35e6-4371-b3c9-6b7cbd7100c0" | ### Request body example

{
    "dependencies": [
        {
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
|   `dependencies` | `task_dependency[]` | Task GUID after removal |
|     `type` | `string` | dependency type **Optional values are**:  - `prev`: prev-dependency - `next`: next-dependency  |
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
| 400 | 1470400 | Incorrect request parameters, such as an invalid `task_guid` passed in. | Refer to the returned error message for the specific error reason. |
| 403 | 1470403 | Missing edit permission for the task. | Removing dependencies requires edit permission for the current task. |
| 404 | 1470404 | The task to remove the dependencies does not exist or has been deleted. | Confirm whether the task exists or is deleted. |
| 500 | 1470500 | Server error. | Internal server error. If there is a persistent error after retrying the call multiple times, you can contact support. |
