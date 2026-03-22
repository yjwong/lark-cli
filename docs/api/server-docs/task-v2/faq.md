---
title: "FAQ"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/task-v2/faq"
service: "task-v2"
resource: "faq"
section: "Tasks v2"
updated: "1710237942000"
---

# FAQ

## Can I Access All Data of the Entire Tenant through tenant_access_token?
No.

`tenant_access_token` corresponds to the application identity. Internally，this identity is no different from ordinary user identities. You can think of it as a special user.

In other words, `tenant_access_token` **has no special privileges** and cannot access the task data of other ordinary users or other applications, let alone be used as a "tenant administrator" role.

## How to Count the Task Data within the Tenant?

There is no "tenant administrator" role, so it cannot directly support counting all data of the entire tenant. We also do not recommend developing such functionality in this way, becasue there are both work-related tasks and personal tasks (recorded in the company's tenant) in Lark. Counting all task data of the tenant would bring about data access volume issues, and it would be difficult to directly improve the overall goals of enterprise task response speed, completion rate, and punctuality rate.

We recommend using tasklists to maintain and manage tasks. The process is as follows:

1. Create several tasklists for some work content through the Lark App manually. Add all employees involved in the work as collaborative members for each tasklist.
2. The tasklist owner or members manually create tasks and fill in necessary information such as descriptions and dues.
3. Use the [Add List Members](/ssl:/ttdoc/uAjLw4CM/ukTMukTMukTM/task-v2/tasklist/add_members) api to add the application to the list of members. If it is only for counting, the "viewer" role is sufficient (see below "Can applications be members of tasks/tasklists?").
4. Use the application to write a program that runs periodically, calling the [Get Tasks Of Tasklist](/ssl:/ttdoc/uAjLw4CM/ukTMukTMukTM/task-v2/tasklist/tasks) api to get all tasks of the tasklist, and count the data based on whether the tasks are completed, overdue, etc.

However, please note that **tasklists are not designed to accommodate a large number of tasks**. Calling the Get Tasks Of Tasklist API for a tasklist containing tens of thousands of tasks may lead to timeouts, rate limiting, or other server errors. Therefore, it is recommended to create multiple tasklists as needed, with each tasklist containing a suitable number of tasks. This also facilitates the actual use of employees.

## Can Applications be Members of Tasks/Tasklists?

Yes.

Both ordinary users and applications are considered "users" within the task system, with no difference. Therefore, an application can be set as a task follower or tasklist member. You can set an application to a member of a task/tasklist by calling add task members or tasklist members api. For example, the following code adds an application as a "viewer" for the tasklist.

```json
POST /task/v2/tasklists/:tasklist_guid/add_members
{
  "members": [
    {
      "id": "cli_a2ee5b213cf8912b",
      "type": "app",
      "role": "viewer"
    }
  ]
}
```

This usage is usually for automation scenarios, such as creating a tasklist through the Lark App, allowing team members to fill tasks in the tasklist, and updating the task progress through task descriptions/comments, eventually completing the task. In this case, an application can be added as a viewer to the tasklist. The application can then use the `tenant_access_token` to call the tasklist and task details api to periodically count the overall progress of the tasklist, sending the results as message cards to the list owner or synchronizing them with third-party systems.

If you do not add the application as member of the tasklist, the application will have no permission to access the tasklist and the task data in that tasklist.

## How to Get the Task GUID in the App UI?
Developers may need to operate a task in the Lark App (such as editing the summary, deleting, etc.). The prerequisite for operation is to obtain the task GUID.

You can enter the task details page, click "share", and enter the "Task Link" page to see the task's share link. The 36-digit string after guid= in the link is the task GUID.
| 1) Click "Share" in the task details | 2) View the task's share link |
|-|-|
|![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/b85d81c949ad42dcbc227de6948ad5b5_4AXKBYbWxH.png?height=706&lazyload=true&width=1188)|![分享3.jpg](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/f6b3618d34235b47e69c364861ce023b_EHTfcWvFUT.jpg?lazyload=true&width=2860&height=1166)| The task GUID in the above picture is `07df1984-e38d-4710-90e6-493a2412dcfe`.

## How to get the tasklist GUID in the Lark App?
Developers may need to operate a list in the Lark App (such as modifying the name, adding members, etc.). The prerequisite for operation is to obtain the tasklist GUID.

You can access a tasklist detail UI, click share.
Enter the checklist interface and click on Share.
|1) Click on "Share" in the checklist interface | 2) Click on Copy Link| 3) Then paste it into a text editor |
|-|-|-|
|![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/3cef29219643d344d7a113feaa9178dd_3AZpvLje4C.png?height=536&lazyload=true&width=1436)|![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/3511be262224e7d7e6047220680d268f_nWR8PsjyzH.png?height=506&lazyload=true&width=1020)|![分享2.jpg](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/fd63ba6e79a326191d57b91d81a1c708_9nkVmFdTU3.jpg?lazyload=true&width=2946&height=364)| The 36-character string following "guid=" in the obtained link is the checklist's GUID.

In the example above, the checklist's GUID is `3606cb71-a942-4098-b7a0-c59869d5144d`.

## Can Task API Support Individual Completion for a Countersign Task?

Currently not supported, but it is in the roadmap.
