---
title: "Add an existing document to Wiki space"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/wiki-v2/space-node/move_docs_to_wiki"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/wiki/v2/spaces/:space_id/nodes/move_docs_to_wiki"
service: "wiki-v2"
resource: "space-node"
section: "Docs"
scopes:
  - "wiki:wiki"
updated: "1660708284000"
---

# Add an existing document to Wiki space

This interface allows to add existing cloud documents to the knowledge base and mount them under the specified parent page

> ### Move operation ###
> 
> After moving, the document will be transferred from "MySpace" or "Shared Space" to "Knowledge Base" and will disappear from the following function entrances:
> 
> - Cloud Space Home Page: Recent Access, Quick Access
> 
> - MySpace
> 
> - Shared Space
> 
> - Favorites
> 
> ### Permission Change ###
> 
> After moving, the document will be displayed to all users who can view the "Page Tree". By default, the permission settings of the parent page will be inherited.

> This method is asynchronous. If the move is complete (or the node is already in the Wiki), the result (Wiki token) is returned directly. If not completed, the task id is returned. Please use the Get Task Results method to query.
> 
> Wiki permission requirements:
> - Document Manageable Permission
> - Original Folder Edit Permission
> - Target Parent Container Edit Permission

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/wiki/v2/spaces/:space_id/nodes/move_docs_to_wiki |
| HTTP Method | POST |
| Supported app types | custom,isv |
| Required scopes | `wiki:wiki` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" Learn more about obtaining and using access_token. |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `space_id` | `string` | Wiki space id **Example value**: "1565676577122621" | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `parent_wiki_token` | `string` | No | Father token of node **Example value**: "WikbcOHIFxB0PJS2UTd2kF2SP6c" |
| `obj_type` | `string` | Yes | Document type **Example value**: "Doc" **Optional values are**:  - `doc`: Documentation - `sheet`: Sheet - `bitable`: Bitable - `mindnote`: Mindnote - `docx`: Docx - `file`: File  |
| `obj_token` | `string` | Yes | Document token **Example value**: "Docbc6e1qBqt1O5mCBVA1QUKVEg" |
| `apply` | `boolean` | No | Whether to apply for moving in documents when there is no permission **Example value**: True | ### Request body example

{
    "parent_wiki_token": "wikbcOHIFxB0PJS2UTd2kF2SP6c",
    "obj_type": "doc",
    "obj_token": "docbc6e1qBqt1O5mCBVA1QUKVEg"
}

## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `wiki_token` | `string` | Wiki token after moving |
|   `task_id` | `string` | Task id |
|   `applied` | `boolean` | Whether a document migration request has been submitted | ### Response body example

// The operation is completed
{
    "code": 0,
    "data": {
        "wiki_token": "wikbcLZuhp4r9QuJumHzV2fzF7T"
    },
    "msg": "success"
}
// Or the operation is not completed yet
{
    "code": 0,
    "data": {
        "task_id": "7037044037068177428-075c9481e6a0007c1df689dfbe5b55a08b6b06f7"
    },
    "msg": "success"
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
