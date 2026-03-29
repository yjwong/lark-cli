---
title: "DeletePermissionMember"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/permission-member/delete"
method: "DELETE"
api_path: "https://open.larksuite.com/open-apis/drive/v1/permissions/:token/members/:member_id"
service: "drive-v1"
resource: "permission-member"
section: "Docs"
scopes:
  - "drive:file"
  - "docs:doc"
  - "sheets:spreadsheet"
  - "drive:drive"
  - "wiki:wiki"
updated: "1647178821000"
---

# Delete permissions of a collaborator

This API is used to delete permissions of a document collaborator based on a filetoken.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/drive/v1/permissions/:token/members/:member_id |
| HTTP Method | DELETE |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `drive:file` `docs:doc` `sheets:spreadsheet` `drive:drive` `wiki:wiki` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" Learn more about obtaining and using access_token. | ::: note
For more information about how to call the AccessToken in the Docs API, see Get started with the Docs API.

### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `token` | `string` | Token of the file. For more information about how to obtain the token, see Overview. **Example value**: "doccnBKgoMyY5OMbUG6FioTXuBe" |
| `member_id` | `string` | ID of the member for whom you need to delete the permissions. This field corresponds to the `member_type` field. **Example value**: "ou_7dab8a3d3cdcc9da365777c7ad535d62" | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `type` | `string` | Yes | The type of file, placed in the query parameter, e.g.`?type = doc` **Example value**: "doc" **Optional values are**: - `doc`: Doc - `sheet`: Sheet - `file`: File in My Space - `wiki`: Wiki node (Partially supported) - `bitable`: Bitable - `docx`: Doc (Not supported) |
| `member_type` | `string` | Yes | Type of the member for whom you need to delete the permissions, for example, `?member_type drive.v1.permission.member.method.delete.query.prop.type.desc=$$$Type of the file, for example, `?type=doc`. This field is included in the query parameters. **Example value**: "openid" **Optional values are**: - `email`: Email address - `openid`: Open ID - `openchat`: Group ID on the Open Platform - `opendepartmentid`: Department ID on the Open Platform - `userid`: Custom user ID | ## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions | ### Response body example

```json
{
    "code": 0,
    "msg": "success"
}
```

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 1061001 | internal error | Internal service error, such as timeout or failure in processing error codes. |
| 403 | 1061002 | params error. | Check whether the request parameters are correct. |
| 400 | 1061003 | not found. | Check whether the destination upload node exists. |
| 403 | 1061004 | forbidden. | Check whether the current user or app has the permissions for the destination upload node. For example, check whether the user has the edit permission for the specified Docs for document upload. |
| 404 | 1061005 | auth failed. | Assume the correct user or app mode to access the API. |
| 500 | 1066001 | Internal Error | Internal service error, such as timeout or failure in processing error codes. |
| 500 | 1066002 | Concurrency error, please retry | Internal service error. Please try again. |
