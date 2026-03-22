---
title: "Resource introduction"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/vc-v1/reserve/schedule-meeting-overview"
service: "vc-v1"
resource: "reserve"
section: "Video Conferencing"
updated: "1692787314000"
---

#  Resource introduction
##  Resource definition
Users can schedule meetings, set participants and meeting permissions in advance, and obtain meeting information, including: schedule a meeting, update a meeting schedule, delete a scheduled meeting, obtain schedule details, and obtain a meeting in progress.

##  Field description
| Parameter | Type | Description |
| --- | --- | --- |
| `reserve_id` | `string` | Schedule ID (unique identifier of a schedule) **Example value**: "6911188411932033028" |
| `user_id_type` | `string` | User ID type **Example value**: "open_id" **Optional values are**: - `open_id`: User's  open id - `union_id`: User's  union id - `user_id`: User's  user id **Default value**: `open_id` **When the value is **`user_id`**, the required field scopes are**:** `contact:user.employee_id:readonly` |
| `end_time` | `string` | Expiration time (Unix time in seconds), which is required for multi-person meetings. **Example value**: "1608888867" |
| `meeting_settings` | `reserve_meeting_setting` | Meeting settings |
| ∟ `topic` | `string` | Meeting topic **Example value**: "my meeting" |
| ∟ `action_permissions` | `reserve_action_permission[]` | List of permission configurations for a meeting. If the same configuration option exists, the relationship between them is "Logical OR" (that is, the permission is obtained if any one of them is true). |
|  ∟ `permission` | `int` | Permission options **Example value**: 1 **Optional values are**: - `1`: Whether a participant can be a host - `2`: Whether a participant can invite other participants - `3`: Whether a participant can join a meeting |
|  ∟ `permission_checkers` | `reserve_permission_checker[]` | List of permission checkers. The relationship between permission checkers is "Logical OR" (that is, the permission is obtained if any one of them is true). |
|   ∟ `check_field` | `int` | Type of the fields to be checked **Example value**: 1 **Optional values are**: - `1`: User ID - `2`: User type - `3`: Tenant ID |
|   ∟ `check_mode` | `int` | Check method **Example value**: 1 **Optional values are**: - `1`: Being in the check_list indicates that the participant has the permission (allowlist) - `2`: Not being in the check_list indicates that the participant has the permission (blocklist) |
|   ∟ `check_list` | `string[]` | Field checklist **Example value**: 123 |
| ∟ `meeting_initial_type` | `int` | Initial type of meeting **Example value**: 1 **Optional values are**: - `1`: Multi-person meeting - `2`: 1v1 calling |
| ∟ `auto_record` | `boolean` | Whether to enable automatic recording when using Lark video conference, the default is false **Example value**: true |
| ∟ `assign_host_list` | `assign_host_list[]` | Assign host list params |
|  ∟ `user_type` | `int` | User type. Only support Lark user in the same tenant **Example Value**：1 **Optional values are**： - `1`：Lark user |
|  ∟ `id` | `string` | User ID **Example Value**："ou_3ec3f6a28a0d08c45d895276e8e5e19b" |
| ∟ `call_setting` | `reserve_call_setting` | 1v1 calling related parameters |
|  ∟ `callee` | `reserve_callee` | User being called |
|  ∟ `id` | `string` | User ID **Example value**: "ou_3ec3f6a28a0d08c45d895276e8e5e19b" |
|  ∟ `user_type` | `int` | User type. Only user type 6 (PSTN user) is supported. **Example value**: 1 **Optional values are**: - `1`: Lark user - `2`: Rooms user - `3`: Docs user - `4`: neo Lark Meetings user - `5`: neo Lark Meetings guest - `6`: PSTN user - `7`: SIP user |
|  ∟ `pstn_sip_info` | `pstn_sip_info` | pstn/sipInformation |
|   ∟ `nickname` | `string` | Temporary nickname set for a PSTN/SIP user **Example value**: "dodo" |
|   ∟ `main_address` | `string` | PSTN/SIP host number, in the format of [International Access Code]-[Area Code][Phone Number]. Only domestic mobile numbers and fixed-line telephone numbers are supported. **Example value**: "+86-02187654321" |
| `with_participants` | `boolean` | `Whether the participant list is required. The default value is false. **Example value**: false` |
| `data` | `-` | `-` |
| ∟ `meeting` | `meeting` | Meeting data |
|  ∟ `id` | `string` | Meeting ID (unique identifier of a video conference, generated after the conference is started) |
|  ∟ `topic` | `string` | Meeting topic |
|  ∟ `url` | `string` | Meeting link (Lark users can join the meeting by clicking the meeting link) |
|  ∟ `meeting_no` | `string` | Meeting number |
|  ∟ `create_time` | `string` | Meeting creation time (Unix time in seconds) |
|  ∟ `start_time` | `string` | Meeting start time (Unix time in seconds) |
|  ∟ `end_time` | `string` | Meeting end time (Unix time in seconds) |
|  ∟ `host_user` | `meeting_user` | Host |
|   ∟ `id` | `string` | User ID |
|   ∟ `user_type` | `int` | User type **Optional values are**: - `1`: Lark user - `2`: Rooms user - `3`: Docs user - `4`: neo Lark Meetings user - `5`: neo Lark Meetings guest - `6`: PSTN user - `7`: SIP user |
| ∟ `status` | `int` | Meeting status **Optional values are**: - `1`: Calling - `2`: In progress - `3`: Ended |
| ∟ `participant_count` | `string` | Number of participants |
| ∟ `participant_count_accumulated` | `string` | Total number of participants |
| ∟ `participants` | `meeting_participant[]` | Participant list |
|  ∟ `id` | `string` | User ID |
|  ∟ `first_join_time` | `string` | First joining time, which is a Unix timestamp in seconds. |
|  ∟ `final_leave_time` | `string` | Final leaving time, which is a Unix timestamp in seconds. |
|  ∟ ` in_meeting_duration` | `string` | Cumulative time in the meeting (in seconds) |
|  ∟ `user_type` | `int` | User type **Optional values are**: - `1`: Lark user - `2`: Rooms user - `3`: Docs user - `4`: neo Lark Meetings user - `5`: neo Lark Meetings guest - `6`: PSTN user - `7`: SIP user |
|  ∟ `is_host` | `boolean` | Indicates whether the participant is a host |
|  ∟ `is_cohost` | `boolean` | Indicates whether the participant is a co-host |
|  ∟ `is_external` | `boolean` | Indicates whether the participant is external |
|  ∟ `status` | `int` | Participant status **Optional values are**: - `1`: Calling - `2`: In the meeting - `3`: Ringing - `4`: Not in the meeting or has left the meeting |
| ∟ `ability` | `meeting_ability` | Capability used in the meeting |
|  ∟ ` use_video` | `boolean` | Indicates whether to use video |
|  ∟ `use_audio` | `boolean` | Indicates whether to use audio |
|  ∟ `use_share_screen` | `boolean` | Indicates whether to use screen sharing |
|  ∟ `use_follow_screen` | `boolean` | Indicates whether to use Magic Share |
|  ∟ `use_recording` | `boolean` | Indicates whether to use recording |
|  ∟ `use_pstn` | `boolean` | Indicates whether to use PSTN | ###  Data example
```json
{
    "reserve_id":"6911188411932033028",
    "user_id_type":"open_id",
    "end_time":"1608888867",
    "meeting_settings":{
        "topic":"my meeting",
        "action_permissions":[
            {
                "permission":1,
                "permission_checkers":[
                    {
                        "check_field":1,
                        "check_mode":1,
                        "check_list": [
                            "ou_3ec3f6a28a0d08c45d895276e8e5e19b"
                        ]
                    }
                ]
            }
        ],
         "meeting_initial_type":1,
         "auto_record":true,
         "call_setting":
         {
         "callee":
             {
                 "id":"ou_3ec3f6a28a0d08c45d895276e8e5e19b",
                 "user_type":1,
                 "pstn_sip_info":{
                     "nickname":"dodo",
                     "main_address":"+86-02187654321"
                 }
             }
         }
    },
    "with_participants":false,
    "data":{
        "meeting":{
            "id": "6911188411934433028",
            "topic": "my meeting",
            "url":"https://vc.larksuite.com/j/337736498",
            "meeting_no": "235812466",
            "create_time":"1608885566",
            "start_time": "1608883322",
            "end_time": "1608883899",
            "host_user": {
                "id":"ou_3ec3f6a28a0d08c45d895276e8e5e19b",
                "user_type": 1
            },
        },
        "status": 2,
            "participant_count": "10",
            "participant_count_accumulated":"15",
            "participants": [
                {
                    "id": "ou_3ec3f6a28a0d08c45d895276e8e5e19b",
                    "first_join_time":"1624438144",
                    "final_leave_time":"1624438144",
                    "in_meeting_duration":"123",
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
```
## Meeting ID description
Learns about the differences between a meeting number and a meeting ID, their respective usage, and the methods for obtaining them.
- What is meeting ID?

Meeting ID is the unique identifier of a meeting, which is different from the 9-digit meeting number seen on the client-side. Meeting ID is used as the unique identifier for calling meeting-related APIs.
- How to obtain a meeting ID？

For a meeting scheduled via the  Schedule a meeting API  API, its meeting ID can be obtained via the  Obtain an active meeting API  API after the meeting starts or by listening on the  Meeting-related events .
