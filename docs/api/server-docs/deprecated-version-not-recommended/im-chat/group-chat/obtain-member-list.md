---
title: "Obtain Member List"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uUzMwUjL1MDM14SNzATN"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/chat/v4/members?chat_id=oc_92c3f700c2ae31369cefee459fb93870&page_token=0&page_size=3"
section: "Deprecated Version (Not Recommended)"
updated: "1626158301000"
---

# Obtain Member List
If the user belongs to the group, the list of all the members of the group will be returned.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/chat/v4/members?chat_id=oc_92c3f700c2ae31369cefee459fb93870&page_token=0&page_size=3 |
| HTTP Method | GET |
| Required scopes | Read group messages(Historic Version) | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `user_access_token` **Value Format**: "Bearer `access_token`" **Example Value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" Learn more about obtaining and using access_token. |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Query parameters
|Parameter|Type|Mandatory|Description|Default|Demo|
-- | -- | -- | -- | -- | --
chat_id| string| Yes | chat ID ||oc_92c3f700c2ae31369cefee459fb93870|
page_size | string | No | The number of members that we will get. The default is 100. The max is 200. Page_size is only a rough number, the actual number of group members returned may be greater or less than page_zise. |100|3|
page_token | string | No | Pagination markup. It is not filled in the first request, indicating traversal from the beginning; when there will be more members, the new page_token will be returned at the same time, and the next traversal can use the page_token to get more members.||0| ## Response

### Response body
|Parameters|Type|Description|
|-|-|-|
code |int| Error code, anything but 0 indicates failure
msg  |string| Error code description
data | - | -
  ∟chat_id|string| Group ID
  ∟has_more|bool|There are more members if has_more is true.
  ∟page_token|string|See request parameter description for details.
  ∟members|list| Member list 
    ∟name|string| name of the user
    ∟open_id|string| open_id of the user
    ∟user_id|string| user_id of the user (only available in Enterprise self-built application)

### Response body example
```json
{
    "code": 0,
    "data": {
        "chat_id": "oc_92c3f700c2ae31369cefee459fb93870",
        "has_more": true,
        "members": [
            {
                "open_id": "ou_56799ac95e82434b49e1cf00c3a3a251",
                "user_id": "1g6gbf73",
                "name": "Zhang San"
            },
            {
                "open_id": "ou_9c7a2ce4f61e78dfe00ffa8b11524e2a",
                "user_id": "296f8dfb",
                "name": "Li Si"
            },
            {
                "open_id": "ou_fdeaf10d447a41d0a8d561454da197c9",
                "user_id": "afc6b4fb",
                "name": "Wang Wu"
            },
            {
                "open_id": "ou_7f8d47e1a788345c0d167abfbf3835b4",
                "user_id": "84eeea8d",
                "name": "Zhao Liu"
            }
        ],
        "page_token": "1559288627"
    },
    "msg": "ok"
}
```

### Error code

For details, please refer to: Service-side error codes
