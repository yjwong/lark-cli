---
title: "Obtain meeting details"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/vc-v1/meeting/get"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/vc/v1/meetings/:meeting_id"
service: "vc-v1"
resource: "meeting"
section: "Video Conferencing"
rate_limit: "1000 per minute、50 per second"
scopes:
  - "vc:meeting:readonly"
field_scopes:
  - "contact:user.employee_id:readonly"
updated: "1692787315000"
---

# Obtain meeting details

Obtains the detailed data of a meeting.

> Users can only obtain their own meeting for the last 90 days.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/vc/v1/meetings/:meeting_id |
| HTTP Method | GET |
| Rate Limit | 1000 per minute、50 per second |
| Supported app types | custom,isv |
| Required scopes | `vc:meeting:readonly` |
| Required field scopes | > The response body of the API contains the following sensitive fields, and they will be returned only after corresponding scopes are added. If you do not need the fields, it is not recommended that you request the scopes. `contact:user.employee_id:readonly` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `meeting_id` | `string` | Meeting ID (Unique identifier of a video conference, generated after the conference is started) **Example value**: "6911188411932033028" | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `with_participants` | `boolean` | No | Indicates whether the participant list is required **Example value**: false |
| `with_meeting_ability` | `boolean` | No | Indicates whether individual statistics is required in the meeting (only for tenant_access_token) **Example value**: false |
| `user_id_type` | `string` | No | User ID categories **Example value**: open_id **Optional values are**:  - `open_id`: Identifies a user to an app. The same user has different Open IDs in different apps. How to get Open ID - `union_id`: Identifies a user to a tenant that acts as a developer. A user has the same Union ID in apps developed by the same developer, and has different Union IDs in apps developed by different developers. A developer can use Union ID to link the same user's identities in multiple apps.How to get Union ID - `user_id`: Identifies a user to a tenant. The same user has different User IDs in different tenants. In one single tenant, a user has the same User ID in all apps （including store apps）. User ID is usually used to communicate user data between different apps. How to get User ID  **Default value**: `open_id` **When the value is `user_id`, the following field scopes are required**: `contact:user.employee_id:readonly` | Note: When using user_id as user_id_type, only lark users can be obtained in the participant list, and users such as rooms/pstn/sip will be filtered. To obtain such users, please use open_id

## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `meeting` | `meeting` | Meeting data |
|     `id` | `string` | Meeting ID (Unique identifier of video conference, generated after the conference is started) |
|     `topic` | `string` | Meeting topic |
|     `url` | `string` | Meeting link (Lark users can join the meeting by clicking the meeting link) |
|     `meeting_no` | `string` | Meeting number |
|     `create_time` | `string` | Meeting creation time (Unix time in seconds) |
|     `start_time` | `string` | Meeting start time (Unix time in seconds) |
|     `end_time` | `string` | Meeting end time (Unix time in seconds) |
|     `host_user` | `meeting_user` | Host |
|       `id` | `string` | User ID |
|       `user_type` | `int` | User type **Optional values are**:  - `1`: Lark user - `2`: Rooms user - `3`: Docs user - `4`: neo Lark Meetings user - `5`: neo Lark Meetings guest - `6`: PSTN user - `7`: SIP user  |
|     `status` | `int` | Meeting status **Optional values are**:  - `1`: Calling - `2`: In progress - `3`: Ended  |
|     `participant_count` | `string` | Maximum number of participants |
|     `participant_count_accumulated` | `string` | Total number of participants |
|     `participants` | `meeting_participant[]` | Participant list |
|       `id` | `string` | User ID |
|       `first_join_time` | `string` | First joining time, which is a UNIX timestamp in seconds. |
|       `final_leave_time` | `string` | Final leaving time, which is a UNIX timestamp in seconds. |
|       `in_meeting_duration` | `string` | Cumulative time in the meeting (in seconds) |
|       `user_type` | `int` | User type **Optional values are**:  - `1`: Lark user - `2`: Rooms user - `3`: Docs user - `4`: neo Lark Meetings user - `5`: neo Lark Meetings guest - `6`: PSTN user - `7`: SIP user  |
|       `is_host` | `boolean` | Whether the participant is a host |
|       `is_cohost` | `boolean` | Whether the participant is a co-host |
|       `is_external` | `boolean` | Whether the participant is external |
|       `status` | `int` | Participant status **Optional values are**:  - `1`: Calling - `2`: In the meeting - `3`: Ringing - `4`: Not in the meeting or has left the meeting  |
|     `ability` | `meeting_ability` | Capability used in the meeting |
|       `use_video` | `boolean` | Indicates whether to use video |
|       `use_audio` | `boolean` | Indicates whether to use audio |
|       `use_share_screen` | `boolean` | Indicates whether to use screen sharing |
|       `use_follow_screen` | `boolean` | Indicates whether to use Magic Share |
|       `use_recording` | `boolean` | Indicates whether to use recording |
|       `use_pstn` | `boolean` | Indicates whether to use PSTN | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
        "meeting": {
            "id": "6911188411934433028",
            "topic": "my meeting",
            "url": "https://vc.larksuite.com/j/337736498",
            "meeting_no": "123456789",
            "create_time": "1608885566",
            "start_time": "1608883322",
            "end_time": "1608888867",
            "host_user": {
                "id": "ou_3ec3f6a28a0d08c45d895276e8e5e19b",
                "user_type": 1
            },
            "status": 2,
            "participant_count": "10",
            "participant_count_accumulated": "10",
            "participants": [
                {
                    "id": "ou_3ec3f6a28a0d08c45d895276e8e5e19b",
                    "first_join_time": "1624438144",
                    "final_leave_time": "1624438144",
                    "in_meeting_duration": "123",
                    "user_type": 1,
                    "is_host": true,
                    "is_cohost": false,
                    "is_external": false,
                    "status": 2
                }
            ],
            "ability": {
                "use_video": true,
                "use_audio": true,
                "use_share_screen": true,
                "use_follow_screen": true,
                "use_recording": true,
                "use_pstn": true
            }
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
