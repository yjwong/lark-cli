---
title: "Obtain Single Approval Form"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uADNyUjLwQjM14CM0ITN"
method: "POST"
api_path: "https://www.larksuite.com/approval/openapi/v2/approval/get"
section: "Approval"
scopes:
  - "approval:approval:readonly"
updated: "1658309940000"
---

# View approval definitions
> This API has been upgraded to enhance its security. We recommend that you migrate to the new version>>

This API is used to get the details of an approval definition based on an Approval Code and create approval instance requests.

## Request
| HTTP URL | https://www.larksuite.com/approval/openapi/v2/approval/get |
| --- | --- |
| HTTP Method | POST |
| Supported app types | custom,isv |
| Required scopes | `approval:approval:readonly` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" Learn more about how to access and use access_token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Request body
|Parameter|Type|Required|Description|
|-|-|-|-|
|approval_code|String|Yes|Approval definition code|
|locale|String|No|zh-CN - Chineseen-US - English  ja-JP - Japanese| ### Request body example

```json
{
    "approval_code": "7C468A54-8745-2245-9675-08B7C63E7A85",
    "locale": "en-US"
}
```

## Response

### Response body
|Parameter|Type|Required|Description|
|-|-|-|-|
|code|int|Yes|Error code. A value other than 0 indicates failure.|
|msg|String|Yes|Return code description|
|data|map|Yes|Returned business information|
|  ∟approval_name|String|Yes|Approval title|
|  ∟status|String|Yes|Approval definition statusACTIVE -EnabledINACTIVE -DisabledDELETED -DeletedUNKNOWN -Unknown|
|  ∟form|String|Yes| JSON array, **widget information**|
|    ∟id|String|Yes|Widget ID|
|    ∟widget_default_value|String|Yes|Default Value of Widget|
|    ∟custom_id|String|No|Custom widget ID|
|    ∟default_value_type|String|Yes|Default Value Type of Widget|
|    ∟name|String|Yes|Widget name|
|    ∟type|String|Yes|Widget type|
|    ∟enable_default_value|bool|Yes|Default Value Enabled or Not|
|    ∟widget_default_value|String|Yes|Default Value of Widget|
|    ∟default_value_type|String|Yes|Default Value Type of Widget|
|    ∟display_condition|String|No|Control explicit and invisible conditions|
|      ∟conditional|String|No||
|      ∟conditions|list|No||
|        ∟conditional|String|No|Multiple conditions are met at the same time|
|        ∟expressions|list|No|Judgment rules|
|          ∟source_widget|string|No|Condition value|
|  ∟node_list|list|Yes|Node information|
|    ∟name|String|Yes|Node name|
|    ∟need_approver|bool|Yes|Whether it is a node selected by the initiatortrue - The initiator must specify approvers|
|    ∟node_id|String|Yes|Node ID|
|    ∟custom_node_id|String|No|Custom node ID. not returned if no value is set|
|    ∟node_type|String|Yes|Approval methodAND -Everyone assignedOR -Anyone assignedSEQUENTIAL -Assigned one by oneCC_NODE -CC node|
|  ∟viewers|list|Yes|Viewer list|
|    ∟type|String|Yes|Viewer type, which can be: TENANT: visible within tenantDEPARTMENT: visible to the specified department USER: visible to the specified user ROLE: visible to the specified role USER_GROUP: visible to the specified user groupNONE: visible to none|
|    ∟open_id|String|No|When the viewer type is DEPARTMENT, open_id indicates open_id of the department. When the viewer type is USER, open_id indicates open_id of the user.When the viewer type is ROLE, open_id indicates open_id of the role.When the viewer type is USER_GROUP, open_id indicates open_id of the user group.|
|    ∟user_id|String|No|User ID of the viewer when the viewer type is USER| **Widget information**: 
|Parameter|Type|
|-|-|
|Single-line text|input|
|Multi-line text|textarea|
|Number|number|
|Amount|amount|
|Date|date|
|Date range|dateInterval|
|Formula|formula|
|Attachment|attachment/attachmentV2|
|Image| image/imageV2|
|Contact|contact|
|Associate approval|connect|
|Address|address|
|Telephone|telephone|
|Leave widgets|leaveGroup|
|Overtime widgets|workGroup|
|Shift swap widgets|shiftGroup|
|Correction widgets|remedyGroup|
|Business trip widgets|tripGroup|
|Outing widgets|outGroup|
```json
{
    "id": "widget1",
    "custom_id": "user_name",
    "name": "Item application",
    "type": "input",
    "printable": true,
    "required":true
}
```
|Parameter|Type|
|-|-|
|Single option|radio/radioV2|
|Multiple options|checkbox/checkboxV2|
```json
{
    "id": "widget1",
    "name": "Item application",
    "type": "radio",
    "printable": true,
    "required":true,
    "option": [
        {
            "text": "1",
            "value": "jxpsebqp-0"
        }
    ]
}
```
|Parameter|Type|
|-|-|
|Details|fieldList|
```json
{
 "id": "widget1",
 "name": "Item application",
 "type": "fieldList",
 "printable": true,
 "required": true,
 "option":{
 	"input_type": "LIST",
 	"print_type": "LIST"
 },
 "children":[
 	{
 	"id": "widget2",
 	"name": "Item name",
 	"type": "input"
 }
 ]
}
```
|Parameter|Type|
|-|-|
|telephone|telephone|
```json
{
    "id":"widget1",
    "name":"Item application",
    "type":"telephone",
    "printable": true,
    "required":true,
    "option":{
        "available_type": "FIXED_LINE_OR_MOBILE"
      }
}
```

### Response body example

```json
{
    "code": 0,
    "msg": "success",
    "data": {
        "approval_name": "Payment",
        "status": "ACTIVE",
        "form": "[{\"id\": \"widget1\", \"custom_id\": \"user_name\",\"name\": \"Item application\",\"type\": \"textarea\",\"printable\": true,\"required\": true}\"]",
        "node_list": [
            {
                "name": "Approval",
                "need_approver": true,
                "node_id": "46e6d96cfa756980907209209ec03b64",
                "node_type": "AND"
            }
        ],
        "viewers": [
            {
                "open_id": "",
                "type": "TENANT",
                "user_id": ""
            }
        ]
    }
}
```
