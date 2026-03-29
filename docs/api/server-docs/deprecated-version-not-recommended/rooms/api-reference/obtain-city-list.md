---
title: "Obtain City List"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uUTNwYjL1UDM24SN1AjN"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/meeting_room/district/list?country_id=1814991"
section: "Deprecated Version (Not Recommended)"
scopes:
  - "calendar:room:readonly"
updated: "1686900968000"
---

# Obtain the list of cities

When creating a building, you need to select the country/region in which it will be located. This API is used to obtain the list of cities available in the system. 

##  Request
| HTTP URL | https://open.larksuite.com/open-apis/meeting_room/district/list?country_id=1814991 |
| --- | --- |
| HTTP Method | GET |
| Supported app types | custom,isv |
| Required scopes | `calendar:room:readonly` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" Learn more about how to obtain and use access_token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Query parameters

| Parameter       | Parameter type | Required | Description                                                         |
| ---------- | -------- | ---- | ------------------------------------------------------------ |
| country_id  | int      | Yes   | Country/region ID | ##  Response

### Response body

| Parameter         | Parameter type | Description                                            |
| ------------ | -------- |---------------------------------------------------- |
| code         | |Return code. A value other than 0 indicates failure.                                |
| msg          |string |Description of the return code. "success" indicates success, and others are error messages. |
| data         | |Returned business information                                         |
| ∟districts   || List of cities                                             |
|    ∟district_id |string |CityID                                            |
|    ∟name |string |City name                                 | ### Response body example

```json
{
    "code":0,
    "msg":"success",
    "data":{
        "districts":[
            {
                "district_id":"1796236",
                "name":"Shanghai"
            }
        ]
    }
}
```

### Error code

For details, refer to Server-side error codes.
