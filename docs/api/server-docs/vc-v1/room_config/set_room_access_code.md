---
title: "set_room_access_code"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/vc-v1/room_config/set_room_access_code"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/vc/v1/room_configs/set_room_access_code"
service: "vc-v1"
resource: "room_config"
section: "Deprecated Version (Not Recommended)"
scopes:
  - "vc:room"
updated: "1686900968000"
---

# Create meeting room access code

Create a range of meeting room access codes

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/vc/v1/room_configs/set_room_access_code |
| HTTP Method | POST |
| Supported app types | custom,isv |
| Required scopes | `vc:room` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `scope` | `int` | Yes | Set node range **Example value**: 5 **Optional values are**:  - `1`: Tenant - `2`: Country/region - `3`: City - `4`: Building - `5`: Floor - `6`: Room  |
| `country_id` | `string` | No | This parameter is required when the country ID scope is 2, 3 **Example value**: "1" |
| `district_id` | `string` | No | This parameter is required when the city ID scope is 3 **Example value**: "2" |
| `building_id` | `string` | No | This parameter is required when the building ID scope is 4, 5 **Example value**: "3" |
| `floor_name` | `string` | No | This parameter is required when the floor scope is 5 **Example value**: "4" |
| `room_id` | `string` | No | This parameter is required when the conference room ID scope is 6 **Example value**: "67687262867363" |
| `valid_day` | `int` | Yes | Effective days **Example value**: 1 **Optional values are**:  - `1`: 1 day - `7`: 7 days - `30`: 30 days  | ### Request body example

{
    "scope": 5,
    "country_id": "1",
    "district_id": "2",
    "building_id": "3",
    "floor_name": "4",
    "room_id": "67687262867363",
    "valid_day": 1
}

## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `access_code` | `string` | access code | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
        "access_code": "YNMLTJOUZ"
    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 500 | 121001 | internal error | Internal server error. If retry still doesn't work, contact the administrator. |
| 400 | 121002 | not support | This function is not supported. |
| 400 | 121003 | param error | Parameter error. Check the value range of parameters (please check yourself according to the field description above). |
| 404 | 121004 | data not exist | The requested data doesn't exist. |
