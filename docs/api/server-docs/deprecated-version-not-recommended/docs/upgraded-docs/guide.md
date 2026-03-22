---
title: "Guide"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/document-docx/guide"
section: "Deprecated Version (Not Recommended)"
updated: "1695369867000"
---

# Access guide

## form
The Upgraded Docs editor has evolved from the document level to the paragraph level. paragraph is usually called **block**, which is defined as **Block** in the data structure of the Upgraded Docs. Block, as an element of content structure, is an information unit with a clear meaning. Multiple blocks can form hierarchical content in a tree-like relationship. The root node of the tree itself is also a special type of block, which is called a page block. For Blocks that form a parent-child relationship, the son Blocks are called Children.

## block list
| **Type** | **Description** |
| --------- | --------------- |
|page | Document Block, which is the root node of the entire document tree |
|text | Text Block|
|heading**N** | Heading Block, heading**N**, **N** value range 1~9|
|bullet | Bullet List Block|
|ordered | Ordered List Block|
|code | Code Block|
|quote | Quote Block|
|todo | TODO Block|
|bitable | Bitable Block |
|callout | Callout Block|
|chat_card | Chat Card Block|
|diagram | UML Diagram Block|
|divider | Divider Block|
|file | File Block|
|grid | Grid Block|
|grid_column|Grid Column Block|
|iframe|Iframe Block|
|image|Image Block|
|isv|ISV Block|
|mindnote|MindNote Block|
|sheet|Sheet Block|
|table|Table Block |
|table_cell|Table Cell Block|
|view|View Block|
|quote_container|Quote Container Block|
|okr|OKR Block|
|okr_objective|OKR Objective Block|
|okr_key_result|OKR Key Result Block|
|okr_progress|OKR Progress Block|
|add_ons|Add-Ons Block|
|wiki_catalog|Wiki Catalog Block|
|jira_issue|Jira Issue Block|
> Block type enumeration definition: Click to view BlockType.

## limit

### QPS Limit
Each API has a call frequency limit. For details, please refer to the frequency control description of the corresponding API documentation.

### Interface error
When the request does not meet some specific conditions, an error will be reported. For details, please refer to the "Error Code" section of the corresponding API document.

### Capability Boundaries
DocX OpenAPI currently operates Block capability boundaries:

| Block | Create | Read Content | Edit Content |
| --------- | --------------- | ------- | ----------- |
|Callout|Support|Support|Support|
|Table|Support|Support|Support|
|Text|Support|Support|Support|
|Divider|Support|-|-|
|Grid|Support|Support|Support|
|Iframe|Supported|Supported|Not Supported|
|ChatCard|Supported|Not Supported|-|
|Image|Supported|Not supported|-|
|File|Supported|Not supported|-|
|TableCell|-|Support|Support|
|GridColumn|-|Support|Support|
|View|-|Supported|Not Supported|
|Diagram|Not supported|Not supported|Not supported|
|ISV|Support|-|-|
|Mindnote|Not supported|Not supported|Not supported|
|Bitable|Support|Not supported|Not supported|
|Sheet|Support|Not supported|Not supported|
|QuoteContainer|Support|Support|Support|
|Task|Not supported|-|-|
|OKR|Support|Support|-|
|OKR Objective|-|Support|-|
|OKR Key Result|-|Support|-|
|OKR Progress|-|Support|-|
|AddOns|Support|Support|-|
|Jira Issue|-|Support|-|
|Wiki Catalog|Support|Support|-|
|Undefined|-|-|-| > 1. - means not supported, for example:
>    - Direct creation of `TableCell` is not supported, because `TableCell` will be created automatically when adding rows or columns to `Table`;
>    - Direct reading of `Bitable` is not supported, and the corresponding `Bitable` OpenAPI can be called according to the returned `Bitable` Token.
> 2. Block type enumeration definition: Click to view BlockType.

## Query Document
Query the **latest version number** and **title** of the document, this API only returns the basic information of the document, not the content of the document.
```bash 
curl --location --request GET 'https://open.larksuite.com/open-apis/docx/v1/documents/doxcnAJ9VRRJqVMYZ1MyKnavXWe' \
--header 'Authorization: Bearer u-xxx'
``` 

## Query Document Raw Content
Gets the plain text content of the document without rich text formatting information.
```bash 
curl --location --request GET 'https://open.larksuite.com/open-apis/docx/v1/documents/doxcnAJ9VRRJqVMYZ1MyKnavXWe/raw_content' \
--header 'Authorization: Bearer u-xxx'
``` 

## Query All Blocks
Perform a deep traversal of all blocks in the specified document and return them in pagination.
```BASH 
curl --location --request GET 'https://open.larksuite.com/open-apis/docx/v1/documents/doxcnAJ9VRRJqVMYZ1MyKnavXWe/blocks?page_size=100&document_revision_id=-1' \
--header 'Authorization: Bearer u-xxx'
```

## Create Document
Supports specifying the space directory through the `folder_token` parameter.
```bash
curl --location --request POST 'https://open.larksuite.com/open-apis/docx/v1/documents?folder_token=fldcnbCHL8OAtkcYHnPzZi1yupN' \
--header 'Authorization: Bearer u-xxx'
``` 

## Query Block
Query the rich text content of the specified block.
> This API will not return children block, if you want to get it, please use "Query Block Children" in combination.

```bash 
curl --location --request GET 'https://open.larksuite.com/open-apis/docx/v1/documents/doxcnAJ9VRRJqVMYZ1MyKnavXWe/blocks/doxcnC4cO4qUui6isgnpofh5edc?document_revision_id=-1' \
--header 'Authorization: Bearer u-xxx'
```  

## Create Block
Creates a batch of children blocks for a given block and inserts them at a specific location. The return value of this API is the rich text content of the newly created children blocks. 
> If no version is specified, the latest version is used by default.
```bash 
curl --location --request POST 'https://open.larksuite.com/open-apis/docx/v1/documents/doxcnAJ9VRRJqVMYZ1MyKnavXWe/blocks/doxcnAJ9VRRJqVMYZ1MyKnavXWe/children?document_revision_id=120' \
--header 'Authorization: Bearer u-xxx' \
--header 'Content-Type: application/json' \
--data-raw '{
    "index": 0,
    "children": [
        {
            "block_type": 2,
            "text": {
                "elements": [
                    {
                        "text_run": {
                            "content": "Multiple people collaborate in real time, inserting all elements. Not just online documentation, but",
                            "text_element_style": {
                                "background_color": 14,
                                "text_color": 5
                            }
                        }
                    },
                    {
                        "text_run": {
                            "content": "Powerful authoring and interactive tools",
                            "text_element_style": {
                                "background_color": 14,
                                "bold": true,
                                "text_color": 5
                            }
                        }
                    }
                ],
                "style": {}
            }
        }
    ]
}'
``` 

## Update Block
Updates the specified block and returns its updated rich text content.
> If no version is specified, the latest version is used by default
```bash 
curl --location --request PATCH 'https://open.larksuite.com/open-apis/docx/v1/documents/doxcnAJ9VRRJqVMYZ1MyKnavXWe/blocks/doxcnC4cO4qUui6isgnpofh5edc' \
--header 'Authorization: Bearer u-xxx' \
--header 'Content-Type: application/json' \
--data-raw '{
    "update_text_elements":{
        "elements":[
            {
                "text_run": {
                        "content": "Space",
                        "text_element_style": {
                            "link": {
                                "url": "https%3A%2F%2Fbytedance.larksuite.com%2Fdrive%2Fhome%2F"
                            }
                        }
                    }
            }
        ]
    }
}'
``` 

## Batch Update Blocks
Batch update the specified block and return its updated rich text content.
> If no version is specified, the latest version is used by default
```bash
curl --location --request PATCH 'https://open.larksuite.com/open-apis/docx/v1/documents/doxcnAJ9VRRJqVMYZ1MyKnavXWe/blocks/batch_update' \
--header 'Authorization: Bearer u-xxx' \ 
--header 'Content-Type: application/json' \
--data-raw '{
    "requests": [
        {
            "block_id": "doxcnk0i44OMOaouw8AdXuXrp6b",
            "merge_table_cells": {
                "column_end_index": 2,
                "column_start_index": 0,
                "row_end_index": 1,
                "row_start_index": 0
            }
        },
        {
            "block_id": "doxcn0K8iGSMW4Mqgs9qlyTP50d",
            "update_text_elements": {
                "elements": [
                    {
                        "text_run": {
                            "content": "Hello",
                            "text_element_style": {
                                "background_color": 2,
                                "bold": true,
                                "italic": true,
                                "strikethrough": true,
                                "text_color": 2,
                                "underline": true                            
                            }
                        }
                    },
                    {
                        "text_run": {
                            "content": "World ",
                            "text_element_style": {
                                "italic": true
                            }
                        }
                    },
                    {
                        "mention_doc": {
                            "obj_type": 22,
                            "token": "doxcnAJ9VRRJqVMYZ1MyKnayXWe",
                            "url": "https://test.larksuite.com/docx/doxcnAJ9VRRJqVMYZ1MyKnayXWe"
                        }
                    }
                ]
            }
        }
    ]
}'
```

## Delete Blocks
Specifies the block that needs to be manipulated, deletes its specified range of children blocks, and returns the version of the document with the action applied. 
> We can regard the Upgraded Docs as a block tree, and the deleted block can be regarded as the child of its parent block, that is, it is clear that the block to be deleted is the first child of its parent block, and then this API is called to delete it.
 ```bash 
curl --location --request DELETE 'https://open.larksuite.com/open-apis/docx/v1/documents/doxcnAJ9VRRJqVMYZ1MyKnavXWe/blocks/doxcnAJ9VRRJqVMYZ1MyKnavXWe/children/batch_delete?document_revision_id=-1' \
--header 'Authorization: Bearer u-xxx' \
--header 'Content-Type: application/json' \
--data-raw '{
    "start_index": 0,
    "end_index": 1
}'
``` 

## Query Block Children
Specify the block to be manipulated, and page through the rich text content of its children blocks.
> This API returns only the rich text content of children blocks.

```bash 
curl --location --request GET 'https://open.larksuite.com/open-apis/docx/v1/documents/doxcnAJ9VRRJqVMYZ1MyKnavXWe/blocks/doxcnAJ9VRRJqVMYZ1MyKnavXWe/children?document_revision_id=-1&page_size=10' \
--header 'Authorization: Bearer u-xxx'
```
