---
title: "Update document sharing settings"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/permission-public/patch"
method: "PATCH"
api_path: "https://open.larksuite.com/open-apis/drive/v1/permissions/:token/public"
service: "drive-v1"
resource: "permission-public"
section: "Docs"
scopes:
  - "drive:file"
  - "docs:doc"
  - "sheets:spreadsheet"
  - "drive:drive"
  - "wiki:wiki"
updated: "1647178821000"
---

# Update common settings of a document

This API is used to update the common settings of a document based on a filetoken.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/drive/v1/permissions/:token/public |
| HTTP Method | PATCH |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `drive:file` `docs:doc` `sheets:spreadsheet` `drive:drive` `wiki:wiki` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" Learn more about obtaining and using access_token. |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ::: note
For more information about how to call the AccessToken in the Docs API, see Get started with the Docs API.

### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `token` | `string` | Token of the file. For more information about how to obtain the token, see Overview. **Example value**: "doccnBKgoMyY5OMbUG6FioTXuBe" | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `type` | `string` | Yes | Type of the permission object, for example, `?type drive.v1.enum.TokenType.option.Wiki.desc=$$$Wiki node (Partially supported) **Example value**: "doc" **Optional values are**: - `doc`: Doc - `sheet`: Sheet - `file`: File in My Space - `wiki`: Wiki node (Partially supported) - `bitable`: Bitable - `docx`: Doc (Not supported) | ::: note
**Note** `wiki`: Wiki node. The following settings are not supported:
- `external_access`: Specifies whether sharing out of the tenant is allowed.
- `share_entity`: Specifies who can add and manage collaborators.
- `invite_external`: Specifies whether users without full access and non-owners are allowed to invite external participants.
- `link_share_entity`: Settings of link sharing.
  - `tenant_readable`: All users who obtain the link can read the document.
  - `tenant_editable`: All users who obtain the link can edit the document.

### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `external_access` | `boolean` | No | Indicates whether sharing out of the tenant is allowed. **Example value**: true |
| `security_entity` | `string` | No | Settings of the permission to make a copy of, print, export, and copy the document. **Example value**: "anyone_can_view" **Optional values are**: - `anyone_can_view`: All users who can access this document. - `anyone_can_edit`: Users with the edit permission. |
| `comment_entity` | `string` | No | Settings of the permission to comment on the document. **Example value**: "anyone_can_view" **Optional values are**: - `anyone_can_view`: All users who can access this document. - `anyone_can_edit`: Users with the edit permission. |
| `share_entity` | `string` | No | Indicates who can add and manage collaborators. **Example value**: "anyone" **Optional values are**: - `anyone`: All users who can read or edit this document. - `same_tenant`: All users who can read or edit this document in the organization. - `only_full_access`: Only users with full access. |
| `link_share_entity` | `string` | No | Link sharing **Example value**: "tenant_readable" **Optional values are**: - `tenant_readable`: Users who obtain the link in the organization can read the document. - `tenant_editable`: Users who obtain the link in the organization can edit the document. - `anyone_readable`: All users who obtain the link can read the document. (This option is valid only when `external_access drive.v1.type.permission_public.prop.link_share_entity.option.Closed.desc=$$$Link sharing is disabled. - `anyone_editable`: All users who obtain the link can edit the document. (This option is valid only when `external_access drive.v1.enum.TokenType.option.Sheet.desc=$$$Sheet - `closed`: Closed share link. |
| `invite_external` | `boolean` | No | Specifies whether users without full access and non-owners are allowed to invite external participants. **Example value**: true | ### Request body example

```json
{
    "external_access": true,
    "security_entity": "anyone_can_view",
    "comment_entity": "anyone_can_view",
    "share_entity": "anyone",
    "link_share_entity": "tenant_readable",
    "invite_external": true
}
```

## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
| ∟ `permission_public` | `permission_public` | New common settings of the document after the current update. |
| ∟ `external_access` | `boolean` | Indicates whether sharing out of the tenant is allowed. |
| ∟ `security_entity` | `string` | Settings of the permission to make a copy of, print, export, and copy the document. **Optional values are**: - `anyone_can_view`: All users who can access this document. - `anyone_can_edit`: Users with the edit permission. |
| ∟ `comment_entity` | `string` | Settings of the permission to comment on the document. **Optional values are**: - `anyone_can_view`: All users who can access this document. - `anyone_can_edit`: Users with the edit permission. |
| ∟ `share_entity` | `string` | Indicates who can add and manage collaborators. **Optional values are**: - `anyone`: All users who can read or edit this document. - `same_tenant`: All users who can read or edit this document in the organization. - `only_full_access`: Only users with full access. |
| ∟ `link_share_entity` | `string` | Link sharing **Optional values are**: - `tenant_readable`: Users who obtain the link in the organization can read the document. - `tenant_editable`: Users who obtain the link in the organization can edit the document. - `anyone_readable`: All users who obtain the link can read the document. (This option is valid only when `external_access drive.v1.type.permission_public.prop.link_share_entity.option.Closed.desc=$$$Link sharing is disabled. - `anyone_editable`: All users who obtain the link can edit the document. (This option is valid only when `external_access drive.v1.enum.TokenType.option.Sheet.desc=$$$Sheet - `closed`: Closed share link. |
| ∟ `invite_external` | `boolean` | Specifies whether users without full access and non-owners are allowed to invite external participants. | ### Response body example

```json
{
    "code": 0,
    "msg": "success",
    "data": {
        "permission_public": {
            "external_access": true,
            "security_entity": "anyone_can_view",
            "comment_entity": "anyone_can_view",
            "share_entity": "anyone",
            "link_share_entity": "tenant_readable",
            "invite_external": true
        }
    }
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
