---
title: "Section Feature Overview"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/task-v2/section/section-feature-overview"
service: "task-v2"
resource: "section"
section: "Tasks v2"
updated: "1699521717000"
---

# Section Feature Overview

Section allows you to easily categorize tasks in the "Owned Tasks" and tasklists that you manage. With sections, you can:
* Group tasks by status: "Pending," "Ongoing," and "Completed"
* Group tasks by priority: "P0 - Urgent and Important," "P1 - Important but not Urgent," and so on
* Group tasks by category: "Marketing related," "HR related," and so on

![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/475bb62ee61228804d67b8475ead3e5c_gRasQdjpN3.png?height=842&lazyload=true&maxWidth=500&width=1280)

Section can be created within a resource. Currently, the resources that support sections are:

* "Owned Tasks" (`resource_type=my_tasks`)
* Tasklists  (`resource_type=tasklist`)

Each resource can have one or more sections. For each resource that supports sections, there must be at least 1 section, which is the "default section". Tasks can be added to a section, and if a section is deleted, the tasks in that section will be automatically moved to the default section of that resource. The default section cannot be deleted but can be renamed. Each section has a global unique ID (Section GUID) for API operations.

## General Usage

* Use the Create Section API to create a section.
* Use the Get Section API to view the GUID, name, resource, and creator of a section.
* Use the Patch Section API to update the section's name and position within the current resource.
* Use the Delete Section API to delete a section. If the section contains tasks, those tasks will be moved to the default section. The default section cannot be deleted.
* Use the List Sections API to get all sections of a resource.
* Use the List Tasks of Section API to get the tasks in a specific section. This API supports simple filtering.

To add a task to a section within a tasklist, use the Add Task to Tasklist API. To move a task from a tasklist to a section, keep the `tasklist_guid` unchanged and specify the new `section_guid`. Moving a task out of a section is equivalent to adding it to another section within the same resource. If you do not want a task to be in any sections, add it to the default section.

## What is the Relationship between Sections and Groups?
"Group" refers to the aggregation of tasks according to a certain field of the task. A section is a task classification maintained by the user manually. In simple terms, section is one of ways how tasks can be grouped. Users can choose to use "sections" to group tasks in display.

![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/961284dd2be76bfa94c72b1fb7588cde_e5jPM0u7T9.png?height=590&lazyload=true&maxWidth=500&width=1092)

## How are Sections Authorized?

Authorization of section is implemented by the resource it belongs to:

* In the tasklist, the permissions of sections follow the tasklist. If the calling identity has the read permission of the tasklist, it can view the details of all sections in the tasklist and the tasks in the sections of that tasklist. If it has the edit permission of the tasklist, it can create/delete sections and add tasks to or remove tasks from the sections.
* In "Owned Tasks", the current user can only access the sections he/she are assigneed.
