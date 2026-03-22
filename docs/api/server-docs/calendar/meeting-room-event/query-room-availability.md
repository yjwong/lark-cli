---
title: "Query room availability"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uIDOyUjLygjM14iM4ITN"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/meeting_room/freebusy/batch_get"
section: "Calendar"
scopes:
  - "calendar:room:readonly"
updated: "1736835306000"
---

# Query room availability

Call this interface to get the busy and free event information of the specified meeting room.

> **Note:** In the query results:
> - Non-repeating events only have unique instance information.
> - Repeating events may have multiple instance information, which will be expanded based on event recurrence rules and time ranges. The recommended query time range is within 30 days.

##  Request
| HTTP URL | https://open.larksuite.com/open-apis/meeting_room/freebusy/batch_get |
| --- | --- |
| HTTP Method | GET |
| Supported app types | custom,isv |
| Required scopes | `calendar:room:readonly` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" Learn more about how to obtain and use access_token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Query parameters
| **Parameter** | **Parameter type** | **Required** | **Description**                                                     |
| -------- | ------------ | -------- | ------------------------------------------------------------ |
| room_ids | array | Yes | Room ID. You can query conference room list or search conference room interface to obtain the specified conference room ID. 1. The number of room_ids should not exceed 20. 2. An example format for passing in multiple room IDs in a GET request is `room_ids=omm_83d09ad4f6896e02029a6a075f71xxxx&room_ids=omm_eada1d61a550955240c28757e7dexxxx`.|
| time_min | string | Yes | The starting time of the query needs to follow the [RFC3339](https://tools.ietf.org/html/rfc3339) format, example: 2019-09-04T08:45:00+08:00. **Note**: URL encoding is required when passing in this parameter. |
| time_max | string | Yes | The end time of the query needs to follow the [RFC3339](https://tools.ietf.org/html/rfc3339) format, example: 2019-09-04T09:45:00+08:00. **Note**: URL encoding is required when passing in this parameter. | ##  Response
### Response body

| **Parameters** |**Type**| **Description** |
| ----------- |-----| ------------------------- |
| code |int| Return code, non-0 indicates failure. |
| msg |string| Description of the return code, `success` means success, others are error messages. |
| data |-| Return information. |
| ∟time_min |string|Query the starting time of conference room busy time. |
| ∟time_max |string| Query the busy end time of the conference room. |
| ∟free_busy |-| Conference room busy information list. |
|    ∟room_id | string| Meeting room ID. |
|       ∟start_time |string| The starting time of the busy event. |
|       ∟end_time |string| The end time of the busy event. |
|       ∟uid |string| The unique ID of the event. |
|       ∟original_time |int| The original time of the event instance. For non-recurring events and recurring events, here is 0; for exception events for recurring events, here is the corresponding original_time value (timestamp type). |
|       ∟organizer_info |-| Event organizer information. Private events do not return this information. |
|         ∟name |string| Name of the organizer. |
|         ∟open_id |string| The open_id of the organizer. | ### Response body example
```json
{
    "code":0,
    "msg":"success",
    "data":{
        "time_max":"2019-09-04T09:45:00+08:00",
        "time_min":"2019-09-04T08:45:00+08:00",
        "free_busy":{
            "omm_83d09ad4f6896e02029a6a075f71c9d1":[
                {
                    "end_time":"2019-09-04T09:30:00+08:00",
                    "organizer_info": {
                        "name": "John",
                        "open_id": "ou_09379914ab6dde664x9daf84d8b13842"
                    },
                    "start_time":"2019-09-04T09:00:00+08:00",
                    "original_time": 0,
                    "uid": "bff6b51f-b7c1-40c6-b8ef-aef966c9ffc7"
                }
            ],
            "omm_c7e9ef91ec9f3c007f1dfb2503918981": [
                {
                    "end_time": "2020-01-14T19:00:00+08:00",
                    "organizer_info": {
                        "name": "xxx",
                        "open_id": "ou_09379914ab6dde66499daf94d8b13842"
                    },
                    "original_time": 1578997800,
                    "start_time": "2020-01-14T18:30:00+08:00",
                    "uid": "5632b176-4f0e-44eb-8331-c4b155ed270f"
                }
           ]
        }
    }
}
```

### Error code

You can go to General Error Code.
