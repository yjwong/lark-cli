---
title: "Create Building"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uATNwYjLwUDM24CM1AjN"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/meeting_room/building/create"
section: "Deprecated Version (Not Recommended)"
scopes:
  - "calendar:room"
updated: "1686900935000"
---

# Create building

This API is used to create building and building floor information, which is equivalent to the functions of adding buildings and floors in the admin console. 

##  Request
| HTTP URL | https://open.larksuite.com/open-apis/meeting_room/building/create |
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
| name  | string      | Yes   | Building name |
| floors | string[]   | Yes   | List of floors |
| country_id   | string   | Yes   | Country/region ID |
| district_id     | string   | Yes   | City ID |
| custom_building_id     | string   | No   | Custom building ID  of the tenant|
### Request body example

```json
{
    "name":"Test building",
    "floors":[
        "F1",
        "F2",
        "F3",
        "F4"
    ],
    "country_id":"1814991",
    "district_id":"2034437",
    "custom_building_id":"test_building_01"
}
```

##  Response
### Response body

| Parameter         | Description                                                 |
| ------------ | ---------------------------------------------------- |
| code         | Return code. A value other than 0 indicates failure.                                |
| msg          | Description of the return code. "success"  indicates success, and others are error messages. |
| data         | Returned business information                                         |
| ∟building_id  | Successfully created building ID                          | ### Response body example

```json
{
    "code":0,
    "msg":"success",
    "data":{
        "building_id":"omb_8ec170b937536a5d87c23b418b83f9bb",
    }
}
```

### Error code

For details, refer to Server-side error codes.
