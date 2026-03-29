---
title: "Status of third-party approval tasks"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/approval-v4/external_task/list"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/approval/v4/external_tasks"
service: "approval-v4"
resource: "external_task"
section: "Approval"
scopes:
  - "approval:approval"
  - "approval:approval:readonly"
updated: "1675167425000"
---

# Status of third-party approval tasks

This API is used to obtain the third-party approval status. After a user enters the query condition, the API returns the status of the approval instances that meet the condition.
The API supports a combination of multiple parameters,  including the following:

 1. Obtain the task status of a specified instance by using instance_ids

 2. Obtain the task status of a specified user by using user_ids

 3. Obtain all the tasks of a specified status by using status

 4. Obtain the next batch of data

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/approval/v4/external_tasks |
| HTTP Method | GET |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `approval:approval` `approval:approval:readonly` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `page_size` | `int` | No | Page size **Example value**: 100 **Data validation rules**: - Maximum value: `500` |
| `page_token` | `string` | No | Page identifier. It is not filled in the first request, indicating traversal from the beginning; when there will be more groups, the new page_token will be returned at the same time, and the next traversal can use the page_token to get more groups **Example value**: "nF1ZXJ5VGhlbkZldGNoCgAAAAAA6PZwFmUzSldvTC1yU" | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `approval_codes` | `string[]` | No | Approval definition codes, which are used to specify that only the data under these definitions will be obtained. **Example value**: B7B65FFE-C2GC-452F-9F0F-9AA8352363D6 **Data validation rules**: - Maximum length: `20` |
| `instance_ids` | `string[]` | No | Approval instance IDs, which are used to specify that only the data under these instances will be obtained. Up to 20 IDs are supported. **Example value**: oa_159160304 |
| `user_ids` | `string[]` | No | Approver's user_ids, which are used to specify that only the data of these users will be obtained. **Example value**: 112321 |
| `status` | `string` | No | Approval task status, which is used to specify that data under this status will be obtained. For status values **Example value**: "PENDING" **Optional values are**:  - `PENDING`: Approving - `APPROVED`: The approval process is over and the result is agreement - `REJECTED`: The approval process ends with a rejection - `TRANSFERRED`: Task transfer - `DONE`: The task passes but the approver does not operate; the approver cannot see the task, if you want to see it, you can cc the person.  | ### Request body example

{
    "approval_codes": [
        "B7B65FFE-C2GC-452F-9F0F-9AA8352363D6"
    ],
    "instance_ids": [
        "oa_159160304"
    ],
    "user_ids": [
        "112321"
    ],
    "status": "PENDING"
}

Remarks: 
  
1. When obtaining the task status of the specified instance through instance_ids, instance_ids a required field; 
```json
{
    "instance_ids": ["oa_159160304"]
}
```
  
2. When obtaining the task status of the specified user through user_ids, approval_codes, user_ids, and status are required fields; 
```json
{
    "approval_codes": ["B7B65FFE-C2GC-452F-9F0F-9AA8352363D6"],
    "user_ids": ["112321"],
    "status": "PENDING"
}
```
  
 3. When obtaining all tasks with the specified status through status, approval_codes and status are required fields; 
```json
{
    "approval_codes": [
        "E78F1022-A166-447C-8320-E151DA90D70F"
    ],
    "status": "PENDING"
}
```
  
 4. page_token a required field when getting the next batch of data through page_token

## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `data` | `external_task_list[]` | Return data |
|     `instance_id` | `string` | Approval instance ID |
|     `approval_id` | `string` | Approval's approval_code |
|     `approval_code` | `string` | Approval ID |
|     `status` | `string` | Current status of an approval instance. **Optional values are**:  - `PENDING`: Approving - `APPROVED`: The approval process is over and the result is agreement - `REJECTED`: The approval process ends with a rejection - `CANCELED`: Approval sponsor withdraws - `DELETED`: Approval removed - `HIDDEN`: Status hide (do not show status)  |
|     `update_time` | `string` | Last update time of an approval instance (in milliseconds) |
|     `tasks` | `external_task_item[]` | Approval tasks under an approval instance |
|       `id` | `string` | Approval task ID |
|       `status` | `string` | Approval task status. For status values **Optional values are**:  - `PENDING`: Approving - `APPROVED`: The approval process is over and the result is agreement - `REJECTED`: The approval process ends with a rejection - `TRANSFERRED`: Task transfer - `DONE`: The task passes but the approver does not operate; the approver cannot see the task, if you want to see it, you can cc the person.  |
|       `update_time` | `string` | Last update time of an approval task (in ms) |
|   `page_token` | `string` | Page identifier, when has_more is true, a new page_token will also be returned. Otherwise, page_token will not be returned |
|   `has_more` | `boolean` | Whether the response body has more parameters | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
        "data": [
            {
                "instance_id": "29075",
                "approval_id": "fwwweffff33111133xxx",
                "approval_code": "B7B65FFE-C2GC-452F-9F0F-9AA8352363D6",
                "status": "PENDING",
                "update_time": "1621863215000",
                "tasks": [
                    {
                        "id": "310",
                        "status": "PENDING",
                        "update_time": "1621863215000"
                    }
                ]
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
