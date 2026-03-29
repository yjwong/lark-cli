---
title: "Event invitee resource introduction"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar-event-attendee/introduction"
service: "calendar-v4"
resource: "calendar-event-attendee"
section: "Calendar"
updated: "1736835321000"
---

# Event invitee resource introduction

Event invitees refer to the resource entities participating in the event. The types of invitees supported in the event include: users, group chats, meeting rooms, mailboxes.

## Usage limit

Each event can include up to 3000 invitees.

## Field description

| Name | Type | Description |
| --- | --- | --- |
| `attendee_id` | `string` | Invitee ID, the unique identifier of the event invitee in the current event. This ID can be used to delete event invitees, or to query the group member information of group type invitees. How to obtain invitee ID: - When calling the Add calendar invitee interface, the invitee ID can be obtained from the returned result. - Call the Get event invitee list interface to query the invitee ID in the event. Different types of invitees have different attendee_id formats. The description is as follows: - `user_xxxxxx` represents user type invitee.  - `chat_xxxxxx` represents group type invitees and can be used to query group member information within the group.  - `resource_xxxxxx` represents conference room type invitees.  - `third_party_xxxxxx` represents an external mailbox type invitee.  **Example value**: "user_xxxxxx" |
| `type` | `string` | Invitee type. **Optional values**: - `user`: user - `chat`: group chat - `resource`: meeting room - `third_party`: email **Example value**: "user" |
| `rsvp_status` | `string` | RSVP status of the invitee. **Optional values**: - `needs_action`: The invitee has not responded yet, or it means that the conference room is being reserved. - `accept`: The invitee replies to accept, or indicates that the conference room reservation is successful - `tentative`: invitee’s reply is pending - `decline`: The invitee refuses or indicates that the conference room reservation failed. - `removed`: The invitee or meeting room has been removed from the event **Example value**: "accept" |
| `is_optional` | `boolean` | Whether the invitee is optional, this parameter value does not take effect on the group members. **Optional values**: - true - false |
| `is_organizer` | `boolean` | Whether the invitee is the event organizer. **Optional values**: - true - false |
| `is_external` | `boolean` | Whether the invitee is an external invitee. Editing is not supported by external invitees. **Optional values**: - true - false |
| `display_name` | `string` | Invitee name. **Example value**: Li Jian |
| `chat_members` | `attendee_chat_member[]` | The group member information of the group, the group type invitee (type is chat) is valid. Group members do not support editing. |
| ∟ `rsvp_status` | `string` | Group invitee RSVP status. **Optional values**: - `needs_action`: The invitee has not responded yet - `accept`: Invitee replies accept - `tentative`: Invitee’s reply is pending - `decline`: The invitee refuses to reply - `removed`: The invitee has been removed from the event **Example value**: "needs_action" |
| ∟ `is_optional` | `boolean` | Whether the invitee is optional. **Optional values**: - true - false |
| ∟ `display_name` | `string` | Invitee name. **Example value**: Li Jian |
| ∟ `is_organizer` | `boolean` | Whether the invitee is the event organizer. **Optional values**: - true - false |
| ∟ `is_external` | `boolean` | Whether the invitee is an external invitee. **Optional values**: - true - false |
| `user_id` | `string` | User type The user ID of the invitee. As a return value, if is_external is true, this parameter will only return the user's open_id or union_id. For information about user IDs, see User-related ID concepts. **Example value**: "ou_xxxxxxxx" |
| `chat_id` | `string` | Group type The group ID of the invitee. For details about group ID, please refer to Group ID Description. **Example value**: "oc_xxxxxxxxx" |
| `room_id` | `string` | Room Type The room ID of the invitee. **Example value**: "omm_xxxxxxxx" |
| `third_party_email` | `string` | External email type The email address of the invitee. **Example value**: "wangwu@email.com" |
| `operate_id` | `string` | If the event is created using the application identity, this parameter is used to specify the room contact ID and the contact will be displayed in the room view information. **Example value**: "ou_xxxxxxxx" |
| `resource_customization` | `calendar.attendee.resource_customization[]` | Personalized configuration of meeting rooms. When the added conference room is configured with a reservation form, you can configure the form information through this parameter. |
| ∟ `index_key` | `string` | A unique ID for each configuration. **Example value**: "16281481596100" |
| ∟ `input_content` | `string` | When the personalized configuration type is fill-in-the-blank, this parameter needs to be filled in. **Example value**: "xxx" |
| ∟ `options` | `customization.option[]` | Options for each configuration. |
| ∟ `option_key` | `string` | A unique ID for each option. **Example value**: "16281481596185" |
| ∟ `others_content` | `string` | When the option type of personalized configuration is other options, this parameter needs to be filled in. **Example value**: "xxx" | ### Data example

```json
{
    "type": "user",
    "attendee_id": "user_xxxxxx",
    "rsvp_status": "needs_action",
    "is_optional": true,
    "is_organizer": true,
    "is_external": false,
    "display_name": "Li Jian",
    "chat_members": [
        {
            "rsvp_status": "needs_action",
            "is_optional": true,
            "display_name": "Group",
            "is_organizer": false,
            "is_external": false
        }
    ],
    "user_id": "ou_xxxxxxxx",
    "chat_id": "oc_xxxxxxxxx",
    "room_id": "omm_xxxxxxxx",
    "third_party_email": "wangwu@email.com",
    "operate_id": "ou_xxxxxxxx",
    "resource_customization": [
        {
            "index_key": "16281481596100",
            "input_content": "Xxx",
            "options": [
                {
                    "option_key": "16281481596185",
                    "others_content": "Xxx"
                }
            ]
        }
    ]
}
```
