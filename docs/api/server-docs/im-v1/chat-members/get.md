---
title: "Obtain group member list"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat-members/get"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/im/v1/chats/:chat_id/members"
service: "im-v1"
resource: "chat-members"
section: "Group Chat"
rate_limit: "1000 per minute、50 per second"
scopes:
  - "im:chat"
  - "im:chat.group_info:readonly"
  - "im:chat.members:read"
  - "im:chat:readonly"
field_scopes:
  - "contact:user.employee_id:readonly"
updated: "1717574953000"
---

# Obtain group member list

Get the list of members of the group the user/bot is in.

> Notes:
> - The bot ability needs to be enabled.
> - The bot or authorized user must be in the group.
> - This API doesn't return the bot in a group.
> - As the group member list returned does not contain the bot, the number of group members returned may be less than the specified page_size value.
> - The members that joined a group at the same time are returned together, so the number of group members returned may be greater than the specified page_size value.
> - When obtaining internal group information, the operator must be under the same tenant as the group.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/im/v1/chats/:chat_id/members |
| HTTP Method | GET |
| Rate Limit | 1000 per minute、50 per second |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `im:chat` `im:chat.group_info:readonly` `im:chat.members:read` `im:chat:readonly` |
| Required field scopes | > The response body of the API contains the following sensitive fields, and they will be returned only after corresponding scopes are added. If you do not need the fields, it is not recommended that you request the scopes. `contact:user.employee_id:readonly` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `chat_id` | `string` | Group ID. For details, refer to Group ID description. **Example value**: "oc_a0553eda9014c201e6969b478895c230" | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `member_id_type` | `string` | No | User ID categories **Example value**: open_id **Optional values are**:  - `open_id`: Identifies members by open_id **When the value is `user_id`, the following field scopes are required:** `contact:user.employee_id:readonly` - `union_id`: Identifies members by union_id - `user_id`: Identifies members by user_id  **Default value**: `open_id` **When the value is `user_id`, the following field scopes are required**: `contact:user.employee_id:readonly` |
| `page_size` | `int` | No | **Example value**: 20 **Default value**: `20` **Data validation rules**: - Maximum value: `100` |
| `page_token` | `string` | No | Page identifier. It is not filled in the first request, indicating traversal from the beginning; when there will be more groups, the new page_token will be returned at the same time, and the next traversal can use the page_token to get more groups **Example value**: WWxHTStrOEs5WHZpNktGbU94bUcvMWlxdDUzTWt1OXNrRmlLaGRNVG0yaz0= | ## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `items` | `list_member[]` | Member list |
|     `member_id_type` | `string` | User ID type of a member, which is consistent with member_id_type in the query parameter. The value can only be `open_id`, `user_id`, or `union_id`. |
|     `member_id` | `string` | User ID of a member. The ID value corresponds to member_id_type in the query parameter. For information about different IDs, see User-related IDs. |
|     `name` | `string` | Name |
|     `tenant_key` | `string` | Tenant Key, which is the unique identifier of the tenant on Lark, which is used to exchange for the corresponding tenant_access_token, and can also be used as the unique identifier of the tenant in the application |
|   `page_token` | `string` | Page identifier, when has_more is true, a new page_token will also be returned. Otherwise, page_token will not be returned |
|   `has_more` | `boolean` | Whether the response body has more parameters |
|   `member_total` | `int` | Amount of members | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
        "items": [
            {
                "member_id_type": "open_id",
                "member_id": "ou_9204a37300b3700d61effaa439f34295",
                "name": "John",
                "tenant_key": "736588c9260f175d"
            }
        ],
        "page_token": "dmJCRHhpd3JRbGV1VEVNRFFyTitRWDY5ZFkybmYrMEUwMUFYT0VMMWdENEtuYUhsNUxGMDIwemtvdE5ORjBNQQ==",
        "has_more": true,
        "member_total": 2
    }
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
| 400 | 232025 | Bot ability is not activated. | The bot ability needs to be enabled for the app. |
| 400 | 232033 | The operator or invited bots does NOT have the authority to manage external chats without the scope. | There is no permission to operate external groups, and application for this permission is not currently supported. |
| 400 | 232034 | The app is unavailable or inactivated by the tenant. | The app is unavailable or inactivated by the tenant. |
