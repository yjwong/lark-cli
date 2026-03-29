---
title: "Resource introduction"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/message-reaction/overview-of-message-reaction-resources"
service: "im-v1"
resource: "message-reaction"
section: "Messaging"
updated: "1717574934000"
---

# Resource introduction

## Resource definition
A reaction reply to a message.

## Field description
| 名称 | 类型 | 描述 |
| --- | --- | --- |
| `reaction_id` | `string` | Reaction resource ID |
| `operator` | `operator` | The operator who added the reaction |
|   `operator_id` | `string` | Operator ID |
|   `operator_type` | `string` | Specifies whether the operator is a user or an app. **可选值有**：  - `app`: "app" - `user`: "user"  |
| `action_time` | `string` | The Unix timestamp of the reaction (in ms) |
| `reaction_type` | `emoji` | Reaction type |
|   `emoji_type` | `string` | emoji类型，请参考emoji类型列举 | ### Data example
```json
{
    "reaction_id": "ZCaCIjUBVVWSrm5L-3ZTw*************sNa8dHVplEzzSfJVUVLMLcS_",
    "operator": {
        "operator_id": "ou_ff0b7ba35fb********67dfc8b885136",
        "operator_type": "app/user"
    },
    "action_time": "1626086391570",
    "reaction_type": {
        "emoji_type": "SMILE"
    }
}
```
