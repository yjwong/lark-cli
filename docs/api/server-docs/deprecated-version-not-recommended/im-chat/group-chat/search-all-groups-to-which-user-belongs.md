---
title: "Search All Groups to Which User Belongs"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uUTOyUjL1kjM14SN5ITN"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/chat/v4/search?page_size=10&page_token=24&query=rd"
section: "Deprecated Version (Not Recommended)"
updated: "1626158304000"
---

# Search All Groups To Which User Belongs
Search all groups which are visiable to the user. These groups include this user or are public to this user.   

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/chat/v4/search?page_size=10&page_token=24&query=rd |
| HTTP Method | GET |
| Required scopes | read group information | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `user_access_token` **Value Format**: "Bearer `access_token`" **Example Value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" Learn more about obtaining and using access_token. |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Query parameters
|Parameter|Type|Mandatory|Description|Default|Demo|
-- | -- | -- | -- | -- | --
query|string|Yes|Keywords used for searching must not be empty.||rd|
page_size | string | No | The number of groups that we will get. The default is 100.|100|100|
page_token | string | No | Pagination markup. It is not filled in the first request, indicating traversal from the beginning; when there will be more groups, the new page_token will be returned at the same time, and the next traversal can use the page_token to get more groups.||6631829734835617796| ## Response
### Response body
|Parameters|Type|Description|
|-|-|-|
code |int| Error code, anything but 0 indicates failure
msg  |string| Error code description
data | - | -
  ∟has_more|bool|There are more groups if has_more is true.
  ∟page_token|string|See request parameter description for details. Valid when has_more is true.
  ∟groups|list|-
    ∟avatar|string| avatar
    ∟chat_id|string| Group ID
    ∟description|string| Group description
    ∟name |string| Group name
    ∟owner_user_id |string| Owner's user_id（If a bot is the group owner, this field does not exist.）
    ∟owner_open_id |string| Owner's open_id

### Response body example
```json
{
    "code": 0,
    "msg": "ok",
    "data": {
        "has_more": true,
        "page_token": "6631829734835617796",
        "groups": [
            {
                "avatar": "http://p2.pstatp.com/origin/78c100066939cf1374f1",
                "chat_id": "oc_41e7bdf4877cfc316136f4ccf6c32613",
                "description": "description 1",
                "name": "group 1",
                "owner_open_id": "ou_f407fcf504d40eac629a91740b4c8ce0",
                "owner_user_id": "deb1gcc7"
            },
            {
                "avatar": "http://p2.pstatp.com/origin/78bb00136a55b07eac95",
                "chat_id": "oc_e7081e51485dce75d6262ceede7a77b8",
                "description": "description 2",
                "name": "group 2",
                "owner_open_id": "ou_50c7e1b638610b781511fadf97118e4a",
                "owner_user_id": "5543fe1d"
            }
        ]
    }
}
```
### Error code

For details, please refer to: Service-side error codes
