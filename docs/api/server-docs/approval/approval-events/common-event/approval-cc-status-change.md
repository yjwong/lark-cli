---
title: "Approval cc status change"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uIDO24iM4YjLygjN/event/common-event/approval-cc-event"
section: "Approval"
updated: "1676430247000"
---

# Approval cc status change
When a CC is created or the Cc has been read, a message is sent to the developer.
1. Create a cc and push the event "operate" to "CREATE".
1. The Cc recipient has read the Cc and pushed the event that "operate" is "READ".

**Approval CC message**:
```json
{    
     "ts": "1502199207.7171419",
     "uuid": "bc447199585340d1f3728d26b1c0297a",
     "token": "41a9425ea7df4536a7623e38fa321bae",
     "type": "event_callback",
     "event": { 
         "app_id": "cli_xxx", 
         "tenant_key":"xxx", 
         "type": "approval_cc", 
         "approval_code": "7C468A54-8745-2245-9675-08B7C63E7A85",
         "instance_code": "81D31358-93AF-92D6-7425-01A5D67C4E71", 
         "id": "12345",
         "user_id": "b613t51g",
         "create_time": 1502199207000, 
         "operate": "CREATE",
         "from": "b613t51g"
    }
}
```
**消息说明**:
| 字段         | 类型           | 描述        |
| --------- | --------------- | --------- |
| `type` | `string` | approval_cc;Fixed field |
| `approval_code` | `string` | Approval definition code |
| `instance_code` | `string` | Approval instance code |
| `id` | `string` | CC ID|
| `user_id` | `string` | CC recipients |
| `create_time` | `long` | CC time |
|`operate` | `string` | Action type CREATE: CC READ: Read |
| `from` | `string` | CC sender, which may be empty |
