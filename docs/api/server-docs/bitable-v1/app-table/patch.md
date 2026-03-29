---
title: "Update data table"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/bitable-v1/app-table/patch"
method: "PATCH"
api_path: "https://open.larksuite.com/open-apis/bitable/v1/apps/:app_token/tables/:table_id"
service: "bitable-v1"
resource: "app-table"
section: "Docs"
rate_limit: "10 per second"
scopes:
  - "base:table:update"
  - "bitable:app"
updated: "1773323337000"
---

# Update data table

This interface is used to update the basic information of the data table, including the name of the data table, etc.

::: note
For the first access, please refer to Cloud Document Interface QuickStart & Base OpenAPI Access Guide 

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/bitable/v1/apps/:app_token/tables/:table_id |
| HTTP Method | PATCH |
| Rate Limit | 10 per second |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `base:table:update` `bitable:app` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `app_token` | `string` | Base unique device identifier app_token description **Example value**: "XrgTb4y1haKYnasu0xXb1g7lcSg" **Data validation rules**: - Minimum length: `1` characters |
| `table_id` | `string` | Base data table unique device identifier table_id description **Example value**: "tbl1TkhyTWDkSoZ3" | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `name` | `string` | No | The new name for the data table. Please note: 1. The first and last spaces in the name will be removed. 2. If the name is empty or the same as the old name, the interface will still return success, but the name will not be changed. **Example value**: "new name" **Data validation rules**: - Length range: `1` ～ `100` characters - Regular expression: `^[^[]\:\\\/\?\*]+$` | ### Request body example

{
    "name": "new name"
}

## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `name` | `string` | The name of the data table | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
        "name": "new name"
    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 1254001 | WrongRequestBody | Request body error |
| 400 | 1254002 | Fail | Internal error, have any questions can be consulting service |
| 400 | 1254003 | WrongBaseToken | AppToken error |
| 400 | 1254004 | WrongTableId | Table id wrong |
| 400 | 1254013 | TableNameDuplicated | TableNameDuplicated |
| 403 | 1254036 | Base is copying, please try again later. | Base copy replicating, try again later |
| 404 | 1254040 | BaseTokenNotFound | AppToken not found |
| 404 | 1254041 | TableIdNotFound | Table not found |
| 429 | 1254290 | TooManyRequest | Request too fast, try again later |
| 403 | 1254302 | Permission denied. | No access rights, usually caused by the table opening of advanced permissions, please add a group containing applications in the advanced permissions settings and give this group read and write permissions |
| 500 | 1255001 | InternalError | Internal error, have any questions can be consulting service |
