---
title: "Query export task results"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/vc-v1/export/get"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/vc/v1/exports/:task_id"
service: "vc-v1"
resource: "export"
section: "Video Conferencing"
rate_limit: "100 per minute"
scopes:
  - "vc:export"
updated: "1689326178000"
---

# Query export task results

View the progress of asynchronous export.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/vc/v1/exports/:task_id |
| HTTP Method | GET |
| Rate Limit | 100 per minute |
| Supported app types | custom,isv |
| Required scopes | `vc:export` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `task_id` | `string` | Task id **Example value**: "7108646852144136212" | ## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `status` | `int` | Task status **Optional values are**:  - `1`: Processing - `2`: Fail - `3`: Done  |
|   `url` | `string` | File download address |
|   `file_token` | `string` | file token |
|   `fail_msg` | `string` | fail message | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
        "status": 3,
        "url": "https://lf1-ttcdn-tos.pstatp.com/obj/xxx",
        "file_token": "6yHu7Igp7Igy62Ez6fLr6IJz7j9i5WMe6fHq5yZeY2Jz6yLqYAMAY46fZfEz64Lr5fYyYQ==",
        "fail_msg": "no permission"
    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 500 | 121001 | internal error | Internal server error. If retry still doesn't work, contact the administrator. |
| 400 | 121002 | not support | This function is not supported. |
| 400 | 121003 | param error | Parameter error. Check the value range of parameters (please check yourself according to the field description above). |
| 404 | 121004 | data not exist | The requested data doesn't exist. |
| 403 | 121005 | no permission | No permission. Check the token type, operator identity and resource. |
