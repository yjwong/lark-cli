---
title: "List roles"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/bitable-v1/app-role/list"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/bitable/v1/apps/:app_token/roles"
service: "bitable-v1"
resource: "app-role"
section: "Docs"
rate_limit: "20 per second"
scopes:
  - "bitable:app"
  - "bitable:app:readonly"
updated: "1694750005000"
---

# List roles

Get all roles according to app_token

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/bitable/v1/apps/:app_token/roles |
| HTTP Method | GET |
| Rate Limit | 20 per second |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `bitable:app` `bitable:app:readonly` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `app_token` | `string` | Base unique device identifier app_token description **Example value**: "appbcbWCzen6D8dezhoCH2RpMAh" | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `page_size` | `int` | No | **Example value**: 10 **Data validation rules**: - Maximum value: `30` |
| `page_token` | `string` | No | Page identifier. It is not filled in the first request, indicating traversal from the beginning; when there will be more groups, the new page_token will be returned at the same time, and the next traversal can use the page_token to get more groups **Example value**: roljRpwIUt | ## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `items` | `app.role[]` | Role information |
|     `role_name` | `string` | Role name |
|     `role_id` | `string` | Role id |
|     `table_roles` | `app.role.table_role[]` | Table role |
|       `table_name` | `string` | Table name |
|       `table_id` | `string` | Table id |
|       `table_perm` | `int` | Table perm **Optional values are**:  - `0`: No permissions - `1`: View only - `2`: Can edit records - `4`: Can edit records and fields  |
|       `rec_rule` | `app.role.table_role.rec_rule` | Record filter rule |
|         `conditions` | `app.role.table_role.rec_rule.condition[]` | Coditions |
|           `field_name` | `string` | Field name |
|           `operator` | `string` | Operator **Optional values are**:  - `is`: Is - `isNot`: Is not - `contains`: Contains - `doesNotContain`: Does not contain - `isEmpty`: Is empty - `isNotEmpty`: Is not empty  |
|           `value` | `string[]` | Option id |
|           `field_type` | `int` | Field type |
|         `conjunction` | `string` | Conjunction **Optional values are**:  - `and`: And - `or`: Or  |
|         `other_perm` | `int` | Other perm **Optional values are**:  - `0`: Cannot view - `1`: View Only  |
|       `field_perm` | `app.role.table_role.field_perm` | Permission of fields, only valid when `table_perm` is 2.  The type is map, key is field name, value is permission of field. **Optional values are**: - `1`: View only - `2`: Can edit records |
|       `allow_add_record` | `boolean` | Added record permission, only meaningful when the table_perm is 2, used to set whether the record can be added |
|       `allow_delete_record` | `boolean` | Delete record permission, meaningful only when the table_perm is 2, used to set whether the record can be deleted |
|     `block_roles` | `app.role.block_role[]` | Block role |
|       `block_id` | `string` | Block id, Such as dashboard block id in list dashboards method |
|       `block_type` | `string` | Block type **Optional values are**:  - `dashboard`: Dashboard  |
|       `block_perm` | `int` | Block perm **Optional values are**:  - `0`: No permissions - `1`: View only  |
|   `page_token` | `string` | Page identifier, when has_more is true, a new page_token will also be returned. Otherwise, page_token will not be returned |
|   `has_more` | `boolean` | Whether the response body has more parameters |
|   `total` | `int` | Total of this Base roles. | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
        "items": [
            {
                "role_id": "rola66uIZT",
                "role_name": "role1",
                "table_roles": [
                    {
                        "table_name": "table1",
                        "table_id": "tblFIgBzKEq75HSE",
                        "table_perm": 2,
                        "allow_add_record": false,
                        "allow_delete_record": true,
                        "rec_rule": {
                            "conjunction": "or",
                            "conditions": [
                                {
                                    "field_name": "Single option",
                                    "field_type": 3,
                                    "operator": "is",
                                    "value": [
                                        "optbdVHf4q"
                                    ]
                                },
                                {
                                    "field_name": "Person", // Person contains current visitor
                                    "field_type": 11,
                                    "operator": "contains",
                                    "value": null
                                },
                                {
                                    "field_name": "", // Created by contains current visitor
                                    "field_type": 1003,
                                    "operator": "contains",
                                    "value": null
                                }
                            ],
                            "other_perm": 0
                        },
                        "field_perm": {
                            "Single option": 1,
                            "Age": 2
                        }
                    },
                    {
                        "table_name": "table2",
                        "table_id": "tblMPI6OC1aWvTvs",
                        "table_perm": 1,
                        "rec_rule": {
                            "conditions": [
                                {
                                    "field_name": "Person",
                                    "field_type": 11,
                                    "operator": "contains",
                                    "value": null
                                },
                                {
                                    "field_name": "Multiple options",
                                    "field_type": 4,
                                    "operator": "is",
                                    "value": [
                                        "opttgKOTSt",
                                        "optWcdXR0W"
                                    ]
                                }
                            ],
                            "other_perm": 0,
                            "conjunction": "and"
                        }
                    },
                    {
                        "table_name": "table3",
                        "table_id": "tblmkLF7Tg6IWbRb",
                        "table_perm": 0
                    },
                    {
                        "table_name": "table4",
                        "table_id": "tbl5VQHDTms19Qe7",
                        "table_perm": 4
                    }
                ],
                "block_roles": [
                    {
                        "block_id": "blknkqrP3RqUkcAW",
                        "block_type": "dashboard",
                        "block_perm": 0
                    },
                    {
                        "block_id": "blkAjxjWKvbBi7EA",
                        "block_type": "dashboard",
                        "block_perm": 1
                    }
                ]
            }
        ],
        "page_token": "rola66uIZT",
        "has_more": false,
        "total": 1
    }
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
| 400 | 1254110 | RoleExceedLimit | Role exceed limit, limited to 30 |
| 200 | 1254290 | TooManyRequest | Request too fast, try again later |
| 200 | 1254291 | Write conflict | The same data table does not support concurrent calls to the write interface, please check whether there is a concurrent call to the write interface. The writing interface includes: adding, modifying, and deleting records; adding, modifying, and deleting fields; modifying forms; modifying views, etc. |
| 400 | 1254301 | OperationTypeError | Base does not have advanced permissions enabled or does not support enabling advanced permissions |
| 403 | 1254302 | Permission denied. | No access rights, usually caused by the table opening of advanced permissions, please add a group containing applications in the advanced permissions settings and give this group read and write permissions |
| 200 | 1255001 | InternalError | Internal error, have any questions can be consulting service |
| 200 | 1255002 | RpcError | Internal error, have any questions can be consulting service |
| 200 | 1255003 | MarshalError | Serialization failed, have any questions can be consulting service |
| 200 | 1255004 | UmMarshalError | Deserialization failed, have any questions can be consulting service |
| 504 | 1255040 | Request timed out, please try again later | Try again |
