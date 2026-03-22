---
title: "Determine whether a user or bot is in a group"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat-members/is_in_chat"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/im/v1/chats/:chat_id/members/is_in_chat"
service: "im-v1"
resource: "chat-members"
section: "Group Chat"
rate_limit: "1000 per minute、50 per second"
scopes:
  - "im:chat"
  - "im:chat.group_info:readonly"
  - "im:chat.members:read"
  - "im:chat:readonly"
updated: "1717574953000"
---

# Determine whether a user or bot is in a group

Determines whether a user or bot is in a group based on their access_token.

> Notes:
> - The bot ability needs to be enabled.
> - When obtaining internal group information, the operator must be under the same tenant as the group.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/im/v1/chats/:chat_id/members/is_in_chat |
| HTTP Method | GET |
| Rate Limit | 1000 per minute、50 per second |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `im:chat` `im:chat.group_info:readonly` `im:chat.members:read` `im:chat:readonly` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `chat_id` | `string` | Group ID. For details, refer to Group ID description. **Example value**: "oc_a0553eda9014c201e6969b478895c230" | **Request body example**
```
curl --location --request GET 'https://open.larksuite.com/open-apis/im/v1/chats/oc_820faa21d7ed275b53d1727a0feaa917/members/is_in_chat' \
--header 'Authorization: Bearer t-39eec8560faa3dded7971873eb649fd40e70e0f1'
```

## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `is_in_chat` | `boolean` | Whether a user or bot is in a group | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
        "is_in_chat": false
    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 232004 | Such an app does NOT exist. | The app_id as the operator doesn't exist. Please contact [technical support](https://applink.larksuite.com/TLJsX982). |
| 400 | 232006 | Your request specifies a chat_id which is invalid. | Invalid chat_id, please check if the chat_id is correct. |
| 400 | 232010 | Operator and chat can NOT be in different tenants. | The operator and the group being operated should be under the same tenant. |
| 400 | 232025 | Bot ability is not activated. | The bot ability needs to be enabled for the app. |
| 400 | 232033 | The operator or invited bots does NOT have the authority to manage external chats without the scope. | No permission to operate external groups. |
| 400 | 232034 | The app is unavailable or inactivated by the tenant. | The app is unavailable or inactivated by the tenant. |
