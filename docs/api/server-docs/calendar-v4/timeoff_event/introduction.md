---
title: "Introduction to leave event resources"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/timeoff_event/introduction"
service: "calendar-v4"
resource: "timeoff_event"
section: "Calendar"
updated: "1736837315000"
---

# Introduction to leave event resources

Leave event is a special event. You can call the leave resource API to create or delete leave events for users. After creating a leave event for a user, in addition to displaying the leave event on the calendar, leave information will also be displayed on the user's personal information.

## Field description

The attributes included in the leave event are described below.

| Name | Type | Description |
| --- | --- | --- |
| `timeoff_event_id` | `string` | Leave event ID. Obtain it through Create leave event, which can be used to delete the leave event later. **Example value**: "timeoff:XXXXXX-XXXX-0917-1623-aa493d591a39-XXXXXX" |
| `user_id` | `string` | User ID. For information about user IDs, see User-related ID concepts. **Example value**: "ou_XXXXXXXXXX" |
| `timezone` | `string` | Time zone information. **Example value**: "Asia/Shanghai" |
| `start_time` | `string` | Leave start time. Any of the following formats are supported: - Second-level timestamp: The leave event set through the timestamp is a normal event, that is, leave is requested by the hour. Value example `1609430400` - Date format: The leave event set by date is a full-day event. Value example `2021-01-01` Note: The time format selected for start_time and end_time must be consistent, otherwise it will be invalid. **Example value**: "2021-01-01" |
| `end_time` | `string` | Leave end time. Any of the following formats are supported: - Second-level timestamp: The leave event set through the timestamp is a normal event, that is, leave is requested by the hour. Value example `1609430400` - Date format: The leave event set by date is a full-day event. Value example `2021-01-01` Note: The time format selected for start_time and end_time must be consistent, otherwise it will be invalid. **Example value**: "2021-01-01" |
| `title` | `string` | Leave event title. |
| `description` | `string` | Leave event description. | ### Data example
```json
{
    "timeoff_event_id": "Timeoff: XXXXXX-XXXX -0917-1623- aa493d591a39-XXXXXX",
    "user_id": "ou_XXXXXXXXXX",
    "timezone": "Asia/Shanghai",
    "start_time": "2021-01-01",
    "end_time": "2021-01-01",
    "title": "Time Off (Full Day)/1-Day Time Off",
    "description": "If this event is deleted, the corresponding "leave" label in the flying book will automatically disappear, and the leave application in the leave application system will not be revoked."
}
```
