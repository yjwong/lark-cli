---
title: "Update Bot messages"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uAjNyYjLwYjM24CM2IjN"
method: "POST"
api_path: "https://www.larksuite.com/approval/openapi/v1/message/update"
section: "Approval"
scopes:
  - "approval:approval:readonly"
updated: "1675167431000"
---

# Update Bot messages

This API is used to update approval bot messages based on the approval bot message id and status. It can only be used to update bot messages of Pending templates. For example, an approval to-do message is pushed to a user, and the previously pushed Bot message can be updated as approved after the user processes the message.

## Request
> **Note:** Only support update the approval bot message within 30 days. 
| HTTP URL | https://www.larksuite.com/approval/openapi/v1/message/update |
| --- | --- |
| HTTP Method | POST |
| Supported app types | custom,isv |
| Required scopes | `approval:approval:readonly` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" Learn more about how to obtain and use access_token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Request body

|Parameter|Type|Required|Description|
|-|-|-|-|
|message_id|String|Yes| Card ID, which is obtained when the card is sent.|
|status|string|Yes|Status type, which is used to update the first action's text. The enum values are: APPROVED:  ApprovedREJECTED:  RejectedCANCELLED:  CanceledFORWARDED:  TransferredROLLBACK:  ReturnedADD:  Added approverDELETED:  DeletedPROCESSED:  Processed| ### Request body example

```json
{
    "message_id":"xxxx",
    "status":"CUSTOM",
    "status_name":"@i18n@status_name", // status=CUSTOM时可以自定义
    "detail_action_name":"@i18n@detail_action_name", // status=CUSTOM时可以自定义
    "i18n_resources":[
        {
          "locale": "zh_cn",
          "texts" : {
              "@i18n@status_name": "已废弃",
              "@i18n@detail_action_name": "已废弃按钮" 
            },
          "is_default": true
        }
    ]
}
```

## Response

### Response body

|Parameter|Type|Required|Description|
|-|-|-|-|
|code|int|Yes|Error code. A value other than 0 indicates failure.|
|msg|string|Yes|Return code description.|
|data|map|Yes|Returned business information|
|  ∟message_id|string|Yes|Message ID, used for card update and recall.| ### Response body example

```json
{
    "code":0,
    "msg":"success",
    "data":{
        "message_id": "xxxx"
    }
}
```
