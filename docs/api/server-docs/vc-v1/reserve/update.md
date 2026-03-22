---
title: "Update a schedule"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/vc-v1/reserve/update"
method: "PUT"
api_path: "https://open.larksuite.com/open-apis/vc/v1/reserves/:reserve_id"
service: "vc-v1"
resource: "reserve"
section: "Video Conferencing"
rate_limit: "1000 per minute、50 per second"
scopes:
  - "vc:reserve"
field_scopes:
  - "contact:user.employee_id:readonly"
updated: "1692787314000"
---

# Update a schedule

Updates a schedule.

> You can only update your own schedules. Do not input the fields that do not need to be updated (If left empty, these fields are not updated). It can be used for renewal, and the expiration date cannot exceed 30 days from the current date.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/vc/v1/reserves/:reserve_id |
| HTTP Method | PUT |
| Rate Limit | 1000 per minute、50 per second |
| Supported app types | custom,isv |
| Required scopes | `vc:reserve` |
| Required field scopes | > The response body of the API contains the following sensitive fields, and they will be returned only after corresponding scopes are added. If you do not need the fields, it is not recommended that you request the scopes. `contact:user.employee_id:readonly` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `reserve_id` | `string` | Schedule ID (unique identifier of a schedule) **Example value**: "6911188411932033028" | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `user_id_type` | `string` | No | User ID categories **Example value**: open_id **Optional values are**:  - `open_id`: Identifies a user to an app. The same user has different Open IDs in different apps. How to get Open ID - `union_id`: Identifies a user to a tenant that acts as a developer. A user has the same Union ID in apps developed by the same developer, and has different Union IDs in apps developed by different developers. A developer can use Union ID to link the same user's identities in multiple apps.How to get Union ID - `user_id`: Identifies a user to a tenant. The same user has different User IDs in different tenants. In one single tenant, a user has the same User ID in all apps （including store apps）. User ID is usually used to communicate user data between different apps. How to get User ID  **Default value**: `open_id` **When the value is `user_id`, the following field scopes are required**: `contact:user.employee_id:readonly` | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `end_time` | `string` | No | Expiration time (Unix time in seconds) **Example value**: "1608888867" |
| `meeting_settings` | `reserve_meeting_setting` | No | Meeting settings |
|   `topic` | `string` | No | Meeting topic **Example value**: "my meeting" |
|   `action_permissions` | `reserve_action_permission[]` | No | List of permission configurations for a meeting. If the same configuration option exists, the relationship between them is "Logical OR" (that is, the permission is obtained if any one of them is true). |
|     `permission` | `int` | Yes | Permission options **Example value**: 1 **Optional values are**:  - `1`: Whether a participant can be a host - `2`: Whether a participant can invite other participants - `3`: Whether a participant can join a meeting  |
|     `permission_checkers` | `reserve_permission_checker[]` | Yes | List of permission checkers. The relationship between permission checkers is "Logical OR" (that is, the permission is obtained if any one of them is true). |
|       `check_field` | `int` | Yes | Type of the fields to be checked **Example value**: 1 **Optional values are**:  - `1`: User ID (check_list fill in the user ID) - `2`: User type (check_list optional values are "1": lark user, "2": rooms users, "6": pstn users, "7": sip user) - `3`: Tenant ID (check_list fill in the tenant_key)  |
|       `check_mode` | `int` | Yes | Check method **Example value**: 1 **Optional values are**:  - `1`: Being in the check_list indicates that the participant has the permission (allowlist) - `2`: Not being in the check_list indicates that the participant has the permission (blocklist)  |
|       `check_list` | `string[]` | Yes | Check the list of fields (fill in the corresponding content according to the type of check_field) **Example value**: ["ou_3ec3f6a28a0d08c45d895276e8e5e19b"] |
|   `meeting_initial_type` | `int` | No | Initial type of meeting **Example value**: 1 **Optional values are**:  - `1`: Multi-person meeting - `2`: 1v1 calling  |
|   `call_setting` | `reserve_call_setting` | No | 1v1 calling related parameters |
|     `callee` | `reserve_callee` | Yes | User being called |
|       `id` | `string` | No | User ID **Example value**: "ou_3ec3f6a28a0d08c45d895276e8e5e19b" |
|       `user_type` | `int` | Yes | User type. Only user type 6 (PSTN user) is supported. **Example value**: 1 **Optional values are**:  - `1`: Lark user - `2`: Rooms user - `3`: Docs user - `4`: neo Lark Meetings user - `5`: neo Lark Meetings guest - `6`: PSTN user - `7`: SIP user  |
|       `pstn_sip_info` | `pstn_sip_info` | No | PSTN/SIP information |
|         `nickname` | `string` | No | Temporary nickname set for a PSTN/SIP user **Example value**: "dodo" |
|         `main_address` | `string` | Yes | PSTN/SIP host number, in the format of [International Access Code]-[Area Code][Phone Number]. Only domestic mobile numbers and fixed-line telephone numbers are supported. **Example value**: "+86-02187654321" |
|   `auto_record` | `boolean` | No | Whether to enable automatic recording when using Lark video conference, the default is false **Example value**: true |
|   `assign_host_list` | `reserve_assign_host[]` | No | Assign host list params |
|     `user_type` | `int` | No | User type **Example value**: 1 **Optional values are**:  - `1`: Lark user  |
|     `id` | `string` | No | User ID **Example value**: "ou_3ec3f6a28a0d08c45d895276e8e5e19b" | ### Request body example

{
    "end_time": "1608888867",
    "meeting_settings": {
        "topic": "my meeting",
        "action_permissions": [
            {
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
            }
        ],
        "meeting_initial_type": 1,
        "call_setting": {
            "callee": {
                "id": "ou_3ec3f6a28a0d08c45d895276e8e5e19b",
                "user_type": 1,
                "pstn_sip_info": {
                    "nickname": "dodo",
                    "main_address": "+86-02187654321"
                }
            }
        },
        "auto_record": true,
        "assign_host_list": [
            {
                "user_type": 1,
                "id": "ou_3ec3f6a28a0d08c45d895276e8e5e19b"
            }
        ]
    }
}

## Response

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
|     `live_link` | `string` | Meeting to livestream link |
|     `end_time` | `string` | Expiration time (Unix time in seconds) |
|     `expire_status` | `int` | Expiration status **Optional values are**:  - `1`: Not expired - `2`: Expired  |
|   `reserve_correction_check_info` | `reserve_correction_check_info` | Reserve params correction check infomation |
|     `invalid_host_id_list` | `string[]` | Invalid host id list | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
        "reserve": {
            "id": "6911188411934973028",
            "meeting_no": "112000358",
            "url": "https://vc.larksuite.com/j/337736498",
            "live_link": "https://meetings.larksuite.com/s/1gub381l4gglv",
            "end_time": "1608883322",
            "expire_status": 0
        },
        "reserve_correction_check_info": {
            "invalid_host_id_list": [
                "ou_3ec3f6a28a0d08c45d895276e8e5e19b"
            ]
        }
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
| 400 | 123002 | record param error | Recording parameter error, 1v1 call does not support automatic recording |
| 400 | 123003 | assign host exceed max limit 10 | assign host exceed max limit, max limit is 10 |
| 400 | 123004 | assign host error | assign host error, check whether id is valid |
