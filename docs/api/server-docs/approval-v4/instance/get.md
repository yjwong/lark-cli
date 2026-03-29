---
title: "Obtain details of an approval instance"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/approval-v4/instance/get"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/approval/v4/instances/:instance_id"
service: "approval-v4"
resource: "instance"
section: "Approval"
scopes:
  - "approval:approval"
  - "approval:approval:readonly"
  - "approval:instance"
field_scopes:
  - "contact:user.employee_id:readonly"
updated: "1675167402000"
---

# Obtain details of an approval instance

This API is used to obtain approval instance details through Instance Code. You can obtain Instance Code via the Obtain approval instances in batches API.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/approval/v4/instances/:instance_id |
| HTTP Method | GET |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `approval:approval` `approval:approval:readonly` `approval:instance` |
| Required field scopes | > The response body of the API contains the following sensitive fields, and they will be returned only after corresponding scopes are added. If you do not need the fields, it is not recommended that you request the scopes. `contact:user.employee_id:readonly` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `instance_id` | `string` | Approval instance code. If you pass in the uuid when creating the approval instance, you can also obtain the code by passing in the uuid. **Example value**: "81D31358-93AF-92D6-7425-01A5D67C4E71" | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `locale` | `string` | No | Language **Example value**: "zh-CN" **Optional values are**:  - `zh-CN`: Chinese - `en-US`: English - `ja-JP`: Japanese  |
| `user_id` | `string` | No | Initiate review user id，Only self-built applications can return **Example value**: "f7cb567e" |
| `user_id_type` | `string` | No | User ID categories **Example value**: "open_id" **Optional values are**:  - `open_id`: Identifies a user to an app. The same user has different Open IDs in different apps. How to get Open ID - `union_id`: Identifies a user to a tenant that acts as a developer. A user has the same Union ID in apps developed by the same developer, and has different Union IDs in apps developed by different developers. A developer can use Union ID to link the same user's identities in multiple apps.How to get Union ID - `user_id`: Identifies a user to a tenant. The same user has different User IDs in different tenants. In one single tenant, a user has the same User ID in all apps （including store apps）. User ID is usually used to communicate user data between different apps. How to get User ID  **Default value**: `open_id` **When the value is `user_id`, the following field scopes are required**: `contact:user.employee_id:readonly` | ## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `approval_name` | `string` | Approval title |
|   `start_time` | `string` | Approval creation time |
|   `end_time` | `string` | Approval completion time. A value of 0 indicates that the approval is not completed. |
|   `user_id` | `string` | Approval initiator |
|   `open_id` | `string` | Open ID of the approval initiator |
|   `serial_number` | `string` | Approval ticket number |
|   `department_id` | `string` | Department of the approval initiator |
|   `status` | `string` | Review instance status **Optional values are**:  - `PENDING`: Under review - `APPROVED`: Pass - `REJECTED`: Reject - `CANCELED`: Withdraw - `DELETED`: Delete  |
|   `uuid` | `string` | Unique ID of the user |
|   `form` | `string` | JSON string, widget value |
|   `task_list` | `instance_task[]` | Approval task list |
|     `id` | `string` | Task ID |
|     `user_id` | `string` | Approver Task user_id is empty for auto-approve and auto-reject tasks. |
|     `open_id` | `string` | Open ID of the approver |
|     `status` | `string` | Task status **Optional values are**:  - `PENDING`: Under review - `APPROVED`: Agree - `REJECTED`: Reject - `TRANSFERRED`: Transferred - `DONE`: Done  |
|     `node_id` | `string` | Node id to which task belongs |
|     `node_name` | `string` | Node name to which task belongs |
|     `custom_node_id` | `string` | The custom id of the node to which the task belongs. If the custom id is not set, this field will not be returned |
|     `type` | `string` | Review method **Optional values are**:  - `AND`: Will sign - `OR`: Individual approval - `AUTO_PASS`: Pass automatically - `AUTO_REJECT`: Automatic rejection - `SEQUENTIAL`: In order  |
|     `start_time` | `string` | Task start time |
|     `end_time` | `string` | Task completion time, unfinished is 0 |
|   `comment_list` | `instance_comment[]` | List of comments |
|     `id` | `string` | Comment id |
|     `user_id` | `string` | Leave a comment user |
|     `open_id` | `string` | Leave a comment User open id |
|     `comment` | `string` | Comment content |
|     `create_time` | `string` | 1564590532967 |
|     `files` | `file[]` | Comment attachments |
|       `url` | `string` | Resource path |
|       `file_size` | `int` | Attachment size |
|       `title` | `string` | Attachment title |
|       `type` | `string` | Attachment type |
|   `timeline` | `instance_timeline[]` | Review dynamic |
|     `type` | `string` | Dynamic typing, user_id_list different meanings in different types ext **Optional values are**:  - `START`: Review begins - `PASS`: Pass - `REJECT`: Reject - `AUTO_PASS`: Pass automatically - `AUTO_REJECT`: Automatic rejection - `REMOVE_REPEAT`: Deduplicate - `TRANSFER`: Transfer - `ADD_APPROVER_BEFORE`: Pre-approver - `ADD_APPROVER`: And sign - `ADD_APPROVER_AFTER`: Post-approver - `DELETE_APPROVER`: Subtract - `ROLLBACK_SELECTED`: Specify fallback - `ROLLBACK`: Fall back all - `CANCEL`: Withdraw - `DELETE`: Delete - `CC`: CC  |
|     `create_time` | `string` | Occurrence time |
|     `user_id` | `string` | Generate users dynamically |
|     `open_id` | `string` | Dynamically generate user open id |
|     `user_id_list` | `string[]` | List of persons to be copied |
|     `open_id_list` | `string[]` | List of persons to be copied |
|     `task_id` | `string` | task_id for generating dynamic associations |
|     `comment` | `string` | Reason |
|     `cc_user_list` | `instance_cc_user[]` | CC list |
|       `user_id` | `string` | Cc user id |
|       `cc_id` | `string` | Review the unique identifier of the CC instance |
|       `open_id` | `string` | Cc open id |
|     `ext` | `string` | Other dynamic information in JSON format, currently including user_id_list, user_id, open_id_list, open_id |
|     `node_key` | `string` | Generate the node key of the task |
|     `files` | `file[]` | Approval attachments |
|       `url` | `string` | Resource path |
|       `file_size` | `int` | Attachment size |
|       `title` | `string` | Attachment title |
|       `type` | `string` | Attachment type |
|   `modified_instance_code` | `string` | Modified original instance code, this field is only displayed when querying modified instances |
|   `reverted_instance_code` | `string` | Revoked original instance code, this field is only displayed when querying revoked instances |
|   `approval_code` | `string` | Review Definition Code |
|   `reverted` | `boolean` | Whether the document has been revoked |
|   `instance_code` | `string` | Approval instance code | **Widget value**: 
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
| attachmentV2 | Attachment widget. ext is the name of the attachment, value is the address of the attachment, separated by commas. | The link remains valid for 12 hours.
| image/imageV2 | Image widget. ext is the name of the image, value is the address of the image, separated by commas. | The link remains valid for 12 hours.

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
{
    "code": 0,
    "msg": "success",
    "data": {
        "approval_name": "Payment",
        "start_time": "1564590532967",
        "end_time": "1564590532967",
        "user_id": "f3ta757q",
        "open_id": "ou_3cda9c969f737aaa05e6915dce306cb9",
        "serial_number": "202102060002",
        "department_id": "od-8ec33ffec336c3a39a278bc25e931676",
        "status": "PENDING",
        "uuid": "1234567",
        "form": "[{\"id\": \"widget1\",\"custom_id\": \"user_info\",\"name\": \"Item application\",\"type\": \"textarea\"},\"value\":\"aaaa\"]",
        "task_list": [
            {
                "id": "1234",
                "user_id": "f7cb567e",
                "open_id": "ou_123457",
                "status": "PENDING",
                "node_id": "46e6d96cfa756980907209209ec03b64",
                "node_name": "开始",
                "custom_node_id": "manager",
                "type": "AND",
                "start_time": "1564590532967",
                "end_time": "0"
            }
        ],
        "comment_list": [
            {
                "id": "1234",
                "user_id": "f7cb567e",
                "open_id": "ou_123456",
                "comment": "ok",
                "create_time": "Comment time",
                "files": [
                    {
                        "url": "https://p3-approval-sign.byteimg.com/lark-approval-attachment/image/20220714/1/332f3596-0845-4746-a4bc-818d54ad435b.png~tplv-ottatrvjsm-image.image?x-expires=1659033558&x-signature=6edF3k%2BaHeAuvfcBRGOkbckoUl4%3D#.png",
                        "file_size": 186823,
                        "title": "e018906140ed9388234bd03b0.png",
                        "type": "image"
                    }
                ]
            }
        ],
        "timeline": [
            {
                "type": "PASS",
                "create_time": "1564590532967",
                "user_id": "f7cb567e",
                "open_id": "ou_123456",
                "user_id_list": [
                    "Eeea5gefe"
                ],
                "open_id_list": [
                    "ou_123456"
                ],
                "task_id": "1234",
                "comment": "ok",
                "cc_user_list": [
                    {
                        "user_id": "eea5gefe",
                        "cc_id": "123445",
                        "open_id": "ou_12345"
                    }
                ],
                "ext": "{\"user_id\":\"62d4a44c\",\"open_id\":\"ou_123456\"}",
                "node_key": "APPROVAL_240330_4058663",
                "files": [
                    {
                        "url": "https://p3-approval-sign.byteimg.com/lark-approval-attachment/image/20220714/1/332f3596-0845-4746-a4bc-818d54ad435b.png~tplv-ottatrvjsm-image.image?x-expires=1659033558&x-signature=6edF3k%2BaHeAuvfcBRGOkbckoUl4%3D#.png",
                        "file_size": 186823,
                        "title": "e018906140ed9388234bd03b0.png",
                        "type": "image"
                    }
                ]
            }
        ],
        "modified_instance_code": "81D31358-93AF-92D6-7425-01A5D67C4E71",
        "reverted_instance_code": "81D31358-93AF-92D6-7425-01A5D67C4E71",
        "approval_code": "7C468A54-8745-2245-9675-08B7C63E7A85",
        "reverted": false,
        "instance_code": "81D31358-93AF-92D6-7425-01A5D67C4E71"
    }
}
```

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 1390001 | param is invalid | Parameter error |
| 400 | 1395001 | There have been some errors. Please try again later | There was an error with the service, please try again later |
| 400 | 1390003 | instance code not found | Check that the review instance code is correct |
| 403 | 1390009 | no operation permission | Check that the operation permissions are correct |
| 400 | 1390004 | user_id or open_id not found | Check user_id open_id correct |
