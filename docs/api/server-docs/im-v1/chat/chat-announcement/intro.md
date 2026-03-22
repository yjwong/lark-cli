---
title: "Resource introduction"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/im-v1/chat/chat-announcement/intro"
service: "im-v1"
resource: "chat"
section: "Group Chat"
updated: "1717574957000"
---

# Resource introduction
## Resource definition
Group announcement refers to the announcement document in a group. It is stored in Lark Docs and is specific to each group.

## Field description
| Parameter | Type | Description |
| --- | --- | --- |
| `content` | `string` | Docs serialization info |
| `revision` | `string` | Current document version number (in pure digits) |
| `create_time` | `string` | Timestamp of document creation (in seconds) |
| `update_time` | `string` | Timestamp of document update (in seconds) |
| `owner_id_type` | `string` | Document owner ID type - If the owner is a user, the value can be `open_id`, `user_id`, or `union_id`. For information about different IDs, see User-related IDs . - If the owner is a bot, the value is the `app_id` of the bot. For details, refer to General parameters. **Optional values are:** - `user_id`: Identifies users by user_id - `union_id`: Identifies users by union_id - `open_id`: Identifies users by open_id - `app_id`: Identifies bots by app_id |
| `owner_id` | `string` | Document owner ID. The ID value corresponds to the ID type of owner_id_type. |
| `modifier_id_type` | `string` | ID type of last document modifier - If the modifier is a user, the value can be `open_id`, `user_id`, or `union_id`. For the information about different IDs, see User-related IDs. - If the modifier is a bot, the value is the `app_id` of the bot. For details, refer to General parameters. **Optional values are:** - `user_id`: Identifies users by user_id - `union_id`: Identifies users by union_id - `open_id`: Identifies users by open_id - `app_id`: Identifies apps by app_id |
| `modifier_id` | `string` | Last document modifier ID. The ID value corresponds to the ID type of modifier_id_type. | ### Data example
```json
{
        "content": "{\"title\":{\"elements\":[{\"type\":\"textRun\",\"textRun\":{\"text\":\"announcement\",\"style\":{},\"location\":{\"zoneId\":\"0\",\"startIndex\":0,\"endIndex\":12}}}],\"location\":{\"zoneId\":\"0\",\"startIndex\":0,\"endIndex\":12},\"lineId\":\"4m10Rp\"},\"body\":{\"blocks\":[{\"type\":\"paragraph\",\"paragraph\":{\"elements\":[{\"type\":\"textRun\",\"textRun\":{\"text\":\"Announcement.\",\"style\":{},\"location\":{\"zoneId\":\"0\",\"startIndex\":13,\"endIndex\":26}}}],\"location\":{\"zoneId\":\"0\",\"startIndex\":13,\"endIndex\":26},\"lineId\":\"5VYRPT\"}},{\"type\":\"paragraph\",\"paragraph\":{\"elements\":[],\"location\":{\"zoneId\":\"0\",\"startIndex\":27,\"endIndex\":27}}}]}}",
        "revision": "12",
        "create_time": "1609296809",
        "update_time": "1609296809",
        "owner_id_type": "open_id",
        "owner_id": "ou_7d8a6e6df7621556ce0d21922b676706ccs",
        "modifier_id_type": "open_id",
        "modifier_id": "ou_7d8a6e6df7621556ce0d21922b676706ccs"
}
```
