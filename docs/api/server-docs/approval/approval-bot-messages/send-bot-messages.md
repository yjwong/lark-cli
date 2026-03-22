---
title: "Send Bot messages"
url: "https://open.larksuite.com/document/ukTMukTMukTM/ugDNyYjL4QjM24CO0IjN"
method: "POST"
api_path: "https://www.larksuite.com/approval/openapi/v1/message/send"
section: "Approval"
scopes:
  - "approval:approval:readonly"
updated: "1675167428000"
---

# Send Bot messages

This API is used to push messages to users through the Bot of Lark Approval. When there is a new approval to-do, or the status of an approval to-do is updated, users can be informed through the Bot of Lark Approval. Developers can also make use of the Open Platform's capability to build a new Bot to push approval-related information.If the push is successful, but no message is received, you should check whether the aggregate push by approval bot is enabled.

## Request
| HTTP URL | https://www.larksuite.com/approval/openapi/v1/message/send |
| --- | --- |
| HTTP Method | POST |
| Supported app types | custom |
| Required scopes | `approval:approval:readonly` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" Learn more about how to obtain and use access_token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Common template - Request body
|Parameter|Type|Required|Description|Example
|-|-|-|-|-
|template_id|String|Yes|Template ID, which identifies the card template to be used. It is an enum value.|1001|
|user_id|string|Yes|Accept bot alert user's user_id, which is the open platform employee id (equivalent to user_id)|b85s39b
|uuid|string|No| The uuid has a maximum length of 64 characters, and is used to ensure that a card is sent once only. If two cards have the same uuid+user_id, only one of them will be sent.|1234567
|approval_name|string|No|{approval definition name} in the Title.  i18n key is supported.|@i18n@1
|title_user_id|string|No|user id of the {requester}, {approver}, {commenter}, or {copy sender} in the Title.|b85s39b
|comment|string|No|Comment. The markdown syntax of message card and the i18n key syntax are supported.|@i18n@2
|content|map|No|Body|
|  ∟user_id|string|No|Requester's user id.When user_id is empty, requester is not displayed.|b85s39b
|  ∟department_id|string|No|Requester's department department_id.If user_id is specified, department_id is required.This field does not support passing open_department_id.|--
|  ∟user_name|string|No|Requester's name. If user_id is specified, user_id is used.This is used for scenarios where the requester is not a real user.|@i18n@3
|  ∟department_name|string|No|Requester's department name. If department_id is specified, department_id is used.|@i18n@4
|  ∟summaries|list|Yes|Approval summaries. A maximum of 5 summaries are supported.|
|    ∟summary|string|No|Approval summary. The markdown syntax of message card and the i18n key syntax are supported.![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/69bada4fe9c3e017bcdaa202831be59a_x0YSohrDJt.png?lazyload=true&width=988&height=1280)|
|note|string|No|Remarks area, in which the approval source, and whether it is required to access in the intranet are indicated.![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/6ffb657012336bab931d3f913ce5d9d4_Hrgxvqb4T9.png?lazyload=true&width=988&height=1280)|@i18n@6
|open_id|string|Yes|Unique identifier of the user in an app.|ou-8ec33278bc2
|sender_user_id|string|No|Sender's employee id, used to send IMcards.|b85s39b
|text|string|No|Message sent with the forwarded IM card.![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/8ef4c5af656f58f92c9cf7721945342f_wLmZNKKY1v.png?lazyload=true&width=622&height=1280)|Please complete the approval as soon as possible.
|actions|list|Yes|Action area, which can contain up to 2 actions. The first action is used to view details and is required, and the second one can be customized and is optional.![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/23a9a6dbba28ae585796d62955610db5_6wtTxCMjwl.png?lazyload=true&width=1046&height=1280)|
|  ∟action_name|string|Yes|Action type. The first action's action_name must be DETAIL, and the custom action's action_name is an i18n key.![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/b08fe3bf4f7316d9bf60cf57ff119c94_JIHcNCWoOV.png?lazyload=true&width=989&height=1280)|@i18n@7
|  ∟url|string|Yes|Default link. Different redirect urls should be configured for different terminals, and the link must contain a schema to be valid, such as https and http.|https://bytedance.com
|  ∟android_url|string|Yes|Link for Android.|https://bytedance.com
|  ∟ios_url|string|Yes|Link for iOS.|https://bytedance.com
|  ∟pc_url|string|Yes|Link for PC.|https://bytedance.com
|action_callback|map|No|Quick approval.Quick approval is available only for orange cards.|
|  ∟action_callback_url|string|No|Action callback url for the third-party system. After the task approver of the "Pending" list approves or rejects, the approval center calls this URL to notify the third-party system.|http://www.feish.cn/approval/openapi/instanceOperate
|  ∟action_callback_token|string|No|Callback token, used for the business system to validate whether the request comes from Approval. Refer to Event subscription overview  for details.|sdjkljkx9lsadf110
|  ∟action_callback_key|string|No|Request parameter encryption key. If this parameter is configured, the request parameters will be encrypted and the business party must decrypt the request. For the encryption and decryption methods, refer to  here.|gfdqedvsadfgfsd
|  ∟action_context|string|No|Action context. This parameter will be included in the callback.|asdasdasdasd
|action_configs|list|No|Quick approval action configuration . Quick approval is available only for orange cards.![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/b08fe3bf4f7316d9bf60cf57ff119c94_949QCJilYA.png?lazyload=true&width=989&height=1280)|
|  ∟action_type|string|Yes|Action type:APPROVE - ApproveREJECT - Reject{KEY} - Any English string. If any string is used,  action_name is required.|APPROVE
|  ∟action_name|string|No|Action name. i18n key  is used for frontend display. If  action_type  is neither  APPROVE nor REJECT, this field is required. i18n key is supported. The text content is provided in  i18n_resource .|@i18n@8
|  ∟is_need_reason|bool|No|Whether comments are required. If yes, the user will be redirected to the comments page after selecting an action.![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/9c8098a3b221efb2662518e09b942f7e_dApPUrfHMh.png?lazyload=true&width=363&height=750)|true
|  ∟is_reason_required|bool|No|Whether comments are required or not.|true
|  ∟is_need_attachment|bool|No|Whether attachments are supported in the comments.|true
|  ∟next_status|string|No|Indicates what status will the card be updated to if the callback is successful. If this is specified, Lark Approval will update the card status to {next_status} after the user selects an action. If this is not specified, the business side instead of Lark needs to update the card status. For the status values, refer to Update approval Bot messages.|APPROVEDREJECTED
|i18n_resources|list|Yes|Internationalized text.|
|  ∟locale|string|Yes|Languagezh-CN - Chineseen-US - Englishja-JP  - Japanese|zh-CN
|  ∟is_default|bool|Yes|Whether to use the default language. If the default language is used, all keys should be contained; if a non-default language is used but its keys do not exist, the default language will be used instead.|true
|  ∟texts|map|Yes|Text key, whose value is in the format of i18n key and starts with  @i18n@ .|{"@i18n@1": "Scope request",   "@i18n@2": "OAApproval","@i18n@3": "Permission"}

### Common template - Request body example

```json
{
    "template_id":"1001",
    "user_id":"b85s39b",
    "uuid":"uuid",
    "approval_name":"@i18n@1",
    "title_user_id":"od-8ec33278bc2",
    "comment":"@i18n@2",
    "content":{
        "user_id":"b85s39b",
        "department_id":"od-8ec33278bc2",
        "summaries":[
            {
                "summary":"@i18n@3"
            }
        ]
    },
    "note":"@i18n@4",
    "actions":[
        {
            "action_name":"DETAIL",
            "url":" https://bytedance.com",
            "android_url":"https://bytedance.com",
            "ios_url":"https://bytedance.com",
            "pc_url":"https://bytedance.com"
        }
    ],
    "action_configs": [
        {
            "action_type": "APPROVE",
            "is_need_reason": true,
            "is_reason_required": true,
            "is_need_attachment": true,
            "next_status": "APPROVED"
        },
        {
            "action_type": "REJECT",
            "action_name": "@i18n@5",
            "next_status": "REJECTED"
        }
    ],
    "action_callback": {
        "action_callback_url":"http://feish.cn/approval/openapi/operate",
        "action_callback_token":"sdjkljkx9lsadf110",
        "action_callback_key":"gfdqedvsadfgfsd",
        "action_context":"acasdasd"
    },
    "i18n_resources":[
        {
            "locale":"en-US",
            "is_default":true,
            "texts":{
                "@i18n@1":"Temporary release",
                "@i18n@2":"Disapproval",
                "@i18n@3":"Need to modify",
                "@i18n@4":"From OA,please access through the internal network.",
                "@i18n@5":"Cancel"
            }
        }
    ]
}
```

### Custom template - Request body 

| Parameter                     | Type   | Required | Description                                                         |
| ------------------------ | ------ | ---- | ------------------------------------------------------------ |
| template_id              | string | Yes   | Template ID, which identifies the card template to be used. It is an enum value.                        |
| user_id                  | string | Yes   | Indicates the user to whom messages are pushed. It is equivalent to the  employee id                        on the Open Platform.|
| uuid                     | string | No   | Idempotent id, with a maximum length of 64 characters. It is used to ensure that a card is sent once only. If two cards have the same uuid+user_id, only one of them will be sent. |
| custom_title             | string | Yes   | Title. i18n key                                                is supported.|
| custom_content           | string | Yes   | Body area. The markdown syntax of message card and the i18n key  syntax are supported.|
| note                     | string | No   | Remarks area, in which the approval source, and whether it is required to access in the intranet are indicated. i18n key          is supported.|
| actions                  | list   | No  | Action area. A maximum of 2 actions are supported.                                              |
|   ∟action_name       | string | Yes   | Action content. i18n key                                            is supported.|
|     ∟url         | string | Yes   | Default link. Different redirect  urls should be configured for different terminals, and the link must contain a schema to be valid, such as https and http.|
|     ∟android_url | string | Yes   | Link for Android.                                                  |
|     ∟ios_url     | string | Yes   | Link for iOS.                                                      |
|     ∟pc_url      | string | Yes   |Link for  PC.                                                       |
| i18n_resources           | list   | Yes   | Internationalized text.                                                   |
|   ∟locale            | string | Yes   | Language.zh-CN - Chineseen-US - Englishja-JP  - Japanese       |
|   ∟is_default        | bool   | Yes   | Whether to use the default language. If the default language is used, all keys should be contained; if a non-default language is used but its keys do not exist, the default language will be used instead. |
|   ∟texts             | map    | Yes   | Text key, whose value is in the format of i18n key and starts with @i18n@.                   | ### Custom template - Request body example
```json
{
    "template_id":"1021",
    "user_id":"employeeId1",
    "uuid":"uuid",
    "custom_title":"@i18n@1",
    "custom_content":"@i18n@2",
    "note":"@i18n@3",
    "actions":[
        {
            "action_name":"@i18n@4",
            "url":" https://bytedance.com",
            "android_url":"https://bytedance.com",
            "ios_url":"https://bytedance.com",
            "pc_url":"https://bytedance.com"
        }
    ],
    "i18n_resources":[
        {
            "locale":"en-US",
            "is_default":true,
            "texts":{
                "@i18n@1":"Custom template",
                "@i18n@2":"Please help process my approval as soon as possible.",
                "@i18n@3":"From OA,please access through the internal network.",
                "@i18n@4":"DETAIL"
            }
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
|  ∟message_id|string|Yes|Message ID, used for card update.|
### Response body example 
```json
{
    "code":0,
    "msg":"success",
    "data":{
        "message_id": "6968359519504171036"
    }
}
```

### Error code
For details, refer to Server-side error codes.

## Template list
|Template No.|Name|Parameter|Notes|
|-|-|-|-
|1001|Comment received|![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/2ac0305622c5501c209053cb0ddcbeaf_HzqM0VhTYy.png?lazyload=true&width=610&height=712)|Required fields:approval_nametitle_user_idsummariesactionscomment|
|1002|Approval saved as to-do|![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/d0fff024f213f964560747b74729bef9_rj2z47IHab.png?lazyload=true&width=614&height=580)|Required fields:approval_nametitle_user_idsummariesactions
|1003|Approval rejected|![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/dff5d01e12328ef48636713b05835c31_XlvlCdkBbc.png?lazyload=true&width=612&height=714)|Required fields:approval_nametitle_user_idsummariesactions|
|1004|Approval approved|![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/1e3c92f6cf391be8c498fe7cdcc63a17_z6ypzczt2W.png?lazyload=true&width=614&height=680)|Required fields:approval_nametitle_user_idsummariesactionscomment|
|1005|Approval initiated successfully|![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/870b77deaea9c757c3952b831abfe138_lsO079Rd9S.png?lazyload=true&width=614&height=546)|Required fields:approval_namesummariesactions|
|1006|Approval to be closed|![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/e82c7e84d3940427a6c8c01b31ef3206_FyYH3Gzy7l.png?lazyload=true&width=616&height=580)|Required fields:approval_namesummariesactions|
|1007|Approval closed|![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/4051b9ec62b8f8ec0c4589fb50ba3bbe_POXtxF8ymf.png?lazyload=true&width=612&height=584)|Required fields:approval_namesummariesactions|
|1008|Approval to-do received|![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/ef8750e1e2a84048733bc118e4757ce6_dRnwHyD79Z.png?lazyload=true&width=612&height=668)|Required fields:approval_nametitle_user_idsummariesactionsQuick approvalaction_configsaction_callbackNote:When title_user_id is empty, the title displayed is "{approval_name}" for your approval.|
|1028|Approved to-do received (no approval initiator)||Required fields:approval_namesummariesactionsQuick approvalaction_configsaction_callback|
|1009|Added as approver|![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/186466158750d0b7d492eba93d840c17_5V9JsvTZs4.png?lazyload=true&width=616&height=800)|Required fields:approval_nametitle_user_idsummariesactionsQuick approvalaction_configsaction_callback|
|1010|Transferred|![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/d20fc202cf69bb780c1ce104d068d181_3TG8VyeUUU.png?lazyload=true&width=616&height=802)|Required fields:approval_nametitle_user_idsummariesactionsQuick approvalaction_configsaction_callback|
|1011|Commissioned|![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/0a35ba22ea04293b845b9fb95b1ec411_WadKLNdJuV.png?lazyload=true&width=614&height=666)|Required fields:approval_nametitle_user_idsummariesactionsQuick approvalaction_configsaction_callback|
|1012|Returned|![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/512755673ac31cbf5ac2ad1d690790f7_JhRosk5M8n.png?lazyload=true&width=612&height=794)|Required fields:approval_nametitle_user_idsummariesactionsQuick approvalaction_configsaction_callback|
|1013|Reminded manually(Personal IM)||Required fields:approval_nametitle_user_idsummariesactionsQuick approvalaction_configsaction_callback|
|1015|Canceled|![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/c834878ecbdd607d40d660640b575463_zFk5p384KH.png?lazyload=true&width=614&height=812)|Required fields:approval_nametitle_user_idsummariesactions|
|1016|CCed|![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/75a57fb74426c022e5c9c3b64bd67d4b_ZrXvfZJC5Y.png?lazyload=true&width=612&height=816)|Required fields:approval_nametitle_user_idsummariesactions|
|1017|Comment replied to|![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/edfc1a23de9effb5c3edea927f2f8aab_UFchSU2M1q.png?lazyload=true&width=610&height=818)|Required fields:approval_nametitle_user_idsummariesactionscomment|
|1018|Comment mentioned|![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/8c26777fc28235a36b3d3eb5008a1b7b_2rNy3g8Syl.png?lazyload=true&width=616&height=818)|Required fields:approval_nametitle_user_idsummariesactionscomment|
|1019|Transferred to manager because of requester's termination|![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/2cc4612c546cc2bc65412724aa230d68_jQPylLKOgi.png?lazyload=true&width=614&height=658)|Required fields:approval_nametitle_user_idsummariesactions|
|1020|CCed to manager because of approver's termination|![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/19065ac376d4369e3529c5928cca1c26_KCFdxj2iti.png?lazyload=true&width=610&height=666)|Required fields:approval_nametitle_user_idsummariesactionsQuick approvalaction_configsaction_callback|
|1021|Custom template|![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/fe4744f5b9a37e265ebe1efacf1e6408_HmPqXUqEgH.png?lazyload=true&width=612&height=472)|Required fields:custom_titlecustom_content|
|1024|Approval shared||Required fields:approval_namesummariesactions|
|1026|CCed by system||Required fields:approval_nametitle_user_idsummariesactions|
|1027|Comment added||Required fields:approval_nametitle_user_idsummariesactionscomment|
