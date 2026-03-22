---
title: "Solve or Restore a Comment"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/file-comment/patch"
method: "PATCH"
api_path: "https://open.larksuite.com/open-apis/drive/v1/files/:file_token/comments/:comment_id"
service: "drive-v1"
resource: "file-comment"
section: "Docs"
rate_limit: "1000 per minute、50 per second"
scopes:
  - "docs:doc"
  - "docs:doc:readonly"
  - "drive:drive"
  - "drive:drive:readonly"
  - "sheets:spreadsheet"
  - "sheets:spreadsheet:readonly"
updated: "1710301985000"
---

# Solve or Restore a Comment

Solves or restores a comment on a document in Docs.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/drive/v1/files/:file_token/comments/:comment_id |
| HTTP Method | PATCH |
| Rate Limit | 1000 per minute、50 per second |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `docs:doc` `docs:doc:readonly` `drive:drive` `drive:drive:readonly` `sheets:spreadsheet` `sheets:spreadsheet:readonly` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `file_token` | `string` | Token of the document **Example value**: "doccnGp4UK1UskrOEJwBXd3****" |
| `comment_id` | `string` | Comment ID **Example value**: "6916106822734578184" | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `file_type` | `string` | Yes | Type of the document **Example value**: doc **Optional values are**:  - `doc`: Doc - `sheet`: Sheet - `file`: File - `docx`: New Doc  | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `is_solved` | `boolean` | Yes | Icon that indicates whether the comment is solved. **Example value**: true | ### Request body example

{
    "is_solved": true
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
