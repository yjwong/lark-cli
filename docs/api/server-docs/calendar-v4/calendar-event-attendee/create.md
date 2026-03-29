---
title: "Create event invitees"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar-event-attendee/create"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/calendar/v4/calendars/:calendar_id/events/:event_id/attendees"
service: "calendar-v4"
resource: "calendar-event-attendee"
section: "Calendar"
rate_limit: "1000 per minute、50 per second"
scopes:
  - "calendar:calendar"
  - "calendar:calendar.event:update"
field_scopes:
  - "contact:user.employee_id:readonly"
updated: "1736835321000"
---

# Create event invitees

Use this interface to add one or more invitees to a specified event with the current identity (app or user). Invitee types include users, groups, meeting rooms, and emails.

> - The current identity is determined by the Token type of Header Authorization. tenant_access_token refers to the application identity, and user_access_token refers to the user identity.
> - If you use the application identity to call this interface, you need to ensure that the application turns on bot capability.
> - The current identity needs to have writer or owner permissions on the calendar, and the calendar type can only be primary or shared. You can call the Query Calendar Information interface to obtain the calendar type and the current identity's access rights to the calendar.
> - The current identity needs to be the organizer of the event, or a invitee of the event and make sure that the event is set with the **Invitees can invite other invitees** permission. You can call the Get event interface to obtain the invitee permissions (attendee_ability) of the event.
> - Newly added event invitees must be in the same enterprise as the event organizer.
> - Each event is limited to 3,000 invitees.
> - After using this interface to add a conference room, the conference room will enter an asynchronous reservation process. That is, the end of the request does not mean that the conference room reservation is successful. You need to check the reservation status of the conference room later.
> - After the conference room administrator ability is enabled, the administrator can reserve a conference room without being restricted by the conference room reservation range (it is currently not supported to reserve a conference room for other members' events as an administrator).

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/calendar/v4/calendars/:calendar_id/events/:event_id/attendees |
| HTTP Method | POST |
| Rate Limit | 1000 per minute、50 per second |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `calendar:calendar` `calendar:calendar.event:update` |
| Required field scopes | > The response body of the API contains the following sensitive fields, and they will be returned only after corresponding scopes are added. If you do not need the fields, it is not recommended that you request the scopes. `contact:user.employee_id:readonly` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `calendar_id` | `string` | The calendar ID corresponding to the event. To learn more, see Calendar ID Introduction. **Example value**: "larksuite.com_xxxxxxxxxx@group.calendar.larksuite.com" |
| `event_id` | `string` | Event ID. The event ID is returned when creating a event. You can also call the following interface to obtain the ID of a certain calendar. - Get event list - Search event **Example value**: "xxxxxxxxx_0" | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `user_id_type` | `string` | No | User ID categories **Example value**: open_id **Optional values are**:  - `open_id`: Identifies a user to an app. The same user has different Open IDs in different apps. How to get Open ID - `union_id`: Identifies a user to a tenant that acts as a developer. A user has the same Union ID in apps developed by the same developer, and has different Union IDs in apps developed by different developers. A developer can use Union ID to link the same user's identities in multiple apps.How to get Union ID - `user_id`: Identifies a user to a tenant. The same user has different User IDs in different tenants. In one single tenant, a user has the same User ID in all apps （including store apps）. User ID is usually used to communicate user data between different apps. How to get User ID  **Default value**: `open_id` **When the value is `user_id`, the following field scopes are required**: `contact:user.employee_id:readonly` | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `attendees` | `calendar.event.attendee[]` | No | Add a new invitee list. **Notice**: - The maximum number of invitees (including meeting rooms) that can be set in a single request is 1,000. - The maximum number of meeting rooms that can be set up in a single request is 100. |
|   `type` | `string` | No | Event invitee type **Example value**: "user" **Optional values are**:  - `user`: User - `chat`: Group - `resource`: Room - `third_party`: Email  |
|   `is_optional` | `boolean` | No | Whether the invitee is optional. **Optional values are**: - true - false **Note**: This field cannot be edited for group invitees. **Example value**: true **Default value**: `false` |
|   `user_id` | `string` | No | User ID. When selecting a user type invitee (the value of type is user), this parameter needs to be passed in. The user ID type passed in needs to be consistent with the value of user_id_type. For information about user IDs, see User-related ID concepts. **Example value**: "ou_xxxxxxxx" |
|   `chat_id` | `string` | No | Group ID. When selecting group type invitees (type value is chat), this parameter needs to be passed in. For details about group ID, please refer to Group ID Description. **Example value**: "oc_xxxxxxxxx" |
|   `room_id` | `string` | No | Meeting room ID. When selecting meeting room type invitees (type value is resource), this parameter needs to be passed in. You can obtain the specified meeting room ID through the following interface: - Query meeting room list - Search meeting room **Example value**: "omm_xxxxxxxx" |
|   `third_party_email` | `string` | No | Email address. When selecting an external mailbox type invitee (type value is third_party), this parameter needs to be passed in. **Example value**: "wangwu@email.com" |
|   `operate_id` | `string` | No | Meeting room contact ID. The user ID type passed in needs to be consistent with the value of user_id_type. For information about user IDs, see User-related ID concepts. **Note**: If the current schedule is created based on the application identity, when adding a meeting room type invitee, you need to specify the contact person of the meeting room through this parameter, and the contact will be displayed in the schedule meeting room information. **Default**: empty **Example value**: "ou_xxxxxxxx" |
|   `resource_customization` | `calendar.attendee.resource_customization[]` | No | Personalized configuration of meeting rooms. - When selecting meeting room type invitees, if the meeting room has a reservation form, you can configure the form information through this parameter. - When the currently added invitee does not involve the personalized configuration of the meeting room, there is no need to set this parameter. |
|     `index_key` | `string` | Yes | Unique ID of each configuration. **Example value**: "16281481596100" |
|     `input_content` | `string` | No | This parameter is required when the type field is left empty. **Example value**: "xxx" |
|     `options` | `customization.option[]` | No | Options of each configuration. **Example value**: 无 |
|       `option_key` | `string` | No | Unique ID of each option. **Example value**: "16281481596185" |
|       `others_content` | `string` | No | This parameter is required when the type field is set to other options. **Example value**: "xxx" |
|   `approval_reason` | `string` | No | The reason for requesting reservation approval for the meeting room. Parameter configuration description: - This field takes effect only when booking an approved meeting room using user identity (user_access_token). - For the scenario of applying for reservation and approval of a meeting room, if this value is not passed, the reservation will directly fail. - If you use the application identity (tenant_access_token) to reserve an approval meeting room, it will directly fail. **Default**: empty **Example value**: "Approval reason" **Data validation rules**: - Maximum length: `200` characters |
| `need_notification` | `boolean` | No | Whether to send bot notifications to invitees. **Optional values are**: - true (default): send - false: do not send **Example value**: false |
| `instance_start_time_admin` | `string` | No | The instance to be modified when accessing as an administrator. **Note**: - This parameter is only used to modify a event instance in a repeating event. This field does not need to be filled in for non-repeating events. - You can call the Get repeating event instance interface to obtain the event_id of a event instance in the repeating event. The value of this parameter is the timestamp suffix of event_id. For example, the queried event instance ID is `2cf525f0-1e67-4b04-ad4d-30b7f003903c_1713168000`, then the current `instance_start_time_admin` value is `1713168000`. **Default**: empty **Example value**: "1647320400" |
| `is_enable_admin` | `boolean` | No | Whether to enable the meeting room administrator status (you need to set a member as the meeting room administrator in the management background first). **Optional values are**: - true: enable - false (default): disable **Note**: After it is turned on, this request will only process meeting room data, and the operations of other invitees will not take effect. **Example value**: false |
| `add_operator_to_attendee` | `boolean` | No | Whether to add the meeting room contact (operate_id) to the schedule invitees. **Optional values are**: - true (default): enabled - false: disable **Example value**: false | ### Request body example

{
    "attendees": [
        {
            "type": "user",
            "is_optional": true,
            "user_id": "ou_xxxxxxxx",
            "chat_id": "oc_xxxxxxxxx",
            "room_id": "omm_xxxxxxxx",
            "third_party_email": "wangwu@email.com",
            "operate_id": "ou_xxxxxxxx",
            "resource_customization": [
                {
                    "index_key": "16281481596100",
                    "input_content": "xxx",
                    "options": [
                        {
                            "option_key": "16281481596185",
                            "others_content": "xxx"
                        }
                    ]
                }
            ]
        }
    ],
    "need_notification": false,
    "is_enable_admin": false,
    "instance_start_time_admin": "1647320400"
}

## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `attendees` | `calendar.event.attendee[]` | List of all event invitees of the event after new event invitees are added. |
|     `type` | `string` | Event invitee type. **Optional values are**:  - `user`: User - `chat`: Group - `resource`: Room - `third_party`: Email  |
|     `attendee_id` | `string` | Invitee ID. The unique ID of the event invitee in the current event. This ID can be used to delete the event invitee later, or to query the group membership information of group type invitees. |
|     `rsvp_status` | `string` | RSVP status of the event invitee. **Optional values are**:  - `needs_action`: The event invitee hasn't replied or a room is being booked. - `accept`: The event invitee accepted or a room was successfully booked. - `tentative`: The event invitee replied pending. - `decline`: The event invitee declined or a room failed to be booked. - `removed`: The event invitee or room has been removed from the event.  |
|     `is_optional` | `boolean` | Whether the event invitee is optional. This field for group event invitees cannot be edited. |
|     `is_organizer` | `boolean` | Whether the event invitee is an event organizer. |
|     `is_external` | `boolean` | Whether the event invitee is external. External invitees have no edit scope. |
|     `display_name` | `string` | Event invitee name. |
|     `chat_members` | `attendee_chat_member[]` | Group members. This field is valid when the type field is Chat. Group members have no edit scope. |
|       `rsvp_status` | `string` | RSVP status of the event invitee. **Optional values are**:  - `needs_action`: The event invitee hasn't replied or a room is being booked. - `accept`: The event invitee accepted or a room was successfully booked. - `tentative`: The event invitee replied pending. - `decline`: The event invitee declined or a room failed to be booked. - `removed`: The event invitee or room has been removed from the event.  |
|       `is_optional` | `boolean` | Whether the event invitee is optional. |
|       `display_name` | `string` | Event invitee name. |
|       `is_organizer` | `boolean` | Whether the event invitee is an event organizer. |
|       `is_external` | `boolean` | Whether the event invitee is external. |
|     `user_id` | `string` | Event invitee's user ID, which depends on the returned user_id_type parameter. When is_external is true, this field returns the open_id or union_id. For more information, see User-related IDs. |
|     `chat_id` | `string` | The group chat_id of a chat-type event invitee. For details, see Group ID description. |
|     `room_id` | `string` | The room_id of a resource-type event invitee. |
|     `third_party_email` | `string` | Email of a third_party type event invitee. |
|     `operate_id` | `string` | If the event was created using App Identity, specify the room contact ID when adding the room. The ID type is consistent with the value of user_id_type. |
|     `resource_customization` | `calendar.attendee.resource_customization[]` | Personalized configurations of the room. |
|       `index_key` | `string` | Unique ID of each configuration. |
|       `input_content` | `string` | This parameter is required when the type field is left empty. |
|       `options` | `customization.option[]` | Options of each configuration. |
|         `option_key` | `string` | Unique ID of each option. |
|         `others_content` | `string` | This parameter is required when the type field is set to other options. |
|     `approval_reason` | `string` | The reason for requesting reservation approval for the meeting room. Parameter configuration description: - This field takes effect only when booking an approved meeting room using user identity (user_access_token). - For the scenario of applying for reservation and approval of a meeting room, if this value is not passed, the reservation will directly fail. - If you use the application identity (tenant_access_token) to reserve an approval meeting room, it will directly fail. | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
        "attendees": [
            {
                "type": "user",
                "attendee_id": "user_xxxxxx",
                "rsvp_status": "needs_action",
                "is_optional": true,
                "is_organizer": true,
                "is_external": false,
                "display_name": "Zhang San",
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
                        "input_content": "xxx",
                        "options": [
                            {
                                "option_key": "16281481596185",
                                "others_content": "xxx"
                            }
                        ]
                    }
                ]
            }
        ]
    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 190002 | invalid parameters in request | Invalid request parameters. Troubleshooting suggestions are as follows: - Confirm that the field name and parameter type of the request parameter are correct. - Confirm that the permissions for the corresponding resource have been applied for. - Confirm that the corresponding resource has not been deleted. |
| 500 | 190003 | internal service error | Internal service error, please contact [Technical Support](https://applink.larksuite.com/TLJpeNdW). |
| 429 | 190004 | method rate limited | Method frequency limit. It is recommended to try again later and reduce the request QPS appropriately. |
| 429 | 190005 | app rate limited | Frequency limiting is applied. We recommend that you try again later and reduce the request QPS appropriately. |
| 403 | 190006 | wrong unit for app tenant | Request error, check whether the App ID and App Secret are correct. If the problem still cannot be solved, please consult [Technical Support](https://applink.larksuite.com/TLJpeNdW). |
| 404 | 190007 | app bot_id not found | The bot_id of the application is not found. You need to make sure that the application has enabled bot capability. If the problem is still not solved, please contact [technical support](https://applink.larksuite.com/TLJpeNdW). |
| 429 | 190010 | current operation rate limited | The current operation is limited, usually because concurrent preemption of public resources fails. You can appropriately reduce the frequency of the current operation and try again. |
| 403 | 190011 | tenant encrypt key has been deleted | The autonomous key in the encryption and decryption state is deleted, and the data encrypted by the key is unavailable. |
| 403 | 190012 | tenant decrypt key has been deleted | Only the autonomous key in the decryption state is deleted, and the data encrypted by the key is unavailable. |
| 400 | 190013 | content being saved is at risk | The incoming content is risk controlled, please check whether the content is legal. |
| 404 | 191000 | calendar not found | Calendar not found. You need to check and change to the correct calendar ID. |
| 400 | 191001 | invalid calendar_id | The calendar_id is invalid. You need to check and change it to the correct calendar ID. |
| 403 | 191002 | no calendar access_role | The current identity does not have permission to access the calendar. If you want to query a calendar, you need to ensure that the current identity has permission to access the calendar. |
| 403 | 191003 | calendar is deleted | The calendar has been deleted. You need to check and change to the correct calendar ID. |
| 403 | 191004 | invalid calendar type | The calendar type is incorrect. You can call the query calendar information interface to obtain the calendar type information, and then ensure that the calendar type is applicable to the current interface. |
| 400 | 193000 | invalid event_id | The event_id is invalid. You need to check and change it to the correct event ID. |
| 404 | 193001 | event not found | Event not found. You need to make sure you passed in the correct event ID. |
| 403 | 193002 | no permission to operate event | No permission to operate. You need to make sure you have permission to edit the calendar and schedule. |
| 403 | 193003 | event is deleted | The event has been deleted. You need to check and change to the correct event ID. |
| 404 | 194000 | attendee not found | No participant was found. You need to make sure the participant parameters are filled in correctly. |
| 403 | 194001 | no permission to list event attendees | No operation permission. Check whether calendar_id is the organizer's calendar of the current schedule, or whether the organizer has permission to view participants. |
| 403 | 194002 | no permission to create event attendees | No operation permission. Check whether calendar_id is the organizer's calendar of the current schedule, or whether the organizer has permission to invite participants. |
| 403 | 194003 | no permission to delete event attendees | No operation permission. Check whether calendar_id is the organizer calendar of the current schedule. |
| 400 | 194004 | invalid attendee type | The participant type is invalid. Please check whether the participant type is entered correctly. |
| 404 | 195100 | user is dismiss or not exist in the tenant | The current identity or specified user has resigned or is no longer in the tenant. Please check and change to the correct identity to call the interface. | For more error code information, see General error codes.
