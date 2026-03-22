---
title: "List document version"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/file-version/list"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/drive/v1/files/:file_token/versions"
service: "drive-v1"
resource: "file-version"
section: "Docs"
rate_limit: "1000 per minute、50 per second"
scopes:
  - "drive:drive:version"
  - "drive:drive:version:readonly"
field_scopes:
  - "contact:user.employee_id:readonly"
updated: "1711433292000"
---

# List document version

Get all versions of the document.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/drive/v1/files/:file_token/versions |
| HTTP Method | GET |
| Rate Limit | 1000 per minute、50 per second |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `drive:drive:version` `drive:drive:version:readonly` |
| Required field scopes | > The response body of the API contains the following sensitive fields, and they will be returned only after corresponding scopes are added. If you do not need the fields, it is not recommended that you request the scopes. `contact:user.employee_id:readonly` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `file_token` | `string` | parent document token, how to get the document Token can refer to: How to get the token of docs resources **Example value**: "doxbcyvqZlSc9WlHvQMlSJwUrsb" | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `page_size` | `int` | Yes | **Example value**: 20 **Data validation rules**: - Maximum value: `100` |
| `page_token` | `string` | No | Page identifier. It is not filled in the first request, indicating traversal from the beginning; when there will be more groups, the new page_token will be returned at the same time, and the next traversal can use the page_token to get more groups **Example value**: 1665739388 |
| `obj_type` | `string` | Yes | document type **Example value**: docx **Optional values are**:  - `docx`: Upgraded Docs - `sheet`: Spreadsheet  |
| `user_id_type` | `string` | No | User ID categories **Example value**: open_id **Optional values are**:  - `open_id`: Identifies a user to an app. The same user has different Open IDs in different apps. How to get Open ID - `union_id`: Identifies a user to a tenant that acts as a developer. A user has the same Union ID in apps developed by the same developer, and has different Union IDs in apps developed by different developers. A developer can use Union ID to link the same user's identities in multiple apps.How to get Union ID - `user_id`: Identifies a user to a tenant. The same user has different User IDs in different tenants. In one single tenant, a user has the same User ID in all apps （including store apps）. User ID is usually used to communicate user data between different apps. How to get User ID  **Default value**: `open_id` **When the value is `user_id`, the following field scopes are required**: `contact:user.employee_id:readonly` | ## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `items` | `version[]` | document version list |
|     `name` | `string` | title |
|     `version` | `string` | version number |
|     `parent_token` | `string` | parent document token,how to get the document Token can refer to: How to get the token of docs resources |
|     `owner_id` | `string` | document version owner uid, user_id |
|     `creator_id` | `string` | document version creator uid, user_id |
|     `create_time` | `string` | document version create time |
|     `update_time` | `string` | document version update time |
|     `status` | `string` | document version stats **Optional values are**:  - `0`: exist status - `1`: deleted status - `2`: trash status  |
|     `obj_type` | `string` | document version type **Optional values are**:  - `docx`: Upgraded Docs - `sheet`: Spreadsheet  |
|     `parent_type` | `string` | parent document token **Optional values are**:  - `docx`: Upgraded Docs - `sheet`: Spreadsheet  |
|   `page_token` | `string` | Page identifier, when has_more is true, a new page_token will also be returned. Otherwise, page_token will not be returned |
|   `has_more` | `boolean` | Whether the response body has more parameters | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
        "items": [
            {
                "name": "new_name",
                "version": "version1",
                "parent_token": "doxbcyvqZlSc9WlHvQMlSJwUrsb",
                "owner_id": "694699009591869450",
                "creator_id": "694699009591869451",
                "create_time": "1660708537",
                "update_time": "1660708537",
                "status": "0",
                "obj_type": "docx",
                "parent_type": "docx"
            }
        ],
        "page_token": "doxbcpM2mm3znrLfWnf4browTYp23",
        "has_more": false
    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 403 | 1068400 | Has no permission, please apply the file permission of reading or edition. | Please apply the file permission of reading or edition. |
| 400 | 1068401 | Review no pass, the title has illegal content. | Please rename the title of document version. |
| 404 | 1068404 | Parent file does not exist, please check the file status. | Please check the file status. |
| 400 | 1068410 | Params error, param [file_token] is wrong. | Please check the `file_token` param. |
| 400 | 1068411 | Params error, param [obj_type] is wrong. | Please check the `obj_type` param. |
| 400 | 1068425 | Request failed, please contact the engineer-https://applink.larksuite.com/TLJsX982. | Please contact the engineer-https://applink.larksuite.com/TLJsX982. |
| 400 | 1068412 | Params error, param [version_id] is wrong. | Please check the `version_id` param. |
