---
title: "View approval definitions"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/approval-v4/approval/get"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/approval/v4/approvals/:approval_code"
service: "approval-v4"
resource: "approval"
section: "Approval"
scopes:
  - "approval:approval"
  - "approval:approval:readonly"
  - "approval:definition"
field_scopes:
  - "contact:user.employee_id:readonly"
updated: "1675167395000"
---

# View approval definitions

This API is used to get the details of an approval definition based on an Approval Code and create approval instance requests.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/approval/v4/approvals/:approval_code |
| HTTP Method | GET |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `approval:approval` `approval:approval:readonly` `approval:definition` |
| Required field scopes | > The response body of the API contains the following sensitive fields, and they will be returned only after corresponding scopes are added. If you do not need the fields, it is not recommended that you request the scopes. `contact:user.employee_id:readonly` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `approval_code` | `string` | Approval definition code **Example value**: "7C468A54-8745-2245-9675-08B7C63E7A85" | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `locale` | `string` | No | Language optional **Example value**: "zh-CN" **Optional values are**:  - `zh-CN`: Chinese - `en-US`: English - `ja-JP`: Japanese  |
| `with_admin_id` | `boolean` | No | Optional Whether to return a list of admin IDs with data permissions for the approval process **Example value**: false |
| `user_id_type` | `string` | No | User ID categories **Example value**: "open_id" **Optional values are**:  - `open_id`: Identifies a user to an app. The same user has different Open IDs in different apps. How to get Open ID - `union_id`: Identifies a user to a tenant that acts as a developer. A user has the same Union ID in apps developed by the same developer, and has different Union IDs in apps developed by different developers. A developer can use Union ID to link the same user's identities in multiple apps.How to get Union ID - `user_id`: Identifies a user to a tenant. The same user has different User IDs in different tenants. In one single tenant, a user has the same User ID in all apps （including store apps）. User ID is usually used to communicate user data between different apps. How to get User ID  **Default value**: `open_id` **When the value is `user_id`, the following field scopes are required**: `contact:user.employee_id:readonly` | ## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `approval_name` | `string` | Approval title |
|   `status` | `string` | Approval definition status **Optional values are**:  - `ACTIVE`: Enabled - `INACTIVE`: Disabled - `DELETED`: Deleted - `UNKNOWN`: Unknown  |
|   `form` | `string` | JSON array, widget information |
|   `node_list` | `approval_node_info[]` | Node information |
|     `name` | `string` | Node name |
|     `need_approver` | `boolean` | Whether it is a node selected by the initiator true - The initiator must specify approvers |
|     `node_id` | `string` | Node ID |
|     `custom_node_id` | `string` | Custom node ID. not returned if no value is set |
|     `node_type` | `string` | Approval method **Optional values are**:  - `AND`: Everyone assigned - `OR`: Anyone assigned - `SEQUENTIAL`: Assigned one by one - `CC_NODE`: Cc node  |
|     `approver_chosen_multi` | `boolean` | Whether to support multiple selection: true-support, the value of initiating and ending nodes is meaningless |
|     `approver_chosen_range` | `approver_chosen_range[]` | Optional range |
|       `approver_range_type` | `int` | Specified range: 0-all, 1-specified role, 2-specified member **Optional values are**:  - `0`: Company-wide - `1`: Specify role scope - `2`: Specify user range  |
|       `approver_range_ids` | `string[]` | According to the above type, store the role id and userid respectively. When type is 0, this field is an empty list |
|     `require_signature` | `boolean` | Whether to sign |
|   `viewers` | `approval_viewer_info[]` | Viewer list |
|     `type` | `string` | Visible person type **Optional values are**:  - `TENANT`: visible within tenant - `DEPARTMENT`: visible to the specified department - `USER`: visible to the specified user - `ROLE`: visible to the specified role - `USER_GROUP`: visible to the specified user group - `NONE`: visible to none  |
|     `id` | `string` | When the visible person type is DEPARTMENT, id is the id of the department; when the visible person type is USER, id is the id of the user; when the visible person type is ROLE, id is the id of the role; when the visible person type is USER_GROUP, id is the id of the user group |
|     `user_id` | `string` | User ID of the viewer when the viewer type is USER |
|   `approval_admin_ids` | `string[]` | Approval process administrator ID with data management authority | **Form field description**
|Parameter|Type|Required|Description|
|-|-|-|-|
form|String|Yes| JSON array, **widget information**| ∟id|String|Yes|Widget ID| ∟widget_default_value|String|Yes|Default Value of Widget| ∟custom_id|String|No|Custom widget ID| ∟default_value_type|String|Yes|Default Value Type of Widget| ∟name|String|Yes|Widget name| ∟type|String|Yes|Widget type| ∟enable_default_value|bool|Yes|Default Value Enabled or Not| ∟widget_default_value|String|Yes|Default Value of Widget| ∟default_value_type|String|Yes|Default Value Type of Widget| ∟display_condition|String|No|Control explicit and invisible conditions| ∟conditional|String|No|| ∟conditions|list|No|| ∟conditional|String|No|Multiple conditions are met at the same time| ∟expressions|list|No|Judgment rules| ∟source_widget|string|No|Condition value| **Widget information**: 
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
                "custom_node_id": "46e6d96cfa756980907209209ec03b64",
                "node_type": "AND",
                "approver_chosen_multi": true,
                "approver_chosen_range": [
                    {
                        "approver_range_type": 2,
                        "approver_range_ids": [
                            "ou_e03053f0541cecc3269d7a9dc34a0b21"
                        ]
                    }
                ],
                "require_signature": false
            }
        ],
        "viewers": [
            {
                "type": "TENANT",
                "id": "ou_e03053f0541cecc3269d7a9dc34a0b21",
                "user_id": "f7cb567e"
            }
        ],
        "approval_admin_ids": [
            "ou_3cda9c969f737aaa05e6915dce306cb9"
        ]
    }
}
```

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 1390001 | param is invalid | Parameter error |
| 400 | 1390002 | approval code not found | Check that the review definition code is correct |
| 400 | 1395001 | There have been some errors. Please try again later | There was an error with the service, please try again later |
| 400 | 1390016 | approval is deleted | Check if the review definition exists |
