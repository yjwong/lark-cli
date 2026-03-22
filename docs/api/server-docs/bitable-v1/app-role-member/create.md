---
title: "Create member"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/bitable-v1/app-role-member/create"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/bitable/v1/apps/:app_token/roles/:role_id/members"
service: "bitable-v1"
resource: "app-role-member"
section: "Docs"
rate_limit: "10 per second"
scopes:
  - "bitable:app"
updated: "1753668118000"
---

# Create member

Create a member

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/bitable/v1/apps/:app_token/roles/:role_id/members |
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
| `app_token` | `string` | Base unique device identifier app_token description **Example value**: "appbcbWCzen6D8dezhoCH2RpMAh" |
| `role_id` | `string` | Role id **Example value**: "roljRpwIUt" | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `member_id_type` | `string` | No | Member id type **Example value**: open_id **Optional values are**:  - `open_id`: open_id - `union_id`: union_id - `user_id`: user_id - `chat_id`: chat_id - `department_id`: department_id. Before using this parameter, make sure the application has departmental visibility, refer to configure application availability scope - `open_department_id`: open_department_id. Before using this parameter, make sure the application has departmental visibility, refer to configure application availability scope  **Default value**: `open_id` | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `member_id` | `string` | Yes | Member id **Example value**: "ou_7dab8a3d3cdcc9da365777c7ad535d62" | ### Request body example

{
    "member_id": "ou_7dab8a3d3cdcc9da365777c7ad535d62"
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
| 200 | 1254000 | WrongRequestJson | Request error |
| 200 | 1254001 | WrongRequestBody | Request body error |
| 200 | 1254002 | Fail | Internal error, have any questions can be consulting service |
| 200 | 1254003 | WrongBaseToken | AppToken error |
| 200 | 1254010 | ReqConvError | Request error |
| 400 | 1254032 | InvalidRoleName | Invalid role name |
| 400 | 1254033 | RoleNameDuplicated | Role name duplicated |
| 400 | 1254036 | Base is copying, please try again later. | Base copy replicating, try again later |
| 200 | 1254040 | BaseTokenNotFound | AppToken not found |
| 404 | 1254047 | RoleIdNotFound | Role not found |
| 404 | 1254048 | MemberNotFound | Member not found |
| 400 | 1254110 | RoleExceedLimit | Role exceed limit, limited to 30 |
| 400 | 1254111 | MemberExceedLimit | Member exceed limit, limited to 200 |
| 200 | 1254290 | TooManyRequest | Request too fast, try again later |
| 200 | 1254291 | Write conflict | The same data table does not support concurrent calls to the write interface, please check whether there is a concurrent call to the write interface. The writing interface includes: adding, modifying, and deleting records; adding, modifying, and deleting fields; modifying forms; modifying views, etc. |
| 400 | 1254301 | OperationTypeError | Base does not have advanced permissions enabled or does not support enabling advanced permissions |
| 403 | 1254302 | Permission denied. | No access rights, usually caused by the table opening of advanced permissions, please add a group containing applications in the advanced permissions settings and give this group read and write permissions |
| 200 | 1255001 | InternalError | Internal error, have any questions can be consulting service |
| 200 | 1255002 | RpcError | Internal error, have any questions can be consulting service |
| 200 | 1255003 | MarshalError | Serialization failed, have any questions can be consulting service |
| 200 | 1255004 | UmMarshalError | Deserialization failed, have any questions can be consulting service |
| 504 | 1255040 | Request timed out, please try again later | Try again |
