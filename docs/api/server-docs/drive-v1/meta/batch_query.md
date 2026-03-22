---
title: "Obtain Metadata"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/meta/batch_query"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/drive/v1/metas/batch_query"
service: "drive-v1"
resource: "meta"
section: "Docs"
rate_limit: "1000 per minute、50 per second"
scopes:
  - "drive:drive"
  - "drive:drive.metadata:readonly"
field_scopes:
  - "contact:user.employee_id:readonly"
  - "drive:file.meta.sec_label.read_only"
updated: "1686034797000"
---

# Obtain metadata

This API is used to obtain the metadata of various files based on a token.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/drive/v1/metas/batch_query |
| HTTP Method | POST |
| Rate Limit | 1000 per minute、50 per second |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `drive:drive` `drive:drive.metadata:readonly` |
| Required field scopes | > The response body of the API contains the following sensitive fields, and they will be returned only after corresponding scopes are added. If you do not need the fields, it is not recommended that you request the scopes. `contact:user.employee_id:readonly` `drive:file.meta.sec_label.read_only` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `user_id_type` | `string` | No | User ID categories **Example value**: open_id **Optional values are**:  - `open_id`: Identifies a user to an app. The same user has different Open IDs in different apps. How to get Open ID - `union_id`: Identifies a user to a tenant that acts as a developer. A user has the same Union ID in apps developed by the same developer, and has different Union IDs in apps developed by different developers. A developer can use Union ID to link the same user's identities in multiple apps.How to get Union ID - `user_id`: Identifies a user to a tenant. The same user has different User IDs in different tenants. In one single tenant, a user has the same User ID in all apps （including store apps）. User ID is usually used to communicate user data between different apps. How to get User ID  **Default value**: `open_id` **When the value is `user_id`, the following field scopes are required**: `contact:user.employee_id:readonly` | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `request_docs` | `request_doc[]` | Yes | Request documents. Up to 200 documents are supported at a time. **Data validation rules**: - Length range: `1` ～ `200` |
|   `doc_token` | `string` | Yes | Token of the file. For more information about how to obtain the token, see How to get the token of docs resources **Example value**: "doccnfYZzTlvXqZIGTdAHKabcef" |
|   `doc_type` | `string` | Yes | Type of the file **Example value**: "doc" **Optional values are**:  - `doc`: Lark Docs - `sheet`: Lark Sheets - `bitable`: Lark Bitable - `mindnote`: Lark MindNotes - `file`: Lark file - `wiki`: Lark wiki - `docx`: Lark Upgraded Docs - `folder`: Lark folder  |
| `with_url` | `boolean` | No | Whether to get File link **Example value**: False | ### Request body example

{
    "request_docs": [
        {
            "doc_token": "doccnfYZzTlvXqZIGTdAHKabcef",
            "doc_type": "doc"
        }
    ],
    "with_url": false
}

## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `metas` | `meta[]` | Metadatas of the file |
|     `doc_token` | `string` | File token |
|     `doc_type` | `string` | File type |
|     `title` | `string` | Title |
|     `owner_id` | `string` | File owner |
|     `create_time` | `string` | Creation time (Unix timestamp) |
|     `latest_modify_user` | `string` | Last editor |
|     `latest_modify_time` | `string` | Last edited time (Unix timestamp) |
|     `url` | `string` | File link |
|     `sec_label_name` | `string` | Security label **Required field scopes**: `drive:file.meta.sec_label.read_only` |
|   `failed_list` | `meta.failed[]` | Failed list for get metadata |
|     `token` | `string` | File Token |
|     `code` | `int` | Error code for failure to get metadata **Optional values are**:  - `970002`: Unsupported doc-type - `970003`: No permission to access meta - `970005`: Record not found (Does not exist or is deleted)  | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
        "metas": [
            {
                "doc_token": "doccnfYZzTlvXqZIGTdAHKabcef",
                "doc_type": "Doc",
                "title": "SampleTitle",
                "owner_id": "ou_b13d41c02edc52ce66aaae67bf1abcef",
                "create_time": "1652066345",
                "latest_modify_user": "ou_b13d41c02edc52ce66aaae67bf1abcef",
                "latest_modify_time": "1652066345",
                "url": "https://sample.larksuite.cn/docs/doccnfYZzTlvXqZIGTdAHKabcef",
                "sec_label_name": "L2"
            }
        ],
        "failed_list": [
            {
                "token": "boxcnrHpsg1QDqXAAAyachabcef",
                "code": 970005
            }
        ]
    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 401 | 1069701 | User identity verification failed | Check if the appid is correct |
| 503 | 1069704 | Internal server error | Server level error |
