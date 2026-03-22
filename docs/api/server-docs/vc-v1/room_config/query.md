---
title: "Query meeting room configuration"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/vc-v1/room_config/query"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/vc/v1/room_configs/query"
service: "vc-v1"
resource: "room_config"
section: "Deprecated Version (Not Recommended)"
scopes:
  - "vc:room:readonly"
field_scopes:
  - "contact:user.employee_id:readonly"
updated: "1686900968000"
---

# Query room configurations

Queries room configurations within a scope.

> Input corresponding parameters according to query range

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/vc/v1/room_configs/query |
| HTTP Method | GET |
| Supported app types | custom,isv |
| Required scopes | `vc:room:readonly` |
| Required field scopes | > The response body of the API contains the following sensitive fields, and they will be returned only after corresponding scopes are added. If you do not need the fields, it is not recommended that you request the scopes. `contact:user.employee_id:readonly` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `scope` | `int` | Yes | Query node scope **Example value**: 5 **Optional values are**:  - `1`: Tenant - `2`: Country/region - `3`: City - `4`: Building - `5`: Floor - `6`: Room  |
| `country_id` | `string` | No | The parameter is required when the country/region ID scope is 2 or 3 **Example value**: "086" |
| `district_id` | `string` | No | The parameter is required when the city ID scope is 3 **Example value**: "001" |
| `building_id` | `string` | No | The parameter is required when the building ID scope is 4 or 5 **Example value**: "22" |
| `floor_name` | `string` | No | The parameter is required when the floor scope is 5 **Example value**: "4" |
| `room_id` | `string` | No | The parameter is required when the room ID scope is 6 **Example value**: "6383786266263" |
| `user_id_type` | `string` | No | User ID categories **Example value**: "open_id" **Optional values are**:  - `open_id`: Identifies a user to an app. The same user has different Open IDs in different apps. How to get Open ID - `union_id`: Identifies a user to a tenant that acts as a developer. A user has the same Union ID in apps developed by the same developer, and has different Union IDs in apps developed by different developers. A developer can use Union ID to link the same user's identities in multiple apps.How to get Union ID - `user_id`: Identifies a user to a tenant. The same user has different User IDs in different tenants. In one single tenant, a user has the same User ID in all apps （including store apps）. User ID is usually used to communicate user data between different apps. How to get User ID  **Default value**: `open_id` **When the value is `user_id`, the following field scopes are required**: `contact:user.employee_id:readonly` | ## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `room_config` | \- |
|   `room_background` | `string` | Lark room background |
|   `display_background` | `string` | Lark sign board background |
|   `digital_signage` | `room_digital_signage` | Lark room digital labels |
|     `enable` | `boolean` | Enable digital label function or not |
|     `mute` | `boolean` | Whether to mute the video |
|     `start_display` | `int` | Start playing n min after the meeting ends，should be between 1~720.(Only available for Lark room digital labels) |
|     `stop_display` | `int` | End playing n min before the meeting starts, should be between 1~720.(Only available for Lark room digital labels) |
|     `materials` | `room_digital_signage_material[]` | Material list |
|       `id` | `string` | Material ID. When uploading a new material, the material ID needs to be set to empty. |
|       `name` | `string` | Material name |
|       `material_type` | `int` | Material type **Optional values are**:  - `1`: Image - `2`: Video - `3`: GIF  |
|       `url` | `string` | Material URL |
|       `duration` | `int` | Play duration (in sec), should be between 1~43200. |
|       `cover` | `string` | Material cover URL |
|       `md5` | `string` | Material file md5 |
|       `vid` | `string` | Material file vid |
|       `size` | `string` | Material desc |
|   `room_box_digital_signage` | `room_digital_signage` | Lark room box digital signage |
|     `enable` | `boolean` | Enable digital label function or not |
|     `mute` | `boolean` | Whether to mute the video |
|     `start_display` | `int` | Start playing n min after the meeting ends，should be between 1~720.(Only available for Lark room digital labels) |
|     `stop_display` | `int` | End playing n min before the meeting starts, should be between 1~720.(Only available for Lark room digital labels) |
|     `materials` | `room_digital_signage_material[]` | Material list |
|       `id` | `string` | Material ID. When uploading a new material, the material ID needs to be set to empty. |
|       `name` | `string` | Material name |
|       `material_type` | `int` | Material type **Optional values are**:  - `1`: Image - `2`: Video - `3`: GIF  |
|       `url` | `string` | Material URL |
|       `duration` | `int` | Play duration (in sec), should be between 1~43200. |
|       `cover` | `string` | Material cover URL |
|       `md5` | `string` | Material file md5 |
|       `vid` | `string` | Material file vid |
|       `size` | `string` | Material desc |
|   `room_status` | `room_status` | Room status |
|     `status` | `boolean` | Available status of room |
|     `schedule_status` | `boolean` | Future available status of room |
|     `disable_start_time` | `string` | Room disable start time |
|     `disable_end_time` | `string` | Room disable end time |
|     `disable_reason` | `string` | Room disable reason |
|     `contact_ids` | `string[]` | Contact user ID |
|     `disable_notice` | `boolean` | Notice content sent to users when room disabled |
|     `resume_notice` | `boolean` | Notice content sent to contact users when room resumed | ### Response body example

{
    "code": 0,
    "msg": "success",
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

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 500 | 121001 | internal error | Internal server error. If retry still doesn't work, contact the administrator. |
| 400 | 121002 | not support | This function is not supported. |
| 400 | 121003 | param error | Parameter error. Check the value range of parameters (please check yourself according to the field description above). |
| 404 | 121004 | data not exist | The requested data doesn't exist. |
