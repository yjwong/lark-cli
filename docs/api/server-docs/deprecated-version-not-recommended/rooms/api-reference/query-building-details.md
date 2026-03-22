---
title: "Query building details"
url: "https://open.larksuite.com/document/ukTMukTMukTM/ukzNyUjL5cjM14SO3ITN"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/meeting_room/building/batch_get?building_ids=omb_8ec170b937536a5d87c23b418b83f9bb&building_ids=omb_38570e4f0fd9ecf15030d3cc8b388f3a&fields=*"
section: "Deprecated Version (Not Recommended)"
scopes:
  - "calendar:room:readonly"
updated: "1686900924000"
---

# Query building details

This API is used to obtain details for specified buildings. 

##  Request
| HTTP URL | https://open.larksuite.com/open-apis/meeting_room/building/batch_get?building_ids=omb_8ec170b937536a5d87c23b418b83f9bb&building_ids=omb_38570e4f0fd9ecf15030d3cc8b388f3a&fields=* |
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
| building_ids | array       | Yes       | Queries the  ID                                       of specified building|
| fields       | string       | No       | Specifies the name of returned field with "," separated between fields, such as "id,name". "*"  indicates to return all fields. The optional fields are: "id,name,description,floors". All fields will be returned by default. | ##  Response

### Response body

| **Parameter**     | **Description**                                             |
| ------------ | ---------------------------------------------------- |
| code         | Return code. A value other than 0 indicates failure.                                |
| msg          | Description of the return code. "success" indicates success, and others are error messages. |
| data         | Returned business information                                         |
| ∟buildings   | List of buildings                                             |
|    ∟building_id | Building ID                                            |
|    ∟description | Description of the building                                     |
|    ∟floors      | List of all floors in current building                         |
|    ∟name        | Building name                                           | |    ∟country_id        | Country  ID                                           |
|    ∟district_id        | City  ID                                           | ### Response body example

```json
{
    "code":0,
    "msg":"success",
    "data":{
        "buildings":[
            {
                "building_id":"omb_8ec170b937536a5d87c23b418b83f9bb",
                "description":"Some description",
                "floors":[
                    "F1"
                ],
                "name":"Building name",
                "country_id": "Country id",
                "district_id": "District id"
            },
            {
                "building_id":"omb_38570e4f0fd9ecf15030d3cc8b388f3a",
                "description":"Some description",
                "floors":[
                    "F1",
                    "F2"
                ],
                "name":"Building name_2",
                "country_id": "Country id",
                "district_id": "District id"
            }
        ]
    }
}
```

### Error code

For details, refer to Server-side error codes.
