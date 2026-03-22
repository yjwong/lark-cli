---
title: "Overview"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/docs-doc-overview"
section: "Deprecated Version (Not Recommended)"
scopes:
  - "drive:drive"
  - "docs:doc"
field_scopes:
  - "drive:drive"
  - "docs:doc:readonly"
updated: "1695369889000"
---

# Docs Overview
Lark Docs allows multiple users to edit online documents in real time. This enables real-time collaborative editing and ensures you always see the latest version. In addition, you can @mention users, add comments, and insert task lists and other content in docs. More than just a document editor, Lark Docs is a rich creative tool.

1. Developers can call the Docs API as a user to create Docs, edit Docs, and perform other actions.  
2. Apps can integrate APIs to automatically generate reports, batch create Docs, and perform other automated actions.

## Terms

### Docs

Docs is a container of rich text content. Each Docs is uniquely identified by a docToken.

### docToken

A docToken is the unique ID of a document. You can use the following methods to get the docToken of a document:  

- Obtains the docToken from the Docs URL: https://xxx.larksuite.com/docs/doccnmPiiJWfj1uk8bV7N5SzXaa
- Obtains the docToken from the return value of the Folder API.

For the content and batch_update methods, you must input the docToken to specify the Docs to operate on. When you successfully create a document using the create method, it returns the corresponding docToken.

### Docs data structure

We have defined a document data structure for the Open API to help developers parse and edit Docs content. Before using the Docs API, we recommend you read the Document Data Structure Overview.  
When you use the content and batch_update APIs to obtain and write Docs content, you must follow this data structure. In addition, if you try to use the batch_update API to edit Docs content but the API reports an error, it won't make any modification to the Docs.

### Docs update workflow

When you create a document using the create API, as there are not yet any other collaborators, you don't have to consider the impact of collaborative editing. You can simply input content to initialize the Docs.  
When you attempt to edit an existing Docs, you must consider the impact of collaborative editing to the Docs by other editors. Because we use "index" to mark the location and area to edit, online edits by other editors may cause the index location to change.  
To solve this problem, we require the content API to obtain the current Docs status and information when submitting content. The content API returns the revision field (Docs version number). When subsequent changes are submitted, you need to pass in the revision field as a parameter. The Docs server automatically computes collaborative changes and inserts the updated content in the correct location.  

Apps can use the following methods to create, read, and edit Docs:  
Create Docs: **documents.create**   
Read Docs: **documents.content**  
Edit Docs: **documents.batch_update** 

The workflows for editing a document after creation and editing an existing Docs are as follows:

![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/c9b9c6ec307ce0cbaef4b9c22431c41e_zFz24Eb0Kh.png?height=324&lazyload=true&width=833)

> **Note:** Note that if the revision field provided when submitting a document update differs too much from the latest revision version of the Docs, the submission may fail. In this case, we recommend you use the content API to obtain the latest version before editing the Docs.

## Resources: Document
Each document is assigned a unique  token, which is its ID.
> **Note:** Because the specifications of online resources have changed over time, different APIs have different names for this  token, including doc_token, document_token, docToken. Carefully read the API documentation to avoid errors due to naming issues.

##  Field description
| Parameter | Type | Description |
| --- | --- | --- |
| `document_token` | `string` | The unique identifier of each document.   **Example value**: "doccnULnB44EMMPSYa3rIb4eJCf" **Required field scopes (Select any)**: `drive:drive` `docs:doc:readonly` |
| `title` | `string` | The document heading. |
| `revision` | `int` | The document version, used to confirm version updates during collaboration. |
| `content` | `string` | The document text. **Structure description**: Refer to Document Data Structure Overview  and  Document Data Structure Reference | ### Method list
>  "Store" represents  store apps; "Custom" represents  Custom apps.
| `Create document `POST` /open-apis/doc/v2/create ` | `drive:drive` `docs:doc` | `tenant_access_token` `user_access_token` | **✓** | **✓** |
| --- | --- | --- | --- | --- |
| `Obtain rich text content of a document` `GET` /open-apis/doc/v2/:docToken/content | `drive:drive` `docs:doc` | `tenant_access_token` `user_access_token` | **✓** | **✓** |
| `Obtain text content of a document` `GET` /open-apis/doc/v2/:docToken/raw_content > Obtains the plain text content of a document | `drive:drive` `docs:doc` | `tenant_access_token` `user_access_token` | **✓** | **✓** |
| `Edit document` `POST` /open-apis/doc/v2/:docToken/batch_update | `drive:drive` `docs:doc` | `tenant_access_token` `user_access_token` | **✓** | **✓** |
