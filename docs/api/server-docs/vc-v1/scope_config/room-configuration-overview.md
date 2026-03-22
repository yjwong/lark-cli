---
title: "Resource introduction"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/vc-v1/scope_config/room-configuration-overview"
service: "vc-v1"
resource: "scope_config"
section: "Video Conferencing"
updated: "1689326170000"
---

# Resource introduction
Meeting room configuration is used to configure the background settings, resource management, etc. of the Fishu meeting room

## Method

> The corresponding original regular muiti-layer meeting room configuration API is as follows：
> * Query meeting room configuration
> * Set up meeting room configuration

**Query Room Configuration**

    This API is used to query a scope of room level or a room configuration.

**Set Room Configuration**

    This API is used to set a scope of room level or a room configuration.

**Get Reserve Limitation**

    This API is used to query a scope of room level or a room reserve limitation.

**Update Reserve Limitation**

    This API is used to set a scope of room level or a room reserve limitation.

**Get Reserve Admin**

    This API is used to query a scope of room level or a room reserve admin.

**Update Reserve Admin**

    This API is used to set a scope of room level or a room reserve admin.

### Field description

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `scope_type` | `int` | Yes | Scope type **Example value**: 1 **Optional values are**:  - `1`: Room level - `2`: Room  |
| `scope_id` | `string` | Yes | Scope ID **Example value**: "omm_608d34d82d531b27fa993902d350a307" |
| `scope_config` | `room_config` | No | Room configuration of a scope |
|   `room_background` | `string` | No | Lark room background **Example value**: "https://lf1-ttcdn-tos.pstatp.com/obj/xxx" |
|   `display_background` | `string` | No | Lark sign board background **Example value**: "https://lf1-ttcdn-tos.pstatp.com/obj/xxx" |
|   `digital_signage` | `room_digital_signage` | No | Lark room digital labels |
|     `enable` | `boolean` | No | Enable digital label function or not **Example value**: true |
|     `mute` | `boolean` | No | Whether to mute the video **Example value**: true |
|     `start_display` | `int` | No | End playing n min before the meeting starts **Example value**: 3 |
|     `stop_display` | `int` | No | Start playing n min after the meeting ends **Example value**: 3 |
|     `materials` | `room_digital_signage_material[]` | No | Material list |
|       `id` | `string` | No | Material ID **Example value**: "7847784676276" |
|       `name` | `string` | No | Material name **Example value**: "name" |
|       `material_type` | `int` | No | Material type **Example value**: 1 **Optional values are**:  - `1`: Image - `2`: Video - `3`: GIF  |
|       `url` | `string` | No | Material URL **Example value**: "url" |
|       `duration` | `int` | No | Play duration (in sec) **Example value**: 15 |
|       `cover` | `string` | No | Material cover URL **Example value**: "url" |
|       `md5` | `string` | No | Material file md5 **Example value**: "md5" |
|       `vid` | `string` | No | Material file vid **Example value**: "vid" |
|       `size` | `string` | No | Material desc **Example value**: "100 |
|   `room_box_digital_signage` | `room_digital_signage` | No | Lark room box digital signage |
|     `enable` | `boolean` | No | Enable digital label function or not **Example value**: true |
|     `mute` | `boolean` | No | Whether to mute the video **Example value**: true |
|     `start_display` | `int` | No | End playing n min before the meeting starts **Example value**: 3 |
|     `stop_display` | `int` | No | Start playing n min after the meeting ends **Example value**: 3 |
|     `materials` | `room_digital_signage_material[]` | No | Material list |
|       `id` | `string` | No | Material ID **Example value**: "7847784676276" |
|       `name` | `string` | No | Material name **Example value**: "name" |
|       `material_type` | `int` | No | Material type **Example value**: 1 **Optional values are**:  - `1`: Image - `2`: Video - `3`: GIF  |
|       `url` | `string` | No | Material URL **Example value**: "url" |
|       `duration` | `int` | No | Play duration (in sec) **Example value**: 15 |
|       `cover` | `string` | No | Material cover URL **Example value**: "url" |
|       `md5` | `string` | No | Material file md5 **Example value**: "md5" |
|       `vid` | `string` | No | Material file vid **Example value**: "vid" |
|       `size` | `string` | No | Material desc **Example value**: "100 |
|   `room_status` | `room_status` | No | Room status |
|     `status` | `boolean` | Yes | Available status of room **Example value**：true |
|     `schedule_status` | `boolean` | No | Future available status of room **Example value**：true |
|     `disable_start_time` | `string` | No | Room disable start time（second） **Example value**："1652356050" |
|     `disable_end_time` | `string` | No | Room disable end time（second） **Example value**："1652442450" |
|     `disable_reason` | `string` | No | Room disable reason **Example value**："Test Reason" |
|     `contact_ids` | `string[]` | No | Contact user ID **Example value**：["ou_3ec3f6a28a0d08c45d895276e8e5e19b"] |
|     `disable_notice` | `boolean` | No | Notice content sent to users when room disabled **Example value**：true |
|     `resume_notice` | `boolean` | No | Notice content sent to users when room resumed **Example value**：true | ### Request body example

{
    "scope_type": 1,
    "scope_id": "omm_608d34d82d531b27fa993902d350a307",
    "scope_config": {
        "room_background": "https://lf1-ttcdn-tos.pstatp.com/obj/xxx",
        "display_background": "https://lf1-ttcdn-tos.pstatp.com/obj/xxx",
        "digital_signage": {
            "enable": true,
            "mute": true,
            "start_display": 3,
            "stop_display": 3,
            "materials": [
                {
                    "id": "7847784676276",
                    "name": "name",
                    "material_type": 0,
                    "url": "url",
                    "duration": 15,
                    "cover": "url",
                    "md5": "md5",
                    "vid": "vid",
                    "size": "100"
                }
            ]
        },
        "room_box_digital_signage": {
            "enable": true,
            "mute": true,
            "start_display": 3,
            "stop_display": 3,
            "materials": [
                {
                    "id": "7847784676276",
                    "name": "name",
                    "material_type": 0,
                    "url": "url",
                    "duration": 15,
                    "cover": "url",
                    "md5": "md5",
                    "vid": "vid",
                    "size": "100"
                }
            ]
        },
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
        }
    }
}
