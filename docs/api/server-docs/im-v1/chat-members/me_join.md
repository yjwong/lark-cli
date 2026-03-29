---
title: "Users or bots join a group voluntarily"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat-members/me_join"
method: "PATCH"
api_path: "https://open.larksuite.com/open-apis/im/v1/chats/:chat_id/members/me_join"
service: "im-v1"
resource: "chat-members"
section: "Group Chat"
rate_limit: "1000 per minute、50 per second"
scopes:
  - "im:chat"
  - "im:chat.members:write_only"
updated: "1717574953000"
---

# Users or bots join a group voluntarily

Users or bots join a group chat voluntarily.

> Notes:
> - The bot ability needs to be enabled.
> - Only the public groups can be joined in.
> - When operating internal group, the operator must be under the same tenant as the group.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/im/v1/chats/:chat_id/members/me_join |
| HTTP Method | PATCH |
| Rate Limit | 1000 per minute、50 per second |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `im:chat` `im:chat.members:write_only` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `chat_id` | `string` | Group ID. For details, refer to Group ID description. **Note**: - Only public group types are supported. - The maximum number of Lark groups for certified companies is: 5,000 members in general groups, 3,000 members in conference groups, and 5,000 members in topic groups. **Example value**: "oc_a0553eda9014c201e6969b478895c230" | ## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- | ### Response body example

{
    "code": 0,
    "data": {},
    "msg": "ok"
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 232004 | Such an app does NOT exist. | The app_id as the operator doesn't exist. Please contact [technical support](https://applink.larksuite.com/TLJsX982). |
| 400 | 232006 | Your request specifies a chat_id which is invalid. | Invalid chat_id, please check if the chat_id is correct. |
| 400 | 232008 | Your request specifies a chat whose type is NOT supported currently. | Unsupported group mode (chat_mode) or group type (chat_type). |
| 400 | 232009 | Your request specifies a chat which has already been dissolved. | The group has been deleted. |
| 400 | 232010 | Operator and chat can NOT be in different tenants. | The operator and the group being operated should be under the same tenant. |
| 400 | 232013 | You have reached the limit of maximum number of members a chat can have. | When joining a group, the maximum number of group members has been reached. The default maximum number of Lark groups for certified companies is: 5,000 members in general groups, 3,000 members in conference groups, and 5,000 members in topic groups. |
| 400 | 232019 | The request has been rate limited. | Rate-limiting is triggered. Pay attention to the request rate. For details, please see Frequency Control Policy. |
| 400 | 232025 | Bot ability is not activated. | The bot ability needs to be enabled for the app. |
| 400 | 232027 | There are no valid members in the ID list specified in your request. | The member ID list is empty or no valid members exist. |
| 400 | 232034 | The app is unavailable or inactivated by the tenant. | The app is unavailable or inactivated by the tenant. |
| 400 | 232075 | The public group needs to apply to join, and users or bots are not allowed to actively join the group. | When the group entry verification is enabled in the public group, you cannot join the group through the API. |
