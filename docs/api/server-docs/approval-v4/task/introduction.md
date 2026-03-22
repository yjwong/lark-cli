---
title: "Resource introduction"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/approval-v4/task/introduction"
service: "approval-v4"
resource: "task"
section: "Approval"
updated: "1675167405000"
---

# Resource introduction
An approval task can't exist alone without an approval node. Each approval node may include one or more approval tasks. The approvers of the current approval node are specified in each task. If the current node includes no more than one approver, there are multiple approval tasks, each of which corresponds to a specific approver. When a node includes multiple approval tasks, some of the tasks are not solely dependent on the approver's actions to change their status, but are affected by the status change of other tasks. (For example, if an anyone assigned node includes multiple approval tasks, after one task is approved, other tasks are automatically converted into the completed status without any action from the corresponding approver.) When an approval instance moves on to a new node, the tasks of the node are created.

| Parameter | Type | Description |
| --- | --- | --- |
| `id` | `string` | Task ID |
| `user_id` | `string` | ApproverTask user_id is empty for auto-approve and auto-reject tasks. |
| `open_id` | `string` | Open ID of the approver |
| `status` | `string` | Task status **Optional values are**:  - `PENDING`: Under review - `APPROVED`: Agree - `REJECTED`: Reject - `TRANSFERRED`: Transferred - `DONE`: Done  |
| `node_id` | `string` | Node id to which task belongs |
| `node_name` | `string` | Node name to which task belongs |
| `custom_node_id` | `string` | The custom id of the node to which the task belongs. If the custom id is not set, this field will not be returned |
| `type` | `string` | Review method **Optional values are**:  - `AND`: Will sign - `OR`: Individual approval - `AUTO_PASS`: Pass automatically - `AUTO_REJECT`: Automatic rejection - `SEQUENTIAL`: In order  |
| `start_time` | `string` | Task start time |
| `end_time` | `string` | Task completion time, unfinished is 0 | ## Data example

```json
{
    "id": "1234",
    "user_id": "f7cb567e",
    "open_id": "ou_123457",
    "status": "PENDING",
    "node_id": "46e6d96cfa756980907209209ec03b64",
    "node_name": "start",
    "custom_node_id": "manager",
    "type": "AND",
    "start_time": "1564590532967",
    "end_time": "0"
}
```
## User ID description
For how to distinguish between and use user_id, open_id, and union_id, refer to User-related IDs.
