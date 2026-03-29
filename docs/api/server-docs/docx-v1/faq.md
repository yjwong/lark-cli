---
title: "Document FAQs"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/document-docx/docx-v1/faq"
service: "docx-v1"
resource: "faq"
section: "Docs"
updated: "1733381764000"
---

# Document FAQs

## 1. How to create a table containing specific content?

Use the Create Nested Block API to create a Table Block under a specified Parent Block. 

```bash
curl --location --request POST 'https://open.larksuite.com/open-apis/docx/v1/documents/:document_id/blocks/:block_id/descendant?document_revision_id=-1' \
--header 'Authorization: Bearer u-xxxx' \
--header 'Content-Type: application/json; charset=utf-8' \
--data-raw '{
    "index": 0,
    "children_id": [
        "headingid_1",
        "table_id_1"
    ],
    "descendants": [
        {
            "block_id": "headingid_1",
            "block_type": 3,
            "heading1": {
                "elements": [{"text_run": {"content" : "Table Block"}}]
            },
            "children": []
        },
        {
            "block_id":"table_id_1",
            "block_type": 31,
            "table":{
                "property" : {
                "row_size": 1,
                "column_size" : 2
                }
            },
            "children": ["table_cell1","table_cell2"]
        },
        {
            "block_id": "table_cell1",
            "block_type": 32,
            "table_cell":{},
            "children": ["table_cell1_child1", "table_cell1_child2"]
        },
        {
            "block_id": "table_cell2",
            "block_type": 32,
            "table_cell":{},
            "children": ["table_cell2_child"]
        },
        {
            "block_id": "table_cell1_child1",
            "block_type": 13,
            "ordered": {
                "elements": [{"text_run": {"content" : "list 1.1"}}]
            },
            "children": []
        },
        {
            "block_id": "table_cell1_child2",
            "block_type": 13,
            "ordered": {
                "elements": [{"text_run": {"content" : "list 1.2"}}]
            },
            "children": []
        },
        {
            "block_id": "table_cell2_child",
            "block_type": 2,
            "text": {
                "elements": [{"text_run": {"content" : ""}}]
            },
            "children": []
        }
    ]
}'
# Before making an API call, replace 'u-xxx' in 'Authorization: Bearer u-xxx' with your actual access token.
```
In the example provided, a Level 1 Heading Block titled "Table Block" was successfully created, along with a Table Block containing specific content.

As shown in the figure below:

![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/12164160bbc6dfee2bf2ebf37eb569b4_d1JineKIz4.png?height=338&lazyload=true&width=1844)

## 2. How to create a sheet block and fill content into table cells?

1. Call the Create Block API to create a Sheet Block under the specified Parent Block, specifying the number of rows and columns for the spreadsheet.
    ```json
    curl --location --request PATCH 'https://{domain}/open-apis/docx/v1/documents/:document_id/blocks/:block_id/children' \
    --header 'Authorization: {Authorization}' \
    --header 'Content-Type: application/json' \
    --data-raw '{
      {
        "index": 0,
        "children": [
          {
            "block_type": 30,
            "sheet": {
              "row_size": 5,
              "column_size": 3
            }
          }
        ]
      }
    }'
    ```	
    In the example above, we create a table with 5 rows and 3 columns. If the call is successful, it is expected to return data in the following format.
    ```json
    {
      "code": 0,
      "data": {
        "children": [
          {
            "block_id": "doxcnx8mv0hzeY07TUlKzpabcef",
            "block_type": 30,
            "parent_id": "UFZvdKi97ojvkzx3ZZocklabcef",
            "sheet": {
              "token": "LxvrsycFwhQYfrt8oYQcwVabcef_QJ6HZR" // Spreadsheet token + worksheet ID format
            }
          }
        ],
        "client_token": "f098d96e-693b-442f-8a7d-82c309ebc500",
        "document_revision_id": 54
      },
      "msg": "success"
    }
    ```

2. The returned `sheet.token` value is a combination of the spreadsheet token and the spreadsheet worksheet ID. You can continue to call the Spreadsheet-related APIs to further operate on the spreadsheet. The following example shows writing data into the spreadsheet.
    ```bash
    curl --location --request PUT 'https://open.larksuite.com/open-apis/sheets/v2/spreadsheets/LxvrsycFwhQYfrt8oYQcwVabcef/values' \
    --header 'Authorization: Bearer t-g10474apW3IFUPQGV362IPSAGELJO2SQWL5abcef' \
    --header 'Content-Type: application/json' \
    --data-raw '{
    "valueRange":{
        "range": "QJ6HZR!A1:B2",
        "values": [
          [
            "Hello", 1
          ],
          [
            "World", 1
          ]
        ]
        }
    }'
    ```
    If the call is successful, it is expected to return the following data:
    ```json
    {
      "code": 0,
      "data": {
        "revision": 2,
        "spreadsheetToken": "LxvrsycFwhQYfrt8oYQcwVabcef",
        "updatedCells": 4,
        "updatedColumns": 2,
        "updatedRange": "QJ6HZR!A1:B2",
        "updatedRows": 2
      },
      "msg": "success"
    }
    ```
## 3. How to insert an image?

**Step 1: Create an image block**

Call the Create blocks API to create an image block under the specified parent block:
```bash
curl --location --request PATCH '{url}' \
--header 'Authorization: {Authorization}' \
--header 'Content-Type: application/json' \
--data-raw '{
    "index": 0,
    "children": [
        {
            "block_type": 27,
            "image": {
                "token": ""
            }
        }
    ]
}'
```
If the call is successful, the expected return is as follows:
```bash
{
    "code": 0,
    "data": {
        "children": [
            {
                "block_id": "doxcnEUmKKppwWrnUIcgZ2ibc9g",
                 // Image block ID
                "block_type": 27,
                "image": {
                    "height": 100,
                    "token": "",
                    "width": 100
                },
                "parent_id": "doxcnQxzmNsMl9rsJRZrCpGx71e"
            }
        ],
        "client_token": "bc25a4f0-9a24-4ade-9ca2-6c1db43fa61d",
        "document_revision_id": 7
    },
    "msg": ""
}
```
**Step 2: Upload image material**

Call the Upload Media API and use the Image block ID returned in step one as the `parent_node` value to upload the image material:

```bash
curl --location --request PATCH '{url}' \
--header 'Authorization: {Authorization}' \
--header 'Content-Type: multipart/form-data; boundary=---7MA4YWxkTrZu0gW' \
--form 'file=@"/tmp/test.PNG"' \ The local path of the image.
--form 'file_name="test.PNG"' \ The name of the image.
--form 'parent_type="docx_image"' \ The material type is docx_image.
--form 'parent_node="doxcnEUmKKppwWrnUIcgZ2ibc9g"' \ Image block ID.
--form 'size="xxx"' \ The size of the image.
```
If the call is successful, the expected return is as follows:
```bash
{
    "code": 0,
    "data": {
        "file_token": "boxbckbfvfcqEg22hAzN8Dh9gJd"  // The image material ID
    },
    "msg": "Success"
}
```
**Step 3: Set the material for the image block**

Call the Update a block or Batch update blocks API to specify the `replace_image` operation and set the material ID returned in step 2 to the corresponding Image block:
```bash
curl --location --request PATCH '{url}' \
--header 'Authorization: {Authorization}' \
--header 'Content-Type: application/json' \
--data-raw '{
    "replace_image": {
        "token": "boxbckbfvfcqEg22hAzN8Dh9gJd" \ The image material ID
    }
}'
```
## 4. How to insert a file?

**Step 1: Create a file block**

Call the Create blocks API to create a file block under the specified parent block:
```bash
curl --location --request PATCH '{url}' \
--header 'Authorization: {Authorization}' \
--header 'Content-Type: application/json' \
--data-raw '{
    "index": 0,
    "children": [
        {
            "block_type": 23,
            "file": {
                "token": ""
            }
        }
    ]
}'
```
If the call is successful, the expected return is as follows:

> After you successfully create a File block, the interface will return a View block. This is because each File block has a corresponding View block to control its visual representation; in other words, the View block is the Parent of the File block.

```bash
{
    "code": 0,
    "data": {
        "children": [
            {
                "block_id": "doxcnIfCrxq7MlhDbj8xCXmPXgf",  // View block ID
                "block_type": 33,
                "children": [
                    "doxcn1Bx1WOlcqzLqTD2UUYiA7g"  // File block ID
                ],
                "parent_id": "doxcnQxzmNsMl9rsJRZrCpGx7ze",
                "view": {
                    "view_type": 1
                }
            }
        ],
        "client_token": "07c56d36-db8b-480f-97f2-7b77a9d3e787",
        "document_revision_id": 8
    },
    "msg": ""
}
```
**Step 2: Upload file material**

Call the Upload Media API and use the File block ID returned in step one as the `parent_node` to upload the material:
```bash
curl --location --request PATCH '{url}' \
--header 'Authorization: {Authorization}' \
--header 'Content-Type: multipart/form-data; boundary=---7MA4YWxkTrZu0gW' \
--form 'file=@"/tmp/test.PNG"' \ The local path of the file.
--form 'file_name="test.PNG"' \ The name of the file.
--form 'parent_type="docx_file"' \ The material type is docx_file.
--form 'parent_node="doxcnEUmKKppwWrnUIcgZ2ibc9g"' \ File block ID.
--form 'size="xxx"' \ The size of the file.
```
If the call is successful, the expected return is as follows:
```bash
{
    "code": 0,
    "data": {
        "file_token": "boxbcXvrJyOMX6EhmGF1bkoQwOb"  // The file material ID
    },
    "msg": "Success"
}
```
**Step 3: Set the material for the file block**

Call the Update a block or Batch update blocks API to specify the `replace_file` operation and set the material ID returned in step 2 to the corresponding File block:
```bash
curl --location --request PATCH '{url}' \
--header 'Authorization: {Authorization}' \
--header 'Content-Type: application/json' \
--data-raw '{
    "replace_file": {
        "token": "boxbcXvrJyOMX6EhmGF1bkoQwOb" \ The file material ID
    }
}'
```
## 5. How to insert a "@user" element?

Note
The "@user" element created by APIs will not send a notification to the mentioned user.
"@User" is defined as an element under the Text block. Its structure is as follows:
```bash
{
    "user_id": string,
    "text_element_style": object(TextElementStyle)
}
```
If you want to mention a user, you can call the Create blocks API to create a Text block under the specified Parent block and specify the user ID of the user to be mentioned in the Text block:
```bash

# https://{domain}/open-apis/docx/v1/documents/:document_id/blocks/:block_id/children?document_revision_id=-1

curl --location --request PATCH '{url}' \
--header 'Content-Type: application/json' \
--header 'Authorization: {Authorization}' \
--data-raw '{
    "children": [
        {
            "block_type": 2,
            "text": {
                "elements": [
                    {
                        "mention_user": {
                            "text_element_style": {
                                "bold": false,
                                "inline_code": false,
                                "italic": false,
                                "strikethrough": false,
                                "underline": false
                            },
                            "user_id": "{user_id}"
                        }
                    }
                ],
                "style": {
                    "align": 1,
                    "folded": false
                }
            }
        }
    ],
    "index": 0
}'
```
In the above example, a Text block is created as the child block of a specific block. The Text block contains a text element `mention_user` that mentions a specified user with `user_id`. If the call is successful, the expected return is as follows:
```bash
{
  "code": 0,
  "data": {
    "children": [
      {
        "block_id": "......",
        "block_type": 2,
        "parent_id": "......",
        "text": {
          "elements": [
            {
              "mention_user": {
                "text_element_style": {
                  "bold": false,
                  "inline_code": false,
                  "italic": false,
                  "strikethrough": false,
                  "underline": false
                },
                "user_id": "......"
              }
            }
          ],
          "style": {
            "align": 1,
            "folded": false
          }
        }
      }
    ],
    ......
  },
  "msg": "success"
}
```
## 6.How to insert a formula?

A formula is an element under the Text block, and its structure is as follows:
```bash
{
    "content": string,
    "text_element_style": object(TextElementStyle)
}
```
If you want to insert a formula into the document, you can call the Create blocks API. The request body is as follows:
```json
{
  "index": 0,
  "children": [
    {
      "block_type": 2,
      "text": {
        "elements": [
          {
            "equation": {
              "content": "1+2=3\n",
              "text_element_style": {
                "bold": false,
                "inline_code": false,
                "italic": false,
                "strikethrough": false,
                "underline": false
              }
            }
          }
        ],
        "style": {
          "align": 1,
          "folded": false
        }
      }
    }
  ]
}
```
## 7. How to fill content into a Callout block?

Call the Create blocks API and fill the`block_id` path parameter with the Callout block ID. Then, enter the content in the `children` array in request body of the Callout block. The following example shows how to insert a Text block in the first line of the Callout block:
```bash
curl --location --request POST '{url}' \
--header 'Authorization: {Authorization}' \ 
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
                            "content": "多人实时协同编辑，一切元素都可插入。",
                            "text_element_style": {
                                "background_color": 14,
                                "text_color": 5
                            }
                        }
                    },
                    {
                        "text_run": {
                            "content": "不仅是在线文档，更是强大的创作和互动工具。",
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
## 8. How to insert a Grid block and add content in the first column?

**Step one: Create a** **Grid** **block**

Call the Create blocks API to create a Grid block under the specified Parent block. The Grid has a total of two columns:

```bash

# https://{domain}/open-apis/docx/v1/documents/:document_id/blocks/:block_id/children

curl --location --request PATCH '{url}' \
--header 'Authorization: {Authorization}' \
--header 'Content-Type: application/json' \
--data-raw '{
  "index": 0,
  "children": [
    {
      "block_type": 24,
      "grid": {
        "column_size": 2
      }
    }
  ]
}'
```
If the call is successful, the expected return is as follows:
> After you create the Grid block, the interface will return the `block_id` of the Grid block and its children, which are the Grid Column blocks. Since the request specifies the creation of two columns, the children array will contain two block_ids. You can then use these block_ids to continue adding Children blocks to the Grid Column blocks.

```json
{
  "code": 0,
  "data": {
    "children": [
      {
        "block_id": "doxcn7VulseZpcWivDsfNiabcef",
        "block_type": 24,
        "children": [
          "doxcnVDmCQuoiQPJUXuaYJabcef", // The first Grid Column block
          "doxcnR4tyA3dJn9MWxa1Vrabcef"  // The second Grid Column block
        ],
        "grid": {
          "column_size": 2
        },
        "parent_id": "Xrt5aEe0DoKTslxIqBRcIEabcef"
      }
    ],
    "client_token": "bef26316-0079-4f26-995e-447004dd996a",
    "document_revision_id": 85
  },
  "msg": "success"
}
```
**Step two: add content in the first column of the** **Grid** **Column block**

Call the Create blocks API to create a Text block under the specified Grid Column block:
```bash

# https://{domain}/open-apis/docx/v1/documents/:document_id/blocks/:block_id/children

curl --location --request PATCH '{url}' \
--header 'Authorization: {Authorization}' \
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
              "content": "多人实时协同，插入一切元素。不仅是在线文档，更是强大的创作和互动工具",
              "text_element_style": {
                "background_color": 14,
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
If the call is successful, the expected return is as follows:
> When creating the Grid block in step 1, the system will automatically add an empty Text block under each Grid Column block. If you don't need the default empty Text block, you can delete it after adding content in this step.

```json
{
  "code": 0,
  "data": {
    "children": [
      {
        "block_id": "doxcnT2booYsWL6XsAcyl958nye",
        "block_type": 2,
        "parent_id": "doxcnVDmCQuoi1PJUXuTYJnEeBe",
        "text": {
          "elements": [
            {
              "text_run": {
                "content": "多人实时协同，插入一切元素。不仅是在线文档，更是强大的创作和互动工具",
                "text_element_style": {
                  "background_color": 14,
                  "bold": false,
                  "inline_code": false,
                  "italic": false,
                  "strikethrough": false,
                  "text_color": 5,
                  "underline": false
                }
              }
            }
          ],
          "style": {
            "align": 1,
            "folded": false
          }
        }
      }
    ],
    "client_token": "b09b2539-487b-42f3-b747-f12ab177bb13",
    "document_revision_id": 86
  },
  "msg": "success"
}
```
## 9. What is the frequency limit for server OpenAPIs?

Refer to the corresponding interface documentation for specific details. For example, the frequency limit for the Update a block API is 3 times per second per application.
## 10. Which types of blocks are supported by the Document OpenAPIs?

Refer to the blockType section in the Block data structure for specific details.
## 11. How to add blocks to a newly created document?

A newly created document already contains a block, which is the Page block.
After you successfully create an empty document, the interface will return the `document_id`, which is also the `block_id` of the Page block in the document. Therefore, you can add blocks by specifying the `document_id` when calling the Create blocks API.
## 12. In what order does the Obtain all blocks of a document API return blocks?

The Obtain all blocks of a document API returns an array of blocks (items) in the order of a pre-order traversal of the document content. The array elements are indexed from 1 to N (N ≥ 1), and the element at index 0 is the root node of the document.

![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/673f2b6f1f5b454495fae746534ce96c_vZ99KtpLpl.png?height=436&lazyload=true&width=1123)

Taking the above figure as an example, the result of blocks pre-order traversal is:

![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/b720c193a1ecfb23ecde4d089eaceb69_ykCnkrgt6G.png?height=181&lazyload=true&width=1013)
## 13. How to update the document title?

The blocks within a document have a tree-like structure, with the root block being the Page block. The document title is a text element of the Page block, and the block ID of the Page block serves as the document ID. To update the document title, you need to use the Update a block API to specify either the `update_text_elements` or `update_text` operation. The path parameters for the request should include:
- `document_id`: Specify the document ID of the document.
- `block_id`: Specify the document ID of the document, which is also the block ID of the Page block.
```bash
curl --location --request PATCH '{url}' \
--header 'Content-Type: application/json' \
--header 'Authorization: {Authorization}' \
--data-raw '{
    "update_text_elements": {
        "elements": [
            {
                "text_run": {
                    "content": "New Title" \ Add your new title here
                }
            }
        ]
    }
}'
```
## 14. Is there a server SDK for the document APIs?

> The development of SDKs may lag behind the latest API updates. For example, the OpenAPI may already support returning a new type of block, but the SDK might not yet support parsing it. However, the SDK is backward compatible.

Currently, SDKs are available for Java, Python, Go, and Node.js. Here are the details for each language:

**Java**

- GitHub Repository：[larksuite/oapi-sdk-java](https://github.com/larksuite/oapi-sdk-java/tree/main/larksuite-oapi/src/main/java/com/larksuite/oapi/service/docx/v1)

**Python**

- GitHub Repository：[larksuite/oapi-sdk-python](https://github.com/larksuite/oapi-sdk-python/tree/main/src/larksuiteoapi/service/docx)
- 
**Go**

- GitHub Repository：[larksuite/oapi-sdk-go](https://github.com/larksuite/oapi-sdk-go/tree/v3_main/service/docx)

**Node.js**

- GitHub Repository：[larksuite/oapi-sdk-node.js](https://github.com/larksuite/node-sdk/blob/main/code-gen/projects/docx.ts)

## 15. How to create a document directly from a document template?

A template is also a document. You can use its document ID and use the Copy File API to create its copy. The document ID of a document is shown below.

![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/fdbc92b302594d28c5fa63c029327645_DNi8RsbJ7b.png?lazyload=true&width=2218&height=479)
## 16. What are the limitations of the Document OpenAPIs?

Compared to the old version of document OpenAPIs, the current version of Document OpenAPI does not support creating documents with content. This is because it is complicated to add blocks with multiple parent-child relationships. It is recommended to use APIs such as import and create a copy to create documents with content.

## 17. How to retrieve images and attachments from a document?

1. Call the List All Blocks in a Document API to paginate and retrieve the rich text content of all blocks in the document.
    ```bash
    curl --location 'https://open.larksuite.com/open-apis/docx/v1/documents/:document_id/blocks' \
    --header 'Authorization: {Authorization}'
    ```
    :::note
    This is a paginated API. If `has_more` is `true`, it indicates that the next pagination of the document content can be queried using the `page_token` returned in the response.
    :::
    
    In the above example, we retrieved the rich text content of the blocks on the first page of the document. If the call is successful, the expected response format is shown below:
    ```json
    {
        "code": 0,
        "data": {
            "has_more": true,
            "page_token": "aw7DoMKBFMOGwqHCrcO8w6jCmMOvw6ILeADCvsKNw57Di8O5XGV3LG4_w5HCqhFxSnDCrCzCn0BgZcOYUg85EMOYcEAcwqYOw4ojw5QFwofCu8KoIMO3K8Ktw4IuNMOBBHNYw4bCgCV3U1zDu8K-J8KSR8Kgw7Y0fsKZdsKvW3d9w53DnkHDrcO5bDkYwrvDisOEPcOtVFJ-I03CnsOILMOoAmLDknd6dsKqG1bClAjDuS3CvcOTwo7Dg8OrwovDsRdqIcKxw5HDohTDtXN9w5rCkWo",
            "items": [
                {
                    "block_id": "AZEUdA02Qo3uuWxjVo7cEyNJnLf",
                    "block_type": 1,
                    "children": [
                        "MQFydWYYCoEDdpxiq4kcjHZ0noW",
                        "OTZtdNzhFoXOWlxd4BkcKO4on2d"
                    ],
                    "page": {
                        "elements": [
                            {
                                "text_run": {
                                    "content": "Beautiful Gallery",
                                    "text_element_style": {
                                        "bold": false,
                                        "inline_code": false,
                                        "italic": false,
                                        "strikethrough": false,
                                        "underline": false
                                    }
                                }
                            }
                        ],
                        "style": {
                            "align": 1
                        }
                    },
                    "parent_id": ""
                },
                {
                    "block_id": "MQFydWYYCoEDdpxiq4kcjHZ0noW",
                    "block_type": 27, // block_type: image
                    "image": {
                        "align": 2,
                        "height": 1200,
                        "token": "HbuhbbMDBoNf1AxZt0Cc6nR6nSe", // image token
                        "width": 4800
                    },
                    "parent_id": "AZEUdA02Qo3uuWxjVo7cEyNJnLf"
                },
                {
                    "block_id": "OTZtdNzhFoXOWlxd4BkcKO4on2d",
                    "block_type": 33, // block_type: view
                    "children": [
                        "I90UdpixCo6ZDOxE7dscMWlRn3e"
                    ],
                    "parent_id": "AZEUdA02Qo3uuWxjVo7cEyNJnLf",
                    "view": {
                        "view_type": 1
                    }
                },
                {
                    "block_id": "I90UdpixCo6ZDOxE7dscMWlRn3e",
                    "block_type": 23, // block_type: file
                    "file": {
                        "name": "image.png",
                        "token": "KNm7bdTXooqUNAx52ZWcBR0Enib" // file token
                    },
                    "parent_id": "OTZtdNzhFoXOWlxd4BkcKO4on2d"
                }
            ]
        },
        "msg": "success"
    }
    ```
2. In the returned data:
    - Blocks with `"block_type": 27` are image blocks, where the `image.token` value is the token of the image.
    - Blocks with `"block_type": 23` are file blocks, where the `file.token` value is the token of the file.
   
  	You can use the tokens of the images and files to call the Download Media API to download the corresponding images and files.

## 18. How to write Markdown content into a document?

Currently, there is no API that supports writing Markdown text directly into a Lark online document. Only the conversion of a Markdown file into a Lark online document is supported. For more details, refer to the File Import Overview.

## 19. How to read detailed content of a spreadsheet block in a document?

1. Call the Get All Blocks of the Document API to retrieve the token of the spreadsheet block. If the call is successful, the returned data will be in the following format.

   The returned `sheet.token` value, `B3hasMxsshByaEtZxAwcVfWxnSe_Ml1QzO`, is a combination of the unique identifier for the spreadsheet (spreadsheet_token) and the unique identifier for the worksheet (sheet_id).

    ```json
    {
      "code": 0,
      "data": {
        "has_more": false,
        "items": [
          {
            "block_id": "RMDydlSOLojiUEx0SROcA4fdn5d",
            "block_type": 1,
            "children": [
              "XpvLdPiaxoBM08xhyfEcVZj7nlc"
            ],
            "page": {
              "elements": [
                {
                  "text_run": {
                    "content": "A Document",
                    "text_element_style": {
                      "bold": false,
                      "inline_code": false,
                      "italic": false,
                      "strikethrough": false,
                      "underline": false
                    }
                  }
                }
              ],
              "style": {
                "align": 1
              }
            },
            "parent_id": ""
          },
          {
            "block_id": "XpvLdPiaxoBM08xhyfEcVZj7nlc",
            "block_type": 30,
            "parent_id": "RMDydlSOLojiUEx0SROcA4fdn5d",
            "sheet": {
              "token": "B3hasMxsshByaEtZxAwcVfWxnSe_Ml1QzO"  // Combination of the spreadsheet's unique identifier (spreadsheet_token) and the worksheet's unique identifier (sheet_id)
            }
          }
        ]
      },
      "msg": "success"
    }
    ```

2. Using the spreadsheet's unique identifier (spreadsheet_token) and the worksheet's unique identifier (sheet_id) obtained in step one, call the Read a Single Range from the Spreadsheet API to retrieve the data from the spreadsheet. For request and response examples, please refer directly to the API documentation.
