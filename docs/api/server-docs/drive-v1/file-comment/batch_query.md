---
title: "Batch Query Comments"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/file-comment/batch_query"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/drive/v1/files/:file_token/comments/batch_query"
service: "drive-v1"
resource: "file-comment"
section: "Docs"
rate_limit: "100 per minute"
scopes:
  - "docs:document.comment:read"
  - "drive:drive"
  - "drive:drive:readonly"
field_scopes:
  - "contact:user.employee_id:readonly"
updated: "1710301984000"
---

# Batch Query Comments

This API is used to obtain cloud document comment information in batches according to comment ID list, including comment and reply ID, reply content, user ID of reviewer and reply person, etc. Supports returning global comments as well as local comments (which can be distinguished by the "is_whole" field).

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/drive/v1/files/:file_token/comments/batch_query |
| HTTP Method | POST |
| Rate Limit | 100 per minute |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `docs:document.comment:read` `drive:drive` `drive:drive:readonly` |
| Required field scopes | > The response body of the API contains the following sensitive fields, and they will be returned only after corresponding scopes are added. If you do not need the fields, it is not recommended that you request the scopes. `contact:user.employee_id:readonly` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `file_token` | `string` | Document token **Example value**: "doxbcdl03Vsxhm7Qmnj110abcef" | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `file_type` | `string` | Yes | Document type **Example value**: docx **Optional values are**:  - `doc`: doc sheet docx file - `sheet`: Form - `file`: File - `docx`: New version of documentation  |
| `user_id_type` | `string` | No | User ID categories **Example value**: open_id **Optional values are**:  - `open_id`: Identifies a user to an app. The same user has different Open IDs in different apps. How to get Open ID - `union_id`: Identifies a user to a tenant that acts as a developer. A user has the same Union ID in apps developed by the same developer, and has different Union IDs in apps developed by different developers. A developer can use Union ID to link the same user's identities in multiple apps.How to get Union ID - `user_id`: Identifies a user to a tenant. The same user has different User IDs in different tenants. In one single tenant, a user has the same User ID in all apps （including store apps）. User ID is usually used to communicate user data between different apps. How to get User ID  **Default value**: `open_id` **When the value is `user_id`, the following field scopes are required**: `contact:user.employee_id:readonly` | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `comment_ids` | `string[]` | Yes | Comment ID list **Example value**: ["1654857036541812356"] | ### Request body example

{
    "comment_ids": [
        "1654857036541812356"
    ]
}

## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `items` | `file.comment[]` | Comment replies |
|     `comment_id` | `string` | Comment ID. This field can be left empty when you create a comment. If this field is specified, the comment is deemed as a reply to an existing comment. |
|     `user_id` | `string` | User ID |
|     `create_time` | `int` | Creation time |
|     `update_time` | `int` | Last update |
|     `is_solved` | `boolean` | Specifies whether the comment is solved. |
|     `solved_time` | `int` | Time when the comment is solved |
|     `solver_user_id` | `string` | User ID of the comment solver |
|     `has_more` | `boolean` | Have more replies |
|     `page_token` | `string` | Reply to the pagination mark |
|     `is_whole` | `boolean` | Whether it is a full text comment |
|     `quote` | `string` | The quote of document |
|     `reply_list` | `reply_list` | List of replies in the comment |
|       `replies` | `file.comment.reply[]` | Reply list |
|         `reply_id` | `string` | Reply ID |
|         `user_id` | `string` | User ID |
|         `create_time` | `int` | Creation time |
|         `update_time` | `int` | Last update |
|         `content` | `reply_content` | Reply content |
|           `elements` | `reply_element[]` | Reply content |
|             `type` | `string` | Content elements of the reply **Optional values are**:  - `text_run`: Plain text - `docs_link`: Link to @mentioning a document in Docs - `person`: Contact to @mention  |
|             `text_run` | `text_run` | Text content |
|               `text` | `string` | Reply with plain text |
|             `docs_link` | `docs_link` | Text content |
|               `url` | `string` | @mentioned document in a reply. |
|             `person` | `person` | Text content |
|               `user_id` | `string` | @mentioned contact in a reply |
|         `extra` | `reply_extra` | Store other type replies item |
|           `image_list` | `string[]` | Image token list | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
        "items": [
            {
                "comment_id": "6916106822734512356",
                "user_id": "ou_cc19b2bfb93f8a44db4b4d6eababcef",
                "create_time": 1610281603,
                "update_time": 1610281603,
                "is_solved": false,
                "solved_time": 1610281603,
                "solver_user_id": "null",
                "has_more": false,
                "page_token": "6916106822734512356",
                "is_whole": true,
                "quote": "The quote of document",
                "reply_list": {
                    "replies": [
                        {
                            "reply_id": "6916106822734512356",
                            "user_id": "ou_cc19b2bfb93f8a44db4b4d6eab2abcef",
                            "create_time": 1610281603,
                            "update_time": 1610281603,
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
                            },
                            "extra": {
                                "image_list": [
                                    "xxfsffsabcef"
                                ]
                            }
                        }
                    ]
                }
            }
        ]
    }
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
| 400 | 1064230 | locked_for_data_migration | Data migrating, temporarily unable to upload. |
