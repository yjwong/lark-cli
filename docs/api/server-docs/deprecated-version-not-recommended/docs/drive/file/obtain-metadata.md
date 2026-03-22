---
title: "Obtain Metadata"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uMjN3UjLzYzN14yM2cTN"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/suite/docs-api/meta"
section: "Deprecated Version (Not Recommended)"
scopes:
  - "drive:drive"
  - "drive:drive.metadata:readonly"
updated: "1703488394000"
---

# Obtain metadata

This API is used to obtain the metadata of various files based on a token. 

> **Note:** Before sending a request, you must be granted the access (read) permission on the file.

## Request
| HTTP URL | https://open.larksuite.com/open-apis/suite/docs-api/meta |
| --- | --- |
| HTTP Method | POST |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `drive:drive` `drive:drive.metadata:readonly` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token`  or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" Learn more about how to obtain and use access_token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Request body
|Parameter|Type|Required|Description|
|--|-----|--|----|
|request_docs||Yes|Request documents. Up to 200 documents are supported at a time.|
|&ensp;∟docs_token|string|Yes|Token of the file. For more information about how to obtain the token, see Overview.|
|&ensp;∟docs_type|string|Yes|Type of the file. Optional values:   1) "doc": Lark Docs2) "sheet": Lark Sheets 3) "bitable": Lark Bitable4) "mindnote": Lark MindNotes 5) "file": Lark Files|
### Request body example
```json
{
    "request_docs": [
        {
            "docs_token": "12345",
            "docs_type": "doc"
        },  
        {
            "docs_token": "12345",
            "docs_type": "sheet"
        }
    ]
}
```

## Response
### Response body
|Parameter|Description|
|--|--|
|docs_metas|Metadata of the file
|&ensp;∟docs_token|Filetoken|
|&ensp;∟docs_type|File type|
|&ensp;∟title|Title|
|&ensp;∟owner_id|File owner|
|&ensp;∟create_time|Creation time. The value is a Unix timestamp. |
|&ensp;∟latest_modify_user|Last editor|
|&ensp;∟latest_modify_time|Last edit time. The value is a Unix timestamp. | ### Response body example
```json
{
    "code": 0, 
    "msg": "Success",
    "data": { 
        "docs_metas": [ { 
                "docs_token": "doc22222",
                "docs_type": "doc",
                "title": "abc", 
                "owner_id": "12345", 
                "create_time": 123456, 
                "latest_modify_user": "12345", 
                "latest_modify_time": 123456
            }
        ]
    }
}
```
### Error code

For details, refer to Server-side error codes.
