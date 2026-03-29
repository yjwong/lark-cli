---
title: "Access to Wiki space information"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/wiki-v2/space/get"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/wiki/v2/spaces/:space_id"
service: "wiki-v2"
resource: "space"
section: "Docs"
scopes:
  - "wiki:wiki"
  - "wiki:wiki:readonly"
updated: "1664358112000"
---

# Access to Wiki space information

This interface is used to query the information of the Wiki space according to the Wiki space ID.

Space type:
-  Person Space: Managed by individuals. One person can only have one personal space, and no other administrators can be added.
- Team Space:  Managed by a team (multiple people), multiple administrators can be added.

Space visibility:
-  Public Space: Visible to all users within the tenant and defaults to member permissions. Additional members cannot be added, but administrators can be added.
- Private Space:  Only visible to knowledge space administrators and members, administrators and members need to be added manually.

> Wiki permission requirements:
> - Is Wiki space member

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/wiki/v2/spaces/:space_id |
| HTTP Method | GET |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `wiki:wiki` `wiki:wiki:readonly` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" Learn more about obtaining and using access_token. | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `space_id` | `string` | Knowledge space id **Example value**: "6870403571079249922" | ## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `space` | `space` | Wiki space |
|     `name` | `string` | Wiki space name |
|     `description` | `string` | Wiki space description |
|     `space_id` | `string` | Wiki space id |
|     `space_type` | `string` | Represents the Wiki space type (team space, or personal space) **Optional values are**:  - `team`: Team space - `person`: Personal space  |
|     `visibility` | `string` | Represents Wiki space visibility (public space, or private space) **Optional values are**:  - `public`: Open space - `private`: Private space  | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
        "space": {
            "name": "Knowledge space",
            "description": "Knowledge space description",
            "space_id": "1565676577122621"
        }
    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 131001 | RPC failed | Call the downstream error, take the x-tt-logid location in the header of the return value |
| 400 | 131002 | param err | Usually, the parameter is incorrect, such as the data type does not match. Please check the **specific interface error information**. If the error is not clear, please consult oncall. |
| 400 | 131003 | out of limit | Operational limits, such as number of nodes, were exceeded. See the table below. - The total number of nodes in the original/target wiki space does not exceed 400,000. - The original/target wiki space directory tree does not exceed 50 layers. - The number of single-layer nodes under the destination parent node does not exceed 2000. - The number of single mobile nodes (with child nodes) does not exceed 2000. |
| 400 | 131004 | invalid user | Illegal user |
| 400 | 131005 | not found | No relevant data was found, for example id does not exist. Reference for related error information: - member not found: The user is not a member (admin) of the wiki space and cannot be deleted. - identity not found: userid does not exist, cannot add/remove members. - space not found: wiki space does not exist - node not found: node does not exist - document not found: document does not exist Please consult oncall when the error is unclear. |
| 400 | 131006 | permission denied | Permission denied, refer to the relevant error information: - wiki space permission denied: need to be a wiki space member (administrator) - node permission denied: Node read permission is required for read operations. Write operations (create, move, etc.) require node container edit permissions. - no source parent node permissions: original parent node container edit permissions are required. - no destination parent node permission: requires destination parent node container edit permission. - only task creators can query status: for task creators (user or app/bot) If using tenant_access_token call, make sure the app/bot is a member of the wiki space. See How to add an app as a wiki base administrator (member). When the interface error message is not clear, please consult oncall. |
| 400 | 131007 | internal err | Internal error, take the x-tt-logid location in the header of the return value |
