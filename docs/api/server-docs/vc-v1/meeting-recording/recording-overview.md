---
title: "Recording overview"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/vc-v1/meeting-recording/recording-overview"
service: "vc-v1"
resource: "meeting-recording"
section: "Video Conferencing"
updated: "1649822478000"
---

#  Recording overview
##  Definition
Users can record a meeting and get the link to the recording after the meeting is over. The events include: Start recording, Stop recording, Access the recording file, and Authorize access to the recording file.

##  Field description
| Name | Type | Description |
| --- | --- | --- |
| `meeting_id` | `string` | Meeting ID (the meeting's unique identifier; generated after the video meeting has started) **Example**："6911188411932033028" |
| `user_id_type` | `string` | User ID types **Example**："open_id" **Available options**： - `open_id`：The user's open ID - `union_id`：The user's union ID - `user_id`：The user's user ID **Default**：`open_id` **When the value is** `user_id` **,then the field permissions require**： `contact:user.employee_id:readonly` |
| `timezone` | `int` | Time zone used when displaying the recording file time: [-12,12] **Example**：8 |
| `data` | `-` | `-` |
| ∟ `recording` | `meeting.recording` | Recording file data |
|  ∟ `url` | `string` | Recording file URL |
|  ∟ `duration` | `string` | Total recording time (unit: msec) |
| `permission_objects` | `recording_permission_object[]` | List of authorized objects |
| ∟ `id` | `string` | Authorized object ID **Example**： "ou_3ec3f6a28a0d08c45d895276e8e5e19b" |
| ∟ `type` | `int` | Authorized object types **Example**：1 **Available options**： - `1`：User authorization - `2`：Group authorization - `3`：Within-tenant authorization (do not fill in the ID field) - `4`：Public network authorization (do not fill in the ID field) |
| ∟ `permission` | `int` | Permission **Example**：1 **Available options**： - `1`：View | ###  Request body example
```json
{
    "meeting_id":"6911188411932033028",
    "user_id_type":"open_id",
    "timezone":8,
    "data":{
        "recording":{
            "url":"https://meetings.larksuite.com/minutes/obcn37dxcftoc3656rgyejm7",
            "duration":"30000",
        }
    },
    "permission_objects":[
    {
        "id":"ou_3ec3f6a28a0d08c45d895276e8e5e19b",
        "type":1,
        "permission":1
    }
    ]
}
```
