---
title: "Obtain Group List to Which Bot Belongs"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uITO5QjLykTO04iM5kDN"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/chat/v4/list?page_size=20&page_token=1530027865231834600"
section: "Deprecated Version (Not Recommended)"
updated: "1626158303000"
---

# Obtain Group List
Obtain a list of groups where the bot exists.

> Bot capability needs to be enabled. Bot must be in the chat.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/chat/v4/list?page_size=20&page_token=1530027865231834600 |
| HTTP Method | GET | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value Format**: "Bearer `access_token`" **Example Value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" Learn more about obtaining and using access_token. |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Query parameters
|Parameter|Type|Mandatory|Description|Default|Demo|
-- | -- | -- | -- | -- | --
page_size | int | No | The number of groups that we will get. The default is 100. max is 200.|100|100|
page_token | string | No | Pagination markup. It is not filled in the first request, indicating traversal from the beginning; when there will be more groups, the new page_token will be returned at the same time, and the next traversal can use the page_token to get more groups.||1530027865231834600| ## Response

### Response body
|Parameters|Type|Description|
|-|-|-|
code |int| Error code, anything but 0 indicates failure
msg  |string| Error code description
data | - | -
  ∟has_more|bool|There are more groups if has_more is true.
  ∟page_token|string|See request parameter description for details.
  ∟groups|-|-
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
    "data": {
        "groups": [
            {
                "avatar": "http://p3.pstatp.com/origin/78c0000676df676a7f6e",
                "chat_id": "oc_9e9619b938c9571c1c3165681cdaead5",
                "description": "description1",
                "name": "test1",
                "owner_open_id": "ou_194911f90c43ec42d1ba0e93f22b8fb1",
                "owner_user_id": "ca51d83b"
            },
            {
                "avatar": "http://p4.pstatp.com/origin/dae10015cfb346541010",
                "chat_id": "oc_5ce6d572455d361153b7cb51da133945",
                "description": "description2",
                "name": "test2",
                "owner_open_id": "ou_194911f90c43ec42d1ba0e93f22b8fb1",
                "owner_user_id": "ca51d83b"
            }
        ],
        "has_more": false,
        "page_token": "0"
    },
    "msg": "ok"
}
```

### Error code

For details, please refer to: Service-side error codes
