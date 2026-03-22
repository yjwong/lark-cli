---
title: "Query Meeting Room Details"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uEDOyUjLxgjM14SM4ITN"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/meeting_room/room/batch_get?room_ids=omm_eada1d61a550955240c28757e7dec3af&room_ids=omm_83d09ad4f6896e02029a6a075f71c9d1&fields=*"
section: "Deprecated Version (Not Recommended)"
scopes:
  - "calendar:room:readonly"
updated: "1686900931000"
---

# Query room details

This API is used to obtain details of specified rooms. 

##  Request
| HTTP URL | https://open.larksuite.com/open-apis/meeting_room/room/batch_get?room_ids=omm_eada1d61a550955240c28757e7dec3af&room_ids=omm_83d09ad4f6896e02029a6a075f71c9d1&fields=* |
| --- | --- |
| HTTP Method | GET |
| Supported app types | custom,isv |
| Required scopes | `calendar:room:readonly` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" Learn more about how to obtain and use access_token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Query parameters

| **Parameter** | **Parameter type** | **Required** | **Description**                                                     |
| -------- | ------------ | -------- | ------------------------------------------------------------ |
| room_ids | array        | Yes       | Queries the  ID                                       of specified room|
| fields   | string       | No       | Specifies the name of returned field with  ","  separated between fields, such as "id,name". "*" indicates to return all fields. The optional fields are:"id,name,description,capacity,building_id,building_name,floor_name,is_disabled,display_id". All fields will be returned by default. | ##  Response

### Response body

| **Parameter**       | **Description**                                                     |
| -------------- | ------------------------------------------------------------ |
| code           | Return code. A value other than 0 indicates failure.                                        |
| msg            | Description of the return code. "success" indicates success, and others are error messages.         |
| data           | Returned business information                                                 |
| ∟rooms         | List of rooms                                                   |
|    ∟room_id       | Rooms ID                                                    |
|    ∟building_id   | Building  ID                                          |of the room
|    ∟building_name | Building name of the room                                         |
|    ∟capacity      | Number of people that the room can hold                                           |
|    ∟description   | Description of the room                                             |
|    ∟display_id    | Display  ID                                              |of the room
|    ∟floor_name    | Floor name of the room                                           |
|    ∟is_disabled   | The room is available or not. If it is unavailable, the value is  true, otherwise the value is false .|
|    ∟name          | Room name                                                   | ### Response body example

```json
{
    "code":0,
    "msg":"success",
    "data":{
        "rooms":[
            {
                "room_id":"omm_eada1d61a550955240c28757e7dec3af",
                "building_id":"omb_8ec170b937536a5d87c23b418b83f9bb",
                "building_name":"Building name",
                "capacity":14,
                "description":"Some description",
                "display_id":"FM537532166",
                "floor_name":"F1",
                "is_disabled":false,
                "name":"Room name"
            },
            {
                "room_id":"omm_83d09ad4f6896e02029a6a075f71c9d1",
                "building_id":"omb_38570e4f0fd9ecf15030d3cc8b388f3a",
                "building_name":"Building name_2",
                "capacity":12,
                "description":"Some description",
                "display_id":"FM746810530",
                "floor_name":"F15",
                "is_disabled":false,
                "name":"Room name_2"
            }
        ]
    }
}
```

### Error code

For details, refer to Server-side error codes.
