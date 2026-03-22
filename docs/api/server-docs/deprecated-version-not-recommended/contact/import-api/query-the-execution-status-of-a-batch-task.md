---
title: "Query the Execution Status of a Batch Task"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uUDOwUjL1gDM14SN4ATN"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/contact/v2/task/get?task_id=0123456789abcdef0123456789abcdef"
section: "Deprecated Version (Not Recommended)"
updated: "1646999111000"
---

# Query the execution status of a batch task

This API is used to query the current execution status and execution result of an asynchronous task in Contacts.

> Apps invoking this API must have permission to update Contacts.
> Marketplace Apps do not have permission to invoke this API.
> The application should apply for `update contacts` and `read contacts by application identity` permissions to invoke this API.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/contact/v2/task/get?task_id=0123456789abcdef0123456789abcdef |
| HTTP Method | GET |
| Required scopes Enable any scope from the list | update contact read contacts by application identity | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value Format**: "Bearer `access_token`" **Example Value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" Learn more about obtaining and using access_token. |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Query parameters
|Parameter|Type|Mandatory|Description|
|-|-|-|-|
|task_id|string|Yes|Asynchronous task ID returned by batch APIs.|
## Response

### Response body
|Parameter|Description|
|-|-|
|code|Error code, anything but 0 indicates failure.|
|msg|Response code description.|
|data|-|
|  ∟task_id|Asynchronous task ID.|
|  ∟type|Task type. There are two types of task currently: add_user, add_department.|
|  ∟status|The current execution status of the task. Lower than 9: Executing, 9: Executed, 10: Execution failed.|
|  ∟progress|Task progress percentage.|
|  ∟total_num|Total item count of task.|
|  ∟success_num|Number of items successfully executed by the current task.|
|  ∟fail_num|Number of items that failed to be executed by the current task.|
|  ∟create_time|Task creation time, Unix timestamp in seconds.|
|  ∟finish_time|Task completion time, Unix timestamp in seconds. If the task has not been executed successfully, this field is not returned.|
|  ∟task_info|Task execution result list. If the task has not been executed successfully, this field is not returned.The execution result uses the same order as the body of the request used to add the task.|
|    ∟code|Subtask response code, anything but 0 indicates failure.|
|    ∟msg|Description of the subtask response code.|
|    ∟action|The operation the subtask is carrying out. 1: Add, 2: Update. If task execution failed, this field is not returned.|
|    ∟name|Subtask request name. If operating a user, it is the user name. If operating a department, it is a department name.|
|==============|The following fields apply to user operations:|
|    ∟email|User email address in the request.|
|    ∟mobile|User mobile number in the request.|
|    ∟user_id|The user's ID within the company in the request, or an automatically generated user ID.|
|    ∟departments|The user’s department in the request.|
|    ∟open_id|The open_id generated for the user. If task execution failed, this field is not returned.|
|==============|The following fields apply to department operations:|
|    ∟department_id|The custom department ID in the request, or an automatically generated department ID.|
|    ∟parent_id|The parent department ID in the request.|
|    ∟chat_id|The department group ID, only returned if it exists.|
### Request body example
When the task is an operation on users:
```json
{
    "code": 0,
    "msg": "success",
    "data": {
        "task_id": "0123456789abcdef0123456789abcdef",
        "type": "add_user",
        "status": 9,
        "progress": 100,
        "total_num": 2,
        "success_num": 1,
        "fail_num": 1,
        "create_time": 1565878519,
        "finish_time": 1565878523,
        "task_info": [
            {
                "code": 0,
                "msg": "success",
                "action": 1,
                "departments": [
                    "custom_1"
                ],
                "email": "custom_1@feishu.cn",
                "mobile": "+8613000000001",
                "name": "custom_1",
                "open_id": "ou_123456784b68a7c89abcdef092dc09ea",
                "user_id": "custom_1"
            },
            {
                "code": 40013,
                "msg": "invalid user_id",
                "departments": [
                    "custom_2"
                ],
                "mobile": "+8613000000002",
                "name": "custom_2",
                "user_id": "_custom_2"
            }
        ]
    }
}
```
When the task is an operation on departments:
```json
{
    "code": 0,
    "msg": "success",
    "data": {
        "task_id": "0123456789abcdef0123456789abcdef",
        "type": "add_department",
        "status": 9,
        "progress": 100,
        "total_num": 2,
        "success_num": 2,
        "fail_num": 0,
        "create_time": 1565878519,
        "finish_time": 1565878523,
        "task_info": [
            {
                "code": 0,
                "msg": "success",
                "action": 1,
                "chat_id": "oc_123456784b68a7c89abcdef092dc09ea",
                "department_id": "custom_1",
                "name": "custom_1",
                "parent_id": "0"
            },
            {
                "code": 0,
                "msg": "success",
                "action": 1,
                "department_id": "custom_2",
                "name": "custom_2",
                "parent_id": "custom_1"
            }
        ]
    }
}
```
### Error code

For details, please refer to: Service-side error codes
