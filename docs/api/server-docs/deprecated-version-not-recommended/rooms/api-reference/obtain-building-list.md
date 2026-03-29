---
title: "Obtain Building List"
url: "https://open.larksuite.com/document/ukTMukTMukTM/ugzNyUjL4cjM14CO3ITN"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/meeting_room/building/list?page_size=1&page_token=0&order_by=name-asc&fields=*"
section: "Deprecated Version (Not Recommended)"
scopes:
  - "calendar:room:readonly"
updated: "1686900920000"
---

#  Obtain the list of buildings

This API is used to obtain buildings (office buildings) of the company. 

##  Request
| HTTP URL | https://open.larksuite.com/open-apis/meeting_room/building/list?page_size=1&page_token=0&order_by=name-asc&fields=* |
| --- | --- |
| HTTP Method | GET |
| Supported app types | custom,isv |
| Required scopes | `calendar:room:readonly` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" Learn more about how to access and use access_token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Query parameters

| Parameter       | Parameter type | Required | Description                                                         |
| ---------- | -------- | ---- | ------------------------------------------------------------ |
| page_size  | int      | No   | Number of buildings expected to be returned by the request. If the quantity is insufficient, all buildings will be returned. The default is 10, and the maximum is 100. |
| page_token | string   | No   | Paging token used to mark the current request. It will return page_size elements from the current paging token. |
| order_by   | string   | No   | Queries names in ascending/descending order. The optional fields are: "name-asc,name-desc". No action will be performed if other strings are passed in. The default is unordered. |
| fields     | string   | No   | Specifies the name of returned field with "," separated between fields, such as "id,name". "*" indicates to return all fields. The optional fields are: "id,name,description,floors". All fields will be returned by default. | ##  Response

### Response body

| Parameter         | Description                                                 |
| ------------ | ---------------------------------------------------- |
| code         | Return code. A value other than 0 indicates failure.                                |
| msg          | Description of the return code. "success"  indicates success, and others are error messages. |
| data         | Returned business information                                         |
| ∟page_token  | Paging token, which returns when the next page exists                           |
| ∟has_more    | When the next page exists, this value is true, otherwise it is  false              .|
| ∟buildings   | List of buildings                                             |
|    ∟building_id | Building ID                                            |
|    ∟description | Description of the building                                     |
|    ∟floors      | List of all floors in current building                         |
|    ∟name        | Building name                                           |
|    ∟country_id        | Country  ID                                           |
|    ∟district_id        | City  ID                                           | ### Response body example

```json
{
    "code":0,
    "msg":"success",
    "data":{
        "page_token":"1",
        "has_more":true,
        "buildings":[
            {
                "building_id":"omb_8ec170b937536a5d87c23b418b83f9bb",
                "description":"Some description",
                "floors":[
                    "F1"
                ],
                "name":"Building name",
                "country_id": "country id",
                "district_id": "district id"
            }
        ]
    }
}
```

### Error code

For details, refer to Server-side error codes.
