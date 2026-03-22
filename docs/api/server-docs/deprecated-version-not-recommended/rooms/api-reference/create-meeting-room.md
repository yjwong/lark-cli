---
title: "Create Meeting Room"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uITNwYjLyUDM24iM1AjN"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/meeting_room/room/create"
section: "Deprecated Version (Not Recommended)"
scopes:
  - "calendar:room"
updated: "1686900949000"
---

# Create room

This API is used to create rooms. 

##  Request
| HTTP URL | https://open.larksuite.com/open-apis/meeting_room/room/create |
| --- | --- |
| HTTP Method | POST |
| Supported app types | custom |
| Required scopes | `calendar:room` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" Learn more about how to obtain and use access_token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Request body
| Parameter       | Parameter type | Required | Description                                                         |
| ---------- | -------- | ---- | ------------------------------------------------------------ |
| building_id  | string      | Yes   | ID |of the building where the room is located
| floor  | string      | Yes   | Floor of the building where the room is located |
| name | string   | Yes   | Room name |
| capacity   | int   | Yes   | Capacity |
| is_disabled     | bool   | Yes   | Disabled or not |
| custom_room_id     | string   | No   |Custom room ID | of the tenant

### Request body example

```json
{
    "building_id":"omb_8d020b12fe49e82847c2af3c193d5754",
    "floor":"F1",
    "name":"Test room",
    "capacity":10,
    "is_disabled": false,
    "custom_room_id":"test_room_01"
}
```
##  Response
### Response body

| Parameter         | Description                                                 |
| ------------ | ---------------------------------------------------- |
| code         | Return code. A value other than 0 indicates failure.                                |
| msg          | Description of the return code. "success" indicates success, and others are error messages. |
| data         | Returned business information                                         |
| ∟room_id  | Successfully created room ID                          |
### Response body example

```json
{
    "code":0,
    "msg":"success",
    "data":{
        "room_id":"omb_8d020b12fe49e82847c2af3c193d5754",
    }
}
```

### Error code

For details, refer to Server-side error codes.
