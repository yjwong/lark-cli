---
title: "Contacts Events"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uITNxYjLyUTM24iM1EjN"
section: "Deprecated Version (Not Recommended)"
updated: "1621758776000"
---

# Contacts Events
> To understand the usage scenarios and configuration process of event subscription, please refer to overview of event subscription.

## Contacts Updates
### Employee Updates
This event is sent when a user joins an organization (user_add), leaves an organization (user_leave) or their personal information is changed (user_update). 
- Permissions required: Access Contacts as an app
- Other requirements: This event is only sent for employee changes within the scope of the organization’s Contacts permission.
- Works with : Obtain employee information

**Callback example:**
```json
{ 
    "ts": "1502199207.7171419", // the time the event is sent; this is usually close to the time the event is generated 
    "uuid": "bc447199585340d1f3728d26b1c0297a",  // unique event ID
    "token": "41a9425ea7df4536a7623e38fa321bae", // authentication token 
    "type": "event_callback", // event callbacks always use event_callback
    "event": { 
         "type": "user_add",    // the event type, i.e. user_add, user_update, or user_leave 
         "app_id": "cli_xxx",   // app ID
         "tenant_key": "xxx",   // organization ID 
         "open_id":"xxx" ,  // the user’s unique ID for this app; each user has a separate open_id for each app
         "employee_id":"xxx",    // i.e. the user’s ID within the organization; only returned for custom apps
         "union_id": "xxx" // the user’s unique ID for this ISV;
    }  
}
```

### Department Changes
This event is sent when a department is created (dept_add), deleted (dept_delete), or edited (dept_update).
- Permissions required: Access Contacts as an app
- Other requirements: This event is only sent for department changes within the scope of the organization’s Contacts permission.
- Works with: Obtain department details

**Callback example:**
```json
{ 
    "ts": "1502199207.7171419", // the time the event is sent; this is usually close to the time the event is generated
    "uuid": "bc447199585340d1f3728d26b1c0297a",  // unique event ID
    "token": "41a9425ea7df4536a7623e38fa321bae", // authentication token 
    "type": "event_callback", // event callbacks always use event_callback
     "event": { 
         "type": "dept_add",  // the event type, i.e. dept_add, dept_update, or dept_delete 
         "app_id": "cli_xxx",  // user ID 
         "tenant_key": "xxx",           // organization ID 
         "open_department_id":"aaa",  // department ID, alreadyabandoned
         "department": {
         	"open_id": "od-xxx",
            "custom_id": "aaa"
         }
     } 
}
```

## User Status Changes
This event is triggered when a user activates, suspends or restores their account or leaves the organization. This event does not require any permissions.

- Note: If you have also subscribed to the [User event](#User Event) that event will also be triggered alongside the user_update event. Make sure the business logic is not duplicated.

**Callback example:**
```json
{
	"ts": "1502199207.7171419", // the time the event is sent; this is usually close to the time the event is generated 
	"uuid": "bc447199585340d1f3728d26b1c0297a",  // unique event ID
	"token": "41a9425ea7df4536a7623e38fa321bae", // authentication token 
	"type": "event_callback", // event callbacks always use event_callback
	"event": {
		"type": "user_status_change",    // event type
		"app_id": "cli_xxx",   // app ID
		"tenant_key": "xxx",   // organization ID 
		"open_id":"xxx" ,  // the user’s unique ID for this app; each user has a separate open_id for each app
		"employee_id":"xxx"    // i.e. the user’s ID within the organization; only returned for custom apps
		"before_status": { // status before the change
			"is_active": false,        // if the account was active
			"is_frozen": false,       // if the account was suspended
			"is_resigned": false    // if the employee has left
		},
		"change_time": "2020-02-21 16:28:48", // the time of the status change
		"current_status": { // the status after the change
			"is_active": true,
			"is_frozen": false,
			"is_resigned": false
		}
	}
}
```

Additional information: The following operations result in the following status changes:
| Operation         | is_active           | is_frozen         | is_resigned |
| --------- | --------------- | -------   | ----------- |
|Account is activated when the user joins the organization| false -> true | - | - |
|Account is suspended| - | false -> true | - |
|Account is restored| - | true -> false | - |
|User has left| - | - | false -> true | ## Permission Scope Change
When the app has the permission to Access Contacts as an app, the administrator can set the scope of the Contacts permission.

![图片名称](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/ark/8206b84f0f264fdeb0e92a4f1bb0ad92.png)
This event is triggered when the permission scope is changed.
- Permissions required: Access Contacts as an app
- Note: You do not need to subscribe to this event. The app is automatically subscribed to it when Access Contacts as an app is enabled.
- Works with: Obtain the Contacts permission scope

**Callback example:**
```json
{ 
    "ts": "1502199207.7171419", // the time the event is sent; this is usually close to the time the event is generated 
    "uuid": "bc447199585340d1f3728d26b1c0297a",  // unique event ID
    "token": "41a9425ea7df4536a7623e38fa321bae", // authentication token 
    "type": "event_callback", // event callbacks always use event_callback
     "event": { 
         "type": "contact_scope_change", // event type 
         "app_id": "cli_xxx",   // app ID
         "tenant_key": "xxx",  // organization ID 
     } 
}
```
