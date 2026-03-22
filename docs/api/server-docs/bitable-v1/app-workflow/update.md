---
title: "Update Workflow status"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/bitable-v1/app-workflow/update"
method: "PUT"
api_path: "https://open.larksuite.com/open-apis/bitable/v1/apps/:app_token/workflows/:workflow_id"
service: "bitable-v1"
resource: "app-workflow"
section: "Docs"
rate_limit: "10 per second"
scopes:
  - "bitable:app"
updated: "1770864569000"
---

# Update Workflow status

This interface can be used to update Workflow status 

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/bitable/v1/apps/:app_token/workflows/:workflow_id |
| HTTP Method | PUT |
| Rate Limit | 10 per second |
| Supported app types | custom,isv |
| Required scopes | `bitable:app` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `app_token` | `string` | Bitable app token **Example value**: "appbcbWCzen6D8dezh" |
| `workflow_id` | `string` | workflow_id **Example value**: "730887xxxx552638996" | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `status` | `string` | Yes | automated state **Example value**: "Enable" | ### Request body example

{
    "status": "Enable"
}

The status of the target automation, currently only "Enable" and "Disable" are available

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
| 200 | 1254000 | WrongRequestJson | request body error |
| 200 | 1254001 | WrongRequestBody | request body error |
| 200 | 1254002 | Fail | Internal error, please contact technical support |
| 200 | 1254003 | WrongBaseToken | app_token mistake |
| 200 | 1254010 | ReqConvError | request error |
| 400 | 1254036 | Bitable is copying, please try again later. | Bitable copy replicating, try again later |
| 200 | 1254040 | BaseTokenNotFound | app_token doesn't exist |
| 200 | 1254290 | TooManyRequest | Request is too fast, try again later |
| 403 | 1254302 | Permission denied. | No access rights, usually caused by the table opening of advanced permissions, please add a group containing applications in the advanced permissions settings and give this group read and write permissions |
| 200 | 1255001 | InternalError | Internal error, please contact technical support |
| 200 | 1255002 | RpcError | Internal error, please contact technical support |
| 200 | 1255003 | MarshalError | Serialization error, please contact technical support |
| 200 | 1255004 | UmMarshalError | deserialization error |
| 504 | 1255040 | Request timed out, please try again later | Try again. |
