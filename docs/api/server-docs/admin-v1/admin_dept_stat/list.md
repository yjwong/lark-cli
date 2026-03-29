---
title: "Obtain active users and function usage data in the department dimension"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/admin-v1/admin_dept_stat/list"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/admin/v1/admin_dept_stats"
service: "admin-v1"
resource: "admin_dept_stat"
section: "Admin"
rate_limit: "Special Rate Limit"
scopes:
  - "admin:admin_dept_stat:readonly"
updated: "1718591664000"
---

# Obtain active users and function usage data in the department dimension

The API is used to obtain active users and function usage data in the department dimension, including the usage data about IM (Instant Messaging), Calendar, Docs, Audio/Video and Email Conferencing.

> - Only custom app has the permission scope to call this API.
> 
> - The data of the day will be generated at 9:30 AM the next day (UTC+0).

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/admin/v1/admin_dept_stats |
| HTTP Method | GET |
| Rate Limit | Special Rate Limit |
| Supported app types | custom |
| Required scopes | `admin:admin_dept_stat:readonly` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `department_id_type` | `string` | Yes | Department ID type **Example value**: open_department_id **Optional values are**:  - `department_id`: Department ID - `open_department_id`: Department Open ID  |
| `start_date` | `string` | Yes | Start date (inclusive) in the format of YYYY-mm-dd **Example value**: 2020-02-15 |
| `end_date` | `string` | Yes | End date (inclusive) in YYYY-mm-dd. The difference between start date and end date shall not exceed 91 days (inclusive) **Example value**: 2020-02-15 |
| `department_id` | `string` | Yes | Department ID, depending on department_id_type and supporting only the root department and its top four sub-departments **Example value**: od-382e2793cfc9471f892e8a672987654c |
| `contains_child_dept` | `boolean` | Yes | Whether sub-departments are included. If the value is false, only the active users and function usage data directly under the department are searched. If the value is true, the active users and function usage data of the department and its sub-departments (the level of sub-department does not exceed the top four levels under the root department) are searched **Example value**: false |
| `page_size` | `int` | No | Page size **Example value**: 10 **Data validation rules**: - Value range: `1` ～ `20` |
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
|   `items` | `admin_dept_stat[]` | Reports |
|     `date` | `string` | Date |
|     `department_id` | `string` | Department department_id or open_department_id |
|     `department_name` | `string` | Department name |
|     `department_path` | `string` | Department path |
|     `total_user_num` | `int` | Total number of department members |
|     `active_user_num` | `int` | Number of activated users |
|     `active_user_rate` | `string` | Activation rate |
|     `suite_dau` | `int` | Number of active users |
|     `suite_active_rate` | `string` | Active rate |
|     `new_user_num` | `int` | Number of new users |
|     `new_active_num` | `int` | Number of new activated users |
|     `resign_user_num` | `int` | Number of terminated staff |
|     `im_dau` | `int` | Number of active message users |
|     `send_messenger_user_num` | `int` | Number of users sending messages |
|     `send_messenger_num` | `int` | Number of messages sent |
|     `avg_send_messenger_num` | `string` | Number of messages sent per person |
|     `docs_dau` | `int` | Number of active Docs users |
|     `create_docs_user_num` | `int` | Number of users creating files |
|     `create_docs_num` | `int` | Number of created files |
|     `avg_create_docs_num` | `string` | Number of created files per person |
|     `cal_dau` | `int` | Number of active Calendar users |
|     `create_cal_user_num` | `int` | Number of users creating events |
|     `create_cal_num` | `int` | Number of created events |
|     `avg_create_cal_num` | `string` | Number of created events per person |
|     `vc_dau` | `int` | Number of active audio/video conferencing users |
|     `vc_duration` | `int` | Meeting duration of all users in the company (minutes,meeting room's duration not included) |
|     `avg_vc_duration` | `string` | Average meeting duration per person (minutes,meeting room's duration not included) |
|     `avg_duration` | `string` | Lark average usage time (min) |
|     `task_dau` | `int` | Number of users using Tasks |
|     `create_task_user_num` | `int` | Number of users who created new tasks |
|     `create_task_num` | `int` | Number of tasks created |
|     `avg_create_task_num` | `string` | Average tasks created per user |
|     `email_send_count` | `string` | Total Mail Shipments |
|     `email_receive_count` | `string` | Total mail receipts |
|     `email_send_ext_count` | `string` | Number of outgoing shipments |
|     `email_receive_ext_count` | `string` | Number of receipts from outside |
|     `email_send_in_count` | `string` | Number of internal senders |
|     `email_receive_in_count` | `string` | Number of receipts from within |
|     `search_active_dau` | `string` | Dasou search active number |
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
                "department_id": "od-382e2793cfc9471f892e8a672987654c",
                "department_name": "subtestkkk",
                "department_path": "testkkk/subtestkkk",
                "total_user_num": 2,
                "active_user_num": 0,
                "active_user_rate": "1.00",
                "suite_dau": 0,
                "suite_active_rate": "0.00",
                "new_user_num": 0,
                "new_active_num": 0,
                "resign_user_num": 0,
                "im_dau": 0,
                "send_messenger_user_num": 0,
                "send_messenger_num": 0,
                "avg_send_messenger_num": "0.00",
                "docs_dau": 0,
                "create_docs_user_num": 0,
                "create_docs_num": 0,
                "avg_create_docs_num": "0.00",
                "cal_dau": 0,
                "create_cal_user_num": 0,
                "create_cal_num": 0,
                "avg_create_cal_num": "0.00",
                "vc_dau": 0,
                "vc_duration": 0,
                "avg_vc_duration": "0.00",
                "avg_duration": "0.00",
                "task_dau": 0,
                "create_task_user_num": 0,
                "create_task_num": 0,
                "avg_create_task_num": "0.00",
                "email_send_count": "2",
                "email_receive_count": "3",
                "email_send_ext_count": "4",
                "email_receive_ext_count": "5",
                "email_send_in_count": "6",
                "email_receive_in_count": "7",
                "search_active_dau": "7",
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
