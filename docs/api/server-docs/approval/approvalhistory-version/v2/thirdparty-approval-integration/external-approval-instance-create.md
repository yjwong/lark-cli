---
title: "External Approval Instance Create"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uczM3UjL3MzN14yNzcTN"
method: "POST"
api_path: "https://www.larksuite.com/approval/openapi/v2/external/instance/create"
section: "Approval"
scopes:
  - "approval:approval:readonly"
updated: "1658309965000"
---

# Sync of third-party approval instances
> This API has been upgraded to enhance its security. We recommend that you migrate to the new version>>

Approval transfer lies in the third-party system instead of the Approval Center. The approval instances, approval tasks, and approval CC data generated after approval transfer are synced from the third-party system to the Approval Center.

Users can browse the instances, tasks, and CC information synced from the third-party system in the Approval Center, and can redirect to the third-party system for more detailed viewing and action. The instance information is in the [Initiated] list, the task information is in the [Pending] and [Approved] lists, and the CC information is in the [CC'd] list.

The third-party system can also configure the callback API for approval tasks, so that the approver can directly perform the approval action in the Approval Center, and the Approval Center will call back the third-party system. After receiving the callback, the third-party system will update the task information and sync the new task information back to the Approval Center to form a closed loop.

## Request
| HTTP URL | https://www.larksuite.com/approval/openapi/v2/external/instance/create |
| --- | --- |
| HTTP Method | POST |
| Supported app types | custom,isv |
| Required scopes | `approval:approval:readonly` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" Learn more about how to access and use access_token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Request body

| `content` `Required` | `string` | JSON string |
| --- | --- | --- | **Third-party approval JSON**
| `approval_code` `Required` | `string` | Approval definition code, which creates the value returned by the approval definition to indicate which process the instance belongs to. This field affects the title of the instance in the list, which is taken from the name field of the corresponding definition. | 81D31358-93AF-92D6-7425-01A5D67C4E71 |
| --- | --- | --- | --- |
| `instance_id` `Required` | `string` | The unique ID of approval instance defined by users. It is required to be unique under tenant and app. | 24492654 |
| `status` `Required` | `string` | Approval instance status **Optional values are:** - `PENDING`: Under review - `APPROVED`: The request is approved when the approval process ends. - `REJECTED`: The request is rejected when the approval process ends. - `CANCELED`: The approval is canceled by the initiator. - `DELETED`: The approval is deleted. - `HIDDEN`: Status is hidden (i.e. not displayed). | PENDING |
| `extra` | `string` | Approval instance extension JSON | {"xxx":"xxx"} |
| `links` `Required` | `map` | Approval instance link collection, which is used to redirect to the third-party system from the [Initiated] list. Either the pc_link field or mobile_link field must be filled in to specify the link for redirection, which is not affected by the platform. |  |
| ∟ `pc_link ``Required` | `string` | Redirection link for PC, which is used for redirection when users use Lark on PC | https://applink.larksuite.com/client/mini_program/open?mode=appCenter&appId=cli_9c9cfc38e07a9101&path=pc/pages/detail?id=1234 |
| ∟ `mobile_link``required` | `string` | Redirection link for mobile, which is used for redirection when users use Lark on mobile | https://applink.larksuite.com/client/mini_program/open?appId=cli_9c9cfc38e07a9101&path=pages/detail?id=1234 |
| `title` | `string` | Approval display name. If this field is filled in, the approval title in the approval list will use this field. If this field isn't filled in, the approval title will use the name of the approval definition. | @i18n@1 |
| `form` | `list` | Form data filled in when a user submits application, which is displayed in all the approval lists. More than one value can be passed. However, in the Approval Center, the first 2 values are displayed on PC, while the first 3 values on mobile. The display length can't exceed 2,048 characters. | [ { "name": "@i18n@2", "value": "@i18n@3" } ] |
| ∟ `name` | `string` | Form field name | @i18n@2 |
| ∟ `value` | `string` | Form value | @i18n@3 |
| `user_id``` | `string` | Approver initiator's user_id. The initiator can view all initiated approvals in the [Initiated] list.In the [Pending], [Approved], and [CC'd] lists, this field shows the initiator of the approval. Either the initiator's open ID or user ID is required. | a987sf9s |
| `user_name` | `string` | Approval initiator's user name. If the initiator isn't a real user (for example, a department) and has no user ID, use this field to pass the name. | @i18n@9 |
| `open_id` | `string` | Either the initiator's open ID or user ID is required. | ou_be73cbc0ee35eb6ca54e9e7cc14998c1 |
| `department_id` | `string` | Initiator department, used in the list to display the department which the initiator belongs to. This field isn't displayed if it isn't passed. If the user hasn't joined any department, the tenant name is displayed when "" is passed. If the department_name field is passed, the department name is displayed. | od-8ec33278bc2 |
| `department_name` | `string` | Approval initiator's department. If the initiator isn't a real user (for example, a department) and has no department_id, use this field to pass the name. | @i18n@10 |
| `start_time` `Required` | `long` | Approval initiation time (Unix timestamp in ms) | 1556468012678 |
| `end_time` `Required` | `long` | Approval instance end time: The value is 0 for uncompleted approvals (Unix timestamp in ms) | 1556468012678 |
| `update_time` `Required` | `long` | Last update time of the approval instance, used to push data version control. If the update_mode value is UPDATE, the approval instance information in the Approval Center will only be updated when the passed update_time changes (becomes larger). This field is mainly used to avoid updating new data with old data when concurrency occurs. | 1556468012678 |
| `display_method` | `string` | The way to open the approval instance on the list page. **Optional values are:** - `BROWSER`: Opens by redirecting to the default system browser - `SIDEBAR`: Opens in the sidebar of Lark - `NORMAL`: Opens in the page embedded in Lark  | BROWSER |
| `update_mode` | `string` | Update method **Optional values are:** - `REPLACE`: Replaces all. The value is set by default. - `UPDATE`: Updates incrementally. When update_mode=REPLACE, each time the currently pushed data is taken as the final data, and the redundant tasks and CC data in the Approval Center will be deleted (not in the data pushed this time). When update_mode=UPDATE, the data in the Approval Center won't be deleted, but only the instance and task data will be added and updated. | UPDATE |
| `task_list` `Required` | `list` | Task list | Notes: The maximum number of tasks is 200. |
| ∟ `task_id` `Required` | `string` | The unique ID within the approval instance, which is used to locate data when updating approval tasks | 112534 |
| ∟ `user_id` `Required` | `string` | Approver's user_id. The task will appear in the [Pending] or [Approved] list of the approver. | a987sf9s |
| ∟ `open_id` | `string` | Either the approver's open ID or user ID is required. | ou_be73cbc0ee35eb6ca54e9e7cc14998c1 |
| ∟ `title` | `string` | Approval task name | i18n1 |
| ∟ `links``required` | `map` | Redirection link in the [Pending] or [Approved], used to redirect to the third-party system Either the pc_link field or mobile_link field must be filled in to specify the link for redirection, which is not affected by the platform. |  |
| ∟ `pc_link``required` | `string` | Redirection link for PC, which is used for redirection when users use Lark on PC | https://applink.larksuite.com/client/mini_program/open?mode=appCenter&appId=cli_9c9cfc38e07a9101&path=pc/pages/detail?id=1234 |
| ∟ `mobile_link``required` | `string` | Redirection link for mobile, which is used for redirection when users use Lark on mobile | https://applink.larksuite.com/client/mini_program/open?appId=cli_9c9cfc38e07a9101&path=pages/detail?id=1234 |
| ∟ `status` `Required` | `string` | Task status **Optional values are:** - `PENDING`: Under review - `APPROVED`: The task is approved - `REJECTED`: The task is rejected - `TRANSFERRED`: The task is transferred - `DONE`: The task is approved but hasn't been processed by the approver. The approver can't see the task. You can CC this individual to make the task visible. |  |
| ∟ `extra` | `string` | Extension JSON |  |
| ∟ `create_time` `Required` | `long` | Task creation time (Unix timestamp in ms) | 1556468012678 |
| ∟ `end_time` `Required` | `long` | Task completion time: The value is 0 for uncompleted approvals (Unix timestamp in ms) | 1556468012678 |
| ∟ `update_time` `required` | `long` | Last update time of the task, used to push data version control. The update policy is the same as the update_time in instance. | 1556468012678 |
| ∟ `action_context` | `string` | Action context, which is taken in the callback request to pass the context data of the task as users perform actions | 123456 |
| ∟ `action_configs` | `list` | Task-level action configuration. Quick approval is available on mobile. |  |
| ∟ ` action_type ` `Required` | `string` | Action type. Each task can be configured with 2 actions, which will be displayed in the approval list. When users perform actions, the callback request will carry this field, indicating whether the user has approved or rejected the action. **APPROVE** - Approved **REJECT** - Rejected **{KEY}** - Arbitrary string. The action name is required when arbitrary string is used. | APPROVE |
| ∟ ` action_name ` | `string` | Action name. The i18n key is used for frontend display. If the action_type field value is neither APPROVAL nor REJECT, this field is required to display a specific action name. | @i18n@5 |
| ∟ ` is_need_reason  ` | `bool` | Whether the comments are required. If true, users will be redirected to the page to fill in comments when they perform actions. | false |
| ∟ ` is_reason_required  ` | `bool` | Whether the approval comments are required. | false |
| ∟ ` is_need_attachment   ` | `bool` | Whether attachments can be uploaded for the comments | false |
| ∟ `display_method` | `string` | The way to open the approval task on the list page. **Optional values are:** - `BROWSER`: Opens by redirecting to the default system browser - `SIDEBAR`: Opens in the sidebar of Lark - `NORMAL`: Opens in the page embedded in Lark  | BROWSER |
| `cc_list` | `list` | CC list | Notes: The maximum number of CCs is 200. |
| ∟ `cc_id` `Required` | `string` | Unique ID in the approval instance | 12345 |
| ∟ `user_id` `Required` | `string` | CC employee ID | 12345 |
| ∟ `open_id` | `string` | Either the CC sender's open ID or user ID is required. | ou_be73cbc0ee35eb6ca54e9e7cc14998c1 |
| ∟ `links` `Required` | `map` | Redirection link, used for redirection in the [CC'd] list Either the pc_link field or mobile_link field must be filled in to specify the link for redirection, which is not affected by the platform. |  |
| ∟  `pc_link` | `string` | Redirection link for PC, which is used for redirection when users use Lark on PC | https://applink.larksuite.com/client/mini_program/open?mode=appCenter&appId=cli_9c9cfc38e07a9101&path=pc/pages/detail?id=1234 |
| ∟  `mobile_link` | `string` | Redirection link for mobile, which is used for redirection when users use Lark on mobile | https://applink.larksuite.com/client/mini_program/open?appId=cli_9c9cfc38e07a9101&path=pages/detail?id=1234 |
| ∟ `read_status` `Required` | `string` | Read status. An empty cell indicates that the read/unread status is unavailable: **Optional values are:** - `READ`: Read - `UNREAD`: Unread | READ |
| ∟ `extra` | `string` | Extension JSON | {"xxx":"xxx"} |
| ∟ `title` | `string` | CC task name | xxx |
| ∟ `create_time``required` | `long` | CC initiation time (Unix timestamp in ms) | 1556468012678 |
| ∟ `update_time``required` | `long` | Last update time of CC, used to push data version control. The update policy is the same as the update_time in instance. | 1556468012678 |
| ∟ `display_method` | `string` | The way to open the approval task on the list page. **Optional values are:** - `BROWSER`: Opens by redirecting to the default system browser - `SIDEBAR`: Opens in the sidebar of Lark - `NORMAL`: Opens in the page embedded in Lark  | BROWSER |
| `i18n_resources``required` | `list` | Internationalized text |  |
| ∟ `locale``required` | `string` | Language - `zh-CN`: Chinese - `en-US`: English - `ja-JP`: Japanese | zh-CN |
| ∟ `is_default``required` | `bool` | Whether the default language | true |
| ∟ `texts``required` | `map` | Key, value, and i18n key of copy starts with  @i18n@. This field is mainly used for internationalized text. Word order users can pass texts in multiple languages at the same time, and the Approval Center will use the corresponding text according to the user's language environment. If the text of the user's language environment isn't passed, the text of default language will be used. | { "@i18n@1": "Permission request",      "@i18n@2": "OA approval", "@i18n@3": "Permission" } | Notes:
- It is required to ensure the uniqueness of various unique IDs in the approval instance, even if they don't belong to the same entity.

```json
{
    "content": "{\"approval_code\":\"60D0F7A2-052D-49E9-B570-C29C836CDF8E\",\"instance_id\":\"216263\",\"status\":\"PENDING\",\"extra\":\"\",\"links\":{\"pc_link\":\"http://182.50.118.46:38080/spa/workflow/static4form/index.html?#/main/workflow/req?requestid=\",\"mobile_link\":\"http://182.50.118.46:38080/spa/workflow/static4form/index.html?#/main/workflow/req?requestid=\"},\"title\":\"@i18n@1\",\"form\":[{\"name\":\"@i18n@2\",\"value\":\"@i18n@3\"}],\"user_id\":\"52d6585f\",\"user_name\":\"342a\",\"open_id\":\"123\",\"department_id\":\"od-8ec33ffec336c3a39a278bc25e931676\",\"department_name\":\"XXXX\",\"start_time\":1622444502000,\"update_time\":1622444502000,\"end_time\":0,\"update_mode\":\"REPLACED\",\"task_list\":[{\"task_id\":\"112253\",\"user_id\":\"52d6585f\",\"links\":{\"pc_link\":\"http://\",\"mobile_link\":\"http://\"},\"status\":\"PENDING\",\"extra\":\"\",\"title\":\"Approve\",\"create_time\":1566468012678,\"end_time\":0,\"update_time\":1566468012678,\"action_context\":\"123456\",\"action_configs\":[{\"action_type\":\"APPROVE\",\"action_name\":\"@i18n@9\",\"is_need_reason\":true,\"is_reason_required\":true,\"is_need_attachment\":true}]}],\"cc_list\":[{\"cc_id\":\"1231243\",\"user_id\":\"a987sf9s\",\"open_id\":\"\",\"links\":{\"pc_link\":\"http://\",\"mobile_link\":\"http://\"},\"read_status\":\"READ\",\"extra\":\"\",\"title\":\"XXX\",\"create_time\":1556468012678,\"update_time\":1566468012678}],\"i18n_resources\":[{\"locale\":\"zh-CN\",\"is_default\":true,\"texts\":{\"@i18n@1\":\"Leave\",\"@i18n@2\":\"Day\",\"@i18n@3\":\"2020-08-01\"}},{\"locale\":\"en-US\",\"is_default\":false,\"texts\":{\"@i18n@1\":\"Leave\",\"@i18n@2\":\"Day\",\"@i18n@3\":\"2020-08-01\"}},{\"locale\":\"ja-JP\",\"is_default\":false,\"texts\":{\"@i18n@1\":\"Leave\",\"@i18n@2\":\"Day\",\"@i18n@3\":\"2020-08-01\"}}]}"
}
```

## Response

### Response body

|Parameter|Type|Required|Description|
|-|-|-|-|
|code|int|Return code. A value other than 0 indicates failure.|
|msg|string|Yes|Return code description|
|data | map |Yes| Returned business information | ### Response body example

```json
{
    "code": 0,
    "msg": "success",
    "data": {
    }
}
```
