---
title: "Update group pinning"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat-top_notice/put_top_notice"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/im/v1/chats/:chat_id/top_notice/put_top_notice"
service: "im-v1"
resource: "chat-top_notice"
section: "Group Chat"
rate_limit: "1000 per minute、50 per second"
scopes:
  - "im:chat"
  - "im:chat.top_notice:write_only"
updated: "1732702087000"
---

# Update group pinning

Updates the pinned information in a group chat. A group message or announcement can be pinned.

## Prerequisites

The app needs to enable bot capability.

## Usage restrictions

The bot or user that calls the interface must be in the group and belong to the same tenant as the group.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/im/v1/chats/:chat_id/top_notice/put_top_notice |
| HTTP Method | POST |
| Rate Limit | 1000 per minute、50 per second |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `im:chat` `im:chat.top_notice:write_only` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `chat_id` | `string` | Group ID. How to get it: - Create a group, get the chat_id of the group from the returned result. - Call the Get the list of groups that the user or bot is in interface to query the chat_id of the group that the user or bot is in. - Call the Search the list of groups visible to the user or bot to search for the chat_id of the group that the user or bot is in and the group that is open to the user or bot. **Example value**: "oc_5ad11d72b830411d72b836c20" | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `chat_top_notice` | `chat.top_notice[]` | Yes | Group Pin configuration |
|   `action_type` | `string` | No | Pin type **Example value**: "2" **Optional values are**:  - `1`: Message type, `message_id` is required - `2`: Group announcement type, no need to fill in `message_id`  **Default value**: `2` |
|   `message_id` | `string` | No | Message ID. How to get the ID: - After calling the Send Message interface, get it from the `message_id` parameter of the response result. - Listen to the Receive Message event. When the event is triggered, you can get the `message_id` of the message from the event body. - Call the Get Session History Messages interface and get it from the `message_id` parameter of the response result. **Example value**: "om_dc13264520392913993dd051dba21dcf" | ### Request body example

{
    "chat_top_notice": [
        {
            "action_type": "2",
            "message_id": "om_dc13264520392913993dd051dba21dcf"
        }
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
| 400 | 230111 | Action unavailable as the message will self-destruct soon. | Action unavailable as the message will self-destruct soon. |
| 400 | 232001 | Your request contains an invalid request parameter. | Parameter error. Check the input parameters according to the appropriate document. |
| 400 | 232009 | Your request specifies a chat which has already been dissolved. | The group has been deleted. |
| 400 | 232010 | Operator and chat can NOT be in different tenants. | The operator and the group being operated should be under the same tenant. |
| 400 | 232011 | Operator can NOT be out of the chat. | The operator is not in the group. You need to add the application or user currently calling the API to the group to be operated and try again. |
| 400 | 232014 | The token used in the request does NOT have the permissions necessary to complete the request. | The operator does not have permission to perform this operation, please check whether it has the corresponding permission. |
| 400 | 232024 | Users do not have the visibility of the app, or the operator does not have collaboration permissions with the target users. | The robot is not visible to the user, or the operator does not have collaboration permissions with the user. - If the robot is not visible to the user, you need to edit the visibility of the application to the user and publish the application in [Developer Console](https://open.larksuite.com/app) > **Application Details Page** > **App Versions** > **Version Management & Release**. For specific operations, refer to Configure Application Availability. - If the operator does not have collaboration permissions with the user, please check whether you have collaboration permissions with the target user, such as blocking or not adding as a contact. |
| 400 | 232025 | Bot ability is not activated. | The app does not have bot capabilities enabled. You need to log in to the [Developer Console](https://open.larksuite.com/app), add the **Bot** capability in **Features** > **Add Features** on the app details page, and publish the app to make the configuration effective. For specific operations, see Bot Capabilities. |
| 400 | 232033 | The operator or invited bots does NOT have the authority to manage external chats without the scope. | No permission to operate external groups. |
| 400 | 232034 | The app is unavailable or inactivate in the tenant. | The app is unavailable or inactivate in the tenant. | For more error code information, see General error codes.
