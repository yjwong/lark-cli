---
title: "Resource introduction"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/vc-v1/room/room-overview"
service: "vc-v1"
resource: "room"
section: "Video Conferencing"
updated: "1689326165000"
---

# Resource introduction

Users can query meeting room levels, create meeting room levels, update meeting room levels, and other operations.

## Methods

> The corresponding original regular muiti-layer meeting room API is as follows：
> * Obtain meeting room list
> * Query Meeting Room Details
> * Create Meeting Room
> * Update Meeting Room
> * Delete Meeting Room
 If the tenant has switched to the flexible level of the meeting room, compatible replacement will be used when calling the fixed level interface, which may lead to a downgraded experience. It is recommended to use the flexible layer interface

**Query Meeting Room List**

    This API is used to query meeting room list at a specified meeting room level.

**Query Meeting Room Details**

    This API is used to query a specified meeting room details by room ID.

**Batch Query Meeting Room Details**

    This API is used to query multi specified meeting room details by room IDs.

**Create Meeting Room**

    This API is used to create meeting room

**Update Meeting Room**

    This API is used to update a specified meeting room details.

**Delete Meeting Room**

    This API is used to delete a specified meeting room.

**Search Meeting Room**

    This API is used to search meeting room. Support keyword or custom room ID to search.

## Events

> Corresponding to the original non-multi-level meeting room events are as follows:
> * Meeting room created
> * Meeting Room Updated
> * Meeting Room Deleted

**Create Meeting Room**

      This event is triggered when a meeting room is created.

**Update Meeting Room**

      This event is triggered when a meeting room is updated.

**Delete Meeting Room**

      This event is triggered when a meeting room is deleted.

## Field description

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `room_id` | `string` | Yes | Room ID **Example value**："omm_4de32cf10a4358788ff4e09e37ebbf9b" |
| `name` | `string` | Yes | Room name **Example value**："TestRoom" |
| `capacity` | `int` | Yes | Room capacity **Example value**：10 |
| `description` | `string` | No | Room description **Example value**："Test Description" |
| `custom_room_id` | `string` | No | Custom room ID **Example value**："1234" |
| `room_level_id` | `string` | Yes | Parent room level ID **Example value**："omb_4ad1a2c7a2fbc5fc9570f38456931293" Reference：Room Level Overview |
| `path` | `string[]` | Yes | Room level path **Example value**：["omb_4ad1a2c7a2fbc5fc9570f38456931293"] |
| `room_status` | `room_status` | No | Room status |
|   `status` | `boolean` | Yes | Available status of room **Example value**：true |
|   `schedule_status` | `boolean` | No | Future available status of room **Example value**：true |
|   `disable_start_time` | `string` | No | Room disable start time（second） **Example value**："1652356050" |
|   `disable_end_time` | `string` | No | Room disable end time（second） **Example value**："1652442450" |
|   `disable_reason` | `string` | No | Room disable reason **Example value**："Test Reason" |
|   `contact_ids` | `string[]` | No | Contact user ID **Example value**：["ou_3ec3f6a28a0d08c45d895276e8e5e19b"] |
|   `disable_notice` | `boolean` | No | Notice content sent to users when room disabled **Example value**：true |
|   `resume_notice` | `boolean` | No | Notice content sent to users when room resumed **Example value**：true |
| `device` | `device[]` | No | device information |
|   `name` | `string` | Yes | device name **Example value**：电话 | ## Request body example

{
  	"room_id": "omm_4de32cf10a4358788ff4e09e37ebbf9b",
    "name": "TestRoom",
    "capacity": 10,
    "description": "Test Description",
    "custom_room_id": "1234",
    "room_level_id": "omb_4ad1a2c7a2fbc5fc9570f38456931293",
  	"path": [
  		"omb_4ad1a2c7a2fbc5fc9570f38456931293"
  	],
    "room_status": {
        "status": true,
        "schedule_status": true,
        "disable_start_time": "1652356050",
        "disable_end_time": "1652442450",
        "disable_reason": "TestReason",
        "contact_ids": [
            "ou_3ec3f6a28a0d08c45d895276e8e5e19b"
        ],
        "disable_notice": true,
        "resume_notice": true
    },
  	"device": [
  		{
  			"name": "电话"
  		}
  	]
}
