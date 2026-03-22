---
title: "Document"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/document-docx/docx-v1/data-structure/document"
service: "docx-v1"
resource: "data-structure"
section: "Docs"
updated: "1695369867000"
---

# Document

Document represents the Upgraded Docs. The hierarchical content that can form a tree-like relationship between multiple blocks. From the structural point of view, the Upgraded Docs is a block tree, which is called Document in OpenAPI.
````json
"document": {
    "document_id": string,
    "revision_id": int,
    "title": string
}
````

| Name | Data Type | Properties | Default Value | Description |
| --------- | --------------- | :-: | :-: | --------- |
|`document_id` | `string` | required |`-`| **Document Unique ID**, and also the ID of the root Block of the document. |
|`revision_id` | `int` | required |`-`| **Document Revision ID**, used to specify the document version to query or update. If the interface is called multiple times and its return version ID has not changed, it means the document has not changed. On the contrary, a change in the version ID usually means that the document has been updated, but it should be noted that the ID change does not necessarily mean that the content of the document has changed. For example, the document may be commented by others. |
|`title` | `string` | required |`-`| **document title**, `Document` only supports returning plain text. |
