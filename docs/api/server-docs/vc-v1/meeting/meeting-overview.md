---
title: "Meeting overview"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/vc-v1/meeting/meeting-overview"
service: "vc-v1"
resource: "meeting"
section: "Video Conferencing"
updated: "1649822474000"
---

#  Meeting overview 
##  Definition
Users can invite and remove participants and set meeting hosts during a meeting. The methods include：Access meeting details, [Access related meetings by meeting number](//ssl:ttdoc/uAjLw4CM/ukTMukTMukTM/reference/vc-v1/meeting/list_by_no), Invite participants, Remove participants, Set the host, and End meetings. The events include: Meeting starts, Meeting ends, Join a meeting, Leave a meeting, Start a Recording, Stop a Recording, Recording completed, Start screen sharing, and End screen sharing.

##  Field description
| Name | Type | Description |
| --- | --- | --- |
| `data` | `-` | `-` |
| ∟ `meeting` | `meeting` | Meeting data |
|  ∟ `id` | `string` | Meeting ID (the meeting's unique identifier; generated after the video meeting has started) |
|  ∟ `topic` | `string` | Meeting topic |
|  ∟ `url` | `string` | Meeting link (Lark users can click the link to join the meeting) |
|  ∟ `create_time` | `string` | Meeting creation time (unix time, unit: sec) |
|  ∟ `start_time` | `string` | Meeting start time (unix time, unit: sec) |
|  ∟ `end_time` | `string` | Meeting end time (unix time, unit: sec) |
|  ∟ `host_user` | `meeting_user` | Host |
|   ∟ `id` | `string` | User ID |
|   ∟ `user_type` | `int` | User types **Available options**： - `1`：lark user - `2`：rooms user - `3`：docs user - `4`：neo single item user - `5`：neo single item guest user - `6`：pstn user - `7`：sip user |
|  ∟ `status` | `int` | Meeting status **Available options**： - `1`：Calling meeting - `2`：Meeting in progress - `3`：Meeting ended |
|  ∟ `participant_count` | `string` | Number of participants |
|  ∟ `participants` | `meeting_participant[]` | Participant list |
|   ∟ `id` | `string` | User ID |
|   ∟ `user_type` | `int` | User types **Available options**： - `1`：lark user - `2`：rooms user - `3`：docs user - `4`：neo single item user - `5`：neo single item guest user - `6`：pstn user - `7`：sip user |
|   ∟ `is_host` | `boolean` | Whether the user is the host |
|   ∟ `is_cohost` | `boolean` | Whether the user is the co-host |
|   ∟ `is_external` | `boolean` | Whether the user is an external participant |
|   ∟ `status` | `int` | Participant status **Available options**： - `1`：Calling - `2`：In-meeting - `3`：Ringing - `4`：Not in meeting or left meeting |
|  ∟ `ability` | `meeting_ability` | Functions that can be used in the meeting |
|   ∟ ` use_video` | `boolean` | Whether can use video |
|   ∟ `use_audio` | `boolean` | Whether the user can use audio |
|   ∟ `use_share_screen` | `boolean` | Whether the user can share screen |
|   ∟ `use_follow_screen` | `boolean` | Whether the user can use Magic Share |
|   ∟ `use_recording` | `boolean` | Whether the user can record |
|   ∟ `use_pstn` | `boolean` | Whether the user can use PSTN | ###  Request body example
```json
{
    "data": {
        "meeting": {
            "id": "6911188411934433028",
            "topic": "my meeting",
            "url": "https://vc.larksuite.com/j/337736498",
            "create_time": "1608885566",
            "start_time": "1608883322",
            "end_time": "1608888867",
            "host_user": {
                "id": "ou_3ec3f6a28a0d08c45d895276e8e5e19b",
                "user_type": 1
            },
            "status": 2,
            "participant_count": "10",
            "participants": [
                {
                    "id": "ou_3ec3f6a28a0d08c45d895276e8e5e19b",
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
```
