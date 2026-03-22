---
title: "Obtain the App Availability in an Organization"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uIjM3UjLyIzN14iMycTN"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/application/v2/app/visibility"
section: "App Information"
scopes:
  - "admin:app.info:readonly"
updated: "1646720009000"
---

# Obtain app availability in a company

This API is used to query the availability of an app in a company. This API is only available to custom apps.

## Request
| HTTP URL | https://open.larksuite.com/open-apis/application/v2/app/visibility |
| --- | --- |
| HTTP Method | GET |
| Supported app types | custom |
| Required scopes | `admin:app.info:readonly` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" Learn more about how to obtain and use access_token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Request parameters
|Parameter|Type|Required|Description|
|--|-----|--|----|----|
|app_id|string|Yes|App  ID|Request  URL|
|user_page_token|string|No|Indicates the start position to pull users that can use the app by pages. The pulling starts from the beginning if this parameter is not specified.|
|user_page_size|int|No|The maximum number of users pulled this time (up to 1000. The value of 0 indicates 1000 users).| ## Response
### Response body

|Parameter|Description|
|--|--|
|code|Return code. A value other than  0  indicates failure.|
|msg|Return code description|
|data|Returned business information|
|  ∟departments|List of departments that can use the app|
|    ∟id|Custom  department_id|
|  ∟users|List of users who can use the app (Only users set individually are returned, and those in the departments and user groups within the availability are not returned.)|
|    ∟user_id|User's user_id. This is only returned for custom apps that have requested the user_id scope.|
|    ∟open_id|User's  open_id|
|  ∟is_visible_to_all|Whether the app is visible to all. 0: No; 1: Yes|
|  ∟has_more_users|Whether there are more users to which the app is visible. 1: Yes, 0: No|
|  ∟user_page_token| user_page_token  used to pull the next page of users| ### Response example

```json
{
    "code": 0,
    "msg": "success",
    "data": {
        "departments": [
            {
                "id": "od-aa2c50a04769feefededb7a05b7525a8"
            },
            {
                "id": "od-bdb63e2f2a339191d6fe51a2deb85fc8"
            }
        ],
        "users": [
            {
                "open_id": "ou_d317f090b7258ad0372aa53963cda70d",
                "user_id": "79affdge"
            },
            {
                "open_id": "ou_f0855b4d2a754a2116a0f723e95420ac",
                "user_id": "f11e4868"
            },
            {
                "open_id": "ou_42ebd86208405e7d4ba6fd36cece410a",
                "user_id": "4ea58f49"
            }
        ],
        "is_visible_to_all": 0,
        "has_more_users": 0,
        "user_page_token": "ou_42ebd86208405e7d4ba6fd36cece410a"
    }
}
```

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 200 | 50003 | invalid app_id | Check whether the app_id is correct and whether the app is installed for the company. |
