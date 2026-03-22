---
title: "Introduction to access control resources"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar-acl/introduction"
service: "calendar-v4"
resource: "calendar-acl"
section: "Calendar"
updated: "1736822128000"
---

# Introduction to access control resources

Access controls (ACL) are used to manage member permissions for calendars. Multiple ACLs can be created within a calendar, and each ACL can set calendar access permissions for a member. The access permissions include:

- Visitors can only see calendar schedule busy and busy information
- Subscribers: can view all schedule details in the calendar.
- Editor: can create or modify schedules in the calendar.
- Administrator: Can manage calendar and sharing settings.

## Field description

The attributes contained in the calendar's access control resources are described below.

| Name | Type | Description |
| --- | --- | --- |
| `acl_id` | `string` | ACL unique identifier within a single calendar, read-only field. The acl_id under different calendar entities may be duplicated. **Example value**: "user_6843287928157667331" **Getting method**: Get access control list |
| `role` | `string` | The access level of the ACL, which is the member's access rights to the calendar.  **Optional values are**: - `free_busy_reader` visitors can only see busy and idle information. - `reader` subscriber can view all schedule details. - `writer` Editor, can create and modify schedules. - `owner` administrator can manage calendar and sharing settings. **Example value**: `owner` |
| `scope` | `acl_scope` | Effective scope of authority.  |
| `    ∟type` | `string` | The type of permission scope. Currently only `user` is supported.  **Optional values are**: - `user`: User, user resource definition can be found in User Resource Overview. **Example value**: `user` |
| `    ∟user_id` | `string` | User ID. When `type=user`, this parameter value must be set. For more information about user ID, please see User-related ID concepts.  **Example value**: "ou_742b382643a4075545422e87ff0d31fb" | ## Data example
```json
{
     "acl_id": "user_6843287928157667331",
     "role": "owner",
     "scope": {
         "type": "user",
         "user_id": "ou_742b382643a4075545422e87ff0d31fb"
     }
}
```
