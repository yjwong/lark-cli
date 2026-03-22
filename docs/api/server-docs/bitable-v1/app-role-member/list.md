---
title: "List members"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/bitable-v1/app-role-member/list"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/bitable/v1/apps/:app_token/roles/:role_id/members"
service: "bitable-v1"
resource: "app-role-member"
section: "Docs"
rate_limit: "20 per second"
scopes:
  - "bitable:app"
  - "bitable:app:readonly"
updated: "1753668118000"
---

# List members

Get all members according to app_token and role_id

## Prerequisites

To call the collaborator-related APIs, you need to ensure that Base has enabled advanced permissions and set custom roles. You can enable advanced permissions through the Update Base Metadata API and set custom roles through the Create Custom Role API.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/bitable/v1/apps/:app_token/roles/:role_id/members |
| HTTP Method | GET |
| Rate Limit | 20 per second |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `bitable:app` `bitable:app:readonly` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `app_token` | `string` | Base unique device identifier app_token description **Example value**: "appbcbWCzen6D8dezhoCH2RpMAh" |
| `role_id` | `string` | Role id **Example value**: "roljRpwIUt" | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `page_size` | `int` | No | Page size **Example value**: 100 **Data validation rules**: - Maximum value: `100` |
| `page_token` | `string` | No | Page identifier. It is not filled in the first request, indicating traversal from the beginning; when there will be more groups, the new page_token will be returned at the same time, and the next traversal can use the page_token to get more groups **Example value**: xxxxx | ## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `items` | `app.role.member[]` | Member information |
|     `open_id` | `string` | User's open_id |
|     `union_id` | `string` | User's union_id |
|     `user_id` | `string` | User's user_id |
|     `chat_id` | `string` | Chat's chat_id |
|     `department_id` | `string` | Department's department_id |
|     `open_department_id` | `string` | Department's open_department_id |
|     `member_name` | `string` | Member name |
|     `member_en_name` | `string` | Member English name |
|     `member_type` | `string` | Member type **Optional values are**:  - `user`: User - `chat`: Chat - `department`: Department  |
|   `has_more` | `boolean` | Whether the response body has more parameters |
|   `page_token` | `string` | Page identifier, when has_more is true, a new page_token will also be returned. Otherwise, page_token will not be returned |
|   `total` | `int` | Total | ### Response body example

{
    "msg": "success",
    "data": {
        "items": [
            {
                "member_type": "user",
                "member_name": "Tom",
                "member_en_name": "Tom",
                "open_id": "ou_xxxxxxxxxxxxxxxx",
                "union_id": "on_xxxxxxxxxxxxxxxx",
                "user_id": "xxxxxx"
            },
            {
                "member_type": "chat",
                "member_name": "design-chat",
                "member_en_name": "design-chat",
                "chat_id": "oc_xxxxxxxxxxxxxxxx"
            },
            {
                "member_type": "department",
                "member_name": "design-center",
                "member_en_name": "design-center",
                "department_id": "xxxxxxxxx",
                "open_department_id": "od-xxxxxxxxxxxxxxxx"
            }
        ],
        "page_token": "xxxxxxxxx",
        "total": 3,
        "has_more": false
    },
    "code": 0
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
| 400 | 1254200 | Something went wrong. Please contact technical support at https://applink.larksuite.com/client/helpdesk/open?id=6626260912531570952 | Internal error, have any questions can be consulting service |
| 200 | 1254290 | TooManyRequest | Request too fast, try again later |
| 200 | 1254291 | Write conflict | The same data table does not support concurrent calls to the write interface, please check whether there is a concurrent call to the write interface. The writing interface includes: adding, modifying, and deleting records; adding, modifying, and deleting fields; modifying forms; modifying views, etc. |
| 400 | 1254301 | OperationTypeError | Base does not have advanced permissions enabled or does not support enabling advanced permissions |
| 403 | 1254302 | Permission denied. | No access rights, usually caused by the table opening of advanced permissions, please add a group containing applications in the advanced permissions settings and give this group read and write permissions |
| 200 | 1255001 | InternalError | Internal error, have any questions can be consulting service |
| 200 | 1255002 | RpcError | Internal error, have any questions can be consulting service |
| 200 | 1255003 | MarshalError | Serialization failed, have any questions can be consulting service |
| 200 | 1255004 | UmMarshalError | Deserialization failed, have any questions can be consulting service |
| 504 | 1255040 | Request timed out, please try again later | Try again |
