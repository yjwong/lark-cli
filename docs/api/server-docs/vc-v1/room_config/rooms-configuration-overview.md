---
title: "Rooms configuration overview"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/vc-v1/room_config/rooms-configuration-overview"
service: "vc-v1"
resource: "room_config"
section: "Deprecated Version (Not Recommended)"
updated: "1686900968000"
---

# Room configuration overview
##  Resource definition
This API is used to set room background and manage resources, including Query room configurations and Set room configurations.

##  Field description
| Parameter | Type | Description |
| --- | --- | --- |
| `scope` | `int` | Query node scope **Example value**: "5" **Optional values are**: - `1`: Tenant - `2`: Country/region - `3`: City - `4`: Building - `5`: Floor - `6`: Room |
| `country_id` | `string` | The parameter is required when the country/region ID scope is 2 or 3 **Example value**: "086" |
| `district_id` | `string` | The parameter is required when the city ID scope is 3 **Example value**: "001 |
| `building_id` | `string` | The parameter is required when the building ID scope is 4 or 5 **Example value**: "22" |
| `floor_name` | `string` | The parameter is required when the floor scope is 5 **Example value**: "4" |
| `room_id` | `string` | The parameter is required when the room ID scope is 6 **Example value**: "6383786266263" |
| `code` | `int` | Error code. A value other than 0 indicates failure. |
| `msg` | `string` | Error description |
| `data` | `room_config` | `-` |
| ∟ `room_background` | `string` | Lark room background |
| ∟ `display_background` | `string` | Lark sign board background |
| ∟ `digital_signage` | `room_digital_signage` | Lark room digital labels |
|  ∟ `enable` | `boolean` | Enable digital label function or not |
|  ∟ `mute` | `boolean` | Whether to mute the video |
|  ∟ `start_display` | `int` | End playing n min before the meeting starts |
|  ∟ `stop_display` | `int` | Start playing n min after the meeting ends |
|  ∟ `materials` | `room_digital_signage_material[]` | Material list |
|  ∟ `id` | `string` | Material ID |
|  ∟ `name` | `string` | Material name |
|  ∟ `material_type` | `int` | Material type **Optional values are**: - `1`: Image - `2`: Video - `3`: GIF |
|  ∟ `url` | `string` | Material URL |
|  ∟ `duration` | `int` | Play duration (in sec) |
|  ∟ `cover` | `string` | Material cover URL |
|  ∟ `md5` | `string` | Material file md5 | ### Data example
```json
{
    "scope":5,
    "country_id":"086",
    "district_id":"001",
    "building_id":"22",
    "floor_name":"4",
    "room_id":"6383786266263",
    "code":0,
    "msg":"success",
    "data": {
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
                    "md5": "md5"
                }
            ]
        }
    }
}
```
