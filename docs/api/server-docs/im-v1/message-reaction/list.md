---
title: "List message reactions"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/message-reaction/list"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/im/v1/messages/:message_id/reactions"
service: "im-v1"
resource: "message-reaction"
section: "Messaging"
rate_limit: "1000 per minute、50 per second"
scopes:
  - "im:message.reactions:read"
  - "im:message:readonly"
field_scopes:
  - "contact:user.employee_id:readonly"
updated: "1717574934000"
---

# List message reactions

Obtains a list of reactions of the specified type for a specified message.

> Notes:
> - Bot ability must be enabled.
> - The message whose reaction information you need to obtain must exist and not be recalled.
> - To obtain the reaction for a message, the authorizer, such as the bot or the user, of the request must be in the chat where the message resides.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/im/v1/messages/:message_id/reactions |
| HTTP Method | GET |
| Rate Limit | 1000 per minute、50 per second |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `im:message.reactions:read` `im:message:readonly` |
| Required field scopes | > The response body of the API contains the following sensitive fields, and they will be returned only after corresponding scopes are added. If you do not need the fields, it is not recommended that you request the scopes. `contact:user.employee_id:readonly` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `message_id` | `string` | Message ID of the reaction to be obtained. For the description, refer to Message ID description. **Example value**: "om_8964d1b4*********2b31383276113" | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `reaction_type` | `string` | No | Type of the message reaction to be queriedreaction type. **Note**: If this parameter is not filled in, all types of reactions will be obtained. **Example value**: LAUGH |
| `page_token` | `string` | No | Page identifier. It is not filled in the first request, indicating traversal from the beginning; when there will be more groups, the new page_token will be returned at the same time, and the next traversal can use the page_token to get more groups **Example value**: YhljsPiGfUgnVAg9urvRFd-BvSqRL20wMZNAWfa9xXkud6UKCybPuUgQ1vM26dj6 |
| `page_size` | `int` | No | **Example value**: 10 **Data validation rules**: - Maximum value: `50` |
| `user_id_type` | `string` | No | User ID categories **Example value**: open_id **Optional values are**:  - `open_id`: Identifies a user to an app. The same user has different Open IDs in different apps. How to get Open ID - `union_id`: Identifies a user to a tenant that acts as a developer. A user has the same Union ID in apps developed by the same developer, and has different Union IDs in apps developed by different developers. A developer can use Union ID to link the same user's identities in multiple apps.How to get Union ID - `user_id`: Identifies a user to a tenant. The same user has different User IDs in different tenants. In one single tenant, a user has the same User ID in all apps （including store apps）. User ID is usually used to communicate user data between different apps. How to get User ID  **Default value**: `open_id` **When the value is `user_id`, the following field scopes are required**: `contact:user.employee_id:readonly` | ## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `items` | `message.reaction[]` | Queries a reaction list based on the specified reaction_type. |
|     `reaction_id` | `string` | Resource ID |
|     `operator` | `operator` | Adds an operator for a reaction. |
|       `operator_id` | `string` | Operator ID |
|       `operator_type` | `string` | Specifies whether the operator is a user or an app. **Optional values are**:  - `app`: "app" - `user`: "user"  |
|     `action_time` | `string` | The Unix timestamp of the reaction in ms. |
|     `reaction_type` | `emoji` | Reaction type |
|       `emoji_type` | `string` | Emoji type. For more information, see Emoji type enumeration. |
|   `has_more` | `boolean` | Whether the response body has more parameters |
|   `page_token` | `string` | Page identifier, when has_more is true, a new page_token will also be returned. Otherwise, page_token will not be returned | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
        "items": [
            {
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
        ],
        "has_more": true,
        "page_token": "YhljsPiGfUgnVAg9urvRFd-BvSqRL*********Wfa9xXkud6UKCybPuUgQ1vM26dj6"
    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 230110 | Action unavailable as the message has been deleted. | Action unavailable as the message has been deleted. |
| 400 | 231001 | reaction type is invalid. | The reaction type is invalid. Make sure that the emoji type is correct. |
| 400 | 231003 | The message is not found, maybe not exist or deleted. | The message for which the reaction is added was not found. The message ID may be incorrect or the message may have been recalled or deleted. |
| 400 | 231008 | The operator has no access to the message. | The operator doesn't have the permission scope to access the message. In most cases, this error is caused because the operator isn't in the chat where the message resides. |
| 400 | 231012 | The request has an invalid pageToken. | The pageToken parameter is invalid. |
| 400 | 231013 | The request has an invalid AuthType. | The request authorization method is invalid. Neither the tenant_access_token nor the user_access_token is used. |
| 400 | 231014 | user_id_type is invalid. | The user_id_type is invalid. One of the open_id, union_id, and user_id must be used. |
| 400 | 231018 | The message is invisible to the operator. | This operation cannot be performed on invisible messages. |
