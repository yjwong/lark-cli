---
title: "Obtain Single Approval Instance Details"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uEDNyUjLxQjM14SM0ITN"
method: "POST"
api_path: "https://www.larksuite.com/approval/openapi/v2/instance/get"
section: "Approval"
scopes:
  - "approval:approval:readonly"
updated: "1705410962000"
---

# Obtain details of a single approval instance
> This API has been upgraded to enhance its security. We recommend that you migrate to the new version>>

This API is used to obtain approval instance details through Instance Code. You can obtain Instance Code via the Obtain approval instances in batches API.

## Request
| HTTP URL | https://www.larksuite.com/approval/openapi/v2/instance/get |
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
| instance_code | String | Yes | Approval instance code. If you pass in the uuid when creating the approval instance, you can also obtain the code by passing in the uuid.|
| locale | String | No | zh-CN - Chineseen-US - Englishja-JP - Japanese |
| user_id | String | No | User ID of the approval initiator, which is used for platform-level approval |
| open_id | String | No | Open ID of the approval initiator|
### Request body example

```json
{
    "instance_code": "81D31358-93AF-92D6-7425-01A5D67C4E71",
    "locale": "zh-CN",
    "open_id": "xxxxx",
    "user_id": "xxxxxxx"
}
```

## Response

### Response body
|Parameter|Type|Required|Description|
|-|-|-|-|
|code|int|Yes|Error code. A value other than 0 indicates failure.|
|msg|String|Yes|Return code description|
|data|map|Yes|Returned business information|
|  ∟approval_code | String | Yes | Approval definition code |
|  ∟approval_name|String|Yes|Approval title|
|  ∟start_time | int | Yes | Approval creation time |
|  ∟end_time | int | Yes | Approval completion time. A value of 0 indicates that the approval is not completed. |
|  ∟user_id | String | Yes | Approval initiator |
|  ∟open_id | String | Yes | Open ID of the approval initiator | 
|  ∟serial_number | String | Yes | Approval ticket number | 
|  ∟department_id | String | Yes | Department of the approval initiator |
|  ∟status | String | Yes | Approval instance statusPENDING - Under reviewAPPROVED - ApprovedREJECTED - RejectedCANCELED - CanceledDELETED - Deleted |
|   ∟uuid | String | Yes | Unique ID of the user |
|  ∟form | String | Yes | JSON string, **widget value**|
|  ∟id | String | Yes | Widget ID |
|  ∟custom_id | String | No | Custom widget ID. This field is not returned if it is not set. |
|    ∟name|String|Yes|Widget name|
|    ∟type|String|Yes|Widget type|
|  ∟value | String | Yes | Widget value. The values of different types use different formats. |
|  ∟ext | object | Yes | Extension field of the widget, which is used for some special widget values. For example, the ext value of an image widget is the file name. |
|   ∟task_list | list | Yes | Approval task list |
|  ∟id | String | Yes | Task ID |
|  ∟user_id | string | No | ApproverTask user_id is empty for auto-approve and auto-reject tasks. |
|  ∟open_id | string | No | Open ID of the approver |
|  ∟status | string | Yes | Task statusPENDING - Under reviewAPPROVED - ApprovedREJECTED - RejectedTRANSFERRED - TransferredDONE - Completed |
|  ∟node_id | String | No | ID of the node to which the task belongs |
|  ∟node_name | String | No | Name of the node to which the task belongs |
|  ∟custom_node_id | String | No | Custom ID of the node to which the task belongs. This field is not returned if it is not set. |
|  ∟type | String | Yes | Approval methodAND - Everyone assignedOR - Anyone assignedAUTO_PASS - Auto-approveAUTO_REJECT - Auto-rejectSEQUENTIAL - In order |
|  ∟start_time | int | Yes | Task start time |
|  ∟end_time | int | Yes | Task completion time. A value of 0 indicates that the task is not completed. |
|  ∟comment_list | list | Yes | Comment list |
|  ∟id | String | Yes | Comment ID |
|  ∟user_id | String | Yes | User who posted the comment |
|  ∟open_id | String | Yes | Open ID of user who posted the comment |
|  ∟comment | String | Yes | Comment content |
|  ∟create_time | int | Yes | Comment time |
|  ∟timeline | list | Yes | Approval update |
|  ∟type | String | Yes | Update type. The user_id_list value has different meanings for different ext types.START - Start reviewPASS - ApproveREJECT - RejectAUTO_PASS - Auto-approveAUTO_REJECT - Auto-rejectREMOVE_REPEAT - DeduplicateTRANSFER - TransferADD_APPROVER_BEFORE - Add pre-approverADD_APPROVER - Add co-approverADD_APPROVER_AFTER - Add post-approverDELETE_APPROVER - Remove approverROLLBACK_SELECTED - Roll back specified requestsROLLBACK - Roll back all requestsCANCEL - CancelDELETE - DeleteCC - CC |
|  ∟create_time | int | Yes | Occurrence time |
|  ∟user_id | String | No | User who generated the update |
|  ∟open_id | String | No | Open ID of the user who generated the update |
|  ∟user_id_list | list | No | CC recipient list |
|  ∟open_id_list | list | No | CC recipient list |
|  ∟task_id | String | No | task_id associated with the update |
|  ∟comment | String | No | Reason |
|  ∟cc_user_list | list | No | CC sender list |
|  ∟user_id | String | No | User ID of the CC sender |
|  ∟cc_id | String | No | Unique CC ID in the approval instance |
|  ∟open_id | String | No | Open ID of the CC sender |
|  ∟ext | String | Yes | Additional update information, including user_id_list and user_id |
|  ∟user_id_list | list | No |**Type** - **user_id_list meaning**TRANSFER - Transfer recipientADD_APPROVER_BEFORE - Added approverADD_APPROVER - Added approverADD_APPROVER_AFTER - Added approverDELETE_APPROVER - Removed approver |
|  ∟open_id_list | list | No | Open ID corresponding to the user_id_list|
|  ∟user_id | String | No |**Type** - **user_id meaning**CC - CC sender |
|  ∟open_id | String | No | open_id corresponding to the user_id |
|  ∟node_key | string | No | Key of the node that generated the task |
|  ∟modified_instance_code|string|否|modify original instance code,Display this field only when searching modified instances|
|  ∟reverted_instance_code|string|否|revert original instance code,Display this field only when searching revoked instances| **Widget value**: 
|Type|Description|
|-|-|
|input|--|
|textarea|--|
| date | RFC3339 format, for example, 2019-10-01T08:12:01+08:00 |
| radio/radioV2 | Text field in "option" |
| address | China/Beijing/Beijing/Chaoyang Qu/chang an jieIf the address widget allows, the last item is the detailed address entered by the user. |
```json
{
    "id": "widget1",
    "custom_id": "user_info",
    "name": "Item application",
    "type": "input",
    "value": "data"
}
```
|Type|Description|
|-|-|
|number|--|
|amount|--|
|formula|--|
```json
{
    "id": "widget1",
    "name": "Item application",
    "type": "number",
    "value": 1234.56
}
```
|Type|Description|
|-|-|
|contact|`user_id`|
```json
{
    "id": "widget1",
    "name": "Item application",
    "type": "contact",
    "value": ["f8ca557e"],
    "open_ids": ["ou_12345"]
}
```

|Type|Description|
|-|-|
|connect|`instance_code`|
| attachmentV2 | Attachment widget. Attachment names in the ext:value are separated by commas. | The link remains valid for 12 hours.
| image/imageV2 | Image widget. Image names in the ext:value are separated by commas. | The link remains valid for 12 hours.

```json
{
    "id": "widget1",
    "name": "Item application",
    "type": "attachmentV2",
    "ext": "'Item 1 name','Item 2 name'",
    "value": ["Item 1", "Item 2"]
}
```

|Type|Description|
|-|-|
|connect|`instance_code`|
| attachment | Attachment widget. It is recommended to use attachmentV2. |
| checkbox/checkboxV2 | Text field in "option" | ```json
{
    "id": "widget1",
    "name": "Item application",
    "type": "checkbox",
    "value": ["Item 1"]
}
```
|Type|Description|
|-|-|
| dateInterval | "start" and "end" must be in RFC3339 format. |
```json
{
    "id": "widget1",
    "name": "Item application",
    "type": "dateInterval",
    "value": {
         "start": "2019-10-01T08:12:01+08:00",
         "end": "2019-10-02T08:12:01+08:00",
         "interval": 2.0
     }
}
```

|Type|Description|
|-|-|
| fieldList | Value is a 2D array. | ```json
{
    "id": "widget1",
    "name": "Item application",
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

|Type|Description|
|-|-|
|document|--| ```json
{
    "id":"widget1",
    "type":"document",
    "value":
        {
          "token": "doxcx7B8OzLFHExkiwYuPGAwf",
          "type": "doc",
          "title": "title",
          "url": "https://bytedance.larksuite.com/docx/doxcx7B8OzLFHExkiwYuPGAwf"
       }
}
```

|Type|Description|
|-|-|
|department|--| ```json
{
    "id":"widget1",
    "type":"department",
    "value":[
        {
            "open_id":"od-xxx"
        }
    ]
}
```
|Field|Description|
|-|-|
| id | Widget ID |
| name | Widget name |
|type|leaveGroup|
|value|--|
|  ∟name | Leave name |
|  ∟start | Start time, in RFC3339 format |
|  ∟end | End time, in RFC3339 format |
|  ∟interval | Duration (in days). Users can manually enter the leave duration for some leave types. | ```json
{
    "id": "widget1",
    "name": "leave",
    "type": "leaveGroup",
    "value": {
         "name": "annual leave"
         "start": "2019-10-01T00:00:00+08:00",
         "end": "2019-10-02T00:0:00+08:00",
         "interval": 2.0
     }
}
```

|Field|Description|
|-|-|
| id | Widget ID |
| name | Widget name |
|type|leaveGroupV2|
|value|--|
|  ∟name | Leave name |
|  ∟start | Start time, in RFC3339 format |
|  ∟end | End time, in RFC3339 format |
|  ∟interval | Duration. Users can manually enter the leave duration for some leave types. |
|  ∟unit | Duration unit (day/hour).Users can manually enter the leave duration for some leave types. |
|  ∟reason | Reason | ```json
{
    "id": "widget1",
    "name": "leave",
    "type": "widgetLeaveGroupV2",
    "value": {
         "name": "annual leave"
         "start": "2019-10-01T00:00:00+08:00",
         "end": "2019-10-02T00:0:00+08:00",
         "interval": 2.0,
         "unit": "DAY",
         "reason": "out going"
     }
}
```

|Type|Description|
|-|-|
| remedyGroup | "time" must be in RFC3339 format. | ```json
{
    "id": "widget1",
    "name": "remedy",
    "type": "remedyGroup",
    "value": {
         "time": "2019-10-01T08:12:01+08:00",
         "reason": "forgot"
     }
}
```

|Type|Description|
|-|-|
| shiftGroup | shiftTime and returnTime must be in RFC3339 format. | ```json
{
    "id": "widget1",
    "name": "shift",
    "type": "shiftGroup",
    "value": {
         "shiftTime": "2019-10-01T08:12:01+08:00",
         "returnTime": "2019-10-02T08:12:01+08:00",
         "reason": "ask for leave"
     }
}
```

|Type|Description|
|-|-|
| workGroup | "start" and "end" must be in RFC3339 format. | ```json
{
    "id": "widget1",
    "name": "work",
    "type": "workGroup",
    "value": {
         "name": "Overtime pay"
         "start": "2019-10-01T08:12:01+08:00",
         "end": "2019-10-02T08:12:01+08:00",
         "interval": 2.0,
         "reason": "ask for leave"
     }
}
```

|Type|Description|
|-|-|
| tripGroup | "start" and "end" must be in RFC3339 format. | ```json
{
    "id": "widget1",
    "name": "trip",
    "type": "tripGroup",
    "value": {
         "schedule": [{
                "start": "2019-10-0T00:00:00Z+08:00",
                "end": "2019-10-01T00:00:00Z+08:00",
                "interval": 123.45,
                "departure": "China/Beijing/Beijing",
                "destination": "China/Shanghai/Shanghai",
                "transportation": "Airplane",
                "oneRound": "One Way",
                "remark": "business",
         }],
         "interval": 2.0,
         "reason": "business",
         "peer": ["f7cb567e"],
    }
}
```
|Type|Description|
|-|-|
|telephone|telephone|
```json
{
    "id":"widget1",
    "type":"telephone",
    "value": {
        "country_code":"+86",
        "national_number":"13122222222"
    }
}
```
### Response body example

```json
Note: The current approver can't be added.
{
    "code": 0,
    "msg": "success",
    "data": {
        "approval_code": "7C468A54-8745-2245-9675-08B7C63E7A85",
        "approval_name": "Payment",
        "start_time": 1564590532967,
        "end_time": 0,
        "user_id": "f3ta757q",
        "open_id": "ou_123456",
        "serial_number":"202102060002",
        "department_id": "od-8ec33ffec336c3a39a278bc25e931676",
        "status": "PENDING",
        "uuid": "xxxxx",
        "form": "[{\"id\": \"widget1\",\"custom_id\": \"user_info\",\"name\": \"Item application\",\"type\": \"textarea\"},\"value\":\"aaaa\"]",
        "task_list":[
          {
            "id": "1234",
            "node_id": "46e6d96cfa756980907209209ec03b64",
            "node_name": "Start",
            "custom_node_id": "manager",
            "user_id": "f7cb567e",
            "open_id": "ou_123457",
            "type": "AND",
            "status": "PENDING",
            "start_time": 1564590532967,
            "end_time": 0
          }
        ],
        "comment_list": [
          {
            "id": "1234",
            "user_id": "f7cb567e",
            "open_id": "ou_123456",
            "comment": "ok",
            "create_time": 1564590532967
          }
        ], 
        "timeline": [
          {
            "type": "START",
            "user_id": "f8ca557e",
            "open_id": "ou_123456",
            "create_time": 1564590532967,
            "comment": "ok",
            "task_id": "1234",
            "cc_user_list": [],
            "ext": "{\"user_id_list\":[\"f7cb567e\"],\"open_id_list\":[\"ou_123456\"]}"
         },
         {
            "type": "CC",
            "user_id_list": [
               "eea5gefe"
            ],
            "open_id_list": [
               "ou_12345"
            ],
            "create_time": 1571913568556,
            "comment": "ccc",
            "ext": "{\"user_id\":\"62d4a44c\",\"open_id\":\"ou_123456\"}"
         },
         {
            "comment": "Defensive play",
            "create_time": 1636616275751,
            "node_key": "APPROVAL_240330_4058663",
            "open_id": "ou_456",
            "task_id": "789",
            "type": "PASS",
            "user_id": "f6bac64a"
           }
       ]
    }
}
```
