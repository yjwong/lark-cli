---
title: "Obtain Meeting Room ID"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uYzMxYjL2MTM24iNzEjN"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/meeting_room/room/batch_get_id?custom_room_ids=test01&custom_room_ids=test02"
section: "Deprecated Version (Not Recommended)"
scopes:
  - "calendar:room:readonly"
updated: "1686900960000"
---

# Query room ID

This API is used to query the custom room IDs of the tenant.

##  Request
| HTTP URL | https://open.larksuite.com/open-apis/meeting_room/room/batch_get_id?custom_room_ids=test01&custom_room_ids=test02 |
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
| custom_room_ids | string       | Yes       | Queries the custom room ID      | of the tenant for specified room

##  Response

### Response body

| **Parameter**       | **Description**                                                     |
| -------------- | ------------------------------------------------------------ |
| code           | Return code. A value other than 0 indicates failure.                                        |
| msg            | Description of the return code. "success" indicates success, and others are error messages.         |
| data           | Returned business information                                                 |
| ∟rooms         | List of rooms                                                   |
|    ∟room_id       | Rooms ID                                                    |
|    ∟custom_room_id   | Custom room  ID                                          | of the tenant

### Response body example

```json
{
    "code":0,
    "msg":"success",
    "data":{
        "rooms":[
            {
                "room_id":"omm_eada1d61a550955240c28757e7dec3af",
                "custom_room_id":"test01"
            },
            {
                "room_id":"omm_83d09ad4f6896e02029a6a075f71c9d1",
                "custom_room_id":"test02"
            }
        ]
    }
}
```

### Error code

For details, refer to Server-side error codes.
