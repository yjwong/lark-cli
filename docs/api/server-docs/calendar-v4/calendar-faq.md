---
title: "Calendar FAQ"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar-faq"
service: "calendar-v4"
resource: "calendar-faq"
section: "Calendar"
updated: "1736822100000"
---

# Calendar FAQ

This article summarizes common problems and solutions you may encounter while using the Calendar API.

## Calendar related

### Q: How to get the calendar subscription list?

You can call Get calendar list. The calendar obtained by this interface is the calendar that the user has currently subscribed to.

### Q: Does Lark Calendar support querying legal working days and holiday schedules?

Lark Calendar does not provide query functions for statutory working days and holidays. If there is a calendar with corresponding content known, you can call the Search calendar interface to find the corresponding calendar, and then call [Get schedule list]( /ssl:ttdoc/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar-event/list) interface to obtain schedule data and determine the day’s schedule based on the schedule data.

### Q: Lark client is bound to three-party calendars (Google, Exchange). Does it support unbinding through API operations?

Currently, only Exchange account binding and unbinding operations are supported, and tenant super administrator rights are required. Google does not support this yet.

## Schedule related

### Q: When obtaining all schedules within a specified time period, the complete data is not obtained in a single call, and calling the sync_token passed in again does not work. Why is it?

In the Get schedule list interface, the start/end time parameters (start_time/end_time) and the incremental synchronization token parameter (sync_token) Cannot be used at the same time. These two types of parameters represent two different synchronization methods:

- Start/end time parameters (start_time/end_time)

	Used to get the schedule list within a specified time period. For example, get all schedules from 2021-11-01 00:00:00 to 2021-11-02 00:00:00.
    
- Incremental synchronization token parameter (sync_token)

	Get the schedule list by schedule version number. The version number will be updated synchronously every time the schedule information is updated. Use sync_token to obtain the schedule information that subsequently changes within the specified version number.

### Q: Why can I see the schedule return, but cannot see the schedule title?

Whether you can see the schedule title is related to the following permission parameters.

- The `visibility` parameter in the Create Schedule interface is used to specify the public scope of the schedule. Ranges:

  - `public`: public, showing schedule details by default.
  - `default`: Default permission, users with `reader` permission can view details.
  - `private`: Private, even if the user has `reader` permission, the details cannot be viewed.

- The `role` parameter in Create access control, which is used to set the user's access permissions to the calendar (can be passed the `role` parameter in Query calendar information obtains the user's access rights to the calendar). `role` parameter value range:

  - `unknown`: Unknown permissions (i.e. the current `role` is not considered).
  - `free_busy_reader`: Visitors, can only see busy/idle information.
  - `reader`: Subscriber, view all schedule details.
  - `writer`: Editor, creates and modifies schedules.
  - `owner`: Administrator, manages calendar and sharing settings.

> **Note:** - When the schedule visibility is `public`, users can see the title regardless of their permissions on the current calendar.
> - When the schedule visibility is `private`, the title is visible to users who are `writer` and `owner` of the current calendar.
> - When the schedule visibility is `default`, the title is visible to users who are `reader`, `writer` and `owner` of the current calendar.

### Q: Why can’t I edit the schedule title or delete the schedule?

When a user is a schedule participant, he or she can obtain the schedule on his or her main calendar, but this does not mean that he or she has editing permissions for the schedule. Related documents:

- In Update schedule or Delete schedule interface documentation to read the notes to understand under what circumstances the schedule can be updated or deleted.
- To determine the current user's permissions on the calendar, you can call the Query Calendar Information interface.
- How to determine whether the current user is a schedule organizer, please refer to **How to determine whether the current user is a schedule organizer** in this article.

### Q: How to determine whether the current user is a schedule organizer?

Currently, you can check whether the participant is the organizer of the schedule through the Get schedule participant list interface.

1. The `calendar_id` in the path parameter needs to be passed in the primary calendar ID of the current user.
2. After obtaining the participant list, if the user with `is_organizer: true` matches the identity of the current user, it means that the current user is the organizer of the schedule.

### Q: Why can only one schedule be obtained for a recurring schedule?

Recurring schedules represent recurring rules through the schedule Rrule field. For detailed definitions, please refer to [RFC5545 protocol](https://datatracker.ietf.org/doc/html/rfc5545). For example, a schedule that repeats once a day corresponds to `RRULE: FREQ=DAILY;UNTIL=20220531T155959Z;INTERVAL=1`.

### Q: The schedule color is set by calling the interface, why does it not take effect?

The schedule color (`color` field) set by the calling interface is only effective for the current identity. If you need to change the color of the schedule seen by others, you need to use the other party's identity to modify it. Fields with similar logic include `reminder` (schedule reminder).

## Related to participants (including conference rooms)

### Q: How to reserve a conference room through the open platform OpenAPI?

You need to call the Add Schedule Participants interface to reserve a meeting room after creating the schedule, instead of passing it in when creating the schedule. `location` field.

When calling the Add Calendar Participant interface to reserve a conference room, you need to pay attention to the following:

- If you make a reservation as an application when calling the interface, you need to upload the `operate_id` field, which is used to specify the contact person of the conference room.
- Whether the reservation is successful in the end depends on whether the administrator has configured conference room reservation restrictions (including approval, time limits, and scope restrictions) in the management background, and whether other users are occupying the conference room at the same time.

### Q: Why does the reservation fail even though there is no schedule conflict when booking a conference room?

There may be a variety of reasons, including but not limited to:

1. The administrator configured **conference room reservation restrictions** in the management background.
    
    - The conference room is an approval conference room. Because OpenAPI does not support the approval function, when the administrator turns on the conference room approval switch in the management background, the user's conference room reservation request will fail.

    - The meeting room has a time range limit. When the administrator turns on the reservation time limit in the management background, the reservation will fail if the time limit is not met. For example: the daily reservation time range of the conference room is 08:00 ~ 23:00, and the reservation will fail when the time period is 03:00 ~ 04:00.

    - The conference room has set limits on the range that users can reserve. When the administrator sets the reservation range in the management background, if the current user is not within the reservation range, the reservation will fail.

    ![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/bfc80460a596b8bd2644f479a6cc835b_mAJ22DdZ2d.png?height=1264&lazyload=true&maxWidth=600&width=2068)

2. Concurrently book meeting rooms. Other users have reserved the current conference room first. Because the interval between the two requests is short, it will cause the user to obtain the conference room availability, but the conference room has not been reserved, but it has actually been reserved.

### Q: After booking a conference room, why can’t I know immediately whether the reservation is successful?

Calling the interface to reserve a conference room is an asynchronous process and may involve competition with other users. Therefore, after submitting the operation of booking a conference room, you need to wait a few seconds before returning to the final booking status.

### Q: Does it support schedule change events in which participants can reply to accept, reject, etc. operations?

You can subscribe to schedule change event. Data changes caused by participants clicking accept, reject, etc. will be included in the schedule within the scope of the change event.

## Error code related

### Q: How to use the error code returned after calling the interface?

You can find the matching error code in the corresponding help document of OpenAPI, or see General error code, and then fix the problem according to the error code description and troubleshooting suggestions.
