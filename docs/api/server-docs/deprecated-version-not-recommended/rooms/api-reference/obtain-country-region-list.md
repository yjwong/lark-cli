---
title: "Obtain Country/Region List"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uQTNwYjL0UDM24CN1AjN"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/meeting_room/country/list"
section: "Deprecated Version (Not Recommended)"
scopes:
  - "calendar:room:readonly"
updated: "1686900964000"
---

# Obtain the list of countries/regions

When creating a building, you need to define the country/region in which it will be located. This API is used to obtain the list of countries/regions available in the system. 

##  Request
| HTTP URL | https://open.larksuite.com/open-apis/meeting_room/country/list |
| --- | --- |
| HTTP Method | GET |
| Supported app types | custom,isv |
| Required scopes | `calendar:room:readonly` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" Learn more about how to obtain and use access_token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ##  Response

### Response body

| Parameter         | Description                                                 |
| ------------ | ---------------------------------------------------- |
| code         | Return code. A value other than 0 indicates failure.                                |
| msg          | Description of the return code. "success" indicates success, and others are error messages. |
| data         | Returned business information                                         |
| ∟countries   | List of countries/regions                                             |
|    ∟country_id | Country/region ID                                            |
|    ∟name | Country/region name                                   | ### Response body example
```json
{
    "code":0,
    "msg":"success",
    "data":{
        "countries":[
            {
                "country_id":"1814991",
                "name":"China"
            }
        ]
    }
}
```

### Error code

For details, refer to Server-side error codes.
