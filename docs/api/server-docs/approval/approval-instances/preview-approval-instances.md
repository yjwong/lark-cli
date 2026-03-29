---
title: "Preview approval instances"
url: "https://open.larksuite.com/document/ukTMukTMukTM/ukTM5UjL5ETO14SOxkTN/approval-preview"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/approval/v4/instances/preview"
section: "Approval"
scopes:
  - "approval:approval:readonly"
updated: "1675167401000"
---

# Preview approval instances

This API is used to preview the approval process before submitting an application. You can also use it to preview the remaining process after an approval is initiated.

##  Request
| HTTP URL | https://open.larksuite.com/open-apis/approval/v4/instances/preview |
| --- | --- |
| HTTP Method | POST |
| Supported app types | custom,isv |
| Required scopes | `approval:approval:readonly` | ###  Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes  `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" Learn more about how to obtain and use access_token |  |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `user_id_type` | `string` | No | User  ID  type **Example value**: "open_id" **Optional values are**: - `open_id`: User's  open id - `union_id`: User's  union id - `user_id`: User's  user id **Default value**: `open_id` **When the value is  `user_id`, the required field scopes are**: `contact:user.employee_id:readonly` | ### Request body
|Parameter|Type|Required|Description|
|-|-|-|-|
|approval_code|string|No|Approval definition code|
|user_id|string|Yes|Approval initiator, employee ID or Open ID|
|department_id|string|No|The department ID of the approval initiator. If the user only belongs to 1 department, this field can be left empty. If the user belongs to multiple departments, you must specify a department.|
|form|string|No|JSON string, widget value. This field is required when previewing the process before submitting an application.|
|  |∟id|string|Yes|Widget ID, or you can use the custom_id|
|  ∟type|string|Yes|Widget type|
|  ∟value|string|Yes|Widget value. The values of different types use different formats.|
|instance_code|string|No|Approval instance code|
|task_id|string|No|If the approval instance already exists, provide the task_id of the current approval task. The user_id of the task assignor also needs to be passed.| #### Description:

-  The request parameters are similar to those of the "Create an approval instance" API (except there is no uuid field). If you haven't initiated an approval instance, pass the values similar to the "Create Approval Instance" API, such as user_id, form, and approval definition approval_code. If you have initiated an approval instance, the process is fixed, and only user_id, instance_code, and task_id are required.

### Request body example

case1: Before approval initiation

```json
{
    "approval_code":"C2CAAA90-70D9-3214-906B-B6FFF947F00D",
    "user_id":"f7cb567e",
    "department_id":"",
    "form":"[{\"id\":\"widget16256287451710001\", \"type\": \"number\", \"value\":\"43\"}]"

}
```

case2:After approval initiation

```json
{
    "instance_code":"12345CA6-97AC-32BB-8231-47C33FFFCCFD",
    "user_id":"f7cb567e",
    "task_id": "6982332863116876308"
}
```

##  Response

###  Response body

|Parameter|Type|Required|Description|
|-|-|-|-|
|code|int|Yes|Error code. A value other than 0 indicates failure.|
|msg|string|Yes|Return code description|
|data|json|Yes|Returned business information|
|∟preview_nodes|list|Yes|Preview node information|
|  ∟user_id_list|list|Yes|Approver ID list|
|  ∟end_cc_id_list| list| Yes  |Approval end CC sender ID list|
|  ∟node_id| string | Yes  |Node ID|
|  ∟node_name     | string | Yes  |Node name|
|  ∟node_type      | string | Yes  |Node type: AND: Everyone assignedOR: Anyone assigned|
|  ∟custom_node_id | string | Yes  |Custom node ID|
|  ∟comments       | list   | Yes  |Node description|
|  ∟is_empty_logic | bool   | Yes  |Whether the approver field is empty. If empty, the user_id_list is used to get the approver ID list.|
|  ∟is_approver_type_free | bool   | Yes  |Whether it is a node selected by the initiator|
|  ∟has_cc_type_free | bool   | Yes  |Whether CC senders can select nodes | ###  Response body example

```json
{ 
    "code":0, 
    "msg":"success", 
    "data": { 
        "preview_nodes":[
            {
                "user_id_list":["ffffffff"],
                "end_cc_id_list":[],
                "node_id":"b078ffd28db767c502ac367053f6e0ac",
                "node_name":"Initiation",
                "node_type":"",
                "comments":[],
                "custom_node_id":""
            },
            {
                "user_id_list":["ffffffff"],
                "end_cc_id_list":[],
                "node_id":"e6ce10282a3cc3bf4a408feffd678dcf",
                "node_name":"Approval",
                "node_type":"AND",
                "comments":[],
                "custom_node_id":"",
                "is_empty_logic":false,
                "is_approver_type_free":false,
                "has_cc_type_free":false
            },
            {
                "user_id_list":[],
                "end_cc_id_list":[],
                "node_id":"b1a326c06d88bf042f73d70f50197905",
                "node_name":"End",
                "node_type":"",
                "comments":[],
                "custom_node_id":""
             }
         ]
    }        
}  
```

###
