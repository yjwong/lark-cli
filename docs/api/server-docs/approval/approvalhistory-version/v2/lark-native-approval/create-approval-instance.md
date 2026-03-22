---
title: "Create Approval Instance"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uIDNyUjLyQjM14iM0ITN"
method: "POST"
api_path: "https://www.larksuite.com/approval/openapi/v2/instance/create"
section: "Approval"
scopes:
  - "approval:approval:readonly"
updated: "1658309943000"
---

# Create an approval instance
> This API has been upgraded to enhance its security. We recommend that you migrate to the new version>>

To create an approval instance, you must have a detailed understanding of the approval definition form. Then, you must pass values through the API based on the form structure.

##  Request
| HTTP URL | https://www.larksuite.com/approval/openapi/v2/instance/create |
| --- | --- |
| HTTP Method | POST |
| Supported app types | custom,isv |
| Required scopes | `approval:approval:readonly` | ###  Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" Learn more about how to obtain and use access_token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Request body
|Parameter|Type|Required|Description|
|-|-|-|-|
|approval_code|String|Yes|Approval definition code|
|user_id|String|No|User who initiates approval|
|open_id|String|Yes|The Open ID of the approval initiator. If a user_ID is provided, the user_ID should be used. |
|department_id|String|No|The department ID of the approval initiator. If the user only belongs to 1 department, this field can be left empty. If the user belongs to multiple departments, the first department in the department list is selected by default.|
|form|String|Yes|JSON array, **Widget value**|
|∟id|String|Yes|Widget ID, or you can use the custom_id|
|∟type|String|Yes|Widget type|
|∟value|String|Yes|Widget value. The values of different types use different formats.|
|node_approver_user_id_list|map|No|For a node selected by the initiator, the code approvers must be entered.key: node id or custom node id, obtained through View approval definition  value: Approver list|
|node_approver_open_id_list|map|No|Open ID selected by the approver and initiator,	Create a union set with the above field "node_approver_user_id_list".|
|node_cc_user_id_list|map|No|For a node selected by the initiator, you can enter CC senders for the nodekey: node id or custom node id, obtained through View approval definition  value: Approver listA single node can have up to 20 CC senders|
|node_cc_open_id_list|map|No| Open ID selected by the CC recipient and initiatorA single node can have up to 20 CC senders.|
|uuid|String|No|Approval instance uuid, used for idempotent action. A single uuid can only be used to create 1 approval instance. In the case of a conflict, the error code 60012 will be returned. The format must be: XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX (case-insensitive).| ### Request body example

```json
{
    "approval_code":"7C468A54-8745-2245-9675-08B7C63E7A85",
    "user_id":"f7cb567e",
    "open_id":"ou_123456",
    "uuid": "",
    "department_id":"",
    "form":"[{\"id\":\"user_name\", \"type\": \"input\", \"value\":\"test\"}]",
    "node_approver_user_id_list": {
        "46e6d96cfa756980907209209ec03b64": ["f7cb567e"],
        "manager_node_id": ["f7cb567e"]
    },
    "node_approver_open_id_list": {
        "46e6d96cfa756980907209209ec03b64": ["ou_12345"],
        "manager_node_id": ["ou_12345"]
    }
}
```
### Widget value description

> Single-line text

```json
{                                     
    "id": "widget1",
    "type": "input",
    "value": "data"
}
```

> Multi-line text

```json
{                                     
    "id": "widget1",
    "type": "textarea",
    "value": "data"
}
```

> Date 
  
    value is in RFC3339 format,

```json
{                                     
    "id": "widget1",
    "type": "date",
    "value": "2019-10-01T08:12:01+08:00"
}
```

> Single select
  
  value is the value field in the option field of the definition, obtained through View approval definition

```json
{                                     
    "id": "widget1",
    "type": "radio",
    "value": "optionA"
}
```

```json
{                                     
    "id": "widget1",
    "type": "radioV2",
    "value": "k2b8mkx0-h71x5gljn3i-1"
}
```

> Number

```json
{
    "id": "widget1",
    "type": "number",
    "value": 1234.5678
}
```

> Amount
```json
{
    "id": "widget1",
    "type": "amount",
    "value": 1234.5678
}
```

> Formula

  The value obtained through the formula when the definition is created. Returns an error if no match is found.
```json
{
    "id": "widget1",
    "type": "formula",
    "value": 1234.5678
}
```

> Contact
 
 value is the user id array (*This API doesn't support creating form contacts)

```json
{
    "id":"widget1",
    "type":"contact",
    "value": ["f8ca557e"],
    "open_ids": ["ou_12345"]
}
```

> Associated approval 
  
  value is the instance code array of associated approvals

```json
{
    "id":"widget1",
    "type":"connect",
    "value": ["19EAC829-F1CB-527F-BE2A-1330422E60C0"]
}
```

> Attachment

 value is the code returned by Upload files

```json
{
    "id":"widget1",
    "type":"attachment",
    "value": ["D93653C3-2609-4EE0-8041-61DC1D84F0B5"]
}
```

```json
{
    "id":"widget1",
    "type":"attachmentV2",
    "value": ["D93653C3-2609-4EE0-8041-61DC1D84F0B5"]
}
```

> Image
 
  value is the code returned by Upload files

```json
{
    "id":"widget1",
    "type":"image",
    "value": ["D93653C3-2609-4EE0-8041-61DC1D84F0B5"]
}
```

```json
{
    "id":"widget1",
    "type":"imageV2",
    "value": ["D93653C3-2609-4EE0-8041-61DC1D84F0B5"]
}
```

> Multiple options
 
 value is the value field in the option field of the definition, obtained through View approval definition
 

```json
{
    "id":"widget1",
    "type":"checkbox",
    "value": ["optionA"]
}
```

```json
{
    "id":"widget1",
    "type":"checkboxV2",
    "value": ["k2b8mkx0-h71x5gljn3i-1"]
}
```

> Date range

   start and end must be in RFC3339 format
   
```json
{
    "id": "widget1",
    "type": "dateInterval",
    "value": {
         "start":"2019-10-01T08:12:01+08:00",
         "end":"2019-10-02T08:12:01+08:00",
         "interval": 2.0
     }
}
```

> Details

  value is a 2D array

```json
{
    "id": "widget1",
    "type": "fieldList",
    "value": [
         [   
            {
                "id": "widget1",
                "type": "checkbox",
                "value": ["jxpsebqp-0"]
            }
         ]
     ]
}
```
> telephone
```json
{
    "id":"widget1",
    "type":"telephone",
    "value": {
        "countryCode":"+86",
        "nationalNumber":"13122222222"
    }
}
```

> Shift swap widgets

  shiftTime and returnTime must be in RFC3339 format

```json
{
    "id": "widget1",
    "type": "shiftGroup",
    "value": {
         "shiftTime": "2019-10-01T08:12:01+08:00",
         "returnTime": "2019-10-02T08:12:01+08:00",
         "reason": "ask for leave"
     }
}
```

> Types created in the Open Platform are not supported 

Widget|Type
--|--
| Address| address|
|Business trip widgets|tripGroup|
|Leave widgets|leaveGroup|
|Overtime widgets|workGroup|
|Correction widgets|remedyGroup| ##  Response

###  Response body

|Parameter|Type|Required|Description|
|-|-|-|-|
|code|int|Yes|Error code. A value other than 0 indicates failure.|
|msg|String|Yes|Return code description.|
|data|map|Yes|Returned business information|
|  ∟instance_code|String|Yes|Approval instance Code| ###  Response body example

```json
{
    "code": 0,
    "msg": "success",
    "data": {
        "instance_code":"81D31358-93AF-92D6-7425-01A5D67C4E71"
    }
}
```
