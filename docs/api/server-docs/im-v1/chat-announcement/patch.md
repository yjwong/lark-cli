---
title: "Update group announcement info"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat-announcement/patch"
method: "PATCH"
api_path: "https://open.larksuite.com/open-apis/im/v1/chats/:chat_id/announcement"
service: "im-v1"
resource: "chat-announcement"
section: "Group Chat"
rate_limit: "1000 per minute、50 per second"
scopes:
  - "im:chat"
  - "im:chat.announcement:write_only"
updated: "1717574957000"
---

# Update group announcement info

Updates the group announcement information in a chat, with the same format as Docs, newer document formats are not supported.

> Notes:
> - The bot ability needs to be enabled.
> - The bot or authorized user must be in the group.
> - The operator needs to have the read permission of the group announcement document.
> - When obtaining internal group information, the operator must be under the same tenant as the group.
> - If Only group owner and group admin can edit group info is enabled for the group, the group owner/administrators or the bot that creates the group and has the Update the information of groups created by app scope can update the group announcement.
> - If Only group owner and group admin can edit group info is not enabled for the group, all the members can update the group announcement.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/im/v1/chats/:chat_id/announcement |
| HTTP Method | PATCH |
| Rate Limit | 1000 per minute、50 per second |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `im:chat` `im:chat.announcement:write_only` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `chat_id` | `string` | The ID of the group with its announcement to be modified. For details, refer to Group ID description **Note**: P2P chat is not supported. **Example value**: "oc_5ad11d72b830411d72b836c20" | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `revision` | `string` | Yes | int64 type of the current document version number, which is returned by calling the "get" API. **Note**：The difference between the incoming version number and the latest version number cannot exceed 100. **Example value**: "12" |
| `requests` | `string[]` | No | Serialization field of the document modification request Updates the group announcement information in a chat, with the same format as Docs **Example value**: ["{\"requestType\":\"InsertBlocksRequestType\",\"insertBlocksRequest\":{\"payload\":\"{\\\"blocks\\\":[{\\\"type\\\":\\\"paragraph\\\",\\\"paragraph\\\":{\\\"elements\\\":[{\\\"type\\\":\\\"textRun\\\",\\\"textRun\\\":{\\\"text\\\":\\\"Docs API Sample Content\\\",\\\"style\\\":{}}}],\\\"style\\\":{}}}]}\",\"location\":{\"zoneId\":\"0\",\"index\":0, \"endOfZone\": true}}}"] | ### Request body example

{
    "revision": "12",
    "requests": [
        "{\"requestType\":\"InsertBlocksRequestType\",\"insertBlocksRequest\":{\"payload\":\"{\\\"blocks\\\":[{\\\"type\\\":\\\"paragraph\\\",\\\"paragraph\\\":{\\\"elements\\\":[{\\\"type\\\":\\\"textRun\\\",\\\"textRun\\\":{\\\"text\\\":\\\"Docs API Sample Content\\\",\\\"style\\\":{}}}],\\\"style\\\":{}}}]}\",\"location\":{\"zoneId\":\"0\",\"index\":0, \"endOfZone\": true}}}"
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
| 400 | 232002 | No Permission: Only chat owner or admin can edit chat information in the current situation. | Only the group owner or administrators are allowed to edit the group information. |
| 400 | 232003 | Chat announcement can NOT be found in chat information. | The group announcement information is abnormal. |
| 400 | 232009 | Your request specifies a chat which has already been dissolved. | The group has been deleted. |
| 400 | 232010 | Operator and chat can NOT be in different tenants. | The operator and the group being operated should be under the same tenant. |
| 400 | 232011 | Operator can NOT be out of the chat. | The operator must be a member of the chat. |
| 400 | 232018 | Updating announcement failed. | The announcement update failed. Check the error message. For details, refer to Edit a document. |
| 400 | 232019 | The request has been rate limited. | Rate-limiting is triggered. Pay attention to the request rate. For details, please see Frequency Control Policy. |
| 400 | 232025 | Bot ability is not activated. | The bot ability needs to be enabled for the app. |
| 400 | 232033 | The operator or invited bots does NOT have the authority to manage external chats without the scope. | The app does not have permission to operate external groups. |
| 400 | 232034 | The app is unavailable or inactivated by the tenant. | The app is unavailable or inactivated by the tenant. |
| 400 | 232066 | The operator does not have the doc permission. | The operator does not have the permission to read the document. Please add the document permission and try again. |
| 400 | 232097 | Unable to operate docx type chat announcement. | Unable to operate docx type chat announcement. Please refer to the update new version group announcement API. |
