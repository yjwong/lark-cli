---
title: "Obtain group information"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat/get"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/im/v1/chats/:chat_id"
service: "im-v1"
resource: "chat"
section: "Group Chat"
rate_limit: "1000 per minute、50 per second"
scopes:
  - "im:chat"
  - "im:chat:read"
  - "im:chat:readonly"
field_scopes:
  - "contact:user.employee_id:readonly"
updated: "1717574947000"
---

# Obtain group information

Obtains basic information such as group name, description, profile photo, and owner ID.

> Notes:
> - The bot ability needs to be enabled for the app.
> - The bot or authorized user must be in the group; otherwise, only the basic information such as group name and profile photo are returned.
> - When obtaining internal group information, the operator must be under the same tenant as the group.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/im/v1/chats/:chat_id |
| HTTP Method | GET |
| Rate Limit | 1000 per minute、50 per second |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `im:chat` `im:chat:read` `im:chat:readonly` |
| Required field scopes | > The response body of the API contains the following sensitive fields, and they will be returned only after corresponding scopes are added. If you do not need the fields, it is not recommended that you request the scopes. `contact:user.employee_id:readonly` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `chat_id` | `string` | Group ID. For details, refer to Group ID description. **Example value**: "oc_a0553eda9014c201e6969b478895c230" | ### Request example
```bash
curl --location --request GET 'https://open.larksuite.com/open-apis/im/v1/chats/oc_84983ff6516d731e5b5f68d4ea2e1da5?user_id_type=open_id' \
--header 'Authorization: Bearer t-f6a6b2a3a0ddee933a770f7d8678f8f7553bb2a8'
```

### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `user_id_type` | `string` | No | User ID categories **Example value**: open_id **Optional values are**:  - `open_id`: Identifies a user to an app. The same user has different Open IDs in different apps. How to get Open ID - `union_id`: Identifies a user to a tenant that acts as a developer. A user has the same Union ID in apps developed by the same developer, and has different Union IDs in apps developed by different developers. A developer can use Union ID to link the same user's identities in multiple apps.How to get Union ID - `user_id`: Identifies a user to a tenant. The same user has different User IDs in different tenants. In one single tenant, a user has the same User ID in all apps （including store apps）. User ID is usually used to communicate user data between different apps. How to get User ID  **Default value**: `open_id` **When the value is `user_id`, the following field scopes are required**: `contact:user.employee_id:readonly` | ## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `avatar` | `string` | URL of group profile photo |
|   `name` | `string` | Group name |
|   `description` | `string` | Group description |
|   `i18n_names` | `i18n_names` | Internationalized group name |
|     `zh_cn` | `string` | Chinese name |
|     `en_us` | `string` | English name |
|     `ja_jp` | `string` | Japanese name |
|   `add_member_permission` | `string` | Permission scopes to add users or bots to a group. **Optional values are**: - `only_owner`: Only the group owner and admins - `all_members`: All members **Note**: P2P chat does not return this field. |
|   `share_card_permission` | `string` | Group sharing scopes. **Optional values are**: - `allowed`: Allowed - `not_allowed`: Not allowed **Note**: P2P chat does not return this field. |
|   `at_all_permission` | `string` | @mention all scopes. **Optional values are**: - `only_owner`: Only the group owner and admins - `all_members`: All members **Note**: P2P chat does not return this field. |
|   `edit_permission` | `string` | Group edit scopes. **Optional values are**: - `only_owner`: Only the group owner and admins - `all_members`: All members |
|   `owner_id_type` | `string` | Group owner ID type, which is consitent with user_id_type in the query parameter. The value can be `open_id`, `user_id`, or `union_id`. **Note**: - If the group owner is a bot, this field is not returned. - P2P chat does not return this field. |
|   `owner_id` | `string` | Group owner ID. The ID value corresponds to user_id_type in the query parameter. For information about different IDs, see User-related IDs. **Note**: - If the group owner is a bot, this field is not returned. - P2P chat does not return this field. |
|   `user_manager_id_list` | `string[]` | List of user managers |
|   `bot_manager_id_list` | `string[]` | List of bot managers |
|   `group_message_type` | `string` | Group message type **Optional values are**: - `chat`: chat - `thread`: thread **Note**: This field is returned only for groups whose `chat_mode` is "group". |
|   `chat_mode` | `string` | Group mode. **Optional values are**: - `group`: A group - `topic`: A topic - `p2p`: A p2p chat |
|   `chat_type` | `string` | Group type. **Optional values are**: - `private`: A private group - `public`: A public group **Note**: P2P chat does not return this field. |
|   `chat_tag` | `string` | Group tag. If there are multiple group tags, return the first one in the following order. **Optional values are**: - `inner`: An inner group - `tenant`: A company group - `department`: A department group - `edu`: An education group - `meeting`: A meeting group - `customer_service`: A customer service group. **Note**: P2P chat does not return this field. |
|   `join_message_visibility` | `string` | Visibility of the messages on joining a group. **Optional values are**: - `only_owner`: Visible to only the group owner and administrators - `all_members`: Visible to all members - `not_anyone`: Invisible to anyone **Note**: P2P chat does not return this field. |
|   `leave_message_visibility` | `string` | Visibility of the messages on leaving a group. **Optional values are**: - `only_owner`: Visible to only the group owner and administrators - `all_members`: Visible to all members - `not_anyone`: Invisible to anyone **Note**: P2P chat does not return this field. |
|   `membership_approval` | `string` | Group joining approval. **Optional values are**: - `no_approval_required`: Approval is not required - `approval_required`: Approval is required **Note**: P2P chat does not return this field. |
|   `moderation_permission` | `string` | Speech scopes **Optional values are**: - `only_owner`: Only the group owner and admins - `all_members`: All members - `moderator_list`: Specified group members |
|   `external` | `boolean` | Whether it is an external group |
|   `tenant_key` | `string` | Tenant Key, which is the unique identifier of the tenant on Lark, which is used to exchange for the corresponding tenant_access_token, and can also be used as the unique identifier of the tenant in the application |
|   `user_count` | `string` | Number of users |
|   `bot_count` | `string` | Number of bots |
|   `restricted_mode_setting` | `restricted_mode_setting` | Restricted mode setting |
|     `status` | `boolean` | Whether the restricted mode is enabled |
|     `screenshot_has_permission_setting` | `string` | Whether to allow screenshots **Optional values are**:  - `all_members`: All members allow screenshots - `not_anyone`: Not anyone allow screenshots  |
|     `download_has_permission_setting` | `string` | Whether to allow downloading pictures, videos and files in messages **Optional values are**:  - `all_members`: All members allow downloading pictures, videos and files in messages - `not_anyone`: Not anyone allow downloading pictures, videos and files in messages  |
|     `message_has_permission_setting` | `string` | Whether to allow copying and forwarding of messages **Optional values are**:  - `all_members`: All members allow copying and forwarding of messages - `not_anyone`: Not anyone allow copying and forwarding of messages  |
|   `urgent_setting` | `string` | Who can buzz others **Optional values are**:  - `only_owner`: Only group owner and admin - `all_members`: Everyone in this group  |
|   `video_conference_setting` | `string` | Who can start video calls **Optional values are**:  - `only_owner`: Only group owner and admin - `all_members`: Everyone in this group  |
|   `hide_member_count_setting` | `string` | Hide group member number setting **Optional values are**:  - `all_members`: visible to all group members - `only_owner`: visible only to the group owner and administrators  |
|   `chat_status` | `string` | Chat status **Optional values are**:  - `normal`: normal - `dissolved`: dissolved - `dissolved_save`: dissolved and saved  | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
        "avatar": "https://p3-lark-file.byteimg.com/img/lark-avatar-staging/default-avatar_44ae0ca3-e140-494b-956f-78091e348435~100x100.jpg",
        "name": "Test group name",
        "description": "Test group description",
        "i18n_names": {
            "zh_cn": "群聊",
            "en_us": "group chat",
            "ja_jp": "グループチャット"
        },
        "add_member_permission": "all_members",
        "share_card_permission": "allowed",
        "at_all_permission": "all_members",
        "edit_permission": "all_members",
        "owner_id_type": "user_id",
        "owner_id": "4d7a3c6g",
        "user_manager_id_list": [
            "ou_9204a37300b3700d61effaa439f34295"
        ],
        "bot_manager_id_list": [
            "cli_a3e157960e7294c"
        ],
        "group_message_type": "chat",
        "chat_mode": "group",
        "chat_type": "private",
        "chat_tag": "inner",
        "join_message_visibility": "only_owner",
        "leave_message_visibility": "only_owner",
        "membership_approval": "no_approval_required",
        "moderation_permission": "all_members",
        "external": false,
        "tenant_key": "736588c9260f175e",
        "user_count": "1",
        "bot_count": "3",
        "restricted_mode_setting": {
            "status": false,
            "screenshot_has_permission_setting": "all_members",
            "download_has_permission_setting": "all_members",
            "message_has_permission_setting": "all_members"
        },
        "urgent_setting": "all_members",
        "video_conference_setting": "all_members",
        "hide_member_count_setting": "all_members",
        "chat_status": "normal"
    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 232004 | Such an app does NOT exist. | The app_id as the operator doesn't exist. Please contact [technical support](https://applink.larksuite.com/TLJsX982). |
| 400 | 232006 | Your request specifies a chat_id which is invalid. | Invalid chat_id, please check if the chat_id is correct. |
| 400 | 232010 | Operator and chat can NOT be in different tenants. | The operator and the group being operated should be under the same tenant. |
| 400 | 232011 | Operator can NOT be out of the chat. | The operator must be a member of the chat. |
| 400 | 232025 | Bot ability is not activated. | The bot ability needs to be enabled for the app. |
| 400 | 232033 | The operator or invited bots does NOT have the authority to manage external chats without the scope. | There is no permission to operate external groups, and application for this permission is not currently supported. |
| 400 | 232034 | The app is unavailable or inactivated by the tenant. | The app is unavailable or inactivated by the tenant. |
