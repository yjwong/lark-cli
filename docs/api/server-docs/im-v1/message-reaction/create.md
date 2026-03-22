---
title: "Add a reaction for a message"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/message-reaction/create"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/im/v1/messages/:message_id/reactions"
service: "im-v1"
resource: "message-reaction"
section: "Messaging"
rate_limit: "1000 per minute、50 per second"
scopes:
  - "im:message"
  - "im:message.reactions:write_only"
updated: "1732702086000"
---

# Add a reaction for a message

Adds a reaction of the specified type for a specified message.

## Prerequisites

- The application needs to enable bot capability.
- The robot or user calling the current interface needs to be in the conversation to which the message to be reacted belongs.

## Usage restrictions

- Retracted messages cannot have reactions added.
- System messages cannot have reactions added.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/im/v1/messages/:message_id/reactions |
| HTTP Method | POST |
| Rate Limit | 1000 per minute、50 per second |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `im:message` `im:message.reactions:write_only` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `message_id` | `string` | Message ID of the reaction to be added.  How to get the ID: - After calling the Send message interface, get it from the `message_id` parameter of the response result. - Listen to the Receive message event. When the event is triggered, you can get the `message_id` of the message from the event body. - Call the Get session history messages interface and get it from the `message_id` parameter of the response result. **Example value**: "om_a8f2294b************a1a38afaac9d" | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `reaction_type` | `emoji` | Yes | Reaction type |
|   `emoji_type` | `string` | Yes | Emoji type. For supported emojis and corresponding emoji_type values, see the emoji introduction. **Example value**: "SMILE" | ### Request body example

{
    "reaction_type": {
        "emoji_type": "SMILE"
    }
}

## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `message.reaction` | \- |
|   `reaction_id` | `string` | Resource ID. After adding a reaction to a message, you will get a unique ID for the reaction. You can use this ID to delete the message reaction. |
|   `operator` | `operator` | Adds an operator for a reaction. |
|     `operator_id` | `string` | Operator ID. The specific value is related to `operator_type`: - When `operator_type` is `app`, the robot's application ID (app_id) is returned. - When `operator_type` is `user`, the user's open_id is returned. |
|     `operator_type` | `string` | Specifies whether the operator is a user or an app. **Optional values are**:  - `app`: "app" - `user`: "user"  |
|   `action_time` | `string` | The time when the message emoticon reply was added. Unix timestamp, unit: ms |
|   `reaction_type` | `emoji` | Reaction type |
|     `emoji_type` | `string` | Emoji type. For the emojis corresponding to the emoji_type value, refer to the Emoji Text Instructions. | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
        "reaction_id": "ZCaCIjUBVVWSrm5L-3ZTw*************sNa8dHVplEzzSfJVUVLMLcS_",
        "operator": {
            "operator_id": "ou_ff0b7ba35fb********67dfc8b885136",
            "operator_type": "user"
        },
        "action_time": "1663054162546",
        "reaction_type": {
            "emoji_type": "SMILE"
        }
    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 230110 | Action unavailable as the message has been deleted. | The message was deleted and cannot be operated. |
| 400 | 231001 | reaction type is invalid. | The emoji type is invalid. Please refer to the emoji introduction to set the correct emoji_type value. |
| 400 | 231002 | The operator has no permission to react on the specific message. | The request user or bot doesn't have the permission scope to add reactions to this message. In most cases, this error is caused because the request sender, such as the user or the bot, is not in the chat where the message resides. |
| 400 | 231003 | The message is not found, maybe not exist or deleted. | The message for which the reaction is added was not found. The message ID may be incorrect or the message may have been recalled or deleted. |
| 400 | 231004 | The chat in which the message exists is not found, maybe not exist, deleted or archived. | The conversation to which the target message belongs does not exist, has been dissolved or archived, and the operation cannot be performed. |
| 400 | 231005 | The thread has been no-trace removed, cannot put reaction. | The message to which the reaction needs to be added is a topic message and this topic message has been recalled without leaving a trace. The operation cannot be performed. |
| 400 | 231006 | The operator_id is invalid. | operator_id is an illegal value and cannot effectively specify a robot or user. Please try again. If the problem is still not resolved, please contact [Technical Support](https://applink.larksuite.com/TLJpeNdW). |
| 400 | 231008 | The operator has no access to the message. | The operator doesn't have the permission scope to access the message. In most cases, this error is caused because the operator isn't in the chat where the message resides. |
| 400 | 231012 | The request has an invalid pageToken. | The page_token parameter is invalid. Please set a correct value according to the page_token parameter description. |
| 400 | 231013 | The request has an invalid AuthType. | The requested authorization method is illegal. No tenant_access_token or user_access_token is used for authorization. |
| 400 | 231014 | user_id_type is invalid. | The user_id_type is invalid. One of the open_id, union_id, and user_id must be used. |
| 400 | 231015 | Act on reaction failed, repeated request is processing. | A repeated reaction request is being processed. |
| 400 | 231017 | This operation is not supported for this message type. | The message type does not support this operation. For details, please refer to Some APIs under message and group domain add unsupported message type verification. |
| 400 | 231018 | The message is invisible to the operator. | This message is not visible to the operator and this operation cannot be performed. |
| 400 | 231019 | Due to the settings of this user, you are temporarily unable to perform this operation. | This operation cannot be performed due to the target user's (message sender) settings (e.g. blocking messages). |
| 400 | 231020 | Due to the configuration of the tenant administrator, you cannot perform this operation on the target user temporarily. | Due to the tenant administrator's configuration, this operation cannot be performed on the target user (message sender) for the time being. Please contact the tenant administrator for assistance. |
| 400 | 231021 | No permission to operate external chats. | No permission to operate external chats. |
| 400 | 231022 | Bot has NO availability to this user. | The single chat user (the message recipient specified by the chat_id of the group chat, but the group chat type corresponding to the chat_id is single chat `p2p`) is not within the available range of the app bot, or is within the disabled range of the app. Solution: 1. Log in to [Developer Console](https://open.larksuite.com/app), find and enter the specified app details page. 2. In the left navigation bar, go to **App Versions** > **Version Management & Release** page, and click **Create a version**. 3. On the **Version Details** page, find the **Availability** area and click **Edit**. 4. In the dialog box that pops up, configure the available range of the app and add the user to the available range. 5. Click **Save** at the bottom of the page and publish the app to make the configuration effective. 6. (Optional) If an error message still appears after completing the above configuration, you need to contact the enterprise administrator to log in to the [Admin](https://www.larksuite.com/admin), enter the specified app details page in **Workplace** > **App Management**, and check whether the user is set as a **Blocked members** in **App availability**. For specific operations, see Configure app availability. |
| 400 | 232009 | Your request specifies a chat which has already been dissolved. | The group has been disbanded and the operation cannot be performed. | For more error code information, see General error codes.
