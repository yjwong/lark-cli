---
title: "Retrieve the result of Wiki task"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/wiki-v2/task/get"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/wiki/v2/tasks/:task_id"
service: "wiki-v2"
resource: "task"
section: "Docs"
scopes:
  - "wiki:wiki"
  - "wiki:wiki:readonly"
updated: "1660708285000"
---

# Retrieve the result of Wiki task

This method is used to retrieve the result of a wiki task.

### Permission requirements ### 
 | Description | Error Message | | --- | --- | | For task creator (user or app/bot) | permission denied: only task creator can query status | ## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/wiki/v2/tasks/:task_id |
| HTTP Method | GET |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `wiki:wiki` `wiki:wiki:readonly` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" Learn more about obtaining and using access_token. | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `task_id` | `string` | Task id **Example value**: "7037044037068177428-075c9481e6a0007c1df689dfbe5b55a08b6b06f7" | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `task_type` | `string` | Yes | Task type **Example value**: "move" **Optional values are**:  - `move`: MoveDocsToWiki Task  | ## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `task` | `task_result` | Task type |
|     `task_id` | `string` | Task id |
|     `move_result` | `move_result[]` | Result for MoveDocsToWiki task |
|       `node` | `node` | Node info |
|         `space_id` | `string` | Knowledge base id |
|         `node_token` | `string` | Node token |
|         `obj_token` | `string` | Document token, which can be determined according to obj_type whether it belongs to doc, sheet or mindnote token (for shortcuts, this field is the obj_token of the corresponding entity) |
|         `obj_type` | `string` | Document type, for shortcuts, this field is the obj_type of the corresponding entity **Optional values are**:  - `doc`: Doc - `sheet`: Sheet - `mindnote`: Mindnote - `bitable`: Bitable - `file`: File - `docx`: Docx  |
|         `parent_node_token` | `string` | Node's father token. If the node is at first level, this field is empty. |
|         `node_type` | `string` | Node type **Optional values are**:  - `origin`: Entity - `shortcut`: Shortcut  |
|         `origin_node_token` | `string` | The entity node_token corresponding to the shortcut needs to pass this value when the node is created as a shortcut |
|         `origin_space_id` | `string` | The spaceid of the entity corresponding to the shortcut |
|         `has_child` | `boolean` | Whether this node has any child nodes. |
|         `title` | `string` | The title of this document. |
|         `obj_create_time` | `string` | Document creation time |
|         `obj_edit_time` | `string` | Document last edited time |
|         `node_create_time` | `string` | Node creation time |
|         `creator` | `string` | Node creator |
|         `owner` | `string` | Node owner |
|       `status` | `int` | status code for task. -1 for fail. 0 for success. 1 for processing. |
|       `status_msg` | `string` | status msg for task. | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
        "task": {
            "task_id": "7037044037068177428-075c9481e6a0007c1df689dfbe5b55a08b6b06f7",
            "move_result": [
                {
                    "node": {
                        "space_id": "6946843325487912356",
                        "node_token": "wikcnKQ1k3pcuo5uSK4t8Vabcef",
                        "obj_token": "DoccnzAaODNqyKC8g9hOWgSpprd",
                        "obj_type": "Doc/sheet/mindnote",
                        "parent_node_token": "wikcnKQ1k3pcuo5uSK4t8Vabcef",
                        "node_type": "Origin/shortcut",
                        "origin_node_token": "wikcnKQ1k3pcuo5uSK4t8Vabcef",
                        "origin_space_id": "6946843325487912356",
                        "has_child": false,
                        "title": "title",
                        "obj_create_time": "1642402428",
                        "obj_edit_time": "1642402428",
                        "node_create_time": "1642402428",
                        "creator": "ou_b13d41c02edc52ce66aaae67bf1abcef",
                        "owner": "ou_b13d41c02edc52ce66aaae67bf1abcef"
                    },
                    "status": 0,
                    "status_msg": "success"
                }
            ]
        }
    }
}

### MoveDocsToWiki Task ###

- Why is the result an array?

Future plans support batch migration of multiple nodes. Therefore, the result is represented by an array, which represents the result of mass migration of multiple nodes. For example, when three nodes are moved in, the result returns an array of length 3, where some nodes may move in successfully and some may fail.

#### status ####

| Value | Description |
| --- | --- |
| 0 | task success |
| 1 | task in progress |
| -1 | task failed | #### status_msg ####

| Value | Description |
| --- | --- |
| success | task success |
| processing | task in progress |
| already in wiki | document already in wiki |
| permission denied | permission denied |
| not support advanced bitable | wiki does not support advanced bitable permission |
| source not exist | source document does not exist |
| not support obj type | the type of document you are operating is not supported |
| tree limit | reach the capacity limit of wiki |
| failure | unknown reason failure | ### Error code
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
