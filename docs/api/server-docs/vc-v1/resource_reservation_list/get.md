---
title: "Get meeting room reservation data"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/vc-v1/resource_reservation_list/get"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/vc/v1/resource_reservation_list"
service: "vc-v1"
resource: "resource_reservation_list"
section: "Video Conferencing"
rate_limit: "100 per minute"
scopes:
  - "vc:room:readonly"
updated: "1690944921000"
---

# Get meeting room reservation data

Get meeting room reservation data.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/vc/v1/resource_reservation_list |
| HTTP Method | GET |
| Rate Limit | 100 per minute |
| Supported app types | custom,isv |
| Required scopes | `vc:room:readonly` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `room_level_id` | `string` | Yes | Room level ID. **Example value**: omb_57c9cc7d9a81e27e54c8fabfd02759e7 |
| `need_topic` | `boolean` | No | Whether the response returns meeting topic. **Example value**: true |
| `start_time` | `string` | Yes | Start time **Example value**: 1655276858 |
| `end_time` | `string` | Yes | End time **Example value**: 1655276858 |
| `room_ids` | `string[]` | Yes | List of meeting room IDs **Example value**: omm_eada1d61a550955240c28757e7dec3af |
| `is_exclude` | `boolean` | No | Whether exclude the list of meeting room IDs **Example value**: false |
| `page_size` | `int` | No | **Example value**: 20 **Default value**: `20` **Data validation rules**: - Value range: `20` ～ `100` |
| `page_token` | `string` | No | Page identifier. It is not filled in the first request, indicating traversal from the beginning; when there will be more groups, the new page_token will be returned at the same time, and the next traversal can use the page_token to get more groups **Example value**: 20 | ## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `room_reservation_list` | `room_meeting_reservation[]` | Meeting room reservation list |
|     `room_id` | `string` | Conference room id |
|     `room_name` | `string` | Conference room name |
|     `event_title` | `string` | Conference theme |
|     `reserver` | `string` | Reserve person |
|     `department_of_reserver` | `string` | The department to which the reserver belongs |
|     `guests_number` | `string` | Number of people invited |
|     `accepted_number` | `string` | Number of recipients |
|     `event_start_time` | `string` | Meeting start time |
|     `event_end_time` | `string` | Meeting end time |
|     `event_duration` | `string` | Duration of the meeting |
|     `reservation_status` | `string` | Meeting room reservation status |
|     `check_in_device` | `string` | Check in device |
|     `room_check_in_status` | `string` | Meeting room check-in status |
|     `check_in_time` | `string` | Meeting room check-in time |
|     `is_release_early` | `string` | Whether to release early |
|     `releasing_person` | `string` | Release person |
|     `releasing_time` | `string` | Release time |
|   `page_token` | `string` | Page identifier, when has_more is true, a new page_token will also be returned. Otherwise, page_token will not be returned |
|   `has_more` | `boolean` | Whether the response body has more parameters | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
        "room_reservation_list": [
            {
                "room_id": "omm_4de32cf10a4358788ff4e09e37ebbf9b",
                "room_name": "VIP Meeting Room",
                "event_title": "test topic",
                "reserver": "kehan",
                "department_of_reserver": "development",
                "guests_number": "5",
                "accepted_number": "2",
                "event_start_time": "2022.12.17 21:00:00 (GMT+08:00)",
                "event_end_time": "2022.12.17 22:00:00 (GMT+08:00)",
                "event_duration": "1:00:00",
                "reservation_status": "reserve success",
                "check_in_device": "display",
                "room_check_in_status": "checkin",
                "check_in_time": "2022.12.09 13:35:30 (GMT+08:00)",
                "is_release_early": "released",
                "releasing_person": "kehan",
                "releasing_time": "2022.12.20 11:25:15 (GMT+08:00)"
            }
        ],
        "page_token": "20",
        "has_more": true
    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 500 | 121001 | internal error | Internal server error. If retry still doesn't work, contact the administrator. |
| 400 | 121002 | not support | This function is not supported. |
| 400 | 121003 | param error | Parameter error. Check the value range of parameters (please check yourself according to the field description above). |
| 400 | 126005 | start time error (cannot larger than now or smaller than 0) | The start time cannot be greater than the current time or less than 0 |
| 400 | 126006 | end time cannot exceed start time for one day or smaller than start time | The end time cannot exceed the start time by one day or be less than the start time |
| 404 | 121004 | data not exist | The requested data doesn't exist. |
| 403 | 121005 | no permission | No permission. Check the token type, operator identity and resource. |
