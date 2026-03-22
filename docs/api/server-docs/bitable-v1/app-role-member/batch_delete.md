---
title: "Batch delete members"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/bitable-v1/app-role-member/batch_delete"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/bitable/v1/apps/:app_token/roles/:role_id/members/batch_delete"
service: "bitable-v1"
resource: "app-role-member"
section: "Docs"
rate_limit: "10 per second"
scopes:
  - "bitable:app"
updated: "1753668119000"
---

# Batch delete members

Batch delete role members

## Prerequisites

To call the collaborator-related APIs, you need to ensure that Base has enabled advanced permissions and set custom roles. You can enable advanced permissions through the Update Base Metadata API and set custom roles through the Create Custom Role API.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/bitable/v1/apps/:app_token/roles/:role_id/members/batch_delete |
| HTTP Method | POST |
| Rate Limit | 10 per second |
| Supported app types | custom,isv |
| Required scopes | `bitable:app` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `app_token` | `string` | Base unique device identifier app_token description **Example value**: "bascnnKKvcoUblgmmhZkYqabcef" |
| `role_id` | `string` | Custom role ID **Example value**: "rolNGhPqks" | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `member_list` | `app.role.member_id[]` | Yes | List of members **Data validation rules**: - Maximum length: `100` |
|   `type` | `string` | No | member id type **Example value**: "open_id" **Optional values are**:  - `open_id`: member ID type is open_id - `union_id`: member ID type is union_id - `user_id`: member ID type is user_id - `chat_id`: member ID type is chat_id - `department_id`: member ID type is department_id - `open_department_id`: member ID type is open_department_id  **Default value**: `open_id` |
|   `id` | `string` | Yes | member ID **Example value**: "ou_35990a9d9052051a2fae9b2f1afabcef" | ### Request body example

{
    "member_list": [
        {
            "type": "open_id",
            "id": "ou_35990a9d9052051a2fae9b2f1afabcef"
        }
    ]
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
| 400 | 1254000 | WrongRequestJson | Request error |
| 400 | 1254001 | WrongRequestBody | Request body error |
| 400 | 1254002 | Fail | Internal error, have any questions can be consulting service |
| 400 | 1254003 | WrongBaseToken | AppToken error |
| 400 | 1254032 | InvalidRoleName | Invalid role name |
| 400 | 1254033 | RoleNameDuplicated | Role name duplicated |
| 404 | 1254040 | BaseTokenNotFound | AppToken not found |
| 404 | 1254047 | RoleIdNotFound | Role not found |
| 400 | 1254048 | MemberNotFound | Member not found |
| 400 | 1254110 | RoleExceedLimit | Role exceed limit, limited to 30 |
| 429 | 1254290 | TooManyRequest | Request too fast, try again later |
| 400 | 1254291 | Write conflict | The same data table does not support concurrent calls to the write interface, please check whether there is a concurrent call to the write interface. The writing interface includes: adding, modifying, and deleting records; adding, modifying, and deleting fields; modifying forms; modifying views, etc. |
| 400 | 1254301 | OperationTypeError | Base does not have advanced permissions enabled or does not support enabling advanced permissions |
| 403 | 1254302 | Permission denied. | No access rights, usually caused by the table opening of advanced permissions, please add a group containing applications in the advanced permissions settings and give this group read and write permissions |
| 400 | 1255001 | InternalError | Internal error, have any questions can be consulting service |
| 504 | 1255040 | Request timed out, please try again later | Try again |
| 400 | 1254036 | Base is copying, please try again later. | Base copy replicating, try again later |
