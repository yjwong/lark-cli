---
title: "Document version changes"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/docs/upgraded-docs-access-guide/upgraded-docs-openapi-access-guide"
section: "Docs"
updated: "1733381809000"
---

# Document version changes

In June 2022, the Lark Open Platform introduced an upgraded version of the document APIs. This article outlines how to distinguish between the old and upgraded versions of documents in links and codes, as well as the differences in the APIs between the old and upgraded versions.

## Distinguish between upgraded and old versions

You can distinguish between upgraded and old documents by document links and document type-related fields in the code.
| **Document category** | **Document link**                 | **File type related fields** |
| --------------------- | --------------------------------- | ---------------------------- |
| Upgraded version      | https://:domain/docx/:docx_token | docx                         |
| Old version           | https://:domain/docs/:docs_token | doc                          | > In different APIs of Docs, the fields related to "file type" are expressed differently, which may be `type`, `types`, `parent_type`, `docs_type`, `file_type`, `obj_type`, and more.

## API changes

For the following capabilities, the APIs of the upgraded version and the old version are different. You cannot mix and use them.

> It is no longer supported to create documents through the old version document API.

| **Capacity** | **Upgraded version APIs** | **Old version APIs**  |
|---|---|---|
| Create documents | Create a document (Creating a document with content is not supported. You can create a document with content by importing, creating a copy, etc.) | Create document (deprecated) |
| Edit document content | Create a blockUpdate a blockBatch update blocksDelete a block | Batch Update Document |
| Get document rich-text content | Obtain the block contentObtain all blocks of a document Obtain all the child blocks| Get Document |
| Get the document's plain text content | Obtain the plain text content of the document | Obtain Document Content|
| Get document metadata | Obtain the basic information of a document  You can also use the Space capability's Obtain Metadata interface to get document metadata | Obtain Document Meta | ## Other Docs API changes

Other open capabilities of Docs, such as cloud space, permissions, comments, etc., support choosing between old or upgraded versions of documents by specifying the document type.

You can pass the document type parameter in the request body to specify the upgraded or old version. Specific fields are as follows:

- `"type": "docx"`, `"file_type": "docx"`, `"parent_type": "docx_file"`, etc., indicate specifying the upgraded version document;
- `"type": "doc"`, `"file_type": "doc"`, `"parent_type": "doc_file"`, etc., indicate specifying the old version document.

If your application uses a request of `doc` type to access the API in the following table, after the document is upgraded to the new version, the document type needs to be replaced with `docx`.

| Docs capabilities | Related interfaces |
|---|---|
| Cloud space | Obtain metadataCopy filesCopy files (Deprecated)Move filesDelete filesGet file statisticsCreate an import taskQuery import resultsCreate an export task Query export task resultsSearch Documents |
| Permissions | Increase collaborator permissionsUpdate collaborator permissionsRemove collaborator permissionsDetermine whether the collaborator has a certain permissionGet the list of collaboratorsTransfer ownershipGet Docs permission settingsUpdate Docs permission settings|
| Subscription | Create a subscriptionGet subscription statusUpdate subscription status |
|Event|Subscribe Docs events|
|Media|Upload mediaUpload media in blocks-Preuploading|
| Wiki | Move cloud document to Wiki Create a node copy in Wiki|
| Comment | Get Document Comments in PagesGet a Global CommentBatch Query CommentsAdd a Global CommentGet Replies ListUpdate ReplyDelete ReplySolve or Restore a Comment | ## FAQs

### 1. Why not make the upgraded document compatible on the basis of the old version API?

The formats and protocols of the old and new versions of docs are completely different. The original data structure and API protocols cannot be compatible. Thus, when operating the upgraded document, the new version API needs to be used.

### 2. Does the upgraded document use Location to mark the position as the old document?

Only the old version use Location to mark the position. For details, see Document Structure Introduction.

There is no Location concept in the upgraded version of the document APIs, which is also one of the main differences between the underlying data structure of the upgraded version and the old version of the document. The upgraded version of the document is based on Block ID and Parent Block ID for positioning, and it can be viewed as a tree structure of blocks.
