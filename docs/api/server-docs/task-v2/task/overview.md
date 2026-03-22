---
title: "Task Feature Overview"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/task-v2/task/overview"
service: "task-v2"
resource: "task"
section: "Tasks v2"
updated: "1699521551000"
---

# Task Features Overview

## Regular Usage

Use Create Task API to create a task. When creating a task, you can add initial task members, start/due time, reminders, etc., and you can also add the created task to one or more tasklists at the same time.

Use the Get Task Details API to get all the data of a task.

Use the Delete Task API to delete a task. No operation can be performed on a deleted task, and a deleted task cannot be recovered.

Use the List Tasks API to get all tasks the caller is responsible for.

Use the Update Task API to modify the properties of a task, including summary, description, due, etc. However, there are separated APIs for adding/removing task members, reminders, and tasklists.

To modify task members, use the Add Task Members and Remove Task Members APIs.

To modify task reminders, use the Add Task Reminders and Remove Task Reminders APIs.

To change the tasklist a task belongs to, use the Add Task to Tasklist and Remove Task from Tasklist APIs. You can view all the tasklists a task belongs to using the List Tasklists of Task API.

To change the task dependencies, use Add Dependency and Remove Dependency API.

## How are Tasks Authorized?
If you get a `1470403` error code (unauthorzied to access) when calling an API to get/edit task data, you can read this section to check if the caller identity corresponding to the access_token has read/edit permission for the task data.

Tasks have two types of permissions:

* **Read**: Can view various data of the task and any sub-tasks. Can initiate comments on tasks;
* **Edit**: Can modify the summary, description, due, etc. of the task, can add/remove task members, can complete the task/restore it to unfinished, can add sub-tasks, etc.

There are two main roles in tasks:

* Assignee, who has read and edit permissions for tasks;
* Follower, who has read permissions for tasks.

In addition, the creator of the task automatically gets read and edit permissions for the task.

In addition to task members, tasks can also be authorized by sharing to chats, adding to tasklists, and other ways. In general, when a user meets the following conditions, they can **read** the task data:

* The user is the creator, assignee, or follower of the task;
* The user received the task's sharing message in a chat;
* The task is a document task (tasks can be created on Lark documents and synced to Lark tasks), and the user is a collaborator of the document;
* The task belongs to a tasklist, and the user is a member of the tasklist, or
* The task is a subtask, and the user has read permission for any task on its parent tasks chain.

When a user meets the following conditions, they can **edit** the task data (e.g., modify the task summary, add/remove task assignees, etc.):

* The user is the creator or assignee of the task;
* The task is a document task (tasks can be created on Lark documents and synced to Lark tasks), and the user is an editable collaborator of the document;
* The task belongs to a tasklist, and the user is an editor or owner of the tasklist;
* The task is a subtask, and the user has editable permission for any task on the parent task chain.

## How to Complete/Uncomplete a Task?

Use the update task API and set the `completed_at` field to a specific timestamp (ms) to complete the task.
```json
PATCH /task/v2/tasks/:task_guid
{
  "task": {
     "completed_at": "1688529478000"
  },
  "update_fields":["completed_at"]
}
```

Setting `completed_at` to 0 means uncompleting the task.

```json
PATCH /task/v2/tasks/:task_guid
{
  "task": {
     "completed_at": "0"
  },
  "update_fields":["completed_at"]
}
```

If a task has been completed, it cannot be completed again. That is, the task's completion time can only be the first time it changes from uncompleted to completed. If you try to set a new completion time for a completed task, you will get an error code=1470400, msg="cannot set non-zero completed_at for a completed task".

To change the completion time of a completed task, you need to set `complete_at` to "0" to restore the task to the uncompleted state, and then complete the task again with the new `complete_at` value.

## How to Use Start and Due Time?

Each task can set a start time and due time for task planning.

The format of start and due time is the same, both including a `timestamp` and a `is_all_day` field:

```json
{
  "timestamp": "1684652400000",
  "is_all_day": false
}
```

`is_all_day` indicates whether the timestamp represents a date. `is_all_day=true` indicates a date; `is_all_day=false` means a moment. If `is_all_day` is not filled in, it defaults to false.

When the start/due time represents a moment, its timestamp is a timestamp since 1970-01-01 00:00:00 UTC in milliseconds. When displaying the moment, the Lark App will use the user's timezone setting to convert this timestamp to the displayed specific timezone.

For example, for a task with a due time of `timestamp=1684652400000, is_all_day=false`. The task has two members, user A using timezone +08:00, and user B using timezone -07:00. User A sees "2023-05-21 15:00:00", while user B  sees "2023-05-21 00:00:00".

When the start/due time represents a date,
the timestamp is a day-precision timestamp since 1970-01-01 00:00:00 UTC in milliseconds. For example:

```json
{
  "timestamp": "1682899200000",
  "is_all_day": true
}
```

When displaying the date, no matter what the user's timezone is, they will see the same date. For example, the due date is `timestamp=1682899200000, is_all_day=true`. The task has two members, user A using timezone +08:00, while user B using timezone -07:00. Both user A and B will see the task due on "2023-05-01".

When calling the create task/update task to set the start/due time, if the numerical precision of the `timestamp` exceeds its business precision, it will be automatically truncated. For moments, only second-precision timestamps are retained; for dates, only day-precision timestamps are retained. 

If the `timestamp` of the start/due time is set to 0, it means clearing the start/due. For example, when updating a task, setting `task.due` to null and setting `task.due.timestamp` to 0 are equivalent.

```json
PATCH /task/v2/tasks/:task_guid
{
  "task": {
    "due": null
  },
  "update_fields": ["due"]
}

PATCH /task/v2/tasks/:task_guid
{
  "task": {
    "due": {
      "timestamp": "0"
    }
  },
  "update_fields": ["due"]
}
```

The `timestamp` cannot be set to a negative number, which means that the start/due time of a task cannot be set to a moment/date before 1970-01-01.

Start and due time can be set independently. A task can have no start/due time; have a start time without a due time; have a due time without a start time; or have both start and due times. However, **if both start and due times are set, the following constraints must be met**:

* The timestamp of the start time must not be later than the timestamp of the due time, i.e., `start.timestamp Only fill in the languages that need to be supported.|
|href.url|Specific link address|It can be left blank, but if it is filled in, it must be a valid url starting with http:// or https://.|
|href.title|Title of the corresponding link|Optional. If both url and title are filled in, the title will be displayed and support click-through as a link.| In an Origin, it is generally recommended to use `platform_i18n_name` to refer to the name of the application itself; `title` + `url` should be used to fill in the name and access url of the specific information related to the task.

For example, if an enterprise has its own ticket system and has developed an application called "Ticket Assistant" associated with the tickets in the ticket system and Lark tasks. When creating a task for a ticket, `platform_i18n_name` should be filled in with the name of the ticket system or the name of the Ticket Assistant application, and `title` should be filled in with the specific name of the ticket, and `url` should be filled in with the specific url of the ticket. This way, tasks can be grouped and filtered according to their sources in the Lark task interface, and clicking on the source of a task can jump to the corresponding page of the ticket system.

Request example:
```
POST https://open.larksuite.com/open-apis/task/v2/tasks
{
    "summary": "The confirmation button is not effective when clicked by the user",
    "origin": {
         "platform_i18n_name": {
             "zh_cn": "Ticket System",
             "en_us": "Ticket System"
         },
         "href": {
             "title": "The confirmation button is not effective when clicked by the user",
             "url": "https://abc.com"
         }
    }
}
```
The effect is as follows:

| Task Details | Task List |
| --- | --- |
|![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/cb6ebe8e7d0dd22284b15a167c5a4c9c_Tcl19aS8rp.png?height=514&lazyload=true&width=786)|![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/a3268cba2924aff063770bde93c9e235_IjLiDAgqP5.png?height=614&lazyload=true&width=1784)| ## How to Use Custom Completion?

"Custom completion" refers to the ability to customize the behavior when the "complete task" button is clcked in App Task Detail UI. This feature is mainly used for synchronizing completion status between Lark tasks and other third-party systems.

When the completion checkbox of a task with custom completion is clicked on the interface, it will either redirect to a third-party interface or generate a prompt to guide the user to the third-party system. After the user completes the task in the third-party system, the third-party system uses the task API to set the task completion status.

For example, a third-party ticket system creates a Lark task through an API and sets a custom completion redirect link. When the user clicks "Complete Task" on the Lark App UI, the task will not be marked as complete, but instead will guide the user to the corresponding ticket web page in the third-party ticket system. On this page, the user can fill in the necessary information according to the ticket settings and complete the ticket. After that, the ticket system can call the task API to mark the task as completed, achieving consistency between the completion status of the Lark Task and the ticket in the third-party ticket system.

To set a task as "custom completion", use the `custom_complete` field of the task. This field can be set by Create Task or Patch Task.

There are two ways to use the `custom_complete` field, in the following formats:

1. Set the redirection to a third-party page after completing the task:
```json
{
"custom_complete":{
  "android":{
    "href":"https://google.com.hk"
    },
   "ios":{
     "href":"https://google.com.hk"
    },
   "pc":{
     "href":"https://google.com.hk"
    }
}
```

Different links can be set for different terminals.

For example, setting the href to the page link `https://google.com.hk` will directly redirect to that page after clicking "Complete Task". The effect is as follows:

![bbbb.gif](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/7f24b1a828722e57d935c89912e5dcad_LEa0tsw9wZ.gif?height=1854&lazyload=true&width=3358)
  
2. Set a prompt to display after completing the task, which can support multiple languages:
```json
{
"custom_complete":{
  "android":{
    "tip":{
      "zh_cn":"请前往指定页面完成任务",
      "en_us":"please go to the designated page to complete the task"
     }
    },
   "ios":{
     "tip":{
       "zh_cn":"请前往指定页面完成任务",
       "en_us":"please go to the designated page to complete the task"
      }
    },
   "pc":{
     "tip":{
       "zh_cn":"请前往指定页面完成任务",
       "en_us":"Please go to the designated page to complete the task"
      }
    }
}
```

The following languages are supported for this format: zh_cn, en_us, it_it, th_th, ko_kr, es_es, ja_jp, id_id, zh_hk, pt_br, de_de, fr_fr, zh_tw, ru_ru, hi_in, vi_vn. Developers only need to fill in the language text that needs to be supported.

For example, by setting "tip" to `{"en_us":"Please go to the designated page to complete the task"}`, the display effect is as follows:

![aaaaa.gif](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/f090d6286904dd4291875c98f1a74c36_2ToYS3pleL.gif?height=1392&lazyload=true&width=2272)

If both the tip and href are filled in, the href will be used, i.e. the "redirect behavior".

## How to Use Recurring Tasks?

If a task has a repeat rule, it will be a "recurring task". When a recurring task is completed, the system will automatically copy the completed task and create a new task with the same details, except that the due of the new task will be set according to the recurrence rule.

![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/85e3eba22c4cff910a0fdf7d87a84ae5_tcjPtkacJG.png?height=906&lazyload=true&width=2608)

For example, if a recurring task T1 is created with a due "2023-03-15 18:00:00" and a repeat rule of "daily recurrence", when T1 is completed, a copy of T1 called T2 will be automatically created, and the deadline of T2 will be set as "2023-03-16 18:00:00" according to the recurrence rule.

The `repeat_rule` field of the task can be set directly when creating a task, or an existing task can be converted to a recurring task by setting the `repeat_rule` field through the Update Task api. However, note that in order to set a task as a recurring task, the task's due time must be set first, otherwise the API will report an error.

To convert a recurring task to a normal task, simply set the `repeat_rule` to an empty string through the Update Task API.

The `repeat_rule` field is a subset of the RRULE string that complies with [RFC5545](https://datatracker.ietf.org/doc/html/rfc5545) specifications. For example:
* `FREQ=DAILY;INTERVAL=1` indicates that the new task's due time will be set to the same time on the next day of the completed task's due time.
* `FREQ=WEEKLY;INTERVAL=1;BYDAY=MO,TU,WE,TH,FR` indicates that the new task's due time will be set to the next working day (Monday to Friday) of the completed task's due time.
* `FREQ=WEEKLY;INTERVAL=2` indicates that the new task's due time will be set to the same time two weeks after the completed task's due time.
* `FREQ=MONTHLY;INTERVAL=1;WKST=MO;BYDAY=2FR` indicates that the new task's due time will be set to the same time on the second Friday of the next month after the completed task's due time.

Due to the fact that Lark tasks only support a subset of the RRULE specification, Lark tasks cannot guarantee support any RRULE string that conforms to RFC5545 specifications. To construct a `repeat_rule` that is compatible with tasks, a simple way is to set a repeat rule on the Lark App Task UI and then invoke the Get Task Details API to obtain the `repeat_rule` string.

When updating a task, you cannot update both `repeat_rule` and `completed_at` at the same time. This is because completing a task automatically deletes its repeat rule (hence the repeat rule icon for completed tasks disappears from the interface). Modifying both will result in undefined behavior, and the interface will report an error for such requests.

Similarly, you cannot clear `due` while setting `repeat_rule` when updating a task, because recurring tasks require `due` to calculate the due time after completion.
