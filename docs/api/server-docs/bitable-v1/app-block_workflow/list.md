---
title: "List Workflows"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/bitable-v1/app-block_workflow/list"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/bitable/v1/apps/:app_token/block_workflows"
service: "bitable-v1"
resource: "app-block_workflow"
section: "Docs"
rate_limit: "20 per minute"
scopes:
  - "bitable:app"
updated: "1770864574000"
---

# List workflow

This method is used to return all workflows in a base table, and base table administrators can manage workflows in the table through this method

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/bitable/v1/apps/:app_token/block_workflows |
| HTTP Method | GET |
| Rate Limit | 20 per minute |
| Supported app types | custom,isv |
| Required scopes | `bitable:app` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `app_token` | `string` | The unique identifier of the Base App. Different forms of Base have different ways of obtaining app_token: - If a Base URL begins with ** feishu.cn/base **, the Base app_token is highlighted in the image below: ! [app_token] (//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/6916f8cfac4045ba6585b90e3afdfb0a_GxbfkJHZBa.png?height=766&lazyload=true&width=3004) - If Base URL starts with ** feishu.cn/wiki **, you need to call the Knowledge Base related [Get Wiki Workspace Node Information] (/ssl: ttdoc/ukTMukTMukTM/uUDN04SN0QjL1QDN/wiki-v2/space/get_node) interface to get Base app_token. When the value of obj_type is Base, the value of the obj_token field is the app_token of Base. For more information, refer to [Base app_token Acquisition Method] (/ssl: ttdoc/ukTMukTMukTM/uUDN04SN0QjL1QDN/bit-overview #-752212c). **Example value**: "U9sGw5wyoiOIqdk1C4mcbYmMnbt" **Data validation rules**: - Length range: `1` ～ `200` characters | ## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `items` | `app.block_workflow[]` | workflow list |
|     `workflow_id` | `string` | workflow unique key |
|     `title` | `string` | workflow title |
|     `status` | `string` | workflow state **Optional values are**:  - `Enable`: enable - `Disable`: disable  | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
        "items": [
            {
                "workflow_id": "12412312421312",
                "title": "工作流",
                "status": "Enable"
            }
        ]
    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 1254000 | WrongRequestJson | The request body format is incorrect. Please check if the request body format complies with the JSON specification |
| 400 | 1254001 | WrongRequestBody | Request body error. Please check that all required parameters have been passed in the request body |
| 400 | 1254002 | Fail | Internal error, please contact [Technical Support] (https://applink.larksuite.com/TLJpeNdW) |
| 400 | 1254003 | WrongBaseToken | App token error, please check if the correct token value was intercepted from the link |
| 400 | 1254004 | WrongTableId | Table id error, please check if the correct table id value was extracted from the link |
| 400 | 1254005 | WrongViewId | View id error, please check if the correct view ID value was extracted from the link |
| 400 | 1254006 | WrongRecordId | Record id error, please check if the correct record ID was obtained from the page |
| 400 | 1254007 | EmptyValue | Empty value, please check the input parameter value |
| 400 | 1254008 | EmptyView | Empty view, please check if the queried view is valid |
| 400 | 1254009 | WrongFieldId | The field ID is incorrect. Please check if the field ID is valid |
| 400 | 1254010 | ReqConvError | Request parameter conversion error |
| 400 | 1254015 | Field types do not match. | Field type mismatch, please check if the field type matches the field type in the base table |
| 403 | 1254027 | UploadAttachNotAllowed | The attachment is not mounted and writing is prohibited. Please upload and mount the attachment first |
| 400 | 1254030 | InvalidPageToken | Page token error, please retrieve a valid next page token from the return body as the requested page token |
| 400 | 1254036 | Base is copying, please try again later. | The Base of the operation is a copy, it is being copied, please try again later. |
| 400 | 1254037 | Invalid client token, make sure that it complies with the specification. | Idempotent key is empty, please check if idempotent key is empty |
| 400 | 1254040 | BaseTokenNotFound | The token does not exist. Please check if the app token passed in is valid |
| 400 | 1254041 | TableIdNotFound | Table id does not exist, please check if the passed table id is valid |
| 400 | 1254042 | ViewIdNotFound | View id does not exist. Please check if the received view ID is valid |
| 400 | 1254043 | RecordIdNotFound | Record ID does not exist. Please check if the passed record ID is valid |
| 400 | 1254044 | FieldIdNotFound | Field id does not exist. Please check if the field ID passed in exists |
| 400 | 1254045 | FieldNameNotFound | The field name does not exist. Please check if the field name passed in is valid |
| 400 | 1254060 | TextFieldConvFail | Multi line text cell format error, please check if the format or value of the multi line text field meets the requirements |
