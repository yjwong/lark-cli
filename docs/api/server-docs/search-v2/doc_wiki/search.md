---
title: "search docs"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/search-v2/doc_wiki/search"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/search/v2/doc_wiki/search"
service: "search-v2"
resource: "doc_wiki"
section: "Search"
rate_limit: "100 per minute"
scopes:
  - "search:docs:read"
updated: "1773323429000"
---

# Document Search

This interface is used to search cloud documents currently visible to the user based on search keywords (query).

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/search/v2/doc_wiki/search |
| HTTP Method | POST |
| Rate Limit | 100 per minute |
| Supported app types | custom,isv |
| Required scopes | `search:docs:read` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `query` | `string` | Yes | search keywords(The query must be accompanied by at least one filter: doc/wiki) **Example value**: "Lark文档使用指南" **Data validation rules**: - Length range: `0` ～ `50` characters |
| `doc_filter` | `doc_filter` | No | Document filter parameters(At least one doc_filter or wiki_filter parameter must be passed.) **Example value**: {"folder_tokens": ["fld_123456"]} |
|   `creator_ids` | `string[]` | No | Document owner OpenID **Example value**: ["ou_789012"] **Data validation rules**: - Length range: `0` ～ `20` |
|   `doc_types` | `string[]` | No | document type **Example value**: ["SHORTCUT"] **Optional values are**:  - `DOC`: Document - `SHEET`: form - `BITABLE`: Bitable - `MINDNOTE`: mind map - `FILE`: file - `WIKI`: Wiki - `DOCX`: New version of the document - `FOLDER`: Space Folder - `CATALOG`: Wiki2.0 folder - `SLIDES`: New version slides - `SHORTCUT`: shortcut  **Data validation rules**: - Length range: `0` ～ `10` |
|   `folder_tokens` | `string[]` | No | Search for documents in a folder (list of folder tokens) Note: The wiki filter will be disabled if this field exists. **Example value**: ["fld_123456"] **Data validation rules**: - Length range: `0` ～ `50` |
|   `only_title` | `boolean` | No | Search only document title **Example value**: false **Default value**: `false` |
|   `open_time` | `time_range` | No | The time range for browsing the document (second timestamp with start and end fields) |
|     `start` | `int` | No | The starting timestamp of the time range **Example value**: 1742348544 **Data validation rules**: - Value range: `0` ～ `9223372036854775807` |
|     `end` | `int` | No | Time range cutoff timestamp **Example value**: 1742348544 **Data validation rules**: - Value range: `0` ～ `9223372036854775807` |
|   `sort_type` | `string` | No | sort by **Example value**: "CREATE_TIME" **Optional values are**:  - `DEFAULT_TYPE`: Default sort - `OPEN_TIME`: User open time sort - `EDIT_TIME`: User edit time descending - `EDIT_TIME_ASC`: User editing time ascending - `ENTITY_CREATE_TIME_ASC`: entity creation time ascending order(deprecated) - `ENTITY_CREATE_TIME_DESC`: entity creation time descending order(deprecated) - `CREATE_TIME`: Sort by document creation time - `CREATE_TIME_ASC`: chronological order created by document(This sorting method is not currently supported)  |
| `wiki_filter` | `wiki_filter` | No | Wiki filter parameters(At least one doc_filter or wiki_filter parameter must be passed.) **Example value**: {"creator_ids": ["ou_789012"], "space_ids": ["space_123456"]} |
|   `creator_ids` | `string[]` | No | Wiki owner OpenID **Example value**: ["ou_7890123456abcdef"] **Data validation rules**: - Length range: `0` ～ `20` |
|   `doc_types` | `string[]` | No | Wiki type **Example value**: ["SHORTCUT"] **Optional values are**:  - `DOC`: Document - `SHEET`: form - `BITABLE`: Bitable - `MINDNOTE`: mind map - `FILE`: file - `WIKI`: Wiki - `DOCX`: New version of the document - `FOLDER`: Space Folder - `CATALOG`: Wiki2.0 folder - `SLIDES`: New version slides - `SHORTCUT`: shortcut  **Data validation rules**: - Length range: `0` ～ `10` |
|   `space_ids` | `string[]` | No | Search for a Wiki under a Space (List of Space IDs) **Example value**: ["space_1234567890fedcba"] **Data validation rules**: - Length range: `0` ～ `50` |
|   `only_title` | `boolean` | No | Search only Wiki titles **Example value**: false **Default value**: `false` |
|   `open_time` | `time_range` | No | The time range for browsing the document (second timestamp with start and end fields) |
|     `start` | `int` | No | The starting timestamp of the time range **Example value**: 1742348544 **Data validation rules**: - Value range: `0` ～ `9223372036854775807` |
|     `end` | `int` | No | Time range cutoff timestamp **Example value**: 1742348544 **Data validation rules**: - Value range: `0` ～ `9223372036854775807` |
|   `sort_type` | `string` | No | sort by **Example value**: "CREATE_TIME" **Optional values are**:  - `DEFAULT_TYPE`: Default sort - `OPEN_TIME`: User open time sort - `EDIT_TIME`: User edit time descending - `EDIT_TIME_ASC`: User editing time ascending - `ENTITY_CREATE_TIME_ASC`: entity creation time ascending order(deprecated) - `ENTITY_CREATE_TIME_DESC`: entity creation time descending order(deprecated) - `CREATE_TIME`: Sort by document creation time - `CREATE_TIME_ASC`: chronological order created by document(This sorting method is not currently supported)  |
| `page_token` | `string` | No | Page identifier. It is not filled in the first request, indicating traversal from the beginning; when there will be more groups, the new page_token will be returned at the same time, and the next traversal can use the page_token to get more groups **Example value**: "token_1234567890fedcba" |
| `page_size` | `int` | No | paging size **Example value**: 15 **Default value**: `0` **Data validation rules**: - Value range: `0` ～ `20` | ### Request body example

{
    "query": "Lark文档使用指南",
    "doc_filter": {
        "creator_ids": [
            "ou_789012"
        ],
        "doc_types": [
            "SHORTCUT"
        ],
        "folder_tokens": [
            "fld_123456"
        ],
        "only_title": false,
        "open_time": {
            "start": 1742348544,
            "end": 1742348544
        },
        "sort_type": "CREATE_TIME"
    },
    "wiki_filter": {
        "creator_ids": [
            "ou_7890123456abcdef"
        ],
        "doc_types": [
            "SHORTCUT"
        ],
        "space_ids": [
            "space_1234567890fedcba"
        ],
        "only_title": false,
        "open_time": {
            "start": 1742348544,
            "end": 1742348544
        },
        "sort_type": "CREATE_TIME"
    },
    "page_token": "token_1234567890fedcba",
    "page_size": 15
}

## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `total` | `int` | Total number of matching results (secondary paging reference) |
|   `has_more` | `boolean` | Whether the response body has more parameters |
|   `res_units` | `doc_res_unit[]` | Search result list |
|     `title_highlighted` | `string` | title highlighting |
|     `summary_highlighted` | `string` | summary highlighting |
|     `entity_type` | `string` | result type **Optional values are**:  - `DOC`: Doc entity - `WIKI`: Wiki type  |
|     `result_meta` | `doc_meta` | Document search meta-information |
|       `doc_types` | `string` | document type **Optional values are**:  - `DOC`: Document - `SHEET`: form - `BITABLE`: Bitable - `MINDNOTE`: mind map - `FILE`: file - `WIKI`: Wiki - `DOCX`: New version of the document - `FOLDER`: Space Folder - `CATALOG`: Wiki2.0 folder - `SLIDES`: New version slides - `SHORTCUT`: shortcut  |
|       `update_time` | `int` | Update timestamp (seconds) |
|       `url` | `string` | Document link |
|       `owner_name` | `string` | owner name |
|       `owner_id` | `string` | Owner OpenID |
|       `is_cross_tenant` | `boolean` | Whether cross-tenant |
|       `create_time` | `int` | Document creation timestamp (seconds) |
|       `last_open_time` | `int` | Last time timestamp was opened (seconds) |
|       `edit_user_id` | `string` | Last edited user OpenID |
|       `edit_user_name` | `string` | Last Edit User Name |
|       `token` | `string` | Document token |
|   `page_token` | `string` | Page identifier, when has_more is true, a new page_token will also be returned. Otherwise, page_token will not be returned | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
        "total": 100,
        "has_more": true,
        "res_units": [
            {
                "title_highlighted": "Lark文档使用指南",
                "summary_highlighted": "本文介绍Lark文档的创建、编辑与分享功能",
                "entity_type": "DOC",
                "result_meta": {
                    "doc_types": "SHORTCUT",
                    "update_time": 1766567446,
                    "url": "https://www.feishu.cn/docs/dox-1234567890abcdef",
                    "owner_name": "张三",
                    "owner_id": "ou-7890123456abcdef",
                    "is_cross_tenant": false,
                    "create_time": 1766567446,
                    "last_open_time": 1766567446,
                    "edit_user_id": "ou-1122334455aabbcc",
                    "edit_user_name": "李四",
                    "token": "dox_9876543210fedcba"
                }
            }
        ],
        "page_token": "token_1234567890fedcba"
    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 1274001 | invalid param: missing required fields | Check whether the request header and request body contain the necessary authentication information and whether the fields are complete |
| 400 | 1274002 | invalid param: illegal enum value | Verify that the value of an enumeration field meets the defined legal enumeration range |
| 401 | 1274011 | user_access_token is invalid or expired | Check if user_access_token is invalid or expired |
| 429 | 1277001 | rate limit exceeded | Check if the current request frequency exceeds the interface limit threshold |
