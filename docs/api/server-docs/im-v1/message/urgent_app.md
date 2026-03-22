---
title: "Send buzz messages in apps"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/message/urgent_app"
method: "PATCH"
api_path: "https://open.larksuite.com/open-apis/im/v1/messages/:message_id/urgent_app"
service: "im-v1"
resource: "message"
section: "Messaging"
scopes:
  - "im:message.urgent"
  - "im:message.urgent:app_send"
field_scopes:
  - "contact:user.employee_id:readonly"
updated: "1676027360000"
---

# Send buzz messages in apps

Performs in-app buzz on the specified message.

> Special notes:
> 
> - The default query limit of the API is 50 QPS. Therefore, call the API with caution.

> Notes:
> - Bot ability must be enabled.
> - Buzzing batch messages are not supported.
> - Only messages sent by the bot can be buzzed.
> - The bot must be in the chat during buzz.
> - The total number of unread buzzed messages for a user cannot exceed 200.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/im/v1/messages/:message_id/urgent_app |
| HTTP Method | PATCH |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `im:message.urgent` `im:message.urgent:app_send` |
| Required field scopes | > The response body of the API contains the following sensitive fields, and they will be returned only after corresponding scopes are added. If you do not need the fields, it is not recommended that you request the scopes. `contact:user.employee_id:readonly` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `message_id` | `string` | ID of the message to be buzzed; for the description, refer to Message ID description **Note**: batch message IDs (bm_xxx) are not supported **Example value**: "om_dc13264520392913993dd051dba21dcf" | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `user_id_type` | `string` | Yes | User ID categories **Example value**: "open_id" **Optional values are**:  - `open_id`: Identifies a user to an app. The same user has different Open IDs in different apps. How to get Open ID - `union_id`: Identifies a user to a tenant that acts as a developer. A user has the same Union ID in apps developed by the same developer, and has different Union IDs in apps developed by different developers. A developer can use Union ID to link the same user's identities in multiple apps.How to get Union ID - `user_id`: Identifies a user to a tenant. The same user has different User IDs in different tenants. In one single tenant, a user has the same User ID in all apps （including store apps）. User ID is usually used to communicate user data between different apps. How to get User ID  **Default value**: `open_id` **When the value is `user_id`, the following field scopes are required**: `contact:user.employee_id:readonly` | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `user_id_list` | `string[]` | Yes | List of target user IDs. The list can't be empty. Open ID is recommended here. For details, refer to How to get Open ID? **Note**: ensure that the user ID is correct and the user is in the group chat where the buzzed message is located. **Example value**: ["ou_6yf8af6bgb9100449565764t3382b168"] | ### Request body example

{
    "user_id_list": [
        "ou_6yf8af6bgb9100449565764t3382b168"
    ]
}

## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `invalid_user_id_list` | `string[]` | Invalid user ID | > **Note:** When some user IDs are valid, these valid users are buzzed, and invalid user IDs will be returned. When all user IDs are invalid, error code 230001 will be returned.

### Response body example

```json
{
    "code": 0,
    "msg": "success",
    "data": {
        "invalid_user_id_list": [
            "ou_6yf8af6bgb9100449565764t3382b168"
        ]
    }
}
```

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 230001 | Your request contains an invalid request parameter. | Parameter error. Please check the input parameter according to the error message returned by the interface and refer to the document. |
| 400 | 230002 | The bot can not be outside the group. | The bot isn't in the group chat. |
| 400 | 230006 | Bot ability is not activated. | Bot ability is not enabled.  In this case, log in to [Developer Console](https://open.larksuite.com/app) and choose "Features" > "Bots" to enable and launch the bot. |
| 400 | 230012 | Bot is NOT the sender of the message. | The bot isn't the message sender. |
| 400 | 230023 | The user has too many unread urgent messages. | The total number of unread buzzed messages for users cannot exceed 200, and users need to read some of the buzzed messages before they can continue to be buzzed. |
| 400 | 230024 | Reach the upper limit of urgent message. | You have reached the buzz limit. In this case, contact the tenant administrator. |
| 400 | 230027 | Lack of necessary permissions. | Please supplement the required permissions according to the permission requirements section in this API document. | > **Note:** For other server-side error codes, refer to Server-side error codes.
