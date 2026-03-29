---
title: "Update group information"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat/update"
method: "PUT"
api_path: "https://open.larksuite.com/open-apis/im/v1/chats/:chat_id"
service: "im-v1"
resource: "chat"
section: "Group Chat"
rate_limit: "1000 per minute、50 per second"
scopes:
  - "im:chat"
  - "im:chat:update"
field_scopes:
  - "contact:user.employee_id:readonly"
updated: "1717574947000"
---

# Update group information

Updates the group profile photo, name, description, and configuration, or assigns new group owner.

> Notes:
> - The bot ability needs to be enabled for the app.
> - For group owners/group administrators or bots who have created a group and have the Update the information of groups created by app permission, all information can be updated.
> - For group members or bots that do not meet the above permission conditions:
> - If not enabled Only group owner and group admin can edit group .infoConfiguration, only group avatar, group name, group description, and group international name information can be updated.
> - If enabled Only group owner and group admin can edit group infoConfiguration, no group information can be modified.
> - If update the Permission to add users/bots to a group and Group sharing permission at the same time, these two settings need to meet the following conditions:
> - If not enabled Only group owner and group admin can invite users or robots into the group, you need to set Group sharing permission for allowed
> - If enabled Only group owner and group admin can invite users or robots into the group, you need to set Group sharing permission for not allowed

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/im/v1/chats/:chat_id |
| HTTP Method | PUT |
| Rate Limit | 1000 per minute、50 per second |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `im:chat` `im:chat:update` |
| Required field scopes | > The response body of the API contains the following sensitive fields, and they will be returned only after corresponding scopes are added. If you do not need the fields, it is not recommended that you request the scopes. `contact:user.employee_id:readonly` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `chat_id` | `string` | Group ID. For details, refer to Group ID description. **Note**: Only support group IDs whose group mode is `group`. **Example value**: "oc_a0553eda9014c201e6969b478895c230" | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `user_id_type` | `string` | No | User ID categories **Example value**: open_id **Optional values are**:  - `open_id`: Identifies a user to an app. The same user has different Open IDs in different apps. How to get Open ID - `union_id`: Identifies a user to a tenant that acts as a developer. A user has the same Union ID in apps developed by the same developer, and has different Union IDs in apps developed by different developers. A developer can use Union ID to link the same user's identities in multiple apps.How to get Union ID - `user_id`: Identifies a user to a tenant. The same user has different User IDs in different tenants. In one single tenant, a user has the same User ID in all apps （including store apps）. User ID is usually used to communicate user data between different apps. How to get User ID  **Default value**: `open_id` **When the value is `user_id`, the following field scopes are required**: `contact:user.employee_id:readonly` | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `avatar` | `string` | No | The image key that corresponds to the group profile photo. It can be obtained using the Upload image API. (Note: The image_type of the uploaded image needs to be specified as avatar) **Example value**: "default-avatar_44ae0ca3-e140-494b-956f-78091e348435" |
| `name` | `string` | No | Group name **Example value**: "Group chat" |
| `description` | `string` | No | Group description **Example value**: "Test group description" |
| `i18n_names` | `i18n_names` | No | Internationalized group name |
|   `zh_cn` | `string` | No | Chinese name **Example value**: "群聊" |
|   `en_us` | `string` | No | English name **Example value**: "group chat" |
|   `ja_jp` | `string` | No | Japanese name **Example value**: "グループチャット" |
| `add_member_permission` | `string` | No | Permission to add users/bots to a group Notes： - If the value is set to `only_owner`, then share_card_permission can only be set to `not_allowed`. - If the value is set to `all_members`, then share_card_permission can only be set to `allowed`. **Optional values are**: - `only_owner`: Only the group owner and admins - `all_members`: All members **Example value**: "all_members" |
| `share_card_permission` | `string` | No | Group sharing permission **Optional values are**: - `allowed`: Allowed - `not_allowed`: Not allowed **Example value**: "allowed" |
| `at_all_permission` | `string` | No | @mention all permission **Optional values are**: - `only_owner`: Only the group owner and admins - `all_members`: All members **Example value**: "all_members" |
| `edit_permission` | `string` | No | Group edit permission **Optional values are**: - `only_owner`: Only the group owner and admins - `all_members`: All members **Example value**: "all_members" |
| `owner_id` | `string` | No | New group owner ID, no need to fill in if the group owner is not transferred. Open ID is recommended here. For details, refer to How to get Open ID? **Example value**: "4d7a3c6g" |
| `join_message_visibility` | `string` | No | Visibility of the messages on joining a group **Optional values are**: - `only_owner`: Visible to only the group owner and administrators - `all_members`: Visible to all members - `not_anyone`: Invisible to anyone **Example value**: "only_owner" |
| `leave_message_visibility` | `string` | No | Visibility of the messages on leaving a group **Optional values are**: - `only_owner`: Visible to only the group owner and administrators - `all_members`: Visible to all members - `not_anyone`: Invisible to anyone **Example value**: "only_owner" |
| `membership_approval` | `string` | No | Group joining approval **Optional values are**: - `no_approval_required`: Approval is not required. - `approval_required`: Approval is required. **Example value**: "no_approval_required" |
| `restricted_mode_setting` | `restricted_mode_setting` | No | Restricted mode setting |
|   `status` | `boolean` | No | Whether the restricted mode is enabled **Example value**: false |
|   `screenshot_has_permission_setting` | `string` | No | Whether to allow screenshots **Example value**: "all_members" **Optional values are**:  - `all_members`: All members allow screenshots - `not_anyone`: Not anyone allow screenshots  |
|   `download_has_permission_setting` | `string` | No | Whether to allow downloading pictures, videos and files in messages **Example value**: "all_members" **Optional values are**:  - `all_members`: All members allow downloading pictures, videos and files in messages - `not_anyone`: Not anyone allow downloading pictures, videos and files in messages  |
|   `message_has_permission_setting` | `string` | No | Whether to allow copying and forwarding of messages **Example value**: "all_members" **Optional values are**:  - `all_members`: All members allow copying and forwarding of messages - `not_anyone`: Not anyone allow copying and forwarding of messages  |
| `chat_type` | `string` | No | Group type **Optional values are**: - `private`: A private group - `public`: A public group **Example value**: "private" |
| `group_message_type` | `string` | No | Group message type **Example value**: "chat" **Optional values are**:  - `chat`: chat - `thread`: thread  |
| `urgent_setting` | `string` | No | Who can buzz others **Example value**: "all_members" **Optional values are**:  - `only_owner`: Only group owner and admin - `all_members`: Everyone in this group  |
| `video_conference_setting` | `string` | No | Who can start video calls **Example value**: "all_members" **Optional values are**:  - `only_owner`: Only group owner and admin - `all_members`: Everyone in this group  |
| `hide_member_count_setting` | `string` | No | Hide member count setting **Example value**: "all_members" **Optional values are**:  - `all_members`: visible to all group members - `only_owner`: visible only to the group owner and group administrators  | > **Note:** If you do not fill in the corresponding fields, keep the original group information.

### Request body example

```json
{
    "avatar": "default-avatar_44ae0ca3-e140-494b-956f-78091e348435",
    "name": "Group chat",
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
    "owner_id": "4d7a3c6g",
    "join_message_visibility": "only_owner",
    "leave_message_visibility": "only_owner",
    "membership_approval": "no_approval_required",
    "restricted_mode_setting": {
        "status": false,
        "screenshot_has_permission_setting": "all_members",
        "download_has_permission_setting": "all_members",
        "message_has_permission_setting": "all_members"
    },
    "chat_type": "private",
    "group_message_type": "chat",
    "urgent_setting": "all_members",
    "video_conference_setting": "all_members",
    "hide_member_count_setting": "all_members"
}
```

## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {}
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 232001 | Your request contains an invalid request parameter. | Parameter error. Check the input parameters according to the appropriate document. |
| 400 | 232002 | No Permission: Only chat owner or admin can edit chat information in the current situation. | Only the group owner or administrators are allowed to edit the group information. |
| 400 | 232004 | Such an app does NOT exist. | The app_id as the operator doesn't exist. Please contact [technical support](https://applink.larksuite.com/TLJsX982). |
| 400 | 232006 | Your request specifies a chat_id which is invalid. | Invalid chat_id, please check if the chat_id is correct. |
| 400 | 232008 | Your request specifies a chat whose type is NOT supported currently. | The group does not support setting to the specified chat_mode, chat_type or group_message_type. |
| 400 | 232009 | Your request specifies a chat which has already been dissolved. | The group has been deleted. |
| 400 | 232010 | Operator and chat can NOT be in different tenants. | The operator and the group being operated should be under the same tenant. |
| 400 | 232011 | Operator can NOT be out of the chat. | The operator must be a member of the chat. |
| 400 | 232012 | New group owner can NOT be out of the chat. | The new group owner is not in the group. |
| 400 | 232016 | Non-chat-owner or Non-chat-admin can only edit certain parts. | The group members other than the group owner and administrators can only modify certain group information (such as avatar, name, description, i18n_names). |
| 400 | 232019 | The request has been rate limited. | Rate-limiting is triggered. Pay attention to the request rate. For details, please see Frequency Control Policy. |
| 400 | 232023 | Chat information review failed while updating the chat. | The review of group-related information failed. Please check whether there is any sensitive content in the group name or group description. |
| 400 | 232025 | Bot ability is not activated. | The bot ability needs to be enabled for the app. |
| 400 | 232026 | This name is already used in an existing public chat. Names of public chats are supposed to be different. | The name of the public group already exists. |
| 400 | 232030 | Your request specifies a user_id which is invalid. | Please check if user_id is correct. |
| 400 | 232031 | The chat and the new designated chat owner must be in the same tenant. | The new group owner cannot be assigned across tenants. |
| 400 | 232033 | The operator or invited bots does NOT have the authority to manage external chats without the scope. | No permission to operate external groups. |
| 400 | 232034 | The app is unavailable or inactivated by the tenant. | The app is unavailable or inactivated by the tenant. |
| 400 | 232035 | Your request specifies an owner_id which is invalid. | Checks if owner_id is correct. |
| 400 | 232037 | The operator or invited bots does NOT have the authority to manage chat labels without the scope. | The operator or the invited bot does not have permission to manage the group label. Please check the permission configuration. |
| 400 | 232041 | The avatar key is illegal. | The key of the group profile photo is invalid. Check it and enter again. |
| 400 | 232047 | The length of the tab name reaches the limit. | Chat tab name is too long. |
| 400 | 232057 | The operator tenant doesn't have the permission to use restricted mode. | The tenant to which the operator belongs cannot use the restricted mode. Please contact the tenant administrator. |
| 400 | 232069 | current chat type unsupported to set public. | External groups whose group type is public are not supported. |
| 400 | 232091 | Due to the security control requirements of this tenant, this tenant does not allow public group. | Due to the security control requirements of this tenant, this tenant does not allow public group. Please contact the tenant administrator |
| 400 | 232093 | This meeting has restricted access. Unable to turn off group membership approval. | In the middle of a meeting, it is not supported to turn off group verification. |
| 400 | 232078 | The operator tenant doesn't have the permission to use hide_member_count_setting. | The operator tenant doesn't have the permission to use hide_member_count_setting. |
