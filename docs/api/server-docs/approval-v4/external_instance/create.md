---
title: "Sync instances"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/approval-v4/external_instance/create"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/approval/v4/external_instances"
service: "approval-v4"
resource: "external_instance"
section: "Approval"
scopes:
  - "approval:approval"
  - "approval:external_instance"
updated: "1752806937000"
---

# Sync instances

Approval transfer lies in the third-party system instead of the Approval Center. The approval instances, approval tasks, and approval CC data generated after approval transfer are synced from the third-party system to the Approval Center.
  
Users can browse the instances, tasks, and CC information synced from the third-party system in the Approval Center, and can redirect to the third-party system for more detailed viewing and action. The instance information is in the [Initiated] list, the task information is in the [Pending] and [Approved] lists, and the CC information is in the [CC'd] list.

The third-party system can also configure the callback API for approval tasks, so that the approver can directly perform the approval action in the Approval Center, and the Approval Center will call back the third-party system. After receiving the callback, the third-party system will update the task information and sync the new task information back to the Approval Center to form a closed loop.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/approval/v4/external_instances |
| HTTP Method | POST |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `approval:approval` `approval:external_instance` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `approval_code` | `string` | Yes | Approval definition code, which creates the value returned by the approval definition to indicate which process the instance belongs to. This field affects the title of the instance in the list, which is taken from the name field of the corresponding definition. **Example value**: "81D31358-93AF-92D6-7425-01A5D67C4E71" |
| `status` | `string` | Yes | Approval instance status **Example value**: "PENDING" **Optional values are**:  - `PENDING`: Under review - `APPROVED`: The request is approved when the approval process ends. - `REJECTED`: The request is rejected when the approval process ends. - `CANCELED`: The approval is canceled by the initiator. - `DELETED`: he approval is deleted. - `HIDDEN`: Status is hidden (i.e. not displayed). - `TERMINATED`: Approval terminated  |
| `extra` | `string` | No | Approval instance extension JSON **Example value**: "{\"xxx\":\"xxx\"}" |
| `instance_id` | `string` | Yes | Approval instance unique identification, user-defined, need to ensure that the tenant under the unique **Example value**: "24492654" |
| `links` | `external_instance_link` | Yes | Approval instance link collection, which is used to redirect to the third-party system from the [Initiated] list. Either the pc_link field or mobile_link field must be filled in to specify the link for redirection, which is not affected by the platform. |
|   `pc_link` | `string` | Yes | Redirection link for PC, which is used for redirection when users use Lark on PC **Example value**: "https://applink.larksuite.com/client/mini_program/open?mode=appCenter&appId=cli_9c90fc38e07a9101&path=pc/pages/detail?id=1234" |
|   `mobile_link` | `string` | No | Redirection link for mobile, which is used for redirection when users use Lark on mobile **Example value**: "https://applink.larksuite.com/client/mini_program/open?appId=cli_9c90fc38e07a9101&path=pages/detail?id=1234" |
| `title` | `string` | No | Approval display name. If this field is filled in, the approval title in the approval list will use this field. If this field isn't filled in, the approval title will use the name of the approval definition. **Example value**: "@i18n@1" |
| `form` | `external_instance_form[]` | No | Form data filled in when a user submits application, which is displayed in all the approval lists. More than one value can be passed. However, in the Approval Center, the first 2 values are displayed on PC, while the first 3 values on mobile. The display length can't exceed 2,048 characters. |
|   `name` | `string` | No | Form field name **Example value**: "@i18n@2" |
|   `value` | `string` | No | Form value **Example value**: "@i18n@3" |
| `user_id` | `string` | No | Approver initiator's user_id. The initiator can view all initiated approvals in the [Initiated] list. In the [Pending], [Approved], and [CC'd] lists, this field shows the initiator of the approval. Either the initiator's open ID or user ID is required. **Example value**: "a987sf9s" |
| `user_name` | `string` | No | Approval initiator's user name. If the initiator isn't a real user (for example, a department) and has no user ID, use this field to pass the name. **Example value**: "@i18n@9" |
| `open_id` | `string` | No | Either the initiator's open ID or user ID is required. **Example value**: "ou_be73cbc0ee35eb6ca54e9e7cc14998c1" |
| `department_id` | `string` | No | Initiator department, used in the list to display the department which the initiator belongs to.This field isn't displayed if it isn't passed.If the user hasn't joined any department, the tenant name is displayed when "" is passed.If the department_name field is passed, the department name is displayed. **Example value**: "od-8ec33278bc2" |
| `department_name` | `string` | No | Approval initiator's department. If the initiator isn't a real user (for example, a department) and has no department_id, use this field to pass the name. **Example value**: "@i18n@10" |
| `start_time` | `string` | Yes | Approval initiation time (Unix timestamp in ms) **Example value**: "1556468012678" |
| `end_time` | `string` | Yes | Approval instance end time: The value is 0 for uncompleted approvals (Unix timestamp in ms) **Example value**: "1556468012678" |
| `update_time` | `string` | Yes | Last update time of the approval instance, unix millisecond timestamp, used to push data version control.If the update_mode value is UPDATE, the approval instance information in the Approval Center will only be updated when the passed update_time changes (becomes larger).This field is mainly used to avoid updating new data with old data when concurrency occurs. **Example value**: "1556468012678" |
| `display_method` | `string` | No | The way to open the approval instance on the list page.Optional values are: **Example value**: "BROWSER" **Optional values are**:  - `BROWSER`: Opens by redirecting to the default system browser - `SIDEBAR`: Opens in the sidebar of Lark - `NORMAL`: Opens in the page embedded in Lark - `TRUSTEESHIP`: Open with hosting  |
| `update_mode` | `string` | No | Update method, When update_mode=REPLACE, each time the currently pushed data is taken as the final data, and the redundant tasks and CC data in the Approval Center will be deleted (not in the data pushed this time). When update_mode=UPDATE, the data in the Approval Center won't be deleted, but only the instance and task data will be added and updated. **Example value**: "UPDATE" **Optional values are**:  - `REPLACE`: Replaces all. The value is set by default. - `UPDATE`: Updates incrementally.  |
| `task_list` | `external_instance_task_node[]` | No | Task list **Data validation rules**: - Maximum length: `200` |
|   `task_id` | `string` | Yes | The unique ID within the approval instance, which is used to locate data when updating approval tasks **Example value**: "112534" |
|   `user_id` | `string` | No | Approver's user_id. The task will appear in the [Pending] or [Approved] list of the approver. **Example value**: "a987sf9s" |
|   `open_id` | `string` | No | Either the approver's open ID or user ID is required. **Example value**: "ou_be73cbc0ee35eb6ca54e9e7cc14998c1" |
|   `title` | `string` | No | Approval task name **Example value**: "i18n1" |
|   `links` | `external_instance_link` | Yes | Redirection link in the [Pending] or [Approved], used to redirect to the third-party systemEither the pc_link field or mobile_link field must be filled in to specify the link for redirection, which is not affected by the platform. |
|     `pc_link` | `string` | Yes | Redirection link for PC, which is used for redirection when users use Lark on PC **Example value**: "https://applink.larksuite.com/client/mini_program/open?mode=appCenter&appId=cli_9c90fc38e07a9101&path=pc/pages/detail?id=1234" |
|     `mobile_link` | `string` | No | Redirection link for mobile, which is used for redirection when users use Lark on mobile **Example value**: "https://applink.larksuite.com/client/mini_program/open?appId=cli_9c90fc38e07a9101&path=pages/detail?id=1234" |
|   `status` | `string` | Yes | Task status **Example value**: "PENDING" **Optional values are**:  - `PENDING`: Under review - `APPROVED`: The task is approved - `REJECTED`: The task is rejected - `TRANSFERRED`: The task is transferred - `DONE`: The task is approved but hasn't been processed by the approver. The approver can't see the task. You can CC this individual to make the task visible.  |
|   `extra` | `string` | No | Extend json, the reason for the end of the task needs to be passed complete_reason field. Enumeration values and corresponding descriptions: - approved: agreed - rejected: rejected - node_auto_reject: Automatic rejection (due to logical judgment) - specific_rollback: return (including return to the originator, return to any intermediate approver) - add: and sign (add a new approver, approve with me) - add_pre: Pre-approver (add new approver, approve before me) - add_post: post-approver (add new approver, approve after me) - delete_assignee: Reductions - forward_resign: transfer (transfer to someone else for approval) - recall: revocation (withdrawal of documents, invalidation of documents) - delete: delete the approval slip - admin_forward: Administrator in the background operation transfer - system_forward: the system automatically transfers - auto_skip: Auto Pass - manual_skip: manually skip - submit_again: resubmit tasks - restart: restart the process - others: others (as the bottom line) **Example value**: "{\"xxx\":\"xxx\",\"complete_reason\":\"approved\"}" |
|   `create_time` | `string` | Yes | Task creation time (Unix timestamp in ms) **Example value**: "1556468012678" |
|   `end_time` | `string` | Yes | Task completion time: The value is 0 for uncompleted approvals (Unix timestamp in ms) **Example value**: "1556468012678" |
|   `update_time` | `string` | No | Last update time of the task, unix millisecond timestamp, used to push data version control. The update policy is the same as the update_time in instance. **Example value**: "1556468012678" |
|   `action_context` | `string` | No | Action context, which is taken in the callback request to pass the context data of the task as users perform actions **Example value**: "123456" |
|   `action_configs` | `action_config[]` | No | Task-level action configuration. Quick approval is available on mobile. |
|     `action_type` | `string` | Yes | Action type. Each task can be configured with 2 actions, which will be displayed in the approval list. When users perform actions, the callback request will carry this field, indicating whether the user has approved or rejected the action. **Optional values are**:  - `APPROVE`: Approved - `REJECT`: Rejected - `{KEY}`: Arbitrary string. The action name is required when arbitrary string is used.  **Example value**: "APPROVE" |
|     `action_name` | `string` | No | Action name. The i18n key is used for frontend display. If the action_type field value is neither APPROVAL nor REJECT, this field is required to display a specific action name. **Example value**: "@i18n@5" |
|     `is_need_reason` | `boolean` | No | Whether the comments are required. If true, users will be redirected to the page to fill in comments when they perform actions. **Example value**: false |
|     `is_reason_required` | `boolean` | No | Whether the approval comments are required. **Example value**: false |
|     `is_need_attachment` | `boolean` | No | Whether attachments can be uploaded for the comments **Example value**: false |
|   `display_method` | `string` | No | The way to open the approval task on the list page.Optional values are **Example value**: "BROWSER" **Optional values are**:  - `BROWSER`: Opens by redirecting to the default system browser - `SIDEBAR`: Opens in the sidebar of Lark - `NORMAL`: Opens in the page embedded in Lark - `TRUSTEESHIP`: Open in managed mode  |
|   `exclude_statistics` | `boolean` | No | Tripartite mission support is not included in efficiency statistics. false: Include efficiency statistics. true: not included in efficiency statistics **Example value**: false **Default value**: `false` |
|   `node_id` | `string` | No | Node id: must meet both - In a process, each node id is unique. For example, the Node_id of each node such as "direct manager" and "next-level superior" in a process are different - In the same process definition, the same node in different approval instances should Node_id remain unchanged. For example, Zhang San and Li Si initiated leave applications respectively, and the node_id of the "direct manager" node in these two approval instances should remain the same **Example value**: "node" |
|   `node_name` | `string` | No | Node names, such as "Financial Approval" and "Legal Approval", support three languages: Chinese, English and Japanese. Example: i18n@name. You need to pass the internationalized copy corresponding to the name in the i18n_resources **Example value**: "i18n@name" |
| `cc_list` | `cc_node[]` | No | CC list **Data validation rules**: - Maximum length: `200` |
|   `cc_id` | `string` | Yes | Unique ID in the approval instance **Example value**: "123456" |
|   `user_id` | `string` | No | CC employee ID **Example value**: "12345" |
|   `open_id` | `string` | No | Either the CC sender's open ID or user ID is required. **Example value**: "ou_be73cbc0ee35eb6ca54e9e7cc14998c1" |
|   `links` | `external_instance_link` | Yes | Redirection link, used for redirection in the [CC'd] listEither the pc_link field or mobile_link field must be filled in to specify the link for redirection, which is not affected by the platform. |
|     `pc_link` | `string` | Yes | Redirection link for PC, which is used for redirection when users use Lark on PC **Example value**: "https://applink.larksuite.com/client/mini_program/open?mode=appCenter&appId=cli_9c90fc38e07a9101&path=pc/pages/detail?id=1234" |
|     `mobile_link` | `string` | No | Redirection link for mobile, which is used for redirection when users use Lark on mobile **Example value**: "https://applink.larksuite.com/client/mini_program/open?appId=cli_9c90fc38e07a9101&path=pages/detail?id=1234" |
|   `read_status` | `string` | Yes | Read status. An empty cell indicates that the read/unread status is unavailable: **Example value**: "READ" **Optional values are**:  - `READ`: Read - `UNREAD`: Unread  |
|   `extra` | `string` | No | Extend json **Example value**: "{\"xxx\":\"xxx\"}" |
|   `title` | `string` | No | CC task name **Example value**: "xxx" |
|   `create_time` | `string` | Yes | CC initiation time (Unix timestamp in ms) **Example value**: "1556468012678" |
|   `update_time` | `string` | Yes | Last update time of CC, unix millisecond timestamp, used to push data version control.The update policy is the same as the update_time in instance. **Example value**: "1556468012678" |
|   `display_method` | `string` | No | The way to open the approval task on the list page.Optional values are: **Example value**: "BROWSER" **Optional values are**:  - `BROWSER`: Opens by redirecting to the default system browser - `SIDEBAR`: Opens in the sidebar of Lark - `NORMAL`: Opens in the page embedded in Lark - `TRUSTEESHIP`: Open in managed mode  |
| `i18n_resources` | `i18n_resource[]` | Yes | Internationalized text |
|   `locale` | `string` | Yes | Language **Example value**: "zh-CN" **Optional values are**:  - `zh-CN`: Chinese - `en-US`: English - `ja-JP`: Japanese  |
|   `texts` | `i18n_resource_text[]` | Yes | Key, value, and i18n key of copy starts with @i18n@. This field is mainly used for internationalized text. Word order users can pass texts in multiple languages at the same time, and the Approval Center will use the corresponding text according to the user's language environment. If the text of the user's language environment isn't passed, the text of default language will be used. |
|     `key` | `string` | Yes | Copywriting key **Example value**: "@i18n@1" |
|     `value` | `string` | Yes | Copywriting **Example value**: "People" |
|   `is_default` | `boolean` | Yes | Whether to use the default language. If the default language is used, all keys should be contained; if a non-default language is used but its keys do not exist, the default language will be used instead. **Example value**: true |
| `trusteeship_url_token` | `string` | No | Document escrow authentication token, escrow callback will be accompanied by this token to help business party authentication **Example value**: "788981c886b1c28ac29d1e68efd60683d6d90dfce80938ee9453e2a5f3e9e306" |
| `trusteeship_user_id_type` | `string` | No | The type of user will affect the selection of the user identification field of the request parameter, including the target user returned by the signature operation. Currently only "user_id" is supported **Example value**: "user_id" |
| `trusteeship_urls` | `trusteeship_urls` | No | The URL address of the interface of the document hosting callback access party |
|   `form_detail_url` | `string` | No | Get the url address of the data related to the form schema **Example value**: "https://#{your_domain}/api/form_detail" |
|   `action_definition_url` | `string` | No | Represents the url address to obtain the data of the approval operation area **Example value**: "https://#{your_domain}/api/action_definition" |
|   `approval_node_url` | `string` | No | Get the url address of the data related to the approval record **Example value**: "https://#{your_domain}/api/approval_node" |
|   `action_callback_url` | `string` | No | The url address of the callback when performing an approval operation **Example value**: "https://#{your_domain}/api/action_callback" |
|   `pull_business_data_url` | `string` | No | Obtain the URL address of the managed dynamic data. When using this interface, you must ensure that the interface address is synchronized in the data of the historical managed document. If there is no interface in the historical document, you need to re-synchronize the data of the historical managed document to update the URL. This interface is used for the interactive use of the Lark approval front end and the line of business. Only the specific components of the approval front end (the components provided by the Lark approval front end and the components that need to interface with the line of business) will be required **Example value**: "https://#{your_domain}/api/pull_business_data" | > **Notes：**
> 
> It is required to ensure the uniqueness of various unique IDs in the approval instance, even if they don't belong to the same entity.Entities include: instance , task , cc . If the instance ID, task ID, and cc ID are duplicated, the corresponding to-do, done, cc, and initiated data will not be visible in the approval center task.

### Request body example

```json
{
    "approval_code": "84C18825-A3D2-41D0-891F-E7A2424C5D48",
    "instance_id": "3162634",
    "status": "PENDING",
    "extra": "",
    "links": {
        "pc_link": "http://applink.larksuite.com/sso/common?redirectUrl=/seeyon/main.do?method=main&client=pc",
        "mobile_link": "http://applink.larksuite.com/sso/common?redirectUrl=/seeyon/main.do?method=main&client=pc"
    },
    "title": "@i18n@1",
    "form": [{
        "name": "@i18n@2",
        "value": "@i18n@3"
    }],
    "user_id": "16fb9ff3",
    "user_name": "张三",
    "open_id": "123",
    "department_id": "",
    "department_name": "hr",
    "start_time": "1657093395000",
    "update_time": "1657093395000",
    "end_time": 0,
    "update_mode": "REPLACE",
    "task_list": [{
        "task_id": "112253",
        "user_id": "16fb9ff3",
        "links": {
            "pc_link": "http://",
            "mobile_link": "http://"
        },
        "status": "PENDING",
        "extra": "",
        "title": "agree",
        "create_time": "1638468921000",
        "end_time": 0,
        "update_time": "1638468921000",
        "action_context": "123456",
        "action_configs": [{
            "action_type": "APPROVE",
            "action_name": "@i18n@1",
            "is_need_reason": true,
            "is_reason_required": true,
            "is_need_attachment": true
        }]
    }],
    "cc_list": [{
        "cc_id": "1231243",
        "user_id": "16fb9ff3",
        "open_id": "",
        "links": {
            "pc_link": "http://",
            "mobile_link": "http://"
        },
        "read_status": "READ",
        "extra": "",
        "title": "XXX",
        "create_time": "1657093395000",
        "update_time": "1657093395000"
    }],
    "i18n_resources": [{
        "locale": "zh-CN",
        "is_default": true,
        "texts": [ {"key":"@i18n@1", "value":"test"}, {"key":"@i18n@2", "value":"天"},{"key":"@i18n@3", "value":"2022-07-06"}]
    }]
}
```

## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `data` | `external_instance` | { "Code": 0, "Msg": "success", "Data": { } } |
|     `approval_code` | `string` | Approval definition code, which creates the value returned by the approval definition to indicate which process the instance belongs to. This field affects the title of the instance in the list, which is taken from the name field of the corresponding definition. |
|     `status` | `string` | Approval instance status **Optional values are**:  - `PENDING`: Under review - `APPROVED`: The request is approved when the approval process ends. - `REJECTED`: The request is rejected when the approval process ends. - `CANCELED`: The approval is canceled by the initiator. - `DELETED`: he approval is deleted. - `HIDDEN`: Status is hidden (i.e. not displayed). - `TERMINATED`: Approval terminated  |
|     `extra` | `string` | Approval instance extension JSON |
|     `instance_id` | `string` | Approval instance unique identification, user-defined, need to ensure that the tenant under the unique |
|     `links` | `external_instance_link` | Approval instance link collection, which is used to redirect to the third-party system from the [Initiated] list. Either the pc_link field or mobile_link field must be filled in to specify the link for redirection, which is not affected by the platform. |
|       `pc_link` | `string` | Redirection link for PC, which is used for redirection when users use Lark on PC |
|       `mobile_link` | `string` | Redirection link for mobile, which is used for redirection when users use Lark on mobile |
|     `title` | `string` | Approval display name. If this field is filled in, the approval title in the approval list will use this field. If this field isn't filled in, the approval title will use the name of the approval definition. |
|     `form` | `external_instance_form[]` | Form data filled in when a user submits application, which is displayed in all the approval lists. More than one value can be passed. However, in the Approval Center, the first 2 values are displayed on PC, while the first 3 values on mobile. The display length can't exceed 2,048 characters. |
|       `name` | `string` | Form field name |
|       `value` | `string` | Form value |
|     `user_id` | `string` | Approver initiator's user_id. The initiator can view all initiated approvals in the [Initiated] list. In the [Pending], [Approved], and [CC'd] lists, this field shows the initiator of the approval. Either the initiator's open ID or user ID is required. |
|     `user_name` | `string` | Approval initiator's user name. If the initiator isn't a real user (for example, a department) and has no user ID, use this field to pass the name. |
|     `open_id` | `string` | Either the initiator's open ID or user ID is required. |
|     `department_id` | `string` | Initiator department, used in the list to display the department which the initiator belongs to.This field isn't displayed if it isn't passed.If the user hasn't joined any department, the tenant name is displayed when "" is passed.If the department_name field is passed, the department name is displayed. |
|     `department_name` | `string` | Approval initiator's department. If the initiator isn't a real user (for example, a department) and has no department_id, use this field to pass the name. |
|     `start_time` | `string` | Approval initiation time (Unix timestamp in ms) |
|     `end_time` | `string` | Approval instance end time: The value is 0 for uncompleted approvals (Unix timestamp in ms) |
|     `update_time` | `string` | Last update time of the approval instance, unix millisecond timestamp, used to push data version control.If the update_mode value is UPDATE, the approval instance information in the Approval Center will only be updated when the passed update_time changes (becomes larger).This field is mainly used to avoid updating new data with old data when concurrency occurs. |
|     `display_method` | `string` | The way to open the approval instance on the list page.Optional values are: **Optional values are**:  - `BROWSER`: Opens by redirecting to the default system browser - `SIDEBAR`: Opens in the sidebar of Lark - `NORMAL`: Opens in the page embedded in Lark - `TRUSTEESHIP`: Open with hosting  |
|     `update_mode` | `string` | Update method, When update_mode=REPLACE, each time the currently pushed data is taken as the final data, and the redundant tasks and CC data in the Approval Center will be deleted (not in the data pushed this time). When update_mode=UPDATE, the data in the Approval Center won't be deleted, but only the instance and task data will be added and updated. **Optional values are**:  - `REPLACE`: Replaces all. The value is set by default. - `UPDATE`: Updates incrementally.  |
|     `task_list` | `external_instance_task_node[]` | Task list |
|       `task_id` | `string` | The unique ID within the approval instance, which is used to locate data when updating approval tasks |
|       `user_id` | `string` | Approver's user_id. The task will appear in the [Pending] or [Approved] list of the approver. |
|       `open_id` | `string` | Either the approver's open ID or user ID is required. |
|       `title` | `string` | Approval task name |
|       `links` | `external_instance_link` | Redirection link in the [Pending] or [Approved], used to redirect to the third-party systemEither the pc_link field or mobile_link field must be filled in to specify the link for redirection, which is not affected by the platform. |
|         `pc_link` | `string` | Redirection link for PC, which is used for redirection when users use Lark on PC |
|         `mobile_link` | `string` | Redirection link for mobile, which is used for redirection when users use Lark on mobile |
|       `status` | `string` | Task status **Optional values are**:  - `PENDING`: Under review - `APPROVED`: The task is approved - `REJECTED`: The task is rejected - `TRANSFERRED`: The task is transferred - `DONE`: The task is approved but hasn't been processed by the approver. The approver can't see the task. You can CC this individual to make the task visible.  |
|       `extra` | `string` | Extend json, the reason for the end of the task needs to be passed complete_reason field. Enumeration values and corresponding descriptions: - approved: agreed - rejected: rejected - node_auto_reject: Automatic rejection (due to logical judgment) - specific_rollback: return (including return to the originator, return to any intermediate approver) - add: and sign (add a new approver, approve with me) - add_pre: Pre-approver (add new approver, approve before me) - add_post: post-approver (add new approver, approve after me) - delete_assignee: Reductions - forward_resign: transfer (transfer to someone else for approval) - recall: revocation (withdrawal of documents, invalidation of documents) - delete: delete the approval slip - admin_forward: Administrator in the background operation transfer - system_forward: the system automatically transfers - auto_skip: Auto Pass - manual_skip: manually skip - submit_again: resubmit tasks - restart: restart the process - others: others (as the bottom line) |
|       `create_time` | `string` | Task creation time (Unix timestamp in ms) |
|       `end_time` | `string` | Task completion time: The value is 0 for uncompleted approvals (Unix timestamp in ms) |
|       `update_time` | `string` | Last update time of the task, unix millisecond timestamp, used to push data version control. The update policy is the same as the update_time in instance. |
|       `action_context` | `string` | Action context, which is taken in the callback request to pass the context data of the task as users perform actions |
|       `action_configs` | `action_config[]` | Task-level action configuration. Quick approval is available on mobile. |
|         `action_type` | `string` | Action type. Each task can be configured with 2 actions, which will be displayed in the approval list. When users perform actions, the callback request will carry this field, indicating whether the user has approved or rejected the action. **Optional values are**:  - `APPROVE`: Approved - `REJECT`: Rejected - `{KEY}`: Arbitrary string. The action name is required when arbitrary string is used.  |
|         `action_name` | `string` | Action name. The i18n key is used for frontend display. If the action_type field value is neither APPROVAL nor REJECT, this field is required to display a specific action name. |
|         `is_need_reason` | `boolean` | Whether the comments are required. If true, users will be redirected to the page to fill in comments when they perform actions. |
|         `is_reason_required` | `boolean` | Whether the approval comments are required. |
|         `is_need_attachment` | `boolean` | Whether attachments can be uploaded for the comments |
|       `display_method` | `string` | The way to open the approval task on the list page.Optional values are **Optional values are**:  - `BROWSER`: Opens by redirecting to the default system browser - `SIDEBAR`: Opens in the sidebar of Lark - `NORMAL`: Opens in the page embedded in Lark - `TRUSTEESHIP`: Open in managed mode  |
|       `exclude_statistics` | `boolean` | Tripartite mission support is not included in efficiency statistics. false: Include efficiency statistics. true: not included in efficiency statistics |
|       `node_id` | `string` | Node id: must meet both - In a process, each node id is unique. For example, the Node_id of each node such as "direct manager" and "next-level superior" in a process are different - In the same process definition, the same node in different approval instances should Node_id remain unchanged. For example, Zhang San and Li Si initiated leave applications respectively, and the node_id of the "direct manager" node in these two approval instances should remain the same |
|       `node_name` | `string` | Node names, such as "Financial Approval" and "Legal Approval", support three languages: Chinese, English and Japanese. Example: i18n@name. You need to pass the internationalized copy corresponding to the name in the i18n_resources |
|     `cc_list` | `cc_node[]` | CC list |
|       `cc_id` | `string` | Unique ID in the approval instance |
|       `user_id` | `string` | CC employee ID |
|       `open_id` | `string` | Either the CC sender's open ID or user ID is required. |
|       `links` | `external_instance_link` | Redirection link, used for redirection in the [CC'd] listEither the pc_link field or mobile_link field must be filled in to specify the link for redirection, which is not affected by the platform. |
|         `pc_link` | `string` | Redirection link for PC, which is used for redirection when users use Lark on PC |
|         `mobile_link` | `string` | Redirection link for mobile, which is used for redirection when users use Lark on mobile |
|       `read_status` | `string` | Read status. An empty cell indicates that the read/unread status is unavailable: **Optional values are**:  - `READ`: Read - `UNREAD`: Unread  |
|       `extra` | `string` | Extend json |
|       `title` | `string` | CC task name |
|       `create_time` | `string` | CC initiation time (Unix timestamp in ms) |
|       `update_time` | `string` | Last update time of CC, unix millisecond timestamp, used to push data version control.The update policy is the same as the update_time in instance. |
|       `display_method` | `string` | The way to open the approval task on the list page.Optional values are: **Optional values are**:  - `BROWSER`: Opens by redirecting to the default system browser - `SIDEBAR`: Opens in the sidebar of Lark - `NORMAL`: Opens in the page embedded in Lark - `TRUSTEESHIP`: Open in managed mode  |
|     `i18n_resources` | `i18n_resource[]` | Internationalized text |
|       `locale` | `string` | Language **Optional values are**:  - `zh-CN`: Chinese - `en-US`: English - `ja-JP`: Japanese  |
|       `texts` | `i18n_resource_text[]` | Key, value, and i18n key of copy starts with @i18n@. This field is mainly used for internationalized text. Word order users can pass texts in multiple languages at the same time, and the Approval Center will use the corresponding text according to the user's language environment. If the text of the user's language environment isn't passed, the text of default language will be used. |
|         `key` | `string` | Copywriting key |
|         `value` | `string` | Copywriting |
|       `is_default` | `boolean` | Whether to use the default language. If the default language is used, all keys should be contained; if a non-default language is used but its keys do not exist, the default language will be used instead. |
|     `trusteeship_url_token` | `string` | Document escrow authentication token, escrow callback will be accompanied by this token to help business party authentication |
|     `trusteeship_user_id_type` | `string` | The type of user will affect the selection of the user identification field of the request parameter, including the target user returned by the signature operation. Currently only "user_id" is supported |
|     `trusteeship_urls` | `trusteeship_urls` | The URL address of the interface of the document hosting callback access party |
|       `form_detail_url` | `string` | Get the url address of the data related to the form schema |
|       `action_definition_url` | `string` | Represents the url address to obtain the data of the approval operation area |
|       `approval_node_url` | `string` | Get the url address of the data related to the approval record |
|       `action_callback_url` | `string` | The url address of the callback when performing an approval operation |
|       `pull_business_data_url` | `string` | Obtain the URL address of the managed dynamic data. When using this interface, you must ensure that the interface address is synchronized in the data of the historical managed document. If there is no interface in the historical document, you need to re-synchronize the data of the historical managed document to update the URL. This interface is used for the interactive use of the Lark approval front end and the line of business. Only the specific components of the approval front end (the components provided by the Lark approval front end and the components that need to interface with the line of business) will be required | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
        "data": {
            "approval_code": "81D31358-93AF-92D6-7425-01A5D67C4E71",
            "status": "PENDING",
            "extra": "{\"xxx\":\"xxx\"}",
            "instance_id": "24492654",
            "links": {
                "pc_link": "https://applink.larksuite.com/client/mini_program/open?mode=appCenter&appId=cli_9c90fc38e07a9101&path=pc/pages/detail?id=1234",
                "mobile_link": "https://applink.larksuite.com/client/mini_program/open?appId=cli_9c90fc38e07a9101&path=pages/detail?id=1234"
            },
            "title": "@i18n@1",
            "form": [
                {
                    "name": "@i18n@2",
                    "value": "@i18n@3"
                }
            ],
            "user_id": "a987sf9s",
            "user_name": "@i18n@9",
            "open_id": "ou_be73cbc0ee35eb6ca54e9e7cc14998c1",
            "department_id": "od-8ec33278bc2",
            "department_name": "@i18n@10",
            "start_time": "1556468012678",
            "end_time": "1556468012678",
            "update_time": "1556468012678",
            "display_method": "BROWSER",
            "update_mode": "UPDATE",
            "task_list": [
                {
                    "task_id": "112534",
                    "user_id": "a987sf9s",
                    "open_id": "ou_be73cbc0ee35eb6ca54e9e7cc14998c1",
                    "title": "i18n1",
                    "links": {
                        "pc_link": "https://applink.larksuite.com/client/mini_program/open?mode=appCenter&appId=cli_9c90fc38e07a9101&path=pc/pages/detail?id=1234",
                        "mobile_link": "https://applink.larksuite.com/client/mini_program/open?appId=cli_9c90fc38e07a9101&path=pages/detail?id=1234"
                    },
                    "status": "PENDING",
                    "extra": "{\"xxx\":\"xxx\",\"complete_reason\":\"approved\"}",
                    "create_time": "1556468012678",
                    "end_time": "1556468012678",
                    "update_time": "1556468012678",
                    "action_context": "123456",
                    "action_configs": [
                        {
                            "action_type": "APPROVE",
                            "action_name": "@i18n@5",
                            "is_need_reason": false,
                            "is_reason_required": false,
                            "is_need_attachment": false
                        }
                    ],
                    "display_method": "BROWSER",
                    "exclude_statistics": false,
                    "node_id": "node",
                    "node_name": "i18n@name"
                }
            ],
            "cc_list": [
                {
                    "cc_id": "123456",
                    "user_id": "12345",
                    "open_id": "ou_be73cbc0ee35eb6ca54e9e7cc14998c1",
                    "links": {
                        "pc_link": "https://applink.larksuite.com/client/mini_program/open?mode=appCenter&appId=cli_9c90fc38e07a9101&path=pc/pages/detail?id=1234",
                        "mobile_link": "https://applink.larksuite.com/client/mini_program/open?appId=cli_9c90fc38e07a9101&path=pages/detail?id=1234"
                    },
                    "read_status": "READ",
                    "extra": "{\"xxx\":\"xxx\"}",
                    "title": "xxx",
                    "create_time": "1556468012678",
                    "update_time": "1556468012678",
                    "display_method": "BROWSER"
                }
            ],
            "i18n_resources": [
                {
                    "locale": "zh-CN",
                    "texts": [
                        {
                            "key": "@i18n@1",
                            "value": "People"
                        }
                    ],
                    "is_default": true
                }
            ],
            "trusteeship_url_token": "788981c886b1c28ac29d1e68efd60683d6d90dfce80938ee9453e2a5f3e9e306",
            "trusteeship_user_id_type": "user_id",
            "trusteeship_urls": {
                "form_detail_url": "https://#{your_domain}/api/form_detail",
                "action_definition_url": "https://#{your_domain}/api/action_definition",
                "approval_node_url": "https://#{your_domain}/api/approval_node",
                "action_callback_url": "https://#{your_domain}/api/action_callback",
                "pull_business_data_url": "https://#{your_domain}/api/pull_business_data"
            }
        }
    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 1390001 | param is invalid | Parameter error |
| 400 | 1395001 | There have been some errors. Please try again later | Service error encountered. Troubleshooting steps: 1. Refer to the parameter descriptions in the API documentation and check whether the parameters passed in the request are correct. If form parameters are passed (form), ensure the form control data is correct. 2. Lower the request frequency, do not synchronize the same approval instance within a short period, please try again later. If the error persists after retrying, please contact Technical Support. |
| 400 | 1390014 | tenant_id not found | No designated members found. |
