---
title: "Resource introduction"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/im-v1/chat/chat-info/intro"
service: "im-v1"
resource: "chat"
section: "Group Chat"
updated: "1717574943000"
---

# Resource introduction
## Resource definition
It refers to the chat groups in Lark.

## Field description
> **Note:** The following fields are not available for all types of groups, and some groups may be missing some fields.
| Parameter | Type | Description |
| --- | --- | --- |
| `chat_id` | `string` | Group ID. For details, refer to [Group ID description]. |
| `avatar` | `string` | URL of group profile photo |
| `name` | `string` | Group name |
| `description` | `string` | Group description |
| `i18n_names` | `i18n_names` | Internationalized group name |
| ∟ `zh_cn` | `string` | Chinese name |
| ∟ `en_us` | `string` | English name |
| ∟ `ja_jp` | `string` | Japanese name |
| `owner_id` | `string` | Group owner ID. The ID value corresponds to owner_id_type. For information about different IDs, see User-related IDs. If the group owner is a bot, this field is not returned. |
| `owner_id_type` | `string` | ID type that corresponds to the owner ID. The value can be `open_id`, `user_id`, or `union_id`. For information about different IDs, see User-related IDs. If the group owner is a bot, this field is not returned. |
| `add_member_permission` | `string` | Permission scopes to add users or bots to a group **Optional values are**: - `only_owner`: Only the group owner and admins - `all_members`: All members |
| `share_card_permission` | `string` | Group sharing scopes **Optional values are**: - `allowed`: Allowed - `not_allowed`: Not allowed |
| `at_all_permission` | `string` | @mention all scopes **Optional values are**: - `only_owner`: Only the group owner and admins - `all_members`: All members |
| `edit_permission` | `string` | Group edit scopes **Optional values are**: - `only_owner`: Only the group owner and admins - `all_members`: All members |
| `group_message_type` | `string` | Group message type **可选值有**： - `chat`：chat - `thread`：thread |
| `chat_mode` | `string` | Group mode **Optional values are**: - `group`: A group - `topic`: A topic group - `p2p`: A private chat |
| `chat_type` | `string` | Group type **Optional values are**: - `private`: A private group - `public`: A public group |
| `chat_tag` | `string` | Group tag. If there are multiple group tags, return the first one in the following order. **Optional values are**: - `inner`: An inner group - `tenant`: A company group - `department`: A department group - `edu`: An education group - `meeting`: A meeting group - `customer_service`: A customer service group |
| `external` | `boolean` | Whether it is an external group |
| `tenant_key` | `string` | A tenant's unique ID in Lark. It can be used to obtain tenant_access_token, and can also be used as a tenant's unique ID in an app. |
| `join_message_visibility` | `string` | Visibility of the messages on joining a group **Optional values are**: - `only_owner`: Visible to only the group owner and administrators - `all_members`: Visible to all members - `not_anyone`: Invisible to anyone |
| `leave_message_visibility` | `string` | Visibility of the messages on leaving a group **Optional values are**: - `only_owner`: Visible to only the group owner and administrators - `all_members`: Visible to all members - `not_anyone`: Invisible to anyone |
| `membership_approval` | `string` | Group joining approval **Optional values are**: - `no_approval_required`: Approval is not required. - `approval_required`: Approval is required. |
| `moderation_permission` | `string` | Speech scopes **Optional values are**: - `only_owner`: Only the group owner and admins - `all_members`: All members - `moderator_list`: Specified group members | ### Data example
```json
{
        "chat_id": "oc_a0553eda9014c201e6969b478895c230",
        "avatar": "https://p3-lark-file.byteimg.com/img/lark-avatar-staging/default-avatar_44ae0ca3-e140-494b-956f-78091e348435~100x100.jpg",
        "name": "test group name",
        "description": "test group description",
        "i18n_names": {
            "zh_cn": "group chat",
            "en_us": "group chat",
            "ja_jp": "グループチャット"
        },
        "owner_id": "4d7a3c6g",
        "owner_id_type": "user_id",
        "add_member_permission": "all members",
        "share_card_permission": "allowed",
        "at_all_permission": "all members",
        "edit_permission": "all members",
        "group_message_type": "chat"
        "chat_mode": "group",
        "chat_type": "private",
        "chat_tag": "inner",
        "external": false,
        "tenant_key": "736588c9260f175e",
        "join_message_visibility": "all_members",
        "leave_message_visibility": "all_members",
        "membership_approval": "no_approval_required",
        "moderation_permission": "all_members",
}
```

## Group ID description
Refer to Group ID description.
