---
title: "Events"
url: "https://open.larksuite.com/document/ukTMukTMukTM/ukjNxYjL5YTM24SO2EjN"
section: "Deprecated Version (Not Recommended)"
updated: "1622191954000"
---

# Group Chat Events

To subscribe to bot related events (when a bot joins or leaves a group or when a person in chat mentions a bot), see Bot and Chat Message Events.

## User Joins or Leave Group
This event is triggered when a user joins or leaves a group.

* Note: This event is only triggered when the bot feature is enabled and when the bot is in the group in which the changes described above occur.

There are three event types:
1. User joins group - add_user_to_chat
2. User leaves group - remove_user_from_chat
3. An add user to group operation is canceled - revoke_add_user_from_chat

**Callback example:**
```json
{
  "ts": "1502199207.7171419", // the time the event is sent; this is usually close to the time the event is generated 
  "uuid": "bc447199585340d1f3728d26b1c0297a", // unique event ID
  "token": "41a9425ea7df4536a7623e38fa321bae", // authentication token 
  "type": "event_callback", // event callbacks always use event_callback
  "event": {
    "app_id": "cli_9c8609450f78d102",
    "chat_id": "oc_9e9619b938c9571c1c3165681cdaead5", // the group chat ID
    "operator": {
      // the person who added or removed the user from the chat. If the user left the chat of their own volition, the operator is the user
      "open_id":"ou_18eac85d35a26f989317ad4f02e8bbbb" // the user’s unique ID for this app; each user has a separate open_id for each app
      "user_id": "ca51d83b"  // user ID; only returned for custom apps 
    },
    "tenant_key": "736588c9260f175d",
    "type": "add_user_to_chat", // the event type, i.e. add_user_to_chat, remove_user_from_chat or revoke_add_user_from_chat
    "users": [{
        "name": "James",
        "open_id": "ou_706adeb944ab1473b9fb3e7da2a40b68",
        "user_id": "51g97a4g"
      },
      {
        "name": "Jason",
        "open_id": "ou_7885357f9922aaa34001b190109e0b48",
        "user_id": "6e125386"
      }
    ]
  }
}
```

##  Group Closed
This event is triggered when the group is closed.

* Note: This event is only triggered when the bot feature is enabled and when the bot is in the group that is closed.

**Callback example:**
```json
{
	"ts": "1502199207.7171419", // the time the event is sent; this is usually close to the time the event is generated 
	"uuid": "bc447199585340d1f3728d26b1c0297a", // unique event ID
	"token": "41a9425ea7df4536a7623e38fa321bae", // authentication token 
	"type": "event_callback", // event callbacks always use event_callback
	"event": {
		"app_id": "cli_9c8609450f78d102",
		"chat_id": "oc_9f2df3c095c9395334bb6e70ced0fa83",
		"operator": {
			"open_id": "ou_18eac85d35a26f989317ad4f02e8bbbb",
			"user_id": "ca51d83b"
		},
		"tenant_key": "736588c9260f175d",
		"type": "chat_disband"
	}
}

```

##  Group Configuration Changes
This event will be triggered in the following circumstances:
|User Behavior | after_change/before_change Fields |
|--|--|
|Change of owner|owner_user_id (only for custom apps), owner_open_id|
|Group owner changes the message notification settings|message_notification|
|Group owner changes the “Only the group owner can add group members” setting|add_member_permission| * Note: This event is only triggered when the bot feature is enabled and when the bot is in the group in which the changes described above occur.

**Callback example:**
```json
{
  "ts": "1502199207.7171419", // the time the event is sent; this is usually close to the time the event is generated 
  "uuid": "bc447199585340d1f3728d26b1c0297a", // unique event ID
  "token": "41a9425ea7df4536a7623e38fa321bae", // authentication token 
  "type": "event_callback", // event callbacks always use event_callback
  'event': {
    'tenant_key': '2d520d3b434f175e',
    'type': 'group_setting_update'
    'app_id': 'cli_9e28cb7ba56a100e',
    'chat_id': 'oc_066cad06159f0752fe02c9af8aebfc5a',
    'after_change': {
      // when the group owner changes, the following two fields will be passed. If the group owner hasn’t changed, these fields won’t exist.
      'owner_open_id': 'ou_b5f49f047428e47925599be05e3255f6',
      'owner_user_id': '35g25a59',
      // when the “Only the group owner can add group members” setting is changed, the following field will be passed
      'add_member_permission': 'everyone',
      // when the “message notifications” setting is changed, this field will be passed
	  'message_notification': False
    },
    'before_change': {
      // this will have the same fields as after_change
    },
    'operator': { // the user who changed the setting
      'open_id': 'ou_55a50f4c6eba6adbfc8b94803fe78825',
      'user_id': 'g6964gd3'
    }
  }
}
```
