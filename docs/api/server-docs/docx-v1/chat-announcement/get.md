---
title: "Obtain the basic information of a group announcement"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/document-docx/docx-v1/chat-announcement/get"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/docx/v1/chats/:chat_id/announcement"
service: "docx-v1"
resource: "chat-announcement"
section: "Group Chat"
rate_limit: "20 per second"
scopes:
  - "im:chat.announcement:read"
field_scopes:
  - "contact:user.employee_id:readonly"
updated: "1739935019000"
---

# Obtain the basic information of a group announcement

Obtain the basic information of the  specified group announcement.

## Prerequisites

- The application needs to enable bot capability.
- The user or robot calling the current interface must be in the corresponding group.
- When obtaining internal group information, the user or robot calling the current interface must be in the same tenant as the corresponding group.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/docx/v1/chats/:chat_id/announcement |
| HTTP Method | GET |
| Rate Limit | 20 per second |
| Supported app types | custom,isv |
| Required scopes | `im:chat.announcement:read` |
| Required field scopes | > The response body of the API contains the following sensitive fields, and they will be returned only after corresponding scopes are added. If you do not need the fields, it is not recommended that you request the scopes. `contact:user.employee_id:readonly` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `chat_id` | `string` | Group ID. How to get it: - Create a group, get the chat_id of the group from the returned result. - Call the Get the list of groups that the user or robot is in interface to query the chat_id of the group that the user or robot is in. - Call the Search the list of groups visible to the user or robot to search for the chat_id of the group that the user or robot is in and the group that is open to the user or robot. **Note**: Single chat (group type is `p2p`) does not support getting group announcements. **Example value**: "oc_5ad11d72b830411d72b836c20" | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `user_id_type` | `string` | No | User ID categories **Example value**: open_id **Optional values are**:  - `open_id`: Identifies a user to an app. The same user has different Open IDs in different apps. How to get Open ID - `union_id`: Identifies a user to a tenant that acts as a developer. A user has the same Union ID in apps developed by the same developer, and has different Union IDs in apps developed by different developers. A developer can use Union ID to link the same user's identities in multiple apps.How to get Union ID - `user_id`: Identifies a user to a tenant. The same user has different User IDs in different tenants. In one single tenant, a user has the same User ID in all apps （including store apps）. User ID is usually used to communicate user data between different apps. How to get User ID  **Default value**: `open_id` **When the value is `user_id`, the following field scopes are required**: `contact:user.employee_id:readonly` | ## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `revision_id` | `int` | Current group announcement version number |
|   `create_time` | `string` | Timestamp of  group announcement creation (in seconds) |
|   `update_time` | `string` | Timestamp of  group announcement updated (in seconds) |
|   `owner_id` | `string` | The group announcement owner ID. The ID value corresponds to the ID type of the owner_id_type. |
|   `owner_id_type` | `string` | Group announcement owner ID type **Optional values are**:  - `user_id`: dentifies a user's identity in a tenant. The User ID of the same user in tenant A and tenant B is different. In the same tenant, a user's User ID is consistent in all applications (including store applications). User ID is mainly used to connect user data between different applications. - `union_id`: Identifies a user's identity under a certain application developer. The Union ID of the same user in applications under the same developer is the same, and the Union ID in applications under different developers is different. Through the Union ID, application developers can associate the identities of the same user in multiple applications. - `open_id`: Identifies a user's identity in a certain application. The Open ID of the same user in different applications is different.  |
|   `modifier_id` | `string` | Last group announcement modifier ID. The ID value corresponds to the ID type of modifier_id_type. |
|   `modifier_id_type` | `string` | ID type of the last group announcement modifier **Optional values are**:  - `user_id`: Identifies a user's identity in a tenant. The User ID of the same user in tenant A and tenant B is different. In the same tenant, a user's User ID is consistent in all applications (including store applications). User ID is mainly used to connect user data between different applications. - `union_id`: Identifies a user's identity under a certain application developer. The Union ID of the same user in applications under the same developer is the same, and the Union ID in applications under different developers is different. Through the Union ID, application developers can associate the identities of the same user in multiple applications. - `open_id`: Identifies a user's identity in a certain application. The Open ID of the same user in different applications is different.  |
|   `announcement_type` | `string` | Group announcement type **Optional values are**:  - `docx`: Upgraded group announcement - `doc`: Old version group announcement  | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
        "revision_id": 1,
        "create_time": "1609296809",
        "update_time": "1609296809",
        "owner_id": "ou_7d8a6e6df7621556ce0d21922b676706ccs",
        "owner_id_type": "user_id",
        "modifier_id": "ou_7d8a6e6df7621556ce0d21922b676706ccs",
        "modifier_id_type": "user_id",
        "announcement_type": "docx"
    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 1770001 | invalid param | Parameter error. Refer to the parameter description provided in the interface document to check whether the input parameters are correct. |
| 404 | 1770002 | not found | The group announcement does not exist, please check that the incoming chat_id is correct. |
| 400 | 1772001 | Chat announcement can NOT be found in chat information. | The group announcement information is abnormal. Please retry the operation. If the issue continues, consult [Technical Support](https://applink.larksuite.com/client/helpdesk/open?id=6626260912531570952&extra=%7B%22channel%22%3A14%2C%22created_at%22%3A1614493146%2C%22scenario_id%22%3A6885151765134622721%2C%22signature%22%3A%22ca94c408b966dc1de2083e5bbcd418294c146e98%22%7D) for assistance. |
| 400 | 1772002 | Operator and chat can NOT be in different tenants. | The operator and the group being operated should be under the same tenant. |
| 400 | 1772003 | Operator can NOT be out of the chat. | The operator is not in the group. You need to add the application or user currently calling the API to the group to be operated and try again. |
| 400 | 1772004 | The operator or invited bots does NOT have the authority to manage external chats without the scope. | The app does not have permission to operate external groups. |
| 400 | 1772025 | Bot ability is not activated. | The app does not have bot capabilities enabled. You need to log in to the [Developer Console](https://open.larksuite.com/app), add the **Bot** capability in **Features** > **Add Features** on the app details page, and publish the app to make the configuration effective. For specific operations, see Bot Capabilities. |
| 400 | 1772006 | announcement type is not supported | Except for the "Obtain the basic information of a group announcement" api, other Upgraded group announcements apis cannot operate doc-type group announcements. If you want to operate doc-type group announcements, please refer to the "Old version of group announcements" api. |
| 400 | 1772034 | The app is unavailable or inactivated by the tenant. | The app is unavailable or inactivated by the tenant. |
| 403 | 1770032 | forbidden | The operator does not have read or edit permissions for group announcements. You need to add document read or edit permissions for the app or user currently calling the API and try again. |
| 500 | 1771001 | server internal error | Internal server error. Please retry the operation. If the issue continues, consult [Technical Support](https://applink.larksuite.com/client/helpdesk/open?id=6626260912531570952&extra=%7B%22channel%22%3A14%2C%22created_at%22%3A1614493146%2C%22scenario_id%22%3A6885151765134622721%2C%22signature%22%3A%22ca94c408b966dc1de2083e5bbcd418294c146e98%22%7D) for assistance. |
| 500 | 1771002 | gateway server internal error | Internal server error. Please retry the operation. If the issue continues, consult [Technical Support](https://applink.larksuite.com/client/helpdesk/open?id=6626260912531570952&extra=%7B%22channel%22%3A14%2C%22created_at%22%3A1614493146%2C%22scenario_id%22%3A6885151765134622721%2C%22signature%22%3A%22ca94c408b966dc1de2083e5bbcd418294c146e98%22%7D) for assistance. |
| 500 | 1771003 | gateway marshal error | Service Serialization Resolution Error. Please retry the operation. If the issue continues, consult [Technical Support](https://applink.larksuite.com/client/helpdesk/open?id=6626260912531570952&extra=%7B%22channel%22%3A14%2C%22created_at%22%3A1614493146%2C%22scenario_id%22%3A6885151765134622721%2C%22signature%22%3A%22ca94c408b966dc1de2083e5bbcd418294c146e98%22%7D) for assistance. |
| 500 | 1771004 | gateway unmarshal error | Service deserialization parsing error. Please retry the operation. If the issue continues, consult [Technical Support](https://applink.larksuite.com/client/helpdesk/open?id=6626260912531570952&extra=%7B%22channel%22%3A14%2C%22created_at%22%3A1614493146%2C%22scenario_id%22%3A6885151765134622721%2C%22signature%22%3A%22ca94c408b966dc1de2083e5bbcd418294c146e98%22%7D) for assistance. |
| 503 | 1771005 | system under maintenance | System service is under maintenance | For more error code information, see General error codes.
