---
title: "Delete Blocks"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/document-docx/docx-v1/document-block-children/batch_delete"
method: "DELETE"
api_path: "https://open.larksuite.com/open-apis/docx/v1/documents/:document_id/blocks/:block_id/children/batch_delete"
service: "docx-v1"
resource: "document-block-children"
section: "Docs"
scopes:
  - "docx:document"
updated: "1695369883000"
---

# Delete Blocks

Specifies the block that needs to be manipulated, deletes its specified range of children blocks, and returns the version of the document with the action applied.
We can regard the Upgraded Docs as a block tree, and the deleted block can be regarded as the child of its parent block, that is, it is clear that the block to be deleted is the first child of its parent block, and then this API is called to delete it.

> **Application frequency limit:** the maximum frequency of a single application call is 3 times per second, beyond which the interface will return HTTP status code 400 and error code 99991400.
> 
> **Document frequency limit**: The maximum number of concurrent edits per second for a single document is 3. If the frequency limit is exceeded, the interface will return HTTP status code 429. The edit operations include:
> 
> - Create Blocks
> - Delete Blocks
> - Update Block
> - Batch Update Blocks
> 
> When a request is frequency-limited, the application needs to handle the frequency-limited status code and use an exponential fallback algorithm or some other frequency-control strategy to reduce the rate of calls to the API.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/docx/v1/documents/:document_id/blocks/:block_id/children/batch_delete |
| HTTP Method | DELETE |
| Supported app types | custom,isv |
| Required scopes | `docx:document` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `document_id` | `string` | Unique identification of the document. Corresponding to the token of the Upgraded Docs, Click to learn how to get document_id. **Example value**: "doxcnePuYufKa49ISjhD8Ih0ikh" |
| `block_id` | `string` | Parent block's unique identity **Example value**: "doxcnO6UW6wAw2qIcYf4hZpFIth" | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `document_revision_id` | `int` | No | The document version of the operation, -1 indicates the latest version of the document. **Example value**: -1 **Default value**: `-1` **Data validation rules**: - Minimum value: `-1` |
| `client_token` | `string` | No | The unique identifier of the operation, corresponding to the client_token of the interface return value, is used for idempotent update operations. This value is null to indicate that a new request will be initiated, and this value is not null to indicate idempotent update operations. **Example value**: "Fe599b60-450f-46ff-b2ef-9f6675625b97" | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `start_index` | `int` | Yes | Deleted starting index (operation interval left closed and right open) **Example value**: 0 **Data validation rules**: - Minimum value: `0` |
| `end_index` | `int` | Yes | Deleted end index (operation range left closed and right open) **Example value**: 1 **Data validation rules**: - Minimum value: `1` | ### Request body example

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
|   `document_revision_id` | `int` | The version number of the document after the current delete operation was successful |
|   `client_token` | `string` | The unique identifier of the operation, which is used in the update request to indicate that the update is idempotent | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
        "document_revision_id": 1,
        "client_token": "fe599b60-450f-46ff-b2ef-9f6675625b97"
    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 1770001 | invalid param | Confirm whether the passed parameters are valid |
| 404 | 1770002 | not found | Confirm that the document has been deleted |
| 400 | 1770003 | resource deleted | Confirm that the resource has been deleted |
| 400 | 1770004 | too many blocks in document | Verify that the maximum number of Document Blocks is exceeded |
| 400 | 1770005 | too deep level in document | Confirm that the document Block level exceeds the upper limit |
| 400 | 1770006 | schema mismatch | Verify that the document structure is valid |
| 400 | 1770007 | too many children in block | Confirm that the maximum number of Children in the specified Block is exceeded |
| 400 | 1770008 | too big file size | Verify that the uploaded file size exceeds the upper limit |
| 400 | 1770010 | too many table column | Confirm whether the number of table columns exceeds the upper limit |
| 400 | 1770011 | too many table cell | Confirm that the number of table cells exceeds the upper limit |
| 400 | 1770012 | too many grid column | Confirm if the maximum number of Grid columns is exceeded |
| 400 | 1770013 | relation mismatch | Confirm whether the relationship between pictures, files and other resources is correct. Before inserting pictures and files, please upload the materials |
| 400 | 1770014 | parent children relation mismatch | Confirm that Block's father-son relationship is correct |
| 400 | 1770015 | single edit with multi document | Verify that the document to which the Block belongs is the same as the specified Document |
| 400 | 1770019 | repeated blockID in document | Verify Duplicate BlockIDs in Documents |
| 400 | 1770020 | operation denied on copying document | Confirm that the Document is creating a copy |
| 400 | 1770021 | too old document | Verify that the specified Document version is old |
| 400 | 1770022 | invalid page token | Confirm whether the page_token in the query parameter is valid |
| 400 | 1770024 | invalid operation | Confirm whether the operation is valid |
| 400 | 1770025 | operation and block not match | Confirm whether the corresponding operation of the specified Block application is valid |
| 400 | 1770026 | row operation over range | Confirm that row operation subscripts are out of bounds |
| 400 | 1770027 | column operation over range | Confirm that column operation subscripts are out of bounds |
| 400 | 1770028 | block not support create children | Confirm that adding Children to the specified Block is valid |
| 400 | 1770029 | block not support to create | Confirm that the specified Block supports creation |
| 400 | 1770030 | invalid parent children relation | Confirm whether the parent-child relationship of the specified operation is valid |
| 400 | 1770031 | block not support to delete children | Confirm that the specified block supports deleting Children |
| 400 | 1770033 | raw content size exceed limited | Raw content size exceed limited |
| 400 | 1770034 | operation count exceed limited | Too many cells are involved in the current request. Please split the request into multiple requests |
| 403 | 1770032 | forbidden | 1. Confirm whether the operator has permissions for the document.  Click to learn about document permissions 2. Confirm whether the operator has view permission for the mentioned document. 3. Confirm whether the mentioned user is employed and is a contact with the current operator. 4. Confirm whether the operator has permission to view and share chat cards. 5. Confirm whether the operator has permission to access the specified Wiki subdirectory. 6. Confirm whether the operator has view permissions for OKR, ISV, and Add-Ons. |
| 500 | 1771001 | server internal error | Server internal error |
| 500 | 1771006 | mount folder failed | Failed to mount document to cloud space directory |
| 500 | 1771002 | gateway server internal error | Gateway service internal error |
| 500 | 1771003 | gateway marshal error | Gateway service analysis error |
| 500 | 1771004 | gateway unmarshal error | Gateway service back analysis error |
| 503 | 1771005 | system under maintenance | System service is under maintenance |
