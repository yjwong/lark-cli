---
title: "Create an approval definition"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/approval-v4/approval/create"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/approval/v4/approvals"
service: "approval-v4"
resource: "approval"
section: "Approval"
scopes:
  - "approval:approval"
  - "approval:definition"
field_scopes:
  - "contact:user.employee_id:readonly"
updated: "1675167395000"
---

# Create an approval definition

This API is used to create a simple approval definition, whose basic information, form and process can be specified. This API is not recommended for custom apps. Contact the administrator to create definitions in the Approval Admin if necessary.

> Please be cautious when calling the API. Created approval definitions cannot be disabled or deleted.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/approval/v4/approvals |
| HTTP Method | POST |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `approval:approval` `approval:definition` |
| Required field scopes | > The response body of the API contains the following sensitive fields, and they will be returned only after corresponding scopes are added. If you do not need the fields, it is not recommended that you request the scopes. `contact:user.employee_id:readonly` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `department_id_type` | `string` | No | The type of department ID used in this call **Example value**: "open_department_id" **Optional values are**:  - `department_id`: Identify departments with custom department_id - `open_department_id`: Identify departments by open_department_id  |
| `user_id_type` | `string` | No | User ID categories **Example value**: "open_id" **Optional values are**:  - `open_id`: Identifies a user to an app. The same user has different Open IDs in different apps. How to get Open ID - `union_id`: Identifies a user to a tenant that acts as a developer. A user has the same Union ID in apps developed by the same developer, and has different Union IDs in apps developed by different developers. A developer can use Union ID to link the same user's identities in multiple apps.How to get Union ID - `user_id`: Identifies a user to a tenant. The same user has different User IDs in different tenants. In one single tenant, a user has the same User ID in all apps （including store apps）. User ID is usually used to communicate user data between different apps. How to get User ID  **Default value**: `open_id` **When the value is `user_id`, the following field scopes are required**: `contact:user.employee_id:readonly` | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `approval_name` | `string` | Yes | Internationalized text key of the approval title. It starts with @i18n@ and has a minimum length of 9 characters. **Example value**: "@i18n@approval_name" |
| `approval_code` | `string` | No | If this is left empty, it indicates creating a new approval. **Example value**: "7C468A54-8745-2245-9675-08B7C63E7A85" |
| `description` | `string` | No | Internationalized text key of the approval description. It starts with @i18n@ and has a minimum length of 9 characters. **Example value**: "@i18n@description" |
| `viewers` | `approval_create_viewers[]` | Yes | Visible scope of Approval Frontend. 1. When viewer_type is USER, viewer_user_id need to be filled in 2. When viewer_type is DEPARTMENT,viewer_department_id need to be filled in 3. When viewer_type is TENANT or NONE, viewer_user_id and viewer_department_id do not need to be filled in |
|   `viewer_type` | `string` | No | Viewer type **Example value**: "USER" **Optional values are**:  - `TENANT`: Users in a tenant - `DEPARTMENT`: Specified departments - `USER`: Specified users - `NONE`: No one  |
|   `viewer_user_id` | `string` | No | When viewer_type is USER, fill in the user id according to user_id_type **Example value**: "19a294c2" |
|   `viewer_department_id` | `string` | No | When viewer_type is DEPARTMENT, fill in the department id according to department_id_type **Example value**: "od-ac9d697abfa990b715dcc33d58a62a9d" |
| `form` | `approval_form` | Yes | Approval definition form. |
|   `form_content` | `string` | Yes | Approval definition form content, which is listed as a JSON array. See the field description below for details. **Example value**: "[{\ "ID\":\ "user_name\",\ "type\":\ "input\",\ "required\": true,\ "name \":\"@ i18n@widget1\"}]" |
| `node_list` | `approval_node[]` | Yes | Approval definition nodes, which are listed as the map array. The start node should be the first element of the list and the end node should be the last element. |
|   `id` | `string` | Yes | Node ID. The start node's ID is START, and the end node's ID is END. name, node_type, and approver are not required for the start and end nodes. **Example value**: "START" |
|   `name` | `string` | No | Internationalized text key of the node name. It starts with @i18n@ and has a minimum length of 9 characters. **Example value**: "@i18n@node_name" |
|   `node_type` | `string` | No | Approval type enumeration. When node_type is SEQUENTIAL, the approver must be selected by the initiator. **Example value**: "AND" **Optional values are**:  - `AND`: Everyone assigned - `OR`: Anyone assigned - `SEQUENTIAL`: Sequential  |
|   `approver` | `approval_approver_ccer[]` | No | List of reviewers |
|     `type` | `string` | Yes | Approver/Cc Type 1. When the type is Supervisor, SupervisorTopDown, DepartmentManager , or DepartmentManagerTopDown, their level should be specified in the level. For example, for approval by third-level managers from bottom up, level is 3. 2. When the type is Personal, User id  fill in according to user_id_type 3. When the approver is Free (selected by the initiator), neither user_id nor level is required. 4. ccer does not support Free. **Example value**: "Supervisor" **Optional values are**:  - `Supervisor`: Manager (bottom up) - `SupervisorTopDown`: Manager (top down) - `DepartmentManager`: Department supervisor (bottom up) - `DepartmentManagerTopDown`: Department supervisor (top down) - `Personal`: Specified member - `Free`: Selected by the initiator  |
|     `user_id` | `string` | No | User id, fill in according to user_id_type **Example value**: "f7cb567e" |
|     `level` | `string` | No | Approval level, when the type is Supervisor, SupervisorTopDown, DepartmentManager, DepartmentManagerTopDown, you need to fill in the corresponding level in the level, for example: review from the bottom to the top three levels, level = 3 **Example value**: "3" |
|   `ccer` | `approval_approver_ccer[]` | No | Copy sender list |
|     `type` | `string` | Yes | Approver/Cc Type 1. When the type is Supervisor, SupervisorTopDown, DepartmentManager , or DepartmentManagerTopDown, their level should be specified in the level. For example, for approval by third-level managers from bottom up, level is 3. 2. When the type is Personal, User id  fill in according to user_id_type 3. When the approver is Free (selected by the initiator), neither user_id nor level is required. 4. ccer does not support Free. **Example value**: "Supervisor" **Optional values are**:  - `Supervisor`: Manager (bottom up) - `SupervisorTopDown`: Manager (top down) - `DepartmentManager`: Department supervisor (bottom up) - `DepartmentManagerTopDown`: Department supervisor (top down) - `Personal`: Specified member - `Free`: Selected by the initiator  |
|     `user_id` | `string` | No | User id, fill in according to user_id_type **Example value**: "f7cb567e" |
|     `level` | `string` | No | Approval level, when the type is Supervisor, SupervisorTopDown, DepartmentManager, DepartmentManagerTopDown, you need to fill in the corresponding level in the level, for example: review from the bottom to the top three levels, level = 3 **Example value**: "3" |
|   `privilege_field` | `field_group` | No | Control permissions for table items |
|     `writable` | `string[]` | Yes | The id list of form entries with writable permissions **Example value**: 9293493 |
|     `readable` | `string[]` | Yes | The id list of table entries with readable permissions **Example value**: 9293493 |
| `settings` | `approval_setting` | No | Other settings of the approval definition. |
|   `revert_interval` | `int` | No | The period of time (in seconds) during which an approval instance can be canceled after it is approved. It defaults to 31 days. 0 indicates not cancelable. **Example value**: 0 |
|   `revert_option` | `int` | No | Whether an approval can be canceled after the first node. It defaults to 1. 0 indicates No. **Example value**: 0 |
| `config` | `approval_config` | No | Approval definition configuration items, used to configure whether an approval definition can be modified by users in the Approval Admin. |
|   `can_update_viewer` | `boolean` | Yes | Whether the visible scope can be modified. It defaults to false. **Example value**: False |
|   `can_update_form` | `boolean` | Yes | Whether the form content can be modified. It defaults to false. **Example value**: False |
|   `can_update_process` | `boolean` | Yes | Whether process nodes can be modified. It defaults to false. **Example value**: False |
|   `can_update_revert` | `boolean` | Yes | Whether the period of time for cancellation can be modified. It defaults to false. **Example value**: False |
|   `help_url` | `string` | No | Used to configure the help page link, by clicking which users can be redirected to the help page from the Approval Admin. **Example value**: "https://www.baidu.com" |
| `icon` | `int` | No | Approval icon enumeration. See the field description below for details. It defaults to 0. **Example value**: 0 **Default value**: `0` |
| `i18n_resources` | `i18n_resource[]` | Yes | Internationalized text. |
|   `locale` | `string` | Yes | Language **Example value**: "zh-CN" **Optional values are**:  - `zh-CN`: Chinese - `en-US`: English - `ja-JP`: Japanese  |
|   `texts` | `i18n_resource_text[]` | Yes | Key, value, and i18n key of copy starts with @i18n@. This field is mainly used for internationalized text. Word order users can pass texts in multiple languages at the same time, and the Approval Center will use the corresponding text according to the user's language environment. If the text of the user's language environment isn't passed, the text of default language will be used. |
|     `key` | `string` | Yes | Copywriting key **Example value**: "@i18n@1" |
|     `value` | `string` | Yes | Copywriting **Example value**: "People" |
|   `is_default` | `boolean` | Yes | Whether to use the default language. If the default language is used, all keys should be contained; if a non-default language is used but its keys do not exist, the default language will be used instead. **Example value**: true |
| `process_manager_ids` | `string[]` | No | Process manager ids, fill in according to user_id_type **Example value**: ["1c5ea995"] | ### Request body example

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
              {"key":"@i18n@recurrence_event","value":"recurrence_event"},
              {"key":"@i18n@single_event","value":"single_event"},
              {"key":"@i18n@attende_count","value":"attende_count"},
              {"key":"@i18n@apply_reason","value":"apply_reason"}
            ],
          "is_default": true
        }],
    "process_manager_ids": ["1c5ea995"]
}

## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `approval_code` | `string` | Review Definition Code |
|   `approval_id` | `string` | Review definition id | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
        "approval_code": "81D31358-93AF-92D6-7425-01A5D67C4E71",
        "approval_id": "7090754740375519252"
    }
}

## Field description 

### viewers field description

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
You can set default values for the Input, Text Area, Radio V2, Contact, and Department widgets when you create the approval definition in Lark Approval. 

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
    "value":"CNY",
    "option":{ "currencyRange": ["CNY","USD"]
    }
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
    "type": "fieldList",
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

### Icon enumeration
 
The following are icons 0-24 from left to right and top to bottom.

![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/3815d7910d3d0c2671eea08f38aa80a9_eoxzGr2c1b.png)

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 1390001 | param is invalid | Parameter error |
| 400 | 1390002 | approval code not found | Check that the review definition code is correct |
| 400 | 1390003 | instance code not found | Check that the review instance code is correct |
| 400 | 1390004 | user_id or open_id not found | Check user_id open_id correct |
| 403 | 1390009 | no operation permission | Check that the operation permissions are correct |
| 400 | 1390013 | unsupported approval for free process | Free flow not supported |
| 400 | 1390015 | approval is not active | Review definition is disabled, check if review definition is enabled |
| 400 | 1395001 | There have been some errors. Please try again later | There was an error with the service, please try again later |
