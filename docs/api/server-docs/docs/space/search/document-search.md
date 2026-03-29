---
title: "Document Search"
url: "https://open.larksuite.com/document/ukTMukTMukTM/ugDM4UjL4ADO14COwgTN"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/suite/docs-api/search/object"
section: "Docs"
scopes:
  - "drive:drive"
  - "drive:drive:readonly"
updated: "1665221890000"
---

# Search for a document

This API is used to search for a document based on search criteria. 

## Request
| HTTP URL | https://open.larksuite.com/open-apis/suite/docs-api/search/object |
| --- | --- |
| HTTP Method | POST |
| Required scopes | `drive:drive` `drive:drive:readonly` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" Learn more about how to access and use access_token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Request body
|Parameter|Type|Required|Description|
|--|-----|--|----|
|search_key|string|Yes|Keyword for the search |
|count|int|No|Number of entries to return for the search. Value range: 0 = 0, and offset + count |No|User ID of the document owner|
|chat_ids|list|No|Chat ID of the group where the document is located|
|docs_types|list|No|Type of the document. Optional values: "doc", "sheet", "slide", "bitable", "mindnote", and "file" | ### Request body example
```json
{
    "search_key": "search key",
    "count": 10, 
    "offset": 0,
    "owner_ids": ["xxx", "xxx"],
    "chat_ids": ["xxx", "xxx"],
    "docs_types": ["doc", "sheet"]
}
```
## Response
### Response body
  |Parameter|Description|
|--|--|
|docs_entities|List of matched documents|
|&ensp;∟docs_token|Token of the document|
|&ensp;∟docs_type|Type of the document|
|&ensp;∟title|Title|
|&ensp;∟owner_id|Owner of the file|
|has_more|Indicates whether more data exists after the search offset results are returned.|
|total|Total number of matched documents|
### Response body example

```json
{
    "code": 0,
    "msg": "Success",
    "data": {
        "docs_entities": [
            {
                "docs_token": "xxx",
                "docs_type": "doc",
                "title": "xxx",
                "owner_id": "xxx"
            },
            {
                "docs_token": "xxx",
                "docs_type": "sheet",
                "title": "xxx",
                "owner_id": "xxx"
            }
        ],
        "has_more": false,
        "total": 10
    }
}
```
### Error code

For details, refer to Server-side error codes.
