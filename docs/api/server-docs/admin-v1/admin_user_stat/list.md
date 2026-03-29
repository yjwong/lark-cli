---
title: "Obtain active users and function usage data in the user dimension"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/admin-v1/admin_user_stat/list"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/admin/v1/admin_user_stats"
service: "admin-v1"
resource: "admin_user_stat"
section: "Admin"
rate_limit: "Special Rate Limit"
scopes:
  - "admin:admin_user_stat:readonly"
field_scopes:
  - "contact:user.employee_id:readonly"
updated: "1718591665000"
---

# Obtain active users and function usage data in the user dimension

This API is used to obtain active users and function usage data in the user dimension. That is, the usage data of IM (Instant messaging), Calendar, Docs, Audio/Video conferencing and Email functions.

> - Only custom app has the permission scope to call this API.
> 
> - The data of the day will be generated at 9:30 am the next day (UTC+0).

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/admin/v1/admin_user_stats |
| HTTP Method | GET |
| Rate Limit | Special Rate Limit |
| Supported app types | custom |
| Required scopes | `admin:admin_user_stat:readonly` |
| Required field scopes | > The response body of the API contains the following sensitive fields, and they will be returned only after corresponding scopes are added. If you do not need the fields, it is not recommended that you request the scopes. `contact:user.employee_id:readonly` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `user_id_type` | `string` | No | User ID categories **Example value**: open_id **Optional values are**:  - `open_id`: Identifies a user to an app. The same user has different Open IDs in different apps. How to get Open ID - `union_id`: Identifies a user to a tenant that acts as a developer. A user has the same Union ID in apps developed by the same developer, and has different Union IDs in apps developed by different developers. A developer can use Union ID to link the same user's identities in multiple apps.How to get Union ID - `user_id`: Identifies a user to a tenant. The same user has different User IDs in different tenants. In one single tenant, a user has the same User ID in all apps （including store apps）. User ID is usually used to communicate user data between different apps. How to get User ID  **Default value**: `open_id` **When the value is `user_id`, the following field scopes are required**: `contact:user.employee_id:readonly` |
| `department_id_type` | `string` | No | Department ID type **Example value**: open_department_id **Optional values are**:  - `department_id`: Department ID - `open_department_id`: Department Open ID  |
| `start_date` | `string` | Yes | Start date (inclusive) in the format of YYYY-mm-dd **Example value**: 2020-02-15 |
| `end_date` | `string` | Yes | End date (inclusive) in YYYY-mm-dd. The difference between start date and end date shall not exceed 31 days (inclusive) **Example value**: 2020-02-15 |
| `department_id` | `string` | No | Department ID, depending on department_id_type **Example value**: od-382e2793cfc9471f892e8a672987654c |
| `user_id` | `string` | No | User's open_id, user_id, or union_id, depending on user_id_type **Example value**: ou_7dab8a3d3cdcc9da365777c7ad535d62 |
| `page_size` | `int` | No | Page size **Example value**: 10 **Data validation rules**: - Value range: `1` ～ `60` |
| `page_token` | `string` | No | Page identifier. It is not filled in the first request, indicating traversal from the beginning; when there will be more groups, the new page_token will be returned at the same time, and the next traversal can use the page_token to get more groups **Example value**: 2 |
| `target_geo` | `string` | No | Geo data that needs to be accessed across domains. Each Geo only contains this Geo data. Local data will be checked by default if not passed. You need to activate FG (cn, us, sg, jp) before calling. Only one Geo data can be checked at a time. **Example value**: cn | ## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `has_more` | `boolean` | Whether the response body has more parameters |
|   `page_token` | `string` | Page identifier, when has_more is true, a new page_token will also be returned. Otherwise, page_token will not be returned |
|   `items` | `admin_user_stat[]` | Report |
|     `date` | `string` | Date |
|     `user_id` | `string` | User ID |
|     `user_name` | `string` | User name |
|     `department_name` | `string` | Department name |
|     `department_path` | `string` | Department path |
|     `create_time` | `string` | Account creation time |
|     `user_active_flag` | `int` | User activation status **Optional values are**:  - `0`: Unactivated - `1`: Activated  |
|     `register_time` | `string` | Activation time |
|     `suite_active_flag` | `int` | Active status of user .Users that sign in to instant messaging, doc, calendar, meeting, open API platform, etc are considered active in Lark Suite **Optional values are**:  - `0`: Inactive - `1`: Active  |
|     `last_active_time` | `string` | Last active time |
|     `im_active_flag` | `int` | Active status of user message. Users that sent or view a message are considered active in Lark Messager. **Optional values are**:  - `0`: Inactive - `1`: Active  |
|     `send_messenger_num` | `int` | Number of messages sent |
|     `docs_active_flag` | `int` | Active status of user Docs. Users that open a doc or space are considered active in Lark Docs **Optional values are**:  - `0`: Inactive - `1`: Active  |
|     `create_docs_num` | `int` | Number of created files |
|     `cal_active_flag` | `int` | Active status of user calendar. Uses that open,create or receive a calendar are considered active in Lark Calendar **Optional values are**:  - `0`: Inactive - `1`: Active  |
|     `create_cal_num` | `int` | Number of created events |
|     `vc_active_flag` | `int` | Active status of user audio/video conferencing. Users that join an audio/video meeting are considered active in Lark Meetings **Optional values are**:  - `0`: Inactive - `1`: Active  |
|     `vc_duration` | `int` | Meeting duration （mins， meeting room's duration not included） |
|     `active_os` | `string` | Active device |
|     `create_task_num` | `int` | Number of tasks created |
|     `vc_num` | `int` | Number of video meetings joined |
|     `app_package_type` | `string` | Lark's application type name |
|     `os_name` | `string` | Operating system name |
|     `email_send_count` | `string` | email send total count |
|     `email_receive_count` | `string` | email receive total count |
|     `email_send_ext_count` | `string` | email send ext count |
|     `email_receive_ext_count` | `string` | email receive ext count |
|     `email_send_in_count` | `string` | email send in count |
|     `email_receive_in_count` | `string` | email receive in count |
|     `search_active_flag` | `int` | Whether to use the big search (0: not used, 1: used) |
|     `total_search_count` | `string` | Total number of searches (the number of sessions that initiated a search request in the Lark main search box) |
|     `quick_search_count` | `string` | Top search count (number of sessions where a search request was initiated by a comprehensive search in the Lark master search box) |
|     `tab_search_count` | `string` | Number of vertical searches (the number of sessions where a search request was initiated in the vertical search tab (e.g. Messages tab, Cloud Documents tab) of the Lark main search box) | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
        "has_more": true,
        "page_token": "3",
        "items": [
            {
                "date": "2020-02-15",
                "user_id": "ou_7dab8a3d3cdcc9da365777c7ad535d62",
                "user_name": "Zhang San",
                "department_name": "testcqlbfaaasdasdasd",
                "department_path": "testkkk/testcqlbfaaasdasdasd",
                "create_time": "2020-09-04 11:17:55",
                "user_active_flag": 1,
                "register_time": "2020-09-04 11:18:32",
                "suite_active_flag": 1,
                "last_active_time": "2020-12-21 22:21:28",
                "im_active_flag": 1,
                "send_messenger_num": 0,
                "docs_active_flag": 1,
                "create_docs_num": 1,
                "cal_active_flag": 1,
                "create_cal_num": 0,
                "vc_active_flag": 1,
                "vc_duration": 0,
                "active_os": "'ios 14.2,-','ios 14.2,lark 3.40.0-alpha'",
                "create_task_num": 0,
                "vc_num": 0,
                "app_package_type": "Lark，Lark",
                "os_name": "iOS,Andorid,Windows",
                "email_send_count": "2",
                "email_receive_count": "3",
                "email_send_ext_count": "4",
                "email_receive_ext_count": "5",
                "email_send_in_count": "6",
                "email_receive_in_count": "7",
                "search_active_flag": 1,
                "total_search_count": "7",
                "quick_search_count": "7",
                "tab_search_count": "7"
            }
        ]
    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 500 | 1051001 | request contain invalid param | Request contain invalid param |
| 400 | 1051002 | request to exceed authority | Request to exceed authority |
