---
title: "Upgrade form"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/bitable-v1/app-table-form/upgrade"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/bitable/v1/apps/:app_token/tables/:table_id/forms/:form_id/upgrade"
service: "bitable-v1"
resource: "app-table-form"
section: "Docs"
rate_limit: "20 per second"
scopes:
  - "base:form:update"
updated: "1772538373000"
---

# Upgrade form

Upgrade legacy forms to collection forms

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/bitable/v1/apps/:app_token/tables/:table_id/forms/:form_id/upgrade |
| HTTP Method | POST |
| Rate Limit | 20 per second |
| Supported app types | custom,isv |
| Required scopes | `base:form:update` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `app_token` | `string` | Base unique device identifier app_token description **Example value**: "bascnv1jIEppJdTCn3jOosabcef" |
| `table_id` | `string` | Base data table unique device identifier table_id description **Example value**: "tblz8nadEUdxNMt5" |
| `form_id` | `string` | The unique identifier of a form in a Bitable. A form is also a type of view and is obtained in the same way as a view_id: - In the URL address bar of Bitable, form_id is the highlighted part in the image below: ![view_id.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/140668632c97e0095832219001d17c54_DJMgVH9x2S.png?height=748&lazyload=true&width=2998) - Obtained through the List View interface. Temporarily unable to get Bitable's form_id embedded in cloud documents **Example value**: "vew6oMbAa4" | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `form_name` | `string` | Yes | Upgraded form name **Example value**: "Feedback on Documentation" **Data validation rules**: - Length range: `1` ～ `100` characters - Regular expression: `^[^[\]]+$` |
| `display_mode` | `string` | Yes | Form layout mode. **Example value**: "one_question_per_page" **Optional values are**:  - `traditional`: traditional layout - `one_question_per_page`: one page layout  | ### Request body example

{
    "form_name": "Feedback on Documentation",
    "display_mode": "one_question_per_page"
}

## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `form` | `upgraded_form` | Upgraded form |
|     `id` | `string` | Upgraded form ID | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
        "form": {
            "id": "vew6oMbAa4"
        }
    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 1254000 | WrongRequestJson | request body error |
| 400 | 1254001 | WrongRequestBody | request body error |
| 500 | 1254002 | Fail | Internal error, please contact [Technical Support](https://applink.larksuite.com/client/helpdesk/open?id=6626260912531570952) |
| 400 | 1254003 | WrongBaseToken | app_token mistake |
| 400 | 1254004 | WrongTableId | table_id mistake |
| 400 | 1254005 | WrongViewId | view_id mistake |
| 400 | 1254008 | EmptyView | empty view |
| 400 | 1254019 | InvalidViewType | Invalid view type |
| 400 | 1254036 | Bitable is copying, please try again later. | Copy Bitable is an asynchronous operation. This error code indicates that the current Bitable is still being copied, and the current Bitable cannot be operated during the copy. You need to wait for the copy to complete before operating. |
| 404 | 1254040 | BaseTokenNotFound | app_token doesn't exist |
| 404 | 1254041 | TableIdNotFound | table_id doesn't exist |
| 404 | 1254042 | ViewIdNotFound | view_id doesn't exist |
| 400 | 1254124 | The name of the new form is empty | The new form name is not allowed to be empty |
| 400 | 1254125 | The name of the new form is invalid. | Invalid new form name |
| 400 | 1254126 | The display mode of the new form is invalid. | The layout mode for the new form is invalid |
| 400 | 1254127 | The name of the new form is duplicated. | The name of the new form duplicates the name of the existing view |
| 400 | 1254128 | The form is already a new version, no upgrade required. | It is already a new version of the form and does not need to be upgraded. |
| 500 | 1254200 | Something went wrong | Something went wrong, please contact [Technical Support](https://applink.larksuite.com/client/helpdesk/open?id=6626260912531570952) |
| 400 | 1254290 | TooManyRequest | Request is too fast, try again later |
| 400 | 1254291 | LockNotObtainedError | The same data table does not support concurrent calls to write interfaces. Please check whether there are concurrent calls to write interfaces. Write interfaces include: add, modify, delete records; add, modify, delete fields; modify forms; modify views; upgrade forms, etc. |
| 403 | 1254302 | RolePermNotAllow | The calling identity lacks Bitable's advanced permissions. You need to grant advanced permissions to the calling identity: - To grant advanced permissions to users, you need to add manageable permissions to the current user in the **Share** entry at the top right of the Bitable page. - To grant advanced permissions to the app, you need to add manageable permissions to the app through the top right of the Bitable page **「...」** - > **「... More 」** ->**「Add Documentation App」** entry. **Attention**: Before **adding a document application**, you need to make sure that the target application has at least one Bitable API permission. Otherwise, you will not be able to search for the target application in the document application window. - You can also add users or a group containing applications in the **Bitable advanced permission settings**, and give this group customized permissions such as reading and writing. |
| 403 | 1254304 | You are not authorized to perform this operation. | Only Enterprise and Ultimate Lark support row and column permissions |
| 403 | 1254305 | PermControl | This operation is subject to restrictions, such as tenant privacy level restrictions, wiki space permission restrictions, etc |
| 500 | 1254607 | Data not ready, please try again later | The error is generally caused by the pre-operation not being completed, or the data of this operation is too large, and the server calculation timed out. When encountering this error code, it is recommended to wait for a period of time and try again. Usually there are the following reasons: - **Frequent editing operations**: Developers' editing operations on Bitable are very frequent. It may lead to timeouts due to waiting for the pre-operation processing to complete taking too long. Bitable's underlying processing of data tables is based on the serial way of the version dimension and does not support concurrency. Therefore, such errors are prone to occur when concurrent requests are made. It is not recommended that developers make concurrent requests for a single data table. - **Batch operation load**: When developers perform batch addition, deletion and other operations in Bitable, if the data volume of the data table is very large, it may cause a single request to take too long, eventually resulting in request timed out. It is recommended that developers appropriately reduce the page_size of batch requests to reduce request time. - **Quotas and computational overhead**: Quotas are based on a single document dimension. If the reading interface involves computational logic such as formula calculation and sorting, it will consume more resources. For example, concurrent reading of multiple data tables under a document may also cause the document to block. |
| 500 | 1255001 | InternalError | Internal error, please contact [Technical Support](https://applink.larksuite.com/client/helpdesk/open?id=6626260912531570952) |
| 500 | 1255002 | RpcError | Internal error, please contact [Technical Support](https://applink.larksuite.com/client/helpdesk/open?id=6626260912531570952) |
| 500 | 1255003 | MarshalError | Serialization error, please contact [Technical Support](https://applink.larksuite.com/client/helpdesk/open?id=6626260912531570952) |
| 500 | 1255004 | UmMarshalError | Deserialization error, please contact [Technical Support](https://applink.larksuite.com/client/helpdesk/open?id=6626260912531570952) |
| 504 | 1255040 | Request timed out, please try again later | Try again. |
