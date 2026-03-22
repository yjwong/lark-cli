---
title: "Obtain groups where the user or bot is a member"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat/list"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/im/v1/chats"
service: "im-v1"
resource: "chat"
section: "Group Chat"
rate_limit: "1000 per minute、50 per second"
scopes:
  - "im:chat"
  - "im:chat.group_info:readonly"
  - "im:chat:read"
  - "im:chat:readonly"
field_scopes:
  - "contact:user.employee_id:readonly"
updated: "1717574948000"
---

# Obtain groups with the user or bot

Get the list of groups where the user or bot represented by access_token is a member.

> Notes:
> - The bot ability needs to be enabled for the app.
> - Be careful not to confuse the request URLs of this API and the Obtain group information API.
> - The obtained group list does not include P2P chat groups.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/im/v1/chats |
| HTTP Method | GET |
| Rate Limit | 1000 per minute、50 per second |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `im:chat` `im:chat.group_info:readonly` `im:chat:read` `im:chat:readonly` |
| Required field scopes | > The response body of the API contains the following sensitive fields, and they will be returned only after corresponding scopes are added. If you do not need the fields, it is not recommended that you request the scopes. `contact:user.employee_id:readonly` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `user_id_type` | `string` | No | User ID categories **Example value**: open_id **Optional values are**:  - `open_id`: Identifies a user to an app. The same user has different Open IDs in different apps. How to get Open ID - `union_id`: Identifies a user to a tenant that acts as a developer. A user has the same Union ID in apps developed by the same developer, and has different Union IDs in apps developed by different developers. A developer can use Union ID to link the same user's identities in multiple apps.How to get Union ID - `user_id`: Identifies a user to a tenant. The same user has different User IDs in different tenants. In one single tenant, a user has the same User ID in all apps （including store apps）. User ID is usually used to communicate user data between different apps. How to get User ID  **Default value**: `open_id` **When the value is `user_id`, the following field scopes are required**: `contact:user.employee_id:readonly` |
| `sort_type` | `string` | No | Sort type of chat list **Example value**: ByCreateTimeAsc **Optional values are**:  - `ByCreateTimeAsc`: Sort by group creation time in ascending order - `ByActiveTimeDesc`: Sort by group active time in descending order  **Default value**: `ByCreateTimeAsc` |
| `page_token` | `string` | No | Page identifier. It is not filled in the first request, indicating traversal from the beginning; when there will be more groups, the new page_token will be returned at the same time, and the next traversal can use the page_token to get more groups **Example value**: dmJCRHhpd3JRbGV1VEVNRFFyTitRWDY5ZFkybmYrMEUwMUFYT0VMMWdENEtuYUhsNUxGMDIwemtvdE5ORjBNQQ im.v1.type.list_chat.prop.avatar.desc=$$$URL of group profile photo |
| `page_size` | `int` | No | **Example value**: 10 **Default value**: `20` **Data validation rules**: - Maximum value: `100` | > **Note:** **Notes**:
> - The query parameter **user_id_type** is used to control the owner_id type of the response body.
> - Due to the frequent changes in the active order of groups, using the `ByActiveTimeDesc` sorting method may cause group omissions. For example, if the page size is set to 10, after the first request is made to obtain the first page of data, and a member of the group originally ranked 11 sends a message, the group will be ranked first. At this time, when a request is made to obtain the data on the second page, the group cannot be obtained, and it needs to be obtained from the first page.

## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `items` | `list_chat[]` | Chat list |
|     `chat_id` | `string` | Group ID. For details, refer to Group ID description. |
|     `avatar` | `string` | URL of group profile photo |
|     `name` | `string` | Group name |
|     `description` | `string` | Group description |
|     `owner_id` | `string` | Group owner ID |
|     `owner_id_type` | `string` | Group owner ID type |
|     `external` | `boolean` | Whether it is an external group |
|     `tenant_key` | `string` | Tenant Key, which is the unique identifier of the tenant on Lark, which is used to exchange for the corresponding tenant_access_token, and can also be used as the unique identifier of the tenant in the application |
|     `chat_status` | `string` | chat status **Optional values are**:  - `normal`: normal - `dissolved`: dissolved - `dissolved_save`: dissolved and saved  |
|   `page_token` | `string` | Page identifier, when has_more is true, a new page_token will also be returned. Otherwise, page_token will not be returned |
|   `has_more` | `boolean` | Whether the response body has more parameters | ### Response body example

{"code":0,
"msg":"success",
"data":{"items":[{
    "chat_id": "oc_a0553eda9014c201e6969b478895c230",
    "avatar": "https://p3-lark-file.byteimg.com/img/lark-avatar-staging/default-avatar_44ae0ca3-e140-494b-956f-78091e348435~100x100.jpg",
    "name": "Test group name",
    "description": "Test group description",
    "owner_id": "4d7a3c6g",
    "owner_id_type": "user_id",
    "external": false,
    "tenant_key": "736588c9260f175e",
    "chat_status": "normal"
}],
"page_token":"dmJCRHhpd3JRbGV1VEVNRFFyTitRWDY5ZFkybmYrMEUwMUFYT0VMMWdENEtuYUhsNUxGMDIwemtvdE5ORjBNQQ

im.v1.type.list_chat.prop.owner_id.desc=$$$Group owner ID",
"has_more":false}}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 232001 | Your request contains an invalid request parameter. | Parameter error. Check the input parameters according to the appropriate document. |
| 400 | 232004 | Such an app does NOT exist. | The app_id as the operator doesn't exist. Please contact [technical support](https://applink.larksuite.com/TLJsX982). |
| 400 | 232025 | Bot ability is not activated. | The bot ability needs to be enabled for the app. |
| 400 | 232034 | The app is unavailable or inactivated by the tenant. | The app is unavailable or inactivated by the tenant. |
