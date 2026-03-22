---
title: "Create an Approval Definition"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uUzNyYjL1cjM24SN3IjN"
method: "POST"
api_path: "https://www.larksuite.com/approval/openapi/v2/approval/create"
section: "Approval"
scopes:
  - "approval:approval:readonly"
updated: "1705410965000"
---

# Create an approval definition
> This API has been upgraded to enhance its security. We recommend that you migrate to the new version>>

This API is used to create a simple approval definition, whose basic information, form and process can be specified. This API is not recommended for custom apps. Contact the administrator to create definitions in the Approval Admin if necessary.

Please be cautious when calling the API. Created approval definitions cannot be disabled or deleted.

## Request
| HTTP URL | https://www.larksuite.com/approval/openapi/v2/approval/create |
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
|approval_name|String|Yes|Internationalized text key of the approval title. It starts with @i18n@ and has a minimum length of 9 characters.|
|approval_code|String|No|If this is left empty, it indicates creating a new approval.|
|description|String|No|Internationalized text key of the approval description. It starts with @i18n@ and has a minimum length of 9 characters.|
|viewers|list|Yes|Visible scope of Approval Frontend. See the field description below for details.|
|form|map|Yes|Approval definition form.|
|  ∟form_content|String|Yes|Approval definition form content, which is listed as a JSON array. See the field description below for details.|
|node_list|list|Yes|Approval definition nodes, which are listed as the map array. The start node should be the first element of the list and the end node should be the last element.|
|  ∟id|String|Yes|Node ID. The start node's ID is START, and the end node's ID is END. name, node_type, and approver are not required for the start and end nodes.|
|  ∟name|String|Yes|Internationalized text key of the node name. It starts with @i18n@ and has a minimum length of 9  characters.|
|  ∟node_type|String|Yes|Approval type enumeration.- AND Everyone assigned- OR Anyone assigned- SEQUENTIAL SequentialWhen node_type is SEQUENTIAL, the approver must be selected by the initiator.|
|  ∟approver|list|Yes|Approver list. See the field description below for details.|
|  ∟ccer|list|No|Copy sender list. See the field description below for details.|
|settings|map|No|Other settings of the approval definition.|
|  ∟revert_interval|int|No|The period of time (in seconds) during which an approval instance can be canceled after it is approved. It defaults to 31 days. 0 indicates not cancelable.|
|  ∟revert_option|int|No|Whether an approval can be canceled after the first node. It defaults to 1. 0 indicates No.|
|config|map|No|Approval definition configuration items, used to configure whether an approval definition can be modified by users in the Approval Admin.|
|∟can_update_viewer|bool|No|Whether the visible scope can be modified. It defaults to false.|
|∟can_update_form|bool|No|Whether the form content can be modified. It defaults to false.|
|∟can_update_process|bool|No|Whether process nodes can be modified. It defaults to false.|
|∟can_update_revert|bool|No|Whether the period of time for cancellation can be modified. It defaults to false.|
|∟help_url|String|No|Used to configure the help page link, by clicking which users can be redirected to the help page from the Approval Admin.|
|icon|int|No|Approval icon enumeration. See the field description below for details. It defaults to 0.|
|i18n_resources|list|Yes|Internationalized text.|
|  ∟locale|String|Yes|Language:zh-CN - Chineseen-US - Englishja-JP   - Japanese|
|  ∟is_default|bool|Yes|Whether to use the default language. If the default language is used, all keys should be contained; if a non-default language is used but its  keys do not exist, the default language will be used instead.|
|  ∟texts|map|Yes|Text key, whose value is in the format of i18n key and starts with @i18n@.| ### Request body example
```json
{
    "approval_name": "@i18n@approval_name",
    "approval_code": "7C468A54-8745-2245-9675-08B7C63E7A85",
    "description": "@i18n@description",
    "viewers":[
        {
            "type":"TENANT",
            "open_id":"ou_e03053f0541cecc3269d7a9dc34a0b21",
            "user_id":"f7cb567e"
        }
    ],
    "form": {
          "form_content": "[{\"id\":\"user_name\", \"type\": \"input\", \"required\":true, \"name\":\"@i18n@widget1\"}]"
        },
    "node_list": [{
          "id": "START"
        },{
          "id": "node_id",
          "name": "@i18n@node_name",
          "node_type": "AND",
          "approver": [
            {
              "type": "Personal",
              "open_id": "ou_e03053f0541cecc3269d7a9dc34a0b21",
              "user_id": "f7cb567e"
            }
          ],
          "ccer": [
            {
              "type": "Supervisor",
              "open_id": "ou_e03053f0541cecc3269d7a9dc34a0b21",
              "user_id": "f7cb567e"
            }
          ]
        },{
          "id": "END"
        }],
    "settings" : {
          "revert_interval":0
        },
    "config" : {
          "can_update_viewer": false,
          "can_update_form": false,
          "can_update_process": false,
          "can_update_revert": false,
          "help_url":"https://www.baidu.com"
        },
    "icon": 0,
    "i18n_resources" : [{
          "locale": "zh-CN",
          "texts" : {
              "@i18n@approval_name": "Approval title",
              "@i18n@description": "Approval description",
              "@i18n@widget1": "Widget 1",
              "@i18n@node_name": "Node name"
            },
          "is_default": true
        }]
}
```
## Response

### Response body
|Parameter|Type|Required|Description|
|-|-|-|-|
|code|int|Yes|Error code. A value other than 0 indicates failure.|
|msg|String|Yes|Return code description.|
|data|map|Yes|Returned business information|
|  ∟approval_code|String|Yes|Approval definition code| **Field description**: 

**viewers field description**

viewers field specifies who can initiate the approval from the frontend of the Approval app.

The viewers field is a list composed of multiple viewers. Each viewer is structured as follows:

```
{
    "type":"TENANT",
    "user_id":"",
    "open_id":""
}
```
The enum values of type are:
- TENANT - Visible to all
- NONE - Invisible to all
- USER - Visible to specified users
- DEPARTMENT - Visible to specified departments
> 1. When the type is USER, either user_id or open_id is required to specify the users to whom the approval is visible.
> 2. When the type is DEPARTMENT, the open_id of departments is required to specify the departments to which the approval is visible.

**approver and ccer field description**

The approver and ccer fields specify the approvers and copy senders respectively at an approval node. This field is a list of users, and the user is structured as follows:

```
{
    "type":"Personal",
    "user_id":"",
    "open_id":""
}
```
The enum values of type are:
- Supervisor - Manager (bottom up)
- SupervisorTopDown - Manager (top down)
- DepartmentManager - Department supervisor (bottom up)
- DepartmentManagerTopDown - Department supervisor (top down)
- Personal - Specified member
- Free - Selected by the initiator
> 1. When the type is Supervisor, SupervisorTopDown, DepartmentManager , or DepartmentManagerTopDown, their level should be specified in the user_id. For example, for approval by third-level managers from bottom up, user_id is 3.
> 2. When the type is Personal, either user_id or open_id is required to specify users.
> 3. When the approver is Free (selected by the initiator), neither user_id nor open_id is required. ccer does not support Free.

**form_content  field description**

Form item's id and name should be specified. The id must not be duplicated and the name is the internationalized text key of the form name, which starts with @i18n@ and has a minimum length of 9 characters.
You can set default values for the Input, Text Area, Radio V2, Contact, and Department widgets when you create the approval definition in Lark Approval. For more information, [see Set default values for widgets API documentation.](https://bytedance.larkoffice.com/docx/doxcndQPK28SEzQTWyeTajWKO2d)

1. Short answer (input), paragraph (textarea), number (number), image (image), attachment (attachmentV2)
```
{
    "id":"",
    "name":"",
    "type":"input",
    "required":true
}
```
2. Amount (amount)
```
{
    "id":"",
    "name":"",
    "type":"amount",
    "required":true,
    "value":"CNY"
}
```
> value  is the enum value of the amount widget, including:
> - CNY - Chinese yuan
> - USD - United States dollar
> - EUR - Euro
> - JPY - Japanese yen
> - CAD - Canadian dollar
> - CHF - Swiss franc
> - SGD - Singapore dollar
> - AUD - Australian dollar
> KRW - South Korean won
> - INR - Indian rupee
> - TWD - New Taiwan dollar
> - HKD - Hong Kong dollar
> - MOP - Macanese pataca
> - THB - Thai baht
> - IDR - Indonesian rupiah
> - PHP - Philippine peso
> - MYR - Malaysian ringgit
3. Description  text
```
{
    "id":"",
    "name":"",
    "type":"text",
    "required":true,
    "value":"@i18n@text"
}
```
> value is the internationalized text key of the description, which starts with @i18n@, and has a minimum length of 9 characters.
4. Checkboxes (radioV2), multiple choice (checkboxV2)

Case 1: Common checkboxes/multiple choice

```
{
    "id":"",
    "name":"",
    "type":"radioV2",
    "required":true,
    "value":[{"key":"1","text":"@i18n@choice1"},{"key":"2","text":"@i18n@choice2"}]
}
```

Case 2: Checkboxes/multiple choice associated with external data sources:Associate external options.
```
{
    "id":"",
    "name":"",
    "type":"radioV2",
    "required":true,
    "value":[],
    "externalData":{
        "externalUrl":"https://www.xxx_bytedance.net/",
        "token":"t",
        "key":"k",
        "linkageConfigs":[
            {
                "linkageWidgetID":"widget1",
                "key":"linkageWidget1",
                "value":"example"
            }
        ],
        "externalDataLinkage":true
    }
}
```

> value is the key and text of each option. The key must not be duplicated, and the text starts with @i18n@ and has a minimum length of 9 characters.
5. Date (date)
```
{
    "id":"",
    "name":"",
    "type":"date",
    "required":true,
    "value": "YYYY-MM-DD"
}
```
> value  is the date format:
> - YYYY-MM-DD: MM-DD-YYYY
> - YYYY-MM-DD a: MM-DD-YYYY AM/PM
> - YYYY-MM-DD hh:mm: MM-DD-YYYY hh:mm
6. Associate (connect) 
```
{
    "id":"",
    "name":"",
    "type":"connect",
    "required":true,
    "value":["code1","code2"]
}
```
> value  is the  code of associated approval definition.
7. Contact (contact)
```
{
    "id":"",
    "name":"",
    "type":"contact",
    "required":true,
    "value":{
        "ignore": true,
        "multi": false
      }
}
```
> value is the configuration items of contact widget:
> - ignore: Whether "can select oneself" is not allowed. It defaults to false, indicating that the user themselves can also be selected.
> - multi: Whether "can select more than one person" is not allowed. It defaults to  false, not available
8. Address (address)
```
{
    "id":"",
    "name":"",
    "type":"address",
    "required":true,
    "value":{
        "enableDetailAddress": false,
        "requiredDetailAddress": false,
        "preLocating": false
      }
}
```
> value is the configuration items of address widget:
> - enableDetailAddress: Whether address details are enabled. It defaults to false.
> - requiredDetailAddress: Whether address details are required or not. It defaults to false.
> - preLocating: Whether auto-positioning is enabled. It defaults to false.
9. Date interval (dateInterval)
```
{
    "id":"",
    "name":"",
    "type":"dateInterval",
    "required":true,
    "value":{
        "format": "YYYY-MM-DD",
        "intervalAllowModify": false,
      }
}
```
> value is the configuration items of date interval widget:
> - format: Date format, which is the same as that of the date widget.
> - intervalAllowModify: Whether the duration can be modified. It defaults to false.
10. telephone
```
{
    "id":"",
    "name":"",
    "type":"telephone",
    "required":true,
    "option":{
        "availableType": "FIXED_LINE_OR_MOBILE"
      }
}
```
>option indicates the configuration items of the Contact No
> - widget. availableType indicates the available types of contact numbers. Valid values are MOBILE, FIXED_LINE, and FIXED_LINE_OR_MOBILE.
11. Details fieldList
```
{
    "id": "",
    "name": "",
    "type": "fieldlist",
    "required": true,
    "value":
    [
        {
            "id": "",
            "name": "",
            "type": "input",
            "required": true
        }
    ],
    "option":
    {
        "inputType": "LIST",
        "printType": "LIST"
    }
}
```
>option indicates the configuration items of the Details widget. 
>- inputType：inputType indicates the format of input text. Valid values are LIST and FORM.
>- printType: printType indicates the print format of the Details widget. Valid values are LIST and FORM.

**Icon enumeration**: 
The following are icons 0-24 from left to right and top to bottom.

![图片名称](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/ark/5bb25389e4d480e45335ab363d45e2e6.png?height=564&lazyload=true&width=564)
### Response body example

```json
{
    "code": 0,
    "msg": "success",
    "data": {
        "approval_code":"81D31358-93AF-92D6-7425-01A5D67C4E71"
    }
}
```
