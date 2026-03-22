---
title: "Updates group speech scopes"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat-moderation/update"
method: "PUT"
api_path: "https://open.larksuite.com/open-apis/im/v1/chats/:chat_id/moderation"
service: "im-v1"
resource: "chat-moderation"
section: "Group Chat"
rate_limit: "1000 per minute、50 per second"
scopes:
  - "im:chat"
  - "im:chat:moderation:write_only"
field_scopes:
  - "contact:user.employee_id:readonly"
updated: "1717574947000"
---

# Updates group speech scopes

Updates the settings of group speech scopes, which can be granted to all members, only admins, or only specified users.

> Notes:
> - The bot ability needs to be enabled.
> - If the API is called with user authorization, the group speech scopes can be updated when **the authorized user is the group owner**.
> - If the API is called with tenant authorization (that is, as a bot), the group speech scopes can be updated when **the bot is the group owner** or **the bot is the group creator, has the Update the information of groups created by app scope, and is still in the group**.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/im/v1/chats/:chat_id/moderation |
| HTTP Method | PUT |
| Rate Limit | 1000 per minute、50 per second |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `im:chat` `im:chat:moderation:write_only` |
| Required field scopes | > The response body of the API contains the following sensitive fields, and they will be returned only after corresponding scopes are added. If you do not need the fields, it is not recommended that you request the scopes. `contact:user.employee_id:readonly` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `chat_id` | `string` | Group ID. For details, refer to Group ID description. **Example value**: "oc_a0553eda9014c201e6969b478895c230" | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `user_id_type` | `string` | No | User ID categories **Example value**: open_id **Optional values are**:  - `open_id`: Identifies a user to an app. The same user has different Open IDs in different apps. How to get Open ID - `union_id`: Identifies a user to a tenant that acts as a developer. A user has the same Union ID in apps developed by the same developer, and has different Union IDs in apps developed by different developers. A developer can use Union ID to link the same user's identities in multiple apps.How to get Union ID - `user_id`: Identifies a user to a tenant. The same user has different User IDs in different tenants. In one single tenant, a user has the same User ID in all apps （including store apps）. User ID is usually used to communicate user data between different apps. How to get User ID  **Default value**: `open_id` **When the value is `user_id`, the following field scopes are required**: `contact:user.employee_id:readonly` | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `moderation_setting` | `string` | No | Group speech mode (The options are `all_members`, `only_owner`, and `moderator_list`, and moderator_list indicates certain users are allowed to speak.) **Example value**: "moderator_list" |
| `moderator_added_list` | `string[]` | No | The list of users who are allowed to speak when moderator_list is selected. Open ID is recommended here. For details, refer to How to get Open ID? (The users not in the group are filtered automatically.) **Example value**: ["4d7a3c6g"] |
| `moderator_removed_list` | `string[]` | No | The list of users who are not allowed to speak when moderator_list is selected. Open ID is recommended here. For details, refer to How to get Open ID? (The users not in the group are filtered automatically.) **Example value**: ["4d7a3c6g"] | ### Request body example

{
    "moderation_setting": "moderator_list",
    "moderator_added_list": [
        "4d7a3c6g"
    ],
    "moderator_removed_list": [
        "4d7a3c6g"
    ]
}

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
| 400 | 232004 | Such an app does NOT exist. | The app_id as the operator doesn't exist. Please contact [technical support](https://applink.larksuite.com/TLJsX982). |
| 400 | 232006 | Your request specifies a chat_id which is invalid. | Invalid chat_id, please check if the chat_id is correct. |
| 400 | 232009 | Your request specifies a chat which has already been dissolved. | The group has been deleted. |
| 400 | 232010 | Operator and chat can NOT be in different tenants. | The operator and the group being operated should be under the same tenant. |
| 400 | 232011 | Operator can NOT be out of the chat. | The operator must be a member of the chat. |
| 400 | 232015 | Your request specifies a member_id_type which is NOT supported. | Invalid user_id_type/member_id_type. |
| 400 | 232017 | No Permission: If the operator is NOT owner or creator with the scope, the operator can NOT complete the request. | The operator is neither the group owner nor the group creator with permission scopes. |
| 400 | 232019 | The request has been rate limited. | Rate-limiting is triggered. Pay attention to the request rate. For details, please see Frequency Control Policy. |
| 400 | 232025 | Bot ability is not activated. | The bot ability needs to be enabled for the app. |
| 400 | 232033 | The operator or invited bots does NOT have the authority to manage external chats without the scope. | No permission to operate external groups. |
| 400 | 232034 | The app is unavailable or inactivated by the tenant. | The app is unavailable or inactivated by the tenant. |
| 400 | 232060 | This chat is banned. | The  group chat is banned, please consult [Technical Support](https://applink.larksuite.com/TLJsX982). |
| 400 | 232092 | Meeting in progress. Unable to modify group posting permissions. | Meeting in progress. Unable to modify group posting permissions. |
