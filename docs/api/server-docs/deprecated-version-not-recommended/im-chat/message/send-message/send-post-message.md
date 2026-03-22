---
title: "Send Post Message"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uMDMxEjLzATMx4yMwETM"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/message/v4/send/"
section: "Deprecated Version (Not Recommended)"
updated: "1626158290000"
---

# Send post message
Sends a post message to a specified user or chat, including private chats and group chats.

> Bot capability needs to be enabled. For private chats, bot need to be visible to the user. For group chats, bot must be in the group.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/message/v4/send/ |
| HTTP Method | POST |
| Required scopes | Send messages as an app(Historic Version) | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value Format**: "Bearer `access_token`" **Example Value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" Learn more about obtaining and using access_token. |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Request body
|Parameter|Type|Mandatory|Description|Default|Demo|
|--|-----|--|----| -- | --
open_id user_id  email  chat_id  | string | No | To send a private chat message to a user, use either open_id, email, or user_id (obtain via the Get user IDs using email addresses or mobile numbers API). To send a message to a group, use the chat_id for the group (obtain via the Obtains Group List API). The server reads the string in the following order: chat_id > open_id > user_id > email||ou_5ad573a6411d72b8305fda3a9c15c70e|
root_id | string | No | To reply to a particular message, add the corresponding message ID||om_40eb06e7b84dc71c03e009ad3c754195|
msg_type | string | Yes | Message type, must be “post”||post|
content | string | Yes | Message content
  ∟post | string | Yes | post content

### Request body example
```json
{
   "open_id":"ou_5ad573a6411d72b8305fda3a9c15c70e",
   "chat_id":"oc_5ad11d72b830411d72b836c20",  
   "root_id":"om_40eb06e7b84dc71c03e009ad3c754195",
    "user_id": "92e39a99",
    "email":"fanlv@gmail.com",
    "msg_type":"post",
    "content":{
        "post":{
            "zh_cn":{}, // option
            "ja_jp":{}, // option
         	"en_us":{}, // option
        }
    }
}
```

### Curl request demo
```curl 
curl -X POST \
  https://open.larksuite.com/open-apis/message/v4/send/ \
  -H 'Authorization: Bearer t-e2d34b8f4882028baea343e414fe71935d9b392f' \
  -H 'Content-Type: application/json' \
  -H 'Postman-Token: e33a7c86-28ed-4b76-aac3-f2ed4d2d064d' \
  -H 'cache-control: no-cache' \
  -d '{
    "chat_id": "oc_5ce6d572455d361153b7cb51da133945",
    "msg_type": "post",
    "content": {
        "post": {
            "zh_cn": {
                "title": "I am a title",
                "content": [
                    [
                        {
                            "tag": "text",
                            "un_escape": true,
                            "text": "line 1:",
                            "lines": 1
                        },
                        {
                            "tag": "a",
                            "text": "super link",
                            "href": "www.baidu.com"
                        },
                        {
                            "tag": "at",
                            "user_id": "39d898aa",
                            "text": "user-employeeid"
                        }
                    ],
                    [
                        {
                            "tag": "text",
                            "un_escape": true,
                            "text": "line 2:",
                            "lines": 1
                        },
                        {
                            "tag": "text",
                            "un_escape": true,
                            "text": "this is a text",
                            "lines": 1
                        }
                    ],
                    [
                        {
                            "tag": "text",
                            "un_escape": true,
                            "text": "",
                            "lines": 1
                        },
                        {
                            "tag": "img",
                            "image_key": "b15a91e5-156d-4d81-8bac-3d27c321001b",
                            "width": 300,
                            "height": 300
                        }
                    ]
                ]
            }
        }
    }
}'
```

## Response
### Response body
|Parameters|Type|Description|
|-|-|-|
|code|int|Error code, anything but 0 indicates failure|
|msg|string|Error code description|
data | - | -
  ∟message_id |string| Message ID

### Response body example
```json
{
    "code": 0,
    "msg": "ok",
    "data":{
       "message_id": "om_92eb70a7120ga8c3ca56a12ee9ba7ca2"
    }
}
```

### Error code

For details, please refer to: Error Codes

## Request body demo

```json
{
    "chat_id": "oc_5ce6d572455d361153b7cb51da133945",
    "msg_type": "post",
    "content": {
        "post": {
            "zh_cn": {
                "title": "I am a title",
                "content": [
                    [
                        {
                            "tag": "text",
                            "un_escape": true,
                            "text": "line 1 :",
                        },
                        {
                            "tag": "a",
                            "text": "super link",
                            "href": "www.baidu.com"
                        },
                        {
                            "tag": "at",
                            "user_id": "ou_18eac85d35a26f989317ad4f02e8bbbb"
                        }
                    ],
                    [
                        {
                            "tag": "text",
                            "text": "line 2 :",
                        },
                        {
                            "tag": "text",
                            "text": "text content",
                        }
                    ],
                    [
                        {
                            "tag": "text",
                            "text": "",
、                        },
                        {
                            "tag": "img",
                            "image_key": "b15a91e5-156d-4d81-8bac-3d27c321001b",
                            "width": 300,
                            "height": 300
                        }
                    ]
                ]
            }
        }
    }
}
```

### Results show

![图片名称](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/ark/a2b0643bae338501e642a246625e045d_en.png)

### Supported tags and parameters
#### text
Field|type|Mandatory| Description|Default|Demo|
| -- |-- | -- |-- | -- |-- |
text| string|Yes | Indicates the default text content. If localized content has been set, the localized content is displayed first||test word|
un_escape|bool| No |  url encode the text||false| #### a
Field|type|Mandatory| Description|Default|Demo|
| -- |-- | -- |-- | -- |-- |
text| string| Yes | Indicates the default text content. If localized content has been set, the localized content is displayed first||test word|
un_escape| bool| No |  url encode the text|false|false|
href | string|  Yes |Default link address||https://test.com| #### at
Field |type|Mandatory| Description|Default|Demo|
-- | -- |-- |-- |-- |-- |
user_id |   string|Yes | open_id or user_id||ou_18eac85d35a26f989317ad4f02e8bbbb| #### img
Field |type|Mandatory| Description|Default|Demo|
-- | -- |-- |-- |-- |-- |
image_key |string |Yes |The unique image identifier. Obtain via the Upload image API. ||b15a91e5-156d-4d81-8bac-3d27c321001b",
height |int | Yes | image's height||300|
width |int |Yes |  image's width||300|
