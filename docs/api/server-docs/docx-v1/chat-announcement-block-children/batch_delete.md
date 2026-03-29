---
title: "Delete blocks in group announcement"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/document-docx/docx-v1/chat-announcement-block-children/batch_delete"
method: "DELETE"
api_path: "https://open.larksuite.com/open-apis/docx/v1/chats/:chat_id/announcement/blocks/:block_id/children/batch_delete"
service: "docx-v1"
resource: "chat-announcement-block-children"
section: "Group Chat"
rate_limit: "5 per second"
scopes:
  - "im:chat.announcement:write_only"
updated: "1739935035000"
---

# Delete blocks in group announcement

Deletes a batch of children blocks for a given parent block at a specific location. The return value of this API is the revision ID of the group announcement after the deletion.

> **Application frequency limit:** The maximum frequency of a single application call is 5 times per second, beyond which the HTTP status code 400 and error code 99991400 will be returned.
> 
> **Group announcement frequency limit**: The maximum number of concurrent edits per second for a single group announcement is 3, beyond which the HTTP status code 429 will be returned. The edit operations include:
> 
> - Create blocks
> - Batch update blocks
> - Delete blocks
> 
> When a request is frequency-limited, the application needs to handle the frequency-limited status code and use an exponential fallback algorithm or some other frequency-control strategy to reduce the rate of calls to the API.

## Prerequisites

- The application needs to enable bot capability.

- The user or robot calling the current interface must be in the corresponding group and must have the permission to read the group announcement document.

## Usage restrictions

- If the group is configured with **Only group owner and group admin can edit group info**, only group owners, group administrators, or robots that create groups and have the **Update the information of groups created by app (im:chat:operate_as_owner)** permission can update group announcement information.

- If the group is not configured with **Only group owner and group admin can edit group info**, all group members can update group announcement information.

- When operating an internal group, the operator and the operated group must be in the same tenant.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/docx/v1/chats/:chat_id/announcement/blocks/:block_id/children/batch_delete |
| HTTP Method | DELETE |
| Rate Limit | 5 per second |
| Supported app types | custom,isv |
| Required scopes | `im:chat.announcement:write_only` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `chat_id` | `string` | Group ID. How to get it: - Create a group, get the chat_id of the group from the returned result. - Call the Get the list of groups that the user or robot is in interface to query the chat_id of the group that the user or robot is in. - Call the Search the list of groups visible to the user or robot to search for the chat_id of the group that the user or robot is in and the group that is open to the user or robot. **Note**: Single chat (group type is `p2p`) does not support getting group announcements. **Example value**: "oc_5ad11d72b830411d72b836c20" |
| `block_id` | `string` | The unique identifier of the parent block. You can get the block_id of the block by calling the Get Group Announces All Blocks API. ** Note**: - This API does not support deleting the rows and columns of the table and deleting the columns of the Grid. You need to complete the relevant operations through the Batch Update Block Contents API. This API does not support deleting all child blocks of Table Cells, Grid Columns, and Callouts. **Example value**: "doxcnO6UW6wAw2qIcYf4hZpFIth" | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `revision_id` | `int` | No | The group announcement version to be queried. -1 indicates the latest version of the group announcement. Once the group announcement is created, the `revision_id` is 1. Make sure you have editing permission for the group announcement. You can use the Obtain the basic information of a group announcement api to get the latest revision ID **Example value**: -1 **Default value**: `-1` **Data validation rules**: - Minimum value: `-1` |
| `client_token` | `string` | No | The unique identifier of the operation, corresponding to the `client_token` returned by the API. It is used for idempotent update operations. When this value is null, a new request will be initiated; when this value is not null, idempotent updates will be performed. **Example value**: fe599b60-450f-46ff-b2ef-9f6675625b97 | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `start_index` | `int` | Yes | Deleted starting index (operation interval left closed right open) **Example value**: 0 **Data validation rules**: - Minimum value: `0` |
| `end_index` | `int` | Yes | Deleted end index (operation interval left closed right open) **Example value**: 1 **Data validation rules**: - Minimum value: `1` | ### Request body example

{
    "start_index": 0,
    "end_index": 1
}

## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `revision_id` | `int` | The version number of the group announcement after the current deletion operation is successful |
|   `client_token` | `string` | The unique identifier of the operation, which is used in the update request to indicate that the update is idempotent | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
        "revision_id": 1,
        "client_token": "fe599b60-450f-46ff-b2ef-9f6675625b97"
    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 1770001 | invalid param | Confirm whether the passed parameters are valid |
| 404 | 1770002 | not found | The `chat_id` of the group announcement does not exist. Please verify that the group has not been disbanded or if the `chat_id` is correctly inputted. |
| 400 | 1770003 | resource deleted | Confirm whether the resource has been deleted |
| 400 | 1770004 | too many blocks in document | Confirm whether the number of group announcement blocks exceeds the limi |
| 400 | 1770005 | too deep level in document | Confirm whether the level of group announcement blocks exceeds the limit |
| 400 | 1770006 | schema mismatch | Confirm whether the group announcement structure is valid |
| 400 | 1770007 | too many children in block | Confirm whether the number of children for a specified block exceeds the limit |
| 400 | 1770008 | too big file size | Confirm whether the uploaded file size exceeds the limit |
| 400 | 1770010 | too many table column | Confirm whether the number of table columns exceeds the limit, the upper limit is 100 columns |
| 400 | 1770011 | too many table cell | Confirm whether the number of table cells exceeds the limit, the upper limit is 2000 cells |
| 400 | 1770012 | too many grid column | Confirm whether the number of grid columns exceeds the limit, the upper limit is 10 columns |
| 400 | 1770013 | relation mismatch | Images, files, and other resources are incorrectly associated. Please make sure that you have uploaded the related image or file material to the corresponding group announcement block at the same time when you create the image or file block. |
| 400 | 1770014 | parent children relation mismatch | Confirm whether the parent-child relationships of blocks are correct |
| 400 | 1770015 | single edit with multi document | Confirm whether the block belongs to the specified group announcement |
| 400 | 1770029 | block not support to create | Confirm whether the specified block supports creation |
| 400 | 1770019 | repeated blockID in document | Confirm whether the BlockID in the group announcement is duplicated |
| 400 | 1770020 | operation denied on copying document | Verify Document is Creating a Copy |
| 400 | 1770021 | too old document | Confirm whether the specified group announcement version (revision_id) is outdated. The difference between the specified version number and the latest version number of the group announcement cannot exceed 1000 |
| 400 | 1770041 | open schema mismatch | Confirm whether the block parent-child relationship is valid |
| 400 | 1770024 | invalid operation | Verify whether the operation is valid |
| 400 | 1770025 | operation and block not match | Confirm whether the corresponding operation of the specified Block application is valid |
| 400 | 1770026 | row operation over range | Confirm whether the line operation subscript is out of bounds |
| 400 | 1770027 | column operation over range | Confirm whether the column operation subscript is out of bounds |
| 400 | 1770028 | block not support create children | Confirm whether it is valid to add Children to the specified Block |
| 400 | 1770030 | invalid parent children relation | Verify whether the parent-child relationship of the specified operation is valid |
| 400 | 1770031 | block not support to delete children | Confirm whether the specified Block supports deleting Children |
| 400 | 1770033 | raw content size exceed limited | Confirm whether plain text content size exceeds limit |
| 400 | 1770034 | operation count exceed limited | There are too many cells involved in the current request, please split it into multiple requests |
| 400 | 1770035 | resource count exceed limit | The number of resources in the current request exceeds the limit, please split it into multiple requests. The upper limit of various resources is: 200 ChatCard, 200 File, 200 MentionDoc, 200 MentionUser, 20 Image, 20 ISV, 5 Sheet and 5 Bitable. |
| 400 | 1770038 | resource not found | The inserted resource was not found or the resource has no permission to insert. Please check whether the resource identifier is correct. |
| 400 | 1772001 | Chat announcement can NOT be found in chat information. | The group announcement information is abnormal. Please retry the operation. If the issue continues, consult [Technical Support](https://applink.larksuite.com/client/helpdesk/open?id=6626260912531570952&extra=%7B%22channel%22%3A14%2C%22created_at%22%3A1614493146%2C%22scenario_id%22%3A6885151765134622721%2C%22signature%22%3A%22ca94c408b966dc1de2083e5bbcd418294c146e98%22%7D) for assistance. |
| 400 | 1772002 | Operator and chat can NOT be in different tenants. | The operator and the group being operated should be under the same tenant. |
| 400 | 1772003 | Operator can NOT be out of the chat. | The operator is not in the group. You need to add the application or user currently calling the API to the group to be operated and try again. |
| 400 | 1772004 | The operator or invited bots does NOT have the authority to manage external chats without the scope. | The app does not have permission to operate external groups. |
| 400 | 1772025 | Bot ability is not activated. | The app does not have bot capabilities enabled. You need to log in to the [Developer Console](https://open.larksuite.com/app), add the **Bot** capability in **Features** > **Add Features** on the app details page, and publish the app to make the configuration effective. For specific operations, see Bot Capabilities. |
| 400 | 1772006 | announcement type is not supported | Except for the "Obtain the basic information of a group announcement" api, other Upgraded group announcements apis cannot operate doc-type group announcements. If you want to operate doc-type group announcements, please refer to the "Old version of group announcements" api. |
| 400 | 1772005 | No Permission: Only chat owner or admin can edit chat information in the current situation. | Only the group owner or administrators are allowed to edit the group information. |
| 400 | 1772034 | The app is unavailable or inactivated by the tenant. | The app is unavailable or inactivated by the tenant. |
| 403 | 1770032 | forbidden | The current operator does not have permission to edit group announcemen. Solution: - Solution 1: Call the Specify group administrator interface to set the current operator as a group administrator, and then try again. - Solution 2: In **Lark Client > Group > Settings > Group Settings**, set **Who can edit group info** to **Everyone in this group**, and then try again. When working with APIs for creating and updating group announcement, ensure the following: - Confirm whether the current access identity has permission to read MentionDoc, i.e., @document. - Confirm whether the user in MentionUser, i.e., @user, is currently employed and is a contact of the current access identity. - Confirm whether the current access identity has permission to view and share group cards. - Confirm whether the current access identity has permission to access specific Wiki subdirectories. - Confirm whether the current access identity has permission to view document blocks such as OKR, ISV, Add-Ons, etc. |
| 500 | 1771001 | server internal error | Server internal error. Please retry the operation. If the issue continues, consult [Technical Support](https://applink.larksuite.com/client/helpdesk/open?id=6626260912531570952&extra=%7B%22channel%22%3A14%2C%22created_at%22%3A1614493146%2C%22scenario_id%22%3A6885151765134622721%2C%22signature%22%3A%22ca94c408b966dc1de2083e5bbcd418294c146e98%22%7D) for assistance. |
| 500 | 1771002 | gateway server internal error | Gateway service internal error. Please retry the operation. If the issue continues, consult [Technical Support](https://applink.larksuite.com/client/helpdesk/open?id=6626260912531570952&extra=%7B%22channel%22%3A14%2C%22created_at%22%3A1614493146%2C%22scenario_id%22%3A6885151765134622721%2C%22signature%22%3A%22ca94c408b966dc1de2083e5bbcd418294c146e98%22%7D) for assistance. |
| 500 | 1771003 | gateway marshal error | Gateway service marshal error. Please retry the operation. If the issue continues, consult [Technical Support](https://applink.larksuite.com/client/helpdesk/open?id=6626260912531570952&extra=%7B%22channel%22%3A14%2C%22created_at%22%3A1614493146%2C%22scenario_id%22%3A6885151765134622721%2C%22signature%22%3A%22ca94c408b966dc1de2083e5bbcd418294c146e98%22%7D) for assistance. |
| 500 | 1771004 | gateway unmarshal error | Gateway service unmarshalling error. Please retry the operation. If the issue continues, consult [Technical Support](https://applink.larksuite.com/client/helpdesk/open?id=6626260912531570952&extra=%7B%22channel%22%3A14%2C%22created_at%22%3A1614493146%2C%22scenario_id%22%3A6885151765134622721%2C%22signature%22%3A%22ca94c408b966dc1de2083e5bbcd418294c146e98%22%7D) for assistance. |
| 503 | 1771005 | system under maintenance | System service is under maintenance |
