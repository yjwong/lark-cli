---
title: "Set up meeting room configuration"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/vc-v1/room_config/set"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/vc/v1/room_configs/set"
service: "vc-v1"
resource: "room_config"
section: "Deprecated Version (Not Recommended)"
scopes:
  - "vc:room"
field_scopes:
  - "contact:user.employee_id:readonly"
updated: "1686900968000"
---

# Set room configurations

Sets room configurations within a scope.

> Input corresponding parameters according to setting scope

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/vc/v1/room_configs/set |
| HTTP Method | POST |
| Supported app types | custom,isv |
| Required scopes | `vc:room` |
| Required field scopes | > The response body of the API contains the following sensitive fields, and they will be returned only after corresponding scopes are added. If you do not need the fields, it is not recommended that you request the scopes. `contact:user.employee_id:readonly` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `user_id_type` | `string` | No | User ID categories **Example value**: "open_id" **Optional values are**:  - `open_id`: Identifies a user to an app. The same user has different Open IDs in different apps. How to get Open ID - `union_id`: Identifies a user to a tenant that acts as a developer. A user has the same Union ID in apps developed by the same developer, and has different Union IDs in apps developed by different developers. A developer can use Union ID to link the same user's identities in multiple apps.How to get Union ID - `user_id`: Identifies a user to a tenant. The same user has different User IDs in different tenants. In one single tenant, a user has the same User ID in all apps （including store apps）. User ID is usually used to communicate user data between different apps. How to get User ID  **Default value**: `open_id` **When the value is `user_id`, the following field scopes are required**: `contact:user.employee_id:readonly` | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `scope` | `int` | Yes | Set node scope **Example value**: 5 **Optional values are**:  - `1`: Tenant - `2`: Country/region - `3`: City - `4`: Building - `5`: Floor - `6`: Room  |
| `country_id` | `string` | No | The parameter is required when the country/region ID scope is 2 or 3 **Example value**: "086" |
| `district_id` | `string` | No | The parameter is required when the city ID scope is 3 **Example value**: "223" |
| `building_id` | `string` | No | The parameter is required when the building ID scope is 4 or 5 **Example value**: "66" |
| `floor_name` | `string` | No | The parameter is required when the floor scope is 5 **Example value**: "3" |
| `room_id` | `string` | No | The parameter is required when the room ID scope is 6 **Example value**: "67687262867363" |
| `room_config` | `room_config` | Yes | Room settings |
|   `room_background` | `string` | No | Lark room background **Example value**: "https://lf1-ttcdn-tos.pstatp.com/obj/xxx" |
|   `display_background` | `string` | No | Lark sign board background **Example value**: "https://lf1-ttcdn-tos.pstatp.com/obj/xxx" |
|   `digital_signage` | `room_digital_signage` | No | Lark room digital labels |
|     `enable` | `boolean` | No | Enable digital label function or not **Example value**: true |
|     `mute` | `boolean` | No | Whether to mute the video **Example value**: true |
|     `start_display` | `int` | No | Start playing n min after the meeting ends，should be between 1~720.(Only available for Lark room digital labels) **Example value**: 3 |
|     `stop_display` | `int` | No | End playing n min before the meeting starts, should be between 1~720.(Only available for Lark room digital labels) **Example value**: 3 |
|     `materials` | `room_digital_signage_material[]` | No | Material list |
|       `id` | `string` | No | Material ID. When uploading a new material, the material ID needs to be set to empty. **Example value**: "7847784676276" |
|       `name` | `string` | No | Material name **Example value**: "name" |
|       `material_type` | `int` | No | Material type **Example value**: 0 **Optional values are**:  - `1`: Image - `2`: Video - `3`: GIF  |
|       `url` | `string` | No | Material URL **Example value**: "url" |
|       `duration` | `int` | No | Play duration (in sec), should be between 1~43200. **Example value**: 15 |
|       `cover` | `string` | No | Material cover URL **Example value**: "url" |
|       `md5` | `string` | No | Material file md5 **Example value**: "md5" |
|       `vid` | `string` | No | Material file vid **Example value**: "vid" |
|       `size` | `string` | No | Material desc **Example value**: "100" |
|   `room_box_digital_signage` | `room_digital_signage` | No | Lark room box digital signage |
|     `enable` | `boolean` | No | Enable digital label function or not **Example value**: true |
|     `mute` | `boolean` | No | Whether to mute the video **Example value**: true |
|     `start_display` | `int` | No | Start playing n min after the meeting ends，should be between 1~720.(Only available for Lark room digital labels) **Example value**: 3 |
|     `stop_display` | `int` | No | End playing n min before the meeting starts, should be between 1~720.(Only available for Lark room digital labels) **Example value**: 3 |
|     `materials` | `room_digital_signage_material[]` | No | Material list |
|       `id` | `string` | No | Material ID. When uploading a new material, the material ID needs to be set to empty. **Example value**: "7847784676276" |
|       `name` | `string` | No | Material name **Example value**: "name" |
|       `material_type` | `int` | No | Material type **Example value**: 0 **Optional values are**:  - `1`: Image - `2`: Video - `3`: GIF  |
|       `url` | `string` | No | Material URL **Example value**: "url" |
|       `duration` | `int` | No | Play duration (in sec), should be between 1~43200. **Example value**: 15 |
|       `cover` | `string` | No | Material cover URL **Example value**: "url" |
|       `md5` | `string` | No | Material file md5 **Example value**: "md5" |
|       `vid` | `string` | No | Material file vid **Example value**: "vid" |
|       `size` | `string` | No | Material desc **Example value**: "100" |
|   `room_status` | `room_status` | No | Room status |
|     `status` | `boolean` | Yes | Available status of room **Example value**: true |
|     `disable_start_time` | `string` | No | Room disable start time **Example value**: "1652356050" |
|     `disable_end_time` | `string` | No | Room disable end time **Example value**: "1652442450" |
|     `disable_reason` | `string` | No | Room disable reason **Example value**: "TestReason" |
|     `contact_ids` | `string[]` | No | Contact user ID **Example value**: ["ou_3ec3f6a28a0d08c45d895276e8e5e19b"] |
|     `disable_notice` | `boolean` | No | Notice content sent to users when room disabled **Example value**: true |
|     `resume_notice` | `boolean` | No | Notice content sent to contact users when room resumed **Example value**: true | ### Request body example

{
    "scope": 5,
    "country_id": "086",
    "district_id": "223",
    "building_id": "66",
    "floor_name": "3",
    "room_id": "67687262867363",
    "room_config": {
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
    "data": {

    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 500 | 121001 | internal error | Internal server error. If retry still doesn't work, contact the administrator. |
| 400 | 121002 | not support | This function is not supported. |
| 400 | 121003 | param error | Parameter error. Check the value range of parameters (please check yourself according to the field description above). |
| 404 | 121004 | data not exist | The requested data doesn't exist. |
