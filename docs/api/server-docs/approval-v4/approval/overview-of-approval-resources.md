---
title: "Resource introduction"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/approval-v4/approval/overview-of-approval-resources"
service: "approval-v4"
resource: "approval"
section: "Approval"
updated: "1675167395000"
---

# Resource introduction
## Resource definition
A single approval flow is composed of forms and approval processes. After you create an approval, employees can fill in widget values when initiating an approval to form an approval instance.
## Field description
| Parameter | Type | Description |
| --- | --- | --- |
| `approval_name` | `string` | Internationalized text key of the approval title. It starts with @i18n@ and has a minimum length of 9 characters. **Example value**: "@i18n@approval_name" |
| `approval_code` | `string` | If this is left empty, it indicates creating a new approval. **Example value**: "7C468A54-8745-2245-9675-08B7C63E7A85" |
| `description` | `string` | Internationalized text key of the approval description. It starts with @i18n@ and has a minimum length of 9 characters. **Example value**: "@i18n@description" |
| `viewers` | `approval_create_viewers[]` | Visible scope of Approval Frontend. When the type is USER, one of the user_id or open_id needs to be filled in to specify which user is visible to; When the type is DEPARTMENT, you need to fill in the open_id of the department, which is used to specify which department is visible |
|   `viewer_type` | `string` | Viewer type **Example value**: "USER" **Optional values are**:  - `TENANT`: Users in a tenant - `DEPARTMENT`: Specified departments - `USER`: Specified users - `NONE`: No one  |
|   `viewer_user_id` | `string` | When view_type is USER, fill in the user id according to user_id_type **Example value**: "19a294c2" |
|   `viewer_department_id` | `string` | When view_type is DEPARTMENT, fill in the department id according to department_id_type **Example value**: "od-ac9d697abfa990b715dcc33d58a62a9d" |
| `form` | `approval_form` | Approval definition form. |
|   `form_content` | `string` | Approval definition form content, which is listed as a JSON array. **Example value**: "[{\ "ID":\ "user_name",\ "type":\ "input",\ "required": true,\ "name ":"@ i18n@widget1"}]"" |
| `node_list` | `approval_node[]` | Approval definition nodes, which are listed as the map array. The start node should be the first element of the list and the end node should be the last element. |
|   `id` | `string` | Node ID. The start node's ID is START, and the end node's ID is END. name, node_type, and approver are not required for the start and end nodes. **Example value**: "START" |
|   `name` | `string` | Internationalized text key of the node name. It starts with @i18n@ and has a minimum length of 9 characters. **Example value**: "@i18n@node_name" |
|   `node_type` | `string` | Approval type enumeration.When node_type is SEQUENTIAL, the approver must be selected by the initiator. **Example value**: "AND" **Optional values are**:  - `AND`: Everyone assigned - `OR`: Anyone assigned - `SEQUENTIAL`: Sequential  |
|   `approver` | `approval_approver_ccer[]` | List of reviewers |
|     `type` | `string` | 1.When the type is Supervisor, SupervisorTopDown, DepartmentManager , or DepartmentManagerTopDown, their level should be specified in the user_id. For example, for approval by third-level managers from bottom up, user_id is 3. 2.When the type is Personal, either user_id or open_id is required to specify users. 3.When the approver is Free (selected by the initiator), neither user_id nor open_id is required. ccer does not support Free. **Example value**: "Supervisor" **Optional values are**:  - `Supervisor`: Manager (bottom up) - `SupervisorTopDown`: Manager (top down) - `DepartmentManager`: Department supervisor (bottom up) - `DepartmentManagerTopDown`: Department supervisor (top down) - `Personal`: Specified member - `Free`: Selected by the initiator  |
|     `user_id` | `string` | User id, fill in according to user_id_type **Example value**: "f7cb567e" |
|     `level` | `string` | Approval level, when the type is Supervisor, SupervisorTopDown, DepartmentManager, DepartmentManagerTopDown, you need to fill in the corresponding level in the level, for example: review from the bottom to the top three levels, level = 3 **Example value**: "3" |
|   `ccer` | `approval_approver_ccer[]` | Copy sender list |
|     `type` | `string` | 1.When the type is Supervisor, SupervisorTopDown, DepartmentManager , or DepartmentManagerTopDown, their level should be specified in the user_id. For example, for approval by third-level managers from bottom up, user_id is 3. 2.When the type is Personal, either user_id or open_id is required to specify users. 3.When the approver is Free (selected by the initiator), neither user_id nor open_id is required. ccer does not support Free. **Example value**: "Supervisor" **Optional values are**:  - `Supervisor`: Manager (bottom up) - `SupervisorTopDown`: Manager (top down) - `DepartmentManager`: Department supervisor (bottom up) - `DepartmentManagerTopDown`: Department supervisor (top down) - `Personal`: Specified member - `Free`: Selected by the initiator  |
|     `user_id` | `string` | User id, fill in according to user_id_type **Example value**: "f7cb567e" |
|     `level` | `string` | Approval level, when the type is Supervisor, SupervisorTopDown, DepartmentManager, DepartmentManagerTopDown, you need to fill in the corresponding level in the level, for example: review from the bottom to the top three levels, level = 3 **Example value**: "3" |
|   `privilege_field` | `field_group` | Control permissions for table items |
|     `writable` | `string[]` | The id list of form entries with writable permissions **Example value**: 9293493 |
|     `readable` | `string[]` | The id list of table entries with readable permissions **Example value**: 9293493 |
| `settings` | `approval_setting` | Other settings of the approval definition. |
|   `revert_interval` | `int` | The period of time (in seconds) during which an approval instance can be canceled after it is approved. It defaults to 31 days. 0 indicates not cancelable. **Example value**: 0 |
|   `revert_option` | `int` | Whether an approval can be canceled after the first node. It defaults to 1. 0 indicates No. **Example value**: 0 |
| `config` | `approval_config` | Approval definition configuration items, used to configure whether an approval definition can be modified by users in the Approval Admin. |
|   `can_update_viewer` | `boolean` | Whether the visible scope can be modified. It defaults to false. **Example value**: False |
|   `can_update_form` | `boolean` | Whether the form content can be modified. It defaults to false. **Example value**: False |
|   `can_update_process` | `boolean` | Whether process nodes can be modified. It defaults to false. **Example value**: False |
|   `can_update_revert` | `boolean` | Whether the period of time for cancellation can be modified. It defaults to false. **Example value**: False |
|   `help_url` | `string` | Used to configure the help page link, by clicking which users can be redirected to the help page from the Approval Admin. **Example value**: "https://www.baidu.com" |
| `icon` | `int` | Approval icon enumeration. See the field description below for details. It defaults to 0. **Example value**: 0 **Default value**: `0` |
| `i18n_resources` | `i18n_resource[]` | Internationalized text. |
|   `locale` | `string` | Language **Example value**: "zh-CN" **Optional values are**:  - `zh-CN`: Chinese - `en-US`: English - `ja-JP`: Japanese  |
|   `texts` | `i18n_resource_text[]` | Key, value, and i18n key of copy starts with @i18n@.This field is mainly used for internationalized text. Word order users can pass texts in multiple languages at the same time, and the Approval Center will use the corresponding text according to the user's language environment. If the text of the user's language environment isn't passed, the text of default language will be used. |
|     `key` | `string` | Copywriting key **Example value**: "@i18n@1" |
|     `value` | `string` | Copywriting **Example value**: "People" |
|   `is_default` | `boolean` | Whether to use the default language. If the default language is used, all keys should be contained; if a non-default language is used but its keys do not exist, the default language will be used instead. **Example value**: true | ## Data example

```json
{
    "approval_name": "@i18n@approval_name",
    "approval_code": "813718CE-F38D-45CA-A5C1-ACF4F564B526",
    "viewers":[
        {
            "viewer_type":"TENANT",
            "viewer_user_id":""
        }
    ],
    "form": {
        "form_content": "[{\"id\":\"111\",\"name\":\"@i18n@event_name\",\"required\":true,\"type\":\"input\"},{\"id\":\"222\",\"name\":\"@i18n@time_interval\",\"required\":true,\"type\":\"dateInterval\",\"value\":{\"format\":\"YYYY-MM-DD hh:mm\",\"intervalAllowModify\":false}},{\"id\":\"333\",\"name\":\"@i18n@event_type\",\"type\":\"radioV2\",\"value\":[{\"key\":\"1\",\"text\":\"@i18n@recurrence_event\"},{\"key\":\"2\",\"text\":\"@i18n@single_event\"}]},{\"id\":\"444\",\"name\":\"@i18n@attende_count\",\"required\":true,\"type\":\"number\"},{\"id\":\"555\",\"name\":\"@i18n@apply_reason\",\"required\":true,\"type\":\"textarea\"}]"
        },

    "node_list": [{
          "id": "START",
          "privilege_field":{ 
				 "writable": ["111","222"],
				 "readable": ["111","222"]
		  }
        },{
          "id": "7106864726566",
          "privilege_field":{ 
				 "writable": ["111","222"],
				 "readable": ["111","222"]
		  },
          "name": "@i18n@node_name",
          "node_type": "AND",
          "approver": [
            {
              "type": "Personal",
              "user_id": "59a92c4a"
            }
          ],
          "ccer": [
            {
              "type": "Supervisor",
              "level": "2"
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
          "can_update_form": true,
          "can_update_process": true,
          "can_update_revert": true,
          "help_url":"https://www.baidu.com"
        },
    "icon": 1,
    "i18n_resources" : [{
          "locale": "zh-CN",
          "texts" : [
              {"key":"@i18n@approval_name","value":"approval_name"},
              {"key":"@i18n@event_name","value":"event_name"},
              {"key":"@i18n@node_name","value":"node_name"},
              {"key":"@i18n@time_interval","value":"time_interval"},
              {"key":"@i18n@event_type","value":"event_type"},
              {"key":"@i18n@recurrence_event","value":"ecurrence_event"},
              {"key":"@i18n@single_event","value":"single_event"},
              {"key":"@i18n@attende_count","value":"attende_count"},
              {"key":"@i18n@apply_reason","value":"apply_reason"}
            ],
          "is_default": true
        }]
}
```

## What is approval code

The unique code for an approval definition, which can be found in the browser URL when you edit an approval. Click the following link to enter the Admin console:

[https://www.larksuite.com/approval/admin/approvalList?devMode=on](https://www.larksuite.com/approval/admin/approvalList?devMode=on)

![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/af9b32b0db7014aabf9c3acc6a9925d3_j6dVkJuMM7.png?lazyload=true&width=1000&height=132)

Click "Edit" in the image above to find the Approval Code in the URL, with the format: `definitionCode=E3254848-D172-4169-B03E-744E7CD11F06`.

## What is form

Forms contain one or more widgets. When an employee initiates an approval, they must fill out the form page. Each approval has its own form.

### 1.Widget

A constituent part of an approval. A form is comprised of one or more approval widgets. Each widget has a basic field:

```js 
struct{
    String Id; //A unique ID that identifies a widget in an approval.
    String CustomId; //A custom widget ID, a unique ID that identifies a widget in an approval.
    String Name; //The name of this widget
    String Type; //The type of this widget
} 
```
### 2.Custom Widget ID

You can specify an ID when editing an approval definition on the Approval Admin. You can subsequently specify this ID to identify the widget in the approval definition.

### 3.Widget description

widget type | description|
| --------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Single-line text (input)| Used to input single-line text |
| Multi-line text (textarea)| Used to input multi-line text|
| Number (number) | Used to input a number.|
| Amount (amount) | Used to input the approval amount and unit (default unit: CNY).|
| Date (date) | Used to input a specific time. |
| Date interval (dateInterval)| Used to input a date interval, including a start time, end time, and duration. |
| Single option (radio, radioV2)| Used to select a single option.|
| Multiple options (checkbox, checkboxV2) | Used to select multiple options. |
| Address (address) | Used to input an address.|
| Contact (contact) | Used to add contacts in the approval.|
| Description (text)| Used to add a description in the approval definition (such as specifications or precautions). This can't be edited when an approval is initiated.|
| Details (fieldList) | Used to input details information. In Details widget, you can add other widgets such as number and amount. When creating an approval definition, design a Details widget to specify what widgets are included in an entry. Initiators can add entries according to their own demands, and each entry is consistent with the initially configured Details widget. |
| Formula (formula) | When creating an approval definition, design a formula to specify how the value of this widget is calculated from other widgets (numbers and amounts). |
| Image (image, imageV2)| Used to add an image in the approval.|
| Attachment (attachment, attachmentV2) | Used to add an attachment, such as a file, to the approval.|
| Associate (connect) | Used to associate to another approval in the current approval. This allows approvers to see information for the associated approval when reviewing the current approval. |
| Leave (leaveGroup)| Used to enter leave approval content, such as to select a leave type (sick leave, maternity leave, etc.; leave types must be set by administrators in Leave Management and selected when initiating an instance for the current approval definition) and enter the leave start time, end time, and leave duration. |
| Overtime (workGroup)| Used to enter overtime request content, such as to select the overtime type (compensatory leave, paid overtime, unpaid overtime, etc.; types must be set by the administrator when creating the approval definition) and enter the overtime start and end times, duration, and reason for overtime.|
| Shift swap (shiftGroup) | Used to enter content for a shift swap request, including the original shifts and reason for the swap. |
| Business trip (tripGroup) | Used to enter business trip approval content, including the trip itinerary (including the itinerary start time, itinerary end time, itinerary duration, departure location, destination location, transportation method, one-way/round-trip, and notes; the itinerary duration is calculated in calendar days; if a trip has multiple itineraries, the initiator can manually add additional itineraries), total trip duration, trip reason, and traveling with. |
| Correction (remedyGroup)| This is used to enter correction approval content, including the correction time and reason for correction.

## What is approval process (Process)

When creating an approval definition, an administrator will design an approval transfer mode, which specifies every step of an instance created by this definition. An approval definition corresponds to an approval process, which is made up of approval nodes. That is, an approval process corresponds to one or more approval nodes. When an employee initiates an approval instance, the instance enters the approval process as designed. When designing an approval process, you must specify the approval actions for each step (For approval nodes, see below for details). This includes the approvers for each step or the information for auto-approval or auto-rejection by the system. If the current step is not automatically performed by the system, approvers need to be specified or selected by the initiator.

### 1.Approvers selected by initiator

If an approver isn't specified for a certain step in the approval definition and "Approver selected by initiator" is set instead, the user must select an approver when initiating an approval instance. When creating such an approval definition, you only need to set if approval is required by everyone assigned or anyone assigned. **When "Approver selected by initiator" is set, the user must specify an approver when initiating an approval instance through the API**.  

### 2.Approval node

Each approval process is composed of one or more approval nodes. Each node specifies a step in the approval flow. After the current node is completed, the approval advances to the next node. When an instance is created, the corresponding approval nodes are generated to form an approval process.

### 3.Custom node ID

Similar to a widget, you can specify a custom node ID. This is a unique identifier used to specify this node in an approval process.

##  User ID description
For how to distinguish between and use user_id, open_id, and union_id, refer to User-related IDs.

##  Department ID description
For how to distinguish between and use department_id,open_department_id,refer to        Department Overview
