---
title: "Update reserve limitation"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/vc-v1/reserve_config/patch"
method: "PATCH"
api_path: "https://open.larksuite.com/open-apis/vc/v1/reserve_configs/:reserve_config_id"
service: "vc-v1"
resource: "reserve_config"
section: "Video Conferencing"
rate_limit: "100 per minute"
scopes:
  - "vc:room"
field_scopes:
  - "contact:user.employee_id:readonly"
updated: "1689326170000"
---

# Update reserve limitation

Update reserve limitation

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/vc/v1/reserve_configs/:reserve_config_id |
| HTTP Method | PATCH |
| Rate Limit | 100 per minute |
| Supported app types | custom,isv |
| Required scopes | `vc:room` |
| Required field scopes | > The response body of the API contains the following sensitive fields, and they will be returned only after corresponding scopes are added. If you do not need the fields, it is not recommended that you request the scopes. `contact:user.employee_id:readonly` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `reserve_config_id` | `string` | room id or room level id **Example value**: "omm_3c5dd7e09bac0c1758fcf9511bd1a771" | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `user_id_type` | `string` | No | User ID categories **Example value**: open_id **Optional values are**:  - `open_id`: Identifies a user to an app. The same user has different Open IDs in different apps. How to get Open ID - `union_id`: Identifies a user to a tenant that acts as a developer. A user has the same Union ID in apps developed by the same developer, and has different Union IDs in apps developed by different developers. A developer can use Union ID to link the same user's identities in multiple apps.How to get Union ID - `user_id`: Identifies a user to a tenant. The same user has different User IDs in different tenants. In one single tenant, a user has the same User ID in all apps （including store apps）. User ID is usually used to communicate user data between different apps. How to get User ID  **Default value**: `open_id` **When the value is `user_id`, the following field scopes are required**: `contact:user.employee_id:readonly` | > Reminder
> 1.  If the configuration of the level is updated, the configuration of all subordinate levels and rooms will be overwritten.
> 2.  When updating, approved_config, time_config, reserve_scope_config can update any one, or two or all of them.

### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `scope_type` | `string` | Yes | 1 means level, 2 means room **Example value**: "2" |
| `approval_config` | `approval_config` | No | reservation approval settings |
|   `approval_switch` | `int` | No | approval_switch: 0 means closed, 1 means open. Note:  1.  If the value is not set, the value of the original switch will not be updated, but the approval_condition is required at this time 2.  When the value is set to 1, approval_condition is required 3.  When the value is set to 0, other fields of the entire approval_config can be omitted. **Example value**: 1 **Data validation rules**: - Value range: `0` ～ `1` |
|   `approval_condition` | `int` | No | approval_condition: 0 means that all reservations require approval, 1 means that approval is required for those that meet the conditions Note: meeting_duration is required when approval_condition is 1 **Example value**: 1 **Data validation rules**: - Value range: `0` ～ `1` |
|   `meeting_duration` | `number(float)` | No | bookings exceeding meeting_duration need to be approved (unit: hour, value range [0.1-99]) Note:  1.  When the approval_condition is 0, if the value is not set when updating, the default update is 99. 2.  If the incoming value exceeds 2 decimal places, it will be automatically rounded to the nearest 2 digits. **Example value**: 3 |
|   `approvers` | `subscribe_user[]` | No | Approvers list, when the approval_switch is 1, at least one approver needs to be set **Example value**: [{user_id:"ou_e8bce6c3935ef1fc1b432992fd9d3db8"}] |
|     `user_id` | `string` | Yes | Reserve admin user ID **Example value**: "ou_a27b07a9071d90577c0177bcec98f856" |
| `time_config` | `time_config` | No | reservation time settings |
|   `time_switch` | `int` | Yes | reservation time switch: 0 for off, 1 for on **Example value**: 1 **Data validation rules**: - Value range: `0` ～ `1` |
|   `days_in_advance` | `int` | No | the earliest time when the conference room can be reserved in advance (unit: day, value range [1-730]) Note: When not filled in, the default update is 365 **Example value**: 30 |
|   `opening_hour` | `string` | No | You can start booking from opening_hour on the opening day (unit: second, value range [0,86400]) Note:  1.  If not filled in, it will be updated to 28800 by default 2.  If the filled value is not a multiple of 60, it will be automatically updated to the nearest integer multiple of 60. **Example value**: "27900" |
|   `start_time` | `string` | No | The start time of the daily scheduled time range (unit: second, value range [0,86400]) Note:  1.  When not filled in, it will be updated to 0 by default, and the end_time filled in at this time must not be less than 30. 2.  When both start_time and end_time are filled in, end_time should be at least 30 times longer than start_time. 3.   If the filled value is not a multiple of 60, it will be automatically updated to the nearest integer multiple of 60. **Example value**: "0" |
|   `end_time` | `string` | No | End time of daily pre-determinable time range (unit: second, value range [0,86400]) Note:  1.  When not filled in, the default update is 86400, and the start_time filled in at this time must not be greater than or equal to 86370. 2.  When both start_time and end_time are filled in, end_time must be at least 30 longer than start_time. 3.  If the filled value is not a multiple of 60, it will be automatically updated to the nearest integer multiple of 60. **Example value**: "86400" |
|   `max_duration` | `int` | No | The upper limit of the duration of a single conference room reservation (unit: hour, value range [1,99]) Note: The default update is 2 when not filled in **Example value**: 24 |
| `reserve_scope_config` | `reserve_scope_config` | No | reservation range settings |
|   `allow_all_users` | `int` | No | The range of members that can book conference: 0 for some members, 1 for all members. Note:  1.  This value is required. 2.  When set to 0, at least 1 reservation department or reservation person is required **Example value**: 0 **Data validation rules**: - Value range: `0` ～ `1` |
|   `allow_users` | `subscribe_user[]` | No | List of available members **Example value**: [{user_id:"ou_e8bce6c3935ef1fc1b432992fd9d3db8"}] |
|     `user_id` | `string` | Yes | Reserve admin user ID **Example value**: "ou_a27b07a9071d90577c0177bcec98f856" |
|   `allow_depts` | `subscribe_department[]` | No | List of available departments **Example value**: [{department_id:"od-5c07f0c117cf8795f25610a69363ce31"}] |
|     `department_id` | `string` | Yes | Reserve admin department ID **Example value**: "od-47d8b570b0a011e9679a755efcc5f61a" | ### Request body example

{
    "scope_type": "2",
    "approval_config": {
        "approval_switch": 1,
        "approval_condition": 1,
        "meeting_duration": 3,
        "approvers": [
            {
                "user_id": "ou_a27b07a9071d90577c0177bcec98f856"
            }
        ]
    },
    "time_config": {
        "time_switch": 1,
        "days_in_advance": 30,
        "opening_hour": "27900",
        "start_time": "0",
        "end_time": "86400",
        "max_duration": 24
    },
    "reserve_scope_config": {
        "allow_all_users": 0,
        "allow_users": [
            {
                "user_id": "ou_a27b07a9071d90577c0177bcec98f856"
            }
        ],
        "allow_depts": [
            {
                "department_id": "od-47d8b570b0a011e9679a755efcc5f61a"
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
| `data` | `\-` | \- | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {}
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 500 | 121001 | internal error | Internal server error. If retry still doesn't work, contact the administrator. |
| 400 | 121002 | not support | This function is not supported. |
| 400 | 121003 | param error | Parameter error. Check the value range of parameters (please check yourself according to the field description above). |
| 404 | 121004 | data not exist | The requested data doesn't exist. |
| 403 | 121005 | no permission | No permission. Check the token type, operator identity and resource. |
