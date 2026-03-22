---
title: "Export meeting room reservation data"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/vc-v1/export/resource_reservation_list"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/vc/v1/exports/resource_reservation_list"
service: "vc-v1"
resource: "export"
section: "Video Conferencing"
rate_limit: "100 per minute"
scopes:
  - "vc:export"
updated: "1689326178000"
---

# Export meeting room reservation data

Export meeting room reservation data. For specific permission requirements, please refer to "Resource introduction".

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/vc/v1/exports/resource_reservation_list |
| HTTP Method | POST |
| Rate Limit | 100 per minute |
| Supported app types | custom,isv |
| Required scopes | `vc:export` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `room_level_id` | `string` | Yes | room level id **Example value**: "omm_608d34d82d531b27fa993902d350a307" |
| `need_topic` | `boolean` | No | Whether to show the conference theme **Example value**: true |
| `start_time` | `string` | Yes | Query start time (unix time, unit sec) **Example value**: "1655276858" |
| `end_time` | `string` | Yes | Query end time (unix time, unit sec) **Example value**: "1655276858" |
| `room_ids` | `string[]` | No | List of meeting room ids to filter **Example value**: ["omm_eada1d61a550955240c28757e7dec3af"] |
| `is_exclude` | `boolean` | No | If true, it means that the conference room outside the scope of the export room_ids, the default is false **Example value**: false | ### Request body example

{
    "room_level_id": "omm_608d34d82d531b27fa993902d350a307",
    "need_topic": true,
    "start_time": "1655276858",
    "end_time": "1655276858",
    "room_ids": [
        "omm_eada1d61a550955240c28757e7dec3af"
    ],
    "is_exclude": false
}

## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `task_id` | `string` | Task id | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
        "task_id": "7111325589855797267"
    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 500 | 121001 | internal error | Internal server error. If retry still doesn't work, contact the administrator. |
| 400 | 121002 | not support | This function is not supported. |
| 400 | 121003 | param error | Parameter error. Check the value range of parameters (please check yourself according to the field description above). |
| 404 | 121004 | data not exist | The requested data doesn't exist. |
| 403 | 121005 | no permission | No permission. Check the token type, operator identity and resource. |
