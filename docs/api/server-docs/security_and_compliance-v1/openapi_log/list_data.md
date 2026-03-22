---
title: "Obtain OpenAPI audit log"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/security_and_compliance-v1/openapi_log/list_data"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/security_and_compliance/v1/openapi_logs/list_data"
service: "security_and_compliance-v1"
resource: "openapi_log"
section: "security_and_compliance"
rate_limit: "100 per minute"
scopes:
  - "security_and_compliance:audit_log.openapi_log:readonly"
updated: "1696011075000"
---

# Obtain OpenAPI audit log

This api is used to obtain the OpenAPI audit log

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/security_and_compliance/v1/openapi_logs/list_data |
| HTTP Method | POST |
| Rate Limit | 100 per minute |
| Supported app types | custom |
| Required scopes | `security_and_compliance:audit_log.openapi_log:readonly` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `api_keys` | `string[]` | No | Lark OpenAPI definition, reference: [API list](https://open.larksuite.com/document/server-docs/getting-started/server-api-list) **Example value**: ["POST/open-apis/authen/v1/access_token"] **Data validation rules**: - Maximum length: `100` |
| `start_time` | `int` | No | Starting timestamp in seconds **Example value**: 1610613336 |
| `end_time` | `int` | No | Termination timestamp in seconds **Example value**: 1610613336 |
| `app_id` | `string` | No | The unique identifier of the application calling OpenAPI, can be obtained by going to [Developer Console](https://open.larksuite.com/app) > Application details page > Certificate and Basic Information **Example value**: "cli_xxx" |
| `page_size` | `int` | No | Paging size **Example value**: 20 **Data validation rules**: - Value range: `1` ～ `100` |
| `page_token` | `string` | No | Page identifier. It is not filled in the first request, indicating traversal from the beginning; when there will be more groups, the new page_token will be returned at the same time, and the next traversal can use the page_token to get more groups **Example value**: "xxx" | ### Request body example

{
    "api_keys": [
        "POST/open-apis/authen/v1/access_token"
    ],
    "start_time": 1610613336,
    "end_time": 1610613336,
    "app_id": "cli_xxx",
    "page_size": 20,
    "page_token": "xxx"
}

## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `items` | `openapi_log[]` | OpenAPI log list |
|     `id` | `string` | OpenAPI log unique identity |
|     `api_key` | `string` | Lark OpenAPI definition |
|     `event_time` | `int` | The time the log was generated, timestamp in seconds |
|     `app_id` | `string` | The unique identifier of the application calling OpenAPI |
|     `ip` | `string` | The ip address that initiates the call to the api |
|     `log_detail` | `openapi_log_detail` | OpenAPI log details |
|       `path` | `string` | HTTP request path |
|       `method` | `string` | HTTP request method |
|       `query_param` | `string` | HTTP query parameters |
|       `payload` | `string` | HTTP request body |
|       `status_code` | `int` | HTTP status code |
|       `response` | `string` | HTTP response body, only return code, msg, error information, etc |
|   `page_token` | `string` | Page identifier, when has_more is true, a new page_token will also be returned. Otherwise, page_token will not be returned |
|   `has_more` | `boolean` | Whether the response body has more parameters | ### Response body example

{"code":0,
"msg":"success",
"data":{"items":[{"id":"111",
"api_key":"POST/open-apis/authen/v1/access_token",
"event_time":1610613336,
"app_id":"cli_xxx",
"ip":"192.123.12.1或fdbd:ff1:ce00:135:c7e:d128:5028:6546",
"log_detail":{"path":"/open-apis/auth/v3/app_access_token",
"method":"POST",
"query_param":"{}",
"payload":"{"app_id": "cli_xxx", "app_secret": "xxx", "app_ticket": "xxx"}",
"status_code":0,
"response":"{"code": 0, "msg": "ok"}"}}],
"page_token":"xxxx",
"has_more":false}}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 500 | 1780001 | Internal Server Error | Please try again. If there is still any abnormality, please contact the open platform technical support |
| 400 | 1780103 | Request Param Error | Check input parameters with documentation |
