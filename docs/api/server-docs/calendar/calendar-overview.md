---
title: "Calendar overview"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uETM3YjLxEzN24SMxcjN"
section: "Calendar"
updated: "1736822095000"
---

# Calendar overview

Based on the Lark calendar function, the calendar API opens up the ability to query calendars, schedules, free and busy entities and other entities. Developers can call the calendar API to operate corresponding data as "applications" or "users". Through the calendar API, you can implement a variety of functions, such as:

- Create an event or meeting room reservation

- Invite Lark users or third-party users to participate in the schedule

- Query user freebusy

- Synchronize user leave status

- Bind or unbind an exchange account

## Customer case

The open platform includes calendar business integration solutions:

- Lalamove: [Lark smart meeting makes managing meetings easy](https://open.larksuite.com/solutions/detail/meetings)

## Access process

The basic access process of the Calendar API is shown in the figure below. For detailed API access process, see Process Overview.

![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/7c27953b2c7bfc8f8766b52721eee857_SDYnPkPfRd.png?height=214&lazyload=true&maxWidth=700&width=2920)

## Development tutorial

Experience the scenario-based sample tutorial and learn how to use the Calendar API to implement calendar schedule management.

- Quickly build calendar events
	 
	![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/d29f4c877c789d86b054b8d348098fa8_XkzMiQGHUG.png?height=1342&lazyload=true&maxWidth=600&width=2738)

- Auto-fetch calendar events

	![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/6f216d21cac15e40840358d947dbf2a0_fDb57YvU7p.png?height=400&lazyload=true&maxWidth=600&width=752)

## Usage restrictions

- The maximum number of calendars that a single user or a single application can subscribe to is 1000.
- The maximum number of participants that can be added to each event is 3,000.

## Resources introduction

The calendar business domain is opened with resources as the center. A calendar is an entity used to aggregate events. After subscribing to the calendar, users can create events and invite participants in the calendar. The resource relationship diagram is as follows:

![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/72f65d2d8a83ea4cdba61f34685e6022_B9MNEw66Ja.png?height=1916&lazyload=true&maxWidth=750&width=2374)

The resources and resource definitions included in the calendar business domain are as follows:

| Calendar | A calendar is a container that contains a series of related events and is an aggregated entity of the same class of events. |
| --- | --- |
| Access control list (ACL) | The member list of the calendar. Different members can be granted different calendar operation permissions. |
| Event | A event contains a certain date or time range and is the basic entity for meeting reservations. It is divided into ordinary events and recurring events. Additionally, the following types of events exist: - Exception event: After a event in a recurring event is modified and is different from other recurring events, the event is called an exception event. - Leave event: A special event that will be displayed on the calendar when the user asks for leave. |
| Event attendee | Object involved in the event, including users, meeting rooms, groups, and third-party emails. > **Note**: The way to add a meeting room to the event is also by adding participants. |
| Event attendee chat member | Group members participating in the event. |
| Freebusy | The calendar's busy and free information for a certain time period. |
| Timeoff | Create leave events for a certain date or time period for users. |
| Setting | Generate a CalDAV account password for the specified device, which is used to synchronize the client calendar information to the local device calendar. |
| Exchange binding and unbinding | Complete the binding or unbinding of the exchange account to the client account. | The calendar resources in Lark client are as follows:

- Area ① in the figure shows user calendars, including user-managed calendars and subscribed calendars.
- The ② area in the figure displays the schedule busy and busy information within the current calendar time period. After selecting a schedule, you can view detailed information such as schedule content, participants, and conference rooms.

![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/89da7d10979d3712e25edfb00cfc06d9_Ke1Le1nLlO.png?height=1344&lazyload=true&maxWidth=600&width=2878)

## Method list

A list of APIs and events for all resources within the calendar business domain is provided below.

> **Note:** The **Public** mentioned in the table in this article refers to the store app, and **Custom** refers to the enterprise's custom app. For application type description, see Application Types.

### Calendar

#### API list

| `Get primary calendar `POST` /open-apis/calendar/v4/calendars/primary > Get the user's primary calendar entity information ` | `calendar:calendar:readonly` | `tenant_access_token` `user_access_token` | **✓** | **✓** |
| --- | --- | --- | --- | --- |
| `Create calendar` `POST` /open-apis/calendar/v4/calendars > Create a calendar entity as the current identity | `calendar:calendar` | `tenant_access_token` `user_access_token` | **✓** | **✓** |
| `Delate Calendar` `DELETE` /open-apis/calendar/v4/calendars/:calendar_id > Delete the specified calendar entity as the current identity | `calendar:calendar` | `tenant_access_token` `user_access_token` | **✓** | **✓** |
| `Get Calendar` `GET` /open-apis/calendar/v4/calendars/:calendar_id > Get the specified calendar entity as the current identity | `calendar:calendar:readonly` `calendar:calendar` | `tenant_access_token` `user_access_token` | **✓** | **✓** |
| `Get Calendar List` `GET` /open-apis/calendar/v4/calendars >Get the list of calendars on the current identity | `calendar:calendar:readonly` `calendar:calendar` | `tenant_access_token` `user_access_token` | **✓** | **✓** |
| `Update Calendar` `PATCH` /open-apis/calendar/v4/calendars/:calendar_id >Update the specified calendar entity information as the current identity | `calendar:calendar` | `tenant_access_token` `user_access_token` | **✓** | **✓** |
| `Search Calendar` `POST` /open-apis/calendar/v4/calendars/search >Search calendar information as an app | `calendar:calendar:readonly` `calendar:calendar` | `tenant_access_token` | **✓** | **✓** |
| `Subscribe Calendar` `POST` /open-apis/calendar/v4/calendars/:calendar_id/subscribe >Subscribe to the specified calendar with the current identity | `calendar:calendar` | `tenant_access_token` `user_access_token` | **✓** | **✓** |
| `Unsubscribe Calendar` `POST` /open-apis/calendar/v4/calendars/:calendar_id/unsubscribe >Unsubscribe the specified calendar with the current identity | `calendar:calendar` | `tenant_access_token` `user_access_token` | **✓** | **✓** |
| `Subscribe Calendar Changes` `POST` /open-apis/calendar/v4/calendars/subscription >Establish an event subscription relationship with the current identity's calendar resource | `calendar:calendar:readonly` `calendar:calendar` | `user_access_token` | **✓** | **✓** |
| `Unsubscribe Calendar Changes` `POST` /open-apis/calendar/v4/calendars/unsubscription >Cancel the event subscription relationship with the current identity's calendar resource | `calendar:calendar:readonly` `calendar:calendar` | `user_access_token` | **✓** | **✓** | #### Event List

| `Calendar is changed` | When there is a calendar change in the subscriber's calendar list | `calendar:calendar:readonly` `calendar:calendar` | calendar.calendar.changed_v4 | **✓** | **✓** |
| --- | --- | --- | --- | --- | --- | ### ACL

#### API list

| `Create ACL `POST` /open-apis/calendar/v4/calendars/:calendar_id/acls > Add access control permissions with the current identity, that is, add calendar members ` | `calendar:calendar` | `tenant_access_token` `user_access_token` | **✓** | **✓** |
| --- | --- | --- | --- | --- |
| `Delete ACL `DELETE` /open-apis/calendar/v4/calendars/:calendar_id/acls/:acl_id > Delete the control permission of the calendar with the current identity, that is, delete the calendar member ` | `calendar:calendar` | `tenant_access_token` `user_access_token` | **✓** | **✓** |
| `Acquire ACL` `GET` /open-apis/calendar/v4/calendars/:calendar_id/acls > Get the list of control permissions of the calendar, that is, the list of calendar members | `calendar:calendar:readonly` `calendar:calendar` | `tenant_access_token` `user_access_token` | **✓** | **✓** |
| `Subscribe ACL Event` `POST` /open-apis/calendar/v4/calendars/:calendar_id/acls/subscription > Calendar member change event on the specified calendar | `calendar:calendar:readonly` `calendar:calendar` | `user_access_token` | **✓** | **✓** |
| `Unsubscribe ACL Event` `POST` /open-apis/calendar/v4/calendars/:calendar_id/acls/unsubscription > Cancel the monitoring of calendar member change events on the specified calendar | `calendar:calendar:readonly` `calendar:calendar` | `user_access_token` | **✓** | **✓** | #### Event List

| `ACL is created` | When an ACL is created on the subscribed calendar | `calendar:calendar:readonly` `calendar:calendar` | calendar.calendar.acl.created_v4 | **✓** | **✓** |
| --- | --- | --- | --- | --- | --- |
| `ACL is deleted` | When an ACL is removed on a subscribed calendar | `calendar:calendar:readonly` `calendar:calendar` | calendar.calendar.acl.deleted_v4 | **✓** | **✓** | ### Event

#### API list

| `Create event `POST` /open-apis/calendar/v4/calendars/:calendar_id/events > Create an event on the calendar ` | `calendar:calendar` | `tenant_access_token` `user_access_token` | **✓** | **✓** |
| --- | --- | --- | --- | --- |
| `Delete Event `DELETE` /open-apis/calendar/v4/calendars/:calendar_id/events/:event_id > Delete an event from the calendar ` | `calendar:calendar` | `tenant_access_token` `user_access_token` | **✓** | **✓** |
| `Get Event` `GET` /open-apis/calendar/v4/calendars/:calendar_id/events/:event_id > Get an event information on the calendar | `calendar:calendar:readonly` `calendar:calendar` | `tenant_access_token` `user_access_token` | **✓** | **✓** |
| `Get Event List` `GET` /open-apis/calendar/v4/calendars/:calendar_id/events > Get the list of events on the calendar | `calendar:calendar:readonly` `calendar:calendar` | `tenant_access_token` `user_access_token` | **✓** | **✓** |
| `Update Event` `PATCH` /open-apis/calendar/v4/calendars/:calendar_id/events/:event_id >Update an event on the calendar | `calendar:calendar` | `tenant_access_token` `user_access_token` | **✓** | **✓** |
| `Search Event` `POST` /open-apis/calendar/v4/calendars/:calendar_id/events/search >Search related events on a calendar as the current user | `calendar:calendar:readonly` `calendar:calendar` | `user_access_token` | **✓** | **✓** |
| `Subscribe Event Changes` `POST` /open-apis/calendar/v4/calendars/:calendar_id/events/subscription >Subscribe to the schedule change event on the specified calendar as a user | `calendar:calendar:readonly` `calendar:calendar` | `user_access_token` | **✓** | **✓** |
| `Unsubscribe Event Changes` `POST` /open-apis/calendar/v4/calendars/:calendar_id/events/unsubscription >Unsubscribe the schedule change event on the specified calendar as a user | `calendar:calendar:readonly` `calendar:calendar` | `user_access_token` | **✓** | **✓** | #### Event List

| `Schedule is changed` | When there is a schedule change on the subscribed user's calendar | `calendar:calendar:readonly` `calendar:calendar` | calendar.calendar.event.changed_v4 | **✓** | **✓** |
| --- | --- | --- | --- | --- | --- | ### Event Attendee

#### API list

| `Create Schedule Participants `POST` /open-apis/calendar/v4/calendars/:calendar_id/events/:event_id/attendees > Add participants to events in batches ` | `calendar:calendar` | `tenant_access_token` `user_access_token` | **✓** | **✓** |
| --- | --- | --- | --- | --- |
| `Get a List of Schedule Participants `GET` /open-apis/calendar/v4/calendars/:calendar_id/events/:event_id/attendees > Get a list of participants for an event ` | `calendar:calendar:readonly` `calendar:calendar` | `tenant_access_token` `user_access_token` | **✓** | **✓** |
| `Delete event invitees` `POST` /open-apis/calendar/v4/calendars/:calendar_id/events/:event_id/attendees/batch_delete > Delete participants from an event in batchs | `calendar:calendar` | `tenant_access_token` `user_access_token` | **✓** | **✓** | ### Event Chat Member

#### API list

| `Get the List of Members of the Schedule Participation Group `GET` /open-apis/calendar/v4/calendars/:calendar_id/events/:event_id/attendees/:attendee_id/chat_members > Get the list of members of the chat participants of the schedule ` | `calendar:calendar:readonly` `calendar:calendar` | `tenant_access_token` `user_access_token` | **✓** | **✓** |
| --- | --- | --- | --- | --- | ### Free Busy

#### API list

| `Query Busy `POST` /open-apis/calendar/v4/freebusy/list > Query the freebusy information of the user's primary calendar or meeting room ` | `calendar:calendar:readonly` `calendar:calendar` | `tenant_access_token` | **✓** | **✓** |
| --- | --- | --- | --- | --- | ### Timeoff

#### API list

| `Create Timeoff Event `POST` /open-apis/calendar/v4/timeoff_events > Create a vacation event for a specified user ` | `calendar:timeoff` | `tenant_access_token` | **✓** | **✓** |
| --- | --- | --- | --- | --- |
| `Delete Timeoff Event `DELETE` /open-apis/calendar/v4/timeoff_events/:timeoff_event_id > Delete a specified vacation event ` | `calendar:timeoff` | `tenant_access_token` | **✓** | **✓** | ### Setting

#### API list

| `Generate CalDAV Configuration `POST` /open-apis/calendar/v4/settings/generate_caldav_conf > Generate a CalDAV account password for the current user ` | `calendar:calendar` `calendar:setting:generate_caldav_conf` | `user_access_token` | **✓** | **✓** |
| --- | --- | --- | --- | --- | ### Exchange Binding

#### API list

| `Exchange Binding `POST` /open-apis/calendar/v4/exchange_bindings > Binding the Exchange account to the Lark account ` | `calendar:calendar:readonly` `calendar:calendar` | `user_access_token` | **✓** | **✓** |
| --- | --- | --- | --- | --- |
| `Exchange Binding Status `GET` /open-apis/calendar/v4/exchange_bindings/:exchange_binding_id > Get the binding status of an Exchange account ` | `calendar:calendar:readonly` | `user_access_token` | **✓** | **✓** |
| `Exchange Unbinding` `DELETE` /open-apis/calendar/v4/exchange_bindings/:exchange_binding_id > Delete the binding relationship between the Exchange account and Lark account | `calendar:calendar` | `user_access_token` | **✓** | **✓** |
