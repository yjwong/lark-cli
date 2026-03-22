---
title: "Search Wiki"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uEzN0YjLxcDN24SM3QjN/search_wiki"
section: "Docs"
updated: "1758004486000"
---

# Search Wiki
Users query Wikis by keywords, and can only find wikis that they can see

> **Note:** **Tips：** If the wiki exists but the user cannot search, it does not necessarily mean that there is a problem with the search. It may be that the user does not have the permission to view the wiki

## Request
|Basic| |
|:-----|:-----|
|HTTP URL|https://open.larksuite.com/open-apis/wiki/v2/nodes/search|
|HTTP Method|POST|
|HTTP Content-Type|application/json; charset=utf-8|
|Credentials Requirements|user_access_token|
|Permission requirements|View all Wiki| ### Request body
|Name|Type|Description|
|:-----|:-----|:-----|
|query|string|Search keyword [*required]|
|space_id|string|The space_id to which the document belongs, empty search all [optional]|
|node_id|string|Search for this node and all its child nodes if it is not empty, search all wikis if it is empty. Use node_id filtering must set space_id [optional]|
|page_token|string|Pagination token on the next page, do not need to fill in the home page, get the next page data according to the returned token [optional]|
|page_size|int|The maximum value of the number returned on this page 0 <page_size <= 50 Default 20 [optional]| ### Request body example
``` json
{
    "query": "search keyword",
    "space_id": "xxxxxxx",
    "node_id": "xxxx",
    "page_token": "xxxxx",
    "page_size": 20
}
```

## Response
### Response body
|Name|Type|Description|
|:-----|:-----|:-----|
|code|int|Error code, non-zero indicates failure.|
|msg|string|Error description.|
|data|-|-|
|  ∟items|list|Array of wiki information.|
|    ∟node_id|string|Token of the wiki node.|
|    ∟space_id|string|ID of the knowledge space the wiki belongs to.|
|    ∟obj_type|int|Type of the wiki, refer to the document type table.|
|    ∟obj_token|string|Token of the actual document for the node. To retrieve or edit the node content, this token must be used to call the corresponding API.|
|    ∟parent_id|string|Not yet effective, always returns an empty value.|
|    ∟sort_id|int|Index of the document in the knowledge base, starting from 1.|
|    ∟title|string|Title of the wiki.|
|    ∟url|string|URL for accessing the wiki.|
|    ∟icon|string|URL of the icon corresponding to the wiki.|
|  ∟page_token|string|Returns the token for the next page if has_more is true.|
|  ∟has_more|bool|Indicates whether there are more pages of data.| #### obj_type table

|obj_type|Description|
|--|--|
|1|Doc|
|2|Sheet|
|3|Bitable|
|4|Mindnote|
|5|File|
|6|Slide|
|7|Wiki|
|8|Docx|
|9|Folder|
|10|Catalog| ### Response body example
``` json
{
  "code": 0,
  "data": {
    "has_more": false,
    "items": [
      {
        "node_id": "BAgPwq6lIi5Nykk0E5fcJeabcef",
        "obj_token": "AcnMdexrlokOShxe40Fc0Oabcef",
        "obj_type": 8,
        "parent_id": "",
        "sort_id": 1,
        "space_id": "7307457194084925443",
        "title": "Welcome to Wiki",
        "url": "https://sample.larksuite.cn/wiki/BAgPwq6lIi5Nykk0E5fcJeabcef"
      }
    ]
  },
  "msg": "success"
}
```

## Error code
|HTTP Status Code|Error Code|Description|Troubleshooting Suggestions|
|:-----|:-----|:-----|:-----|
|200|10001|invalid param|Parameter error, refer to the document to check the input parameter|
|200|10002|network anomaly, please try again|The back-end service is abnormal or the network is abnormal, you can request again|
