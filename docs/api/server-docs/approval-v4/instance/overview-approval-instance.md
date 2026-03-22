---
title: "Resource introduction"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/approval-v4/instance/overview-approval-instance"
service: "approval-v4"
resource: "instance"
section: "Approval"
updated: "1675167398000"
---

# Resource introduction

## Resource definition
The object generated when an employee initiates an approval. It includes more than one approval task.

## Field description
| Parameter | Type | Description |
| --- | --- | --- |
| `instance_code` | `string` | instance Code |
| `approval_name` | `string` | Approval title |
| `approval_code` | `string` | Review Definition Code |
| `start_time` | `string` | Approval creation time |
| `end_time` | `string` | Approval completion time. A value of 0 indicates that the approval is not completed. |
| `user_id` | `string` | Approval initiator |
| `open_id` | `string` | Open ID of the approval initiator |
| `serial_number` | `string` | Approval ticket number |
| `department_id` | `string` | Department of the approval initiator |
| `status` | `string` | Review instance status **Optional values are**:  - `PENDING`: Under review - `APPROVED`: Pass - `REJECTED`: Reject - `CANCELED`: Canceled - `DELETED`: Delete  |
| `uuid` | `string` | Unique ID of the user |
| `form` | `string` | JSON string, widget value |
| `task_list` | `instance_task[]` | Approval task list |
|   `id` | `string` | Task ID |
|   `user_id` | `string` | ApproverTask user_id is empty for auto-approve and auto-reject tasks. |
|   `open_id` | `string` | Open ID of the approver |
|   `status` | `string` | Task status **Optional values are**:  - `PENDING`: Under review - `APPROVED`: Agree - `REJECTED`: Reject - `TRANSFERRED`: Transferred - `DONE`: Done  |
|   `node_id` | `string` | Node id to which task belongs |
|   `node_name` | `string` | Node name to which task belongs |
|   `custom_node_id` | `string` | The custom id of the node to which the task belongs. If the custom id is not set, this field will not be returned |
|   `type` | `string` | Review method **Optional values are**:  - `AND`: Will sign - `OR`: Individual approval - `AUTO_PASS`: Pass automatically - `AUTO_REJECT`: Automatic rejection - `SEQUENTIAL`: In order  |
|   `start_time` | `string` | Task start time |
|   `end_time` | `string` | Task completion time, unfinished is 0 |
| `comment_list` | `instance_comment[]` | List of comments |
|   `id` | `string` | Comment id |
|   `user_id` | `string` | Leave a comment user |
|   `open_id` | `string` | Leave a comment User open id |
|   `comment` | `string` | Comment content |
|   `create_time` | `string` | 1564590532967 |
| `timeline` | `instance_timeline[]` | Review dynamic |
|   `type` | `string` | Dynamic typing, user_id_list different meanings in different types ext **Optional values are**:  - `START`: Review begins - `PASS`: Pass - `REJECT`: Reject - `AUTO_PASS`: Pass automatically - `AUTO_REJECT`: Automatic rejection - `REMOVE_REPEAT`: Deduplicate - `TRANSFER`: Transfer - `ADD_APPROVER_BEFORE`: Pre-approver - `ADD_APPROVER`: And sign - `ADD_APPROVER_AFTER`: Post-approver - `DELETE_APPROVER`: Subtract - `ROLLBACK_SELECTED`: Specify fallback - `ROLLBACK`: Fall back all - `CANCEL`: Cancel - `DELETE`: Delete - `CC`: CC  |
|   `create_time` | `string` | Occurrence time |
|   `user_id` | `string` | Generate users dynamically |
|   `open_id` | `string` | Dynamically generate user open id |
|   `user_id_list` | `string[]` | List of persons to be copied |
|   `open_id_list` | `string[]` | List of persons to be copied |
|   `task_id` | `string` | task_id for generating dynamic associations |
|   `comment` | `string` | Reason |
|   `cc_user_list` | `instance_cc_user[]` | CC list |
|       `user_id` | `string` | Cc user id |
|       `cc_id` | `string` | Review the unique identifier of the CC instance |
|       `open_id` | `string` | Cc open id |
|   `ext` | `string` | Other dynamic information in JSON format, currently including user_id_list, user_id, open_id_list, open_id |
|   `node_key` | `string` | Generate the node key of the task |
| `modified_instance_code` | `string` | Modified original instance code, this field is only displayed when querying modified instances |
| `reverted_instance_code` | `string` | Revoked original instance code, this field is only displayed when querying revoked instances |
| `reverted` | `boolean` | Whether the document has been revoked | ## Data example
```json
{
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
                "node_name": "start",
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
                "create_time": "1564590532967"
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
                "node_key": "APPROVAL_240330_4058663"
            }
        ],
        "modified_instance_code": "81D31358-93AF-92D6-7425-01A5D67C4E71",
        "reverted_instance_code": "81D31358-93AF-92D6-7425-01A5D67C4E71",
        "approval_code": "7C468A54-8745-2245-9675-08B7C63E7A85",
        "reverted": false
}
```

## Form description
Understand the definition of form， refer to Overview of Approval Definitions

## Approval updates
The approval updates record all actions and status changes for an approval instance from its creation to the present. It is used to track changes to instance details.

## User ID description
For how to distinguish between and use user_id, open_id, and union_id, refer to User-related IDs.

## Department ID description
For how to distinguish between and use department_id,open_department_id,refer to        Department Overview
