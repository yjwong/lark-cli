---
title: "Search for groups visible to a user or bot"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat/search"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/im/v1/chats/search"
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

# Search for groups visible to a user or bot

Searches for the list of groups visible to a user or bot based on the access_token used, including the groups that the user or bot is in, and groups that are open to the user or bot. 
The group information that can be obtained include group ID (chat_id), group name, and group description.

> Notes:
> - The bot ability needs to be enabled for the app.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/im/v1/chats/search |
| HTTP Method | GET |
| Rate Limit | 1000 per minute、50 per second |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `im:chat` `im:chat.group_info:readonly` `im:chat:read` `im:chat:readonly` |
| Required field scopes | > The response body of the API contains the following sensitive fields, and they will be returned only after corresponding scopes are added. If you do not need the fields, it is not recommended that you request the scopes. `contact:user.employee_id:readonly` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token | > **Note:** The access token corresponds to the entity that initiates a search.

### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `user_id_type` | `string` | No | User ID categories **Example value**: open_id **Optional values are**:  - `open_id`: Identifies a user to an app. The same user has different Open IDs in different apps. How to get Open ID - `union_id`: Identifies a user to a tenant that acts as a developer. A user has the same Union ID in apps developed by the same developer, and has different Union IDs in apps developed by different developers. A developer can use Union ID to link the same user's identities in multiple apps.How to get Union ID - `user_id`: Identifies a user to a tenant. The same user has different User IDs in different tenants. In one single tenant, a user has the same User ID in all apps （including store apps）. User ID is usually used to communicate user data between different apps. How to get User ID  **Default value**: `open_id` **When the value is `user_id`, the following field scopes are required**: `contact:user.employee_id:readonly` |
| `query` | `string` | No | Keyword. **Notes**: - The keyword supports matching the internationalized name of the group and the name of the group member - Supports searching in multiple languages - Support fuzzy search such as pinyin, prefix, etc. - An empty result will be returned if the keyword is empty or longer than `64` characters **Example value**: abc |
| `page_token` | `string` | No | Page identifier. It is not filled in the first request, indicating traversal from the beginning; when there will be more groups, the new page_token will be returned at the same time, and the next traversal can use the page_token to get more groups **Example value**: dmJCRHhpd3JRbGV1VEVNRFFyTitRWDY5ZFkybmYrMEUwMUFYT0VMMWdENEtuYUhsNUxGMDIwemtvdE5ORjBNQQ im.v1.errorcode.232001.desc=$$$Your request contains an invalid request parameter. |
| `page_size` | `int` | No | **Example value**: 10 **Default value**: `20` **Data validation rules**: - Maximum value: `100` | ## Response

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

im.v1.type.list_chat.prop.description.desc=$$$Group description",
"has_more":true}}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 232001 | Your request contains an invalid request parameter. | Parameter error. Check the input parameters according to the appropriate document. |
| 400 | 232004 | Such an app does NOT exist. | The app_id as the operator doesn't exist. Please contact [technical support](https://applink.larksuite.com/TLJsX982). |
| 400 | 232025 | Bot ability is not activated. | The bot ability needs to be enabled for the app. |
| 400 | 232034 | The app is unavailable or inactivate in the tenant. | The app is unavailable or inactivate in the tenant. |
