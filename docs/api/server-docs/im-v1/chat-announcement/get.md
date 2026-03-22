---
title: "Obtain group announcement information"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat-announcement/get"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/im/v1/chats/:chat_id/announcement"
service: "im-v1"
resource: "chat-announcement"
section: "Group Chat"
rate_limit: "1000 per minute、50 per second"
scopes:
  - "im:chat"
  - "im:chat.announcement:read"
  - "im:chat:readonly"
field_scopes:
  - "contact:user.employee_id:readonly"
updated: "1717574957000"
---

# Obtain group announcement information

Obtains the group announcement in a chat, with the same format as Docs.

> Notes:
> - The bot ability needs to be enabled.
> - The bot or authorized user must be in the group.
> - When obtaining internal group information, the operator must be under the same tenant as the group.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/im/v1/chats/:chat_id/announcement |
| HTTP Method | GET |
| Rate Limit | 1000 per minute、50 per second |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `im:chat` `im:chat.announcement:read` `im:chat:readonly` |
| Required field scopes | > The response body of the API contains the following sensitive fields, and they will be returned only after corresponding scopes are added. If you do not need the fields, it is not recommended that you request the scopes. `contact:user.employee_id:readonly` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `chat_id` | `string` | The ID of the group with its announcement to be obtained. For details, refer to Group ID description. **Note**: P2P chat is not supported. **Example value**: "oc_5ad11d72b830411d72b836c20" | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `user_id_type` | `string` | No | User ID categories **Example value**: open_id **Optional values are**:  - `open_id`: Identifies a user to an app. The same user has different Open IDs in different apps. How to get Open ID - `union_id`: Identifies a user to a tenant that acts as a developer. A user has the same Union ID in apps developed by the same developer, and has different Union IDs in apps developed by different developers. A developer can use Union ID to link the same user's identities in multiple apps.How to get Union ID - `user_id`: Identifies a user to a tenant. The same user has different User IDs in different tenants. In one single tenant, a user has the same User ID in all apps （including store apps）. User ID is usually used to communicate user data between different apps. How to get User ID  **Default value**: `open_id` **When the value is `user_id`, the following field scopes are required**: `contact:user.employee_id:readonly` | ## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `content` | `string` | Docs serialization info |
|   `revision` | `string` | Current Docs version number (in pure digits) |
|   `create_time` | `string` | Timestamp of Docs creation (in seconds) |
|   `update_time` | `string` | Timestamp of Docs update (in seconds) |
|   `owner_id_type` | `string` | Docs owner ID type - If the owner is a user, the value is identical to the user_id_type value in the query parameter, which can be `open_id`, `user_id`, or `union_id`. For information about different IDs, see User-related ID. - If the owner is a bot, the value is the `app_id` of the bot. For details, refer to Obtain the token to access as an app. **Optional values are**:  - `user_id`: Identifies users by user_id - `union_id`: Identifies users by union_id - `open_id`: Identifies users by open_id - `app_id`: Identifies bots by app_id  |
|   `owner_id` | `string` | Docs owner ID. The ID value corresponds to the ID type of the owner_id_type. |
|   `modifier_id_type` | `string` | ID Type of the last Docs modifier - If the modifier is a user, the value is identical to the user_id_type value in the query parameter, which can be `open_id`, `user_id`, or `union_id`. For information about different IDs, see User-related ID. - If the modifier is a bot, the value is the `app_id` of the bot. For details, refer to Obtain the token to access as an app. **Optional values are**:  - `user_id`: Identifies users by user_id - `union_id`: Identifies users by union_id - `open_id`: Identifies users by open_id - `app_id`: Identifies apps by app_id  |
|   `modifier_id` | `string` | Last Docs modifier ID. The ID value corresponds to the ID type of modifier_id_type. | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
        "content": "xxx",
        "revision": "12",
        "create_time": "1609296809",
        "update_time": "1609296809",
        "owner_id_type": "open_id",
        "owner_id": "ou_7d8a6e6df7621556ce0d21922b676706ccs",
        "modifier_id_type": "open_id",
        "modifier_id": "ou_7d8a6e6df7621556ce0d21922b676706ccs"
    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 232001 | Your request contains an invalid request parameter. | Parameter error. Check the input parameters according to the appropriate document. |
| 400 | 232003 | Chat announcement can NOT be found in chat information. | The group announcement information is abnormal. |
| 400 | 232010 | Operator and chat can NOT be in different tenants. | The operator and the group being operated should be under the same tenant. |
| 400 | 232011 | Operator can NOT be out of the chat. | The operator must be a member of the chat. |
| 400 | 232019 | The request has been rate limited. | Rate-limiting is triggered. Pay attention to the request rate. For details, please see Frequency Control Policy. |
| 400 | 232025 | Bot ability is not activated. | The bot ability needs to be enabled for the app. |
| 400 | 232033 | The operator or invited bots does NOT have the authority to manage external chats without the scope. | The app does not have permission to operate external groups. |
| 400 | 232034 | The app is unavailable or inactivated by the tenant. | The app is unavailable or inactivated by the tenant. |
| 400 | 232066 | The operator does not have the doc permission. | The operator does not have the permission to read the document. Please add the document permission and try again. |
| 400 | 232097 | Unable to operate docx type chat announcement. | Unable to operate docx type chat announcement. Please refer to the  obtaining new version group announcement information API. |
