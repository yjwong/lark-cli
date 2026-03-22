---
title: "Get User IDs Using Email Addresses or Mobile Numbers"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uUzMyUjL1MjM14SNzITN"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/user/v1/batch_get_id?emails=lisi@z.com&emails=wangwu@z.com&mobiles=13812345678&mobiles=%2b12126668888"
section: "Deprecated Version (Not Recommended)"
updated: "1646999075000"
---

# Get user IDs using email addresses or mobile numbers

Query a user’s open_id and user_id using the user’s email address or mobile number. Batch queries are supported.
> The `Obtain user ID via email or mobile number` permission must have been granted to invoke this API.Only the IDs of users who have permission to access the app could be queried.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/user/v1/batch_get_id?emails=lisi@z.com&emails=wangwu@z.com&mobiles=13812345678&mobiles=%2b12126668888 |
| HTTP Method | GET |
| Required scopes | Obtain user ID via email or mobile number |
| Required field scopes Enable the scopes for fields you want to return | Obtain user ID Obtain user's mobile number Obtain user's email information | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value Format**: "Bearer `access_token`" **Example Value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" Learn more about obtaining and using access_token. |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Query parameters
Parameter | Type | Mandatory | Example | Description
-- | -- | -- | -- | --
emails | string | Optional | lisi@z.com | The email addresses of the users to be queried, max. 50.
mobiles | string | Optional | 13812345678%2b12126668888 | The mobile numbers of the users to be queried, max. 50.A "+" symbol before country calling code is required if the mobile number is not in mainland China, and it need to be URL encoded.
## Response
### Response body
Parameter | Description
-- | --
code | Error code, anything but 0 indicates failure.
msg | Response code description.
data | -
  ∟email_users | The users found based on the email addresses. The key is email address and the value is users found with that address.Pay attention that multiple users may be found for the same email address.
    ∟open_id | User's open_id. open_id Description
    ∟user_id | User's user_id, only returned for Custom Apps who has permission of `Obtain users' userid`. user_id Description
  ∟emails_not_exist | Email addresses not found.
  ∟mobile_users | The users found based on the mobile numbers. The key is mobile numbers and the value is users found with that mobile number.Pay attention that multiple users may be found for the same mobile number.
  ∟mobiles_not_exist | Mobile numbers not found.

### Response body example
```json
{
    "code": 0,
    "msg": "success",
    "data": {
        "email_users": {
            "lisi@z.com": [
                {
                    "open_id": "ou_979112345678741d29069abcdef089d4",
                    "user_id": "a7eb3abe"
                },
                {
                    "open_id": "ou_1725134fb47412345678dd40589abc",
                    "user_id": "38e7d231"
                }
            ]
        },
        "emails_not_exist": [
            "wangwu@z.com"
        ],
        "mobile_users": {
            "13812345678": [
                {
                    "open_id": "ou_46a087654321a1dc920ffab8fedc823f",
                    "user_id": "18fg19d4"
                }
            ]
        },
        "mobiles_not_exist": [
            "13912345678",
            "+12126668888"
        ]
    }
}
```
### Error code

For details, please refer to: Service-side error codes
