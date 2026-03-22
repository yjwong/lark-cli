---
title: "Update Reply"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/file-comment-reply/update"
method: "PUT"
api_path: "https://open.larksuite.com/open-apis/drive/v1/files/:file_token/comments/:comment_id/replies/:reply_id"
service: "drive-v1"
resource: "file-comment-reply"
section: "Docs"
rate_limit: "1000 per minute、50 per second"
scopes:
  - "docs:doc"
  - "docs:doc:readonly"
  - "drive:drive"
  - "drive:drive:readonly"
  - "sheets:spreadsheet"
  - "sheets:spreadsheet:readonly"
field_scopes:
  - "contact:user.employee_id:readonly"
updated: "1710301985000"
---

# Update Reply

Updates a reply to a document in Docs.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/drive/v1/files/:file_token/comments/:comment_id/replies/:reply_id |
| HTTP Method | PUT |
| Rate Limit | 1000 per minute、50 per second |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `docs:doc` `docs:doc:readonly` `drive:drive` `drive:drive:readonly` `sheets:spreadsheet` `sheets:spreadsheet:readonly` |
| Required field scopes | > The response body of the API contains the following sensitive fields, and they will be returned only after corresponding scopes are added. If you do not need the fields, it is not recommended that you request the scopes. `contact:user.employee_id:readonly` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `file_token` | `string` | Token of the document **Example value**: "doccnHh7U87HOFpii5u5G*****" |
| `comment_id` | `string` | Comment ID **Example value**: "6916106822734578184" |
| `reply_id` | `string` | Reply ID **Example value**: "6916106822734594568" | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `file_type` | `string` | Yes | Type of the document **Example value**: doc **Optional values are**:  - `doc`: Doc - `sheet`: Sheet - `file`: File - `docx`: New Doc  |
| `user_id_type` | `string` | No | User ID categories **Example value**: open_id **Optional values are**:  - `open_id`: Identifies a user to an app. The same user has different Open IDs in different apps. How to get Open ID - `union_id`: Identifies a user to a tenant that acts as a developer. A user has the same Union ID in apps developed by the same developer, and has different Union IDs in apps developed by different developers. A developer can use Union ID to link the same user's identities in multiple apps.How to get Union ID - `user_id`: Identifies a user to a tenant. The same user has different User IDs in different tenants. In one single tenant, a user has the same User ID in all apps （including store apps）. User ID is usually used to communicate user data between different apps. How to get User ID  **Default value**: `open_id` **When the value is `user_id`, the following field scopes are required**: `contact:user.employee_id:readonly` | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `content` | `reply_content` | Yes | Reply content |
|   `elements` | `reply_element[]` | Yes | Reply content |
|     `type` | `string` | Yes | Content elements of the reply **Example value**: "text_run" **Optional values are**:  - `text_run`: Plain text - `docs_link`: Link to @mentioning a document in Docs - `person`: Contact to @mention  |
|     `text_run` | `text_run` | No | Text content |
|       `text` | `string` | Yes | Reply with plain text **Example value**: "comment text" |
|     `docs_link` | `docs_link` | No | Text content |
|       `url` | `string` | Yes | @mentioned document in a reply. **Example value**: "https://example.larksuite.com/docs/doccnHh7U87HOFpii5u5Gabcef" |
|     `person` | `person` | No | Text content |
|       `user_id` | `string` | Yes | @mentioned contact in a reply **Example value**: "ou_cc19b2bfb93f8a44db4b4d6eababcef" | ### Request body example

{
    "content": {
        "elements": [
            {
                "type": "text_run",
                "text_run": {
                    "text": "comment text"
                },
                "docs_link": {
                    "url": "https://example.larksuite.com/docs/doccnHh7U87HOFpii5u5Gabcef"
                },
                "person": {
                    "user_id": "ou_cc19b2bfb93f8a44db4b4d6eababcef"
                }
            }
        ]
    }
}

## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {}
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 1069301 | fail | Please try again. If the problem persists, contact the oncall personnel of the related service party. |
| 400 | 1069302 | param error | Check whether the parameter is valid. |
| 403 | 1069303 | forbidden | Check whether you have the permission to comment on the document. |
| 400 | 1069304 | docs had been deleted | Check whether the document to comment on has been deleted. |
| 400 | 1069305 | docs not exist | Check whether you have access to the document to comment on. |
| 400 | 1069306 | content review not pass | Check whether the comment contains invalid content. |
| 404 | 1069307 | not exist | Check whether you have access to the document to comment on and whether the person that is @mentioned by the comment or the document in Docs to comment on exists. |
| 400 | 1069308 | exceeded limit | The comment data has reached the limit. For more information, contact the Customer Service. |
| 400 | 1069399 | internal error | Please try again. If the problem persists, contact the oncall personnel of the related service party. |
