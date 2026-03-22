---
title: "Delete a group"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat/delete"
method: "DELETE"
api_path: "https://open.larksuite.com/open-apis/im/v1/chats/:chat_id"
service: "im-v1"
resource: "chat"
section: "Group Chat"
rate_limit: "1000 per minute、50 per second"
scopes:
  - "im:chat"
  - "im:chat:delete"
updated: "1717574947000"
---

# Delete a group

Deletes a group.

> Notes:
> - The bot ability needs to be enabled for the app.
> - If tenant_access_token is used, the bot must meet any of the following conditions to delete the group:
> - The bot is the group owner.
> - The bot is the creator of the group and has the Update the information of groups created by app scope.
> - If user_access_token is used, the group can be deleted only when the user is the group owner.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/im/v1/chats/:chat_id |
| HTTP Method | DELETE |
| Rate Limit | 1000 per minute、50 per second |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `im:chat` `im:chat:delete` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `chat_id` | `string` | Group ID. For details, refer to Group ID description. **Example value**: "oc_a0553eda9014c201e6969b478895c230" | ## Response

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
| 400 | 232004 | Such an app does NOT exist. | The app_id as the operator doesn't exist. Please contact [technical support](https://applink.larksuite.com/TLJsX982). |
| 400 | 232006 | Your request specifies a chat_id which is invalid. | Invalid chat_id, please check if the chat_id is correct. |
| 400 | 232008 | Your request specifies a chat whose type is NOT supported currently. | Unsupported group mode (chat_mode) or group type (chat_type). |
| 400 | 232009 | Your request specifies a chat which has already been dissolved. | The group has been deleted. |
| 400 | 232010 | Operator and chat can NOT be in different tenants. | The operator and the group being operated should be under the same tenant. |
| 400 | 232011 | Operator can NOT be out of the chat. | The operator must be a member of the chat. |
| 400 | 232017 | No Permission: If the operator is NOT owner or creator with the scope, the operator can NOT complete the request. | The operator is neither the group owner nor the group creator with permission scopes. |
| 400 | 232025 | Bot ability is not activated. | The bot ability needs to be enabled for the app. |
| 400 | 232033 | The operator or invited bots does NOT have the authority to manage external chats without the scope. | No permission to operate external groups. |
| 400 | 232034 | The app is unavailable or inactivated by the tenant. | The app is unavailable or inactivated by the tenant. |
