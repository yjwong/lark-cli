---
title: "Get participant details"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/vc-v1/participant_list/get"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/vc/v1/participant_list"
service: "vc-v1"
resource: "participant_list"
section: "Video Conferencing"
rate_limit: "100 per minute"
scopes:
  - "vc:room:readonly"
field_scopes:
  - "contact:user.employee_id:readonly"
updated: "1689326174000"
---

# Get participant details

Get participant details.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/vc/v1/participant_list |
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
| `meeting_start_time` | `string` | Yes | Meeting start time(need to be accurate to the minute,unix time,unit second) **Example value**: 1655276858 |
| `meeting_end_time` | `string` | Yes | Meeting end time **Example value**: 1655276858 |
| `meeting_no` | `string` | Yes | Meeting number **Example value**: 123456789 |
| `user_id` | `string` | No | User ID **Example value**: ou_3ec3f6a28a0d08c45d895276e8e5e19b |
| `room_id` | `string` | No | Room ID **Example value**: omm_eada1d61a550955240c28757e7dec3af |
| `page_size` | `int` | No | **Example value**: 20 **Default value**: `20` **Data validation rules**: - Value range: `20` ～ `100` |
| `page_token` | `string` | No | Page identifier. It is not filled in the first request, indicating traversal from the beginning; when there will be more groups, the new page_token will be returned at the same time, and the next traversal can use the page_token to get more groups **Example value**: 20 |
| `user_id_type` | `string` | No | User ID categories **Example value**: open_id **Optional values are**:  - `open_id`: Identifies a user to an app. The same user has different Open IDs in different apps. How to get Open ID - `union_id`: Identifies a user to a tenant that acts as a developer. A user has the same Union ID in apps developed by the same developer, and has different Union IDs in apps developed by different developers. A developer can use Union ID to link the same user's identities in multiple apps.How to get Union ID - `user_id`: Identifies a user to a tenant. The same user has different User IDs in different tenants. In one single tenant, a user has the same User ID in all apps （including store apps）. User ID is usually used to communicate user data between different apps. How to get User ID  **Default value**: `open_id` **When the value is `user_id`, the following field scopes are required**: `contact:user.employee_id:readonly` | ## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `participants` | `participant[]` | List of participants |
|     `participant_name` | `string` | Attendee |
|     `department` | `string` | Department |
|     `user_id` | `string` | User ID |
|     `employee_id` | `string` | Employee number |
|     `phone` | `string` | Telephone |
|     `email` | `string` | E-mail |
|     `device` | `string` | Equipment |
|     `app_version` | `string` | Client side version |
|     `public_ip` | `string` | Public IP |
|     `internal_ip` | `string` | Intranet IP |
|     `use_rtc_proxy` | `boolean` | Proxy service |
|     `location` | `string` | Location |
|     `network_type` | `string` | Network type |
|     `protocol` | `string` | Connection type |
|     `microphone` | `string` | Microphone |
|     `speaker` | `string` | Loudspeaker |
|     `camera` | `string` | Camera |
|     `audio` | `boolean` | Audio |
|     `video` | `boolean` | Video |
|     `sharing` | `boolean` | Share |
|     `join_time` | `string` | Time of joining the meeting |
|     `leave_time` | `string` | Time of departure |
|     `time_in_meeting` | `string` | Duration of participation |
|     `leave_reason` | `string` | Reason for leaving |
|   `page_token` | `string` | Page identifier, when has_more is true, a new page_token will also be returned. Otherwise, page_token will not be returned |
|   `has_more` | `boolean` | Whether the response body has more parameters | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
        "participants": [
            {
                "participant_name": "kehan",
                "department": "development",
                "user_id": "8efq90",
                "employee_id": "202205789",
                "phone": "021-883889",
                "email": "xxxx@163.com",
                "device": "windows",
                "app_version": "5.26.0-alpha.38",
                "public_ip": "27.xx.xx.183",
                "internal_ip": "192.xx.xx.13",
                "use_rtc_proxy": false,
                "location": "dongguan",
                "network_type": "wifi",
                "protocol": "udp",
                "microphone": "麦克风阵列(Realtek(R) Audio)",
                "speaker": "扬声器(Realtek(R) Audio)",
                "camera": "HD Camera",
                "audio": true,
                "video": true,
                "sharing": false,
                "join_time": "2022.12.23 11:16:59 (GMT+08:00)",
                "leave_time": "2022.12.23 11:18:51 (GMT+08:00)",
                "time_in_meeting": "00:01:52",
                "leave_reason": "主持人结束会议"
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
| 400 | 126006 | end time cannot exceed start time for one day or smaller than start time | The end time cannot exceed the start time by one day or be less than the start time |
| 400 | 126007 | userID and roomID cannot be set concurrently | Do not set userID and roomID at the same time |
| 400 | 126009 | cannot find meeting according to the params you provide | Recheck the param information |
| 404 | 121004 | data not exist | The requested data doesn't exist. |
| 403 | 121005 | no permission | No permission. Check the token type, operator identity and resource. |
