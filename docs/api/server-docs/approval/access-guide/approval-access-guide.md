---
title: "Approval access guide"
url: "https://open.larksuite.com/document/ukTMukTMukTM/ukDNyUjL5QjM14SO0ITN"
section: "Approval"
updated: "1675167379000"
---

# Approval access guide

## What is the approval API

Approval is an official app provided to companies as part of the Lark suite. You can use it to quickly create internal approvals processes, such as for leave requests and business trips.  
The approval open API is used to query and create approval instances. It is used to connect companies' existing platforms to Lark Approval.

## Terms

### 1. Approval

A single approval flow is composed of forms and approval processes (see below for details). After you create an approval, employees can fill in widget values when initiating an approval to form an approval instance.

### 1) Approval Code

The unique code for an approval definition, which can be found in the browser URL when you edit an approval. Click the following link to enter the Admin console:

[https://www.larksuite.com/approval/admin/approvalList?devMode=on](https://www.larksuite.com/approval/admin/approvalList?devMode=on)

![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/d02eb479bb69d97f88a7e118bb79e0d5_vJ7QNmLumi.png?lazyload=true&width=1640&height=207)

Click "Edit" in the image above to find the Approval Code in the URL, with the format: `definitionCode=E3254848-D172-4169-B03E-744E7CD11F06`.

### 2) Form

Forms contain one or more widgets. When an employee initiates an approval, they must fill out the form page. Each approval has its own form.

#### Widget

A constituent part of an approval. A form is comprised of one or more approval widgets. Each widget has a basic field:

```js 
struct{
    String Id; //A unique ID that identifies a widget in an approval.
    String CustomId; //A custom widget ID, a unique ID that identifies a widget in an approval.
    String Name; //The name of this widget
    String Type; //The type of this widget
} 
```
#### Custom ID

You can specify an ID when editing an approval definition on the Approval Admin. You can subsequently specify this ID to identify the widget in the approval definition.

#### Widget description

1. Single-line text (input)
- Used to input single-line text.

2. Multi-line text (textarea)
- Used to input multi-line text.

3. Number (number)
- Used to input a number.

4. Amount (amount)
- Used to input the approval amount and unit (default unit: CNY).

5. Date (date)
- Used to input a specific time.

6. Date interval (dateInterval)
- Used to input a date interval, including a start time, end time, and duration.

7. Single option (radio, radioV2)
- Used to select a single option.

8. Multiple options (checkbox, checkboxV2)
- Used to select multiple options.

9. Address (address)
- Used to input an address.

10. Contact (contact)
- Used to add contacts in the approval.

11. Description (text)
- Used to add a description in the approval definition (such as specifications or precautions). This can't be edited when an approval is initiated.

12. Details (fieldList)
- Used to input details information. In Details widget, you can add other widgets such as number and amount. When creating an approval definition, design a Details widget to specify what widgets are included in an entry. Initiators can add entries according to their own demands, and each entry is consistent with the initially configured Details widget.

13. Formula (formula)
- When creating an approval definition, design a formula to specify how the value of this widget is calculated from other widgets (numbers and amounts).

14. Image (image, imageV2)
- Used to add an image in the approval.

15. Attachment (attachment, attachmentV2)
- Used to add an attachment, such as a file, to the approval.

16. Associate (connect)
- Used to associate to another approval in the current approval. This allows approvers to see information for the associated approval when reviewing the current approval.

17. Leave (leaveGroup)
- Used to enter leave approval content, such as to select a leave type (sick leave, maternity leave, etc.; leave types must be set by administrators in Leave Management and selected when initiating an instance for the current approval definition) and enter the leave start time, end time, and leave duration.

18. Overtime (workGroup)
- Used to enter overtime request content, such as to select the overtime type (compensatory leave, paid overtime, unpaid overtime, etc.; types must be set by the administrator when creating the approval definition) and enter the overtime start and end times, duration, and reason for overtime.

19. Shift swap (shiftGroup)
- Used to enter content for a shift swap request, including the original shifts and reason for the swap.

20. Business trip (tripGroup)
- Used to enter business trip approval content, including the trip itinerary (including the itinerary start time, itinerary end time, itinerary duration, departure location, destination location, transportation method, one-way/round-trip, and notes; the itinerary duration is calculated in calendar days; if a trip has multiple itineraries, the initiator can manually add additional itineraries), total trip duration, trip reason, and traveling with.

21. Correction (remedyGroup)
- This is used to enter correction approval content, including the correction time and reason for correction.

### 3) Approval process (Process)

When creating an approval definition, an administrator will design an approval transfer mode, which specifies every step of an instance created by this definition. An approval definition corresponds to an approval process, which is made up of approval nodes. That is, an approval process corresponds to one or more approval nodes. When an employee initiates an approval instance, the instance enters the approval process as designed. When designing an approval process, you must specify the approval actions for each step (For approval nodes, see below for details). This includes the approvers for each step or the information for auto-approval or auto-rejection by the system. If the current step is not automatically performed by the system, approvers need to be specified or selected by the initiator.

#### Approvers selected by initiator

If an approver isn't specified for a certain step in the approval definition and "Approver selected by initiator" is set instead, the user must select an approver when initiating an approval instance. When creating such an approval definition, you only need to set if approval is required by everyone assigned or anyone assigned. **When "Approver selected by initiator" is set, the user must specify an approver when initiating an approval instance through the API**.  For more information about how the API initiates an "Approver selected by initiator" instance, see the API method description.

#### Approval node (Node)

Each approval process is composed of one or more approval nodes. Each node specifies a step in the approval flow. After the current node is completed, the approval advances to the next node. When an instance is created, the corresponding approval nodes are generated to form an approval process.

#### Custom ID

Similar to a widget, you can specify a custom node ID. This is a unique identifier used to specify this node in an approval process.

#### Approval node type

1. Everyone assigned (AND)
- This is a logical relationship that indicates all approvers selected for the current node must approve the instance for it to move to the next node.

2. Anyone assigned (OR)
- This is a logical relationship that indicates that only one approver selected for the current node must approve the instance for it to move to the next node.

3. Auto-approve (AUTO_PASS)
- This indicates that the system will automatically approve the instance in the current node to allow it to move to the next node.

4. Auto-reject (AUTO_REJECT)
- This indicates that the system will automatically reject the instance in the current node.

### 2. Approval instance (Instance)

The object generated when an employee initiates an approval. It includes more than one approval task (see below for details).

#### Approval instance status
1. In progress (PENDING)
- This indicates the current approval instance is ongoing and doesn't yet have a result.

2. Approved (APPROVED)
- This indicates the current approval instance has been approved.

3. Rejected (REJECTED)
- This indicates the current approval instance has been rejected.

4. Canceled (CANCELED)
- This indicates the current approval instance has been canceled by the initiator.

5. Deleted (DELETED)
- This indicates the current approval is ongoing, but the associated approval definition has been disabled or deleted by an administrator so that the instance is in the deleted status.

### Approval task (Task)

An approval task can't exist alone without an approval node. Each approval node may include one or more approval tasks. The approvers of the current approval node are specified in each task. If the current node includes no more than one approver, there are multiple approval tasks, each of which corresponds to a specific approver. When a node includes multiple approval tasks, some of the tasks are not solely dependent on the approver's actions to change their status, but are affected by the status change of other tasks. (For example, if an anyone assigned node includes multiple approval tasks, after one task is approved, other tasks are automatically converted into the completed status without any action from the corresponding approver.) When an approval instance moves on to a new node, the tasks of the node are created.

#### Approval task status

1. In progress (PENDING)
- This indicates the task is in progress and doesn't yet have a result.

2. Approved (APPROVED)
- This indicates the approver for the approval task has approved the instance.

3. Rejected (REJECTED)
- This indicates the approver for the approval task has rejected the instance.

4. Transferred (TRANSFERRED)
- This indicates that the approver for the approval task has transferred the instance to another approver.

5. Completed (DONE)
- This indicates the approval task has been completed.

::: note
Some of the situations that result in a status of DONE are as follows:
- In an OR node, one approver has approved the instance, in which case the task for this approver is APPROVED and the other tasks are DONE.
- In an AND or OR node, one approver has rejected the instance, in which case the task for this approver is REJECTED and the other tasks are DONE.
- When a task is PENDING and the initiator cancels the approval, the task status changes to DONE.
- When a task is PENDING and an administrator deletes the corresponding approval definition, the task status changes to DONE.
- When a task status is DONE, this generally means the approver for the task hasn't completed the task while the statuses of other tasks in the node have changed.

### 3. Approval updates (Timeline)

The approval updates record all actions and status changes for an approval instance from its creation to the present. It is used to track changes to instance details.

## How can I use the approval API

### 1. Create an approval

After the approval function is enabled, go to the Approval Admin and enter the required approval information, form, and process to create an approval.

![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/0f82470fd78078fc2bb10ee0cdfe37ff_uZyYBGKrp9.png?lazyload=true&width=1640&height=769)

![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/66b8a297dc56bfe1b74c94f0b18df6a0_942kVPhgcL.png?lazyload=true&width=1640&height=731)

### 2. Call the API

After an approval is created, employees can go to the App Center > Approval Gadget to initiate an approval instance. Developers can call the API to initiate an approval or obtain approval instance data.

The Approval app supports the following APIs:

* View approval definition
* Obtain approval instances in batches
* Obtain details of a single approval instance
* Create an approval instance
* Approve an approval task
* Reject an approval task
* Transfer an approval task
* Revoke an approval task

The tenant_access_token is use to authenticate approval API scope. You can obtain a token as follows:

* [Obtain tenant_access_token (for custom apps)](https://lark-open.bytedance.net/document/ukTMukTMukTM/ukDNz4SO0MjL5QzM/auth-v3/auth/tenant_access_token_internal)
* [Obtain tenant_access_token (for store apps)](https://lark-open.bytedance.net/document/ukTMukTMukTM/ukDNz4SO0MjL5QzM/auth-v3/auth/tenant_access_token)

### 3. Event listening

* Approval Event Listening Developer's Guide
