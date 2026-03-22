---
title: "Document overview"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/document-docx/docx-overview"
section: "Docs"
updated: "1733381760000"
---

# Document overview

This document introduces the basic concepts, limitations, access process, and API list of the document capabilities in the Lark Open Platform Docs.

> This document introduces the capabilities of Upgraded Docs. To understand the differences between Upgraded Docs and Docs features, see Document version changes.

## Basic concepts

There are two basic concepts in the document APIs: document and block.

### Document

A document is an online document created by the user in the Docs. Each document has a unique `document_id` as its identifier. As shown below, you can open the document in your browser and retrieve the `document_id` from the address bar.

> For documents in the Wiki, you need to call the Get Wiki node information API to obtain the `obj_token` for the document. In this context, the `obj_token` serves as the `document_id` for that specific document.

![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/fdbc92b302594d28c5fa63c029327645_Azqt8eY6Mq.png?lazyload=true&width=2218&height=479)
The basic metadata structure of a document is as follows:

```json
"document": {
    "document_id": string, // The unique identifier of the document.
    "revision_id": int,   // The identifier for the document version, which can be used to query or update a particular document version.
    "title": string, // The title of the document.
    "display_setting": { // The display setting of the document.
        "show_authors": boolean, // Whether the author is displayed in the document details.
        "show_comment_count": boolean, // Whether the total number of comments is displayed in the document details.
        "show_create_time": boolean, // Whether the creation time is displayed in the document details.
        "show_like_count": boolean, // Whether the total number of likes is displayed in the document details.
        "show_pv": boolean, // Whether the pv is displayed in the document details.
        "show_uv": boolean // Whether the uv is displayed in the document details.
    },
    "cover": {  // The cover of the document.
        "token": string, // Image token
        "offset_ratio_x": float, // The view's horizontal offset ratio.
        "offset_ratio_y": float // The view's vertical offset ratio.
    }
}
```

### Block

Within a document, there are multiple types of paragraphs, each defined as a block. A block is the smallest building unit in a document, serving as a structured component with a clear meaning. Blocks can take various forms, such as a text paragraph, a sheet, an image, or a Base. Each block has a unique `block_id` for identification.

Every document has a root block called Page block. The `block_id` of the Page block is identical to the `document_id` of the document it belongs to. In the data structure, the Page block forms a parent-child relationship with other blocks, where the Page block is the parent (referred to as Parent), and other blocks are children (referred to as Children). Parent-child relationships can also exist among other blocks.

![](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/9f5fdc25d4172bbf3cb312405c799fe7_SERAIx3wHQ.png?height=1981&lazyload=true&width=3164)

![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/65fd5d1dd1695c43daafcf2b95d57689_SYSTPEiZic.png?height=548&lazyload=true&width=1497)

**Block category**

From a functional perspective, blocks can be categorized into the following types. To know about the specific blocks, refer to BlockType enums.
| Functional category | Typical blocks | Example |
| --- | --- | --- |
| Text type | Page, Text, Heading, Bullet list, Ordered list, Code, Todo, etc. |  |
| Data type | Bitable, Sheet, Mindnote, etc. |  |
| Visual type | Divider |  |
| Media type | Image, File, Iframe, etc. |  |
| Collaboration type | ChatCard |  |
| Container type | TableCell, GridColumn, Callout, View, QuoteContainer, etc. |  |
| Vertical type | Diagram (Flowchart & UML Diagram) |  |
| Auxiliary type | Table, Grid, etc. |  |
| Third-party type | ISV (Independent Software Vendor) block |  |
| Undefined type | Undefined block | / | **Parent-and-child relationships of block**

Parent-child relationships can exist among blocks. Before you use the Create blocks API, refer to the following rules to specify the parent block and child block.

**Rules for parent blocks**

Only blocks with the capacity to contain can serve as parent blocks to which child blocks can be added. These blocks include:

- Text-type functional blocks: Page, Text, Heading, Bullet, Ordered, Task, and Todo blocks.
- Container-type functional blocks: TableCell, GridColumn, Callout, and QuoteContainer blocks.

**Rules for child blocks**

The following blocks cannot be added as child blocks to a parent block:

| Block      | Description                                                                                                             |
| ---------- | ----------------------------------------------------------------------------------------------------------------------- |
| Page       | A document has only one page block, which is automatically generated during document creation.                          |
| GridColumn | Can only be added through the InsertGridColumnRequest of Update a block interface.                                      |
| Mindnote   | Not supported.                                                                                                          |
| TableCell  | Can only be added through the InsertTableRowRequest or InsertTableColumnRequest operations of Update a block interface. |
| View       | When adding a File block, a default view block is automatically generated.                                              |
| Undefined  | Invalid operation.                                                                                                      |
| Diagram    | Not supported.                                                                                                          |
**Limitations for parent-child relations**

Other limitations for Table Cell, Grid Column, and Callout blocks are described below:

| Block | Restriction |
| --- | --- |
| Table Cell | It is not allowed to add the following blocks as child blocks to a Table Cell block:• Table• Sheet• Bitable• OKR |
| Grid Column | It is not allowed to add the following blocks as child blocks to a Grid Column block:• Grid• Bitable• OKR |
| Callout | Only the following blocks are allowed as child blocks to a Callout block:• Text• HeadingN• OrderedBullet• Task• Todo• Quote• QuoteContainer | ## Limitations

You can use document API to operate on different types of blocks, including creating, reading, and editing block content. The support for different blocks varies in the document APIs. For details, refer to the table below.

In the table, "/" represents that the corresponding operation is not required or has been covered in other capabilities, as follows:

- Capability not required. For example, the Divider block does not contain content, so there is no need to provide the ability to read or edit its content.
- Capability alternatively supported. For instance, for the TableCell block, you will directly create or delete the TableCell block while editing the content of the Table block. Therefore, the platform does not need to provide another creation capability for the TableCell block.
- Capability covered in other open capabilities. For example, for the Bitable block, you can create an empty Bitable block and then use the Base capabilities based on the returned token to use the corresponding read and write APIs. To understand Base capabilities, refer to Base method list.

  | Block | Create block | Read content | Edit content | |---|---|---|---| | Callout | Supported | Supported | Supported | | Table | Supported | Supported | Supported | | Text | Supported | Supported | Supported | | Divider | Supported | / | / | | Grid | Supported | Supported | Supported | | Iframe | Supported | Supported | Not supported | | ChatCard | Supported | / | / | | Image | Supported | / | / | | File | Supported | / | / | | TableCell | / | Supported | Supported | | GridColumn | / | Supported | Supported | | View | / | Supported | Not supported | | ISV | Supported | / | / | | Bitable | Supported | / | / | | Sheet | Supported | / | / | | Mindnote | Not supported | / | / | | Diagram | Not supported | Not supported | Not supported | | QuoteContainer | Supported | Supported | Supported | | Task | Not supported | / | / | | OKR | Supported | Supported | / | | OkrObjective | / | Supported | / | | OkrKeyResult | / | Supported | / | | OkrProgress | / | Supported | / | | AddOns | Supported | Supported | / | | JiraIssue | / | Supported | / | | WikiCatalog | Supported | / | / | | Board | Supported | Supported | / | | Undefined | / | / | / | ## Integration process

The process for integrating with the Document API is illustrated in the diagram below. For details, refer to the "How to integrate" section in Docs Overview.

![whiteboard_exported_image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/d8882f258a8b7834cf5bdf25be05051b_G4W02Ism2N.png?height=546&lazyload=true&width=8202)

## API list

The following is a list of APIs for Document and Block.

### Document
| `GET` Obatain the basic information of a document/open-apis/docx/v1/documents/:document_id | `docx:document` `docx:document:readonly` | `tenant_access_token` `user_access_token` | **✓** | **✓** |
| --- | --- | --- | --- | --- |
| `GET` Obtain the plain text content of the document/open-apis/docx/v1/documents/:document_id/raw_content | `docx:document` `docx:document:readonly` | `tenant_access_token` `user_access_token` | **✓** | **✓** |
| `GET` Obtain all blocks of a document/open-apis/docx/v1/documents/:document_id/blocks | `docx:document` `docx:document:readonly` | `tenant_access_token` `user_access_token` | **✓** | **✓** |
| `POST` Create a document/open-apis/docx/v1/documents | `docx:document` | `tenant_access_token` `user_access_token` | **✓** | **✓** | ## Block
Blocks, as structural elements of document content, are units of information with clear meanings. Blocks come in many forms, and can be a piece of text, a spreadsheet, or an image, etc. Each block is identified by a unique `block_id`.

### Method list
| `GET` 	Obtain the block content/open-apis/docx/v1/documents/:document_id/blocks/:block_id | `docx:document` `docx:document:readonly` | `tenant_access_token` `user_access_token` | **✓** | **✓** |
| --- | --- | --- | --- | --- |
| `POST` 	Create blocks/open-apis/docx/v1/documents/:document_id/blocks/:block_id/children | `docx:document` | `tenant_access_token` `user_access_token` | **✓** | **✓** |
| `POST` 	Create nested blocks/open-apis/docx/v1/documents/:document_id/blocks/:block_id/descendant | `docx:document` | `tenant_access_token` `user_access_token` | **✓** | **✓** |
| `PATCH` 	Update a block/open-apis/docx/v1/documents/:document_id/blocks/:block_id | `docx:document` | `tenant_access_token` `user_access_token` | **✓** | **✓** |
| `PATCH` Batch update blocks/open-apis/docx/v1/documents/:document_id/blocks/batch_update | `docx:document` | `tenant_access_token` `user_access_token` | **✓** | **✓** |
| `DELETE` Delete blocks/open-apis/docx/v1/documents/:document_id/blocks/:block_id/children/batch_delete | `docx:document` | `tenant_access_token` `user_access_token` | **✓** | **✓** |
| `GET` Obtain all the child blocks/open-apis/docx/v1/documents/:document_id/blocks/:block_id/children | `docx:document` `docx:document:readonly` | `tenant_access_token` `user_access_token` | **✓** | **✓** |
