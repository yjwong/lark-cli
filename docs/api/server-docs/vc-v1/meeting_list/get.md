---
title: "Get meeting details"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/vc-v1/meeting_list/get"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/vc/v1/meeting_list"
service: "vc-v1"
resource: "meeting_list"
section: "Video Conferencing"
rate_limit: "100 per minute"
scopes:
  - "vc:room:readonly"
field_scopes:
  - "contact:user.employee_id:readonly"
updated: "1689326174000"
---

# Get meeting details

Get meeting details.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/vc/v1/meeting_list |
| HTTP Method | GET |
| Rate Limit | 100 per minute |
| Supported app types | custom,isv |
| Required scopes | `vc:room:readonly` |
| Required field scopes | > The response body of the API contains the following sensitive fields, and they will be returned only after corresponding scopes are added. If you do not need the fields, it is not recommended that you request the scopes. `contact:user.employee_id:readonly` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `start_time` | `string` | Yes | Start time (unix time, in sec) **Example value**: 1655276858 |
| `end_time` | `string` | Yes | End time (unix time, in sec) **Example value**: 1655276858 |
| `meeting_no` | `string` | No | Filter by 9-digit meeting number (at most one filter) **Example value**: 123456789 |
| `user_id` | `string` | No | Filter by participating Lark users (at most one filter) **Example value**: ou_3ec3f6a28a0d08c45d895276e8e5e19b |
| `room_id` | `string` | No | Filter by Participating Rooms (at most one filter) **Example value**: omm_eada1d61a550955240c28757e7dec3af |
| `page_size` | `int` | No | **Example value**: 20 **Default value**: `20` **Data validation rules**: - Value range: `20` ～ `100` |
| `page_token` | `string` | No | Page identifier. It is not filled in the first request, indicating traversal from the beginning; when there will be more groups, the new page_token will be returned at the same time, and the next traversal can use the page_token to get more groups **Example value**: 20 |
| `user_id_type` | `string` | No | User ID categories **Example value**: open_id **Optional values are**:  - `open_id`: Identifies a user to an app. The same user has different Open IDs in different apps. How to get Open ID - `union_id`: Identifies a user to a tenant that acts as a developer. A user has the same Union ID in apps developed by the same developer, and has different Union IDs in apps developed by different developers. A developer can use Union ID to link the same user's identities in multiple apps.How to get Union ID - `user_id`: Identifies a user to a tenant. The same user has different User IDs in different tenants. In one single tenant, a user has the same User ID in all apps （including store apps）. User ID is usually used to communicate user data between different apps. How to get User ID  **Default value**: `open_id` **When the value is `user_id`, the following field scopes are required**: `contact:user.employee_id:readonly` | ## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `meeting_list` | `meeting_info[]` | Meeting list |
|     `meeting_id` | `string` | 9 digits meeting Number |
|     `meeting_topic` | `string` | Conference theme |
|     `organizer` | `string` | Organizer |
|     `department` | `string` | Department |
|     `user_id` | `string` | User ID |
|     `employee_id` | `string` | Employee number |
|     `email` | `string` | E-mail |
|     `mobile` | `string` | Phone Number |
|     `meeting_start_time` | `string` | Meeting start time |
|     `meeting_end_time` | `string` | Meeting end time |
|     `meeting_duration` | `string` | Duration of the meeting |
|     `number_of_participants` | `string` | Number of participants |
|     `audio` | `boolean` | Audio |
|     `video` | `boolean` | Video |
|     `sharing` | `boolean` | Share |
|     `recording` | `boolean` | Recording |
|     `telephone` | `boolean` | Telephone |
|   `page_token` | `string` | Page identifier, when has_more is true, a new page_token will also be returned. Otherwise, page_token will not be returned |
|   `has_more` | `boolean` | Whether the response body has more parameters | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
        "meeting_list": [
            {
                "meeting_id": "705605196",
                "meeting_topic": "seminar",
                "organizer": "kehan",
                "department": "development",
                "user_id": "92f879",
                "employee_id": "202105149765",
                "email": "xxxx@163.com",
                "mobile": "021-673288",
                "meeting_start_time": "2022.12.23 11:16:59 (GMT+08:00)",
                "meeting_end_time": "2022.12.23 11:18:51 (GMT+08:00)",
                "meeting_duration": "00:01:52",
                "number_of_participants": "1",
                "audio": true,
                "video": true,
                "sharing": false,
                "recording": false,
                "telephone": false
            }
        ],
        "page_token": "20",
        "has_more": true
    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 500 | 121001 | internal error | Internal server error. If retry still doesn't work, contact the administrator. |
| 400 | 121002 | not support | This function is not supported. |
| 400 | 121003 | param error | Parameter error. Check the value range of parameters (please check yourself according to the field description above). |
| 400 | 126003 | meeting number invalid (cannot smaller than 0) | Invalid meeting number (cannot be less than 0) |
| 400 | 126004 | meeting number length not correct (supposed to be 9) | The meeting number should be 9 digits |
| 400 | 126005 | start time error (cannot larger than now or smaller than 0) | The start time cannot be greater than the current time or less than 0 |
| 400 | 126006 | end time cannot exceed start time for one year or smaller than start time | The end time cannot exceed the start time by one year or be less than the start time |
| 400 | 126008 | userID and roomID and meeting number should be set only one | userID, roomID and meeting number can only be set one at the same time |
| 404 | 121004 | data not exist | The requested data doesn't exist. |
| 403 | 121005 | no permission | No permission. Check the token type, operator identity and resource. |
