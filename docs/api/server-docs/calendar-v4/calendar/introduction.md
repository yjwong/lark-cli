---
title: "Calendar resources introduction"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar/introduction"
service: "calendar-v4"
resource: "calendar"
section: "Calendar"
updated: "1736822105000"
---

# Calendar resource introduction

Calendar resources include the resources of the calendar itself and the schedule resources it contains. Multiple calendars can be created, and each calendar has attributes such as title, color, type, and public scope. At the same time, it supports querying the busy and free information of the schedules in each calendar. With the series of APIs provided by the open platform for creating, subscribing, and querying, you can manage calendar resources and query schedule status.

## Basic concept

Before officially using the API, you need to understand the following basic concepts.

| Concept | Introduction |
| --- | --- |
| Identity | Identity is divided into user identity and application identity. A user or an application (the robot capability needs to be turned on) can create or subscribe to one or more calendars. > The open platform has designed an access token (access_token) mechanism to ensure the security of API calls. For details, please see Get Access Credentials. When you actually call the API, you need to use access credentials to indicate your current identity. |
| Calendar | Lark calendars are mainly divided into two categories: main calendar and shared calendar. Synchronization of third-party calendars (e.g. Exchange) is also supported. - Main calendar: The default calendar owned by the user or application (the robot capability needs to be turned on), has the same name as the user or application, and cannot be deleted. - Shared calendar: Calendars created by users or applications (bot capabilities need to be turned on) can be customized with attributes such as name, description, and public scope when they are created. The calendar supports adjusting the visibility, so you can use the specified calendar as a personal private calendar, or a public calendar shared with other members of the team. |
| Subscribe | Users can subscribe to public calendars. After subscribing, users can view or manage the schedules in the calendar, facilitating team collaboration and communication. |
| Working hours | The working time customized by the user in the client calendar settings is used to mark working time periods and non-working time periods. |
| Busy | Calendar busy and free information within a certain period of time. |
| Event | The calendar resource contains a calendar change event. By subscribing to this event, you can obtain the user's calendar changes in a timely manner and perform corresponding business processing. | ## Usage restrictions

The maximum number of calendars that a single user or application can subscribe to is 1,000.

## Field description

The properties contained in the calendar resource itself are described below.

| Name | Type | Description |
| --- | --- | --- |
| `calendar_id` | `string` | Calendar ID is the unique identifier of the calendar entity and is a read-only field. You can query, delete or update calendar information through this ID.  **Example value**: "larksuite.com_xxxxxxxxxx@group.calendar.larksuite.com" **method of obtaining**: - Create shared calendar - Query primary calendar information - Query calendar list - Search calendar |
| `summary` | `string` | Calendar title. **Example value**: "Test Calendar" **Data verification rules**: - **Maximum length**: `255` characters |
| `description` | `string` | Calendar description. **Example value**: "Create calendar using open interface" **Data verification rules**: - **Maximum length**: `255` characters |
| `color` | `int` | Calendar color, the value is represented by an int32 color RGB value, in which bits 24 ~ 31 are transparency, bits 16 ~ 23 are red, bits 8 ~ 15 are green, and bits 0 ~ 7 are blue. For example, -11034625 represents RGB values (87, 159, 255). **Notice**: - The calendar color will be mapped to the closest color on the Lark client color palette for display. - This color only takes effect for the current identity (the current identity is determined by the access credentials when calling the API). **Example value**: -11034625 |
| `summary_alias` | `string` | Calendar note name. **Note**: Adding or modifying the remark name only takes effect for the current identity (the current identity is determined by the access credentials when calling the API). **Example value**: "Calendar note name" **Data verification rules**: - **Maximum length**: `255` characters |
| `type` | `string` | Calendar type, read-only field. **Optional values are**: - `primary`: The primary calendar of the user or application - `shared`: shared calendar created by user or application - `google`: User-bound Google calendar - `resource`: conference room calendar - `exchange`: User-bound Exchange calendar - `unknown`: unknown type **Example value**: `primary` |
| `permissions` | `string` | Calendar visibility.  **Optional values are**: - `private`: private - `show_only_free_busy`: Only display busy and busy information - `public`: Others can view schedule details **Example value**: `show_only_free_busy` |
| `role` | `string` | The current identity (the current identity is determined by the access credentials when calling the API) has access to this calendar, read-only field.  **Optional values are**: - `free_busy_reader`: Visitors can only see busy and idle information - `reader`: Subscriber, can view all schedule details - `writer`: Editor, can create and modify schedules - `owner`: Administrator, can manage calendar and sharing settings - `unknown`: unknown permissions **Example value**: `owner` |
| `is_deleted` | `boolean` | Whether the calendar has been marked for deletion for the current identity (the current identity is determined by the access credentials when calling the API). Read-only field.  **Example value**: false |
| `is_third_party` | `boolean` | Whether the current calendar is third-party data. Three-party calendars and schedules only support reading, not writing, that is, this field is a read-only field.  **Example value**: false | ## Data example

```json
{
     "calendar_id": "larksuite.com_xxxxxxxxxx@group.calendar.larksuite.com",
     "color": -11034625,
     "description": "Create calendar using open interface",
     "permissions": "private",
     "role": "owner",
     "summary": "Test Calendar",
     "summary_alias": "Calendar note name",
     "type": "primary",
     "is_deleted": false,
     "is_third_party": false
}
```
