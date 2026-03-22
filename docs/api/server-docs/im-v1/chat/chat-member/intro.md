---
title: "Resource introduction"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/im-v1/chat/chat-member/intro"
service: "im-v1"
resource: "chat"
section: "Group Chat"
updated: "1717574952000"
---

# Resource introduction
## Resource definition
Group members are a collection of members in a group (including users and bots), which describes the relationship between a group and its members.

## Field description
| Parameter | Type | Description |
| --- | --- | --- |
| `items` | `list_member[]` | Member list |
| ∟ `member_id_type` | `string` | User ID type of a member. The value can be `open_id`, `user_id`, or `union_id`. |
| ∟ `member_id` | `string` | User ID of a member. The ID value corresponds to member_id_type. For information about different IDs, see User-related IDs. |
| ∟ `name` | `string` | Name | ### Data example
```json
{
        "items": [
            {
                "member_id_type": "user_id",
                "member_id": "4d7a3c6g",
                "name": "John"
            }
        ]
}
```
