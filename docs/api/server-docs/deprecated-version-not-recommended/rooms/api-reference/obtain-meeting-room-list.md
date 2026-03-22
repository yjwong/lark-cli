---
title: "Obtain meeting room list"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uADOyUjLwgjM14CM4ITN"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/meeting_room/room/list?building_id=omb_8ec170b937536a5d87c23b418b83f9bb&page_size=1&page_token=0&order_by=name-asc&fields=*"
section: "Deprecated Version (Not Recommended)"
scopes:
  - "calendar:room:readonly"
updated: "1686900928000"
---

# Obtain the list of rooms

This API is used to obtain the rooms of the specified building. 

##  Request
| HTTP URL | https://open.larksuite.com/open-apis/meeting_room/room/list?building_id=omb_8ec170b937536a5d87c23b418b83f9bb&page_size=1&page_token=0&order_by=name-asc&fields=* |
| --- | --- |
| HTTP Method | GET |
| Supported app types | custom,isv |
| Required scopes | `calendar:room:readonly` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" Learn more about how to obtain and use access_token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Query parameters

| **Parameter**    | **Parameter type** | **Required** | **Description**                                                     |
| ----------- | ------------ | -------- | ------------------------------------------------------------ |
| building_id | string       | Yes       | Building  ID                                            | to be queried
| page_size   | int          | No       | Number of rooms expected to be returned by the request. If the quantity is insufficient, all rooms will be returned. The default is 100 and the maximum is 1000. |
| page_token  | string       | No       | Paging token used to mark the current request. It will return page_size elements from the current paging token. |
| order_by    | string       | No       | Queries names/floors in ascending/descending order. The optional fields are: "name-asc,name-desc,floor_name-asc,floor_name-desc". No action will be performed if other strings are passed in. The default is unordered. |
| fields      | string       | No       | Specifies the name of returned field with  ","  separated between fields, such as "id,name". "*"  indicates to return all fields. The optional fields are: "id,name,description,capacity,building_id,building_name,floor_name,is_disabled,display_id". All fields will be returned by default. | ##  Response

### Response body

| **Parameter**       | **Description**                                                     |
| -------------- | ------------------------------------------------------------ |
| code           | Return code. A value other than 0 indicates failure.                                        |
| msg            | Description of the return code. "success" indicates success, and others are error messages.         |
| data           | Returned business information                                                 |
| ∟page_token    | Paging token, which returns when the next page exists                                   |
| ∟has_more      | When the next page exists, this value is true, otherwise it is  false                      .|
| ∟rooms         | List of rooms                                                   |
|    ∟room_id       | Rooms ID                                                    |
|    ∟building_id   | Building  ID                                          |of the room
|    ∟building_name | Building name of the room                                         |
|    ∟capacity      | Number of people that the room can hold                                           |
|    ∟description   | Description of the room                                             |
|    ∟display_id    | Display  ID                                              |of the room
|    ∟floor_name    | Floor name of the room                                           |
|    ∟is_disabled   | The room is available or not. If it is unavailable, the value is  true, otherwise the value is  false .|
|    ∟name          | Room name                                                   | ### Response body example

```json
{
    "code":0,
    "msg":"success",
    "data":{
        "page_token":"1",
        "has_more":true,
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
            }
        ]
    }
}
```

### Error code

For details, refer to Server-side error codes.
