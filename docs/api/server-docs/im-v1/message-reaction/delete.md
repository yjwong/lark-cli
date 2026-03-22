---
title: "Delete a reaction for a message"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/message-reaction/delete"
method: "DELETE"
api_path: "https://open.larksuite.com/open-apis/im/v1/messages/:message_id/reactions/:reaction_id"
service: "im-v1"
resource: "message-reaction"
section: "Messaging"
rate_limit: "1000 per minute、50 per second"
scopes:
  - "im:message"
  - "im:message.reactions:write_only"
updated: "1732702087000"
---

# Delete a reaction for a message

Deletes a reaction for a specified message.

## Prerequisites

- The application needs to enable bot capability.
- The robot or user calling the current interface needs to be in the conversation to which the message to be deleted belongs.

## Usage restrictions

- Retracted messages cannot add reactions.
- The robot or user calling the current interface can only delete reactions added by themselves, and must ensure that the reaction actually exists in the message.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/im/v1/messages/:message_id/reactions/:reaction_id |
| HTTP Method | DELETE |
| Rate Limit | 1000 per minute、50 per second |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `im:message` `im:message.reactions:write_only` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `message_id` | `string` | Message ID of the reaction to be deleted. How to get the ID: - After calling the Send message interface, get it from the `message_id` parameter of the response result. - Listen to the Receive message event. When the event is triggered, you can get the `message_id` of the message from the event body. - Call the Get session history messages interface and get it from the `message_id` parameter of the response result. **Example value**: "om_8964d1b4*********2b31383276113" |
| `reaction_id` | `string` | ID of the reaction to be deleted, how to get the ID: - Call the Add message reaction interface to add an emoji reply, and get it in the returned result. - Call the Get message reaction interface to get the ID of a certain emoji reply. **Example value**: "ZCaCIjUBVVWSrm5L-3ZTw*************sNa8dHVplEzzSfJVUVLMLcS_" | ## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `message.reaction` | \- |
|   `reaction_id` | `string` | Resource ID |
|   `operator` | `operator` | Adds an operator for a reaction. |
|     `operator_id` | `string` | Operator ID. The specific value is related to `operator_type`: - When `operator_type` is `app`, the robot's application ID (app_id) is returned. - When `operator_type` is `user`, the user's open_id is returned. |
|     `operator_type` | `string` | Specifies whether the operator is a user or an app. **Optional values are**:  - `app`: "app" - `user`: "user"  |
|   `action_time` | `string` | The time when the message emoticon reply was added. Unix timestamp, unit: ms |
|   `reaction_type` | `emoji` | Reaction type |
|     `emoji_type` | `string` | Emoji type. For emojis corresponding to the emoji_type value, refer to the Emojis Introduction. | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
        "reaction_id": "ZCaCIjUBVVWSrm5L-3ZTw*************sNa8dHVplEzzSfJVUVLMLcS_",
        "operator": {
            "operator_id": "ou_ff0b7ba35fb********67dfc8b885136",
            "operator_type": "user"
        },
        "action_time": "1626086391570",
        "reaction_type": {
            "emoji_type": "SMILE"
        }
    }

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 230110 | Action unavailable as the message has been deleted. | The message was deleted and cannot be operated. |
| 400 | 231001 | reaction type is invalid. | The emoji type to be operated does not belong to the current tenant and cannot be deleted. |
| 400 | 231002 | The operator has no permission to react on the specific message. | The request user or bot doesn't have the permission scope to add reactions to this message. In most cases, this error is caused because the request sender, such as the user or the bot, is not in the chat where the message resides. |
| 400 | 231003 | The message is not found, maybe not exist or deleted. | The message for which the reaction is added was not found. The message ID may be incorrect or the message may have been recalled or deleted. |
| 400 | 231004 | The chat in which the message exists is not found, maybe not exist, deleted or archived. | The chat where the message for the new reaction resides doesn't exist or has been deleted or archived. |
| 400 | 231005 | The thread has been no-trace removed, cannot put reaction. | The message to which the reaction needs to be added is a topic message and this topic message has been recalled without leaving a trace. |
| 400 | 231007 | The operator has no permission to delete this reaction. | The operator doesn't have the permission scope to delete a reaction because the reaction is not originally added by the operator. |
| 400 | 231008 | The operator has no access to the message. | The operator doesn't have the permission scope to access the message. In most cases, this error is caused because the operator isn't in the chat where the message resides. |
| 400 | 231010 | The reaction does not belong to the message. | The reaction to be deleted doesn't belong to the specified message. Check whether the reaction_id matches the message_id. |
| 400 | 231011 | The request has an invalid reaction_id. | The reaction_id parameter is invalid and the actual reaction cannot be located. Please obtain and pass in the correct reaction_id and try again. |
| 400 | 231012 | The request has an invalid pageToken. | The page_token parameter is invalid. Please set a correct value according to the page_token parameter description. |
| 400 | 231013 | The request has an invalid AuthType. | The request authorization method is invalid. Neither the tenant_access_token nor the user_access_token is used. |
| 400 | 231015 | Act on reaction failed, repeated request is processing. | A repeated reaction request is being processed. |
| 400 | 231018 | The message is invisible to the operator. | This operation cannot be performed on invisible messages. |
| 400 | 231019 | Due to the settings of this user, you are temporarily unable to perform this operation. | This operation cannot be performed due to the target user's (message sender's) settings (e.g. blocking messages). |
| 400 | 231020 | Due to the configuration of the tenant administrator, you cannot perform this operation on the target user temporarily. | Due to the tenant administrator's configuration, this operation cannot be performed on the target user (message sender) for the time being. Please contact the tenant administrator for assistance. |
| 400 | 231021 | No permission to operate external chats. | No permission to operate external chats. |
| 400 | 231022 | Bot has NO availability to this user. | The single chat user (the message recipient specified by the chat_id of the group chat, but the group chat type corresponding to the chat_id is single chat `p2p`) is not within the available range of the app bot, or is within the disabled range of the app. Solution: 1. Log in to [Developer Console](https://open.larksuite.com/app), find and enter the specified app details page. 2. In the left navigation bar, go to **App Versions** > **Version Management & Release** page, and click **Create a version**. 3. On the **Version Details** page, find the **Availability** area and click **Edit**. 4. In the dialog box that pops up, configure the available range of the app and add the user to the available range. 5. Click **Save** at the bottom of the page and publish the app to make the configuration effective. 6. (Optional) If an error message still appears after completing the above configuration, you need to contact the enterprise administrator to log in to the [Admin](https://www.larksuite.com/admin), enter the specified app details page in **Workplace** > **App Management**, and check whether the user is set as a **Blocked members** in **App availability**. For specific operations, see Configure app availability. |
| 400 | 232009 | Your request specifies a chat which has already been dissolved. | The group has been disbanded and the operation cannot be performed. | For more error code information, see General error codes.
