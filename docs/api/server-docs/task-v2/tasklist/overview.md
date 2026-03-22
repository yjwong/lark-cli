---
title: "Tasklist Features Overview"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/task-v2/tasklist/overview"
service: "task-v2"
resource: "tasklist"
section: "Tasks v2"
updated: "1699521618000"
---

# Tasklist Features Overview

## General Usage

Use the Create Tasklist API to create a tasklist. The creator is automatically set as the owner of the tasklist. At the time of creation, initial members can also be added to the tasklist.

Use the Get Tasklist Details API to get all the data for a tasklist.

Use the Delete Tasklist API to delete a tasklist. No operations can be performed on a deleted tasklist and it cannot be restored.

Use the List Tasklists API to get all the tasklists that the calling identity can list, returned in a paginated format.

Use the Update Tasklist API to modify the name and owner of a tasklist. For changing the owner, only the current owner can change the owner of the tasklist. Apart from the owner, the Update Tasklist API cannot add or remove ordinary members from the tasklist.

To modify the tasklist members, use the [Add Tasklist Members](/ssl:ttdoc:/uAjLw4CM/ukTMukTMukTM/task-v2/tasklist/add_members) and [Remove Tasklist Members](/ssl:ttdoc:/uAjLw4CM/ukTMukTMukTM/task-v2/tasklist/remove_members) APIs.

To list all the tasks in a tasklist, use the [List Tasklist Tasks](/ssl:ttdoc:/uAjLw4CM/ukTMukTMukTM/task-v2/tasklist/tasks) API.

## How are Tasklists Authorized?

If you receive error code `1470403` (authorized to access) when accessing/editing tasklist data using an API, you can read this section to check if the access_token used by the API corresponds to an identity with read/edit permissions for the tasklist data.

Tasklists have three types of permissions:

* **Read**: Can view the tasklist and all tasks within the tasklist;
* **Edit**: Can modify the tasklist name, add/remove tasklist members, add/remove tasks in the tasklist, or adjust their order;
* **Manage**: Can modify the tasklist owner and delete the tasklist.

Tasklist authorization is based on the role of caller identity. Each tasklist has only one owner and multiple members. Each members can be an editor or a viewer. Owners and members has such permissions.

* **Owner**: Read + Edit + Manage
* **Editor**: Read + Edit
* **Viewer**: Read

Members can also be a chat. When a chat is used, all chatters in that chat have the role assigned to the chat.

The creator of the tasklist is **not** a role, unlike the case the task creator automatically gains edit permissions. When a tasklist is created, the user who created the tasklist is automatically set as the owner of the tasklist. However, if the creator transfer his owner role to another user, and cannot get tasklist permission through chats, he will not have permission to access the tasklist.

In summary, a user can **read** the tasklist if meet the following conditions:

* The user is the owner of the tasklist;
* The user is a tasklist editor or viewer member;
* The tasklist has set a chat as a editor or viewer member , and the user is in that chat.

A user can **edit** the tasklist (e.g. modify the tasklist name, add tasks to the tasklist, or adjust the order of tasks in the tasklist) if meet the following conditions:

* The user is the owner of the tasklist;
* The user is an editor member of the tasklist;
* The tasklist has set a chat as an editor, and the user is in that chat.

A user can **manage** the tasklist (e.g. modify the tasklist owner, delete the tasklist) if meet the following condition:

* The user is the owner of the tasklist.
