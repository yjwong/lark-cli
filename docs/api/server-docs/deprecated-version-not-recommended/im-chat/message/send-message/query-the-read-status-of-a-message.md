---
title: "Query the Read Status of a Message"
url: "https://open.larksuite.com/document/ukTMukTMukTM/ukTM2UjL5EjN14SOxYTN"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/message/v4/read_info/"
section: "Deprecated Version (Not Recommended)"
updated: "1626158293000"
---

# Query the read status of a message

> The bot feature must be enabled. You can only query the read status for messages sent by the bot.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/message/v4/read_info/ |
| HTTP Method | POST | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value Format**: "Bearer `access_token`" **Example Value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" Learn more about obtaining and using access_token. |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Request body
|Parameter|Type|Required/Optional|Description|Default Value|Example|
|--|--|--|--|--|--  |
|message_id | string | Required | The message ID | | om_4d9bed57fa51660c96fcda6238d0e84e | ### Request body example

```json
{
   "message_id": "om_4d9bed57fa51660c96fcda6238d0e84e"
}
```

## Response
### Response body
|Parameter|Type|Description|
|-|-|-|
|code|int|Response code, anything but 0 indicates failure|
|msg|string|Description of the response code|
|data | - | - |
|  ∟read_users | list  | Information on users who have read the message |
|    ∟open_id | string  | User open_id |
|    ∟user_id | string  | User ID; this field is not returned for ISVs |
|    ∟timestamp | string  | Time the message was read | ### Response body example
```json
{
    "code": 0,
    "data": {
        "read_users": [
            {
                "open_id": "ou_18eac85d35a26f989317ad4f02e8bbbb",
                "timestamp": "1570697776",
                "user_id": "ca51d83b"
            }
        ]
    },
    "msg": "ok"
}
```

### Error code

For details, please refer to: Error Codes
