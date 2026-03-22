---
title: "List automations"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/bitable-v1/app-workflow/list"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/bitable/v1/apps/:app_token/workflows"
service: "bitable-v1"
resource: "app-workflow"
section: "Docs"
rate_limit: "20 per second"
scopes:
  - "bitable:app"
  - "bitable:app:readonly"
updated: "1770864564000"
---

# List automations

This interface is used to list the automations of base

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/bitable/v1/apps/:app_token/workflows |
| HTTP Method | GET |
| Rate Limit | 20 per second |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `bitable:app` `bitable:app:readonly` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `app_token` | `string` | Base app token **Example value**: "appbcbWCzen6D8dezhoCH2RpMAh" | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `page_token` | `string` | No | Page identifier. It is not filled in the first request, indicating traversal from the beginning; when there will be more groups, the new page_token will be returned at the same time, and the next traversal can use the page_token to get more groups **Example value**: eVQrYzJBNDNONlk4VFZBZVlSdzlKdFJ4bVVHVExENDNKVHoxaVdiVnViQT0= |
| `page_size` | `int` | No | paging size **Example value**: 10 **Default value**: `20` | ## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `workflows` | `app.workflow[]` | Automation information |
|     `workflow_id` | `string` | Automation id |
|     `status` | `string` | the status of automation |
|     `title` | `string` | the name of automation | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
        "workflows": [
            {
                "workflow_id": "72934597xxxx9998484",
                "status": "Enable",
                "title": "automation"
            }
        ]
    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 1254000 | WrongRequestJson | request body error |
| 400 | 1254001 | WrongRequestBody | request body error |
| 400 | 1254002 | Fail | Internal error, if you have any questions, please consult customer service. |
| 400 | 1254003 | WrongBaseToken | app_token mistake |
| 400 | 1254004 | WrongTableId | table_id mistake |
| 400 | 1254010 | ReqConvError | request error |
| 404 | 1254040 | BaseTokenNotFound | app_token doesn't exist |
| 404 | 1254041 | TableIdNotFound | table_id doesn't exist |
| 429 | 1254290 | TooManyRequest | Request is too fast, try again later |
| 400 | 1254301 | OperationTypeError | Bitable does not enable advanced permissions or does not support enabling advanced permissions |
| 400 | 1255001 | InternalError | Internal error, please contact [Technical Support] (https://applink.larksuite.com/TLJpeNdW) |
| 400 | 1255002 | RpcError | Internal error, please contact [Technical Support] (https://applink.larksuite.com/TLJpeNdW) |
| 400 | 1255003 | MarshalError | Serialization error, please contact [Technical Support] (https://applink.larksuite.com/TLJpeNdW) Deserialization error |
| 400 | 1255004 | UmMarshalError | deserialization error |
| 400 | 1255005 | ConvError | Internal error, please contact [Technical Support] (https://applink.larksuite.com/TLJpeNdW) |
| 504 | 1255040 | Request timed out, please try again later | request timed out |
| 400 | 1254036 | Bitable is copying, please try again later. | Copy Bitable is an asynchronous operation. This error code indicates that the current Bitable is still being copied, and the current Bitable cannot be operated during the copy. You need to wait for the copy to complete before operating. |
| 500 | 1254200 | Something went wrong. Please contact technical support at https://applink.larksuite.com/client/helpdesk/open?id=6626260912531570952 | Internal error, please contact [Technical Support] (https://applink.larksuite.com/TLJpeNdW) |
