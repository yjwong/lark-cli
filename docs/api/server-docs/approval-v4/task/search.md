---
title: "List of tasks"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/approval-v4/task/search"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/approval/v4/tasks/search"
service: "approval-v4"
resource: "task"
section: "Approval"
scopes:
  - "approval:approval:readonly"
  - "approval:approval.list:readonly"
field_scopes:
  - "contact:user.employee_id:readonly"
updated: "1676878819000"
---

# List of tasks

The interface queries the list of qualified review tasks in the review system through different conditions.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/approval/v4/tasks/search |
| HTTP Method | POST |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `approval:approval:readonly` `approval:approval.list:readonly` |
| Required field scopes | > The response body of the API contains the following sensitive fields, and they will be returned only after corresponding scopes are added. If you do not need the fields, it is not recommended that you request the scopes. `contact:user.employee_id:readonly` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `page_size` | `int` | No | Page size **Example value**: 10 **Default value**: `10` **Data validation rules**: - Value range: `5` ～ `200` |
| `page_token` | `string` | No | Pagination mark, the first request is not filled, indicating that the traversal is started from the beginning; if there are more items in the paging query result, a new page_token will be returned at the same time, and the next traversal can use this page_token to obtain the query result **Example value**: "nF1ZXJ5VGhlbkZldGNoCgAAAAAA6PZwFmUzSldvTC1yU" |
| `user_id_type` | `string` | No | User ID categories **Example value**: "open_id" **Optional values are**:  - `open_id`: Identifies a user to an app. The same user has different Open IDs in different apps. How to get Open ID - `union_id`: Identifies a user to a tenant that acts as a developer. A user has the same Union ID in apps developed by the same developer, and has different Union IDs in apps developed by different developers. A developer can use Union ID to link the same user's identities in multiple apps.How to get Union ID - `user_id`: Identifies a user to a tenant. The same user has different User IDs in different tenants. In one single tenant, a user has the same User ID in all apps （including store apps）. User ID is usually used to communicate user data between different apps. How to get User ID  **Default value**: `open_id` **When the value is `user_id`, the following field scopes are required**: `contact:user.employee_id:readonly` | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `user_id` | `string` | No | Fill in the Approver id according to the x_user_type **Example value**: "lwiu098wj" |
| `approval_code` | `string` | No | Review the code **Example value**: "EB828003-9FFE-4B3F-AA50-2E199E2ED942" |
| `instance_code` | `string` | No | Review example code **Example value**: "EB828003-9FFE-4B3F-AA50-2E199E2ED943" |
| `instance_external_id` | `string` | No | Review example third party id note: and approval_code and set **Example value**: "EB828003-9FFE-4B3F-AA50-2E199E2ED976" |
| `group_external_id` | `string` | No | Review Definition Group 3rd party id Note: and instance_code take union **Example value**: "1234567" |
| `task_title` | `string` | No | Review task title (only for third-party reviews) **Example value**: "test" |
| `task_status` | `string` | No | Review task status, note: if not in the collection, report an error **Example value**: "PENDING" **Optional values are**:  - `PENDING`: Under review - `REJECTED`: Reject - `APPROVED`: Pass - `TRANSFERRED`: Transfer - `DONE`: Completed - `RM_REPEAT`: Deduplicate - `PROCESSED`: Processed - `ALL`: All states  |
| `task_start_time_from` | `string` | No | task query start time (unix millisecond timestamp) **Example value**: "1547654251506" |
| `task_start_time_to` | `string` | No | task query end time (unix millisecond timestamp) **Example value**: "1547654251506" |
| `locale` | `string` | No | Area **Example value**: "zh-CN" **Optional values are**:  - `zh-CN`: Chinese - `en-US`: English - `ja-JP`: Japanese  |
| `task_status_list` | `string[]` | No | Multiple states in the task_status can be selected, when this parameter is filled, the task_status invalidated **Example value**: PENDING |
| `order` | `int` | No | Sort by task time **Example value**: 2 **Optional values are**:  - `0`: Reverse update_time - `1`: Line up update_time - `2`: Reverse start_time - `3`: Line up start_time  **Default value**: `2` | > Note:
> 1. user_id, approval_code, instance_code, instance_external_id, and group_external_id can't be empty at the same time.
> 2. Take the union of the query results of approval_code and group_external_id , take the union of the query results of instance_code and instance_external_id , and take the intersection of other query conditions.
> 3. The query time span can't exceed 30 days. You must set either both the start and end times or neither. 

### Request body example

```json
{
    "user_id": "lwiu098wj",
    "approval_code": "EB828003-9FFE-4B3F-AA50-2E199E2ED942",
    "instance_code": "EB828003-9FFE-4B3F-AA50-2E199E2ED943",
    "instance_external_id": "EB828003-9FFE-4B3F-AA50-2E199E2ED976",
    "group_external_id": "1234567",
    "task_title": "test",
    "task_status": "PENDING",
    "task_start_time_from": "1547654251506",
    "task_start_time_to": "1547654251506",
    "locale": "zh-CN",
    "task_status_list": [
        "PENDING"
    ],
    "order": 2
}
```

## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `count` | `int` | Query return number |
|   `task_list` | `task_search_item[]` | Review task list |
|     `approval` | `instance_search_approval` | Review Definition |
|       `code` | `string` | Review the code |
|       `name` | `string` | Review Definition Name |
|       `is_external` | `boolean` | Whether it is a third-party review |
|       `external` | `instance_search_approval_external` | Third party review information |
|         `batch_cc_read` | `boolean` | Whether to support batch reading |
|       `approval_id` | `string` | Approval Definition Id |
|       `icon` | `string` | Approval definition icon information |
|     `group` | `instance_search_group` | Review definition grouping |
|       `external_id` | `string` | Review definition group external id |
|       `name` | `string` | Review definition group name |
|     `instance` | `instance_search_node` | Review example information |
|       `code` | `string` | Review example code |
|       `external_id` | `string` | Review instance external id |
|       `user_id` | `string` | Review instance originator id |
|       `start_time` | `string` | Review instance start time |
|       `end_time` | `string` | Review instance end time |
|       `status` | `string` | Review instance status **Optional values are**:  - `reject`: Reject - `pending`: Under review - `recall`: Withdraw - `deleted`: Deleted - `approved`: Pass  |
|       `title` | `string` | Review instance name (only third-party reviews have it) |
|       `extra` | `string` | Review instance extension field, string type json |
|       `serial_id` | `string` | Review serial number |
|       `link` | `instance_search_link` | Review example link (only third-party reviews have) |
|         `pc_link` | `string` | Review example pc link |
|         `mobile_link` | `string` | Review example mobile end link |
|     `task` | `task_search_node` | Review tasks |
|       `user_id` | `string` | Review task approver id |
|       `start_time` | `string` | Review task start time |
|       `end_time` | `string` | Review task end time |
|       `status` | `string` | Review task status **Optional values are**:  - `rejected`: Reject - `pending`: Under review - `approved`: Pass - `transferred`: Transfer - `done`: Completed - `rm_repeat`: Deduplicate - `processed`: Processed - `hidden`: Hidden  |
|       `title` | `string` | Review task name (only third-party reviews have it) |
|       `extra` | `string` | Review task extension field, string type json |
|       `link` | `instance_search_link` | Review example link (only third-party reviews have) |
|         `pc_link` | `string` | Review example pc link |
|         `mobile_link` | `string` | Review example mobile end link |
|       `task_id` | `string` | Task id |
|       `update_time` | `string` | Approval task update time |
|       `task_external_id` | `string` | External id |
|   `page_token` | `string` | Page Turn Token |
|   `has_more` | `boolean` | Are there more tasks to pull | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
        "count": 10,
        "task_list": [
            {
                "approval": {
                    "code": "EB828003-9FFE-4B3F-AA50-2E199E2ED943",
                    "name": "approval",
                    "is_external": true,
                    "external": {
                        "batch_cc_read": false
                    }
                },
                "group": {
                    "external_id": "0004",
                    "name": "groupA"
                },
                "instance": {
                    "code": "EB828003-9FFE-4B3F-AA50-2E199E2ED943",
                    "external_id": "0004_3ED52DC1-AA6C",
                    "user_id": "lwiu098wj",
                    "start_time": "1547654251506",
                    "end_time": "1547654251506",
                    "status": "pending",
                    "title": "test",
                    "extra": "{}",
                    "serial_id": "201902020001",
                    "link": {
                        "pc_link": "https://www.baidu.com/",
                        "mobile_link": "https://www.baidu.com/"
                    }
                },
                "task": {
                    "user_id": "lwiu098wj",
                    "start_time": "1547654251506",
                    "end_time": "1547654251506",
                    "status": "pending",
                    "title": "test",
                    "extra": "{}",
                    "link": {
                        "pc_link": "https://www.baidu.com/",
                        "mobile_link": "https://www.baidu.com/"
                    },
                    "task_id": "7110153401253494803",
                    "update_time": "1547654251506"
                }
            }
        ],
        "page_token": "nF1ZXJ5VGhlbkZldGNoCgAAAAAA6PZwFmUzSldvTC1yU",
        "has_more": false
    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 1390001 | param is invalid | Parameter error |
| 400 | 1395001 | There have been some errors. Please try again later | There was an error with the service, please try again later |
| 400 | 1390003 | instance code not found | Check that the review instance code is correct |
| 400 | 1390002 | approval code not found | Check that the review definition code is correct |
