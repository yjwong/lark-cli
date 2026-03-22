---
title: "Overview"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/approval-v4/approval-overview"
service: "approval-v4"
resource: "approval-overview"
section: "Approval"
scopes:
  - "approval:approval"
updated: "1706621674000"
---

# Overview
## Services
Lark Approval By providing a one-stop, efficient approval solution, it helps enterprises solve various approval problems and easily unlocks an efficient and pleasant approver experience. Lark Approval can quickly establish internal approval processes for enterprises, such as leave, business trips, etc. The approval open interface can query and create approval instances, which can be used for the original platform of enterprises to get through with approval. We provide a series of safe and reliable APIs to facilitate your operation of approval information. Through the approval API, you can realize various functions, such as:
- Lark native approval integration
- Third-party approval integration
- Lark Store app integration

### How to integrate
| Number | Steps | Introduction |
| --- | --- | --- |
| `1` | `Create an app` | - For creating a custom app, see Custom app development process. - For creating a store app, seeevelop and publish store apps. |
| `2` | `evelop and publish store apps` | Before calling APIs, you need to obtain the access token and enable corresponding permission scopes. For details, refer to How to call server-side APIs. |
| `3` | `Listen events and obtain changes in Contacts.` | You need to apply for relevant permission scope before listening on events. For details, see，Event subscription overview. | ### Development tutorial
| Tutorial name | Outline |
| --- | --- |
| Quickly develop three-party approvals ![](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/bfa904a141993b5b8cdb5f2242c2eaa8_MG6K8mOOmj.jpeg?height=600&lazyload=true&width=1128) | 1. Introduction 2. Prep work 3. Creating applications and requesting permissions 4. Get access token 5. Create and Update three-party Approval Definitions 6. Three-party approval instance synchronization 7. Send and update approval bot messages 8. Three-party Expedited Approval 9. Three-Party Approval Example Verification | ## Resources
The following table describes different resources.：
| Resource | Resource definition |
| --- | --- |
| Approval | A single approval flow is composed of forms and approval processes. After you create an approval, employees can fill in widget values when initiating an approval to form an instance。 |
| Instance | The object generated when an employee initiates an approval. It includes more than one task |
| Task | An approval task can't exist alone without an approval node. Each approval node may include one or more approval tasks. The approvers of the current approval node are specified in each task. If the current node includes no more than one approver, there are multiple approval tasks, each of which corresponds to a specific approver. |
| Comment | Comments or comment responses made by employees in approval instances. |
| File | When there is a picture or attachment control in the review form, the developer needs to upload the file to the review system through the review upload file interface before creating the review instance. |
| External approval | Approval definition is the description of an approval, including the approval title, icon, description, group, and other basic information.The three parties will create and synchronize three-party approval instances |
| External instance | The approval flow generated when an employee initiates an approval. Includes information such as multiple approval tasks, approval cc, etc |
| External task | Each approval operation of the approver corresponds to an approval task | Below are details for fields, methods and events of each resource.

### Resource：Approval
View  Resource fields and examples.

#### Methods
| `create approval definition` `POST` /open-apis/approval/v4/approval | `approval:approval` | `tenant_access_token` |
| --- | --- | --- |
| `Obtain Single Approval Form` `GET` /open-apis/approval/v4/approvals/:approval_code | `approval:approval:readonly` | `tenant_access_token` |
| `subscribe approval event` `POST` /open-apis/approval/v4/:approval_code/subscribe | `approval:approval` | `tenant_access_token` |
| `Cancel a Subscription to an Approvals Event` `POST` /open-apis/approval/v4/:approval_code/unsubscribe | `approval:approval` | `tenant_access_token` | ### Resource：Instance
View Resource fields and examples

#### Methods
| `	Create Approval Instance` `POST` /open-apis/approval/v4/instances | `approval:approval` | `tenant_access_token` |
| --- | --- | --- |
| `Batch Obtain Approval Instance ID` `GET` /open-apis/approval/v4/instances | `approval:approval:readonly` | `tenant_access_token` |
| `Obtain Single Approval Instance Details` `GET` /open-apis/approval/v4/instances/:instance_id | `approval:approval:readonly` | `tenant_access_token` |
| `CC Instance` `POST` /open-apis/approval/v4/instances/cc | `approval:approval:readonly` | `tenant_access_token` |
| `Approval Instance Cancel` `POST` /open-apis/approval/v4/instances/cancel | `approval:approval:readonly` | `tenant_access_token` |
| `Approval Preview` `POST` /open-apis/approval/v4/instances/preview | `approval:approval:readonly` | `tenant_access_token` | #### Event list
| `Approval Update` | `approval:approval:readonly` | `When the "approval" definition is updated` |
| --- | --- | --- |
| `Instance Update` | `approval:approval:readonly` | `When the “approval” instance status change` |
| `Task Update` | `approval:approval:readonly` | `When the “approval” task status change` |
| `Approval for leave` | `approval:approval:readonly` | `If the form of the "Approval" app contains the leave widgets, this event will be triggered after the form is approved` |
| `Approval for overtime` | `approval:approval:readonly` | `If the form of the "Approval" app contains the overtime widgets, this event will be triggered after the form is approved` |
| `Approval for shift swap` | `approval:approval:readonly` | `This event is triggered after the shift swap application is approved` |
| `Correction Approval` | `approval:approval:readonly` | `This event is triggered after the correction application is approved. You can submit an correction application in the "Attendance" app` |
| `Approval for business trip` | `approval:approval:readonly` | `If the form of the "Approval" app contains the business trip widgets, this event will be triggered after the form is approved` |
| `Offsite approval ` | `approval:approval:readonly` | `If the form of the "Approval" app contains the offsite widgets, this event will be triggered after the form is approved` | ### Resource：Task
View Resource fields and examples

#### Methods
| `Approval Task Approve` `POST` /open-apis/approval/v4/tasks/approve | `approval:approval:readonly` | `tenant_access_token` |
| --- | --- | --- |
| `Approval Task Reject` `POST` /open-apis/approval/v4/tasks/reject | `approval:approval:readonly` | `tenant_access_token` |
| `	Approval Task Transfer` `POST` /open-apis/approval/v4/tasks/transfer | `approval:approval:readonly` | `tenant_access_token` |
| `Approval Task Return` `POST` /open-apis/approval/v4/tasks/specified_rollback | `approval:approval:readonly` | `tenant_access_token` |
| `	Approval Task AddSign` `POST` /open-apis/approval/v4/tasks/add_sign | `approval:approval:readonly` | `tenant_access_token` | ### Resource：Comment
View Resource fields and examples
#### Methods
| `Create Comment` `POST` /open-apis/approval/v4/instances/:instance_id/comments | `approval:approval` | `tenant_access_token` |
| --- | --- | --- |
| `Obtain comment` `GET` /open-apis/approval/v4/instances/:instance_id/comments | `approval:approval` `approval:approval:readonly` | `tenant_access_token` |
| `Delete Comment` `DELETE` /open-apis/approval/v4/instances/:instance_id/comments/:comment_id | `approval:approval` | `tenant_access_token` |
| `Clear Comment` `POST` /open-apis/approval/v4/instances/:instance_id/comments/remove | `approval:approval` | `tenant_access_token` | ### Resource：External Approval
View Resource fields and examples

#### Methods
| `External Approval Create` `POST` /open-apis/approval/v4/external_approvals | `approval:approval` | `tenant_access_token` |
| --- | --- | --- | ### Resource：External Instance
View Resource fields and examples

#### Methods
| `External Approval Instance Create` `POST` /open-apis/approval/v4/external_instances | `approval:approval` | `tenant_access_token` |
| --- | --- | --- |
| `External Approval Instance Check` `POST` /open-apis/approval/v4/external_instances/check | `approval:approval:readonly` | `tenant_access_token` | ### Resource：Exteranl Task
View Resource fields and examples

#### Methods
| `Get Third-party Approval Task Status` `GET` /open-apis/approval/v4/external_tasks | `approval:approval:readonly` | `tenant_access_token` |
| --- | --- | --- |
