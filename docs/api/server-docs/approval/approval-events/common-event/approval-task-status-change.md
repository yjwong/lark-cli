---
title: "Approval task status change"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uIDO24iM4YjLygjN/event/common-event/approval-task-event"
section: "Approval"
updated: "1676430247000"
---

# Approval task status change
After the approver approves, rejects, transfers, or returns the approval task, the approval task status message will be pushed to the developer.
1. After you create an approval, the [PENDING] status of the approval task in the first approval step will be pushed.
2. If the current step is everyone assigned (AND):
   - 	If any approval task is approved, the [APPROVED] status of the task will be pushed.
   - 	If any approval task is rejected, the [REJECTED] status of the task and the [DONE] status of the remaining tasks in the current step will be pushed.
3. If the current step is anyone assigned (OR):
   - If any approval task is approved, the [APPROVED] status of the task, the [DONE] status of the remaining tasks in the current step, and the [PENDING] status of all tasks in the next step will be pushed.
   - If any approval task is rejected, the [REJECTED] status of the task and the [DONE] status of all remaining tasks in the current step will be pushed.
4. If you transfer the approval task, the [TRANSFERRED] status of the task and the [PENDING] status of the transferee task will be pushed.
5. After the approval initiator cancels the approval, the [DONE] status of all remaining tasks will be pushed.
6. If the approval definition is deleted by the administrator, the [DONE] status of all remaining tasks will be pushed.
7. If the user returns the approval task, push the [ROLLBACK] status of the task and the [PENDING] status of the returned task

**Approval task status message**:
```json
{    
     "ts": "1502199207.7171419",
     "uuid": "bc447199585340d1f3728d26b1c0297a",
     "token": "41a9425ea7df4536a7623e38fa321bae",
     "type": "event_callback",
     "event": { 
         "app_id": "cli_xxx", 
         "open_id": "ou_123456",
         "tenant_key":"xxx", 
         "type": "approval_task", 
         "approval_code": "7C468A54-8745-2245-9675-08B7C63E7A85",
         "instance_code": "81D31358-93AF-92D6-7425-01A5D67C4E71", 
         "task_id": "12345",
         "user_id": "b613t51g",
         "status": "PENDING", 
         "operate_time": "1502199207000",
         "custom_key": "xxx",
         "def_key": "xxx",
         "extra":"{\"rollback_node_ids\":[\"nodeid\"],\"rollback_custom_node_ids\":[\"customnodeid\"]}"
    }
}
```
**Message description**:
| Field         | Type          | Description          |
| --------- | --------------- | --------- |
| `type` | `string` | approval_task;Fixed field |
| `approval_code` | `string` | Approval definition code | 
| `instance_code` | `string` | Approval instance code |
| `task_id` | `string` | Approval task ID |
| `user_id` | `string` | Operator ID (if it is an auto-approve task, user_id is empty) |
| `status` | `string` | Task status;PENDING - In progress;APPROVED - Approved;REJECTED - Rejected;TRANSFERRED - Transferred;ROLLBACK - returned;DONE - Completed |
| `operate_time` | `string` | Event occurrence time |
|`extra` | `string` | Extended data, currently only rollback events have this field, rollback_node_ids rollback node list, rollback_custom_node_ids user-defined node list|
|`custom_key` | `string` | Node custom ID |
|`def_key` | `string` | Node system generates unique ID |
