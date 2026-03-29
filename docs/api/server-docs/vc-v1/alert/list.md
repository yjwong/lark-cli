---
title: "Get alert center history"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/vc-v1/alert/list"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/vc/v1/alerts"
service: "vc-v1"
resource: "alert"
section: "Video Conferencing"
rate_limit: "100 per minute"
scopes:
  - "vc:alert:readonly"
updated: "1689326174000"
---

# Get alert center history

Get tenant's equipment warning record under specific conditions.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/vc/v1/alerts |
| HTTP Method | GET |
| Rate Limit | 100 per minute |
| Supported app types | custom |
| Required scopes | `vc:alert:readonly` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `start_time` | `string` | Yes | Start time (UNIX time, unit SEC) **Example value**: 1608888867 |
| `end_time` | `string` | Yes | End time (UNIX time, unit SEC) **Example value**: 1608888867 |
| `query_type` | `int` | No | Query object type **Example value**: 1 **Optional values are**:  - `1`: Rooms - `2`: Erc  |
| `query_value` | `string` | No | Query object ID **Example value**: 6911188411932033028 |
| `page_size` | `int` | No | **Example value**: 100 **Data validation rules**: - Maximum value: `200` |
| `page_token` | `string` | No | Page identifier. It is not filled in the first request, indicating traversal from the beginning; when there will be more groups, the new page_token will be returned at the same time, and the next traversal can use the page_token to get more groups **Example value**: 0 | ## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `has_more` | `boolean` | Whether the response body has more parameters |
|   `page_token` | `string` | Page identifier, when has_more is true, a new page_token will also be returned. Otherwise, page_token will not be returned |
|   `items` | `alert[]` | Alarm record |
|     `alert_id` | `string` | Alarm ID |
|     `resource_scope` | `string` | The specific name of the conference room/server that trigger the alarm rules |
|     `monitor_target` | `int` | Monitoring object to trigger alarm rules **Optional values are**:  - `1`: Rooms - `2`: Checkboard - `3`: Roombox - `4`: Room tv share - `5`: Sip - `6`: Erc - `7`: Room sensor  |
|     `alert_strategy` | `string` | Description of the rules of the alarm rules |
|     `alert_time` | `string` | Alarm notification time (UNIX time, unit SEC) |
|     `alert_level` | `int` | Alarm level: serious/warning/reminder **Optional values are**:  - `0`: Reminder - `1`: Warning - `2`: Serious  |
|     `contacts` | `contact[]` | Alarm contact |
|       `contact_type` | `int` | Alarm contact type **Optional values are**:  - `1`: User - `2`: User group - `3`: Department  |
|       `contact_name` | `string` | Contact name |
|     `notifyMethods` | `int[]` | notify methods **Optional values are**:  - `0`: bot - `1`: email  |
|     `alertRule` | `string` | Alert rule |
|     `process_time` | `string` | Alarm process time (UNIX time, unit SEC) |
|     `recover_time` | `string` | Alarm recover time (UNIX time, unit SEC) |
|     `process_status` | `int` | Alarm process status **Optional values are**:  - `0`: Wait process - `1`: Wait process - `2`: Processing - `3`: Recover - `4`: Recover  | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
        "has_more": true,
        "page_token": "50",
        "items": [
            {
                "alert_id": "7115030004018184212",
                "resource_scope": "xxx level",
                "monitor_target": 2,
                "alert_strategy": "For 1 continuous cycle (1 minute in total), if the power of the controller is 

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 500 | 121001 | internal error | Internal server error. If retry still doesn't work, contact the administrator. |
| 400 | 121002 | not support | This function is not supported. |
| 400 | 121003 | param error | Parameter error. Check the value range of parameters (please check yourself according to the field description above). |
| 404 | 121004 | data not exist | The requested data doesn't exist. |
