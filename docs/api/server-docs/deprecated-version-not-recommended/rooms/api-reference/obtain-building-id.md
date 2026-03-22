---
title: "Obtain Building ID"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uQzMxYjL0MTM24CNzEjN"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/meeting_room/building/batch_get_id?custom_building_ids=test01&custom_building_ids=test02"
section: "Deprecated Version (Not Recommended)"
scopes:
  - "calendar:room:readonly"
updated: "1686900946000"
---

# Query building ID

This API is used to query the custom building IDs of the tenant. 

##  Request
| HTTP URL | https://open.larksuite.com/open-apis/meeting_room/building/batch_get_id?custom_building_ids=test01&custom_building_ids=test02 |
| --- | --- |
| HTTP Method | GET |
| Supported app types | custom,isv |
| Required scopes | `calendar:room:readonly` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" Learn more about how to obtain and use access_token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Query parameters

| **Parameter**     | **Parameter type** | **Required** | **Description**                                                     |
| ------------ | ------------ | -------- | ------------------------------------------------------------ |
| custom_building_ids | string       | Yes       | Queries the custom building ID | of the tenant for specified building

##  Response

### Response body

| **Parameter**     | **Description**                                             |
| ------------ | ---------------------------------------------------- |
| code         | Return code. A value other than 0 indicates failure.                                |
| msg          | Description of the return code. "success" indicates success, and others are error messages. |
| data         | Returned business information                                         |
| ∟buildings   | List of buildings                                             |
|    ∟building_id | BuildingID                                            |
|    ∟custom_building_id | Custom building ID              | of the tenant

### Response body example

```json
{
    "code":0,
    "msg":"success",
    "data":{
        "buildings":[
            {
                "building_id":"omb_8ec170b937536a5d87c23b418b83f9bb",
                "custom_bulding_id":"test01"
            },
            {
                "building_id":"omb_38570e4f0fd9ecf15030d3cc8b388f3a",
                "custom_bulding_id":"test02"
            }
        ]
    }
}
```

### Error code

For details, refer to Server-side error codes.
