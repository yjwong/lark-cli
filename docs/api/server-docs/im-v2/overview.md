---
title: "Overview"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/group/im-v2/overview"
service: "im-v2"
resource: "overview"
section: "Feed"
updated: "1711958923000"
---

# Feed Opening overview

## Introduction

**Feed Opening** means displaying feed cards with richer form and content in the message list. This makes it easier to reach important messages by customizing the labels, buttons, beeps, temporary top statuses, etc. of the feed cards.

![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/2c9c47e11c874b1c989131714537f6e6_0nt9qqRw54.png?height=2994&lazyload=true&maxWidth=648&width=5648)

 
**Feed Opening advantages at a glance:**

| Before | After |
| --- | --- |
| -   Various types of messages are piling up, and important information is hard to reach. - New messages constantly refresh the message list, and important messages are drowned out. - Long message processing path leads to cumbersome processing and is easy to forget. - Notification tones are indistinguishable, so it is impossible to recognize the receipt of important messages. ![](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/a1b7c9623f9bd3326ebac8a0e7e3d3a8_cx9jhTCSWG.png?height=1800&lazyload=true&width=4000) | -   Eye-catching **buttons and labels** on feed cards make important information visible at a glance. - Important information can be **continuously displayed at the top of the message list** to prevent omission. - Exposed buttons help to handle important matters **with one click**. - **Customizable notification beeps** allow you to be the first to know about important information. ![](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/d378acda670b4b95d37f0376b11483c7_obuHQ9xrfZ.png?height=1800&lazyload=true&width=4000) | ## Application Scenarios

Feed Opening has various application scenarios. Below are a few typical application scenarios.

| Scenarios | Description | Illustration |
| --- | --- | --- |
| **Attendance clock-in reminder** | Pain point: There are many things to do during commuting time, so it is easy to forget to clock in and out, resulting in missing cards and absenteeism.  After access: -   Ten minutes before clocking in and out of work, the clocking robot displays on top of the message list, forming a strong reminder to avoid forgetting; - Configure the quick punch button to make the operation simple and convenient; - Cancel the top display after clocking in and out, so as not to disturb the daily work. |  |
| **Important schedule reminders** | Pain point: A large number of schedules leads to confusion in remembering which schedules must be attended, and it is easy to miss important meetings.  After access:： -   Schedule top display helps to get the schedule information for the first time; - Label the schedule to indicate the start time of the schedule, so as to arrange it in advance; - Display the Join Meeting button to join the meeting with one click. |  |
| **Complaint handling** | Pain point: Complaints need to be handled in a timely manner, but with a large volume of messages, it is difficult to find the corresponding information quickly.         After access: -   Complaints are displayed on top of the message, and other messages are clearly distinguished; - Set up customized beeps, so that the equipment standby state can also be the first time to understand the information touch. |  |
| **Employee training tasks** | Pain point: Employees have more trivial work, which can easily lead to forgetting important messages at a glance.  fter access: -   Training notifications are resident in the employee message list as a reminder, and are constantly on top, preventing employees from forgetting about training as they deal with other matters; - Configure buttons so that employees can go to the study with one click, shortening the operation path; - Reminders disappear automatically only after employees complete the training, greatly improving the task completion rate. |  |
| **Notification of goods distribution** | Pain point: Employees need to know the time and progress of goods delivery several times, so they need to enter the system several times. The operation is not convenient and flexible.  After access: -   When the goods start to be delivered, the delivery status and estimated delivery time are automatically displayed in the message list; - Employees can quickly sign for the delivery, check the incoming goods list and contact the driver through buttons, which makes the operation faster and more convenient; - The notification disappears automatically after confirming receipt. |  | ## Introduction of Terms

| Terms | Parameter | Description | Illustration |
| --- | --- | --- | --- |
| Feed card | ```app_feed_card``` | -   A special type of message display in the message list to make the message more visible. - There are 2 types of feed cards: - **App Feed Card:** A feed card created by the application. You can customize the appearance of the feed, the link to jump to when you click it, etc. - **Group Chat/Bot Feed Card:** Updates app bot sessions and bot group chats directly into feed cards. Only the card button and instant alert status can be customized. | App Feed Card   Group Chat / Bot Feed Card    |
| Instant reminder | ```time_sensitive``` | -   When the instant reminder state is open, the feed card is always displayed at the top of the message list and remains at the first position in the message list, without being flooded by new messages. - When instant reminder is closed, the feed is no longer displayed on top. - Application: Application feed cards, group & bot session feed cards. - Instant reminder status, true - open, false - closed. |       Instant reminder status open        Instant reminder status close    | ## Access Process
### App Feed Card:
| Steps | Introduction |
| --- | --- |
| 1. Create an app. | -   For creating a custom app, see Custom app development process. - For creating a store app, see Develop and publish store apps. |
| 2. Enable required scopes. | Before calling APIs, you need to obtain the access token and enable corresponding permission scopes. - Create, update, and delete app feed card - Activate instant notification capability |
| 3.  Get the user id of the target push user and create the feed card. | - You can customize the appearance of the feed cards, the links that jump when you click on them, etc. |
| 4.  Call the API to update the feed card. | -   This interface is used to update the avatar, title, preview, label status, button and other information of the information flow card. - **Field Update Strategy:** The specific fields to be updated are based on `update_fields`. The fields represented by `update_fields` are extracted from `app_feed_card` for updates. Fields not included in `update_fields` will not be updated. |
| 5.  Call the API to delete the feed card. | - After the notification task is completed, call this API to delete the corresponding feed card. | ### Group Chat/Bot Feed Cards:
| Steps | Introduction |
| --- | --- |
| 1. Create an app. | -   For creating a custom app, see Custom app development process. - For creating a store app, see Develop and publish store apps. |
| 2. Enable required scopes. | Before calling APIs, you need to obtain the access token and enable corresponding permission scopes.-   Create, update, and delete app feed card - Activate instant notification capability |
| 3. Choose an existing group/bot session and get the corresponding chat_id. | -   If you choose a group, you need to pull the app bot into the target group. |
| 4. Call the API to update the feed card buttons. | - You can add shortcut operation buttons, update buttons, or delete buttons to Groups & Bots feed card. |
| 5. Turn on "instant reminder" if you want to stay on top of the list. | -   You need to get the user_id of the user you want to feed to, and when you call this interface, the card will be fed to the top of the list of those users. - If you want to cancel the topping, call it again and close the instant reminder state. |
