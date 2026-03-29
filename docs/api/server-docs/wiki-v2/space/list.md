---
title: "Get a list of Wiki spaces"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/wiki-v2/space/list"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/wiki/v2/spaces"
service: "wiki-v2"
resource: "space"
section: "Docs"
scopes:
  - "wiki:wiki"
  - "wiki:wiki:readonly"
updated: "1664358112000"
---

# Get a list of Wiki spaces

This interface is used to get a list of Wiki spaces that have permission to access.

This interface is a paging interface. Due to permission filtering, the return list may be empty, but the paging flag (has_more) is true, and can continue the paging request.

For the description of each attribute of the Wiki space, please refer to Access to Wiki space information

> When calling with tenant access token, please make sure that the application/bot has access to part of the Wiki space, otherwise the return list is easy to be empty.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/wiki/v2/spaces |
| HTTP Method | GET |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `wiki:wiki` `wiki:wiki:readonly` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" Learn more about obtaining and using access_token. | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `page_size` | `int` | No | Page size **Example value**: 10 **Data validation rules**: - Maximum value: `50` |
| `page_token` | `string` | No | Page identifier. It is not filled in the first request, indicating traversal from the beginning; when there will be more groups, the new page_token will be returned at the same time, and the next traversal can use the page_token to get more groups **Example value**: "1565676577122621" | ## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `items` | `space[]` | Data list |
|     `name` | `string` | Wiki space name |
|     `description` | `string` | Wiki space description |
|     `space_id` | `string` | Wiki space id |
|     `space_type` | `string` | Represents the Wiki space type (team space, or personal space) **Optional values are**:  - `team`: Team space - `person`: Personal space  |
|     `visibility` | `string` | Represents Wiki space visibility (public space, or private space) **Optional values are**:  - `public`: Open space - `private`: Private space  |
|   `page_token` | `string` | Page identifier, when has_more is true, a new page_token will also be returned. Otherwise, page_token will not be returned |
|   `has_more` | `boolean` | Whether the response body has more parameters | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
        "items": [
            {
                "name": "Knowledge space",
                "description": "Knowledge space description",
                "space_id": "1565676577122621"
            }
        ],
        "page_token": "1565676577122621",
        "has_more": true
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
| 400 | 131008 | Already exists | The data already exists, do not repeat the operation. |
