---
title: "Obtain a schedule"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/vc-v1/reserve/get"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/vc/v1/reserves/:reserve_id"
service: "vc-v1"
resource: "reserve"
section: "Video Conferencing"
rate_limit: "1000 per minute、50 per second"
scopes:
  - "vc:reserve:readonly"
field_scopes:
  - "contact:user.employee_id:readonly"
updated: "1692787314000"
---

# Obtain a schedule

Obtains details about a schedule.

> You can only obtain your own schedules.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/vc/v1/reserves/:reserve_id |
| HTTP Method | GET |
| Rate Limit | 1000 per minute、50 per second |
| Supported app types | custom,isv |
| Required scopes | `vc:reserve:readonly` |
| Required field scopes | > The response body of the API contains the following sensitive fields, and they will be returned only after corresponding scopes are added. If you do not need the fields, it is not recommended that you request the scopes. `contact:user.employee_id:readonly` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `reserve_id` | `string` | Schedule ID (unique identifier of a schedule) **Example value**: "6911188411932033028" | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `user_id_type` | `string` | No | User ID categories **Example value**: open_id **Optional values are**:  - `open_id`: Identifies a user to an app. The same user has different Open IDs in different apps. How to get Open ID - `union_id`: Identifies a user to a tenant that acts as a developer. A user has the same Union ID in apps developed by the same developer, and has different Union IDs in apps developed by different developers. A developer can use Union ID to link the same user's identities in multiple apps.How to get Union ID - `user_id`: Identifies a user to a tenant. The same user has different User IDs in different tenants. In one single tenant, a user has the same User ID in all apps （including store apps）. User ID is usually used to communicate user data between different apps. How to get User ID  **Default value**: `open_id` **When the value is `user_id`, the following field scopes are required**: `contact:user.employee_id:readonly` | ## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `reserve` | `reserve` | Schedule data |
|     `id` | `string` | Reservation ID (unique identifier of the reservation, not meeting ID, meeting ID is only generated after the meeting starts) |
|     `meeting_no` | `string` | 9-digit meeting number (Lark users can join the meeting by entering this number) |
|     `url` | `string` | Meeting link (Lark users can join the meeting by clicking the meeting link) |
|     `app_link` | `string` | APPLink is used to open Lark App for joining. "{?}" is the placeholder, which is used to configure the parameters for joining the meeting and should be replaced with a specific value: 0 indicates Disabled and 1 indicates Enabled. Users can set the status of their microphone, speaker, and camera before joining the meeting in the Preview page. |
|     `live_link` | `string` | Meeting to livestream link |
|     `end_time` | `string` | Expiration time (Unix time in seconds) |
|     `expire_status` | `int` | Expiration status **Optional values are**:  - `1`: Not expired - `2`: Expired  |
|     `reserve_user_id` | `string` | ID of the user who scheduled |
|     `meeting_settings` | `reserve_meeting_setting` | Meeting settings |
|       `topic` | `string` | Meeting topic |
|       `action_permissions` | `reserve_action_permission[]` | List of permission configurations for a meeting. If the same configuration option exists, the relationship between them is "Logical OR" (that is, the permission is obtained if any one of them is true). |
|         `permission` | `int` | Permission options **Optional values are**:  - `1`: Whether a participant can be a host - `2`: Whether a participant can invite other participants - `3`: Whether a participant can join a meeting  |
|         `permission_checkers` | `reserve_permission_checker[]` | List of permission checkers. The relationship between permission checkers is "Logical OR" (that is, the permission is obtained if any one of them is true). |
|           `check_field` | `int` | Type of the fields to be checked **Optional values are**:  - `1`: User ID (check_list fill in the user ID) - `2`: User type (check_list optional values are "1": lark user, "2": rooms users, "6": pstn users, "7": sip user) - `3`: Tenant ID (check_list fill in the tenant_key)  |
|           `check_mode` | `int` | Check method **Optional values are**:  - `1`: Being in the check_list indicates that the participant has the permission (allowlist) - `2`: Not being in the check_list indicates that the participant has the permission (blocklist)  |
|           `check_list` | `string[]` | Check the list of fields (fill in the corresponding content according to the type of check_field) |
|       `meeting_initial_type` | `int` | Initial type of meeting **Optional values are**:  - `1`: Multi-person meeting - `2`: 1v1 calling  |
|       `call_setting` | `reserve_call_setting` | 1v1 calling related parameters |
|         `callee` | `reserve_callee` | User being called |
|           `id` | `string` | User ID |
|           `user_type` | `int` | User type. Only user type 6 (PSTN user) is supported. **Optional values are**:  - `1`: Lark user - `2`: Rooms user - `3`: Docs user - `4`: neo Lark Meetings user - `5`: neo Lark Meetings guest - `6`: PSTN user - `7`: SIP user  |
|           `pstn_sip_info` | `pstn_sip_info` | PSTN/SIP information |
|             `nickname` | `string` | Temporary nickname set for a PSTN/SIP user |
|             `main_address` | `string` | PSTN/SIP host number, in the format of [International Access Code]-[Area Code][Phone Number]. Only domestic mobile numbers and fixed-line telephone numbers are supported. |
|       `auto_record` | `boolean` | Whether to enable automatic recording when using Lark video conference, the default is false |
|       `assign_host_list` | `reserve_assign_host[]` | Assign host list params |
|         `user_type` | `int` | User type **Optional values are**:  - `1`: Lark user  |
|         `id` | `string` | User ID | ### Response body example

{"code":0,
"msg":"success",
"data":{"reserve":{"id":"6911188411934973028",
"meeting_no":"112000358",
"url":"https://vc.larksuite.com/j/337736498",
"app_link":"https://applink.larksuite.com/client/videochat/open?source

vc.v1.reserve.method.apply.error-mapping.table.desc-column.width=$$$20%",
"live_link":"https://meetings.larksuite.com/s/1gub381l4gglv",
"end_time":"1608883322",
"expire_status":0,
"reserve_user_id":"ou_3ec3f6a28a0d08c45d895276e8e5e19b",
"meeting_settings":{"topic":"my meeting",
"action_permissions":[{
    "permission": 1,
    "permission_checkers": [
        {
            "check_field": 1,
            "check_mode": 1,
            "check_list": [
                "ou_3ec3f6a28a0d08c45d895276e8e5e19b"
            ]
        }
    ]
}],
"meeting_initial_type":1,
"call_setting":{"callee":{"id":"ou_3ec3f6a28a0d08c45d895276e8e5e19b",
"user_type":1,
"pstn_sip_info":{"nickname":"dodo",
"main_address":"+86-02187654321"}}},
"auto_record":true,
"assign_host_list":[{
    "user_type": 1,
    "id": "ou_3ec3f6a28a0d08c45d895276e8e5e19b"
}]}}}}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 500 | 121001 | internal error | Internal server error. If retry still doesn't work, contact the administrator. |
| 400 | 121002 | not support | This function is not supported. |
| 400 | 121003 | param error | Parameter error. Check the value range of parameters (please check yourself according to the field description above). |
| 404 | 121004 | data not exist | The requested data doesn't exist. |
| 403 | 121005 | no permission | No permission. Check the token type, operator identity and resource. |
