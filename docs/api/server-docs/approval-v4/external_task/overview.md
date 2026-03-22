---
title: "Resource introduction"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/approval-v4/external_task/overview"
service: "approval-v4"
resource: "external_task"
section: "Approval"
updated: "1675167425000"
---

# Resource introduction

Each approval operation of the approver corresponds to an approval task

| Parameter | Type | Description |
| --- | --- | --- |
| `task_id` | `string` | The unique ID within the approval instance, which is used to locate data when updating approval tasks **Example value**："112534" |
| `user_id` | `string` | Approver's user_id. The task will appear in the [Pending] or [Approved] list of the approver. **Example value**："a987sf9s" |
| `open_id` | `string` | Either the approver's open ID or user ID is required. **Example value**："ou_be73cbc0ee35eb6ca54e9e7cc14998c1" |
| `title` | `string` | Approval task name **Example value**："i18n1" |
| `links` | `external_instance_link` | Redirection link in the [Pending] or [Approved], used to redirect to the third-party systemEither the pc_link field or mobile_link field must be filled in to specify the link for redirection, which is not affected by the platform. |
|   `pc_link` | `string` | **Example value**："https://applink.larksuite.com/client/mini_program/open?mode=appCenter&appId=cli_9c90fc38e07a9101&path=pc/pages/detail?id=1234" |
|   `mobile_link` | `string` | Redirection link for mobile, which is used for redirection when users use Lark on mobile **Example value**："https://applink.larksuite.com/client/mini_program/open?appId=cli_9c90fc38e07a9101&path=pages/detail?id=1234" |
| `status` | `string` | Task status **Example value**："PENDING待审批" **Optional values are**：  - `PENDING`: Under review - `APPROVED`: The task is approved - `REJECTED`: The task is rejected - `TRANSFERRED`: The task is transferred - `DONE`: The task is approved but hasn't been processed by the approver. The approver can't see the task. You can CC this individual to make the task visible.  |
| `extra` | `string` | Extend json **Example value**："{\"xxx\":\"xxx\"}" |
| `create_time` | `string` | Task creation time (Unix timestamp in ms) **Example value**："1556468012678" |
| `end_time` | `string` | Task completion time: The value is 0 for uncompleted approvals (Unix timestamp in ms) **Example value**："1556468012678" |
| `update_time` | `string` | Last update time of the task, used to push data version control.The update policy is the same as the update_time in instance. **Example value**："1556468012678" |
| `action_context` | `string` | Action context, which is taken in the callback request to pass the context data of the task as users perform actions **Example value**："123456" |
| `action_configs` | `action_config[]` | Task-level action configuration. Quick approval is available on mobile. |
|   `action_type` | `string` | Action type. Each task can be configured with 2 actions, which will be displayed in the approval list. When users perform actions, the callback request will carry this field, indicating whether the user has approved or rejected the action. **Example value**："APPROVE" **Optional values are**：  - `APPROVE`: Approved - `REJECT`: Rejected - `{KEY}`: Arbitrary string. The action name is required when arbitrary string is used.  |
|   `action_name` | `string` | Action name. The i18n key is used for frontend display. If the action_type field value is neither APPROVAL nor REJECT, this field is required to display a specific action name. **Example value**："@i18n@5" |
|   `is_need_reason` | `boolean` | Whether the comments are required. If true, users will be redirected to the page to fill in comments when they perform actions. **Example value**：false |
|   `is_reason_required` | `boolean` | Whether the approval comments are required. **Example value**：false |
|   `is_need_attachment` | `boolean` | Whether attachments can be uploaded for the comments **Example value**：false |
| `display_method` | `string` | The way to open the approval task on the list page.Optional values are **Example value**："BROWSER" **Optional values are**：  - `BROWSER`: Opens by redirecting to the default system browser - `SIDEBAR`: Opens in the sidebar of Lark - `NORMAL`: Opens in the page embedded in Lark  | ## Data example
```json
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
```

## User ID description
For how to distinguish between and use user_id, open_id, and union_id, refer to User-related IDs.
