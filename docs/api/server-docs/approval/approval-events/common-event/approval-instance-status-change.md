---
title: "Approval instance status change"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uIDO24iM4YjLygjN/event/common-event/approval-instance-event"
section: "Approval"
updated: "1676430247000"
---

# Approval instance status change
By subscribing to the event of approval instance status change, the approval instance status message will be pushed to the developer after the instance status changes.
1. After you create an approval, the [PENDING] status will be pushed.
2. If any approver rejects the approval, the [REJECTED] status will be pushed.
3. After all approvers permit, the [APPROVED] status will be pushed.
4. If the approval initiator cancels the approval, the [CANCELED] status will be pushed.
5. If the approval definition is deleted by the administrator, the [DELETED] status will be pushed.
6. When approved instance is reverted by creator, status is REVERTED.

**Approval instance status message**:
```json
{
     "ts": "1502199207.7171419",
     "uuid": "bc447199585340d1f3728d26b1c0297a", 
     "token": "41a9425ea7df4536a7623e38fa321bae",
     "type": "event_callback",
     "event": { 
         "app_id": "cli_xxx", 
         "tenant_key":"xxx", 
         "type": "approval_instance", 
         "approval_code": "7C468A54-8745-2245-9675-08B7C63E7A85",
         "instance_code": "81D31358-93AF-92D6-7425-01A5D67C4E71", 
         "status": "PENDING",
         "instance_operate_time": "1666079207003",
         "uuid":"6525bffb"
    }
}
```
**Message description**:
| Field         | Type          | Description          |
| --------- | --------------- | --------- |
| `type` | `string` | approval_instance;Fixed field |
| `approval_code` | `string` | Approval definition code | 
| `instance_code` | `string` | Approval instance code |
| `status` | `string` | Instance status;PENDING - In progress;APPROVED - Approved;REJECTED - Rejected;CANCELED - Canceled;DELETED - Deleted;REVERTED - Reverted |
| `instance_operate_time` | `string` | Event occurrence time |
| `uuid` | `string` | The approval instance defines a custom unique ID, which is passed in when the interface is created for approval. |
