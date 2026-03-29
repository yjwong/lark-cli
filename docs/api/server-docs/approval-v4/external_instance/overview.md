---
title: "Resource introduction"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/approval-v4/external_instance/overview"
service: "approval-v4"
resource: "external_instance"
section: "Approval"
updated: "1675167421000"
---

# Resource introduction
The approval flow generated when an employee initiates an approval. Includes information such as multiple approval tasks, approval cc, etc

| Parameter | Type | Description |
| --- | --- | --- |
| `approval_code` | `string` | Approval definition code, which creates the value returned by the approval definition to indicate which process the instance belongs to. This field affects the title of the instance in the list, which is taken from the name field of the corresponding definition. **Example value**："81D31358-93AF-92D6-7425-01A5D67C4E71" |
| `status` | `string` | Approval instance status **Example value**："PENDING" **Optional values are**：  - `PENDING`: Under review - `APPROVED`: The request is approved when the approval process ends. - `REJECTED`: The request is rejected when the approval process ends. - `CANCELED`: The approval is canceled by the initiator. - `DELETED`: he approval is deleted. - `HIDDEN`: Status is hidden (i.e. not displayed).  |
| `extra` | `string` | Approval instance extension JSON **Example value**："{\"xxx\":\"xxx\"}" |
| `instance_id` | `string` | The unique ID of approval instance defined by users. It is required to be unique under tenant and app. **Example value**："24492654" |
| `links` | `external_instance_link` | Approval instance link collection, which is used to redirect to the third-party system from the [Initiated] list.Either the pc_link field or mobile_link field must be filled in to specify the link for redirection, which is not affected by the platform. |
|   `pc_link` | `string` | Redirection link for PC, which is used for redirection when users use Lark on PC **Example value**："https://applink.larksuite.com/client/mini_program/open?mode=appCenter&appId=cli_9c90fc38e07a9101&path=pc/pages/detail?id=1234" |
|   `mobile_link` | `string` | Redirection link for mobile, which is used for redirection when users use Lark on mobile **Example value**："https://applink.larksuite.com/client/mini_program/open?appId=cli_9c90fc38e07a9101&path=pages/detail?id=1234" |
| `title` | `string` | Approval display name. If this field is filled in, the approval title in the approval list will use this field. If this field isn't filled in, the approval title will use the name of the approval definition. **Example value**："@i18n@1" |
| `form` | `external_instance_form[]` | Form data filled in when a user submits application, which is displayed in all the approval lists. More than one value can be passed. However, in the Approval Center, the first 2 values are displayed on PC, while the first 3 values on mobile. The display length can't exceed 2,048 characters. |
|   `name` | `string` | Form field name **Example value**："@i18n@2" |
|   `value` | `string` | Approver initiator's user_id. The initiator can view all initiated approvals in the [Initiated] list.In the [Pending], [Approved], and [CC'd] lists, this field shows the initiator of the approval. Either the initiator's open ID or user ID is required. **Example value**："a987sf9s" |
| `user_id` | `string` | Approver initiator's user_id. The initiator can view all initiated approvals in the [Initiated] list.In the [Pending], [Approved], and [CC'd] lists, this field shows the initiator of the approval. Either the initiator's open ID or user ID is required. **Example value**："a987sf9s" |
| `user_name` | `string` | Approval initiator's user name. If the initiator isn't a real user (for example, a department) and has no user ID, use this field to pass the name. **Example value**："@i18n@9" |
| `open_id` | `string` | Either the initiator's open ID or user ID is required. **Example value**："ou_be73cbc0ee35eb6ca54e9e7cc14998c1" |
| `department_id` | `string` | Initiator department, used in the list to display the department which the initiator belongs to.This field isn't displayed if it isn't passed.If the user hasn't joined any department, the tenant name is displayed when "" is passed.If the department_name field is passed, the department name is displayed. **Example value**："od-8ec33278bc2" |
| `department_name` | `string` | Approval initiator's department. If the initiator isn't a real user (for example, a department) and has no department_id, use this field to pass the name. **Example value**："@i18n@10" |
| `start_time` | `string` | Approval initiation time (Unix timestamp in ms) **Example value**："1556468012678" |
| `end_time` | `string` | Approval instance end time: The value is 0 for uncompleted approvals (Unix timestamp in ms) **Example value**："1556468012678" |
| `update_time` | `string` | Last update time of the approval instance, used to push data version control.If the update_mode value is UPDATE, the approval instance information in the Approval Center will only be updated when the passed update_time changes (becomes larger).This field is mainly used to avoid updating new data with old data when concurrency occurs. **Example value**："1556468012678" |
| `display_method` | `string` | The way to open the approval instance on the list page.Optional values are: **Example value**："BROWSER" **Optional values are**：  - `BROWSER`: Opens by redirecting to the default system - `SIDEBAR`: Opens in the sidebar of Lark - `NORMAL`: Opens in the page embedded in Lark  |
| `update_mode` | `string` | Update method, When update_mode=REPLACE, each time the currently pushed data is taken as the final data, and the redundant tasks and CC data in the Approval Center will be deleted (not in the data pushed this time).When update_mode=UPDATE, the data in the Approval Center won't be deleted, but only the instance and task data will be added and updated. **Example value**："UPDATE" **Optional values are**：  - `REPLACE`: Replaces all. The value is set by default. - `UPDATE`: Updates incrementally.  |
| `task_list` | `external_instance_task_node[]` | Task list **Data validation rules**： - Maximum length：`200` |
|   `task_id` | `string` | The unique ID within the approval instance, which is used to locate data when updating approval tasks **Example value**："112534" |
|   `user_id` | `string` | Approver's user_id. The task will appear in the [Pending] or [Approved] list of the approver. **Example value**："a987sf9s" |
|   `open_id` | `string` | Either the approver's open ID or user ID is required. **Example value**："ou_be73cbc0ee35eb6ca54e9e7cc14998c1" |
|   `title` | `string` | Approval task name **Example value**："i18n1" |
|   `links` | `external_instance_link` | Redirection link in the [Pending] or [Approved], used to redirect to the third-party systemEither the pc_link field or mobile_link field must be filled in to specify the link for redirection, which is not affected by the platform. |
|     `pc_link` | `string` | **Example value**："https://applink.larksuite.com/client/mini_program/open?mode=appCenter&appId=cli_9c90fc38e07a9101&path=pc/pages/detail?id=1234" |
|     `mobile_link` | `string` | Redirection link for mobile, which is used for redirection when users use Lark on mobile **Example value**："https://applink.larksuite.com/client/mini_program/open?appId=cli_9c90fc38e07a9101&path=pages/detail?id=1234" |
|   `status` | `string` | Task status **Example value**："PENDING待审批" **Optional values are**：  - `PENDING`: Under review - `APPROVED`: The task is approved - `REJECTED`: The task is rejected - `TRANSFERRED`: The task is transferred - `DONE`: The task is approved but hasn't been processed by the approver. The approver can't see the task. You can CC this individual to make the task visible.  |
|   `extra` | `string` | Extend json **Example value**："{\"xxx\":\"xxx\"}" |
|   `create_time` | `string` | Task creation time (Unix timestamp in ms) **Example value**："1556468012678" |
|   `end_time` | `string` | Task completion time: The value is 0 for uncompleted approvals (Unix timestamp in ms) **Example value**："1556468012678" |
|   `update_time` | `string` | Last update time of the task, used to push data version control.The update policy is the same as the update_time in instance. **Example value**："1556468012678" |
|   `action_context` | `string` | Action context, which is taken in the callback request to pass the context data of the task as users perform actions **Example value**："123456" |
|   `action_configs` | `action_config[]` | Task-level action configuration. Quick approval is available on mobile. |
|     `action_type` | `string` | Action type. Each task can be configured with 2 actions, which will be displayed in the approval list. When users perform actions, the callback request will carry this field, indicating whether the user has approved or rejected the action. **Example value**："APPROVE" **Optional values are**：  - `APPROVE`: Approved - `REJECT`: Rejected - `{KEY}`: Arbitrary string. The action name is required when arbitrary string is used.  |
|     `action_name` | `string` | Action name. The i18n key is used for frontend display. If the action_type field value is neither APPROVAL nor REJECT, this field is required to display a specific action name. **Example value**："@i18n@5" |
|     `is_need_reason` | `boolean` | Whether the comments are required. If true, users will be redirected to the page to fill in comments when they perform actions. **Example value**：false |
|     `is_reason_required` | `boolean` | Whether the approval comments are required. **Example value**：false |
|     `is_need_attachment` | `boolean` | Whether attachments can be uploaded for the comments **Example value**：false |
|   `display_method` | `string` | The way to open the approval task on the list page.Optional values are **Example value**："BROWSER" **Optional values are**：  - `BROWSER`: Opens by redirecting to the default system browser - `SIDEBAR`: Opens in the sidebar of Lark - `NORMAL`: Opens in the page embedded in Lark  |
| `cc_list` | `cc_node[]` | CC list **Data validation rules**： - Maximum length：`200` |
|   `cc_id` | `string` | Unique ID in the approval instance **Example value**："123456" |
|   `user_id` | `string` | CC employee ID **Example value**："12345" |
|   `open_id` | `string` | Either the CC sender's open ID or user ID is required. **Example value**："ou_be73cbc0ee35eb6ca54e9e7cc14998c1" |
|   `links` | `external_instance_link` | Redirection link, used for redirection in the [CC'd] listEither the pc_link field or mobile_link field must be filled in to specify the link for redirection, which is not affected by the platform. |
|     `pc_link` | `string` | Redirection link for PC, which is used for redirection when users use Lark on PC **Example value**："https://applink.larksuite.com/client/mini_program/open?mode=appCenter&appId=cli_9c90fc38e07a9101&path=pc/pages/detail?id=1234" |
|     `mobile_link` | `string` | Redirection link for mobile, which is used for redirection when users use Lark on mobile **Example value**："https://applink.larksuite.com/client/mini_program/open?appId=cli_9c90fc38e07a9101&path=pages/detail?id=1234" |
|   `read_status` | `string` | Read status. An empty cell indicates that the read/unread status is unavailable: **Example value**："READ" **Optional values are**：  - `READ`: 已读 - `UNREAD`: 未读  |
|   `extra` | `string` | Extend json **Example value**："{\"xxx\":\"xxx\"}" |
|   `title` | `string` | CC task name **Example value**："title" |
|   `create_time` | `string` | CC initiation time (Unix timestamp in ms) **Example value**："1556468012678" |
|   `update_time` | `string` | Last update time of CC, used to push data version control.The update policy is the same as the update_time in instance. **Example value**："instance 的update_time" |
|   `display_method` | `string` | The way to open the approval task on the list page.Optional values are: **Example value**："BROWSER" **Optional values are**：  - `BROWSER`: Opens by redirecting to the default system browser - `SIDEBAR`: Opens in the sidebar of Lark - `NORMAL`: Opens in the page embedded in Lark  |
| `i18n_resources` | `i18n_resource[]` | Internationalized text |
|   `locale` | `string` | Language **Example value**："zh-CN" **Optional values are**：  - `zh-CN`: Chinese - `en-US`: English - `ja-JP`: Japanese  |
|   `texts` | `i18n_resource_text[]` | Key, value, and i18n key of copy starts with @i18n@.This field is mainly used for internationalized text. Word order users can pass texts in multiple languages at the same time, and the Approval Center will use the corresponding text according to the user's language environment. If the text of the user's language environment isn't passed, the text of default language will be used. |
|     `key` | `string` | \Copywriting key **Example value**："@i18n@1" |
|     `value` | `string` | Copywriting **Example value**："people" |
|   `is_default` | `boolean` | Whether to use the default language. If the default language is used, all keys should be contained; if a non-default language is used but its keys do not exist, the default language will be used instead. **Example value**：true | ## 数据示例
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
    "user_name": "zhangsan",
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
        "texts": [ {"key":"@i18n@1", "value":"test"}, {"key":"@i18n@2", "value":"day"},{"key":"@i18n@3", "value":"2022-07-06"}]
    }]
}
```

## Form description
Understand the definition of form， refer to Overview of Approval Definitions

## User ID description
For how to distinguish between and use user_id, open_id, and union_id, refer to User-related IDs.

## Department ID description
For how to distinguish between and use department_id,open_department_id,refer to        Department Overview
