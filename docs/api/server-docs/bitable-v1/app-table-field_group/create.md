---
title: "Create field group"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/bitable-v1/app-table-field_group/create"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/bitable/v1/apps/:app_token/tables/:table_id/field_groups"
service: "bitable-v1"
resource: "app-table-field_group"
section: "Docs"
rate_limit: "10 per second"
scopes:
  - "base:field_group:create"
updated: "1772538369000"
---

# Create field group

Create group for fields of the Bitable data table

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/bitable/v1/apps/:app_token/tables/:table_id/field_groups |
| HTTP Method | POST |
| Rate Limit | 10 per second |
| Supported app types | custom,isv |
| Required scopes | `base:field_group:create` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `app_token` | `string` | Base unique device identifier app_token description **Example value**: "bascnv1jIEppJdTCn3jOosabcef" |
| `table_id` | `string` | Base data table unique device identifier table_id description **Example value**: "tblz8nadEUdxNMt5" | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `field_groups` | `field_group[]` | Yes | To add a field group list **Data validation rules**: - Length range: `1` ～ `300` |
|   `id` | `string` | No | ID of the field group **Example value**: "fldPTb0U2y" |
|   `name` | `string` | Yes | The name of the field group **Example value**: "Customer Information" **Data validation rules**: - Length range: `1` ～ `100` characters - Regular expression: `^[^[\]]+$` |
|   `children` | `field_group_child[]` | Yes | Members of a field group **Data validation rules**: - Length range: `1` ～ `300` |
|     `type` | `string` | Yes | group member type **Example value**: "field" **Optional values are**:  - `field`: field  |
|     `id` | `string` | Yes | Group Member ID **Example value**: "fldPTb0U2y" |
|   `description` | `string` | No | Description of field grouping **Example value**: "Customer Information Field Grouping" **Data validation rules**: - Length range: `0` ～ `2000` characters | ### Request body example

{
    "field_groups": [
        {
            "id": "fldPTb0U2y",
            "name": "Customer Information",
            "children": [
                {
                    "type": "field",
                    "id": "fldPTb0U2y"
                }
            ],
            "description": "Customer Information Field Grouping"
        }
    ]
}

## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `field_groups` | `string` | Contents of field grouping | ### Response body example

{"code":0,
"msg":"success",
"data":{"field_groups":"[   {     "id": "fldjX7dUj5",     "name": "编组1"   },   {     "id": "fldjX7dUj6",     "name": "编组2"   } ]"}}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 1254000 | WrongRequestJson | request body error |
| 400 | 1254001 | WrongRequestBody | request body error |
| 500 | 1254002 | Fail | Internal error, please contact [Technical Support](https://applink.larksuite.com/client/helpdesk/open?id=6626260912531570952) |
| 400 | 1254003 | WrongBaseToken | app_token mistake |
| 400 | 1254004 | WrongTableId | table_id mistake |
| 400 | 1254009 | WrongFieldId | field_id mistake |
| 400 | 1254036 | Bitable is copying, please try again later. | Copy Bitable is an asynchronous operation. This error code indicates that the current Bitable is still being copied, and the current Bitable cannot be operated during the copy. You need to wait for the copy to complete before operating. |
| 404 | 1254040 | BaseTokenNotFound | app_token doesn't exist |
| 404 | 1254041 | TableIdNotFound | table_id doesn't exist |
| 404 | 1254044 | FieldIdNotFound | field_id doesn't exist |
| 400 | 1254114 | The name of field group is empty | The name of the field group cannot be empty |
| 400 | 1254115 | The name of field group is invalid | The name of the field group is invalid |
| 400 | 1254116 | The description of field group is too long | The description of the field group is too long |
| 400 | 1254117 | The name of field group is duplicated | Duplicate names for field groups |
| 400 | 1254118 | The children of field group is empty | No child members are included in the field group |
| 400 | 1254119 | The children of field group is too large | A field group can contain up to 300 child elements |
| 400 | 1254120 | The field groups to be created is empty | The number of field group to be created is empty |
| 400 | 1254121 | The field groups to be created is too large | Create up to 300 field groups at a time |
| 400 | 1254122 | Child belongs to multiple field groups | Child elements are not allowed to be referenced by multiple field groups |
| 400 | 1254123 | The child type of field group is invalid | Invalid child element type for field group |
| 400 | 1254290 | TooManyRequest | Request is too fast, try again later |
| 400 | 1254291 | Write conflict | The same data table does not support concurrent calls to write interfaces. Please check whether there are concurrent calls to write interfaces. Write interfaces include: add, modify, delete records; add, modify, delete fields; modify forms; modify views; upgrade forms, etc. |
| 403 | 1254302 | RolePermNotAllow | The calling identity lacks Bitable's advanced permissions. You need to grant advanced permissions to the calling identity: - To grant advanced permissions to users, you need to add manageable permissions to the current user in the **Share** entry at the top right of the Bitable page. - To grant advanced permissions to the app, you need to add manageable permissions to the app through the top right of the Bitable page **「...」** - > **「... More 」** ->**「 Add Documentation App」** entry. **Attention**: Before  **adding a document application**, you need to make sure that the target application has at least one Bitable API permission. Otherwise, you will not be able to search for the target application in the document application window. - You can also add users or a group containing applications in the **Bitable advanced permission settings**, and give this group customized permissions such as reading and writing. |
| 403 | 1254304 | You are not authorized to perform this operation. | Only Enterprise and Ultimate Lark support row and column permissions |
| 500 | 1254607 | Data not ready, please try again later | The error is generally caused by the pre-operation not being completed, or the data of this operation is too large, and the server calculation timed out. When encountering this error code, it is recommended to wait for a period of time and try again. Usually there are the following reasons: - **Frequent editing operations**: Developers' editing operations on Bitable are very frequent. It may lead to timeouts due to waiting for the pre-operation processing to complete taking too long. Bitable's underlying processing of data tables is based on the serial way of the version dimension and does not support concurrency. Therefore, such errors are prone to occur when concurrent requests are made. It is not recommended that developers make concurrent requests for a single data table. - **Batch operation load**: When developers perform batch addition, deletion and other operations in Bitable, if the data volume of the data table is very large, it may cause a single request to take too long, eventually resulting in request timed out. It is recommended that developers appropriately reduce the page_size of batch requests to reduce request time. - **Quotas and computational overhead**: Quotas are based on a single document dimension. If the reading interface involves computational logic such as formula calculation and sorting, it will consume more resources. For example, concurrent reading of multiple data tables under a document may also cause the document to block. |
| 500 | 1255001 | InternalError | Internal error, please contact [Technical Support](https://applink.larksuite.com/client/helpdesk/open?id=6626260912531570952) |
| 500 | 1254200 | Something went wrong | Internal error, please contact [Technical Support](https://applink.larksuite.com/client/helpdesk/open?id=6626260912531570952) |
| 500 | 1255003 | MarshalError | Serialization error, please contact [Technical Support](https://applink.larksuite.com/client/helpdesk/open?id=6626260912531570952) |
| 500 | 1255004 | UmMarshalError | Deserialization error, please contact [Technical Support](https://applink.larksuite.com/client/helpdesk/open?id=6626260912531570952) |
| 504 | 1255040 | Request timed out, please try again later | Try again. |
