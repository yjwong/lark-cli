---
title: "Get top user list"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/vc-v1/report/get_top_user"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/vc/v1/reports/get_top_user"
service: "vc-v1"
resource: "report"
section: "Video Conferencing"
rate_limit: "Special Rate Limit"
scopes:
  - "vc:report:readonly"
field_scopes:
  - "contact:user.employee_id:readonly"
updated: "1690944921000"
---

# Get top user list

Obtains the list of top users who use the Meetings feature most in an organization for a period of time.

> Data for the last 90 days can be queried. The top 10 users are returned by default. A maximum of top 100 users can be queried.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/vc/v1/reports/get_top_user |
| HTTP Method | GET |
| Rate Limit | Special Rate Limit |
| Supported app types | custom,isv |
| Required scopes | `vc:report:readonly` |
| Required field scopes | > The response body of the API contains the following sensitive fields, and they will be returned only after corresponding scopes are added. If you do not need the fields, it is not recommended that you request the scopes. `contact:user.employee_id:readonly` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `start_time` | `string` | Yes | Start time (unix time in sec) **Example value**: 1608888867 |
| `end_time` | `string` | Yes | End time (unix time in sec) **Example value**: 1608889966 |
| `limit` | `int` | Yes | Number of users to be obtained **Example value**: 10 |
| `order_by` | `int` | Yes | Sorting criteria (descending) **Example value**: 1 **Optional values are**:  - `1`: Number of meetings - `2`: Meeting duration  |
| `user_id_type` | `string` | No | User ID categories **Example value**: open_id **Optional values are**:  - `open_id`: Identifies a user to an app. The same user has different Open IDs in different apps. How to get Open ID - `union_id`: Identifies a user to a tenant that acts as a developer. A user has the same Union ID in apps developed by the same developer, and has different Union IDs in apps developed by different developers. A developer can use Union ID to link the same user's identities in multiple apps.How to get Union ID - `user_id`: Identifies a user to a tenant. The same user has different User IDs in different tenants. In one single tenant, a user has the same User ID in all apps （including store apps）. User ID is usually used to communicate user data between different apps. How to get User ID  **Default value**: `open_id` **When the value is `user_id`, the following field scopes are required**: `contact:user.employee_id:readonly` | ## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `top_user_report` | `report_top_user[]` | Top user list |
|     `id` | `string` | User ID |
|     `name` | `string` | User name |
|     `user_type` | `int` | User type **Optional values are**:  - `1`: Lark user - `2`: Rooms user - `3`: Docs user - `4`: neo Lark Meetings user - `5`: neo Lark Meetings guest - `6`: PSTN user - `7`: SIP user  |
|     `meeting_count` | `string` | Number of meetings |
|     `meeting_duration` | `string` | Meeting duration (in min) | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
        "top_user_report": [
            {
                "id": "ou_3ec3f6a28a0d08c45d895276e8e5e19b",
                "name": "name",
                "user_type": 1,
                "meeting_count": "100",
                "meeting_duration": "3000"
            }
        ]
    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 500 | 121001 | internal error | Internal server error. If retry still doesn't work, contact the administrator. |
| 400 | 121002 | not support | This function is not supported. |
| 400 | 121003 | param error | Parameter error. Check the value range of parameters (please check yourself according to the field description above). |
