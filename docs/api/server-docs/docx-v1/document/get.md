---
title: "Query Document"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/document-docx/docx-v1/document/get"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/docx/v1/documents/:document_id"
service: "docx-v1"
resource: "document"
section: "Docs"
rate_limit: "Special Rate Limit"
scopes:
  - "docx:document"
  - "docx:document:readonly"
updated: "1695369871000"
---

# Query Document

Query the latest version number, title, etc. of the document

> **Application frequency limit:** the maximum frequency of a single application call is 5 times per second, beyond which the interface will return HTTP status code 400 and error code 99991400. When a request is frequency-limited, the application needs to handle the frequency-limited status code and use an exponential fallback algorithm or some other frequency-control strategy to reduce the rate of calls to the API.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/docx/v1/documents/:document_id |
| HTTP Method | GET |
| Rate Limit | Special Rate Limit |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `docx:document` `docx:document:readonly` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `document_id` | `string` | Unique identification of the document. Corresponding to the token of the Upgraded Docs, Click to learn how to get document_id. **Example value**: "doxcnePuYufKa49ISjhD8Ih0ikh" **Data validation rules**: - Length range: `27` ～ `27` characters | ## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `document` | `document` | Document Information |
|     `document_id` | `string` | Unique identification of the document. Corresponding to the token of the Upgraded Docs, Click to learn how to get document_id. |
|     `revision_id` | `int` | Document Version ID |
|     `title` | `string` | Document Title | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
        "document": {
            "document_id": "doxcni6mOy7jLRWbEylaKKC7K88",
            "revision_id": 1,
            "title": "title"
        }
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
| 403 | 1770032 | forbidden | Confirm whether the operator has permissions for the document. Click to learn about document permissions |
| 500 | 1771001 | server internal error | Server internal error |
| 500 | 1771006 | mount folder failed | Failed to mount document to cloud space directory |
| 500 | 1771002 | gateway server internal error | Gateway service internal error |
| 500 | 1771003 | gateway marshal error | Gateway service analysis error |
| 500 | 1771004 | gateway unmarshal error | Gateway service back analysis error |
| 503 | 1771005 | system under maintenance | System service is under maintenance |
